package repository

import (
	"context"
)

// FirewallMode نوع حالت امنیتی فایروال
type FirewallMode string

const (
	FirewallModeLow    FirewallMode = "low"
	FirewallModeMedium FirewallMode = "medium"
	FirewallModeHigh   FirewallMode = "high"
	FirewallModeStrict FirewallMode = "strict"
)

// FirewallStatus وضعیت کلی فایروال
type FirewallStatus struct {
	Enabled bool         `json:"enabled"`
	Mode    FirewallMode `json:"mode"`
}

// FirewallRule مدل قانون فایروال (ابسترکت مستقل از OS)
type FirewallRule struct {
	Name        string `json:"name"`
	Direction   string `json:"direction"` // inbound | outbound
	Action      string `json:"action"`    // allow | deny | drop (OS: Block)
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Port        string `json:"port"`
	Enabled     bool   `json:"enabled"`
}

type CreateFirewallRule struct {
	Name        string `json:"name"`
	Direction   string `json:"direction"`
	Action      string `json:"action"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Port        string `json:"port"`
}

type UpdateFirewallRule struct {
	Direction   *string `json:"direction"`
	Action      *string `json:"action"`
	Source      *string `json:"source"`
	Destination *string `json:"destination"`
	Port        *string `json:"port"`
	Enabled     *bool   `json:"enabled"`
}

type FirewallLog struct {
	Timestamp     string `json:"timestamp"`
	SourceIP      string `json:"sourceIP"`
	DestinationIP string `json:"destinationIP"`
	Port          string `json:"port"`
	TrafficType   string `json:"trafficType"`
	Action        string `json:"action"`
}

// FirewallLoggingConfig تنظیمات ثبت لاگ فایروال
type FirewallLoggingConfig struct {
	EnabledDropped bool   `json:"enabledDropped"` // droppedconnections enable/disable
	EnabledAllowed bool   `json:"enabledAllowed"` // allowed connections (successfulconnections)
	FilePath       string `json:"filePath"`
	MaxFileKB      int    `json:"maxFileKB"`
}

// FirewallRepositoryInterface اینترفیس ریپازیتوری فایروال (CQRS-friendly)
type FirewallRepositoryInterface interface {
	// Query
	GetStatus(ctx context.Context) (*FirewallStatus, error)
	ListRules(ctx context.Context) ([]FirewallRule, error)
	ReadLogs(ctx context.Context, limit int) ([]FirewallLog, error)
	GetLogging(ctx context.Context) (*FirewallLoggingConfig, error)

	// Command
	SetEnabled(ctx context.Context, enabled bool) error
	SetMode(ctx context.Context, mode FirewallMode) error
	CreateRule(ctx context.Context, rule CreateFirewallRule) error
	UpdateRule(ctx context.Context, name string, rule UpdateFirewallRule) error
	ToggleRule(ctx context.Context, name string, enabled bool) error
	DeleteRule(ctx context.Context, name string) error
	ClearLogs(ctx context.Context) error
	Backup(ctx context.Context, destPath string) (string, error)
	Restore(ctx context.Context, sourcePath string) error
	SetLogging(ctx context.Context, cfg FirewallLoggingConfig) error
}
