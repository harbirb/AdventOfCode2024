package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	tm := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		numlist := make([]int, len(row))
		for i, num := range row {
			numlist[i], _ = strconv.Atoi(num)
		}
		tm = append(tm, numlist)
	}

	thm := make(map[string]int)

	for i := range tm {
		for j := range tm[i] {
			if tm[i][j] == 0 {
				// found trailhead
				// make new map to track peaks
				phm := make(map[string]int)
				dfs(tm, phm, i, j)
				coords := strconv.Itoa(i) + "," + strconv.Itoa(j)
				rating := 0
				for _, v := range phm {
					rating += v
				}
				thm[coords] = rating
			}
		}
	}
	sum := 0
	for _, v := range thm {
		sum += v
	}
	fmt.Println(sum)

}

// return num of peaks reachable from a trailhead
func dfs(tm [][]int, phm map[string]int, i, j int) {
	// recurse until 9 reached
	rows, cols := len(tm), len(tm[0])
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	if tm[i][j] == 9 {
		coords := strconv.Itoa(i) + "," + strconv.Itoa(j)
		phm[coords]++
	}
	for _, dir := range dirs {
		newi, newj := i+dir[0], j+dir[1]
		if newi >= 0 && newi < rows && newj >= 0 && newj < cols && tm[newi][newj] == tm[i][j]+1 {
			dfs(tm, phm, newi, newj)
		}
	}
}
