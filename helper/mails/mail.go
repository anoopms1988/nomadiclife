package mails

import (
	"net/mail"
	"net/smtp"
)

func ValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

type smtpServer struct {
	host string
	port string
}

// Address URI to smtp server
func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

func SendMail(from_address string, to_address string, mail_msg string) error {
	from := from_address
	password := "#Testing123@"
	to := []string{to_address}
	// smtp server configuration.
	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}
	message := []byte(
		"Subject: Nomadic life\r\n\r\n" +
			mail_msg)
	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpServer.host)
	// Sending email.
	err := smtp.SendMail(smtpServer.Address(), auth, from, to, message)
	if err != nil {
		return err
	}
	return nil
}
