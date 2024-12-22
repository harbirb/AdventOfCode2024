package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var track [][]string
var pm = make(map[string]int)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		track = append(track, strings.Split(scanner.Text(), ""))
	}
	// rows, cols := len(track), len(track[0])

	var i, j int
	for y, row := range track {
		fmt.Println(row)
		for x := range row {
			if row[x] == "S" {
				i, j = y, x
				track[i][j] = "."
			}
		}
	}
	dfs(i, j, 0)
	cheats := 0
	cm := make(map[string]int)

	// from each path block, visit all other path blocks greater than it
	// if accessible by cheating (manhattan dist of 20 or less), add to cheatmap

	for k, v := range pm {
		for k2, v2 := range pm {
			if v2 > v {
				dist := manhattan(k, k2)
				if dist <= 20 {
					cm[k+","+k2] = v2 - v - dist
				}
			}
		}
	}

	// cmm := make(map[int]int)
	for _, v := range cm {
		if v >= 100 {
			// cmm[v]++
			cheats++
		}
	}
	// fmt.Println(cmm)
	fmt.Println(cheats)
}

func dfs(i, j, ps int) {
	if track[i][j] == "#" {
		return
	}
	if track[i][j] == "." || track[i][j] == "E" {
		track[i][j] = strconv.Itoa(ps)
		y, x := strconv.Itoa(i), strconv.Itoa(j)
		pm[y+","+x] = ps
		// fmt.Println(track[i][j])
		dfs(i+1, j, ps+1)
		dfs(i-1, j, ps+1)
		dfs(i, j+1, ps+1)
		dfs(i, j-1, ps+1)
		return
	}
}

func manhattan(c1, c2 string) int {
	coords1 := strings.Split(c1, ",")
	coords2 := strings.Split(c2, ",")
	y1, _ := strconv.Atoi(coords1[0])
	x1, _ := strconv.Atoi(coords1[1])
	y2, _ := strconv.Atoi(coords2[0])
	x2, _ := strconv.Atoi(coords2[1])
	dy := y1 - y2
	dx := x1 - x2
	if dy < 0 {
		dy *= -1
	}
	if dx < 0 {
		dx *= -1
	}
	return dy + dx
}
