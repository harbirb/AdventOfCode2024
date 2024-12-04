package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read the input file, strip the whitespace to obtain a slice of strings
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// create 2d array, populating line by line
	var data [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		var report []int
		for _, part := range parts {
			num, _ := strconv.Atoi(part)
			report = append(report, num)
		}
		data = append(data, report)
	}

	var numSafe = 0

	// go through each report, determine ordering
	for _, report := range data {
		var inc = report[0] < report[1]
		valid := true
		// invalidate report if ordering doesnt match or change is too great
		for j := range len(report) - 1 {
			if inc {
				if report[j] >= report[j+1] || report[j]+3 < report[j+1] {
					valid = false
				}
			} else {
				if report[j] <= report[j+1] || report[j] > report[j+1]+3 {
					valid = false
				}
			}
		}
		if valid {
			numSafe++
		} else {
			// check the report using dampener
			numSafe += damp(report)
		}
	}
	fmt.Printf("Number of safe reports: %d", numSafe)

}

// PART TWO
// dampener should check validity after removing one level at a time from an unsafe report
// return 1 if safe after dampening
func damp(currReport []int) int {
	for i := range currReport {
		report := append([]int{}, currReport[:i]...)
		report = append(report, currReport[i+1:]...)
		// fmt.Println(report)
		var inc = report[0] < report[1]
		valid := true
		for j := range len(report) - 1 {
			if inc {
				if report[j] >= report[j+1] || report[j]+3 < report[j+1] {
					valid = false
				}
			} else {
				if report[j] <= report[j+1] || report[j] > report[j+1]+3 {
					valid = false
				}
			}
		}
		if valid {
			return 1
		}
	}
	return 0
}
