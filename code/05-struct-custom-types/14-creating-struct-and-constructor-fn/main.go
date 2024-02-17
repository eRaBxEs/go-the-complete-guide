package main

import (
	"fmt"

	"example.com/structs/14-creating-struct-and-constructor-fn/note"
)

func main() {

	title, content := getNoteData()

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(userNote)

}

func getNoteData() (string, string) {
	title := getuserInput("Note title:")

	content := getuserInput("Note content:")

	return title, content
}

func getuserInput(prompt string) string {
	fmt.Print(prompt)
	var value string
	fmt.Scanln(&value)

	return value
}
