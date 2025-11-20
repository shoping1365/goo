import { defineEventHandler, getRouterParam } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

interface FooterDeleteResponse {
  success: boolean
  message?: string
}

export default defineEventHandler(async (event): Promise<FooterDeleteResponse> => {
  try {
    const id = getRouterParam(event, 'id')

    if (!id) {
      return {
        success: false,
        message: 'شناسه فوتر الزامی است'
      }
    }

    // console.log('Deleting admin footer:', id)

    try {
      await fetchGo(event, `/api/admin/footer-settings/${id}`, {
        method: 'DELETE'
      })
    } catch (fetchErr: unknown) {
      const err = fetchErr as { data?: { message?: string; error?: string }; message?: string }
      // console.error('Admin footer delete failed:', {
      //   statusCode: err?.statusCode,
      //   status: err?.status,
      //   message: err?.message,
      //   data: err?.data
      // })
      return {
        success: false,
        message: err?.data?.message || err?.data?.error || err?.message || 'خطا در حذف فوتر'
      }
    }

    // console.log('Admin footer deleted successfully.')

    return {
      success: true,
      message: 'فوتر با موفقیت حذف شد'
    }
  } catch (error: unknown) {
    const err = error as { data?: { message?: string; error?: string }; message?: string }
    // console.error('Unexpected error during admin footer delete:', error)
    return {
      success: false,
      message: err?.data?.message || err?.data?.error || err?.message || 'خطا در حذف فوتر'
    }
  }
})
