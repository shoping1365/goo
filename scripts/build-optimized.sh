#!/bin/bash
# تنظیمات بهینه برای build سریع با تمام هسته‌های CPU

# تعداد هسته‌های CPU
CPU_CORES=$(nproc 2>/dev/null || sysctl -n hw.ncpu 2>/dev/null || echo 4)

echo "=========================================="
echo "⚡ CPU Optimization Settings"
echo "=========================================="
echo "CPU Cores: $CPU_CORES"
echo ""

# تنظیم متغیرهای محیطی
export UV_THREADPOOL_SIZE=$((CPU_CORES * 8))
export NODE_OPTIONS="--max-old-space-size=16384 --max-semi-space-size=1024"
export NUXT_TELEMETRY_DISABLED=1

# محاسبه تعداد بهینه worker ها
OPTIMAL_WORKERS=$((CPU_CORES - 1))
if [ $OPTIMAL_WORKERS -lt 1 ]; then
    OPTIMAL_WORKERS=1
fi

export VITE_WORKER_THREADS=$OPTIMAL_WORKERS

echo "UV_THREADPOOL_SIZE: $UV_THREADPOOL_SIZE"
echo "VITE_WORKER_THREADS: $VITE_WORKER_THREADS"
echo "NODE_OPTIONS: $NODE_OPTIONS"
echo ""
echo "=========================================="
echo "🚀 Starting Build..."
echo "=========================================="

# اجرای build با تنظیمات بهینه
time npm run build

echo ""
echo "=========================================="
echo "✅ Build Complete!"
echo "=========================================="
