import type { H3Event } from 'h3';
import { createError, defineEventHandler, readFormData } from 'h3';

declare const useRuntimeConfig: () => { public: { goApiBase: string } }
declare const $fetch: <T = unknown>(url: string, options?: { method?: string; headers?: Record<string, string>; body?: unknown }) => Promise<T>

interface ReviewRequestBody {
  rating: number
  comment: string
  product_id: number
  anonymous: boolean
  title?: string
  pros?: string
  cons?: string
}

export default defineEventHandler(async (event: H3Event) => {
  try {
    const formData = await readFormData(event)

    // Extract form data
    const rating = formData.get('rating')
    const comment = formData.get('comment')
    const productId = formData.get('product_id')
    const title = formData.get('title')
    const pros = formData.get('pros')
    const cons = formData.get('cons')
    const anonymous = formData.get('anonymous') === 'true'

    if (!rating || !comment || !productId) {
      throw createError({
        statusCode: 400,
        message: 'امتیاز، نظر و شناسه محصول الزامی است'
      })
    }

    // دریافت توکن کاربر
    const authHeader = getHeader(event, 'authorization') || ''
    if (!authHeader) {
      throw createError({
        statusCode: 401,
        message: 'برای ثبت نظر باید وارد شوید.'
      })
    }

    const config = useRuntimeConfig()
    const baseURL = config.public.goApiBase

    // Prepare request body
    const requestBody: ReviewRequestBody = {
      rating: parseInt(rating as string),
      comment: comment as string,
      product_id: parseInt(productId as string),
      anonymous: anonymous
    }

    if (title) requestBody.title = title as string
    if (pros) requestBody.pros = pros as string
    if (cons) requestBody.cons = cons as string

    // Handle file uploads if any
    const images = formData.getAll('images')
    const videos = formData.getAll('videos')

    if (images.length > 0 || videos.length > 0) {
      // For now, we'll skip file uploads and just send the review data
      // File uploads can be implemented later with proper multipart handling
    }

    const response = await $fetch(`${baseURL}/api/reviews`, {
      method: 'POST',
      body: requestBody,
      headers: {
        Authorization: authHeader,
        'Content-Type': 'application/json'
      }
    })

    return response
  } catch (error: unknown) {
    console.error('Error creating review:', error)
    const errorWithStatus = error as { statusCode?: number; message?: string }
    throw createError({
      statusCode: errorWithStatus?.statusCode || 500,
      message: errorWithStatus?.message || 'خطا در ثبت نظر'
    })
  }
})