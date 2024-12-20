package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	fullgrid := []string{}
	end := 70
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fullgrid = append(fullgrid, scanner.Text())
	}

	i, j := 0, len(fullgrid)
	for i < j-1 {
		grid := make(map[string]int)
		for k := range (i + j) / 2 {
			grid[fullgrid[k]] = 1
		}
		success := false
		q := [][]int{{0, 0}}
		visited := make(map[string]bool)
		steps := 0
	outer:
		for len(q) > 0 {
			levelsize := len(q)
			newq := [][]int{}
			for i := range levelsize {
				curr := q[i]
				x := strconv.Itoa(curr[0])
				y := strconv.Itoa(curr[1])
				if curr[0] < 0 || curr[0] > end || curr[1] < 0 || curr[1] > end {
					continue
				}
				if curr[0] == end && curr[1] == end {
					// fmt.Println(i, steps)
					success = true
					break outer
				}
				if grid[x+","+y] == 1 || visited[x+","+y] {
					continue
				} else {
					visited[x+","+y] = true
					newq = append(newq, []int{curr[0] + 1, curr[1]})
					newq = append(newq, []int{curr[0] - 1, curr[1]})
					newq = append(newq, []int{curr[0], curr[1] + 1})
					newq = append(newq, []int{curr[0], curr[1] - 1})
				}
			}
			steps++
			q = newq
		}
		if success {
			i = (i + j) / 2
		} else {
			j = (i + j) / 2
		}
	}
	fmt.Println(fullgrid[i])

}
