package main

import (
	"fmt"
	"math"
)

func main() {
	numberOfBulbs := 100
	root := math.Sqrt(float64(numberOfBulbs))
	for i := 1; i < int(root)+1; i++ {
		fmt.Println("Light bulb ", (i * i), " will be on")
	}

	fmt.Println("Total ", root, " light bulbs will be on in the end out of ", numberOfBulbs, " light bulbs")
}
