package helm

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/mei/helm-release/pkg/logger"
)

// UpgradeInstall executes the 'helm upgrade --install' command with the given parameters.
func UpgradeInstall(releaseName, chartPath, namespace string) error {
	// If namespace is empty, get it from kubeconfig
	if namespace == "" {
		// Get the namespace from kubeconfig using kubectl
		cmd := exec.Command("kubectl", "config", "view", "--minify", "--output", "jsonpath={..namespace}")
		output, err := cmd.Output()
		if err == nil && len(output) > 0 {
			namespace = string(output)
			logger.Info("Using namespace %q from kubeconfig", namespace)
		} else {
			// Default to "default" namespace if not specified and can't be determined
			namespace = "default"
			logger.Info("Using default namespace %q", namespace)
		}
	}

	// Prepare the command
	args := []string{"upgrade", "--install", releaseName, chartPath, "--namespace", namespace}
	
	logger.Debug("Executing command: helm %v", args)
	
	// Create the command
	cmd := exec.Command("helm", args...)
	
	// Set the command's stdout and stderr to the current process's stdout and stderr
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	// Execute the command
	if err := cmd.Run(); err != nil {
		logger.Error("Helm upgrade --install command failed: %v", err)
		return fmt.Errorf("helm upgrade --install command failed: %w", err)
	}
	
	logger.Info("Release %q successfully deployed to namespace %q", releaseName, namespace)
	return nil
}
