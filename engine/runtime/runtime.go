package runtime

import (
	"log"
	"os"
	"strconv"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type Runtime struct {
	window *glfw.Window
	// profiler *Profiler
	// renderer *Renderer
	// tree *NodeTree
}

func NewRuntime() (*Runtime, error) {
	//initialize glfw
	err := glfw.Init()
	if err != nil {
		log.Fatal("Error initializing glfw:", err)
	}

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

	return &Runtime{
		window: window,
	}, nil
}

func keyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		window.SetShouldClose(true)
	}
}
