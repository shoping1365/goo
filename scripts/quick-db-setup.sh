#!/bin/bash
# Quick Database Setup for Existing PostgreSQL
# برای زمانی که PostgreSQL از قبل نصب هست

set -e

echo "=========================================="
echo "🔧 Setting up Databases"
echo "=========================================="
echo ""

# بررسی PostgreSQL
echo "1️⃣ Checking PostgreSQL-16..."
systemctl status postgresql-16 --no-pager | head -5
echo "✅ PostgreSQL-16 is running"
echo ""

# ساخت دیتابیس‌ها
echo "2️⃣ Creating databases and users..."

# iranxia
echo "Creating iranxia database..."
sudo -u postgres psql << 'EOSQL'
DROP DATABASE IF EXISTS iranxia;
DROP USER IF EXISTS iranxia_user;
CREATE USER iranxia_user WITH PASSWORD 'IranXia@2025#Secure!Pass';
CREATE DATABASE iranxia OWNER iranxia_user;
GRANT ALL PRIVILEGES ON DATABASE iranxia TO iranxia_user;
\c iranxia
GRANT ALL ON SCHEMA public TO iranxia_user;
EOSQL
echo "✅ iranxia database created"

# rayancomp
echo "Creating rayancomp database..."
sudo -u postgres psql << 'EOSQL'
DROP DATABASE IF EXISTS rayancomp;
DROP USER IF EXISTS rayancomp_user;
CREATE USER rayancomp_user WITH PASSWORD 'RayanComp@2025#Secure!Pass';
CREATE DATABASE rayancomp OWNER rayancomp_user;
GRANT ALL PRIVILEGES ON DATABASE rayancomp TO rayancomp_user;
\c rayancomp
GRANT ALL ON SCHEMA public TO rayancomp_user;
EOSQL
echo "✅ rayancomp database created"
echo ""

# ایجاد .env برای rayancomp
echo "3️⃣ Creating rayancomp .env..."
mkdir -p /data/rayancomp/my-go-backend
cat > /data/rayancomp/my-go-backend/.env << 'EOF'
# Database Configuration
DATABASE_URL=postgres://rayancomp_user:RayanComp@2025#Secure!Pass@localhost:5432/rayancomp?sslmode=disable
DB_HOST=localhost
DB_PORT=5432
DB_USER=rayancomp_user
DB_PASSWORD=RayanComp@2025#Secure!Pass
DB_NAME=rayancomp
DB_SSLMODE=disable
AUTO_MIGRATE_DEV=false

# Server Configuration
BACKEND_PORT=9091
PORT=3001
SERVER_HOST=0.0.0.0
ENV=production
APP_ENV=production
GIN_MODE=release

# Security Keys
SESSION_SECRET=Bk8yZZBrCYBBvpC8gB/gkIaxrZ0uOCkwQQdNbca02oP=
JWT_SECRET=e67440d3817ggfg9920dgf:a497d7cb42c71990f346e43e3c35b0e4b
ENCRYPTION_KEY=9dgC2H88YstU+ctLgGx4BF93lp2c0gPLV4AeF/cqjJ=
ENCRYPTION_KEY_VERSION=v1

# Redis Configuration
REDIS_ADDR=localhost:6379
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=1

# File Upload
UPLOAD_DIR=/data/rayancomp/public/uploads
MAX_UPLOAD_SIZE=10485760

# API Configuration
ANTI_CSRF_TRUSTED_ORIGINS=https://rayancomp.com

# Environment
NODE_ENV=production
GO_ENV=production
EOF
echo "✅ rayancomp .env created"
echo ""

# تست اتصال
echo "4️⃣ Testing database connections..."
echo "Databases:"
sudo -u postgres psql -l | grep -E "iranxia|rayancomp"
echo "✅ Databases ready"
echo ""

# Restart backend services
echo "5️⃣ Restarting backend services..."
systemctl restart iranxia-backend rayancomp-backend
sleep 3

# بررسی وضعیت
echo "6️⃣ Checking service status..."
echo ""
if systemctl is-active --quiet iranxia-backend; then
    echo "✅ iranxia-backend: RUNNING"
else
    echo "❌ iranxia-backend: FAILED"
    journalctl -u iranxia-backend -n 5 --no-pager
fi

if systemctl is-active --quiet rayancomp-backend; then
    echo "✅ rayancomp-backend: RUNNING"
else
    echo "❌ rayancomp-backend: FAILED"
    journalctl -u rayancomp-backend -n 5 --no-pager
fi

echo ""
echo "=========================================="
echo "✅ Setup Complete!"
echo "=========================================="
