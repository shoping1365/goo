import { defineEventHandler } from 'h3'

export default defineEventHandler(async (event) => {
  const query = getQuery(event)
  const page = query.page || '1'
  const limit = query.limit || '10'
  
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  
  try {
    const response = await fetch(`${base}/api/ippanel-outbox?page=${page}&limit=${limit}`, {
      method: 'GET',
      headers: { 'Accept': 'application/json' }
    })
    
    if (!response.ok) {
      console.error('❌ خطا در دریافت outbox IPPanel:', response.status, response.statusText)
      return {
        status: 'error',
        message: 'خطا در دریافت پیامک‌های ارسالی IPPanel',
        data: { messages: [], pagination: { total: 0, page: 1, limit: 10, total_pages: 1 } }
      }
    }
    
    const json = await response.json()
    
    return json
  } catch (error) {
    console.error('❌ خطا در ارتباط با سرور IPPanel:', error)
    return {
      status: 'error',
      message: 'خطا در ارتباط با سرور IPPanel',
      data: { messages: [], pagination: { total: 0, page: 1, limit: 10, total_pages: 1 } }
    }
  }
})