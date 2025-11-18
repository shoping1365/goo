import { createError, defineEventHandler, getCookie, getQuery } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

export default defineEventHandler(async (event) => {
  try {
    const query = getQuery(event)
    const page = query.page || '1'
    const limit = query.limit || '20'

    const queryString = new URLSearchParams({
      page: String(page),
      limit: String(limit)
    }).toString()

    const response = await fetchGo(event, `/api/admin/payment-gateways/paypal/transactions?${queryString}`, {
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
      message: errorObj.statusMessage || 'خطا در دریافت تراکنش‌های درگاه پی‌پال'
    })
  }
}) 