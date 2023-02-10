package mymath

import (
	"testing"
)

// Los archivos _test nunca van incluidos en la build al compilar
func TestSum(t *testing.T) {
	sum := Sum([]int{10, -2, 3})
	if sum != 11 {
		t.Errorf("FAIL: Wanted 11 but received %d", sum)
	}
}
