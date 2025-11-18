#!/bin/bash
# Install and Setup PostgreSQL on AlmaLinux
# این اسکریپت PostgreSQL رو نصب و راه‌اندازی میکنه

set -e

echo "=========================================="
echo "📦 Installing PostgreSQL"
echo "=========================================="
echo ""

# تشخیص نسخه OS
if [ -f /etc/os-release ]; then
    . /etc/os-release
    OS=$ID
    VER=$VERSION_ID
fi

echo "Detected OS: $OS $VER"
echo ""

# نصب PostgreSQL
echo "1️⃣ Installing PostgreSQL..."
if [ "$OS" = "almalinux" ] || [ "$OS" = "rocky" ] || [ "$OS" = "rhel" ]; then
    # برای AlmaLinux/Rocky/RHEL
    dnf install -y postgresql-server postgresql-contrib
else
    echo "⚠️  Unknown OS. Trying generic installation..."
    dnf install -y postgresql-server postgresql-contrib || yum install -y postgresql-server postgresql-contrib
fi
echo "✅ PostgreSQL installed"
echo ""

# Initialize database
echo "2️⃣ Initializing PostgreSQL database..."
postgresql-setup --initdb || /usr/bin/postgresql-setup initdb
echo "✅ Database initialized"
echo ""

# Start and enable service
echo "3️⃣ Starting PostgreSQL service..."
systemctl enable postgresql
systemctl start postgresql
sleep 2
systemctl status postgresql --no-pager | head -10
echo "✅ PostgreSQL is running"
echo ""

# تنظیم trust برای local connections (برای راحتی setup)
echo "4️⃣ Configuring PostgreSQL authentication..."
PG_HBA="/var/lib/pgsql/data/pg_hba.conf"
if [ -f "$PG_HBA" ]; then
    # بک‌آپ
    cp "$PG_HBA" "${PG_HBA}.backup"
    
    # تغییر local به trust برای setup
    sed -i 's/local   all             all                                     peer/local   all             all                                     trust/g' "$PG_HBA"
    sed -i 's/host    all             all             127.0.0.1\/32            ident/host    all             all             127.0.0.1\/32            md5/g' "$PG_HBA"
    sed -i 's/host    all             all             ::1\/128                 ident/host    all             all             ::1\/128                 md5/g' "$PG_HBA"
    
    echo "✅ Authentication configured"
    
    # Restart برای اعمال تغییرات
    systemctl restart postgresql
    sleep 2
fi
echo ""

# تنظیم پسورد برای postgres user
echo "5️⃣ Setting postgres user password..."
sudo -u postgres psql -c "ALTER USER postgres WITH PASSWORD '1365';"
echo "✅ Password set for postgres user"
echo ""

# 🆕 فعال‌سازی Extensions در template1 (مثل Windows!)
echo "5.5️⃣ Setting up global extensions in template1..."
echo "   (مثل Windows - یکبار برای همیشه!)"
sudo -u postgres psql -d template1 << 'EOSQL'
CREATE EXTENSION IF NOT EXISTS pg_trgm;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS btree_gin;
CREATE EXTENSION IF NOT EXISTS btree_gist;
EOSQL
echo "✅ Global extensions enabled in template1"
echo "   💡 از این به بعد همه دیتابیس‌های جدید این extensionها را دارند!"
echo ""

# ساخت دیتابیس‌ها
echo "6️⃣ Creating databases and users..."

# iranxia
sudo -u postgres psql << 'EOSQL'
DROP DATABASE IF EXISTS iranxia;
DROP USER IF EXISTS iranxia_user;
CREATE USER iranxia_user WITH PASSWORD 'IranXia@2025#Secure!Pass';
CREATE DATABASE iranxia OWNER iranxia_user;
GRANT ALL PRIVILEGES ON DATABASE iranxia TO iranxia_user;
\c iranxia
GRANT ALL ON SCHEMA public TO iranxia_user;
EOSQL
echo "✅ iranxia database created"

# rayancomp
sudo -u postgres psql << 'EOSQL'
DROP DATABASE IF EXISTS rayancomp;
DROP USER IF EXISTS rayancomp_user;
CREATE USER rayancomp_user WITH PASSWORD 'RayanComp@2025#Secure!Pass';
CREATE DATABASE rayancomp OWNER rayancomp_user;
GRANT ALL PRIVILEGES ON DATABASE rayancomp TO rayancomp_user;
\c rayancomp
GRANT ALL ON SCHEMA public TO rayancomp_user;
EOSQL
echo "✅ rayancomp database created"
echo ""

# تست اتصال
echo "7️⃣ Testing database connections..."
sudo -u postgres psql -l | grep -E "iranxia|rayancomp"
echo "✅ Databases are ready"
echo ""

# بررسی extensions
echo "8️⃣ Verifying extensions..."
echo "   Extensions in template1:"
sudo -u postgres psql -d template1 -c "\dx" | head -10
echo ""
echo "   Extensions in iranxia:"
sudo -u postgres psql -d iranxia -c "\dx" | head -10
echo ""

echo "=========================================="
echo "✅ PostgreSQL Setup Complete!"
echo "=========================================="
echo ""
echo "📊 Database Information:"
echo "  - iranxia database: iranxia_user / IranXia@2025#Secure!Pass"
echo "  - rayancomp database: rayancomp_user / RayanComp@2025#Secure!Pass"
echo "  - postgres superuser: postgres / 1365"
echo ""
echo "🎉 Global Extensions:"
echo "  - pg_trgm (برای جستجوی سریع)"
echo "  - uuid-ossp (برای UUID)"
echo "  - btree_gin, btree_gist (برای indexهای پیشرفته)"
echo ""
echo "💡 از این به بعد هر دیتابیس جدید این extensionها را دارد!"
echo ""
echo "🔧 Service commands:"
echo "  systemctl status postgresql"
echo "  systemctl restart postgresql"
echo "  sudo -u postgres psql"
