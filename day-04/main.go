package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	input := "ckczppom"

	fmt.Println(runner(input, 5))
	fmt.Println(runner(input, 6))
}

func runner(input string, zeroCount int) int {
	count := 0
	for {
		str := fmt.Sprintf("%s%d", input, count)
		data := md5.Sum([]byte(str))

		var zeroCountStr string
		for i := 0; i < zeroCount; i++ {
			zeroCountStr += "0"
		}

		result := hex.EncodeToString(data[:])
		if result[:zeroCount] == zeroCountStr {
			break
		}
		count++
	}

	return count
}
