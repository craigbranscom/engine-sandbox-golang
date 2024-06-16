package nodes

import (
	components "github.com/craigbranscom/engine-sandbox-golang/engine/components"
)

type Node3D struct {
	BaseNode
	components.Position3D
}

func NewNode3D() *Node3D {
	baseNode := NewBaseNode()
	position3d := components.NewPosition3DComponent()

	return &Node3D{
		BaseNode:   *baseNode,
		Position3D: *position3d,
	}
}

func BuildNode3D(baseNode BaseNode, position3d components.Position3D) *Node3D {
	return &Node3D{
		BaseNode:   baseNode,
		Position3D: position3d,
	}
}
