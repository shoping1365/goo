import { createError, defineEventHandler, getRouterParam } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

interface MobileAppHeaderDeleteResponse {
     success: boolean
     message?: string
}

export default defineEventHandler(async (event): Promise<MobileAppHeaderDeleteResponse> => {
     try {
          const id = getRouterParam(event, 'id')

          if (!id) {
               throw createError({
                    statusCode: 400,
                    message: 'شناسه هدر موبایل الزامی است'
               })
          }

          // console.log('درخواست حذف هدر موبایل با ID:', id)

          const responseData = await fetchGo(event, `/api/admin/mobile-app-header-settings/${id}`, {
               method: 'DELETE'
          })

          // console.log('پاسخ حذف هدر موبایل:', responseData)

          if (responseData && typeof responseData === 'object') {
               const payload = responseData as Record<string, unknown>
               return {
                    success: payload.success !== false,
                    message: (payload.message as string) || 'هدر موبایل با موفقیت حذف شد'
               }
          }

          return {
               success: true,
               message: 'هدر موبایل با موفقیت حذف شد'
          }

     } catch (error: unknown) {
          const err = error as { statusCode?: number; data?: { message?: string; error?: string } }
          // console.error('خطا در حذف هدر موبایل:', error)

          // اگر خطا از سرور Go آمده باشد
          if (err.data) {
               throw createError({
                    statusCode: err.statusCode || 500,
                    message: err.data.message || err.data.error || 'خطا در حذف هدر موبایل',
                    data: err.data
               })
          }

          // اگر خطای شبکه باشد
          throw createError({
               statusCode: 500,
               message: 'خطا در اتصال به سرور'
          })
     }
})