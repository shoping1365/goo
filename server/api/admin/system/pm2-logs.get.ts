import { exec } from 'child_process'
import { promisify } from 'util'

const _execAsync = promisify(exec)

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  return await $fetch(`${base}/api/admin/system/pm2-logs`, { headers: event.node.req.headers as Record<string, string | string[] | undefined> })
})