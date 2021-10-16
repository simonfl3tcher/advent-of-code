package main

import (
	"reflect"
	"testing"
)

func TestNewBoxDimensions(t *testing.T) {
	result := newBoxDimensions("2x3x4")
	expected := boxDimensions{2, 3, 4}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v to be %v", result, expected)
	}
}

func TestBoxDimensions(t *testing.T) {
	t.Run("calculate paper should return correct result", func(t *testing.T) {
		b := boxDimensions{2, 3, 4}
		result := b.CalculatePaper()
		expected := 58

		if result != expected {
			t.Errorf("expected result to be %d, got %d", expected, result)
		}
	})

	t.Run("calculate ribbon should return correct result", func(t *testing.T) {
		b := boxDimensions{2, 3, 4}
		result := b.CalculateRibbon()
		expected := 34

		if result != expected {
			t.Errorf("expected result to be %d, got %d", expected, result)
		}
	})
}
