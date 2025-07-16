#!/bin/bash
set -e

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${YELLOW}Helm Release Homebrew Formula Update Script${NC}"
echo "This script will update the Homebrew formula with the correct SHA256 hash."
echo

# Check if brew is installed
if ! command -v brew &> /dev/null; then
    echo -e "${RED}Homebrew is not installed. Please install it first:${NC}"
    echo "  https://brew.sh/"
    exit 1
fi

echo -e "${GREEN}Step 1: Updating the Homebrew formula${NC}"
echo "Updating the formula with the correct SHA256 hash..."

# Get the correct SHA256 hash
SHA256="cc86e974b6f3e9cf03b9d4e7722228ba5f4b68004b5b4ff8a2d9f304fa66cfd0"

# Check if the tap is already installed
if ! brew tap | grep -q "meithon/tap"; then
    echo "Adding tap meithon/tap..."
    brew tap meithon/tap
fi

# Get the formula path
FORMULA_PATH=$(brew --repository)/Library/Taps/meithon/homebrew-tap/Formula/helm-release.rb
if [ ! -f "$FORMULA_PATH" ]; then
    FORMULA_PATH=$(brew --repository)/Library/Taps/meithon/homebrew-tap/helm-release.rb
fi

if [ ! -f "$FORMULA_PATH" ]; then
    echo -e "${RED}Could not find the formula file. Please make sure the tap is installed correctly.${NC}"
    exit 1
fi

echo "Formula path: $FORMULA_PATH"

# Update the formula
echo "Updating the formula with SHA256: $SHA256"
sed -i '' 's/sha256 ".*"/sha256 "'$SHA256'"/' "$FORMULA_PATH"

echo -e "${GREEN}Formula updated successfully!${NC}"

echo
echo -e "${GREEN}Step 2: Testing the installation${NC}"
echo "Now you can install helm-release with:"
echo -e "${YELLOW}  brew install meithon/tap/helm-release${NC}"
echo
echo "Or if you already have it installed, you can upgrade it with:"
echo -e "${YELLOW}  brew upgrade meithon/tap/helm-release${NC}"
