package main

import (
	"booking-app/helper"
	"fmt"
)

const conferenceTickets uint = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50

// type UserData struct {
// 	firstName   string
// 	lastName    string
// 	email       string
// 	userTickets uint
// }

func main() {

	greetUsers()

	var bookings = make([]map[string]string, 0)

	for remainingTickets > 0 && len(bookings) <= 50 {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTickets := helper.ValidateConditions(firstName, lastName, email, userTickets, remainingTickets)

		if userTickets <= remainingTickets {
			helper.BookTickets(remainingTickets, userTickets, bookings, firstName, lastName, email)

			// firstNames := []string{}
			// for _, booking := range bookings {
			// 	var names = strings.Fields(booking)
			// 	firstNames = append(firstNames, names[0])
			// }

			if remainingTickets == 0 {
				fmt.Println("We are sold out, come back next year!")
				break
			}

			fmt.Printf("User %v %v has bought %v tickets using %v email\n", firstName, lastName, userTickets, email)
			fmt.Printf("Total remaining tickets %v\n", remainingTickets)
		} else {
			if !isValidName {
				fmt.Println("First or Last name is too short")
			}
			if !isValidEmail {
				fmt.Println("Email doesn't contain @ sign")
			}
			if !isValidTickets {
				fmt.Println("Invalid number of tickets entered")
			}
		}

	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets from which %v are remaining\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets to attend\n")
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter the first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter the last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter the email: ")
	fmt.Scan(&email)

	fmt.Println("Enter the no. of tickets to buy: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}
