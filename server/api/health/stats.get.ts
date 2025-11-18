import type { H3Event } from 'h3'
import { createError, defineEventHandler } from 'h3'
import { getHeader } from '../_utils/getHeader'

declare const useRuntimeConfig: () => { public: { goApiBase: string } }
declare const $fetch: typeof fetch

interface MonitorHealth {
  health?: {
    uptime?: string
    active_sessions?: number
    total_connections?: number
  }
}

export default defineEventHandler(async (event: H3Event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase

  try {
    const [, monitor] = await Promise.all([
      $fetch(`${base}/health`, { method: 'GET' }),
      $fetch(`${base}/api/admin/monitoring/health`, {
        method: 'GET',
        credentials: 'include',
        headers: { cookie: getHeader(event, 'cookie') || '' }
      }).catch(() => null)
    ])

    const monitorData = monitor as MonitorHealth | null
    const data = monitorData?.health || null

    return {
      status: 'ok',
      time: new Date().toISOString(),
      data: {
        uptime: data?.uptime || '-',
        database_size: '-',
        active_connections: data?.active_sessions || data?.total_connections || '-',
      }
    }
  } catch (err: unknown) {
    console.error('Error fetching health stats:', err)
    throw createError({ statusCode: 500, message: 'health stats failed' })
  }
})