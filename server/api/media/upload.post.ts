import { createError, getCookie, type H3Event } from 'h3'

interface UploadResponse {
  success: boolean
  files?: unknown[]
  error?: string
}

export default defineEventHandler(async (event: H3Event): Promise<UploadResponse> => {
  try {
    const formData = await readFormData(event)
    const authToken = getCookie(event, 'access_token') || getCookie(event, 'auth-token')

    // اضافه کردن فیلدهای اجباری اگر وجود ندارند
    if (!formData.has('category')) {
      formData.append('category', 'banners')
    }
    if (!formData.has('purpose')) {
      formData.append('purpose', 'general')
    }

    // اضافه کردن user_id اگر در context موجود است
    const userID = getCookie(event, 'user_id')
    if (userID && !formData.has('uploaded_by')) {
      formData.append('uploaded_by', userID)
    }

    const config = useRuntimeConfig()
    const base = config.public.goApiBase
    
    // استفاده از fetch native به جای $fetch برای جلوگیری از خطای type instantiation
    const response = await fetch(`${base}/api/media/upload`, {
      method: 'POST',
      body: formData,
      headers: {
        'Authorization': authToken ? `Bearer ${authToken}` : ''
      }
    })

    if (!response.ok) {
      throw new Error(`Upload failed: ${response.statusText}`)
    }

    return await response.json() as UploadResponse
  } catch (error: unknown) {
    const err = error as Error & { status?: number }
    console.error('Media upload error:', err)
    throw createError({
      statusCode: err.status || 500,
      message: err.message || 'خطا در آپلود فایل'
    })
  }
})