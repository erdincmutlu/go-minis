package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolutionOne(t *testing.T) {
	tests := []struct {
		name string
		in   string
		k    int
		out  string
	}{
		{
			name: "a",
			in:   "Codility we test coders",
			k:    14,
			out:  "Codility we",
		},

		{
			name: "b",
			in:   "why not",
			k:    100,
			out:  "why not",
		},
		{
			name: "3",
			in:   "",
			k:    100,
			out:  "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := SolutionOne(test.in, test.k)
			require.Equal(t, test.out, out)
		})
	}
}
