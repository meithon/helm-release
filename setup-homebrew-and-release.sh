#!/bin/bash
set -e

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${YELLOW}Helm Release Homebrew Setup and Release Script${NC}"
echo "This script will help you set up the Homebrew tap and create the first release."
echo

# Check if GitHub CLI is installed
if ! command -v gh &>/dev/null; then
  echo -e "${RED}GitHub CLI (gh) is not installed. Please install it first:${NC}"
  echo "  https://cli.github.com/manual/installation"
  exit 1
fi

# Check if user is logged in to GitHub
if ! gh auth status &> /dev/null; then
    echo -e "${RED}You are not logged in to GitHub. Please login first:${NC}"
    echo "  gh auth login"
    exit 1
fi

echo -e "${GREEN}Step 1: Creating homebrew-tap repository on GitHub${NC}"
echo "Creating repository: meithon/homebrew-tap..."

if gh repo view meithon/homebrew-tap &>/dev/null; then
  echo -e "${YELLOW}Repository already exists. Skipping creation.${NC}"
else
  gh repo create meithon/homebrew-tap --public --description "Homebrew tap for meithon's tools" --confirm
  echo -e "${GREEN}Repository created successfully!${NC}"
fi

echo
echo -e "${GREEN}Step 2: Pushing initial files to homebrew-tap repository${NC}"
echo "Initializing local repository..."

cd homebrew-tap

# Update the SHA256 hash in the formula
sed -i '' 's/REPLACE_WITH_ACTUAL_SHA256_AFTER_FIRST_RELEASE/cc86e974b6f3e9cf03b9d4e7722228ba5f4b68004b5b4ff8a2d9f304fa66cfd0/' helm-release.rb

git init
git add .
git commit -m "Initial commit"

echo "Setting remote origin..."
git remote add origin https://github.com/meithon/homebrew-tap.git

echo "Pushing to GitHub..."
git push -u origin main

echo
echo -e "${GREEN}Step 3: Creating Personal Access Token for GitHub Actions${NC}"
echo "You need to create a Personal Access Token with 'repo' scope."
echo "This token will be used by GitHub Actions to update the Homebrew formula."
echo
echo -e "${YELLOW}Please go to: https://github.com/settings/tokens/new${NC}"
echo "1. Set a note like 'HOMEBREW_TAP_TOKEN'"
echo "2. Select the 'repo' scope"
echo "3. Click 'Generate token'"
echo "4. Copy the generated token"
echo

read -p "Have you created and copied the token? (y/n) " -n 1 -r
echo
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
  echo -e "${RED}Please create the token before continuing.${NC}"
  exit 1
fi

echo
echo -e "${GREEN}Step 4: Adding the token as a secret to helm-release repository${NC}"
echo "Adding HOMEBREW_TAP_TOKEN secret to meithon/helm-release repository..."

read -p "Paste your token here (it will not be displayed): " -s TOKEN
echo

if [ -z "$TOKEN" ]; then
  echo -e "${RED}Token cannot be empty.${NC}"
  exit 1
fi

gh secret set HOMEBREW_TAP_TOKEN -b "$TOKEN" -R meithon/helm-release

echo -e "${GREEN}Secret added successfully!${NC}"

echo
echo -e "${GREEN}Step 5: Updating CHANGELOG.md for the first release${NC}"
echo "Updating CHANGELOG.md..."

cd ..
sed -i '' 's/## \[Unreleased\]/## [Unreleased]\n\n## [0.1.0] - '$(date +%Y-%m-%d)'/' CHANGELOG.md

git add CHANGELOG.md
git commit -m "Update CHANGELOG for v0.1.0 release"
git push origin main

echo
echo -e "${GREEN}Step 6: Creating and pushing the first release tag${NC}"
echo "Creating tag v0.1.0..."

git tag -a v0.1.0 -m "Release v0.1.0"
git push origin v0.1.0

echo
echo -e "${GREEN}All done! The GitHub Actions workflow should now be running.${NC}"
echo "You can check the progress at: https://github.com/meithon/helm-release/actions"
echo
echo "Once the workflow completes, you can install helm-release with:"
echo -e "${YELLOW}  brew tap meithon/tap${NC}"
echo -e "${YELLOW}  brew install helm-release${NC}"
echo
echo "Or in one command:"
echo -e "${YELLOW}  brew install meithon/tap/helm-release${NC}"
