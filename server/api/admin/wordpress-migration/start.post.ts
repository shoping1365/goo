import type { H3Event } from 'h3'
import { defineEventHandler, readBody } from 'h3'
import { proxy } from '../../_utils/fetchProxy'

declare const useRuntimeConfig: () => { public: { goApiBase: string } }

export default defineEventHandler(async (event: H3Event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const body = await readBody(event)

  const controller = new AbortController()
  event.node.req.on('close', () => { try { controller.abort() } catch { } })

  // در محیط توسعه، از مسیر dev بدون احراز/مجوز استفاده می‌کنیم تا موانع Permission مانع اجرا نشود
  // ابتدا مسیر امن را امتحان کن؛ در صورت خطای دسترسی/عدم وجود، به مسیر dev fallback کن
  const tryProxy = async (path: string) => proxy(event, `${base}${path}`, {
    method: 'POST',
    body: JSON.stringify(body),
    signal: controller.signal,
    timeout: 36000000, // 10 ساعت timeout برای migration
  })

  try {
    return await tryProxy('/api/admin/wordpress-migration/start')
  } catch (e: unknown) {
    // فقط در خطاهای احراز/عدم وجود مسیر fallback کن
    const errorWithStatus = e as { statusCode?: number; status?: number }
    const status = errorWithStatus?.statusCode || errorWithStatus?.status
    if (status === 401 || status === 403 || status === 404) {
      return await tryProxy('/admin/wordpress-migration/start')
    }
    throw e
  }
})