import type { H3Event } from 'h3'
import { createError, defineEventHandler, readBody } from 'h3'
import { proxy } from '~/server/api/_utils/fetchProxy'

declare const useRuntimeConfig: () => { public: { goApiBase: string } }

export default defineEventHandler(async (event: H3Event) => {
     try {
          const config = useRuntimeConfig()
          const base = config.public.goApiBase

          const body = await readBody(event)

          // fetchProxy automatically forwards cookies and creates Authorization header
          return proxy(event, `${base}/api/admin/widgets`, {
               method: 'POST',
               body: JSON.stringify(body)
          })
     } catch (error: unknown) {
          console.error('خطا در ایجاد ویجت:', error)
          const errorWithStatus = error as { statusCode?: number; message?: string; statusMessage?: string }
          throw createError({
               statusCode: errorWithStatus.statusCode || 500,
               message: errorWithStatus.message || errorWithStatus.statusMessage || 'خطا در ایجاد ویجت'
          })
     }
})
