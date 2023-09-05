package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	engine "github.com/Dappetizer/engine-sandbox-golang/engine"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"

	// "github.com/go-gl/mathgl/mgl32"

	"github.com/go-yaml/yaml"
	"github.com/joho/godotenv"
	// _ "github.com/lib/pq"
	// amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	//lock thread since opengl isnt thread-safe
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

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

	//TODO: run render loop

	//initialize glfw
	// err = glfw.Init()
	// if err != nil {
	// 	log.Fatal("Error initializing glfw:", err)
	// }
	// defer glfw.Terminate()

	//set glfw options
	// glfw.WindowHint(glfw.ContextVersionMajor, 4)
	// glfw.WindowHint(glfw.ContextVersionMinor, 1)
	// glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	// glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	//create a window
	// windowWidth := os.Getenv("WINDOW_WIDTH")
	// windowWidthInt, err := strconv.Atoi(windowWidth)
	// if err != nil {
	// 	log.Fatal("Error converting window width to int", err)
	// }
	// windowHeight := os.Getenv("WINDOW_HEIGHT")
	// windowHeightInt, err := strconv.Atoi(windowHeight)
	// if err != nil {
	// 	log.Fatal("Error converting window height to int", err)
	// }
	// applicationName := os.Getenv("APPLICATION_NAME")
	// window, err := glfw.CreateWindow(windowWidthInt, windowHeightInt, applicationName, nil, nil)
	// if err != nil {
	// 	log.Fatal("Error creating glfw window:", err)
	// 	return
	// }
	// defer window.Destroy()

	//configure window
	// window.SetKeyCallback(keyCallback)

	//set context
	// window.MakeContextCurrent()

	//initialize opengl functions
	err = gl.Init()
	if err != nil {
		log.Fatal("Error initializing OpenGL:", err)
		return
	}

	program, err := engine.NewCompiledShaderProgram("engine/shaders/vertex.glsl", "engine/shaders/fragment.glsl")
	if err != nil {
		log.Fatal("Error creating compiled shader program:", err)
	}

	//load shaders from file
	// vertexShaderSource, err := importer.LoadShaderSourceFromFile("engine/shaders/vertex.glsl")
	// if err != nil {
	// 	log.Fatal("Error loading vertex shader from file:", err)
	// 	return
	// }
	// fragmentShaderSource, err := importer.LoadShaderSourceFromFile("engine/shaders/fragment.glsl")
	// if err != nil {
	// 	log.Fatal("Error loading fragment shader from file:", err)
	// 	return
	// }

	//compile and link shaders
	// program, err := newProgram(vertexShaderSource, fragmentShaderSource)
	// if err != nil {
	// 	log.Fatal("Error creating shader program:", err)
	// 	return
	// }
	// defer gl.DeleteProgram(program)

	//declare and generate vertex arrays and vertex buffer objs
	var vertexArrayObject, vertexBufferOject uint32
	gl.GenVertexArrays(1, &vertexArrayObject)
	gl.GenBuffers(1, &vertexBufferOject)

	//bind vertex array and vertex buffer objs
	gl.BindVertexArray(vertexArrayObject)
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBufferOject)

	//get aggregate vertex position data from node tree
	vertices := eng.Tree().GetVertexPositionData(eng.Tree().RootNode())

	//push vertices into array buffer
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	//configure and enable vertex attributes
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	var cameraXPos, cameraYPos, cameraZPos float32 = 0.0, 0.0, 0.0
	var rotationX, rotationY, rotationZ float32 = 0.0, 0.0, 0.0

	//render loop
	for !eng.Window().ShouldClose() {

		//clear frame
		gl.Clear(gl.COLOR_BUFFER_BIT)

		//check for keyboard input
		if eng.Window().GetKey(glfw.KeyW) == glfw.Press {
			cameraYPos += 0.01
		}
		if eng.Window().GetKey(glfw.KeyS) == glfw.Press {
			cameraYPos -= 0.01
		}
		if eng.Window().GetKey(glfw.KeyD) == glfw.Press {
			cameraXPos += 0.01
		}
		if eng.Window().GetKey(glfw.KeyA) == glfw.Press {
			cameraXPos -= 0.01
		}
		if eng.Window().GetKey(glfw.KeyQ) == glfw.Press {
			cameraZPos += 0.01
		}
		if eng.Window().GetKey(glfw.KeyD) == glfw.Press {
			cameraZPos -= 0.01
		}

		if eng.Window().GetKey(glfw.KeyU) == glfw.Press {
			rotationX -= 0.01
		}
		if eng.Window().GetKey(glfw.KeyI) == glfw.Press {
			rotationX += 0.01
		}
		if eng.Window().GetKey(glfw.KeyJ) == glfw.Press {
			rotationY -= 0.01
		}
		if eng.Window().GetKey(glfw.KeyK) == glfw.Press {
			rotationY += 0.01
		}
		if eng.Window().GetKey(glfw.KeyN) == glfw.Press {
			rotationZ -= 0.01
		}
		if eng.Window().GetKey(glfw.KeyM) == glfw.Press {
			rotationZ += 0.01
		}

		//update uniforms
		cameraUniform := gl.GetUniformLocation(program.ProgramId, gl.Str("camera\x00"))
		gl.Uniform3f(cameraUniform, cameraXPos, cameraYPos, cameraZPos)

		rotationUniform := gl.GetUniformLocation(program.ProgramId, gl.Str("rotation\x00"))
		gl.Uniform3f(rotationUniform, rotationX, rotationY, rotationZ)

		//select shader program
		// gl.UseProgram(program.ProgramId)
		program.UseShaderProgram()

		//bind and draw
		gl.BindVertexArray(vertexArrayObject)
		gl.DrawArrays(gl.TRIANGLES, 0, 3)
		gl.BindVertexArray(0)

		//update profiler
		eng.Profiler.UpdateProfiler()
		fmt.Println("FPS:", eng.Profiler.FramesPerSecond())

		//swap buffers and poll events
		eng.Window().SwapBuffers()
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
