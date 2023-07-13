package email

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
	"peony/utils"
)

var (
	Dialer *gomail.Dialer
)

// InitEmail 初始化邮件服务
func InitEmail() {
	dialer := gomail.NewDialer(
		utils.V.GetString("email.host"),
		utils.V.GetInt("email.port"),
		utils.V.GetString("email.username"),
		utils.V.GetString("email.password"),
	)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	Dialer = dialer
}

func SendEmail(toEmail string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", utils.V.GetString("email.username"))
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "邮箱验证")
	m.SetBody("text/html", body)

	err := Dialer.DialAndSend(m)
	return err
}
