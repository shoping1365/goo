import { exec } from 'child_process'
import { promisify } from 'util'

const _execAsync = promisify(exec)

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  return await $fetch(`${base}/api/admin/monitoring/overview`, { headers: event.node.req.headers as Record<string, string | string[] | undefined> });
})

// تابع پردازش اطلاعات حافظه
function _parseMemoryInfo(_physicalMemory: string, _virtualMemory: string, _memoryUsage: string) {
  // پردازش حافظه فیزیکی
  const physicalLines = physicalMemory.split('\n').filter(line => line.trim())
  const physicalData = physicalLines[1]?.split(',')[1] || '0'
  const totalPhysicalMB = Math.round(parseInt(physicalData) / (1024 * 1024))
  
  // پردازش حافظه مجازی
  const virtualLines = virtualMemory.split('\n').filter(line => line.trim())
  const virtualData = virtualLines[1]?.split(',') || []
  const allocatedVirtualMB = parseInt(virtualData[1] || '0')
  const usedVirtualMB = parseInt(virtualData[2] || '0')
  
  // پردازش استفاده حافظه
  const usageLines = memoryUsage.split('\n').filter(line => line.trim())
  const usageData = usageLines[1]?.split(',') || []
  const totalVisibleMB = parseInt(usageData[1] || '0')
  const freePhysicalMB = parseInt(usageData[2] || '0')
  const usedPhysicalMB = totalVisibleMB - freePhysicalMB
  
  return {
    physical: {
      total: totalPhysicalMB,
      used: usedPhysicalMB,
      free: freePhysicalMB,
      usagePercent: Math.round((usedPhysicalMB / totalVisibleMB) * 100)
    },
    virtual: {
      allocated: allocatedVirtualMB,
      used: usedVirtualMB,
      free: allocatedVirtualMB - usedVirtualMB,
      usagePercent: Math.round((usedVirtualMB / allocatedVirtualMB) * 100)
    },
    usage: {
      total: totalVisibleMB,
      used: usedPhysicalMB,
      free: freePhysicalMB,
      percent: Math.round((usedPhysicalMB / totalVisibleMB) * 100)
    }
  }
}