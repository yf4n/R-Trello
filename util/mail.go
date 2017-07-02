package util

import (
	"log"
	"net/mail"
	"net/smtp"

	"github.com/scorredoira/email"
)

func SendMail(attachPath string) {
	subject := GetIniConfig("mail", "subject")
	from := GetIniConfig("mail", "from")
	to := GetIniConfig("mail", "to")
	pwd := GetIniConfig("mail", "pwd")
	host := GetIniConfig("mail", "host")
	port := GetIniConfig("mail", "port")
	alias := GetIniConfig("mail", "alias")
	body := GetIniConfig("mail", "body")
	servername := host + ":" + port

	log.Println(servername, subject, from, to, pwd, host, port, alias, body, attachPath)
	m := email.NewMessage(subject, body)
	m.From = mail.Address{Name: alias, Address: from}
	m.To = []string{to}
	log.Println("Add Attach")

	if err := m.Attach(attachPath); err != nil {
		log.Fatal(err)
	}
	log.Println("add ok, ready to send")

	// send it
	auth := smtp.PlainAuth("", from, pwd, host)
	log.Println(" au ok")
	if err := email.Send(servername, auth, m); err != nil {
		log.Fatal(err)
	}
}
