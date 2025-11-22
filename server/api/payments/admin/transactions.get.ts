import { defineEventHandler, getQuery, createError } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

export default defineEventHandler(async (event) => {
  try {
    const query = getQuery(event)
    
    // پارامترهای فیلتر
    const page = parseInt(query.page as string) || 1
    const limit = parseInt(query.limit as string) || 10
    const status = query.status as string || ''
    const gateway_id = query.gateway_id as string || ''
    const search = query.search as string || ''
    const date = query.date as string || ''

    // ساخت query string
    const params = new URLSearchParams()
    if (page) params.append('page', page.toString())
    if (limit) params.append('limit', limit.toString())
    if (status) params.append('status', status)
    if (gateway_id) params.append('gateway_id', gateway_id)
    if (search) params.append('search', search)
    if (date) params.append('date', date)
    
    const queryString = params.toString()
    const url = `/api/payments/admin/transactions${queryString ? `?${queryString}` : ''}`
    
    // ارسال درخواست به backend از طریق fetchGo
    const response = await fetchGo(event, url, {
      method: 'GET'
    })

    return response
  } catch (error: unknown) {
    console.error('خطا در دریافت لیست تراکنش‌ها:', error)
    
    const err = error as { statusCode?: number; message?: string; statusMessage?: string }
    if (err.statusCode) {
      throw createError({
        statusCode: err.statusCode,
        message: err.message || err.statusMessage || 'خطا در دریافت لیست تراکنش‌ها'
      })
    }
    
    throw createError({
      statusCode: 500,
      message: 'خطای داخلی سرور'
    })
  }
}) 