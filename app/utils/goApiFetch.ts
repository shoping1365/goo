/**
 * Client-side wrapper for calling Go backend API from browser
 * This utility uses native fetch with proper authentication
 * 
 * Usage:
 * const data = await goApiFetch<ResponseType>('/api/endpoint', { method: 'POST', body: {...} })
 */

type FetchMethod = 'GET' | 'POST' | 'PUT' | 'DELETE' | 'PATCH'

interface GoApiFetchOptions {
  method?: FetchMethod
  body?: any
  headers?: Record<string, string>
  query?: Record<string, any>
}

/**
 * Get the Go API base URL from runtime config or environment
 */
function getGoApiBaseUrl(): string {
  // Try to get from window.__NUXT__ if available (client-side)
  if (typeof window !== 'undefined' && (window as any).__NUXT__?.config?.public?.goApiBase) {
    return (window as any).__NUXT__.config.public.goApiBase
  }
  
  // Fallback to environment variable or default
  return process.env.NUXT_PUBLIC_GO_API_BASE || 'http://localhost:8080'
}

/**
 * Make authenticated requests to Go backend API
 * Uses native fetch with cookies
 */
export async function goApiFetch<T = any>(
  endpoint: string,
  options: GoApiFetchOptions = {}
): Promise<T> {
  const baseURL = getGoApiBaseUrl()
  
  if (!baseURL) {
    throw new Error('Go API base URL is not configured')
  }

  // Build full URL
  let url = endpoint.startsWith('http') ? endpoint : `${baseURL}${endpoint}`
  
  // Add query parameters if provided
  if (options.query) {
    const params = new URLSearchParams()
    Object.entries(options.query).forEach(([key, value]) => {
      if (value !== undefined && value !== null) {
        params.append(key, String(value))
      }
    })
    const queryString = params.toString()
    if (queryString) {
      url += (url.includes('?') ? '&' : '?') + queryString
    }
  }

  // Prepare headers
  const headers: HeadersInit = {
    'Content-Type': 'application/json',
    ...options.headers
  }

  // Prepare request options
  const requestOptions: RequestInit = {
    method: options.method || 'GET',
    headers,
    credentials: 'include', // Important: include cookies
  }

  // Add body if present
  if (options.body) {
    if (typeof options.body === 'object') {
      requestOptions.body = JSON.stringify(options.body)
    } else {
      requestOptions.body = options.body
    }
  }

  try {
    const response = await fetch(url, requestOptions)

    if (!response.ok) {
      const errorData = await response.json().catch(() => ({}))
      
      // Handle 401 Unauthorized
      if (response.status === 401) {
        // Redirect to login on client-side
        if (typeof window !== 'undefined') {
          window.location.href = '/login'
        }
        throw new Error('Unauthorized')
      }
      
      throw new Error(
        errorData.message || 
        errorData.error || 
        `HTTP ${response.status}: ${response.statusText}`
      )
    }

    // Return JSON response
    return await response.json()
  } catch (error: any) {
    console.error('goApiFetch error:', error)
    throw error
  }
}
