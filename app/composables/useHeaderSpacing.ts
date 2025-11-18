import { ref, computed, watch } from 'vue'

// متغیر global برای نگهداری ارتفاع هدر
const globalHeaderHeight = ref(0)

export const useHeaderSpacing = () => {
  const headerHeight = ref(0)
  
  // محاسبه فاصله از هدر
  const mainMarginTop = computed(() => {
    return `${headerHeight.value}px`
  })
  
  // به‌روزرسانی ارتفاع هدر
  const updateHeaderHeight = (height: number) => {
    headerHeight.value = height
    globalHeaderHeight.value = height
  }
  
  // دریافت ارتفاع فعلی هدر
  const getCurrentHeaderHeight = () => {
    return headerHeight.value
  }
  
  return {
    headerHeight,
    mainMarginTop,
    updateHeaderHeight,
    getCurrentHeaderHeight
  }
}

// Export برای استفاده در layout
export const getGlobalHeaderHeight = () => globalHeaderHeight