#!/bin/bash
# Setup Systemd Services for All Projects
# این اسکریپت برای هر پروژه یک systemd service میسازه

set -e

# Load configuration
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${SCRIPT_DIR}/sites-config.sh"

echo "=========================================="
echo "🔧 Setting up Systemd Services"
echo "=========================================="

for project_dir in "${PROJECTS[@]}"; do
    project_name=$(basename "$project_dir")
    
    print_info "Creating systemd service for ${project_name}-backend..."
    
    # بررسی وجود فایل .env بک‌اند
    if [ ! -f "${project_dir}/my-go-backend/.env" ]; then
        print_warning "${project_dir}/my-go-backend/.env not found!"
        print_info "Creating from main .env..."
        
        # اگر .env اصلی وجود داره، کپی کن
        if [ -f "/opt/shared/my-go-backend/.env" ]; then
            cp /opt/shared/my-go-backend/.env ${project_dir}/my-go-backend/.env
            print_success "Copied .env from /opt/shared"
        fi
    fi
    
    cat > "/etc/systemd/system/${project_name}-backend.service" << EOF
[Unit]
Description=${project_name} Go Backend API
After=network.target postgresql.service postgresql-16.service
Wants=postgresql.service postgresql-16.service

[Service]
Type=simple
User=root
Group=root
WorkingDirectory=${project_dir}/my-go-backend

# Load environment from .env file
EnvironmentFile=-${project_dir}/my-go-backend/.env

# اجرای باینری Go
ExecStart=${project_dir}/my-go-backend/api

# Restart policy
Restart=always
RestartSec=5s

# Timeout settings
TimeoutStartSec=30s
TimeoutStopSec=30s

# Resource limits
LimitNOFILE=65536
LimitNPROC=4096

# Logging
StandardOutput=append:${project_dir}/logs/backend.log
StandardError=append:${project_dir}/logs/backend-error.log

# Security
NoNewPrivileges=true
PrivateTmp=true

# Environment
Environment="GIN_MODE=release"
Environment="ENV=production"

[Install]
WantedBy=multi-user.target
EOF

    print_success "Service file created: /etc/systemd/system/${project_name}-backend.service"
done

# Reload systemd
print_info "Reloading systemd daemon..."
systemctl daemon-reload
print_success "Systemd daemon reloaded"

echo ""
echo "=========================================="
print_success "✅ All systemd services created!"
echo "=========================================="
echo ""
print_info "Starting and enabling services..."
echo ""

# Enable و Start همه services
for project_dir in "${PROJECTS[@]}"; do
    project_name=$(basename "$project_dir")
    
    print_info "Enabling ${project_name}-backend..."
    systemctl enable "${project_name}-backend" 2>/dev/null || print_warning "Could not enable"
    
    # فقط اگه باینری وجود داشته باشه، start کن
    if [ -f "${project_dir}/my-go-backend/api" ]; then
        print_info "Starting ${project_name}-backend..."
        systemctl start "${project_name}-backend" 2>/dev/null || print_warning "Could not start - maybe need to build first"
        
        sleep 1
        if systemctl is-active --quiet "${project_name}-backend"; then
            print_success "${project_name}-backend is RUNNING ✅"
        else
            print_warning "${project_name}-backend is NOT RUNNING - will start after build"
        fi
    else
        print_warning "Binary not found: ${project_dir}/my-go-backend/api"
        print_info "Will start after first build (deploy)"
    fi
    echo ""
done

echo "=========================================="
print_info "Service management commands:"
echo ""
for project_dir in "${PROJECTS[@]}"; do
    project_name=$(basename "$project_dir")
    echo "  # ${project_name}:"
    echo "    systemctl status ${project_name}-backend"
    echo "    systemctl restart ${project_name}-backend"
    echo "    journalctl -u ${project_name}-backend -f"
    echo ""
done