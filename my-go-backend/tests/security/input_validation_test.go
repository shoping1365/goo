package security_test

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"strings"
	"testing"
	"time"

	"my-go-backend/internal/handlers"
	"my-go-backend/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SetupValidationRouter sets up a router with Review and Media handlers for testing
func SetupValidationRouter(db *gorm.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(gin.Recovery())
	
	// Mock Auth Middleware for simplicity in some tests, or use the real one
	// Here we use a simple middleware that extracts user_id from header if present
	// or relies on the real middleware if we want full integration.
	// For validation tests, we often want to bypass complex auth unless auth itself is being tested.
	// However, CreateReview requires user_id in context.
	r.Use(func(c *gin.Context) {
		// Mock auth: if Authorization header is present, parse it or just set a dummy user
		// For this test, we'll assume the test sets the user_id in context manually 
		// or we use the real auth middleware.
		// Let's use a helper middleware for testing that trusts a specific header
		if uid := c.GetHeader("X-Test-User-ID"); uid != "" {
			// fmt.Println("Setting Mock User ID:", uid)
			// Convert to uint
			var id uint
			fmt.Sscanf(uid, "%d", &id)
			c.Set("user_id", id)
		}
		c.Next()
	})

	reviewHandler := handlers.NewReviewHandler(db)
	mediaHandler := &handlers.MediaHandler{DB: db}

	r.POST("/api/reviews", reviewHandler.CreateReview)
	// UploadMediaHandler is http.HandlerFunc, wrap it for Gin
	r.POST("/api/media/upload", gin.WrapF(mediaHandler.UploadMediaHandler))

	return r
}

func TestXSSInReviews(t *testing.T) {
	db := SetupTestDB()
	// Ensure tables exist and have necessary columns
	// Since we are running on a real DB that might have old migrations, we force add the column if missing
	db.Exec("ALTER TABLE reviews ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMPTZ")
	db.AutoMigrate(&models.Customer{}, &models.Review{}, &models.Product{})
	
	r := SetupValidationRouter(db)

	// Create a test user
	userID, _ := CreateTestUser(db, "test_xss_user", "user")

	// Clean up previous reviews for this user to avoid false positives from old runs
	db.Where("customer_id = ?", userID).Delete(&models.Review{})

	// XSS Payload
	xssPayload := "<script>alert('XSS')</script>"
	
	// Prepare form data
	form := map[string]string{
		"rating":     "5",
		"product_id": "1", // Assuming product 1 exists or FK check is skipped/mocked. 
		                   // If FK check exists, we might need to create a product.
		"comment":    xssPayload,
		"title":      "Great Product",
		"pros":       "Good",
		"cons":       "None",
	}

	// We need to create a product if the handler checks for it.
	// ReviewHandler.CreateReview doesn't seem to check product existence explicitly in the snippet I saw,
	// but it might fail on DB constraint. Let's create a dummy product.
	product := models.Product{
		Name: "Test Product", 
		Slug: fmt.Sprintf("test-product-xss-%d", time.Now().UnixNano()),
		UUID: uuid.New().String(),
	}
	db.Create(&product)
	form["product_id"] = fmt.Sprintf("%d", product.ID)

	// Create Request
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for k, v := range form {
		_ = writer.WriteField(k, v)
	}
	_ = writer.Close()

	req, _ := http.NewRequest("POST", "/api/reviews", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("X-Test-User-ID", fmt.Sprintf("%d", userID)) // Mock Auth

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check response
	if w.Code != http.StatusOK && w.Code != http.StatusCreated {
		t.Logf("Review creation failed with status %d: %s", w.Code, w.Body.String())
		// If it failed, maybe validation caught it?
	}

	// Check Database for the comment
	var review models.Review
	err := db.Where("customer_id = ? AND product_id = ?", userID, product.ID).First(&review).Error
	if err != nil {
		t.Fatalf("Failed to find created review: %v", err)
	}

	// Check if XSS payload is stored raw
	if strings.Contains(review.Comment, "<script>") {
		t.Errorf("⚠️  SECURITY WARNING: XSS Payload stored raw in database! Comment: %s", review.Comment)
		t.Log("Recommendation: Use a sanitization library like 'bluemonday' before saving content.")
	} else {
		t.Log("✅ XSS Payload appears to be sanitized or modified.")
	}
}

func TestFileUploadValidation(t *testing.T) {
	db := SetupTestDB()
	r := SetupValidationRouter(db)

	// 1. Test Dangerous Extension (.exe)
	t.Run("DangerousExtension", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		
		// Create a fake exe file
		part, _ := writer.CreateFormFile("file", "malware.exe")
		part.Write([]byte("MZ...")) // Fake EXE header
		writer.Close()

		req, _ := http.NewRequest("POST", "/api/media/upload", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Expecting failure (400 or 403 or 500 depending on handler)
		// If it returns 200, it's a security risk.
		if w.Code == http.StatusOK {
			t.Errorf("⚠️  SECURITY WARNING: Server accepted .exe file upload! Status: %d", w.Code)
		} else {
			t.Logf("✅ Server rejected .exe file with status: %d", w.Code)
		}
	})

	// 2. Test Fake Image (Text file renamed to .jpg)
	t.Run("FakeImage", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		
		// Create a file named image.jpg but with text content
		h := make(textproto.MIMEHeader)
		h.Set("Content-Disposition", `form-data; name="file"; filename="fake_image.jpg"`)
		h.Set("Content-Type", "image/jpeg")
		part, _ := writer.CreatePart(h)
		part.Write([]byte("This is just text, not an image."))
		writer.Close()

		req, _ := http.NewRequest("POST", "/api/media/upload", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// The handler should ideally detect content type mismatch
		if w.Code == http.StatusOK {
			// Check response to see if it actually saved it
			t.Logf("Server accepted fake image. Response: %s", w.Body.String())
			// If the server relies only on extension, this is a weakness, but less critical than .exe if it doesn't execute it.
			// However, robust servers should check magic bytes.
			t.Log("ℹ️  Note: Ensure server checks file magic bytes, not just extension.")
		} else {
			t.Logf("✅ Server rejected fake image with status: %d", w.Code)
		}
	})
}
