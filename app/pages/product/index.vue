<template>
  <div class="min-h-screen bg-white" dir="rtl">
    <!-- Breadcrumb -->
    <div class="bg-white shadow-sm">
      <div class="container mx-auto px-4 py-3">
        <nav class="flex items-center gap-2 text-sm text-gray-600">
          <NuxtLink to="/" class="hover:text-blue-600 transition-colors">خانه</NuxtLink>
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"></path>
          </svg>
          <NuxtLink to="/products" class="hover:text-blue-600 transition-colors">محصولات</NuxtLink>
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"></path>
          </svg>
          <span class="text-gray-900 font-medium truncate">{{ product?.name || 'بارگذاری...' }}</span>
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
        <NuxtLink to="/products" class="inline-flex items-center gap-2 bg-blue-600 text-white px-6 py-3 rounded-lg hover:bg-blue-700 transition-colors">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path>
          </svg>
          بازگشت به محصولات
        </NuxtLink>
      </div>
    </div>

    <!-- Product Content -->
    <div v-else-if="product" class="container mx-auto px-4 py-6">
      <!-- Main Product Section -->
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
          <component :is="currentTabComponent" :product="product" />
        </div>
      </div>
    </div>

    <!-- Floating Action Buttons -->
    <div class="fixed bottom-6 left-6 flex flex-col gap-3 z-40">
      <!-- Add to Cart -->
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

      <!-- Add to Wishlist -->
      <button
        :class="[
          'px-4 py-4 rounded-full shadow-lg hover:shadow-xl transition-all duration-200 group',
          isInWishlist ? 'bg-red-600 hover:bg-red-700' : 'bg-gray-600 hover:bg-gray-700'
        ]"
        title="افزودن به علاقه‌مندی‌ها"
        @click="toggleWishlist"
      >
        <svg class="w-6 h-6 group-hover:scale-110 transition-transform text-white" fill="currentColor" viewBox="0 0 24 24">
          <path d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z"/>
        </svg>
      </button>

      <!-- Share -->
      <button
        class="bg-blue-600 hover:bg-blue-700 text-white px-4 py-4 rounded-full shadow-lg hover:shadow-xl transition-all duration-200 group"
        title="اشتراک‌گذاری"
        @click="shareProduct"
      >
        <svg class="w-6 h-6 group-hover:scale-110 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.367 2.684 3 3 0 00-5.367-2.684z"></path>
        </svg>
      </button>
    </div>

    <!-- Toast Notifications -->
    <div
v-if="toast.show" :class="[
      'fixed topx-4 py-4 right-6 z-50 px-6 py-4 rounded-lg shadow-lg transition-all duration-300',
      toast.type === 'success' ? 'bg-green-100 text-green-800 border border-green-200' :
      toast.type === 'error' ? 'bg-red-100 text-red-800 border border-red-200' :
      'bg-blue-100 text-blue-800 border border-blue-200'
    ]">
      <div class="flex items-center gap-3">
        <svg v-if="toast.type === 'success'" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
        </svg>
        <svg v-else-if="toast.type === 'error'" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
        </svg>
        <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
        </svg>
        <span>{{ toast.message }}</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import ProductComplements from './ProductComplements.vue'
import ProductDescription from './ProductDescription.vue'
import ProductFAQ from './ProductFAQ.vue'
import ProductMainInfo from './ProductMainInfo.vue'
import ProductRelated from './ProductRelated.vue'
import ProductReviews from './ProductReviews.vue'
import ProductSpecifications from './ProductSpecifications.vue'

// Meta
definePageMeta({
  layout: 'default'
})

const route = useRoute()
const productId = computed(() => route.query.id)

// State
const product = ref(null)
const loading = ref(true)
const error = ref('')
const activeTab = ref('description')
const isInWishlist = ref(false)

// SEO
useHead(() => ({
  title: product.value ? `${product.value.name} - فروشگاه` : 'محصول - فروشگاه',
  meta: [
    {
      name: 'description',
      content: product.value?.description || 'مشاهده جزئیات محصول در فروشگاه'
    },
    {
      property: 'og:title',
      content: product.value?.name || 'محصول'
    },
    {
      property: 'og:description',
      content: product.value?.description || 'مشاهده جزئیات محصول در فروشگاه'
    },
    {
      property: 'og:image',
      content: product.value?.main_image || '/default-product.jpg'
    }
  ]
}))

// Toast
const toast = reactive({
  show: false,
  message: '',
  type: 'success'
})

// Tab configuration
const tabs = [
  { 
    label: 'توضیحات', 
    value: 'description'
  },
  { 
    label: 'مشخصات فنی', 
    value: 'specifications'
  },
  { 
    label: 'نظرات', 
    value: 'reviews'
  },
  { 
    label: 'سوالات متداول', 
    value: 'faq'
  },
  { 
    label: 'محصولات مرتبط', 
    value: 'related'
  },
  { 
    label: 'محصولات مکمل', 
    value: 'complements'
  }
]

// Tab components
const tabComponents = {
  description: ProductDescription,
  specifications: ProductSpecifications,
  reviews: ProductReviews,
    faq: ProductFAQ,
    related: ProductRelated,
    complements: ProductComplements
}

const currentTabComponent = computed(() => {
  return tabComponents[activeTab.value] || tabComponents.description
})

// Functions
const fetchProduct = async () => {
  if (!productId.value) {
    error.value = 'شناسه محصول معتبر نیست'
    loading.value = false
    return
  }

  try {
    loading.value = true
    const response = await $fetch(`/api/products/${productId.value}`)
    
    if (!response) {
      throw new Error('محصول یافت نشد')
    }
    
    product.value = response
    
    // Check if product is in wishlist
    await checkWishlistStatus()
    
  } catch (err) {
    console.error('Error fetching product:', err)
    error.value = err.message || 'خطا در بارگذاری محصول'
  } finally {
    loading.value = false
  }
}

const checkWishlistStatus = async () => {
  // فقط اگر کاربر احراز هویت شده باشد، وضعیت wishlist را بررسی کن
  const { isAuthenticated } = useAuthState()
  if (!isAuthenticated.value) {
    isInWishlist.value = false
    return
  }

  try {
    const response = await $fetch(`/api/wishlist/check/${productId.value}`)
    isInWishlist.value = response?.isInWishlist || false
  } catch {
    // Silently fail for non-authenticated users
    isInWishlist.value = false
  }
}

// ردیابی زمان ماندن در صفحه برای تحلیل بازاریابی
const viewStartTime = ref(null)
const currentProductId = ref(null)
let durationUpdateTimeout = null

// ثبت بازدید محصول (فقط برای کاربران لاگین شده)
const trackProductView = async (prodId) => {
  if (!prodId || !import.meta.client) return
  try {
    await $fetch(`/api/recent-views/product/${prodId}`, {
      method: 'POST',
      credentials: 'include',
      // Silent fail: اگر کاربر لاگین نباشد یا خطایی رخ داد، چیزی نمایش ندهیم
      onResponseError: () => {}
    })
    // شروع ردیابی زمان برای این محصول
    viewStartTime.value = Date.now()
    currentProductId.value = prodId
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
    updateViewDuration()
  }
}

const handleBeforeUnload = () => {
  updateViewDuration()
}

const addToCart = async () => {
  try {
    await $fetch('/api/cart/add', {
      method: 'POST',
      body: {
        product_id: product.value.id,
        quantity: 1
      }
    })
    
    showToast('محصول با موفقیت به سبد خرید اضافه شد', 'success')
  } catch (err) {
    console.error('Add to cart error:', err)
    showToast('خطا در افزودن به سبد خرید', 'error')
  }
}

const toggleWishlist = async () => {
  try {
    if (isInWishlist.value) {
      await $fetch(`/api/wishlist/remove/${product.value.id}`, {
        method: 'DELETE'
      })
      isInWishlist.value = false
      showToast('محصول از علاقه‌مندی‌ها حذف شد', 'success')
    } else {
      await $fetch('/api/wishlist/add', {
        method: 'POST',
        body: {
          product_id: product.value.id
        }
      })
      isInWishlist.value = true
      showToast('محصول به علاقه‌مندی‌ها اضافه شد', 'success')
    }
  } catch (err) {
    console.error('Wishlist error:', err)
    showToast('خطا در عملیات علاقه‌مندی‌ها', 'error')
  }
}

const shareProduct = async () => {
  try {
    if (navigator.share) {
      await navigator.share({
        title: product.value.name,
        text: product.value.description,
        url: window.location.href
      })
    } else {
      // Fallback: copy to clipboard
      await navigator.clipboard.writeText(window.location.href)
      showToast('لینک محصول کپی شد', 'success')
    }
  } catch (err) {
    console.error('Share error:', err)
    showToast('خطا در اشتراک‌گذاری', 'error')
  }
}

const showToast = (message, type = 'success') => {
  toast.message = message
  toast.type = type
  toast.show = true
  
  setTimeout(() => {
    toast.show = false
  }, 3000)
}

// Lifecycle
onMounted(async () => {
  await fetchProduct()
  // ثبت بازدید بعد از لود موفق محصول
  if (product.value?.id) {
    await trackProductView(product.value.id)
  }
  
  // اضافه کردن event listener ها برای ردیابی زمان
  if (import.meta.client) {
    document.addEventListener('visibilitychange', handleVisibilityChange)
    window.addEventListener('beforeunload', handleBeforeUnload)
    
    // هر 30 ثانیه یک بار duration را بروزرسانی کن
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

// Watch for route changes
watch(() => productId.value, async () => {
  if (productId.value) {
    // بروزرسانی duration برای محصول قبلی
    if (currentProductId.value && currentProductId.value !== product.value?.id) {
      await updateViewDuration()
    }
    
    await fetchProduct()
    // ثبت بازدید برای محصول جدید
    if (product.value?.id) {
      await trackProductView(product.value.id)
    }
  }
})
</script>

<style scoped>
/* Custom styles for RTL layout and animations */
.container {
  max-width: 1200px;
}

@media (max-width: 640px) {
  .container {
    padding-left: 1rem;
    padding-right: 1rem;
  }
}
</style> 
