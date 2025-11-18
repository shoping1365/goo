#!/bin/bash
# Debug and fix backend binary issues

echo "🔍 Debugging backend binary issue..."
echo ""

PROJECT_DIR="/data/iranxia"

echo "1️⃣ Checking if binary file exists..."
if [ -f "${PROJECT_DIR}/my-go-backend/api" ]; then
    echo "✅ Binary exists"
    ls -lh "${PROJECT_DIR}/my-go-backend/api"
else
    echo "❌ Binary NOT found at ${PROJECT_DIR}/my-go-backend/api"
fi

echo ""
echo "2️⃣ Checking if main.go exists..."
if [ -f "${PROJECT_DIR}/my-go-backend/cmd/api/main.go" ]; then
    echo "✅ main.go exists"
    ls -lh "${PROJECT_DIR}/my-go-backend/cmd/api/main.go"
else
    echo "❌ main.go NOT found"
fi

echo ""
echo "3️⃣ Checking directory permissions..."
ls -ld "${PROJECT_DIR}/my-go-backend"
ls -ld "${PROJECT_DIR}/my-go-backend/cmd" 2>/dev/null || echo "cmd directory not found"

echo ""
echo "4️⃣ Checking SELinux context..."
ls -Z "${PROJECT_DIR}/my-go-backend/" 2>/dev/null | grep api || echo "No binary found or SELinux not enabled"

echo ""
echo "5️⃣ Attempting to build binary..."
cd "${PROJECT_DIR}/my-go-backend"

if [ -f "cmd/api/main.go" ]; then
    echo "Building Go binary..."
    go build -o api ./cmd/api/main.go || {
        echo "❌ Build failed!"
        exit 1
    }
    
    echo "✅ Build successful!"
    
    # Set permissions
    chmod +x api
    chown root:root api
    
    # Fix SELinux context
    restorecon -v api 2>/dev/null || chcon -t bin_t api 2>/dev/null || echo "SELinux fix skipped"
    
    echo ""
    echo "6️⃣ Final check..."
    ls -lZ api 2>/dev/null || ls -l api
    
    echo ""
    echo "7️⃣ Testing if binary runs..."
    ./api --version 2>&1 || ./api -h 2>&1 | head -5 || echo "Binary test complete (no --version flag)"
    
else
    echo "❌ Cannot build - main.go not found!"
    exit 1
fi

echo ""
echo "✅ All checks complete!"
echo ""
echo "Now restart the service:"
echo "  systemctl restart iranxia-backend"
echo "  systemctl status iranxia-backend"
