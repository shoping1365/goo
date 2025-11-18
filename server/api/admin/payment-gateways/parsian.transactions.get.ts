import { defineEventHandler, getQuery, getCookie } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

interface TransactionsResponse {
  success: boolean
  data?: unknown
  message?: string
}

export default defineEventHandler(async (event) => {
  try {
    const query = getQuery(event)
    const page = query.page || '1'
    const limit = query.limit || '20'

    // دریافت تراکنش‌های درگاه پارسیان
    const response = await fetchGo(event, `/api/admin/parsian/transactions?page=${page}&limit=${limit}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${getCookie(event, 'admin_token') || ''}`
      }
    }) as TransactionsResponse

    return {
      success: true,
      data: response.data,
      message: 'تراکنش‌های درگاه پارسیان با موفقیت دریافت شد'
    }
  } catch (error: unknown) {
    console.error('خطا در دریافت تراکنش‌های درگاه پارسیان:', error)
    const errorObj = error as { data?: { message?: string } }
    
    return {
      success: false,
      message: errorObj.data?.message || 'خطا در دریافت تراکنش‌های درگاه پارسیان'
    }
  }
}) 
 