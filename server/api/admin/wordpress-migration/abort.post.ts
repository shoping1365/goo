import { defineEventHandler } from 'h3'
import { proxy } from '../../_utils/fetchProxy'

export default defineEventHandler(async (event) => {
     const config = useRuntimeConfig()
     const base = config.public.goApiBase

     const tryProxy = async (path: string) => proxy(event, `${base}${path}`, {
          method: 'POST'
     })

     try {
          return await tryProxy('/api/admin/wordpress-migration/abort')
     } catch (e: any) {
          const status = e?.statusCode || e?.status
          if (status === 401 || status === 403 || status === 404) {
               return await tryProxy('/admin/wordpress-migration/abort')
          }
          throw e
     }
})