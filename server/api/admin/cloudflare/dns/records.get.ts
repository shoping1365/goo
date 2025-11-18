// server/api/admin/cloudflare/dns/records.get.ts
// لیست رکوردهای DNS در Cloudflare

import { createError, defineEventHandler } from 'h3'

interface CFDNSRecord {
  id: string
  type: string
  name: string
  content: string
  proxied?: boolean
  ttl?: number
}

interface CloudflareResultWrapper<T> {
  result?: T | CFDNSRecord[]
  success?: boolean
  [key: string]: unknown
}

declare const cloudflareRequest: <T = CloudflareResultWrapper<{ result: CFDNSRecord[] }>>(endpoint: string, options?: { method?: string; body?: unknown }) => Promise<T>
declare const getCloudflareZoneId: () => string

export default defineEventHandler(async () => {
  try {
    const zoneId = getCloudflareZoneId()
    const res = await cloudflareRequest<CloudflareResultWrapper<{ result: CFDNSRecord[] }>>(`/zones/${zoneId}/dns_records`)
    // Cloudflare معمولا به صورت { result: [...], success: true } بازمی‌گرداند
    // در اینجا برای سازگاری، داده را صاف می‌کنیم
    const recordsArray = (res.result && Array.isArray(res.result))
      ? res.result
      : ((res.result as { result?: CFDNSRecord[] })?.result || [])
    const records = recordsArray as CFDNSRecord[]
    return { success: true, data: records }
  } catch (error: unknown) {
    const errorWithMessage = error as { message?: string }
    throw createError({ statusCode: 500, message: errorWithMessage?.message || 'خطا در دریافت رکوردهای DNS' })
  }
})


