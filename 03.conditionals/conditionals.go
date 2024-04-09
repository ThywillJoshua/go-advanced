package main

import "fmt"

const (
	Monday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

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

func main() {
	today, role := Tuesday, Contractor

	if role == Admin || role == Manager {
		accessGranted()
		return
	}

	if role == Contractor && today >= Saturday {
		accessGranted()
		return
	}

	if role == Member && today <= Friday {
		accessGranted()
		return
	}

	if role == Guest && (today == Monday || today == Wednesday || today == Friday) {
		accessGranted()
		return
	}

	accessDenied()
}
