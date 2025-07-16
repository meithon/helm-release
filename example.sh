#!/bin/bash

# Example script to demonstrate how to use the helm-release CLI tool

# Build the CLI tool
echo "Building helm-release CLI tool..."
go build -o helm-release

# Make the script executable
chmod +x helm-release

# Use the CLI tool to deploy the test resources with debug logging
echo "Deploying test resources with debug logging..."
./helm-release --file=test-resources.yaml --release=test-release --log-level=debug
# Note: Namespace will be automatically determined from your kubeconfig

# Note: Since we're using the Helm SDK directly, you can check the release using:
# echo "Checking Helm release..."
# helm list --namespace=default

# Check the deployed resources
echo "Checking deployed resources..."
kubectl get configmap,deployment -n default | grep test-
