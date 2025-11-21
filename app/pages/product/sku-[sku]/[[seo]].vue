<template>
  <div class="min-h-screen bg-white" dir="rtl">
    <!-- Breadcrumb -->
    <div v-if="product" class="bg-white shadow-sm">
      <div class="container mx-auto px-4 py-3">
        <nav class="flex items-center gap-2 text-sm text-gray-600">
          <NuxtLink to="/" class="hover:text-blue-600 transition-colors">خانه</NuxtLink>
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"></path>
          </svg>
          <NuxtLink 
            v-if="product.category && typeof product.category === 'object'" 
            :to="`/product-category/${product.category.slug}`"
            class="hover:text-blue-600 transition-colors truncate"
          >
            {{ product.category.name }}
          </NuxtLink>
          <template v-if="product.category">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"></path>
            </svg>
            <span class="text-gray-900 font-medium truncate">{{ product.name }}</span>
          </template>
        </nav>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="container mx-auto px-4 py-8">
      <div class="bg-white rounded-lg shadow-lg px-4 py-4">
        <div class="animate-pulse">
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
            <div class="space-y-4">
              <div class="h-64 bg-gray-200 rounded-lg"></div>
              <div class="flex gap-2">
                <div v-for="i in 4" :key="i" class="h-16 w-16 bg-gray-200 rounded-lg"></div>
              </div>
            </div>
            <div class="space-y-4">
              <div class="h-8 bg-gray-200 rounded"></div>
              <div class="h-4 bg-gray-200 rounded w-3/4"></div>
              <div class="h-6 bg-gray-200 rounded w-1/2"></div>
              <div class="h-12 bg-gray-200 rounded"></div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Error State -->
    <div v-else-if="error" class="container mx-auto px-4 py-8">
      <div class="bg-white rounded-lg shadow-lg px-4 py-4 text-center">
        <div class="text-red-500 mb-4">
          <svg class="w-16 h-16 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L4.081 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
          </svg>
        </div>
        <h2 class="text-xl font-bold text-gray-900 mb-2">محصول یافت نشد</h2>
        <p class="text-gray-600 mb-4">{{ error }}</p>
        <NuxtLink to="/" class="inline-flex items-center gap-2 bg-blue-600 text-white px-6 py-3 rounded-lg hover:bg-blue-700 transition-colors">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path>
          </svg>
          بازگشت به خانه
        </NuxtLink>
      </div>
    </div>

    <!-- Product Content -->
    <div v-else-if="product" class="container mx-auto px-4 py-6">
      <div class="bg-white rounded-lg shadow-lg mb-6">
        <ProductMainInfo :product="product" @add-to-cart="addToCart" />
      </div>

      <!-- Tabs Navigation -->
      <div class="bg-white rounded-t-lg shadow-lg">
        <div class="border-b border-gray-200">
          <nav class="flex gap-8 px-6">
            <button
              v-for="tab in tabs"
              :key="tab.value"
              :class="[
                'py-4 px-2 border-b-2 font-medium text-sm transition-colors',
                activeTab === tab.value
                  ? 'border-blue-500 text-blue-600'
                  : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
              ]"
              @click="activeTab = tab.value"
            >
              <div class="flex items-center gap-2">
                {{ tab.label }}
              </div>
            </button>
          </nav>
        </div>

        <!-- Tab Content -->
        <div class="px-4 py-4">
          <component
            :is="currentTabComponent"
            :product="product"
            :product-id="product?.id"
          />
        </div>
      </div>
    </div>

    <!-- Floating Actions (same as اصلی) -->
    <div class="fixed bottom-6 left-6 flex flex-col gap-3 z-40">
      <button
        v-if="product"
        :disabled="!(product.in_stock ?? product.stock_quantity > 0)"
        class="bg-green-600 hover:bg-green-700 disabled:bg-gray-400 text-white px-4 py-4 rounded-full shadow-lg hover:shadow-xl transition-all duration-200 group"
        :title="(product.in_stock ?? product.stock_quantity > 0) ? 'افزودن به سبد خرید' : 'ناموجود'"
        @click="addToCart"
      >
        <svg class="w-6 h-6 group-hover:scale-110 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4m-2.4 0L3 3m4 10v6a1 1 0 001 1h8a1 1 0 001-1v-6m-9 0V9a1 1 0 011-1h6a1 1 0 011 1v4.01"></path>
        </svg>
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
// Force TS mode
import ProductQnA from '~/components/product/ProductQnA.vue'
import ProductVideos from '~/components/product/ProductVideos.vue'
import ProductComplements from '../ProductComplements.vue'
import ProductDescription from '../ProductDescription.vue'
import ProductFAQ from '../ProductFAQ.vue'
import ProductMainInfo from '../ProductMainInfo.vue'
import ProductRelated from '../ProductRelated.vue'
import ProductReviews from '../ProductReviews.vue'
import ProductSpecifications from '../ProductSpecifications.vue'

import { computed, onBeforeUnmount, onMounted, reactive, ref, watch, type Component, type Ref } from 'vue'
import { useCart } from '~/composables/useCart'
import type { Product } from '~/types/product'

// Explicitly declare Nuxt auto-imports to satisfy TS
declare const definePageMeta: (meta: Record<string, unknown>) => void
declare const useRoute: () => { params: Record<string, string>; query: Record<string, string> }
declare const useAsyncData: <T>(key: string, handler: () => Promise<T>, options?: Record<string, unknown>) => Promise<{ data: Ref<T | null>; error: Ref<Error | null>; pending: Ref<boolean> }>
declare const useRuntimeConfig: () => { public: { siteUrl: string } }
declare const useHead: (meta: () => Record<string, unknown>) => void
declare const navigateTo: (to: string) => Promise<void>
declare const $fetch: <T>(url: string, options?: Record<string, unknown>) => Promise<T>

definePageMeta({ layout: 'default' })

const route = useRoute()
// شناسه برای فراخوانی API: اگر پارامتر با پیشوند 'sku-' آمد، آن را حذف می‌کنیم
const identifier = computed(() => {
  const rawParam = route.params.sku ?? route.params.id ?? route.query.id ?? ''
  const raw = String(rawParam)
  if (!raw) {
    return ''
  }
  if (raw.startsWith('sku-')) {
    return raw.substring(4)
  }
  return raw
})

// استفاده از useAsyncData برای SSR صحیح
const { data: product, error: fetchError, pending: loading } = await useAsyncData<Product>(
  `product-${identifier.value}`,
  async () => {
    if (!identifier.value) {
      throw new Error('شناسه محصول معتبر نیست')
    }
    try {
      const isPreview = route.query.preview === '1' || route.query.preview === 'true'
      const apiUrl = isPreview ? `/api/products/${identifier.value}?preview=1` : `/api/products/${identifier.value}`
      const response = await $fetch<Product>(apiUrl)
      if (!response) throw new Error('محصول یافت نشد')
      return response
    } catch (err) {
      const e = err as { message?: string }
      throw new Error(e.message || 'خطا در بارگذاری محصول')
    }
  },
  {
    watch: [identifier]
  }
)

const error = computed(() => fetchError.value?.message || '')
const activeTab = ref('description')

// SEO
const config = useRuntimeConfig()
const baseUrl = process.client ? window.location.origin : config.public.siteUrl

useHead(() => ({
  title: product.value ? `${product.value.name} - فروشگاه` : 'محصول - فروشگاه',
  meta: !product.value ? [] : [
    // متا توضیحات از فیلد meta_description یا description محصول
    {
      name: 'description',
      content: product.value.meta_description || product.value.description || `${product.value.name} - فروشگاه`
    },
    // سایر متاتگ‌های OG برای شبکه‌های اجتماعی
    {
      property: 'og:title',
      content: product.value.seo_title || product.value.name
    },
    {
      property: 'og:description',
      content: product.value.meta_description || product.value.description || `${product.value.name} - فروشگاه`
    },
    {
      property: 'og:type',
      content: 'product'
    },
    {
      property: 'og:url',
      content: `${baseUrl}/product/sku-${product.value.sku}${product.value.slug ? '/' + product.value.slug : ''}`
    },
    // تصویر محصول برای نمایش در شبکه‌های اجتماعی
    ...(product.value.image ? [{
      property: 'og:image',
      content: product.value.image.startsWith('http') ? product.value.image : `${baseUrl}${product.value.image}`
    }] : [])
  ],
  // افزودن لینک canonical
  link: !product.value ? [] : [
    {
      rel: 'canonical',
      href: `${baseUrl}/product/sku-${product.value.sku}${product.value.slug ? '/' + product.value.slug : ''}`
    }
  ]
}))

const toast = reactive({ show: false, message: '', type: 'success' })

const tabs = [
  { label: 'توضیحات', value: 'description' },
  { label: 'ویدیوها', value: 'videos' },
  { label: 'مشخصات فنی', value: 'specifications' },
  { label: 'نظرات', value: 'reviews' },
  { label: 'سوالات متداول', value: 'faq' },
  { label: 'پرسش و پاسخ', value: 'qna' },
  { label: 'محصولات مرتبط', value: 'related' },
  { label: 'محصولات مکمل', value: 'complements' }
]

const tabComponents: Record<string, Component> = {
  description: ProductDescription,
  videos: ProductVideos,
  reviews: ProductReviews,
  faq: ProductFAQ,
  related: ProductRelated,
  complements: ProductComplements,
  specifications: ProductSpecifications,
  qna: ProductQnA
}

const currentTabComponent = computed(() => tabComponents[activeTab.value] || tabComponents.description)

// ردیابی زمان ماندن در صفحه برای تحلیل بازاریابی
const viewStartTime = ref(null)
const currentProductId = ref(null)
let durationUpdateTimeout = null

// ثبت بازدید محصول (فقط برای کاربران لاگین شده)
const trackProductView = async (productId) => {
  if (!productId || !import.meta.client) return
  try {
    await $fetch(`/api/recent-views/product/${productId}`, {
      method: 'POST',
      credentials: 'include',
      // Silent fail: اگر کاربر لاگین نباشد یا خطایی رخ داد، چیزی نمایش ندهیم
      onResponseError: () => {}
    })
    // شروع ردیابی زمان برای این محصول
    viewStartTime.value = Date.now()
    currentProductId.value = productId
  } catch {
    // Silent fail - بازدید ثبت نشد ولی مشکلی ایجاد نمی‌کنیم
  }
}

// بروزرسانی مدت زمان ماندن در صفحه
const updateViewDuration = async () => {
  if (!currentProductId.value || !viewStartTime.value || !import.meta.client) return
  
  const duration = Math.floor((Date.now() - viewStartTime.value) / 1000)
  
  // فقط اگر بیشتر از 3 ثانیه مانده باشد، ذخیره کنیم
  if (duration < 3) return
  
  try {
    await $fetch(`/api/recent-views/product/${currentProductId.value}/duration`, {
      method: 'PATCH',
      body: { duration },
      credentials: 'include',
      onResponseError: () => {}
    })
  } catch {
    // Silent fail
  }
}

// ردیابی خروج از صفحه یا تغییر visibility
const handleVisibilityChange = () => {
  if (document.hidden) {
    // کاربر از صفحه رفت یا tab را عوض کرد
    updateViewDuration()
  }
}

const handleBeforeUnload = () => {
  // کاربر در حال بستن صفحه است
  updateViewDuration()
}

// SSR + CSR: useAsyncData handles both server and client
onMounted(async () => {
  // ثبت بازدید بعد از لود موفق محصول
  if (product.value?.id) {
    await trackProductView(product.value.id)
  }
  
  // اضافه کردن event listener ها برای ردیابی زمان
  if (import.meta.client) {
    document.addEventListener('visibilitychange', handleVisibilityChange)
    window.addEventListener('beforeunload', handleBeforeUnload)
    
    // هر 30 ثانیه یک بار duration را بروزرسانی کن (برای جلوگیری از از دست رفتن داده)
    durationUpdateTimeout = setInterval(() => {
      if (!document.hidden) {
        updateViewDuration()
      }
    }, 30000)
  }
})

// پاک‌سازی event listener ها
onBeforeUnmount(() => {
  if (import.meta.client) {
    updateViewDuration() // بروزرسانی نهایی
    document.removeEventListener('visibilitychange', handleVisibilityChange)
    window.removeEventListener('beforeunload', handleBeforeUnload)
    if (durationUpdateTimeout) {
      clearInterval(durationUpdateTimeout)
    }
  }
})

// وقتی محصول تغییر کرد، بازدید را ثبت کن
watch(() => product.value?.id, async (newId, oldId) => {
  if (newId && newId !== oldId) {
    // بروزرسانی duration برای محصول قبلی
    if (oldId) {
      await updateViewDuration()
    }
    // ثبت بازدید برای محصول جدید
    await trackProductView(newId)
  }
})

const { addToCart: addToCartComposable } = useCart()
const showToast = (message, type) => { toast.message = message; toast.type = type; toast.show = true; setTimeout(() => { toast.show = false }, 3000) }
const addToCart = async () => {
  if (!product.value?.id) { showToast('اطلاعات محصول در دسترس نیست', 'error'); return }
  try {
    const result = await addToCartComposable(product.value.id, 1, product.value)
    if (result?.success) {
      showToast(result.message || 'محصول به سبد خرید اضافه شد', 'success')
      await navigateTo('/Cart')
    } else {
      showToast(result?.message || 'خطا در افزودن به سبد خرید', 'error')
    }
  } catch { showToast('خطا در افزودن به سبد خرید', 'error') }
}
</script>

<style scoped>
.container { max-width: 1200px; }
@media (max-width: 640px) { .container { padding-left: 1rem; padding-right: 1rem; } }
</style>
