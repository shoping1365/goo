<template>
  <div class="single-slider-side" :class="containerClass" :style="containerStyle">
    <div :class="innerContainerClass">
      <div v-if="config.slides?.length > 0" class="slider-container">
        <div class="relative overflow-hidden" :style="sliderContainerStyle">
          <div class="simple-slider">
            <div class="slides-container" :style="slidesContainerStyle">
              <div
                v-for="(slide, index) in config.slides"
                :key="`slide-${index}`"
                class="slide"
                :class="{
                  'active': currentSlideIndex === index,
                  'prev': index < currentSlideIndex
                }"
                :style="slideStyle"
              >
                <a
                  :href="slide.link || '#'"
                  class="block relative h-full"
                  :class="{ 'pointer-events-none': !slide.link }"
                >
                  <img
                    :src="slide.image || '/default-product.svg'"
                    :alt="slide.title"
                    class="w-full h-full object-contain object-center"
                    style="image-rendering: auto;"
                  />
                  <div
                    v-if="slide.title && config.show_title"
                    class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-6"
                  >
                    <h3 class="text-white text-lg md:text-xl font-bold">{{ slide.title }}</h3>
                    <p v-if="slide.description && config.show_description" class="text-white/90 text-sm mt-1">{{ slide.description }}</p>
                  </div>
                </a>
              </div>
            </div>

            <!-- Navigation Controls -->
            <div class="slider-controls">
              <button
                v-if="config.slides.length > 1 && config.show_navigation"
                class="navigation-btn navigation-btn-prev"
                aria-label="Next"
                @click="nextSlide"
              >
                ←
              </button>
              <button
                v-if="config.slides.length > 1 && config.show_navigation"
                class="navigation-btn navigation-btn-next"
                aria-label="Prev"
                @click="prevSlide"
              >
                →
              </button>

              <div
                v-if="config.slides.length > 1 && config.show_pagination"
                v-show="config.slides.length > 1 && config.show_pagination"
                class="navigation-dots"
              >
                <button
                  v-for="(slide, index) in config.slides"
                  :key="`dot-${index}`"
                  class="dot"
                  :class="{ active: currentSlideIndex === index }"
                  @click="goToSlide(index)"
                ></button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div v-else class="bg-gray-200 rounded-lg flex items-center justify-center" :style="`height: ${config.height}px`">
        <p class="text-gray-500 text-center">هیچ اسلایدری یافت نشد</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, watch, onMounted, onUnmounted } from 'vue'
import type { Widget, SliderConfig } from '~/types/widget'
import type { SlideItem } from '~/types/widget'

interface Props {
  widget: Widget
}

const props = defineProps<Props>()

// Config
const config = computed(() => {
  const finalConfig = {
    height: 400,
    slides: [],
    autoplay_enabled: true,
    autoplay_delay: 5,
    show_navigation: true,
    show_pagination: true,
    show_title: true,
    show_description: true,
    slider_width: 800,
    center_width: 1000,
    ...((props.widget.config as SliderConfig) || {})
  }

  // Debug config

  return finalConfig
})

// Styles
const containerClass = computed(() => {
  const classes = ['relative']
  
  // اگر عرض اسلایدر وسط باشد، محدودیت عرض اعمال کن
  if (config.value.slider_width === 600) {
    classes.push('max-w-7xl', 'mx-auto')
  }
  
  return classes
})

const innerContainerClass = computed(() => {
  const classes = []
  
  // اگر عرض اسلایدر وسط باشد، پدینگ اضافه کن
  if (config.value.slider_width === 600) {
    classes.push('mx-auto', 'px-4', 'py-6')
  }
  
  return classes
})

const containerStyle = computed(() => {
  const styles: Record<string, string> = {}
  
  // عرض اسلایدر
  if (config.value.slider_width !== undefined) {
    if (config.value.slider_width === 800) {
      styles.width = '100%'
      styles.maxWidth = '100vw'
    } else if (config.value.slider_width === 600) {
      const centerWidth = config.value.center_width || 1000
      styles.width = `${centerWidth}px`
      styles.marginLeft = 'auto'
      styles.marginRight = 'auto'
    }
  } else {
    // اگر عرض تنظیم نشده، از عرض کامل استفاده کن
    styles.width = '100%'
    styles.maxWidth = '100vw'
  }
  
  // پس‌زمینه
  if (config.value.bg_enabled && config.value.bg_color) {
    styles.backgroundColor = config.value.bg_color
  }
  
  // پدینگ و مارجین - فقط اگر عرض اسلایدر وسط نباشد
  if (config.value.slider_width !== 600) {
    if (config.value.padding_top) styles.paddingTop = `${config.value.padding_top}px`
    if (config.value.padding_bottom) styles.paddingBottom = `${config.value.padding_bottom}px`
    if (config.value.margin_left) styles.marginLeft = `${config.value.margin_left}px`
    if (config.value.margin_right) styles.marginRight = `${config.value.margin_right}px`
  }
  
  return styles
})

// Reactive window width
const windowWidth = ref(0)

const slideStyle = computed(() => {
  const styles: Record<string, string> = {}
  
  // ارتفاع ریسپانسیو - بر اساس عرض محاسبه کن
  if (windowWidth.value <= 768) {
    // در موبایل، ارتفاع کمتر برای کاهش فضای خالی
    styles.height = '40vw' // کاهش نسبت از 56.25vw به 40vw
    styles.maxHeight = '250px'
    styles.minHeight = '150px'
  } else {
    // در دسکتاپ هم ارتفاع کمتر
    styles.height = '40vw'
    styles.maxHeight = `${config.value.height}px`
    styles.minHeight = '200px'
  }
  
  return styles
})

const slidesContainerStyle = computed(() => {
  const styles: Record<string, string> = {}
  
  // ارتفاع ریسپانسیو - بر اساس عرض محاسبه کن
  if (windowWidth.value <= 768) {
    // در موبایل، ارتفاع کمتر برای کاهش فضای خالی
    styles.height = '40vw' // کاهش نسبت از 56.25vw به 40vw
    styles.maxHeight = '250px'
    styles.minHeight = '150px'
  } else {
    // در دسکتاپ هم ارتفاع کمتر
    styles.height = '40vw'
    styles.maxHeight = `${config.value.height}px`
    styles.minHeight = '200px'
  }
  
  return styles
})

const sliderContainerStyle = computed(() => {
  const styles: Record<string, string> = {}
  if (config.value.border_width) {
    styles.borderWidth = `${config.value.border_width}px`
    styles.borderStyle = 'solid'
    styles.borderColor = config.value.border_color || '#e5e7eb'
  }
  return styles
})



// Simple slider state
const currentSlideIndex = ref(0)

// Watch for debugging
watch(currentSlideIndex, (newIndex, oldIndex) => {
  // Slide index changed
})

// Simple slider methods
const nextSlide = () => {
  if (config.value.slides && config.value.slides.length > 0) {
    const newIndex = currentSlideIndex.value === 0
      ? config.value.slides.length - 1
      : currentSlideIndex.value - 1
    currentSlideIndex.value = newIndex
  }
}

const prevSlide = () => {
  if (config.value.slides && config.value.slides.length > 0) {
    const newIndex = (currentSlideIndex.value + 1) % config.value.slides.length
    currentSlideIndex.value = newIndex
  }
}

const goToSlide = (index: number) => {
  currentSlideIndex.value = index
}

// Auto play for simple slider
let autoplayInterval: NodeJS.Timeout | null = null

const startAutoplay = () => {
  if (config.value.autoplay_enabled && config.value.slides && config.value.slides.length > 1) {
    autoplayInterval = setInterval(() => {
      nextSlide()
    }, config.value.autoplay_delay * 1000)
  }
}

const stopAutoplay = () => {
  if (autoplayInterval) {
    clearInterval(autoplayInterval)
    autoplayInterval = null
  }
}

watch(() => config.value.autoplay_enabled, (newValue) => {
  newValue ? startAutoplay() : stopAutoplay()
})

onMounted(() => {
  // تنظیم عرض اولیه
  windowWidth.value = window.innerWidth
  
  // اضافه کردن event listener برای تغییر سایز
  const handleResize = () => {
    windowWidth.value = window.innerWidth
  }
  
  window.addEventListener('resize', handleResize)
  
  // تمیز کردن event listener در onUnmounted
  onUnmounted(() => {
    window.removeEventListener('resize', handleResize)
  })
  
  if (config.value.autoplay_enabled && config.value.slides?.length > 1) {
    startAutoplay()
  }
})

onUnmounted(() => stopAutoplay())



const handleImageError = (event: Event) => {
  (event.target as HTMLImageElement).src = '/default-product.svg'
}

// Watch for slides changes to reset current slide index
watch(() => config.value.slides, (newSlides, oldSlides) => {
  if (newSlides && newSlides.length > 0) {
    currentSlideIndex.value = 0
    
    // Restart autoplay if enabled
    if (config.value.autoplay_enabled && newSlides.length > 1) {
      stopAutoplay()
      startAutoplay()
    }
  }
}, { deep: true })


</script>

<style scoped>
.simple-slider {
  position: relative;
  width: 100%;
  height: 100%;
  overflow: hidden;
}

.slides-container {
  position: relative;
  width: 100%;
  height: 100%;
}

.slide {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  opacity: 0;
  transform: translateX(-100%);
  transition: transform 0.5s, opacity 0.5s;
}

.slide.active {
  opacity: 1;
  transform: translateX(0);
  z-index: 1;
}

.slide.prev {
  transform: translateX(100%);
}

.slide img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  object-position: center;
}

.slider-controls {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  pointer-events: none;
}

.slider-controls > * {
  pointer-events: auto;
}

.navigation-dots {
  position: absolute;
  bottom: 1rem;
  left: 50%;
  transform: translateX(-50%);
  display: flex;
  gap: 0.5rem;
  z-index: 20;
}

.dot {
  width: 12px;
  height: 12px;
  border-radius: 50%;
  cursor: pointer;
  background-color: rgba(255, 255, 255, 0.5);
  transition: background-color 0.3s;
  border: none;
  padding: 0;
  margin: 0;
}

.dot.active {
  background-color: white;
  box-shadow: 0 0 4px rgba(255, 255, 255, 0.8);
}

.navigation-btn {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  background-color: rgba(0, 0, 0, 0.5);
  color: white;
  border: none;
  width: 40px;
  height: 40px;
  border-radius: 50%;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 18px;
  font-weight: bold;
  transition: background-color 0.3s;
}

.navigation-btn:hover {
  background-color: rgba(0, 0, 0, 0.7);
}

.navigation-btn-prev {
  left: 1rem;
}

.navigation-btn-next {
  right: 1rem;
}

@media (max-width: 768px) {
  .navigation-btn {
    width: 36px;
    height: 36px;
    font-size: 16px;
  }
}

</style>