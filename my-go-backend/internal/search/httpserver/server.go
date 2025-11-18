package httpserver

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"my-go-backend/internal/search/models"
	"my-go-backend/internal/search/service"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Server لایه HTTP را مدیریت می‌کند و درخواست‌ها را به سرویس می‌سپارد.
type Server struct {
	addr          string
	metricsAddr   string
	svc           *service.Service
	enableMetrics bool
	httpServer    *http.Server
	metricsServer *http.Server
}

// New ایجاد سرور جدید با gin و routeهای لازم.
func New(addr, metricsAddr string, enableMetrics bool, svc *service.Service) *Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(requestLogger())

	s := &Server{
		addr:          addr,
		metricsAddr:   metricsAddr,
		svc:           svc,
		enableMetrics: enableMetrics,
	}

	router.POST("/v1/search", s.handleSearch)
	router.POST("/v1/suggest", s.handleSuggest)
	router.GET("/healthz", s.handleHealth)
	router.GET("/ready", s.handleHealth)
	router.GET("/live", s.handleHealth)

	s.httpServer = &http.Server{
		Addr:              addr,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	if enableMetrics {
		mux := http.NewServeMux()
		mux.Handle("/metrics", promhttp.Handler())
		s.metricsServer = &http.Server{
			Addr:         metricsAddr,
			Handler:      mux,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
		}
	}

	return s
}

// Run سرور را اجرا می‌کند و تا دریافت سیگنال توقف زنده نگه می‌دارد.
func (s *Server) Run(ctx context.Context) error {
	errCh := make(chan error, 2)

	go func() {
		log.Printf("search http server listening on %s", s.addr)
		err := s.httpServer.ListenAndServe()
		errCh <- err
	}()

	if s.metricsServer != nil {
		go func() {
			log.Printf("search metrics server listening on %s", s.metricsAddr)
			err := s.metricsServer.ListenAndServe()
			errCh <- err
		}()
	}

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_ = s.httpServer.Shutdown(shutdownCtx)
		if s.metricsServer != nil {
			_ = s.metricsServer.Shutdown(shutdownCtx)
		}
		return nil
	case err := <-errCh:
		if err == nil || errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
}

func (s *Server) handleSearch(c *gin.Context) {
	var req models.SearchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid payload", "details": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	resp, err := s.svc.Search(ctx, req)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, service.ErrEmptyQuery()) || errors.Is(err, service.ErrInvalidLimit()) {
			status = http.StatusBadRequest
		}
		c.JSON(status, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": resp})
}

func (s *Server) handleSuggest(c *gin.Context) {
	var req models.SuggestRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid payload", "details": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	resp, err := s.svc.Suggest(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": resp})
}

func (s *Server) handleHealth(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	resp, err := s.svc.Health(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": resp.Status, "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func requestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		status := c.Writer.Status()
		log.Printf("%s %s -> %d (%s)", c.Request.Method, c.FullPath(), status, duration)
	}
}
