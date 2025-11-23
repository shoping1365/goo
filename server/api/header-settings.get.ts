import { createError, defineEventHandler, setResponseHeader } from 'h3'
import type { ServerResponse } from 'http'
import { fetchGo } from './_utils/fetchGo'

interface PublicHeaderResponse {
  success: boolean
  data: unknown[]
}

export default defineEventHandler(async (event): Promise<PublicHeaderResponse> => {
  try {
    const responseData = await fetchGo(event, '/api/header-settings', {
      method: 'GET'
    })

    // جلوگیری از کش شدن پاسخ برای SSR (سازگار با هر محیط)
    const cacheControl = 'no-store, max-age=0, must-revalidate'
    const nodeRes = event.node?.res as ServerResponse
    if (nodeRes && typeof nodeRes.setHeader === 'function') {
      nodeRes.setHeader('Cache-Control', cacheControl)
    } else {
      try {
        setResponseHeader(event, 'Cache-Control', cacheControl)
      } catch (headerErr) {
        console.warn('[header-settings] Unable to set Cache-Control header', headerErr)
      }
    }

    if (responseData && typeof responseData === 'object') {
      const payload = responseData as Record<string, unknown>
      return {
        success: payload.success !== false,
        data: (payload.data as unknown[]) || []
      }
    }

    return {
      success: true,
      data: Array.isArray(responseData) ? responseData : []
    }

  } catch (error: unknown) {
    const err = error as { statusCode?: number; status?: number; message?: string; data?: unknown; error?: string }
    console.error('خطا در دریافت تنظیمات هدر:', {
      statusCode: err?.statusCode,
      status: err?.status,
      message: err?.message,
      data: err?.data
    })

    throw createError({
      statusCode: err?.statusCode || err?.status || 500,
      message: err?.data?.message || err?.data?.error || err?.message || 'خطا در دریافت تنظیمات هدر',
      data: err?.data
    })
  }
}) 