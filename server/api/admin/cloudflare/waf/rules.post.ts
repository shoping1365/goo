// server/api/admin/cloudflare/waf/rules.post.ts
// ایجاد قانون جدید WAF (Firewall Rule)

import { createError, defineEventHandler, readBody } from 'h3'

interface CloudflareResponse {
  result?: {
    id?: string
    [key: string]: unknown
  }[]
  [key: string]: unknown
}

interface WAFRuleBody {
  filter?: {
    id?: string
    expression?: string
    description?: string
    [key: string]: unknown
  }
  action?: string
  description?: string
  paused?: boolean
  [key: string]: unknown
}

interface FilterResponse {
  result?: {
    id?: string
    [key: string]: unknown
  }[]
  [key: string]: unknown
}

declare const cloudflareRequest: <T = CloudflareResponse>(endpoint: string, options?: { method?: string; body?: unknown }) => Promise<T>
declare const getCloudflareZoneId: () => string

export default defineEventHandler(async (event) => {
  try {
    const zoneId = getCloudflareZoneId()
    const body = await readBody(event) as WAFRuleBody
    if (!body?.filter || !body?.action) {
      throw createError({ statusCode: 400, message: 'فیلدهای filter و action الزامی است' })
    }

    // ساخت فیلتر اگر فقط expression داده شده باشد
    let filterId = body.filter?.id
    if (!filterId && body.filter?.expression) {
      const filterRes = await cloudflareRequest<FilterResponse>(`/zones/${zoneId}/filters`, {
        method: 'POST',
        body: [{ expression: String(body.filter.expression), description: String(body.filter.description || '') }]
      })
      filterId = filterRes?.result?.[0]?.id
    }

    if (!filterId) throw createError({ statusCode: 400, message: 'شناسه یا عبارت فیلتر معتبر نیست' })

    const payload = [{
      action: String(body.action),
      description: String(body.description || ''),
      filter: { id: String(filterId) },
      paused: body.paused === true ? true : false,
    }]

    const res = await cloudflareRequest<CloudflareResponse>(`/zones/${zoneId}/firewall/rules`, { method: 'POST', body: payload })
    return { success: true, data: res?.result?.[0] }
  } catch (error: unknown) {
    const errorWithMessage = error as { message?: string }
    throw createError({ statusCode: 500, message: errorWithMessage?.message || 'خطا در ایجاد قانون WAF' })
  }
})


