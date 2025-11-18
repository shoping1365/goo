// ابزار هماهنگ برای ارتباط با Go backend – نسخه جدید مبتنی بر fetchGo
import type { H3Event } from 'h3'
import { fetchGo } from '../api/_utils/fetchGo'

export interface GoApiOptions {
  method?: 'GET' | 'POST' | 'PUT' | 'DELETE' | 'PATCH'
  body?: unknown
  headers?: Record<string, string>
  timeout?: number
  baseUrl?: string // Optional base URL override
}

/**
 * ارسال درخواست به Go Backend با forward کردن authentication cookies
 */
export async function goApiFetch<T = unknown>(
  event: H3Event,
  path: string,
  options: GoApiOptions = {}
): Promise<T> {
  const method = options.method || 'GET'
  try {
    const data = await fetchGo(event, path, {
      method,
      body: options.body,
      headers: options.headers,
      timeout: options.timeout,
      baseUrl: options.baseUrl
    })
    return data as T
  } catch (error: unknown) {
    const errorObj = error as { message?: string; statusCode?: number; status?: number }
    console.error('[goApiFetch] Upstream error', { path, method, message: errorObj?.message, status: errorObj?.statusCode || errorObj?.status })
    throw error
  }
}

/**
 * ارسال درخواست به Go Backend و forward کردن Set-Cookie headers
 */
export async function goApiFetchWithCookies<T = unknown>(
  event: H3Event,
  path: string,
  options: GoApiOptions = {}
): Promise<{ data: T; headers: Headers }> {
  const method = options.method || 'GET'
  try {
    const data = await fetchGo(event, path, {
      method,
      body: options.body,
      headers: options.headers,
      timeout: options.timeout,
      baseUrl: options.baseUrl
    })

    // fetchGo خودش Set-Cookie ها را روی پاسخ رو به جلو تنظیم می‌کند
    return {
      data: data as T,
      headers: new Headers()
    }
  } catch (error: unknown) {
    const errorObj = error as { message?: string; statusCode?: number; status?: number }
    console.error('[goApiFetchWithCookies] Upstream error', { path, method, message: errorObj?.message, status: errorObj?.statusCode || errorObj?.status })
    throw error
  }
}