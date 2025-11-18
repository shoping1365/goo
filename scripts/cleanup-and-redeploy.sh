#!/bin/bash
# پاکسازی PM2 و راه‌اندازی صحیح با systemd

echo "=========================================="
echo "🧹 پاکسازی PM2"
echo "=========================================="

# حذف همه process های بک‌اند از PM2
echo "حذف بک‌اند‌ها از PM2..."
pm2 delete iranxia-backend 2>/dev/null || true
pm2 delete rayancomp-backend 2>/dev/null || true

# حذف فرانت‌اندهای errored
pm2 delete iranxia-frontend 2>/dev/null || true
pm2 delete rayancomp-frontend 2>/dev/null || true

# حذف همه
pm2 delete all 2>/dev/null || true

pm2 save

echo "✅ PM2 پاکسازی شد"
echo ""

echo "=========================================="
echo "🔧 راه‌اندازی Systemd Services"
echo "=========================================="

cd /opt/shared
bash scripts/setup-systemd-services.sh

echo ""
echo "=========================================="
echo "🚀 Deploy مجدد"
echo "=========================================="

bash scripts/deploy-all-sites.sh

echo ""
echo "=========================================="
echo "✅ تمام!"
echo "=========================================="
echo ""
echo "بررسی وضعیت:"
echo "  systemctl status iranxia-backend"
echo "  systemctl status rayancomp-backend"
echo "  pm2 list"
