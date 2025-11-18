<template>
  <!-- کانتینر تب‌های دسکتاپ و موبایل -->
  <div class="mt-6 mx-4">
    <!-- تب‌ها -->
    <div class="mb-6">
      <nav class="flex rounded-lg overflow-hidden">
        <button
          @click="activeTab = 'desktop'"
          :class="[
            'w-1/2 text-center py-3 font-medium text-sm transition-colors',
            activeTab === 'desktop'
              ? 'bg-purple-200 text-purple-800'
              : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
          ]"
        >
          دسکتاپ
        </button>
        <button
          @click="activeTab = 'mobile'"
          :class="[
            'w-1/2 text-center py-3 font-medium text-sm transition-colors',
            activeTab === 'mobile'
              ? 'bg-purple-200 text-purple-800'
              : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
          ]"
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
                  height: `${bannerConfig.mobile_height || 150}px`,
                  backgroundColor: bannerConfig.bg_enabled ? bannerConfig.bg_color : 'transparent'
                }"
              >
                <img
                  v-if="bannerConfig.banners && bannerConfig.banners.length > 0"
                  :src="getMobileImageUrl(bannerConfig.banners[0])"
                  :alt="bannerConfig.banners[0].title"
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
                  v-if="bannerConfig.show_title || bannerConfig.show_description"
                  class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-6"
                >
                  <h4
                    v-if="bannerConfig.show_title && bannerConfig.banners && bannerConfig.banners[0] && bannerConfig.banners[0].title"
                    class="text-white text-lg font-bold mb-2"
                  >
                    {{ bannerConfig.banners[0].title }}
                  </h4>
                  <p
                    v-if="bannerConfig.show_description && bannerConfig.banners && bannerConfig.banners[0] && bannerConfig.banners[0].description"
                    class="text-white/90 text-sm"
                  >
                    {{ bannerConfig.banners[0].description }}
                  </p>
                </div>
              </div>
            </div>

            <!-- بنر دوم -->
            <div v-if="bannerConfig.banners && bannerConfig.banners.length > 1" class="relative overflow-hidden rounded-lg shadow-lg">
              <div 
                class="relative"
                :style="{
                  height: `${bannerConfig.mobile_height || 150}px`,
                  backgroundColor: bannerConfig.bg_enabled ? bannerConfig.bg_color : 'transparent'
                }"
              >
                <img
                  :src="getMobileImageUrl(bannerConfig.banners[1])"
                  :alt="bannerConfig.banners[1].title"
                  class="w-full h-full object-cover"
                />
                <div
                  v-if="bannerConfig.show_title || bannerConfig.show_description"
                  class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-6"
                >
                  <h4
                    v-if="bannerConfig.show_title && bannerConfig.banners[1].title"
                    class="text-white text-lg font-bold mb-2"
                  >
                    {{ bannerConfig.banners[1].title }}
                  </h4>
                  <p
                    v-if="bannerConfig.show_description && bannerConfig.banners[1].description"
                    class="text-white/90 text-sm"
                  >
                    {{ bannerConfig.banners[1].description }}
                  </p>
                </div>
              </div>
            </div>

            <!-- بنر سوم -->
            <div v-if="bannerConfig.banners && bannerConfig.banners.length > 2" class="relative overflow-hidden rounded-lg shadow-lg">
              <div 
                class="relative"
                :style="{
                  height: `${bannerConfig.mobile_height || 150}px`,
                  backgroundColor: bannerConfig.bg_enabled ? bannerConfig.bg_color : 'transparent'
                }"
              >
                <img
                  :src="getMobileImageUrl(bannerConfig.banners[2])"
                  :alt="bannerConfig.banners[2].title"
                  class="w-full h-full object-cover"
                />
                <div
                  v-if="bannerConfig.show_title || bannerConfig.show_description"
                  class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-6"
                >
                  <h4
                    v-if="bannerConfig.show_title && bannerConfig.banners[2].title"
                    class="text-white text-lg font-bold mb-2"
                  >
                    {{ bannerConfig.banners[2].title }}
                  </h4>
                  <p
                    v-if="bannerConfig.show_description && bannerConfig.banners[2].description"
                    class="text-white/90 text-sm"
                  >
                    {{ bannerConfig.banners[2].description }}
                  </p>
                </div>
              </div>
            </div>

            <!-- بنر چهارم -->
            <div v-if="bannerConfig.banners && bannerConfig.banners.length > 3" class="relative overflow-hidden rounded-lg shadow-lg">
              <div 
                class="relative"
                :style="{
                  height: `${bannerConfig.mobile_height || 150}px`,
                  backgroundColor: bannerConfig.bg_enabled ? bannerConfig.bg_color : 'transparent'
                }"
              >
                <img
                  :src="getMobileImageUrl(bannerConfig.banners[3])"
                  :alt="bannerConfig.banners[3].title"
                  class="w-full h-full object-cover"
                />
                <div
                  v-if="bannerConfig.show_title || bannerConfig.show_description"
                  class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-6"
                >
                  <h4
                    v-if="bannerConfig.show_title && bannerConfig.banners[3].title"
                    class="text-white text-lg font-bold mb-2"
                  >
                    {{ bannerConfig.banners[3].title }}
                  </h4>
                  <p
                    v-if="bannerConfig.show_description && bannerConfig.banners[3].description"
                    class="text-white/90 text-sm"
                  >
                    {{ bannerConfig.banners[3].description }}
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
                type="radio"
                v-model="bannerConfig.mobile_image_mode"
                value="auto"
                class="mr-3 text-purple-600 focus:ring-purple-500"
              />
              <span class="text-sm text-gray-700">برش خودکار عکس</span>
            </label>
            <label class="flex items-center">
              <input
                type="radio"
                v-model="bannerConfig.mobile_image_mode"
                value="separate"
                class="mr-3 text-purple-600 focus:ring-purple-500"
              />
              <span class="text-sm text-gray-700">عکس جداگانه موبایل</span>
            </label>
          </div>
        </div>

        <!-- تنظیمات برش عکس -->
        <div v-if="bannerConfig.mobile_image_mode === 'auto'" class="mb-6">
          <div class="grid grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">عرض برش (پیکسل)</label>
              <input
                type="number"
                v-model.number="bannerConfig.mobile_crop_width"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-purple-500 focus:border-purple-500"
                placeholder="400"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">ارتفاع برش (پیکسل)</label>
              <input
                type="number"
                v-model.number="bannerConfig.mobile_crop_height"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-purple-500 focus:border-purple-500"
                placeholder="200"
              />
            </div>
          </div>
          
          <!-- دکمه اعمال برش -->
          <div class="mt-4">
            <button
              @click="applyMobileCrop"
              :disabled="!bannerConfig.banners || bannerConfig.banners.length === 0"
              class="px-4 py-2 bg-purple-600 text-white rounded-md hover:bg-purple-700 disabled:bg-gray-400 disabled:cursor-not-allowed"
            >
              اعمال برش موبایل
            </button>
          </div>
        </div>

        <!-- آپلود عکس موبایل جداگانه -->
        <div v-if="bannerConfig.mobile_image_mode === 'separate'" class="mb-6">
          <label class="block text-sm font-medium text-gray-700 mb-3">عکس موبایل</label>
          <div class="border-2 border-dashed border-gray-300 rounded-lg p-6 text-center">
            <div v-if="bannerConfig.mobile_image" class="mb-4">
              <img
                :src="bannerConfig.mobile_image"
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
              @click="openMobileImageSelector"
              class="px-4 py-2 bg-purple-600 text-white rounded-md hover:bg-purple-700"
            >
              انتخاب عکس موبایل
            </button>
            <button
              v-if="bannerConfig.mobile_image"
              @click="removeMobileImage"
              class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-700 mr-2"
            >
              حذف عکس
            </button>
          </div>
        </div>

        <!-- تنظیمات ارتفاع موبایل -->
        <div class="mb-6">
          <label class="block text-sm font-medium text-gray-700 mb-2">ارتفاع موبایل (پیکسل)</label>
          <input
            type="number"
            v-model.number="bannerConfig.mobile_height"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-purple-500 focus:border-purple-500"
            placeholder="150"
          />
        </div>

        <!-- تنظیمات فاصله موبایل -->
        <div class="grid grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">فاصله بالا (پیکسل)</label>
            <input
              type="number"
              v-model.number="bannerConfig.mobile_padding_top"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-purple-500 focus:border-purple-500"
              placeholder="0"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">فاصله پایین (پیکسل)</label>
            <input
              type="number"
              v-model.number="bannerConfig.mobile_padding_bottom"
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
import { ref, computed } from 'vue'

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

// Tab state
const activeTab = ref<'desktop' | 'mobile'>('desktop')

// Expose activeTab for parent component
defineExpose({
  activeTab
})

// Computed values with default fallbacks
const mobileBannerHeight = computed({
  get: () => props.bannerConfig.mobile_height,
  set: val => {
    if (props.bannerConfig) {
      props.bannerConfig.mobile_height = val
    }
  }
})

// Mobile crop function
const applyMobileCrop = async () => {
  if (!props.bannerConfig.banners || props.bannerConfig.banners.length === 0) {
    alert('ابتدا باید حداقل یک بنر اضافه کنید')
    return
  }

  const banner = props.bannerConfig.banners[0] // استفاده از اولین بنر
  const cropWidth = props.bannerConfig.mobile_crop_width || 400
  const cropHeight = props.bannerConfig.mobile_crop_height || 200

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
      props.bannerConfig.mobile_cropped_image = response.data.cropped_url
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
  if (props.bannerConfig.mobile_image_mode === 'separate' && banner.mobile_image) {
    return banner.mobile_image
  }
  
  // اگر حالت برش خودکار انتخاب شده و عکس برش شده وجود دارد
  if (props.bannerConfig.mobile_image_mode === 'auto' && props.bannerConfig.mobile_cropped_image) {
    return props.bannerConfig.mobile_cropped_image
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
    props.bannerConfig.mobile_image = ''
  }
}
</script>
