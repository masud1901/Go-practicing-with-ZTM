//--Summary:
//  The existing program is used to restrict access to a resource
//  based on day of the week and user role. Currently, the program
//  allows any user to access the resource. Implement the functionality
//  needed to grant resource access using any combination of `if`, `else if`,
//  and `else`.
//
//--Requirements:
//* Use the accessGranted() and accessDenied() functions to display
//  informational messages

package main

import "fmt"

// Days of the week
const (
	Monday    = 0
	Tuesday   = 1
	Wednesday = 2
	Thursday  = 3
	Friday    = 4
	Saturday  = 5
	Sunday    = 6
)

// User roles
const (
	Admin      = 10
	Manager    = 20
	Contractor = 30
	Member     = 40
	Guest      = 50
)

func accessGranted() {
	fmt.Println("Granted")
}

func accessDenied() {
	fmt.Println("Denied")
}

func access(day, role int) bool {

	if role == Admin || role == Manager {
		//* Access at any time: Admin, Manager

		return true

	} else if role == Contractor {
		//* Access weekends: Contractor
		if day == Saturday {
			return true
		} else {
			return false
		}
	} else if role == Member {
		//* Access weekdays: Member
		if day != Saturday {
			return true
		} else {
			return false
		}
	} else {
		//* Access Mondays, Wednesdays, and Fridays: Guest
		if day%2 == 0 && day != Sunday {
			return true
		} else {
			return false
		}

	}
}

func accessOptimized(day, role int) bool {
	switch role {
	case Admin, Manager:
		return true // Always allowed

	case Contractor:
		return day == Saturday || day == Sunday // Weekends only

	case Member:
		return day >= Monday && day <= Friday // Weekdays only

	case Guest:
		return day == Monday || day == Wednesday || day == Friday // Specific days

	default:
		return false // Unknown role = denied
	}
}

func main() {
	// The day and role. Change these to check your work.
	today, role := Saturday, Contractor

	if access(today, role) {

		accessGranted()
	} else {
		accessDenied()
	}
}
