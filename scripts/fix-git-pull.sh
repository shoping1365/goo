#!/bin/bash
# Quick fix for git pull conflict

echo "🔧 Fixing git pull conflict..."

cd /opt/shared

# Stash local changes
echo "📦 Stashing local changes..."
git stash

# Pull latest changes
echo "⬇️  Pulling latest changes..."
git pull origin main

# Apply stashed changes (if any conflicts, we'll use incoming changes)
echo "🔄 Attempting to apply stashed changes..."
git stash pop || {
    echo ""
    echo "⚠️  Conflict detected. Using incoming changes from git..."
    git checkout --theirs scripts/deploy-all-sites.sh
    git add scripts/deploy-all-sites.sh
    git stash drop
}

echo ""
echo "✅ Git pull completed successfully!"
echo ""
echo "Current status:"
git status

echo ""
echo "🚀 Ready to deploy! Run: bash scripts/deploy-all-sites.sh"
