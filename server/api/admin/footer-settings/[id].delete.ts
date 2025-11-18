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

    console.log('Deleting admin footer:', id)

    try {
      await fetchGo(event, `/api/admin/footer-settings/${id}`, {
        method: 'DELETE'
      })
    } catch (fetchErr: any) {
      console.error('Admin footer delete failed:', {
        statusCode: fetchErr?.statusCode,
        status: fetchErr?.status,
        message: fetchErr?.message,
        data: fetchErr?.data
      })
      return {
        success: false,
        message: fetchErr?.data?.message || fetchErr?.data?.error || fetchErr?.message || 'خطا در حذف فوتر'
      }
    }

    console.log('Admin footer deleted successfully.')

    return {
      success: true,
      message: 'فوتر با موفقیت حذف شد'
    }
  } catch (error: any) {
    console.error('Unexpected error during admin footer delete:', error)
    return {
      success: false,
      message: error?.data?.message || error?.data?.error || error?.message || 'خطا در حذف فوتر'
    }
  }
})
