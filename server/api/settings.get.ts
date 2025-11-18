import { defineEventHandler, getQuery } from 'h3'
import { proxy } from './_utils/fetchProxy'

export default defineEventHandler(async (event) => {
     const config = useRuntimeConfig()
     const q = getQuery(event)
     const category = (q.category || '').toString().trim()
     if (category) {
          return await proxy(event, `${config.public.goApiBase}/api/admin/settings/category/${encodeURIComponent(category)}`)
     }
     return await proxy(event, `${config.public.goApiBase}/api/admin/settings`)
})



