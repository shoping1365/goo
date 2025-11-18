package handlers

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// generateUniqueName تولید نام یکتا برای فایل‌ها
func generateUniqueName(ext string) string {
	ts := time.Now().UnixNano() / 1e6
	randPart := rand.Intn(0xffff)
	return fmt.Sprintf("%s-%04x%s", strconv.FormatInt(ts, 36), randPart, ext)
}
