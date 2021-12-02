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

// UserRegistrationComplete user registration
func UserRegistrationComplete(to string) error {
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

	msg, err := getUserRegistrationCompleteMessage(to, from)
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

func getUserRegistrationCompleteMessage(to string, from string) (string, error) {
	bu := os.Getenv("BASE_URL")
	if bu == "" {
		return "", fmt.Errorf("env BASE_URL is not set")
	}

	data := map[string]string{
		"To":      to,
		"From":    from,
		"Url":     bu,
		"Subject": encodeSubject("【ﾉﾎﾞﾘｽﾄ】会員登録完了"),
	}

	t, err := template.New("user_registration_complete.tmpl").ParseFiles("template/mail/user_registration_complete.tmpl")
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)

	if err = t.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
