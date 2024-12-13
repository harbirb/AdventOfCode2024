package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var rows, cols int
var data [][]string

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	data = [][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, strings.Split(scanner.Text(), ""))
	}

	rows, cols = len(data), len(data[0])

	sum := 0
	for i := range data {
		for j := range data[i] {
			if data[i][j] != "" {
				v := make(map[string]int)
				f := make(map[string]int)
				dfs(data[i][j], "START", v, f, i, j)
				// hashmap of fence location and direction. Iterate through map and condense fences with touching edges to get numsides
				for k := range f {
					fence := strings.Split(k, ",")
					fi, fj, fdir := fence[0], fence[1], fence[2]
					if fdir == "up" || fdir == "down" {
						origin, _ := strconv.Atoi(fj)
						dj := 1
						for f[fi+","+strconv.Itoa(origin+dj)+","+fdir] > 0 {
							delete(f, fi+","+strconv.Itoa(origin+dj)+","+fdir)
							dj++
						}
						dj = 1
						for f[fi+","+strconv.Itoa(origin-dj)+","+fdir] > 0 {
							delete(f, fi+","+strconv.Itoa(origin-dj)+","+fdir)
							dj++
						}
					} else if fdir == "right" || fdir == "left" {
						origin, _ := strconv.Atoi(fi)
						di := 1
						for f[strconv.Itoa(origin+di)+","+fj+","+fdir] > 0 {
							delete(f, strconv.Itoa(origin+di)+","+fj+","+fdir)
							di++
						}
						di = 1
						for f[strconv.Itoa(origin-di)+","+fj+","+fdir] > 0 {
							delete(f, strconv.Itoa(origin-di)+","+fj+","+fdir)
							di++
						}
					}
				}
				sum += len(f) * len(v)
			}
		}
	}
	fmt.Println(sum)
}

func dfs(curr, dir string, v, f map[string]int, i, j int) {
	coords := strconv.Itoa(i) + "," + strconv.Itoa(j)
	if i < 0 || i >= rows || j < 0 || j >= cols || (data[i][j] != curr && v[coords] == 0) {
		f[coords+","+dir]++
		return
	}
	if data[i][j] == curr && v[coords] == 0 {
		v[coords]++
		data[i][j] = ""
		dfs(curr, "down", v, f, i+1, j)
		dfs(curr, "up", v, f, i-1, j)
		dfs(curr, "right", v, f, i, j+1)
		dfs(curr, "left", v, f, i, j-1)
	}
}
