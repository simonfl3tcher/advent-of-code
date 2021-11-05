package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type strChecker struct {
	str string
}

func (s *strChecker) threeVowels() bool {
	vowels := "aeiou"
	var vowelCount int
	for _, sv := range s.str {
		for _, v := range vowels {
			if sv == v {
				vowelCount++
			}
		}
	}
	return vowelCount >= 3
}

func (s *strChecker) oneLetterAppearsTwice() bool {
	var lastChar rune
	for _, v := range s.str {
		if lastChar == v {
			return true
		}

		lastChar = v
	}
	return false
}

func (s *strChecker) doesNotContainInvalidStrings() bool {
	invalidStrings := [4]string{"ab", "cd", "pq", "xy"}
	for _, v := range invalidStrings {
		if strings.Contains(s.str, v) {
			return false
		}
	}
	return true
}

func (s *strChecker) repeatingLetterWithGap() bool {
	for i := 0; i < len(s.str)-2; i++ {
		if s.str[i] == s.str[i+2] {
			return true
		}
	}

	return false
}

func (s *strChecker) twoPairs() bool {
	for i := 0; i < len(s.str)-2; i++ {
		if strings.Count(s.str, s.str[i:i+2]) >= 2 {
			return true
		}
	}

	return false
}

func (s *strChecker) valid(part int) bool {
	if part == 1 {
		return s.threeVowels() && s.oneLetterAppearsTwice() && s.doesNotContainInvalidStrings()
	}
	temp1 := s.twoPairs()
	temp2 := s.repeatingLetterWithGap()
	return temp1 && temp2
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var validCountPart1 int
	var validCountPart2 int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := strChecker{scanner.Text()}
		if str.valid(1) {
			validCountPart1++
		}
		if str.valid(2) {
			validCountPart2++
		}
	}

	fmt.Println(validCountPart1)
	fmt.Println(validCountPart2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
