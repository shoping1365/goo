import { ref, onUnmounted } from 'vue'

/**
 * Composable برای مدیریت صحیح interval ها و جلوگیری از نشت حافظه
 * این composable تضمین می‌کند که interval ها در زمان unmount کامپوننت پاک شوند
 */
export const useInterval = () => {
  const intervals = ref<Set<NodeJS.Timeout>>(new Set())

  /**
   * ایجاد interval جدید با مدیریت خودکار
   * @param callback تابعی که باید اجرا شود
   * @param delay تاخیر به میلی‌ثانیه
   * @returns شناسه interval
   */
  const setInterval = (callback: () => void, delay: number): NodeJS.Timeout => {
    const intervalId = globalThis.setInterval(callback, delay)
    intervals.value.add(intervalId)
    return intervalId
  }

  /**
   * پاکسازی interval خاص
   * @param intervalId شناسه interval
   */
  const clearInterval = (intervalId: NodeJS.Timeout): void => {
    if (intervals.value.has(intervalId)) {
      globalThis.clearInterval(intervalId)
      intervals.value.delete(intervalId)
    }
  }

  /**
   * پاکسازی همه interval ها
   */
  const clearAllIntervals = (): void => {
    intervals.value.forEach(intervalId => {
      globalThis.clearInterval(intervalId)
    })
    intervals.value.clear()
  }

  /**
   * بررسی اینکه آیا interval خاصی فعال است
   * @param intervalId شناسه interval
   * @returns true اگر interval فعال باشد
   */
  const hasInterval = (intervalId: NodeJS.Timeout): boolean => {
    return intervals.value.has(intervalId)
  }

  /**
   * تعداد interval های فعال
   */
  const activeIntervalsCount = ref(0)

  // به‌روزرسانی تعداد interval های فعال
  const updateActiveCount = () => {
    activeIntervalsCount.value = intervals.value.size
  }

  // پاکسازی خودکار در زمان unmount
  onUnmounted(() => {
    clearAllIntervals()
  })

  return {
    setInterval,
    clearInterval,
    clearAllIntervals,
    hasInterval,
    activeIntervalsCount,
    updateActiveCount
  }
} 