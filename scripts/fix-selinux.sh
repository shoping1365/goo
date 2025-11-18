#!/bin/bash
# رفع مشکل SELinux برای فایل api

echo "=========================================="
echo "🔒 Fixing SELinux Context for API Binary"
echo "=========================================="
echo ""

# بررسی وضعیت SELinux
if command -v getenforce &> /dev/null; then
    selinux_status=$(getenforce)
    echo "🔍 SELinux Status: $selinux_status"
    
    if [ "$selinux_status" = "Enforcing" ]; then
        echo "⚠️  SELinux is enforcing - fixing context..."
        
        # لیست پروژه‌ها
        PROJECTS=("/data/iranxia" "/data/rayancomp")
        
        for project_dir in "${PROJECTS[@]}"; do
            project_name=$(basename "$project_dir")
            api_path="${project_dir}/my-go-backend/api"
            
            if [ -f "$api_path" ]; then
                echo "🔧 Fixing context for ${project_name}..."
                
                # تنظیم مجوزها
                chmod 750 "$api_path"
                chown root:root "$api_path"
                
                # تنظیم SELinux context
                chcon -t systemd_unit_file_exec_t "$api_path" 2>/dev/null || \
                chcon -t bin_t "$api_path" 2>/dev/null || \
                chcon -t unconfined_exec_t "$api_path" 2>/dev/null || \
                chcon -t default_t "$api_path" 2>/dev/null || true
                
                # نمایش مجوزها
                echo "  📋 Permissions:"
                ls -la "$api_path"
                
                # راه‌اندازی مجدد سرویس
                echo "  🔄 Restarting service..."
                systemctl daemon-reload
                systemctl restart "${project_name}-backend" 2>/dev/null || {
                    systemctl start "${project_name}-backend" 2>/dev/null || {
                        echo "  ❌ Failed to start service"
                        systemctl status "${project_name}-backend" --no-pager
                    }
                }
                
                sleep 2
                if systemctl is-active --quiet "${project_name}-backend"; then
                    echo "  ✅ Service is running"
                else
                    echo "  ❌ Service failed to start"
                fi
                echo ""
            fi
        done
        
    else
        echo "✅ SELinux is not enforcing - no action needed"
    fi
else
    echo "⚠️  SELinux not found - skipping context fix"
fi

echo "=========================================="
echo "✅ SELinux fix completed!"
echo "=========================================="
