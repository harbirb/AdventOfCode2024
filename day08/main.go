package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var rows, cols int

func main() {
	// Strategy:
	// visit whole map, fill a hashmap with antenna locations for each freq
	// for each freq, get each pairwise antenna and determine antinode locations
	// count antinode if its in bounds and unique

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	m := [][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "")
		m = append(m, parts)
	}
	rows, cols = len(m), len(m[0])
	a := make(map[string][][]int)
	for i, row := range m {
		for j, freq := range row {
			if freq == "." {
				continue
			} else {
				a[freq] = append(a[freq], []int{i, j})
			}
		}
	}
	ans := make(map[string]int)
	for _, vs := range a {
		for i := range vs {
			for j := range vs {
				if j != i {
					dy, dx := vs[j][0]-vs[i][0], vs[j][1]-vs[i][1]
					// pt 2 first inode on top of antenna
					any, anx := vs[j][0], vs[j][1]
					for inBounds(any, anx) {
						an1_coords := strconv.Itoa(any) + "," + strconv.Itoa(anx)
						ans[an1_coords]++
						// PART2, keep adding antinodes at repeating distances
						any += dy
						anx += dx
					}
				}
			}
		}
	}
	fmt.Println(len(ans))
}

func inBounds(y, x int) bool {
	return y >= 0 && x >= 0 && y < rows && x < cols
}
