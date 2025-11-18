import { defineEventHandler } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

interface FooterListResponse {
  success: boolean
  data: any[]
  message?: string
}

export default defineEventHandler(async (event): Promise<FooterListResponse> => {
  try {
    console.log('Fetching admin footer list...')
    const responseData = await fetchGo(event, '/api/admin/footer-settings', {
      method: 'GET'
    })

    console.log('Admin footer list fetched successfully.', {
      hasData: Boolean(responseData),
      keys: responseData && typeof responseData === 'object' ? Object.keys(responseData) : undefined
    })

    if (responseData && typeof responseData === 'object' && 'data' in responseData) {
      const payload = responseData as Record<string, any>
      return {
        success: payload.success !== false,
        data: payload.data || []
      }
    }

    return {
      success: true,
      data: Array.isArray(responseData) ? responseData : []
    }
  } catch (error: any) {
    console.error('Unexpected error while fetching admin footer list:', {
      statusCode: error?.statusCode,
      status: error?.status,
      message: error?.message,
      data: error?.data
    })
    return {
      success: false,
      data: [],
      message: error?.data?.message || error?.data?.error || error?.message || 'خطا در اتصال به سرور'
    }
  }
})
