# Release Process

This document explains the process for creating a new release of the helm-release tool.

## Prerequisites

- You have push access to the main repository
- You have set up the `HOMEBREW_TAP_TOKEN` secret in the GitHub repository settings

## Release Steps

1. **Update CHANGELOG.md**

   Move items from the "Unreleased" section to a new version section:

   ```markdown
   ## [1.0.0] - YYYY-MM-DD

   ### Added
   - Feature 1
   - Feature 2

   ### Changed
   - Change 1
   - Change 2

   ### Fixed
   - Fix 1
   - Fix 2
   ```

2. **Commit the CHANGELOG update**

   ```bash
   git add CHANGELOG.md
   git commit -m "Update CHANGELOG for v1.0.0 release"
   git push origin main
   ```

3. **Create and push a new tag**

   ```bash
   git tag -a v1.0.0 -m "Release v1.0.0"
   git push origin v1.0.0
   ```

4. **Monitor the GitHub Actions workflow**

   - Go to the "Actions" tab in the GitHub repository
   - You should see the "Release" workflow running
   - This workflow will:
     - Build the binaries for multiple platforms
     - Create a GitHub release with the binaries
     - Update the Homebrew formula

5. **Verify the release**

   - Check the GitHub Releases page to ensure the release was created with the correct assets
   - Verify that the Homebrew formula was updated in the homebrew-tap repository
   - Test the Homebrew installation:

   ```bash
   brew update
   brew install meithon/tap/helm-release
   # Or if already installed
   brew upgrade meithon/tap/helm-release
   ```

## Versioning Guidelines

This project follows [Semantic Versioning](https://semver.org/):

- **MAJOR** version when you make incompatible API changes
- **MINOR** version when you add functionality in a backwards compatible manner
- **PATCH** version when you make backwards compatible bug fixes

## Hotfix Releases

For urgent fixes that need to be released outside the normal release cycle:

1. Create a branch from the latest release tag:

   ```bash
   git checkout -b hotfix/v1.0.1 v1.0.0
   ```

2. Make the necessary fixes and commit them:

   ```bash
   git add .
   git commit -m "Fix critical issue X"
   ```

3. Update the CHANGELOG.md with the hotfix details:

   ```markdown
   ## [1.0.1] - YYYY-MM-DD

   ### Fixed
   - Critical issue X
   ```

4. Commit the CHANGELOG update:

   ```bash
   git add CHANGELOG.md
   git commit -m "Update CHANGELOG for v1.0.1 hotfix"
   ```

5. Create and push a new tag:

   ```bash
   git tag -a v1.0.1 -m "Hotfix v1.0.1"
   git push origin v1.0.1
   ```

6. Merge the hotfix back to main:

   ```bash
   git checkout main
   git merge hotfix/v1.0.1
   git push origin main
   ```

## Troubleshooting

If the release workflow fails:

1. Check the workflow logs for errors
2. Fix any issues in the main branch
3. Delete the failed tag:

   ```bash
   git tag -d v1.0.0
   git push --delete origin v1.0.0
   ```

4. Retry the release process from step 3
