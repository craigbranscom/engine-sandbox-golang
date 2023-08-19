package engine

import (
	"fmt"
	"strings"

	nodes "github.com/Dappetizer/engine-sandbox-golang/engine/nodes"
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
