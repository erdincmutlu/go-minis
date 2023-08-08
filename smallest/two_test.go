package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolutionTwo(t *testing.T) {
	tests := []struct {
		name string
		in1  []int
		in2  []int
		out  int
	}{
		{
			name: "1",
			in1:  []int{1, 4, 1},
			in2:  []int{1, 5, 1},
			out:  2,
		},
		{
			name: "2",
			in1:  []int{4, 4, 2, 4},
			in2:  []int{5, 5, 2, 5},
			out:  3,
		},
		{
			name: "3",
			in1:  []int{2, 3, 4, 2},
			in2:  []int{2, 5, 7, 2},
			out:  2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := SolutionTwo(test.in1, test.in2)
			require.Equal(t, test.out, out)
		})
	}
}
