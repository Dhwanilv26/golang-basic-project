package main

import (
	"fmt"

	"example.com/note/note"
)

func getNoteData() (string, string) {
	title := getUserInput("note title:")

	content := getUserInput("note content:")

	return title, content
}

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("%+v", userNote)
}

func getUserInput(prompt string) string {
	fmt.Println(prompt)
	var value string
	fmt.Scanln(&value)
	if value == "" {
		return ""
	}
	return value

}
