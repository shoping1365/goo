import { defineEventHandler, getQuery } from 'h3'
import { getCookie } from 'h3'

export default defineEventHandler(async (event) => {
  try {
    // دریافت پارامترهای query
    const query = getQuery(event)
    const page = query.page || '1'
    const limit = query.limit || '10'

    // دریافت آمار ترافیک کلی از بک‌اند
    const response = await $fetch(`${process.env.BACKEND_URL}/api/admin/traffic/general-stats`, {
      query: {
        page,
        limit,
        adType: query.adType || '',
        ipAddress: query.ipAddress || '',
        location: query.location || '',
        pagePath: query.page || '',
        status: query.status || '',
        timeRange: query.timeRange || ''
      },
      headers: {
        'Authorization': `Bearer ${getCookie(event, 'auth-token')}`,
        'Content-Type': 'application/json'
      }
    })

    return response
  } catch (error: any) {
    console.error('خطا در دریافت آمار ترافیک کلی:', error)
    
    // در صورت خطا، داده‌های نمونه برگردان
    return {
      status: 'success',
      data: {
        traffic_logs: [],
        pagination: {
          current_page: 1,
          total_pages: 1,
          total_items: 0,
          items_per_page: 10,
          has_next: false,
          has_prev: false
        }
      }
    }
  }
}) 