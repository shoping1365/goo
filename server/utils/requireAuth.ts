import type { H3Event } from 'h3'
import { getCookie, createError } from 'h3'

export interface AuthUser {
  id: number
  role: string
  email?: string
  [key: string]: unknown
}

/**
 * Require authentication for API routes
 * Returns the authenticated user or throws an error
 */
export async function requireAuth(event: H3Event): Promise<AuthUser> {
  // Get token from cookies
  const token = getCookie(event, 'access_token') || getCookie(event, 'auth-token')
  
  if (!token) {
    throw createError({
      statusCode: 401,
      statusMessage: 'Unauthorized',
      message: 'احراز هویت نشده است'
    })
  }

  // Try to get user from /api/auth/me endpoint
  try {
    const config = useRuntimeConfig()
    const base = config.public.goApiBase
    
    const response = await $fetch<{ id?: number; ID?: number; role?: string; effective_role_name?: string; email?: string }>(`${base}/api/auth/me`, {
      headers: {
        Authorization: `Bearer ${token}`,
        Cookie: `access_token=${token}`
      }
    })

    if (!response || !response.id) {
      throw createError({
        statusCode: 401,
        statusMessage: 'Unauthorized',
        message: 'کاربر یافت نشد'
      })
    }

    return {
      id: response.id || response.ID,
      role: response.role || response.effective_role_name || 'user',
      email: response.email,
      ...response
    }
  } catch (_error) {
    throw createError({
      statusCode: 401,
      statusMessage: 'Unauthorized',
      message: 'احراز هویت نشده است یا توکن منقضی شده'
    })
  }
}
