<template>
  <div class="product-carousel-create-page min-h-screen bg-gray-50">
    <!-- هدر صفحه -->
    <div class="bg-white shadow-sm border-b">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex items-center justify-between h-16">
          <div class="flex items-center">
            <NuxtLink to="/admin/content/banners" class="text-gray-500 hover:text-gray-700 mr-4">
              <i class="fas fa-arrow-right"></i>
            </NuxtLink>
            <h1 class="text-xl font-semibold text-gray-900">ایجاد ویجت کاروسل محصولات</h1>
          </div>
          <div class="flex items-center space-x-4">
            <button
              @click="saveAsDraft"
              :disabled="isSaving"
              class="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-purple-500"
            >
              ذخیره پیش‌نویس
            </button>
            <button
              @click="saveWidget"
              :disabled="isSaving"
              class="px-4 py-2 text-sm font-medium text-white bg-purple-600 border border-transparent rounded-md hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-purple-500"
            >
              {{ isSaving ? 'در حال ذخیره...' : 'ذخیره و انتشار' }}
            </button>
          </div>
        </div>
      </div>
    </div>

    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
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
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import type { ProductCarouselConfig } from '~/types/widget'
import { useWidgetRegistry } from '~/composables/useWidgetRegistry'
import { useWidget } from '~/composables/useWidget'
import ProductCarouselSettings from './components/ProductCarouselSettings.vue'
import ProductCarouselPreview from './components/ProductCarouselPreview.vue'
import { useRouter } from 'vue-router'

// تعریف definePageMeta و navigateTo برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
declare const navigateTo: (to: string) => Promise<void>

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

const router = useRouter()

// استفاده از composables
const { createWidget } = useWidget()

// State
const isSaving = ref(false)
const categories = ref<any[]>([])

// فرم داده‌ها
const formData = ref({
  title: '',
  description: '',
  type: 'product-carousel' as const,
  status: 'active' as const,
  page: 'home' as const
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

    // ایجاد ویجت
    const createdWidget = await createWidget(widgetData)

    if (createdWidget) {
      // هدایت به صفحه ویرایش ویجت
      await router.push(`/admin/content/banners/ProductCarousel/edit/${createdWidget.id}`)
    }

  } catch (error: any) {
    console.error('خطا در ذخیره ویجت:', error)

    // نمایش پیام خطای مناسب به کاربر
    let errorMessage = 'خطا در ذخیره ویجت. لطفاً دوباره تلاش کنید.'

    if (error?.statusCode === 400) {
      errorMessage = 'داده‌های وارد شده نامعتبر است. لطفاً بررسی کنید.'
    } else if (error?.statusCode === 401) {
      errorMessage = 'دسترسی غیرمجاز. لطفاً دوباره وارد شوید.'
    } else if (error?.statusCode === 500) {
      errorMessage = 'خطای سرور. لطفاً بعداً تلاش کنید.'
    } else if (error?.message) {
      errorMessage = error.message
    }

    alert(errorMessage)
  } finally {
    isSaving.value = false
  }
}

const saveAsDraft = async () => {
  try {
    isSaving.value = true

    // ذخیره به عنوان پیش‌نویس
    const widgetData = {
      title: formData.value.title?.trim() || 'پیش‌نویس ویجت کاروسل محصولات',
      description: formData.value.description?.trim() || '',
      type: formData.value.type,
      status: 'draft' as const,
      page: formData.value.page,
      config: {
        ...config.value,
        title: config.value.title?.trim() || 'محصولات پیشنهادی',
        subtitle: config.value.subtitle?.trim() || ''
      }
    }

    await createWidget(widgetData)
    alert('پیش‌نویس با موفقیت ذخیره شد')

  } catch (error: any) {
    console.error('خطا در ذخیره پیش‌نویس:', error)

    // نمایش پیام خطای مناسب به کاربر
    let errorMessage = 'خطا در ذخیره پیش‌نویس. لطفاً دوباره تلاش کنید.'

    if (error?.statusCode === 400) {
      errorMessage = 'داده‌های وارد شده نامعتبر است. لطفاً بررسی کنید.'
    } else if (error?.statusCode === 401) {
      errorMessage = 'دسترسی غیرمجاز. لطفاً دوباره وارد شوید.'
    } else if (error?.statusCode === 500) {
      errorMessage = 'خطای سرور. لطفاً بعداً تلاش کنید.'
    } else if (error?.message) {
      errorMessage = error.message
    }

    alert(errorMessage)
  } finally {
    isSaving.value = false
  }
}

// Methods
const loadCategories = async () => {
  try {
    console.log('در حال بارگذاری دسته‌بندی‌ها...')
    const response = await $fetch('/api/admin/product-categories') as any
    console.log('پاسخ دسته‌بندی‌ها:', response)

    if (response?.data) {
      categories.value = response.data
      console.log('دسته‌بندی‌ها بارگذاری شد:', categories.value.length, 'دسته‌بندی')
    } else if (Array.isArray(response)) {
      categories.value = response
      console.log('دسته‌بندی‌ها بارگذاری شد:', categories.value.length, 'دسته‌بندی')
    }
  } catch (error) {
    console.error('خطا در بارگذاری دسته‌بندی‌ها:', error)
    alert('خطا در بارگذاری دسته‌بندی‌ها. لطفاً صفحه را مجدداً بارگذاری کنید.')
  }
}

// Lifecycle
onMounted(() => {
  loadCategories()
})
</script>

<style scoped>
.product-carousel-create-page {
  background-color: #ffffff;
  min-height: 100vh;
  font-family: 'Vazir', 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}

/* Smooth transitions */
.product-carousel-create-page {
  transition: all 0.3s ease;
}

/* Custom scrollbar */
.product-carousel-create-page::-webkit-scrollbar {
  width: 8px;
}

.product-carousel-create-page::-webkit-scrollbar-track {
  background: #f1f5f9;
}

.product-carousel-create-page::-webkit-scrollbar-thumb {
  background: #cbd5e1;
  border-radius: 4px;
}

.product-carousel-create-page::-webkit-scrollbar-thumb:hover {
  background: #94a3b8;
}

/* Focus styles for accessibility */
.product-carousel-create-page *:focus {
  outline: 2px solid #3b82f6;
  outline-offset: 2px;
}

/* Loading animation */
.product-carousel-create-page.loading {
  pointer-events: none;
  opacity: 0.7;
}

/* Responsive typography */
@media (max-width: 768px) {
  .product-carousel-create-page {
    font-size: 14px;
  }
}

/* Dark mode support */
@media (prefers-color-scheme: dark) {
  .product-carousel-create-page {
    background-color: #0f172a;
    color: #e2e8f0;
  }
}
</style>
