import { defineEventHandler, readBody, getCookie, createError } from 'h3'
import { useRuntimeConfig } from '#imports'

export default defineEventHandler(async (event) => {
    const config = useRuntimeConfig()
    const apiBaseUrl = config.public.goApiBase

    try {
        // خواندن body درخواست
        const body = await readBody(event)

        if (!body?.ids || !Array.isArray(body.ids) || body.ids.length === 0) {
            throw createError({
                statusCode: 400,
                message: 'آرایه ids الزامی است'
            })
        }

        // بررسی احراز هویت
        const authToken = getCookie(event, 'auth-token') || getCookie(event, 'access_token')

        if (!authToken) {
            throw createError({
                statusCode: 401,
                message: 'لطفاً وارد شوید'
            })
        }

        const url = `${apiBaseUrl}/api/admin/seo/redirects/bulk-delete`

        // درخواست به Go backend
        const response = await $fetch(url, {
            method: 'DELETE',
            headers: {
                'Authorization': `Bearer ${authToken}`,
                'Content-Type': 'application/json'
            },
            body: body
        })

        return response

    } catch (error: unknown) {
        const errorObj = error as { statusCode?: number; statusMessage?: string }

        throw createError({
            statusCode: errorObj.statusCode || 500,
            message: errorObj.statusMessage || 'خطا در حذف ریدایرکت‌ها'
        })
    }
})