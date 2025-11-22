<template>
  <div class="bg-white rounded-lg border border-gray-200 p-6">
    <div class="flex items-center justify-between mb-4">
      <h3 class="text-lg font-semibold text-gray-900">پیش‌نمایش کاروسل محصولات</h3>
      <div class="flex items-center space-x-2">
        <span class="text-sm text-gray-700">زنده</span>
        <div class="w-2 h-2 bg-green-500 rounded-full animate-pulse"></div>
      </div>
    </div>

    <!-- پیش‌نمایش کاروسل -->
    <div class="bg-gray-50 rounded-lg p-6 min-h-[300px]">
      <!-- هدر کاروسل -->
      <div class="text-center mb-6">
        <h2 class="text-2xl font-bold text-gray-900 mb-2">{{ config.title || 'محصولات پیشنهادی' }}</h2>
        <p v-if="config.subtitle" class="text-gray-800">{{ config.subtitle }}</p>
      </div>

      <!-- اطلاعات دسته‌بندی انتخاب شده -->
      <div v-if="config.categoryId" class="mb-4 p-3 bg-blue-50 rounded-lg border border-blue-200">
        <div class="flex items-center">
          <i class="fas fa-folder text-blue-600 ml-2"></i>
          <span class="text-sm font-medium text-gray-900">دسته‌بندی انتخاب شده:</span>
        </div>
        <p class="text-sm text-gray-800 mt-1">ID: {{ config.categoryId }}</p>
      </div>

      <!-- کاروسل -->
      <div class="relative">
        <!-- دکمه‌های ناوبری -->
        <div v-if="config.showNavigation" class="absolute left-0 top-1/2 transform -translate-y-1/2 z-10">
          <button class="bg-white shadow-lg rounded-full p-2 hover:bg-gray-50 transition-colors">
            <i class="fas fa-chevron-left text-gray-800"></i>
          </button>
        </div>
        <div v-if="config.showNavigation" class="absolute right-0 top-1/2 transform -translate-y-1/2 z-10">
          <button class="bg-white shadow-lg rounded-full p-2 hover:bg-gray-50 transition-colors">
            <i class="fas fa-chevron-right text-gray-800"></i>
          </button>
        </div>

        <!-- اسلایدها -->
        <div class="overflow-hidden">
          <div class="flex space-x-4" :style="{ transform: 'translateX(0px)' }">
            <!-- محصولات واقعی -->
            <div 
              v-for="product in displayProducts" 
              :key="product.id"
              class="flex-shrink-0 bg-white rounded-lg shadow-sm border border-gray-200 p-6"
              :style="{ width: `${100 / (config.slidesPerView || 4)}%` }"
            >
              <!-- تصویر محصول -->
              <div class="aspect-square bg-gray-200 rounded-lg mb-3 flex items-center justify-center overflow-hidden">
                <img 
                  v-if="product.image" 
                  :src="product.image" 
                  :alt="product.name"
                  class="w-full h-full object-cover"
                />
                <i v-else class="fas fa-image text-gray-400 text-2xl"></i>
              </div>
              
              <!-- اطلاعات محصول -->
              <div class="space-y-2">
                <h4 class="font-medium text-gray-900 text-sm line-clamp-2">
                  {{ product.name }}
                </h4>
                
                <!-- قیمت -->
                <div v-if="config.showPrice && product.price" class="flex items-center space-x-2">
                  <span class="text-lg font-bold text-green-600">{{ formatPrice(product.price) }}</span>
                  <span v-if="(product as { original_price?: number }).original_price && ((product as { original_price?: number }).original_price as number) > product.price" class="text-sm text-gray-500 line-through">
                    {{ formatPrice((product as { original_price?: number }).original_price as number) }}
                  </span>
                </div>
                
                <!-- امتیاز -->
                <div v-if="config.showRating && product.rating" class="flex items-center space-x-1">
                  <div class="flex text-yellow-400">
                    <i v-for="i in 5" :key="i" class="fas fa-star text-xs" :class="{ 'text-gray-300': i > (product.rating || 0) }"></i>
                  </div>
                  <span class="text-xs text-gray-700">({{ (product as { rating_count?: number }).rating_count || 0 }})</span>
                </div>
                
                <!-- تخفیف -->
                <div v-if="config.showDiscount && (product as { discount_percentage?: number }).discount_percentage" class="bg-red-100 text-red-800 text-xs px-2 py-1 rounded-full inline-block">
                  {{ (product as { discount_percentage?: number }).discount_percentage }}% تخفیف
                </div>
              </div>
            </div>

            <!-- پیام عدم وجود محصول -->
            <div v-if="displayProducts.length === 0" class="w-full text-center py-8">
              <i class="fas fa-box-open text-gray-400 text-4xl mb-4"></i>
              <p class="text-gray-600 mb-2">محصولی یافت نشد</p>
              <p class="text-sm text-gray-500">
                {{ config.categoryId ? 'در دسته‌بندی انتخاب شده محصولی وجود ندارد' : 'لطفاً دسته‌بندی را انتخاب کنید' }}
              </p>
            </div>
          </div>
        </div>

        <!-- نشانگرها -->
        <div v-if="config.showIndicators && displayProducts.length > 0" class="flex justify-center mt-4 space-x-2">
          <div 
            v-for="i in Math.ceil(displayProducts.length / (config.slidesPerView || 4))" 
            :key="i"
            class="w-2 h-2 rounded-full"
            :class="i === 1 ? 'bg-blue-500' : 'bg-gray-300'"
          ></div>
        </div>
      </div>

      <!-- اطلاعات اضافی -->
      <div class="mt-6 text-center text-sm text-gray-700">
        <p>نمایش {{ displayProducts.length }} محصول در {{ config.slidesPerView || 4 }} اسلاید</p>
        <p v-if="config.autoplay">پخش خودکار: {{ config.autoplayDelay || 3000 }}ms</p>
        <p v-if="config.categoryId" class="text-blue-600 font-medium">
          فیلتر شده بر اساس دسته‌بندی ID: {{ config.categoryId }}
        </p>
        <p v-else class="text-gray-800">
          نمایش همه محصولات (بدون فیلتر دسته‌بندی)
        </p>
        <p v-if="isLoading" class="text-blue-600">
          <i class="fas fa-spinner fa-spin ml-1"></i>
          در حال بارگذاری محصولات...
        </p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch, onMounted } from 'vue'
import type { ProductCarouselConfig } from '~/types/widget'

interface Props {
  config: ProductCarouselConfig
}

const props = defineProps<Props>()

import type { Product } from '~/types/product'

// State
const products = ref<Product[]>([])
const isLoading = ref(false)

// محاسبه محصولات نمایشی
const displayProducts = computed(() => {
  const limit = Math.min(props.config.productCount || 8, products.value.length)
  return products.value.slice(0, limit)
})

// دریافت محصولات از API
const fetchProducts = async () => {
  if (!props.config.categoryId) {
    // اگر دسته‌بندی انتخاب نشده، محصولات عمومی دریافت کن
    await fetchGeneralProducts()
    return
  }

  isLoading.value = true
  try {
    const query = new URLSearchParams({
      limit: String(props.config.productCount || 8),
      category_id: String(props.config.categoryId)
    })

    const response = await $fetch(`/api/products/public?${query}`)
    
    interface ProductResponse {
      data?: {
        products?: Product[]
      } | Product[]
      products?: Product[]
    }
    
    if (response && typeof response === 'object' && 'data' in response) {
      const data = (response as ProductResponse).data
      if (data && typeof data === 'object' && !Array.isArray(data)) {
        const dataObj = data as { products?: Product[] }
        if ('products' in dataObj && Array.isArray(dataObj.products)) {
          products.value = dataObj.products
        } else {
          products.value = []
        }
      } else if (Array.isArray(data)) {
        products.value = data
      } else {
        products.value = []
      }
    } else if (response && typeof response === 'object' && 'products' in response && Array.isArray((response as ProductResponse).products)) {
      products.value = (response as ProductResponse).products as Product[]
    } else if (Array.isArray(response)) {
      products.value = response
    } else {
      products.value = []
    }
  } catch (error) {
    console.error('خطا در دریافت محصولات:', error)
    products.value = []
  } finally {
    isLoading.value = false
  }
}

// دریافت محصولات عمومی
const fetchGeneralProducts = async () => {
  isLoading.value = true
  try {
    const query = new URLSearchParams({
      limit: String(props.config.productCount || 8)
    })

    const response = await $fetch(`/api/products/public?${query}`)
    
    interface ProductResponse {
      data?: {
        products?: Product[]
      } | Product[]
      products?: Product[]
    }
    
    if (response && typeof response === 'object' && 'data' in response) {
      const data = (response as ProductResponse).data
      if (data && typeof data === 'object' && !Array.isArray(data)) {
        const dataObj = data as { products?: Product[] }
        if ('products' in dataObj && Array.isArray(dataObj.products)) {
          products.value = dataObj.products
        } else {
          products.value = []
        }
      } else if (Array.isArray(data)) {
        products.value = data
      } else {
        products.value = []
      }
    } else if (response && typeof response === 'object' && 'products' in response && Array.isArray((response as ProductResponse).products)) {
      products.value = (response as ProductResponse).products as Product[]
    } else if (Array.isArray(response)) {
      products.value = response
    } else {
      products.value = []
    }
  } catch (error) {
    console.error('خطا در دریافت محصولات عمومی:', error)
    products.value = []
  } finally {
    isLoading.value = false
  }
}

// فرمت قیمت
const formatPrice = (price: number): string => {
  if (!price) return '0 تومان'
  return new Intl.NumberFormat('fa-IR').format(price) + ' تومان'
}

// نظارت بر تغییرات config
watch(() => props.config.categoryId, () => {
  fetchProducts()
}, { immediate: false })

watch(() => props.config.productCount, () => {
  fetchProducts()
}, { immediate: false })

// بارگذاری اولیه
onMounted(() => {
  fetchProducts()
})
</script>

<style scoped>
/* استایل‌های خاص برای پیش‌نمایش */
.aspect-square {
  aspect-ratio: 1 / 1;
}

/* انیمیشن برای دکمه‌های ناوبری */
button:hover {
  transform: scale(1.05);
}

/* انیمیشن برای نشانگرها */
.bg-blue-500 {
  animation: pulse 2s infinite;
}

@keyframes pulse {
  0%, 100% {
    opacity: 1;
  }
  50% {
    opacity: 0.5;
  }
}

/* محدود کردن متن به دو خط */
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

/* انیمیشن بارگذاری */
.fa-spinner {
  animation: spin 1s linear infinite;
}

@keyframes spin {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}

/* بهبود تصاویر محصولات */
img {
  transition: transform 0.3s ease;
}

img:hover {
  transform: scale(1.05);
}
</style>
