package main

import (
	"fmt"

	"example.com/structs/10-struct-embedding/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User
	appUser, err := user.New(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("texst@example.com", "test123")
	admin.OutputUserDetails() // instead of admin.User.OutputUserDetails()
	admin.ClearUserName()     // instead of admin.User.ClearUserName()
	admin.OutputUserDetails()

	// ... do something awesome with that gathered data!

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
