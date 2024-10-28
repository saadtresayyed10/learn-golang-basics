package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	bookings := []string{}

	for remainingTickets > 0 && len(bookings) <= 50 {
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

		isValidName, isValidEmail, isValidTickets := validateConditions(firstName, lastName, email, userTickets, remainingTickets)

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Println("Whole slice: ", bookings)
			fmt.Println("1st element in slice: ", bookings[0])
			fmt.Println("Size of the slice: ", len(bookings))

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

func greetUsers(conferenceName string, conferenceTickets uint, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets from which %v are remaining\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets to attend\n")
}

func validateConditions(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTickets

}
