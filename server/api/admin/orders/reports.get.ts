import type { H3Event } from 'h3'
import { defineEventHandler, getQuery } from 'h3'
import { proxy } from '../../_utils/fetchProxy'

declare const useRuntimeConfig: () => { public: { goApiBase: string } }

export default defineEventHandler(async (event: H3Event) => {
  const config = useRuntimeConfig()
  const query = getQuery(event)

  // ایجاد URL با پارامترهای query
  const params = new URLSearchParams()
  if (query.period) params.append('period', query.period as string)
  if (query.year) params.append('year', query.year as string)
  if (query.month) params.append('month', query.month as string)

  const url = `${config.public.goApiBase}/api/admin/orders/reports?${params.toString()}`

  return await proxy(event, url, { method: 'GET' })
})
