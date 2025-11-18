<template>
  <div class="product-carousel-preview relative" :class="containerClass" :style="containerStyle">
    <div class="w-full">
      <!-- عنوان کاروسل -->
      <div v-if="config.title" class="text-center mb-8">
        <h2 class="text-2xl md:text-3xl font-bold text-gray-900 mb-2">{{ config.title }}</h2>
        <p v-if="config.subtitle" class="text-gray-600">{{ config.subtitle }}</p>
        <div v-if="config.subtitle" class="w-24 h-1 bg-blue-600 mx-auto mt-2"></div>
      </div>

      <!-- کاروسل محصولات -->
      <div class="relative">
        <!-- Loading state -->
        <div v-if="isLoading" class="flex items-center justify-center py-12">
          <div class="text-gray-500">در حال بارگذاری محصولات...</div>
        </div>
        
        <!-- Error state -->
        <div v-else-if="error" class="flex items-center justify-center py-12">
          <div class="text-red-500">{{ error }}</div>
        </div>
        
        <!-- Empty state -->
        <div v-else-if="!hasProducts" class="flex items-center justify-center py-12">
          <div class="text-gray-500">محصولی یافت نشد</div>
        </div>
        
        <!-- Products carousel -->
        <template v-else>
        <!-- دکمه‌های ناوبری -->
        <button
          v-show="showNavigation"
          @click="prevSlide"
          class="absolute left-0 top-1/2 -translate-y-1/2 z-10 bg-white shadow-lg rounded-full p-3 hover:bg-gray-50 transition-colors"
          aria-label="Previous products"
        >
          <svg class="w-5 h-5 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/>
          </svg>
        </button>

        <button
          v-show="showNavigation"
          @click="nextSlide"
          class="absolute right-0 top-1/2 -translate-y-1/2 z-10 bg-white shadow-lg rounded-full p-3 hover:bg-gray-50 transition-colors"
          aria-label="Next products"
        >
          <svg class="w-5 h-5 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
          </svg>
        </button>

        <!-- کانتینر اسلایدها -->
        <div class="overflow-hidden" ref="carouselContainer">
          <div
            class="flex transition-transform duration-300 ease-in-out"
            :style="{ transform: `translateX(-${slideTransform}%)` }"
          >
            <div
              v-for="(product, index) in visibleProducts"
              :key="product.id || index"
              :class="`flex-none px-2 ${slideWidth}`"
            >
              <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
                <!-- تصویر محصول -->
                <div class="aspect-square mb-4 bg-gray-100 rounded-lg overflow-hidden">
                  <img
                    :src="getProductImage(product)"
                    :alt="product.name"
                    class="w-full h-full object-cover"
                  />
                </div>
                
                <!-- اطلاعات محصول -->
                <div class="space-y-2">
                  <h3 class="text-sm font-medium text-gray-900 line-clamp-2">{{ product.name }}</h3>
                  
                  <!-- قیمت -->
                  <div v-if="config.showPrice" class="flex items-center space-x-2">
                    <span class="text-lg font-bold text-green-600">{{ formatPrice(getProductPrice(product)) }}</span>
                    <span v-if="config.showDiscount && product.original_price && product.original_price > product.price" class="text-sm text-gray-500 line-through">
                      {{ formatPrice(product.original_price) }}
                    </span>
                  </div>
                  
                  <!-- امتیاز -->
                  <div v-if="config.showRating && product.rating" class="flex items-center space-x-1">
                    <div class="flex text-yellow-400">
                      <i v-for="i in 5" :key="i" class="fas fa-star text-xs"></i>
                    </div>
                    <span class="text-xs text-gray-500">({{ product.rating }})</span>
                  </div>
                  
                  <!-- تخفیف -->
                  <div v-if="config.showDiscount && product.discount_percentage" class="text-xs text-red-600 font-medium">
                    {{ product.discount_percentage }}% تخفیف
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- نشانگرها -->
        <div 
          v-if="showNavigation" 
          class="flex justify-center mt-6 space-x-2"
        >
          <button
            v-for="n in totalSlides"
            :key="`indicator-${n}`"
            @click="goToSlide(n - 1)"
            class="w-2 h-2 rounded-full transition-colors"
            :class="isActiveSlide(n - 1) ? 'bg-blue-600' : 'bg-gray-300'"
            :aria-label="`Go to slide ${n}`"
          ></button>
        </div>
        </template>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch, onMounted, type ComputedRef } from 'vue'
import type { Product, ProductCarouselConfig } from '~/types/widget'

interface Props {
  config: ProductCarouselConfig
}

const props = defineProps<Props>()

// Reactive data
const currentSlide = ref(0)
const products = ref<Product[]>([])
const isLoading = ref(false)
const error = ref<string | null>(null)

// Computed properties
const containerClass = computed(() => {
  return props.config.wide_bg ? 'w-full' : 'container mx-auto px-4'
})

const containerStyle = computed(() => {
  return {
    backgroundColor: props.config.background_color || props.config.bg_color || '#ffffff',
    padding: props.config.padding,
    margin: props.config.margin
  }
})

const slideWidth = computed(() => {
  return `w-1/${props.config.slidesPerView}`
})

// دریافت محصولات از API
const fetchProducts = async () => {
  isLoading.value = true
  error.value = null
  
  try {
    const query = new URLSearchParams({
      limit: String(props.config.productCount || 12)
    })

    if (props.config.categoryId) {
      query.append('category_id', String(props.config.categoryId))
    }

    const response = await $fetch(`/api/products/public?${query}`)
    
    // پردازش پاسخ API
    if (response && typeof response === 'object' && 'data' in response) {
      const data = (response as any).data
      if (Array.isArray(data.products)) {
        products.value = data.products
      } else if (Array.isArray(data)) {
        products.value = data
      } else {
        products.value = []
      }
    } else if (Array.isArray(response)) {
      products.value = response
    } else {
      products.value = []
    }
  } catch (err: any) {
    console.error('خطا در دریافت محصولات:', err)
    error.value = 'خطا در دریافت محصولات'
    products.value = []
  } finally {
    isLoading.value = false
  }
}

// استفاده از computed با نوع صریح برای رفع مشکل استنتاج نوع در v-for
const visibleProducts = computed(() => {
  return products.value.slice(0, props.config.productCount) as Product[]
}) as ComputedRef<Product[]>

const hasProducts = computed(() => {
  return products.value.length > 0
})

const showNavigation = computed(() => {
  return visibleProducts.value.length > props.config.slidesPerView
})

const slideTransform = computed(() => {
  return currentSlide.value * (100 / props.config.slidesPerView)
})

const totalSlides = computed(() => {
  return Math.ceil(visibleProducts.value.length / props.config.slidesPerView)
})

// Methods
const formatPrice = (price: number) => {
  if (!price) return '0 تومان'
  return new Intl.NumberFormat('fa-IR').format(price) + ' تومان'
}

// دریافت قیمت محصول (قیمت فروش یا قیمت اصلی)
const getProductPrice = (product: Product): number => {
  return product.price || 0
}

// دریافت تصویر محصول
const getProductImage = (product: Product): string => {
  if (product.image) {
    return product.image
  }
  // تصویر پیش‌فرض در صورت نبودن تصویر
  return '/images/placeholder-product.jpg'
}

// Watch برای تغییرات config
watch(() => [props.config.categoryId, props.config.productCount], () => {
  fetchProducts()
}, { deep: true })

// بارگذاری محصولات در mount
onMounted(() => {
  fetchProducts()
})

const nextSlide = () => {
  if (currentSlide.value < visibleProducts.value.length - props.config.slidesPerView) {
    currentSlide.value++
  } else {
    currentSlide.value = 0
  }
}

const prevSlide = () => {
  if (currentSlide.value > 0) {
    currentSlide.value--
  } else {
    currentSlide.value = visibleProducts.value.length - props.config.slidesPerView
  }
}

const goToSlide = (index: number) => {
  currentSlide.value = index
}

// Helper function برای مقایسه slide در template
const isActiveSlide = (slideIndex: number): boolean => {
  return currentSlide.value === slideIndex
}
</script>

<style scoped>
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.aspect-square {
  aspect-ratio: 1 / 1;
}
</style>
