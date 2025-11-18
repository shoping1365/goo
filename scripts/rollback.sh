#!/bin/bash
# Rollback Script for Deployment
# Date: October 8, 2025
# Purpose: Restore from backup if deployment fails

# Don't exit on error - we want to handle errors gracefully
set +e

BACKUP_BASE_DIR="/opt/deployment/backup"
TIMESTAMP=$(date +%Y%m%d_%H%M%S)
BACKUP_DIR="${BACKUP_BASE_DIR}/${TIMESTAMP}"
PROJECT_DIR="/opt/shared"

if [ -z "$BACKUP_DIR" ]; then
    echo "❌ Error: Backup directory not specified!"
    echo "Usage: $0 <backup_directory>"
    exit 1
fi

if [ ! -d "$BACKUP_DIR" ]; then
    echo "❌ Error: Backup directory does not exist: $BACKUP_DIR"
    echo "🔍 Looking for latest backup..."

    # Try to find latest backup
    LATEST_BACKUP=$(ls -t /opt/deployment/backup 2>/dev/null | head -1)
    if [ ! -z "$LATEST_BACKUP" ]; then
        BACKUP_DIR="/opt/deployment/backup/$LATEST_BACKUP"
        echo "📦 Found latest backup: $BACKUP_DIR"
    else
        echo "❌ No backups found!"
        exit 1
    fi
fi

echo "🔄 Starting rollback process..."
echo "📂 Restoring from: $BACKUP_DIR"

# 1. Stop services
echo "⏸️ Stopping services..."
pm2 stop all 2>/dev/null || true

# 2. Restore code
echo "📦 Restoring code..."
if [ -d "${BACKUP_DIR}/code" ]; then
    rsync -a --delete "${BACKUP_DIR}/code/" "${PROJECT_DIR}/"
else
    echo "⚠️ Warning: Code backup not found, skipping..."
fi

# 3. Restore Go backend binary
echo "🔨 Restoring Go backend binary..."
if [ -f "${BACKUP_DIR}/api.backup" ]; then
    cp "${BACKUP_DIR}/api.backup" "${PROJECT_DIR}/my-go-backend/api"
    chmod +x "${PROJECT_DIR}/my-go-backend/api"
fi

# 4. Restore uploaded files
echo "📁 Restoring uploaded files..."
if [ -d "${BACKUP_DIR}/uploads" ]; then
    rsync -a "${BACKUP_DIR}/uploads/" "${PROJECT_DIR}/my-go-backend/public/uploads/"
fi

# 5. Restore database (if exists)
echo "💾 Restoring database..."
# TODO: اگر از PostgreSQL یا MySQL استفاده می‌کنی، اینجا restore کن
# Example for PostgreSQL:
# psql -U username dbname < "${BACKUP_DIR}/database.sql"

# 6. Restore PM2 configuration
echo "⚙️ Restoring PM2 configuration..."
if [ -f "${BACKUP_DIR}/pm2_dump.pm2" ]; then
    cp "${BACKUP_DIR}/pm2_dump.pm2" ~/.pm2/dump.pm2
    pm2 resurrect
fi

# 7. Restart services
echo "🚀 Restarting services..."
pm2 restart all

echo "✅ Rollback completed successfully!"
echo "📍 Restored from: $BACKUP_DIR"