// server/api/admin/cloudflare/dns/[id].delete.ts
// حذف رکورد DNS بر اساس شناسه

import { createError, defineEventHandler, getRouterParam } from 'h3'

interface CloudflareResponse {
  result?: {
    [key: string]: unknown
  }
  [key: string]: unknown
}

declare const cloudflareRequest: <T = CloudflareResponse>(endpoint: string, options?: { method?: string; body?: unknown }) => Promise<T>
declare const getCloudflareZoneId: () => string

export default defineEventHandler(async (event) => {
  try {
    const zoneId = getCloudflareZoneId()
    const id = getRouterParam(event, 'id')
    if (!id) throw createError({ statusCode: 400, message: 'شناسه رکورد الزامی است' })

    const res = await cloudflareRequest<CloudflareResponse>(`/zones/${zoneId}/dns_records/${id}`, { method: 'DELETE' })
    return { success: true, data: res?.result }
  } catch (error: unknown) {
    const errorWithMessage = error as { message?: string }
    throw createError({ statusCode: 500, message: errorWithMessage?.message || 'خطا در حذف رکورد DNS' })
  }
})


