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
	cmds := [][]string{}
	cmds1 := [][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cmds = append(cmds, strings.Split(scanner.Text(), ""))
	}
	kp := map[string][]int{"0": {1, 0}, "A": {2, 0}, "1": {0, 1}, "2": {1, 1}, "3": {2, 1}, "4": {0, 2}, "5": {1, 2}, "6": {2, 2}, "7": {0, 3}, "8": {1, 3}, "9": {2, 3}}
	// dp := map[string][]int{"<": {0, 0}, "v": {1, 0}, ">": {2, 0}, "^": {1, 1}, "A": {2, 1}}

	for _, cmd := range cmds {
		cmd1 := ""
		x, y := 2, 0
		for _, c := range cmd {
			nextx, nexty := kp[c][0], kp[c][1]
			// fmt.Println(nextx, nexty)
			for x != nextx || y != nexty {
				if x < nextx {
					x++
					cmd1 += ">"
				} else if y < nexty {
					y++
					cmd1 += "^"
				} else if x > nextx && (y != 0 || x-1 != 0) {
					x--
					cmd1 += "<"
				} else if y > nexty && (x != 0 || y-1 != 0) {
					y--
					cmd1 += "v"
				}
			}
			if x == nextx && y == nexty {
				cmd1 += "A"
				continue
			}
		}
		fmt.Println(cmd1)
		cmds1 = append(cmds1, strings.Split(cmd1, ""))
	}

	// for _, cmd1 := range cmds1 {
	// 	cmd2 := ""
	// 	x, y := 2, 1
	// 	for _, c := range cmd1 {
	// 		// fmt.Print(c)
	// 		nextx, nexty := dp[c][0], dp[c][1]
	// 		// fmt.Println(nextx, nexty)
	// 		for x != nextx || y != nexty {
	// 			if x < nextx {
	// 				x++
	// 				cmd2 += ">"
	// 			} else if y > nexty {
	// 				y--
	// 				cmd2 += "v"
	// 			} else if x > nextx && (y != 1 || x-1 != 0) {
	// 				x--
	// 				cmd2 += "<"
	// 			} else if y < nexty && x != 0 {
	// 				y++
	// 				cmd2 += "^"
	// 			}
	// 		}
	// 		if x == nextx && y == nexty {
	// 			cmd2 += "A"
	// 			continue
	// 		}
	// 	}

	// }
	cmds2 := t2(cmds1)
	fmt.Println(cmds2)
	cmds3 := t2(cmds2)
	fmt.Println(cmds3)
	for _, cmd := range cmds3 {
		fmt.Println(len(cmd))
	}

	numeric := []int{}
	for _, cmd := range cmds {
		strtoint := ""
		for _, c := range cmd {
			_, e := strconv.Atoi(c)
			if e == nil {
				strtoint += c
			}
		}
		v, _ := strconv.Atoi(strtoint)
		numeric = append(numeric, v)
	}
	fmt.Println(numeric)

	sample := [][]string{}
	s := strings.Split("<A>Av<<AA>^AA>AvAA^Av<AAA>^A", "")
	sample = append(sample, s)
	fmt.Println(t2(sample))

}

func t2(input [][]string) [][]string {
	dp := map[string][]int{"<": {0, 0}, "v": {1, 0}, ">": {2, 0}, "^": {1, 1}, "A": {2, 1}}

	output := [][]string{}
	for _, cmd1 := range input {
		cmd2 := ""
		x, y := 2, 1
		for _, c := range cmd1 {
			// fmt.Print(c)
			nextx, nexty := dp[c][0], dp[c][1]
			// fmt.Println(nextx, nexty)
			for x != nextx || y != nexty {
				if x < nextx {
					x++
					cmd2 += ">"
				} else if y > nexty {
					y--
					cmd2 += "v"
				} else if x > nextx && (y != 1 || x-1 != 0) {
					x--
					cmd2 += "<"
				} else if y < nexty && x != 0 {
					y++
					cmd2 += "^"
				}
			}
			if x == nextx && y == nexty {
				cmd2 += "A"
				continue
			}
		}
		fmt.Println(cmd2)
		output = append(output, strings.Split(cmd2, ""))
	}
	return output
}
