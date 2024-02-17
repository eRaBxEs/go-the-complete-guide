package note

import (
	"errors"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
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
