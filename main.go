package main

import (
	"log"

	engine "github.com/Dappetizer/engine-sandbox-golang/engine"
	"github.com/Dappetizer/engine-sandbox-golang/engine/nodes"
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

	//create engine instance
	eng, err := engine.NewEngine()
	if err != nil {
		log.Fatal("Error creating new engine", err)
	}
	_ = eng
	//create root node
	root := nodes.NewBaseNode("Root", nil)
	if err != nil {
		log.Fatal("Error creating node", err)
	}
	eng.Tree().SetRootNode(root)

	//create example scene
	sub2d := nodes.NewNode2D("Sub2D", root)
	if err != nil {
		log.Fatal("Error creating node", err)
	}
	_ = sub2d
	sub3d := nodes.NewNode3D("Sub3D", root)
	if err != nil {
		log.Fatal("Error creating node", err)
	}
	_ = sub3d

	//print node tree
	eng.Tree().PrintNodeTree(eng.Tree().RootNode(), 0)

	log.Println("Engine startup complete")

	for {
	}
}
