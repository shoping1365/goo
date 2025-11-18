package middleware

import (
	"net/http"
)

// NOTE: JWT Middleware is already implemented in auth.go
// This file is kept for reference only
//
// The following middleware functions are available:
// - Auth() - Validates JWT AccessToken from Authorization header or cookie
// - OptionalAuth() - Validates JWT if provided (optional)
// - AdminMiddleware() - Checks if user has admin role
// - RequirePermission(permission) - Checks if user has specific permission
// - GetUserID(c) - Helper to get user_id from context
// - GetMobile(c) - Helper to get mobile from context
//
// Usage:
//   authRequired := r.Group("/api/protected")
//   authRequired.Use(middleware.Auth())
//   {
//     authRequired.POST("/logout", handler.Logout)
//   }
//
// For admin routes:
//   admin := r.Group("/api/admin")
//   admin.Use(middleware.Auth(), middleware.AdminMiddleware())
//   {
//     admin.GET("/users", handler.GetUsers)
//   }

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement JWT validation
		next.ServeHTTP(w, r)
	})
}
