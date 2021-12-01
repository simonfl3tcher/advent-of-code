package main

import (
	"advent-of-code/utils"
	"fmt"
)

func part1(lines []string) int {
	consecutives := utils.IntCons(lines, 2)
	var sum int
	for _, consecutive := range consecutives {
		if consecutive[0] < consecutive[1] {
			sum += 1
		}
	}
	return sum
}

func part2(lines []string) int {
	consecutives := utils.IntCons(lines, 3)
	var sum int
	for i, consecutive := range consecutives {
		if i+1 > len(consecutives)-1 {
			break
		}

		if utils.Sum(consecutive) < utils.Sum(consecutives[i+1]) {
			sum += 1
		}
	}
	return sum
}

func main() {
	lines := utils.FileLinesToSlice("input.txt")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}
