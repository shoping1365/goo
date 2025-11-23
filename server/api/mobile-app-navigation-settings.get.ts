import { getHeader } from './_utils/getHeader';

export default defineEventHandler(async (event): Promise<{ success: boolean; data: unknown[] }> => {
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

          return {
               success: true,
               data: (response.data as unknown[]) || []
          }
     } catch (error: unknown) {
          const err = error as { statusCode?: number; statusMessage?: string }
          console.error('Error fetching public navigation settings:', err)

          throw createError({
               statusCode: err.statusCode || 500,
               message: err.statusMessage || 'خطا در دریافت تنظیمات ناوبری'
          })
     }
})