import { defineEventHandler, readBody, getCookie, createError } from 'h3'
import { useRuntimeConfig } from '#imports'

export default defineEventHandler(async (event) => {
    const config = useRuntimeConfig()
    const apiBaseUrl = config.public.goApiBase

    try {
        console.log('🗑️ Starting bulk delete...')

        // خواندن body درخواست
        const body = await readBody(event)
        console.log('📝 Request body:', body)

        if (!body?.ids || !Array.isArray(body.ids) || body.ids.length === 0) {
            console.log('❌ Validation failed: missing or empty ids array')
            throw createError({
                statusCode: 400,
                message: 'آرایه ids الزامی است'
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

        const url = `${apiBaseUrl}/api/admin/seo/redirects/bulk-delete`
        console.log('🎯 Request URL:', url)

        // درخواست به Go backend
        const response = await $fetch(url, {
            method: 'DELETE',
            headers: {
                'Authorization': `Bearer ${authToken}`,
                'Content-Type': 'application/json'
            },
            body: body
        })

        console.log('✅ Bulk delete successful:', response)
        return response

    } catch (error: unknown) {
        console.error('❌ Error in bulk delete:', error)
        const errorObj = error as { statusCode?: number; statusMessage?: string }

        throw createError({
            statusCode: errorObj.statusCode || 500,
            message: errorObj.statusMessage || 'خطا در حذف ریدایرکت‌ها'
        })
    }
})