<template>
  <!-- Optimized for performance - Removed ClientOnly for faster loading -->
  <!-- @ts-ignore -->
  <div
    class="product-carousel-widget relative"
    :class="containerClass"
    :style="containerStyle"
    style="font-family: 'IranYekan', 'Vazir', 'IRANSansX', 'Tahoma', sans-serif;"
  >
      <div class="w-full">
        <!-- عنوان کاروسل -->
        <div v-if="getConfig().title" class="text-center mb-8">
          <h2 class="text-2xl md:text-3xl font-bold text-gray-900 mb-2">{{ getConfig().title }}</h2>
          <p v-if="getConfig().subtitle" class="text-gray-600">{{ getConfig().subtitle }}</p>
          <div v-if="getConfig().subtitle" class="w-24 h-1 bg-blue-600 mx-auto mt-2"></div>
        </div>

        <!-- Loading State -->
        <div v-if="pending" class="flex justify-center items-center h-64">
          <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
        </div>

        <!-- کاروسل محصولات -->
        <div v-else class="relative">
          <!-- دکمه‌های ناوبری -->
          <button
            v-show="getProducts().length > getConfig().slidesPerView"
            class="absolute left-0 top-1/2 -translate-y-1/2 z-10 bg-white shadow-lg rounded-full p-3 hover:bg-gray-50 transition-colors"
            aria-label="Previous products"
            @click="prevSlide"
          >
            <svg class="w-5 h-5 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/>
            </svg>
          </button>

          <button
            v-show="getProducts().length > getConfig().slidesPerView"
            class="absolute right-0 top-1/2 -translate-y-1/2 z-10 bg-white shadow-lg rounded-full p-3 hover:bg-gray-50 transition-colors"
            aria-label="Next products"
            @click="nextSlide"
          >
            <svg class="w-5 h-5 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
            </svg>
          </button>

          <!-- کانتینر اسلایدها -->
          <div ref="carouselContainer" class="overflow-hidden">
            <div
              class="flex transition-transform duration-300 ease-in-out"
              :style="{ transform: `translateX(-${getCurrentSlide() * (100 / getConfig().slidesPerView)}%)` }"
            >
              <div
                v-for="product in getProducts()"
                :key="product.id"
                :class="`flex-none px-1 ${slideWidth}`"
              >
                <!-- کارت محصول -->
                <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden hover:shadow-md transition-shadow max-w-[200px] mx-auto">
                  <!-- تصویر محصول -->
                  <NuxtLink :to="buildProductLink(product)" class="block">
                    <div class="aspect-square bg-gray-200 relative overflow-hidden">
                      <img
                        v-if="product.image"
                        :src="product.image"
                        :alt="product.name"
                        class="w-full h-full object-cover"
                      />
                      <div v-else class="w-full h-full flex items-center justify-center">
                        <i class="fas fa-image text-gray-400 text-lg"></i>
                      </div>

                      <!-- تخفیف -->
                      <div
v-if="getConfig().showDiscount && product.discount_percentage"
                           class="absolute top-1 right-1 bg-red-500 text-white text-xs px-1 py-0.5 rounded-full">
                        {{ product.discount_percentage }}%
                      </div>
                    </div>
                  </NuxtLink>

                  <!-- اطلاعات محصول -->
                  <div class="p-2">
                    <!-- نام محصول -->
                    <NuxtLink :to="buildProductLink(product)" class="block">
                      <h3 class="font-medium text-gray-900 text-xs line-clamp-2 mb-1 leading-tight hover:text-blue-600 transition-colors">
                        {{ product.name }}
                      </h3>
                    </NuxtLink>

                    <!-- قیمت -->
                    <div v-if="getConfig().showPrice && product.price && isProductInStock(product) && !product.call_for_price" class="mb-1">
                      <div class="flex items-center space-x-1">
                        <span class="text-xs font-bold text-green-600">{{ formatPrice(product.price) }}</span>
                        <span
v-if="product.original_price && product.original_price > product.price"
                              class="text-xs text-gray-500 line-through">
                          {{ formatPrice(product.original_price) }}
                        </span>
                      </div>
                    </div>
                    <!-- پیام تماس برای قیمت -->
                    <div v-else-if="product.call_for_price" class="mb-1">
                      <span class="text-xs font-bold text-blue-600">قیمت تماس بگیرید</span>
                    </div>
                    <!-- پیام ناموجود -->
                    <div v-else-if="!isProductInStock(product)" class="mb-1">
                      <span class="text-xs font-bold text-red-600">ناموجود</span>
                    </div>

                    <!-- امتیاز -->
                    <div v-if="getConfig().showRating && product.rating" class="flex items-center space-x-1 mb-1">
                      <div class="flex text-yellow-400">
                        <i
v-for="i in 5" :key="i" class="fas fa-star text-xs"
                           :class="{ 'text-gray-300': i > product.rating }"></i>
                      </div>
                      <span class="text-xs text-gray-500">({{ product.rating_count || 0 }})</span>
                    </div>

                    <!-- دکمه افزودن به سبد خرید -->
                    <button
                      :disabled="getAddingToCart() === product.id || !isProductInStock(product) || product.call_for_price"
                      class="w-full py-1 px-2 rounded-md text-xs font-medium transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
                      :class="(isProductInStock(product) && !product.call_for_price)
                        ? 'bg-blue-600 text-white hover:bg-blue-700' 
                        : 'bg-gray-400 text-gray-600 cursor-not-allowed'"
                      @click="addToCart(product)"
                    >
                      <span v-if="getAddingToCart() === product.id" class="flex items-center justify-center">
                        <div class="animate-spin rounded-full h-3 w-3 border-b-2 border-white mr-1"></div>
                        در حال افزودن...
                      </span>
                      <span v-else-if="product.call_for_price">تماس برای قیمت</span>
                      <span v-else-if="!isProductInStock(product)">ناموجود</span>
                      <span v-else>افزودن به سبد خرید</span>
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- نشانگرها -->
          <div
            v-if="getProducts().length > getConfig().slidesPerView"
            class="flex justify-center mt-6 space-x-2"
          >
            <button
              v-for="n in Math.ceil(getProducts().length / getConfig().slidesPerView)"
              :key="`indicator-${n}`"
              class="w-2 h-2 rounded-full transition-colors"
              :class="getCurrentSlide() === n - 1 ? 'bg-blue-600' : 'bg-gray-300'"
              :aria-label="`Go to slide ${n}`"
              @click="goToSlide(n - 1)"
            ></button>
          </div>
        </div>
      </div>
    </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue';
import type { Product, ProductCarouselConfig, Widget } from '~/types/widget';

interface Props {
  widget: Widget
}

const props = defineProps<Props>()

// Import composables
// @ts-ignore
const { buildProductLink } = useProductLink()
// @ts-ignore
const { addToCart: addToCartComposable } = useCart()

// استخراج تنظیمات از widget config
const config = computed((): ProductCarouselConfig => {
  const defaultConfig: ProductCarouselConfig = {
    title: 'محصولات پیشنهادی',
    subtitle: 'بهترین محصولات برای شما',
    categoryId: null,
    productCount: 12,
    slidesPerView: 6,
    showPrice: true,
    showRating: true,
    showDiscount: true,
    autoplay: true,
    autoplayDelay: 3000,
    wide_bg: false,
    bg_color: '#ffffff',
    background_color: '#ffffff',
    padding: '24px',
    margin: '0px',
    showNavigation: true,
    showIndicators: true,
    showControls: false,
    navigationStyle: 'default',
    navigationSize: 40,
    indicatorStyle: 'default',
    indicatorColor: '#3b82f6'
  }
  
  return { ...defaultConfig, ...props.widget.config }
})

// Reactive data
const currentSlide = ref(0)
const pending = ref(false)
const products = ref<Product[]>([])
const addingToCart = ref<number | null>(null)

// دریافت محصولات از API
const fetchProducts = async () => {
  pending.value = true
  try {
    const query = new URLSearchParams({
      limit: String(config.value.productCount || 12)
    })

    if (config.value.categoryId) {
      query.append('category_id', String(config.value.categoryId))
    }

    const response = await $fetch(`/api/products/public?${query}`)
    
    interface ProductsResponse {
      data?: {
        products?: Array<{
          id?: number | string
          [key: string]: unknown
        }>
        [key: string]: unknown
      }
      [key: string]: unknown
    }
    if (response && typeof response === 'object' && 'data' in response) {
      const data = (response as ProductsResponse).data
      if (data && Array.isArray(data.products)) {
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
  } catch (error) {
    console.error('خطا در دریافت محصولات:', error)
    products.value = []
  } finally {
    pending.value = false
  }
}

// Computed properties
const containerClass = computed(() => {
  return config.value.wide_bg ? 'w-full' : 'container mx-auto px-4'
})

const containerStyle = computed(() => {
  return {
    backgroundColor: config.value.background_color || config.value.bg_color || '#ffffff',
    padding: config.value.padding,
    margin: config.value.margin
  }
})

const slideWidth = computed(() => {
  return `w-1/${config.value.slidesPerView}`
})

const visibleProducts = computed(() => {
  return products.value.slice(0, config.value.productCount)
})

// Helper functions for template
const getConfig = () => config.value
const getProducts = () => products.value
const getCurrentSlide = () => currentSlide.value
const getAddingToCart = () => addingToCart.value

// Methods
const formatPrice = (price: number) => {
  if (!price) return '0 تومان'
  return new Intl.NumberFormat('fa-IR').format(price) + ' تومان'
}

const isProductInStock = (product: Product) => {
  // بررسی موجودی محصول
  if (product.stock_quantity !== undefined) {
    return product.stock_quantity > 0
  }
  if (product.in_stock !== undefined) {
    return product.in_stock
  }
  // اگر اطلاعات موجودی موجود نیست، فرض بر موجود بودن
  return true
}

const addToCart = async (product: Product) => {
  try {
    // اطمینان از وجود شناسه محصول
    if (!product.id) {
      console.error('شناسه محصول موجود نیست')
      return
    }

    // تنظیم loading state
    addingToCart.value = product.id

    // استفاده از composable سبد خرید
    const result = await addToCartComposable(product.id, 1, product)
    
    if (result.success) {
      // نمایش پیام موفقیت (اختیاری)
    } else {
      console.error('خطا در افزودن به سبد خرید:', result.message)
    }
  } catch (error) {
    console.error('خطا در افزودن به سبد خرید:', error)
  } finally {
    // پاک کردن loading state
    addingToCart.value = null
  }
}

const nextSlide = () => {
  if (currentSlide.value < Math.ceil(visibleProducts.value.length / config.value.slidesPerView) - 1) {
    currentSlide.value++
  } else {
    currentSlide.value = 0
  }
}

const prevSlide = () => {
  if (currentSlide.value > 0) {
    currentSlide.value--
  } else {
    currentSlide.value = Math.ceil(visibleProducts.value.length / config.value.slidesPerView) - 1
  }
}

const goToSlide = (index: number) => {
  currentSlide.value = index
}

// Autoplay functionality
let autoplayInterval: NodeJS.Timeout | null = null

const startAutoplay = () => {
  if (config.value.autoplay && visibleProducts.value.length > config.value.slidesPerView) {
    autoplayInterval = setInterval(() => {
      nextSlide()
    }, config.value.autoplayDelay)
  }
}

const stopAutoplay = () => {
  if (autoplayInterval) {
    clearInterval(autoplayInterval)
    autoplayInterval = null
  }
}

// Lifecycle
// @ts-ignore
onMounted(() => {
  fetchProducts()
})

// @ts-ignore
onUnmounted(() => {
  stopAutoplay()
})

// Watch for config changes
// @ts-ignore
watch(() => config.value.categoryId, () => {
  fetchProducts()
})

// @ts-ignore
watch(() => config.value.productCount, () => {
  fetchProducts()
})

// Watch for products changes to start autoplay
// @ts-ignore
watch(() => products.value.length, () => {
  if (products.value.length > 0) {
    startAutoplay()
  }
})
</script>

<style scoped>
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
  line-clamp: 2;
}

.aspect-square {
  aspect-ratio: 1 / 1;
}

.line-clamp {
  overflow: hidden;
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 2;
  line-clamp: 2;
}
</style>