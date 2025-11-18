#!/bin/bash
# Debug script for iranxia backend issues
# این اسکریپت مشکلات backend iranxia را بررسی می‌کند

set -e

echo "=========================================="
echo "🔍 DEBUG: Iranxia Backend Issues"
echo "=========================================="
echo ""

# تابع‌های کمکی
print_success() { echo -e "\033[0;32m✅ $1\033[0m"; }
print_error() { echo -e "\033[0;31m❌ $1\033[0m"; }
print_info() { echo -e "\033[0;34mℹ️  $1\033[0m"; }
print_warning() { echo -e "\033[0;33m⚠️  $1\033[0m"; }

PROJECT_NAME="iranxia"
PROJECT_DIR="/data/iranxia"

echo "1️⃣ Checking project directory..."
if [ -d "$PROJECT_DIR" ]; then
    print_success "Project directory exists: $PROJECT_DIR"
    ls -la "$PROJECT_DIR" | head -5
else
    print_error "Project directory not found: $PROJECT_DIR"
    exit 1
fi
echo ""

echo "2️⃣ Checking backend directory..."
if [ -d "$PROJECT_DIR/my-go-backend" ]; then
    print_success "Backend directory exists"
    ls -la "$PROJECT_DIR/my-go-backend" | head -5
else
    print_error "Backend directory not found: $PROJECT_DIR/my-go-backend"
    exit 1
fi
echo ""

echo "3️⃣ Checking binary file..."
if [ -f "$PROJECT_DIR/my-go-backend/api" ]; then
    print_success "Binary file exists"
    ls -la "$PROJECT_DIR/my-go-backend/api"
    file "$PROJECT_DIR/my-go-backend/api"
else
    print_error "Binary file not found: $PROJECT_DIR/my-go-backend/api"
fi
echo ""

echo "4️⃣ Checking .env file..."
ENV_FILE="$PROJECT_DIR/my-go-backend/.env"
if [ -f "$ENV_FILE" ]; then
    print_success ".env file exists"
    echo "Backend port:"
    grep "BACKEND_PORT" "$ENV_FILE" || echo "BACKEND_PORT not found"
    echo "Database settings:"
    grep -E "DB_(HOST|PORT|USER|NAME)" "$ENV_FILE" || echo "DB settings not found"
else
    print_error ".env file not found: $ENV_FILE"
fi
echo ""

echo "5️⃣ Checking PostgreSQL services..."
print_info "Available PostgreSQL services:"
systemctl list-units --type=service | grep postgresql || echo "No PostgreSQL services found"

print_info "Active PostgreSQL services:"
systemctl list-units --type=service --state=active | grep postgresql || echo "No active PostgreSQL services"

# چک کردن نسخه‌های مختلف
for version in postgresql postgresql-16 postgresql-15 postgresql-14; do
    if systemctl list-units --type=service | grep -q "$version.service"; then
        print_info "Found: $version.service"
        systemctl is-active --quiet "$version.service" && print_success "$version is active" || print_warning "$version is inactive"
    fi
done
echo ""

echo "6️⃣ Checking systemd service file..."
SERVICE_PATH="/etc/systemd/system/${PROJECT_NAME}-backend.service"
if [ -f "$SERVICE_PATH" ]; then
    print_success "Service file exists: $SERVICE_PATH"
    echo "Service file content:"
    cat "$SERVICE_PATH"
else
    print_error "Service file not found: $SERVICE_PATH"
fi
echo ""

echo "7️⃣ Checking service status..."
if systemctl list-units --type=service | grep -q "${PROJECT_NAME}-backend.service"; then
    print_info "Service exists in systemd"
    systemctl status "${PROJECT_NAME}-backend" --no-pager || true
else
    print_error "Service not found in systemd"
fi
echo ""

echo "8️⃣ Checking recent logs..."
if systemctl list-units --type=service | grep -q "${PROJECT_NAME}-backend.service"; then
    print_info "Recent logs:"
    journalctl -u "${PROJECT_NAME}-backend" -n 20 --no-pager || true
else
    print_warning "Cannot check logs - service not found"
fi
echo ""

echo "9️⃣ Checking ports..."
print_info "Checking backend port..."
if [ -f "$ENV_FILE" ]; then
    BACKEND_PORT=$(grep "BACKEND_PORT" "$ENV_FILE" | cut -d'=' -f2 | tr -d ' ')
    if [ -n "$BACKEND_PORT" ]; then
        print_info "Backend port: $BACKEND_PORT"
        ss -tulpn | grep ":$BACKEND_PORT " || print_success "Port $BACKEND_PORT is free"
    else
        print_warning "BACKEND_PORT not found in .env"
    fi
else
    print_warning "Cannot check port - .env file not found"
fi
echo ""

echo "🔟 Testing binary execution..."
if [ -f "$PROJECT_DIR/my-go-backend/api" ] && [ -x "$PROJECT_DIR/my-go-backend/api" ]; then
    print_info "Testing binary execution..."
    cd "$PROJECT_DIR/my-go-backend"
    timeout 5s ./api --help 2>&1 | head -5 || print_warning "Binary execution test failed or timed out"
else
    print_warning "Cannot test binary - file not found or not executable"
fi
echo ""

echo "=========================================="
echo "🔍 DEBUG COMPLETE"
echo "=========================================="
echo ""
echo "📋 Summary:"
echo "  - Project directory: $([ -d "$PROJECT_DIR" ] && echo "✅ EXISTS" || echo "❌ MISSING")"
echo "  - Backend directory: $([ -d "$PROJECT_DIR/my-go-backend" ] && echo "✅ EXISTS" || echo "❌ MISSING")"
echo "  - Binary file: $([ -f "$PROJECT_DIR/my-go-backend/api" ] && echo "✅ EXISTS" || echo "❌ MISSING")"
echo "  - .env file: $([ -f "$ENV_FILE" ] && echo "✅ EXISTS" || echo "❌ MISSING")"
echo "  - Service file: $([ -f "$SERVICE_PATH" ] && echo "✅ EXISTS" || echo "❌ MISSING")"
echo "  - Service status: $(systemctl is-active --quiet "${PROJECT_NAME}-backend" 2>/dev/null && echo "✅ ACTIVE" || echo "❌ INACTIVE")"
echo ""
