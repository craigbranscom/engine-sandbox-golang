package components

import (
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type OpenGLProcess struct {
}

func (proc *OpenGLProcess) Process() {
	fmt.Println(">>>")
	//TODO: implement opengl rendering
	gl.Clear(gl.COLOR_BUFFER_BIT)
}
