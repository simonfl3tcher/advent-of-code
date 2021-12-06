package main

import (
	"advent-of-code/utils"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
	values [5][]string
	winner bool
}

const moveRegex = `\s?([0-9]*)\s+([0-9]*)\s+([0-9]*)\s+([0-9]*)\s+([0-9]*)`

var regx *regexp.Regexp

func init() {
	regx = regexp.MustCompile(moveRegex)
}

func (c *Card) checkCall(str string) {
	for lineIndex, line := range c.values {
		for charIndex, char := range line {
			if char == str {
				c.values[lineIndex][charIndex] = "X"
			}
		}
	}
}

func (c *Card) bingo() bool {
	bingo := false
	columns := [5][5]string{}
	for lineIndex, line := range c.values {
		if reflect.DeepEqual(line, []string{"X", "X", "X", "X", "X"}) {
			bingo = true
			break
		}
		for charIndex, char := range line {
			if char == "X" {
				columns[charIndex][lineIndex] = "X"
			}
		}
	}

	if bingo {
		return bingo
	}

	for _, column := range columns {
		if reflect.DeepEqual(column, [5]string{"X", "X", "X", "X", "X"}) {
			bingo = true
			break
		}
	}

	return bingo
}

func (c *Card) calculateScore(winningNumber string) int {
	var nums int
	for _, line := range c.values {
		for _, char := range line {
			if char != "X" {
				v, _ := strconv.Atoi(char)
				nums += v
			}
		}
	}
	wn, _ := strconv.Atoi(winningNumber)
	return nums * wn
}

type Winner struct {
	card          *Card
	winningNumber string
}

func part1(i *Input) int {
	var winner Winner

out:
	for _, call := range i.calls {
		for _, bc := range i.boards {
			bc.checkCall(call)
		}
		for _, bc := range i.boards {
			if bc.bingo() {
				fmt.Println("in here")
				winner.card = bc
				winner.winningNumber = call
				break out
			}
		}
	}

	return winner.card.calculateScore(winner.winningNumber)
}

func part2(i *Input) int {
	var winners []Winner

	for _, call := range i.calls {
		for _, bc := range i.boards {
			if !bc.winner {
				bc.checkCall(call)
			}
		}
		for _, bc := range i.boards {
			if !bc.winner {
				if bc.bingo() {
					bc.winner = true
					winners = append(winners, Winner{card: bc, winningNumber: call})
				}
			}
		}
	}

	lastWinner := winners[len(winners)-1]
	return lastWinner.card.calculateScore(lastWinner.winningNumber)
}

func main() {
	lines := utils.FileLinesToSlice("input.txt")
	input := getInput(lines)

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

type Input struct {
	calls  []string
	boards []*Card
}

func getInput(lines []string) *Input {
	calls := strings.Split(lines[0], ",")
	f := lines[1:]
	cards := [][]string{}
	i := 0
	for {
		if i >= len(f) {
			break
		}
		cards = append(cards, f[i+1:i+6])
		i += 6
	}
	return &Input{calls: calls, boards: getBoards(cards)}
}

func getBoards(cards [][]string) []*Card {
	boardCards := []*Card{}
	for _, card := range cards {
		ca := Card{}
		for index, line := range card {
			matches := regx.FindStringSubmatch(line)
			ca.values[index] = []string{matches[1], matches[2], matches[3], matches[4], matches[5]}
		}
		boardCards = append(boardCards, &ca)
	}
	return boardCards
}
