package main

import (
	"advent-of-code/utils"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func part1(nums []int) int {
	nums = utils.Quicksort(nums)
	median := nums[len(nums)/2]

	f := 0
	for _, num := range nums {
		if num < median {
			f += median - num
		} else {
			f += num - median
		}
	}
	return f
}

func part2(nums []int) int {
	count := 0.0
	for _, num := range nums {
		count += float64(num)
	}
	mean := int(math.Floor(count / float64(len(nums))))

	counter := 0
	for _, num := range nums {
		var distance int
		if num < mean {
			distance += mean - num
		} else {
			distance += num - mean
		}

		fuel := 0
		for i := 1; i <= distance; i++ {
			fuel += i
		}
		counter += fuel
	}
	return counter
}

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	nums := []int{}
	for _, c := range strings.Split(string(bytes), ",") {
		i, _ := strconv.Atoi(c)
		nums = append(nums, i)
	}

	fmt.Printf("Part 1: %d\n", part1(nums))
	fmt.Printf("Part 2: %d\n", part2(nums))
}
