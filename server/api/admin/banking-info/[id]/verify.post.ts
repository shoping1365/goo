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
    await goApiFetch(event, `/api/admin/banking-info/${bankingInfoId}/verify`, {
      method: 'POST',
      body: {
        note: body.note || ''
      }
    })

    return {
      success: true,
      message: 'اطلاعات بانکی با موفقیت تایید شد'
    }

  } catch (error: unknown) {
    // console.error('خطا در تایید اطلاعات بانکی:', error)

    if (error && typeof error === 'object' && 'statusCode' in error) {
      throw error
    }

    throw createError({
      statusCode: 500,
      message: 'خطا در تایید اطلاعات بانکی'
    })
  }
})



