package main

import (
	"fmt"
	"sync"
	"time"
)

//sync pacakage is related to concurrensy. Contains Wait method that has add, wait, and done.
//declaring these vars with package scope
//before when we created funcs outside of main, we created new vars as params for the funcs and our func logic reference the vars from the params
//now with global scope we can just use these vars directly instead of passing them into funcs as params. For example, below is the old code for
//...the greet users function and we called it in main like this: greetUsers(conferenceName, totalTickets, remainingTickets)

/*
func greetUsers(conName string, totalTix int, remainingTix uint) {
	fmt.Printf("Welcome to the %v booking application\n", conName)
	fmt.Printf("Total tickets available: %v & Remaining tickets available: %v \n", totalTix, remainingTix)
	fmt.Println("Get your tickets here to attend")
}

*/
var conferenceName string = "Go Conference"
const conferenceTickets int = 50
var totalTickets = 50
var remainingTickets uint = 50
//var bookings = make([]map[string]string, 0) //before we made a slice of type map, we had a slice of type string. We imported package Strings to use Fields method.
//The method signature for Fields is func Fields(s string) []string. It's used for splitting a string into a substring based on white spaces and obviously returns a 
//..string slice with each substring as an individual element. Useful for splitting a string sentence into a string slice of each word from the sentence. 

var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()
	
	// %t rather than %v will print the var type in a fmt.Printf("") statement. It will literally return string, int, etc. 

	
		
		//fmt.scan needs a pointer signified by & infront of the variable that is going to return a value to. 
		//it returns the value directly to the var so you need to have the var declared before you use scan. 
		//don't do var name string = scan(&name) ...... do var name string and then in another line do fmt.scan(&name). It's gonna find name and return the value to it

		//something awesome about Go is that you can have multiple return values. List multiple comma separated values in return line 
		//see below for how to assign those values. Notice that we use := because we're actually delcaring those variables on the left side. 
		//we used those same var names in the func that we declared outside of main. But the contents of that func 
		//...are out of scope within main. 
		//hence why we are declaring and not reassigning the value of an existing var using = .....we are declaring them so we use :=

		firstName, lastName, email, userTickets := getUserInput()

		

		//isValidCity := strings.EqualFold("Singapore", city) || strings.EqualFold(city, "London")

		isValidName, isValidEmail, isValidTickets := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTickets{
			
			wg.Add(1)
			bookTickets(userTickets, firstName, lastName, email)
			go sendTicket(userTickets, firstName, lastName, email) //go keyword enables concurrensy. We have a time delay of 10 seconds within this function. 
			//go keyword enables the rest of the program to continue forward while we wait on the 10 second timer of the function.

			names := getFirstnames()
			fmt.Printf("The first names of all in attendence so far: %v\n", names)

			if remainingTickets == 0 {
				fmt.Println("Conference is full. No tickets available.")
			}

		} else {
			if !isValidName {
				fmt.Printf("%v %v is not valid.\n", firstName, lastName)
			}
			if !isValidEmail {
				fmt.Printf("%v is not a valid email.\n", email)
			}
			if !isValidTickets {
				fmt.Printf("%v is not a valid option for tickets. Ticket selection must be under %v.\n", userTickets, remainingTickets)
			}
		}
	wg.Wait()

	
}

func greetUsers() {
	fmt.Printf("Welcome to the %v booking application\n", conferenceName)
	fmt.Printf("Total tickets available: %v & Remaining tickets available: %v \n", totalTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

//below we use "_" rather than "index". Index returns the index # and we don't need that. So we just use "_" as a placeholder. can't leave it blank.
//booking is item and bookings is slice. each booking in bookings
func getFirstnames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {

	remainingTickets = remainingTickets - userTickets
	//create a map for a user, we will store multiple maps in a slice. Each map represents a specific user's total data

	

	var userData = UserData {
		firstName: firstName,
		lastName:  lastName,
		email: email,
		numberOfTickets: userTickets,
	}
	//below we were setting key values pairs for our userData map. Above we are setting the fields in our UserData struct
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings: %v\n", bookings)

	fmt.Printf("Thanks for your order %v %v, you will receive %v tickets in your email: %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastname string, email string) {
	time.Sleep(10* time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastname)
	fmt.Println("################################")
	fmt.Printf("Sending ticket:\n %v to email address: %v \n", ticket, email)
	fmt.Println("################################")
	wg.Done()
}

func calculateRemainingTickets(remainingTickets uint, userTickets uint) uint {
	remainingTickets -= userTickets
	return remainingTickets
}



func getUserInput() (string, string, string, uint) {
	//we have to declare these vars within the function. We don't want them as params because we want the user to enter them in
	//the way we have it set up here, we call this func in main. User is prompted for info -> we use scan with &pointer ->
	//...the scan result is stored in local func vars (lastName, email, etc) and the values within the vars are returned. 
	//hence the (string, string...) return types in the func delcaration. in main we just declare multiple vars to capature return vals
	//I explain more of this in main

	var firstName string
		var lastName string
		var email string
		var userTickets uint

	fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter the number of tickets you would like to purchase: ")
		fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}