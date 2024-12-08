package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime/trace"
	"strconv"
	"strings"
)

var rows, cols int

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	trace.Start(f)
	defer trace.Stop()
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	labmap := [][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "")
		labmap = append(labmap, row)
	}

	rows, cols = len(labmap), len(labmap[0])
	// find guard's position gi, gj
	var initial_i, initial_j int
	for i := range labmap {
		for j := range labmap[i] {
			if labmap[i][j] == "^" {
				initial_i, initial_j = i, j
			}
		}
	}
	gi, gj := initial_i, initial_j
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	diridx := 0
	dir := dirs[0]
	visited := make(map[string]int)
	for inBounds(gi, gj) {
		nexti, nextj := gi+dir[0], gj+dir[1]
		if inBounds(nexti, nextj) {
			// obstacle ahead
			if labmap[nexti][nextj] == "#" {
				diridx++
				diridx %= 4
				dir = dirs[diridx]
				continue
			} else {
				// empty space ahead
				// simulate a fake obstacle ahead
				coords := strconv.Itoa(nexti) + "," + strconv.Itoa(nextj)
				if visited[coords] == 0 {
					labmap[nexti][nextj] = "O"
					if isLoop(labmap, initial_i, initial_j) {
						visited[coords]++
					}
					labmap[nexti][nextj] = "."
				}
				gi += dir[0]
				gj += dir[1]
			}
		} else {
			gi += dir[0]
			gj += dir[1]
		}
	}

	// when guard is off the map, sum and return number of X's
	sum := 0
	for i := range labmap {
		for j := range labmap[i] {
			if labmap[i][j] == "X" {
				sum++
			}
		}
	}
	fmt.Println(sum)
	fmt.Println(len(visited))
}

func isLoop(labmap [][]string, gi, gj int) bool {
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	diridx := 0
	dir := dirs[diridx]
	dirchar := "urdl"
	visited := make(map[string]int)
	for inBounds(gi, gj) {
		nexti, nextj := gi+dir[0], gj+dir[1]
		if inBounds(nexti, nextj) {
			if labmap[nexti][nextj] == "#" || labmap[nexti][nextj] == "O" {
				coords := strconv.Itoa(gi) + "," + strconv.Itoa(gj) + "," + string(dirchar[diridx])
				if visited[coords] > 3 {
					return true
				}
				visited[coords]++
				diridx++
				diridx %= 4
				dir = dirs[diridx]
				continue
			} else {
				gi += dir[0]
				gj += dir[1]
			}
		} else {
			gi += dir[0]
			gj += dir[1]
		}
	}
	return false
}

func inBounds(i, j int) bool {
	return i >= 0 && j >= 0 && i < rows && j < cols
}
