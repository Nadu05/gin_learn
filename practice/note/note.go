package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	CreateDate  string `json:"createDate"`
}

func New(title, description string) (*Note, error) {

	if title == "" || description == "" {
		return &Note{}, errors.New("title or description is empty")
	}
	return &Note{
		Title:       title,
		Description: description,
		CreateDate:  time.Now().Format("2006-01-02 15:04:05"),
	}, nil
}

func (note Note) DisplayNote() {
	fmt.Printf("Note title is : %v  and note content is %v \n\n",
		note.Title, note.Description)

}
