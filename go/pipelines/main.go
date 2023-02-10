package main

import (
	"fmt"
	"strings"
)

func topStream(downstream chan string) {
	names := []string{"test1", "placa", "jaja"}

	for _, v := range names {
		downstream <- v
	}

	close(downstream)
}

func intermediateStream(upstream, downstream chan string) {
	for {
		v, ok := <-upstream

		if !ok {
			close(upstream)
			return
		}

		if strings.Contains(v, "placa") {
			downstream <- v
		}
	}
}

func bottomStream(upstream chan string) {
	for _, v := range <-upstream {
		fmt.Println(v)
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go topStream(ch1)
	go intermediateStream(ch1, ch2)
	bottomStream(ch2)
}
