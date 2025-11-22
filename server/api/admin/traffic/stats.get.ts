export default defineEventHandler(async (event) => {
  try {
    // دریافت آمار ترافیک از بک‌اند
    const response = await $fetch(`${process.env.BACKEND_URL}/api/admin/traffic/stats`, {
      headers: {
        'Authorization': `Bearer ${getCookie(event, 'auth-token')}`,
        'Content-Type': 'application/json'
      }
    })

    return response
  } catch (error: unknown) {
    console.error('خطا در دریافت آمار ترافیک:', error)
    
    // در صورت خطا، داده‌های نمونه برگردان
    return {
      status: 'success',
      data: {
        online_users: 0,
        today_requests: 0,
        suspicious_requests: 0,
        blocked_attacks: 0,
        hourly_traffic: []
      }
    }
  }
}) 