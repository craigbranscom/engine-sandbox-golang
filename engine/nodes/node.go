package nodes

import (
	"github.com/google/uuid"
)

type Node interface {
	Id() uuid.UUID
	Name() string
	Parent() Node
	Children() []Node

	RegisterChild(childNode Node)
}

type BaseNode struct {
	id       uuid.UUID
	name     string
	parent   Node
	children []Node
}

func NewBaseNode(name string, parent Node) *BaseNode {
	//generate unique node id
	nodeId := uuid.New()

	baseNode := &BaseNode{
		id:       nodeId,
		name:     name,
		parent:   parent,
		children: nil,
	}

	if parent != nil {
		//register new child on parent node
		parent.RegisterChild(baseNode)
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

func (node *BaseNode) RegisterChild(childNode Node) {
	//TODO: validate against duplicate children
	//TODO: validate not circular parent-child relationship

	node.children = append(node.Children(), childNode)
}
