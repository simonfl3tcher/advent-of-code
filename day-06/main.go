package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type CommandPosition struct {
	x int
	y int
}

type Command struct {
	fullCommand string
	action      string
	start       CommandPosition
	end         CommandPosition
}

func stringToCommand(str string) Command {
	r, _ := regexp.Compile(`([a-zA-Z]+) (\d+,\d+) through (\d+,\d+)`)
	matches := r.FindStringSubmatch(str)

	startPosition := strings.Split(matches[2], ",")
	endPosition := strings.Split(matches[3], ",")

	startX, err := strconv.Atoi(startPosition[0])
	if err != nil {
		panic("cannot find position")
	}
	startY, err := strconv.Atoi(startPosition[1])
	if err != nil {
		panic("cannot find position")
	}
	endX, err := strconv.Atoi(endPosition[0])
	if err != nil {
		panic("cannot find position")
	}
	endY, err := strconv.Atoi(endPosition[1])
	if err != nil {
		panic("cannot find position")
	}

	return Command{
		fullCommand: matches[0],
		action:      matches[1],
		start:       CommandPosition{x: startX, y: startY},
		end:         CommandPosition{x: endX, y: endY},
	}
}

type LightsType [1000][1000]bool
type LightsBrightnessType [1000][1000]int

func runCommand(c Command, l *LightsType, lb *LightsBrightnessType) {
	for x := c.start.x; x <= c.end.x; x++ {
		for y := c.start.y; y <= c.end.y; y++ {
			switch c.action {
			case "toggle":
				l[x][y] = !l[x][y]
				lb[x][y] += 2
			case "on":
				l[x][y] = true
				lb[x][y] += 1
			case "off":
				l[x][y] = false
				if lb[x][y] > 0 {
					lb[x][y] -= 1
				}
			}
		}
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lights LightsType
	var lightBrightness LightsBrightnessType

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Do the thing with the command
		command := stringToCommand(scanner.Text())
		runCommand(command, &lights, &lightBrightness)
	}

	var lightCount int
	for i := 0; i < len(lights); i++ {
		for j := 0; j < len(lights[i]); j++ {
			if lights[i][j] {
				lightCount += 1
			}
		}
	}

	var lightBrightnessCount int
	for i := 0; i < len(lightBrightness); i++ {
		for j := 0; j < len(lightBrightness[i]); j++ {
			if lightBrightness[i][j] > 0 {
				lightBrightnessCount += lightBrightness[i][j]
			}
		}
	}

	fmt.Printf("Part 1: %d\n", lightCount)
	fmt.Printf("Part 2: %d\n", lightBrightnessCount)
}
