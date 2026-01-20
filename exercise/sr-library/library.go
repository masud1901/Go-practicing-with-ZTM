//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:

//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned

type Title string
type Name string

type LendAudit struct {
	checkOut time.Time
	checkIn  time.Time
}

type member struct {
	name  Name
	books map[Title]LendAudit
}

type bookEntry struct {
	total  int
	lended int
}

type Library struct {
	menbers map[Name]member
	books   map[Title]bookEntry
}

func printMemberAudit(member *member) {
	for title, audit := range member.books {
		var returnTime string
		if audit.checkIn.IsZero() {
			returnTime = "Not returned yet"
		} else {
			returnTime = audit.checkIn.String()
		}
		fmt.Printf("Book: %s, Checked out at: %s, Returned at: %s\n", title, audit.checkOut.String(), returnTime)
	}
}

func printAllMemberAudit(library *Library) {
	for _, member := range library.menbers {
		printMemberAudit(&member)
	}
}

func printLibraryBooks(library *Library) {
	fmt.Println()
	for title, book := range library.books {
		fmt.Printf("Title: %s, Total: %d, Lended: %d\n", title, book.total, book.lended)
	}
	fmt.Println()
}

func checkoutBook(library *Library, title Title, member *member) bool {
	book, found := library.books[title]
	if !found || book.lended >= book.total {
		return false
	}

	book.lended++
	library.books[title] = book

	member.books[title] = LendAudit{
		checkOut: time.Now(),
	}
	return true
}

func returnBook(library *Library, title Title, member *member) bool {
	book, found := library.books[title]
	if !found {
		return false
	}

	audit, found := member.books[title]
	if !found || !audit.checkIn.IsZero() {
		return false
	}

	book.lended--
	library.books[title] = book

	audit.checkIn = time.Now()
	member.books[title] = audit
	return true
}

func main() {
	Library := Library{
		menbers: make(map[Name]member),
		books:   make(map[Title]bookEntry),
	}

	Library.books["Book A"] = bookEntry{total: 2}
	Library.books["Book B"] = bookEntry{total: 1}
	Library.books["Book C"] = bookEntry{total: 3}
	Library.books["Book D"] = bookEntry{total: 1}

	Library.menbers["Alice"] = member{name: "Alice", books: make(map[Title]LendAudit)}
	Library.menbers["Bob"] = member{name: "Bob", books: make(map[Title]LendAudit)}
	Library.menbers["Charlie"] = member{name: "Charlie", books: make(map[Title]LendAudit)}

	fmt.Println("Initial Library State:")
	printLibraryBooks(&Library)
	printAllMemberAudit(&Library)

	fmt.Println("Alice checks out Book A:")
	aliceMember := Library.menbers["Alice"]
	checkoutBook(&Library, "Book A", &aliceMember)
	Library.menbers["Alice"] = aliceMember
	printLibraryBooks(&Library)
	printAllMemberAudit(&Library)

	fmt.Println("Bob checks out Book B:")
	bobMember := Library.menbers["Bob"]
	checkoutBook(&Library, "Book B", &bobMember)
	Library.menbers["Bob"] = bobMember
	printLibraryBooks(&Library)
	printAllMemberAudit(&Library)

	fmt.Println("Alice returns Book A:")
	aliceMember = Library.menbers["Alice"]
	returnBook(&Library, "Book A", &aliceMember)
	Library.menbers["Alice"] = aliceMember
	printLibraryBooks(&Library)
	printAllMemberAudit(&Library)

}
