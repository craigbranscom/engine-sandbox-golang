package importer

import (
	"log"
	"os"
	"strconv"

	components "github.com/Dappetizer/engine-sandbox-golang/engine/components"
	nodes "github.com/Dappetizer/engine-sandbox-golang/engine/nodes"
)

func BuildNodeFromYaml(data map[interface{}]interface{}) nodes.Node {
	//TODO: get parent data
	switch data["type"] {
	case "BaseNode":
		node := nodes.BuildBaseNode(data["name"].(string), nil)
		BuildNodeChildrenFromYaml(node, data)
		return node
	case "Node2D":
		baseNode := nodes.BuildBaseNode(data["name"].(string), nil)
		pos2DComponent := BuildPosition2DComponentFromYaml(data)
		node := nodes.BuildNode2D(*baseNode, *pos2DComponent)
		BuildNodeChildrenFromYaml(node, data)
		return node
	case "Line2D":
		baseNode := nodes.BuildBaseNode(data["name"].(string), nil)
		pos2DComponent := BuildPosition2DComponentFromYaml(data)
		node2d := nodes.BuildNode2D(*baseNode, *pos2DComponent)
		pointsIfaceSlice := data["points"].([]interface{})
		var points []components.Position2D
		for _, pointIface := range pointsIfaceSlice {
			point := BuildPosition2DComponentFromYaml(pointIface.(map[interface{}]interface{}))
			points = append(points, *point)
		}
		width, widthErr := strconv.ParseUint(data["width"].(string), 10, 32)
		if widthErr != nil {
			log.Fatalf("Error parsing width value: %v", widthErr)
		}
		node := nodes.BuildLine2D(*node2d, points, uint(width))
		BuildNodeChildrenFromYaml(node, data)
		return node
	default:
		log.Fatalf("Unknown node type: %s", data["type"])
		return nil
	}
}

func BuildNodeChildrenFromYaml(node nodes.Node, data map[interface{}]interface{}) {
	children := data["children"].([]interface{})
	for _, childData := range children {
		child := BuildNodeFromYaml(childData.(map[interface{}]interface{}))
		node.AppendChild(child)
	}
}

func BuildPosition2DComponentFromYaml(data map[interface{}]interface{}) *components.Position2D {
	xPos, xErr := strconv.ParseFloat(data["xPos"].(string), 64)
	if xErr != nil {
		log.Fatalf("Error parsing x value: %v", xErr)
	}
	yPos, yErr := strconv.ParseFloat(data["yPos"].(string), 64)
	if yErr != nil {
		log.Fatalf("Error parsing y value: %v", yErr)
	}
	return components.BuildPosition2DComponent(xPos, yPos)
}

func LoadShaderSourceFromFile(filename string) (string, error) {
	shaderSource, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(shaderSource), nil
}
