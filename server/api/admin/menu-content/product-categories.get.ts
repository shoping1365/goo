import { defineEventHandler } from 'h3'

interface ProductCategory {
  id: number
  name: string
  slug: string
  [key: string]: unknown
}

interface RawProductCategoryResponse {
  data?: ProductCategory[]
  [key: string]: unknown
}

const normalizePayload = (payload: unknown): ProductCategory[] => {
  if (Array.isArray(payload)) return payload as ProductCategory[]
  const response = payload as RawProductCategoryResponse
  if (response && Array.isArray(response.data)) return response.data
  return []
}

declare const useRuntimeConfig: () => { public: { goApiBase: string } }
declare const $fetch: <T = unknown>(url: string, options?: { query?: Record<string, string | number> }) => Promise<T>

export default defineEventHandler(async () => {
  try {
    const config = useRuntimeConfig()
    const base = config.public.goApiBase
    const raw = await $fetch<RawProductCategoryResponse>(`${base}/api/product-categories`, {
      query: {
        limit: 100
      }
    })

    return {
      success: true,
      data: normalizePayload(raw)
    }
  } catch (error: unknown) {
    console.error('Error fetching product categories for menu content:', error)
    return {
      success: false,
      data: [],
      error: 'خطا در دریافت دسته‌بندی محصولات'
    }
  }
})