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
      <!-- تنظیمات اسلایدر -->
      <div class="bg-white rounded-xl shadow p-8 mt-6 mx-4">
        <div class="flex justify-between items-center mb-6">
          <h3 class="text-lg font-bold text-gray-700">تنظیمات موبایل</h3>
          <div class="flex items-center gap-2 border-2 border-blue-200 rounded-lg p-1 bg-blue-50">
            <input
              type="checkbox"
              id="easyLoadMobile"
              v-model="props.sliderConfig.easy_load_enabled"
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
              v-model="props.sliderConfig.bg_enabled"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
            >
              <option :value="true">فعال</option>
              <option :value="false">غیرفعال</option>
            </select>
          </div>

          <!-- عریض پس‌زمینه -->
          <div v-if="props.sliderConfig.bg_enabled">
            <label class="block mb-2 text-sm font-medium text-gray-700">عریض پس‌زمینه</label>
            <select
              v-model="props.sliderConfig.wide_bg"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
            >
              <option :value="true">بله</option>
              <option :value="false">خیر</option>
            </select>
          </div>

          <!-- رنگ پس‌زمینه -->
          <div v-if="props.sliderConfig.bg_enabled">
            <label class="block mb-2 text-sm font-medium text-gray-700">رنگ پس‌زمینه</label>
            <input
              type="color"
              v-model="props.sliderConfig.bg_color"
              class="w-full h-10 border border-gray-300 rounded-md"
            />
          </div>

          <!-- پدینگ بالا -->
          <div>
            <label class="block mb-2 text-sm font-medium text-gray-700">پدینگ بالا (px)</label>
            <input
              type="number"
              :value="props.sliderConfig.padding_top !== undefined ? props.sliderConfig.padding_top : ''"
              @input="(e) => props.sliderConfig.padding_top = (e.target as HTMLInputElement).value === '' ? undefined : Number((e.target as HTMLInputElement).value)"
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
              type="number"
              :value="props.sliderConfig.padding_bottom !== undefined ? props.sliderConfig.padding_bottom : ''"
              @input="(e) => props.sliderConfig.padding_bottom = (e.target as HTMLInputElement).value === '' ? undefined : Number((e.target as HTMLInputElement).value)"
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
              type="number"
              :value="props.sliderConfig.margin_right !== undefined ? props.sliderConfig.margin_right : ''"
              @input="(e) => props.sliderConfig.margin_right = (e.target as HTMLInputElement).value === '' ? undefined : Number((e.target as HTMLInputElement).value)"
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
              type="number"
              :value="props.sliderConfig.margin_left !== undefined ? props.sliderConfig.margin_left : ''"
              @input="(e) => props.sliderConfig.margin_left = (e.target as HTMLInputElement).value === '' ? undefined : Number((e.target as HTMLInputElement).value)"
              min="0"
              max="100"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
              placeholder="0"
            />
          </div>



          <!-- عرض اسلایدر در موبایل -->
          <div>
            <label class="block mb-2 text-sm font-medium text-gray-700">عرض اسلایدر در موبایل</label>
            <select
              v-model="props.sliderConfig.mobile_slider_width"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
            >
              <option :value="100">کوچک (100px)</option>
              <option :value="300" selected>متوسط (300px)</option>
              <option :value="500">بزرگ (500px)</option>
            </select>
          </div>

          <!-- ارتفاع بنر موبایل -->
          <div>
            <label class="block mb-2 text-sm font-medium text-gray-700">ارتفاع بنر موبایل (پیکسل)</label>
            <select
              v-model="mobileSliderHeight"
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
            <div class="flex flex-col gap-3 p-2 relative w-64 mx-auto">
              <!-- بنرهای دوتایی در موبایل - عمودی -->
              <div v-if="props.sliderConfig.mobile_vertical_display" class="flex flex-col gap-3">
                <!-- بنر اول -->
                <div
                  class="relative overflow-hidden rounded-lg w-full"
                  :style="{
                    height: `${props.sliderConfig.mobile_height || 150}px`
                  }"
                >
                  <img
                    v-if="activeBanners.length > 0"
                    :src="activeBanners[0].image"
                    :alt="activeBanners[0].title"
                    class="w-full h-full object-cover absolute inset-0"
                  />
                  <div v-else class="flex items-center justify-center h-full text-gray-400">
                    <div class="text-center">
                      <svg class="w-8 h-8 mx-auto mb-2 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                      </svg>
                      <p class="text-xs">بنر اول</p>
                    </div>
                  </div>
                  <div
                    v-if="props.sliderConfig.show_title || props.sliderConfig.show_description"
                    class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-2 z-10"
                  >
                    <h4
                      v-if="props.sliderConfig.show_title && activeBanners[0] && activeBanners[0].title"
                      class="text-white text-sm font-bold mb-1"
                    >
                      {{ activeBanners[0].title }}
                    </h4>
                    <p
                      v-if="props.sliderConfig.show_description && activeBanners[0] && activeBanners[0].description"
                      class="text-white/90 text-xs"
                    >
                      {{ activeBanners[0].description }}
                    </p>
                  </div>
                </div>

                <!-- بنر دوم -->
                <div
                  class="relative overflow-hidden rounded-lg w-full"
                  :style="{
                    height: `${props.sliderConfig.mobile_height || 150}px`
                  }"
                >
                  <img
                    v-if="activeBanners.length > 1"
                    :src="activeBanners[1].image"
                    :alt="activeBanners[1].title"
                    class="w-full h-full object-cover absolute inset-0"
                  />
                  <div v-else class="flex items-center justify-center h-full text-gray-400">
                    <div class="text-center">
                      <svg class="w-8 h-8 mx-auto mb-2 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                      </svg>
                      <p class="text-xs">بنر دوم</p>
                    </div>
                  </div>
                  <div
                    v-if="props.sliderConfig.show_title || props.sliderConfig.show_description"
                    class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-2 z-10"
                  >
                    <h4
                      v-if="props.sliderConfig.show_title && activeBanners[1] && activeBanners[1].title"
                      class="text-white text-sm font-bold mb-1"
                    >
                      {{ activeBanners[1].title }}
                    </h4>
                    <p
                      v-if="props.sliderConfig.show_description && activeBanners[1] && activeBanners[1].description"
                      class="text-white/90 text-xs"
                    >
                      {{ activeBanners[1].description }}
                    </p>
                  </div>
                </div>
              </div>

              <!-- نمایش افقی در موبایل -->
              <div v-else class="grid grid-cols-2 gap-2">
                <!-- بنر اول -->
                <div
                  class="relative overflow-hidden rounded-lg w-full"
                  :style="{
                    height: `${props.sliderConfig.mobile_height || 150}px`
                  }"
                >
                  <img
                    v-if="activeBanners.length > 0"
                    :src="activeBanners[0].image"
                    :alt="activeBanners[0].title"
                    class="w-full h-full object-cover absolute inset-0"
                  />
                  <div v-else class="flex items-center justify-center h-full text-gray-400">
                    <div class="text-center">
                      <svg class="w-6 h-6 mx-auto mb-1 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                      </svg>
                      <p class="text-xs">1</p>
                    </div>
                  </div>
                </div>

                <!-- بنر دوم -->
                <div
                  class="relative overflow-hidden rounded-lg w-full"
                  :style="{
                    height: `${props.sliderConfig.mobile_height || 150}px`
                  }"
                >
                  <img
                    v-if="activeBanners.length > 1"
                    :src="activeBanners[1].image"
                    :alt="activeBanners[1].title"
                    class="w-full h-full object-cover absolute inset-0"
                  />
                  <div v-else class="flex items-center justify-center h-full text-gray-400">
                    <div class="text-center">
                      <svg class="w-6 h-6 mx-auto mb-1 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                      </svg>
                      <p class="text-xs">2</p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- اطلاعات تنظیمات موبایل -->
          <div class="mt-4 text-sm text-gray-600 text-center border-t border-gray-200 pt-4">
            <p>ارتفاع: {{ props.sliderConfig.mobile_height || 150 }}px |
              نمایش: {{ props.sliderConfig.mobile_vertical_display ? 'عمودی' : 'افقی' }} |
              پس‌زمینه: {{ props.sliderConfig.bg_enabled ? 'فعال' : 'غیرفعال' }}</p>
          </div>
        </div>

        <!-- نمایش و مدیریت اسلایدها -->
        <div class="bg-white rounded-xl shadow p-6 flex-1 min-h-[400px]">
          <div class="flex justify-between items-center mb-6">
            <h3 class="text-lg font-bold text-gray-700">اسلایدر اصلی</h3>
            <button
              class="bg-cyan-500 text-white px-4 py-2 rounded-md font-bold hover:bg-cyan-600 transition-colors"
              @click="props.openAddSliderModal"
            >
              افزودن اسلایدر جدید
            </button>
          </div>

          <div v-if="activeBanners.length === 0" class="text-gray-400 text-center py-8">
            چیزی برای نمایش وجود ندارد!
          </div>

          <div v-else class="flex flex-col gap-6">
            <div
              v-for="(slide, idx) in activeBanners"
              :key="idx"
              class="flex items-center gap-3 p-2 rounded-lg bg-gray-50 w-full"
            >
              <img
                :src="slide.image"
                alt="اسلایدر"
                class="w-28 h-20 object-cover rounded border-2 border-purple-200"
              />
              <div class="flex flex-col flex-1">
                <div class="font-bold text-sm text-gray-700 mb-1">{{ slide.title }}</div>
                <div v-if="slide.description" class="text-xs text-gray-600 mb-1">{{ slide.description }}</div>
                <div v-if="slide.link" class="text-xs text-blue-600 break-all">{{ slide.link }}</div>
              </div>
              <div class="flex gap-2">
                <button
                  @click="props.editSlide(idx)"
                  class="text-blue-500 hover:text-blue-700 p-1"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536M9 11l6 6M3 17v4h4l10.293-10.293a1 1 0 00-1.414-1.414L3 17z"></path>
                  </svg>
                </button>
                <button
                  @click="props.removeSlide(idx)"
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
  sliderConfig: any
  currentPreviewSlide: number
  openAddSliderModal: () => void
  editSlide: (index: number) => void
  removeSlide: (index: number) => void
  openAddBannerModal: () => void
  editBanner: (index: number) => void
  removeBanner: (index: number) => void
  getSliderClasses: () => string
  getSliderWidth: () => string
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
    return props.sliderConfig.mobile_banners || []
  }
  return props.sliderConfig.banners || []
})

// Computed values with default fallbacks
const mobileSliderHeight = computed({
  get: () => props.sliderConfig.mobile_height,
  set: val => {
    if (props.sliderConfig) {
      props.sliderConfig.mobile_height = val
    }
  }
})
</script>
