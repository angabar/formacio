package example

import "fmt"

func ExampleHello() {
	greeting := Hello("test")
	fmt.Println(greeting)

	// Output:
	// Hello test
}

func ExamplePage() {
	checkIns := map[string]bool{
		"Alice":   false,
		"Vinsent": true,
		"Paco":    false,
	}

	Page(checkIns)

	// Unordered output:
	// Paging Alice; Please see the front
	// Paging Paco; Please see the front
}
