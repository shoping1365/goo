import { defineEventHandler, getRouterParam } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

interface FooterGetResponse {
  success: boolean
  data?: unknown
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

    // console.log('Fetching admin footer by id:', id)

    const responseData = await fetchGo(event, `/api/admin/footer-settings/${id}`, {
      method: 'GET'
    })

    // console.log('Admin footer fetched successfully.', {
    //   hasData: Boolean(responseData)
    // })

    if (responseData && typeof responseData === 'object' && 'data' in responseData) {
      const payload = responseData as { success?: boolean; data?: unknown }
      return {
        success: payload.success !== false,
        data: payload.data
      }
    }

    return {
      success: true,
      data: responseData
    }
  } catch (error: unknown) {
    const err = error as { data?: { message?: string; error?: string }; message?: string }
    // console.error('Unexpected error while fetching admin footer:', {
    //   statusCode: err?.statusCode,
    //   status: err?.status,
    //   message: err?.message,
    //   data: err?.data
    // })
    return {
      success: false,
      message: err?.data?.message || err?.data?.error || err?.message || 'خطا در دریافت فوتر'
    }
  }
})



