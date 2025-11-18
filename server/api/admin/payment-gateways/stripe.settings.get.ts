import { defineEventHandler, getCookie, createError } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

export default defineEventHandler(async (event) => {
  try {
    const response = await fetchGo(event, '/api/admin/payment-gateways/stripe/settings', {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${getCookie(event, 'admin_token') || ''}`
      }
    })
    
    return response
  } catch (error: unknown) {
    const errorObj = error as { statusCode?: number; statusMessage?: string }
    throw createError({
      statusCode: errorObj.statusCode || 500,
      message: errorObj.statusMessage || 'خطا در دریافت تنظیمات درگاه استرایپ'
    })
  }
}) 