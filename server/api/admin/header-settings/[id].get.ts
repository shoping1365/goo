import { createError, defineEventHandler, getRouterParam, parseCookies } from 'h3'

declare const useRuntimeConfig: () => { public: { goApiBase: string } }

interface HeaderData {
  [key: string]: unknown
}

interface HeaderResponse {
  success: boolean
  data: HeaderData
}

export default defineEventHandler(async (event): Promise<HeaderResponse> => {
  try {
    const id = getRouterParam(event, 'id')
    const config = useRuntimeConfig()

    console.log('درخواست دریافت هدر:', id)

    // دریافت cookies از درخواست
    const cookies = parseCookies(event)
    const cookieHeader = Object.entries(cookies)
      .map(([key, value]) => `${key}=${value}`)
      .join('; ')

    // ارسال درخواست به Go backend با fetch معمولی
    const response = await fetch(`${config.public.goApiBase}/api/admin/header-settings/${id}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'Cookie': cookieHeader
      }
    })

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }

    const responseData = await response.json()

    console.log('پاسخ دریافت هدر:', responseData)

    return {
      success: true,
      data: responseData?.data || responseData
    }

  } catch (error: unknown) {
    console.error('خطا در دریافت هدر:', error)

    // اگر خطا از سرور Go آمده باشد
    const errorWithData = error as { data?: { message?: string; error?: string }; statusCode?: number }
    if (errorWithData?.data) {
      throw createError({
        statusCode: errorWithData.statusCode || 500,
        message: errorWithData.data.message || errorWithData.data.error || 'خطا در دریافت هدر',
        data: errorWithData.data
      })
    }

    // اگر خطای شبکه باشد
    throw createError({
      statusCode: 500,
      message: 'خطا در اتصال به سرور'
    })
  }
}) 