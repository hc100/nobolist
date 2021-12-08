package sendmail

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"text/template"
)

// ResetPassword user registration
func ResetPassword(to string) error {
	sc := os.Getenv("SENDMAIL_COMMAND")
	if sc == "" {
		return fmt.Errorf("env SENDMAIL_COMMAND is not set")
	}

	from := os.Getenv("SYSTEM_MAIL_FROM")
	if from == "" {
		return fmt.Errorf("env SYSTEM_MAIL_FROM is not set")
	}

	sendmail := exec.Command(sc, "-f", from, to)
	stdin, err := sendmail.StdinPipe()
	if err != nil {
		return err
	}

	stdout, err := sendmail.StdoutPipe()
	if err != nil {
		return err
	}

	msg, err := getResetPasswordMessage(to, from)
	if err != nil {
		return err
	}

	sendmail.Start()
	stdin.Write([]byte(msg))
	stdin.Close()
	ioutil.ReadAll(stdout)
	sendmail.Wait()

	return nil
}

func getResetPasswordMessage(to string, from string) (string, error) {
	bu := os.Getenv("BASE_URL")
	if bu == "" {
		return "", fmt.Errorf("env BASE_URL is not set")
	}

	data := map[string]string{
		"To":      to,
		"From":    from,
		"Url":     bu,
		"Subject": encodeSubject("【ﾉﾎﾞﾘｽﾄ】パスワードの再設定完了"),
	}

	t, err := template.New("reset_password.tmpl").ParseFiles("template/mail/reset_password.tmpl")
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)

	if err = t.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
