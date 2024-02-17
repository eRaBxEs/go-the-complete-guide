package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/structs/15-handling-long-user-input/note"
)

func main() {

	title, content := getNoteData()

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()

}

func getNoteData() (string, string) {
	title := getuserInput("Note title:")

	content := getuserInput("Note content:")

	return title, content
}

func getuserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	// rather use fmt.Scanln use buffer Reader
	reader := bufio.NewReader(os.Stdin)  // os.Stdin indicates that we are reading from the command line
	text, err := reader.ReadString('\n') // delimiter : to indicate where we want to stop reading

	if err != nil {
		return ""
	}

	// Now after getting the read data in text, we will need to remove the line break from it (\n)
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r") // removing \r because on windows each line break comes with a carriage return "\r"

	return text
}
