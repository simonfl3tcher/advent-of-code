package utils

import (
	"bufio"
	"os"
	"strconv"
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

func Cons(values []string, num int) [][]string {
	result := make([][]string, 0)
	for i := 0; i < len(values); i++ {
		if i+num <= len(values) {
			result = append(result, values[i:i+num])
		}
	}
	return result
}

func IntSlice(values []string) []int {
	result := make([]int, 0, len(values))
	for _, value := range values {
		v, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		result = append(result, v)
	}
	return result
}

func IntCons(values []string, num int) [][]int {
	result := make([][]int, 0)
	for i := 0; i < len(values); i++ {
		if i+num <= len(values) {
			result = append(result, IntSlice(values[i:i+num]))
		}
	}
	return result
}

func Sum(values []int) int {
	result := 0
	for _, value := range values {
		result += value
	}
	return result
}

func Quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]

	less := []int{}
	greater := []int{}

	for _, v := range arr[1:] {
		if v <= pivot {
			less = append(less, v)
		}
		if v > pivot {
			greater = append(greater, v)
		}
	}

	response := []int{}
	response = append(response, Quicksort(less)...)
	response = append(response, pivot)
	response = append(response, Quicksort(greater)...)

	return response
}
