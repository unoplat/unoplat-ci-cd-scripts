package main

import (
	"fmt"
	"os"

	dockerimageprocessing "github.com/unoplat/unoplat-ci-cd-scripts/code/image-scan/dockerimageprocessing"
	"github.com/unoplat/unoplat-ci-cd-scripts/code/image-scan/utils"
)

func main() {
	valueFileName := os.Getenv("VALUES_FILE")
	data, err := utils.ReadFile(valueFileName)
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
		return
	}

	imageInfo := dockerimageprocessing.NewImageInfo()
	if err := imageInfo.UnmarshalYAML(data); err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
		return
	}

	jsonData, err := imageInfo.MarshalJSON()
	if err != nil {
		fmt.Printf("Error marshalling to JSON: %s\n", err)
		return
	}

	fileName := "images.json"
	if err := utils.WriteFile(fileName, jsonData); err != nil {
		fmt.Printf("Error writing JSON to file: %s\n", err)
		return
	}

	fmt.Printf("Successfully written image data to %s\n", fileName)
}
