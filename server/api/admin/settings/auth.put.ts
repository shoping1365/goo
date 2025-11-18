import { fetchGo } from '../../_utils/fetchGo'

export default defineEventHandler(async (event) => {
  try {
    const body = await readBody(event)
    
    // تبدیل تنظیمات به فرمت مناسب برای ذخیره
    const settingsToUpdate = [
      { key: 'mobile_auth_enabled', value: body.mobile_auth_enabled?.toString() || 'true', category: 'auth', description: 'فعال‌سازی احراز هویت با موبایل' },
      { key: 'otp_length', value: body.otp_length?.toString() || '5', category: 'auth', description: 'طول کد تایید' },
      { key: 'otp_expiry_minutes', value: body.otp_expiry_minutes?.toString() || '5', category: 'auth', description: 'مدت اعتبار کد تایید' },
      { key: 'max_otp_attempts', value: body.max_otp_attempts?.toString() || '3', category: 'auth', description: 'حداکثر تلاش برای کد تایید' },
      { key: 'jwt_expiry_hours', value: body.jwt_expiry_hours?.toString() || '24', category: 'auth', description: 'مدت اعتبار توکن JWT' },
      { key: 'refresh_token_expiry_days', value: body.refresh_token_expiry_days?.toString() || '30', category: 'auth', description: 'مدت اعتبار توکن تازه‌سازی' },
      { key: 'jwt_secret', value: body.jwt_secret || '', category: 'auth', description: 'کلید مخفی JWT' },
      { key: 'username_auth_enabled', value: body.username_auth_enabled?.toString() || 'true', category: 'auth', description: 'فعال‌سازی ورود با یوزرنیم' },
      { key: 'min_password_length', value: body.min_password_length?.toString() || '8', category: 'auth', description: 'حداقل طول پسورد' },
      { key: 'max_login_attempts', value: body.max_login_attempts?.toString() || '5', category: 'auth', description: 'حداکثر تلاش ورود' },
      { key: 'account_lockout_minutes', value: body.account_lockout_minutes?.toString() || '15', category: 'auth', description: 'مدت قفل حساب' },
      { key: 'default_user_role', value: body.default_user_role || 'user', category: 'auth', description: 'نقش پیش‌فرض کاربران جدید' },
      { key: 'email_verification_enabled', value: body.email_verification_enabled?.toString() || 'false', category: 'auth', description: 'فعال‌سازی تایید ایمیل' }
    ]
    
    // بروزرسانی تنظیمات در Go backend (اجتناب از فراخوانی نسبی /api/... که باعث IPC socket می‌شود)
    for (const setting of settingsToUpdate) {
      await fetchGo(event, `/api/admin/settings/${setting.key}`, {
        method: 'PUT',
        body: {
          value: setting.value,
          description: setting.description,
          category: setting.category
        }
      })
    }
    
    return {
      success: true,
      message: 'تنظیمات احراز هویت با موفقیت بروزرسانی شد'
    }
  } catch (error) {
    console.error('خطا در بروزرسانی تنظیمات احراز هویت:', error)
    throw createError({
      statusCode: 500,
      message: 'خطا در بروزرسانی تنظیمات احراز هویت'
    })
  }
}) 