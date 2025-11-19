import { createError, defineEventHandler, readBody } from 'h3'
import { getCookieValue } from '../_utils/cookieHelper'
import { getHeader } from '../_utils/getHeader'

// API ایجاد سفارش جدید
export default defineEventHandler(async (event) => {
  try {
    // خواندن بدنه درخواست سفارش
    const body = await readBody(event)
    
    // بررسی احراز هویت کاربر
    const userId = event.context?.user?.id
    if (!userId) {
      throw createError({
        statusCode: 401,
        message: 'کاربر احراز هویت نشده است'
      })
    }

    // دریافت اتصال پایگاه داده
    const { getDatabase } = await import('../_utils/database')

    // استخراج فیلدهای اصلی سفارش از بدنه درخواست
    const { 
      notes, 
      paymentMethod,
      items,
      total
    } = body

    console.log('ایجاد سفارش جدید:', { notes, paymentMethod, items, total })

    if (!items || !Array.isArray(items) || items.length === 0) {
      throw createError({
        statusCode: 400,
        message: 'هیچ آیتمی برای سفارش ارسال نشده است'
      })
    }

    // محاسبه مجموع مبلغ سفارش بر اساس قیمت محصولات
    const db = await getDatabase()
    let calculatedTotal = 0
    
    for (const item of items) {
      const product = await db.query(
        'SELECT id, name, price, stock_quantity FROM products WHERE id = $1',
        [item.product_id]
      )
      
      if (!product || product.length === 0) {
        throw createError({
          statusCode: 400,
          message: `محصول ${item.product?.name || 'نامشخص'} یافت نشد`
        })
      }
      
      if ((product[0] as any).stock_quantity < item.quantity) {
        throw createError({
          statusCode: 400,
          message: `موجودی محصول ${(product[0] as any).name} کافی نیست`
        })
      }
      
      // اضافه کردن مبلغ این آیتم به مجموع کل
      calculatedTotal += (product[0] as any).price * item.quantity
    }
    
    // در حال حاضر تخفیف و هزینه اضافی در نظر گرفته نشده است
    const finalTotal = calculatedTotal

    // دریافت آدرس کاربر از profile_data
    console.log('دریافت آدرس از profile_data برای کاربر:', userId)
    const userResult = await db.query(
      'SELECT profile_data FROM users WHERE id = $1',
      [userId]
    )
    
    console.log('نتیجه دریافت profile_data کاربر:', userResult)
    
    if (!userResult || userResult.length === 0) {
      throw createError({
        statusCode: 400,
        message: 'کاربر یافت نشد'
      })
    }
    
    const userData = userResult[0] as any
    const profileData = userData.profile_data
    
    if (!profileData || !profileData.selected_address) {
      throw createError({
        statusCode: 400,
        message: 'آدرس پیش‌فرض برای کاربر تنظیم نشده است. لطفاً ابتدا آدرس خود را تکمیل کنید'
      })
    }
    
    const address = {
      id: 1, // شناسه نمادین آدرس بر اساس profile_data
      full_address: profileData.selected_address,
      recipient_name: profileData.first_name + ' ' + profileData.last_name,
      recipient_mobile: event.context?.user?.mobile || '',
      postal_code: profileData.postal_code || ''
    }

    // تولید شماره سفارش یکتا
    const random = Math.floor(Math.random() * 10000).toString().padStart(4, '0')
    const orderNumber = `ORD-${random}`

    // خواندن IP و User Agent کاربر
    const clientIP = getClientIP(event) || 'unknown'
    const userAgent = getHeader(event, 'user-agent') || 'unknown'

    // ایجاد رکورد سفارش در پایگاه داده
    const orderResult = await db.query(
      `INSERT INTO orders (
        user_id, 
        shipping_address_id, 
        payment_method, 
        total_amount, 
        final_amount,
        status, 
        payment_status,
        notes,
        order_number,
        created_at,
        customer_ip,
        user_agent
      ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id`,
      [
        userId,
        address.id,
        paymentMethod,
        finalTotal,
        finalTotal, // final_amount
        'pending',
        'pending',
        notes || null,
        orderNumber,
        new Date().toISOString(), // created_at
        clientIP, // customer_ip
        userAgent // user_agent
      ]
    )

    const orderId = (orderResult[0] as any).id

    // ثبت آیتم‌های سفارش در جدول order_items
    for (const item of items) {
      // دریافت اطلاعات محصول برای هر آیتم سفارش
      const productResult = await db.query(
        'SELECT id, name, image, sku, price FROM products WHERE id = $1',
        [item.product_id]
      )
      
      const product = productResult[0] as any
      
      await db.query(
        `INSERT INTO order_items (
          order_id, 
          product_id, 
          product_name,
          product_image,
          product_sku,
          quantity, 
          unit_price, 
          total_price,
          final_price
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
        [
          orderId,
          item.product_id,
          product.name,
          product.image,
          product.sku,
          item.quantity,
          product.price,
          product.price * item.quantity,
          product.price * item.quantity
        ]
      )

      // کاهش موجودی انبار محصول
      await db.query(
        'UPDATE products SET stock_quantity = stock_quantity - $1 WHERE id = $2',
        [item.quantity, item.product_id]
      )
    }

    // خالی کردن سبد خرید کاربر (در دیتابیس و session)
    try {
      await db.query('DELETE FROM cart_items WHERE cart_id IN (SELECT id FROM carts WHERE user_id = $1)', [userId])
      
      // حذف سبد خرید از session storage
      const { clearCart } = await import('../cart/data')
      const sessionId = getCookieValue(event, 'session_id')
      if (sessionId) {
        clearCart(sessionId)
      }
    } catch (cartError) {
      console.error('خطا در خالی کردن سبد خرید:', cartError)
      // در صورت خطا در خالی کردن سبد خرید، ادامه فرایند سفارش را متوقف نکن
    }

    // در صورت پرداخت آنلاین، ثبت تراکنش پرداخت
    if (paymentMethod === 'online') {
      // در صورت نبود درگاه، ایجاد یک درگاه پیش‌فرض
      let gatewayId = 1
      
      try {
        // بررسی وجود درگاه پرداخت
        const gatewayResult = await db.query(
          'SELECT id FROM payment_gateways WHERE id = $1',
          [gatewayId]
        )
        
        if (!gatewayResult || gatewayResult.length === 0) {
          // ایجاد رکورد درگاه پرداخت پیش‌فرض
          const newGatewayResult = await db.query(
            `INSERT INTO payment_gateways (
              name, 
              english_name, 
              type, 
              status, 
              color, 
              fee, 
              min_amount, 
              max_amount, 
              priority, 
              is_test_mode, 
              description, 
              created_at
            ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, NOW()) RETURNING id`,
            [
              'درگاه پیش‌فرض',
              'default_gateway',
              'online',
              'active',
              'bg-blue-500',
              0,
              0,
              999999999,
              1,
              true,
              'درگاه پیش‌فرض سیستم پرداخت'
            ]
          )
          gatewayId = (newGatewayResult[0] as any).id
        }
      } catch (gatewayError) {
        console.error('خطا در دریافت/ایجاد درگاه پرداخت:', gatewayError)
        // در صورت خطا تلاش می‌کنیم از درگاه پیش‌فرض استفاده کنیم (id = 1)
      }

      // ایجاد رکورد تراکنش پرداخت برای سفارش
      const paymentResult = await db.query(
        `INSERT INTO payment_transactions (
          gateway_id, 
          order_id, 
          amount, 
          payment_method, 
          status, 
          created_at
        ) VALUES ($1, $2, $3, $4, $5, NOW()) RETURNING id`,
        [gatewayId, orderId, finalTotal, 'online', 'pending']
      )

      return {
        success: true,
        message: 'سفارش با موفقیت ثبت شد. در حال انتقال به درگاه پرداخت',
        data: {
          orderId,
          orderNumber,
          paymentTransactionId: (paymentResult[0] as any).id,
          redirectToPayment: true
        }
      }
    }

    // در صورت پرداخت غیر آنلاین، فقط سفارش ثبت می‌شود
    return {
      success: true,
      message: 'سفارش با موفقیت ثبت شد. پرداخت در محل یا روش آفلاین انتخاب شده است',
      data: {
        orderId,
        orderNumber,
        redirectToPayment: false
      }
    }

  } catch (error: any) {
    console.error('خطا در ایجاد سفارش:', error)
    
    if (error.statusCode) {
      throw error
    }
    
    throw createError({
      statusCode: 500,
      message: 'خطای داخلی در ایجاد سفارش'
    })
  }
})

// دریافت کاربر از session (در حال حاضر غیرفعال)
async function getUserFromSession(event: any) {
  try {
    // قبلاً از requireAuth در _utils/auth.ts استفاده می‌شد
    // بررسی‌های احراز هویت در این نسخه غیرفعال شده است
    
    return event.context?.user || null
  } catch (error: any) {
    console.error('خطا در دریافت کاربر از session:', error)
    return null
  }
}

// دریافت IP کاربر از هدرهای درخواست
function getClientIP(event: any): string {
  return event.node.req.headers['x-forwarded-for'] || 
         event.node.req.headers['x-real-ip'] || 
         event.node.req.connection?.remoteAddress || 
         'unknown'
} 
