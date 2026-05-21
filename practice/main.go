package main

import (
	"bufio"
	"fmt"
	"os"
	"practice/note"
	"strings"
)

func main() {
	title, description := getNoteData()
	userNote, err := note.New(title, description)

	userNote.DisplayNote()
	if err != nil {
		fmt.Println(err)
		return
	}
}
func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	description := getUserInput("Note description:")
	return title, description
}

func getUserInput(prompt string) string {
	fmt.Println(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
