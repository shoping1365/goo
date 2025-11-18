import { defineEventHandler } from 'h3'

export default defineEventHandler(async (event) => {
  try {
    // بررسی احراز هویت کاربر
    const userId = event.context?.user?.id
    
    if (!userId) {
      throw createError({
        statusCode: 401,
        message: 'کاربر احراز هویت نشده است'
      })
    }

    // دریافت اتصال دیتابیس
    const { getDatabase } = await import('../_utils/database')
    const db = await getDatabase()

    // دریافت آدرس‌های کاربر از جدول user_addresses
    const addresses = await db.query(
      'SELECT * FROM user_addresses WHERE user_id = $1 ORDER BY is_default DESC, created_at DESC',
      [userId]
    )

    // اگر آدرسی در user_addresses نیست، از profile_data استفاده کن
    if (!addresses || addresses.length === 0) {
      const userResult = await db.query(
        'SELECT profile_data FROM users WHERE id = $1',
        [userId]
      )
      
      if (userResult && userResult.length > 0) {
        const userData = userResult[0]
        const profileData = userData.profile_data
        
        if (profileData && profileData.selected_address) {
          // ایجاد آدرس مجازی از profile_data
          const virtualAddress = {
            id: 999, // شناسه مجازی
            user_id: userId,
            full_address: profileData.selected_address,
            street: profileData.selected_address,
            postal_code: profileData.postal_code || '',
            recipient_name: (profileData.first_name || '') + ' ' + (profileData.last_name || ''),
            recipient_mobile: event.context?.user?.mobile || '',
            is_default: true,
            created_at: new Date().toISOString(),
            updated_at: new Date().toISOString()
          }
          
          return {
            success: true,
            data: [virtualAddress]
          }
        }
      }
      
      return {
        success: true,
        data: []
      }
    }

    // تبدیل آدرس‌های دیتابیس به فرمت مورد انتظار
    const formattedAddresses = addresses.map(addr => ({
      id: addr.id,
      user_id: addr.user_id,
      full_address: addr.street || addr.full_address,
      street: addr.street,
      postal_code: addr.postal_code,
      recipient_name: addr.recipient_name,
      recipient_mobile: addr.recipient_mobile,
      phone: addr.phone,
      province: addr.province,
      city: addr.city,
      province_id: addr.province_id,
      city_id: addr.city_id,
      is_default: addr.is_default,
      created_at: addr.created_at,
      updated_at: addr.updated_at
    }))

    return {
      success: true,
      data: formattedAddresses
    }

  } catch (error) {
    console.error('خطا در دریافت آدرس‌های کاربر:', error)
    
    if (error.statusCode) {
      throw error
    }
    
    throw createError({
      statusCode: 500,
      message: 'خطا در دریافت آدرس‌های کاربر'
    })
  }
}) 