import { defineEventHandler } from 'h3'

interface Post {
  id: number
  title: string
  slug: string
  [key: string]: unknown
}

interface RawPostResponse {
  data?: Post[]
  [key: string]: unknown
}

const normalizePayload = (payload: unknown): Post[] => {
  if (Array.isArray(payload)) return payload as Post[]
  const response = payload as RawPostResponse
  if (response && Array.isArray(response.data)) return response.data
  return []
}

declare const useRuntimeConfig: () => { public: { goApiBase: string } }
declare const $fetch: <T = unknown>(url: string, options?: { query?: Record<string, string | number> }) => Promise<T>

export default defineEventHandler(async () => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  try {
    const raw = await $fetch<RawPostResponse>(`${base}/api/posts`, {
      query: {
        type: 'post',
        status: 'published',
        limit: 100
      }
    })

    return {
      success: true,
      data: normalizePayload(raw)
    }
  } catch (error: unknown) {
    console.error('Error fetching posts for menu content:', error)
    return {
      success: false,
      data: [],
      error: 'خطا در دریافت نوشته‌ها'
    }
  }
})