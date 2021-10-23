package main

import (
	"advent-of-code-2015/utils"
	"fmt"
	"strconv"
)

func part1(lines []string) int {
	var strLenCounter int
	var unquoteStrLenCounter int
	for _, v := range lines {
		unquoteStr, err := strconv.Unquote(v)
		if err != nil {
			panic(err)
		}

		strLenCounter += len(v)
		unquoteStrLenCounter += len(unquoteStr)
	}

	return strLenCounter - unquoteStrLenCounter
}

func part2(lines []string) int {
	var strLenCounter int
	var quotedStrLiteralCount int
	fmt.Println(lines)
	for _, v := range lines {
		quoteStr := strconv.Quote(v)

		strLenCounter += len(v)
		quotedStrLiteralCount += len(quoteStr)
	}

	return quotedStrLiteralCount - strLenCounter
}

func main() {
	lines := utils.FileLinesToSlice("input.txt")

	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}
