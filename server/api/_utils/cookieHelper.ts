// h3 v2 compatible cookie helper
import type { H3Event } from 'h3'

// Helper to parse cookies from header string
function parseCookieHeader(header?: string | null): Map<string, string> {
  const map = new Map<string, string>()
  if (!header) return map

  const parts = header.split(';')
  for (const part of parts) {
    const [rawName, ...rawValueParts] = part.split('=')
    if (!rawName) continue
    const name = rawName.trim()
    if (!name) continue
    const rawValue = rawValueParts.join('=').trim()
    if (!rawValue) {
      map.set(name, '')
      continue
    }
    try {
      map.set(name, decodeURIComponent(rawValue))
    } catch {
      map.set(name, rawValue)
    }
  }

  return map
}

/**
 * Get cookie value - h3 v2 compatible
 * Uses event.req.headers.get() or fallback to manual access
 */
export function getCookieValue(event: H3Event, name: string): string | undefined {
  try {
    // Try h3 v2 way: event.req.headers (Headers object)
    const reqHeaders = (event as unknown as { req?: { headers?: Headers | Record<string, string | string[]> } }).req?.headers || event.node?.req?.headers
    
    if (!reqHeaders) return undefined
    
    let cookieHeader: string | undefined
    
    // Try Headers API (h3 v2)
    if (typeof reqHeaders.get === 'function') {
      cookieHeader = reqHeaders.get('cookie') || undefined
    } else {
      // Fallback: plain object (legacy)
      const cookie = reqHeaders.cookie || reqHeaders.Cookie
      cookieHeader = cookie ? (Array.isArray(cookie) ? cookie.join('; ') : cookie) : undefined
    }
    
    if (cookieHeader) {
      return parseCookieHeader(cookieHeader).get(name)
    }
  } catch (_e) {
    // Ignore errors
  }
  
  return undefined
}
