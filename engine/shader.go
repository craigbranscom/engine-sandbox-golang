package engine

import (
	"fmt"
	"log"

	"github.com/craigbranscom/engine-sandbox-golang/engine/importer"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type CompiledShaderProgram struct {
	ProgramId uint32
}

func NewCompiledShaderProgram(vertexShaderPath string, fragmentShaderPath string) (*CompiledShaderProgram, error) {
	//TODO: manage uniform variables

	//load shaders from file
	vertexShaderSource, err := importer.LoadShaderSourceFromFile(vertexShaderPath)
	if err != nil {
		log.Fatal("Error loading vertex shader from file:", vertexShaderPath, err)
		return nil, err
	}
	fragmentShaderSource, err := importer.LoadShaderSourceFromFile(fragmentShaderPath)
	if err != nil {
		log.Fatal("Error loading fragment shader from file:", fragmentShaderPath, err)
		return nil, err
	}

	//compile and link shaders
	program, err := newProgram(vertexShaderSource, fragmentShaderSource)
	if err != nil {
		log.Fatal("Error creating shader program:", err)
		return nil, err
	}

	return &CompiledShaderProgram{
		ProgramId: program,
	}, nil
}

func (program *CompiledShaderProgram) UseShaderProgram() {
	gl.UseProgram(program.ProgramId)
}

func newProgram(vertexSource, fragmentSource string) (uint32, error) {
	vertexShader, err := compileShader(vertexSource, gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}

	fragmentShader, err := compileShader(fragmentSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return 0, err
	}

	program := gl.CreateProgram()
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)
		log := make([]byte, logLength)
		gl.GetProgramInfoLog(program, logLength, nil, &log[0])
		return 0, fmt.Errorf("Error linking program: %v", string(log))
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return program, nil
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)
	csource, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csource, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)
		log := make([]byte, logLength)
		gl.GetShaderInfoLog(shader, logLength, nil, &log[0])
		return 0, fmt.Errorf("Error compiling %v shader: %v", shaderType, string(log))
	}

	return shader, nil
}
