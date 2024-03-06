package dockerimageprocessing

import (
	"encoding/json"
	"fmt"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

// ImageInfo is a struct that encapsulates information about images.
// It contains a map of image names to their respective file paths and a logger instance.
type ImageInfo struct {
	Images map[string][]string
	logger *zap.Logger
}

func NewImageInfo() *ImageInfo {
	// Initialize the zap logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	return &ImageInfo{
		Images: make(map[string][]string),
		logger: logger,
	}
}

func (ii *ImageInfo) UnmarshalYAML(data []byte) error {
	var values map[string]interface{}
	if err := yaml.Unmarshal(data, &values); err != nil {
		ii.logger.Error("Failed to unmarshal YAML!", zap.Error(err))
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

// findImages is a method of the ImageInfo struct that traverses a given node tree
// (represented as an interface{}) in search of image specifications.
// The method uses a stack to perform a depth-first search of the tree.
// If an image specification is found, its path and corresponding image path are logged and stored in the images map.
// The method also handles nested structures, such as maps and slices, by iterating over their elements and pushing them to the stack.
//
// Parameters:
// node: the root node of the tree to be searched.
// path: the initial path to the node.
// images: a pointer to a map where the paths (as keys) and image paths (as values) of found images will be stored.
//
// This method does not return any value.
func (ii *ImageInfo) findImages(node interface{}, path string, images *map[string][]string) {
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
				ii.logger.Info("Found Docker Image Specification :", zap.String("path", currentNode.Path+".tag"), zap.String("imagePath", imagePath))
				// Access the slice of strings stored in the map 'images' using 'imagePath' as the key.
				// Append the value of 'currentNode.Path' concatenated with ".tag" to this slice.
				// Store the updated slice back in the map at the same key.
				(*images)[imagePath] = append((*images)[imagePath], currentNode.Path+".tag")

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
