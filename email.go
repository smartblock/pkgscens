package pkgscens

import (
	"encoding/base64"
	"fmt"
	"net/smtp"
	"strings"
)

//SMTPAuth def
type SMTPAuth struct {
	Identity string
	Username string
	Password string
	Host     string
}

//SendMailInput struct
type SendMailInput struct {
	Addr     string
	SMTPAuth SMTPAuth
	FromName string
	FromMail string
	ToMail   []string
	ToName   []string
	Subject  string
	MsgType  string
	Message  string
}

//SendMail func
func SendMail(sendMailInput SendMailInput) error {
	auth := smtp.PlainAuth(
		"",
		sendMailInput.SMTPAuth.Username,
		sendMailInput.SMTPAuth.Password,
		sendMailInput.SMTPAuth.Host,
	)

	if sendMailInput.FromMail == "" {
		return &PkgError{
			Msg: "empty from mail",
		}
	}

	if len(sendMailInput.ToMail) != len(sendMailInput.ToName) {
		return &PkgError{
			Msg: "email_to_name_must_match_with_to_email",
		}
	}

	if sendMailInput.Subject == "" {
		return &PkgError{
			Msg: "Empty Subject",
		}
	}

	msgType := ""

	if sendMailInput.MsgType != "TEXT" && sendMailInput.MsgType != "HTML" {
		if sendMailInput.MsgType == "" {
			msgType = "TEXT"
		} else {
			return &PkgError{
				Msg: "MsgType Error",
			}
		}
	} else {
		msgType = sendMailInput.MsgType
	}

	var recipientMail []string
	var toHeaderMail []string

	// to email process
	if len(sendMailInput.ToMail) != 0 {
		for index, to := range sendMailInput.ToMail {
			str := sendMailInput.ToName[index] + " <" + to + ">"
			toHeaderMail = append(toHeaderMail, str)
			recipientMail = append(recipientMail, to)
		}
	}

	toHeader := strings.Join(toHeaderMail, ",")

	header := make(map[string]string)
	header["MIME-Version"] = "1.0"
	header["Content-Transfer-Encoding"] = "base64"
	header["From"] = sendMailInput.FromName + " <" + sendMailInput.FromMail + ">"
	header["Subject"] = sendMailInput.Subject

	if msgType == "HTML" {
		header["Content-Type"] = "text/html; charset=\"utf-8\""
	} else {
		header["Content-Type"] = "text/plain; charset=\"utf-8\""
	}

	if toHeader != "" {
		header["To"] = toHeader
	}

	msg := ""
	for k, v := range header {
		msg += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	msg += "\r\n" + base64.StdEncoding.EncodeToString([]byte(sendMailInput.Message))

	err := smtp.SendMail(
		sendMailInput.Addr,     // server:port
		auth,                   // auth
		sendMailInput.FromMail, // from email_address
		recipientMail,          // to []email_address
		[]byte(msg),            // msg content_here
	)

	if err != nil {
		return err
	}

	return nil
}

//PkgError Type
type PkgError struct {
	Msg string
}

func (m *PkgError) Error() string {
	return m.Msg
}
