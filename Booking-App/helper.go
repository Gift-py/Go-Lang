package main

import (
	"fmt"
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

	from := "AKIA47F3KO7RYQVZBU2H"
	password := "BOHwds3TD4bh90PX/6PCMRi68qiR3odxyFtKZsNe8xle"

	toEmailAddress := email
	to := []string{toEmailAddress}

	host := "email-smtp.us-east-1.amazonaws.com"
	port := "587"
	address := host + ":" + port

	subject := fmt.Sprintf("Subject: %v tickets for %v %v\n", userTickets, firstName, lastName)
	body := fmt.Sprintf("Thank you %v for your purchase. This is a list of your ticket ID(s): \n %v", firstName, ticketId)

	message := []byte(subject + body)

	config := mail.MailerConfig{
		Host: host,
		Port: port,
		Username: from,
		Password: password,
		Timeout: 5 * time.Second,
		Sender: "Go Conference"
	}


}
