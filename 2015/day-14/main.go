package main

import (
	"advent-of-code/utils"
	"fmt"
	"regexp"
	"strconv"
)

type Raindeer struct {
	name     string
	speed    int
	flyTime  int
	restTime int

	currentlyFlying    bool
	currentPosition    int
	currentWaitingTime int
	bestPositionPoints int
}

func raindeerFromLine(str string) Raindeer {
	regex := regexp.MustCompile(`(\w+) can fly (\d+) km\/s for (\d+) seconds, but then must rest for (\d+) seconds.`)

	matches := regex.FindStringSubmatch(str)

	r := Raindeer{name: matches[1]}
	r.speed, _ = strconv.Atoi(matches[2])
	r.flyTime, _ = strconv.Atoi(matches[3])
	r.restTime, _ = strconv.Atoi(matches[4])
	r.currentWaitingTime = r.flyTime
	r.currentlyFlying = true

	return r
}

func (r *Raindeer) move() {
	if r.currentlyFlying {
		r.currentPosition += r.speed
	}
	r.currentWaitingTime -= 1
	if r.currentWaitingTime <= 0 {
		r.currentlyFlying = !r.currentlyFlying
		if r.currentlyFlying {
			r.currentWaitingTime = r.flyTime
		} else {
			r.currentWaitingTime = r.restTime
		}
	}
}

func main() {
	lines := utils.FileLinesToSlice("input.txt")

	var raindeer []*Raindeer
	for _, line := range lines {
		r := raindeerFromLine(line)
		raindeer = append(raindeer, &r)
	}

	for i := 0; i < 2503; i++ {
		for _, r := range raindeer {
			r.move()
		}

		var bestDistance int
		var bestDistanceSlice []int
		for _, r := range raindeer {
			bestDistanceSlice = append(bestDistanceSlice, r.currentPosition)
		}

		bestDistance = utils.MaxFromSlice(bestDistanceSlice)
		for _, r := range raindeer {
			if r.currentPosition == bestDistance {
				r.bestPositionPoints += 1
			}
		}
	}

	winnerPart1 := raindeer[0]
	for _, r := range raindeer {
		if r.currentPosition > winnerPart1.currentPosition {
			winnerPart1 = r
		}
	}

	winnerPart2 := raindeer[0]
	for _, r := range raindeer {
		if r.bestPositionPoints > winnerPart2.bestPositionPoints {
			winnerPart2 = r
		}
	}

	fmt.Printf("Part 1: %d\n", winnerPart1.currentPosition)
	fmt.Printf("Part 2: %d\n", winnerPart2.bestPositionPoints)
}
