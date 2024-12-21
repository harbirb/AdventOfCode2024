package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// var pm = make(map[string]int)
var tm = make(map[string]int)
var ps []string

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	ts := []string{}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), ",") {
			ps = strings.Split(scanner.Text(), ", ")
		} else if scanner.Text() != "" {
			ts = append(ts, scanner.Text())
		}
	}
	num := 0
	for _, t := range ts {
		num += dfs(t)
		fmt.Println(t, tm[t])
	}
	fmt.Println(num)

}

func dfs(t string) int {
	if tm[t] > 0 {
		return tm[t]
	} else if tm[t] < 0 {
		return 0
	}
	w := 0
	for _, p := range ps {
		if p == t {
			w += 1
		} else if len(p) < len(t) && t[:len(p)] == p {
			w += dfs(t[len(p):])
		}
	}
	if w == 0 {
		tm[t] = -1
	} else {
		tm[t] = w
	}
	return w
}
