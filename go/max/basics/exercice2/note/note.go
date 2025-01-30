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

func (note *Note) Display() {
	fmt.Printf("Your note title: %s, has the following content:\n\n %s", note.title, note.content)
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Input cant be empty")
	}

	return Note{
		title: title, content: content, createdAt: time.Now(),
	}, nil
}
