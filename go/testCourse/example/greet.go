package example

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

func Page(checkIns map[string]bool) {
	for name, checkIn := range checkIns {
		if !checkIn {
			fmt.Printf("Paging %s; Please see the front\n", name)
		}
	}
}
