import { defineEventHandler, createError } from 'h3'
import { goApiFetch } from '~/server/utils/goApi'

interface BankingInfoResponse {
  success: boolean
  data?: unknown[]
}

export default defineEventHandler(async (event) => {
  try {
    // ارسال درخواست به بک‌اند Go
    const response = await goApiFetch<BankingInfoResponse>(event, '/api/user/banking-info', {
      method: 'GET'
    })

    return {
      success: true,
      data: response.data || []
    }
  } catch (error: unknown) {
    console.error('خطا در دریافت اطلاعات بانکی:', error)
    
    if (error && typeof error === 'object' && 'statusCode' in error && error.statusCode === 404) {
      return {
        success: true,
        data: []
      }
    }
    
    if (error && typeof error === 'object' && 'statusCode' in error) {
      throw error
    }
    
    throw createError({
      statusCode: 500,
      message: 'خطا در دریافت اطلاعات بانکی'
    })
  }
})