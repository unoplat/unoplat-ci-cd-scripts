package dockerimageprocessing

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

type ImageInfo struct {
	Images map[string]string
}

func NewImageInfo() *ImageInfo {
	return &ImageInfo{
		Images: make(map[string]string),
	}
}

func (ii *ImageInfo) UnmarshalYAML(data []byte) error {
	var values map[string]interface{}
	if err := yaml.Unmarshal(data, &values); err != nil {
		return err
	}
	ii.findImages(values, "", &ii.Images)
	return nil
}

func (ii *ImageInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(ii.Images)
}

type Node struct {
	Value interface{}
	Path  string
}

func (ii *ImageInfo) findImages(node interface{}, path string, images *map[string]string) {
	stack := []Node{{Value: node, Path: path}}

	for len(stack) > 0 {
		// Pop the last node from the stack
		var currentNode Node
		currentNode, stack = stack[len(stack)-1], stack[:len(stack)-1]
		path := currentNode.Path

		switch node := currentNode.Value.(type) {
		case map[string]interface{}:
			// Check for image specification directly here
			if registry, okR := node["registry"]; okR {
				if repository, okRepo := node["repository"]; okRepo {
					if tag, okTag := node["tag"]; okTag {
						fullImagePath := fmt.Sprintf("%v/%v:%v", registry, repository, tag)
						(*images)[path] = fullImagePath
						// Do not continue processing this node
						break
					}
				}
			}
			// If not an image spec, or partially so, iterate over map and push to stack
			for key, value := range node {
				newPath := ii.appendPath(path, key)
				stack = append(stack, Node{Value: value, Path: newPath})
			}
		case []interface{}:
			for i, value := range node {
				newPath := ii.appendPath(path, fmt.Sprintf("[%d]", i))
				stack = append(stack, Node{Value: value, Path: newPath})
			}
		}
	}
}

// Helper function to append new segments to a path
func (ii *ImageInfo) appendPath(basePath, addition string) string {
	if basePath == "" {
		return addition
	}
	return basePath + "." + addition
}

// findImages and appendPath functions remain unchanged, but are methods of ImageInfo struct.
