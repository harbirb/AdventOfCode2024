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
	scanner := bufio.NewScanner(file)
	var line []string
	for scanner.Scan() {
		line = strings.Split(scanner.Text(), "")
	}

	disk := []string{}
	isFree := false
	for i, j := range line {
		blk_size, _ := strconv.Atoi(j)
		for range blk_size {
			if isFree {
				disk = append(disk, ".")
			} else {
				disk = append(disk, strconv.Itoa(i/2))
			}
		}
		isFree = !isFree
	}

	h, t := 0, len(disk)-1

	for h < t {
		// loop over file blks, using tail pointers
		t_front := t
		for disk[t_front] == disk[t] {
			t_front--
		}
		sblk := t - t_front
		h_front := h
		// loop over all possible free blocks, use temp head pointers
		for h_front < t {
			h_end := h_front
			for disk[h_end] == "." {
				h_end++
			}
			sfree := h_end - h_front
			// if free space is enough, swap it in
			if sfree >= sblk {
				id := disk[t]
				for i := range sblk {
					disk[h_front+i] = id
					disk[t-i] = "."
				}
				h = 0
				for disk[h] != "." {
					h++
				}
				t = t_front
				for disk[t] == "." {
					t--
				}
				// break out of free block loop, and into file block loop
				break

			} else {
				// move temp head pointer to the next free block
				h_front = h_end
				for disk[h_front] != "." {
					h_front++
				}
			}
		}
		// no space for curr block, move tail back to decrement file-id
		if h_front > t {
			t = t_front
			for disk[t] == "." {
				t--
			}

		}

	}

	// for h < t {
	// 	if disk[h] == "." && disk[t] != "." {
	// 		disk[h] = disk[t]
	// 		disk[t] = "."
	// 	} else {
	// 		for disk[h] != "." {
	// 			h++
	// 		}
	// 		for disk[t] == "." {
	// 			t--
	// 		}
	// 	}
	// }
	fmt.Println(disk)

	// checksum
	sum := 0
	for i, id := range disk {
		if id != "." {
			idnum, _ := strconv.Atoi(id)
			sum += i * idnum
		}
	}
	fmt.Println(sum)

}
