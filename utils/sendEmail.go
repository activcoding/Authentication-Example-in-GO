package utils

import (
	"fmt"
	"github.com/mailjet/mailjet-apiv3-go/v4"
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

func SendMailWithMailJet(recipientName string, recipientEmail string, subject string, message string) {
	apiKey := GetVariable("MAIL_JET_API_KEY")
	apiSecret := GetVariable("MAIL_JET_SECRET_KEY")

	mailClient := mailjet.NewMailjetClient(apiKey, apiSecret)
	messageInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: "tomludwig@duck.com",
				Name:  "Tom Ludwig",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: recipientEmail,
					Name:  recipientName,
				},
			},
			Subject:  subject,
			TextPart: "This is a test email",
			HTMLPart: message,
		},
	}

	messages := mailjet.MessagesV31{Info: messageInfo}
	res, err := mailClient.SendMailV31(&messages)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
