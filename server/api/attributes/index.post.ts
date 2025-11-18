import { defineEventHandler, readBody } from 'h3'
import { proxy } from '../_utils/fetchProxy'

export default defineEventHandler(async (event) => {
  const base = getGoApiBaseUrl()
  const payload = await readBody(event)

  // با استفاده از util‌ proxy خطاهای بک‌اند (کد و پیام) بدون تغییر به کلاینت برمی‌گردند.
  return proxy(event, `${base}/api/attributes`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload)
  })
})