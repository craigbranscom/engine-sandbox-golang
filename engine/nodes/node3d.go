package nodes

type Position3D struct {
	xPos float64
	yPos float64
	zPos float64
}

// type INode3D interface {
// 	XPos()
// 	YPos()
//  ZPos()
// 	SetPosition3D(xPos int64, yPos int64, zPos int64)
// }

type Node3D struct {
	BaseNode
	Position3D
}

func NewNode3D(name string, parent Node, xPos float64, yPos float64, zPos float64) *Node3D {
	baseNode := NewBaseNode(name, parent)

	return &Node3D{
		BaseNode: *baseNode,
		Position3D: Position3D{
			xPos: xPos,
			yPos: yPos,
			zPos: zPos,
		},
	}
}

func (node *Position3D) XPos() float64 {
	return node.xPos
}

func (node *Position3D) YPos() float64 {
	return node.yPos
}

func (node *Position3D) ZPos() float64 {
	return node.zPos
}

func (node *Position3D) SetPosition3D(xPos float64, yPos float64, zPos float64) {
	node.xPos = xPos
	node.yPos = yPos
	node.zPos = zPos
}
