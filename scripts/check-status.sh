#!/bin/bash
# راهنمای سریع برای راه‌اندازی اولین بار

echo "=========================================="
echo "🔍 بررسی وضعیت فعلی"
echo "=========================================="

# بررسی وجود systemd services
echo ""
echo "1️⃣ بررسی Systemd Services:"
for service in iranxia-backend rayancomp-backend; do
    if [ -f "/etc/systemd/system/${service}.service" ]; then
        echo "  ✅ ${service}.service وجود دارد"
        systemctl is-active --quiet ${service} && echo "     └─ وضعیت: RUNNING ✅" || echo "     └─ وضعیت: STOPPED ⚠️"
    else
        echo "  ❌ ${service}.service وجود ندارد - باید بسازیم"
    fi
done

# بررسی PM2
echo ""
echo "2️⃣ بررسی PM2 Processes:"
if command -v pm2 &> /dev/null; then
    pm2 list
else
    echo "  ⚠️ PM2 نصب نیست"
fi

# بررسی Go
echo ""
echo "3️⃣ بررسی Go:"
if command -v go &> /dev/null; then
    go version
else
    echo "  ❌ Go نصب نیست"
fi

# بررسی Node
echo ""
echo "4️⃣ بررسی Node.js:"
if command -v node &> /dev/null; then
    node --version
    npm --version
else
    echo "  ❌ Node.js نصب نیست"
fi

# بررسی فایل‌های .env
echo ""
echo "5️⃣ بررسی فایل‌های .env:"
for dir in /data/iranxia /data/rayancomp; do
    if [ -f "${dir}/.env" ]; then
        echo "  ✅ ${dir}/.env وجود دارد"
    else
        echo "  ❌ ${dir}/.env وجود ندارد"
    fi
done

# بررسی باینری‌های Go
echo ""
echo "6️⃣ بررسی باینری‌های Go:"
for dir in /data/iranxia /data/rayancomp; do
    if [ -f "${dir}/my-go-backend/api" ]; then
        echo "  ✅ ${dir}/my-go-backend/api موجود است"
    else
        echo "  ❌ ${dir}/my-go-backend/api موجود نیست - باید build کنیم"
    fi
done

echo ""
echo "=========================================="
echo "✅ بررسی تمام شد"
echo "=========================================="
