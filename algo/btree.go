package algo

import "fmt"

type BNode struct {
	Val         interface{}
	Left, Right *BNode
}

// Assumes int node type
func (n *BNode) count() int {
	sum := 0
	if n.Right != nil {
		sum += n.Right.count()
	}
	if n.Left != nil {
		sum += n.Left.count()
	}
	if n.Val != -1 {
		sum += n.Val.(int)
	}
	return sum
}

func (n *BNode) Print() {
	fmt.Printf("%v, ", n.Val)
	n.Left.Print()
	n.Right.Print()
}
