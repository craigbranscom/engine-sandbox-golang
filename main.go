package main

import (
	"log"
	"os"
	"runtime"

	engine "github.com/Dappetizer/engine-sandbox-golang/engine"

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

	//run render loop
	eng.StartRenderLoop()
}
