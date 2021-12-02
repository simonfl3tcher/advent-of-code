package main

import (
	"advent-of-code/utils"
	"fmt"
	"regexp"
	"strconv"
)

const moveRegex = `^(down|up|forward)\s(\d)$`

var regx *regexp.Regexp

func init() {
	regx = regexp.MustCompile(moveRegex)
}

type Direction int

const (
	forward Direction = iota
	up
	down
)

type Move struct {
	direction Direction
	distance  int
}

type Tracker struct {
	depth     int
	horizonal int
	aim       int
}

func (t *Tracker) totalDepth() int {
	return t.horizonal * t.depth
}

func stringToDirection(str string) Direction {
	switch str {
	case "forward":
		return forward
	case "up":
		return up
	case "down":
		return down
	default:
		panic("not represented")
	}
}

func linesToMoves(lines []string) []Move {
	var moves []Move
	for _, move := range lines {
		matches := regx.FindStringSubmatch(move)
		direction := matches[1]
		distance, _ := strconv.Atoi(matches[2])
		moves = append(moves, Move{direction: stringToDirection(direction), distance: distance})
	}

	return moves
}

func part2(lines []string) int {
	t := Tracker{}
	for _, m := range linesToMoves(lines) {
		switch m.direction {
		case forward:
			t.horizonal += m.distance
			t.depth += t.aim * m.distance
		case up:
			t.aim -= m.distance
		case down:
			t.aim += m.distance
		}
	}
	return t.totalDepth()
}

func part1(lines []string) int {
	t := Tracker{}
	for _, m := range linesToMoves(lines) {
		switch m.direction {
		case forward:
			t.horizonal += m.distance
		case up:
			t.depth -= m.distance
		case down:
			t.depth += m.distance
		}
	}

	return t.totalDepth()
}

func main() {
	lines := utils.FileLinesToSlice("input.txt")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}
