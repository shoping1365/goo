// server/api/admin/cloudflare/waf/[id].delete.ts
// حذف قانون WAF

import { createError, defineEventHandler, getRouterParam } from 'h3'

interface CloudflareResponse {
  result?: {
    [key: string]: unknown
  }[]
  [key: string]: unknown
}

declare const cloudflareRequest: <T = CloudflareResponse>(endpoint: string, options?: { method?: string; body?: unknown }) => Promise<T>
declare const getCloudflareZoneId: () => string

export default defineEventHandler(async (event) => {
  try {
    const zoneId = getCloudflareZoneId()
    const id = getRouterParam(event, 'id')
    if (!id) throw createError({ statusCode: 400, message: 'شناسه قانون الزامی است' })

    // Cloudflare API حذف تکی را با ارسال آرایه از اشیاء پشتیبانی می‌کند
    const res = await cloudflareRequest<CloudflareResponse>(`/zones/${zoneId}/firewall/rules`, {
      method: 'DELETE',
      body: [{ id }],
    })
    return { success: true, data: res?.result }
  } catch (error: unknown) {
    const errorWithMessage = error as { message?: string }
    throw createError({ statusCode: 500, message: errorWithMessage?.message || 'خطا در حذف قانون WAF' })
  }
})


