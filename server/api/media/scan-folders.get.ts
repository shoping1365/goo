import { useRuntimeConfig } from '#imports'
import { createError, defineEventHandler, getQuery } from 'h3'

export default defineEventHandler(async (event) => {
  try {
    const query = getQuery(event)
    const category = query.category || 'products'
    
    if (category !== 'products') {
      return { success: true, data: [] }
    }

    const { public: { goApiBase } } = useRuntimeConfig()
    const backendUrl = goApiBase
    
    try {
      // ارسال درخواست به Go backend برای اسکن پوشه‌ها
      const response = await fetch(`${backendUrl}/api/media/scan-folders?category=${encodeURIComponent(category)}`, {
        headers: {
          'Content-Type': 'application/json',
          'Cookie': event.node.req.headers.cookie || ''
        }
      })

      if (!response.ok) {
        console.error('Backend responded with status:', response.status)
        throw new Error(`Backend error ${response.status}`)
      }

      const json = await response.json()
      return json
    } catch (backendError) {
      console.error('Backend error:', backendError)
      // اگر backend در دسترس نباشد، آرایه خالی برگردان
      return {
        success: true,
        data: []
      }
    }
  } catch (error: any) {
    console.error('Folder scan error:', error)
    throw createError({
      statusCode: 500,
      message: error.message || 'خطا در اسکن پوشه‌ها'
    })
  }
})