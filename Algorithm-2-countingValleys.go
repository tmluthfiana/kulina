package main

import (
	"fmt"
)

func countingValleys(n int, s string) int {
        isSeaLevel := true
        altitude := 0
        valleys := 0

        for i:=0; i < n; i++ {
            if string(s)[i] == 'D' {
                altitude--
            } else {
                altitude++
            }

            if isSeaLevel && altitude < 0 {
                valleys++
            }

            isSeaLevel = altitude == 0
        }
        return valleys

    }

func main() {
   n := 10
   s := "UDDDUDUUDU"
   fmt.Println(countingValleys(n, s))
}