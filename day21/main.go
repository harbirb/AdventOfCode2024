package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type path struct {
	x        int
	y        int
	currpath string
}

func main() {
	start := time.Now()
	file, _ := os.Open("input.txt")
	defer file.Close()
	codes := []string{}
	seq_lengths := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		codes = append(codes, scanner.Text())
	}
	kp := map[string][]int{"0": {1, 0}, "A": {2, 0}, "1": {0, 1}, "2": {1, 1}, "3": {2, 1}, "4": {0, 2}, "5": {1, 2}, "6": {2, 2}, "7": {0, 3}, "8": {1, 3}, "9": {2, 3}}

	for _, code := range codes {
		seq_length := 0
		x, y := 2, 0
		for _, c := range strings.Split(code, "") {
			poss := []string{}
			// bfs to generate all possible paths to the next key on the keypad
			q := []path{{x, y, ""}}
			nextx, nexty := kp[c][0], kp[c][1]
			for len(q) > 0 {
				curr := q[0]
				nextq := q[1:]
				if curr.x != nextx || curr.y != nexty {
					if curr.x == 0 && curr.y == 0 {
						q = nextq
						continue
					}
					if curr.x < nextx {
						nextq = append(nextq, path{curr.x + 1, curr.y, curr.currpath + ">"})
					}
					if curr.y < nexty {
						nextq = append(nextq, path{curr.x, curr.y + 1, curr.currpath + "^"})
					}
					if curr.x > nextx {
						nextq = append(nextq, path{curr.x - 1, curr.y, curr.currpath + "<"})
					}
					if curr.y > nexty {
						nextq = append(nextq, path{curr.x, curr.y - 1, curr.currpath + "v"})
					}
				} else {
					poss = append(poss, curr.currpath+"A")
				}
				q = nextq
			}
			x, y = nextx, nexty
			shortest_len := 0
			for _, seq := range poss {
				curr_len := t(seq, 25)
				if shortest_len == 0 {
					shortest_len = curr_len
				} else if curr_len < shortest_len {
					shortest_len = curr_len
				}
			}
			seq_length += shortest_len
		}
		seq_lengths = append(seq_lengths, seq_length)
	}

	numeric := []int{}
	for _, code := range codes {
		numChars := ""
		for _, c := range strings.Split(code, "") {
			_, e := strconv.Atoi(c)
			if e == nil {
				numChars += c
			}
		}
		v, _ := strconv.Atoi(numChars)
		numeric = append(numeric, v)
	}
	sum := 0
	for i := range numeric {
		sum += numeric[i] * seq_lengths[i]
	}
	fmt.Println(sum)
	fmt.Printf("Program took %s to run.\n", time.Since(start))

}

var dp = map[string][]int{"<": {0, 0}, "v": {1, 0}, ">": {2, 0}, "^": {1, 1}, "A": {2, 1}}

// memo maps from a depth+input to the output length after performing n translations where n=depth
var memo = make(map[string]int)

// function to translate an input recursively, depth times. Returns length of translated sequence
// depth specifies number of calls. Depth == 1 means the input is translated only once.
func t(input string, depth int) int {
	if depth == 0 {
		return len(input)
	}
	outputlen := 0
	di := strconv.Itoa(depth) + input
	if memo[di] != 0 {
		return memo[di]
	} else {
		x, y := 2, 1
		for _, c := range strings.Split(input, "") {
			poss := []string{}
			q := []path{{x, y, ""}}
			nextx, nexty := dp[c][0], dp[c][1]
			for len(q) > 0 {
				curr := q[0]
				nextq := q[1:]
				if curr.x != nextx || curr.y != nexty {
					if curr.x == 0 && curr.y == 1 {
						q = nextq
						continue
					}
					if curr.x < nextx {
						nextq = append(nextq, path{curr.x + 1, curr.y, curr.currpath + ">"})
					}
					if curr.y > nexty {
						nextq = append(nextq, path{curr.x, curr.y - 1, curr.currpath + "v"})
					}
					if curr.x > nextx {
						nextq = append(nextq, path{curr.x - 1, curr.y, curr.currpath + "<"})
					}
					if curr.y < nexty {
						nextq = append(nextq, path{curr.x, curr.y + 1, curr.currpath + "^"})
					}
				} else {
					poss = append(poss, curr.currpath+"A")
				}
				q = nextq
			}
			x, y = nextx, nexty
			shortestlen := 0
			for _, p := range poss {
				nextp := t(p, depth-1)
				if shortestlen == 0 {
					shortestlen = nextp
				} else if nextp < shortestlen {
					shortestlen = nextp
				}
			}
			outputlen += shortestlen
		}
		memo[di] = outputlen
		return outputlen
	}
}
