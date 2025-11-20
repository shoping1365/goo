<template>
  <!-- بنر تکی -->
  <div
class="full-banner relative" :class="{
    'full-width': config.banner_width === 800,
    'center-width': config.banner_width === 600
  }" :style="{ 
    height: `${config.height || 400}px`,
    width: config.banner_width === 800 ? '100%' : config.banner_width === 600 ? `${config.center_width || 1000}px` : '100%',
    maxWidth: config.banner_width === 600 ? `${config.center_width || 1000}px` : '100%',
    margin: config.banner_width === 600 ? '0 auto' : '0',
    backgroundColor: config.bg_enabled ? config.bg_color : 'transparent',
    paddingTop: config.padding_top ? `${config.padding_top}px` : '0',
    paddingBottom: config.padding_bottom ? `${config.padding_bottom}px` : '0',
    marginLeft: config.banner_width === 800 && config.margin_left ? `${config.margin_left}px` : '0',
    marginRight: config.banner_width === 800 && config.margin_right ? `${config.margin_right}px` : '0',
    '--desktop-height': `${config.height || 400}px`,
    '--mobile-height': `${config.mobile_height || config.height || 400}px`,
    '--mobile-width': config.mobile_banner_width ? `${config.mobile_banner_width}px` : '100%'
  }">

    <!-- عنوان بنر (در صورت وجود) -->
    <div v-if="config.title" class="w-full text-center font-bold text-xl md:text-2xl mb-2 mt-4 text-gray-800">
      {{ config.title }}
    </div>

    <!-- Desktop Layout -->
    <div class="hidden md:block">
      <div v-if="desktopBanners && desktopBanners.length > 0" class="relative">
        <!-- اگر لینک وجود داشت، کل بنر کلیک‌پذیر باشد -->
        <component
          :is="desktopBanners[0].link ? 'a' : 'div'"
          :href="desktopBanners[0].link"
          :target="desktopBanners[0].link && desktopBanners[0].openInNewTab ? '_blank' : undefined"
          :rel="desktopBanners[0].link && desktopBanners[0].openInNewTab ? 'noopener noreferrer' : undefined"
          class="block w-full"
        >
          <img
            :src="desktopBanners[0].image"
            :alt="desktopBanners[0].title || desktopBanners[0].description || ''"
            class="w-full object-cover lazy-image responsive-banner-image"
            :class="{ 'lazy-loaded': !isLazyLoadingEnabled }"
            loading="eager"
            fetchpriority="high"
          />
        </component>

        <!-- نمایش عنوان و توضیحات بنر -->
        <div
          v-if="(config.show_title && desktopBanners[0].title) || (config.show_description && desktopBanners[0].description)"
          class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-6"
        >
          <h4
            v-if="config.show_title && desktopBanners[0].title"
            class="text-white text-lg font-bold mb-2"
          >
            {{ desktopBanners[0].title }}
          </h4>
          <p
            v-if="config.show_description && desktopBanners[0].description"
            class="text-white/90 text-sm"
          >
            {{ desktopBanners[0].description }}
          </p>
        </div>
      </div>
    </div>

    <!-- Mobile Layout -->
    <div class="md:hidden">
      <div v-if="mobileBanners && mobileBanners.length > 0" class="relative">
        <!-- اگر لینک وجود داشت، کل بنر کلیک‌پذیر باشد -->
        <component
          :is="mobileBanners[0].link ? 'a' : 'div'"
          :href="mobileBanners[0].link"
          :target="mobileBanners[0].link && mobileBanners[0].openInNewTab ? '_blank' : undefined"
          :rel="mobileBanners[0].link && mobileBanners[0].openInNewTab ? 'noopener noreferrer' : undefined"
          class="block w-full"
        >
          <img
            :src="mobileBanners[0].image"
            :alt="mobileBanners[0].title || mobileBanners[0].description || ''"
            class="w-full object-cover lazy-image responsive-banner-image"
            :class="{ 'lazy-loaded': !isLazyLoadingEnabled }"
            loading="eager"
            fetchpriority="high"
          />
        </component>

        <!-- نمایش عنوان و توضیحات بنر -->
        <div
          v-if="(config.show_title && mobileBanners[0].title) || (config.show_description && mobileBanners[0].description)"
          class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-6"
        >
          <h4
            v-if="config.show_title && mobileBanners[0].title"
            class="text-white text-lg font-bold mb-2"
          >
            {{ mobileBanners[0].title }}
          </h4>
          <p
            v-if="config.show_description && mobileBanners[0].description"
            class="text-white/90 text-sm"
          >
            {{ mobileBanners[0].description }}
          </p>
        </div>
      </div>
    </div>

  </div>
</template>

<script setup>
import { computed, ref, onMounted } from 'vue'

const props = defineProps({
  widget: {
    type: Object,
    required: true
  }
})

// Config
const config = computed(() => {
  const widgetConfig = props.widget.config || {}
  
  const finalConfig = {
    banners: widgetConfig.banners || widgetConfig.slides || [],
    mobile_banners: widgetConfig.mobile_banners || [],
    banner_width: widgetConfig.banner_width || 800,
    center_width: widgetConfig.center_width || 1000,
    slider_width: 800,
    height: widgetConfig.height || 400,
    mobile_height: widgetConfig.mobile_height || widgetConfig.height || 400,
    mobile_banner_width: widgetConfig.mobile_banner_width || 0,
    show_title: widgetConfig.show_title !== undefined ? widgetConfig.show_title : true,
    show_description: widgetConfig.show_description !== undefined ? widgetConfig.show_description : true,
    bg_enabled: widgetConfig.bg_enabled || false,
    bg_color: widgetConfig.bg_color || '#ffffff',
    easy_load_enabled: widgetConfig.easy_load_enabled || false,
    padding_top: widgetConfig.padding_top || 0,
    padding_bottom: widgetConfig.padding_bottom || 0,
    margin_left: widgetConfig.margin_left || 0,
    margin_right: widgetConfig.margin_right || 0,
    title: widgetConfig.title || ''
  }
  
  return finalConfig
})

const bannerImage = ref(null)
const isLazyLoadingEnabled = ref(false)

// بنرهای دسکتاپ و موبایل - از آرایه‌های جداگانه banners و mobile_banners استفاده می‌کنیم
const desktopBanners = computed(() => config.value.banners ?? [])
const mobileBanners = computed(() => config.value.mobile_banners ?? [])

onMounted(() => {
  // فعال کردن lazy loading بعد از mount
  setTimeout(() => {
    isLazyLoadingEnabled.value = config.value.easy_load_enabled === true
  }, 100)
})

</script>

<style scoped>
.full-banner {
  overflow: hidden;
}

/* تنظیم ارتفاع بنر */
.responsive-banner-image {
  width: 100%;
  object-fit: cover;
  height: var(--desktop-height);
}

@media (max-width: 767px) {
  .responsive-banner-image {
    height: var(--mobile-height);
  }
  
  .full-banner {
    width: var(--mobile-width) !important;
    max-width: var(--mobile-width) !important;
    margin: 0 auto !important;
  }
}

/* برای حالت تمام عرض */
.full-banner.full-width {
  width: 100vw ;
  max-width: 100vw;
  margin-left: calc(-50vw + 50%) ;
  margin-right: calc(-50vw + 50%) ;
}

/* برای حالت وسط */
.full-banner.center-width {
  margin-left: auto ;
  margin-right: auto ;
  display: block ;
}
</style>
