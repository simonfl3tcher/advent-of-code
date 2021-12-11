package main

import (
	"advent-of-code/utils"
	"fmt"
	"regexp"
	"strconv"
)

const lineRegex = `(Alice|Bob|Carol|David|Eric|Frank|George|Mallory).*(gain|lose)\s(\d+).*(Alice|Bob|Carol|David|Eric|Frank|George|Mallory)`

var regx *regexp.Regexp

var names = []string{"Alice", "Bob", "Carol", "David", "Eric", "Frank", "George", "Mallory"}

func init() {
	regx = regexp.MustCompile(lineRegex)
}

func peopleMapping(lines []string) map[string]map[string]int {
	people := map[string]map[string]int{}
	for _, line := range lines {
		matches := regx.FindStringSubmatch(line)
		if people[matches[1]] == nil {
			people[matches[1]] = map[string]int{}
		}
		num := ""
		if matches[2] == "gain" {
			num = fmt.Sprintf("%s", matches[3])
		} else {
			num = fmt.Sprintf("-%s", matches[3])
		}

		x, _ := strconv.Atoi(num)
		people[matches[1]][matches[4]] = x
	}

	return people
}

func findHappinessScore(people map[string]map[string]int, perms [][]string) int {
	counts := []int{}
	for _, c := range perms {
		cons := utils.Cons(c, 2)
		add := []string{c[len(c)-1], c[0]}
		cons = append(cons, add)

		acc := 0
		for _, f := range cons {
			acc += people[f[0]][f[1]] + people[f[1]][f[0]]
		}
		counts = append(counts, acc)
	}
	counts = utils.Quicksort(counts)
	return counts[len(counts)-1]
}

func part1(lines []string) int {
	people := peopleMapping(lines)
	return findHappinessScore(people, utils.Permutation(names))
}

func part2(lines []string) int {
	people := peopleMapping(lines)
	namesWithMe := names
	namesWithMe = append(namesWithMe, "Me")
	people["Me"] = map[string]int{}
	for _, h := range namesWithMe {
		people[h]["Me"] = 0
		people["Me"][h] = 0
	}

	return findHappinessScore(people, utils.Permutation(namesWithMe))
}

func main() {
	lines := utils.FileLinesToSlice("input.txt")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}
