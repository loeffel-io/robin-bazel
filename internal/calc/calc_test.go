package calc

import (
	"testing"
)

func TestCalc_Add(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Error("1 + 2 != 3")
	}
}
