package position2d

type Position2D struct {
	xPos float64
	yPos float64
}

func NewPosition2DComponent() *Position2D {
	return &Position2D{
		xPos: 0,
		yPos: 0,
	}
}

func BuildPosition2DComponent(xPos float64, yPos float64) *Position2D {
	return &Position2D{
		xPos: xPos,
		yPos: yPos,
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
