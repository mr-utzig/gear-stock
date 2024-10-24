package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

type response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse(status bool, message string, data interface{}) *response {
	return &response{Status: status, Message: message, Data: data}
}

func SendMail(to string, subject string, body string) error {
	auth := smtp.PlainAuth("", os.Getenv("SMTP_USERNAME"), os.Getenv("SMTP_PASSWORD"), os.Getenv("SMTP_HOST"))
	from := "gear.stock@lutzig.com"
	msg := []byte(fmt.Sprintf("From: %v\r\nTo: %v\r\nSubject: %v\r\n\r\n%v\r\n", from, to, subject, body))

	return smtp.SendMail(os.Getenv("SMTP_ADDR"), auth, from, []string{to}, msg)
}
