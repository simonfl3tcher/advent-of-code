package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", numSteps(string(bytes), 1))
	fmt.Printf("Part 2: %d\n", numSteps(string(bytes), 2))
}

func numSteps(str string, numberOfSantas int) int {
	steps := make(map[[2]int]int)
	santaACurrentPosition := [2]int{0, 0}
	santaBCurrentPosition := [2]int{0, 0}

	steps[santaACurrentPosition] = 1
	if numberOfSantas == 2 {
		steps[santaBCurrentPosition] = 1
	}

	for i, char := range str {
		var currentStep *[2]int
		if numberOfSantas == 2 && i%2 == 0 {
			currentStep = &santaBCurrentPosition
		} else {
			currentStep = &santaACurrentPosition
		}

		switch char {
		case '^':
			*currentStep = [2]int{currentStep[0] + 1, currentStep[1]}
			steps[*currentStep] += 1
		case '>':
			*currentStep = [2]int{currentStep[0], currentStep[1] + 1}
			steps[*currentStep] += 1
		case 'v':
			*currentStep = [2]int{currentStep[0] - 1, currentStep[1]}
			steps[*currentStep] += 1
		case '<':
			*currentStep = [2]int{currentStep[0], currentStep[1] - 1}
			steps[*currentStep] += 1
		}
	}

	return len(steps)
}
