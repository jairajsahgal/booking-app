package main

import "fmt"

func main() {
	var conferenceName string= "Go Conference"
	const conferenceTickets = 50
	var remainingTickets int = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets,conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int
	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets", userName, userTickets)
}