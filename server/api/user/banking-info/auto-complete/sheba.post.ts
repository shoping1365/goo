import { defineEventHandler, createError, readBody } from 'h3'
import { goApiFetch } from '~/server/utils/goApi'

interface AutoCompleteResponse {
  success: boolean
  message?: string
  data?: {
    bank_name?: string
    card_number?: string
    sheba_number?: string
  }
}

export default defineEventHandler(async (event): Promise<AutoCompleteResponse> => {
  try {
    const body = await readBody(event)
    
    if (!body.sheba_number) {
      throw createError({
        statusCode: 400,
        message: 'شماره شبا الزامی است'
      })
    }

    const response = await goApiFetch<AutoCompleteResponse>(event, '/api/user/banking-info/auto-complete/sheba', {
      method: 'POST',
      body: {
        sheba_number: body.sheba_number
      }
    })

    return {
      success: true,
      data: response.data
    }
  } catch (error: unknown) {
    console.error('خطا در تکمیل خودکار شبا:', error)
    
    if (error && typeof error === 'object' && 'statusCode' in error) {
      throw error
    }
    
    throw createError({
      statusCode: 500,
      message: 'خطا در تکمیل خودکار شبا'
    })
  }
})