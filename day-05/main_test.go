package main

import (
	"testing"
)

func TestNumSteps(t *testing.T) {
	t.Run("part 1", func(t *testing.T) {
		t.Run("ugknbfddgicrmopn is nice", func(t *testing.T) {
			checker := strChecker{"ugknbfddgicrmopn"}
			result := checker.valid(1)
			expected := true

			if result != expected {
				t.Errorf("expected %t, got %t", expected, result)
			}
		})

		t.Run("aaa is nice", func(t *testing.T) {
			checker := strChecker{"aaa"}
			result := checker.valid(1)
			expected := true

			if result != expected {
				t.Errorf("expected %t, got %t", expected, result)
			}
		})

		t.Run("jchzalrnumimnmhp is naughty because it has no double letters", func(t *testing.T) {
			checker := strChecker{"jchzalrnumimnmhp"}
			result := checker.valid(1)
			expected := false

			if result != expected {
				t.Errorf("expected %t, got %t", expected, result)
			}
		})

		t.Run("haegwjzuvuyypxyu is naughty because it contains the string xy", func(t *testing.T) {
			checker := strChecker{"haegwjzuvuyypxyu"}
			result := checker.valid(1)
			expected := false

			if result != expected {
				t.Errorf("expected %t, got %t", expected, result)
			}
		})

		t.Run("dvszwmarrgswjxmb is naughty because it contains only one vowel", func(t *testing.T) {
			checker := strChecker{"dvszwmarrgswjxmb"}
			result := checker.valid(1)
			expected := false

			if result != expected {
				t.Errorf("expected %t, got %t", expected, result)
			}
		})
	})

	// Refactor to table test???
	t.Run("part 2", func(t *testing.T) {
		t.Run("qjhvhtzxzqqjkmpb is nice", func(t *testing.T) {
			checker := strChecker{"qjhvhtzxzqqjkmpb"}
			result := checker.valid(2)
			expected := true

			if result != expected {
				t.Errorf("expected %t, got %t", expected, result)
			}
		})

		t.Run("xxyxx is nice", func(t *testing.T) {
			checker := strChecker{"xxyxx"}
			result := checker.valid(2)
			expected := true

			if result != expected {
				t.Errorf("expected %t, got %t", expected, result)
			}
		})

		t.Run("uurcxstgmygtbstg is naughty", func(t *testing.T) {
			checker := strChecker{"uurcxstgmygtbstg"}
			result := checker.valid(2)
			expected := false

			if result != expected {
				t.Errorf("expected %t, got %t", expected, result)
			}
		})

		t.Run("ieodomkazucvgmuy is naughty", func(t *testing.T) {
			checker := strChecker{"uurcxstgmygtbstg"}
			result := checker.valid(2)
			expected := false

			if result != expected {
				t.Errorf("expected %t, got %t", expected, result)
			}

		})
	})
}
