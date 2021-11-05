package main

import (
	"testing"
)

func TestNumSteps(t *testing.T) {
	t.Run("for abcdef", func(t *testing.T) {
		result := runner("abcdef", 5)
		expected := 609043

		if result != expected {
			t.Errorf("expected %d, got %d", expected, result)
		}
	})
}
