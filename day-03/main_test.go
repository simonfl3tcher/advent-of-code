package main

import (
	"testing"
)

func TestNumSteps(t *testing.T) {
	t.Run("for a single santa", func(t *testing.T) {
		result := numSteps("^>v<", 1)
		expected := 4

		if result != expected {
			t.Errorf("expected %d, got %d", expected, result)
		}
	})

	t.Run("for two santas", func(t *testing.T) {
		result := numSteps("^>v<", 2)
		expected := 3

		if result != expected {
			t.Errorf("expected %d, got %d", expected, result)
		}
	})
}
