package main

import (
	"advent-of-code-2015/utils"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

const defaultInt int64 = 1000 // default value because 0 is a valid value so we can't use this as a default

type Aunt struct {
	Id          int64 `compound:"id"`
	Children    int64 `compound:"children"`
	Cats        int64 `compound:"cats"`
	Samoyeds    int64 `compound:"samoyeds"`
	Pomeranians int64 `compound:"pomeranians"`
	Akitas      int64 `compound:"akitas"`
	Vizslas     int64 `compound:"vizslas"`
	Goldfish    int64 `compound:"goldfish"`
	Trees       int64 `compound:"trees"`
	Cars        int64 `compound:"cars"`
	Perfumes    int64 `compound:"perfumes"`
}

func main() {
	lines := utils.FileLinesToSlice("input.txt")
	var aunts []Aunt
	err := Unmarshal(lines, &aunts)
	if err != nil {
		panic(err)
	}

	tickerTapeSue := Aunt{
		Children:    3,
		Cats:        7,
		Samoyeds:    2,
		Pomeranians: 3,
		Akitas:      0,
		Vizslas:     0,
		Goldfish:    5,
		Trees:       3,
		Cars:        2,
		Perfumes:    1,
	}

	match1 := part1Match(aunts, tickerTapeSue)
	match2 := part2Match(aunts, tickerTapeSue)

	fmt.Printf("Part 1: %d\n", match1.Id)
	fmt.Printf("Part 2: %d\n", match2.Id)
}

func part1Match(aunts []Aunt, tickerTapeSue Aunt) Aunt {
	for _, sue := range aunts {
		var matchCount int
		if sue.Children == tickerTapeSue.Children {
			matchCount++
		}
		if sue.Cats == tickerTapeSue.Cats {
			matchCount++
		}
		if sue.Samoyeds == tickerTapeSue.Samoyeds {
			matchCount++
		}
		if sue.Pomeranians == tickerTapeSue.Pomeranians {
			matchCount++
		}
		if sue.Akitas == tickerTapeSue.Akitas {
			matchCount++
		}
		if sue.Vizslas == tickerTapeSue.Vizslas {
			matchCount++
		}
		if sue.Goldfish == tickerTapeSue.Goldfish {
			matchCount++
		}
		if sue.Trees == tickerTapeSue.Trees {
			matchCount++
		}
		if sue.Cars == tickerTapeSue.Cars {
			matchCount++
		}
		if sue.Perfumes == tickerTapeSue.Perfumes {
			matchCount++
		}

		if matchCount >= 3 {
			return sue
		}
	}
	return Aunt{}
}

func part2Match(aunts []Aunt, tickerTapeSue Aunt) Aunt {
	for _, sue := range aunts {
		var matchCount2 int
		if sue.Children == tickerTapeSue.Children {
			matchCount2++
		}
		if sue.Cats > tickerTapeSue.Cats && sue.Cats != defaultInt {
			matchCount2++
		}
		if sue.Samoyeds == tickerTapeSue.Samoyeds {
			matchCount2++
		}
		if sue.Pomeranians < tickerTapeSue.Pomeranians && sue.Pomeranians != defaultInt {
			matchCount2++
		}
		if sue.Akitas == tickerTapeSue.Akitas {
			matchCount2++
		}
		if sue.Vizslas == tickerTapeSue.Vizslas {
			matchCount2++
		}
		if sue.Goldfish < tickerTapeSue.Goldfish && sue.Goldfish != defaultInt {
			matchCount2++
		}
		if sue.Trees > tickerTapeSue.Trees && sue.Trees != defaultInt {
			matchCount2++
		}
		if sue.Cars == tickerTapeSue.Cars {
			matchCount2++
		}
		if sue.Perfumes == tickerTapeSue.Perfumes {
			matchCount2++
		}

		if matchCount2 >= 3 {
			return sue
		}
	}

	return Aunt{}
}

func Unmarshal(data []string, v interface{}) error {
	sliceValPtr := reflect.ValueOf(v)
	if sliceValPtr.Kind() != reflect.Ptr {
		return errors.New("must be a pointer to a slice of structs")
	}
	sliceVal := sliceValPtr.Elem()
	if sliceVal.Kind() != reflect.Slice {
		return errors.New("must be a pointer to a slice of structs")
	}
	structType := sliceVal.Type().Elem()
	if structType.Kind() != reflect.Struct {
		return errors.New("must be a pointer to a slice of structs")
	}

	for _, line := range data {
		newVal := reflect.New(structType).Elem()
		err := unmarshalOne(line, newVal)
		if err != nil {
			return err
		}
		sliceVal.Set(reflect.Append(sliceVal, newVal))
	}
	return nil
}

func unmarshalOne(row string, value reflect.Value) error {
	idRegex, _ := regexp.Compile(`Sue\s(\d+):`)

	vt := value.Type()
	for i := 0; i < value.NumField(); i++ {
		typeField := vt.Field(i)
		tagName := typeField.Tag.Get("compound")

		val := defaultInt
		if tagName == "id" {
			matches := idRegex.FindStringSubmatch(row)
			i, err := strconv.ParseInt(matches[1], 10, 64)
			if err != nil {
				return err
			}
			val = i
		} else {
			r, _ := regexp.Compile(`^Sue\s\d+:.*` + tagName + `:\s(\d+).*$`)
			matches := r.FindStringSubmatch(row)
			if len(matches) > 0 {
				i, err := strconv.ParseInt(matches[1], 10, 64)
				if err != nil {
					return err
				}
				val = i
			}
		}

		field := value.Field(i)
		switch field.Kind() {
		case reflect.Int64:
			field.SetInt(val)
		default:
			return fmt.Errorf("cannot handle field of kind %v", field.Kind())
		}
	}
	return nil
}
