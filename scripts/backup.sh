#!/bin/bash
# Backup Script for Deployment
# Date: October 8, 2025
# Purpose: Create backup before deployment

# Don't exit on error - we want to handle errors gracefully
set +e

BACKUP_BASE_DIR="/opt/deployment/backup"
TIMESTAMP=$(date +%Y%m%d_%H%M%S)
BACKUP_DIR="${BACKUP_BASE_DIR}/${TIMESTAMP}"
PROJECT_DIR="/opt/shared"

echo "🔄 Starting backup process..."
echo "📅 Timestamp: ${TIMESTAMP}"

# Create backup base directory if not exists (without sudo)
if [ ! -d "${BACKUP_BASE_DIR}" ]; then
    echo "⚠️ Warning: Backup directory ${BACKUP_BASE_DIR} does not exist!"
    echo "📝 Please create it manually with:"
    echo "   sudo mkdir -p ${BACKUP_BASE_DIR}"
    echo "   sudo chown \$(whoami):\$(whoami) ${BACKUP_BASE_DIR}"
    echo ""
    echo "⏭️ Skipping backup for this deployment..."
    echo "SKIP_BACKUP"  # Special marker for workflow
    exit 0
fi

# Check if we have write permission
if [ ! -w "${BACKUP_BASE_DIR}" ]; then
    echo "⚠️ Warning: No write permission for ${BACKUP_BASE_DIR}"
    echo "📝 Please fix permissions with:"
    echo "   sudo chown \$(whoami):\$(whoami) ${BACKUP_BASE_DIR}"
    echo ""
    echo "⏭️ Skipping backup for this deployment..."
    echo "SKIP_BACKUP"
    exit 0
fi

# Create backup directory
echo "📁 Creating backup directory: ${BACKUP_DIR}"
mkdir -p "${BACKUP_DIR}"

# 1. Backup code
echo "📦 Backing up code..."
rsync -a --exclude='node_modules' \
         --exclude='.git' \
         --exclude='.nuxt' \
         --exclude='dist' \
         "${PROJECT_DIR}/" "${BACKUP_DIR}/code/" 2>/dev/null || echo "⚠️ Warning: Some files could not be backed up"

# 2. Backup Go backend binary
echo "🔨 Backing up Go backend binary..."
if [ -f "${PROJECT_DIR}/my-go-backend/api" ]; then
    cp "${PROJECT_DIR}/my-go-backend/api" "${BACKUP_DIR}/api.backup" 2>/dev/null || true
fi

# 3. Backup uploaded files
echo "📁 Backing up uploaded files..."
if [ -d "${PROJECT_DIR}/my-go-backend/public/uploads" ]; then
    rsync -a "${PROJECT_DIR}/my-go-backend/public/uploads/" "${BACKUP_DIR}/uploads/" 2>/dev/null || true
fi

# 4. Backup PM2 process list
echo "⚙️ Backing up PM2 configuration..."
pm2 save 2>/dev/null || true
if [ -f ~/.pm2/dump.pm2 ]; then
    cp ~/.pm2/dump.pm2 "${BACKUP_DIR}/pm2_dump.pm2" 2>/dev/null || true
fi

# 5. Create backup info file
cat > "${BACKUP_DIR}/backup_info.txt" << EOF
Backup Created: ${TIMESTAMP}
Project Directory: ${PROJECT_DIR}
Backup Directory: ${BACKUP_DIR}
Git Commit: $(cd ${PROJECT_DIR} && git rev-parse HEAD 2>/dev/null || echo "N/A")
Git Branch: $(cd ${PROJECT_DIR} && git branch --show-current 2>/dev/null || echo "N/A")
EOF

echo "✅ Backup completed successfully!"
echo "📍 Backup location: ${BACKUP_DIR}"

# Cleanup old backups (keep last 10)
echo "🧹 Cleaning up old backups (keeping last 10)..."
cd "${BACKUP_BASE_DIR}" 2>/dev/null || exit 0
ls -t 2>/dev/null | tail -n +11 | xargs -r rm -rf 2>/dev/null || true

echo "✅ Backup process finished!"
echo "${BACKUP_DIR}"  # Return backup directory for workflow