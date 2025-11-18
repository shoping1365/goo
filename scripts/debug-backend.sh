#!/bin/bash
# Debug Backend Issues
# این اسکریپت مشکلات بک‌اند رو بررسی و نمایش میده

echo "=========================================="
echo "🔍 Debugging Backend Services"
echo "=========================================="
echo ""

# بررسی iranxia-backend
echo "1️⃣ Checking iranxia-backend..."
echo "---"
echo "Service status:"
systemctl status iranxia-backend --no-pager -l || true
echo ""
echo "Last 30 lines of logs:"
journalctl -u iranxia-backend -n 30 --no-pager || true
echo ""
echo "Binary exists?"
ls -lh /data/iranxia/my-go-backend/api 2>&1 || echo "Binary NOT FOUND!"
echo ""
echo ".env exists?"
ls -lh /data/iranxia/my-go-backend/.env 2>&1 || echo ".env NOT FOUND!"
echo ""
echo "=========================================="
echo ""

# بررسی rayancomp-backend
echo "2️⃣ Checking rayancomp-backend..."
echo "---"
echo "Service status:"
systemctl status rayancomp-backend --no-pager -l || true
echo ""
echo "Last 30 lines of logs:"
journalctl -u rayancomp-backend -n 30 --no-pager || true
echo ""
echo "Binary exists?"
ls -lh /data/rayancomp/my-go-backend/api 2>&1 || echo "Binary NOT FOUND!"
echo ""
echo ".env exists?"
ls -lh /data/rayancomp/my-go-backend/.env 2>&1 || echo ".env NOT FOUND!"
echo ""
echo "=========================================="
echo ""

# بررسی PostgreSQL
echo "3️⃣ Checking PostgreSQL..."
systemctl status postgresql postgresql-16 --no-pager || systemctl status postgresql --no-pager || true
echo ""

# بررسی پورت‌ها
echo "4️⃣ Checking ports..."
echo "Port 9090 (iranxia):"
netstat -tulpn | grep 9090 || echo "Port 9090 is FREE"
echo ""
echo "Port 9091 (rayancomp):"
netstat -tulpn | grep 9091 || echo "Port 9091 is FREE"
echo ""

echo "=========================================="
echo "✅ Debug complete!"
echo "=========================================="
