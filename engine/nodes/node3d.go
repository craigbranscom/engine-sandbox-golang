package nodes

type Position3D struct {
	x int64
	y int64
	z int64
}

// type INode3D interface {
// 	X()
// 	Y()
//  Z()
// 	SetPosition3D(x int64, y int64, z int64)
// }

type Node3D struct {
	BaseNode
	Position3D
}

func NewNode3D(name string, parent Node) *Node3D {
	baseNode := NewBaseNode(name, parent)

	return &Node3D{
		BaseNode: *baseNode,
		Position3D: Position3D{
			x: 0,
			y: 0,
			z: 0,
		},
	}
}

func (node *Position3D) X() int64 {
	return node.x
}

func (node *Position3D) Y() int64 {
	return node.y
}

func (node *Position3D) Z() int64 {
	return node.z
}

func (node *Position3D) SetPosition3D(x int64, y int64, z int64) {
	node.x = x
	node.y = y
	node.z = z
}
