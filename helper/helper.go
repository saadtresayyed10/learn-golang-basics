package helper

import (
	"fmt"
	"strconv"
	"strings"
)

func ValidateConditions(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTickets

}

func BookTickets(remainingTickets uint, userTickets uint, bookings []map[string]string, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)

	fmt.Println("Whole slice: ", bookings)
	fmt.Println("1st element in slice: ", bookings[0])
	fmt.Println("Size of the slice: ", len(bookings))
}
