#!/bin/bash

# Read the current version from the version file
VERSION=$(cat version)

# If the version file is empty or doesn't exist, start at v0.0.0
if [[ -z "$VERSION" ]]; then
    VERSION="v0.0.0"
fi

# Extract the major, minor, and patch numbers
MAJOR=$(echo $VERSION | cut -d. -f1)
MINOR=$(echo $VERSION | cut -d. -f2)
PATCH=$(echo $VERSION | cut -d. -f3)

# Increment the patch version by 1
PATCH=$((PATCH + 1))

# Create the new version number
NEW_VERSION="$MAJOR.$MINOR.$PATCH"

echo "Current version: $VERSION"
echo "New version: $NEW_VERSION"

# Update the version file
echo $NEW_VERSION > version

echo "Updated version file to: $NEW_VERSION"

# Create a new Git tag
git tag $NEW_VERSION