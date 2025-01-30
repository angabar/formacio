package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := user.New(firstName, lastName, birthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("admin@admin.com", "123admin")

	admin.OutputData()

	appUser.OutputData()
	appUser.ClearUserName()
	appUser.OutputData()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

/*

Los metodos tambien pueden asignarse a los tipos custom

type str string

func (text str) log() {
	fmt.Println(text)
}

...

name := "Max"
name.log()

*/
