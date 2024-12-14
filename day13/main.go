package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	a, b, prize := [][]int{}, [][]int{}, [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Prize:") {
			coords := strings.Fields(line)
			x, _ := strconv.Atoi(coords[1][2 : len(coords[1])-1])
			y, _ := strconv.Atoi(coords[2][2:])
			prize = append(prize, []int{x + 10000000000000, y + 10000000000000})
		} else if strings.Contains(line, "Button") {
			coords := strings.Fields(line)
			x, _ := strconv.Atoi(coords[2][2 : len(coords[2])-1])
			y, _ := strconv.Atoi(coords[3][2:])
			if strings.Contains(line, "A") {
				a = append(a, []int{x, y})
			} else {
				b = append(b, []int{x, y})
			}
		}
	}
	fmt.Println(a, b, prize)

	tokens := 0
	for i, target := range prize {
		paths := []int{}
		k := (target[1]*a[i][0] - target[0]*a[i][1]) / (a[i][0]*b[i][1] - b[i][0]*a[i][1])
		if k == 0 {
			fmt.Println("no paths exist", a[i], b[i], target)
		}
		if (target[0]-k*b[i][0])%a[i][0] == 0 && (target[1]-k*b[i][1])%a[i][1] == 0 {
			paths = append(paths, k+3*(target[0]-k*b[i][0])/a[i][0])
		}
		if len(paths) > 0 {
			tokens += slices.Max(paths)
		}
		if len(paths) > 1 {
			fmt.Println("many paths exist", a[i], b[i], target)
		}
	}
	fmt.Println(tokens)

}
