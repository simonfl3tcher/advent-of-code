package main

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"
)

type AdjacentNumbers struct {
	north *Number
	east  *Number
	south *Number
	west  *Number
}

type Number struct {
	num      int
	adjacent AdjacentNumbers
}

func newNumber(str string) *Number {
	i, _ := strconv.Atoi(str)
	return &Number{num: i}
}

type lineNumbers [][]*Number

func (ln *lineNumbers) lowPoints() []*Number {
	response := []*Number{}
	for _, line := range *ln {
		for _, n := range line {
			check := true
			if n.adjacent.north != nil {
				if n.num >= n.adjacent.north.num {
					check = false
				}
			}
			if n.adjacent.east != nil {
				if n.num >= n.adjacent.east.num {
					check = false
				}
			}
			if n.adjacent.south != nil {
				if n.num >= n.adjacent.south.num {
					check = false
				}
			}
			if n.adjacent.west != nil {
				if n.num >= n.adjacent.west.num {
					check = false
				}
			}
			if check {
				response = append(response, n)
			}
		}
	}
	return response
}

func main() {
	lines := utils.FileLinesToSlice("input.txt")
	linesInNumbers := lineNumbers{}
	for _, line := range lines {
		var l []*Number
		nums := strings.Split(line, "")
		for _, num := range nums {
			l = append(l, newNumber(num))
		}
		linesInNumbers = append(linesInNumbers, l)
	}

	for lineIndex, k := range linesInNumbers {
		for numIndex, j := range k {
			if numIndex > 0 {
				j.adjacent.west = k[numIndex-1]
			}
			if numIndex < len(k)-1 {
				j.adjacent.east = k[numIndex+1]
			}
			if lineIndex > 0 {
				j.adjacent.north = linesInNumbers[lineIndex-1][numIndex]
			}
			if lineIndex < len(linesInNumbers)-1 {
				j.adjacent.south = linesInNumbers[lineIndex+1][numIndex]
			}
		}
	}

	acc := 0
	for _, v := range linesInNumbers.lowPoints() {
		acc += 1 + v.num
	}

	fmt.Println(acc)
}
