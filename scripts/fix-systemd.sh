#!/bin/bash
# Fix systemd service - Remove problematic settings

echo "Fixing systemd services..."

# Stop services
systemctl stop iranxia-backend rayancomp-backend 2>/dev/null || true

# Create fixed iranxia-backend service
cat > /etc/systemd/system/iranxia-backend.service << 'EOF'
[Unit]
Description=Iranxia Go Backend API
After=network.target postgresql.service postgresql-16.service
Wants=postgresql.service postgresql-16.service

[Service]
Type=simple
User=root
Group=root
WorkingDirectory=/data/iranxia/my-go-backend
EnvironmentFile=-/data/iranxia/my-go-backend/.env
ExecStart=/data/iranxia/my-go-backend/api
Restart=always
RestartSec=5s

[Install]
WantedBy=multi-user.target
EOF

# Create fixed rayancomp-backend service
cat > /etc/systemd/system/rayancomp-backend.service << 'EOF'
[Unit]
Description=RayanComp Go Backend API
After=network.target postgresql.service postgresql-16.service
Wants=postgresql.service postgresql-16.service

[Service]
Type=simple
User=root
Group=root
WorkingDirectory=/data/rayancomp/my-go-backend
EnvironmentFile=-/data/rayancomp/my-go-backend/.env
ExecStart=/data/rayancomp/my-go-backend/api
Restart=always
RestartSec=5s

[Install]
WantedBy=multi-user.target
EOF

# Reload and restart
systemctl daemon-reload
systemctl enable iranxia-backend rayancomp-backend
systemctl restart iranxia-backend rayancomp-backend

echo "Waiting for services to start..."
sleep 3

echo ""
echo "=== iranxia-backend Status ==="
systemctl status iranxia-backend --no-pager || true

echo ""
echo "=== rayancomp-backend Status ==="
systemctl status rayancomp-backend --no-pager || true

echo ""
echo "=== Logs ==="
journalctl -u iranxia-backend -n 10 --no-pager || true
