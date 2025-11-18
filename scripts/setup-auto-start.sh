#!/bin/bash
# اسکریپت راه‌اندازی خودکار سرویس‌ها بعد از ریبوت سرور
# این اسکریپت باید یک بار در هر پروژه اجرا شود

set -e

# رنگ‌ها برای خروجی
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

print_success() {
    echo -e "${GREEN}✅ $1${NC}"
}

print_error() {
    echo -e "${RED}❌ $1${NC}"
}

print_info() {
    echo -e "${BLUE}ℹ️  $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

# بررسی اجرا با root
if [ "$EUID" -ne 0 ]; then 
    print_error "این اسکریپت باید با دسترسی root اجرا شود"
    exit 1
fi

# دریافت مسیر پروژه از آرگومان یا استفاده از مسیر فعلی
PROJECT_DIR="${1:-$(pwd)}"
PROJECT_NAME=$(basename "$PROJECT_DIR")

print_info "راه‌اندازی auto-start برای پروژه: $PROJECT_NAME در $PROJECT_DIR"

# بارگذاری فایل .env
if [ -f "$PROJECT_DIR/.env" ]; then
    print_info "بارگذاری تنظیمات از .env..."
    export $(cat "$PROJECT_DIR/.env" | grep -v '^#' | xargs)
else
    print_warning "فایل .env یافت نشد، از تنظیمات پیش‌فرض استفاده می‌شود"
fi

FRONTEND_PORT=${FRONTEND_PORT:-3000}
BACKEND_PORT=${BACKEND_PORT:-9090}

# ========================================
# 1. راه‌اندازی PM2 Startup
# ========================================
print_info "تنظیم PM2 برای راه‌اندازی خودکار..."

# بررسی اینکه PM2 نصب است
if ! command -v pm2 &> /dev/null; then
    print_error "PM2 نصب نیست. لطفاً ابتدا PM2 را نصب کنید: npm install -g pm2"
    exit 1
fi

# تنظیم PM2 startup
pm2 startup systemd -u root --hp /root

# ذخیره لیست فعلی PM2
pm2 save

print_success "PM2 برای راه‌اندازی خودکار تنظیم شد"

# ========================================
# 2. ایجاد systemd service برای Backend
# ========================================
print_info "ایجاد systemd service برای Backend..."

cat > /etc/systemd/system/${PROJECT_NAME}-backend.service << EOF
[Unit]
Description=${PROJECT_NAME} Go Backend API
After=network.target postgresql.service postgresql-16.service

[Service]
Type=simple
User=root
WorkingDirectory=$PROJECT_DIR/my-go-backend
Environment="PORT=$BACKEND_PORT"
Environment="GIN_MODE=release"
ExecStart=/usr/local/go/bin/go run ./cmd/api
Restart=always
RestartSec=10
StandardOutput=append:$PROJECT_DIR/logs/backend.log
StandardError=append:$PROJECT_DIR/logs/backend-error.log

# محدودیت‌های امنیتی
NoNewPrivileges=true
PrivateTmp=true

# تنظیمات Resource Limits
LimitNOFILE=65536
LimitNPROC=4096

[Install]
WantedBy=multi-user.target
EOF

print_success "systemd service برای Backend ایجاد شد"

# ========================================
# 3. فعال‌سازی و شروع سرویس‌ها
# ========================================
print_info "فعال‌سازی سرویس‌ها..."

# Reload systemd
systemctl daemon-reload

# فعال کردن backend service برای auto-start
systemctl enable ${PROJECT_NAME}-backend.service

# شروع backend service
systemctl restart ${PROJECT_NAME}-backend.service

print_success "Backend service فعال و راه‌اندازی شد"

# ========================================
# 4. فعال‌سازی Caddy (اگر نصب است)
# ========================================
if systemctl list-unit-files | grep -q caddy.service; then
    print_info "فعال‌سازی Caddy..."
    systemctl enable caddy
    systemctl restart caddy
    print_success "Caddy فعال و راه‌اندازی شد"
else
    print_warning "Caddy نصب نیست، این مرحله رد شد"
fi

# ========================================
# 5. بررسی وضعیت سرویس‌ها
# ========================================
print_info "بررسی وضعیت سرویس‌ها..."

echo ""
print_info "=== وضعیت PM2 ==="
pm2 list

echo ""
print_info "=== وضعیت Backend Service ==="
systemctl status ${PROJECT_NAME}-backend --no-pager | head -10

if systemctl list-unit-files | grep -q caddy.service; then
    echo ""
    print_info "=== وضعیت Caddy ==="
    systemctl status caddy --no-pager | head -10
fi

# ========================================
# 6. تست دسترسی به سرویس‌ها
# ========================================
echo ""
print_info "تست دسترسی به سرویس‌ها..."

sleep 5  # صبر برای بالا آمدن کامل سرویس‌ها

# تست Frontend
if curl -f -s -o /dev/null http://localhost:$FRONTEND_PORT; then
    print_success "Frontend روی پورت $FRONTEND_PORT در دسترس است"
else
    print_error "Frontend روی پورت $FRONTEND_PORT در دسترس نیست"
fi

# تست Backend
if curl -f -s -o /dev/null http://localhost:$BACKEND_PORT/health; then
    print_success "Backend روی پورت $BACKEND_PORT در دسترس است"
else
    print_error "Backend روی پورت $BACKEND_PORT در دسترس نیست"
fi

# ========================================
# خلاصه
# ========================================
echo ""
print_success "========================================="
print_success "راه‌اندازی خودکار با موفقیت تنظیم شد!"
print_success "========================================="
echo ""
print_info "سرویس‌های زیر بعد از ریبوت سرور خودکار بالا می‌آیند:"
echo "  1. PM2 (Frontend: ${PROJECT_NAME}-frontend)"
echo "  2. Backend (systemd: ${PROJECT_NAME}-backend)"
if systemctl list-unit-files | grep -q caddy.service; then
    echo "  3. Caddy (Web Server & Reverse Proxy)"
fi
echo ""
print_info "برای تست ریبوت می‌توانید این دستور را اجرا کنید:"
echo "  sudo reboot"
echo ""
print_info "برای مشاهده لاگ‌ها:"
echo "  pm2 logs ${PROJECT_NAME}-frontend"
echo "  journalctl -u ${PROJECT_NAME}-backend -f"
if systemctl list-unit-files | grep -q caddy.service; then
    echo "  journalctl -u caddy -f"
fi
echo ""
