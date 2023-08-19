package nodes

import (
	"github.com/google/uuid"
)

type Node interface {
	Id() uuid.UUID
	Name() string
	Parent() Node
	Children() []Node

	SetName(name string)
	AppendChild(childNode Node)
}

type BaseNode struct {
	id       uuid.UUID
	name     string
	parent   Node
	children []Node
}

func NewBaseNode() *BaseNode {
	//generate unique node id
	nodeId := uuid.New()

	return &BaseNode{
		id:       nodeId,
		name:     "",
		parent:   nil,
		children: nil,
	}
}

func BuildBaseNode(name string, parent Node) *BaseNode {
	baseNode := NewBaseNode()
	baseNode.SetName(name)

	if parent != nil {
		//register new child on parent node
		parent.AppendChild(baseNode)
	}

	return baseNode
}

func (node *BaseNode) Id() uuid.UUID {
	return node.id
}

func (node *BaseNode) Name() string {
	return node.name
}

func (node *BaseNode) Parent() Node {
	return node.parent
}

func (node *BaseNode) Children() []Node {
	return node.children
}

func (node *BaseNode) SetName(name string) {
	node.name = name
}

func (node *BaseNode) AppendChild(childNode Node) {
	node.children = append(node.Children(), childNode)
}
