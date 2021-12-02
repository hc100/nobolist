package sendmail

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// Convert UTF-8 to ISO2022JP
func toISO2022JP(str string) ([]byte, error) {
	reader := strings.NewReader(str)
	transformer := japanese.ISO2022JP.NewEncoder()

	return ioutil.ReadAll(transform.NewReader(reader, transformer))
}

// 76バイト毎にCRLFを挿入する
func add76crlf(msg string) string {
	var buffer bytes.Buffer
	for k, c := range strings.Split(msg, "") {
		buffer.WriteString(c)
		if k%76 == 75 {
			buffer.WriteString("\r\n")
		}
	}
	return buffer.String()
}

// UTF8文字列を指定文字数で分割
func utf8Split(utf8string string, length int) []string {
	resultString := []string{}
	var buffer bytes.Buffer
	for k, c := range strings.Split(utf8string, "") {
		buffer.WriteString(c)
		if k%length == length-1 {
			resultString = append(resultString, buffer.String())
			buffer.Reset()
		}
	}
	if buffer.Len() > 0 {
		resultString = append(resultString, buffer.String())
	}
	return resultString
}

// サブジェクトをMIMEエンコードする
func encodeSubject(subject string) string {
	var buffer bytes.Buffer
	for _, line := range utf8Split(subject, 13) {
		buffer.WriteString(" =?iso-2022-jp?B?")
		iso2022jp, err := toISO2022JP(line)
		if err == nil {
			buffer.WriteString(base64.StdEncoding.EncodeToString(iso2022jp))
		}

		buffer.WriteString("?=\r\n")
	}
	return buffer.String()
}
