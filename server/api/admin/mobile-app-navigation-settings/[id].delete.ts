import { useRuntimeConfig } from '#imports'
import type { H3Event } from 'h3'
import { createError, defineEventHandler, getRouterParam } from 'h3'
import { proxy } from '~/server/api/_utils/fetchProxy'

declare const requireAuth: (event: H3Event) => Promise<{ [key: string]: unknown } | null>

export default defineEventHandler(async (event: H3Event) => {
     // احراز هویت
     await requireAuth(event)

     // دریافت ID از پارامترها
     const id = getRouterParam(event, 'id')
     if (!id) {
          throw createError({
               statusCode: 400,
               message: 'ID is required'
          })
     }

     // ارسال درخواست به backend
     const config = useRuntimeConfig()
     const base = config.public.goApiBase
     return proxy(event, `${base}/api/admin/mobile-app-navigation-settings/${id}`, {
          method: 'DELETE'
     })
})

