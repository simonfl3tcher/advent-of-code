package main

import (
	"strconv"
	"testing"
)

func TestPart1(t *testing.T) {
	result := part1([]string{strconv.Quote(""), strconv.Quote("abc"), strconv.Quote("aaa\"aaa"), strconv.Quote("\x27")})
	expected := 12

	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestPart2(t *testing.T) {
	result := part2([]string{strconv.Quote(""), strconv.Quote("abc"), strconv.Quote("aaa\"aaa"), strconv.Quote("\x27")})
	expected := 19

	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
