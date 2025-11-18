package services

import (
	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/models"
)

// این فایل وظیفه مدیریت صف و درج انبوه لاگ‌های ترافیک را بر عهده دارد

// [TEMP DISABLED] – no channel/once to avoid unused warnings

// StartTrafficLogWorker راه‌اندازی صف و goroutine برای درج batch در دیتابیس
func StartTrafficLogWorker(uowFactory unitofwork.UnitOfWorkFactory) {
	// [TEMP DISABLED]
}

// EnqueueTrafficLog افزودن رکورد ترافیک به صف درج انبوه
func EnqueueTrafficLog(log models.TrafficLog) {
	// [TEMP DISABLED]
}
