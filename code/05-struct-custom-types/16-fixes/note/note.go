package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string
	Content   string
	CreatedAt time.Time
}

func (n Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v\n\n", n.Title, n.Content)
}

func (n Note) Save() error { // we dont need a pointer for the receiver because the method does not edit the note-app
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.MarshalIndent(n, "", "\t")
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644) // note that 0644 assigns read and edit permission to the owner of the file

}

// creating a constructor
func New(title, content string) (Note, error) {
	// add validations
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
