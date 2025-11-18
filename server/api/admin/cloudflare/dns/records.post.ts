// server/api/admin/cloudflare/dns/records.post.ts
// ایجاد رکورد DNS جدید

import { createError, defineEventHandler, readBody } from 'h3'

interface CloudflareResponse {
  result?: {
    [key: string]: unknown
  }
  [key: string]: unknown
}

interface DNSRecordBody {
  type?: string | number
  name?: string
  content?: string
  proxied?: boolean
  ttl?: number | string
  [key: string]: unknown
}

interface DNSPayload {
  type: string
  name: string
  content: string
  ttl: number
  proxied: boolean
}

declare const cloudflareRequest: <T = CloudflareResponse>(endpoint: string, options?: { method?: string; body?: unknown }) => Promise<T>
declare const getCloudflareZoneId: () => string

export default defineEventHandler(async (event) => {
  try {
    const zoneId = getCloudflareZoneId()
    const body = await readBody(event) as DNSRecordBody
    // اعتبارسنجی مقدماتی ورودی
    if (!body?.type || !body?.name || !body?.content) {
      throw createError({ statusCode: 400, message: 'فیلدهای type, name, content اجباری هستند' })
    }

    const payload: DNSPayload = {
      type: String(body.type),
      name: String(body.name),
      content: String(body.content),
      ttl: body.ttl ? Number(body.ttl) : 1, // 1 = Auto
      proxied: typeof body.proxied === 'boolean' ? body.proxied : false,
    }

    const res = await cloudflareRequest<CloudflareResponse>(`/zones/${zoneId}/dns_records`, {
      method: 'POST',
      body: payload,
    })

    return { success: true, data: res?.result }
  } catch (error: unknown) {
    const errorWithStatus = error as { statusCode?: number; message?: string }
    throw createError({ statusCode: errorWithStatus?.statusCode || 500, message: errorWithStatus?.message || 'خطا در ایجاد رکورد DNS' })
  }
})


