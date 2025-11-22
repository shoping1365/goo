import { exec } from 'child_process'
import { promisify } from 'util'
import { defineEventHandler, readBody } from 'h3'
import { useRuntimeConfig } from '#imports'

const _execAsync = promisify(exec)

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const body = await readBody(event)
  return await $fetch(`${base}/api/admin/system/pm2-update`, { method: 'POST', body, headers: event.node.req.headers as Record<string, string> })
})