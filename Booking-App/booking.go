package main

import (
	"fmt"
)

/*
	To create a ticketing app
	- welcome screen
	- show amount of available ticket
	- collect userInput
	- save bookings (name, email, amount-of-tiickets)
	- send ticket id (multithreading)

*/
const conferenceTickets = 50

var remainingTickets uint = 50     //uint datatype does not accept negetive values
var bookings = make([]UserData, 0) //a list of the bookings
//the "make" keyword creates a list whose data type is the "UserData" type.
//The 0 is the present length of the list... It is dynamic

//this is a custom data structure (almost like a dictionary)
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
		bookTicket(firstName, lastName, email, userTickets)
		sendTicket(firstName, lastName, email, userTickets)

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

	fmt.Println("Enter your first name -> ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name -> ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email -> ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets to purchase -> ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets

}

func bookTicket(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("List of booking %v\n:", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for the Go Confrernce \n", remainingTickets)
}

func sendTicket(firstName string, lastName string, email string, userTickets uint) {

}

func getFirstName() []string {
	firstNames := []string{} //a list of strings
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}
