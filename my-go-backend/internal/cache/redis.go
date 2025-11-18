package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

/*
پکیج کش

این پکیج برای مدیریت کش Redis و ذخیره‌سازی موقت داده‌ها استفاده می‌شود.
*/

var (
	// کلاینت Redis
	client *redis.Client

	// کانتکست پیش‌فرض
	ctx = context.Background()
)

// Initialize کلاینت Redis را مقداردهی اولیه می‌کند
func Initialize(addr, password string, db int) error {
	client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	// تست اتصال
	_, err := client.Ping(ctx).Result()
	return err
}

// Set یک مقدار را در کش ذخیره می‌کند
func Set(key string, value interface{}, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return client.Set(ctx, key, data, expiration).Err()
}

// Get یک مقدار را از کش بازیابی می‌کند
func Get(key string, dest interface{}) error {
	data, err := client.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}

	return json.Unmarshal(data, dest)
}

// Delete یک مقدار را از کش حذف می‌کند
func Delete(key string) error {
	return client.Del(ctx, key).Err()
}

// Clear تمام مقادیر کش را پاک می‌کند
func Clear() error {
	return client.FlushDB(ctx).Err()
}

// Close اتصال به Redis را می‌بندد
func Close() error {
	return client.Close()
}
