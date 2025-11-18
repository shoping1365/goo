#!/bin/bash
# GitHub Actions Self-Hosted Runner Installation Script
# Date: October 8, 2025
# Repository: arash-mohammadlou/go-nuxt
# Installation Path: /opt/deployment/githubrunner

echo "🚀 Starting GitHub Actions Self-Hosted Runner installation..."

# Step 1: Create folder with proper permissions
echo "📁 Creating /opt/deployment/githubrunner directory..."
sudo mkdir -p /opt/deployment/githubrunner
sudo chown $USER:$USER /opt/deployment/githubrunner
cd /opt/deployment/githubrunner

# Step 2: Download the latest runner package
echo "⬇️ Downloading runner package..."
curl -o actions-runner-linux-x64-2.328.0.tar.gz -L https://github.com/actions/runner/releases/download/v2.328.0/actions-runner-linux-x64-2.328.0.tar.gz

# Step 3: Optional: Validate the hash
echo "🔐 Validating hash..."
echo "01066fad3a2893e63e6ca880ae3a1fad5bf93929d60e77ee15f2b97c148c3cd4e  actions-runner-linux-x64-2.328.0.tar.gz" | shasum -a 256 -c

# Step 4: Extract the installer
echo "📦 Extracting installer..."
tar xzf ./actions-runner-linux-x64-2.328.0.tar.gz

echo ""
echo "✅ Download and extraction completed!"
echo ""
echo "📝 Next steps:"
echo "==============================================="
echo "1. Configure the runner:"
echo "   ./config.sh --url https://github.com/arash-mohammadlou/go-nuxt --token AYBUUDP53H72WPA7D7DJG2DI4M6KQ"
echo ""
echo "2. Install as a service (recommended):"
echo "   sudo ./svc.sh install"
echo "   sudo ./svc.sh start"
echo ""
echo "3. Or run interactively for testing:"
echo "   ./run.sh"
echo ""
echo "4. Check status:"
echo "   sudo ./svc.sh status"
echo "==============================================="
