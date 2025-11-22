// Rate Limiting برای جلوگیری از حملات Brute Force
import type { H3Event } from 'h3'
import { getRequestHeader } from 'h3'

interface RateLimitEntry {
  count: number
  resetAt: number
  blockedUntil?: number
}

// ذخیره در حافظه (برای production باید از Redis استفاده کرد)
const rateLimitStore = new Map<string, RateLimitEntry>()

export interface RateLimitOptions {
  maxAttempts: number      // تعداد تلاش مجاز
  windowMs: number         // بازه زمانی به میلی‌ثانیه
  blockDurationMs?: number // مدت زمان مسدود شدن (اختیاری)
  keyPrefix?: string       // پیشوند کلید
}

/**
 * بررسی rate limit برای یک کلید خاص
 * @returns true اگر درخواست مجاز باشد، false اگر محدود شده باشد
 */
export function checkRateLimit(
  event: H3Event,
  key: string,
  options: RateLimitOptions
): { allowed: boolean; remaining: number; resetAt: number } {
  // اگر امنیت موقتاً غیرفعال باشد، rate limit هم غیرفعال می‌شود
  const cfg = useRuntimeConfig(event) as { public?: { disableSecurity?: boolean } }
  if (cfg?.public?.disableSecurity) {
    return { allowed: true, remaining: options.maxAttempts, resetAt: Date.now() + options.windowMs }
  }

  const now = Date.now()
  const fullKey = options.keyPrefix ? `${options.keyPrefix}:${key}` : key
  
  // پاکسازی ورودی‌های منقضی شده
  cleanupExpiredEntries()
  
  let entry = rateLimitStore.get(fullKey)
  
  // بررسی مسدود بودن
  if (entry?.blockedUntil && entry.blockedUntil > now) {
    return {
      allowed: false,
      remaining: 0,
      resetAt: entry.blockedUntil
    }
  }
  
  // ایجاد یا ریست entry
  if (!entry || entry.resetAt <= now) {
    entry = {
      count: 0,
      resetAt: now + options.windowMs
    }
  }
  
  entry.count++
  
  // بررسی محدودیت
  if (entry.count > options.maxAttempts) {
    if (options.blockDurationMs) {
      entry.blockedUntil = now + options.blockDurationMs
    }
    
    rateLimitStore.set(fullKey, entry)
    
    return {
      allowed: false,
      remaining: 0,
      resetAt: entry.blockedUntil || entry.resetAt
    }
  }
  
  rateLimitStore.set(fullKey, entry)
  
  return {
    allowed: true,
    remaining: options.maxAttempts - entry.count,
    resetAt: entry.resetAt
  }
}

/**
 * دریافت IP واقعی کاربر
 */
export function getClientIp(event: H3Event): string {
  const xForwardedFor = getRequestHeader(event, 'x-forwarded-for')
  const xRealIp = getRequestHeader(event, 'x-real-ip')
  const cfConnectingIp = getRequestHeader(event, 'cf-connecting-ip')
  const remote = event.node?.req?.socket?.remoteAddress || ''

  // Helper: normalize IPv6 loopback and IPv4-mapped
  const normalize = (ip?: string | null): string | null => {
    if (!ip) return null
    let v = ip
    // اگر چند IP پشت سر هم هستند، اولین را بگیر
    if (v.includes(',')) v = v.split(',')[0].trim()
    // IPv4-mapped IPv6 ::ffff:127.0.0.1
    if (v.startsWith('::ffff:')) v = v.replace('::ffff:', '')
    // Loopback IPv6
    if (v === '::1' || v === '0:0:0:0:0:0:0:1') v = '127.0.0.1'
    // حذف براکت و پورت احتمالی [::1]:1234
    v = v.replace(/^\[/, '').replace(/\]$/, '')
    v = v.split('%')[0] // zone index
    return v
  }

  // اولویت‌ها
  const ip = normalize(cfConnectingIp) || normalize(xForwardedFor) || normalize(xRealIp) || normalize(remote)
  return ip || '127.0.0.1'
}

/**
 * پاکسازی ورودی‌های منقضی شده از حافظه
 */
function cleanupExpiredEntries() {
  const now = Date.now()
  
  for (const [key, entry] of rateLimitStore.entries()) {
    if (entry.resetAt <= now && (!entry.blockedUntil || entry.blockedUntil <= now)) {
      rateLimitStore.delete(key)
    }
  }
}

/**
 * ریست کردن rate limit برای یک کلید
 */
export function resetRateLimit(key: string, keyPrefix?: string) {
  const fullKey = keyPrefix ? `${keyPrefix}:${key}` : key
  rateLimitStore.delete(fullKey)
}

// پاکسازی خودکار هر 5 دقیقه
if (process.server) {
  setInterval(cleanupExpiredEntries, 5 * 60 * 1000)
}