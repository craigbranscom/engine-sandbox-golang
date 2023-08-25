package engine

import (
	importer "github.com/Dappetizer/engine-sandbox-golang/engine/importer"
	// gl "github.com/go-gl/gl/v4.1-core/gl"
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

func (engine *Engine) StartRenderLoop() {
	var renderLoopShouldClose bool
	for !renderLoopShouldClose {
		// engine.tree.RootNode().Process()
		//TODO: swap buffers, poll events
	}
}
