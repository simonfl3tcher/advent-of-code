package main

import (
	"advent-of-code/utils"
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

type Result struct {
	minDistance int
	maxDistance int
}

func distancesKey(k1, k2 string) [2]string {
	key := []string{k1, k2}
	sort.Strings(key)
	return [2]string{key[0], key[1]}
}

func run(lines []string) Result {
	r, _ := regexp.Compile(`(\w+) to (\w+) = (\d+)`)

	var cities []string
	distances := make(map[[2]string]int)

	for _, line := range lines {
		from := r.FindStringSubmatch(line)[1]
		to := r.FindStringSubmatch(line)[2]
		cities = append(cities, from)
		cities = append(cities, to)

		distance := r.FindStringSubmatch(line)[3]
		i, err := strconv.Atoi(distance)
		if err != nil {
			panic(err)
		}
		distances[distancesKey(from, to)] = i
	}

	permutations := utils.Permutation(utils.Unique(cities))

	calculatedDistance := make([]int, 0)
	for _, permutation := range permutations {
		total := 0
		for i := 0; i < len(permutation)-1; i++ {
			total += distances[distancesKey(permutation[i], permutation[i+1])]
		}

		calculatedDistance = append(calculatedDistance, total)
	}

	return Result{minDistance: utils.MinFromSlice(calculatedDistance), maxDistance: utils.MaxFromSlice(calculatedDistance)}
}

func main() {
	lines := utils.FileLinesToSlice("input.txt")

	result := run(lines)

	fmt.Printf("Part 1: %d\n", result.minDistance)
	fmt.Printf("Part 2: %d\n", result.maxDistance)
}
