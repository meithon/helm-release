# Homebrew Tap Setup Guide

This document explains how to set up the Homebrew tap repository for the helm-release tool.

## Overview

A Homebrew tap is a repository that contains Homebrew formulas. To distribute helm-release via Homebrew, we need to create a tap repository.

## Steps to Set Up the Homebrew Tap

1. Create a new GitHub repository named `homebrew-tap`
   - The repository name must start with `homebrew-`
   - Example: `github.com/meithon/homebrew-tap`

2. Clone the repository locally:
   ```bash
   git clone https://github.com/meithon/homebrew-tap.git
   cd homebrew-tap
   ```

3. Add the formula file:
   - Create a file named `helm-release.rb` in the root of the repository
   - Use the template from `homebrew-formula-template.rb` in the main project
   - For the initial setup, you can use a placeholder SHA256 value, which will be updated by the CI/CD pipeline after the first release

4. Commit and push the formula:
   ```bash
   git add helm-release.rb
   git commit -m "Add helm-release formula"
   git push
   ```

5. Create a Personal Access Token (PAT) for GitHub Actions:
   - Go to GitHub Settings > Developer settings > Personal access tokens
   - Create a new token with `repo` scope
   - Copy the token value

6. Add the token as a secret in the helm-release repository:
   - Go to the helm-release repository on GitHub
   - Navigate to Settings > Secrets and variables > Actions
   - Create a new repository secret named `HOMEBREW_TAP_TOKEN` with the PAT value

## Testing the Tap Locally

You can test the tap locally before pushing it to GitHub:

```bash
# Create a local tap
brew tap-new meithon/tap

# Copy your formula to the local tap
cp helm-release.rb $(brew --repository)/Library/Taps/meithon/homebrew-tap/Formula/

# Install from the local tap
brew install meithon/tap/helm-release
```

## Updating the Formula

The GitHub Actions workflow in the helm-release repository will automatically update the formula in the homebrew-tap repository when a new release is created. The workflow:

1. Calculates the SHA256 hash of the release tarball
2. Updates the formula with the new version and hash
3. Commits and pushes the changes to the homebrew-tap repository

## Manual Updates (if needed)

If you need to manually update the formula:

```bash
# Calculate the SHA256 hash of the tarball
curl -sL https://github.com/meithon/helm-release/archive/refs/tags/v0.1.0.tar.gz | shasum -a 256

# Update the formula with the new hash
# Edit helm-release.rb and update the url and sha256 fields

# Commit and push the changes
git add helm-release.rb
git commit -m "Update helm-release to v0.1.0"
git push
```

## User Installation

Once the tap is set up, users can install helm-release with:

```bash
brew tap meithon/tap
brew install helm-release

# Or in one command
brew install meithon/tap/helm-release
```
