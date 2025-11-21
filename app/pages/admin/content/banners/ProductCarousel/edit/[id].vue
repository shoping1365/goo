<template>
  <div class="product-carousel-edit-page min-h-screen bg-gray-50">
    <!-- هدر صفحه -->
    <div class="bg-white shadow-sm border-b">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex items-center justify-between h-16">
          <div class="flex items-center">
            <NuxtLink to="/admin/content/banners" class="text-gray-500 hover:text-gray-700 mr-4">
              <i class="fas fa-arrow-right"></i>
            </NuxtLink>
            <h1 class="text-xl font-semibold text-gray-900">ویرایش ویجت کاروسل محصولات</h1>
          </div>
          <div class="flex items-center space-x-4">
            <button
              :disabled="isSaving"
              class="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-purple-500"
              @click="saveAsDraft"
            >
              ذخیره پیش‌نویس
            </button>
            <button
              :disabled="isSaving"
              class="px-4 py-2 text-sm font-medium text-white bg-purple-600 border border-transparent rounded-md hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-purple-500"
              @click="saveWidget"
            >
              {{ isSaving ? 'در حال ذخیره...' : 'ذخیره و انتشار' }}
            </button>
          </div>
        </div>
      </div>
    </div>

    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <!-- نمایش loading -->
      <div v-if="loading" class="text-center py-8">
        <i class="fas fa-spinner fa-spin text-2xl text-gray-400"></i>
        <p class="text-gray-600 mt-2">در حال بارگذاری ویجت...</p>
      </div>

      <!-- نمایش خطا -->
      <div v-else-if="error" class="bg-red-50 border border-red-200 rounded-md p-6 mb-6">
        <div class="flex">
          <i class="fas fa-exclamation-triangle text-red-400 mt-1"></i>
          <div class="mr-3">
            <h3 class="text-sm font-medium text-red-800">خطا در بارگذاری ویجت</h3>
            <p class="text-sm text-red-700 mt-1">{{ error }}</p>
          </div>
        </div>
      </div>

      <!-- فرم ویرایش -->
      <div v-else-if="widget">
        <!-- تنظیمات ویجت -->
        <ProductCarouselSettings
          :form-data="formData"
          :config="config"
          :categories="categories"
        />

        <!-- پیش‌نمایش ویجت -->
        <div class="mt-8">
          <ProductCarouselPreview :config="config" />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'
import { useWidget } from '~/composables/useWidget'
import type { ProductCarouselConfig, Widget, WidgetType, WidgetStatus, WidgetPage, CategoryItem } from '~/types/widget'
import ProductCarouselPreview from '../create/components/ProductCarouselPreview.vue'
import ProductCarouselSettings from '../create/components/ProductCarouselSettings.vue'

// تعریف definePageMeta برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// دریافت ID از route
const route = useRoute()

// استفاده از composables
const { fetchWidget, updateWidget } = useWidget()
const widgetId = parseInt(route.params.id as string)

// State
const loading = ref(true)
const isSaving = ref(false)
const error = ref<string | null>(null)
const widget = ref<Widget | null>(null)
const categories = ref<CategoryItem[]>([])

// فرم داده‌ها
const formData = ref({
  title: '',
  description: '',
  type: 'product-carousel' as WidgetType,
  status: 'active' as WidgetStatus,
  page: 'home' as WidgetPage
})

// تنظیمات ویجت
const config = ref<ProductCarouselConfig>({
  title: 'محصولات پیشنهادی',
  subtitle: 'بهترین محصولات برای شما',
  categoryId: null,
  productCount: 12,
  slidesPerView: 4,
  showPrice: true,
  showRating: true,
  showDiscount: true,
  autoplay: true,
  autoplayDelay: 3000,
  wide_bg: false,
  bg_color: '#ffffff',
  padding: '24px',
  margin: '0px',
  showNavigation: true,
  showIndicators: true,
  showControls: false,
  navigationStyle: 'default',
  navigationSize: 40,
  indicatorStyle: 'default',
  indicatorColor: '#3b82f6'
})

// بارگذاری ویجت
const loadWidget = async () => {
  try {
    loading.value = true
    error.value = null
    
    const widgetData = await fetchWidget(widgetId)
    widget.value = widgetData
    
    // پر کردن فرم داده‌ها
    formData.value = {
      title: widgetData.title,
      description: widgetData.description || '',
      type: widgetData.type,
      status: widgetData.status,
      page: widgetData.page
    }
    
    // پر کردن تنظیمات ویجت
    if (widgetData.config) {
      const parsedConfig = typeof widgetData.config === 'string' 
        ? JSON.parse(widgetData.config) 
        : widgetData.config
      
      config.value = {
        ...config.value,
        ...parsedConfig
      }
    }
    
  } catch (err: unknown) {
    const errorMessage = err instanceof Error ? err.message : 'خطا در بارگذاری ویجت'
    error.value = errorMessage
    console.error('خطا در بارگذاری ویجت:', err)
  } finally {
    loading.value = false
  }
}

// ذخیره ویجت
const saveWidget = async () => {
  try {
    isSaving.value = true

    // اعتبارسنجی
    if (!formData.value.title.trim()) {
      alert('لطفاً عنوان ویجت را وارد کنید')
      return
    }

    if (!config.value.productCount || config.value.productCount <= 0) {
      alert('لطفاً تعداد محصولات را به درستی مشخص کنید')
      return
    }

    if (!config.value.slidesPerView || config.value.slidesPerView <= 0) {
      alert('لطفاً تعداد اسلاید در هر ویو را به درستی مشخص کنید')
      return
    }

    // آماده‌سازی داده‌ها برای ذخیره
    const widgetData = {
      title: formData.value.title.trim(),
      description: formData.value.description.trim(),
      type: formData.value.type,
      status: formData.value.status,
      page: formData.value.page,
      config: {
        ...config.value,
        title: config.value.title?.trim() || 'محصولات پیشنهادی',
        subtitle: config.value.subtitle?.trim() || ''
      }
    }

    // به‌روزرسانی ویجت
    const updatedWidget = await updateWidget(widgetId, widgetData)

    if (updatedWidget) {
      // نمایش پیام موفقیت
      alert('ویجت با موفقیت به‌روزرسانی شد')
      
      // به‌روزرسانی داده‌های محلی
      widget.value = updatedWidget
    }

  } catch (err: unknown) {
    const errorMessage = err instanceof Error ? err.message : 'خطای نامشخص'
    console.error('خطا در ذخیره ویجت:', err)
    alert('خطا در ذخیره ویجت: ' + errorMessage)
  } finally {
    isSaving.value = false
  }
}

// ذخیره پیش‌نویس
const saveAsDraft = async () => {
  formData.value.status = 'draft'
  await saveWidget()
}

// بارگذاری اولیه
onMounted(() => {
  loadWidget()
})
</script>

