package main

import (
	"fmt"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	fmt.Printf("Conference type %T, Conference Tickets type %T\n", conferenceName, conferenceTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets from which %v are remaining\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets to attend\n")

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
			fmt.Printf("Only %v are remaining, can't book %v tickets\n", remainingTickets, userTickets)
		}

	}

}
