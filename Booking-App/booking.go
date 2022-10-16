package main

import (
	"fmt"
)

const conferenceTickets = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	greetUser()

	if remainingTickets == 0 {
		fmt.Println("Sorry, all tickets are booked. Come back next year")
	}

	fmt.Println("\nUSER DATA COLLECTION")

	firstName, lastName, email, userTickets := getUserData()
	//to validate user input
	isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		var userData = bookTicket(firstName, lastName, email, userTickets)
		SendTicket(firstName, lastName, email, userTickets, userData)

		firstNames := getFirstName()
		fmt.Printf("The names of the bookings are %v\n", firstNames)

	} else {
		if !isValidName {
			fmt.Println("first name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("email address you entered doesn't contain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("number of tickets you entered is invalid")
		}
	}
}

func greetUser() {
	fmt.Println("Hello Friend! Welcome to Go Conference Ticket Booking App")
	fmt.Printf("We have a  total of %v tickets, and %v tickets are available. \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets now!!")
}

func getUserData() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("Enter your first name -> ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name -> ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email -> ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets to purchase -> ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets

}

func bookTicket(firstName string, lastName string, email string, userTickets uint) UserData {
	remainingTickets = remainingTickets - userTickets
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("List of booking %v:\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for the Go Conference \n", remainingTickets)

	return userData
}

func getFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}
