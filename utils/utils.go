package utils

import (
	"bufio"
	"os"
)

func MinFromSlice(slice []int) int {
	smallestNumber := slice[0]
	for _, element := range slice {
		if element < smallestNumber {
			smallestNumber = element
		}
	}
	return smallestNumber
}

func MaxFromSlice(slice []int) int {
	largestNumber := slice[0]
	for _, element := range slice {
		if element > largestNumber {
			largestNumber = element
		}
	}
	return largestNumber
}

func FileLinesToSlice(str string) []string {
	file, err := os.Open(str)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var result []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result
}

func Unique(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func Permutation(values []string) (result [][]string) {
	if len(values) == 1 {
		result = append(result, values)
		return
	}

	for i, current := range values {
		others := make([]string, 0, len(values)-1)
		others = append(others, values[:i]...)
		others = append(others, values[i+1:]...)
		for _, route := range Permutation(others) {
			result = append(result, append(route, current))
		}
	}
	return
}
