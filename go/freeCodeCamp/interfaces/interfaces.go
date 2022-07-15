package main

// Las interfaces definen un comportamiento, no un tipo, estas son para definir
// que requisitos debe tener un struct u otro tipo de datos para poder usarla
type User interface {
	Permissions() int
	Name() string
}

type Admin struct {
	name string
}

// Como vemos las interfaces garantizan que los métodos que se necesitan vayan a
// existir, por ejemplo aquí solicitmos Permissions y Name, cualquier tipo que
// vaya a utilizar este método deberá primero implementar la interfaz

// No es necesario usar implements, ni nada por el estilo, Go entiende que un
// struct cumple con una interfaz si tiene todos los criterios de la misma
func getAuth(u User) string {
	if u.Permissions() > 4 {
		return u.Name() + " tiene permisos de admin"
	}

	return ""
}

func (ad Admin) Permissions() int {
	return 5
}

func (ad Admin) Name() string {
	return ad.name
}

func main() {
	admin := Admin{name: "test"}
	getAuth(admin)
}
