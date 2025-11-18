// server/api/admin/cloudflare/ddos/settings.put.ts
// تنظیم سطح امنیتی (low/medium/high/under_attack)

import { createError, defineEventHandler, readBody } from 'h3'

interface CloudflareResponse {
  result?: {
    [key: string]: unknown
  }
  [key: string]: unknown
}

interface SecurityLevelBody {
  securityLevel?: string
  [key: string]: unknown
}

declare const cloudflareRequest: <T = CloudflareResponse>(endpoint: string, options?: { method?: string; body?: unknown }) => Promise<T>
declare const getCloudflareZoneId: () => string

export default defineEventHandler(async (event) => {
  try {
    const zoneId = getCloudflareZoneId()
    const body = await readBody(event) as SecurityLevelBody
    const allowed = ['off', 'essentially_off', 'low', 'medium', 'high', 'under_attack']
    const value = String(body?.securityLevel || '')
    if (!allowed.includes(value)) {
      throw createError({ statusCode: 400, message: 'securityLevel نامعتبر است' })
    }

    const res = await cloudflareRequest<CloudflareResponse>(`/zones/${zoneId}/settings/security_level`, {
      method: 'PATCH',
      body: { value },
    })

    return { success: true, data: res?.result }
  } catch (error: unknown) {
    const errorWithMessage = error as { message?: string }
    throw createError({ statusCode: 500, message: errorWithMessage?.message || 'خطا در بروزرسانی تنظیمات DDoS' })
  }
})


