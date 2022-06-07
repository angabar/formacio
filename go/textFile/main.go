package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	arguments := os.Args

	if len(arguments) >= 2 {
		fileName := arguments[1]

		file, err := os.Open(fileName)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(file)

		// Read does not push the byte slice, so we need to create it with
		// enough spaces to fill
		data := make([]byte, 100)

		count, err := file.Read(data)

		if err != nil {
			log.Fatal(err)
		}

		// if we use data instead of data[:count] we can see all the remaining
		// empty spaces
		fmt.Printf("read %d bytes: %q\n", count, data[:count])
	}
}
