package main

import (
	"fmt"
	"time"
)

type (
	Title string
	Name  string
)

type LendAudit struct {
	checkOut time.Time
	checkIn  time.Time
}

type Member struct {
	name  Name
	books map[Title]LendAudit
}

type BookEntry struct {
	total  int // total owned by library
	lended int // total lended out
}

type Library struct {
	members map[Name]Member
	books   map[Title]BookEntry
}

func printMemberAudit(m *Member) {
	for title, audit := range m.books {
		var returnTime string
		if audit.checkIn.IsZero() {
			returnTime = "[not returned yet]"
		} else {
			returnTime = audit.checkIn.String()
		}

		fmt.Println(m.name, ":", title, ":", audit.checkOut.String(), "through", returnTime)
	}
}

func printMemberAudits(l *Library) {
	for _, member := range l.members {
		printMemberAudit(&member)
	}
}

func printLibraryBooks(l *Library) {
	fmt.Println()
	for title, book := range l.books {
		fmt.Println(title, "/ total:", book.total, "/lended:", book.lended)
	}
	fmt.Println()
}

func checkoutBook(l *Library, title Title, m *Member) bool {
	// Check book exists
	book, found := l.books[title]
	if !found {
		fmt.Println("Book doesn't exist in this Library")
		return false
	}

	// Check if there are availaible copies
	if book.lended == book.total {
		fmt.Println("No more books available to lend")
		return false
	}

	book.lended += 1
	l.books[title] = book

	m.books[title] = LendAudit{checkOut: time.Now()}
	return true
}

func returnBook(l *Library, title Title, m *Member) bool {
	book, found := l.books[title]
	if !found {
		fmt.Println("Book doesn't exist in this Library")
		return false
	}

	audit, found := m.books[title]
	if !found {
		fmt.Println("Member did not check this book out")
		return false
	}

	book.lended -= 1
	l.books[title] = book

	audit.checkIn = time.Now()
	m.books[title] = audit

	return true
}

func main() {
	library := Library{
		books:   make(map[Title]BookEntry),
		members: make(map[Name]Member),
	}

	library.books["Webapp in Go"] = BookEntry{
		total:  4,
		lended: 0,
	}
	library.books["Java Complete"] = BookEntry{
		total:  2,
		lended: 0,
	}
	library.books["Wild Wild West"] = BookEntry{
		total:  1,
		lended: 0,
	}
	library.books["Harry Potter"] = BookEntry{
		total:  1,
		lended: 0,
	}

	library.members["Jayson"] = Member{"Jayson", make(map[Title]LendAudit)}
	library.members["Anna"] = Member{"Anna", make(map[Title]LendAudit)}
	library.members["Oana"] = Member{"Oana", make(map[Title]LendAudit)}

	member := library.members["Jayson"]

	fmt.Println("\n Initial")
	printLibraryBooks(&library)
	printMemberAudit(&member)

	checkedOut := checkoutBook(&library, "Harry Potter", &member)
	fmt.Println("\nCheck out a book: ")
	if checkedOut {
		printLibraryBooks(&library)
		printMemberAudit(&member)
	}

	returned := returnBook(&library, "Harry Potter", &member)
	if returned {
		printLibraryBooks(&library)
		printMemberAudit(&member)
	}
}
