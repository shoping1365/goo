<template>
  <!-- بنر سه تایی -->
  <div class="triple-banner" :style="containerStyle">
    <div class="py-6" :class="containerPaddingClass">
      <!-- Desktop Layout -->
      <div class="hidden md:block">
        <div v-if="desktopBanners.length > 0" class="grid grid-cols-1 md:grid-cols-3 gap-6" :style="gridStyle">
          <div 
            v-for="(banner, index) in desktopBanners.slice(0, 3)" 
            :key="`desktop-${banner.id || index}`"
            class="relative overflow-hidden rounded-lg"
            :style="{ height: `${config.height || 400}px` }"
          >
            <!-- اگر لینک وجود داشت، کل بنر کلیک‌پذیر باشد -->
            <a v-if="banner.link" :href="banner.link" target="_blank" rel="noopener" class="block w-full h-full">
              <img
                :src="banner.image"
                :alt="banner.title"
                class="w-full h-full object-cover absolute inset-0"
              />
              <!-- عنوان و توضیحات روی تصویر -->
              <div 
                v-if="(config.show_title && banner.title) || (config.show_description && banner.description)" 
                class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/80 via-black/40 to-transparent p-3 z-10"
              >
                <h3 v-if="config.show_title && banner.title" class="text-white text-base md:text-lg font-bold mb-1">
                  {{ banner.title }}
                </h3>
                <p v-if="config.show_description && banner.description" class="text-white/90 text-xs md:text-sm">
                  {{ banner.description }}
                </p>
              </div>
            </a>
            <!-- اگر لینک نبود، فقط تصویر بنر -->
            <div v-else class="w-full h-full">
              <img
                :src="banner.image"
                :alt="banner.title"
                class="w-full h-full object-cover absolute inset-0"
              />
              <!-- عنوان و توضیحات روی تصویر -->
              <div 
                v-if="(config.show_title && banner.title) || (config.show_description && banner.description)" 
                class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/80 via-black/40 to-transparent p-3 z-10"
              >
                <h3 v-if="config.show_title && banner.title" class="text-white text-base md:text-lg font-bold mb-1">
                  {{ banner.title }}
                </h3>
                <p v-if="config.show_description && banner.description" class="text-white/90 text-xs md:text-sm">
                  {{ banner.description }}
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Mobile Layout -->
      <div v-if="widget.show_on_mobile !== false" class="md:hidden">
        <div v-if="mobileBanners.length > 0" class="grid gap-3" :class="config.mobile_vertical_display ? 'grid-cols-1' : 'grid-cols-2'">
          <div 
            v-for="(banner, index) in mobileBanners.slice(0, 3)" 
            :key="`mobile-${banner.id || index}`"
            class="relative overflow-hidden rounded-lg"
            :style="{
              height: `${config.mobile_height || 150}px`
            }"
          >
            <!-- اگر لینک وجود داشت، کل بنر کلیک‌پذیر باشد -->
            <a v-if="banner.link" :href="banner.link" target="_blank" rel="noopener" class="block w-full h-full">
              <img
                :src="banner.image"
                :alt="banner.title"
                class="w-full h-full object-cover absolute inset-0"
              />
              <!-- عنوان و توضیحات روی تصویر -->
              <div 
                v-if="(config.show_title && banner.title) || (config.show_description && banner.description)" 
                class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/80 via-black/40 to-transparent p-4 z-10"
              >
                <h3 v-if="config.show_title && banner.title" class="text-white text-sm font-bold mb-1">
                  {{ banner.title }}
                </h3>
                <p v-if="config.show_description && banner.description" class="text-white/90 text-xs">
                  {{ banner.description }}
                </p>
              </div>
            </a>
            <!-- اگر لینک نبود، فقط تصویر بنر -->
            <div v-else class="w-full h-full">
              <img
                :src="banner.image"
                :alt="banner.title"
                class="w-full h-full object-cover absolute inset-0"
              />
              <!-- عنوان و توضیحات روی تصویر -->
              <div 
                v-if="(config.show_title && banner.title) || (config.show_description && banner.description)" 
                class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/80 via-black/40 to-transparent p-4 z-10"
              >
                <h3 v-if="config.show_title && banner.title" class="text-white text-sm font-bold mb-1">
                  {{ banner.title }}
                </h3>
                <p v-if="config.show_description && banner.description" class="text-white/90 text-xs">
                  {{ banner.description }}
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
      
      <!-- اگر بنر وجود نداشت، پیام جایگزین -->
      <div v-if="desktopBanners.length === 0" class="bg-gray-200 rounded-lg flex items-center justify-center h-64">
        <div class="text-center">
          <div class="text-gray-500 text-lg mb-2">هیچ بنری یافت نشد</div>
          <div class="text-gray-400 text-sm">لطفاً بنرهای مورد نظر را اضافه کنید</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import type { Widget, BannerConfig } from '~/types/widget'

interface Props {
  widget: Widget
}

const props = defineProps<Props>()

// Get config with proper typing
const config = computed(() => {
  const widgetConfig: any = props.widget.config || {}
  
  return {
    banners: widgetConfig.banners || [],
    mobile_banners: widgetConfig.mobile_banners || [],
    height: widgetConfig.height || 400,
    mobile_height: widgetConfig.mobile_height || 150,
    mobile_vertical_display: widgetConfig.mobile_vertical_display || false,
    bg_enabled: widgetConfig.bg_enabled || false,
    bg_color: widgetConfig.bg_color || '#ffffff',
    show_title: widgetConfig.show_title !== undefined ? widgetConfig.show_title : true,
    show_description: widgetConfig.show_description !== undefined ? widgetConfig.show_description : true,
    padding_top: widgetConfig.padding_top || 0,
    padding_bottom: widgetConfig.padding_bottom || 0,
    margin_left: widgetConfig.margin_left || 0,
    margin_right: widgetConfig.margin_right || 0,
    banner1_ratio: widgetConfig.banner1_ratio || 33,
    banner2_ratio: widgetConfig.banner2_ratio || 33,
    banner3_ratio: widgetConfig.banner3_ratio || 33,
    banner_width: widgetConfig.banner_width || 800,
    center_width: widgetConfig.center_width || 1000
  }
})

// Desktop banners computed
const desktopBanners = computed(() => {
  return config.value.banners || []
})

// Mobile banners computed - fallback to desktop if not set
const mobileBanners = computed(() => {
  return config.value.mobile_banners || []
})

// Container style for width and padding/margin
const containerStyle = computed(() => {
  const styles: Record<string, string> = {}
  const configValue = config.value
  
  // عرض بنر
  if (configValue.banner_width !== undefined) {
    if (configValue.banner_width === 800) {
      styles.width = '100%'
    } else if (configValue.banner_width === 600) {
      styles.maxWidth = '1200px'
      styles.marginLeft = 'auto'
      styles.marginRight = 'auto'
    }
  }
  
  // پدینگ و مارجین
  if (configValue.padding_top !== undefined) {
    styles.paddingTop = `${configValue.padding_top}px`
  }
  if (configValue.padding_bottom !== undefined) {
    styles.paddingBottom = `${configValue.padding_bottom}px`
  }
  if (configValue.margin_left !== undefined) {
    styles.marginLeft = `${configValue.margin_left}px`
  }
  if (configValue.margin_right !== undefined) {
    styles.marginRight = `${configValue.margin_right}px`
  }
  
  return styles
})

// Container padding class
const containerPaddingClass = computed(() => {
  const configValue = config.value
  
  if (configValue.banner_width === 600) {
    return 'px-4 sm:px-6 lg:px-8'
  } else {
    return 'px-4'
  }
})

// Grid style for banner ratios
const gridStyle = computed(() => {
  const styles: Record<string, string> = {}
  const configValue = config.value
  
  if (configValue.banner1_ratio && configValue.banner2_ratio && configValue.banner3_ratio) {
    styles.gridTemplateColumns = `${configValue.banner1_ratio}fr ${configValue.banner2_ratio}fr ${configValue.banner3_ratio}fr`
  }
  
  return styles
})
</script>

<style scoped>
.triple-banner {
  max-width: 1200px !important;
  margin-left: auto !important;
  margin-right: auto !important;
}

.triple-banner img {
  transition: transform 0.3s ease;
}

.triple-banner a:hover img {
  transform: scale(1.05);
}
</style> 

