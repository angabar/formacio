package utils

import (
	"testing"
)

func TestSum(t *testing.T) {
	sum := Sum([]int{1, 2, 3})
	if sum != 6 {
		t.Errorf("Wanted 11 but got %d", sum)
	}
}
