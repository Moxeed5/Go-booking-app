package main

import (
	"fmt"
	"strings"
)


func main(){
	conferenceName := "Go Conference" //cannot explicitly declare type when using := and can only be used with a func. 
	const conferenceTickets int = 50
	var totalTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}
	
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to the %v booking application\n", conferenceName)
	fmt.Printf("Total tickets available: %v & Remaining tickets available: %v \n",totalTickets,remainingTickets)
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

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName + " " + lastName)

	
	// fmt.Printf("The whole array: %v \n", bookings)
	// fmt.Printf("Index 0: %v \n", bookings[0])
	// fmt.Printf("The array type is: %T \n", bookings)
	// fmt.Printf("The length of the bookings array is %v \n", len(bookings))
	
	fmt.Printf("User %v has booked %v tickets\n", bookings[0], userTickets)
	fmt.Printf("Thanks for your order %v, you will receive %v tickets in your email: %v\n", bookings[0], userTickets, email)

	firstNames := []string{}

	//the below loop is a foreach. Foreach index within bookings slice (declared above & not an array bc we don't specify length).
	//booking is the temp var within the foreach. Similar to var item in List in c#
	//range is used to return element within the index and store it in booking var
	//fields is a method within strings package. Fields separates elements within a single element on the space. 
	//****update**** replaced index with underscore which tells go we are not using the variable. 
	//if we need to use the index of the slice, use index. If you just need the element, use an underscore in place of index. 
	//I can think of scenerios where every xth element needs to be manipulated. In that case you'd want to use index to track position.
	for _, booking := range bookings{ 

		var names = strings.Fields(booking) 
		firstNames = append(firstNames, names[0])
		fmt.Printf("%T", names)
	} 
	fmt.Printf("First names of all current bookings %v\n", firstNames)
	
	}
	
	
}

