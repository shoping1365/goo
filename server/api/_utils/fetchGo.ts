import { H3Event, createError } from 'h3'
import { getGoApiBaseUrl } from '../../utils/api-config'

// Cache the base URL to avoid repeated config lookups
let cachedBaseUrl: string | undefined

function resolveGoApiBase(): string {
  if (!cachedBaseUrl) {
    const base = getGoApiBaseUrl()
    if (!base) {
      throw new Error('GO API base URL is not configured')
    }
    cachedBaseUrl = base
  }
  return cachedBaseUrl
}

// Allow JSON-like bodies in addition to standard RequestInit body types
type FetchGoOptions = Omit<RequestInit, 'body'> & {
  body?: any
  retry?: number
  retryDelayMs?: number
  timeout?: number
  baseUrl?: string // Optional base URL override
}

const transientNetworkCodes = new Set([
  'ECONNRESET',
  'ECONNREFUSED',
  'ECONNABORTED',
  'EPIPE',
  'UND_ERR_CONNECT_TIMEOUT',
  'UND_ERR_HEADERS_TIMEOUT',
  'UND_ERR_ABORTED',
  'UND_ERR_BODY_TIMEOUT',
  'ETIMEDOUT'
])

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

function resolveAuthToken(cookieMap: Map<string, string>): string | undefined {
  const preferredKeys = [
    'admin_token',
    'access_token',
    'auth-token',
    'auth_token',
    'token',
    'session_token'
  ]

  for (const key of preferredKeys) {
    const value = cookieMap.get(key)
    if (value) return value
  }
  return undefined
}

export async function fetchGo(event: H3Event, path: string, opts: FetchGoOptions = {}) {
  let url: string

  // اگر path یک URL کامل است (شروع با http:// یا https://)، مستقیماً از آن استفاده کن
  if (path.startsWith('http://') || path.startsWith('https://')) {
    try {
      // Validate the URL
      const urlObj = new URL(path)
      url = urlObj.toString()
    } catch (urlError: any) {
      console.error('[fetchGo] Invalid absolute URL:', { path, error: urlError?.message })
      throw createError({
        statusCode: 500,
        message: `Invalid URL: ${urlError?.message || 'Unknown error'}`,
        data: { path, error: urlError?.message }
      })
    }
  } else {
    // اگر path نسبی است، از base URL استفاده کن
    const base = opts.baseUrl || resolveGoApiBase()

    if (!base || typeof base !== 'string' || base.trim().length === 0) {
      console.error('[fetchGo] Backend base URL is not configured', { base, type: typeof base })
      throw createError({ statusCode: 500, message: 'Backend base URL is not configured' })
    }

    try {
      // Normalize base and path
      const normalizedBase = base.replace(/\/+$/, '')
      const normalizedPath = path.startsWith('/') ? path : `/${path}`

      // Construct URL
      url = `${normalizedBase}${normalizedPath}`

      // Validate the constructed URL
      const urlObj = new URL(url)
      url = urlObj.toString()
    } catch (urlError: any) {
      console.error('[fetchGo] Failed to construct URL', {
        base,
        path,
        error: urlError?.message
      })
      throw createError({
        statusCode: 500,
        message: `Invalid upstream URL: ${urlError?.message || 'Unknown error'}`,
        data: { base, path, error: urlError?.message }
      })
    }
  }

  console.log('[fetchGo] Making request to:', url)

  const maxRetry = typeof opts.retry === 'number' ? opts.retry : 3
  const retryDelay = typeof opts.retryDelayMs === 'number' ? opts.retryDelayMs : 1000

  // forward cookies and standard auth header for session continuity
  const headers: Record<string, string> = {}

  // Get headers using h3 v2 API - event.req.headers.get() or fallback
  try {
    // Try h3 v2 way: event.req.headers (Headers object)
    const reqHeaders = (event as any).req?.headers || event.node?.req?.headers

    if (reqHeaders) {
      // Try Headers API (h3 v2)
      if (typeof reqHeaders.get === 'function') {
        const cookieHeader = reqHeaders.get('cookie')
        if (cookieHeader) {
          headers.cookie = cookieHeader
          const cookieMap = parseCookieHeader(cookieHeader)
          const authToken = resolveAuthToken(cookieMap)
          if (authToken) {
            headers.authorization = `Bearer ${authToken}`
          }
        }

        const authorization = reqHeaders.get('authorization')
        if (authorization && !headers.authorization) {
          headers.authorization = authorization
        }
      } else {
        // Fallback: plain object (legacy)
        const cookieHeader = reqHeaders.cookie || reqHeaders.Cookie
        if (cookieHeader) {
          headers.cookie = Array.isArray(cookieHeader) ? cookieHeader.join('; ') : cookieHeader
          const cookieMap = parseCookieHeader(headers.cookie)
          const authToken = resolveAuthToken(cookieMap)
          if (authToken) {
            headers.authorization = `Bearer ${authToken}`
          }
        }

        const authHeader = reqHeaders.authorization || reqHeaders.Authorization
        if (authHeader && !headers.authorization) {
          headers.authorization = Array.isArray(authHeader) ? authHeader.join(' ') : authHeader
        }
      }
    }
  } catch (headerError) {
    console.warn('[fetchGo] Failed to get headers:', headerError)
  }

  // Merge with custom headers
  if (opts.headers) {
    Object.assign(headers, opts.headers)
  }

  let lastError: any = null
  for (let attempt = 0; attempt <= maxRetry; attempt++) {
    try {
      console.log('[fetchGo] Attempting request:', { url, method: opts.method || 'GET', headers })

      // Use native fetch instead of $fetch for better h3 v2 compatibility
      const fetchOptions: RequestInit = {
        method: opts.method || 'GET',
        headers,
      }

      if (opts.body) {
        if (typeof opts.body === 'string') {
          fetchOptions.body = opts.body
        } else {
          fetchOptions.body = JSON.stringify(opts.body)
          if (!headers['content-type'] && !headers['Content-Type']) {
            (fetchOptions.headers as any)['content-type'] = 'application/json'
          }
        }
      }

      const timeout = typeof opts.timeout === 'number' ? opts.timeout : 60000
      const controller = new AbortController()
      const timeoutId = setTimeout(() => controller.abort(), timeout)
      fetchOptions.signal = controller.signal

      const res = await fetch(url, fetchOptions)
      clearTimeout(timeoutId)

      console.log('[fetchGo] Response received:', { status: res.status, statusText: res.statusText })

      // Forward Set-Cookie headers from Go backend to client
      const setCookieHeader = res.headers.get('set-cookie')
      if (setCookieHeader && event.node?.res) {
        const nodeRes = event.node.res as any
        if (typeof nodeRes.appendHeader === 'function') {
          nodeRes.appendHeader('set-cookie', setCookieHeader)
        } else {
          const existing = nodeRes.getHeader('set-cookie')
          if (!existing) {
            nodeRes.setHeader('set-cookie', setCookieHeader)
          } else {
            const updated = Array.isArray(existing)
              ? [...existing, setCookieHeader]
              : [existing, setCookieHeader]
            nodeRes.setHeader('set-cookie', updated)
          }
        }
      }

      // Check if response is ok
      if (!res.ok) {
        const errorText = await res.text()
        console.error('[fetchGo] Response error:', { status: res.status, statusText: res.statusText, body: errorText })
        throw createError({
          statusCode: res.status,
          message: res.statusText || 'Upstream error',
          data: errorText
        })
      }

      // Parse response
      const contentType = res.headers.get('content-type')
      let response: any
      if (contentType?.includes('application/json')) {
        response = await res.json()
      } else {
        response = await res.text()
      }

      console.log('[fetchGo] Success:', { url, responseType: typeof response })
      return response
    } catch (err: any) {
      lastError = err
      const code = (err?.code || err?.statusCode || '').toString()
      const causeCode = (err?.cause?.code || '').toString()
      const message: string = (typeof err === 'string' ? err : (err?.message || err?.statusMessage || ''))
      const isAbortName = err?.name === 'AbortError' || err?.cause?.name === 'AbortError'
      const isTransient =
        transientNetworkCodes.has(code) ||
        transientNetworkCodes.has(causeCode) ||
        isAbortName ||
        /ECONNABORTED|socket hang up|aborted|ECONNREFUSED/i.test(message)
      if (attempt < maxRetry && isTransient) {
        // exponential backoff: 1s, 2s, 4s
        const delay = retryDelay * Math.pow(2, attempt);
        console.warn(`[fetchGo] Attempt ${attempt + 1}/${maxRetry + 1} failed. Retrying in ${delay}ms...`, { code, message });
        await new Promise(res => setTimeout(res, delay));
        continue;
      }
      console.error('[fetchGo] Upstream request failed', {
        url,
        method: opts.method || 'GET',
        status: err?.status ?? err?.statusCode,
        code: err?.code,
        message: err?.message,
        responseStatus: err?.response?.status,
        responseBody: err?.response?._data,
        cause: {
          code: err?.cause?.code,
          message: err?.cause?.message
        }
      })
      // throw normalized error for Nitro
      throw createError({
        statusCode: err?.statusCode || err?.status || err?.response?.status || 502,
        message: err?.message || err?.statusMessage || 'Upstream error',
        data: {
          code: err?.code,
          message: err?.message,
          status: err?.status ?? err?.response?.status,
          responseBody: err?.response?._data,
          cause: err?.cause?.code || err?.cause?.message ? {
            code: err?.cause?.code,
            message: err?.cause?.message
          } : undefined
        }
      })
    }
  }
  throw createError({ statusCode: 502, message: lastError?.message || 'Upstream error' })
}