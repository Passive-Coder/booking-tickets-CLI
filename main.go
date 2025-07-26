package main

import (
	"booking-app/asset"
	"booking-app/helper"
	"booking-app/shared"
	"fmt"
	"strings"
)

var bookings = make([]shared.UserData, 0)

func main() {
	var line string = "-"
	var increment = 1
	for increment <= 100 {
		line += "-"
		increment++
	}
	fmt.Println()
	fmt.Println(line)
	fmt.Println("Welcome to Conference Booking CLI!!")
	fmt.Println(line)
	fmt.Println()

	for {
		fmt.Println()
		fmt.Println("Conferences Available are: ")
		for index, document := range asset.Conferences {
			fmt.Printf("Choice %v\n", index+1)
			fmt.Print("   Conference Name: ")
			fmt.Println(document.Name)
			fmt.Print("   Seats Available: ")
			fmt.Println(document.SeatsAvailable)
		}
		var choice uint
		fmt.Println()
		fmt.Printf("Enter choice (1-%d): ", len(asset.Conferences))
		fmt.Scan(&choice)

		if choice >= 1 && choice <= uint(len(asset.Conferences)) {
			var remainingTickets uint = asset.Conferences[choice-1].SeatsAvailable
			name, email, tickets := helper.GetInput()
			isValidName, isValidEmail, isValidTicket := isValid(name, email, tickets, remainingTickets)
			if isValidName && isValidEmail && isValidTicket {
				remainingTickets -= tickets
				asset.Conferences[choice-1].SeatsAvailable -= tickets
				var userData = shared.UserData{
					Name:            name,
					Email:           email,
					NumberOfTickets: tickets,
				}

				bookings = append(bookings, userData)

				fmt.Printf("%v ticket/s for %v\n", tickets, name)

				fmt.Printf("Remaining Tickets: %v\n", remainingTickets)
				if remainingTickets == 0 {
					fmt.Printf("All tickets are booked")
					break
				}

				go helper.SendTicket(asset.Conferences[choice-1], userData)

			} else {
				if !isValidName {
					fmt.Println("Name was not valid")
				}
				if !isValidEmail {
					fmt.Println("Email was not valid")
				}
				if !isValidTicket {
					fmt.Println("Excess tickets requested")
				}
			}
		} else {
			fmt.Println("Wrong choice. Try again!")
			continue
		}
	}
}

func isValid(name string, email string, tickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(name) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := !(tickets > remainingTickets)
	return isValidName, isValidEmail, isValidTicket
}
