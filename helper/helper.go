package helper

import (
	"booking-app/asset"
	"booking-app/shared"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var reader = bufio.NewReader(os.Stdin)

func GetFirstNames(bookings []map[string]string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var splittedName = strings.Fields(booking["name"])
		firstNames = append(firstNames, splittedName[0])
	}
	return firstNames
}

func GetInput() (string, string, uint) {

	fmt.Printf("Enter name: ")
	var name string
	name, _ = reader.ReadString('\n')

	var email string
	fmt.Printf("Enter email: ")
	fmt.Scan(&email)

	fmt.Printf("Enter no.of tickets to book: ")
	var tickets uint
	fmt.Scan(&tickets)

	return name, email, tickets
}

func SendTicket(details asset.Conference, userDetails shared.UserData) {
	time.Sleep(time.Second * 10)
	fmt.Println()
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println("Ticket sent")
	fmt.Println("Ticker details:")
	fmt.Println("Conference name: ", details.Name)
	fmt.Println("Name of participant: ", userDetails.Name)
	fmt.Println("Number of tickets: ", userDetails.NumberOfTickets)
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------------------------")
}
