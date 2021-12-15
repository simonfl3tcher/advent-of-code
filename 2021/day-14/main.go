package main

import (
	"advent-of-code/utils"
	"fmt"
	"regexp"
	"strings"
)

const insertionRuleRegex = `^(\w+) -> (\w)$`

var insertionRegex *regexp.Regexp

func init() {
	insertionRegex = regexp.MustCompile(insertionRuleRegex)
}

func runner(template []string, rules map[string]string, counter int) int {
	currentPairs := map[string]int{}
	currentCounts := map[string]int{}

	for index, t := range template {
		if index+1 < len(template) {
			currentPairs[fmt.Sprintf("%s%s", t, template[index+1])] += 1
		}
		currentCounts[t] += 1
	}

	for counter > 0 {
		newMap := map[string]int{}
		for k, v := range currentPairs {
			newMap[k] = v
		}
		for k, v := range currentPairs {
			rule := rules[k]
			if rule == "" || v == 0 {
				continue
			}

			if newMap[k] > 0 {
				newMap[k] -= v
			}

			keys := strings.Split(k, "")
			for index, g := range keys {
				if index == 0 {
					newMap[fmt.Sprintf("%s%s", g, rule)] += v
				} else {
					newMap[fmt.Sprintf("%s%s", rule, g)] += v
				}
			}
			currentCounts[rule] += v
		}
		currentPairs = newMap
		counter--
	}

	lowest := -1
	highest := -1
	for _, v := range currentCounts {
		if v < lowest || lowest < 0 {
			lowest = v
		}
		if v > highest || highest < 0 {
			highest = v
		}
	}

	return highest - lowest
}

func main() {
	lines := utils.FileLinesToSlice("input.txt")
	template := strings.Split(lines[0], "")
	pairInsertionRules := lines[2:]

	rules := map[string]string{}
	for _, v := range pairInsertionRules {
		matches := insertionRegex.FindStringSubmatch(v)
		rules[matches[1]] = matches[2]
	}

	fmt.Printf("Part 1: %d\n", runner(template, rules, 10))
	fmt.Printf("Part 2: %d\n", runner(template, rules, 40))
}
