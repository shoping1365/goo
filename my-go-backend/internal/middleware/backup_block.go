package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// BlockBackupAccess denies any direct HTTP access to files under /media/uploads/media/backup
// Returns 500 on any such request so that even if path is guessed nothing is revealed.
func BlockBackupAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/media/uploads/media/backup") {
			// Return generic server error
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.Next()
	}
}
