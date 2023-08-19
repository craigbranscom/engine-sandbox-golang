package nodes

type Position2D struct {
	x int64
	y int64
}

// type INode2D interface {
// 	X()
// 	Y()
// 	SetPosition2D(x int64, y int64)
// }

type Node2D struct {
	BaseNode
	Position2D
}

func NewNode2D(name string, parent Node) *Node2D {
	baseNode := NewBaseNode(name, parent)

	return &Node2D{
		BaseNode: *baseNode,
		Position2D: Position2D{
			x: 0,
			y: 0,
		},
	}
}

func (node *Position2D) X() int64 {
	return node.x
}

func (node *Position2D) Y() int64 {
	return node.y
}

func (node *Position2D) SetPosition2D(x int64, y int64) {
	node.x = x
	node.y = y
}
