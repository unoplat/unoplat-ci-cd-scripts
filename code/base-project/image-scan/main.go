package main

import (
	"os"

	dockerimageprocessing "github.com/unoplat/unoplat-ci-cd-scripts/code/image-scan/dockerimageprocessing"
	"github.com/unoplat/unoplat-ci-cd-scripts/code/image-scan/utils"
	"go.uber.org/zap"
)

func main() {
	valueFileName := os.Getenv("HELM_VALUES_FILE_PATH")

	logger, _ := zap.NewProduction()
	data, err := utils.ReadFile(valueFileName)
	if err != nil {
		logger.Error("Error reading YAML file:", zap.Error(err))
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

	//
	fileName := "docker_images.json"
	if err := utils.WriteFile(fileName, jsonData); err != nil {
		logger.Error("Error writing to json", zap.Error(err))
		return
	}

	logger.Info("Successfully written image data to", zap.String("filename", fileName))
}
