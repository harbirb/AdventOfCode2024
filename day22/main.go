package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	secrets := []int{}
	for scanner.Scan() {
		secret, _ := strconv.Atoi(scanner.Text())
		secrets = append(secrets, secret)
	}
	// fmt.Println(secrets)

	prices := [][]int{}
	changes := [][]string{}
	for _, s := range secrets {
		prices = append(prices, []int{s % 10})
		changes = append(changes, []string{})
	}

	for range 2000 {
		for i, secret := range secrets {
			init_secret := secret
			secret64 := 64 * secret
			secret = secret ^ secret64
			secret = secret % 16777216
			secret32 := secret / 32
			secret = secret ^ secret32
			secret = secret % 16777216
			secret2048 := secret * 2048
			secret = secret ^ secret2048
			secret = secret % 16777216
			secrets[i] = secret
			price := secret % 10
			prices[i] = append(prices[i], price)
			change := price - init_secret%10
			chg_str := ""
			chg_str += strconv.Itoa(change)
			changes[i] = append(changes[i], chg_str)
		}
	}
	// fmt.Println(secrets)
	// fmt.Println(prices)
	// fmt.Println(changes)

	sm := make(map[string]int)

	for i := range changes {
		visited := make(map[string]bool)
		for j := 0; j < len(changes[i])-4; j++ {
			seq := changes[i][j : j+4]
			seqstring := strings.Join(seq, ",")
			if visited[seqstring] {
				// not the first occurrence
				continue
			}
			visited[seqstring] = true
			price := prices[i][j+4]
			sm[seqstring] += price
		}
	}

	max := 0
	for _, v := range sm {
		if v > max {
			max = v
		}
	}
	fmt.Println(max)

	// sum := 0
	// for _, s := range secrets {
	// 	sum += s
	// }
	// fmt.Println(sum)

	fmt.Printf("Program took %s to run.\n", time.Since(start))
}
