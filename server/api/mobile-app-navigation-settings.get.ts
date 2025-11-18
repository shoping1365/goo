import { getHeader } from './_utils/getHeader'

export default defineEventHandler(async (event): Promise<{ success: boolean; data: any[] }> => {
     try {
          const config = useRuntimeConfig()
          const baseURL = config.public.goApiBase
          const fetchResponse = await fetch(`${baseURL}/api/mobile-app-navigation-settings`, {
               method: 'GET',
               headers: {
                    'User-Agent': getHeader(event, 'user-agent') || '',
                    'Content-Type': 'application/json'
               }
          })

          if (!fetchResponse.ok) {
               throw createError({
                    statusCode: fetchResponse.status,
                    message: 'Failed to fetch navigation settings'
               })
          }

          const response = await fetchResponse.json()
          console.log('Public navigation settings response:', response)

          return {
               success: true,
               data: response.data || []
          }
     } catch (error: any) {
          console.error('Error fetching public navigation settings:', error)

          throw createError({
               statusCode: error.statusCode || 500,
               message: error.statusMessage || 'خطا در دریافت تنظیمات ناوبری'
          })
     }
})