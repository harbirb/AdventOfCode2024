package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// create 2d array, populating line by line
	var data [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		letters := strings.Split(line, "")
		data = append(data, letters)
	}
	rows, cols := len(data), len(data[0])
	mas := []string{"M", "A", "S"}
	dirs := [][]int{{0, 1}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, {1, 0}, {1, 1}}
	sum := 0

	// Visit each letter. If we hit an "X", start the search. VIsit one of 8 directions around the X.
	// calculate new indices in that direction, check if its in bounds
	// check letter at the new index matches next letter in sequence "MAS"
	// If new indices are out of bounds, or one letter doesn't match, break out and pick a new direction
	for i, row := range data {
		for j, letter := range row {
			if letter == "X" {
				for _, l := range dirs {
					for k := range 3 {
						newi := i + l[0]*(1+k)
						newj := j + l[1]*(1+k)
						if newi >= rows || newi < 0 || newj >= cols || newj < 0 {
							break
							// out of range
						}
						if data[newi][newj] != mas[k] {
							break
							// letter doesn't match
						} else if k == 2 {
							// all letters in XMAS match
							sum++
						}
					}
				}
			}
		}
	}
	fmt.Println(sum)

	// PART TWO STRATEGY
	// search and look for A's
	// "A" cannot be at any edges
	// check that both diagonals contain either MAS or SAM
	// +1,+1 is M and -1,-1 is S or other way around
	// +1,-1 is M and -1,+1 is S or other way around
	exes := 0
	for i, row := range data {
		for j, letter := range row {
			if letter == "A" && i > 0 && j > 0 && i < rows-1 && j < cols-1 {
				ul := data[i-1][j-1]
				ur := data[i-1][j+1]
				dl := data[i+1][j-1]
				dr := data[i+1][j+1]
				if (ul == "M" && dr == "S" || ul == "S" && dr == "M") && (ur == "M" && dl == "S" || ur == "S" && dl == "M") {
					exes++
				}
			}
		}
	}
	fmt.Println(exes)
}
