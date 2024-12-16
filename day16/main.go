package main

import (
	"bufio"
	"container/heap"
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
	maze := [][]string{}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		maze = append(maze, strings.Split(scanner.Text(), ""))
	}

	var start string
	for i := range maze {
		for j := range maze[i] {
			if maze[i][j] == "S" {
				start = strconv.Itoa(i) + "," + strconv.Itoa(j) + "," + "E"
				maze[i][j] = "O"
			}
		}
	}
	visited := make(map[string]int)
	dmap := map[string][]int{"E": {0, 1}, "W": {0, -1}, "N": {-1, 0}, "S": {1, 0}}
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Node{tile: start, score: 0})

	var sol Node
	for pq.Len() > 0 {
		curr := heap.Pop(pq).(Node)
		if visited[curr.tile] == 0 || visited[curr.tile] > curr.score {
			visited[curr.tile] = curr.score
		}
		coords := strings.Split(curr.tile, ",")
		i, _ := strconv.Atoi(coords[0])
		j, _ := strconv.Atoi(coords[1])
		dir := coords[2]
		// rotate in all directions
		for k := range dmap {
			if visited[coords[0]+","+coords[1]+","+k] == 0 && k != dir {
				heap.Push(pq, Node{tile: coords[0] + "," + coords[1] + "," + k, score: curr.score + 1000})
			}
		}

		// move in current direction
		newcoords := strconv.Itoa(i+dmap[dir][0]) + "," + strconv.Itoa(j+dmap[dir][1])
		if maze[i+dmap[dir][0]][j+dmap[dir][1]] != "#" && visited[newcoords+","+dir] == 0 {
			heap.Push(pq, Node{tile: newcoords + "," + dir, score: curr.score + 1})
		}

		if maze[i][j] == "E" {
			fmt.Println(curr.score)
			sol = curr
			break
		}
	}

	// go to node, check visited in all possible directions, and check by moving back in current facing direction
	// if that node's score is less than current node's score, it is part of a best path
	q := []string{sol.tile}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		coords := strings.Split(curr, ",")
		i, _ := strconv.Atoi(coords[0])
		j, _ := strconv.Atoi(coords[1])
		dir := coords[2]
		// move back in current direction
		prevtile := strconv.Itoa(i-dmap[dir][0]) + "," + strconv.Itoa(j-dmap[dir][1]) + "," + dir
		prevscore := visited[prevtile]
		if visited[curr] > prevscore && prevscore > 0 {
			q = append(q, prevtile)
		}
		// check different facing direction
		for k := range dmap {
			if k != dir {
				prevtile = coords[0] + "," + coords[1] + "," + k
				prevscore = visited[prevtile]
				if visited[curr] > prevscore && prevscore > 0 {
					q = append(q, prevtile)
				}
			}
		}
		maze[i][j] = "O"
	}

	sumo := 0
	for i := range maze {
		for j := range maze[i] {
			if maze[i][j] == "O" {
				sumo += 1
			}
		}
	}
	fmt.Println(sumo)

}

type Node struct {
	tile  string
	score int
}

type PriorityQueue []Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest score first
	return pq[i].score < pq[j].score
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(Node))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
