//go:build !windows
// +build !windows

package repository

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

// FirewallRepositoryLinux پیاده‌سازی ریپازیتوری فایروال برای Linux (firewalld)
type FirewallRepositoryWindows struct{}

func NewFirewallRepositoryWindows() *FirewallRepositoryWindows {
	return &FirewallRepositoryWindows{}
}

// execCommand اجرای امن دستورات سیستم‌عامل با timeouts
func (r *FirewallRepositoryWindows) execCommand(ctx context.Context, name string, args ...string) (string, error) {
	// مدت زمان پیش‌فرض 10 ثانیه
	if _, hasDeadline := ctx.Deadline(); !hasDeadline {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
	}

	cmd := exec.CommandContext(ctx, name, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("command failed: %s, error: %w, output: %s", name, err, string(output))
	}

	return strings.TrimSpace(string(output)), nil
}

// GetStatus دریافت وضعیت فایروال
func (r *FirewallRepositoryWindows) GetStatus(ctx context.Context) (*FirewallStatus, error) {
	if _, err := exec.LookPath("firewall-cmd"); err != nil {
		return &FirewallStatus{Enabled: false, Mode: FirewallModeLow}, nil
	}

	output, err := r.execCommand(ctx, "firewall-cmd", "--state")
	enabled := err == nil && strings.Contains(output, "running")

	mode := FirewallModeMedium
	if enabled {
		mode = FirewallModeHigh
	}

	return &FirewallStatus{
		Enabled: enabled,
		Mode:    mode,
	}, nil
}

// ListRules لیست کردن قوانین فایروال
func (r *FirewallRepositoryWindows) ListRules(ctx context.Context) ([]FirewallRule, error) {
	if _, err := exec.LookPath("firewall-cmd"); err != nil {
		return []FirewallRule{}, nil
	}

	output, err := r.execCommand(ctx, "firewall-cmd", "--list-all")
	if err != nil {
		return nil, err
	}

	var rules []FirewallRule
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if strings.Contains(line, "ports:") {
			parts := strings.Fields(line)
			for _, part := range parts[1:] {
				rules = append(rules, FirewallRule{
					Name:      part,
					Direction: "inbound",
					Action:    "allow",
					Port:      part,
					Enabled:   true,
				})
			}
		}
	}

	return rules, nil
}

// ReadLogs خواندن لاگ‌های فایروال
func (r *FirewallRepositoryWindows) ReadLogs(ctx context.Context, limit int) ([]FirewallLog, error) {
	return []FirewallLog{}, nil
}

// GetLogging دریافت تنظیمات لاگ
func (r *FirewallRepositoryWindows) GetLogging(ctx context.Context) (*FirewallLoggingConfig, error) {
	return &FirewallLoggingConfig{
		EnabledDropped: false,
		EnabledAllowed: false,
		FilePath:       "/var/log/firewalld",
		MaxFileKB:      1024,
	}, nil
}

// SetEnabled فعال/غیرفعال کردن فایروال
func (r *FirewallRepositoryWindows) SetEnabled(ctx context.Context, enabled bool) error {
	if _, err := exec.LookPath("firewall-cmd"); err != nil {
		return nil
	}

	var cmd string
	if enabled {
		cmd = "start"
	} else {
		cmd = "stop"
	}

	_, err := r.execCommand(ctx, "systemctl", cmd, "firewalld")
	return err
}

// SetMode تنظیم mode فایروال
func (r *FirewallRepositoryWindows) SetMode(ctx context.Context, mode FirewallMode) error {
	return nil
}

// CreateRule ایجاد قانون جدید
func (r *FirewallRepositoryWindows) CreateRule(ctx context.Context, rule CreateFirewallRule) error {
	if _, err := exec.LookPath("firewall-cmd"); err != nil {
		return nil
	}

	// استفاده از Port به جای LocalPort
	portRule := rule.Port
	_, err := r.execCommand(ctx, "firewall-cmd", "--permanent", "--add-port="+portRule)
	if err != nil {
		return err
	}

	_, err = r.execCommand(ctx, "firewall-cmd", "--reload")
	return err
}

// UpdateRule آپدیت قانون موجود
func (r *FirewallRepositoryWindows) UpdateRule(ctx context.Context, name string, rule UpdateFirewallRule) error {
	return nil
}

// ToggleRule فعال/غیرفعال کردن قانون
func (r *FirewallRepositoryWindows) ToggleRule(ctx context.Context, name string, enabled bool) error {
	return nil
}

// DeleteRule حذف قانون
func (r *FirewallRepositoryWindows) DeleteRule(ctx context.Context, name string) error {
	if _, err := exec.LookPath("firewall-cmd"); err != nil {
		return nil
	}

	_, err := r.execCommand(ctx, "firewall-cmd", "--permanent", "--remove-port="+name)
	if err != nil {
		return err
	}

	_, err = r.execCommand(ctx, "firewall-cmd", "--reload")
	return err
}

// ClearLogs پاک کردن لاگ‌ها
func (r *FirewallRepositoryWindows) ClearLogs(ctx context.Context) error {
	return nil
}

// Backup پشتیبان‌گیری از تنظیمات فایروال
func (r *FirewallRepositoryWindows) Backup(ctx context.Context, destPath string) (string, error) {
	if _, err := exec.LookPath("firewall-cmd"); err != nil {
		return "", nil
	}

	output, err := r.execCommand(ctx, "firewall-cmd", "--list-all")
	if err != nil {
		return "", err
	}

	return output, nil
}

// Restore بازگردانی از پشتیبان
func (r *FirewallRepositoryWindows) Restore(ctx context.Context, sourcePath string) error {
	return nil
}

// SetLogging تنظیم لاگینگ
func (r *FirewallRepositoryWindows) SetLogging(ctx context.Context, cfg FirewallLoggingConfig) error {
	return nil
}
