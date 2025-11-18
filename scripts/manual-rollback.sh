#!/bin/bash
# Manual Rollback Script
# اسکریپت دستی برای بازگردانی از بک‌آپ
# استفاده: bash scripts/manual-rollback.sh [BACKUP_TIMESTAMP]

set -e

# Load configuration
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${SCRIPT_DIR}/sites-config.sh"

# تابع برای نمایش لیست بک‌آپ‌ها
list_backups() {
    echo "=========================================="
    echo "📦 Available Backups"
    echo "=========================================="
    
    if [ ! -d "${BACKUP_BASE_DIR}" ]; then
        print_error "Backup directory does not exist: ${BACKUP_BASE_DIR}"
        return 1
    fi
    
    local backups=($(ls -t "${BACKUP_BASE_DIR}" 2>/dev/null))
    
    if [ ${#backups[@]} -eq 0 ]; then
        print_warning "No backups found"
        return 1
    fi
    
    local i=1
    for backup in "${backups[@]}"; do
        echo "${i}. ${backup}"
        if [ -f "${BACKUP_BASE_DIR}/${backup}/backup_info.txt" ]; then
            echo "   $(grep 'Git Commit' "${BACKUP_BASE_DIR}/${backup}/backup_info.txt" || echo '')"
        fi
        i=$((i + 1))
    done
    echo "=========================================="
}

# تابع برای انتخاب بک‌آپ
select_backup() {
    list_backups
    
    echo ""
    echo "Enter backup number or timestamp (or press Enter for latest):"
    read -r selection
    
    if [ -z "$selection" ]; then
        # استفاده از آخرین بک‌آپ
        SELECTED_BACKUP=$(ls -t "${BACKUP_BASE_DIR}" 2>/dev/null | head -1)
    elif [[ "$selection" =~ ^[0-9]+$ ]]; then
        # انتخاب بر اساس شماره
        SELECTED_BACKUP=$(ls -t "${BACKUP_BASE_DIR}" 2>/dev/null | sed -n "${selection}p")
    else
        # استفاده مستقیم از timestamp
        SELECTED_BACKUP="$selection"
    fi
    
    if [ -z "$SELECTED_BACKUP" ]; then
        print_error "Invalid selection"
        return 1
    fi
    
    if [ ! -d "${BACKUP_BASE_DIR}/${SELECTED_BACKUP}" ]; then
        print_error "Backup not found: ${SELECTED_BACKUP}"
        return 1
    fi
    
    echo "$SELECTED_BACKUP"
}

# Main Script
echo "=========================================="
echo "🔄 Manual Rollback Tool"
echo "=========================================="
echo ""

# دریافت بک‌آپ مورد نظر
if [ -n "$1" ]; then
    BACKUP_TIMESTAMP="$1"
    if [ ! -d "${BACKUP_BASE_DIR}/${BACKUP_TIMESTAMP}" ]; then
        print_error "Backup not found: ${BACKUP_TIMESTAMP}"
        exit 1
    fi
else
    BACKUP_TIMESTAMP=$(select_backup)
    if [ $? -ne 0 ]; then
        exit 1
    fi
fi

BACKUP_DIR="${BACKUP_BASE_DIR}/${BACKUP_TIMESTAMP}"

echo ""
print_info "Selected Backup: ${BACKUP_TIMESTAMP}"
echo "📂 Location: ${BACKUP_DIR}"
echo ""

# نمایش اطلاعات بک‌آپ
if [ -f "${BACKUP_DIR}/backup_info.txt" ]; then
    print_info "Backup Information:"
    cat "${BACKUP_DIR}/backup_info.txt"
    echo ""
fi

# تأیید کاربر
echo "⚠️  WARNING: This will restore all sites from the selected backup!"
echo "Current running sites will be stopped and replaced."
echo ""
read -p "Are you sure you want to continue? (yes/no): " confirm

if [ "$confirm" != "yes" ]; then
    print_warning "Rollback cancelled by user"
    exit 0
fi

echo ""
echo "=========================================="
echo "🔄 Starting Rollback Process"
echo "=========================================="

ROLLBACK_SUCCESS=true

# Rollback هر پروژه
for project_dir in "${PROJECTS[@]}"; do
    parse_project_config "$project_dir"
    
    echo ""
    echo "----------------------------------------"
    print_info "Rolling back: ${PROJECT_NAME}"
    echo "----------------------------------------"
    
    PROJECT_BACKUP_DIR="${BACKUP_DIR}/${PROJECT_NAME}"
    
    if [ ! -d "${PROJECT_BACKUP_DIR}" ]; then
        print_warning "No backup found for ${PROJECT_NAME}, skipping..."
        continue
    fi
    
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
    
    # بازگردانی uploads
    if [ -d "${PROJECT_BACKUP_DIR}/uploads" ]; then
        print_info "Restoring uploaded files..."
        mkdir -p "${PROJECT_DIR}/my-go-backend/public/uploads"
        rsync -a "${PROJECT_BACKUP_DIR}/uploads/" \
            "${PROJECT_DIR}/my-go-backend/public/uploads/" || true
        print_success "Uploads restored"
    fi
    
    # بازگردانی database
    if [ -f "${PROJECT_BACKUP_DIR}/database.db.backup" ]; then
        print_info "Restoring database..."
        cp "${PROJECT_BACKUP_DIR}/database.db.backup" \
            "${PROJECT_DIR}/my-go-backend/database.db" || true
        print_success "Database restored"
    fi
    
    # بازگردانی .env
    if [ -f "${PROJECT_BACKUP_DIR}/.env.backup" ]; then
        print_info "Restoring environment..."
        cp "${PROJECT_BACKUP_DIR}/.env.backup" "${PROJECT_DIR}/.env" || true
    fi
    
    # Restart سرویس‌ها
    print_info "Restarting services..."
    pm2 restart "${PM2_FRONTEND}" 2>/dev/null || true
    pm2 restart "${PM2_BACKEND}" 2>/dev/null || true
    
    print_success "${PROJECT_NAME} rolled back successfully!"
done

echo ""
echo "=========================================="
if [ "$ROLLBACK_SUCCESS" = true ]; then
    print_success "ROLLBACK COMPLETED SUCCESSFULLY!"
    echo ""
    echo "📊 PM2 Status:"
    pm2 list
    echo ""
    print_info "Rollback Summary:"
    echo "📂 Restored from: ${BACKUP_DIR}"
    echo "📅 Backup date: ${BACKUP_TIMESTAMP}"
    echo "🌐 Projects restored: ${#PROJECTS[@]}"
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