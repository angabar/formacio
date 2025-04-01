package maps

import "fmt"

func main() {
	websites := map[string]string{
		"Google":              "google.com",
		"Amazon Web Services": "aws.com",
	}

	fmt.Println(websites["Google"])

	websites["Forocoches"] = "forocoches.es"

	fmt.Println(websites)

	// Con la funcion delete borramos un par de clave/valor del map
	delete(websites, "Google")

	fmt.Println(websites)
}
