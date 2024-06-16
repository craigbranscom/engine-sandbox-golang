package engine

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/craigbranscom/engine-sandbox-golang/engine/components"
	nodes "github.com/craigbranscom/engine-sandbox-golang/engine/nodes"
)

type NodeTree struct {
	rootNode  nodes.Node
	nodeCount uint32
	nodeDepth uint32
}

func NewNodeTree() *NodeTree {
	return &NodeTree{
		rootNode:  nil,
		nodeCount: 0,
		nodeDepth: 0,
	}
}

func (tree *NodeTree) RootNode() nodes.Node {
	return tree.rootNode
}

func (tree *NodeTree) NodeCount() uint32 {
	return tree.nodeCount
}

func (tree *NodeTree) NodeDepth() uint32 {
	return tree.nodeDepth
}

func (tree *NodeTree) SetRootNode(node nodes.Node) {
	tree.rootNode = node
}

func (tree *NodeTree) SetNodeCount(count uint32) {
	tree.nodeCount = count
}

func (tree *NodeTree) SetNodeDepth(depth uint32) {
	tree.nodeDepth = depth
}

func (tree *NodeTree) PrintNodeTree(node nodes.Node, indent int) {
	indents := strings.Repeat("\t", indent)
	fmt.Println(indents, node.Name(), ":", node.Id())

	children := node.Children()
	for _, child := range children {
		tree.PrintNodeTree(child, indent+1)
	}
}

func (tree *NodeTree) GetVertexPositionData(node nodes.Node) []float32 {
	var vertices []float32

	//TODO: use reflection to check for components with vertex position data
	// hasComponent := ContainsComponent(node)
	// fmt.Println(">>>", hasComponent)

	if node.Name() == "Triangle" {
		triangle := node.(*nodes.Triangle3D)
		pos := triangle.VertexPositions()
		floatArr := []float32{pos[0].XPos(), pos[0].YPos(), pos[0].ZPos(), pos[1].XPos(), pos[1].YPos(), pos[1].ZPos(), pos[2].XPos(), pos[2].YPos(), pos[2].ZPos()}
		vertices = append(floatArr)
	}

	children := node.Children()
	for _, child := range children {
		vertices = append(tree.GetVertexPositionData(child))
	}

	return vertices
}

func ContainsComponent(obj interface{}) bool {
	componentType := reflect.TypeOf((*components.Position3DInterface)(nil)).Elem()
	objectType := reflect.TypeOf(obj)
	return objectType.Implements(componentType)
}
