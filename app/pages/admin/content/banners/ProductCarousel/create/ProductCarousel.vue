<template>
  <div class="product-carousel relative w-full">
    <!-- کانتینر اصلی کاروسل -->
    <div 
      class="carousel-container overflow-hidden relative"
      :style="{ height: `${carouselHeight}px` }"
    >
      <!-- کانتینر اسلایدها -->
      <div 
        class="carousel-slides flex transition-transform duration-500 ease-in-out"
        :style="{ 
          transform: `translateX(-${currentIndex * slideWidth}px)`,
          width: `${totalSlidesWidth}px`
        }"
        @touchstart="handleTouchStart"
        @touchmove="handleTouchMove"
        @touchend="handleTouchEnd"
      >
        <!-- اسلایدهای محصولات -->
        <ProductCarouselItem
          v-for="(product, index) in visibleProducts"
          :key="product.id || index"
          :product="product"
          :slide-width="slideWidth"
          :is-active="index === currentIndex"
          @click="handleProductClick(product)"
        />
      </div>

      <!-- Overlay برای نمایش اطلاعات محصول -->
      <div 
        v-if="showProductOverlay"
        class="product-overlay absolute inset-0 bg-black bg-opacity-50 flex items-center justify-center z-10"
        @click="closeProductOverlay"
      >
        <div class="bg-white p-6 rounded-lg max-w-md mx-4">
          <h3 class="text-xl font-bold mb-2">{{ selectedProduct?.name }}</h3>
          <p class="text-gray-600 mb-4">{{ selectedProduct?.description }}</p>
          <div class="text-2xl font-bold text-green-600 mb-4">
            {{ selectedProduct?.price }} تومان
          </div>
          <button 
            class="bg-blue-600 text-white px-6 py-2 rounded-lg hover:bg-blue-700 transition-colors"
            @click="addToCart(selectedProduct)"
          >
            افزودن به سبد خرید
          </button>
        </div>
      </div>
    </div>

    <!-- ناوبری کاروسل -->
    <CarouselNavigation
      v-if="showNavigation"
      :can-go-prev="canGoPrev"
      :can-go-next="canGoNext"
      :navigation-style="navigationStyle"
      :navigation-size="navigationSize"
      @prev="prevSlide"
      @next="nextSlide"
    />

    <!-- نشانگرهای موقعیت -->
    <CarouselIndicators
      v-if="showIndicators"
      :total-slides="totalSlides"
      :current-index="currentIndex"
      :indicator-style="indicatorStyle"
      :indicator-color="indicatorColor"
      @change="goToSlide"
    />

    <!-- کنترل‌های اضافی -->
    <CarouselControls
      v-if="showControls"
      :autoplay-enabled="autoplay"
      :autoplay-delay="autoplayDelay"
      :is-playing="isPlaying"
      @toggle-autoplay="toggleAutoplay"
      @change-delay="changeAutoplayDelay"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted, watch, nextTick } from 'vue'
import type { Product } from '~/types/product'
import ProductCarouselItem from './ProductCarouselItem.vue'
import CarouselNavigation from './CarouselNavigation.vue'
import CarouselIndicators from './CarouselIndicators.vue'
import CarouselControls from './CarouselControls.vue'

// Props
interface Props {
  products: Product[]
  slidesPerView?: number
  autoplay?: boolean
  autoplayDelay?: number
  loop?: boolean
  showNavigation?: boolean
  showIndicators?: boolean
  showControls?: boolean
  navigationStyle?: 'default' | 'minimal' | 'bold'
  navigationSize?: number
  indicatorStyle?: 'default' | 'small' | 'large'
  indicatorColor?: string
  carouselHeight?: number
  slideGap?: number
  pauseOnHover?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  slidesPerView: 3,
  autoplay: true,
  autoplayDelay: 5000,
  loop: true,
  showNavigation: true,
  showIndicators: true,
  showControls: false,
  navigationStyle: 'default',
  navigationSize: 40,
  indicatorStyle: 'default',
  indicatorColor: '#3b82f6',
  carouselHeight: 400,
  slideGap: 20,
  pauseOnHover: true
})

// Emits
const emit = defineEmits<{
  'product-click': [product: Product]
  'slide-change': [index: number]
  'autoplay-toggle': [enabled: boolean]
}>()

// State
const currentIndex = ref(0)
const isPlaying = ref(props.autoplay)
const autoplayTimer = ref<NodeJS.Timeout | null>(null)
const touchStartX = ref(0)
const touchStartY = ref(0)
const isDragging = ref(false)
const _dragStartX = ref(0)
const dragOffset = ref(0)
const selectedProduct = ref<Product | null>(null)
const showProductOverlay = ref(false)

// Computed
const totalSlides = computed(() => props.products.length)
const slideWidth = computed(() => {
  const containerWidth = typeof window !== 'undefined' ? window.innerWidth : 1200
  return (containerWidth - (props.slidesPerView - 1) * props.slideGap) / props.slidesPerView
})
const totalSlidesWidth = computed(() => totalSlides.value * slideWidth.value)
const canGoPrev = computed(() => props.loop || currentIndex.value > 0)
const canGoNext = computed(() => props.loop || currentIndex.value < totalSlides.value - props.slidesPerView)
const visibleProducts = computed(() => props.products)

// Methods
const nextSlide = () => {
  if (currentIndex.value < totalSlides.value - props.slidesPerView) {
    currentIndex.value++
  } else if (props.loop) {
    currentIndex.value = 0
  }
  emit('slide-change', currentIndex.value)
}

const prevSlide = () => {
  if (currentIndex.value > 0) {
    currentIndex.value--
  } else if (props.loop) {
    currentIndex.value = totalSlides.value - props.slidesPerView
  }
  emit('slide-change', currentIndex.value)
}

const goToSlide = (index: number) => {
  if (index >= 0 && index <= totalSlides.value - props.slidesPerView) {
    currentIndex.value = index
    emit('slide-change', currentIndex.value)
  }
}

const startAutoplay = () => {
  if (!props.autoplay || !isPlaying.value) return
  
  autoplayTimer.value = setInterval(() => {
    nextSlide()
  }, props.autoplayDelay)
}

const stopAutoplay = () => {
  if (autoplayTimer.value) {
    clearInterval(autoplayTimer.value)
    autoplayTimer.value = null
  }
}

const toggleAutoplay = () => {
  isPlaying.value = !isPlaying.value
  if (isPlaying.value) {
    startAutoplay()
  } else {
    stopAutoplay()
  }
  emit('autoplay-toggle', isPlaying.value)
}

const changeAutoplayDelay = (_delay: number) => {
  if (isPlaying.value) {
    stopAutoplay()
    startAutoplay()
  }
}

// Touch events for mobile
const handleTouchStart = (event: TouchEvent) => {
  touchStartX.value = event.touches[0].clientX
  touchStartY.value = event.touches[0].clientY
  isDragging.value = true
  stopAutoplay()
}

const handleTouchMove = (event: TouchEvent) => {
  if (!isDragging.value) return
  
  const touchX = event.touches[0].clientX
  const touchY = event.touches[0].clientY
  
  const deltaX = touchX - touchStartX.value
  const deltaY = touchY - touchStartY.value
  
  // Check if it's a horizontal swipe
  if (Math.abs(deltaX) > Math.abs(deltaY)) {
    event.preventDefault()
    dragOffset.value = deltaX
  }
}

const handleTouchEnd = () => {
  if (!isDragging.value) return
  
  isDragging.value = false
  
  if (Math.abs(dragOffset.value) > 50) {
    if (dragOffset.value > 0) {
      prevSlide()
    } else {
      nextSlide()
    }
  }
  
  dragOffset.value = 0
  
  if (isPlaying.value) {
    startAutoplay()
  }
}

// Mouse events for desktop
const handleMouseEnter = () => {
  if (props.pauseOnHover) {
    stopAutoplay()
  }
}

const handleMouseLeave = () => {
  if (props.pauseOnHover && isPlaying.value) {
    startAutoplay()
  }
}

// Product interaction
const handleProductClick = (product: Product) => {
  selectedProduct.value = product
  showProductOverlay.value = true
  emit('product-click', product)
}

const closeProductOverlay = () => {
  showProductOverlay.value = false
  selectedProduct.value = null
}

const addToCart = (_product: Product) => {
  // اینجا منطق افزودن به سبد خرید را پیاده‌سازی کنید

  closeProductOverlay()
}

// Lifecycle
onMounted(() => {
  if (isPlaying.value) {
    startAutoplay()
  }
  
  // Add mouse events
  const container = document.querySelector('.carousel-container')
  if (container) {
    container.addEventListener('mouseenter', handleMouseEnter)
    container.addEventListener('mouseleave', handleMouseLeave)
  }
})

onUnmounted(() => {
  stopAutoplay()
  
  const container = document.querySelector('.carousel-container')
  if (container) {
    container.removeEventListener('mouseenter', handleMouseEnter)
    container.removeEventListener('mouseleave', handleMouseLeave)
  }
})

// Watchers
watch(() => props.autoplay, (newValue) => {
  if (newValue && isPlaying.value) {
    startAutoplay()
  } else {
    stopAutoplay()
  }
})

watch(() => props.autoplayDelay, (_newValue) => {
  if (isPlaying.value) {
    stopAutoplay()
    startAutoplay()
  }
})

// Responsive handling
const handleResize = () => {
  // Recalculate slide width on resize
  nextTick(() => {
    // Force re-render
  })
}

onMounted(() => {
  window.addEventListener('resize', handleResize)
})

onUnmounted(() => {
  window.removeEventListener('resize', handleResize)
})
</script>

<style scoped>
.product-carousel {
  position: relative;
  width: 100%;
}

.carousel-container {
  position: relative;
  width: 100%;
  overflow: hidden;
  border-radius: 12px;
  background: #f9fafb;
}

.carousel-slides {
  display: flex;
  transition: transform 0.5s cubic-bezier(0.4, 0, 0.2, 1);
  will-change: transform;
}

.product-overlay {
  background: rgba(0, 0, 0, 0.8);
  backdrop-filter: blur(8px);
  z-index: 50;
}

.product-overlay > div {
  background: white;
  border-radius: 12px;
  box-shadow: 0 25px 50px rgba(0, 0, 0, 0.25);
  max-width: 400px;
  width: 90%;
  margin: 0 auto;
}

.product-overlay h3 {
  color: #1f2937;
  font-weight: 700;
}

.product-overlay p {
  color: #6b7280;
  line-height: 1.6;
}

.product-overlay .text-green-600 {
  color: #059669;
  font-weight: 700;
}

.product-overlay button {
  background: #3b82f6;
  border: none;
  border-radius: 8px;
  font-weight: 600;
  padding: 12px 24px;
  transition: all 0.2s ease;
  cursor: pointer;
}

.product-overlay button:hover {
  background: #2563eb;
  transform: translateY(-1px);
}

/* Touch and drag feedback */
.carousel-slides.dragging {
  transition: none;
  cursor: grabbing;
}

/* Smooth transitions */
.carousel-container,
.carousel-slides,
.product-carousel-item {
  transition: transform 0.3s ease, opacity 0.3s ease;
}

/* Loading state */
.carousel-container.loading {
  opacity: 0.6;
  pointer-events: none;
}

/* Responsive styles */
@media (max-width: 768px) {
  .carousel-container {
    border-radius: 8px;
  }

  .product-overlay {
    padding: 16px;
  }

  .product-overlay > div {
    padding: 20px;
    margin: 16px;
  }
}

@media (max-width: 480px) {
  .carousel-container {
    border-radius: 6px;
  }

  .product-overlay > div {
    padding: 16px;
    margin: 12px;
  }

  .product-overlay h3 {
    font-size: 18px;
  }

  .product-overlay p {
    font-size: 14px;
  }
}

/* Accessibility */
.product-carousel:focus-within {
  outline: 2px solid #3b82f6;
  outline-offset: 2px;
}

/* Animation classes */
.carousel-enter-active,
.carousel-leave-active {
  transition: all 0.5s ease;
}

.carousel-enter-from,
.carousel-leave-to {
  opacity: 0;
  transform: translateX(30px);
}

/* Dark mode support */
@media (prefers-color-scheme: dark) {
  .carousel-container {
    background: #1f2937;
  }

  .product-overlay > div {
    background: #374151;
    color: #f9fafb;
  }

  .product-overlay h3 {
    color: #f9fafb;
  }

  .product-overlay p {
    color: #d1d5db;
  }
}
</style>
