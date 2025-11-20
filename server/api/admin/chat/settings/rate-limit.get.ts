import type { H3Event } from 'h3'

interface User {
  token?: string
  [key: string]: unknown
}

declare const defineEventHandler: (handler: (event: H3Event) => unknown | Promise<unknown>) => unknown
declare const requireAuth: (event: H3Event) => Promise<User | null>
declare const createError: (options: { statusCode: number; message: string }) => Error
declare const hasPermission: (user: User, permission: string) => boolean
declare const useRuntimeConfig: () => { public: { goApiBase: string } }
declare const $fetch: (url: string, options?: { headers?: Record<string, string> }) => Promise<{ data?: unknown }>

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
    if (!hasPermission(user, 'chat_settings.read')) {
      throw createError({
        statusCode: 404,
        message: 'دسترسی غیرمجاز'
      })
    }

    // دریافت تنظیمات از سرور Go
    const config = useRuntimeConfig()
    const response = await $fetch(`${config.public.goApiBase}/api/chat/admin/settings/rate-limit`, {
      headers: {
        'Authorization': `Bearer ${user.token}`,
        'Content-Type': 'application/json'
      }
    })

    return {
      success: true,
      data: response.data
    }

  } catch (error: unknown) {
    // console.error('خطا در دریافت تنظیمات نرخ پیام:', error)

    const errorWithStatus = error as { statusCode?: number; statusMessage?: string }
    throw createError({
      statusCode: errorWithStatus?.statusCode || 500,
      message: errorWithStatus?.statusMessage || 'خطا در دریافت تنظیمات'
    })
  }
}) 
