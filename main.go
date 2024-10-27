package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets from which %v are remaining\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets to attend")
}
