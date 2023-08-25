package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	engine "github.com/Dappetizer/engine-sandbox-golang/engine"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-yaml/yaml"
	"github.com/joho/godotenv"
	// _ "github.com/lib/pq"
	// amqp "github.com/rabbitmq/amqp091-go"
)

const WINDOW_HEIGHT = 600
const WINDOW_WIDTH = 800

var vertices = []float32{
	-0.3, -0.3, 0.0,
	0.3, -0.3, 0.0,
	0.0, 0.5, 0.0,
}

var vertexShaderSource = `
	#version 410 core
	layout (location = 0) in vec3 aPos;
	uniform vec2 camera;
	uniform float rotationX;
	uniform float rotationY;
	uniform float rotationZ;
	mat4 rotateX(float angle) {
		float c = cos(angle);
		float s = sin(angle);
		return mat4(
			1.0, 0.0, 0.0, 0.0,
			0.0, c, -s, 0.0,
			0.0, s, c, 0.0,
			0.0, 0.0, 0.0, 1.0
		);
	}
	mat4 rotateY(float angle) {
		float c = cos(angle);
		float s = sin(angle);
		return mat4(
			c, 0.0, s, 0.0,
			0.0, 1.0, 0.0, 0.0,
			-s, 0.0, c, 0.0,
			0.0, 0.0, 0.0, 1.0
		);
	}
	mat4 rotateZ(float angle) {
		float c = cos(angle);
		float s = sin(angle);
		return mat4(
			c, -s, 0.0, 0.0,
			s, c, 0.0, 0.0,
			0.0, 0.0, 1.0, 0.0,
			0.0, 0.0, 0.0, 1.0
		);
	}
	void main()
	{
		mat4 rotationMat = rotateX(rotationX) * rotateY(rotationY) * rotateZ(rotationZ);
		vec4 rotatedPos = rotationMat * vec4(aPos, 1.0);
		gl_Position = vec4(rotatedPos.x - camera.x, rotatedPos.y - camera.y, rotatedPos.z, 1.0);
	}`

var fragmentShaderSource = `
	#version 410 core
	out vec4 FragColor;
	void main()
	{
		FragColor = vec4(1.0, 0.5, 0.2, 1.0);
	}`

func main() {
	//lock thread
	runtime.LockOSThread()

	//load project env
	err := godotenv.Load("env.yaml")
	if err != nil {
		log.Fatal("Error loading env file", err)
	}

	//load and parse yaml scene file
	yamlFile, err := os.ReadFile("scene.yaml")
	if err != nil {
		log.Fatal("Error parsing yaml", err)
	}
	var nodeTreeYaml []map[interface{}]interface{}
	unmarshalErr := yaml.Unmarshal(yamlFile, &nodeTreeYaml)
	if unmarshalErr != nil {
		log.Fatal("Error unmarshalling yaml", unmarshalErr)
	}

	//create engine instance
	eng, err := engine.NewEngine()
	if err != nil {
		log.Fatal("Error creating new engine", err)
	}

	//build node tree from yaml file
	eng.BuildNodeTreeFromYaml(nodeTreeYaml)

	//print node tree
	eng.Tree().PrintNodeTree(eng.Tree().RootNode(), 0)

	//initialize glfw
	err = glfw.Init()
	if err != nil {
		log.Fatal("Error initializing glfw:", err)
	}
	defer glfw.Terminate()

	//set glfw options
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	//create a window
	window, err := glfw.CreateWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "OpenGL Triangle", nil, nil)
	if err != nil {
		log.Fatal("Error creating glfw window:", err)
		return
	}
	defer window.Destroy()

	//configure window
	window.SetKeyCallback(keyCallback)

	//set context
	window.MakeContextCurrent()

	//initialize opengl functions
	err = gl.Init()
	if err != nil {
		log.Fatal("Error initializing OpenGL:", err)
		return
	}

	//compile and link shaders
	program, err := newProgram(vertexShaderSource, fragmentShaderSource)
	if err != nil {
		log.Fatal("Error creating shader program:", err)
		return
	}
	defer gl.DeleteProgram(program)

	//create and bind vertex array and buffer objs
	var vao, vbo uint32
	gl.GenVertexArrays(1, &vao)
	gl.GenBuffers(1, &vbo)

	gl.BindVertexArray(vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	//configure vertex attributes
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	log.Println("Engine startup complete")

	log.Println("Starting render loop")

	var cameraXPos, cameraYPos float32 = 0.0, 0.0
	var rotationX, rotationY, rotationZ float32 = 0.0, 0.0, 0.0

	//render loop
	for !window.ShouldClose() {

		//clear frame
		gl.Clear(gl.COLOR_BUFFER_BIT)

		//check for keyboard input
		if window.GetKey(glfw.KeyW) == glfw.Press {
			cameraYPos += 0.01
		}
		if window.GetKey(glfw.KeyS) == glfw.Press {
			cameraYPos -= 0.01
		}
		if window.GetKey(glfw.KeyA) == glfw.Press {
			cameraXPos -= 0.01
		}
		if window.GetKey(glfw.KeyD) == glfw.Press {
			cameraXPos += 0.01
		}

		if window.GetKey(glfw.KeyU) == glfw.Press {
			rotationX -= 0.01
		}
		if window.GetKey(glfw.KeyI) == glfw.Press {
			rotationX += 0.01
		}
		if window.GetKey(glfw.KeyJ) == glfw.Press {
			rotationY -= 0.01
		}
		if window.GetKey(glfw.KeyK) == glfw.Press {
			rotationY += 0.01
		}
		if window.GetKey(glfw.KeyN) == glfw.Press {
			rotationZ -= 0.01
		}
		if window.GetKey(glfw.KeyM) == glfw.Press {
			rotationZ += 0.01
		}

		//update uniforms
		rotationUniformX := gl.GetUniformLocation(program, gl.Str("rotationX\x00"))
		rotationUniformY := gl.GetUniformLocation(program, gl.Str("rotationY\x00"))
		rotationUniformZ := gl.GetUniformLocation(program, gl.Str("rotationZ\x00"))
		gl.Uniform1f(rotationUniformX, rotationX)
		gl.Uniform1f(rotationUniformY, rotationY)
		gl.Uniform1f(rotationUniformZ, rotationZ)

		cameraUniform := gl.GetUniformLocation(program, gl.Str("camera\x00"))
		gl.Uniform2f(cameraUniform, cameraXPos, cameraYPos)

		//select shader program
		gl.UseProgram(program)

		//bind and draw
		gl.BindVertexArray(vao)
		gl.DrawArrays(gl.TRIANGLES, 0, 3)
		gl.BindVertexArray(0)

		//swap buffers and poll events
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func keyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		window.SetShouldClose(true)
	}
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
