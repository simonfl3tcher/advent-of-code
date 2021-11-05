package main

import (
	"testing"
)

var dummyInput = []string{
	"London to Dublin = 464",
	"London to Belfast = 518",
	"Dublin to Belfast = 141",
}

func TestMinimumValue(t *testing.T) {
	result := run(dummyInput)
	expected := 605

	if result.minDistance != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestMaximumValue(t *testing.T) {
	result := run(dummyInput)
	expected := 982

	if result.maxDistance != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
