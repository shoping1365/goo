import type { H3Event } from 'h3';
import { defineEventHandler } from 'h3';

declare const useRuntimeConfig: () => { public: { goApiBase: string } }
declare const createError: (options: { statusCode: number; message: string }) => Error
declare const getHeader: (event: H3Event, name: string) => string | undefined

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const backendUrl = config.public.goApiBase

  try {
    const response = await fetch(`${backendUrl}/api/admin/api-settings`, {
      headers: {
        'Content-Type': 'application/json',
        ...getHeaders(event)
      }
    })

    if (!response.ok) {
      throw createError({
        statusCode: response.status,
        message: response.statusText
      })
    }

    return await response.json()
  } catch (error: unknown) {
    console.error('Error fetching API settings:', error)
    throw createError({
      statusCode: 500,
      message: 'خطا در دریافت تنظیمات API'
    })
  }
})

function getHeaders(event: H3Event): Record<string, string> {
  const headers: Record<string, string> = {}

  // Forward important headers
  const cookie = getHeader(event, 'cookie')
  if (cookie) headers.cookie = cookie

  const authorization = getHeader(event, 'authorization')
  if (authorization) headers.authorization = authorization

  return headers
} 