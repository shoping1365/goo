// Composable برای تشخیص نوع دستگاه (دسکتاپ، موبایل، تبلت)
import { ref, computed, readonly, onUnmounted } from 'vue'

export const useDeviceDetection = () => {
     const userAgent = ref('')
     const windowWidth = ref(0)

     // تشخیص User-Agent در سمت کلاینت
     if (import.meta.client) {
          userAgent.value = navigator.userAgent
          windowWidth.value = window.innerWidth

          // گوش دادن به تغییرات اندازه پنجره
          const updateWindowWidth = () => {
               windowWidth.value = window.innerWidth
          }

          window.addEventListener('resize', updateWindowWidth)

          // پاک کردن event listener در unmount
          if (typeof window !== 'undefined') {
               onUnmounted(() => {
                    window.removeEventListener('resize', updateWindowWidth)
               })
          }
     }

     // تشخیص دستگاه موبایل بر اساس User-Agent
     const isMobile = computed(() => {
          if (!userAgent.value) return false

          const mobileKeywords = [
               'mobile', 'android', 'iphone', 'ipod', 'blackberry',
               'windows phone', 'opera mini', 'mobile safari', 'webos',
               'palm', 'symbian', 'kindle', 'silk', 'fennec', 'maemo'
          ]

          return mobileKeywords.some(keyword =>
               userAgent.value.toLowerCase().includes(keyword)
          )
     })

     // تشخیص دستگاه تبلت بر اساس User-Agent
     const isTablet = computed(() => {
          if (!userAgent.value) return false

          const tabletKeywords = ['ipad', 'tablet', 'kindle', 'silk']

          return tabletKeywords.some(keyword =>
               userAgent.value.toLowerCase().includes(keyword)
          )
     })

     // تشخیص دستگاه دسکتاپ
     const isDesktop = computed(() => {
          return !isMobile.value && !isTablet.value
     })

     // تشخیص بر اساس عرض صفحه (برای responsive design)
     const isMobileByWidth = computed(() => {
          return windowWidth.value <= 768
     })

     const isTabletByWidth = computed(() => {
          return windowWidth.value > 768 && windowWidth.value <= 1024
     })

     const isDesktopByWidth = computed(() => {
          return windowWidth.value > 1024
     })

     // تشخیص نهایی (ترجیح User-Agent بر عرض صفحه)
     const deviceType = computed(() => {
          if (isMobile.value) return 'mobile'
          if (isTablet.value) return 'tablet'
          if (isDesktop.value) return 'desktop'

          // اگر User-Agent تشخیص ندهد، از عرض صفحه استفاده کن
          if (isMobileByWidth.value) return 'mobile'
          if (isTabletByWidth.value) return 'tablet'
          return 'desktop'
     })

     return {
          userAgent: readonly(userAgent),
          windowWidth: readonly(windowWidth),
          isMobile,
          isTablet,
          isDesktop,
          isMobileByWidth,
          isTabletByWidth,
          isDesktopByWidth,
          deviceType
     }
}
