package utils

import "testing"

func TestMinFromSlice(t *testing.T) {
	result := MinFromSlice([]int{9, 4, 1})
	expected := 1

	if result != expected {
		t.Errorf("expected result to be %d, got%d", result, expected)
	}
}
