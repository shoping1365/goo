#!/bin/bash
# Complete Database and Service Setup
# این اسکریپت دیتابیس رو میسازه و سرویس‌ها رو راه‌اندازی میکنه

set -e

echo "=========================================="
echo "🔧 Complete Setup: Database + Services"
echo "=========================================="
echo ""

# قدم 1: بررسی PostgreSQL
echo "1️⃣ Checking PostgreSQL..."
if ! systemctl is-active --quiet postgresql postgresql-16; then
    echo "⚠️  PostgreSQL is not running. Starting..."
    systemctl start postgresql postgresql-16 2>/dev/null || systemctl start postgresql 2>/dev/null || systemctl start postgresql-16
    sleep 2
fi
systemctl status postgresql postgresql-16 --no-pager | head -5 || systemctl status postgresql --no-pager | head -5
echo "✅ PostgreSQL is running"
echo ""

# قدم 2: ساخت دیتابیس و یوزر برای iranxia
echo "2️⃣ Setting up iranxia database..."
sudo -u postgres psql << 'EOSQL'
-- Drop if exists (for clean setup)
DROP DATABASE IF EXISTS iranxia;
DROP USER IF EXISTS iranxia_user;

-- Create user
CREATE USER iranxia_user WITH PASSWORD 'IranXia@2025#Secure!Pass';

-- Create database
CREATE DATABASE iranxia OWNER iranxia_user;

-- Grant privileges
GRANT ALL PRIVILEGES ON DATABASE iranxia TO iranxia_user;

\c iranxia
GRANT ALL ON SCHEMA public TO iranxia_user;

-- نصب extension pgcrypto (ضروری!)
CREATE EXTENSION IF NOT EXISTS pgcrypto;

\q
EOSQL
echo "✅ iranxia database created with pgcrypto extension"
echo ""

# قدم 3: ساخت دیتابیس و یوزر برای rayancomp
echo "3️⃣ Setting up rayancomp database..."
sudo -u postgres psql << 'EOSQL'
-- Drop if exists (for clean setup)
DROP DATABASE IF EXISTS rayancomp;
DROP USER IF EXISTS rayancomp_user;

-- Create user
CREATE USER rayancomp_user WITH PASSWORD 'RayanComp@2025#Secure!Pass';

-- Create database
CREATE DATABASE rayancomp OWNER rayancomp_user;

-- Grant privileges
GRANT ALL PRIVILEGES ON DATABASE rayancomp TO rayancomp_user;

\c rayancomp
GRANT ALL ON SCHEMA public TO rayancomp_user;

-- نصب extension pgcrypto (ضروری!)
CREATE EXTENSION IF NOT EXISTS pgcrypto;

\q
EOSQL
echo "✅ rayancomp database created with pgcrypto extension"
echo ""

# قدم 4: ایجاد .env برای rayancomp
echo "4️⃣ Creating rayancomp .env..."
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
SERVER_PORT=9091
PORT=9091
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
UPLOAD_DIR=/data/rayancomp/uploads
MAX_UPLOAD_SIZE=10485760

# API Configuration
ANTI_CSRF_TRUSTED_ORIGINS=https://rayancomp.net

# Environment
NODE_ENV=production
GO_ENV=production
EOF
echo "✅ rayancomp .env created"
echo ""

# قدم 5: تست اتصال به دیتابیس
echo "5️⃣ Testing database connections..."
echo "Testing iranxia database:"
sudo -u postgres psql -d iranxia -c "SELECT current_database(), current_user;" || echo "❌ Failed"
echo "Testing pgcrypto extension:"
sudo -u postgres psql -d iranxia -c "SELECT gen_random_uuid();" || echo "❌ pgcrypto failed"
echo ""
echo "Testing rayancomp database:"
sudo -u postgres psql -d rayancomp -c "SELECT current_database(), current_user;" || echo "❌ Failed"
echo "Testing pgcrypto extension:"
sudo -u postgres psql -d rayancomp -c "SELECT gen_random_uuid();" || echo "❌ pgcrypto failed"
echo ""

# قدم 6: Restart backend services
echo "6️⃣ Restarting backend services..."
systemctl restart iranxia-backend rayancomp-backend
sleep 3

# قدم 7: بررسی وضعیت
echo "7️⃣ Checking service status..."
echo ""
echo "=== iranxia-backend ==="
systemctl status iranxia-backend --no-pager -n 0 || journalctl -u iranxia-backend -n 5 --no-pager
echo ""
echo "=== rayancomp-backend ==="
systemctl status rayancomp-backend --no-pager -n 0 || journalctl -u rayancomp-backend -n 5 --no-pager

echo ""
echo "=========================================="
if systemctl is-active --quiet iranxia-backend && systemctl is-active --quiet rayancomp-backend; then
    echo "✅ ALL SERVICES ARE RUNNING!"
    echo ""
    echo "🌐 URLs:"
    echo "  iranxia frontend: http://localhost:3000"
    echo "  iranxia backend: http://localhost:9090"
    echo "  rayancomp frontend: http://localhost:3001"
    echo "  rayancomp backend: http://localhost:9091"
else
    echo "⚠️  Some services failed. Check logs:"
    echo "  journalctl -u iranxia-backend -f"
    echo "  journalctl -u rayancomp-backend -f"
fi
echo "=========================================="
