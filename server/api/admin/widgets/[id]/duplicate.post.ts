import { defineEventHandler, getRouterParam } from 'h3'
import { useRuntimeConfig } from '#imports'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase

  const { proxy } = await import('~/server/api/_utils/fetchProxy')
  const id = getRouterParam(event, 'id')

  // fetchProxy automatically forwards cookies and creates Authorization header
  return proxy(event, `${base}/api/admin/widgets/${id}/duplicate`, {
    method: 'POST'
  })
})