package engine

import (
	importer "github.com/Dappetizer/engine-sandbox-golang/engine/importer"
)

type Engine struct {
	tree *NodeTree
}

func NewEngine() (*Engine, error) {
	nodeTree := NewNodeTree()

	return &Engine{
		tree: nodeTree,
	}, nil
}

func (engine *Engine) Tree() *NodeTree {
	return engine.tree
}

func (engine *Engine) BuildNodeTreeFromYaml(data []map[interface{}]interface{}) {
	//recursively build node tree
	rootNode := importer.BuildNodeFromYaml(data[0])
	//set root node
	engine.Tree().SetRootNode(rootNode)
}
