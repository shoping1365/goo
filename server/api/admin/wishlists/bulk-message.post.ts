import type { H3Event } from 'h3'
import { createError, defineEventHandler, readBody } from 'h3'

interface BulkMessageBody {
  recipientUserIds?: number[]
  productIds?: number[]
  type?: string
  message?: string
}

export default defineEventHandler(async (event: H3Event) => {
  try {
    const body = await readBody(event) as BulkMessageBody
    const recipientUserIds: number[] = body?.recipientUserIds || []
    const productIds: number[] = body?.productIds || []
    const type: string = body?.type || 'custom'
    const message: string = body?.message || ''

    if (!Array.isArray(recipientUserIds) || recipientUserIds.length === 0) {
      throw createError({ statusCode: 400, message: 'لیست گیرندگان خالی است' })
    }

    // شبیه‌سازی ثبت Job ارسال (اتصال بعدی به سیستم اعلان موجود)
    await new Promise((r) => setTimeout(r, 200))

    return {
      success: true,
      queued: true,
      recipients: recipientUserIds.length,
      productIds,
      type,
      messagePreview: (message || '').slice(0, 160)
    }
  } catch (error: unknown) {
    const errorWithStatus = error as { statusCode?: number; message?: string; statusMessage?: string }
    throw createError({
      statusCode: errorWithStatus.statusCode || 500,
      message: errorWithStatus.message || errorWithStatus.statusMessage || 'خطا در ثبت ارسال گروهی'
    })
  }
})


