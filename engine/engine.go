package engine

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/craigbranscom/engine-sandbox-golang/engine/importer"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Engine struct {
	window   *glfw.Window
	profiler *Profiler
	// renderer *Renderer
	tree *NodeTree
}

func NewEngine() (*Engine, error) {
	//set glfw options
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	//create a window
	windowWidth := os.Getenv("WINDOW_WIDTH")
	windowWidthInt, err := strconv.Atoi(windowWidth)
	if err != nil {
		log.Fatal("Error converting window width to int", err)
	}
	windowHeight := os.Getenv("WINDOW_HEIGHT")
	windowHeightInt, err := strconv.Atoi(windowHeight)
	if err != nil {
		log.Fatal("Error converting window height to int", err)
	}
	applicationName := os.Getenv("APPLICATION_NAME")
	window, err := glfw.CreateWindow(windowWidthInt, windowHeightInt, applicationName, nil, nil)
	if err != nil {
		log.Fatal("Error creating glfw window:", err)
		return nil, err
	}

	//configure key callbacks
	window.SetKeyCallback(keyCallback)

	//set context
	window.MakeContextCurrent()

	//initialize opengl functions
	err = gl.Init()
	if err != nil {
		log.Fatal("Error initializing OpenGL:", err)
		return nil, err
	}

	//initialize node tree
	nodeTree := NewNodeTree()

	//initialize profiler
	targetFPS := float64(60.0)
	profiler := NewProfiler(targetFPS)

	return &Engine{
		window:   window,
		profiler: profiler,
		tree:     nodeTree,
	}, nil
}

func (engine *Engine) Tree() *NodeTree {
	return engine.tree
}

func (engine *Engine) Window() *glfw.Window {
	return engine.window
}

func (engine *Engine) Profiler() *Profiler {
	return engine.profiler
}

func (engine *Engine) BuildNodeTreeFromYaml(data []map[interface{}]interface{}) {
	//recursively build node tree
	rootNode := importer.BuildNodeFromYaml(data[0])
	//set root node
	engine.Tree().SetRootNode(rootNode)
}

func (engine *Engine) StartRenderLoop() {
	//defer cleanup functions in render loop scope
	defer glfw.Terminate()
	defer engine.window.Destroy()

	//import and compile shaders
	program, err := NewCompiledShaderProgram("engine/shaders/vertex.glsl", "engine/shaders/fragment.glsl")
	if err != nil {
		log.Fatal("Error creating compiled shader program:", err)
	}
	defer gl.DeleteProgram(program.ProgramId)

	//declare and generate vertex arrays and vertex buffer objs
	var vertexArrayObject, vertexBufferOject uint32
	gl.GenVertexArrays(1, &vertexArrayObject)
	gl.GenBuffers(1, &vertexBufferOject)

	//bind vertex array and vertex buffer objs
	gl.BindVertexArray(vertexArrayObject)
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBufferOject)

	//get aggregate vertex position data from node tree
	vertices := engine.tree.GetVertexPositionData(engine.tree.RootNode())

	//push vertices into array buffer
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	//configure and enable vertex attributes
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	//define camera and rotation vars
	var cameraXPos, cameraYPos, cameraZPos float32 = 0.0, 0.0, 0.0
	var rotationX, rotationY, rotationZ float32 = 0.0, 0.0, 0.0

	//main render loop
	for !engine.window.ShouldClose() {
		//set frame time start
		frameStartTime := glfw.GetTime()

		//clear frame
		gl.Clear(gl.COLOR_BUFFER_BIT)

		//check for keyboard input
		if engine.window.GetKey(glfw.KeyW) == glfw.Press {
			cameraYPos += 0.01
		}
		if engine.window.GetKey(glfw.KeyS) == glfw.Press {
			cameraYPos -= 0.01
		}
		if engine.window.GetKey(glfw.KeyD) == glfw.Press {
			cameraXPos += 0.01
		}
		if engine.window.GetKey(glfw.KeyA) == glfw.Press {
			cameraXPos -= 0.01
		}
		if engine.window.GetKey(glfw.KeyQ) == glfw.Press {
			cameraZPos += 0.01
		}
		if engine.window.GetKey(glfw.KeyD) == glfw.Press {
			cameraZPos -= 0.01
		}

		if engine.window.GetKey(glfw.KeyU) == glfw.Press {
			rotationX -= 0.01
		}
		if engine.window.GetKey(glfw.KeyI) == glfw.Press {
			rotationX += 0.01
		}
		if engine.window.GetKey(glfw.KeyJ) == glfw.Press {
			rotationY -= 0.01
		}
		if engine.window.GetKey(glfw.KeyK) == glfw.Press {
			rotationY += 0.01
		}
		if engine.window.GetKey(glfw.KeyN) == glfw.Press {
			rotationZ -= 0.01
		}
		if engine.window.GetKey(glfw.KeyM) == glfw.Press {
			rotationZ += 0.01
		}

		//update uniforms
		cameraUniform := gl.GetUniformLocation(program.ProgramId, gl.Str("camera\x00"))
		gl.Uniform3f(cameraUniform, cameraXPos, cameraYPos, cameraZPos)
		rotationUniform := gl.GetUniformLocation(program.ProgramId, gl.Str("rotation\x00"))
		gl.Uniform3f(rotationUniform, rotationX, rotationY, rotationZ)

		//select shader program
		program.UseShaderProgram()

		//bind and draw
		gl.BindVertexArray(vertexArrayObject)
		gl.DrawArrays(gl.TRIANGLES, 0, 3)
		gl.BindVertexArray(0)

		//update profiler
		engine.profiler.UpdateProfiler()
		// fmt.Println("FPS:", engine.profiler.FramesPerSecond())

		// Render the scene
		// app.renderer.Render()

		//swap buffers and poll events
		engine.window.SwapBuffers()
		glfw.PollEvents()

		//delay next frame if necessary to hit target frame time
		//TODO: perform additional compute if needed instead of sleeping
		frameElapsedTime := glfw.GetTime() - frameStartTime
		if frameElapsedTime < engine.profiler.TargetFrameTime() {
			sleepDuration := time.Duration((engine.profiler.TargetFrameTime() - frameElapsedTime) * float64(time.Second))
			time.Sleep(sleepDuration)
		}
	}
}

func keyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		window.SetShouldClose(true)
	}
}
