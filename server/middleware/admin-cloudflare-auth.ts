// میان‌افزار احراز هویت برای مسیرهای Cloudflare ادمین
// اگر کاربر کوکی admin_token نداشته باشد، دسترسی منع می‌شود
import { eventHandler, getCookie, createError } from 'h3'

export default eventHandler((event) => {
  const url = event.path || event.node?.req?.url || ''
  if (!url.startsWith('/api/admin/cloudflare')) return

  const token = getCookie(event, 'admin_token')
  if (!token) {
    throw createError({ statusCode: 401, message: 'دسترسی غیرمجاز' })
  }
})