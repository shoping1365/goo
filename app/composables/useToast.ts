import { ref, readonly } from 'vue'

interface Toast {
  id: number
  message: string
  type: 'success' | 'error' | 'warning' | 'info'
  duration: number
}

const toasts = ref<Toast[]>([])
let nextId = 1

export const useToast = () => {
  const addToast = (message: string, type: 'success' | 'error' | 'warning' | 'info' = 'info', duration: number = 3000) => {
    const toast: Toast = {
      id: nextId++,
      message,
      type,
      duration
    }
    toasts.value.push(toast)
    
    // حذف خودکار بعد از مدت زمان مشخص
    if (duration > 0) {
      setTimeout(() => {
        removeToast(toast.id)
      }, duration)
    }
  }

  const removeToast = (id: number) => {
    const index = toasts.value.findIndex(toast => toast.id === id)
    if (index > -1) {
      toasts.value.splice(index, 1)
    }
  }

  const showSuccess = (message: string, duration?: number) => {
    addToast(message, 'success', duration)
  }

  const showError = (message: string, duration?: number) => {
    addToast(message, 'error', duration)
  }

  const showWarning = (message: string, duration?: number) => {
    addToast(message, 'warning', duration)
  }

  const showInfo = (message: string, duration?: number) => {
    addToast(message, 'info', duration)
  }

  return {
    toasts: readonly(toasts),
    addToast,
    removeToast,
    showSuccess,
    showError,
    showWarning,
    showInfo
  }
}