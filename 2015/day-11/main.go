package main

import (
	"fmt"
	"regexp"
	"strings"
)

type Password struct {
	password []byte
}

func (p *Password) includesThreeStraightLetters() bool {
	alphabet := alphabet()
	var cons []string
	for i := 0; i < len(alphabet)-2; i++ {
		cons = append(cons, fmt.Sprintf("%s%s%s", alphabet[i], alphabet[i+1], alphabet[i+2]))
	}

	for _, chunk := range cons {
		if strings.Contains(string(p.password), chunk) {
			return true
		}
	}

	return false
}

func (p *Password) doesNotContainIOL() bool {
	match, _ := regexp.MatchString("^[^iolIOL]+$", string(p.password))
	return match
}

func (p *Password) includesTwoPairs() bool {
	alphabet := alphabet()
	var chunks []string
	for i := 0; i < len(alphabet); i++ {
		chunks = append(chunks, fmt.Sprintf("%s%s", alphabet[i], alphabet[i]))
	}

	var containsCount int
	for _, chunk := range chunks {
		if strings.Contains(string(p.password), chunk) {
			if containsCount == 0 {
				containsCount++
			} else {
				return true
			}
		}
	}

	return false
}

func (p *Password) isValid() bool {
	return p.includesThreeStraightLetters() && p.doesNotContainIOL() && p.includesTwoPairs()
}

func (p *Password) increment() {
	for i := len(p.password) - 1; i >= 0; i-- {
		if p.password[i] != 'z' {
			p.password[i]++
			break
		}
		p.password[i] = 'a'
	}
}

func newPassword(str string) Password {
	return Password{password: []byte(str)}
}

func main() {
	p1 := newPassword("cqjxjnds")
	for {
		p1.increment()
		if p1.isValid() {
			fmt.Printf("Part 1: %s\n", string(p1.password))
			break
		}
	}

	p2 := newPassword("cqjxxyzz")
	for {
		p2.increment()
		if p2.isValid() {
			fmt.Printf("Part 2: %s\n", string(p2.password))
			break
		}
	}
}

func alphabet() []string {
	a := []string{}
	for i := 'a'; i <= 'z'; i++ {
		a = append(a, string(i))
	}
	return a
}
