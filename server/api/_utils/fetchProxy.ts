import { appendHeader, createError, H3Event } from 'h3'

function parseCookieHeader(header?: string | null) {
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

const getCookieCompat = (event: H3Event, name: string) => {
  try {
    // Try h3 v2 way: event.req.headers (Headers object)
    const reqHeaders = (event as unknown as Record<string, any>).req?.headers || event.node?.req?.headers
    
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

export async function proxy(
  event: H3Event,
  url: string,
  opts: RequestInit & { retry?: number; timeout?: number } = {}
) {
  // درخواست اصلی از Nuxt شامل body (stream) و رأس‌های اصلی است؛ آن‌ها را به backend فوروارد می‌کنیم
  const nodeReq = event.node?.req || (event.req as unknown as Record<string, any>)
  if (!nodeReq || !nodeReq.headers) {
    throw createError({ statusCode: 500, message: 'Request headers are unavailable for proxying' })
  }

  // تنظیمات پیش‌فرض: همان متد، همان body، و کپی هدرها (به جز host و content-length که توسط fetch تنظیم می‌شود)
  const { host: _host, "content-length": _cl, ...forwardHeaders } = nodeReq.headers as Record<string, string>

  // انتقال کل هدر Cookie ورودی (برای دسترسی به refresh_token و سایر کوکی‌ها)
  if (nodeReq.headers.cookie) {
    forwardHeaders['Cookie'] = nodeReq.headers.cookie as string
  }

  // هدرها را با هدرهای سفارشی ادغام می‌کنیم تا کوکی/IP/UA حفظ شوند و بتوانیم هدرهای اضافه مانند X-User را بیافزاییم
  const { headers: overrideHeaders, method: overrideMethod, body: providedBody, timeout, ...restOpts } = opts
  const customHeaders = (overrideHeaders || {}) as Record<string, string>
  const mergedHeaders: Record<string, string> = { ...forwardHeaders, ...customHeaders }

  // اگر فراخواننده Authorization را ننوشته بود، از access_token موجود در کوکی بسازیم
  if (!('authorization' in mergedHeaders) && !('Authorization' in mergedHeaders)) {
    const authToken = getCookieCompat(event, 'access_token') || getCookieCompat(event, 'auth-token')
    if (authToken) {
      mergedHeaders['Authorization'] = `Bearer ${authToken}`
    }
  }

  // ممنوعیت هدرهای سفارشی (X-User-ID, X-User-Role) طبق سیاست امنیتی


  const method = (overrideMethod || nodeReq.method || 'GET').toUpperCase()
  const requestInit: RequestInit & { timeout?: number } = {
    ...restOpts,
    method,
    headers: mergedHeaders,
    timeout: timeout || 60000
  }

  // اگر opts.body ارائه شده، از آن استفاده کن، در غیر این صورت از nodeReq
  if (providedBody) {
    requestInit.body = providedBody
  } else if (method !== 'GET' && method !== 'HEAD') {
    requestInit.body = nodeReq as unknown as BodyInit
    // Required for streaming bodies with undici/fetch in Node 18+
    // @ts-ignore
    requestInit.duplex = 'half'
  }

  try {
    // Validate URL before making request
    try {
      new URL(url)
    } catch (urlError: unknown) {
      const err = urlError as Record<string, any>
      // console.error('[fetchProxy] Invalid URL:', { url, error: err?.message })
      throw createError({
        statusCode: 500,
        message: `Invalid URL: ${err?.message || 'Unknown error'}`,
        data: { url, error: err?.message }
      })
    }
    
    // @ts-ignore - $fetch provided by Nitro runtime
    const res = await $fetch.raw(url, requestInit)

    return processResponse(event, res)
  } catch (error: unknown) {
    const err = error as Record<string, any>
    // اگر خطا از createError است، مستقیماً throw کن
    if (err?.statusCode) {
      throw error
    }
    
    const msg = String(err?.message || '')
    const streamLike = msg.includes('premature close') || msg.includes('Cannot pipe to a closed') || msg.includes('write after end') || msg.includes('destroyed stream')
    if (streamLike) {
      // تبدیل به خطای بی‌خطر 499 (Client Closed Request)
      throw createError({ statusCode: 499, message: 'Client closed request' })
    }
    // console.error('❌ fetchProxy error:', error)
    // console.error('URL:', url)
    // try { console.error('Options:', JSON.stringify({ ...requestInit, body: undefined }, null, 2)) } catch {}
    throw error
  }
}

function processResponse(event: H3Event, res: unknown) {
  // Forward Set-Cookie headers from upstream so browser stores auth cookies
  try {
    const anyHeaders = (res as Record<string, any>).headers
    const setCookies: string[] = typeof anyHeaders?.getSetCookie === 'function'
      ? anyHeaders.getSetCookie()
      : ([] as string[])
    if (Array.isArray(setCookies) && setCookies.length > 0) {
      for (const sc of setCookies) appendHeader(event, 'set-cookie', sc)
    } else {
      const single = anyHeaders.get('set-cookie')
      if (single) appendHeader(event, 'set-cookie', single)
    }
  } catch { }

  const data = (res as Record<string, any>)._data
  const response = res as Response
  if (!response.ok) {
    const msg = data?.message || data?.error || response.statusText || 'خطا'
    throw createError({
      statusCode: response.status,
      message: msg,
      data,
    })
  }
  return data
}
