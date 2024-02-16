package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func (u user) outputUserDetails() { // here we have (u user) just before the name of the function and this is called receiver
	fmt.Println(u.firstName, u.lastName, u.birthdate) // go allows u to access the fields without the need to dereference
}

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// now we have a utility function which can noq create a struct
func newUser(firstName, lastName, birthdate string) *user { // In order to avoid a copy to be returned but rather the original value
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user = newUser(userFirstName, userLastName, userBirthdate)

	// ... do something awesome with that gathered data!

	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
