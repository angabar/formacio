package user

import (
	"errors"
	"fmt"
	"time"
)

// La exportacion de un struct tambien se hace por campos, es decir, puedo
// exportar un strcut como User, pero no sus campos, esto se determina en la
// primera letra, si es mayuscula se exporta, sino no

// Un uso normal es exportar el User para ser usado como tipo por otras
// variables, pero no los campos
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	User
	email    string
	password string
}

// Lo normal es poner solo "New" porque ya viene acompañado del nombre del
// paquete: user.New
func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("All data is mandatory")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string) *Admin {
	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthDate: "--/--/----",
			createdAt: time.Now(),
		},
	}
}

func (u *User) OutputData() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}
