package engine

import (
	"log"
	"os"
	"strconv"

	importer "github.com/Dappetizer/engine-sandbox-golang/engine/importer"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Engine struct {
	window   *glfw.Window
	Profiler *Profiler
	// renderer *Renderer
	tree *NodeTree
}

func NewEngine() (*Engine, error) {
	//TODO: initialize OpenGL, create window, and set up key callbacks

	//initialize glfw
	err := glfw.Init()
	if err != nil {
		log.Fatal("Error initializing glfw:", err)
	}
	// defer glfw.Terminate()

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
	// defer window.Destroy()

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
	profiler := NewProfiler()

	return &Engine{
		window:   window,
		Profiler: profiler,
		tree:     nodeTree,
	}, nil
}

func (engine *Engine) Tree() *NodeTree {
	return engine.tree
}

func (engine *Engine) Window() *glfw.Window {
	return engine.window
}

func (engine *Engine) BuildNodeTreeFromYaml(data []map[interface{}]interface{}) {
	//recursively build node tree
	rootNode := importer.BuildNodeFromYaml(data[0])
	//set root node
	engine.Tree().SetRootNode(rootNode)
}

func (engine *Engine) StartRenderLoop() {
	//main render loop
	// for !engine.window.ShouldClose() {
	// 	// Handle user input, update game logic, etc.
	//     // ...

	//     // Render the scene
	//     app.renderer.Render()

	//     // Swap buffers and poll events
	//     app.window.SwapBuffers()
	//     glfw.PollEvents()
	// }
}

func keyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		window.SetShouldClose(true)
	}
}
