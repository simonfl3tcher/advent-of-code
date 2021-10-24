package main

import (
	"fmt"
)

func runner(tail []rune, head rune, currentBuiltString string, currentCountOfChar int) string {
	if len(tail) == 0 {
		if currentCountOfChar == 0 {
			return currentBuiltString
		} else {
			return currentBuiltString + fmt.Sprintf("%d%s", currentCountOfChar+1, string(head))
		}
	}

	if head == tail[0] {
		return runner(tail[1:], head, currentBuiltString, currentCountOfChar+1)
	}

	fmt.Println("1231")
	f := fmt.Sprintf("%d%s", currentCountOfChar+1, string(tail[0]))
	return runner(tail[1:], tail[0], currentBuiltString+f, currentCountOfChar)
}

func run(input string) string {
	tail := []rune(input)
	return runner(tail[1:], tail[0], "", 0)
}

func main() {
	input := "21"
	result := run(input)

	fmt.Printf("Part 1: %s\n", result)
}
