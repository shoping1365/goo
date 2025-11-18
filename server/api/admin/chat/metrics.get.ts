import type { H3Event } from 'h3';
import { proxy } from '../../_utils/fetchProxy';

declare const defineEventHandler: (handler: (event: H3Event) => unknown | Promise<unknown>) => unknown
declare const useRuntimeConfig: () => { public: { goApiBase: string } }
declare const createError: (options: { statusCode: number; message: string }) => Error

interface MetricsResponse {
  metrics?: {
    [key: string]: unknown
  }
  health?: {
    [key: string]: unknown
  }
  alerts?: {
    [key: string]: unknown
  }[]
  [key: string]: unknown
}

export default defineEventHandler(async (event) => {
  try {
    // موقتاً غیرفعال برای تست (auth removed)
    // if (!hasPermission(user, 'chat_metrics.read')) {
    //   throw createError({ statusCode: 404, message: 'دسترسی غیرمجاز' })
    // }

    const config = useRuntimeConfig()
    const url = `${config.public.goApiBase}/api/chat/admin/metrics`
    const response = await proxy(event, url, {
      method: 'GET',
      headers: {
        // Authentication headers will be added by proxy function
      }
    }) as MetricsResponse

    return {
      success: true,
      metrics: response?.metrics,
      health: response?.health,
      alerts: response?.alerts
    }
  } catch (error: unknown) {
    console.error('خطا در دریافت متریک‌های چت:', error)
    const errorWithStatus = error as { statusCode?: number; statusMessage?: string }
    throw createError({ statusCode: errorWithStatus?.statusCode || 500, message: errorWithStatus?.statusMessage || 'خطا در دریافت متریک‌ها' })
  }
})

