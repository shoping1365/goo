<template>
  <div class="main-slider-side-banner relative" :style="containerStyle">
    <div class="h-full">
      <!-- Desktop Layout -->
      <div class="hidden md:flex gap-2 md:gap-3 lg:gap-6 h-full" :class="widgetConfig.banner_position === 'right' ? 'md:flex-row-reverse lg:flex-row-reverse' : ''">
        
        <!-- اسلایدر اصلی دسکتاپ -->
        <div v-if="slides.length > 0" class="slider-container relative" :style="sliderStyle">
          <!-- Simple CSS Slider -->
          <div class="simple-slider">
            <div class="slides-container" :style="`height: ${height}px`">
              <div
                v-for="(slide, index) in slides"
                :key="`slide-${index}`"
                class="slide"
                :class="{ 'active': currentSlideIndex === index }"
                :style="`height: ${height}px`"
              >
                <component
                  :is="slide.link ? 'a' : 'div'"
                  :href="slide.link"
                  :target="slide.link && slide.openInNewTab ? '_blank' : undefined"
                  :rel="slide.link && slide.openInNewTab ? 'noopener noreferrer' : undefined"
                  class="block relative h-full"
                >
                  <img
                    :src="slide.image"
                    :alt="slide.title"
                    class="w-full h-full object-cover object-center"
                    :loading="index === 0 ? 'eager' : 'lazy'"
                    :fetchpriority="index === 0 ? 'high' : 'auto'"
                    style="image-rendering: auto;"
                  />
                  <div
                    v-if="slide.title && slide.showTitle"
                    class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-6"
                  >
                    <h3 class="text-white text-lg md:text-xl font-bold">{{ slide.title }}</h3>
                    <p v-if="slide.description" class="text-white/90 text-sm mt-1">{{ slide.description }}</p>
                  </div>
                </component>
              </div>
            </div>

            <!-- Navigation Buttons Desktop -->
            <button
              v-if="slides.length > 1 && showNavigation"
              class="navigation-btn navigation-btn-prev"
              aria-label="Previous slide"
              @click="prevSlide"
            >
              ←
            </button>
            <button
              v-if="slides.length > 1 && showNavigation"
              class="navigation-btn navigation-btn-next"
              aria-label="Next slide"
              @click="nextSlide"
            >
              →
            </button>

            <!-- Navigation Dots Desktop -->
            <div v-if="slides.length > 1 && showPagination" class="navigation-dots">
              <button
                v-for="(slide, index) in slides"
                :key="`dot-${index}`"
                class="dot"
                :class="currentSlideIndex === index ? 'active' : 'inactive'"
                :aria-label="`Go to slide ${index + 1}`"
                @click="goToSlide(index)"
              ></button>
            </div>
          </div>
        </div>

        <!-- بنرهای دسکتاپ -->
        <div v-if="sideBanners.length > 0" class="banner-container flex flex-col gap-3 lg:gap-6" :style="bannerStyle">
          <component
            :is="banner.link ? 'a' : 'div'"
            v-for="(banner, index) in sideBanners"
            :key="banner.id || index"
            :href="banner.link"
            :target="banner.link && banner.openInNewTab ? '_blank' : undefined"
            :rel="banner.link && banner.openInNewTab ? 'noopener noreferrer' : undefined"
            :class="['banner-item relative flex-1 min-h-0']"
          >
            <img
              :src="banner.image"
              :alt="banner.title"
              class="w-full h-full object-cover object-center rounded-md"
              loading="lazy"
              style="image-rendering: auto;"
            />
          </component>
        </div>
      </div>

      <!-- Mobile Layout -->
      <div class="md:hidden flex flex-col gap-2" :class="currentMobileBannerPosition === 'top' ? 'flex-col-reverse' : ''">
        
        <!-- اسلایدر اصلی موبایل -->
        <div v-if="mobileSlides.length > 0" ref="mobileSliderWrapperRef" class="slider-container relative w-full">
          <!-- Simple CSS Slider -->
          <div class="simple-slider">
            <div class="slides-container" :style="`height: ${currentHeight}px`">
              <div
                v-for="(slide, index) in mobileSlides"
                :key="`mobile-slide-${index}`"
                class="slide"
                :class="{ 'active': currentSlideIndex === index }"
                :style="`height: ${currentHeight}px`"
              >
                <component
                  :is="slide.link ? 'a' : 'div'"
                  :href="slide.link"
                  :target="slide.link && slide.openInNewTab ? '_blank' : undefined"
                  :rel="slide.link && slide.openInNewTab ? 'noopener noreferrer' : undefined"
                  class="block relative h-full"
                >
                  <img
                    :src="slide.mobile_image"
                    :alt="slide.title"
                    :class="mobileImageClass"
                    :loading="index === 0 ? 'eager' : 'lazy'"
                    :fetchpriority="index === 0 ? 'high' : 'auto'"
                    style="image-rendering: auto;"
                  />
                  <div
                    v-if="slide.title && slide.showTitle"
                    class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-6"
                  >
                    <h3 class="text-white text-lg md:text-xl font-bold">{{ slide.title }}</h3>
                    <p v-if="slide.description" class="text-white/90 text-sm mt-1">{{ slide.description }}</p>
                  </div>
                </component>
              </div>
            </div>

            <!-- Navigation Buttons Mobile -->
            <button
              v-if="mobileSlides.length > 1 && showNavigation"
              class="navigation-btn navigation-btn-prev"
              aria-label="Previous slide"
              @click="prevSlide"
            >
              ←
            </button>
            <button
              v-if="mobileSlides.length > 1 && showNavigation"
              class="navigation-btn navigation-btn-next"
              aria-label="Next slide"
              @click="nextSlide"
            >
              →
            </button>

            <!-- Navigation Dots Mobile -->
            <div v-if="mobileSlides.length > 1 && showPagination" class="navigation-dots">
              <button
                v-for="(slide, index) in mobileSlides"
                :key="`mobile-dot-${index}`"
                class="dot"
                :class="currentSlideIndex === index ? 'active' : 'inactive'"
                :aria-label="`Go to slide ${index + 1}`"
                @click="goToSlide(index)"
              ></button>
            </div>
          </div>
        </div>

        <!-- بنرهای موبایل -->
        <div v-if="mobileBanners.length > 0" class="banner-container-mobile flex flex-col gap-2 sm:gap-3 w-full">
          <component
            :is="banner.link ? 'a' : 'div'"
            v-for="(banner, index) in mobileBanners"
            :key="`mobile-${banner.id || index}`"
            :href="banner.link"
            :target="banner.link && banner.openInNewTab ? '_blank' : undefined"
            :rel="banner.link && banner.openInNewTab ? 'noopener noreferrer' : undefined"
            class="banner-item-mobile relative"
            :style="`height: ${mobileBannerHeight}px;`"
          >
            <img
              :src="banner.mobile_image"
              :alt="banner.title"
              class="w-full h-full object-cover object-center rounded-md"
              loading="lazy"
              style="image-rendering: auto;"
            />
          </component>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, nextTick, onMounted, onUnmounted, ref, watch } from 'vue'
import type { SliderConfig, Widget } from '~/types/widget'

interface Props {
  widget: Widget
}

const props = defineProps<Props>()

// --- بهینه‌سازی: استفاده از computedهای کوچک و متمرکز به جای یک آبجکت بزرگ ---

const widgetConfig = computed((): SliderConfig => {
  const config = props.widget.config as SliderConfig || {}
  return {
    wide_bg: false,
    bg_color: '',
    bg_enabled: false,
    banner_position: 'right',
    slider_banner_ratio: '16-8',
    height: 400,
    padding_top: 0,
    padding_bottom: 0,
    margin_left: 0,
    margin_right: 0,
    slider_width: 800,
    center_width: 1000,
    banner_width: 0,
    autoplay_delay: 5,
    autoplay_enabled: true,
    navigation_enabled: true,
    mobile_height: 300,
    mobile_padding_top: 0,
    mobile_padding_bottom: 0,
    mobile_margin_left: 0,
    mobile_margin_right: 0,
    mobile_slider_width: 800,
    mobile_banner_width: 0,
    mobile_banner_position: 'top',
    slides: [],
    side_banners: [],
    ...config
  }
})

// پراپرتی‌های اصلی با مقادیر پیش‌فرض
const allSlides = computed(() => widgetConfig.value.slides || [])
const allBanners = computed(() => widgetConfig.value.side_banners || [])

const slides = computed(() => allSlides.value.filter(slide => Boolean(slide.image)))
const sideBanners = computed(() => allBanners.value.filter(banner => Boolean(banner.image)))

// جایگاه بنر بر اساس دستگاه
const currentMobileBannerPosition = computed(() => widgetConfig.value.mobile_banner_position || 'top')

// تشخیص موبایل
const isMobile = ref(false)

// اسلایدها و بنرهای موبایل - از آرایه‌های جداگانه mobile_slides و mobile_side_banners استفاده می‌کنیم
const mobileSlides = computed(() => widgetConfig.value.mobile_slides ?? [])
const mobileBanners = computed(() => widgetConfig.value.mobile_side_banners ?? [])

const mobileSliderWrapperRef = ref<HTMLElement | null>(null)
const imageRatioCache = new Map<string, number>()
const DEFAULT_AUTO_RATIO = 9 / 16
const DEFAULT_AUTO_HEIGHT = 220

const rawMobileHeight = computed(() => widgetConfig.value.mobile_height)
const isMobileHeightAuto = computed(() => {
  const raw = rawMobileHeight.value
  if (typeof raw === 'string') {
    return raw === 'auto' || raw === ''
  }
  return raw === null || raw === undefined
})

const explicitMobileHeight = computed<number | null>(() => {
  if (!rawMobileHeight.value && rawMobileHeight.value !== 0) {
    return null
  }
  if (isMobileHeightAuto.value) {
    return null
  }

  const raw = rawMobileHeight.value
  if (typeof raw === 'number' && Number.isFinite(raw) && raw > 0) {
    return raw
  }

  const parsed = Number(raw)
  if (Number.isFinite(parsed) && parsed > 0) {
    return parsed
  }

  return null
})

const autoMobileHeight = ref<number | null>(null)

const resolveWidth = () => {
  if (typeof window === 'undefined') {
    return 0
  }
  const el = mobileSliderWrapperRef.value
  if (!el) {
    return 0
  }
  return el.clientWidth || el.offsetWidth || 0
}

const applyAutoHeightFromRatio = (ratio: number) => {
  const width = resolveWidth()
  if (!width) {
    return
  }
  const safeRatio = ratio > 0 ? ratio : DEFAULT_AUTO_RATIO
  autoMobileHeight.value = Math.max(Math.round(width * safeRatio), 1)
}

const computeAutoMobileHeight = async () => {
  if (!isMobileHeightAuto.value) {
    autoMobileHeight.value = null
    return
  }

  if (typeof window === 'undefined') {
    return
  }

  await nextTick()

  const width = resolveWidth()
  if (!width) {
    return
  }

  const firstSlide = mobileSlides.value.find(slide => Boolean(slide.mobile_image))
  if (!firstSlide) {
    applyAutoHeightFromRatio(DEFAULT_AUTO_RATIO)
    return
  }

  const cachedRatio = imageRatioCache.get(firstSlide.mobile_image)
  if (cachedRatio) {
    applyAutoHeightFromRatio(cachedRatio)
    return
  }

  const img = new Image()
  img.src = firstSlide.mobile_image

  if (img.complete && img.naturalWidth > 0) {
    const ratio = img.naturalHeight / img.naturalWidth || DEFAULT_AUTO_RATIO
    imageRatioCache.set(firstSlide.mobile_image, ratio)
    applyAutoHeightFromRatio(ratio)
    return
  }

  img.onload = () => {
    const ratio = img.naturalHeight / img.naturalWidth || DEFAULT_AUTO_RATIO
    imageRatioCache.set(firstSlide.mobile_image, ratio)
    applyAutoHeightFromRatio(ratio)
  }

  img.onerror = () => {
    applyAutoHeightFromRatio(DEFAULT_AUTO_RATIO)
  }
}

const resolvedMobileSliderHeight = computed(() => {
  if (explicitMobileHeight.value !== null) {
    return explicitMobileHeight.value
  }
  if (autoMobileHeight.value !== null) {
    return autoMobileHeight.value
  }

  const width = resolveWidth()
  if (width) {
    return Math.max(Math.round(width * DEFAULT_AUTO_RATIO), DEFAULT_AUTO_HEIGHT)
  }

  return DEFAULT_AUTO_HEIGHT
})

const height = computed(() => widgetConfig.value.height || 400)
const autoplayDelay = computed(() => {
  // استفاده از تنظیمات جداگانه برای موبایل یا دسکتاپ
  if (isMobile.value) {
    const mobileDelay = widgetConfig.value.mobile_autoplay_delay ?? widgetConfig.value.autoplay_delay ?? 5
    return mobileDelay * 1000
  }
  return (widgetConfig.value.autoplay_delay || 5) * 1000
})
const isAutoplayEnabled = computed(() => {
  // استفاده از تنظیمات جداگانه برای موبایل یا دسکتاپ
  if (isMobile.value) {
    return widgetConfig.value.mobile_autoplay_enabled ?? widgetConfig.value.autoplay_enabled ?? true
  }
  return widgetConfig.value.autoplay_enabled ?? true
})
const isLoopEnabled = computed(() => {
  // استفاده از تنظیمات جداگانه برای موبایل یا دسکتاپ
  if (isMobile.value) {
    return widgetConfig.value.mobile_loop_enabled ?? widgetConfig.value.loop_enabled ?? true
  }
  return widgetConfig.value.loop_enabled ?? true
})

const showNavigation = computed(() => {
  // استفاده از تنظیمات جداگانه برای موبایل یا دسکتاپ
  if (isMobile.value) {
    return widgetConfig.value.mobile_show_navigation ?? widgetConfig.value.show_navigation ?? true
  }
  return widgetConfig.value.show_navigation ?? true
})

const showPagination = computed(() => {
  // استفاده از تنظیمات جداگانه برای موبایل یا دسکتاپ
  if (isMobile.value) {
    return widgetConfig.value.mobile_show_pagination ?? widgetConfig.value.show_pagination ?? true
  }
  return widgetConfig.value.show_pagination ?? true
})

// Mobile image class based on height mode
const mobileImageClass = computed(() => 
  isMobileHeightAuto.value 
    ? 'w-full h-full object-contain object-center'
    : 'w-full h-full object-cover object-center'
)

// محاسبه نسبت عرض اسلایدر و بنر
const sliderBannerRatio = computed(() => {
  if (widgetConfig.value.slider_banner_ratio) {
    const [slider, banner] = widgetConfig.value.slider_banner_ratio.split('-').map(Number)
    if (!isNaN(slider) && !isNaN(banner) && slider > 0 && banner > 0) {
      return { slider, banner }
    }
  }
  return { slider: 16, banner: 8 } // نسبت پیش‌فرض
})

const sliderFlexBasis = computed(() => sliderBannerRatio.value.slider)
const bannerFlexBasis = computed(() => sliderBannerRatio.value.banner)

// استایل‌های جداگانه برای اسلایدر و بنر
const sliderStyle = computed(() => {
  const total = sliderFlexBasis.value + bannerFlexBasis.value
  const sliderPercent = (sliderFlexBasis.value / total) * 100
  
  const styles: Record<string, string> = {
    height: `${currentHeight.value}px`,
    flexBasis: `${sliderPercent}%`,
    flexShrink: '0',
    flexGrow: '0',
    minWidth: '0'
  }
  
  return styles
})

const bannerStyle = computed(() => {
  const total = sliderFlexBasis.value + bannerFlexBasis.value
  const bannerPercent = (bannerFlexBasis.value / total) * 100
  
  return {
    flexBasis: `${bannerPercent}%`,
    flexShrink: '0',
    flexGrow: '0',
    minWidth: '0',
    height: `${currentHeight.value}px`
  }
})

// --- پایان بخش بهینه‌سازی ---

// Container style for width and padding/margin
const containerStyle = computed(() => {
  const styles: Record<string, string> = {}

  // انتخاب عرض بر اساس دستگاه
  const currentSliderWidth = isMobile.value 
    ? (widgetConfig.value.mobile_slider_width ?? 800)
    : (widgetConfig.value.slider_width ?? 800)
  
  const currentPaddingTop = isMobile.value
    ? (widgetConfig.value.mobile_padding_top ?? 0)
    : (widgetConfig.value.padding_top ?? 0)
  
  const currentPaddingBottom = isMobile.value
    ? (widgetConfig.value.mobile_padding_bottom ?? 0)
    : (widgetConfig.value.padding_bottom ?? 0)
  
  const currentMarginLeft = isMobile.value
    ? (widgetConfig.value.mobile_margin_left ?? 0)
    : (widgetConfig.value.margin_left ?? 0)
  
  const currentMarginRight = isMobile.value
    ? (widgetConfig.value.mobile_margin_right ?? 0)
    : (widgetConfig.value.margin_right ?? 0)

  // عرض اسلایدر - محدود به حداکثر عرض صفحه
  if (isMobile.value) {
    const targetWidth = Number(currentSliderWidth)
    if (Number.isFinite(targetWidth) && targetWidth > 0) {
      styles.width = '100%'
      styles.maxWidth = `${targetWidth}px`
      styles.boxSizing = 'border-box'
    } else {
      styles.width = '100%'
      styles.maxWidth = '100vw'
    }
  } else if (currentSliderWidth === 800) {
    // عرض کامل (پیش‌فرض دسکتاپ)
    styles.width = '100%'
    styles.maxWidth = '100vw'
  } else if (currentSliderWidth === 600) {
    // عرض محدود به center_width (دسکتاپ)
    const centerWidth = widgetConfig.value.center_width || 1000
    styles.width = `${Math.min(centerWidth, typeof window !== 'undefined' ? window.innerWidth - 32 : centerWidth)}px`
    styles.maxWidth = '100vw'
  }

  // پدینگ - همیشه اعمال شود
  styles.paddingTop = `${currentPaddingTop}px`
  styles.paddingBottom = `${currentPaddingBottom}px`
  
  // مارجین - همیشه اعمال شود
  if (currentMarginLeft !== 0 || currentMarginRight !== 0) {
    // حداقل یکی از مارجین‌ها مقدار داره
    styles.marginLeft = `${currentMarginLeft}px`
    styles.marginRight = `${currentMarginRight}px`
  } else {
    // هیچ مارجین کاستومی نیست
    // در موبایل مرکز کن، در دسکتاپ margin پیش‌فرض (0) بمونه
    if (isMobile.value) {
      styles.marginLeft = 'auto'
      styles.marginRight = 'auto'
    }
  }

  return styles
})

// Responsive heights from config
const mobileHeight = computed(() => {
  if (isMobileHeightAuto.value) {
    return resolvedMobileSliderHeight.value
  }
  const val = widgetConfig.value.mobile_height || height.value
  return typeof val === 'number' ? val : parseInt(String(val)) || height.value
})
const mobileBannerHeight = computed(() => {
  const h = mobileHeight.value
  return typeof h === 'number' ? Math.min(h * 0.35, 150) : 100
})

// ارتفاع فعلی بر اساس دستگاه
const currentHeight = computed(() => isMobile.value ? mobileHeight.value : height.value)

// Calculate individual banner height so total banners + gaps = configured height
const individualBannerHeight = computed(() => {
  const bannerCount = sideBanners.value.length
  if (bannerCount === 0) return 0

  // If only 1 banner, it should have full height
  if (bannerCount === 1) {
    return height.value
  }

  // For multiple banners, divide height among them
  const totalHeight = height.value
  // gap-3 = 12px for md, gap-6 = 16px for lg - use average
  const totalGaps = (bannerCount - 1) * 14 // average of 12px and 16px
  const availableHeight = totalHeight - totalGaps
  return Math.max(0, availableHeight / bannerCount)
})

// Simple slider state
const currentSlideIndex = ref(0)

// Helper to get current slides array based on device
const getCurrentSlides = () => {
  return isMobile.value ? mobileSlides.value : slides.value
}

// Simple slider methods
const nextSlide = () => {
  const currentSlides = getCurrentSlides()
  if (currentSlides.length > 0) {
    if (isLoopEnabled.value) {
      // حالت حلقه فعال: به اولین اسلاید برمی‌گردیم
      currentSlideIndex.value = (currentSlideIndex.value + 1) % currentSlides.length
    } else {
      // حالت حلقه غیرفعال: در آخرین اسلاید می‌ایستیم
      if (currentSlideIndex.value < currentSlides.length - 1) {
        currentSlideIndex.value++
      }
    }
  }
}

const prevSlide = () => {
  const currentSlides = getCurrentSlides()
  if (currentSlides.length > 0) {
    if (isLoopEnabled.value) {
      // حالت حلقه فعال: به آخرین اسلاید برمی‌گردیم
      currentSlideIndex.value = (currentSlideIndex.value - 1 + currentSlides.length) % currentSlides.length
    } else {
      // حالت حلقه غیرفعال: در اولین اسلاید می‌ایستیم
      if (currentSlideIndex.value > 0) {
        currentSlideIndex.value--
      }
    }
  }
}

const goToSlide = (index: number) => {
  currentSlideIndex.value = index
}

// Auto play for simple slider
let autoplayInterval: NodeJS.Timeout | null = null

const startAutoplay = () => {
  const currentSlides = getCurrentSlides()
  if (isAutoplayEnabled.value && currentSlides.length > 1) {
    autoplayInterval = setInterval(() => {
      nextSlide()
    }, autoplayDelay.value)
  }
}

const stopAutoplay = () => {
  if (autoplayInterval) {
    clearInterval(autoplayInterval)
    autoplayInterval = null
  }
}

// Watch for autoplay changes
watch(isAutoplayEnabled, (newValue) => {
  if (newValue) {
    startAutoplay()
  } else {
    stopAutoplay()
  }
})

// تشخیص موبایل
const checkMobile = () => {
  isMobile.value = window.innerWidth < 768
}

// Initialize autoplay
onMounted(() => {
  checkMobile()
  window.addEventListener('resize', checkMobile)
  
  // Compute auto mobile height if needed
  if (isMobile.value && isMobileHeightAuto.value) {
    computeAutoMobileHeight()
  }
  
  const currentSlides = getCurrentSlides()
  if (isAutoplayEnabled.value && currentSlides.length > 1) {
    startAutoplay()
  }
})

onUnmounted(() => {
  window.removeEventListener('resize', checkMobile)
  stopAutoplay()
})

// Watch for device changes to reset slider and restart autoplay
watch(isMobile, () => {
  currentSlideIndex.value = 0
  stopAutoplay()
  
  // Compute auto mobile height when switching to mobile
  if (isMobile.value && isMobileHeightAuto.value) {
    computeAutoMobileHeight()
  }
  
  const currentSlides = getCurrentSlides()
  if (isAutoplayEnabled.value && currentSlides.length > 1) {
    startAutoplay()
  }
})

// Watch for slides changes to reset current slide index
watch([slides, mobileSlides], () => {
  currentSlideIndex.value = 0
  
  // Recompute auto height if mobile slides changed
  if (isMobile.value && isMobileHeightAuto.value) {
    computeAutoMobileHeight()
  }
  
  // Restart autoplay if enabled
  const currentSlides = getCurrentSlides()
  if (isAutoplayEnabled.value && currentSlides.length > 1) {
    stopAutoplay()
    startAutoplay()
  }
}, { deep: true })


</script>

<style scoped>
.main-slider-side-banner {
  box-sizing: border-box;
  overflow: hidden;
}

.slider-container {
  position: relative;
  box-sizing: border-box;
  overflow: hidden;
}

.banner-container {
  box-sizing: border-box;
  overflow: hidden;
}

.simple-slider {
  position: relative;
  width: 100%;
  height: 100%;
  overflow: hidden;
  border-radius: 0.375rem;
  box-sizing: border-box;
}

.slides-container {
  position: relative;
  width: 100%;
  height: 100%;
  overflow: hidden;
  box-sizing: border-box;
}

.slide {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  opacity: 0;
  transition: opacity 0.6s ease-in-out;
  visibility: hidden;
  overflow: hidden;
  box-sizing: border-box;
}

.slide.active {
  opacity: 1;
  visibility: visible;
}

.slide img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  object-position: center;
  border-radius: 0.375rem;
  image-rendering: auto;
  image-rendering: -webkit-optimize-contrast;
  image-rendering: crisp-edges;
}

.banner-item {
  border-radius: 0.375rem;
  overflow: hidden;
  box-sizing: border-box;
}

.banner-item img,
.banner-item-mobile img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  object-position: center;
  image-rendering: auto;
  image-rendering: -webkit-optimize-contrast;
  image-rendering: crisp-edges;
}

.navigation-btn {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  z-index: 20;
  background-color: rgba(255, 255, 255, 0.5);
  border: none;
  border-radius: 50%;
  width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 24px;
  color: #333;
  cursor: pointer;
  transition: background-color 0.3s;
}

.navigation-btn:hover {
  background-color: rgba(255, 255, 255, 0.8);
}

.navigation-btn-prev {
  left: 1rem;
}

.navigation-btn-next {
  right: 1rem;
}

.navigation-dots {
  position: absolute;
  bottom: 1rem;
  left: 50%;
  transform: translateX(-50%);
  z-index: 20;
  display: flex;
  gap: 0.5rem;
}

.dot {
  width: 12px;
  height: 12px;
  border-radius: 50%;
  background-color: rgba(255, 255, 255, 0.5);
  border: 1px solid rgba(0, 0, 0, 0.2);
  cursor: pointer;
  transition: background-color 0.3s;
}

.dot.active {
  background-color: white;
}

/* موبایل */
@media (max-width: 767px) {
  .slider-container {
    width: 100% !important;
    box-sizing: border-box;
  }
  
  .simple-slider,
  .slides-container,
  .slide {
    border-radius: 0.5rem;
  }
  
  /* نقطه‌های کوچکتر برای موبایل */
  .navigation-dots {
    gap: 0.375rem;
  }
  
  .dot {
    width: 8px;
    height: 8px;
  }
  
  .banner-item-mobile {
    min-height: 120px;
  }
  
  .banner-item-mobile img {
    border-radius: 0.5rem;
  }
}
</style> 