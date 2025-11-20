/**
 * Admin Authentication Middleware
 * 
 * Server-side and client-side protection based on ADMIN_MIDDLEWARE_SECURITY.md
 * Prevents unauthorized users from seeing admin pages content.
 * Real authentication and authorization handled by backend API.
 */

declare const useRequestHeaders: (names?: string[]) => Record<string, string | undefined>

// @ts-ignore
export default defineNuxtRouteMiddleware(async (to, _from) => {
  // اجرا روی هم server و هم client برای امنیت کامل
  // Server-side execution prevents page content from being rendered for unauthorized users

  try {
    const headers: Record<string, string> = {
      Accept: 'application/json',
    }

    // ensure SSR requests forward session cookies to the backend auth check
    if (import.meta.server) {
      const { cookie } = useRequestHeaders(['cookie'])
      if (cookie) headers.Cookie = cookie
    }

    // Get user info from backend using global $fetch
    // @ts-ignore
    const response = await $fetch('/api/auth/me', {
      credentials: 'include',
      headers
    }) as { authenticated: boolean; user?: { role: string }; role?: string }

    // Check if user is authenticated
    if (!response?.authenticated) {
      throw new Error('User not authenticated')
    }

    // Check if user has admin role
    const userRole = response?.user?.role || response?.role || ''
    const adminRoles = ['admin', 'developer']
    const isAdmin = adminRoles.includes(userRole.toLowerCase())

    if (!isAdmin) {
      throw new Error('User is not admin')
    }

    // Admin access confirmed - allow navigation
    return

  } catch (_error) {
    // If not authenticated or not admin, redirect to login
    const loginUrl = `/auth/login?redirect=${encodeURIComponent(to.path)}`

    // Use navigateTo with external option to prevent SSR of protected content
    // @ts-ignore
    return navigateTo(loginUrl, {
      redirectCode: 302,
      external: false
    })
  }
})