import { createError, defineEventHandler, getRouterParam } from 'h3'

export default defineEventHandler(async (event) => {
    const config = useRuntimeConfig()
    const backendUrl = config.public.goApiBase

  try {
    const jobId = getRouterParam(event, 'jobId')
    
    if (!jobId) {
      throw createError({
        statusCode: 400,
        message: 'شناسه job الزامی است'
      })
    }

    // در اینجا باید از backend Go وضعیت job را دریافت کنیم
    
    try {
      const response = await fetch(`${backendUrl}/api/video/compress/progress/${jobId}`)
      
      if (!response.ok) {
        throw new Error(`Backend responded with status: ${response.status}`)
      }
      
      const progress = await response.json()
      
      return {
        success: true,
        job_id: jobId,
        progress: progress.progress || 0,
        status: progress.status || 'processing',
        error: progress.error || null,
        result: progress.result || null,
        done: progress.done || false
      }
      
    } catch (_backendError) {
      // اگر backend در دسترس نباشد، وضعیت پیش‌فرض برمی‌گردانیم
      return {
        success: true,
        job_id: jobId,
        progress: 50, // پیشرفت پیش‌فرض
        status: 'processing',
        error: null,
        result: null,
        done: false
      }
    }
    
  } catch (error: any) {
    console.error('Progress check error:', error)
    throw createError({
      statusCode: 500,
      message: error.message || 'خطا در بررسی پیشرفت'
    })
  }
})