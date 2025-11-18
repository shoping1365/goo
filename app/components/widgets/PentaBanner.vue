<template>
  <!-- بنر پنج تایی -->
  <div class="penta-banner" :style="containerStyle">
    <!-- اگر بنرها وجود داشتند -->
    <div v-if="config.banners && config.banners.length > 0" :style="gridStyle">
        <div 
          v-for="(banner, index) in config.banners.slice(0, 5)" 
          :key="banner.id || index"
          class="relative overflow-hidden rounded-lg shadow-lg hover:shadow-xl transition-shadow duration-300"
          :style="{
            height: `${config.height || 400}px`,
            backgroundColor: config.bg_enabled ? config.bg_color : 'transparent'
          }"
        >
          <!-- اگر لینک وجود داشت، کل بنر کلیک‌پذیر باشد -->
          <a v-if="banner.link" :href="banner.link" target="_blank" rel="noopener" class="block w-full h-full">
            <img
              :src="getMobileImageUrl(banner)"
              :alt="banner.title"
              class="w-full h-full object-cover"
            />
            <!-- عنوان و توضیحات روی تصویر -->
            <div 
              v-if="(config.show_title && banner.title) || (config.show_description && banner.description)" 
              class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-6"
            >
              <h3 v-if="config.show_title && banner.title" class="text-white text-lg font-bold mb-2">
                {{ banner.title }}
              </h3>
              <p v-if="config.show_description && banner.description" class="text-white/90 text-sm">
                {{ banner.description }}
              </p>
            </div>
          </a>
          <!-- اگر لینک نبود، فقط تصویر بنر -->
          <div v-else class="w-full h-full">
            <img
              :src="getMobileImageUrl(banner)"
              :alt="banner.title"
              class="w-full h-full object-cover"
            />
            <!-- عنوان و توضیحات روی تصویر -->
            <div 
              v-if="(config.show_title && banner.title) || (config.show_description && banner.description)" 
              class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-6"
            >
              <h3 v-if="config.show_title && banner.title" class="text-white text-lg font-bold mb-2">
                {{ banner.title }}
              </h3>
              <p v-if="config.show_description && banner.description" class="text-white/90 text-sm">
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
</template>

<script setup>
const props = defineProps(['widget'])

// Get config with proper typing
const config = computed(() => {
  const widgetConfig = props.widget.config || {}
  
  const finalConfig = {
    banners: widgetConfig.banners || widgetConfig.slides || [],
    height: widgetConfig.height || 400,
    mobile_height: widgetConfig.mobile_height || widgetConfig.height || 400,
    mobile_image_mode: widgetConfig.mobile_image_mode || 'auto',
    mobile_cropped_image: widgetConfig.mobile_cropped_image || '',
    mobile_crop_width: widgetConfig.mobile_crop_width || 375,
    mobile_crop_height: widgetConfig.mobile_crop_height || 150,
    bg_enabled: widgetConfig.bg_enabled || false,
    bg_color: widgetConfig.bg_color || '#ffffff',
    show_title: widgetConfig.show_title !== undefined ? widgetConfig.show_title : true,
    show_description: widgetConfig.show_description !== undefined ? widgetConfig.show_description : true,
    banner1_ratio: widgetConfig.banner1_ratio || 20,
    banner2_ratio: widgetConfig.banner2_ratio || 20,
    banner3_ratio: widgetConfig.banner3_ratio || 20,
    banner4_ratio: widgetConfig.banner4_ratio || 20,
    banner5_ratio: widgetConfig.banner5_ratio || 20,
    padding_top: widgetConfig.padding_top || 0,
    padding_bottom: widgetConfig.padding_bottom || 0,
    margin_left: widgetConfig.margin_left || 0,
    margin_right: widgetConfig.margin_right || 0
  }
  
  return finalConfig
})

// تابع تشخیص نوع دستگاه و دریافت URL مناسب
const getMobileImageUrl = (banner) => {
  // اگر حالت عکس جداگانه انتخاب شده و عکس موبایل وجود دارد
  if (config.value.mobile_image_mode === 'separate' && banner.mobile_image) {
    return banner.mobile_image
  }

  // اگر برش خودکار اعمال شده و در موبایل هستیم
  if (config.value.mobile_cropped_image && typeof window !== 'undefined' && window.innerWidth < 768) {
    return config.value.mobile_cropped_image
  }

  // در غیر این صورت از عکس اصلی استفاده کن
  return banner.image || ''
}

// Grid style for dynamic banner ratios
const gridStyle = computed(() => {
  const styles = {
    display: 'grid',
    gap: '8px',
    gridTemplateColumns: `${config.value.banner1_ratio}fr ${config.value.banner2_ratio}fr ${config.value.banner3_ratio}fr ${config.value.banner4_ratio}fr ${config.value.banner5_ratio}fr`,
    boxSizing: 'border-box'
  }
  
  return styles
})

// Container style for width and padding/margin
const containerStyle = computed(() => {
  const styles = {
    overflow: 'hidden',
    boxSizing: 'border-box'
  }
  
  // پدینگ
  if (config.value.padding_top !== undefined) {
    styles.paddingTop = `${config.value.padding_top}px`
  }
  if (config.value.padding_bottom !== undefined) {
    styles.paddingBottom = `${config.value.padding_bottom}px`
  }
  
  // مارجین
  if (config.value.margin_left !== undefined) {
    styles.marginLeft = `${config.value.margin_left}px`
  }
  if (config.value.margin_right !== undefined) {
    styles.marginRight = `${config.value.margin_right}px`
  }
  
  return styles
})
</script>

<style scoped>

</style>