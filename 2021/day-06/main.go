package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func runner(bytes []byte, days int) int {
	fish := [9]int{}
	for _, c := range strings.Split(string(bytes), ",") {
		i, _ := strconv.Atoi(c)
		fish[i] += 1
	}

	for day := 0; day < days; day++ {
		newFish := [9]int{}
		for k, v := range fish {
			if k == 0 && v > 0 {
				newFish[6] += v
				newFish[8] += v
				continue
			}
			if k > 0 {
				newFish[k-1] += v
			}
		}
		fish = newFish
	}

	count := 0
	for _, v := range fish {
		count += v
	}
	return count
}

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", runner(bytes, 80))
	fmt.Printf("Part 2: %d\n", runner(bytes, 256))
}
