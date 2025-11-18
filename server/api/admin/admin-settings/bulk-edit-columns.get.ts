import { defineEventHandler } from 'h3'
import { proxy } from '../../_utils/fetchProxy'

declare const useRuntimeConfig: () => { public: { goApiBase: string } }

export default defineEventHandler(async (event) => {
  try {
    const config = useRuntimeConfig()
    // استفاده از پروکسی استاندارد تا JWT از کوکی به Authorization منتقل شود
    const data = await proxy(event, `${config.public.goApiBase}/admin/admin-settings/bulk-edit-columns`, { method: 'GET' })
    return data
  } catch (error: unknown) {
    console.error('Error fetching bulk edit columns:', error)

    // در صورت عدم دسترسی/یافت‌نشدن یا خطای سرور، اجازه بده UI پیش‌فرض‌ها را ست کند
    // با برگرداندن آبجکتی بدون فیلد value، شرط (response && response.value) فالس می‌شود
    const errorWithStatus = error as { statusCode?: number; status?: number }
    const status = errorWithStatus?.statusCode || errorWithStatus?.status || 0
    if ([401, 403, 404, 500].includes(Number(status))) {
      return { message: 'fallback: let-ui-set-defaults' }
    }

    // سایر خطاها
    throw error
  }
}) 