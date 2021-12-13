package main

import (
	"advent-of-code/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const foldRegxExp = `^fold along (x|y)=(\d+)$`

var foldRegex *regexp.Regexp

func init() {
	foldRegex = regexp.MustCompile(foldRegxExp)
}

type GridElement struct {
	value string
	x     int
	y     int
}

type Fold struct {
	axis   string
	number int
}

func GridAndFoldsFromLines(lines []string) ([][]string, []Fold) {
	parseFolds := false
	folds := []Fold{}
	gridElements := []GridElement{}
	highestX := 0
	highestY := 0
	for _, line := range lines {
		if line == "" {
			parseFolds = true
			continue
		}

		if !parseFolds {
			splits := strings.Split(line, ",")
			x, _ := strconv.Atoi(splits[0])
			y, _ := strconv.Atoi(splits[1])
			ge := GridElement{x: x, y: y, value: "#"}
			gridElements = append(gridElements, ge)
			if x > highestX {
				highestX = x
			}
			if y > highestY {
				highestY = y
			}
		} else {
			matches := foldRegex.FindStringSubmatch(line)
			i, _ := strconv.Atoi(matches[2])
			fold := Fold{axis: matches[1], number: i}
			folds = append(folds, fold)
		}
	}

	grid := make([][]string, highestY+1)
	for i := range grid {
		grid[i] = make([]string, highestX+1)
	}

	for y := 0; y <= highestY; y++ {
		for x := 0; x <= highestX; x++ {
			grid[y][x] = "."
		}
	}

	for _, ge := range gridElements {
		grid[ge.y][ge.x] = "#"
	}

	return grid, folds
}

func runFolds(grid *[][]string, folds []Fold) {
	for _, fold := range folds {
		if fold.axis == "x" {
			for gridI, gy := range *grid {
				iteration := 1
				for i := fold.number + 1; i < len(gy); i++ {
					if (*grid)[gridI][i] == "#" || (*grid)[gridI][fold.number-iteration] == "#" {
						(*grid)[gridI][fold.number-iteration] = "#"
					}
					iteration++
				}
				(*grid)[gridI] = (*grid)[gridI][:fold.number]
			}
		} else if fold.axis == "y" {
			iteration := 1
			for i := fold.number + 1; i < len(*grid); i++ {
				for charI, char := range (*grid)[i] {
					if char == "#" || (*grid)[fold.number-iteration][charI] == "#" {
						(*grid)[fold.number-iteration][charI] = "#"
					}
				}
				iteration++
			}

			*grid = (*grid)[:fold.number]
		}
	}
}

func part1(lines []string) {
	grid, folds := GridAndFoldsFromLines(lines)
	runFolds(&grid, []Fold{folds[0]})

	acc := 0
	for _, y := range grid {
		for _, x := range y {
			if x == "#" {
				acc++
			}
		}
	}
	fmt.Printf("Part 1: %d\n", acc)
}

func part2(lines []string) {
	grid, folds := GridAndFoldsFromLines(lines)
	runFolds(&grid, folds)

	fmt.Println("Part 2:")
	for _, f := range grid {
		fmt.Println(f)
	}
}

func main() {
	lines := utils.FileLinesToSlice("input.txt")
	part1(lines)
	part2(lines)
}
