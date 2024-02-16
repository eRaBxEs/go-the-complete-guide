package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func (u User) OutputUserDetails() { // here we have (u user) just before the name of the function and this is called receiver
	fmt.Println(u.firstName, u.lastName, u.birthdate) // go allows u to access the fields without the need to dereference
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthdate: "----",
		},
	}
}

func New(firstName, lastName, birthdate string) (*User, error) { // In order to avoid a copy to be returned but rather the original value
	// using this constructor kind of functions enables us to add validations,
	// so we can always properly check for expectations at the point of creation
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("firstname, lastname and birthdate are required")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
