// server/api/admin/cloudflare/ddos/settings.get.ts
// دریافت تنظیمات حفاظتی (Under Attack Mode و غیره)

import { createError, defineEventHandler } from 'h3'

interface CloudflareResponse {
  result?: {
    value?: string
    [key: string]: unknown
  }
  [key: string]: unknown
}

declare const cloudflareRequest: <T = CloudflareResponse>(endpoint: string, options?: { method?: string; body?: unknown }) => Promise<T>
declare const getCloudflareZoneId: () => string

export default defineEventHandler(async () => {
  try {
    const zoneId = getCloudflareZoneId()

    // مثال: خواندن حالت Under Attack Mode از /zones/:id/settings/security_level
    const secLevel = await cloudflareRequest<CloudflareResponse>(`/zones/${zoneId}/settings/security_level`)

    return {
      success: true,
      data: {
        securityLevel: secLevel?.result?.value || 'medium',
      }
    }
  } catch (error: unknown) {
    const errorWithMessage = error as { message?: string }
    throw createError({ statusCode: 500, message: errorWithMessage?.message || 'خطا در دریافت تنظیمات DDoS' })
  }
})


