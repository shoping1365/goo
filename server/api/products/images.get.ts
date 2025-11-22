interface ApiResponse {
  success: boolean
  data?: unknown
  error?: string
}

export default defineEventHandler(async (_event) => {
  try {
    const config = useRuntimeConfig()
    const base = config.public.goApiBase
    
    // فراخوانی API بک‌اند برای دریافت تصاویر محصولات
    const response = await fetch(`${base}/api/products/images`).then(res => res.json())
    
    // اگر response خودش success و data دارد، آن را مستقیماً برگردان
    if (response && typeof response === 'object' && 'success' in response && 'data' in response) {
      return response as ApiResponse
    }
    
    // در غیر این صورت، آن را wrap کن
    return {
      success: true,
      data: response
    } as ApiResponse
  } catch (error) {
    console.error('Error fetching product images:', error)
    return {
      success: false,
      error: 'خطا در دریافت تصاویر محصولات',
      data: []
    } as ApiResponse
  }
})
