import { defineEventHandler, createError, getRouterParam } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

interface MobileAppHeaderResponse {
     success: boolean
     data: any
     message?: string
}

export default defineEventHandler(async (event): Promise<MobileAppHeaderResponse> => {
     try {
          const id = getRouterParam(event, 'id')

          if (!id) {
               throw createError({
                    statusCode: 400,
                    message: 'شناسه هدر موبایل الزامی است'
               })
          }

          console.log('درخواست دریافت هدر موبایل با ID:', id)

          const responseData = await fetchGo(event, `/api/admin/mobile-app-header-settings/${id}`, {
               method: 'GET'
          })

          console.log('پاسخ هدر موبایل:', responseData)

          if (responseData && typeof responseData === 'object') {
               const payload = responseData as Record<string, any>
               return {
                    success: payload.success !== false,
                    data: payload.data || null
               }
          }

          return {
               success: true,
               data: responseData || null
          }

     } catch (error: any) {
          console.error('خطا در دریافت هدر موبایل:', error)

          // اگر خطا از سرور Go آمده باشد
          if (error.data) {
               throw createError({
                    statusCode: error.statusCode || 500,
                    message: error.data.message || error.data.error || 'خطا در دریافت هدر موبایل',
                    data: error.data
               })
          }

          // اگر خطای شبکه باشد
          throw createError({
               statusCode: 500,
               message: 'خطا در اتصال به سرور'
          })
     }
})