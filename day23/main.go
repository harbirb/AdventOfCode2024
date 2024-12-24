package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sm := make(map[string][]string)
	for scanner.Scan() {
		pairs := strings.Split(scanner.Text(), "-")
		sm[pairs[0]] = append(sm[pairs[0]], pairs[1])
		sm[pairs[1]] = append(sm[pairs[1]], pairs[0])
	}
	// fmt.Println(sm)

	conn := map[string]bool{}

	for k, vs := range sm {
		for _, v1 := range vs {
			for _, v2 := range vs {
				if slices.Contains(sm[v1], k) && slices.Contains(sm[v2], k) && slices.Contains(sm[v1], v2) && slices.Contains(sm[v2], v1) {
					if k[0] == 't' || v1[0] == 't' || v2[0] == 't' {
						sl := []string{k, v1, v2}
						sort.Strings(sl)
						set := strings.Join(sl, ",")
						conn[set] = true
					}
				}
			}
		}
	}

	// for k := range conn {
	// 	fmt.Println(k)
	// }
	fmt.Println(len(conn))

	visited := make(map[string]bool)
	for k, vs := range sm {
		visited[k] = true
		currset := make(map[string]bool)
		biggest_setlen := 0
		for _, v := range vs {
			// perform intersection of edges
			intsec := intersection(sm[k], sm[v])
			if len(intsec) > biggest_setlen {
				biggest_setlen = len(intsec)
			}

		}
	}

	fmt.Printf("Program took %s to run.\n", time.Since(start))
}

func intersection(s1, s2 []string) []string {
	intsec := []string{}
	for _, s := range s1 {
		if slices.Contains(s2, s) {
			intsec = append(intsec, s)
		}
	}
	return intsec
}
