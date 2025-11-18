// server/api/admin/cloudflare/waf/rules.get.ts
// لیست قوانین WAF سفارشی (Firewall Rules)

import { createError, defineEventHandler } from 'h3'

interface CloudflareResponse {
  result?: {
    [key: string]: unknown
  }[]
  [key: string]: unknown
}

declare const cloudflareRequest: <T = CloudflareResponse>(endpoint: string, options?: { method?: string; body?: unknown }) => Promise<T>
declare const getCloudflareZoneId: () => string

export default defineEventHandler(async () => {
  try {
    const zoneId = getCloudflareZoneId()
    const res = await cloudflareRequest<CloudflareResponse>(`/zones/${zoneId}/firewall/rules`)
    return { success: true, data: res?.result || [] }
  } catch (error: unknown) {
    const errorWithMessage = error as { message?: string }
    throw createError({ statusCode: 500, message: errorWithMessage?.message || 'خطا در دریافت قوانین WAF' })
  }
})


