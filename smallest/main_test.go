package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		out  int
	}{
		{
			name: "a",
			in:   []int{1, 3, 6, 4, 1, 2, 6, 4,2, 5, 2, 1, 23,9, 14, 12, 8, 23, 23, 1,2, 243, 5, 6, 7, 8},
			out:  10,
		},
		{
			name: "1",
			in:   []int{1, 3, 6, 4, 1, 2},
			out:  5,
		},

		{
			name: "2",
			in:   []int{1, 3, 6, 4, 1},
			out:  2,
		},

		{
			name: "3",
			in:   []int{1},
			out:  2,
		},

		{
			name: "4",
			in:   []int{5},
			out:  1,
		},

		{
			name: "5",
			in:   []int{},
			out:  1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := Solution(test.in)
			require.Equal(t, test.out, out)
		})
	}
}
