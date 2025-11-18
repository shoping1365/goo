import { defineEventHandler } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

interface Category {
  id: number
  name: string
  slug: string
  [key: string]: unknown
}

interface RawCategoryResponse {
  data?: Category[]
  [key: string]: unknown
}

const normalizePayload = (payload: unknown): Category[] => {
  if (Array.isArray(payload)) return payload as Category[]
  const response = payload as RawCategoryResponse
  if (response && Array.isArray(response.data)) return response.data
  return []
}

export default defineEventHandler(async (event) => {
  try {
    const raw = await fetchGo(event, '/api/post-categories')
    const categories = normalizePayload(raw)

    return {
      success: true,
      data: categories,
      message: 'دسته‌ها با موفقیت دریافت شدند'
    }
  } catch (error: unknown) {
    console.error('Error fetching categories:', error)

    return {
      success: false,
      data: [],
      message: 'خطا در دریافت دسته‌ها'
    }
  }
})