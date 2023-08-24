package main

import "fmt"

func main(){
	conferenceName := "Go Conference" //cannot explicitly declare type when using := and can only be used with a func. 
	const conferenceTickets int = 50
	var totalTickets = 50
	var remainingTickets uint = 50
	
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to the %v booking application\n", conferenceName)
	fmt.Printf("Total tickets available: %v & Remaining tickets available: %v \n",totalTickets,remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var bookings []string
	


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

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName + " " + lastName)

	for i := 0; i < len(bookings); i++ {
	fmt.Printf("The whole array: %v \n", bookings)
	fmt.Printf("Index 0: %v \n", bookings[i])
	fmt.Printf("The array type is: %T \n", bookings)
	fmt.Printf("The length of the bookings array is %v \n", len(bookings))
	
	fmt.Printf("User %v has booked %v tickets\n", bookings[i], userTickets)
	fmt.Printf("Thanks for your order %v, you will receive %v tickets in your email: %v\n", bookings[i], userTickets, email)
	fmt.Printf("Total tickets remaining %v\n", remainingTickets)
	}

	
	
	

	
}

