package cci

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Bubble
func TestBubbleSort(t *testing.T) {
	cases := []struct {
		in   []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{
			[]int{6, 3, 4, 2, 5, 1},
			[]int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, c := range cases {
		in := make([]int, len(c.in))
		copy(in, c.in)
		bubbleSort(in)
		require.Equal(t, c.want, in, "Sort it out!")
	}
}
