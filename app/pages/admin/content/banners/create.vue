<template>
  <div class="min-h-screen w-full px-0 py-12">
    <!-- تایتل بازگشت به بنرها -->
    <div class="w-full bg-white px-8 py-4 flex items-center justify-start shadow-sm mb-8">
      <NuxtLink to="/admin/content/banners" class="text-blue-600 font-bold text-lg hover:underline">
        بازگشت به ابزارک‌ها
      </NuxtLink>
    </div>

    <div class="max-w-6xl mx-auto">
      <div class="bg-white rounded-xl shadow p-8">
        <div class="flex justify-between items-center mb-8">
          <h2 class="text-xl font-bold">ایجاد ابزارک جدید</h2>
        </div>

        <!-- نمایش خطا -->
        <div v-if="error" class="mb-6 p-6 bg-red-50 border border-red-200 rounded-lg">
          <p class="text-red-800">{{ error }}</p>
        </div>

        <!-- فرم اولیه -->
        <form class="flex items-end gap-6" @submit.prevent="handleWidgetCreate">
          <!-- نوع ابزارک -->
          <div class="flex-1">
            <label class="block mb-2 text-sm font-medium text-gray-700">نوع ابزارک</label>
            <select
              v-model="formData.type"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
              required
            >
              <option value="">انتخاب کنید</option>
              <option v-for="(label, type) in WIDGET_TYPE_LABELS" :key="type" :value="type">
                {{ label }}
              </option>
            </select>
          </div>

          <!-- بخش -->
          <div class="flex-1">
            <label class="block mb-2 text-sm font-medium text-gray-700">بخش</label>
            <select
              v-model="formData.page"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
            >
              <option value="home">صفحه اصلی</option>
              <option value="other">سایر بخش‌ها</option>
            </select>
          </div>

          <!-- دکمه -->
          <div class="flex-shrink-0">
            <button
              type="submit"
              :disabled="loading"
              class="bg-purple-600 text-white px-6 py-2 rounded-md font-bold hover:bg-purple-700 transition-colors disabled:opacity-50 h-11"
            >
              {{ loading ? 'در حال ایجاد...' : 'ایجاد ابزارک' }}
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { WIDGET_TYPE_LABELS } from '@/types/widget'
import { useRouter } from 'vue-router'
import { useAuth } from '~/composables/useAuth'

// تعریف definePageMeta برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string }) => void

const router = useRouter()

definePageMeta({
  layout: 'admin-main'
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user: _user, hasPermission: _hasPermission } = useAuth()

// State
const loading = ref(false)
const error = ref('')
const formData = ref({
  type: '',
  page: 'home'
})

// Methods
const handleWidgetCreate = async () => {
  try {
    loading.value = true
    error.value = ''

    if (!formData.value.type) {
      error.value = 'لطفاً نوع ابزارک را انتخاب کنید.'
      return
    }

    // Navigate to the specific create page based on type
    const routes = {
      'main-slider-side-banner': '/admin/content/banners/MainSliderSideBanner/create',
      'single-slider-side': '/admin/content/banners/SingleSliderSide/create',
      'full-banner': '/admin/content/banners/FullBanner/create',

      'double-banner': '/admin/content/banners/doublebanner/create',
      'triple-banner': '/admin/content/banners/triplebanner/create',
      'quad-banner': '/admin/content/banners/quadbanner/create',
      'penta-banner': '/admin/content/banners/pentabanner/create',

      'brands-slider': '/admin/content/banners/BrandsSlider/create',
      'category-box': '/admin/content/banners/CategoryBox/create',
      'categories': '/admin/content/banners/CategoryBox/create',
      'product-carousel': '/admin/content/banners/ProductCarousel/create',
      'products': '/admin/content/banners/Products/create',
      'products-slider': '/admin/content/banners/ProductsSlider/create',
      'services-slider': '/admin/content/banners/ServicesSlider/create'
    }

    const route = routes[formData.value.type as keyof typeof routes]
    if (route) {
      await router.push(route)
    } else {
      error.value = 'نوع ابزارک انتخاب شده پشتیبانی نمی‌شود.'
    }
  } catch (_err) {
    // خطا در ایجاد ابزارک
    error.value = 'خطا در ایجاد ابزارک. لطفاً دوباره تلاش کنید.'
    loading.value = false
  }
}
</script>

