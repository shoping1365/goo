#!/bin/bash

set -e

# رنگ‌ها
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

TESTS_PASSED=0
TESTS_FAILED=0
LOG_FILE="pre-deploy-tests_$(date +%Y%m%d_%H%M%S).log"
SECURITY_REPORT="security-report_$(date +%Y%m%d_%H%M%S).txt"

# خواندن متغیرهای محیطی
load_env() {
    if [ -f .env ]; then
        export $(cat .env | grep -v '^#' | xargs)
    fi
    
    # بررسی وجود متغیرهای ضروری
    if [ -z "$BACKEND_URL" ]; then
        echo -e "${RED}[ERROR]${NC} متغیر BACKEND_URL تنظیم نشده است!"
        echo "لطفاً در فایل .env مقدار BACKEND_URL را تنظیم کنید"
        exit 1
    fi
    
    if [ -z "$FRONTEND_URL" ]; then
        echo -e "${RED}[ERROR]${NC} متغیر FRONTEND_URL تنظیم نشده است!"
        echo "لطفاً در فایل .env مقدار FRONTEND_URL را تنظیم کنید"
        exit 1
    fi
}

log() {
    echo -e "${BLUE}[$(date +'%H:%M:%S')]${NC} $1" | tee -a "$LOG_FILE"
}

log_error() {
    echo -e "${RED}[✗]${NC} $1" | tee -a "$LOG_FILE"
    ((TESTS_FAILED++))
}

log_success() {
    echo -e "${GREEN}[✓]${NC} $1" | tee -a "$LOG_FILE"
    ((TESTS_PASSED++))
}

log_warning() {
    echo -e "${YELLOW}[!]${NC} $1" | tee -a "$LOG_FILE"
}

# بررسی موفقیت تست
check_test() {
    if [ $? -eq 0 ]; then
        log_success "$1"
        return 0
    else
        log_error "$1"
        return 1
    fi
}

echo "======================================"
echo "  Pre-Deploy Tests"
echo "======================================"
log "شروع تست‌های پیش از دیپلوی..."

# بارگذاری متغیرهای محیطی
load_env
log "Backend URL: $BACKEND_URL"
log "Frontend URL: $FRONTEND_URL"

# 1. بررسی کدهای Go
log "\n📦 تست بک‌اند Go..."
if [ -d "backend" ]; then
    cd backend

    log "  → لینتر Go..."
    if command -v golangci-lint &> /dev/null; then
        golangci-lint run --timeout 5m >> "../$LOG_FILE" 2>&1
        check_test "Go Linter"
    else
        go fmt ./... >> "../$LOG_FILE" 2>&1
        check_test "Go Format"
        go vet ./... >> "../$LOG_FILE" 2>&1
        check_test "Go Vet"
    fi

    log "  → تست‌های واحد Go..."
    go test -v -race -coverprofile=coverage.out ./... >> "../$LOG_FILE" 2>&1
    check_test "Go Unit Tests"

    log "  → بررسی پوشش تست..."
    if [ -f "coverage.out" ]; then
        COVERAGE=$(go tool cover -func=coverage.out | grep total | awk '{print $3}' | sed 's/%//')
        if (( $(echo "$COVERAGE >= 70" | bc -l) )); then
            log_success "Coverage: $COVERAGE% (حداقل: 70%)"
        else
            log_error "Coverage: $COVERAGE% (حداقل: 70%)"
        fi
    else
        log_warning "فایل coverage یافت نشد"
    fi

    log "  → بررسی امنیت Go (gosec)..."
    if command -v gosec &> /dev/null; then
        echo "=== Go Security Scan ===" >> "../$SECURITY_REPORT"
        gosec -fmt=text ./... >> "../$SECURITY_REPORT" 2>&1
        
        # شمارش آسیب‌پذیری‌ها
        GOSEC_ISSUES=$(grep -c "Severity:" "../$SECURITY_REPORT" || echo "0")
        
        if [ "$GOSEC_ISSUES" -eq 0 ]; then
            log_success "Go Security Scan - هیچ آسیب‌پذیری یافت نشد"
        else
            log_error "Go Security Scan - $GOSEC_ISSUES آسیب‌پذیری یافت شد!"
            echo ""
            echo -e "${YELLOW}نمونه آسیب‌پذیری‌ها:${NC}"
            grep -A 3 "Severity: HIGH\|Severity: MEDIUM" "../$SECURITY_REPORT" | head -20
        fi
    else
        log_warning "gosec نصب نیست"
    fi

    log "  → بررسی آسیب‌پذیری‌های کتابخانه‌ها..."
    if command -v govulncheck &> /dev/null; then
        echo "" >> "../$SECURITY_REPORT"
        echo "=== Vulnerability Check ===" >> "../$SECURITY_REPORT"
        govulncheck -json ./... > "../vulncheck.json" 2>&1
        
        # تحلیل نتایج
        if grep -q "\"Vulns\"" "../vulncheck.json"; then
            VULN_COUNT=$(grep -o "\"Vulns\"" "../vulncheck.json" | wc -l)
            log_error "Vulnerability Check - $VULN_COUNT آسیب‌پذیری CVE یافت شد!"
            
            # نمایش جزئیات
            govulncheck ./... >> "../$SECURITY_REPORT" 2>&1
            echo ""
            echo -e "${YELLOW}لیست CVE ها:${NC}"
            grep "GO-" "../vulncheck.json" | head -10
        else
            log_success "Vulnerability Check - هیچ CVE یافت نشد"
        fi
    else
        log_warning "govulncheck نصب نیست"
    fi

    cd ..
else
    log_warning "دایرکتوری backend یافت نشد"
fi

# 2. بررسی کدهای Nuxt
log "\n🎨 تست فرانت‌اند Nuxt..."
if [ -d "frontend" ]; then
    cd frontend

    log "  → نصب وابستگی‌ها..."
    npm ci --silent >> "../$LOG_FILE" 2>&1
    check_test "NPM Dependencies"

    log "  → لینتر ESLint..."
    if grep -q "\"lint\"" package.json; then
        npm run lint >> "../$LOG_FILE" 2>&1
        check_test "ESLint"
    else
        log_warning "اسکریپت lint تعریف نشده"
    fi

    log "  → تست‌های واحد..."
    if grep -q "\"test\"" package.json; then
        npm run test >> "../$LOG_FILE" 2>&1
        check_test "Frontend Unit Tests"
    else
        log_warning "تست فرانت‌اند تعریف نشده"
    fi

    log "  → بررسی آسیب‌پذیری‌های NPM..."
    echo "" >> "../$SECURITY_REPORT"
    echo "=== NPM Security Audit ===" >> "../$SECURITY_REPORT"
    
    npm audit --audit-level=high --json > "../npm-audit.json" 2>&1
    
    # تحلیل نتایج
    if [ -f "../npm-audit.json" ]; then
        HIGH_VULNS=$(grep -o "\"severity\":\"high\"" "../npm-audit.json" | wc -l)
        CRITICAL_VULNS=$(grep -o "\"severity\":\"critical\"" "../npm-audit.json" | wc -l)
        
        if [ "$HIGH_VULNS" -gt 0 ] || [ "$CRITICAL_VULNS" -gt 0 ]; then
            log_error "NPM Security - $CRITICAL_VULNS Critical, $HIGH_VULNS High آسیب‌پذیری!"
            
            # گزارش کامل
            npm audit --audit-level=high >> "../$SECURITY_REPORT" 2>&1
            
            echo ""
            echo -e "${YELLOW}آسیب‌پذیری‌های بحرانی:${NC}"
            npm audit --audit-level=critical 2>&1 | grep -A 5 "Package\|Severity\|Recommendation" | head -20
        else
            log_success "NPM Security Audit - هیچ آسیب‌پذیری بحرانی یافت نشد"
        fi
    fi

    cd ..
else
    log_warning "دایرکتوری frontend یافت نشد"
fi

# 3. تست سرویس‌های در حال اجرا
log "\n🔧 تست سرویس‌های در حال اجرا..."

log "  → Health Check Backend ($BACKEND_URL)..."
RETRY=0
MAX_RETRY=30
while [ $RETRY -lt $MAX_RETRY ]; do
    if curl -sf "$BACKEND_URL/health" > /dev/null 2>&1; then
        log_success "Backend Health Check"
        break
    fi
    RETRY=$((RETRY+1))
    sleep 2
done
if [ $RETRY -eq $MAX_RETRY ]; then
    log_error "Backend در دسترس نیست"
fi

log "  → Health Check Frontend ($FRONTEND_URL)..."
RETRY=0
while [ $RETRY -lt $MAX_RETRY ]; do
    if curl -sf "$FRONTEND_URL" > /dev/null 2>&1; then
        log_success "Frontend Health Check"
        break
    fi
    RETRY=$((RETRY+1))
    sleep 2
done
if [ $RETRY -eq $MAX_RETRY ]; then
    log_error "Frontend در دسترس نیست"
fi

log "  → تست API Endpoints..."
if curl -sf "$BACKEND_URL/api/todos" > /dev/null 2>&1; then
    log_success "API Endpoint Test"
else
    log_error "API Endpoint Test"
fi

# 4. تست پرفورمنس
log "\n⚡ تست پرفورمنس..."

if command -v ab &> /dev/null; then
    log "  → تست بار Backend با 10,000 کاربر همزمان..."
    log "    URL: $BACKEND_URL/api/todos"
    log "    درخواست‌ها: 100,000 | کاربران همزمان: 10,000"
    
    ab -n 100000 -c 10000 -t 30 "$BACKEND_URL/api/todos" >> "$LOG_FILE" 2>&1
    
    if [ $? -eq 0 ]; then
        REQUESTS_PER_SEC=$(grep "Requests per second" "$LOG_FILE" | tail -1 | awk '{print $4}')
        TIME_PER_REQUEST=$(grep "Time per request" "$LOG_FILE" | tail -1 | awk '{print $4}')
        FAILED_REQUESTS=$(grep "Failed requests" "$LOG_FILE" | tail -1 | awk '{print $3}')
        
        log_success "Performance Test Completed"
        log "    Requests/sec: $REQUESTS_PER_SEC"
        log "    Time/request: ${TIME_PER_REQUEST}ms"
        log "    Failed: $FAILED_REQUESTS"
        
        if [ "$FAILED_REQUESTS" -gt 100 ]; then
            log_error "تعداد درخواست‌های ناموفق زیاد است: $FAILED_REQUESTS"
        fi
    else
        log_error "Backend Load Test Failed"
    fi
    
    log "  → تست بار Frontend..."
    ab -n 100000 -c 10000 -t 30 "$FRONTEND_URL/" >> "$LOG_FILE" 2>&1
    check_test "Frontend Load Test"
else
    log_warning "Apache Bench نصب نیست"
fi

# نتیجه‌گیری
log "\n======================================"
log "  نتایج تست‌ها"
log "======================================"
log_success "موفق: $TESTS_PASSED"
if [ $TESTS_FAILED -gt 0 ]; then
    log_error "ناموفق: $TESTS_FAILED"
    echo ""
    log_error "❌ برخی تست‌ها ناموفق بودند!"
    echo ""
    echo -e "${YELLOW}📄 گزارش‌های امنیتی:${NC}"
    echo "  - لاگ کامل: $LOG_FILE"
    echo "  - گزارش امنیت: $SECURITY_REPORT"
    
    if [ -f "npm-audit.json" ]; then
        echo "  - گزارش NPM: npm-audit.json"
    fi
    
    if [ -f "vulncheck.json" ]; then
        echo "  - گزارش CVE: vulncheck.json"
    fi
    
    exit 1
else
    echo ""
    log_success "✅ همه تست‌ها موفق بود! آماده دیپلوی است."
    echo ""
    echo -e "${BLUE}📄 گزارش‌ها ذخیره شد:${NC}"
    echo "  - لاگ کامل: $LOG_FILE"
    echo "  - گزارش امنیت: $SECURITY_REPORT"
    exit 0
fi