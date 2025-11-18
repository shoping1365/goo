package services

import (
	"context"
	"time"
)

// Cache defines the interface for cache operations
type Cache interface {
	// Get retrieves a value from cache
	Get(ctx context.Context, key string) (string, error)

	// Set stores a value in cache with expiration
	Set(ctx context.Context, key string, value string, expiration time.Duration) error

	// Delete removes a value from cache
	Delete(ctx context.Context, key string) error

	// Clear removes all values from cache
	Clear(ctx context.Context) error
}
