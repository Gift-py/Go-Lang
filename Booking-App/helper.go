package main

import (
	"log"
	"regexp"
	"time"
)

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	isValidEmail := re.MatchString(email)
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber

}

func SendTicket(firstName string, lastName string, email string, userTickets uint) {
	//to generate ticket IDs
	var ticketId = make([]int, 0)
	for idx := 0; idx <= int(userTickets); idx++ {
		ticketId = append(ticketId, idx+6661)
	}

	from := Api_username
	password := Api_password

	host := Api_host
	port := Api_port

	//subject := fmt.Sprintf("Subject: %v tickets for %v %v\n", userTickets, firstName, lastName)
	//body := fmt.Sprintf("Thank you %v for your purchase. This is a list of your ticket ID(s): \n %v", firstName, ticketId)

	config := MailerConfig{
		Host:     host,
		Port:     port,
		Username: from,
		Password: password,
		Timeout:  5 * time.Second,
		Sender:   "Go Conference",
	}

	sender := New(config)

	err := sender.Send(email, "go_conference.html", nil)
	if err != nil {
		log.Fatal(err)
	}

}
