import { createError, defineEventHandler, getHeader, getRouterParam } from 'h3'

interface NavigationData {
     id: number
     name: string
     platform: string
     [key: string]: unknown
}

declare const useRuntimeConfig: () => { public: { goApiBase: string } }

export default defineEventHandler(async (event): Promise<{ success: boolean; data: NavigationData }> => {
     try {
          const config = useRuntimeConfig()
          const baseURL = config.public.goApiBase
          const id = getRouterParam(event, 'id')

          if (!id) {
               throw createError({
                    statusCode: 400,
                    message: 'Navigation ID is required'
               })
          }

          // Call the Go backend API directly
          const fetchResponse = await fetch(`${baseURL}/api/admin/mobile-app-navigation-settings/${id}`, {
               method: 'GET',
               headers: {
                    'Authorization': getHeader(event, 'authorization') || '',
                    'Content-Type': 'application/json',
                    'Cookie': getHeader(event, 'cookie') || ''
               }
          })

          if (!fetchResponse.ok) {
               throw createError({
                    statusCode: fetchResponse.status,
                    message: 'Failed to fetch navigation data'
               })
          }

          const response = await fetchResponse.json() as NavigationData


          return {
               success: true,
               data: response
          }
     } catch (error: unknown) {
          console.error('Error fetching navigation:', error)

          const errorWithStatus = error as { statusCode?: number; statusMessage?: string }
          throw createError({
               statusCode: errorWithStatus?.statusCode || 500,
               message: errorWithStatus?.statusMessage || 'خطا در دریافت اطلاعات ناوبری'
          })
     }
})