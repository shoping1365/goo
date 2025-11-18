import { createError, defineEventHandler, parseCookies, readMultipartFormData } from 'h3'

interface MobileAppHeaderUploadResponse {
     success: boolean
     data: unknown
     message?: string
}

declare const useRuntimeConfig: () => { public: { goApiBase: string } }

export default defineEventHandler(async (event): Promise<MobileAppHeaderUploadResponse> => {
     try {
          const config = useRuntimeConfig()

          console.log('درخواست آپلود لوگو هدر موبایل')

          // دریافت cookies از درخواست
          const cookies = parseCookies(event)
          const cookieHeader = Object.entries(cookies)
               .map(([key, value]) => `${key}=${value}`)
               .join('; ')

          // دریافت فایل از درخواست
          const formData = await readMultipartFormData(event)

          if (!formData || formData.length === 0) {
               throw createError({
                    statusCode: 400,
                    message: 'فایل لوگو الزامی است'
               })
          }

          // ایجاد FormData برای ارسال به Go backend
          const goFormData = new FormData()
          for (const field of formData) {
               if (field.type === 'file') {
                    // Convert Buffer to ArrayBuffer for Blob
                    // Create a new Uint8Array copy to ensure we have a proper ArrayBuffer
                    const dataArray = field.data instanceof Uint8Array
                         ? new Uint8Array(field.data)
                         : new Uint8Array(field.data as ArrayLike<number>)
                    // Create a new ArrayBuffer from the Uint8Array
                    const arrayBuffer = dataArray.buffer.slice(dataArray.byteOffset, dataArray.byteOffset + dataArray.byteLength) as ArrayBuffer
                    goFormData.append('logo', new Blob([arrayBuffer]), field.filename || 'logo.png')
               } else {
                    goFormData.append(field.name!, field.data.toString())
               }
          }

          // ارسال درخواست به Go backend با fetch معمولی
          const response = await fetch(`${config.public.goApiBase}/api/admin/mobile-app-header-settings/upload-logo`, {
               method: 'POST',
               headers: {
                    'Cookie': cookieHeader
               },
               body: goFormData
          })

          if (!response.ok) {
               throw new Error(`HTTP error! status: ${response.status}`)
          }

          const responseData = await response.json()

          console.log('پاسخ آپلود لوگو هدر موبایل:', responseData)

          return {
               success: true,
               data: responseData?.data || null
          }

     } catch (error: unknown) {
          console.error('خطا در آپلود لوگو هدر موبایل:', error)

          const errorWithStatus = error as { statusCode?: number; message?: string; data?: { message?: string; error?: string } }
          // اگر خطا از سرور Go آمده باشد
          if (errorWithStatus?.data) {
               throw createError({
                    statusCode: errorWithStatus.statusCode || 500,
                    message: errorWithStatus.data.message || errorWithStatus.data.error || 'خطا در آپلود لوگو هدر موبایل',
                    data: errorWithStatus.data
               })
          }

          // اگر خطای شبکه باشد
          throw createError({
               statusCode: errorWithStatus?.statusCode || 500,
               message: errorWithStatus?.message || 'خطا در اتصال به سرور'
          })
     }
})

