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
	vals, nums := []int{}, [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":")
		v, _ := strconv.Atoi(parts[0])
		vals = append(vals, v)
		line := strings.Fields(parts[1])
		numlist := make([]int, len(line))
		for i, num := range line {
			numlist[i], _ = strconv.Atoi(num)
		}
		nums = append(nums, numlist)
	}

	sum1, sum2 := 0, 0

	for i, ns := range nums {
		if dfs(ns, 1, ns[0], vals[i]) {
			sum1 += vals[i]
		}
		if bfsIterative(ns, vals[i]) {
			sum2 += vals[i]
		}
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
}

func dfs(nums []int, i, currval, val int) bool {
	if i == len(nums) {
		return currval == val
	}
	if currval > val {
		return false
	}
	return dfs(nums, i+1, currval+nums[i], val) || dfs(nums, i+1, currval*nums[i], val) || dfs(nums, i+1, concat(currval, nums[i]), val)
}

func concat(first, second int) int {
	n1, n2 := strconv.Itoa(first), strconv.Itoa(second)
	result, _ := strconv.Atoi(n1 + n2)
	return result

}

func bfsIterative(nums []int, val int) bool {
	q := []int{nums[0]}
	nums = nums[1:]
	for len(nums) > 0 {
		nextq := []int{}
		for _, currVal := range q {
			addVal := currVal + nums[0]
			mulVal := currVal * nums[0]
			concatVal := concat(currVal, nums[0])
			nextq = append(nextq, addVal, mulVal, concatVal)
		}
		nums = nums[1:]
		q = nextq
	}
	for _, currval := range q {
		if currval == val {
			return true
		}
	}
	return false
}
