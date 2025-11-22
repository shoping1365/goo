import { defineEventHandler, getRouterParam, createError } from 'h3'

export default defineEventHandler(async (event) => {
  try {
    // دریافت ID از پارامترهای مسیر
    const id = getRouterParam(event, 'id')
    
    if (!id) {
      throw createError({
        statusCode: 400,
        message: 'شناسه درگاه پرداخت الزامی است',
      })
    }

    const config = useRuntimeConfig()
    const base = config.public.goApiBase
    // درخواست به بک‌اند
    const response = await $fetch(`${base}/api/payment-gateways/${id}`, {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json',
      },
    })

    return response
  } catch (error: unknown) {
    console.error('خطا در حذف درگاه پرداخت:', error)
    throw createError({
      statusCode: 500,
      message: 'خطا در حذف درگاه پرداخت',
    })
  }
})