package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/interface/06-special-any-value-interface/note"
	"example.com/interface/06-special-any-value-interface/todo"
)

type saver interface { // using a convention in go to name an interface based on the method contract it posseses
	Save() error
}

// type displayer interface {
// 	Display()
// }

type outputtable interface {
	saver
	Display()
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

func main() {

	printSomething(1)
	printAnything(1)
	printSomething(1.5)
	printAnything(1.5)
	printSomething("Learn Go now")
	printAnything("Learn Go now")

	title, content := getNoteData()

	todoText := getUserInput("Todo text:")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	printSomething(todo) // does nothing because the type switch defined in the printSomething does nothing for value of this type.

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}

	err = outputData(userNote)
	if err != nil {
		return
	}

}

func printSomething(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("Integer:", value)
	case float64:
		fmt.Println("Float:", value)
	case string:
		fmt.Println("String:", value)
	}
}

func printAnything(value any) { // where type any is an alias for interface{}
	fmt.Println(value)
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {

	err := data.Save()
	if err != nil {
		fmt.Println("Saving the data failed!")
		return err
	}

	fmt.Println("Saving the data succeeded!")
	return nil
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
