<template>
  <!-- کانتینر تب‌های دسکتاپ و موبایل -->
  <div class="mt-6 mx-4">
    <!-- تب‌ها -->
    <div class="mb-6">
      <nav class="flex rounded-lg overflow-hidden">
        <button
          :class="[
            'w-1/2 text-center py-3 font-medium text-sm transition-colors',
            activeTab === 'desktop'
              ? 'bg-purple-200 text-purple-800'
              : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
          ]"
          @click="activeTab = 'desktop'"
        >
          دسکتاپ
        </button>
        <button
          :class="[
            'w-1/2 text-center py-3 font-medium text-sm transition-colors',
            activeTab === 'mobile'
              ? 'bg-purple-200 text-purple-800'
              : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
          ]"
          @click="activeTab = 'mobile'"
        >
          موبایل
        </button>
      </nav>
    </div>

    <!-- محتوای تب دسکتاپ -->
    <div v-if="activeTab === 'desktop'">
      <slot name="desktop-content"></slot>
    </div>

    <!-- محتوای تب موبایل -->
    <div v-if="activeTab === 'mobile'">
      <slot name="mobile-content"></slot>
      
      <!-- پیش‌نمایش موبایل -->
      <div class="bg-white rounded-xl shadow p-8 mt-6">
        <h3 class="text-lg font-bold text-gray-700 mb-6">پیش‌نمایش موبایل</h3>
        
        <!-- پیش‌نمایش بنرهای چهارتایی در موبایل -->
        <div class="bg-gray-50 rounded-lg p-6 border-2 border-dashed border-gray-300">
          <div class="space-y-4">
            <!-- بنر اول -->
            <div class="relative overflow-hidden rounded-lg shadow-lg">
              <div 
                class="relative"
                :style="{
                  height: `${localBannerConfig.mobile_height || 150}px`,
                  backgroundColor: localBannerConfig.bg_enabled ? localBannerConfig.bg_color : 'transparent'
                }"
              >
                <img
                  v-if="localBannerConfig.banners && localBannerConfig.banners.length > 0"
                  :src="getMobileImageUrl(localBannerConfig.banners[0])"
                  :alt="localBannerConfig.banners[0].title"
                  class="w-full h-full object-cover"
                />
                <div v-else class="flex items-center justify-center h-full text-gray-400">
                  <div class="text-center">
                    <svg class="w-12 h-12 mx-auto mb-2 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                    </svg>
                    <p class="text-sm">بنر اول</p>
                  </div>
                </div>
                <div
                  v-if="localBannerConfig.show_title || localBannerConfig.show_description"
                  class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-6"
                >
                  <h4
                    v-if="localBannerConfig.show_title && localBannerConfig.banners && localBannerConfig.banners[0] && localBannerConfig.banners[0].title"
                    class="text-white text-lg font-bold mb-2"
                  >
                    {{ localBannerConfig.banners[0].title }}
                  </h4>
                  <p
                    v-if="localBannerConfig.show_description && localBannerConfig.banners && localBannerConfig.banners[0] && localBannerConfig.banners[0].description"
                    class="text-white/90 text-sm"
                  >
                    {{ localBannerConfig.banners[0].description }}
                  </p>
                </div>
              </div>
            </div>

            <!-- بنر دوم -->
            <div v-if="localBannerConfig.banners && localBannerConfig.banners.length > 1" class="relative overflow-hidden rounded-lg shadow-lg">
              <div 
                class="relative"
                :style="{
                  height: `${localBannerConfig.mobile_height || 150}px`,
                  backgroundColor: localBannerConfig.bg_enabled ? localBannerConfig.bg_color : 'transparent'
                }"
              >
                <img
                  :src="getMobileImageUrl(localBannerConfig.banners[1])"
                  :alt="localBannerConfig.banners[1].title"
                  class="w-full h-full object-cover"
                />
                <div
                  v-if="localBannerConfig.show_title || localBannerConfig.show_description"
                  class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-6"
                >
                  <h4
                    v-if="localBannerConfig.show_title && localBannerConfig.banners[1].title"
                    class="text-white text-lg font-bold mb-2"
                  >
                    {{ localBannerConfig.banners[1].title }}
                  </h4>
                  <p
                    v-if="localBannerConfig.show_description && localBannerConfig.banners[1].description"
                    class="text-white/90 text-sm"
                  >
                    {{ localBannerConfig.banners[1].description }}
                  </p>
                </div>
              </div>
            </div>

            <!-- بنر سوم -->
            <div v-if="localBannerConfig.banners && localBannerConfig.banners.length > 2" class="relative overflow-hidden rounded-lg shadow-lg">
              <div 
                class="relative"
                :style="{
                  height: `${localBannerConfig.mobile_height || 150}px`,
                  backgroundColor: localBannerConfig.bg_enabled ? localBannerConfig.bg_color : 'transparent'
                }"
              >
                <img
                  :src="getMobileImageUrl(localBannerConfig.banners[2])"
                  :alt="localBannerConfig.banners[2].title"
                  class="w-full h-full object-cover"
                />
                <div
                  v-if="localBannerConfig.show_title || localBannerConfig.show_description"
                  class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-6"
                >
                  <h4
                    v-if="localBannerConfig.show_title && localBannerConfig.banners[2].title"
                    class="text-white text-lg font-bold mb-2"
                  >
                    {{ localBannerConfig.banners[2].title }}
                  </h4>
                  <p
                    v-if="localBannerConfig.show_description && localBannerConfig.banners[2].description"
                    class="text-white/90 text-sm"
                  >
                    {{ localBannerConfig.banners[2].description }}
                  </p>
                </div>
              </div>
            </div>

            <!-- بنر چهارم -->
            <div v-if="localBannerConfig.banners && localBannerConfig.banners.length > 3" class="relative overflow-hidden rounded-lg shadow-lg">
              <div 
                class="relative"
                :style="{
                  height: `${localBannerConfig.mobile_height || 150}px`,
                  backgroundColor: localBannerConfig.bg_enabled ? localBannerConfig.bg_color : 'transparent'
                }"
              >
                <img
                  :src="getMobileImageUrl(localBannerConfig.banners[3])"
                  :alt="localBannerConfig.banners[3].title"
                  class="w-full h-full object-cover"
                />
                <div
                  v-if="localBannerConfig.show_title || localBannerConfig.show_description"
                  class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-6"
                >
                  <h4
                    v-if="localBannerConfig.show_title && localBannerConfig.banners[3].title"
                    class="text-white text-lg font-bold mb-2"
                  >
                    {{ localBannerConfig.banners[3].title }}
                  </h4>
                  <p
                    v-if="localBannerConfig.show_description && localBannerConfig.banners[3].description"
                    class="text-white/90 text-sm"
                  >
                    {{ localBannerConfig.banners[3].description }}
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      
      <!-- تنظیمات موبایل -->
      <div class="bg-white rounded-xl shadow p-6 mt-6">
        <h3 class="text-lg font-bold text-gray-700 mb-6">تنظیمات موبایل</h3>
        
        <!-- حالت نمایش موبایل -->
        <div class="mb-6">
          <label class="block text-sm font-medium text-gray-700 mb-3">حالت نمایش موبایل</label>
          <div class="space-y-3">
            <label class="flex items-center">
              <input
                v-model="localBannerConfig.mobile_image_mode"
                type="radio"
                value="auto"
                class="mr-3 text-purple-600 focus:ring-purple-500"
              />
              <span class="text-sm text-gray-700">برش خودکار عکس</span>
            </label>
            <label class="flex items-center">
              <input
                v-model="localBannerConfig.mobile_image_mode"
                type="radio"
                value="separate"
                class="mr-3 text-purple-600 focus:ring-purple-500"
              />
              <span class="text-sm text-gray-700">عکس جداگانه موبایل</span>
            </label>
          </div>
        </div>

        <!-- تنظیمات برش عکس -->
        <div v-if="localBannerConfig.mobile_image_mode === 'auto'" class="mb-6">
          <div class="grid grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">عرض برش (پیکسل)</label>
              <input
                v-model.number="localBannerConfig.mobile_crop_width"
                type="number"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-purple-500 focus:border-purple-500"
                placeholder="400"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">ارتفاع برش (پیکسل)</label>
              <input
                v-model.number="localBannerConfig.mobile_crop_height"
                type="number"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-purple-500 focus:border-purple-500"
                placeholder="200"
              />
            </div>
          </div>
          
          <!-- دکمه اعمال برش -->
          <div class="mt-4">
            <button
              :disabled="!localBannerConfig.banners || localBannerConfig.banners.length === 0"
              class="px-4 py-2 bg-purple-600 text-white rounded-md hover:bg-purple-700 disabled:bg-gray-400 disabled:cursor-not-allowed"
              @click="applyMobileCrop"
            >
              اعمال برش موبایل
            </button>
          </div>
        </div>

        <!-- آپلود عکس موبایل جداگانه -->
        <div v-if="localBannerConfig.mobile_image_mode === 'separate'" class="mb-6">
          <label class="block text-sm font-medium text-gray-700 mb-3">عکس موبایل</label>
          <div class="border-2 border-dashed border-gray-300 rounded-lg p-6 text-center">
            <div v-if="localBannerConfig.mobile_image" class="mb-4">
              <img
                :src="localBannerConfig.mobile_image"
                alt="عکس موبایل"
                class="max-h-32 mx-auto rounded-lg shadow-sm"
              />
            </div>
            <div v-else class="text-gray-400 mb-4">
              <svg class="w-12 h-12 mx-auto mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
              </svg>
              <p class="text-sm">عکس موبایل انتخاب نشده</p>
            </div>
            <button
              class="px-4 py-2 bg-purple-600 text-white rounded-md hover:bg-purple-700"
              @click="openMobileImageSelector"
            >
              انتخاب عکس موبایل
            </button>
            <button
              v-if="localBannerConfig.mobile_image"
              class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-700 mr-2"
              @click="removeMobileImage"
            >
              حذف عکس
            </button>
          </div>
        </div>

        <!-- تنظیمات ارتفاع موبایل -->
        <div class="mb-6">
          <label class="block text-sm font-medium text-gray-700 mb-2">ارتفاع موبایل (پیکسل)</label>
          <input
            v-model.number="localBannerConfig.mobile_height"
            type="number"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-purple-500 focus:border-purple-500"
            placeholder="150"
          />
        </div>

        <!-- تنظیمات فاصله موبایل -->
        <div class="grid grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">فاصله بالا (پیکسل)</label>
            <input
              v-model.number="localBannerConfig.mobile_padding_top"
              type="number"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-purple-500 focus:border-purple-500"
              placeholder="0"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">فاصله پایین (پیکسل)</label>
            <input
              v-model.number="localBannerConfig.mobile_padding_bottom"
              type="number"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-purple-500 focus:border-purple-500"
              placeholder="0"
            />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// Vue composables
import { computed, ref } from 'vue'

// Props
interface Props {
  bannerConfig: any
  currentPreviewBanner: number
  openAddBannerModal: () => void
  editBanner: (index: number) => void
  removeBanner: (index: number) => void
  getBannerClasses: () => string
  getBannerWidth: () => string
  openMobileImageSelector?: () => void
  removeMobileImage?: () => void
}

const props = defineProps<Props>()

const emit = defineEmits(['update:bannerConfig'])

// Tab state
const activeTab = ref<'desktop' | 'mobile'>('desktop')

// Expose activeTab for parent component
defineExpose({
  activeTab
})

// Computed values with default fallbacks
const localBannerConfig = computed({
  get: () => new Proxy(props.bannerConfig, {
    set(target, key, value) {
      emit('update:bannerConfig', { ...target, [key]: value })
      return true
    }
  }),
  set: (val) => {
    emit('update:bannerConfig', val)
  }
})

const mobileBannerHeight = computed({
  get: () => localBannerConfig.value.mobile_height,
  set: val => {
    if (localBannerConfig.value) {
      localBannerConfig.value.mobile_height = val
    }
  }
})

// Mobile crop function
const applyMobileCrop = async () => {
  if (!localBannerConfig.value.banners || localBannerConfig.value.banners.length === 0) {
    alert('ابتدا باید حداقل یک بنر اضافه کنید')
    return
  }

  const banner = localBannerConfig.value.banners[0] // استفاده از اولین بنر
  const cropWidth = localBannerConfig.value.mobile_crop_width || 400
  const cropHeight = localBannerConfig.value.mobile_crop_height || 200

  if (!banner.image) {
    alert('بنر باید عکس داشته باشد')
    return
  }

  try {
    const response = await $fetch('/api/media/image-crop', {
      method: 'POST',
      body: {
        src: banner.image,
        crop_width: cropWidth,
        crop_height: cropHeight,
        mode: 'cover',
        device: 'mobile',
        quality: 85
      }
    }) as any
    
    if (response.success) {
      localBannerConfig.value.mobile_cropped_image = response.data.cropped_url
      alert('برش عکس موبایل با موفقیت اعمال شد!')
    } else {
      throw new Error('API response failed')
    }
  } catch (error) {
    console.error('Backend crop API error:', error)
    alert('خطا در برش عکس موبایل')
  }
}

// Helper functions for mobile image handling
const getMobileImageUrl = (banner) => {
  // اگر حالت عکس جداگانه انتخاب شده و عکس موبایل وجود دارد
  if (localBannerConfig.value.mobile_image_mode === 'separate' && banner.mobile_image) {
    return banner.mobile_image
  }
  
  // اگر حالت برش خودکار انتخاب شده و عکس برش شده وجود دارد
  if (localBannerConfig.value.mobile_image_mode === 'auto' && localBannerConfig.value.mobile_cropped_image) {
    return localBannerConfig.value.mobile_cropped_image
  }
  
  // در غیر این صورت از عکس اصلی استفاده کن
  return banner.image
}

// Default functions if not provided
const openMobileImageSelector = () => {
  if (props.openMobileImageSelector) {
    props.openMobileImageSelector()
  } else {
    alert('تابع انتخاب عکس موبایل تعریف نشده')
  }
}

const removeMobileImage = () => {
  if (props.removeMobileImage) {
    props.removeMobileImage()
  } else {
    localBannerConfig.value.mobile_image = ''
  }
}
</script>
