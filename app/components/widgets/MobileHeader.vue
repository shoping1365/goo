<template>
  <header v-if="shouldShowMobileHeader && activeMobileHeader" :style="mobileHeaderStyle" class="mobile-sticky-header">
    <!-- نمایش هدر بر اساس نوع از دیتابیس -->
    <div class="mobile-header-content">
      
      <!-- تشخیص نوع هدر بر اساس نام -->
      <MinimalTemplate v-if="getHeaderType(activeMobileHeader) === 'minimal'" :header="activeMobileHeader" />
      
      <!-- هدر بنر -->
      <BannerTemplate v-else-if="getHeaderType(activeMobileHeader) === 'banner'" :header="activeMobileHeader" />
      
      <!-- هدر کلاسیک -->
      <ClassicTemplate v-else :header="activeMobileHeader" />

    </div>
  </header>
</template>

<script setup lang="ts">
import { useRoute, useState } from 'nuxt/app'
import { computed, defineAsyncComponent, onMounted, ref, watch } from 'vue'
import { useDeviceDetection } from '~/composables/useDeviceDetection'
import { usePublicMobileHeaders } from '~/composables/usePublicMobileHeaders'

// Import تمپلیت‌ها با defineAsyncComponent
const MinimalTemplate = defineAsyncComponent(() => import('~/pages/admin/content/mobile-app-header-management/templates/minimal-template.vue'))
const BannerTemplate = defineAsyncComponent(() => import('~/pages/admin/content/mobile-app-header-management/templates/banner-template.vue'))
const ClassicTemplate = defineAsyncComponent(() => import('~/pages/admin/content/mobile-app-header-management/templates/classic-template.vue'))

const route = useRoute()
const { loadMobileHeaders, getActiveMobileHeader, shouldShowMobileHeader: checkShouldShowMobileHeader } = usePublicMobileHeaders()
const { isMobile } = useDeviceDetection()

// تشخیص نوع هدر بر اساس نام
interface Header {
  id?: number
  name?: string
  type?: string
  is_active?: boolean
  page_selection?: 'all' | 'specific' | 'exclude'
  specific_pages?: string | null
  excluded_pages?: string | null
  description?: string
  [key: string]: unknown
}

const activeMobileHeader = ref<Header | null>(null)

// بررسی اینکه آیا هدر موبایل باید نمایش داده شود یا نه
const shouldShowMobileHeader = computed(() => {
  // هدر موبایل فقط در دستگاه‌های موبایل نمایش داده می‌شود
  if (!isMobile.value) {
    return false
  }

  // استفاده از state میدل‌ور
  const mobileHeaderDisplayState = useState('mobileHeaderDisplayState', () => ({
    shouldShow: true,
    headerData: null
  }))
  
  // اگر state میدل‌ور موجود باشد، از آن استفاده کن
  if (mobileHeaderDisplayState.value && mobileHeaderDisplayState.value.shouldShow !== undefined) {
    return mobileHeaderDisplayState.value.shouldShow
  }
  
  // بررسی بر اساس صفحه فعلی
  const currentPath = route.path
  return checkShouldShowMobileHeader(currentPath)
})

const mobileHeaderStyle = computed((): Record<string, string | number> => {
  return {}
})
const getHeaderType = (header: Header) => {
  if (!header?.name) return 'classic'
  
  const name = header.name.toLowerCase()
  const description = header.description?.toLowerCase() || ''
  
  if (name.includes('مینیمال') || description.includes('مینیمال')) {
    return 'minimal'
  } else if (name.includes('بنر') || description.includes('بنر')) {
    return 'banner'
  } else {
    return 'classic'
  }
}

// بارگذاری هدرهای موبایل
onMounted(async () => {
  try {
    await loadMobileHeaders()
    const header = getActiveMobileHeader.value
    if (header) {
      const headerRecord = header as unknown as Record<string, unknown>
      activeMobileHeader.value = {
        id: headerRecord.id as number | undefined,
        name: headerRecord.name as string | undefined,
        type: headerRecord.type as string | undefined,
        is_active: headerRecord.is_active as boolean | undefined,
        page_selection: headerRecord.page_selection as 'all' | 'specific' | 'exclude' | undefined,
        specific_pages: headerRecord.specific_pages as string | undefined,
        excluded_pages: headerRecord.excluded_pages as string | undefined,
        description: headerRecord.description as string | undefined,
        ...headerRecord
      } as Header
    } else {
      activeMobileHeader.value = null
    }
  } catch (error) {
    console.error('خطا در بارگذاری هدرهای موبایل:', error)
  }
})

// نظارت بر تغییرات هدر فعال
watch(getActiveMobileHeader, (newHeader) => {
  if (newHeader) {
    const headerRecord = newHeader as unknown as Record<string, unknown>
    activeMobileHeader.value = {
      id: headerRecord.id as number | undefined,
      name: headerRecord.name as string | undefined,
      type: headerRecord.type as string | undefined,
      is_active: headerRecord.is_active as boolean | undefined,
      page_selection: headerRecord.page_selection as 'all' | 'specific' | 'exclude' | undefined,
      specific_pages: headerRecord.specific_pages as string | undefined,
      excluded_pages: headerRecord.excluded_pages as string | undefined,
      description: headerRecord.description as string | undefined,
      ...headerRecord
    } as Header
  } else {
    activeMobileHeader.value = null
  }
}, { immediate: true })

// نظارت بر تغییرات state میدلور
const mobileHeaderDisplayState = useState('mobileHeaderDisplayState', () => ({
  shouldShow: true,
  headerData: null
}))

watch(mobileHeaderDisplayState, (newState) => {
  if (newState?.headerData) {
    activeMobileHeader.value = newState.headerData
  }
}, { deep: true })
</script>

<style scoped>


/* Responsive */
@media (min-width: 769px) {
  .mobile-sticky-header {
    display: none;
  }
}
</style>