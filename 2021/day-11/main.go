package main

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
)

type GridElement struct {
	Value   int
	Flashed bool
	X       int
	Y       int
}

type GridLine []*GridElement
type Grid []*GridLine

func (g *Grid) neighbours(y, x int) []*GridElement {
	positions := [][2]int{
		{y - 1, x - 1},
		{y - 1, x},
		{y - 1, x + 1},
		{y, x - 1},
		{y, x + 1},
		{y + 1, x - 1},
		{y + 1, x},
		{y + 1, x + 1},
	}

	neighbours := []*GridElement{}
	for _, v := range positions {
		if v[0] < 0 || v[1] < 0 || v[1] >= len(*(*g)[0]) || v[0] >= len(*g) {
			continue
		}
		neighbours = append(neighbours, (*(*g)[v[0]])[v[1]])
	}
	return neighbours
}

func (g *Grid) performIterations(iterations int, flashes int) int {
	if iterations < 1 {
		return flashes
	}

	for yIndex, gr := range *g {
		for xIndex, gc := range *gr {
			gc.performAction(*g, yIndex, xIndex)
		}
	}

	for _, gr := range *g {
		for _, gc := range *gr {
			if gc.Flashed {
				flashes += 1
			}
			gc.Flashed = false
		}
	}

	return g.performIterations(iterations-1, flashes)
}

func (ge *GridElement) performAction(grid Grid, yIndex, xIndex int) {
	if ge.Flashed {
		return
	}

	if ge.Value <= 9 {
		ge.Value++
	}

	if ge.Value > 9 {
		ge.Value = 0
		ge.Flashed = true

		for _, n := range grid.neighbours(yIndex, xIndex) {
			n.performAction(grid, n.Y, n.X)
		}
	}
}

func (g *Grid) allOctopusFlash() bool {
	t := true
	for _, a := range *g {
		for _, b := range *a {
			if b.Value != 0 {
				t = false
			}
		}
	}
	return t
}

func newGrid(lines []string) Grid {
	grid := Grid{}
	for yIndex, line := range lines {
		gridLine := GridLine{}
		for xIndex, i := range line {
			value, _ := strconv.Atoi(string(i))
			gridLine = append(gridLine, &GridElement{Value: value, X: xIndex, Y: yIndex})
		}
		grid = append(grid, &gridLine)
	}
	return grid
}

func part1(lines []string) int {
	grid := newGrid(lines)
	return grid.performIterations(100, 0)
}

func part2(lines []string) int {
	grid := newGrid(lines)
	i := 0
	for {
		i++
		grid.performIterations(1, 0)
		if grid.allOctopusFlash() {
			break
		}
	}
	return i
}

func main() {
	lines := utils.FileLinesToSlice("input.txt")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}
