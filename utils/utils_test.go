package utils

import (
	"reflect"
	"testing"
)

func TestMinFromSlice(t *testing.T) {
	result := MinFromSlice([]int{9, 4, 1})
	expected := 1

	if result != expected {
		t.Errorf("expected result to be %d, got%d", result, expected)
	}
}

func TestMaxFromSlice(t *testing.T) {
	result := MaxFromSlice([]int{9, 10, 1})
	expected := 10

	if result != expected {
		t.Errorf("expected result to be %d, got%d", result, expected)
	}
}

func TestUnique(t *testing.T) {
	result := Unique([]string{"a", "b", "c", "a"})
	expected := []string{"a", "b", "c"}

	if reflect.DeepEqual(result, expected) == false {
		t.Errorf("expected result to be %v, got %v", result, expected)
	}
}

func TestPermutation(t *testing.T) {
	result := Permutation([]string{"a", "b", "c"})
	expected := [][]string{{"a", "b", "c"}, {"a", "c", "b"}, {"b", "a", "c"}, {"b", "c", "a"}, {"c", "a", "b"}, {"c", "b", "a"}}

	var matchCount int
	for _, r := range result {
		for _, e := range expected {
			if reflect.DeepEqual(r, e) {
				matchCount++
			}
		}
	}

	if matchCount != len(expected) {
		t.Errorf("expected result to be %v, got %v", result, expected)
	}
}
