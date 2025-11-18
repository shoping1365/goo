import { ref } from 'vue'

/**
 * Composable برای debounce کردن توابع
 * این composable تعداد اجرای توابع را کاهش می‌دهد و عملکرد را بهبود می‌بخشد
 */
export const useDebounce = <T extends (...args: any[]) => any>(
  fn: T,
  delay: number = 300
) => {
  const timeoutId = ref<NodeJS.Timeout | null>(null)

  const debouncedFn = (...args: Parameters<T>) => {
    if (timeoutId.value) {
      clearTimeout(timeoutId.value)
    }

    timeoutId.value = setTimeout(() => {
      fn(...args)
    }, delay)
  }

  const cancel = () => {
    if (timeoutId.value) {
      clearTimeout(timeoutId.value)
      timeoutId.value = null
    }
  }

  return {
    debouncedFn,
    cancel
  }
} 