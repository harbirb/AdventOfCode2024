package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	data := string(dat)
	r := regexp.MustCompile(`mul\(\d+\,\d+\)|do\(\)|don't\(\)`)
	instrs := r.FindAllString(data, -1)

	result := 0
	do := true
	for _, instr := range instrs {
		if instr == "do()" {
			do = true
			continue
		} else if instr == "don't()" {
			do = false
			continue
		} else {
			if do {
				// valid mul instruction
				strs := strings.Split(instr, ",")
				n1, _ := strconv.Atoi(strs[0][4:])
				n2, _ := strconv.Atoi(strs[1][:len(strs[1])-1])

				result += n1 * n2
			}
		}
	}

	fmt.Println(result)
}
