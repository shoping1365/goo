/**
 * API Configuration Utility
 * مدیریت مرکزی URL های API backend
 */

/**
 * normalize and validate a backend base URL read from config/env
 */
const sanitizeBaseUrl = (raw?: string | null): string | undefined => {
  if (!raw) return undefined

  const trimmed = raw.trim()
  if (!trimmed) return undefined

  const lowered = trimmed.toLowerCase()
  if (lowered === 'undefined' || lowered === 'null' || lowered === 'false') {
    return undefined
  }

  const hasProtocol = /^[a-z][a-z0-9+.-]*:\/\//i.test(trimmed)
  const candidate = hasProtocol ? trimmed : `http://${trimmed}`

  try {
    const url = new URL(candidate)
    url.search = ''
    url.hash = ''
    url.pathname = url.pathname.replace(/\/+$/, '')
    const normalized = url.toString().replace(/\/$/, '')
    return normalized
  } catch (error: any) {
    console.warn('⚠️ Ignoring invalid GO API base URL:', { raw, error: error?.message })
    return undefined
  }
}

/**
 * دریافت Base URL برای Go API Backend
 * @returns {string} Base URL of Go API
 */
export function getGoApiBaseUrl(): string {
  const config = useRuntimeConfig()

  const sources: Array<{ label: string; value?: string | null }> = [
    { label: 'config.public.goApiBase', value: config.public.goApiBase },
    { label: 'process.env.GO_API_BASE_URL', value: process.env.GO_API_BASE_URL },
    { label: 'process.env.NUXT_PUBLIC_API_BASE', value: process.env.NUXT_PUBLIC_API_BASE },
    { label: 'process.env.NUXT_PUBLIC_GO_API_BASE', value: process.env.NUXT_PUBLIC_GO_API_BASE }
  ]

  for (const source of sources) {
    const sanitized = sanitizeBaseUrl(source.value)
    if (sanitized) {
      if (source.label === 'config.public.goApiBase') {
        console.log('🔍 getGoApiBaseUrl - Using config.public.goApiBase:', sanitized)
      }
      return sanitized
    }
  }
  throw new Error('❌ GO API base URL is not configured. Please set NUXT_PUBLIC_GO_API_BASE or GO_API_BASE_URL.')
}

/**
 * دریافت Base URL برای WebSocket
 * @returns {string} Base URL of WebSocket
 */
export function getWsBaseUrl(): string {
  const config = useRuntimeConfig()

  // اولویت 1: تنظیمات عمومی Nuxt
  if (config.public.wsBase) {
    return config.public.wsBase
  }

  // اولویت 2: متغیر محیطی مستقیم
  if (process.env.NUXT_PUBLIC_WS_BASE) {
    return process.env.NUXT_PUBLIC_WS_BASE
  }

  // در صورت عدم وجود هیچ تنظیمی، خطا بده
  throw new Error('❌ WS_BASE_URL is not configured! Please set NUXT_PUBLIC_WS_BASE in your .env file')
}