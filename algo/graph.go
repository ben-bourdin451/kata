package algo

type Graph interface {
	// returns the neighbors of the node alongside the distance to them
	GetNeighbors() ([]Graph, []float64)
}

// Simple implementation of a graph node
type GNode struct {
	Name      string
	Neighbors []*Edge
}

type Edge struct {
	Node *GNode
	Val  float64
}

func (node *GNode) GetNeighbors() ([]Graph, []float64) {
	nodes, dist := []Graph{}, []float64{}

	for _, n := range node.Neighbors {
		nodes = append(nodes, n.Node)
		dist = append(dist, n.Val)
	}

	return nodes, dist
}
