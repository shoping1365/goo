#!/bin/bash
# 🚀 ONE-COMMAND COMPLETE DEPLOYMENT
# این اسکریپت همه چیز رو از صفر تا صد راه‌اندازی میکنه - بدون هیچ دخالت دستی!

set -e

echo "=========================================="
echo "🚀 COMPLETE AUTOMATED DEPLOYMENT"
echo "=========================================="
echo "Starting at: $(date)"
echo ""

# تابع‌های کمکی
print_success() { echo -e "\033[0;32m✅ $1\033[0m"; }
print_error() { echo -e "\033[0;31m❌ $1\033[0m"; }
print_info() { echo -e "\033[0;34mℹ️  $1\033[0m"; }

# مرحله 1: پاکسازی PM2
print_info "Step 1: Cleaning PM2..."
pm2 delete all 2>/dev/null || true
pm2 save --force 2>/dev/null || true
print_success "PM2 cleaned"
echo ""

# مرحله 2: به‌روزرسانی سرور
print_info "Step 2: Updating server..."
apt-get update -y && apt-get upgrade -y
print_success "Server updated"
echo ""

# مرحله 3: نصب وابستگی‌ها
print_info "Step 3: Installing dependencies..."
apt-get install -y \
  postgresql-16 \
  postgresql-client-16 \
  golang-go \
  pm2 \
  nginx \
  certbot \
  python3-certbot-nginx
print_success "Dependencies installed"
echo ""

# مرحله 4: ساخت Systemd Services
print_info "Step 4: Creating systemd services..."

cat > /etc/systemd/system/iranxia-backend.service << 'EOF'
[Unit]
Description=Iranxia Go Backend API
After=network.target postgresql-16.service

[Service]
Type=simple
WorkingDirectory=/data/iranxia/my-go-backend
EnvironmentFile=-/data/iranxia/my-go-backend/.env
ExecStart=/data/iranxia/my-go-backend/cmd/api
Restart=always
RestartSec=5s

[Install]
WantedBy=multi-user.target
EOF

cat > /etc/systemd/system/rayancomp-backend.service << 'EOF'
[Unit]
Description=RayanComp Go Backend API
After=network.target postgresql-16.service

[Service]
Type=simple
WorkingDirectory=/data/rayancomp/my-go-backend
EnvironmentFile=-/data/rayancomp/my-go-backend/.env
ExecStart=/data/rayancomp/my-go-backend/api
Restart=always
RestartSec=5s

[Install]
WantedBy=multi-user.target
EOF

systemctl daemon-reload
systemctl enable iranxia-backend rayancomp-backend
print_success "Systemd services created"
echo ""

# مرحله 5: راه‌اندازی Nginx
print_info "Step 5: Setting up Nginx..."

cat > /etc/nginx/sites-available/iranxia << 'EOF'
server {
    listen 80;
    server_name iranxia.com www.iranxia.com;

    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
EOF

cat > /etc/nginx/sites-available/rayancomp << 'EOF'
server {
    listen 80;
    server_name rayancomp.com www.rayancomp.com;

    location / {
        proxy_pass http://localhost:3001;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
EOF

ln -s /etc/nginx/sites-available/iranxia /etc/nginx/sites-enabled/
ln -s /etc/nginx/sites-available/rayancomp /etc/nginx/sites-enabled/

nginx -t
systemctl restart nginx
print_success "Nginx set up"
echo ""

# مرحله 6: صدور گواهی SSL
print_info "Step 6: Issuing SSL certificates..."
certbot --nginx -d iranxia.com -d www.iranxia.com
certbot --nginx -d rayancomp.com -d www.rayancomp.com
print_success "SSL certificates issued"
echo ""

# مرحله 7: ساخت پایگاه داده
print_info "Step 7: Setting up databases..."

su - postgres -c "createuser --interactive"
su - postgres -c "createdb iranxia_production"
su - postgres -c "createdb rayancomp_production"

print_success "Databases set up"
echo ""

# مرحله 8: راه‌اندازی مجدد سرورها
print_info "Step 8: Starting servers..."
systemctl start iranxia-backend rayancomp-backend
print_success "Servers started"
echo ""

# مرحله 9: بررسی وضعیت
print_info "Step 9: Checking status..."
systemctl status iranxia-backend rayancomp-backend
print_success "Status checked"
echo ""

echo "=========================================="
echo "🚀 DEPLOYMENT COMPLETED"
echo "=========================================="
echo "Finished at: $(date)"
