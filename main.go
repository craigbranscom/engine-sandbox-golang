package main

import (
	"log"
	"os"

	engine "github.com/Dappetizer/engine-sandbox-golang/engine"
	"github.com/go-yaml/yaml"
	"github.com/joho/godotenv"
	// _ "github.com/lib/pq"
	// amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	//start context
	// ctx := context.Background()

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
	_ = eng
	eng.BuildNodeTreeFromYaml(nodeTreeYaml)
	//create root node
	// root := nodes.NewBaseNode("Root", nil)
	// if err != nil {
	// 	log.Fatal("Error creating node", err)
	// }
	// eng.Tree().SetRootNode(root)

	//print node tree
	eng.Tree().PrintNodeTree(eng.Tree().RootNode(), 0)

	log.Println("Engine startup complete")

	for {
	}
}
