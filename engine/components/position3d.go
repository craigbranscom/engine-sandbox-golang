package components

type Position3DInterface interface {
	XPos()
	YPos()
	ZPos()
}

type Position3D struct {
	xPos float32
	yPos float32
	zPos float32
}

func NewPosition3DComponent() *Position3D {
	return &Position3D{
		xPos: 0,
		yPos: 0,
		zPos: 0,
	}
}

func BuildPosition3DComponent(xPos float32, yPos float32, zPos float32) *Position3D {
	return &Position3D{
		xPos: xPos,
		yPos: yPos,
		zPos: zPos,
	}
}

func (node *Position3D) XPos() float32 {
	return node.xPos
}

func (node *Position3D) YPos() float32 {
	return node.yPos
}

func (node *Position3D) ZPos() float32 {
	return node.zPos
}

func (node *Position3D) SetPosition3D(xPos float32, yPos float32, zPos float32) {
	node.xPos = xPos
	node.yPos = yPos
	node.zPos = zPos
}
