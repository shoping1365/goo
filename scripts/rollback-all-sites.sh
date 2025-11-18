#!/bin/bash
# Rollback All Production Projects
# تاریخ: 8 اکتبر 2025
# بازگردانی تمام پروژه‌ها از بک‌آپ

set -e

# Load configuration
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${SCRIPT_DIR}/sites-config.sh"

# دریافت آخرین بک‌آپ
LATEST_BACKUP=$(ls -t "${BACKUP_BASE_DIR}" 2>/dev/null | head -1)

if [ -z "$LATEST_BACKUP" ]; then
    print_error "No backup found in ${BACKUP_BASE_DIR}"
    exit 1
fi

BACKUP_DIR="${BACKUP_BASE_DIR}/${LATEST_BACKUP}"

echo "=========================================="
echo "🔄 Starting Rollback Process"
echo "=========================================="
echo "📂 Backup Source: ${BACKUP_DIR}"
echo "📅 Backup Date: ${LATEST_BACKUP}"
echo ""

# خواندن اطلاعات بک‌آپ
if [ -f "${BACKUP_DIR}/backup_info.txt" ]; then
    print_info "Backup Information:"
    cat "${BACKUP_DIR}/backup_info.txt"
    echo ""
fi

ROLLBACK_SUCCESS=true

# Rollback هر پروژه
for project_dir in "${PROJECTS[@]}"; do
    parse_project_config "$project_dir"
    
    echo "=========================================="
    print_info "Rolling back: ${PROJECT_NAME}"
    echo "=========================================="
    
    PROJECT_BACKUP_DIR="${BACKUP_DIR}/${PROJECT_NAME}"
    
    # بررسی وجود بک‌آپ
    if [ ! -d "${PROJECT_BACKUP_DIR}" ]; then
        print_warning "No backup found for ${PROJECT_NAME}, skipping..."
        continue
    fi
    
    # بررسی وضعیت بک‌آپ
    if [ -f "${PROJECT_BACKUP_DIR}/status.txt" ]; then
        BACKUP_STATUS=$(cat "${PROJECT_BACKUP_DIR}/status.txt")
        if [ "$BACKUP_STATUS" = "skip" ]; then
            print_warning "Backup was skipped for ${PROJECT_NAME}, skipping rollback..."
            continue
        fi
    fi
    
    # خواندن اطلاعات پروژه از بک‌آپ
    if [ -f "${PROJECT_BACKUP_DIR}/.env.backup" ]; then
        read_env_file "${PROJECT_BACKUP_DIR}/.env.backup"
    fi
    
    # توقف سرویس‌ها
    print_info "Stopping services..."
    pm2 stop "${PM2_FRONTEND}" 2>/dev/null || true
    pm2 stop "${PM2_BACKEND}" 2>/dev/null || true
    
    # بازگردانی فایل‌ها
    print_info "Restoring files..."
    if [ -d "${PROJECT_BACKUP_DIR}/project" ]; then
        rsync -a --delete "${PROJECT_BACKUP_DIR}/project/" "${PROJECT_DIR}/" || {
            print_error "Failed to restore files for ${PROJECT_NAME}"
            ROLLBACK_SUCCESS=false
            continue
        }
        print_success "Files restored"
    fi
    
    # بازگردانی فایل‌های آپلود شده
    if [ -d "${PROJECT_BACKUP_DIR}/uploads" ]; then
        print_info "Restoring uploaded files..."
        mkdir -p "${PROJECT_DIR}/my-go-backend/public/uploads"
        rsync -a "${PROJECT_BACKUP_DIR}/uploads/" \
            "${PROJECT_DIR}/my-go-backend/public/uploads/" || true
        print_success "Uploads restored"
    fi
    
    # بازگردانی دیتابیس
    if [ -f "${PROJECT_BACKUP_DIR}/database.db.backup" ]; then
        print_info "Restoring database..."
        cp "${PROJECT_BACKUP_DIR}/database.db.backup" \
            "${PROJECT_DIR}/my-go-backend/database.db" || true
        print_success "Database restored"
    fi
    
    # بازگردانی .env
    if [ -f "${PROJECT_BACKUP_DIR}/.env.backup" ]; then
        print_info "Restoring environment file..."
        cp "${PROJECT_BACKUP_DIR}/.env.backup" "${PROJECT_DIR}/.env" || true
    fi
    
    # Restart سرویس‌ها
    print_info "Restarting services..."
    pm2 restart "${PM2_FRONTEND}" 2>/dev/null || true
    pm2 restart "${PM2_BACKEND}" 2>/dev/null || true
    
    print_success "✅ ${PROJECT_NAME} rolled back successfully!"
    echo ""
done

echo "=========================================="
if [ "$ROLLBACK_SUCCESS" = true ]; then
    print_success "ALL PROJECTS ROLLED BACK SUCCESSFULLY!"
    echo ""
    echo "📊 PM2 Status:"
    pm2 list
    echo ""
    print_info "Rollback Details:"
    echo "📂 Restored from: ${BACKUP_DIR}"
    echo "📅 Backup timestamp: ${LATEST_BACKUP}"
    echo "=========================================="
    exit 0
else
    print_error "SOME ROLLBACKS FAILED!"
    echo ""
    echo "📊 PM2 Status:"
    pm2 list
    echo "=========================================="
    exit 1
fi