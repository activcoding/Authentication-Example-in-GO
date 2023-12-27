package utils

import (
	"fmt"
	"net/smtp"
)

func SendEmail(to string, subject string, body string) {
	from := "veriy.c1f354@gmail.com"
	password := "zywtuz-kewqoF-1qaqwo"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, password, "smtp.gmail.com"), from, []string{to}, []byte(msg))
	if err != nil {
		panic(err)
		return
	}

	fmt.Println("Email sent")
}
