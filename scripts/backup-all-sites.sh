#!/bin/bash
# Backup All Production Projects
# تاریخ: 8 اکتبر 2025

set -e

# Load configuration
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${SCRIPT_DIR}/sites-config.sh"

TIMESTAMP=$(date +%Y%m%d_%H%M%S)
BACKUP_ROOT="${BACKUP_BASE_DIR}/${TIMESTAMP}"

echo "=========================================="
echo "🔄 Starting Backup Process"
echo "=========================================="
echo "📅 Timestamp: ${TIMESTAMP}"
echo "📂 Backup Root: ${BACKUP_ROOT}"
echo ""

# ایجاد دایرکتوری اصلی بک‌آپ
mkdir -p "${BACKUP_ROOT}"

# ذخیره اطلاعات کلی
cat > "${BACKUP_ROOT}/backup_info.txt" << EOF
Backup Created: ${TIMESTAMP}
Shared Directory: ${SHARED_DIR}
Git Commit: $(cd ${SHARED_DIR} 2>/dev/null && git rev-parse HEAD 2>/dev/null || echo "N/A")
Git Branch: $(cd ${SHARED_DIR} 2>/dev/null && git branch --show-current 2>/dev/null || echo "N/A")
Number of Projects: ${#PROJECTS[@]}
EOF

BACKUP_SUCCESS=true

# بک‌آپ هر پروژه
for project_dir in "${PROJECTS[@]}"; do
    parse_project_config "$project_dir"
    
    echo "----------------------------------------"
    print_info "Backing up: ${PROJECT_NAME}"
    echo "📁 Directory: ${PROJECT_DIR}"
    
    PROJECT_BACKUP_DIR="${BACKUP_ROOT}/${PROJECT_NAME}"
    mkdir -p "${PROJECT_BACKUP_DIR}"
    
    if [ ! -d "${PROJECT_DIR}" ]; then
        print_warning "Project directory does not exist: ${PROJECT_DIR}"
        echo "skip" > "${PROJECT_BACKUP_DIR}/status.txt"
        continue
    fi
    
    # خواندن .env برای اطلاعات
    if [ -f "${PROJECT_DIR}/.env" ]; then
        read_env_file "${PROJECT_DIR}/.env"
    fi
    
    # بک‌آپ کد و فایل‌های بیلد شده
    print_info "Backing up built files..."
    rsync -a \
        --exclude='node_modules' \
        --exclude='.git' \
        "${PROJECT_DIR}/" "${PROJECT_BACKUP_DIR}/project/" 2>/dev/null || {
        print_error "Failed to backup project files"
        BACKUP_SUCCESS=false
        continue
    }
    
    # بک‌آپ فایل‌های آپلود شده
    if [ -d "${PROJECT_DIR}/my-go-backend/public/uploads" ]; then
        print_info "Backing up uploaded files..."
        rsync -a "${PROJECT_DIR}/my-go-backend/public/uploads/" \
            "${PROJECT_BACKUP_DIR}/uploads/" 2>/dev/null || true
    fi
    
    # بک‌آپ دیتابیس SQLite (اگر وجود داره)
    if [ -f "${PROJECT_DIR}/my-go-backend/database.db" ]; then
        print_info "Backing up SQLite database..."
        cp "${PROJECT_DIR}/my-go-backend/database.db" \
            "${PROJECT_BACKUP_DIR}/database.db.backup" 2>/dev/null || true
    fi
    
    # ذخیره وضعیت PM2
    print_info "Saving PM2 status..."
    pm2 describe "${PM2_FRONTEND}" > "${PROJECT_BACKUP_DIR}/pm2_frontend.txt" 2>/dev/null || true
    pm2 describe "${PM2_BACKEND}" > "${PROJECT_BACKUP_DIR}/pm2_backend.txt" 2>/dev/null || true
    
    # ذخیره .env
    if [ -f "${PROJECT_DIR}/.env" ]; then
        cp "${PROJECT_DIR}/.env" "${PROJECT_BACKUP_DIR}/.env.backup" 2>/dev/null || true
    fi
    
    # ثبت اطلاعات بک‌آپ پروژه
    cat > "${PROJECT_BACKUP_DIR}/project_info.txt" << EOF
Project Name: ${PROJECT_NAME}
Project Directory: ${PROJECT_DIR}
Frontend Port: ${FRONTEND_PORT:-N/A}
Backend Port: ${BACKEND_PORT:-N/A}
PM2 Frontend: ${PM2_FRONTEND}
PM2 Backend: ${PM2_BACKEND}
Backup Time: ${TIMESTAMP}
EOF
    
    echo "success" > "${PROJECT_BACKUP_DIR}/status.txt"
    print_success "Backup completed for ${PROJECT_NAME}"
done

echo ""
echo "=========================================="

if [ "$BACKUP_SUCCESS" = true ]; then
    print_success "All Projects Backed Up Successfully!"
    echo "📍 Backup Location: ${BACKUP_ROOT}"
    
    # پاکسازی بک‌آپ‌های قدیمی
    echo ""
    print_info "Cleaning up old backups (keeping last ${KEEP_BACKUPS})..."
    cd "${BACKUP_BASE_DIR}" 2>/dev/null || exit 0
    ls -t 2>/dev/null | tail -n +$((KEEP_BACKUPS + 1)) | xargs -r rm -rf 2>/dev/null || true
    print_success "Cleanup completed"
    
    echo "=========================================="
    exit 0
else
    print_error "Some Backups Failed!"
    echo "=========================================="
    exit 1
fi