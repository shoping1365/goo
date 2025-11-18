import { useRuntimeConfig } from '#imports'
import type { H3Event } from 'h3'
import { defineEventHandler } from 'h3'
import { proxy } from '~/server/api/_utils/fetchProxy'

export default defineEventHandler(async (event: H3Event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  return proxy(event, `${base}/api/admin/mobile-app-navigation-settings`)
})


