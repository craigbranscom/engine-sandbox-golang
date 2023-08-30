package nodes

import (
	components "github.com/Dappetizer/engine-sandbox-golang/engine/components"
)

type Triangle3D struct {
	Node3D
	vertexPositions []components.Position3D
}

func NewTriangle3D() *Triangle3D {
	node3d := NewNode3D()
	return &Triangle3D{
		Node3D:          *node3d,
		vertexPositions: nil,
	}
}

func BuildTriangle3D(node3d Node3D, vertices []components.Position3D) *Triangle3D {
	return &Triangle3D{
		Node3D:          node3d,
		vertexPositions: vertices,
	}
}

func (tri *Triangle3D) VertexPositions() []components.Position3D {
	return tri.vertexPositions
}
