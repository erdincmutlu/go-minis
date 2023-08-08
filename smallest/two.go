package main

//package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")
import "sort"

func SolutionTwo(P []int, S []int) int {
	// write your code in Go 1.4

	passengers := 0
	for _, p := range P {
		passengers += p
	}

	sort.Ints(S)
	cars := 0
	for i := len(S) - 1; i >= 0; i-- {
		if passengers <= 0 {
			break
		}
		passengers -= S[i]
		cars++
	}

	return cars
}
