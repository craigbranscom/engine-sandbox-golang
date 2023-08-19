package nodes

type Line2D struct {
	Node2D
}

func NewLine2D(name string, parent Node, xPos float64, yPos float64) *Line2D {
	node2d := NewNode2D(name, parent, xPos, yPos)

	return &Line2D{
		Node2D: *node2d,
	}
}
