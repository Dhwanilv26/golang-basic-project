package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

func getNoteData() (string, string) {
	title := getUserInput("note title : ")

	content := getUserInput("note content : ")

	return title, content
}

func getTodoData() string {
	text := getUserInput("todo text:")
	return text
}

func main() {
	title, content := getNoteData()
	todoText := getTodoData()

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}
	err = outputData(todo)

	if err != nil {
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Print(err)
		return
	}

	err = outputData(userNote)

	if err != nil {
		return
	}
}

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputtable interface {
	saver
	displayer
}

func saveData(data saver) error {

	err := data.Save()

	if err != nil {
		fmt.Print(err)
		return err
	}
	fmt.Print("saving the data is successful")
	return nil
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}
func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r") // because the text will still contain the \r character which is present in windows

	return text

}

// we are using T as a generic so the type of a b and result will be determined in runtime only, and we dont have to deal with the value.(type) or value.(int), value.(float64) and all of that mess if we used a long interface code
func add[T int | float64 | string](a, b T) T {
	return a + b
}
