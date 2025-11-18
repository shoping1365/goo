#!/bin/bash
# Configuration for Multiple Production Sites
# هر پروژه تنظیمات خودش را از .env می‌خواند

# پوشه مشترک سورس کد
SHARED_DIR="/opt/shared"

# تعریف پروژه‌ها - فقط مسیر دایرکتوری
# نام پروژه از نام دایرکتوری گرفته می‌شود
# پورت‌ها از فایل .env هر پروژه خوانده می‌شوند

PROJECTS=(
    "/data/iranxia"
    "/data/rayancomp"
)

# دایرکتوری بک‌آپ
BACKUP_BASE_DIR="/opt/deployment/backup"

# تعداد بک‌آپ‌هایی که نگه داشته میشه
KEEP_BACKUPS=10

# رنگ‌ها برای خروجی
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# تابع برای چاپ پیام‌های رنگی
print_success() {
    echo -e "${GREEN}✅ $1${NC}"
}

print_error() {
    echo -e "${RED}❌ $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

print_info() {
    echo -e "${BLUE}ℹ️  $1${NC}"
}

# تابع برای گرفتن نام پروژه از مسیر
get_project_name() {
    local project_dir=$1
    # گرفتن نام آخرین دایرکتوری از مسیر
    basename "$project_dir"
}

# تابع برای parse کردن اطلاعات پروژه
parse_project_config() {
    local project_dir=$1
    export PROJECT_DIR="$project_dir"
    export PROJECT_NAME=$(get_project_name "$project_dir")
}

# تابع برای خواندن پورت‌ها از .env
read_env_file() {
    local env_file=$1
    
    if [ ! -f "$env_file" ]; then
        return 1
    fi
    
    # خواندن PORT (frontend) و BACKEND_PORT
    export FRONTEND_PORT=$(grep -E '^PORT=' "$env_file" | cut -d'=' -f2 | tr -d ' "' || echo "3000")
    export BACKEND_PORT=$(grep -E '^BACKEND_PORT=' "$env_file" | cut -d'=' -f2 | tr -d ' "' || echo "9090")
    
    # نام‌های PM2 (می‌توان از env خواند یا از نام پروژه ساخت)
    export PM2_FRONTEND="${PROJECT_NAME}-frontend"
    export PM2_BACKEND="${PROJECT_NAME}-backend"
}