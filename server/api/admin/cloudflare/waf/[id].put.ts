// server/api/admin/cloudflare/waf/[id].put.ts
// بروزرسانی قانون WAF

import { createError, defineEventHandler, getRouterParam, readBody } from 'h3'

interface CloudflareResponse {
  result?: {
    [key: string]: unknown
  }[]
  [key: string]: unknown
}

interface WAFRuleBody {
  [key: string]: unknown
}

declare const cloudflareRequest: <T = CloudflareResponse>(endpoint: string, options?: { method?: string; body?: unknown }) => Promise<T>
declare const getCloudflareZoneId: () => string

export default defineEventHandler(async (event) => {
  try {
    const zoneId = getCloudflareZoneId()
    const id = getRouterParam(event, 'id')
    if (!id) throw createError({ statusCode: 400, message: 'شناسه قانون الزامی است' })

    const body = await readBody(event) as WAFRuleBody

    // Cloudflare API برای بروزرسانی نیاز به آرایه دارد
    const payload = [{ id, ...body }]
    const res = await cloudflareRequest<CloudflareResponse>(`/zones/${zoneId}/firewall/rules`, { method: 'PUT', body: payload })
    return { success: true, data: res?.result?.[0] }
  } catch (error: unknown) {
    const errorWithMessage = error as { message?: string }
    throw createError({ statusCode: 500, message: errorWithMessage?.message || 'خطا در بروزرسانی قانون WAF' })
  }
})


