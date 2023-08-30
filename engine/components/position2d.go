package components

type Position2D struct {
	xPos float32
	yPos float32
}

func NewPosition2DComponent() *Position2D {
	return &Position2D{
		xPos: 0,
		yPos: 0,
	}
}

func BuildPosition2DComponent(xPos float32, yPos float32) *Position2D {
	return &Position2D{
		xPos: xPos,
		yPos: yPos,
	}
}

func (node *Position2D) XPos() float32 {
	return node.xPos
}

func (node *Position2D) YPos() float32 {
	return node.yPos
}

func (node *Position2D) SetPosition2D(xPos float32, yPos float32) {
	node.xPos = xPos
	node.yPos = yPos
}
