# Helm Release CLI

A simple CLI tool that performs Helm releases using standard Kubernetes resource YAML files as input.

## Overview

This tool allows you to apply Kubernetes resources using Helm's release management capabilities without requiring a pre-existing Helm chart. It creates a temporary chart on the fly and uses it to deploy the resources via Helm's `upgrade --install` command.

## Installation

### Using Homebrew (recommended)

```bash
# Install from the tap
brew tap meithon/tap
brew install helm-release

# Or in one command
brew install meithon/tap/helm-release
```

### From Source

```bash
# Clone the repository
git clone https://github.com/meithon/helm-release.git
cd helm-release

# Build the binary
go build -o helm-release

# Optionally, move the binary to a directory in your PATH
mv helm-release /usr/local/bin/
```

### From Binary Releases

You can also download pre-built binaries from the [GitHub Releases page](https://github.com/meithon/helm-release/releases).

## Usage

```bash
# Basic usage
helm-release --file=k8s-resource.yaml --release=my-release --namespace=default

# Using short options
helm-release -f k8s-resource.yaml -r my-release -n default
```

### Command Line Options

- `--file`, `-f`: Path to the Kubernetes resource YAML file (required)
- `--release`, `-r`: Helm release name (required)
- `--namespace`, `-n`: Kubernetes namespace (optional, defaults to the namespace from your kubeconfig)
- `--log-level`, `-l`: Log level (optional, defaults to "info"). Available levels: debug, info, warn, error. Log messages are color-coded for better readability.

## How It Works

1. The CLI reads and parses the Kubernetes resource YAML file
2. It adds Helm-specific labels to each Kubernetes resource:
   - `app.kubernetes.io/managed-by: Helm`
   - `app.kubernetes.io/instance: <release-name>`
3. It creates a temporary Helm chart with the following structure:
   ```
   temp-chart/
   ├── Chart.yaml
   └── templates/
       └── resources.yaml  # Contains the K8s resources with Helm labels
   ```
4. It uses the Helm SDK to perform the equivalent of `helm upgrade --install [release] [chart_path] --namespace [namespace]`
5. The temporary chart is cleaned up after the Helm operation completes

## Implementation Details

This tool uses the official Helm SDK (helm.sh/helm/v3) to perform Helm operations directly from Go code, without requiring the Helm CLI to be installed. Key components used:

- `action.Configuration`: Initializes the Helm client with Kubernetes configuration
- `action.NewHistory`: Checks if a release already exists
- `action.NewInstall`: Creates a new install action for new releases
- `action.NewUpgrade`: Creates a new upgrade action for existing releases
- `chart/loader.Load`: Loads the chart from the temporary directory
- `cli.New`: Creates a new Helm settings object using the current environment

The tool intelligently determines whether to install a new release or upgrade an existing one by checking the release history first, providing a seamless experience similar to `helm upgrade --install` but with more robust error handling.

## Requirements

- Go 1.16 or later
- Kubernetes cluster configured (via kubeconfig)

Note: This tool uses the Helm SDK directly, so you don't need to have the Helm CLI installed.

## Development

### Versioning

This project uses [Semantic Versioning](https://semver.org/). Version information is embedded in the binary and can be accessed with:

```bash
helm-release --version
```

### CI/CD

This project uses GitHub Actions for continuous integration and delivery:

- **Build Workflow**: Runs on every push to main and pull requests, testing the code across multiple Go versions
- **Release Workflow**: Triggered when a new tag is pushed, automatically:
  - Builds binaries for multiple platforms (macOS, Linux, Windows)
  - Creates a GitHub release with the binaries
  - Updates the Homebrew formula

### Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

MIT
