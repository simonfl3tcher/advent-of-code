package main

import (
	"testing"
)

func TestPositionTrackerMove(t *testing.T) {
	t.Run("move up", func(t *testing.T) {
		pt := PositionTracker{}
		pt.Move('(', 0)

		if pt.sum != 1 {
			t.Errorf("expected sum to be 1, got %d", pt.sum)
		}
	})

	t.Run("move down", func(t *testing.T) {
		pt := PositionTracker{}
		pt.Move(')', 0)

		if pt.sum != -1 {
			t.Errorf("expected sum to be -1, got %d", pt.sum)
		}
	})

	t.Run("combines multiple moves", func(t *testing.T) {
		pt := PositionTracker{}
		pt.Move('(', 0)
		pt.Move(')', 1)
		pt.Move('(', 2)
		pt.Move('(', 3)

		if pt.sum != 2 {
			t.Errorf("expected sum to be 2, got %d", pt.sum)
		}
	})

	t.Run("sets the first index which causes -1 sum", func(t *testing.T) {
		pt := PositionTracker{}
		pt.Move('(', 0)
		pt.Move('(', 1)
		pt.Move(')', 2)
		pt.Move(')', 3)
		pt.Move(')', 4)

		if pt.indexWhichCausedBasement != 5 {
			t.Errorf("expected indexWhichCausedBasement to be 5, got %d", pt.indexWhichCausedBasement)
		}
	})

	t.Run("does not override first index which causes -1 sum if there are multiple", func(t *testing.T) {
		pt := PositionTracker{}
		pt.Move(')', 0)
		pt.Move('(', 1)
		pt.Move(')', 2)
		pt.Move(')', 3)
		pt.Move(')', 4)

		if pt.indexWhichCausedBasement != 1 {
			t.Errorf("expected indexWhichCausedBasement to be 1, got %d", pt.indexWhichCausedBasement)
		}
	})
}
