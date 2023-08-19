package importer

import (
	"log"
	"strconv"

	nodes "github.com/Dappetizer/engine-sandbox-golang/engine/nodes"
	components "github.com/Dappetizer/engine-sandbox-golang/engine/nodes/components"
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
		xPos, xErr := strconv.ParseFloat(data["xPos"].(string), 64)
		if xErr != nil {
			log.Fatalf("Error parsing x value: %v", xErr)
		}
		yPos, yErr := strconv.ParseFloat(data["yPos"].(string), 64)
		if yErr != nil {
			log.Fatalf("Error parsing y value: %v", yErr)
		}
		pos2DComponent := components.BuildPosition2DComponent(xPos, yPos)
		node := nodes.BuildNode2D(*baseNode, *pos2DComponent)
		BuildNodeChildrenFromYaml(node, data)
		return node
	case "Line2D":
		baseNode := nodes.BuildBaseNode(data["name"].(string), nil)
		xPos, xErr := strconv.ParseFloat(data["xPos"].(string), 64)
		if xErr != nil {
			log.Fatalf("Error parsing x value: %v", xErr)
		}
		yPos, yErr := strconv.ParseFloat(data["yPos"].(string), 64)
		if yErr != nil {
			log.Fatalf("Error parsing y value: %v", yErr)
		}
		pos2DComponent := components.BuildPosition2DComponent(xPos, yPos)
		node2d := nodes.BuildNode2D(*baseNode, *pos2DComponent)

		pointsIfaceSlice := data["points"].([]interface{})
		var points []components.Position2D
		for _, pointIface := range pointsIfaceSlice {
			m := pointIface.(map[interface{}]interface{})
			x, xErr := strconv.ParseFloat(m["xPos"].(string), 64)
			if xErr != nil {
				log.Fatalf("Error parsing x value: %v", xErr)
			}
			y, yErr := strconv.ParseFloat(m["yPos"].(string), 64)
			if yErr != nil {
				log.Fatalf("Error parsing y value: %v", yErr)
			}
			pos := components.BuildPosition2DComponent(x, y)
			points = append(points, *pos)
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
