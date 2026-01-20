package main

import "fmt"

type Passenger struct {
	name         string
	ticketNumber int
	boarded      bool
}

type Bus struct {
	frontSeat Passenger
}

func main() {
	casey := Passenger{name: "Casey", ticketNumber: 1234, boarded: false}
	// bus := Bus{frontSeat: casey}
	fmt.Println(casey)
	bill := Passenger{name: "Bill", ticketNumber: 5678, boarded: true}
	ella := Passenger{name: "Ella", ticketNumber: 9101, boarded: true}
	fmt.Println(bill, ella)

	var heidi Passenger
	heidi.name = "Heidi"
	heidi.ticketNumber = 1121
	heidi.boarded = false
	fmt.Println(heidi)
	casey.boarded = true
	bill.boarded = true

	if bill.boarded {
		fmt.Printf("%s has boarded the bus.\n", bill.name)
	}

	
}
