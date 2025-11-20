import { defineEventHandler, parseCookies } from 'h3'

declare const useRuntimeConfig: () => { public: { goApiBase: string } }

interface HeaderItem {
  [key: string]: unknown
}

interface HeaderListResponse {
  success: boolean
  data: HeaderItem[]
  message?: string
}

export default defineEventHandler(async (event): Promise<HeaderListResponse> => {
  try {
    const config = useRuntimeConfig()

    // console.log('درخواست دریافت لیست هدرها (ادمین)')

    // دریافت cookies از درخواست
    const cookies = parseCookies(event)
    const cookieHeader = Object.entries(cookies)
      .map(([key, value]) => `${key}=${value}`)
      .join('; ')

    // ارسال درخواست به Go backend با fetch معمولی
    const response = await fetch(`${config.public.goApiBase}/api/admin/header-settings`, {
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

    // console.log('پاسخ لیست هدرها:', responseData)

    return {
      success: true,
      data: responseData?.data || []
    }

  } catch (error: unknown) {
    // console.error('خطا در دریافت لیست هدرها:', error)

    // اگر خطا از سرور Go آمده باشد
    const errorWithData = error as { data?: { message?: string; error?: string } }
    if (errorWithData?.data) {
      return {
        success: false,
        data: [],
        message: errorWithData.data.message || errorWithData.data.error || 'خطا در دریافت لیست هدرها'
      }
    }

    // اگر خطای شبکه باشد
    return {
      success: false,
      data: [],
      message: 'خطا در اتصال به سرور'
    }
  }
})