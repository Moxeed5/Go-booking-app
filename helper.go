package main

import "strings"


func ValidateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >=2
	isValidEmail := strings.Contains(email,"@") && strings.Contains(email,".") 
	isValidTickets := userTickets >=1 && userTickets <= RemainingTickets

	return isValidName, isValidEmail, isValidTickets
}
