package handlers

import (
	"github.com/go-mail/mail"
	"github.com/ABHILASHREDDY-123/ECommerce_Website.git/models"
)



func SendEmail(email *models.Email) error {
    m := mail.NewMessage()
	m.SetHeader("From", email.SenderEmail)
	m.SetHeader("To", email.ReceiverEmail)
	m.SetHeader("Subject", email.Subject)
	m.SetBody("text/plain", email.Body)

	d := mail.NewDialer("smtp.gmail.com", 587, email.SenderEmail, email.Password)

	err := d.DialAndSend(m)
	return err
}