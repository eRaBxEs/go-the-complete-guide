package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/interface/08-extracting-type-information/note"
	"example.com/interface/08-extracting-type-information/todo"
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

	intVal, ok := value.(int)
	if ok {
		fmt.Println("Integer:", intVal)
		return
	}

	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("Float:", floatVal)
		return
	}

	stringVal, ok := value.(string)
	if ok {
		fmt.Println("String:", stringVal)
		return
	}

	// switch value.(type) {
	// case int:
	// 	fmt.Println("printSomething(value interface{}) :: Integer:", value)
	// case float64:
	// 	fmt.Println("printSomething(value interface{}) :: Float:", value)
	// case string:
	// 	fmt.Println("printSomething(value interface{}) :: String:", value)
	// }
}

func printAnything(value any) { // where type any is an alias for interface{}
	switch value.(type) {
	case int:
		fmt.Println("printAnything(value any) :: Integer:", value)
	case float64:
		fmt.Println("printAnything(value any) :: Float:", value)
	case string:
		fmt.Println("printAnything(value any) :: String:", value)
	}
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
