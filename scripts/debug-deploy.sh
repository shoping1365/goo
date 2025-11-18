#!/bin/bash
# Debug script to find deployment issue
# این اسکریپت مشکل کپی فایل‌ها رو پیدا می‌کنه

set -e

echo "=========================================="
echo "🔍 Debugging Deployment Issue"
echo "=========================================="

# تنظیمات
SHARED_DIR="/opt/shared"
PROJECT_DIR="/data/iranxia"

echo ""
echo "1️⃣ Checking source files in SHARED_DIR..."
echo "----------------------------------------"
if [ -f "${SHARED_DIR}/my-go-backend/cmd/api/main.go" ]; then
    echo "✅ Source main.go exists: ${SHARED_DIR}/my-go-backend/cmd/api/main.go"
    ls -lh "${SHARED_DIR}/my-go-backend/cmd/api/main.go"
else
    echo "❌ Source main.go NOT found: ${SHARED_DIR}/my-go-backend/cmd/api/main.go"
fi

echo ""
echo "2️⃣ Checking directory structure in SHARED_DIR..."
echo "----------------------------------------"
ls -la "${SHARED_DIR}/my-go-backend/" 2>/dev/null || echo "❌ my-go-backend directory not found"
echo ""
ls -la "${SHARED_DIR}/my-go-backend/cmd/" 2>/dev/null || echo "❌ cmd directory not found"
echo ""
ls -la "${SHARED_DIR}/my-go-backend/cmd/api/" 2>/dev/null || echo "❌ cmd/api directory not found"

echo ""
echo "3️⃣ Testing rsync command..."
echo "----------------------------------------"
echo "Running: rsync -av --dry-run --exclude='node_modules' --exclude='.nuxt' --exclude='.output' --exclude='dist' --exclude='my-go-backend/api' --exclude='.git' --exclude='logs' --exclude='.env' '${SHARED_DIR}/my-go-backend/' '${PROJECT_DIR}/my-go-backend/'"

rsync -av --dry-run \
    --exclude='node_modules' \
    --exclude='.nuxt' \
    --exclude='.output' \
    --exclude='dist' \
    --exclude='my-go-backend/api' \
    --exclude='.git' \
    --exclude='logs' \
    --exclude='.env' \
    "${SHARED_DIR}/my-go-backend/" "${PROJECT_DIR}/my-go-backend/" | head -50

echo ""
echo "4️⃣ Checking what would be copied for cmd/api/..."
echo "----------------------------------------"
rsync -av --dry-run \
    --exclude='node_modules' \
    --exclude='my-go-backend/api' \
    "${SHARED_DIR}/my-go-backend/cmd/api/" "${PROJECT_DIR}/my-go-backend/cmd/api/" 2>&1 | head -20

echo ""
echo "5️⃣ Checking current state in PROJECT_DIR..."
echo "----------------------------------------"
if [ -d "${PROJECT_DIR}/my-go-backend" ]; then
    echo "✅ my-go-backend exists in project"
    ls -la "${PROJECT_DIR}/my-go-backend/" 2>/dev/null | head -20
else
    echo "❌ my-go-backend NOT found in project"
fi

echo ""
if [ -f "${PROJECT_DIR}/my-go-backend/cmd/api/main.go" ]; then
    echo "✅ Target main.go exists: ${PROJECT_DIR}/my-go-backend/cmd/api/main.go"
    ls -lh "${PROJECT_DIR}/my-go-backend/cmd/api/main.go"
else
    echo "❌ Target main.go NOT found: ${PROJECT_DIR}/my-go-backend/cmd/api/main.go"
fi

echo ""
echo "6️⃣ Checking file permissions..."
echo "----------------------------------------"
stat "${SHARED_DIR}/my-go-backend/cmd/api/main.go" 2>/dev/null || echo "Cannot stat source file"

echo ""
echo "=========================================="
echo "🔍 Debug Complete"
echo "=========================================="
