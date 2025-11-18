#!/bin/bash
# Quick Setup and Deploy Script
# این اسکریپت همه کارها رو یکجا انجام میده

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${SCRIPT_DIR}/sites-config.sh"

echo "=========================================="
echo "🔧 Step 1: Setting up Systemd Services"
echo "=========================================="

# ایجاد systemd services
bash "${SCRIPT_DIR}/setup-systemd-services.sh"

echo ""
echo "=========================================="
echo "✅ Step 2: Enabling Services"
echo "=========================================="

for project_dir in "${PROJECTS[@]}"; do
    project_name=$(basename "$project_dir")
    print_info "Enabling ${project_name}-backend..."
    systemctl enable "${project_name}-backend" 2>/dev/null || true
done

echo ""
echo "=========================================="
echo "🚀 Step 3: Deploying All Projects"
echo "=========================================="

# اجرای deploy
bash "${SCRIPT_DIR}/deploy-all-sites.sh"

echo ""
echo "=========================================="
echo "✅ SETUP COMPLETE!"
echo "=========================================="
