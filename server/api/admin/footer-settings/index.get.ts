import { defineEventHandler } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

interface FooterListResponse {
  success: boolean
  data: unknown[]
  message?: string
}

export default defineEventHandler(async (event): Promise<FooterListResponse> => {
  try {
    // console.log('Fetching admin footer list...')
    const responseData = await fetchGo(event, '/api/admin/footer-settings', {
      method: 'GET'
    })

    // console.log('Admin footer list fetched successfully.', {
    //   hasData: Boolean(responseData),
    //   keys: responseData && typeof responseData === 'object' ? Object.keys(responseData) : undefined
    // })

    if (responseData && typeof responseData === 'object' && 'data' in responseData) {
      const payload = responseData as { success?: boolean; data?: unknown[] }
      return {
        success: payload.success !== false,
        data: payload.data || []
      }
    }

    return {
      success: true,
      data: Array.isArray(responseData) ? responseData : []
    }
  } catch (error: unknown) {
    const err = error as { data?: { message?: string; error?: string }; message?: string }
    // console.error('Unexpected error while fetching admin footer list:', {
    //   statusCode: err?.statusCode,
    //   status: err?.status,
    //   message: err?.message,
    //   data: err?.data
    // })
    return {
      success: false,
      data: [],
      message: err?.data?.message || err?.data?.error || err?.message || 'خطا در اتصال به سرور'
    }
  }
})
