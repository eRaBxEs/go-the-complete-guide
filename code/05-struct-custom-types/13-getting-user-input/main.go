package main

import (
	"errors"
	"fmt"
)

func main() {

	title, content, err := getNoteData()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(title, content)

}

func getNoteData() (string, string, error) {
	title, err := getuserInput("Note title:")
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	content, err := getuserInput("Note content:")
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	return title, content, nil
}

func getuserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	var value string
	fmt.Scanln(&value)

	if value == "" {
		return "", errors.New("invalid input")
	}

	return value, nil
}
