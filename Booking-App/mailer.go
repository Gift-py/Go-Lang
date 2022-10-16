package main

import (
	"bytes"
	"html/template"
	"log"
	"time"

	"github.com/go-mail/mail/v2"
)

type MailerConfig struct {
	Timeout      time.Duration
	Host         string
	Port         int
	Username     string
	Password     string
	Sender       string
	TemplatePath string
}

type Mailer struct {
	dialer *mail.Dialer
	config MailerConfig
	sender string
}

func New(config MailerConfig) Mailer {
	dialer := mail.NewDialer(config.Host, config.Port, config.Username, config.Password)
	dialer.Timeout = config.Timeout

	return Mailer{
		dialer: dialer,
		sender: config.Sender,
		config: config,
	}
}

func (m Mailer) Send(to, templateFile string, data interface{}) error {
	t := template.New("go_conference.html")
	var err error

	t, err = t.ParseFiles(templateFile)
	if err != nil {
		log.Println(err)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		log.Println(err)
	}

	result := tpl.String()

	msg := mail.NewMessage()
	msg.SetHeader("From", "giftabumere247@gmail.com")
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", "Go-Conference")
	msg.SetBody("text/html", result)

	return m.dialer.DialAndSend(msg)
}
