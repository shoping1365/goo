import { useNuxtApp } from 'nuxt/app'

type Level = 'success' | 'error' | 'warning' | 'info'

interface ToastOptions {
  autoClose: number
}

interface Toast {
  success: (msg: string, opts?: ToastOptions) => void
  error: (msg: string, opts?: ToastOptions) => void
  warning?: (msg: string, opts?: ToastOptions) => void
  info: (msg: string, opts?: ToastOptions) => void
}

interface CustomWindow extends Window {
  __lastError?: { title: string; message: string; meta?: string; ts: number }
  $toast?: Toast
}

function emitLastErrorEvent(title: string, message: string, meta?: string) {
  if (typeof window === 'undefined') return
  try {
    const win = window as unknown as CustomWindow
    win.__lastError = { title, message, meta, ts: Date.now() }
    window.dispatchEvent(new CustomEvent('app:last-error', { detail: { title, message, meta } }))
  } catch { }
}

export function useNotifier() {
  const nuxtApp = useNuxtApp()
  const $toast = nuxtApp.$toast as Toast | undefined

  const getToast = (): Toast | null => {
    if (typeof window === 'undefined') return null
    const win = window as unknown as CustomWindow
    const t = $toast || win.$toast
    return t && typeof t.success === 'function' ? t : null
  }

  const show = (level: Level, message?: string, title?: string, meta?: string) => {
    const toast = getToast()
    const ttl = 6000
    const text = `${title ? title + ': ' : ''}${message || ''}`.trim() || 'خطای نامشخص'
    if (toast) {
      const opts = { autoClose: ttl }
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


