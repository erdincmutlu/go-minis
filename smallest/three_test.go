package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolutionThree(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		out  int
	}{
		{
			name: "1",
			in:   []int{10, 10},
			out:  2,
		},
		{
			name: "2",
			in:   []int{5,19,8,1},
			out:  3,
		},
		{
			name: "3",
			in:   []int{3,0,5},
			out:  2,
		},
		{
			name: "4",
			in:   []int{5},
			out:  1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := SolutionThree(test.in)
			require.Equal(t, test.out, out)
		})
	}
}
