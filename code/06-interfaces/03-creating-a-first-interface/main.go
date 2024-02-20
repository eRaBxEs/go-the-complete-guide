package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/interface/02-finishing-interface-preparations/note"
	"example.com/interface/02-finishing-interface-preparations/todo"
)

type saver interface { // using a convention in go to name an interface based on the method contract it posseses
	Save() error
}

func main() {

	title, content := getNoteData()

	todoText := getUserInput("Todo text:")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	todo.Display()
	err = todo.Save()
	if err != nil {
		fmt.Println("Saving the todo failed!")
		return
	}

	fmt.Println("Saving the todo succeeded!")

	userNote.Display()
	err = userNote.Save()
	if err != nil {
		fmt.Println("Saving the note failed!")
		return
	}

	fmt.Println("Saving the note succeeded!")

}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
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
