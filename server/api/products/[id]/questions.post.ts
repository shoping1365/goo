export default defineEventHandler(async (event) => {
  const { getRouterParam, readBody, createError } = await import('h3')
  const { fetchGo } = await import('../../_utils/fetchGo')
  try {
    const productId = getRouterParam(event, 'id')
    const body = await readBody(event)
    
    if (!productId) {
      throw createError({
        statusCode: 400,
        message: 'شناسه محصول الزامی است'
      })
    }

    if (!body.question || body.question.trim() === '') {
      throw createError({
        statusCode: 400,
        message: 'متن پرسش الزامی است'
      })
    }

    // ارسال اطلاعات دستگاه و IP
    const deviceInfo = {
      ua: event.node.req.headers['user-agent'] || '',
      lang: event.node.req.headers['accept-language'] || '',
    }
    
    const response = await fetchGo(event, `/api/questions/product/${productId}`, {
      method: 'POST',
      body: {
        question: body.question.trim(),
        is_anonymous: body.is_anonymous || false,
        name: body.customer_name,
        phone: body.customer_mobile,
        device_info: deviceInfo
      }
    })

    return response
  } catch (error) {
    console.error('Error creating product question:', error)
    throw createError({
      statusCode: 500,
      message: 'خطا در ثبت پرسش'
    })
  }
}) 