package chart

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mei/helm-release/pkg/logger"
)

const chartYamlTemplate = `apiVersion: v2
name: temp-chart
description: A temporary chart created by helm-release
type: application
version: 0.1.0
appVersion: "1.0.0"
`

// CreateTempChart creates a temporary Helm chart with the given Kubernetes resources.
// It returns the path to the temporary chart.
func CreateTempChart(resources string) (string, error) {
	logger.Debug("Creating temporary Helm chart")
	
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "helm-release-chart-*")
	if err != nil {
		logger.Error("Failed to create temporary directory: %v", err)
		return "", fmt.Errorf("failed to create temporary directory: %w", err)
	}
	logger.Debug("Created temporary directory: %s", tempDir)

	// Create Chart.yaml
	chartYamlPath := filepath.Join(tempDir, "Chart.yaml")
	logger.Debug("Creating Chart.yaml at %s", chartYamlPath)
	if err := os.WriteFile(chartYamlPath, []byte(chartYamlTemplate), 0644); err != nil {
		logger.Error("Failed to create Chart.yaml: %v", err)
		os.RemoveAll(tempDir)
		return "", fmt.Errorf("failed to create Chart.yaml: %w", err)
	}

	// Create templates directory
	templatesDir := filepath.Join(tempDir, "templates")
	logger.Debug("Creating templates directory at %s", templatesDir)
	if err := os.Mkdir(templatesDir, 0755); err != nil {
		logger.Error("Failed to create templates directory: %v", err)
		os.RemoveAll(tempDir)
		return "", fmt.Errorf("failed to create templates directory: %w", err)
	}

	// Create resources.yaml in templates directory
	resourcesPath := filepath.Join(templatesDir, "resources.yaml")
	logger.Debug("Creating resources.yaml at %s", resourcesPath)
	if err := os.WriteFile(resourcesPath, []byte(resources), 0644); err != nil {
		logger.Error("Failed to create resources.yaml: %v", err)
		os.RemoveAll(tempDir)
		return "", fmt.Errorf("failed to create resources.yaml: %w", err)
	}

	logger.Debug("Temporary chart created successfully at %s", tempDir)
	return tempDir, nil
}

// CleanupTempChart removes the temporary chart directory.
func CleanupTempChart(chartPath string) error {
	logger.Debug("Cleaning up temporary chart at %s", chartPath)
	if err := os.RemoveAll(chartPath); err != nil {
		logger.Error("Failed to remove temporary chart directory: %v", err)
		return fmt.Errorf("failed to remove temporary chart directory: %w", err)
	}
	logger.Debug("Temporary chart cleaned up successfully")
	return nil
}
