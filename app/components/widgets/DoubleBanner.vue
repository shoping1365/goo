<template>
  <!-- بنر دو تایی -->
  <div class="double-banner" :style="containerStyle">
    <div>
      <!-- Desktop Layout -->
      <div class="hidden md:block">
        <div
v-if="desktopBanners.length > 0" class="grid gap-6" :style="{
          '--banner1-ratio': `${config.banner1_ratio || 50}fr`,
          '--banner2-ratio': `${config.banner2_ratio || 50}fr`
        }">
          <div 
            v-for="(banner, index) in desktopBanners.slice(0, 2)" 
            :key="banner.id || index"
            class="relative overflow-hidden rounded-lg"
            :style="{
              height: `${config.height || 400}px`
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
                class="absolute bottom-0 left-0 right-0 p-6 z-10"
              >
                <h3 v-if="config.show_title && banner.title" class="text-white text-lg md:text-xl font-bold mb-1">
                  {{ banner.title }}
                </h3>
                <p v-if="config.show_description && banner.description" class="text-white/90 text-sm md:text-base">
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
                class="absolute bottom-0 left-0 right-0 p-6 z-10"
              >
                <h3 v-if="config.show_title && banner.title" class="text-white text-lg md:text-xl font-bold mb-1">
                  {{ banner.title }}
                </h3>
                <p v-if="config.show_description && banner.description" class="text-white/90 text-sm md:text-base">
                  {{ banner.description }}
                </p>
              </div>
            </div>
          </div>
        </div>
        
        <!-- اگر بنر وجود نداشت، پیام جایگزین -->
        <div v-else class="bg-gray-200 rounded-lg flex items-center justify-center h-64">
          <div class="text-center">
            <div class="text-gray-500 text-lg mb-2">هیچ بنری یافت نشد</div>
            <div class="text-gray-400 text-sm">لطفاً بنرهای مورد نظر را اضافه کنید</div>
          </div>
        </div>
      </div>

      <!-- Mobile Layout -->
      <div v-if="widget.show_on_mobile !== false" class="md:hidden">
        <div v-if="mobileBanners.length > 0" class="grid grid-cols-1 gap-3" :class="config.mobile_vertical_display ? 'grid-cols-1' : 'grid-cols-2'">
          <div 
            v-for="(banner, index) in mobileBanners.slice(0, 2)" 
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
                class="absolute bottom-0 left-0 right-0 p-4 z-10"
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
                class="absolute bottom-0 left-0 right-0 p-4 z-10"
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
        
        <!-- اگر بنر موبایل وجود نداشت، پیام جایگزین -->
        <div v-else class="bg-gray-200 rounded-lg flex items-center justify-center h-40">
          <div class="text-center">
            <div class="text-gray-500 text-sm mb-1">هیچ بنری یافت نشد</div>
            <div class="text-gray-400 text-xs">لطفاً بنرهای مورد نظر را اضافه کنید</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import type { Widget, DoubleBannerConfig } from '~/types/widget'

interface Props {
  widget: Widget
}

const props = defineProps<Props>()

// Get config with proper typing
const config = computed(() => {
  const widgetConfig = props.widget.config as DoubleBannerConfig
  
  const finalConfig = {
    banners: (widgetConfig as any).banners || (widgetConfig as any).slides || [],
    mobile_banners: (widgetConfig as any).mobile_banners || [],
    height: (widgetConfig as any).height || 400,
    mobile_height: (widgetConfig as any).mobile_height || (widgetConfig as any).height || 150,
    mobile_vertical_display: (widgetConfig as any).mobile_vertical_display || false,
    bg_enabled: (widgetConfig as any).bg_enabled || false,
    bg_color: (widgetConfig as any).bg_color || '#ffffff',
    show_title: (widgetConfig as any).show_title !== undefined ? (widgetConfig as any).show_title : true,
    show_description: (widgetConfig as any).show_description !== undefined ? (widgetConfig as any).show_description : true,
    padding_top: (widgetConfig as any).padding_top || 0,
    padding_bottom: (widgetConfig as any).padding_bottom || 0,
    margin_left: (widgetConfig as any).margin_left || 0,
    margin_right: (widgetConfig as any).margin_right || 0,
    banner1_ratio: (widgetConfig as any).banner1_ratio || 50,
    banner2_ratio: (widgetConfig as any).banner2_ratio || 50,
    banner_width: (widgetConfig as any).banner_width || 800,
    center_width: (widgetConfig as any).center_width || 1000
  }
  
  return finalConfig
})

// بنرهای دسکتاپ و موبایل جداگانه
const desktopBanners = computed(() => config.value.banners.filter(banner => Boolean(banner.image)))
const mobileBanners = computed(() => {
  // اگر آرایه mobile_banners وجود داشت از اون استفاده کن، وگرنه از banners
  const mobileBannersArray = (config.value.mobile_banners && config.value.mobile_banners.length > 0)
    ? config.value.mobile_banners 
    : config.value.banners
  return mobileBannersArray.filter(banner => Boolean(banner.image))
})

// Container style for padding and margin
const containerStyle = computed(() => {
  const styles: Record<string, string> = {}
  
  // عرض بنر
  if (config.value.banner_width !== undefined) {
    if (config.value.banner_width === 800) {
      styles.width = '100%'
    } else if (config.value.banner_width === 600) {
      const centerWidth = config.value.center_width || 1000
      styles.width = `${centerWidth}px`
      styles.marginLeft = 'auto'
      styles.marginRight = 'auto'
    }
  }
  
  // پدینگ و مارجین - فقط اگر عرض بنر وسط نباشد
  if (config.value.banner_width !== 600) {
    if (config.value.padding_top !== undefined) styles.paddingTop = `${config.value.padding_top}px`
    if (config.value.padding_bottom !== undefined) styles.paddingBottom = `${config.value.padding_bottom}px`
    if (config.value.margin_left !== undefined) styles.marginLeft = `${config.value.margin_left}px`
    if (config.value.margin_right !== undefined) styles.marginRight = `${config.value.margin_right}px`
  }
  
  return styles
})

</script>

<style scoped>


/* نسبت بنرها در دسکتاپ */
@media (min-width: 768px) {
  .double-banner .grid {
    grid-template-columns: var(--banner1-ratio, 1fr) var(--banner2-ratio, 1fr) !important;
  }
}
</style> 