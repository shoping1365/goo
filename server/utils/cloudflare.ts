// ابزارهای مشترک ارتباط با Cloudflare API
// تمام توضیحات به زبان فارسی نوشته شده است

import { ofetch } from 'ofetch'

/**
 * ساخت هدرهای احراز هویت برای Cloudflare
 * از متغیر محیطی CLOUDFLARE_API_TOKEN استفاده می‌کند
 */
function buildAuthHeaders (): Record<string, string> {
  const apiToken = process.env.CLOUDFLARE_API_TOKEN
  if (!apiToken) {
    throw new Error('متغیر محیطی CLOUDFLARE_API_TOKEN تنظیم نشده است')
  }
  return {
    'Authorization': `Bearer ${apiToken}`,
    'Content-Type': 'application/json'
  }
}

/**
 * دریافت Zone ID از متغیر محیطی
 */
export function getCloudflareZoneId (): string {
  const zoneId = process.env.CLOUDFLARE_ZONE_ID
  if (!zoneId) {
    throw new Error('متغیر محیطی CLOUDFLARE_ZONE_ID تنظیم نشده است')
  }
  return zoneId
}

/**
 * فراخوانی Cloudflare API
 * @param endpoint مسیر نسبی، مثال: `/zones/{zoneId}/dns_records`
 * @param options تنظیمات درخواست
 */
export async function cloudflareRequest<T = unknown> (
  endpoint: string,
  options: { method?: string, query?: Record<string, unknown>, body?: unknown } = {}
): Promise<T> {
  const baseURL = 'https://api.cloudflare.com/client/v4'
  const headers = buildAuthHeaders()

  try {
    const response = await ofetch<{ success?: boolean; errors?: Array<{ code?: number; message?: string }>; messages?: Array<{ code?: number; message?: string }>; result?: T }>(endpoint, {
      baseURL,
      method: options.method || 'GET',
      headers,
      query: options.query,
      body: options.body ? JSON.stringify(options.body) : undefined,
    })

    // Cloudflare پاسخ را به شکل { success, result, errors, messages } برمی‌گرداند
    if (response && response.success) {
      return response.result as T
    }

    // اگر success=false بود، خطا را با جزئیات برگردانیم
    const firstError = response?.errors?.[0]
    const message = firstError?.message || 'خطای ناشناخته از Cloudflare'
    const code = firstError?.code
    const err = new Error(`${message}${code ? ` (code: ${code})` : ''}`)
    // @ts-ignore
    err.data = response
    throw err
  } catch (error) {
    // خطاهای سطح شبکه یا پارسینگ
    throw error
  }
}

// انواع کمکی ساده برای استفاده در هندلرها
export interface CloudflareResultWrapper<T> {
  success: boolean
  errors: Array<{ code?: number, message: string }>
  messages: Array<{ code?: number, message: string }>
  result: T
}


