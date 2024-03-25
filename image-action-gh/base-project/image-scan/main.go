package main

import (
	"os"
	"path/filepath"

	dockerimageprocessing "github.com/unoplat/unoplat-ci-cd-scripts/code/image-scan/dockerimageprocessing"
	"github.com/unoplat/unoplat-ci-cd-scripts/code/image-scan/utils"
	"go.uber.org/zap"
)

func main() {

	valueFileName := os.Getenv("HELM_VALUES_FILE_PATH")
	jsonFilePath := os.Getenv("DOCKER_IMAGES_JSON_PATH")
	logger, _ := zap.NewProduction()

	if jsonFilePath == "" {
		logger.Error("DOCKER_IMAGES_JSON_PATH environment variable is not set")
		return
	}

	data, err := utils.ReadFile(valueFileName)
	if err != nil {
		logger.Error("Error reading YAML file::", zap.Error(err))
		return
	}

	imageInfo := dockerimageprocessing.NewImageInfo()
	if err := imageInfo.UnmarshalYAML(data); err != nil {
		logger.Error("Error parsing YAML file:", zap.Error(err))
		return
	}

	jsonData, err := imageInfo.MarshalJSON()
	if err != nil {
		logger.Error("Error marshalling json:", zap.Error(err))
		return
	}

	fileName := "docker_images.json"
	fullPath := filepath.Join(jsonFilePath, fileName)

	if err := utils.WriteFile(fullPath, jsonData); err != nil {
		logger.Error("Error writing to json", zap.Error(err))
		return
	}

	logger.Info("Successfully written image data to", zap.String("filename", fileName))
}
