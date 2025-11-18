export type WsCloseInfo = {
  code: number
  reason: string
  wasClean: boolean
  userMessage: string
}

const codeToTitle: Record<number, string> = {
  1000: 'اتصال به‌طور عادی بسته شد',
  1001: 'سرور در حال قطع اتصال است',
  1002: 'خطای پروتکل وب‌سوکت',
  1003: 'داده نامعتبر دریافت شد',
  1005: 'کد بسته شدن دریافت نشد',
  1006: 'اتصال به‌صورت غیرمنتظره قطع شد',
  1007: 'فرمت پیام نامعتبر بود',
  1008: 'نقض خط‌مشی',
  1009: 'پیام بسیار بزرگ بود',
  1010: 'سرور از پروتکل‌های درخواستی پشتیبانی نکرد',
  1011: 'خطای داخلی سرور',
  1012: 'سرور در حال راه‌اندازی مجدد است',
  1013: 'سرور موقتاً شلوغ است',
  1015: 'خطای TLS'
}

export function parseWsClose(event: CloseEvent): WsCloseInfo {
  const code = event.code
  const base = codeToTitle[code] || 'اتصال قطع شد'
  let hint = ''

  switch (code) {
    case 1000:
      hint = 'اتصال با موفقیت پایان یافت.'
      break
    case 1001:
      hint = 'سرور در حال خاتمه اتصال است.'
      break
    case 1006:
      hint = 'اتصال بدون پیام خروج قطع شد. احتمال مشکل شبکه یا فایروال.'
      break
    case 1008:
      hint = 'به دلیل نقض خط‌مشی امنیتی قطع شد.'
      break
    case 1009:
      hint = 'حجم پیام ارسالی زیاد بود.'
      break
    case 1013:
      hint = 'سرور موقتاً شلوغ است. لطفاً بعداً تلاش کنید.'
      break
    default:
      hint = event.reason || ''
  }

  const userMessage = [base, hint].filter(Boolean).join(' - ')
  return { code, reason: event.reason || base, wasClean: event.wasClean, userMessage }
}

export function parseWsError(error: Event | any): string {
  // Browsers often hide details; provide a generic message
  const message = typeof error?.message === 'string' ? error.message : ''
  return message || 'خطا در اتصال وب‌سوکت. لطفاً اتصال اینترنت خود را بررسی کنید.'
}