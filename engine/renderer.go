package engine

type Renderer struct {
	ShaderProgram *CompiledShaderProgram
	Objects       []*Mesh
}

func NewRenderer(shader *CompiledShaderProgram) *Renderer {
	// Initialize the renderer with the shader program
	var objects []*Mesh

	return &Renderer{
		ShaderProgram: shader,
		Objects:       objects,
	}
}

func (renderer *Renderer) AddObject(mesh *Mesh) {
	//add 3d object to be rendered
	renderer.Objects = append(renderer.Objects, mesh)
}

func (renderer *Renderer) RenderObjects() {
	// Set up view and projection matrices, render objects, etc.
}
