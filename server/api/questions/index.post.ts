import { defineEventHandler, readBody, createError, H3Event } from 'h3'
import { goApiFetch } from '~/server/utils/goApi'

interface QuestionBody {
  question: string
  product_id: number
  customer_name?: string
  customer_mobile?: string
  category?: string
  is_anonymous?: boolean
}

interface QuestionResponse {
  success: boolean
  message: string
  data?: any
}

export default defineEventHandler(async (event: H3Event): Promise<QuestionResponse> => {
  const body = await readBody(event) as QuestionBody

  try {
    const response = await goApiFetch<QuestionResponse>(event, '/api/questions', {
      method: 'POST',
      body: {
        question: body.question,
        product_id: body.product_id,
        customer_name: body.customer_name || '',
        customer_mobile: body.customer_mobile || '',
        category: body.category || 'general',
        is_anonymous: body.is_anonymous || false,
        ip_address: getClientIP(event)
      }
    })
    return response
  } catch (error: any) {
    console.error('Error creating question:', error)
    throw createError({
      statusCode: error.statusCode || 500,
      message: error.message || error.statusMessage || 'خطا در ثبت سوال'
    })
  }
})

/**
 * استخراج IP آدرس کلاینت از هدرهای مختلف
 * اولویت: x-forwarded-for > x-real-ip > connection.remoteAddress
 */
function getClientIP(event: H3Event): string {
  const headers = event.node.req.headers

  // بررسی x-forwarded-for (ممکن است شامل چند IP باشد)
  const xForwardedFor = headers['x-forwarded-for'] as string
  if (xForwardedFor) {
    // اگر چند IP وجود دارد، اولین IP را برمی‌گردانیم
    return xForwardedFor.split(',')[0].trim()
  }

  // بررسی x-real-ip
  const xRealIP = headers['x-real-ip'] as string
  if (xRealIP) {
    return xRealIP.trim()
  }

  // بررسی connection.remoteAddress
  const remoteAddress = event.node.req.connection?.remoteAddress
  if (remoteAddress) {
    return remoteAddress
  }

  // بررسی socket.remoteAddress (برای HTTP/2)
  const socketAddress = (event.node.req as any).socket?.remoteAddress
  if (socketAddress) {
    return socketAddress
  }

  return 'unknown'
}