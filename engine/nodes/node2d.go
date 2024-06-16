package nodes

import (
	components "github.com/craigbranscom/engine-sandbox-golang/engine/components"
)

type Node2D struct {
	BaseNode
	components.Position2D
}

func NewNode2D() *Node2D {
	baseNode := NewBaseNode()
	position2d := components.NewPosition2DComponent()

	return &Node2D{
		BaseNode:   *baseNode,
		Position2D: *position2d,
	}
}

func BuildNode2D(baseNode BaseNode, position2d components.Position2D) *Node2D {
	return &Node2D{
		BaseNode:   baseNode,
		Position2D: position2d,
	}
}
