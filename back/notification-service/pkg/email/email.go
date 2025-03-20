package email

import (
	"fmt"
	"log"
	"net/smtp"
	"notification/pkg/configs"
)

func SendEmail(email string, verificationCode string) error {
	config := configs.GetConfig()
	if config == nil {
		log.Fatal("send email config not loaded")
	}

	auth := smtp.PlainAuth("", config.EmailUsername, config.EmailPassword, config.EmailHost)

	msg := "From: " + config.EmailFrom + "\n" +
		"To: " + email + "\n" +
		"Subject: Verification Code\n\n" +
		"Your verification code is: " + verificationCode

	addr := fmt.Sprintf("%s:%s", config.EmailHost, config.EmailPort)

	err := smtp.SendMail(addr, auth, config.EmailFrom, []string{email}, []byte(msg))
	if err != nil {
		log.Printf("Error sending email: %v", err)
		return err
	}

	log.Printf("Email sent to %s with verification code %s", email, verificationCode)

	return nil
}
