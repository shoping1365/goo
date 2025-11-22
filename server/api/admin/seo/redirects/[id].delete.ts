import { defineEventHandler, getRouterParam, getCookie, createError } from 'h3'
import { useRuntimeConfig } from '#imports'

export default defineEventHandler(async (event) => {
    const config = useRuntimeConfig()
    const apiBaseUrl = config.public.goApiBase

    try {
        // گرفتن ID از URL
        const id = getRouterParam(event, 'id')

        if (!id) {
            throw createError({
                statusCode: 400,
                message: 'ID ریدایرکت الزامی است'
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

        const url = `${apiBaseUrl}/api/admin/seo/redirects/${id}`

        // درخواست به Go backend
        const response = await $fetch(url, {
            method: 'DELETE',
            headers: {
                'Authorization': `Bearer ${authToken}`,
                'Content-Type': 'application/json'
            }
        })

        return response

    } catch (error: unknown) {
        const errorObj = error as { statusCode?: number; statusMessage?: string }

        throw createError({
            statusCode: errorObj.statusCode || 500,
            message: errorObj.statusMessage || 'خطا در حذف ریدایرکت'
        })
    }
})