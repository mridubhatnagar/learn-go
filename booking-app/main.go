package main

import (
	"fmt"
	"strings"
)


func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	// remaining tickets cannot be -ve hence uint
	var remainingTickets uint = 50
	// creates a bookings array of size 50 and type string
	var bookings[]string 

	fmt.Printf("conferenceTickets is %T\nconferenceName is %T\nremainingTickets is %T\n", conferenceTickets, conferenceName, remainingTickets)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string 
		var userTickets uint 


		// ask user for input. Whatever value user enters. Gets assigned to userName
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName + " " + lastName) 

			fmt.Printf("The whole slice: %v\n", bookings)
			fmt.Printf("The first value: %v\n", bookings[0])
			fmt.Printf("The type of slice: %T\n", bookings)
			fmt.Printf("Size of the slice: %v\n", len(bookings))

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("Remaining tickets are %v for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])

			} 
			if remainingTickets == 0 {
				fmt.Printf("Conference is sold out.")
				break 
			}
			fmt.Printf("These are all the bookings in the application: %v\n", firstNames)
			
			
		} else {
			fmt.Printf("Only %v tickets available.And you have booked %v\n", remainingTickets, userTickets)
		
		}
		
		

		
	}
	
	
}

