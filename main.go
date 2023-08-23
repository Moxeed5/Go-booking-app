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

	var userName string
	var userTickets int
	// ask user for name & tickets

	userName = "Tom"
	fmt.Printf("%v has booked %v tickets\n", userName, userTickets)

	
}

