package engine

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
