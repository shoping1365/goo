import { defineEventHandler, createError } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

interface MobileAppHeaderListResponse {
     success: boolean
     data: any[]
     message?: string
}

export default defineEventHandler(async (event): Promise<MobileAppHeaderListResponse> => {
     try {
          console.log('درخواست دریافت لیست هدرهای موبایل (ادمین)')

          const responseData = await fetchGo(event, '/api/admin/mobile-app-header-settings', {
               method: 'GET'
          })

          console.log('پاسخ لیست هدرهای موبایل:', responseData)

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
          console.error('خطا در دریافت لیست هدرهای موبایل:', error)

          // اگر خطا از سرور Go آمده باشد
          if (error.data) {
               return {
                    success: false,
                    data: [],
                    message: error.data.message || error.data.error || 'خطا در دریافت لیست هدرهای موبایل'
               }
          }

          // اگر خطای شبکه باشد
          return {
               success: false,
               data: [],
               message: 'خطا در اتصال به سرور'
          }
     }
})