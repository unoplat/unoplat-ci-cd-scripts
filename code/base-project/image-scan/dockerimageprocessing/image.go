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
		currentNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		switch node := currentNode.Value.(type) {
		case map[string]interface{}:
			// Directly check for image specification here
			imagePath, isImage := ii.constructImagePath(node)
			if isImage {
				(*images)[currentNode.Path+".tag"] = imagePath
				continue // Found an image specification, no need to go deeper in this branch
			}

			// If not an image spec, iterate over map and push to stack for further exploration
			for key, value := range node {
				newPath := ii.appendPath(currentNode.Path, key)
				stack = append(stack, Node{Value: value, Path: newPath})
			}
		case []interface{}:
			for i, value := range node {
				newPath := ii.appendPath(currentNode.Path, fmt.Sprintf("[%d]", i))
				stack = append(stack, Node{Value: value, Path: newPath})
			}
		}
	}
}

// Helper function to construct the full image path if possible
func (ii *ImageInfo) constructImagePath(node map[string]interface{}) (string, bool) {
	registry, hasRegistry := node["registry"].(string)
	repository, hasRepository := node["repository"].(string)
	tag, hasTag := node["tag"].(string)

	if hasRegistry && hasRepository && hasTag {
		return fmt.Sprintf("%s/%s:%s", registry, repository, tag), true
	}
	return "", false
}

// Helper function to append new segments to a path
func (ii *ImageInfo) appendPath(basePath, addition string) string {
	if basePath == "" {
		return addition
	}
	return basePath + "." + addition
}

// findImages and appendPath functions remain unchanged, but are methods of ImageInfo struct.
