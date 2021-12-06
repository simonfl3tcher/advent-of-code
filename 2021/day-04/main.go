package main

import (
	"advent-of-code/utils"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Card [5][]string

const moveRegex = `\s?([0-9]*)\s+([0-9]*)\s+([0-9]*)\s+([0-9]*)\s+([0-9]*)`

var regx *regexp.Regexp

func init() {
	regx = regexp.MustCompile(moveRegex)
}

func (c *Card) checkCall(str string) {
	for lineIndex, line := range c {
		for charIndex, char := range line {
			if char == str {
				c[lineIndex][charIndex] = "X"
			}
		}
	}
}

func (c *Card) bingo() bool {
	bingo := false
	columns := [5][5]string{}
	for lineIndex, line := range c {
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
	for _, line := range c {
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

func part1(cards [][]string, calls []string) int {
	boardCards := []*Card{}
	for _, card := range cards {
		ca := Card{}
		for index, line := range card {
			matches := regx.FindStringSubmatch(line)
			ca[index] = []string{matches[1], matches[2], matches[3], matches[4], matches[5]}
		}
		boardCards = append(boardCards, &ca)
	}

	var winningCard *Card
	var winningNumber string
out:
	for _, call := range calls {
		for _, bc := range boardCards {
			bc.checkCall(call)
		}
		for _, bc := range boardCards {
			if bc.bingo() {
				fmt.Println("in here")

				winningCard = bc
				winningNumber = call
				break out
			}
		}
	}

	return winningCard.calculateScore(winningNumber)
}

func part2(cards [][]string, calls []string) int {
	return 1
}

func main() {
	lines := utils.FileLinesToSlice("input.txt")
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

	fmt.Printf("Part 1: %d\n", part1(cards, calls))
	fmt.Printf("Part 2: %d\n", part2(cards, calls))
}
