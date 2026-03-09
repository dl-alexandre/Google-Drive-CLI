#!/bin/bash

# Update package declarations in all subpackages

# Drive package
for f in internal/cli/drive/*.go; do
    sed -i '' 's/^package cli$/package drive/' "$f"
done

# Workspace package
for f in internal/cli/workspace/*.go; do
    sed -i '' 's/^package cli$/package workspace/' "$f"
done

# Admin package
for f in internal/cli/admin/*.go; do
    sed -i '' 's/^package cli$/package admin/' "$f"
done

# Chat package
for f in internal/cli/chat/*.go; do
    sed -i '' 's/^package cli$/package chat/' "$f"
done

# Gmail package
for f in internal/cli/gmail/*.go; do
    sed -i '' 's/^package cli$/package gmail/' "$f"
done

# People package
for f in internal/cli/people/*.go; do
    sed -i '' 's/^package cli$/package people/' "$f"
done

# Calendar package
for f in internal/cli/calendar/*.go; do
    sed -i '' 's/^package cli$/package calendar/' "$f"
done

# Activity package
for f in internal/cli/activity/*.go; do
    sed -i '' 's/^package cli$/package activity/' "$f"
done

# Sync package
for f in internal/cli/sync/*.go; do
    sed -i '' 's/^package cli$/package sync/' "$f"
done

# System package
for f in internal/cli/system/*.go; do
    sed -i '' 's/^package cli$/package system/' "$f"
done

echo "Package declarations updated"
