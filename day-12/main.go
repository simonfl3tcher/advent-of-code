package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	var input interface{}
	err = json.Unmarshal(file, &input)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %g\n", count(input, true))
	fmt.Printf("Part 2: %g\n", count(input, false))
}

func count(input interface{}, countRed bool) (total float64) {
	switch value := input.(type) {
	case float64:
		return value
	case string:
		return 0
	case []interface{}:
		for _, child := range value {
			total += count(child, countRed)
		}
		return
	case map[string]interface{}:
		if !countRed {
			for _, child := range value {
				if text, _ := child.(string); text == "red" {
					return 0
				}
			}
		}
		// check for red here!
		for _, child := range value {
			total += count(child, countRed)
		}
		return
	default:
		panic(fmt.Sprintf("unknown type: %T", value))
	}
}
