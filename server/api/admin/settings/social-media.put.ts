import { fetchGo } from '../../_utils/fetchGo'

// Test deployment - trigger rebuild
export default defineEventHandler(async (event) => {
  try {
    const body = await readBody(event)
    console.log('📥 Received body:', body)
    
    // Validation
    if (!body || typeof body !== 'object') {
      throw createError({
        statusCode: 400,
        statusMessage: 'بدنه درخواست نامعتبر است'
      })
    }
    
    // تبدیل تنظیمات به فرمت مناسب برای ذخیره
    const platforms = [
      { key: 'instagram', description: 'آدرس اینستاگرام', enabledDescription: 'فعال‌سازی اینستاگرام' },
      { key: 'telegram', description: 'آدرس تلگرام', enabledDescription: 'فعال‌سازی تلگرام' },
      { key: 'whatsapp', description: 'آدرس واتساپ', enabledDescription: 'فعال‌سازی واتساپ' },
      { key: 'twitter', description: 'آدرس توییتر', enabledDescription: 'فعال‌سازی توییتر' },
      { key: 'linkedin', description: 'آدرس لینکدین', enabledDescription: 'فعال‌سازی لینکدین' },
      { key: 'facebook', description: 'آدرس فیسبوک', enabledDescription: 'فعال‌سازی فیسبوک' },
      { key: 'youtube', description: 'آدرس یوتیوب', enabledDescription: 'فعال‌سازی یوتیوب' },
      { key: 'aparat', description: 'آدرس آپارات', enabledDescription: 'فعال‌سازی آپارات' },
      { key: 'rubika', description: 'آدرس روبیکا', enabledDescription: 'فعال‌سازی روبیکا' },
      { key: 'eitaa', description: 'آدرس ایتا', enabledDescription: 'فعال‌سازی ایتا' },
      { key: 'bale', description: 'آدرس بله', enabledDescription: 'فعال‌سازی بله' },
      { key: 'tiktok', description: 'آدرس تیک‌تاک', enabledDescription: 'فعال‌سازی تیک‌تاک' },
      { key: 'pinterest', description: 'آدرس پینترست', enabledDescription: 'فعال‌سازی پینترست' },
      { key: 'discord', description: 'آدرس دیسکورد', enabledDescription: 'فعال‌سازی دیسکورد' },
      { key: 'github', description: 'آدرس گیت‌هاب', enabledDescription: 'فعال‌سازی گیت‌هاب' }
    ]

    const settingsToUpdate = platforms.flatMap(({ key, description, enabledDescription }) => {
      const value = body[key] || ''
      const enabledValue = body[`${key}_enabled`]?.toString() || 'false'

      return [
        {
          key: `social-media.${key}`,
          value,
          category: 'social-media',
          description,
          type: 'string',
          is_public: true
        },
        {
          key: `social-media.${key}_enabled`,
          value: enabledValue,
          category: 'social-media',
          description: enabledDescription,
          type: 'boolean',
          is_public: true
        }
      ]
    })

    settingsToUpdate.push({
      key: 'social-media.custom_links',
      value: JSON.stringify(Array.isArray(body.customLinks) ? body.customLinks : []),
      category: 'social-media',
      description: 'لینک‌های سفارشی شبکه‌های اجتماعی',
      type: 'json',
      is_public: true
    })
    
    // بروزرسانی تنظیمات در Go backend (سازگار با نسخه‌های قدیمی که bulk PUT ندارند)
    console.log('📤 Updating', settingsToUpdate.length, 'social media settings individually')
    const results: Array<{ key: string; success: boolean; error?: any }> = []

    for (const setting of settingsToUpdate) {
      const payload: Record<string, any> = {
        value: setting.value,
        description: setting.description,
        category: setting.category
      }

      if (setting.type) payload.type = setting.type
      if (typeof setting.is_public === 'boolean') payload.is_public = setting.is_public

      try {
        const endpoint = `/api/admin/settings/${encodeURIComponent(setting.key)}`
        console.log('➡️ Updating setting:', setting.key, 'payload:', payload)
        const response = await fetchGo(event, endpoint, {
          method: 'PUT',
          body: payload
        })
        console.log('✅ Setting updated:', setting.key, 'response:', response)
        results.push({ key: setting.key, success: true })
      } catch (error: any) {
        console.error('❌ Failed to update setting:', setting.key, error)
        console.error('❌ Error details:', {
          key: setting.key,
          statusCode: error?.statusCode,
          statusMessage: error?.statusMessage,
          data: error?.data,
          message: error?.message
        })
        results.push({ key: setting.key, success: false, error })
      }
    }

    const failed = results.filter(result => !result.success)
    if (failed.length) {
      throw createError({
        statusCode: 500,
        statusMessage: 'برخی از تنظیمات شبکه‌های اجتماعی ذخیره نشد',
        data: failed.map(item => ({ key: item.key, error: item.error?.data || item.error?.message }))
      })
    }

    return {
      success: true,
      message: 'تنظیمات شبکه‌های اجتماعی با موفقیت ذخیره شد',
      data: { updated: results.length }
    }
    
  } catch (error: any) {
    console.error('❌ خطا در بروزرسانی تنظیمات شبکه‌های اجتماعی:', error)
    console.error('❌ Error stack:', error?.stack)
    
    throw createError({
      statusCode: error?.statusCode || 500,
      statusMessage: error?.statusMessage || 'خطا در بروزرسانی تنظیمات شبکه‌های اجتماعی',
      data: {
        message: error?.message,
        details: error?.data
      }
    })
  }
})
