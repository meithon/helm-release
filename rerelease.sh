#!/bin/bash
set -e

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${YELLOW}Helm Release 再リリーススクリプト${NC}"
echo "このスクリプトは既存のリリースを削除し、新しいリリースを作成します。"
echo

# Check if GitHub CLI is installed
if ! command -v gh &> /dev/null; then
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

echo -e "${GREEN}Step 1: 既存のリリースタグを削除${NC}"
echo "タグ v0.1.0 を削除しています..."

# Delete the tag locally
git tag -d v0.1.0 2>/dev/null || true

# Delete the tag remotely
git push --delete origin v0.1.0 2>/dev/null || true

echo -e "${GREEN}Step 2: CHANGELOGを更新${NC}"
echo "CHANGELOGを更新しています..."

# Check if the CHANGELOG already has a 0.1.0 entry
if grep -q "\[0\.1\.0\]" CHANGELOG.md; then
    echo -e "${YELLOW}CHANGELOGには既に0.1.0のエントリがあります。スキップします。${NC}"
else
    # Update CHANGELOG.md
    sed -i '' 's/## \[Unreleased\]/## [Unreleased]\n\n## [0.1.0] - '$(date +%Y-%m-%d)'/' CHANGELOG.md
    
    # Commit the changes
    git add CHANGELOG.md
    git commit -m "Update CHANGELOG for v0.1.0 release"
    git push origin main
fi

echo -e "${GREEN}Step 3: 新しいリリースタグを作成${NC}"
echo "タグ v0.1.0 を作成しています..."

# Create and push the tag
git tag -a v0.1.0 -m "Release v0.1.0"
git push origin v0.1.0

echo
echo -e "${GREEN}完了！GitHub Actionsワークフローが実行されています。${NC}"
echo "進捗状況は以下のURLで確認できます: https://github.com/meithon/helm-release/actions"
echo
echo "ワークフローが完了したら、以下のコマンドでhelm-releaseをインストールできます:"
echo -e "${YELLOW}  brew tap meithon/tap${NC}"
echo -e "${YELLOW}  brew install helm-release${NC}"
echo
echo "または一行で:"
echo -e "${YELLOW}  brew install meithon/tap/helm-release${NC}"
