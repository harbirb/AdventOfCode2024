package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var wires = make(map[string]int)
var gates = make(map[string]string)
var swaps = make(map[string]bool)

// check if hasValue value has been computed
var hasValue = make(map[string]bool)

func main() {
	start := time.Now()
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), ":") {
			wire := strings.Split(scanner.Text(), ": ")
			name := wire[0]
			val, _ := strconv.Atoi(wire[1])
			wires[name] = val
			hasValue[name] = true
		}
		if strings.Contains(scanner.Text(), " -> ") {
			gate := strings.Split(scanner.Text(), " -> ")
			gates[gate[1]] = gate[0]
			hasValue[gate[1]] = false
		}

	}
	for wire := range hasValue {
		if !hasValue[wire] {
			compute(wire)
		}
	}

	zwires := []string{}
	for k := range wires {
		if k[0] == 'z' {
			zwires = append(zwires, k)
		}
	}
	sort.Slice(zwires, func(a, b int) bool {
		return zwires[a] > zwires[b]
	})
	fmt.Println(zwires)

	res := 0
	for _, v := range zwires {
		res = res << 1
		res += wires[v]
	}
	fmt.Println(res)

	for k := range gates {
		d(k)
	}
	sw := []string{}
	for k := range swaps {
		sw = append(sw, k)
	}
	sort.Strings(sw)
	fmt.Println(strings.Join(sw, ","))

	fmt.Printf("Program took %s to run.\n", time.Since(start))
}

func compute(wire string) {
	gate := gates[wire]
	w1 := gate[0:3]
	w2 := gate[len(gate)-3:]
	if !hasValue[w1] {
		compute(w1)
	}
	if !hasValue[w2] {
		compute(w2)
	}
	if strings.Contains(gate, " AND ") {
		wires[wire] = wires[w1] & wires[w2]
	} else if strings.Contains(gate, " OR ") {
		wires[wire] = wires[w1] | wires[w2]
	} else if strings.Contains(gate, " XOR ") {
		wires[wire] = wires[w1] ^ wires[w2]
	}
	hasValue[wire] = true
}

func d(r string) {
	gate := gates[r]
	w1 := gate[0:3]
	w2 := gate[len(gate)-3:]
	op := strings.Fields(gate[3 : len(gate)-3])[0]
	if r[0] == 'z' {
		// if zwire comes from a gate without a XOR, that z wire is out of place
		if op != "XOR" && r != "z45" {
			swaps[r] = true
		}
	}
	// if the wire is not a z and thegate contains XOR, and one of the children also contain a XOR, that wire is out of place
	if op == "XOR" && !strings.ContainsAny(w1+w2, "xy") {
		if r[0] != 'z' {
			swaps[r] = true
		}
		op1 := strings.Fields(gates[w1])[1]
		op2 := strings.Fields(gates[w2])[1]
		if op1 != "OR" && op1 != "XOR" {
			swaps[w1] = true
		}
		if op2 != "OR" && op2 != "XOR" && !strings.Contains(gates[w2], "x00") {
			swaps[w2] = true
		}
	}
	// if the gate is an OR and one side does not contain an and, that side's wire is out of place
	if op == "OR" {
		if !strings.Contains(gates[w1], "AND") {
			swaps[w1] = true
		}
		if !strings.Contains(gates[w2], "AND") {
			swaps[w2] = true
		}
	}
}
