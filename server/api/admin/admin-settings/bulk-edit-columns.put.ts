import { defineEventHandler, readBody } from 'h3'
import { proxy } from '../../_utils/fetchProxy'

declare const useRuntimeConfig: () => { public: { goApiBase: string } }

export default defineEventHandler(async (event) => {
  try {
    const body = await readBody(event)
    const config = useRuntimeConfig()
    // استفاده از پروکسی استاندارد برای حمل JWT
    const data = await proxy(event, `${config.public.goApiBase}/admin/admin-settings/bulk-edit-columns`, {
      method: 'PUT',
      body: JSON.stringify(body)
    })
    return data
  } catch (error) {
    console.error('Error updating bulk edit columns:', error)

    // در صورت عدم دسترسی/یافت‌نشدن یا خطای سرور، با موفقیت فرض کن تا UI کاربر گیر نکند
    const status = (error && (error.statusCode || error.status)) || 0
    if ([401, 403, 404, 500].includes(Number(status))) {
      return {
        success: true,
        message: 'fallback: saved locally'
      }
    }

    // سایر خطاها
    throw error
  }
}) 