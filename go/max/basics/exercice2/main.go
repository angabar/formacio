package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note"
	"example.com/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	saver
	Display()
}

func main() {
	printSomething(1)
	fmt.Println(add(1, 3))

	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(&todo)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(&userNote)

	if err != nil {
		fmt.Println(err)
		return
	}
}

// interface{} es como any en typescript
func printSomething(value interface{}) {
	// Con esta sintaxis comprobamos que el tipo de la variable entrante, aunque
	// interesante, es mejor el sistema de switch para hacer estas
	// comprobaciones
	intVal, ok := value.(int)

	if ok {
		fmt.Println(intVal + 1)
		return
	}

	float64Val, ok := value.(float64)

	if ok {
		fmt.Println(float64Val + 1)
		return
	}

	// switch value.(type) {
	// case int:
	// 	fmt.Println("int")
	// case float64:
	// 	fmt.Println("float64")
	// case string:
	// 	fmt.Println("string")
	// }
}

func outputData(data displayer) error {
	data.Display()
	return saveData(data)

}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Saving the note succeeded!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")

	return input
}

// Esto es una funcion generica, donde establecemos entre | los tipos que puede
// tener el tipo generico
func add[T int | float64 | string](a, b T) T {
	return a + b
}
