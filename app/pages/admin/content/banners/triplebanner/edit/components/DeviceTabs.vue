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
  <div class="bg-white rounded-xl p-8 mt-6">
        <h3 class="text-lg font-bold text-gray-700 mb-6">پیش‌نمایش موبایل</h3>
        
        <!-- پیش‌نمایش بنرهای سه‌تایی در موبایل -->
        <div class="bg-gray-50 rounded-lg p-6 border-2 border-dashed border-gray-300">
          <div class="space-y-4">
            <!-- بنر اول -->
            <div class="relative overflow-hidden rounded-lg">
              <div 
                class="relative"
                :style="{
                  height: `${bannerConfig.mobile_height || 150}px`,
                  backgroundColor: bannerConfig.bg_enabled ? bannerConfig.bg_color : 'transparent'
                }"
              >
                <img
                  v-if="activeBanners.length > 0"
                  :src="activeBanners[0].image"
                  :alt="activeBanners[0].title"
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
                    v-if="bannerConfig.show_title && activeBanners[0] && activeBanners[0].title"
                    class="text-white text-lg font-bold mb-2"
                  >
                    {{ activeBanners[0].title }}
                  </h4>
                  <p
                    v-if="bannerConfig.show_description && activeBanners[0] && activeBanners[0].description"
                    class="text-white/90 text-sm"
                  >
                    {{ activeBanners[0].description }}
                  </p>
                </div>
              </div>
            </div>

            <!-- بنر دوم -->
            <div v-if="activeBanners.length > 1" class="relative overflow-hidden rounded-lg">
              <div 
                class="relative"
                :style="{
                  height: `${bannerConfig.mobile_height || 150}px`,
                  backgroundColor: bannerConfig.bg_enabled ? bannerConfig.bg_color : 'transparent'
                }"
              >
                <img
                  :src="activeBanners[1].image"
                  :alt="activeBanners[1].title"
                  class="w-full h-full object-cover"
                />
                <div
                  v-if="bannerConfig.show_title || bannerConfig.show_description"
                  class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-6"
                >
                  <h4
                    v-if="bannerConfig.show_title && activeBanners[1].title"
                    class="text-white text-lg font-bold mb-2"
                  >
                    {{ activeBanners[1].title }}
                  </h4>
                  <p
                    v-if="bannerConfig.show_description && activeBanners[1].description"
                    class="text-white/90 text-sm"
                  >
                    {{ activeBanners[1].description }}
                  </p>
                </div>
              </div>
            </div>

            <!-- بنر سوم -->
            <div v-if="activeBanners.length > 2" class="relative overflow-hidden rounded-lg">
              <div 
                class="relative"
                :style="{
                  height: `${bannerConfig.mobile_height || 150}px`,
                  backgroundColor: bannerConfig.bg_enabled ? bannerConfig.bg_color : 'transparent'
                }"
              >
                <img
                  :src="activeBanners[2].image"
                  :alt="activeBanners[2].title"
                  class="w-full h-full object-cover"
                />
                <div
                  v-if="bannerConfig.show_title || bannerConfig.show_description"
                  class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-6"
                >
                  <h4
                    v-if="bannerConfig.show_title && activeBanners[2].title"
                    class="text-white text-lg font-bold mb-2"
                  >
                    {{ activeBanners[2].title }}
                  </h4>
                  <p
                    v-if="bannerConfig.show_description && activeBanners[2].description"
                    class="text-white/90 text-sm"
                  >
                    {{ activeBanners[2].description }}
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      
      <!-- تنظیمات موبایل -->
  <div class="bg-white rounded-xl p-6 mt-6">
        <h3 class="text-lg font-bold text-gray-700 mb-6">تنظیمات موبایل</h3>

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
}

const props = defineProps<Props>()

const emit = defineEmits(['update:bannerConfig'])

// Tab state
const activeTab = ref<'desktop' | 'mobile'>('desktop')

// Expose activeTab for parent component
defineExpose({
  activeTab
})

const activeBanners = computed(() => {
  if (activeTab.value === 'mobile') {
    return props.bannerConfig?.mobile_banners || []
  }
  return props.bannerConfig?.banners || []
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
</script>
