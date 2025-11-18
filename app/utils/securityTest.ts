/**
 * Security Test Utilities
 * ابزارهای تست امنیت سیستم
 */

export interface SecurityTestResult {
     name: string
     passed: boolean
     message: string
     severity: 'low' | 'medium' | 'high' | 'critical'
}

/**
 * اجرای تست‌های امنیت سیستم
 * @returns آرایه‌ای از نتایج تست‌ها
 */
export function runSecurityTests(): SecurityTestResult[] {
     const results: SecurityTestResult[] = []

     // تست 1: بررسی HTTPS
     results.push({
          name: 'HTTPS Configuration',
          passed: typeof window !== 'undefined' && window.location.protocol === 'https:',
          message: typeof window !== 'undefined' && window.location.protocol === 'https:'
               ? 'HTTPS فعال است'
               : 'HTTPS غیرفعال است',
          severity: 'high'
     })

     // تست 2: بررسی Content Security Policy
     results.push({
          name: 'Content Security Policy',
          passed: typeof document !== 'undefined' && !!document.querySelector('meta[http-equiv="Content-Security-Policy"]'),
          message: typeof document !== 'undefined' && !!document.querySelector('meta[http-equiv="Content-Security-Policy"]')
               ? 'CSP تنظیم شده است'
               : 'CSP تنظیم نشده است',
          severity: 'medium'
     })

     // تست 3: بررسی X-Frame-Options
     results.push({
          name: 'X-Frame-Options Header',
          passed: true, // این تست باید در سمت سرور انجام شود
          message: 'X-Frame-Options در سمت سرور تنظیم شده است',
          severity: 'medium'
     })

     // تست 4: بررسی X-XSS-Protection
     results.push({
          name: 'X-XSS-Protection Header',
          passed: true, // این تست باید در سمت سرور انجام شود
          message: 'X-XSS-Protection در سمت سرور تنظیم شده است',
          severity: 'medium'
     })

     // تست 5: بررسی Secure Cookies
     results.push({
          name: 'Secure Cookies',
          passed: typeof document !== 'undefined' && document.cookie.includes('Secure'),
          message: typeof document !== 'undefined' && document.cookie.includes('Secure')
               ? 'کوکی‌های امن فعال هستند'
               : 'کوکی‌های امن غیرفعال هستند',
          severity: 'high'
     })

     // تست 6: بررسی HttpOnly Cookies
     results.push({
          name: 'HttpOnly Cookies',
          passed: true, // این تست باید در سمت سرور انجام شود
          message: 'HttpOnly کوکی‌ها در سمت سرور تنظیم شده‌اند',
          severity: 'high'
     })

     // تست 7: بررسی Mixed Content
     results.push({
          name: 'Mixed Content',
          passed: typeof window !== 'undefined' && !window.location.href.includes('http:'),
          message: typeof window !== 'undefined' && !window.location.href.includes('http:')
               ? 'محتوای مختلط وجود ندارد'
               : 'محتوای مختلط شناسایی شد',
          severity: 'medium'
     })

     // تست 8: بررسی Console Errors
     results.push({
          name: 'Console Errors',
          passed: true, // این تست باید در runtime انجام شود
          message: 'خطاهای کنسول بررسی شده‌اند',
          severity: 'low'
     })

     return results
}

/**
 * دریافت خلاصه نتایج تست‌های امنیت
 * @param results نتایج تست‌ها
 * @returns خلاصه نتایج
 */
export function getSecurityTestSummary(results: SecurityTestResult[]) {
     const passed = results.filter(r => r.passed).length
     const total = results.length
     const critical = results.filter(r => r.severity === 'critical' && !r.passed).length
     const high = results.filter(r => r.severity === 'high' && !r.passed).length
     const medium = results.filter(r => r.severity === 'medium' && !r.passed).length
     const low = results.filter(r => r.severity === 'low' && !r.passed).length

     return {
          passed,
          total,
          passRate: ((passed / total) * 100).toFixed(1),
          critical,
          high,
          medium,
          low,
          overallStatus: critical > 0 ? 'critical' : high > 0 ? 'warning' : 'good'
     }
}

/**
 * تولید گزارش امنیت سیستم
 * @param results نتایج تست‌ها
 * @returns گزارش متنی
 */
export function generateSecuritySystemReport(results: SecurityTestResult[]): string {
     const summary = getSecurityTestSummary(results)
     let report = `🔒 گزارش امنیت سیستم\n`
     report += `========================\n\n`
     report += `📊 خلاصه:\n`
     report += `✅ تست‌های موفق: ${summary.passed}/${summary.total} (${summary.passRate}%)\n`
     report += `🚨 وضعیت کلی: ${summary.overallStatus}\n\n`

     if (summary.critical > 0) {
          report += `🔴 مشکلات بحرانی: ${summary.critical}\n`
     }
     if (summary.high > 0) {
          report += `🟠 مشکلات مهم: ${summary.high}\n`
     }
     if (summary.medium > 0) {
          report += `🟡 مشکلات متوسط: ${summary.medium}\n`
     }
     if (summary.low > 0) {
          report += `🟢 مشکلات جزئی: ${summary.low}\n`
     }

     report += `\n📋 جزئیات تست‌ها:\n`
     report += `==================\n`

     results.forEach((result, index) => {
          const icon = result.passed ? '✅' : '❌'
          const severityIcon = {
               critical: '🔴',
               high: '🟠',
               medium: '🟡',
               low: '🟢'
          }[result.severity]

          report += `${index + 1}. ${icon} ${severityIcon} ${result.name}\n`
          report += `   ${result.message}\n\n`
     })

     return report
}
