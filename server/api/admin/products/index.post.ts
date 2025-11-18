import { defineEventHandler, readBody } from 'h3'
import { useRuntimeConfig } from '#imports'
import { proxy } from '~/server/api/_utils/fetchProxy'

export default defineEventHandler(async (event) => {
     const config = useRuntimeConfig()
     const base = config.public.goApiBase
     const body = await readBody(event)

     return proxy(event, `${base}/api/products`, {
          method: 'POST',
          body,
          headers: {
               'Content-Type': 'application/json'
          }
     })
})