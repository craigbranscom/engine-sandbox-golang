package nodes

import components "github.com/Dappetizer/engine-sandbox-golang/engine/nodes/components"

type Line2D struct {
	Node2D
	Points []components.Position2D
	Width  int
}

func NewLine2D(name string, parent Node, xPos float64, yPos float64) *Line2D {
	node2d := NewNode2D(name, parent, xPos, yPos)

	return &Line2D{
		Node2D: *node2d,
		Points: nil,
		Width:  0,
	}
}
