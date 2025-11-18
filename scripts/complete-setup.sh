#!/bin/bash
# One-Command Complete Setup and Deploy
# این اسکریپت همه چیز رو از صفر تا صد راه‌اندازی میکنه

set -e

echo "=========================================="
echo "🚀 Complete Setup and Deploy"
echo "=========================================="
echo ""

# بررسی که در مسیر درست هستیم
if [ ! -d "/opt/shared" ]; then
    echo "❌ Error: /opt/shared not found!"
    exit 1
fi

cd /opt/shared

# قدم 1: پاکسازی PM2
echo "1️⃣ Cleaning PM2..."
pm2 delete all 2>/dev/null || true
pm2 save --force 2>/dev/null || true
echo "✅ PM2 cleaned"
echo ""

# قدم 2: ایجاد systemd services
echo "2️⃣ Creating systemd services..."

# iranxia-backend
cat > /etc/systemd/system/iranxia-backend.service << 'EOF'
[Unit]
Description=Iranxia Go Backend API
After=network.target postgresql.service postgresql-16.service
Wants=postgresql.service postgresql-16.service

[Service]
Type=simple
User=root
Group=root
WorkingDirectory=/data/iranxia/my-go-backend
EnvironmentFile=-/data/iranxia/my-go-backend/.env
ExecStart=/data/iranxia/my-go-backend/cmd/api
Restart=always
RestartSec=5s
TimeoutStartSec=30s
TimeoutStopSec=30s
LimitNOFILE=65536
LimitNPROC=4096
NoNewPrivileges=true
Environment="GIN_MODE=release"
Environment="ENV=production"

[Install]
WantedBy=multi-user.target
EOF

# rayancomp-backend
cat > /etc/systemd/system/rayancomp-backend.service << 'EOF'
[Unit]
Description=RayanComp Go Backend API
After=network.target postgresql.service postgresql-16.service
Wants=postgresql.service postgresql-16.service

[Service]
Type=simple
User=root
Group=root
WorkingDirectory=/data/rayancomp/my-go-backend
EnvironmentFile=-/data/rayancomp/my-go-backend/.env
ExecStart=/data/rayancomp/my-go-backend/api
Restart=always
RestartSec=5s
TimeoutStartSec=30s
TimeoutStopSec=30s
LimitNOFILE=65536
LimitNPROC=4096
NoNewPrivileges=true
Environment="GIN_MODE=release"
Environment="ENV=production"

[Install]
WantedBy=multi-user.target
EOF

systemctl daemon-reload
systemctl enable iranxia-backend 2>/dev/null || true
systemctl enable rayancomp-backend 2>/dev/null || true

echo "✅ Systemd services created"
echo ""

# قدم 3: ایجاد دایرکتوری‌های لازم
echo "3️⃣ Creating directories..."
mkdir -p /data/iranxia/logs /data/iranxia/my-go-backend/public/uploads
mkdir -p /data/rayancomp/logs /data/rayancomp/my-go-backend/public/uploads
chmod 755 /data/iranxia /data/rayancomp
echo "✅ Directories created"
echo ""

# قدم 4: کپی .env فایل‌های بک‌اند
echo "4️⃣ Setting up backend .env files..."
if [ -f "/opt/shared/my-go-backend/.env" ]; then
    cp /opt/shared/my-go-backend/.env /data/iranxia/my-go-backend/.env 2>/dev/null || true
    cp /opt/shared/my-go-backend/.env /data/rayancomp/my-go-backend/.env 2>/dev/null || true
    echo "✅ Backend .env files copied"
else
    echo "⚠️  Warning: /opt/shared/my-go-backend/.env not found"
fi
echo ""

# قدم 5: Deploy
echo "5️⃣ Running deployment..."
bash /opt/shared/scripts/deploy-all-sites.sh

echo ""
echo "=========================================="
echo "✅ SETUP COMPLETE!"
echo "=========================================="
echo ""
echo "📊 Check status:"
echo "  systemctl status iranxia-backend"
echo "  systemctl status rayancomp-backend"
echo "  pm2 list"
echo ""
echo "📋 View logs:"
echo "  journalctl -u iranxia-backend -f"
echo "  pm2 logs iranxia-frontend"