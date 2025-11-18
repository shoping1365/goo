#!/bin/bash
# اسکریپت تنظیم Auto-Start برای همه پروژه‌ها
# این اسکریپت برای همه پروژه‌ها در /data کار میکنه

set -e

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

print_success() { echo -e "${GREEN}✅ $1${NC}"; }
print_error() { echo -e "${RED}❌ $1${NC}"; }
print_info() { echo -e "${BLUE}ℹ️  $1${NC}"; }
print_warning() { echo -e "${YELLOW}⚠️  $1${NC}"; }

# بررسی root
if [ "$ERID" -ne 0 ]; then 
    print_error "این اسکریپت باید با root اجرا شود"
    exit 1
fi

print_info "========================================="
print_info "تنظیم Auto-Start برای همه پروژه‌ها"
print_info "========================================="

# 1. تنظیم PM2 Startup
print_info "1. تنظیم PM2 برای راه‌اندازی خودکار..."
pm2 startup systemd -u root --hp /root
pm2 save --force
print_success "PM2 تنظیم شد"

# 2. تنظیم هر پروژه
for PROJECT_DIR in /data/iranxia /data/rayancomp; do
    if [ ! -d "$PROJECT_DIR" ]; then
        print_warning "پوشه $PROJECT_DIR وجود ندارد، رد شد"
        continue
    fi
    
    PROJECT_NAME=$(basename "$PROJECT_DIR")
    print_info "2. تنظیم پروژه: $PROJECT_NAME"
    
    # بارگذاری .env
    if [ -f "$PROJECT_DIR/.env" ]; then
        export $(cat "$PROJECT_DIR/.env" | grep -v '^#' | xargs)
    fi
    
    BACKEND_PORT=${BACKEND_PORT:-9090}
    
    # ایجاد systemd service برای Backend
    print_info "   ساخت systemd service برای Backend..."
    
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
LimitNOFILE=65536

[Install]
WantedBy=multi-user.target
EOF
    
    # Reload و فعال‌سازی
    systemctl daemon-reload
    systemctl enable ${PROJECT_NAME}-backend.service
    systemctl restart ${PROJECT_NAME}-backend.service
    
    print_success "   Backend service برای $PROJECT_NAME فعال شد"
    
    # استارت Frontend با PM2
    if [ -f "$PROJECT_DIR/.output/server/index.mjs" ]; then
        print_info "   راه‌اندازی Frontend با PM2..."
        cd "$PROJECT_DIR"
        PROJECT_DIR="$PROJECT_DIR" pm2 start ecosystem.config.cjs
        pm2 save --force
        print_success "   Frontend برای $PROJECT_NAME راه‌اندازی شد"
    else
        print_warning "   فایل .output/server/index.mjs یافت نشد، ابتدا build بگیرید"
    fi
    
    # Reset متغیرها برای پروژه بعدی
    unset BACKEND_PORT FRONTEND_PORT
done

# 3. فعال‌سازی Caddy
print_info "3. فعال‌سازی Caddy..."
if systemctl list-unit-files | grep -q caddy.service; then
    systemctl enable caddy
    systemctl restart caddy
    print_success "Caddy فعال شد"
else
    print_warning "Caddy نصب نیست"
fi

# خلاصه
echo ""
print_success "========================================="
print_success "همه سرویس‌ها برای Auto-Start تنظیم شدند!"
print_success "========================================="
echo ""
print_info "لیست سرویس‌های فعال:"
systemctl list-unit-files | grep -E 'iranxia|rayancomp|caddy|pm2' | grep enabled
echo ""
print_info "وضعیت PM2:"
pm2 list
echo ""
print_info "برای تست ریبوت: sudo reboot"
