// server/api/admin/cloudflare/overview.get.ts
// دریافت نمای کلی از وضعیت Cloudflare (ترافیک، کش، تهدیدات)

import { createError, defineEventHandler } from 'h3'

interface CFZoneAnalyticsSummary {
  totals?: {
    requests?: {
      all?: number | string
      cached?: number | string
      [key: string]: unknown
    }
    threats?: {
      all?: number | string
      [key: string]: unknown
    }
    bandwidth?: {
      all?: number | string
      [key: string]: unknown
    }
    [key: string]: unknown
  }
  [key: string]: unknown
}

interface CloudflareResultWrapper<T> {
  result?: T
  success?: boolean
  [key: string]: unknown
}

declare const cloudflareRequest: <T = CloudflareResultWrapper<CFZoneAnalyticsSummary>>(endpoint: string, options?: { method?: string; body?: unknown }) => Promise<T>
declare const getCloudflareZoneId: () => string

export default defineEventHandler(async () => {
  try {
    const zoneId = getCloudflareZoneId()

    // نمونه ساده از آمار: در API واقعی باید به /zones/:id/analytics/dashboard مراجعه کرد
    // برای سادگی و جلوگیری از پیچیدگی، از endpoint summary/day استفاده می‌کنیم (دموی تقریبی)
    const analytics = await cloudflareRequest<CloudflareResultWrapper<CFZoneAnalyticsSummary>>(`/zones/${zoneId}/analytics/summary`)

    const totals = analytics?.result?.totals || {}
    const requests = Number(totals.requests?.all || 0)
    const cached = Number(totals.requests?.cached || 0)
    const threats = Number(totals.threats?.all || 0)
    const bandwidth = Number(totals.bandwidth?.all || 0)

    return {
      success: true,
      data: {
        requests,
        cacheHitRate: requests > 0 ? Math.round((cached / requests) * 1000) / 10 : 0,
        threats,
        avgResponseMs: 45, // اگر نیاز بود بعدا از API واقعی محاسبه می‌شود
        bandwidth,
      }
    }
  } catch (error: unknown) {
    const errorWithMessage = error as { message?: string }
    throw createError({
      statusCode: 500,
      message: errorWithMessage?.message || 'خطا در دریافت نمای کلی کلادفلر'
    })
  }
})


