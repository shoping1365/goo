import { defineEventHandler, getRouterParam } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

interface FooterGetResponse {
  success: boolean
  data?: any
  message?: string
}

export default defineEventHandler(async (event): Promise<FooterGetResponse> => {
  try {
    const id = getRouterParam(event, 'id')

    if (!id) {
      return {
        success: false,
        message: 'شناسه فوتر الزامی است'
      }
    }

    console.log('Fetching admin footer by id:', id)

    const responseData = await fetchGo(event, `/api/admin/footer-settings/${id}`, {
      method: 'GET'
    })

    console.log('Admin footer fetched successfully.', {
      hasData: Boolean(responseData)
    })

    if (responseData && typeof responseData === 'object' && 'data' in responseData) {
      const payload = responseData as Record<string, any>
      return {
        success: payload.success !== false,
        data: payload.data
      }
    }

    return {
      success: true,
      data: responseData
    }
  } catch (error: any) {
    console.error('Unexpected error while fetching admin footer:', {
      statusCode: error?.statusCode,
      status: error?.status,
      message: error?.message,
      data: error?.data
    })
    return {
      success: false,
      message: error?.data?.message || error?.data?.error || error?.message || 'خطا در دریافت فوتر'
    }
  }
})



