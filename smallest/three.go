package main

//package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")
import (
	"sort"
)

func SolutionThreeOld(A []int) int {
	// write your code in Go 1.4

	AFloat := []float64{}
	polution := 0.0
	for i := range A {
		AFloat = append(AFloat, float64(A[i]))
		polution += AFloat[i]
	}
	expected := polution / 2.0

	sort.Float64s(AFloat)
	highestIndex := len(AFloat) - 1
	filters := 0
	for {
		// fmt.Printf("Befor: polution:%f, AFloat:%+v, expected:%f\n", polution, AFloat, expected)

		// sort.Float64s(AFloat)
		// fmt.Printf("Sorte: polution:%f, AFloat:%+v, expected:%f\n", polution, AFloat, expected)

		// AFloat[len(AFloat)-1] = AFloat[len(AFloat)-1] / 2.0
		// polution -= AFloat[len(AFloat)-1]
		// filters++

		// fmt.Printf("After: polution:%f, AFloat:%+v, expected:%f\n", polution, AFloat, expected)

		// // find highest
		// max :=0
		// maxIndex :=
		// for i := 0; i < len(AFloat); i++ {
		// 	if AFloat[i] > max {
		// 		max=AFloat[i]
		// 		maxIndex=i
		// 	}
		// }

		if AFloat[highestIndex] < AFloat[highestIndex-1] {
			// Value is no longer highest. Swap with the highest
			AFloat[highestIndex-1], AFloat[highestIndex] = AFloat[highestIndex], AFloat[highestIndex-1]
		}

		// half the polution
		AFloat[highestIndex] = AFloat[highestIndex] / 2.0
		// update polution
		polution -= AFloat[highestIndex]
		filters++

		if polution <= expected {
			break
		}
	}

	return filters
}

func SolutionThree(A []int) int {
	// 	// write your code in Go 1.4

	sort.Ints(A)
	pollution := []float64{}
	totalPollution := 0.0
	for i := len(A) - 1; i >= 0; i-- {
		pollution = append(pollution, float64(A[i]))
		totalPollution += float64(A[i])
	}
	expected := totalPollution / 2.0

	filters := 0
	for {
		if len(A) >= 2 && pollution[0] < pollution[1] {
			// Value is no longer highest. Swap with the highest
			pollution[0], pollution[1] = pollution[1], pollution[0]
		}

		// half the polution
		pollution[0] = pollution[0] / 2.0
		// update polution
		totalPollution -= pollution[0]
		filters++

		if totalPollution <= expected {
			break
		}
	}

	return filters
}
