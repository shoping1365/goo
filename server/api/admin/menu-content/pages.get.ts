import { defineEventHandler } from 'h3'

interface Page {
    id: number
    title: string
    slug: string
    [key: string]: unknown
}

interface RawPageResponse {
    data?: Page[]
    [key: string]: unknown
}

const normalizePayload = (payload: unknown): Page[] => {
    if (Array.isArray(payload)) return payload as Page[]
    const response = payload as RawPageResponse
    if (response && Array.isArray(response.data)) return response.data
    return []
}

declare const useRuntimeConfig: () => { public: { goApiBase: string } }
declare const $fetch: <T = unknown>(url: string, options?: { query?: Record<string, string | number> }) => Promise<T>

export default defineEventHandler(async () => {
    const config = useRuntimeConfig()
    const base = config.public.goApiBase
    try {
        const raw = await $fetch<RawPageResponse>(`${base}/api/posts`, {
            query: {
                type: 'page',
                status: 'published',
                limit: 100
            }
        })

        return {
            success: true,
            data: normalizePayload(raw)
        }
    } catch (error: unknown) {
        console.error('Error fetching pages for menu content:', error)
        return {
            success: false,
            data: [],
            error: 'خطا در دریافت صفحات'
        }
    }
})