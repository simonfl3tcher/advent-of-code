package main

import (
	"advent-of-code/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const reg = `(\d+,\d+)\s->\s(\d+,\d+)`

var regx *regexp.Regexp

func init() {
	regx = regexp.MustCompile(reg)
}

type Command struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func NewCommands(lines []string) []*Command {
	commands := []*Command{}
	for _, line := range lines {
		matches := regx.FindStringSubmatch(line)
		commands = append(commands, newCommand(matches[1], matches[2]))
	}
	return commands
}

func newCommand(cmd1, cmd2 string) *Command {
	cmd1Slice := strings.Split(cmd1, ",")
	cmd2Slice := strings.Split(cmd2, ",")
	x1, _ := strconv.Atoi(cmd1Slice[0])
	y1, _ := strconv.Atoi(cmd1Slice[1])
	x2, _ := strconv.Atoi(cmd2Slice[0])
	y2, _ := strconv.Atoi(cmd2Slice[1])

	return &Command{x1: x1, y1: y1, x2: x2, y2: y2}
}

func (c *Command) horizontal() bool {
	return c.y1 == c.y2
}

func (c *Command) vertical() bool {
	return c.x1 == c.x2
}

func (c *Command) diagonal() bool {
	return !c.horizontal() && !c.vertical()
}

type Board [1000][1000]int

func (b *Board) runCommand(command *Command, includeDiagonal bool) {
	if command.vertical() {
		var start int
		var end int
		if command.y1 < command.y2 {
			start = command.y1
			end = command.y2
		} else {
			start = command.y2
			end = command.y1
		}
		for i := start; i <= end; i++ {
			b[i][command.x1] += 1
		}
	}
	if command.horizontal() {
		var start int
		var end int
		if command.x1 < command.x2 {
			start = command.x1
			end = command.x2
		} else {
			start = command.x2
			end = command.x1
		}
		for i := start; i <= end; i++ {
			b[command.y1][i] += 1
		}
	}
	if includeDiagonal && command.diagonal() {
		x := command.x1
		y := command.y1
		var xDirection string
		var yDirection string
		if x < command.x2 {
			xDirection = "positive"
		} else {
			xDirection = "negative"
		}

		if y < command.y2 {
			yDirection = "positive"
		} else {
			yDirection = "negative"
		}

		if yDirection == "positive" {
			for y <= command.y2 {
				if xDirection == "positive" && x > command.x2 {
					break
				} else if xDirection == "negative" && x < command.x2 {
					break
				}

				b[y][x] += 1
				if xDirection == "positive" {
					x++
				} else if xDirection == "negative" {
					x--
				}
				y++
			}
		}
		if yDirection == "negative" {
			for y >= command.y2 {
				if xDirection == "positive" && x > command.x2 {
					break
				} else if xDirection == "negative" && x < command.x2 {
					break
				}

				b[y][x] += 1
				if xDirection == "positive" {
					x++
				} else if xDirection == "negative" {
					x--
				}
				y--
			}
		}
	}
}

func (b *Board) overlappingLinesCount() int {
	overlappingLinesCount := 0
	for _, line := range b {
		for _, pos := range line {
			if pos >= 2 {
				overlappingLinesCount += 1
			}
		}
	}
	return overlappingLinesCount
}

func part1(lines []string) int {
	board := Board{}
	for _, command := range NewCommands(lines) {
		board.runCommand(command, false)
	}
	return board.overlappingLinesCount()
}

func part2(lines []string) int {
	board := Board{}
	for _, command := range NewCommands(lines) {
		board.runCommand(command, true)
	}
	return board.overlappingLinesCount()
}

func main() {
	lines := utils.FileLinesToSlice("input.txt")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}
