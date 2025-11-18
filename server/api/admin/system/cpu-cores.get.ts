import { exec } from 'child_process'
import { promisify } from 'util'

const execAsync = promisify(exec)

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  return await $fetch(`${base}/api/admin/monitoring/performance`, { headers: event.node.req.headers as any });
})

// تابع پردازش اطلاعات CPU
function parseCPUInfo(cpuInfo) {
  const lines = cpuInfo.split('\n').filter(line => line.trim())
  const data = lines[1]?.split(',') || []
  
  return {
    totalCores: parseInt(data[1] || '0'),
    logicalCores: parseInt(data[2] || '0')
  }
}

// تابع پردازش درصد استفاده CPU
function parseCPUUsage(cpuUsage) {
  const lines = cpuUsage.split('\n').filter(line => line.trim())
  const usage = lines[1]?.split(',')[1] || '0'
  
  return parseInt(usage)
}