package main

import (
	"os"
	"strings"
	"testing"
)

// go test -bench=.
func BenchmarkSlice(b *testing.B) {
	strings.Join(os.Args, "++")
}

func BenchmarkCustomSlice(b *testing.B) {
	customSlice(os.Args, "++")
}
