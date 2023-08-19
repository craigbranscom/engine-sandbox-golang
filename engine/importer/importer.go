package importer

import (
	"log"
	"strconv"

	nodes "github.com/Dappetizer/engine-sandbox-golang/engine/nodes"
)

func BuildNodeFromYaml(data map[interface{}]interface{}) nodes.Node {
	switch data["type"] {
	case "BaseNode":
		node := nodes.NewBaseNode(data["name"].(string), nil)
		children := data["children"].([]interface{})
		for _, childData := range children {
			child := BuildNodeFromYaml(childData.(map[interface{}]interface{}))
			node.AppendChild(child)
		}
		return node
	case "Node2D":
		xPos, xErr := strconv.ParseFloat(data["xPos"].(string), 64)
		if xErr != nil {
			log.Fatalf("Error parsing x value: %v", xErr)
		}
		yPos, yErr := strconv.ParseFloat(data["yPos"].(string), 64)
		if yErr != nil {
			log.Fatalf("Error parsing y value: %v", yErr)
		}
		node := nodes.NewNode2D(data["name"].(string), nil, xPos, yPos)
		children := data["children"].([]interface{})
		for _, childData := range children {
			child := BuildNodeFromYaml(childData.(map[interface{}]interface{}))
			node.AppendChild(child)
		}
		return node
	case "Node3D":
		xPos, xErr := strconv.ParseFloat(data["xPos"].(string), 64)
		if xErr != nil {
			log.Fatalf("Error parsing x value: %v", xErr)
		}
		yPos, yErr := strconv.ParseFloat(data["yPos"].(string), 64)
		if yErr != nil {
			log.Fatalf("Error parsing y value: %v", yErr)
		}
		zPos, zErr := strconv.ParseFloat(data["zPos"].(string), 64)
		if zErr != nil {
			log.Fatalf("Error parsing z value: %v", zErr)
		}
		node := nodes.NewNode3D(data["name"].(string), nil, xPos, yPos, zPos)
		children := data["children"].([]interface{})
		for _, childData := range children {
			child := BuildNodeFromYaml(childData.(map[interface{}]interface{}))
			node.AppendChild(child)
		}
		return node
	default:
		log.Fatalf("Unknown node type: %s", data["type"])
		return nil
	}
}
