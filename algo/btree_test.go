package algo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var tree1 *BNode = &BNode{
	Val: 3,
	Left: &BNode{
		Val:  6,
		Left: &BNode{Val: 9},
	},
	Right: &BNode{
		Val:  2,
		Left: &BNode{Val: 10},
	},
}

func TestTreeSum(t *testing.T) {
	got := tree1.count()
	require.Equal(t, 30, got, "got %v, expected %v", got, 30)
}
