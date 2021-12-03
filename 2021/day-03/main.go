package main

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"
)

type Bit struct {
	zeroCount int
	oneCount  int
}

func (b *Bit) mostCommon() int {
	if b.oneCount > b.zeroCount {
		return 1
	}
	return 0
}

func (b *Bit) leastCommon() int {
	if b.oneCount < b.zeroCount {
		return 1
	}
	return 0
}

func (b *Bit) equal() bool {
	return b.oneCount == b.zeroCount
}

type reporter [12]Bit
type RatingType int

const (
	oxygen RatingType = iota
	co2
)

func (r *reporter) calulateMostCommonPositions(str string) {
	chars := strings.Split(str, "")

	for i, c := range chars {
		if c == "0" {
			r[i].zeroCount += 1
		} else {
			r[i].oneCount += 1
		}
	}
}

func (r *reporter) calculatePowerConsumption(lines []string) int64 {
	for _, l := range lines {
		r.calulateMostCommonPositions(l)
	}

	var mstCommon string
	var lstCommon string
	for _, t := range r {
		mstCommon = fmt.Sprintf("%s%d", mstCommon, t.mostCommon())
		lstCommon = fmt.Sprintf("%s%d", lstCommon, t.leastCommon())
	}

	x, _ := strconv.ParseInt(mstCommon, 2, 64)
	y, _ := strconv.ParseInt(lstCommon, 2, 64)

	return x * y
}

func (r *reporter) calculateLifeSupportRating(lines []string) int64 {
	for _, l := range lines {
		r.calulateMostCommonPositions(l)
	}

	oxygenGeneratorRating := generatorRating(r, 0, lines, oxygen)
	co2ScrubberRating := generatorRating(r, 0, lines, co2)

	x, _ := strconv.ParseInt(oxygenGeneratorRating, 2, 64)
	y, _ := strconv.ParseInt(co2ScrubberRating, 2, 64)

	return x * y
}

type charCandidates []string

func generatorRating(r *reporter, index int, candidates charCandidates, rt RatingType) string {
	if len(candidates) < 1 {
		panic("cannot have candidates less than 1")
	}

	if len(candidates) == 1 || index > len(r)-1 {
		return candidates[0]
	}

	var charToCheck string
	if rt == oxygen {
		charToCheck = fmt.Sprintf("%d", r[index].mostCommon())
	} else {
		charToCheck = fmt.Sprintf("%d", r[index].leastCommon())
	}

	newCandidates := charCandidates{}
	for ind, k := range candidates {
		if index > len(k)-1 {
			break
		}

		if r[index].equal() && rt == oxygen {
			charToCheck = "1"
		} else if r[index].equal() && rt == co2 {
			charToCheck = "0"
		}

		if string(k[index]) != charToCheck {
			continue
		}
		newCandidates = append(newCandidates, candidates[ind])
	}

	reporter := reporter{}
	for _, l := range newCandidates {
		reporter.calulateMostCommonPositions(l)
	}

	index++
	return generatorRating(&reporter, index, newCandidates, rt)
}

func part1(lines []string) int64 {
	reporter := reporter{}
	return reporter.calculatePowerConsumption(lines)
}

func part2(lines []string) int64 {
	reporter := reporter{}
	return reporter.calculateLifeSupportRating(lines)
}

func main() {
	lines := utils.FileLinesToSlice("input.txt")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}
