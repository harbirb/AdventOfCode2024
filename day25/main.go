package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	objs := [][][]string{}
	obj := [][]string{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			objs = append(objs, obj)
			obj = [][]string{}
		} else {
			obj = append(obj, strings.Split(scanner.Text(), ""))
		}

	}
	objs = append(objs, obj)

	k := [][5]int{}
	l := [][5]int{}
	for _, obj := range objs {
		isLock := obj[0][0] == "#"
		hts := [5]int{-1, -1, -1, -1, -1}
		for h, o := range obj {
			for i := range o {
				if isLock {
					if o[i] == "." && hts[i] == -1 {
						hts[i] = h - 1
					}
				} else {
					if o[i] == "#" && hts[i] == -1 {
						hts[i] = 6 - h
					}
				}
			}
		}
		if isLock {
			l = append(l, hts)
		} else {
			k = append(k, hts)
		}

	}

	pairs := 0
	for _, key := range k {
		for _, lock := range l {
			fits := true
			for i := range key {
				if key[i]+lock[i] >= 6 {
					fits = false
					break
				}
			}
			if fits {
				pairs++
			}
		}
	}

	fmt.Println(k, l)
	fmt.Println(pairs)
	fmt.Println(time.Since(start))
}
