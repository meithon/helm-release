package kubernetes

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/mei/helm-release/pkg/logger"
	"gopkg.in/yaml.v3"
)

// K8sResource represents a Kubernetes resource with metadata
type K8sResource struct {
	APIVersion string                 `yaml:"apiVersion"`
	Kind       string                 `yaml:"kind"`
	Metadata   map[string]interface{} `yaml:"metadata"`
	Spec       map[string]interface{} `yaml:"spec,omitempty"`
	// Other fields are kept as is
}

// ParseResourceFile reads and parses a Kubernetes resource YAML file.
// It returns the raw YAML content as a string, which will be used
// to create the templates in the temporary Helm chart.
func ParseResourceFile(filePath string) (string, error) {
	logger.Debug("Parsing Kubernetes resource file: %s", filePath)
	
	// Read the file
	logger.Debug("Reading file: %s", filePath)
	data, err := os.ReadFile(filePath)
	if err != nil {
		logger.Error("Failed to read file %s: %v", filePath, err)
		return "", fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	// Validate that the file contains valid YAML
	logger.Debug("Validating YAML content")
	var obj interface{}
	if err := yaml.Unmarshal(data, &obj); err != nil {
		logger.Error("Invalid YAML in file %s: %v", filePath, err)
		return "", fmt.Errorf("invalid YAML in file %s: %w", filePath, err)
	}

	logger.Debug("Successfully parsed Kubernetes resource file: %s", filePath)
	// Return the raw YAML content
	return string(data), nil
}

// AddHelmLabels adds Helm-specific labels to Kubernetes resources.
// It takes the raw YAML content and a release name, and returns the modified YAML content.
func AddHelmLabels(yamlContent string, releaseName string) (string, error) {
	logger.Debug("Adding Helm labels to Kubernetes resources")
	
	// Split the YAML content into multiple documents (for multi-resource YAML files)
	yamlDocs := strings.Split(yamlContent, "---\n")
	var modifiedDocs []string
	
	for _, doc := range yamlDocs {
		// Skip empty documents
		if strings.TrimSpace(doc) == "" {
			continue
		}
		
		// Parse the YAML document
		var resource K8sResource
		if err := yaml.Unmarshal([]byte(doc), &resource); err != nil {
			logger.Error("Failed to parse YAML document: %v", err)
			return "", fmt.Errorf("failed to parse YAML document: %w", err)
		}
		
		// Skip documents without apiVersion or kind (not K8s resources)
		if resource.APIVersion == "" || resource.Kind == "" {
			modifiedDocs = append(modifiedDocs, doc)
			continue
		}
		
		// Initialize metadata if it doesn't exist
		if resource.Metadata == nil {
			resource.Metadata = make(map[string]interface{})
		}
		
		// Initialize labels if they don't exist
		labels, ok := resource.Metadata["labels"].(map[string]interface{})
		if !ok {
			labels = make(map[string]interface{})
			resource.Metadata["labels"] = labels
		}
		
		// Add Helm labels
		labels["app.kubernetes.io/managed-by"] = "Helm"
		labels["app.kubernetes.io/instance"] = releaseName
		
		// Convert the modified resource back to YAML
		var buf bytes.Buffer
		encoder := yaml.NewEncoder(&buf)
		encoder.SetIndent(2)
		if err := encoder.Encode(resource); err != nil {
			logger.Error("Failed to encode modified resource: %v", err)
			return "", fmt.Errorf("failed to encode modified resource: %w", err)
		}
		
		modifiedDocs = append(modifiedDocs, buf.String())
	}
	
	// Join the modified documents back together
	result := strings.Join(modifiedDocs, "---\n")
	logger.Debug("Successfully added Helm labels to Kubernetes resources")
	return result, nil
}
