package util

import (
	"log"
	"net/mail"
	"net/smtp"

	"github.com/scorredoira/email"
)

func SendMail(attachPath string, dateRange string) {
	subject := GetIniConfig("mail", "subject")
	from := GetIniConfig("mail", "from")
	to := GetIniConfig("mail", "to")
	pwd := GetIniConfig("mail", "pwd")
	host := GetIniConfig("mail", "host")
	port := GetIniConfig("mail", "port")
	alias := GetIniConfig("mail", "alias")
	body := GetIniConfig("mail", "body")
	servername := host + ":" + port

	m := email.NewMessage(subject+dateRange, body)
	m.From = mail.Address{Name: alias, Address: from}
	m.To = []string{to}

	if err := m.Attach(attachPath); err != nil {
		log.Fatal(err)
	}

	auth := smtp.PlainAuth("", from, pwd, host)

	if err := email.Send(servername, auth, m); err != nil {
		log.Fatal(err)
	}
}
