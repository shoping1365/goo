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
      <!-- تنظیمات بنر -->
      <div class="bg-white rounded-xl shadow p-8 mt-6 mx-4">
        <div class="flex justify-between items-center mb-6">
          <h3 class="text-lg font-bold text-gray-700">تنظیمات موبایل</h3>
          <div class="flex items-center gap-2 border-2 border-blue-200 rounded-lg p-1 bg-blue-50">
            <input
              type="checkbox"
              id="easyLoadMobile"
              v-model="props.bannerConfig.easy_load_enabled"
              class="w-4 h-4 text-blue-600 bg-blue-100 border-blue-300 rounded focus:ring-blue-500 focus:ring-2"
            />
            <label for="easyLoadMobile" class="text-sm font-medium text-blue-700">لیزی لود</label>
          </div>
        </div>
        
        <div class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-6 gap-6">
          <!-- پس‌زمینه فعال -->
          <div>
            <label class="block mb-2 text-sm font-medium text-gray-700">پس‌زمینه فعال</label>
            <select
              v-model="props.bannerConfig.bg_enabled"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
            >
              <option :value="true">فعال</option>
              <option :value="false">غیرفعال</option>
            </select>
          </div>

          <!-- عریض پس‌زمینه -->
          <div v-if="props.bannerConfig.bg_enabled">
            <label class="block mb-2 text-sm font-medium text-gray-700">عریض پس‌زمینه</label>
            <select
              v-model="props.bannerConfig.wide_bg"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
            >
              <option :value="true">بله</option>
              <option :value="false">خیر</option>
            </select>
          </div>

          <!-- رنگ پس‌زمینه -->
          <div v-if="props.bannerConfig.bg_enabled">
            <label class="block mb-2 text-sm font-medium text-gray-700">رنگ پس‌زمینه</label>
            <input
              type="color"
              v-model="props.bannerConfig.bg_color"
              class="w-full h-10 border border-gray-300 rounded-md"
            />
          </div>

          <!-- پدینگ بالا -->
          <div>
            <label class="block mb-2 text-sm font-medium text-gray-700">پدینگ بالا (px)</label>
            <input
              v-model="props.bannerConfig.padding_top"
              type="number"
              min="0"
              max="100"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
              placeholder="0"
            />
          </div>

          <!-- پدینگ پایین -->
          <div>
            <label class="block mb-2 text-sm font-medium text-gray-700">پدینگ پایین (px)</label>
            <input
              v-model="props.bannerConfig.padding_bottom"
              type="number"
              min="0"
              max="100"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
              placeholder="0"
            />
          </div>

          <!-- مارجین راست -->
          <div>
            <label class="block mb-2 text-sm font-medium text-gray-700">مارجین راست (px)</label>
            <input
              v-model="props.bannerConfig.margin_right"
              type="number"
              min="0"
              max="100"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
              placeholder="0"
            />
          </div>

          <!-- مارجین چپ -->
          <div>
            <label class="block mb-2 text-sm font-medium text-gray-700">مارجین چپ (px)</label>
            <input
              v-model="props.bannerConfig.margin_left"
              type="number"
              min="0"
              max="100"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
              placeholder="0"
            />
          </div>

          <!-- عرض بنر در موبایل -->
          <div>
            <label class="block mb-2 text-sm font-medium text-gray-700">عرض بنر در موبایل</label>
            <select
              v-model="props.bannerConfig.mobile_banner_width"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
            >
              <option :value="100">کوچک (100px)</option>
              <option :value="300" selected>متوسط (300px)</option>
              <option :value="500">بزرگ (500px)</option>
            </select>
          </div>

          <!-- ارتفاع بنر -->
          <div>
            <label class="block mb-2 text-sm font-medium text-gray-700">ارتفاع بنر (پیکسل)</label>
            <select
              v-model="mobileBannerHeight"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
            >
              <option :value="100">100 پیکسل</option>
              <option :value="150" selected>150 پیکسل</option>
              <option :value="200">200 پیکسل</option>
              <option :value="250">250 پیکسل</option>
            </select>
          </div>
        </div>
      </div>

      <!-- کانتینر پیش‌نمایش و مدیریت -->
      <div class="flex gap-6 mt-4 w-full">
        <!-- پیش‌نمایش زنده موبایل - فقط در تب موبایل -->
        <div v-if="activeTab === 'mobile'" class="rounded-xl shadow p-2 ml-auto max-w-xs">
          <h3 class="text-lg font-bold text-gray-700 mb-2">پیش‌نمایش زنده موبایل</h3>
          
          <div class="bg-white rounded-lg p-2 border-2 border-dashed border-gray-300">
            <!-- Layout container for mobile - vertical stack -->
            <div class="flex flex-col gap-3 p-2 min-h-[500px] relative w-64 mx-auto">
              <!-- Banner Section - Full width for mobile -->
              <div
                class="relative overflow-hidden rounded-lg w-full"
                :style="{
                  height: `${props.bannerConfig.mobile_height || 150}px`,
                  position: 'relative',
                }"
              >
                <!-- نمایش بنرها -->
                <div v-if="activeBanners.length > 0" class="relative h-full">
                  <div 
                    v-for="(banner, index) in activeBanners" 
                    :key="index"
                    class="absolute inset-0 transition-opacity duration-500"
                    :class="{ 'opacity-100': index === props.currentPreviewBanner, 'opacity-0': index !== props.currentPreviewBanner }"
                  >
                    <img 
                      :src="banner.image" 
                      :alt="banner.title"
                      class="w-full h-full object-cover"
                    />
                    <div
                      v-if="props.bannerConfig.show_title || props.bannerConfig.show_description"
                      class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-3"
                    >
                      <h4
                        v-if="props.bannerConfig.show_title && banner.title && (banner.showTitle !== false)"
                        class="text-white text-base font-bold mb-1"
                      >
                        {{ banner.title }}
                      </h4>
                      <p
                        v-if="props.bannerConfig.show_description && banner.description"
                        class="text-white/90 text-xs"
                      >
                        {{ banner.description }}
                      </p>
                    </div>
                  </div>
                </div>
                
                <!-- پیام عدم وجود بنر -->
                <div v-else class="flex items-center justify-center h-full text-gray-400">
                  <div class="text-center">
                    <svg class="w-12 h-12 mx-auto mb-3 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                    </svg>
                    <p class="text-sm">هیچ بنری برای نمایش وجود ندارد</p>
                    <p class="text-xs text-gray-500">برای شروع، یک بنر جدید اضافه کنید</p>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- اطلاعات تنظیمات موبایل -->
          <div class="mt-4 text-sm text-gray-600 text-center border-t border-gray-200 pt-4">
            <p>ارتفاع بنر: {{ props.bannerConfig.mobile_height || 150 }}px |
              پس‌زمینه: {{ props.bannerConfig.bg_enabled ? 'فعال' : 'غیرفعال' }}</p>
          </div>
        </div>

        <!-- نمایش و مدیریت بنرها -->
        <div class="bg-white rounded-xl shadow p-6 flex-1 min-h-[400px]">
          <div class="flex justify-between items-center mb-6">
            <h3 class="text-lg font-bold text-gray-700">بنر اصلی</h3>
            <button
              class="bg-cyan-500 text-white px-4 py-2 rounded-md font-bold hover:bg-cyan-600 transition-colors"
              @click="props.openAddBannerModal"
            >
              افزودن بنر جدید
            </button>
          </div>

          <div v-if="activeBanners.length === 0" class="text-gray-400 text-center py-8">
            چیزی برای نمایش وجود ندارد!
          </div>

          <div v-else class="flex flex-col gap-6">
            <div
              v-for="(banner, idx) in activeBanners"
              :key="idx"
              class="flex items-center gap-3 p-2 rounded-lg bg-gray-50 w-full"
            >
              <img
                :src="banner.image"
                alt="بنر"
                class="w-28 h-20 object-cover rounded border-2 border-purple-200"
              />
              <div class="flex flex-col flex-1">
                <div class="font-bold text-sm text-gray-700 mb-1">{{ banner.title }}</div>
                <div v-if="banner.description" class="text-xs text-gray-600 mb-1">{{ banner.description }}</div>
                <div v-if="banner.link" class="text-xs text-blue-600 break-all">{{ banner.link }}</div>
              </div>
              <div class="flex gap-2">
                <button
                  @click="props.editBanner(idx)"
                  class="text-blue-500 hover:text-blue-700 p-1"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536M9 11l6 6M3 17v4h4l10.293-10.293a1 1 0 00-1.414-1.414L3 17z"></path>
                  </svg>
                </button>
                <button
                  @click="props.removeBanner(idx)"
                  class="text-red-500 hover:text-red-700 p-1"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                  </svg>
                </button>
              </div>
            </div>
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
}

const props = defineProps<Props>()

// Tab state
const activeTab = ref<'desktop' | 'mobile'>('desktop')

// Expose activeTab for parent component
defineExpose({
  activeTab
})

// Computed برای بنرهای دسکتاپ و موبایل
const activeBanners = computed(() => {
  if (activeTab.value === 'mobile') {
    return props.bannerConfig.mobile_banners || []
  }
  return props.bannerConfig.banners || []
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
</script>
