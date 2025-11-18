export default defineEventHandler(async (event) => {
  try {
    const body = await readBody(event)
    
    // بررسی پارامترهای ورودی
    if (!body.src || !body.crop_width || !body.crop_height) {
      throw createError({
        statusCode: 400,
        message: 'پارامترهای src، crop_width و crop_height الزامی هستند'
      })
    }

    // فراخوانی API backend Go
    const config = useRuntimeConfig()
    const backendUrl = config.public.goApiBase
    const response = await $fetch(`${backendUrl}/api/media/image-crop`, {
      method: 'POST',
      body: {
        src: body.src,
        crop_width: body.crop_width,
        crop_height: body.crop_height,
        mode: body.mode || 'cover',
        device: body.device || 'mobile',
        quality: body.quality || 85
      }
    })

    return response
  } catch (error) {
    console.error('خطا در برش تصویر:', error)
    
    if (error.statusCode) {
      throw error
    }
    
    throw createError({
      statusCode: 500,
      message: 'خطا در برش تصویر'
    })
  }
})