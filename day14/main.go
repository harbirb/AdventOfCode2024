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
	// p=2,4 v=2,-3
	// p=9,5 v=-3,-3
	defer file.Close()
	ps := [][]int{}
	vs := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pos := strings.Split(strings.Fields(scanner.Text())[0], ",")
		x, _ := strconv.Atoi(pos[0][2:])
		y, _ := strconv.Atoi(pos[1])
		ps = append(ps, []int{x, y})
		vel := strings.Split(strings.Fields(scanner.Text())[1], ",")
		dx, _ := strconv.Atoi(vel[0][2:])
		dy, _ := strconv.Atoi(vel[1])
		vs = append(vs, []int{dx, dy})
	}
	w, h := 101, 103
	secs := 5
	var q1, q2, q3, q4 int
	for s := range secs {
		pm := make(map[string]int)
		for i := range ps {
			ps[i][0] = ((ps[i][0]+vs[i][0])%w + w) % w
			ps[i][1] = ((ps[i][1]+vs[i][1])%h + h) % h
			pm[strconv.Itoa(ps[i][0])+","+strconv.Itoa(ps[i][1])]++
		}
		q1, q2, q3, q4 = 0, 0, 0, 0
		for _, p := range ps {
			if p[0] < w/2 && p[1] < h/2 {
				q1++
			} else if p[0] > w/2 && p[1] < h/2 {
				q2++
			} else if p[0] < w/2 && p[1] > h/2 {
				q3++
			} else if p[0] > w/2 && p[1] > h/2 {
				q4++
			}
		}
		// check for tree in pm
		maxline := 0
		for k := range pm {
			x, _ := strconv.Atoi(strings.Split(k, ",")[0])
			line := 0
			for x < w {
				nextbot := strconv.Itoa(x) + "," + strings.Split(k, ",")[1]
				x++
				line++
				if pm[nextbot] == 0 {
					break
				}

			}
			maxline = max(maxline, line)

		}
		if maxline > 9 {
			draw(ps)
			fmt.Println(s + 1)
		}

	}
	fmt.Println(q1 * q2 * q3 * q4)
}

func draw(ps [][]int) {
	var b [103][101]string
	for i := range b {
		for j := range b[i] {
			b[i][j] = "."
		}
	}
	for _, p := range ps {
		b[p[1]][p[0]] = "#"
	}
	for i := range b {
		fmt.Println(b[i])
	}

	// Open or create a file for writing
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write the 2D slice to the file row by row
	for _, row := range b {
		line := ""
		for i, cell := range row {
			if i > 0 {
			}
			line += cell
		}
		line += "\n" // End each row with a newline
		if _, err := file.WriteString(line); err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}
