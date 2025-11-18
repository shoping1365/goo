import { defineEventHandler, createError, getRouterParam } from 'h3'
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
    await goApiFetch(event, `/api/user/banking-info/${id}`, {
      method: 'DELETE'
    })

    return {
      success: true,
      message: 'حساب با موفقیت حذف شد'
    }
  } catch (error: any) {
    console.error('خطا در حذف حساب بانکی:', error)
    
    if (error && typeof error === 'object' && 'statusCode' in error) {
      throw error
    }
    
    throw createError({
      statusCode: 500,
      message: 'خطا در حذف حساب بانکی'
    })
  }
})