package main

import (
	"fmt"
	"strconv"
	"strings"
)

func expandNum(number int) {
	numStr := strings.Split(strconv.Itoa(number), "")

	for n := 0; n < len(numStr); n++ {
		zeroCount := len(numStr) - n - 1
		x := 0
		subNum := numStr[n]

		for x < zeroCount {
			subNum = subNum + "0"
			x++
		}

		fmt.Println(subNum)
	}
}

func main() {
	expandNum(1079)
	expandNum(1345679)
}
