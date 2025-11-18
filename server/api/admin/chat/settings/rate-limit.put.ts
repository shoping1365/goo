import type { H3Event } from 'h3'
import { createError, defineEventHandler, readBody } from 'h3'

interface User {
  token?: string
  [key: string]: unknown
}

interface RateLimitBody {
  customer_limit?: number
  operator_limit?: number
  admin_limit?: number
  window_seconds?: number
  [key: string]: unknown
}

declare const requireAuth: (event: H3Event) => Promise<User | null>
declare const hasPermission: (user: User, permission: string) => boolean
declare const useRuntimeConfig: () => { public: { goApiBase: string } }
declare const $fetch: <T = unknown>(url: string, options?: { method?: string; headers?: Record<string, string>; body?: unknown }) => Promise<{ data?: unknown }>

export default defineEventHandler(async (event: H3Event) => {
  try {
    // بررسی احراز هویت
    const user = await requireAuth(event)
    if (!user) {
      throw createError({
        statusCode: 401,
        message: 'احراز هویت مورد نیاز است'
      })
    }

    // بررسی مجوز
    if (!hasPermission(user, 'chat_settings.write')) {
      throw createError({
        statusCode: 404,
        message: 'دسترسی غیرمجاز'
      })
    }

    // دریافت داده‌های ورودی
    const body = await readBody(event) as RateLimitBody

    // اعتبارسنجی داده‌ها
    if (!body.customer_limit || !body.operator_limit || !body.admin_limit || !body.window_seconds) {
      throw createError({
        statusCode: 400,
        message: 'تمام فیلدها الزامی هستند'
      })
    }

    if (body.customer_limit < 1 || body.customer_limit > 1000) {
      throw createError({
        statusCode: 400,
        message: 'محدودیت مشتری باید بین 1 تا 1000 باشد'
      })
    }

    if (body.operator_limit < 1 || body.operator_limit > 1000) {
      throw createError({
        statusCode: 400,
        message: 'محدودیت اپراتور باید بین 1 تا 1000 باشد'
      })
    }

    if (body.admin_limit < 1 || body.admin_limit > 1000) {
      throw createError({
        statusCode: 400,
        message: 'محدودیت ادمین باید بین 1 تا 1000 باشد'
      })
    }

    if (body.window_seconds < 10 || body.window_seconds > 3600) {
      throw createError({
        statusCode: 400,
        message: 'بازه زمانی باید بین 10 تا 3600 ثانیه باشد'
      })
    }

    // ارسال به سرور Go
    const config = useRuntimeConfig()
    const response = await $fetch(`${config.public.goApiBase}/api/chat/admin/settings/rate-limit`, {
      method: 'PUT',
      headers: {
        'Authorization': `Bearer ${user.token}`,
        'Content-Type': 'application/json'
      },
      body: {
        customer_limit: body.customer_limit,
        operator_limit: body.operator_limit,
        admin_limit: body.admin_limit,
        window_seconds: body.window_seconds
      }
    })

    return {
      success: true,
      message: 'تنظیمات با موفقیت بروزرسانی شد',
      data: response.data
    }

  } catch (error: unknown) {
    console.error('خطا در بروزرسانی تنظیمات نرخ پیام:', error)

    const errorWithStatus = error as { statusCode?: number; statusMessage?: string }
    throw createError({
      statusCode: errorWithStatus?.statusCode || 500,
      message: errorWithStatus?.statusMessage || 'خطا در بروزرسانی تنظیمات'
    })
  }
}) 
