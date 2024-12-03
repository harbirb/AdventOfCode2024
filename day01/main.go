package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	// Read the input file, strip the whitespace to obtain a slice of strings
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	arr := strings.Fields(string(dat))

	// Divide into 2 slices of numbers, sort each column
	var firstArr, secondArr []int
	for i, val := range arr {
		ival, _ := strconv.ParseInt(val, 10, 0)
		if i%2 == 0 {
			secondArr = append(secondArr, int(ival))
		} else {
			firstArr = append(firstArr, int(ival))
		}
	}
	slices.Sort(firstArr)
	slices.Sort(secondArr)

	// Sum the distances between each point
	result := 0
	for i := range firstArr {
		dist := firstArr[i] - secondArr[i]
		if dist > 0 {
			result += dist
		} else {
			result -= dist
		}
	}
	fmt.Printf("Sum of distances: %d", result)

	// PART TWO: SIMILARITY SCORE

	// make a map of how many times each num appears in each array
	// iterate over the keys of the map. If the key appears in both, add key * val1 * val2
	m1 := make(map[int]int)
	m2 := make(map[int]int)
	for i := range firstArr {
		m1[firstArr[i]]++
		m2[secondArr[i]]++
	}

	simScore := 0
	for k, v := range m1 {
		if m2[k] > 0 {
			simScore += k * v * m2[k]
		}
	}
	fmt.Printf("\nSimilarity: %d", simScore)

}
