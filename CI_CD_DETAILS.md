# CI/CD Pipeline Details

This document provides detailed information about the CI/CD pipeline set up for the helm-release project.

## Overview

The CI/CD pipeline is implemented using GitHub Actions and consists of two main workflows:

1. **Build and Test Workflow**: Runs on every push to the main branch and on pull requests
2. **Release Workflow**: Triggered when a new tag is pushed, handles the release process

## Build and Test Workflow

File: `.github/workflows/build.yml`

### Triggers

- Push to the `main` branch
- Pull requests to the `main` branch

### Jobs

#### Build and Test

- Runs on Ubuntu latest
- Tests across multiple Go versions (1.21, 1.22, 1.24)
- Steps:
  1. Checkout code
  2. Set up Go environment
  3. Install dependencies
  4. Run linting with golangci-lint
  5. Build the project
  6. Run tests

## Release Workflow

File: `.github/workflows/release.yml`

### Triggers

- Push of a tag matching the pattern `v*` (e.g., v1.0.0)

### Jobs

#### GoReleaser

- Runs on Ubuntu latest
- Steps:
  1. Checkout code with full history (for changelog generation)
  2. Set up Go environment
  3. Run GoReleaser to:
     - Build binaries for multiple platforms
     - Create archives and checksums
     - Generate changelog
     - Create GitHub release with assets

#### Update Homebrew Formula

- Runs after the GoReleaser job completes
- Steps:
  1. Checkout the homebrew-tap repository
  2. Get release information (tag, version, tarball URL, SHA256)
  3. Update the formula with the new version and SHA256
  4. Commit and push changes to the homebrew-tap repository

## GoReleaser Configuration

File: `.goreleaser.yml`

### Key Features

- **Hooks**: Runs `go mod tidy` before building
- **Builds**:
  - Builds for Linux, Windows, and macOS
  - Supports amd64 and arm64 architectures
  - Embeds version information using ldflags
- **Archives**:
  - Creates tar.gz archives for Linux and macOS
  - Creates zip archives for Windows
  - Uses a standardized naming convention
- **Checksums**: Generates SHA256 checksums for all archives
- **Changelog**: Automatically generates a changelog from Git commits

## Required Secrets

- **GITHUB_TOKEN**: Automatically provided by GitHub Actions, used for creating releases
- **HOMEBREW_TAP_TOKEN**: Personal Access Token with repo scope, used for pushing to the homebrew-tap repository

## Setting Up the HOMEBREW_TAP_TOKEN

1. Create a Personal Access Token (PAT) with `repo` scope:
   - Go to GitHub Settings > Developer settings > Personal access tokens
   - Create a new token with `repo` scope
   - Copy the token value

2. Add the token as a secret in the helm-release repository:
   - Go to the helm-release repository on GitHub
   - Navigate to Settings > Secrets and variables > Actions
   - Create a new repository secret named `HOMEBREW_TAP_TOKEN` with the PAT value

## Customizing the CI/CD Pipeline

### Adding More Test Platforms

To test on more platforms, modify the matrix in the build workflow:

```yaml
strategy:
  matrix:
    go-version: ['1.21', '1.22', '1.24']
    os: [ubuntu-latest, macos-latest, windows-latest]
```

### Adding Code Coverage

To add code coverage reporting, add these steps to the build workflow:

```yaml
- name: Run tests with coverage
  run: go test -v -coverprofile=coverage.txt -covermode=atomic ./...

- name: Upload coverage to Codecov
  uses: codecov/codecov-action@v3
  with:
    file: ./coverage.txt
    fail_ci_if_error: true
```

### Adding More Release Targets

To add more architectures or operating systems, modify the GoReleaser configuration:

```yaml
builds:
  - env:
      - CGO_ENABLED=0
    goos:
      - linux
      - windows
      - darwin
      - freebsd
    goarch:
      - amd64
      - arm64
      - arm
      - 386
```

## Troubleshooting

### Common Issues

1. **Release workflow fails to push to homebrew-tap**:
   - Check that the `HOMEBREW_TAP_TOKEN` secret is set correctly
   - Verify that the token has the necessary permissions
   - Ensure the homebrew-tap repository exists and is accessible

2. **GoReleaser fails to build**:
   - Check the GoReleaser logs for specific errors
   - Verify that the Go version is compatible
   - Ensure all dependencies are available

3. **Linting fails**:
   - Run golangci-lint locally to identify and fix issues
   - Consider updating the linter configuration if necessary

### Debugging Workflows

To debug GitHub Actions workflows:

1. Enable debug logging:
   - Set the secret `ACTIONS_RUNNER_DEBUG` to `true`
   - Set the secret `ACTIONS_STEP_DEBUG` to `true`

2. Use the GitHub Actions UI to view detailed logs

3. For local testing, consider using [act](https://github.com/nektos/act) to run workflows locally
