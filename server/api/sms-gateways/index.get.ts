import { defineEventHandler } from 'h3'

export default defineEventHandler(async (_event) => {
  try {
    const config = useRuntimeConfig()
    const base = config.public.goApiBase
    const response = await fetch(`${base}/api/sms-gateways`, {
      method: 'GET',
      headers: {
        'Accept': 'application/json',
      },
    })
    if (!response.ok) {
      return {
        status: 'error',
        message: 'خطا در دریافت لیست درگاه‌ها',
        data: []
      }
    }
    const json = await response.json()

    // پشتیبانی از هر دو ساختار خروجی (data یا خود آرایه)
    const data = Array.isArray(json.data) ? json.data : (Array.isArray(json) ? json : [])

    // فیلتر کردن فقط درگاه‌های فعال
    const activeGateways = data.filter((gateway: { is_active?: boolean }) => gateway.is_active === true)

    return {
      status: 'success',
      data: activeGateways
    }
  } catch (error) {
    console.error('❌ خطا در ارتباط با سرور پیامک:', error)
    return {
      status: 'error',
      message: 'خطا در ارتباط با سرور پیامک',
      data: []
    }
  }
})