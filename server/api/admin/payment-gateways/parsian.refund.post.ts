import { defineEventHandler, getCookie, readBody } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

interface RefundBody {
  transaction_id: string
  amount: number
}

interface RefundResponse {
  success: boolean
  message?: string
  data?: unknown
}

export default defineEventHandler(async (event) => {
  try {
    const body = await readBody(event) as RefundBody
    const { transaction_id, amount } = body

    if (!transaction_id || !amount) {
      return {
        success: false,
        message: 'شناسه تراکنش و مبلغ الزامی است'
      }
    }

    // بازگشت وجه درگاه پارسیان از طریق fetchGo
    await fetchGo(event, '/api/admin/parsian/refund', {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${getCookie(event, 'admin_token') || ''}`,
        'Content-Type': 'application/json'
      },
      body: {
        transaction_id,
        amount
      }
    })

    return {
      success: true,
      message: 'بازگشت وجه پارسیان با موفقیت انجام شد'
    }
  } catch (error: unknown) {
    console.error('خطا در بازگشت وجه درگاه پارسیان:', error)
    const errorObj = error as { data?: { message?: string } }

    return {
      success: false,
      message: errorObj.data?.message || 'خطا در بازگشت وجه درگاه پارسیان'
    }
  }
})
