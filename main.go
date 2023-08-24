package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var totalTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}

	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to the %v booking application\n", conferenceName)
	fmt.Printf("Total tickets available: %v & Remaining tickets available: %v \n", totalTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		// ask user for name & tickets
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter the number of tickets you would like to purchase: ")
		fmt.Scan(&userTickets)

		if userTickets < remainingTickets {
			remainingTickets -= userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("User %v has booked %v tickets\n", firstName, userTickets)
			fmt.Printf("Thanks for your order %v, you will receive %v tickets in your email: %v\n", firstName, userTickets, email)

			firstNames := []string{}
			for _, booking := range bookings {
				names := strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("First names of all current bookings %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Conference is full. No tickets available.")
				break
			}
		} else if userTickets == remainingTickets {
			fmt.Println("Oh jeez you've bought all the tickets. You must be rich") //fact check: I am not rich
			fmt.Printf("Thanks for your order %v, you will receive %v tickets in your email: %v\n", firstName, userTickets, email)
			break;
		} else {
			fmt.Printf("We only have %v tickets remaining. Your order of %v is too large\n", remainingTickets, userTickets)
			continue
		}
	}
}
