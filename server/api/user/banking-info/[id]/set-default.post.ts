import { createError, defineEventHandler, getRouterParam } from 'h3'
import { goApiFetch } from '~/server/utils/goApi'

export default defineEventHandler(async (event) => {
  try {
    const id = getRouterParam(event, 'id')

    if (!id) {
      throw createError({
        statusCode: 400,
        message: 'شناسه حساب الزامی است'
      })
    }

    // ارسال درخواست به بک‌اند Go
    await goApiFetch(event, `/api/user/banking-info/${id}/set-default`, {
      method: 'POST'
    })

    return {
      success: true,
      message: 'حساب با موفقیت به عنوان پیشفرض تنظیم شد'
    }
  } catch (error: unknown) {
    console.error('خطا در تنظیم حساب پیشفرض:', error)

    if (error && typeof error === 'object' && 'statusCode' in error) {
      throw error
    }

    throw createError({
      statusCode: 500,
      message: 'خطا در تنظیم حساب پیشفرض'
    })
  }
})