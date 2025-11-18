import { defineEventHandler, getRouterParam, getCookie, createError } from 'h3'
import { useRuntimeConfig } from '#imports'

export default defineEventHandler(async (event) => {
    const config = useRuntimeConfig()
    const apiBaseUrl = config.public.goApiBase

    try {
        console.log('🗑️ Starting single redirect delete...')

        // گرفتن ID از URL
        const id = getRouterParam(event, 'id')
        console.log('📝 Redirect ID:', id)

        if (!id) {
            console.log('❌ Validation failed: missing id')
            throw createError({
                statusCode: 400,
                message: 'ID ریدایرکت الزامی است'
            })
        }

        console.log('🌐 API Base URL:', apiBaseUrl)

        // بررسی احراز هویت
        const authToken = getCookie(event, 'auth-token') || getCookie(event, 'access_token')
        console.log('🔑 Auth token exists:', !!authToken)

        if (!authToken) {
            throw createError({
                statusCode: 401,
                message: 'لطفاً وارد شوید'
            })
        }

        const url = `${apiBaseUrl}/api/admin/seo/redirects/${id}`
        console.log('🎯 Request URL:', url)

        // درخواست به Go backend
        const response = await $fetch(url, {
            method: 'DELETE',
            headers: {
                'Authorization': `Bearer ${authToken}`,
                'Content-Type': 'application/json'
            }
        })

        console.log('✅ Delete successful:', response)
        return response

    } catch (error: unknown) {
        console.error('❌ Error in delete:', error)
        const errorObj = error as { statusCode?: number; statusMessage?: string }

        throw createError({
            statusCode: errorObj.statusCode || 500,
            message: errorObj.statusMessage || 'خطا در حذف ریدایرکت'
        })
    }
})