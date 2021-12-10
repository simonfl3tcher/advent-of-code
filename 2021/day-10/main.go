package main

import (
	"advent-of-code/utils"
	"errors"
	"fmt"
)

func findIlegalChar(text string) (rune, error) {
	var stack []rune
	for _, c := range text {
		if c == '(' || c == '[' || c == '{' || c == '<' {
			stack = append(stack, c)
		} else if c == ')' || c == ']' || c == '}' || c == '>' {
			if braceMap(c) == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				return c, errors.New("found an illegal brace")
			}
		}
	}
	return '0', nil
}

func braceMap(char rune) rune {
	v := map[rune]rune{')': '(', ']': '[', '}': '{', '>': '<', '(': ')', '[': ']', '{': '}', '<': '>'}
	return v[char]
}

func findCompetionCombinationForLine(text string) []rune {
	var stack []rune
	for _, c := range text {
		if c == '(' || c == '[' || c == '{' || c == '<' {
			stack = append(stack, c)
		} else if c == ')' || c == ']' || c == '}' || c == '>' {
			if braceMap(c) == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			}
		}
	}

	g := []rune{}
	if len(stack) > 0 {
		for i := len(stack) - 1; i >= 0; i-- {
			g = append(g, braceMap(stack[i]))
		}
	}

	return g
}

func part1(lines []string) int {
	charScores := map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
	ilegalChars := map[rune]int{}
	for _, line := range lines {
		ilegal, err := findIlegalChar(line)
		if err != nil {
			ilegalChars[ilegal]++
		}
	}
	acc := 0
	for k, v := range ilegalChars {
		acc += v * charScores[k]
	}
	return acc
}

func part2(lines []string) int {
	charScores := map[rune]int{')': 1, ']': 2, '}': 3, '>': 4}
	incompleteLines := []string{}
	for _, line := range lines {
		_, err := findIlegalChar(line)
		if err == nil {
			incompleteLines = append(incompleteLines, line)
		}
	}

	completions := [][]rune{}
	for _, lin := range incompleteLines {
		completions = append(completions, findCompetionCombinationForLine(lin))
	}

	scores := []int{}
	for _, c := range completions {
		acc := 0
		for _, char := range c {
			acc = (acc * 5) + charScores[char]
		}
		scores = append(scores, acc)
	}

	middle := int(len(scores) / 2)
	scores = utils.Quicksort(scores)
	return scores[middle]
}

func main() {
	lines := utils.FileLinesToSlice("input.txt")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}
