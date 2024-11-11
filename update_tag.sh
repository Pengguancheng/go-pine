#!/bin/bash

git fetch

# Get the highest tag number
VERSION=$(git describe --tags --abbrev=0 2>/dev/null)

# If there are no tags, start at v0.0.0
if [[ -z "$VERSION" ]]; then
    VERSION="v0.0.0"
fi

# Extract the major, minor, and patch numbers
MAJOR=$(echo $VERSION | cut -d. -f1)
MINOR=$(echo $VERSION | cut -d. -f2)
PATCH=$(echo $VERSION | cut -d. -f3)

# If PATCH is empty, set it to 0
if [[ -z "$PATCH" ]]; then
    PATCH=0
fi

# Increment the patch version by 0.01
PATCH=$(echo "$PATCH + 1" | bc)

# Create the new version number
NEW_TAG="$MAJOR.$MINOR.$PATCH"

echo "Current version: $VERSION"
echo "New version: $NEW_TAG"

# Create a new annotated tag
git tag $NEW_TAG

echo "Created new tag: $NEW_TAG"