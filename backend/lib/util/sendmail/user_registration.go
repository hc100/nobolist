package sendmail

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"text/template"
)

// EXAMPLE: echo "Subject: TestnHello" | sendmail -f you@domain.com you@domain.com
// Useful Links: https://gobyexample.com/spawning-processes

// UserRegistration user registration
func UserRegistration(to string, key string) error {
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

	msg, err := getUserRegistrationMessage(to, from, key)
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

func getUserRegistrationMessage(to string, from string, key string) (string, error) {
	bu := os.Getenv("BASE_URL")
	if bu == "" {
		return "", fmt.Errorf("env BASE_URL is not set")
	}

	data := map[string]string{
		"To":      to,
		"From":    from,
		"Key":     key,
		"Url":     bu,
		"Subject": encodeSubject("【ﾉﾎﾞﾘｽﾄ】会員仮登録の確認"),
	}

	t, err := template.New("user_registration.tmpl").ParseFiles("template/mail/user_registration.tmpl")
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)

	if err = t.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
