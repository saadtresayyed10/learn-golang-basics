package helper

import (
	"fmt"
	"strings"
)

func ValidateConditions(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTickets

}

func BookTickets(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Println("Whole slice: ", bookings)
	fmt.Println("1st element in slice: ", bookings[0])
	fmt.Println("Size of the slice: ", len(bookings))
}
