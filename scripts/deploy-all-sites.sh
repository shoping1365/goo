#!/bin/bash
# Deploy to All Production Projects
# تاریخ: 8 اکتبر 2025
# این اسکریپت سورس کد را از /opt/shared به تمام پروژه‌ها کپی و بیلد می‌کند
# هر پروژه تنظیمات خودش را از .env می‌خواند
# Backend با systemd و Frontend با PM2 مدیریت می‌شود
# noop: trigger deploy (comment-only change)

set -e
set -o pipefail

# Load configuration
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${SCRIPT_DIR}/sites-config.sh"

# Helper: safe value extraction from .env (KEY=VALUE), returns empty if not found
get_env_value() {
    # $1: key, $2: env file path
    local key="$1"; local env_file="$2"
    # match beginning of line KEY=... and extract after '=' while trimming quotes/spaces
    grep -E "^${key}=" "$env_file" 2>/dev/null | head -n1 | cut -d'=' -f2- | sed 's/^\s\+//;s/\s\+$//' | sed 's/^\"//;s/\"$//' | sed "s/^'//;s/'$//"
}

# Helper: تشخیص نسخه PostgreSQL نصب شده
detect_postgresql_service() {
    # چک کردن نسخه‌های مختلف PostgreSQL
    if systemctl list-units --type=service | grep -q "postgresql-16.service"; then
        echo "postgresql-16.service"
    elif systemctl list-units --type=service | grep -q "postgresql-15.service"; then
        echo "postgresql-15.service"
    elif systemctl list-units --type=service | grep -q "postgresql-14.service"; then
        echo "postgresql-14.service"
    elif systemctl list-units --type=service | grep -q "postgresql.service"; then
        echo "postgresql.service"
    else
        echo "postgresql.service"  # fallback
    fi
}

# Preflight: ensure required database envs exist in .env
validate_db_envs() {
    # $1: env file path
    local env_file="$1"
    local missing=()
    local DB_HOST DB_PORT DB_USER DB_NAME
    DB_HOST=$(get_env_value DB_HOST "$env_file")
    DB_PORT=$(get_env_value DB_PORT "$env_file")
    DB_USER=$(get_env_value DB_USER "$env_file")
    DB_NAME=$(get_env_value DB_NAME "$env_file")
    # Accept DB_PASSWORD empty, DB_SSLMODE optional
    if [ -z "$DB_HOST" ]; then missing+=(DB_HOST); fi
    if [ -z "$DB_PORT" ]; then missing+=(DB_PORT); fi
    if [ -z "$DB_USER" ]; then missing+=(DB_USER); fi
    if [ -z "$DB_NAME" ]; then missing+=(DB_NAME); fi
    if [ ${#missing[@]} -gt 0 ]; then
        print_error "${PROJECT_NAME}: Missing required DB envs in ${env_file}: ${missing[*]}"
        echo "  📝 Tip: Define these keys in ${env_file} (they are project-specific and not overwritten by deploy)."
        return 1
    fi
    return 0
}

echo "=========================================="
echo "🚀 Starting Deployment to All Projects"
echo "=========================================="
echo "📂 Shared Source: ${SHARED_DIR}"
echo "🌐 Number of Projects: ${#PROJECTS[@]}"
echo "🕒 Run timestamp: $(date -u '+%Y-%m-%dT%H:%M:%SZ')"
if [ -d "${SHARED_DIR}/.git" ]; then
    GIT_HASH=$(git -C "${SHARED_DIR}" rev-parse --short HEAD 2>/dev/null || echo "unknown")
    echo "🔐 Commit: ${GIT_HASH}"
fi
echo ""

# پاکسازی کامل PM2 از تمام پردازش‌های قدیمی
print_info "Cleaning up ALL old PM2 processes (complete cleanup)..."

# حذف اپ‌های قدیمی PM2 که با ecosystem.* استارت شده‌اند
pm2 delete ecosystem.production 2>/dev/null || true
pm2 delete ecosystem.config 2>/dev/null || true

# حذف تمام PM2 processes برای هر پروژه (backend و frontend)
for project_dir in "${PROJECTS[@]}"; do
    project_name=$(basename "$project_dir")
    
    print_info "Cleaning up PM2 processes for ${project_name}..."
    
    # حذف backend processes (که اشتباهی با PM2 استارت شده‌اند - حالا با systemd هستند)
    pm2 stop "${project_name}-backend" 2>/dev/null || true
    pm2 delete "${project_name}-backend" 2>/dev/null || true
    
    # حذف frontend processes
    pm2 stop "${project_name}-frontend" 2>/dev/null || true
    pm2 delete "${project_name}-frontend" 2>/dev/null || true
    
    # حذف با نام‌های مختلف احتمالی
    pm2 stop "${project_name}" 2>/dev/null || true
    pm2 delete "${project_name}" 2>/dev/null || true
    
    print_success "PM2 processes cleaned for ${project_name}"
done

# حذف پردازش‌های errored یا stopped
print_info "Cleaning up errored/stopped PM2 processes..."
pm2 list | grep -E "errored|stopped" | awk '{print $2}' | xargs -r pm2 delete 2>/dev/null || true

# نمایش وضعیت نهایی PM2
print_info "Final PM2 status:"
pm2 list || true

# ذخیره وضعیت PM2
pm2 save --force 2>/dev/null || true

print_success "PM2 cleanup completed"
echo ""

# پاکسازی اولیه Backend Services - قبل از دیپلوی
print_info "Pre-deployment backend cleanup..."
for project_dir in "${PROJECTS[@]}"; do
    project_name=$(basename "$project_dir")
    
    print_info "Cleaning up ${project_name}-backend service..."
    
    # اگر سرویس وجود دارد، stop کن
    if systemctl list-units --full -all | grep -q "${project_name}-backend.service"; then
        print_info "Stopping ${project_name}-backend service..."
        systemctl stop "${project_name}-backend" 2>/dev/null || true
        
        # صبر کوتاه
        sleep 2
        
        # اگر هنوز active است، force kill
        if systemctl is-active --quiet "${project_name}-backend" 2>/dev/null; then
            print_warning "${project_name}-backend still active, sending SIGKILL..."
            systemctl kill --signal=SIGKILL "${project_name}-backend" 2>/dev/null || true
            sleep 2
        fi
        
        # چک کردن مجدد و kill مستقیم اگر لازم باشد
        if systemctl is-active --quiet "${project_name}-backend" 2>/dev/null; then
            print_warning "${project_name}-backend still running, direct process kill..."
            BACKEND_PID=$(systemctl show -p MainPID --value "${project_name}-backend" 2>/dev/null | grep -v '^0$' || echo "")
            if [ -n "$BACKEND_PID" ] && kill -0 "$BACKEND_PID" 2>/dev/null; then
                kill -TERM "$BACKEND_PID" 2>/dev/null || true
                sleep 1
                if kill -0 "$BACKEND_PID" 2>/dev/null; then
                    kill -KILL "$BACKEND_PID" 2>/dev/null || true
                fi
            fi
        fi
        
        print_success "${project_name}-backend service stopped"
    else
        print_info "${project_name}-backend service not found (first deployment)"
    fi
done

# صبر بیشتر برای آزاد شدن پورت‌ها
print_info "Waiting for all backend ports to be freed..."
sleep 5

# چک کردن پورت‌های backend و آزاد کردن آنها
print_info "Checking and freeing any remaining backend ports..."
for project_dir in "${PROJECTS[@]}"; do
    project_name=$(basename "$project_dir")
    ENV_FILE="${project_dir}/my-go-backend/.env"
    
    if [ -f "$ENV_FILE" ]; then
        BACKEND_PORT=$(get_env_value BACKEND_PORT "$ENV_FILE")
        if [ -n "$BACKEND_PORT" ]; then
            if ss -tulpn 2>/dev/null | grep -q ":${BACKEND_PORT} "; then
                print_warning "Port ${BACKEND_PORT} still occupied for ${project_name}, freeing..."
                
                # آزاد کردن پورت
                if command -v fuser >/dev/null 2>&1; then
                    fuser -k -TERM ${BACKEND_PORT}/tcp 2>/dev/null || true
                    sleep 1
                    if ss -tulpn 2>/dev/null | grep -q ":${BACKEND_PORT} "; then
                        fuser -k -KILL ${BACKEND_PORT}/tcp 2>/dev/null || true
                    fi
                fi
                
                # صبر تا پورت آزاد شود
                for i in {1..5}; do
                    if ! ss -tulpn 2>/dev/null | grep -q ":${BACKEND_PORT} "; then
                        print_success "Port ${BACKEND_PORT} freed for ${project_name}"
                        break
                    fi
                    sleep 1
                done
            else
                print_success "Port ${BACKEND_PORT} already free for ${project_name}"
            fi
        fi
    fi
done

print_success "Backend services pre-cleanup completed"
echo ""

DEPLOY_SUCCESS=true

# Track per-project deployment results for better diagnostics
declare -A PROJECT_RESULTS

record_result() {
    # $1: project name, $2: status string
    local _p="$1"; local _s="$2"
    PROJECT_RESULTS["${_p}"]="${_s}"
}

# Track used ports to detect duplicates
declare -A USED_FRONTEND_PORTS
declare -A USED_BACKEND_PORTS

# Deploy به هر پروژه
for project_dir in "${PROJECTS[@]}"; do
    parse_project_config "$project_dir"
    
    echo "=========================================="
    print_info "Deploying to: ${PROJECT_NAME}"
    echo "=========================================="
    echo "📁 Target: ${PROJECT_DIR}"
    echo ""
    
    # ایجاد دایرکتوری پروژه
    mkdir -p "${PROJECT_DIR}"
    
    # کپی سورس کد از shared به پروژه
    print_info "Copying source code from shared..."
    
    # بررسی وجود فایل در مبدا
    if [ ! -f "${SHARED_DIR}/my-go-backend/cmd/api/main.go" ]; then
        print_error "Source main.go not found at ${SHARED_DIR}/my-go-backend/cmd/api/main.go"
        print_error "Cannot proceed without source files!"
        DEPLOY_SUCCESS=false
        continue
    fi
    
    print_success "Source main.go found at ${SHARED_DIR}/my-go-backend/cmd/api/main.go"
    
    # کپی فایل‌ها با rsync - مثل Git عمل میکنه و فایل‌های قدیمی رو حذف میکنه
    # ولی فایل‌ها و پوشه‌های مخصوص production محافظت میشن
    print_info "Syncing files with rsync (Git-like behavior with --delete)..."
    rsync -av --delete --progress \
        --exclude='node_modules/' \
        --exclude='.nuxt/' \
        --exclude='.output/' \
        --exclude='dist/' \
        --exclude='.git/' \
        --exclude='logs/' \
        --exclude='.env' \
        --exclude='.env.local' \
        --exclude='.env.production' \
        --exclude='uploads/' \
        --exclude='storage/' \
        --exclude='temp/' \
        --exclude='cache/' \
        --exclude='sessions/' \
        --exclude='*.pem' \
        --exclude='*.key' \
        --exclude='*.crt' \
        --exclude='ssl/' \
        --exclude='*.db' \
        --exclude='*.sqlite' \
        --exclude='*.log' \
        --exclude='backup/' \
        --exclude='backups/' \
        --exclude='my-go-backend/api' \
        --exclude='my-go-backend/api.exe' \
        --exclude='my-go-backend/main' \
        --exclude='my-go-backend/server' \
        "${SHARED_DIR}/" "${PROJECT_DIR}/" || {
        print_error "rsync failed!"
        record_result "${PROJECT_NAME}" "FAILED: rsync copy from ${SHARED_DIR}"
        DEPLOY_SUCCESS=false
        continue
    }
    print_success "Source code copied via rsync"
    
    # بررسی وجود فایل در مقصد
    if [ ! -f "${PROJECT_DIR}/my-go-backend/cmd/api/main.go" ]; then
        print_error "❌ CRITICAL: main.go was NOT copied to destination!"
        print_error "Source: ${SHARED_DIR}/my-go-backend/cmd/api/main.go"
        print_error "Destination: ${PROJECT_DIR}/my-go-backend/cmd/api/main.go"
        echo ""
        echo "Checking what was copied:"
        ls -la "${PROJECT_DIR}/my-go-backend/" 2>/dev/null || echo "my-go-backend directory is empty"
        ls -la "${PROJECT_DIR}/my-go-backend/cmd/" 2>/dev/null || echo "cmd directory not found"
        record_result "${PROJECT_NAME}" "FAILED: source verification (missing main.go)"
        DEPLOY_SUCCESS=false
        continue
    fi
    
    print_success "✅ Verified: main.go successfully copied to ${PROJECT_DIR}/my-go-backend/cmd/api/main.go"
    
    cd "${PROJECT_DIR}"
    
    # سیاست پروژه: فایل env برای بک‌اند و دیتابیس فقط در مسیر my-go-backend/.env نگهداری می‌شود
    ENV_FILE="${PROJECT_DIR}/my-go-backend/.env"
    if [ ! -f "$ENV_FILE" ]; then
        print_error ".env file not found at ${ENV_FILE} (backend env must live under my-go-backend/.env)"
        record_result "${PROJECT_NAME}" "FAILED: missing my-go-backend/.env"
        DEPLOY_SUCCESS=false
        continue
    fi

    # خواندن تنظیمات از .env انتخاب‌شده (my-go-backend/.env)
    read_env_file "$ENV_FILE"
    print_success "Configuration loaded from: $ENV_FILE"
    echo "  🔌 Frontend Port (PORT): ${FRONTEND_PORT}"
    echo "  🔌 Backend Port (BACKEND_PORT): ${BACKEND_PORT}"

    # Validate uniqueness of ports across projects (warn, not fail)
    if [ -n "${USED_FRONTEND_PORTS[${FRONTEND_PORT}]}" ]; then
        print_warning "Frontend PORT ${FRONTEND_PORT} already used by ${USED_FRONTEND_PORTS[${FRONTEND_PORT}]} (consider changing to avoid conflicts)"
    else
        USED_FRONTEND_PORTS[${FRONTEND_PORT}]="${PROJECT_NAME}"
    fi
    if [ -n "${USED_BACKEND_PORTS[${BACKEND_PORT}]}" ]; then
        print_warning "Backend BACKEND_PORT ${BACKEND_PORT} already used by ${USED_BACKEND_PORTS[${BACKEND_PORT}]} (consider changing to avoid conflicts)"
    else
        USED_BACKEND_PORTS[${BACKEND_PORT}]="${PROJECT_NAME}"
    fi

    # پیش از ادامه، اعتبارسنجی مقادیر دیتابیس در .env
    if ! validate_db_envs "$ENV_FILE"; then
        record_result "${PROJECT_NAME}" "FAILED: invalid/missing DB_* keys in .env"
        DEPLOY_SUCCESS=false
        continue
    fi
    
    # پاک کردن فایل‌های قدیمی قبل از نصب - مثل Git
    print_info "Cleaning old build files and dependencies (Git-like cleanup)..."
    
    # حذف کامل node_modules و lock files
    rm -rf node_modules package-lock.json yarn.lock pnpm-lock.yaml
    
    # حذف کامل build artifacts فرانت‌اند
    rm -rf .nuxt .output dist .cache
    
    # حذف فایل‌های موقت و cache
    rm -rf .turbo .vite .rollup.cache
    
    # حذف binary قدیمی بک‌اند
    rm -f my-go-backend/api my-go-backend/api.exe my-go-backend/main my-go-backend/server
    
    # حذف Go build cache (اختیاری - فقط در صورت مشکل)
    # cd my-go-backend && go clean -cache -modcache -testcache 2>/dev/null || true && cd ..
    
    print_success "Cleaned all old build files (Git-like cleanup complete)"
    
    # نصب dependencies فرانت‌اند
    print_info "Installing frontend dependencies..."
    # استفاده از تمام هسته‌های CPU برای npm install
    # Skip postinstall scripts during npm install to avoid nuxt prepare errors
    npm install --no-audit --legacy-peer-deps --ignore-scripts --also=dev || {
        print_error "Failed to install dependencies for ${PROJECT_NAME}"
        record_result "${PROJECT_NAME}" "FAILED: npm install"
        DEPLOY_SUCCESS=false
        continue
    }
    # Now run postinstall scripts explicitly after dependencies are installed
    npm run postinstall || {
        print_warning "postinstall script failed (non-fatal), continuing..."
    }
    print_success "Dependencies installed"
    
    # بیلد فرانت‌اند
    print_info "Building frontend (Nuxt.js)..."
    # تنظیم متغیرهای محیطی برای استفاده بهینه از CPU
    export UV_THREADPOOL_SIZE=128
    export NODE_OPTIONS="--max-old-space-size=16384 --max-semi-space-size=1024"
    
    npm run build || {
        print_error "Failed to build frontend for ${PROJECT_NAME}"
        record_result "${PROJECT_NAME}" "FAILED: npm run build"
        DEPLOY_SUCCESS=false
        continue
    }
    print_success "Frontend built successfully"
    
    # بیلد بک‌اند Go
    print_info "Building backend (Go)..."
    cd "${PROJECT_DIR}/my-go-backend"
    
    # همگام‌سازی vendor با go.mod
    print_info "Syncing vendor dependencies..."
    go mod vendor || {
        print_error "Failed to sync vendor for ${PROJECT_NAME}"
        record_result "${PROJECT_NAME}" "FAILED: go mod vendor"
        DEPLOY_SUCCESS=false
        continue
    }
    
    # Build Go binary - باید کل package رو build کنه نه فقط main.go
    # Pass -buildvcs=false to avoid 'error obtaining VCS status' when .git is missing or inaccessible
    go build -buildvcs=false -o api ./cmd/api || {
        print_error "Failed to build backend for ${PROJECT_NAME}"
        echo "Build output:"
        go build -v -buildvcs=false -o api ./cmd/api 2>&1 | tail -30
        record_result "${PROJECT_NAME}" "FAILED: go build backend"
        DEPLOY_SUCCESS=false
        continue
    }
    
    # بررسی وجود فایل binary
    if [ ! -f "api" ]; then
        print_error "Binary file 'api' was not created after build!"
        record_result "${PROJECT_NAME}" "FAILED: missing backend binary after build"
        DEPLOY_SUCCESS=false
        continue
    fi
    
    # تنظیم مجوزهای فایل باینری (فقط مالک و گروه)
    chmod 750 api
    chown root:root api 2>/dev/null || true
    
    # اصلاح context امنیتی SELinux
    if command -v restorecon &> /dev/null; then
        restorecon -v api 2>/dev/null || chcon -t bin_t api 2>/dev/null || true
    fi
    
    # اگر SELinux فعال است، context را به systemd تنظیم کن
    if command -v getenforce &> /dev/null && [ "$(getenforce)" = "Enforcing" ]; then
        echo "  🔒 SELinux is enforcing, setting proper context..."
        chcon -t systemd_unit_file_exec_t api 2>/dev/null || \
        chcon -t bin_t api 2>/dev/null || \
        chcon -t unconfined_exec_t api 2>/dev/null || true
    fi
    
    # بررسی مجوزهای فایل
    echo "  📋 File permissions:"
    ls -la api
    echo "  📋 File type:"
    file api
    
    # بررسی نهایی
    ls -lh api
    print_success "Backend built successfully: $(pwd)/api"
    
    # ایجاد دایرکتوری‌های لازم
    mkdir -p "${PROJECT_DIR}/logs"
    mkdir -p "${PROJECT_DIR}/my-go-backend/public/uploads"
    chmod 755 "${PROJECT_DIR}/logs"
    chmod 755 "${PROJECT_DIR}/my-go-backend"
    chmod 755 "${PROJECT_DIR}/my-go-backend/cmd"
    
    # بررسی و ایجاد/به‌روزرسانی systemd service (اگر قدیمی یا دارای Environment غلط باشد، بازنویسی کن)
    SERVICE_PATH="/etc/systemd/system/${PROJECT_NAME}-backend.service"
    RECREATE_SERVICE=false
    if [ ! -f "$SERVICE_PATH" ]; then
        print_warning "Systemd service not found. Creating..."
        RECREATE_SERVICE=true
    else
        if grep -q 'Environment="PORT=' "$SERVICE_PATH"; then
            print_warning "Legacy service file still uses PORT= instead of BACKEND_PORT=. Recreating..."
            RECREATE_SERVICE=true
        fi
    fi

    if [ "$RECREATE_SERVICE" = true ]; then
        # تشخیص نسخه PostgreSQL
        POSTGRESQL_SERVICE=$(detect_postgresql_service)
        print_info "Detected PostgreSQL service: ${POSTGRESQL_SERVICE}"
        
        cat > "$SERVICE_PATH" << EOF
[Unit]
Description=${PROJECT_NAME} Go Backend API
After=network.target ${POSTGRESQL_SERVICE}
Wants=${POSTGRESQL_SERVICE}

[Service]
Type=simple
User=root
Group=root
WorkingDirectory=${PROJECT_DIR}/my-go-backend
Environment="BACKEND_PORT=${BACKEND_PORT}"
Environment="GIN_MODE=release"
ExecStart=${PROJECT_DIR}/my-go-backend/api
Restart=always
RestartSec=10
StandardOutput=journal
StandardError=journal
LimitNOFILE=65536
# اگر سرویس fail شد، منتظر بمان تا پورت آزاد شود
StartLimitInterval=200
StartLimitBurst=5

[Install]
WantedBy=multi-user.target
EOF
        
        systemctl daemon-reload
        systemctl enable "${PROJECT_NAME}-backend" 2>/dev/null || true
        print_success "Systemd service (re)created"
    else
        print_info "Systemd service already up-to-date"
    fi
    
    # تنظیم مجوزهای نهایی قبل از راه‌اندازی سرویس (فقط مالک و گروه)
    print_info "Setting final permissions..."
    chmod 750 "${PROJECT_DIR}/my-go-backend/api"
    chown root:root "${PROJECT_DIR}/my-go-backend/api" 2>/dev/null || true
    
    # اصلاح SELinux context نهایی
    if command -v getenforce &> /dev/null && [ "$(getenforce)" = "Enforcing" ]; then
        echo "  🔒 Setting SELinux context for systemd..."
        chcon -t systemd_unit_file_exec_t "${PROJECT_DIR}/my-go-backend/api" 2>/dev/null || \
        chcon -t bin_t "${PROJECT_DIR}/my-go-backend/api" 2>/dev/null || \
        chcon -t unconfined_exec_t "${PROJECT_DIR}/my-go-backend/api" 2>/dev/null || true
    fi
    
    # بررسی مجوزهای نهایی
    echo "  📋 Final permissions:"
    ls -la "${PROJECT_DIR}/my-go-backend/api"
    
    # تست اجرای فایل
    if [ -x "${PROJECT_DIR}/my-go-backend/api" ]; then
        print_success "Binary file is executable"
    else
        print_error "Binary file is NOT executable!"
        chmod +x "${PROJECT_DIR}/my-go-backend/api"
        ls -la "${PROJECT_DIR}/my-go-backend/api"
    fi
    
    # Restart بک‌اند - با مدیریت کامل پورت
    print_info "Restarting backend service..."
    
    # ابتدا سرویس قدیمی را AGGRESSIVE متوقف کن
    print_info "Stopping old backend service (aggressive mode)..."
    
    # گرفتن PID اگر سرویس در حال اجراست
    BACKEND_PID=$(systemctl show -p MainPID --value "${PROJECT_NAME}-backend" 2>/dev/null | grep -v '^0$' || echo "")
    
    # اول stop معمولی
    systemctl stop "${PROJECT_NAME}-backend" 2>/dev/null || true
    sleep 2
    
    # اگر سرویس هنوز active است، kill aggressive
    if systemctl is-active --quiet "${PROJECT_NAME}-backend" 2>/dev/null; then
        print_warning "Service still active after stop, sending KILL signal..."
        systemctl kill --signal=SIGKILL "${PROJECT_NAME}-backend" 2>/dev/null || true
        sleep 2
    fi
    
    # اگر PID داشتیم و هنوز زنده است، مستقیم kill کن
    if [ -n "$BACKEND_PID" ] && kill -0 "$BACKEND_PID" 2>/dev/null; then
        print_warning "Process $BACKEND_PID still alive, killing directly..."
        kill -TERM "$BACKEND_PID" 2>/dev/null || true
        sleep 1
        if kill -0 "$BACKEND_PID" 2>/dev/null; then
            kill -KILL "$BACKEND_PID" 2>/dev/null || true
            sleep 1
        fi
    fi
    
    # بررسی و آزادسازی پورت بک‌اند
    print_info "Checking backend port ${BACKEND_PORT}..."
    if command -v ss >/dev/null 2>&1; then
        if ss -tulpn | grep -q ":${BACKEND_PORT} "; then
            print_warning "Port ${BACKEND_PORT} still in use after service stop"
            ss -tulpn | grep ":${BACKEND_PORT} " || true
            
            # تشخیص socket activation و dependency issues
            print_info "Checking for socket units and dependencies..."
            
            # چک کردن socket units
            if systemctl list-sockets | grep -q "${PROJECT_NAME}-backend.socket"; then
                print_warning "Found socket unit, disabling..."
                systemctl disable --now "${PROJECT_NAME}-backend.socket" 2>/dev/null || true
                sleep 2
            fi
            
            # چک کردن تمام socket units که ممکن است پورت را نگه دارند
            SOCKET_UNITS=$(systemctl list-sockets --no-pager | grep ":${BACKEND_PORT}" | awk '{print $1}' 2>/dev/null || echo "")
            if [ -n "$SOCKET_UNITS" ]; then
                print_warning "Found socket units holding port ${BACKEND_PORT}: $SOCKET_UNITS"
                for socket in $SOCKET_UNITS; do
                    print_info "Disabling socket: $socket"
                    systemctl disable --now "$socket" 2>/dev/null || true
                done
                sleep 3
            fi
            
            # چک کردن service dependencies
            print_info "Checking service dependencies..."
            DEPENDENT_SERVICES=$(systemctl list-dependencies --reverse "${PROJECT_NAME}-backend" 2>/dev/null | grep -v "${PROJECT_NAME}-backend" | grep -E "\.service$" | awk '{print $1}' || echo "")
            if [ -n "$DEPENDENT_SERVICES" ]; then
                print_warning "Found dependent services: $DEPENDENT_SERVICES"
                for service in $DEPENDENT_SERVICES; do
                    if systemctl is-active --quiet "$service" 2>/dev/null; then
                        print_info "Stopping dependent service: $service"
                        systemctl stop "$service" 2>/dev/null || true
                    fi
                done
                sleep 2
            fi
            
            # kill کردن پردازش‌های روی پورت - چند روش مختلف
            print_info "Killing processes on port ${BACKEND_PORT}..."
            
            # روش 1: fuser
            if command -v fuser >/dev/null 2>&1; then
                fuser -k -TERM ${BACKEND_PORT}/tcp 2>/dev/null || true
                sleep 2
                if ss -tulpn | grep -q ":${BACKEND_PORT} "; then
                    print_warning "Port still busy, sending SIGKILL via fuser..."
                    fuser -k -KILL ${BACKEND_PORT}/tcp 2>/dev/null || true
                    sleep 2
                fi
            fi
            
            # روش 2: lsof (backup method)
            if ss -tulpn | grep -q ":${BACKEND_PORT} " && command -v lsof >/dev/null 2>&1; then
                print_warning "Trying lsof method..."
                PIDS=$(lsof -ti:${BACKEND_PORT} 2>/dev/null || echo "")
                if [ -n "$PIDS" ]; then
                    echo "  Found PIDs: $PIDS"
                    for pid in $PIDS; do
                        kill -TERM "$pid" 2>/dev/null || true
                    done
                    sleep 2
                    # SIGKILL اگر هنوز زنده‌اند
                    for pid in $PIDS; do
                        if kill -0 "$pid" 2>/dev/null; then
                            kill -KILL "$pid" 2>/dev/null || true
                        fi
                    done
                    sleep 1
                fi
            fi
            
            # روش 3: ss + awk (آخرین راه)
            if ss -tulpn | grep -q ":${BACKEND_PORT} "; then
                print_warning "Trying ss+awk method..."
                PIDS=$(ss -tulpn | grep ":${BACKEND_PORT} " | awk '{print $7}' | grep -oP 'pid=\K[0-9]+' 2>/dev/null || echo "")
                if [ -n "$PIDS" ]; then
                    echo "  Found PIDs via ss: $PIDS"
                    for pid in $PIDS; do
                        kill -KILL "$pid" 2>/dev/null || true
                    done
                    sleep 2
                fi
            fi
            
            # روش 4: systemd reload (اگر PID 1 است)
            if ss -tulpn | grep -q ":${BACKEND_PORT} "; then
                PORT_HOLDER=$(ss -tulpn | grep ":${BACKEND_PORT} " | grep -o 'pid=[0-9]*' | cut -d= -f2 | head -1)
                if [ "$PORT_HOLDER" = "1" ]; then
                    print_warning "Port ${BACKEND_PORT} held by systemd (PID 1) - reloading systemd..."
                    systemctl daemon-reload
                    sleep 3
                    
                    # اگر هنوز مشغول است، سعی کن service را disable کن
                    if ss -tulpn | grep -q ":${BACKEND_PORT} "; then
                        print_warning "Still busy after daemon-reload, trying to disable service..."
                        systemctl disable "${PROJECT_NAME}-backend" 2>/dev/null || true
                        sleep 2
                    fi
                fi
            fi
        fi
    fi
    
    # صبر تا پورت کاملاً آزاد شود
    print_info "Waiting for backend port ${BACKEND_PORT} to be completely free..."
    sleep 3  # صبر اولیه برای اینکه OS پورت را آزاد کند
    
    for i in {1..20}; do
        if ! ss -tulpn 2>/dev/null | grep -q ":${BACKEND_PORT} "; then
            print_success "Backend port ${BACKEND_PORT} is now free"
            break
        fi
        echo "  Attempt $i/20: Port still busy, waiting..."
        sleep 3
    done
    
    # چک نهایی پورت - اگر هنوز مشغول است، پورت را تغییر بده
    if ss -tulpn 2>/dev/null | grep -q ":${BACKEND_PORT} "; then
        print_error "Backend port ${BACKEND_PORT} STILL in use after all cleanup attempts!"
        ss -tulpn | grep ":${BACKEND_PORT} " || true
        
        # راه‌حل نهایی: پیدا کردن پورت آزاد
        print_warning "Attempting to find alternative port..."
        ORIGINAL_BACKEND_PORT="$BACKEND_PORT"
        
        for alt_port in $(seq $((BACKEND_PORT + 1)) $((BACKEND_PORT + 10))); do
            if ! ss -tulpn 2>/dev/null | grep -q ":${alt_port} "; then
                print_success "Found alternative port: ${alt_port}"
                BACKEND_PORT="$alt_port"
                
                # به‌روزرسانی .env فایل
                sed -i "s/BACKEND_PORT=${ORIGINAL_BACKEND_PORT}/BACKEND_PORT=${BACKEND_PORT}/" "${ENV_FILE}" 2>/dev/null || true
                print_info "Updated .env file with new port: ${BACKEND_PORT}"
                break
            fi
        done
        
        # اگر هیچ پورت آزادی پیدا نشد، fail کن
        if [ "$BACKEND_PORT" = "$ORIGINAL_BACKEND_PORT" ]; then
            print_error "No alternative ports available (${ORIGINAL_BACKEND_PORT}-$((ORIGINAL_BACKEND_PORT + 10)))"
            print_error "Cannot start backend on occupied port."
            record_result "${PROJECT_NAME}" "FAILED: backend port ${BACKEND_PORT} occupied, no alternatives"
            DEPLOY_SUCCESS=false
            continue
        fi
    fi
    
    # حالا systemd را reload و سرویس را استارت کن
    systemctl daemon-reload
    
    # چک کردن dependencies قبل از start
    print_info "Checking PostgreSQL dependency..."
    POSTGRESQL_SERVICE=$(detect_postgresql_service)
    if ! systemctl is-active --quiet "$POSTGRESQL_SERVICE"; then
        print_warning "PostgreSQL service ($POSTGRESQL_SERVICE) is not active. Starting..."
        systemctl start "$POSTGRESQL_SERVICE" || {
            print_error "Failed to start PostgreSQL service: $POSTGRESQL_SERVICE"
            print_error "Available PostgreSQL services:"
            systemctl list-units --type=service | grep postgresql || true
        }
        sleep 3
    fi
    
    # چک کردن WorkingDirectory
    print_info "Checking WorkingDirectory..."
    if [ ! -d "${PROJECT_DIR}/my-go-backend" ]; then
        print_error "WorkingDirectory does not exist: ${PROJECT_DIR}/my-go-backend"
        record_result "${PROJECT_NAME}" "FAILED: WorkingDirectory missing"
        DEPLOY_SUCCESS=false
        continue
    fi
    
    # چک کردن ExecStart path
    print_info "Checking ExecStart path..."
    if [ ! -f "${PROJECT_DIR}/my-go-backend/api" ]; then
        print_error "ExecStart binary does not exist: ${PROJECT_DIR}/my-go-backend/api"
        record_result "${PROJECT_NAME}" "FAILED: ExecStart binary missing"
        DEPLOY_SUCCESS=false
        continue
    fi
    
    # چک کردن permissions
    print_info "Checking binary permissions..."
    ls -la "${PROJECT_DIR}/my-go-backend/api"
    
    # چک کردن environment variables
    print_info "Checking environment variables..."
    echo "BACKEND_PORT: ${BACKEND_PORT}"
    echo "GIN_MODE: release"
    
    # چک کردن service file
    print_info "Checking service file..."
    if [ ! -f "$SERVICE_PATH" ]; then
        print_error "Service file not found: $SERVICE_PATH"
        record_result "${PROJECT_NAME}" "FAILED: service file missing"
        DEPLOY_SUCCESS=false
        continue
    fi
    
    # نمایش محتوای service file برای debug
    print_info "Service file content:"
    cat "$SERVICE_PATH" | head -10
    
    systemctl start "${PROJECT_NAME}-backend" || {
        print_error "Failed to start backend service"
        print_error "Checking service status:"
        systemctl status "${PROJECT_NAME}-backend" --no-pager || true
        print_error "Checking recent logs:"
        journalctl -u "${PROJECT_NAME}-backend" -n 50 --no-pager || true
        print_error "Checking PostgreSQL status:"
        systemctl status "$POSTGRESQL_SERVICE" --no-pager || true
        record_result "${PROJECT_NAME}" "FAILED: systemd backend start"
        DEPLOY_SUCCESS=false
        continue
    }
    
    # صبر بیشتر برای اینکه سرویس کاملاً استارت شود
    sleep 5
    
    # چک کردن وضعیت سرویس با جزئیات بیشتر
    print_info "Checking service status after start..."
    
    # چک کردن وضعیت سرویس با روش‌های مختلف
    SERVICE_STATUS="unknown"
    
    # روش 1: systemctl is-active (اگر موجود باشد)
    if command -v systemctl >/dev/null 2>&1 && systemctl is-active "${PROJECT_NAME}-backend" >/dev/null 2>&1; then
        SERVICE_STATUS=$(systemctl is-active "${PROJECT_NAME}-backend" 2>/dev/null || echo "unknown")
    fi
    
    # روش 2: systemctl status (fallback)
    if [ "$SERVICE_STATUS" = "unknown" ]; then
        if systemctl status "${PROJECT_NAME}-backend" --no-pager | grep -q "Active: active"; then
            SERVICE_STATUS="active"
        elif systemctl status "${PROJECT_NAME}-backend" --no-pager | grep -q "Active: inactive"; then
            SERVICE_STATUS="inactive"
        elif systemctl status "${PROJECT_NAME}-backend" --no-pager | grep -q "Active: failed"; then
            SERVICE_STATUS="failed"
        else
            SERVICE_STATUS="unknown"
        fi
    fi
    
    # روش 3: چک کردن process با pgrep
    if [ "$SERVICE_STATUS" = "unknown" ]; then
        if pgrep -f "${PROJECT_NAME}-backend" >/dev/null 2>&1; then
            SERVICE_STATUS="active"
        else
            SERVICE_STATUS="inactive"
        fi
    fi
    
    # روش 4: چک کردن process با ps (backup)
    if [ "$SERVICE_STATUS" = "unknown" ]; then
        if ps aux | grep -v grep | grep -q "${PROJECT_NAME}-backend"; then
            SERVICE_STATUS="active"
        else
            SERVICE_STATUS="inactive"
        fi
    fi
    
    print_info "Service status detected: $SERVICE_STATUS"
    
    if [ "$SERVICE_STATUS" = "active" ]; then
        print_success "Backend service is running"
    else
        print_error "Backend service status: $SERVICE_STATUS"
        print_error "Service failed to stay active!"
        
        # نمایش جزئیات بیشتر
        print_error "Detailed service status:"
        systemctl status "${PROJECT_NAME}-backend" --no-pager || true
        
        print_error "Recent logs:"
        journalctl -u "${PROJECT_NAME}-backend" -n 20 --no-pager || true
        
        # تشخیص علت مشکل و تلاش برای حل آن
        print_info "🔍 Diagnosing service failure cause..."
        
        # چک کردن آیا مشکل پورت است
        PORT_ERROR=$(journalctl -u "${PROJECT_NAME}-backend" -n 10 --no-pager | grep -i "address already in use\|bind.*already in use\|port.*in use" || echo "")
        if [ -n "$PORT_ERROR" ]; then
            print_warning "🚨 Port conflict detected! Attempting to free the occupied port..."
            echo "Port error details: $PORT_ERROR"
            
            # چک کردن پورت فعلی
            if ss -tulpn 2>/dev/null | grep -q ":${BACKEND_PORT} "; then
                print_warning "Port ${BACKEND_PORT} is occupied. Freeing it..."
                
                # نمایش چه پروسه‌ای پورت را اشغال کرده
                print_info "Processes occupying port ${BACKEND_PORT}:"
                ss -tulpn | grep ":${BACKEND_PORT} " || true
                
                # آزاد کردن پورت با روش‌های مختلف
                print_info "Attempting to free port ${BACKEND_PORT}..."
                
                # روش 1: fuser
                if command -v fuser >/dev/null 2>&1; then
                    print_info "Using fuser to kill processes on port ${BACKEND_PORT}..."
                    fuser -k -TERM ${BACKEND_PORT}/tcp 2>/dev/null || true
                    sleep 2
                    if ss -tulpn 2>/dev/null | grep -q ":${BACKEND_PORT} "; then
                        print_warning "Port still busy, sending SIGKILL via fuser..."
                        fuser -k -KILL ${BACKEND_PORT}/tcp 2>/dev/null || true
                        sleep 2
                    fi
                fi
                
                # روش 2: lsof
                if ss -tulpn 2>/dev/null | grep -q ":${BACKEND_PORT} " && command -v lsof >/dev/null 2>&1; then
                    print_warning "Using lsof method..."
                    PIDS=$(lsof -ti:${BACKEND_PORT} 2>/dev/null || echo "")
                    if [ -n "$PIDS" ]; then
                        echo "  Found PIDs: $PIDS"
                        for pid in $PIDS; do
                            kill -TERM "$pid" 2>/dev/null || true
                        done
                        sleep 2
                        # SIGKILL اگر هنوز زنده‌اند
                        for pid in $PIDS; do
                            if kill -0 "$pid" 2>/dev/null; then
                                kill -KILL "$pid" 2>/dev/null || true
                            fi
                        done
                        sleep 1
                    fi
                fi
                
                # روش 3: ss + awk
                if ss -tulpn 2>/dev/null | grep -q ":${BACKEND_PORT} "; then
                    print_warning "Using ss+awk method..."
                    PIDS=$(ss -tulpn | grep ":${BACKEND_PORT} " | awk '{print $7}' | grep -oP 'pid=\K[0-9]+' 2>/dev/null || echo "")
                    if [ -n "$PIDS" ]; then
                        echo "  Found PIDs via ss: $PIDS"
                        for pid in $PIDS; do
                            kill -KILL "$pid" 2>/dev/null || true
                        done
                        sleep 2
                    fi
                fi
                
                # صبر تا پورت کاملاً آزاد شود
                print_info "Waiting for port ${BACKEND_PORT} to be completely free..."
                for i in {1..10}; do
                    if ! ss -tulpn 2>/dev/null | grep -q ":${BACKEND_PORT} "; then
                        print_success "Port ${BACKEND_PORT} is now free!"
                        break
                    fi
                    echo "  Attempt $i/10: Port still busy, waiting..."
                    sleep 2
                done
                
                # چک نهایی پورت
                if ss -tulpn 2>/dev/null | grep -q ":${BACKEND_PORT} "; then
                    print_error "Port ${BACKEND_PORT} STILL occupied after cleanup attempts!"
                    ss -tulpn | grep ":${BACKEND_PORT} " || true
                    record_result "${PROJECT_NAME}" "FAILED: unable to free backend port ${BACKEND_PORT}"
                    DEPLOY_SUCCESS=false
                    continue
                fi
                
                # حالا که پورت آزاد شد، سرویس را مجدداً استارت کن
                print_info "Port ${BACKEND_PORT} is free. Restarting service..."
                systemctl start "${PROJECT_NAME}-backend"
                sleep 3
                
                # چک کردن مجدد وضعیت
                if systemctl is-active --quiet "${PROJECT_NAME}-backend"; then
                    print_success "✅ Service started successfully after freeing port ${BACKEND_PORT}"
                else
                    print_error "Service still failed after freeing port"
                    record_result "${PROJECT_NAME}" "FAILED: systemd backend still inactive after port cleanup"
                    DEPLOY_SUCCESS=false
                    continue
                fi
            else
                print_warning "Port ${BACKEND_PORT} appears free now. Retrying service start..."
                systemctl start "${PROJECT_NAME}-backend"
                sleep 3
                
                if systemctl is-active --quiet "${PROJECT_NAME}-backend"; then
                    print_success "✅ Service started successfully on retry"
                else
                    record_result "${PROJECT_NAME}" "FAILED: systemd backend retry failed"
                    DEPLOY_SUCCESS=false
                    continue
                fi
            fi
        else
            # چک کردن dependency issues
            print_error "Checking dependencies:"
            systemctl list-dependencies "${PROJECT_NAME}-backend" --no-pager || true
            
            record_result "${PROJECT_NAME}" "FAILED: systemd backend inactive (status: $SERVICE_STATUS)"
            DEPLOY_SUCCESS=false
            continue
        fi
    fi
    
    # اطمینان از وجود خروجی بیلد فرانت‌اند
    if [ ! -f "${PROJECT_DIR}/.output/server/index.mjs" ]; then
        print_error "Frontend build output missing: ${PROJECT_DIR}/.output/server/index.mjs"
        record_result "${PROJECT_NAME}" "FAILED: frontend build output missing (.output/server/index.mjs)"
        DEPLOY_SUCCESS=false
        continue
    fi

    # Restart فرانت‌اند با PM2 
    print_info "Restarting frontend with PM2..."
    
    # حذف کامل instance‌های قدیمی این پروژه
    print_info "Stopping old PM2 instances..."
    pm2 stop "${PM2_FRONTEND}" 2>/dev/null || true
    sleep 1
    pm2 delete "${PM2_FRONTEND}" 2>/dev/null || true
    
    # اطمینان از آزاد بودن پورت frontend
    print_info "Checking frontend port ${FRONTEND_PORT}..."
    if command -v ss >/dev/null 2>&1; then
        if ss -tulpn | grep -q ":${FRONTEND_PORT} "; then
            print_warning "Port ${FRONTEND_PORT} still in use. Killing processes..."
            if command -v fuser >/dev/null 2>&1; then
                fuser -k -TERM ${FRONTEND_PORT}/tcp 2>/dev/null || true
                sleep 2
                # اگر باز هم مشغول بود، SIGKILL بزن
                if ss -tulpn | grep -q ":${FRONTEND_PORT} "; then
                    fuser -k -KILL ${FRONTEND_PORT}/tcp 2>/dev/null || true
                    sleep 1
                fi
            fi
        fi
    fi
    
    # صبر تا پورت کاملاً آزاد شود
    print_info "Waiting for port ${FRONTEND_PORT} to be completely free..."
    for i in {1..10}; do
        if ! ss -tulpn 2>/dev/null | grep -q ":${FRONTEND_PORT} "; then
            print_success "Port ${FRONTEND_PORT} is now free"
            break
        fi
        echo "  Attempt $i/10: Port still busy, waiting..."
        sleep 2
    done
    
    # چک نهایی پورت
    if ss -tulpn 2>/dev/null | grep -q ":${FRONTEND_PORT} "; then
        print_error "Port ${FRONTEND_PORT} STILL in use after cleanup attempts!"
        ss -tulpn | grep ":${FRONTEND_PORT} " || true
        print_error "Cannot start frontend on occupied port."
        record_result "${PROJECT_NAME}" "FAILED: frontend port ${FRONTEND_PORT} occupied"
        DEPLOY_SUCCESS=false
        continue
    fi
    
    cd "${PROJECT_DIR}"
    # ست کردن متغیرهای محیطی پورت برای جلوگیری از کانفلیکت بین پروژه‌ها
    export NODE_ENV=production
    export HOST=0.0.0.0
    export PORT="${FRONTEND_PORT}"
    export NITRO_PORT="${FRONTEND_PORT}"
    export NUXT_PORT="${FRONTEND_PORT}"

    pm2 start .output/server/index.mjs \
        --name "${PM2_FRONTEND}" \
        -i 8 \
        --env "NODE_ENV=production,HOST=0.0.0.0,PORT=${FRONTEND_PORT},NUXT_PORT=${FRONTEND_PORT},NITRO_PORT=${FRONTEND_PORT}" \
        --update-env || {
        print_error "Failed to start frontend"
        echo "Recent PM2 logs for ${PM2_FRONTEND}:"
        pm2 logs "${PM2_FRONTEND}" --lines 50 --nostream 2>/dev/null || true
        record_result "${PROJECT_NAME}" "FAILED: pm2 frontend start"
        DEPLOY_SUCCESS=false
        continue
    }
    
    # چک کردن اینکه واقعاً استارت شد
    sleep 3
    if ! pm2 list | grep -q "${PM2_FRONTEND}.*online"; then
        print_error "Frontend process not online after start!"
        pm2 logs "${PM2_FRONTEND}" --lines 30 --nostream 2>/dev/null || true
        record_result "${PROJECT_NAME}" "FAILED: pm2 frontend not online"
        DEPLOY_SUCCESS=false
        continue
    fi
    
    print_success "Frontend started successfully with 8 instances"
    pm2 save
    
    echo ""
    print_success "✅ ${PROJECT_NAME} deployed successfully!"
    echo "🌐 Frontend: http://localhost:${FRONTEND_PORT}"
    echo "🔌 Backend API: http://localhost:${BACKEND_PORT}"
    echo ""

    # Mark project success
    record_result "${PROJECT_NAME}" "OK"
done

# پاکسازی نهایی PM2
print_info "Final PM2 cleanup and optimization..."
pm2 delete ecosystem.production 2>/dev/null || true
pm2 delete ecosystem.config 2>/dev/null || true
pm2 list | grep -E "errored|stopped" | awk '{print $2}' | xargs -r pm2 delete 2>/dev/null || true
pm2 save --force
print_success "PM2 state saved"

echo "=========================================="
if [ "$DEPLOY_SUCCESS" = true ]; then
    print_success "ALL PROJECTS DEPLOYED SUCCESSFULLY!"
    echo ""
    echo "📊 PM2 Status:"
    pm2 list
    echo ""
    echo "📊 Systemd Status:"
    for project_dir in "${PROJECTS[@]}"; do
        project_name=$(basename "$project_dir")
        systemctl is-active --quiet "${project_name}-backend" && \
            echo "  ✅ ${project_name}-backend: RUNNING" || \
            echo "  ❌ ${project_name}-backend: STOPPED"
    done
    echo "=========================================="
    exit 0
else
    echo ""
    echo "📋 Deployment summary:"
    for project_dir in "${PROJECTS[@]}"; do
        project_name=$(basename "$project_dir")
        status_str="${PROJECT_RESULTS[${project_name}]}"
        if [ -z "$status_str" ]; then status_str="FAILED: unknown (no status recorded)"; fi
        echo "  - ${project_name}: ${status_str}"
    done
    echo ""
    print_error "SOME DEPLOYMENTS FAILED!"
    exit 1
fi