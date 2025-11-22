import { exec } from 'child_process'
import { promisify } from 'util'

const _execAsync = promisify(exec)

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  return await $fetch(`${base}/api/admin/monitoring/performance`, { headers: event.node.req.headers as Record<string, string | string[] | undefined> });
})

// تابع پردازش اطلاعات CPU
function _parseCPUInfo(_cpuInfo: string) {
  const lines = cpuInfo.split('\n').filter(line => line.trim())
  const data = lines[1]?.split(',') || []
  
  return {
    totalCores: parseInt(data[1] || '0'),
    logicalCores: parseInt(data[2] || '0')
  }
}

// تابع پردازش درصد استفاده CPU
function _parseCPUUsage(_cpuUsage: string) {
  const lines = cpuUsage.split('\n').filter(line => line.trim())
  const usage = lines[1]?.split(',')[1] || '0'
  
  return parseInt(usage)
}