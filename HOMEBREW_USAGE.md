# Using helm-release with Homebrew

This document provides instructions for installing and using helm-release via Homebrew.

## Installation

### Prerequisites

- macOS or Linux with Homebrew installed
- If you don't have Homebrew installed, you can install it from [brew.sh](https://brew.sh/)

### Installing helm-release

```bash
# Add the tap (only needed once)
brew tap meithon/tap

# Install helm-release
brew install helm-release

# Or in one command
brew install meithon/tap/helm-release
```

## Updating

To update to the latest version:

```bash
# Update Homebrew
brew update

# Upgrade helm-release
brew upgrade helm-release
```

## Uninstalling

To uninstall helm-release:

```bash
brew uninstall helm-release
```

## Troubleshooting

### Checking the installed version

```bash
helm-release --version
```

### Checking the installation path

```bash
which helm-release
```

### Reinstalling

If you encounter issues, you can try reinstalling:

```bash
brew uninstall helm-release
brew install meithon/tap/helm-release
```

### Checking for issues with Homebrew

```bash
brew doctor
```

### Clearing Homebrew cache

```bash
brew cleanup
```

## Using helm-release

Once installed via Homebrew, you can use helm-release as described in the main README:

```bash
# Basic usage
helm-release --file=k8s-resource.yaml --release=my-release --namespace=default

# Using short options
helm-release -f k8s-resource.yaml -r my-release -n default
```

## Getting Help

```bash
# Display help information
helm-release --help
```

## Additional Resources

- [helm-release GitHub Repository](https://github.com/meithon/helm-release)
- [Homebrew Documentation](https://docs.brew.sh/)
