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

func (ii *ImageInfo) findImages(node interface{}, path string, images *map[string]string) {
	switch node := node.(type) {
	case map[string]interface{}:
		for key, value := range node {
			newPath := ii.appendPath(path, key)
			// Check if the current map is an image specification
			if registry, ok := node["registry"]; ok {
				if repository, ok := node["repository"]; ok {
					if tag, ok := node["tag"]; ok {
						// Construct the full image path and add it to the map
						fullImagePath := fmt.Sprintf("%v/%v:%v", registry, repository, tag)
						(*images)[newPath] = fullImagePath
						return // Skip further searching in this branch
					}
				}
			}
			// Otherwise, recursively search each value
			ii.findImages(value, newPath, images)
		}
	case []interface{}:
		for i, value := range node {
			newPath := ii.appendPath(path, fmt.Sprintf("[%d]", i))
			ii.findImages(value, newPath, images)
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
