package main

import (
	"time"

	"github.com/go-mail/mail/v2"
)

type MailerConfig struct {
	Timeout  time.Duration
	Host     string
	Port     int
	Username string
	Password string
	Sender   string
}

type Mailer struct {
	dailer *mail.Dailer
	config MailerConfig
	sender string
}
