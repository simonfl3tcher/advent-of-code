package main

import (
	"fmt"
	"io/ioutil"
)

type PositionTracker struct {
	sum                      int
	indexWhichCausedBasement int
}

func (p *PositionTracker) Move(direction rune, currentMove int) {
	switch direction {
	case '(':
		p.sum++
	case ')':
		p.sum--
	}

	if p.sum == -1 && p.indexWhichCausedBasement == 0 {
		p.indexWhichCausedBasement = currentMove + 1
	}
}

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	pt := PositionTracker{}
	for index, char := range string(bytes) {
		pt.Move(char, index)
	}

	fmt.Printf("Final position: %d\n", pt.sum)
	fmt.Printf("Position of char which causes basement entry: %d\n", pt.indexWhichCausedBasement)
}
