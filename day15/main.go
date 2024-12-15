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
	w := [][]string{}
	ms := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "#") {
			pos := strings.Split(scanner.Text(), "")
			w = append(w, pos)
			continue
		}
		ms = append(ms, strings.Split(scanner.Text(), "")...)
	}

	move := map[string][]int{"^": {-1, 0}, "v": {1, 0}, "<": {0, -1}, ">": {0, 1}}

	// for _, m := range ms {
	// 	dir := move[m]
	// 	boxes := 1
	// 	for w[bot[0]+boxes*dir[0]][bot[1]+boxes*dir[1]] != "#" {
	// 		if w[bot[0]+boxes*dir[0]][bot[1]+boxes*dir[1]] == "." {
	// 			w[bot[0]+boxes*dir[0]][bot[1]+boxes*dir[1]] = w[bot[0]+dir[0]][bot[1]+dir[1]]
	// 			w[bot[0]][bot[1]] = "."
	// 			w[bot[0]+dir[0]][bot[1]+dir[1]] = "@"
	// 			bot[0] += dir[0]
	// 			bot[1] += dir[1]
	// 			break
	// 		}

	// 		boxes++
	// 	}
	// fmt.Println(m)
	// for i := range w {
	// 	fmt.Println(w[i])
	// }
	// }

	// sum := 0
	// for i := range w {
	// 	for j := range w[i] {
	// 		if w[i][j] == "O" {
	// 			sum += 100*i + j
	// 		}
	// 	}
	// }
	// fmt.Println(sum)

	// Resize things in the warehouse
	w2 := [][]string{}
	for i := range w {
		row := []string{}
		for j := range w[i] {
			if w[i][j] == "O" {
				row = append(row, "[", "]")
			} else if w[i][j] == "@" {
				row = append(row, "@", ".")
			} else {
				row = append(row, w[i][j], w[i][j])
			}
		}
		w2 = append(w2, row)
	}

	bot := []int{}
	for i := range w2 {
		for j := range w2[i] {
			if w2[i][j] == "@" {
				bot = append(bot, i, j)
			}
		}
	}

	for i := range w2 {
		fmt.Println(w2[i])
	}

	for _, m := range ms {
		dir := move[m]
		boxes := 1
		if m == "<" || m == ">" {
			pushing := false
			for w2[bot[0]+boxes*dir[0]][bot[1]+boxes*dir[1]] != "#" {
				if w2[bot[0]+boxes*dir[0]][bot[1]+boxes*dir[1]] == "." {
					// if pushing boxes left/right
					if pushing {
						w2[bot[0]][bot[1]+boxes*dir[1]] = w2[bot[0]][bot[1]+(boxes-1)*dir[1]]
						for i := range boxes {
							if w2[bot[0]][bot[1]+i*dir[1]] == "[" {
								w2[bot[0]][bot[1]+i*dir[1]] = "]"
							} else {
								w2[bot[0]][bot[1]+i*dir[1]] = "["
							}
						}
					}
					w2[bot[0]][bot[1]] = "."
					w2[bot[0]+dir[0]][bot[1]+dir[1]] = "@"
					bot[0] += dir[0]
					bot[1] += dir[1]
					break
				}
				pushing = true
				boxes++
			}
		} else {
			//if moving one box up/down, need to check 2 wide area in direction. If box in the way, need to check the next boxes areas ahead
			if isClear(w2, bot[0]+dir[0], bot[1], dir[0]) {
				shift(w2, bot[0]+dir[0], bot[1], dir[0])
				w2[bot[0]][bot[1]] = "."
				w2[bot[0]+dir[0]][bot[1]+dir[1]] = "@"
				bot[0] += dir[0]
			}
		}
	}

	for i := range w2 {
		fmt.Println(w2[i])
	}

	sum := 0
	for i := range w2 {
		for j := range w2[i] {
			if w2[i][j] == "[" {
				sum += 100*i + j
			}
		}
	}
	fmt.Println(sum)

}

// take in any part of a box and look ahead to see if it's clear
func isClear(w2 [][]string, i, j, dir int) bool {
	if w2[i][j] == "#" {
		return false
	}
	if w2[i][j] == "." {
		return true
	}
	if w2[i][j] == "[" {
		return isClear(w2, i+dir, j, dir) && isClear(w2, i+dir, j+1, dir)
	}
	if w2[i][j] == "]" {
		return isClear(w2, i+dir, j-1, dir) && isClear(w2, i+dir, j, dir)
	}
	return true
}

// take in any part of a box, clear the path ahead, then shift the current box
func shift(w2 [][]string, i, j, dir int) {
	if w2[i][j] == "[" {
		shift(w2, i+dir, j, dir)
		shift(w2, i+dir, j+1, dir)
		w2[i+dir][j] = "["
		w2[i+dir][j+1] = "]"
		w2[i][j] = "."
		w2[i][j+1] = "."

	} else if w2[i][j] == "]" {
		shift(w2, i, j-1, dir)
	}
}
