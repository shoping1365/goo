export default defineEventHandler(async (event) => {
  try {
    // دریافت پارامترهای صفحه‌بندی از query string
    const query = getQuery(event)
    const page = query.page || '1'
    const limit = query.limit || '10'
    
    // دریافت کاربران آنلاین از بک‌اند با صفحه‌بندی
    const response = await $fetch(`${process.env.BACKEND_URL}/api/admin/traffic/online-users`, {
      query: {
        page,
        limit
      },
      headers: {
        'Authorization': `Bearer ${getCookie(event, 'auth-token')}`,
        'Content-Type': 'application/json'
      }
    })

    return response
  } catch (error: any) {
    console.error('خطا در دریافت کاربران آنلاین:', error)
    
    // در صورت خطا، لیست خالی برگردان
    return {
      status: 'success',
      data: [],
      pagination: {
        current_page: 1,
        total_pages: 0,
        total_items: 0,
        items_per_page: 10,
        has_next_page: false,
        has_prev_page: false
      }
    }
  }
}) 