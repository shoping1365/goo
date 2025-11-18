import { defineEventHandler, getRouterParam, createError } from 'h3'
import { proxy } from '../../../_utils/fetchProxy'
import { getCookieValue } from '../../../_utils/cookieHelper'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const category = getRouterParam(event, 'category')
  const token = getCookieValue(event, 'access_token') || getCookieValue(event, 'auth-token')
  
  if (!token) {
    throw createError({ statusCode: 401, message: 'برای ادامه باید وارد شوید.' })
  }
  
  try {
    // دریافت تنظیمات از backend
    const response = await proxy(event, `${config.public.goApiBase}/api/admin/settings/category/${category}`, {
      headers: { Authorization: `Bearer ${token}` }
    })
    
    return response
  } catch (error) {
    // در صورت خطا یا عدم وجود backend، تنظیمات mock برمی‌گرداند
    const mockSettings: Record<string, any> = {
      auth: [
        { key: 'auth.session_timeout', value: '3600', category: 'auth', type: 'string' },
        { key: 'auth.max_login_attempts', value: '5', category: 'auth', type: 'string' },
        { key: 'auth.lockout_duration', value: '900', category: 'auth', type: 'string' },
        { key: 'auth.require_email_verification', value: 'true', category: 'auth', type: 'boolean' },
        { key: 'auth.password_min_length', value: '8', category: 'auth', type: 'string' },
        { key: 'auth.enable_two_factor', value: 'false', category: 'auth', type: 'boolean' }
      ],
      'social-media': [
        { key: 'social-media.instagram', value: '', category: 'social-media', type: 'string' },
        { key: 'social-media.instagram_enabled', value: 'true', category: 'social-media', type: 'boolean' },
        { key: 'social-media.telegram', value: '', category: 'social-media', type: 'string' },
        { key: 'social-media.telegram_enabled', value: 'true', category: 'social-media', type: 'boolean' },
        { key: 'social-media.whatsapp', value: '', category: 'social-media', type: 'string' },
        { key: 'social-media.whatsapp_enabled', value: 'true', category: 'social-media', type: 'boolean' },
        { key: 'social-media.twitter', value: '', category: 'social-media', type: 'string' },
        { key: 'social-media.twitter_enabled', value: 'true', category: 'social-media', type: 'boolean' },
        { key: 'social-media.linkedin', value: '', category: 'social-media', type: 'string' },
        { key: 'social-media.linkedin_enabled', value: 'true', category: 'social-media', type: 'boolean' },
        { key: 'social-media.facebook', value: '', category: 'social-media', type: 'string' },
        { key: 'social-media.facebook_enabled', value: 'true', category: 'social-media', type: 'boolean' },
        { key: 'social-media.youtube', value: '', category: 'social-media', type: 'string' },
        { key: 'social-media.youtube_enabled', value: 'true', category: 'social-media', type: 'boolean' },
        { key: 'social-media.aparat', value: '', category: 'social-media', type: 'string' },
        { key: 'social-media.aparat_enabled', value: 'true', category: 'social-media', type: 'boolean' },
        { key: 'social-media.rubika', value: '', category: 'social-media', type: 'string' },
        { key: 'social-media.rubika_enabled', value: 'true', category: 'social-media', type: 'boolean' },
        { key: 'social-media.eitaa', value: '', category: 'social-media', type: 'string' },
        { key: 'social-media.eitaa_enabled', value: 'true', category: 'social-media', type: 'boolean' },
        { key: 'social-media.bale', value: '', category: 'social-media', type: 'string' },
        { key: 'social-media.bale_enabled', value: 'true', category: 'social-media', type: 'boolean' },
        { key: 'social-media.tiktok', value: '', category: 'social-media', type: 'string' },
        { key: 'social-media.tiktok_enabled', value: 'true', category: 'social-media', type: 'boolean' },
        { key: 'social-media.pinterest', value: '', category: 'social-media', type: 'string' },
        { key: 'social-media.pinterest_enabled', value: 'true', category: 'social-media', type: 'boolean' },
        { key: 'social-media.discord', value: '', category: 'social-media', type: 'string' },
        { key: 'social-media.discord_enabled', value: 'true', category: 'social-media', type: 'boolean' },
        { key: 'social-media.github', value: '', category: 'social-media', type: 'string' },
        { key: 'social-media.github_enabled', value: 'true', category: 'social-media', type: 'boolean' },
        { key: 'social-media.custom_links', value: '[]', category: 'social-media', type: 'json' }
      ]
    }
    
    return {
      success: true,
      data: mockSettings[category || ''] || []
    }
  }
})