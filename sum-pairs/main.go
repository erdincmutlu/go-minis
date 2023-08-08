package main

import "fmt"

// Given a set of numbers -50 to 50, find all pairs that add up to a certain sum
// that is passed in.
// Implement the O(n) solution

func sumPairs(nums []int, sum int) [][]int {
	var out [][]int
	numbers := make(map[int]bool, len(nums))
	for _, n := range nums {
		if numbers[sum-n] {
			out = append(out, []int{n, sum - n})
		} else {
			numbers[n] = true
		}
	}

	return out
}

func main() {
	fmt.Printf("%+v", sumPairs([]int{-20, -10, -11, 1, 5, 20, 12, 15, 35, 13, 45, 41}, 25))
}
