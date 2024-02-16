package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func (u User) OutputUserDetails() { // here we have (u user) just before the name of the function and this is called receiver
	fmt.Println(u.FirstName, u.lastName, u.birthdate) // go allows u to access the fields without the need to dereference
}

func (u *User) ClearUserName() {
	u.FirstName = ""
	u.lastName = ""
}

func NewUser(firstName, lastName, birthdate string) (*User, error) { // In order to avoid a copy to be returned but rather the original value
	// using this constructor kind of functions enables us to add validations,
	// so we can always properly check for expectations at the point of creation
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("firstname, lastname and birthdate are required")
	}
	return &User{
		FirstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
