import { createError, defineEventHandler, getQuery } from 'h3'
import { fetchGo } from '../../../_utils/fetchGo'

interface TransactionsResponse {
  success: boolean
  data?: unknown
  message?: string
}

export default defineEventHandler(async (event) => {
  try {
    // بررسی احراز هویت ادمین
    const user = event.context.user
    if (!user || !user.is_admin) {
      throw createError({
        statusCode: 403,
        message: 'دسترسی غیرمجاز'
      })
    }

    const query = getQuery(event)
    const page = parseInt(query.page as string) || 1
    const limit = parseInt(query.limit as string) || 20

    // ارسال درخواست به backend از طریق fetchGo
    const response = await fetchGo(event, `/api/admin/payment-gateways/melli/transactions?page=${page}&limit=${limit}`) as TransactionsResponse

    return {
      success: true,
      data: response.data,
      message: 'تراکنش‌های درگاه ملی با موفقیت دریافت شد'
    }

  } catch (error: unknown) {
    console.error('خطا در دریافت تراکنش‌های ملی:', error)

    const errorObj = error as { statusCode?: number; statusMessage?: string }
    throw createError({
      statusCode: errorObj.statusCode || 500,
      message: errorObj.statusMessage || 'خطا در دریافت تراکنش‌های درگاه ملی'
    })
  }
}) 