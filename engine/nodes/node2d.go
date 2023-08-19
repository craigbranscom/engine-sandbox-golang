package nodes

type Position2D struct {
	xPos float64
	yPos float64
}

// type INode2D interface {
// 	XPos()
// 	YPos()
// 	SetPosition2D(x int64, y int64)
// }

type Node2D struct {
	BaseNode
	Position2D
}

func NewNode2D(name string, parent Node, xPos float64, yPos float64) *Node2D {
	baseNode := NewBaseNode(name, parent)

	return &Node2D{
		BaseNode: *baseNode,
		Position2D: Position2D{
			xPos: xPos,
			yPos: yPos,
		},
	}
}

func (node *Position2D) XPos() float64 {
	return node.xPos
}

func (node *Position2D) YPos() float64 {
	return node.yPos
}

func (node *Position2D) SetPosition2D(xPos float64, yPos float64) {
	node.xPos = xPos
	node.yPos = yPos
}
