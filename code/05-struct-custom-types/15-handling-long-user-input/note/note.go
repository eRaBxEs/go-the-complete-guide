package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func (n Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v\n\n", n.title, n.content)
}

// creating a constructor
func New(title, content string) (Note, error) {
	// add validations
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}
	return Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}
