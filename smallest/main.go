package main

import "fmt"

func main() {
	fmt.Printf("Started\n")
}

func Solution(A []int) int {
	// write your code in Go 1.4

	// Keep a list of seen number.
	// Not every seen number to be stored, only store a new seen number which is greater than current min.
	seen := make(map[int]bool)
	min := 1
	for _, n := range A {
		if n < min {
			// Min is already bigger than this number, no need to store
			continue
		} else if n > min {
			// Store this seen number for checking afterwards to determine a new min
			seen[n] = true
		} else { // n == min
			// Find the next possible minimum
			for {
				min++
				if seen[min] == false {
					break
				}
			}
		}
	}
	return min
}
