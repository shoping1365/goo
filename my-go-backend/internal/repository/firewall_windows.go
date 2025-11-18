package repository

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// FirewallRepositoryWindows پیاده‌سازی ریپازیتوری فایروال برای Windows (netsh / PowerShell)
type FirewallRepositoryWindows struct{}

func NewFirewallRepositoryWindows() *FirewallRepositoryWindows { return &FirewallRepositoryWindows{} }

// execCommand اجرای امن دستورات سیستم‌عامل با timeouts
func (r *FirewallRepositoryWindows) execCommand(ctx context.Context, name string, args ...string) (string, error) {
	// مدت زمان پیش‌فرض 10 ثانیه
	if _, hasDeadline := ctx.Deadline(); !hasDeadline {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
	}
	cmd := exec.CommandContext(ctx, name, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("command failed: %s %s: %w - %s", name, strings.Join(args, " "), err, string(out))
	}
	return string(out), nil
}

// GetStatus وضعیت فایروال (Windows Defender Firewall)
func (r *FirewallRepositoryWindows) GetStatus(ctx context.Context) (*FirewallStatus, error) {
	// netsh advfirewall show allprofiles
	out, err := r.execCommand(ctx, "netsh", "advfirewall", "show", "allprofiles")
	if err != nil {
		return nil, err
	}
	low := strings.ToLower(out)
	enabled := strings.Contains(low, "state                    on") || strings.Contains(low, "state on")
	// حالت امنیتی را از یک policy ساده استخراج می‌کنیم (Custom mapping based on rules density)
	// اینجا به صورت حداقلی: اگر inbound default = block و outbound default = allow → medium
	mode := FirewallModeMedium
	if strings.Contains(low, "inbound policy              block") && strings.Contains(low, "outbound policy             block") {
		mode = FirewallModeStrict
	} else if strings.Contains(low, "inbound policy              allow") {
		mode = FirewallModeLow
	}
	return &FirewallStatus{Enabled: enabled, Mode: mode}, nil
}

// ListRules لیست قوانین (نام، جهت، اکشن، پورت‌ها)
func (r *FirewallRepositoryWindows) ListRules(ctx context.Context) ([]FirewallRule, error) {
	out, err := r.execCommand(ctx, "netsh", "advfirewall", "firewall", "show", "rule", "name=all")
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(strings.NewReader(out))
	scanner.Split(bufio.ScanLines)
	var rules []FirewallRule
	var cur FirewallRule
	reKV := regexp.MustCompile(`^([^:]+):\s*(.*)$`)
	flush := func() {
		if cur.Name != "" {
			if cur.Direction == "In" {
				cur.Direction = "inbound"
			}
			if cur.Direction == "Out" {
				cur.Direction = "outbound"
			}
			cur.Enabled = strings.EqualFold(cur.Action, cur.Action) // keep as parsed; Enabled may be set below
			rules = append(rules, cur)
			cur = FirewallRule{}
		}
	}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			flush()
			continue
		}
		m := reKV.FindStringSubmatch(line)
		if len(m) != 3 {
			continue
		}
		key := strings.ToLower(strings.TrimSpace(m[1]))
		val := strings.TrimSpace(m[2])
		switch key {
		case "rule name":
			// شروع قانون جدید
			flush()
			cur.Name = val
		case "direction":
			cur.Direction = val
		case "action":
			v := strings.ToLower(val)
			if v == "allow" {
				cur.Action = "allow"
			} else {
				cur.Action = "deny"
			}
		case "localport":
			if val == "Any" {
				cur.Port = "any"
			} else {
				cur.Port = val
			}
		case "enabled":
			cur.Enabled = strings.EqualFold(val, "Yes")
		case "remoteip":
			cur.Source = val
		case "localip":
			cur.Destination = val
		}
	}
	flush()
	return rules, nil
}

// ReadLogs خواندن لاگ‌ها (Windows Event Log نداریم اینجا؛ از "netsh trace" استفاده نمی‌کنیم). Placeholder خالی.
func (r *FirewallRepositoryWindows) ReadLogs(ctx context.Context, limit int) ([]FirewallLog, error) {
	if limit <= 0 {
		limit = 200
	}
	// مسیر پیش‌فرض لاگ فایروال ویندوز
	systemRoot := os.Getenv("SystemRoot")
	if systemRoot == "" {
		systemRoot = `C:\\Windows`
	}
	logPath := filepath.Join(systemRoot, "System32", "LogFiles", "Firewall", "pfirewall.log")
	f, err := os.Open(logPath)
	if err != nil {
		// اگر لاگ فعال نیست، خطا نمی‌دهیم و لیست خالی برمی‌گردانیم
		if errors.Is(err, os.ErrNotExist) {
			return []FirewallLog{}, nil
		}
		return nil, err
	}
	defer f.Close()

	// خواندن انتهای فایل برای آخرین ردیف‌ها
	// به سادگی کل فایل را می‌خوانیم و آخرین N رکورد غیر # را برمی‌داریم
	content, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	var logs []FirewallLog
	for i := len(lines) - 1; i >= 0 && len(logs) < limit; i-- {
		line := strings.TrimSpace(lines[i])
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		// فیلدها با فاصله جدا شده‌اند
		parts := strings.Fields(line)
		if len(parts) < 8 {
			continue
		}
		// حداقل: date time action protocol src-ip dst-ip src-port dst-port
		dateStr := parts[0]
		timeStr := parts[1]
		action := strings.ToLower(parts[2]) // allow | drop
		protocol := strings.ToUpper(parts[3])
		srcIP := parts[4]
		dstIP := parts[5]
		// پورت‌ها ممکن است '-'
		dstPort := parts[7]
		if dstPort == "-" {
			dstPort = ""
		}
		log := FirewallLog{
			Timestamp:     fmt.Sprintf("%s %s", dateStr, timeStr),
			SourceIP:      srcIP,
			DestinationIP: dstIP,
			Port:          dstPort,
			TrafficType:   protocol,
			Action:        map[string]string{"allow": "allowed", "drop": "blocked"}[action],
		}
		if log.Action == "" {
			log.Action = action
		}
		logs = append(logs, log)
	}
	// معکوس برای ترتیب زمانی قدیم → جدید
	for i, j := 0, len(logs)-1; i < j; i, j = i+1, j-1 {
		logs[i], logs[j] = logs[j], logs[i]
	}
	return logs, nil
}

// SetEnabled فعال/غیرفعال‌سازی فایروال سیستم
func (r *FirewallRepositoryWindows) SetEnabled(ctx context.Context, enabled bool) error {
	state := "off"
	if enabled {
		state = "on"
	}
	_, err := r.execCommand(ctx, "netsh", "advfirewall", "set", "allprofiles", fmt.Sprintf("state %s", state))
	return err
}

// SetMode تنظیم حالت امنیتی ساده: low/medium/high/strict → روی inbound/outbound policy
func (r *FirewallRepositoryWindows) SetMode(ctx context.Context, mode FirewallMode) error {
	switch mode {
	case FirewallModeLow:
		// inbound allow, outbound allow (کمترین محدودیت)
		if _, err := r.execCommand(ctx, "netsh", "advfirewall", "set", "allprofiles", "firewallpolicy", "allowinbound,allowoutbound"); err != nil {
			return err
		}
	case FirewallModeMedium:
		// inbound block, outbound allow
		if _, err := r.execCommand(ctx, "netsh", "advfirewall", "set", "allprofiles", "firewallpolicy", "blockinbound,allowoutbound"); err != nil {
			return err
		}
	case FirewallModeHigh, FirewallModeStrict:
		// inbound block, outbound block
		if _, err := r.execCommand(ctx, "netsh", "advfirewall", "set", "allprofiles", "firewallpolicy", "blockinbound,blockoutbound"); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported mode: %s", mode)
	}
	return nil
}

// CreateRule افزودن قانون جدید
func (r *FirewallRepositoryWindows) CreateRule(ctx context.Context, rule CreateFirewallRule) error {
	dir := "in"
	if strings.EqualFold(rule.Direction, "outbound") {
		dir = "out"
	}
	action := "allow"
	if strings.ToLower(rule.Action) != "allow" {
		action = "block"
	}
	localIP := rule.Destination
	remoteIP := rule.Source
	localPort := rule.Port
	if localPort == "" {
		localPort = "any"
	}
	args := []string{"advfirewall", "firewall", "add", "rule", fmt.Sprintf("name=%s", rule.Name), fmt.Sprintf("dir=%s", dir), fmt.Sprintf("action=%s", action)}
	if localIP != "" {
		args = append(args, fmt.Sprintf("localip=%s", localIP))
	}
	if remoteIP != "" {
		args = append(args, fmt.Sprintf("remoteip=%s", remoteIP))
	}
	if localPort != "" {
		args = append(args, fmt.Sprintf("localport=%s", localPort))
	}
	_, err := r.execCommand(ctx, "netsh", args...)
	return err
}

// UpdateRule بروزرسانی قانون بر اساس نام (netsh حذف و ایجاد مجدد)
func (r *FirewallRepositoryWindows) UpdateRule(ctx context.Context, name string, upd UpdateFirewallRule) error {
	// ساده‌ترین راه: حذف و ایجاد مجدد با پارامترهای جدید
	if err := r.DeleteRule(ctx, name); err != nil {
		return err
	}
	newRule := CreateFirewallRule{Name: name}
	if upd.Direction != nil {
		newRule.Direction = *upd.Direction
	}
	if upd.Action != nil {
		newRule.Action = *upd.Action
	}
	if upd.Source != nil {
		newRule.Source = *upd.Source
	}
	if upd.Destination != nil {
		newRule.Destination = *upd.Destination
	}
	if upd.Port != nil {
		newRule.Port = *upd.Port
	}
	return r.CreateRule(ctx, newRule)
}

// ToggleRule فعال/غیرفعال با Enable/Disable rule name
func (r *FirewallRepositoryWindows) ToggleRule(ctx context.Context, name string, enabled bool) error {
	state := "disable"
	if enabled {
		state = "enable"
	}
	_, err := r.execCommand(ctx, "netsh", "advfirewall", "firewall", state, fmt.Sprintf("rule name=%s", name))
	return err
}

// DeleteRule حذف قانون بر اساس نام
func (r *FirewallRepositoryWindows) DeleteRule(ctx context.Context, name string) error {
	_, err := r.execCommand(ctx, "netsh", "advfirewall", "firewall", "delete", "rule", fmt.Sprintf("name=%s", name))
	return err
}

// ClearLogs پاک‌سازی لاگ‌ها (بدون منبع خاص OS → no-op)
func (r *FirewallRepositoryWindows) ClearLogs(ctx context.Context) error {
	systemRoot := os.Getenv("SystemRoot")
	if systemRoot == "" {
		systemRoot = `C:\\Windows`
	}
	logPath := filepath.Join(systemRoot, "System32", "LogFiles", "Firewall", "pfirewall.log")
	// truncate
	return os.Truncate(logPath, 0)
}

// Backup پشتیبان قوانین به فایل WFW با استفاده از netsh advfirewall export
func (r *FirewallRepositoryWindows) Backup(ctx context.Context, destPath string) (string, error) {
	if destPath == "" {
		destPath = "firewall-backup.wfw"
	}
	_, err := r.execCommand(ctx, "netsh", "advfirewall", "export", destPath)
	if err != nil {
		return "", err
	}
	return destPath, nil
}

// Restore بازگردانی قوانین از فایل WFW
func (r *FirewallRepositoryWindows) Restore(ctx context.Context, sourcePath string) error {
	if sourcePath == "" {
		return fmt.Errorf("sourcePath is required")
	}
	_, err := r.execCommand(ctx, "netsh", "advfirewall", "import", sourcePath)
	return err
}

// GetLogging دریافت تنظیمات لاگ
func (r *FirewallRepositoryWindows) GetLogging(ctx context.Context) (*FirewallLoggingConfig, error) {
	out, err := r.execCommand(ctx, "netsh", "advfirewall", "show", "currentprofile")
	if err != nil {
		return nil, err
	}
	low := strings.ToLower(out)
	cfg := &FirewallLoggingConfig{FilePath: "", MaxFileKB: 4096}
	if strings.Contains(low, "log dropped packets:               yes") {
		cfg.EnabledDropped = true
	}
	if strings.Contains(low, "log successful connections:        yes") {
		cfg.EnabledAllowed = true
	}
	for _, line := range strings.Split(out, "\n") {
		l := strings.TrimSpace(line)
		if strings.HasPrefix(strings.ToLower(l), "file name:") {
			cfg.FilePath = strings.TrimSpace(strings.TrimPrefix(l, "File name:"))
		}
		if strings.HasPrefix(strings.ToLower(l), "maximum file size:") {
			v := strings.TrimSpace(strings.TrimPrefix(l, "Maximum file size:"))
			if n, err := strconv.Atoi(strings.Fields(v)[0]); err == nil {
				cfg.MaxFileKB = n
			}
		}
	}
	return cfg, nil
}

// SetLogging تنظیم لاگ فایروال ویندوز
func (r *FirewallRepositoryWindows) SetLogging(ctx context.Context, cfg FirewallLoggingConfig) error {
	if cfg.MaxFileKB <= 0 {
		cfg.MaxFileKB = 4096
	}
	if _, err := r.execCommand(ctx, "netsh", "advfirewall", "set", "allprofiles", "logging", fmt.Sprintf("droppedconnections=%s", ternary(cfg.EnabledDropped, "enable", "disable"))); err != nil {
		return err
	}
	if _, err := r.execCommand(ctx, "netsh", "advfirewall", "set", "allprofiles", "logging", fmt.Sprintf("successfulconnections=%s", ternary(cfg.EnabledAllowed, "enable", "disable"))); err != nil {
		return err
	}
	if cfg.FilePath != "" {
		if _, err := r.execCommand(ctx, "netsh", "advfirewall", "set", "allprofiles", "logging", fmt.Sprintf("filename=%s", cfg.FilePath)); err != nil {
			return err
		}
	}
	if _, err := r.execCommand(ctx, "netsh", "advfirewall", "set", "allprofiles", "logging", fmt.Sprintf("maxfilesize=%d", cfg.MaxFileKB)); err != nil {
		return err
	}
	return nil
}

func ternary[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}

var _ FirewallRepositoryInterface = (*FirewallRepositoryWindows)(nil)
