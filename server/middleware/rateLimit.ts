/**
 * Rate Limiting Middleware
 * Prevent API abuse and DDoS attacks
 */

import { eventHandler, getRequestHeader, createError } from 'h3'
import { requireAuth, type AuthUser } from '../utils/requireAuth'
import { getCookieValue } from '../api/_utils/cookieHelper'

interface RateLimitConfig {
  maxRequests: number
  windowMs: number
}

interface ClientRequest {
  count: number
  resetTime: number
}

// Store for tracking requests
const clientRequests = new Map<string, ClientRequest>()

// Default rate limit: 100 requests per 15 minutes
const defaultConfig: RateLimitConfig = {
  maxRequests: 100,
  windowMs: 15 * 60 * 1000,
}

// Different limits for different endpoints
const rateLimitConfig: Record<string, RateLimitConfig> = {
  '/api/auth/login': {
    maxRequests: 5,
    windowMs: 15 * 60 * 1000, // 5 requests per 15 minutes
  },
  '/api/auth/register': {
    maxRequests: 3,
    windowMs: 60 * 60 * 1000, // 3 requests per hour
  },
  '/api/auth/otp': {
    maxRequests: 10,
    windowMs: 15 * 60 * 1000, // 10 requests per 15 minutes
  },
  '/api/admin': {
    maxRequests: 200,
    windowMs: 60 * 1000, // 200 requests per minute
  },
}

/**
 * Get client IP address
 */
const getClientIp = (event: any): string => {
  try {
    // Try to get forwarded IP
    const xForwardedFor = getRequestHeader(event, 'x-forwarded-for')
    if (xForwardedFor) {
      const ip = Array.isArray(xForwardedFor) ? xForwardedFor[0] : xForwardedFor
      return ip.split(',')[0].trim()
    }

    const xRealIp = getRequestHeader(event, 'x-real-ip')
    if (xRealIp) {
      return Array.isArray(xRealIp) ? xRealIp[0] : xRealIp
    }
  } catch (e) {
    // Fallback if getRequestHeader fails
  }

  // Fallback to context or node
  return event.context?.clientAddress || event.node?.req?.socket?.remoteAddress || 'unknown'
}

/**
 * Get rate limit config for path
 */
const getConfigForPath = (path: string): RateLimitConfig => {
  for (const [pattern, config] of Object.entries(rateLimitConfig)) {
    if (path.startsWith(pattern)) {
      return config
    }
  }
  return defaultConfig
}

/**
 * Check if request exceeds rate limit
 */
export const checkRateLimit = (clientId: string, config: RateLimitConfig): boolean => {
  const now = Date.now()
  const request = clientRequests.get(clientId)

  if (!request || now > request.resetTime) {
    // Create new window
    clientRequests.set(clientId, {
      count: 1,
      resetTime: now + config.windowMs,
    })
    return true
  }

  if (request.count < config.maxRequests) {
    request.count++
    return true
  }

  return false
}

/**
 * Clear old entries to prevent memory leak
 */
const cleanupOldEntries = () => {
  const now = Date.now()
  for (const [clientId, request] of clientRequests.entries()) {
    if (now > request.resetTime) {
      clientRequests.delete(clientId)
    }
  }
}

// Cleanup every 10 minutes
if (typeof process !== 'undefined' && (process as any).server) {
  setInterval(cleanupOldEntries, 10 * 60 * 1000)
}

/**
 * Check if user is admin/developer
 * Bypass rate limiting for admin and developer users
 * 
 * ساده‌سازی: فقط بررسی می‌کنیم که توکن موجود باشد و از path های admin استفاده شود
 */
const ELEVATED_ROLES = new Set(['admin', 'super_admin', 'superadmin', 'developer', 'devops', 'dev', 'staff', 'manager'])

const hasElevatedRole = (user: AuthUser | undefined | null): boolean => {
  if (!user) return false

  const primaryRole = (user.role || user.effective_role_name || '').toString().toLowerCase()
  if (primaryRole && ELEVATED_ROLES.has(primaryRole)) {
    return true
  }

  const roleList = Array.isArray((user as any).roles) ? (user as any).roles : []
  return roleList.some((roleEntry: any) => {
    if (!roleEntry) return false
    const value = typeof roleEntry === 'string'
      ? roleEntry
      : (roleEntry.name || roleEntry.role || roleEntry.key || '')
    return ELEVATED_ROLES.has(value.toString().toLowerCase())
  })
}

const isAdminUser = async (event: any): Promise<boolean> => {
  const path = event.node.req.url || ''

  try {
    // Admin front-end routes should be unthrottled
    if (path.startsWith('/admin')) {
      return true
    }

    const accessToken = getCookieValue(event, 'access_token') || getCookieValue(event, 'auth-token')
    if (!accessToken) {
      return false
    }

    // مسیرهای مدیریتی/دولوپری را بدون محدودیت نگه داریم
    if (path.includes('/api/auth/me') || path.startsWith('/api/admin')) {
      return true
    }

    const user = await requireAuth(event)
    if (hasElevatedRole(user)) {
      return true
    }

    // اگر کاربر دارای فلگ‌های صریح باشد
    const flags = ['is_admin', 'isAdmin', 'is_developer', 'isDeveloper']
    if (flags.some(flag => Boolean((user as any)?.[flag]))) {
      return true
    }

    return false
  } catch {
    return false
  }
}




export default eventHandler(async (event) => {
  const path = event.path || event.node?.req?.url || ''

  // Skip rate limiting for static files
  if (path.startsWith('/public/') || path.startsWith('/_nuxt/')) {
    return
  }

  // Skip rate limiting for health check
  if (path === '/health') {
    return
  }

  // Skip rate limiting for admin users
  const isAdmin = await isAdminUser(event)
  if (isAdmin) {
    return
  }

  // Get client IP
  const clientIp = getClientIp(event)
  const clientId = `${clientIp}:${path}`

  // Get config for this path
  const config = getConfigForPath(path)

  // Check rate limit
  if (!checkRateLimit(clientId, config)) {
    throw createError({
      statusCode: 429,
      message: 'Too Many Requests',
      data: {
        userMessage: 'شما بیش از حد درخواست ارسال کردید. لطفاً بعداً دوباره تلاش کنید.',
        retryAfter: config.windowMs / 1000,
      },
    })
  }
})
