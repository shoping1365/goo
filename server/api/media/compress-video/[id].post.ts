import { defineEventHandler, getQuery, getRouterParam, createError } from 'h3'

export default defineEventHandler(async (event) => {
    const config = useRuntimeConfig()
    const backendUrl = config.public.goApiBase

  try {
    const videoId = getRouterParam(event, 'id')
    const query = getQuery(event)
    const isSmartCompression = query.smart === 'true'
    const workerCount = parseInt(query.worker_count as string) || 4 // تعداد ورکر از تنظیمات
    const frameAnalysisMode = query.frame_analysis_mode as string || 'smart' // روش تحلیل فریم
    
    // بررسی وجود ویدیو
    if (!videoId) {
      throw createError({
        statusCode: 400,
        message: 'شناسه ویدیو الزامی است'
      })
    }

    // ایجاد job ID منحصر به فرد
    const jobId = `compression_${videoId}_${Date.now()}_${Math.random().toString(36).substr(2, 9)}`
    
    // شروع فرآیند فشرده‌سازی در background
    setTimeout(async () => {
      try {
        // فراخوانی سرویس فشرده‌سازی هوشمند یا معمولی
        const compressionType = isSmartCompression ? 'smart' : 'normal'
        
        // در اینجا باید به backend Go متصل شویم
        const response = await fetch(`${backendUrl}/api/video/compress`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({
            video_id: videoId,
            compression_type: compressionType,
            job_id: jobId,
            settings: {
              // تنظیمات فشرده‌سازی
              quality: isSmartCompression ? 'smart' : '720p',
              format: 'mp4',
              codec: 'h264',

              worker_count: workerCount, // تعداد ورکر از تنظیمات مرکزی
              frame_analysis_mode: frameAnalysisMode // روش تحلیل فریم برای ویدیوهای طولانی
            }
          })
        })
        
        if (!response.ok) {
          throw new Error(`Backend responded with status: ${response.status}`)
        }
        
        const result = await response.json()
        
        // ذخیره نتیجه در کش یا دیتابیس
        console.log(`Compression completed for video ${videoId}:`, result)
        
      } catch (error) {
        console.error(`Compression failed for video ${videoId}:`, error)
      }
    }, 100)

    return {
      success: true,
      job_id: jobId,
      message: isSmartCompression 
        ? 'فشرده‌سازی هوشمند شروع شد' 
        : 'فشرده‌سازی معمولی شروع شد',
      compression_type: isSmartCompression ? 'smart' : 'normal'
    }
    
  } catch (error: any) {
    console.error('Video compression error:', error)
    throw createError({
      statusCode: 500,
      message: error.message || 'خطا در فشرده‌سازی ویدیو'
    })
  }
})