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

var sm = make(map[string][]string)

func main() {
	start := time.Now()
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
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
				if slices.Contains(sm[v1], v2) {
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
	fmt.Println(len(conn))

	cand := map[string]bool{}
	for s := range sm {
		cand[s] = true
	}
	poo := cf(map[string]bool{}, cand, map[string]bool{})
	names := []string{}
	for k := range poo {
		names = append(names, k)
	}
	sort.Strings(names)
	fmt.Println(strings.Join(names, ","))

	fmt.Printf("Program took %s to run.\n", time.Since(start))
}

// Visit each node v. Candidates = Neighbors of v.
// Recursively call w next candidates = intersection of (v and N(candidate))

// function(clique, candidates, visited)
// start with empty clique, all v's as candidates, empty visited
//  pick a v from candidates
//  add v to clique
//  add v to visited
//  new candidates = intersection of candidates and neighbors of v

func cf(cl, cand, visited map[string]bool) map[string]bool {
	newcl := make(map[string]bool)
	for i := range cl {
		newcl[i] = true
	}
	newv := make(map[string]bool)
	for j := range visited {
		newv[j] = true
	}
	maxcl := make(map[string]bool)
	if len(cand) == 0 {
		return cl
	}
	for k := range cand {
		newcl[k] = true
		newv[k] = true
		newcand := make(map[string]bool)
		for _, neigh := range sm[k] {
			if newv[neigh] {
				continue
			}
			newcand[neigh] = true
		}
		intsec := intersection(cand, newcand)
		newmaxcl := cf(newcl, intsec, newv)
		if len(maxcl) == 0 || len(newmaxcl) > len(maxcl) {
			// copy elementwise
			for f := range maxcl {
				delete(maxcl, f)
			}
			for f := range newmaxcl {
				maxcl[f] = true
			}
		}
		delete(newcl, k)
	}
	return maxcl

}

func intersection(s1, s2 map[string]bool) map[string]bool {
	intsec := make(map[string]bool)
	for k := range s1 {
		if s2[k] {
			intsec[k] = true
		}
	}
	return intsec
}

// func pk(m map[string]bool) {
// 	for k := range m {
// 		fmt.Print(k, " ")
// 	}
// 	fmt.Print("\n")
// }
