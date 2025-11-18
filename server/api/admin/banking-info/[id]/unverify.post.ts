import type { H3Event } from 'h3';
import { createError, defineEventHandler, getRouterParam, readBody } from 'h3';

declare const requireAdminAuth: (event: H3Event) => Promise<{ [key: string]: unknown } | null>
declare const goApiFetch: (event: H3Event, endpoint: string, options?: { method?: string; body?: unknown }) => Promise<unknown>

export default defineEventHandler(async (event) => {
  try {
    // احراز هویت ادمین
    const adminUser = await requireAdminAuth(event)
    if (!adminUser) {
      throw createError({
        statusCode: 401,
        message: 'احراز هویت ادمین الزامی است'
      })
    }

    // دریافت پارامترهای URL
    const bankingInfoId = getRouterParam(event, 'id')
    if (!bankingInfoId) {
      throw createError({
        statusCode: 400,
        message: 'شناسه اطلاعات بانکی الزامی است'
      })
    }

    // دریافت body
    const body = await readBody(event)

    // ارسال درخواست به بک‌اند Go
    const response = await goApiFetch(event, `/api/admin/banking-info/${bankingInfoId}/unverify`, {
      method: 'POST',
      body: {
        note: body.note || ''
      }
    })

    return {
      success: true,
      message: 'تایید اطلاعات بانکی لغو شد'
    }

  } catch (error: unknown) {
    console.error('خطا در لغو تایید اطلاعات بانکی:', error)

    const errorWithStatus = error as { statusCode?: number; message?: string }
    if (error && typeof error === 'object' && 'statusCode' in error) {
      throw error
    }

    throw createError({
      statusCode: 500,
      message: 'خطا در لغو تایید اطلاعات بانکی'
    })
  }
})



