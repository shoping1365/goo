package main

import (
	"context"
	"crypto/tls"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"my-go-backend/internal/repository"
	"my-go-backend/internal/router"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"my-go-backend/internal/database"
	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/handlers"
	searchbootstrap "my-go-backend/internal/search/bootstrap"
	searchconfig "my-go-backend/internal/search/config"
	searchengine "my-go-backend/internal/search/engine"
	searchhttp "my-go-backend/internal/search/httpserver"
	searchservice "my-go-backend/internal/search/service"
	searchworker "my-go-backend/internal/search/worker"
	"my-go-backend/internal/services"
	"my-go-backend/internal/sms"

	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"golang.org/x/net/http2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	gormDB *gorm.DB
	dbSize int64 // database size in bytes
)

// maskPassword hides passwords before logging
func maskPassword(password string) string {
	if password == "" {
		return "(empty)"
	}
	if len(password) <= 4 {
		return "****"
	}
	return password[:2] + "****" + password[len(password)-2:]
}

func dbMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "db", gormDB)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func main() {
	// Load environment variables from .env file
	// Try loading the .env file from several locations
	envPaths := []string{
		".env",
		filepath.Join("..", ".env"),
		filepath.Join("..", "..", ".env"),
		"/data/iranxia/.env",
		"/opt/shared/.env",
	}

	envLoaded := false
	for _, envPath := range envPaths {
		if err := godotenv.Load(envPath); err == nil {
			log.Printf("✅ .env file loaded from '%s'", envPath)
			envLoaded = true
			break
		}
	}

	if !envLoaded {
		log.Println("⚠️  Warning: .env file not found in any of the following paths:")
		for _, envPath := range envPaths {
			log.Printf("   - %s", envPath)
		}
		log.Println("   Continuing with system environment variables...")
	}

	// Force build connection string from individual DB variables
	// to avoid parsing issues with DATABASE_URL.
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	// Debug: print the resolved database variables
	log.Printf("🔍 Database environment variables:")
	log.Printf("   DB_HOST: '%s'", dbHost)
	log.Printf("   DB_PORT: '%s'", dbPort)
	log.Printf("   DB_USER: '%s'", dbUser)
	log.Printf("   DB_PASSWORD: '%s'", maskPassword(dbPassword))
	log.Printf("   DB_NAME: '%s'", dbName)
	log.Printf("   DB_SSLMODE: '%s'", dbSSLMode)

	if dbUser == "" || dbHost == "" || dbPort == "" || dbName == "" {
		log.Fatal("❌ Error: One or more required database environment variables (DB_USER, DB_HOST, DB_PORT, DB_NAME) are not set.")
	}

	// Construct the database URL manually with URL encoding
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser,
		url.QueryEscape(dbPassword), // URL encode the password
		dbHost,
		dbPort,
		dbName,
		dbSSLMode,
	)
	log.Printf("📝 DATABASE_URL constructed: '%s'", databaseURL)

	log.Println("========================================")
	log.Println("🚀 Starting server bootstrap...")
	log.Println("========================================")

	mainGin()
}

func apiStatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	apis := []map[string]interface{}{
		{"id": 1, "name": "Primary API", "url": os.Getenv("MAIN_API_URL"), "status": "success", "lastChecked": time.Now()},
		{"id": 2, "name": "Payment API", "url": os.Getenv("PAYMENT_API_URL"), "status": "success", "lastChecked": time.Now()},
	}
	json.NewEncoder(w).Encode(apis)
}

func databaseStatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := gormDB.Raw("SELECT pg_database_size(current_database())").Scan(&dbSize); err != nil {
		http.Error(w, "Database connection is down", http.StatusInternalServerError)
		return
	}

	var tableCount int
	gormDB.Raw("SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = 'public'").Scan(&tableCount)

	var dbName string
	gormDB.Raw("SELECT current_database()").Scan(&dbName)

	var uptimeStr string
	gormDB.Raw("SELECT now() - pg_postmaster_start_time() AS uptime").Scan(&uptimeStr)

	sizeMB := float64(dbSize) / (1024 * 1024)

	json.NewEncoder(w).Encode([]map[string]interface{}{{
		"id": 1, "name": dbName, "type": "PostgreSQL", "status": "success", "size": fmt.Sprintf("%.2f MB", sizeMB), "tables": tableCount, "uptime": uptimeStr,
	}})
}

func serverStatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cpuPercent, _ := cpu.Percent(time.Second, false)
	memInfo, _ := mem.VirtualMemory()
	diskInfo, _ := disk.Usage("/")
	var cpuVal float64
	if len(cpuPercent) > 0 {
		cpuVal = cpuPercent[0]
	}
	json.NewEncoder(w).Encode([]map[string]interface{}{{
		"id":     1,
		"name":   getHostname(),
		"ip":     os.Getenv("SERVER_IP"),
		"status": "success",
		"cpu":    cpuVal,
		"ram":    memInfo.UsedPercent,
		"disk":   diskInfo.UsedPercent,
	}})
}

func getHostname() string {
	host, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return host
}

func logsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rows, err := gormDB.Raw("SELECT id, timestamp, level, message, details FROM system_logs ORDER BY timestamp DESC LIMIT 100").Rows()
	if err != nil {
		http.Error(w, "Failed to query logs", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var logs []map[string]interface{}
	for rows.Next() {
		var id int
		var timestamp time.Time
		var level, message, details string
		if err := rows.Scan(&id, &timestamp, &level, &message, &details); err != nil {
			continue
		}
		logs = append(logs, map[string]interface{}{
			"id": id, "timestamp": timestamp, "level": level, "message": message, "details": details,
		})
	}

	if err = rows.Err(); err != nil {
		http.Error(w, "Error iterating log rows", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(logs)
}

func clearLogsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := gormDB.Exec("TRUNCATE TABLE system_logs"); err != nil {
		http.Error(w, "Failed to clear logs", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Logs cleared"})
}

func testAPIHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Path[len("/api/admin/system/test-api/"):]
	var apiURL string

	if id == "1" {
		apiURL = os.Getenv("MAIN_API_URL")
	} else if id == "2" {
		apiURL = os.Getenv("PAYMENT_API_URL")
	} else {
		http.Error(w, "Invalid API ID", http.StatusBadRequest)
		return
	}

	resp, err := http.Get(apiURL)
	status := "success"
	if err != nil || (resp != nil && resp.StatusCode >= 400) {
		status = "error"
	}
	if resp != nil {
		defer resp.Body.Close()
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"id": id, "status": status, "lastChecked": time.Now()})
}

func testDatabaseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := gormDB.Raw("SELECT 1").Error; err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"id": 1, "status": "error", "details": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"id": 1, "status": "success"})
}

func testServerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cpuPercent, _ := cpu.Percent(time.Second, false)
	memInfo, _ := mem.VirtualMemory()
	diskInfo, _ := disk.Usage("/")
	status := "success"
	var cpuVal float64
	if len(cpuPercent) > 0 {
		cpuVal = cpuPercent[0]
		if cpuVal > 90 || memInfo.UsedPercent > 90 || diskInfo.UsedPercent > 90 {
			status = "warning"
		}
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"id": 1, "status": status, "cpu": cpuVal, "ram": memInfo.UsedPercent, "disk": diskInfo.UsedPercent})
}

func insertSystemLog(level, message, details string) {
	if err := gormDB.Exec("INSERT INTO system_logs (timestamp, level, message, details) VALUES (NOW(), $1, $2, $3)", level, message, details).Error; err != nil {
		// ERROR: Failed to insert system log
	}
}

func clearCacheHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cmd := exec.Command("bash", "clear_redis_cache.sh")
	output, err := cmd.CombinedOutput()
	if err != nil {
		insertSystemLog("error", "Cache clear failed", err.Error()+"\n"+string(output))
		http.Error(w, "Cache clear failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	insertSystemLog("info", "Cache cleared", string(output))
	json.NewEncoder(w).Encode(map[string]string{"message": "Cache cleared successfully", "output": string(output)})
}

func resetDatabaseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cmd := exec.Command("bash", "reset_db.sh")
	output, err := cmd.CombinedOutput()
	if err != nil {
		insertSystemLog("error", "Database reset failed", err.Error()+"\n"+string(output))
		http.Error(w, "Database reset failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	insertSystemLog("info", "Database reset completed", string(output))
	json.NewEncoder(w).Encode(map[string]string{"message": "Database reset succeeded", "output": string(output)})
}

func restartServerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cmd := exec.Command("bash", "-c", "sudo reboot")
	output, err := cmd.CombinedOutput()
	if err != nil {
		insertSystemLog("error", "Server reboot failed", err.Error()+"\n"+string(output))
		http.Error(w, "Server reboot failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	insertSystemLog("info", "Server reboot requested", string(output))
	json.NewEncoder(w).Encode(map[string]string{"message": "Server reboot requested", "output": string(output)})
}

func dbMonitorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := gormDB.Raw(`
		SELECT pid, usename, client_addr, state, query, now() - query_start AS duration, query_start
		FROM pg_stat_activity
		WHERE state != 'idle' AND query NOT ILIKE '%pg_stat_activity%' AND query NOT ILIKE '%pg_stat_statements%'
		ORDER BY duration DESC LIMIT 20
	`).Rows()
	if err != nil {
		http.Error(w, "Error fetching active queries", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var activeQueries []map[string]interface{}
	for rows.Next() {
		var pid int
		var usename, state, query string
		var client_addr sql.NullString
		var duration string
		var queryStart time.Time
		if err := rows.Scan(&pid, &usename, &client_addr, &state, &query, &duration, &queryStart); err == nil {
			activeQueries = append(activeQueries, map[string]interface{}{
				"pid":         pid,
				"usename":     usename,
				"client_addr": client_addr.String,
				"state":       state,
				"query":       query,
				"duration":    duration,
				"query_start": queryStart,
			})
		}
	}

	var slowQueries []map[string]interface{}
	slowRows, err := gormDB.Raw(`
		SELECT queryid, calls,
			ROUND(CAST(total_exec_time / NULLIF(calls,0) AS numeric), 2) AS mean_time,
			total_exec_time,
			query
		FROM pg_stat_statements
		ORDER BY mean_time DESC
		LIMIT 20
	`).Rows()
	if err != nil {
		// Could not fetch slow queries (is pg_stat_statements enabled?)
	} else {
		defer slowRows.Close()
		for slowRows.Next() {
			var queryid int64
			var calls int64
			var mean_time, total_exec_time float64
			var query string
			if err := slowRows.Scan(&queryid, &calls, &mean_time, &total_exec_time, &query); err == nil {
				slowQueries = append(slowQueries, map[string]interface{}{
					"queryid":         queryid,
					"calls":           calls,
					"mean_time":       mean_time,
					"total_exec_time": total_exec_time,
					"query":           query,
				})
			}
		}
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"activeQueries": activeQueries,
		"slowQueries":   slowQueries,
	})
}

func mainGin() {
	log.Println("========================================")
	log.Println("📡 STEP 1: Reading database configuration...")
	log.Println("========================================")

	// Database connection
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	if dbUser == "" || dbHost == "" || dbPort == "" || dbName == "" {
		log.Fatal("❌ Error: One or more required database environment variables (DB_USER, DB_HOST, DB_PORT, DB_NAME) are not set.")
	}

	// Construct the database URL
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, url.QueryEscape(dbPassword), dbHost, dbPort, dbName, dbSSLMode)

	log.Println("========================================")
	log.Println("📡 STEP 2: Configuring GORM logger...")
	log.Println("========================================")

	// Configure GORM logger
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent, // تغییر به Silent
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	log.Println("========================================")
	log.Println("📡 STEP 3: Connecting to database...")
	log.Println("========================================")

	// Connect to database using GORM
	var err error
	gormDB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	log.Println("✅ Database connection established")

	log.Println("========================================")
	log.Println("📡 STEP 4: Tuning connection pool...")
	log.Println("========================================")

	// Database pool tuning
	if sqlDB, err := gormDB.DB(); err == nil {
		cpu := runtime.NumCPU()
		if cpu < 1 {
			cpu = 1
		}
		sqlDB.SetMaxIdleConns(cpu)
		sqlDB.SetMaxOpenConns(cpu * 2)
		sqlDB.SetConnMaxLifetime(1 * time.Hour)
		log.Printf("   MaxIdleConns: %d", cpu)
		log.Printf("   MaxOpenConns: %d", cpu*2)
		log.Printf("   ConnMaxLifetime: 1 hour")
	}

	database.GormDB = gormDB

	log.Println("========================================")
	log.Println("📡 STEP 5: Testing database connection...")
	log.Println("========================================")

	// Test database connection
	if err := gormDB.Raw("SELECT 1").Error; err != nil {
		log.Fatalf("❌ Failed to ping database: %v", err)
	}
	log.Println("✅ Database ping successful")

	log.Println("========================================")
	log.Println("📡 STEP 6: Running migrations...")
	log.Println("⚠️  This step may take a few minutes")
	log.Println("========================================")

	// Run migrations
	startTime := time.Now()
	if err := database.RunMigrations(gormDB); err != nil {
		log.Fatalf("❌ Migration error: %v", err)
	}
	log.Printf("✅ All migrations executed successfully (total time: %v)", time.Since(startTime))

	log.Println("========================================")
	log.Println("📡 STEP 6.5: Preparing search indexes...")
	log.Println("========================================")

	if err := ensureSearchIndexes(); err != nil {
		log.Fatalf("❌ Failed to prepare search indexes: %v", err)
	}

	log.Println("========================================")
	log.Println("📡 STEP 7: Bootstrapping services...")
	log.Println("========================================")

	// --------------------------------------------------------------------
	// Initialize Redis (for settings cache)
	// --------------------------------------------------------------------
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDB := 0
	redisClient := redis.NewClient(&redis.Options{Addr: redisAddr, Password: redisPassword, DB: redisDB})

	// Instantiate Setting service & handler
	settingService := services.NewSettingService(gormDB, redisClient)
	settingsService := services.NewSettingsService(gormDB)
	settingHandler := handlers.NewSettingHandler(settingService, settingsService)
	adminSettingHandler := handlers.NewAdminSettingHandler(gormDB)

	// Initialize payment services and repositories
	paymentGatewayRepo := repository.NewPaymentGatewayRepository(gormDB)
	paymentTransactionRepo := repository.NewPaymentTransactionRepository(gormDB)
	paymentServiceFactory := services.NewPaymentServiceFactory()
	paymentHandler := handlers.NewPaymentHandler(paymentGatewayRepo, paymentTransactionRepo, paymentServiceFactory)
	mellatHandler := handlers.NewMellatHandler(paymentGatewayRepo, paymentTransactionRepo)

	// Initialize SMS Gateway Manager
	_, err = sms.NewGatewayManager(gormDB)
	if err != nil {
		// Failed to create SMS Gateway Manager
		return
	}

	// Initialize chat UnitOfWork factory and service
	uowFactory := unitofwork.NewUnitOfWorkFactory(gormDB)
	chatService, err := services.NewChatService(uowFactory)
	if err != nil {
		// Failed to create chat service
		return
	}

	// Encryption service (for WebSocket messages)
	encryptionSvc, err := services.NewEncryptionService()
	if err != nil {
		// Failed to create encryption service
		return
	}

	// WebSocket manager
	wsManager, err := services.NewWebSocketManager(gormDB, encryptionSvc)
	if err != nil {
		// Failed to create WebSocket manager
		return
	}

	// Chat handler (needs WebSocket manager and UnitOfWorkFactory)
	chatHandler := handlers.NewChatHandler(chatService, wsManager, uowFactory)

	// Initialize monitoring services and handlers
	monitoringService := services.NewMonitoringService(gormDB)
	monitoringHandler := handlers.NewMonitoringHandler(monitoringService)

	// Start automatic monitoring (enabled by default)
	monitoringService.EnableMonitoring()

	// --------------------------------------------------------------------
	// راه‌اندازی سرویس جستجو (موتور داخلی Bleve)
	// --------------------------------------------------------------------
	searchCfg, err := searchconfig.Load()
	if err != nil {
		log.Fatalf("❌ Search service configuration error: %v", err)
	}

	searchClient, err := searchengine.NewBleveClient(searchengine.Options{
		Directory:      searchCfg.IndexDir,
		AllowedIndexes: searchCfg.AllowedIndexes,
		SuggestFields:  searchCfg.SuggestFields,
	})
	if err != nil {
		log.Fatalf("❌ Failed to open search indexes: %v", err)
	}
	defer func() {
		if err := searchClient.Close(); err != nil {
			log.Printf("⚠️  search engine close warning: %v", err)
		}
	}()

	searchSvc := searchservice.New(searchClient, searchCfg.AllowedIndexes, searchCfg.EnableMetrics)

	var (
		productIndex string
		postIndex    string
	)
	for _, idx := range searchCfg.AllowedIndexes {
		trimmed := strings.TrimSpace(idx)
		lower := strings.ToLower(trimmed)
		switch {
		case productIndex == "" && (lower == "products" || lower == "product" || strings.Contains(lower, "product")):
			productIndex = trimmed
		case postIndex == "" && (lower == "content" || lower == "posts" || lower == "post" || strings.Contains(lower, "content")):
			postIndex = trimmed
		}
	}
	if productIndex == "" && len(searchCfg.AllowedIndexes) > 0 {
		productIndex = searchCfg.AllowedIndexes[0]
		log.Printf("⚠️  No explicit products index configured; falling back to %s", productIndex)
	}
	if postIndex == "" {
		log.Printf("⚠️  Content index not configured; post events will fail until SEARCH_ALLOWED_INDEXES includes one")
	}

	workerLogger := log.New(os.Stdout, "[search-worker] ", log.LstdFlags)
	workerCfg := searchworker.Config{
		PollInterval:    2 * time.Second,
		MaxBatchSize:    25,
		MaxAttempts:     5,
		ContextTimeout:  5 * time.Second,
		IndexResolver:   searchworker.ResolverFromConfig(productIndex, postIndex),
		DocumentBuilder: searchworker.BuildDocumentFromDB(gormDB),
	}
	searchWorker, err := searchworker.NewWorker(gormDB, searchSvc, workerCfg, workerLogger)
	if err != nil {
		log.Fatalf("❌ Failed to initialize search worker: %v", err)
	}
	workerLogger.Printf("initializing worker (product index=%s, content index=%s, poll=%s, batch=%d)",
		productIndex,
		postIndex,
		workerCfg.PollInterval,
		workerCfg.MaxBatchSize,
	)
	searchWorker.Start()
	defer searchWorker.Stop()
	log.Printf("🔄 Search worker started (product index=%s, content index=%s)", productIndex, postIndex)
	searchHTTPServer := searchhttp.New(searchCfg.HTTPAddr, searchCfg.MetricsAddr, searchCfg.EnableMetrics, searchSvc)
	log.Printf("🔍 Search service listening on %s (metrics %s, indexes=%d)", searchCfg.HTTPAddr, searchCfg.MetricsAddr, len(searchCfg.AllowedIndexes))

	searchCtx, searchCancel := context.WithCancel(context.Background())
	defer searchCancel()
	searchErrCh := make(chan error, 1)

	go func() {
		searchErrCh <- searchHTTPServer.Run(searchCtx)
	}()

	go func() {
		if err := <-searchErrCh; err != nil {
			log.Fatalf("❌ Search service stopped: %v", err)
		}
	}()

	// Create router using the router package (now with settingHandler)
	r := router.SetupRouter(settingHandler, settingService, gormDB, redisClient, adminSettingHandler, paymentHandler, mellatHandler, chatHandler, wsManager, monitoringHandler, searchSvc, searchCfg.AllowedIndexes)

	// چاپ لیست routeها برای عیب‌یابی
	// for _, ri := range r.Routes() {
	// 	fmt.Printf("%s %s\n", ri.Method, ri.Path)
	// }

	// Start server
	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		log.Fatal("❌ Error: BACKEND_PORT environment variable is not set.")
	}

	log.Printf("🚀 Starting server on port %s...", port)

	// Create HTTP server with HTTP/2 support
	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
		// تنظیمات timeout برای افزایش زمان اجرای درخواست‌های طولانی (مانند انتقال محصولات)
		ReadTimeout:       30 * time.Minute,  // افزایش به 30 دقیقه
		WriteTimeout:      30 * time.Minute,  // افزایش به 30 دقیقه
		IdleTimeout:       120 * time.Second, // زمان idle connection
		ReadHeaderTimeout: 30 * time.Second,  // زمان خواندن هدرها
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
			MaxVersion: tls.VersionTLS13,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			},
		},
	}

	// Enable HTTP/2
	if err := http2.ConfigureServer(server, &http2.Server{}); err != nil {
		// Warning: HTTP/2 configuration failed
	}

	// Graceful shutdown setup
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		// Shutting down server...

		// Create a context with timeout for graceful shutdown
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		searchCancel()

		if err := server.Shutdown(ctx); err != nil {
			// Server forced to shutdown
		}

		// Server exited
		os.Exit(0)
	}()

	// For development, run without TLS (HTTP/2 will work over h2c)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		// Explicitly log the startup error so systemd journal shows root cause (e.g., address in use)
		log.Fatalf("❌ Server failed to start on port %s: %v", port, err)
	}
}

func ensureSearchIndexes() error {
	indexDir := strings.TrimSpace(os.Getenv("SEARCH_INDEX_DIR"))
	if indexDir == "" {
		indexDir = "var/search/indexes"
	}
	indexDir = filepath.Clean(indexDir)

	allowed := normalizeIndexList(os.Getenv("SEARCH_ALLOWED_INDEXES"))
	if indexesPresent(indexDir, allowed) {
		return nil
	}

	log.Printf("🔧 search indexes missing at %s; running bootstrap", indexDir)

	dataDir := strings.TrimSpace(os.Getenv("SEARCH_SAMPLE_DATA_DIR"))
	if dataDir == "" {
		dataDir = "data/search"
	}
	dataDir = filepath.Clean(dataDir)

	if err := searchbootstrap.Run(searchbootstrap.Options{
		IndexDir: indexDir,
		DataDir:  dataDir,
		Force:    false,
		Indexes:  allowed,
	}); err != nil {
		return fmt.Errorf("run search bootstrap: %w", err)
	}

	log.Println("✅ search indexes ready")
	return nil
}

func normalizeIndexList(raw string) []string {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil
	}
	parts := strings.Split(raw, ",")
	seen := make(map[string]struct{})
	indexes := make([]string, 0, len(parts))
	for _, part := range parts {
		name := strings.TrimSpace(part)
		if name == "" {
			continue
		}
		lower := strings.ToLower(name)
		if _, ok := seen[lower]; ok {
			continue
		}
		seen[lower] = struct{}{}
		indexes = append(indexes, lower)
	}
	return indexes
}

func indexesPresent(indexDir string, indexes []string) bool {
	info, err := os.Stat(indexDir)
	if err != nil || !info.IsDir() {
		return false
	}

	if len(indexes) == 0 {
		return hasAnyBleve(indexDir)
	}

	for _, idx := range indexes {
		if idx == "" {
			continue
		}
		path := filepath.Join(indexDir, idx+".bleve")
		if fileInfo, statErr := os.Stat(path); statErr != nil || fileInfo.IsDir() {
			return false
		}
	}
	return true
}

func hasAnyBleve(indexDir string) bool {
	entries, err := os.ReadDir(indexDir)
	if err != nil {
		return false
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if strings.HasSuffix(strings.ToLower(entry.Name()), ".bleve") {
			return true
		}
	}
	return false
}
