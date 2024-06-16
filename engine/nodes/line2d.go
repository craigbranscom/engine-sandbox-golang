package nodes

import (
	components "github.com/craigbranscom/engine-sandbox-golang/engine/components"
)

type Line2D struct {
	Node2D
	points []components.Position2D
	width  uint
}

func NewLine2D() *Line2D {
	node2d := NewNode2D()
	return &Line2D{
		Node2D: *node2d,
		points: nil,
		width:  0,
	}
}

func BuildLine2D(node2d Node2D, points []components.Position2D, width uint) *Line2D {
	return &Line2D{
		Node2D: node2d,
		points: points,
		width:  width,
	}
}

func (line *Line2D) Points() []components.Position2D {
	return line.points
}

func (line *Line2D) Width() uint {
	return line.width
}
