import { useNuxtApp } from 'nuxt/app'

type Level = 'success' | 'error' | 'warning' | 'info'

function emitLastErrorEvent(title: string, message: string, meta?: string) {
  if (typeof window === 'undefined') return
  try {
    ; (window as any).__lastError = { title, message, meta, ts: Date.now() }
    window.dispatchEvent(new CustomEvent('app:last-error', { detail: { title, message, meta } }))
  } catch { }
}

export function useNotifier() {
  const { $toast } = useNuxtApp() as any
  const getToast = () => {
    const t = (typeof window !== 'undefined' && (($toast as any) || (window as any).$toast)) as any
    return t && typeof t.success === 'function' ? t : null
  }

  const show = (level: Level, message?: string, title?: string, meta?: string) => {
    const toast = getToast()
    const ttl = 6000
    const text = `${title ? title + ': ' : ''}${message || ''}`.trim() || 'خطای نامشخص'
    if (toast) {
      const opts: any = { autoClose: ttl }
      if (level === 'success') toast.success(text, opts)
      else if (level === 'error') toast.error(text, opts)
      else if (level === 'warning' && typeof toast.warning === 'function') toast.warning(text, opts)
      else toast.info(text, opts)
    }
    if (level === 'error' || level === 'warning') {
      emitLastErrorEvent(title || (level === 'error' ? 'خطا' : 'هشدار'), message || 'خطای نامشخص', meta)
    }
  }

  return {
    success: (message: string, title = 'موفقیت') => show('success', message, title),
    error: (message: string, title = 'خطا', meta?: string) => show('error', message, title, meta),
    warning: (message: string, title = 'هشدار', meta?: string) => show('warning', message, title, meta),
    info: (message: string, title = 'اطلاع‌رسانی') => show('info', message, title)
  }
}


