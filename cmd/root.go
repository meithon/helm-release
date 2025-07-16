package cmd

import (
	"fmt"

	"github.com/meithon/helm-release/pkg/chart"
	"github.com/meithon/helm-release/pkg/helm"
	"github.com/meithon/helm-release/pkg/kubernetes"
	"github.com/meithon/helm-release/pkg/logger"
	"github.com/spf13/cobra"
)

// Version is the version of the application.
// This will be overwritten during build by the -X linker flag.
var Version = "dev"

var rootCmd = &cobra.Command{
	Version: Version,
	Use:   "helm-release",
	Short: "A CLI tool to perform Helm releases with Kubernetes resource YAML files",
	Long: `helm-release is a CLI tool that allows you to perform Helm releases
using standard Kubernetes resource YAML files as input.

It creates a temporary Helm chart on the fly and uses it to deploy
the Kubernetes resources via Helm's release functionality.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		file, _ := cmd.Flags().GetString("file")
		release, _ := cmd.Flags().GetString("release")
		namespace, _ := cmd.Flags().GetString("namespace")

		logger.Info("Starting Helm release process")
		logger.Debug("Parameters: file=%s, release=%s, namespace=%s", file, release, namespace)

		// Read and parse the Kubernetes resource YAML file
		logger.Debug("Parsing Kubernetes resource file: %s", file)
		resources, err := kubernetes.ParseResourceFile(file)
		if err != nil {
			logger.Error("Failed to parse resource file: %v", err)
			return fmt.Errorf("failed to parse resource file: %w", err)
		}

		// Add Helm labels to the resources
		logger.Debug("Adding Helm labels to resources")
		resources, err = kubernetes.AddHelmLabels(resources, release)
		if err != nil {
			logger.Error("Failed to add Helm labels: %v", err)
			return fmt.Errorf("failed to add Helm labels: %w", err)
		}

		// Create a temporary chart
		logger.Debug("Creating temporary Helm chart")
		chartPath, err := chart.CreateTempChart(resources)
		if err != nil {
			logger.Error("Failed to create temporary chart: %v", err)
			return fmt.Errorf("failed to create temporary chart: %w", err)
		}
		defer chart.CleanupTempChart(chartPath)
		logger.Debug("Temporary chart created at: %s", chartPath)

		// Execute Helm upgrade --install
		logger.Info("Executing Helm release for '%s'", release)
		err = helm.UpgradeInstall(release, chartPath, namespace)
		if err != nil {
			logger.Error("Failed to execute Helm command: %v", err)
			return fmt.Errorf("failed to execute Helm command: %w", err)
		}

		logger.Info("Successfully deployed release '%s' to namespace '%s'", release, namespace)
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Define flags
	rootCmd.Flags().StringP("file", "f", "", "Path to Kubernetes resource YAML file")
	rootCmd.Flags().StringP("release", "r", "", "Helm release name")
	rootCmd.Flags().StringP("namespace", "n", "", "Kubernetes namespace (defaults to the namespace from your kubeconfig)")
	rootCmd.Flags().StringP("log-level", "l", "info", "Log level (debug, info, warn, error)")

	// Mark required flags
	rootCmd.MarkFlagRequired("file")
	rootCmd.MarkFlagRequired("release")

	// Add a PreRun hook to initialize the logger
	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		// Get the log level from the flag
		logLevelStr, _ := cmd.Flags().GetString("log-level")
		logLevel, err := logger.ParseLevel(logLevelStr)
		if err != nil {
			return err
		}
		
		// Set the log level
		logger.SetDefaultLevel(logLevel)
		logger.Debug("Log level set to %s", logLevelStr)
		
		return nil
	}
}
