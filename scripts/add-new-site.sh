#!/bin/bash
# اضافه کردن سایت جدید به صورت خودکار
# استفاده: bash add-new-site.sh SITE_NAME FRONTEND_PORT BACKEND_PORT

set -e

if [ "$#" -ne 3 ]; then
    echo "Usage: bash add-new-site.sh SITE_NAME FRONTEND_PORT BACKEND_PORT"
    echo "Example: bash add-new-site.sh site3 3002 9092"
    exit 1
fi

SITE_NAME=$1
FRONTEND_PORT=$2
BACKEND_PORT=$3
SITE_DIR="/data/${SITE_NAME}"

echo "=========================================="
echo "🚀 Adding New Site: $SITE_NAME"
echo "=========================================="
echo "Frontend Port: $FRONTEND_PORT"
echo "Backend Port: $BACKEND_PORT"
echo "Directory: $SITE_DIR"
echo ""

# قدم 1: ساخت دایرکتوری
echo "1️⃣ Creating directories..."
mkdir -p ${SITE_DIR}/my-go-backend
mkdir -p ${SITE_DIR}/logs
chmod 755 ${SITE_DIR}
echo "✅ Directories created"
echo ""

# قدم 2: ایجاد .env Frontend
echo "2️⃣ Creating frontend .env..."
cat > ${SITE_DIR}/.env << EOF
PORT=${FRONTEND_PORT}
BACKEND_PORT=${BACKEND_PORT}
NUXT_PUBLIC_GO_API_BASE=http://localhost:${BACKEND_PORT}
NUXT_PUBLIC_WS_BASE=ws://localhost:${BACKEND_PORT}
EOF
echo "✅ Frontend .env created"
echo ""

# قدم 3: ایجاد .env Backend
echo "3️⃣ Creating backend .env..."
DB_PASSWORD="${SITE_NAME^}@2025#Secure!Pass"
cat > ${SITE_DIR}/my-go-backend/.env << EOF
DATABASE_URL=postgres://${SITE_NAME}_user:${DB_PASSWORD}@localhost:5432/${SITE_NAME}?sslmode=disable
DB_HOST=localhost
DB_PORT=5432
DB_USER=${SITE_NAME}_user
DB_PASSWORD=${DB_PASSWORD}
DB_NAME=${SITE_NAME}
DB_SSLMODE=disable
AUTO_MIGRATE_DEV=false

SERVER_PORT=${BACKEND_PORT}
PORT=${BACKEND_PORT}
SERVER_HOST=0.0.0.0
ENV=production
APP_ENV=production
GIN_MODE=release

SESSION_SECRET=$(openssl rand -base64 32)
JWT_SECRET=$(openssl rand -hex 32)
ENCRYPTION_KEY=$(openssl rand -base64 32)
ENCRYPTION_KEY_VERSION=v1

REDIS_ADDR=localhost:6379
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=${BACKEND_PORT: -1}

UPLOAD_DIR=${SITE_DIR}/uploads
MAX_UPLOAD_SIZE=10485760

ANTI_CSRF_TRUSTED_ORIGINS=https://${SITE_NAME}.com

NODE_ENV=production
GO_ENV=production
EOF
echo "✅ Backend .env created"
echo ""

# قدم 4: ساخت دیتابیس
echo "4️⃣ Creating PostgreSQL database..."
sudo -u postgres psql << EOSQL
DROP DATABASE IF EXISTS ${SITE_NAME};
DROP USER IF EXISTS ${SITE_NAME}_user;
CREATE USER ${SITE_NAME}_user WITH PASSWORD '${DB_PASSWORD}';
CREATE DATABASE ${SITE_NAME} OWNER ${SITE_NAME}_user;
GRANT ALL PRIVILEGES ON DATABASE ${SITE_NAME} TO ${SITE_NAME}_user;
\c ${SITE_NAME}
GRANT ALL ON SCHEMA public TO ${SITE_NAME}_user;
EOSQL
echo "✅ Database created"
echo ""

# قدم 5: ساخت Systemd Service
echo "5️⃣ Creating systemd service..."
cat > /etc/systemd/system/${SITE_NAME}-backend.service << EOF
[Unit]
Description=${SITE_NAME^} Go Backend API
After=network.target postgresql-16.service

[Service]
Type=simple
WorkingDirectory=${SITE_DIR}/my-go-backend
EnvironmentFile=-${SITE_DIR}/my-go-backend/.env
ExecStart=${SITE_DIR}/my-go-backend/cmd/api
Restart=always
RestartSec=5s

[Install]
WantedBy=multi-user.target
EOF

systemctl daemon-reload
systemctl enable ${SITE_NAME}-backend
echo "✅ Systemd service created"
echo ""

# قدم 6: اضافه کردن به sites-config.sh
echo "6️⃣ Adding to sites-config.sh..."
if ! grep -q "${SITE_DIR}" /opt/shared/scripts/sites-config.sh; then
    sed -i "/^PROJECTS=($/a\    \"${SITE_DIR}\"" /opt/shared/scripts/sites-config.sh
    echo "✅ Added to config"
else
    echo "⚠️  Already in config"
fi
echo ""

echo "=========================================="
echo "✅ Site Added Successfully!"
echo "=========================================="
echo ""
echo "📊 Site Information:"
echo "  Name: $SITE_NAME"
echo "  Frontend Port: $FRONTEND_PORT"
echo "  Backend Port: $BACKEND_PORT"
echo "  Database: $SITE_NAME"
echo "  Directory: $SITE_DIR"
echo ""
echo "🚀 Next Steps:"
echo "  1. Deploy: bash /opt/shared/scripts/deploy-all-sites.sh"
echo "  2. Check status: systemctl status ${SITE_NAME}-backend"
echo "  3. Check frontend: pm2 list | grep ${SITE_NAME}"
echo ""