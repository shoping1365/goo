package utils

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

// این فایل مسئول تشخیص خودکار نوع مقدار ورودی (عددی/بولی/تاریخ/متن) است
// و مقادیر پارس‌شده را برای ذخیره در ستون‌های بهینه دیتابیس برمی‌گرداند.

var (
	numberRe = regexp.MustCompile(`^[-+]?\d*(?:[\.,]?\d+)?$`)
)

// DetectScalarType نوع مقدار ورودی را تشخیص می‌دهد و در صورت امکان مقدار پارس‌شده را نیز برمی‌گرداند.
// خروجی typ یکی از مقادیر "number" | "boolean" | "date" | "text" است.
// اگر قابل پارس باشد، یکی از num/b/bt مقداردهی می‌شود.
func DetectScalarType(input string) (typ string, num *float64, b *bool, dt *time.Time) {
	s := strings.TrimSpace(input)
	s = strings.ReplaceAll(s, ",", "") // حذف جداکننده‌های هزارگان معمول
	lower := strings.ToLower(s)

	// 1) بولی (فارسی/انگلیسی)
	switch lower {
	case "true", "false", "yes", "no", "on", "off", "بله", "خیر", "درست", "نادرست":
		val := lower == "true" || lower == "yes" || lower == "on" || lower == "بله" || lower == "درست"
		typ = "boolean"
		b = &val
		return
	}

	// 2) عددی
	if numberRe.MatchString(s) && s != "" {
		// جایگزینی نقطهٔ اعشار
		normalized := strings.ReplaceAll(s, ",", ".")
		if n, err := strconv.ParseFloat(normalized, 64); err == nil {
			typ = "number"
			num = &n
			return
		}
	}

	// 3) تاریخ/زمان – چند الگو رایج میلادی
	layouts := []string{
		time.RFC3339, // 2006-01-02T15:04:05Z07:00
		"2006-01-02 15:04:05",
		"2006-01-02 15:04",
		"2006-01-02",
		"02-01-2006",
		"02/01/2006",
		"2006/01/02",
	}
	for _, layout := range layouts {
		if t, err := time.Parse(layout, s); err == nil {
			typ = "date"
			dt = &t
			return
		}
	}

	// 4) پیش‌فرض: متن
	typ = "text"
	return
}

// تابع کمکی اضافه نیاز نیست؛ از strconv.ParseFloat استفاده شد
