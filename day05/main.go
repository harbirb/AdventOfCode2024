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

	r := make(map[int][]int) //map of rules r[X] = [numbers that must appear after X]
	updates := [][]int{}     // 2d array of updates
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			nums := strings.Split(line, "|")
			before, _ := strconv.Atoi(nums[0])
			after, _ := strconv.Atoi(nums[1])
			r[before] = append(r[before], after)
		} else if strings.Contains(line, ",") {
			nums := strings.Split(line, ",")
			update := []int{}
			for _, num := range nums {
				n, _ := strconv.Atoi(num)
				update = append(update, n)
			}
			updates = append(updates, update)
		}
	}

	sum := 0
	sum2 := 0

	for _, update := range updates {
		valid := true
		swapped := false
		for i, _ := range update {
			for j := 0; j < i; j++ {
				if slices.Contains(r[update[i]], update[j]) {
					valid = false
					temp := update[i]
					update[i] = update[j]
					update[j] = temp
					swapped = true
				}
			}
		}
		if valid {
			sum += update[len(update)/2]
		}
		if swapped {
			sum2 += update[len(update)/2]
		}

	}
	fmt.Println(sum)
	fmt.Println(sum2)
}
