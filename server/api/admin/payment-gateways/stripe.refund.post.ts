import { createError, defineEventHandler, getCookie, readBody } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

export default defineEventHandler(async (event) => {
  try {
    const body = await readBody(event)

    const response = await fetchGo(event, '/api/admin/payment-gateways/stripe/refund', {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${getCookie(event, 'admin_token') || ''}`,
        'Content-Type': 'application/json'
      },
      body
    })

    return response
  } catch (error: unknown) {
    const errorObj = error as { statusCode?: number; statusMessage?: string }
    throw createError({
      statusCode: errorObj.statusCode || 500,
      message: errorObj.statusMessage || 'خطا در بازگشت وجه درگاه استرایپ'
    })
  }
}) 