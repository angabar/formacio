package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 1.1
	/* for i := 0; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	} */

	// 1.2 Lo podemos hacer de dos formas
	/* for i := 0; i < len(os.Args); i++ {
		fmt.Println(i, os.Args[i])
	}

	for i, v := range os.Args {
		fmt.Println(i, v)
	} */

	// 1.3
	fmt.Println(customSlice([]string{"test", "uno", "dos", "tres"}, ",,,,"))

	fmt.Println(countDup())
}

func customSlice(customArray []string, separator string) string {
	result := ""
	for i, v := range customArray {
		if i != len(customArray)-1 {
			result += v + separator
			continue
		}
		result += v
	}
	return result
}

func countDup() map[string]int {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	return counts
}
