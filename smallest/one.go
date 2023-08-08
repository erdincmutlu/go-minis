package main

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")
import "strings"

func SolutionOne(message string, K int) string {
	// write your code in Go 1.4

	words := strings.Fields(message)
	total := 0
	out := ""
	for _, w := range words {
		addition := 1 + len(w)
		if out == "" {
			addition--
		}

		if total+addition <= K {
			if out != "" {
				out += " "
			}
			out += w
			total += addition
		} else {
			break
		}
	}
	return out

}
