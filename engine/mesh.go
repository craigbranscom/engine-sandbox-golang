package engine

import (
	"github.com/craigbranscom/engine-sandbox-golang/engine/components"
	"github.com/go-gl/gl/v4.1-core/gl"
)

//TODO: make Mesh a node type?

type Mesh struct {
	Vertices     []components.Position3D
	VertexBuffer uint32
	VertexArray  uint32
}

func NewMesh(vertices []components.Position3D) *Mesh {
	// Create VBO and VAO, upload vertex data, configure attribute pointers, etc.

	//declare and generate vertex arrays and vertex buffer objs
	var vertexArrayObject, vertexBufferOject uint32
	gl.GenVertexArrays(1, &vertexArrayObject)
	gl.GenBuffers(1, &vertexBufferOject)

	//bind vertex array and vertex buffer objs
	gl.BindVertexArray(vertexArrayObject)
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBufferOject)

	return &Mesh{
		Vertices:     vertices,
		VertexBuffer: vertexBufferOject,
		VertexArray:  vertexArrayObject,
	}
}

func (mesh *Mesh) DrawMesh() {
	//bind vao and issue draw command
	gl.BindVertexArray(mesh.VertexArray)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(mesh.Vertices)))
}
