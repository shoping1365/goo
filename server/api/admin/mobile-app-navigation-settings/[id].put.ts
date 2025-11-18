import { useRuntimeConfig } from '#imports'
import { createError, defineEventHandler, getRouterParam, readBody } from 'h3'
import { proxy } from '~/server/api/_utils/fetchProxy'

export default defineEventHandler(async (event) => {
     // احراز هویت (removed - requireAuth)

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
     const baseURL = config.public.goApiBase
     const body = await readBody(event)
     return proxy(event, `${baseURL}/api/admin/mobile-app-navigation-settings/${id}`, {
          method: 'PUT',
          body: JSON.stringify(body)
     })
})