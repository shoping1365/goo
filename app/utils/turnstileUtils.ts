/**
 * اعتبارسنجی توکن Turnstile با Cloudflare API
 * @param token - توکن دریافتی از Turnstile widget
 * @param secretKey - کلید مخفی Turnstile
 * @returns نتیجه اعتبارسنجی
 */
export async function verifyTurnstileToken(token: string, secretKey: string): Promise<{
  success: boolean
  message: string
  data?: any
  error?: string
}> {
  try {
    if (!token) {
      return {
        success: false,
        message: 'توکن Turnstile ارائه نشده است'
      }
    }

    if (!secretKey) {
      return {
        success: false,
        message: 'کلید مخفی Turnstile تنظیم نشده است'
      }
    }

    // ارسال درخواست به Cloudflare برای اعتبارسنجی
    const response = await fetch('https://challenges.cloudflare.com/turnstile/v0/siteverify', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
      },
      body: new URLSearchParams({
        secret: secretKey,
        response: token,
      }),
    })

    const result = await response.json()

    if (result.success) {
      return {
        success: true,
        message: 'اعتبارسنجی Turnstile موفق بود',
        data: result
      }
    } else {
      return {
        success: false,
        message: 'اعتبارسنجی Turnstile ناموفق بود',
        error: result['error-codes']?.join(', ') || 'خطای نامشخص'
      }
    }
  } catch (error) {
    return {
      success: false,
      message: 'خطا در ارتباط با Cloudflare',
      error: error.message
    }
  }
}

/**
 * دریافت کلیدهای Turnstile از متغیرهای محیطی
 */
export function getTurnstileKeys() {
  const siteKey = process.env.NUXT_PUBLIC_TURNSTILE_SITE_KEY || '1x00000000000000000000AA'
  const secretKey = process.env.NUXT_TURNSTILE_SECRET_KEY || ''
  
  return { siteKey, secretKey }
} 