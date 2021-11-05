package main

import (
	"advent-of-code/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type boxDimensions struct {
	length, width, height int
}

func (b *boxDimensions) CalculatePaper() int {
	s := []int{b.length * b.width, b.width * b.height, b.height * b.length}
	return 2*s[0] + 2*s[1] + 2*s[2] + utils.MinFromSlice(s)
}

func (b *boxDimensions) CalculateRibbon() int {
	perimeter1 := b.length + b.length + b.width + b.width
	perimeter2 := b.length + b.length + b.height + b.height
	perimeter3 := b.height + b.height + b.width + b.width

	wrap := utils.MinFromSlice([]int{perimeter1, perimeter2, perimeter3})
	ribbon := b.length * b.width * b.height

	return wrap + ribbon
}

func newBoxDimensions(str string) boxDimensions {
	s := strings.Split(str, "x")
	l, _ := strconv.Atoi(s[0])
	w, _ := strconv.Atoi(s[1])
	h, _ := strconv.Atoi(s[2])

	return boxDimensions{l, w, h}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var totalPaper int
	var totalRibbon int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		box := newBoxDimensions(scanner.Text())
		totalPaper += box.CalculatePaper()
		totalRibbon += box.CalculateRibbon()
	}

	fmt.Printf("Total paper: %d\n", totalPaper)
	fmt.Printf("Total ribbon: %d\n", totalRibbon)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
