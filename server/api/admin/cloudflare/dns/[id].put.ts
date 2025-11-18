// server/api/admin/cloudflare/dns/[id].put.ts
// ویرایش رکورد DNS بر اساس شناسه

import { createError, defineEventHandler, getRouterParam, readBody } from 'h3'

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
  type?: string
  name?: string
  content?: string
  proxied?: boolean
  ttl?: number
}

declare const cloudflareRequest: <T = CloudflareResponse>(endpoint: string, options?: { method?: string; body?: unknown }) => Promise<T>
declare const getCloudflareZoneId: () => string

export default defineEventHandler(async (event) => {
  try {
    const zoneId = getCloudflareZoneId()
    const id = getRouterParam(event, 'id')
    if (!id) throw createError({ statusCode: 400, message: 'شناسه رکورد الزامی است' })

    const body = await readBody(event) as DNSRecordBody
    const payload: DNSPayload = {}
    if (body?.type) payload.type = String(body.type)
    if (body?.name) payload.name = String(body.name)
    if (body?.content) payload.content = String(body.content)
    if (typeof body?.proxied === 'boolean') payload.proxied = body.proxied
    if (body?.ttl !== undefined) payload.ttl = Number(body.ttl)

    const res = await cloudflareRequest<CloudflareResponse>(`/zones/${zoneId}/dns_records/${id}`, {
      method: 'PUT',
      body: payload,
    })

    return { success: true, data: res?.result }
  } catch (error: unknown) {
    const errorWithMessage = error as { message?: string }
    throw createError({ statusCode: 500, message: errorWithMessage?.message || 'خطا در بروزرسانی رکورد DNS' })
  }
})


