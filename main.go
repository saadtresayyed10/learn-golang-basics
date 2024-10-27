package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	fmt.Printf("Conference type %T, Conference Tickets type %T\n", conferenceName, conferenceTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets from which %v are remaining\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets to attend\n")

	var username string
	var userTickets int

	username = "Saad"
	userTickets = 2

	fmt.Printf("User %v has booked %v tickets\n", username, userTickets)

}
