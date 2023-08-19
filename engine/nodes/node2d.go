package nodes

import (
	components "github.com/Dappetizer/engine-sandbox-golang/engine/nodes/components"
)

type Node2D struct {
	BaseNode
	components.Position2D
}

func NewNode2D(name string, parent Node, xPos float64, yPos float64) *Node2D {
	baseNode := NewBaseNode(name, parent)
	position2d := components.NewPosition2DComponent()

	return &Node2D{
		BaseNode:   *baseNode,
		Position2D: *position2d,
	}
}
