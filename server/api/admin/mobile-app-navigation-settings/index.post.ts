import { useRuntimeConfig } from '#imports'
import type { H3Event } from 'h3'
import { defineEventHandler, readBody } from 'h3'
import { proxy } from '~/server/api/_utils/fetchProxy'

export default defineEventHandler(async (event: H3Event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const body = await readBody(event)
  return proxy(event, `${base}/api/admin/mobile-app-navigation-settings`, {
    method: 'POST',
    body: JSON.stringify(body)
  })
})


