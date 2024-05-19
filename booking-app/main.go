package main

import (
	"fmt"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = conferenceTickets

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")

	var bookings = [50]string{} //syntaxt: [size]dataType{Value}

	var firstName string
	var lastName string
	var email string

	var userTickets uint

	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName) // & means memory pointer - passing reference

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	bookings[0] = firstName + " " + lastName

	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("This first value: %v\n", bookings[0])
	fmt.Printf("Array length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v ticktets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
