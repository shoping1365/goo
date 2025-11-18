import { defineEventHandler, readBody, createError } from 'h3'

export default defineEventHandler(async (event) => {
  try {
    const body = await readBody(event)

    if (!body.cart_item_id) {
      throw createError({
        statusCode: 400,
        message: 'شناسه آیتم سبد خرید الزامی است'
      })
    }

    // شبیه‌سازی حذف از سبد خرید
    await new Promise(resolve => setTimeout(resolve, 200))

    // آیتم از سبد خرید حذف شد

    return {
      success: true,
      message: 'محصول با موفقیت از سبد خرید حذف شد'
    }

  } catch (error) {
    // خطا در حذف از سبد خرید

    throw createError({
      statusCode: error.statusCode || 500,
      message: error.statusMessage || 'خطا در حذف از سبد خرید'
    })
  }
})