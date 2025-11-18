#!/bin/bash
# Setup passwordless sudo for GitHub Actions runner user
# Run this on the server as root or with sudo

set -e

RUNNER_USER="arash"

echo "=========================================="
echo "🔧 Setting up passwordless sudo for ${RUNNER_USER}"
echo "=========================================="

# Check if running as root
if [ "$EUID" -ne 0 ]; then 
    echo "❌ Please run as root or with sudo"
    echo "Usage: sudo bash scripts/setup-runner-sudo.sh"
    exit 1
fi

# Create sudoers file for runner user
SUDOERS_FILE="/etc/sudoers.d/github-actions-runner"

echo "Creating sudoers configuration..."
cat > "${SUDOERS_FILE}" << 'EOF'
# GitHub Actions Runner - Passwordless sudo for deployment
# Created automatically - DO NOT EDIT MANUALLY

# Allow runner user to execute all commands without password
arash ALL=(ALL) NOPASSWD: ALL

# Specific commands (more secure alternative - uncomment if preferred)
# arash ALL=(ALL) NOPASSWD: /bin/mkdir, /bin/chown, /bin/chmod, /usr/bin/rsync
# arash ALL=(ALL) NOPASSWD: /bin/bash /opt/shared/scripts/*.sh
# arash ALL=(ALL) NOPASSWD: /usr/bin/systemctl
# arash ALL=(ALL) NOPASSWD: /usr/bin/pm2
EOF

# Set correct permissions (sudoers files must be 0440)
chmod 0440 "${SUDOERS_FILE}"

# Validate sudoers syntax
if visudo -c -f "${SUDOERS_FILE}"; then
    echo "✅ Sudoers configuration created successfully"
    echo "📁 Location: ${SUDOERS_FILE}"
else
    echo "❌ Invalid sudoers syntax! Removing file..."
    rm -f "${SUDOERS_FILE}"
    exit 1
fi

# Test sudo access
echo ""
echo "Testing sudo access for ${RUNNER_USER}..."
if sudo -u "${RUNNER_USER}" sudo -n true 2>/dev/null; then
    echo "✅ Passwordless sudo is working for ${RUNNER_USER}"
else
    echo "⚠️  Test failed, but configuration is in place"
    echo "Try logging out and back in, or restart the runner service"
fi

echo ""
echo "=========================================="
echo "✅ Setup completed!"
echo "=========================================="
echo ""
echo "Next steps:"
echo "1. Restart GitHub Actions runner service:"
echo "   sudo systemctl restart actions.runner.*.service"
echo ""
echo "2. Or if using svc.sh:"
echo "   cd /opt/deployment/actions-runner"
echo "   sudo ./svc.sh restart"
echo ""
echo "3. Test with: sudo -u arash sudo -n whoami"
echo "   Should return 'root' without asking for password"
echo ""
