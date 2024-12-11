package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	stones := strings.Split(string(file), " ")

	sm := make(map[string]int)
	for _, s := range stones {
		sm[s]++
	}

	for d := range 75 {
		fmt.Println(d)
		newsm := make(map[string]int)
		for k, v := range sm {
			num, _ := strconv.Atoi(k)
			if num == 0 {
				newsm["1"] += v
			} else if len(k)%2 == 0 {
				lhalf, _ := strconv.Atoi(k[:len(k)/2])
				rhalf, _ := strconv.Atoi(k[len(k)/2:])
				newsm[strconv.Itoa(lhalf)] += v
				newsm[strconv.Itoa(rhalf)] += v
			} else {
				newsm[strconv.Itoa(num*2024)] += v
			}
		}
		sm = newsm
		// fmt.Println(sm)
	}

	sum := 0
	for _, v := range sm {
		sum += v
	}
	fmt.Println(sum)

}
