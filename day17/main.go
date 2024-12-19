package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var a, b, c int

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	instrs := []int{}
	// ip := 0
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "Register A:") {
			a, _ = strconv.Atoi(strings.Split(scanner.Text(), ": ")[1])
		}
		if strings.Contains(scanner.Text(), "Register B:") {
			b, _ = strconv.Atoi(strings.Split(scanner.Text(), ": ")[1])
		}
		if strings.Contains(scanner.Text(), "Register C:") {
			c, _ = strconv.Atoi(strings.Split(scanner.Text(), ": ")[1])
		}
		if strings.Contains(scanner.Text(), "Program:") {
			for _, v := range strings.Split(strings.Split(scanner.Text(), ": ")[1], ",") {
				i, _ := strconv.Atoi(v)
				instrs = append(instrs, i)
			}
		}
	}

	// ip = 0
	// a, b, c = 0, 0, 0
	// for ip < len(instrs)-1 {
	// 	switch instrs[ip] {
	// 	case 0:
	// 		num := a
	// 		den := int(math.Pow(2, float64(combo(instrs[ip+1]))))
	// 		a = num / den
	// 		ip += 2
	// 	case 1:
	// 		b = b ^ instrs[ip+1]
	// 		ip += 2
	// 	case 2:
	// 		b = combo(instrs[ip+1]) % 8
	// 		ip += 2
	// 	case 3:
	// 		if a != 0 {
	// 			ip = instrs[ip+1]
	// 		} else {
	// 			ip += 2
	// 		}
	// 	case 4:
	// 		b = b ^ c
	// 		ip += 2
	// 	case 5:
	// 		output = append(output, combo(instrs[ip+1])%8)
	// 		ip += 2
	// 	case 6:
	// 		num := a
	// 		den := int(math.Pow(2, float64(combo(instrs[ip+1]))))
	// 		b = num / den
	// 		ip += 2
	// 	case 7:
	// 		num := a
	// 		den := int(math.Pow(2, float64(combo(instrs[ip+1]))))
	// 		c = num / den
	// 		ip += 2
	// 	}

	a = 0
	for i := 1; i <= len(instrs); i++ {
		for {
			match := true
			for j := range i {
				next := len(instrs) - i + j
				denom := int(math.Pow(8, float64(j)))
				if !it(a/denom, instrs[next]) {
					match = false
					a++
					break
				}
			}
			if match {
				fmt.Println(a, instrs[len(instrs)-i])
				prog(a)
				a *= 8
				break
			}
		}

	}

	// to get x, b = 8m +x. a*8

	// for i := 177313 * 8; i < 177313*8*8; i++ {
	// 	if it(i, 4) && it(i/8, 0) && it(i/64, 3) {
	// 		fmt.Print(i)
	// 		prog(i)
	// 		break
	// 	}
	// }

}

func it(a, target int) bool {
	b = 0
	c = 0
	b = a % 8
	b = b ^ 2
	c = a / int(math.Pow(2, float64(b)))
	b = b ^ 7
	b = b ^ c
	return target == b%8
}

func prog(a int) {

	for {
		b = a % 8
		b = b ^ 2
		c = a / int(math.Pow(2.0, float64(b)))
		b = b ^ 7
		b = b ^ c
		a = a / 8
		fmt.Print(b % 8)
		fmt.Print(",")
		if a == 0 {
			fmt.Println()
			break
		}
	}
}
