package main

import (
	"advent-of-code/utils"
	"fmt"
	"regexp"
	"strconv"
)

type Ingredient struct {
	name       string
	capacity   int64
	durability int64
	flavor     int64
	texture    int64
	calories   int64
}

var (
	regex = regexp.MustCompile(`^(\w+):\s\w+\s(-?\d+),\s\w+\s(-?\d+),\s\w+\s(-?\d+),\s\w+\s(-?\d+),\s\w+\s(-?\d+)$`)
)

func main() {
	input := utils.FileLinesToSlice("input.txt")

	var ingredients []Ingredient
	for _, line := range input {
		ingredients = append(ingredients, parseIngredient(line))
	}

	// Part 1
	fmt.Printf("Part 1: %v", findBestScore(ingredients, 100))
}

func findBestScore(ingredients []Ingredient, max int64) int64 {
	return 1
}

func parseIngredient(line string) Ingredient {
	matches := regex.FindStringSubmatch(line)
	var i Ingredient
	i.name = matches[1]
	i.capacity, _ = strconv.ParseInt(matches[2], 10, 64)
	i.durability, _ = strconv.ParseInt(matches[3], 10, 64)
	i.flavor, _ = strconv.ParseInt(matches[4], 10, 64)
	i.texture, _ = strconv.ParseInt(matches[5], 10, 64)
	i.calories, _ = strconv.ParseInt(matches[6], 10, 64)

	return i
}
