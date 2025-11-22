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
      <!-- تنظیمات اسلایدر -->
      <div class="bg-white rounded-xl shadow p-8 mt-6 mx-4">
        <div class="flex justify-between items-center mb-6">
          <h3 class="text-lg font-bold text-gray-700">تنظیمات موبایل</h3>
          <div class="flex items-center gap-2 border-2 border-blue-200 rounded-lg p-1 bg-blue-50">
            <input
              id="easyLoadMobile"
              v-model="localSliderConfig.easy_load_enabled"
              type="checkbox"
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
              v-model="localSliderConfig.bg_enabled"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
            >
              <option :value="true">فعال</option>
              <option :value="false">غیرفعال</option>
            </select>
          </div>

          <!-- عریض پس‌زمینه -->
          <div v-if="localSliderConfig.bg_enabled">
            <label class="block mb-2 text-sm font-medium text-gray-700">عریض پس‌زمینه</label>
            <select
              v-model="localSliderConfig.wide_bg"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
            >
              <option :value="true">بله</option>
              <option :value="false">خیر</option>
            </select>
          </div>

          <!-- رنگ پس‌زمینه -->
          <div v-if="localSliderConfig.bg_enabled">
            <label class="block mb-2 text-sm font-medium text-gray-700">رنگ پس‌زمینه</label>
            <input
              v-model="localSliderConfig.bg_color"
              type="color"
              class="w-full h-10 border border-gray-300 rounded-md"
            />
          </div>

          <!-- پدینگ بالا -->
          <div>
            <label class="block mb-2 text-sm font-medium text-gray-700">پدینگ بالا (px)</label>
            <input
              v-model="localSliderConfig.padding_top"
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
              v-model="localSliderConfig.padding_bottom"
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
              v-model="localSliderConfig.margin_right"
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
              v-model="localSliderConfig.margin_left"
              type="number"
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
              v-model="localSliderConfig.mobile_slider_width"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
            >
              <option :value="100">کوچک (100px)</option>
              <option :value="300" selected>متوسط (300px)</option>
              <option :value="500">بزرگ (500px)</option>
            </select>
          </div>

          <!-- ارتفاع اسلایدر -->
          <div>
            <label class="block mb-2 text-sm font-medium text-gray-700">ارتفاع اسلایدر (پیکسل)</label>
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
        <!-- پیش‌نمایش زنده -->
        <div class="rounded-xl shadow p-2 ml-auto max-w-xs">
          <h3 class="text-lg font-bold text-gray-700 mb-2">پیش‌نمایش زنده موبایل</h3>
          
          <div class="bg-white rounded-lg p-2 border-2 border-dashed border-gray-300">
            <!-- Layout container for mobile - vertical stack -->
            <div class="flex flex-col gap-3 p-2 min-h-[500px] relative w-64 mx-auto">
              <!-- بنرهای دوتایی در موبایل - عمودی -->
              <div v-if="localSliderConfig.mobile_vertical_display" class="flex flex-col gap-3">
                <!-- بنر اول -->
                <div
                  class="relative overflow-hidden rounded-lg w-full"
                  :style="{
                    height: `${localSliderConfig.mobile_height || 150}px`,
                    backgroundColor: localSliderConfig.bg_enabled ? localSliderConfig.bg_color : 'transparent'
                  }"
                >
                  <img
                    v-if="sliderSlides.length > 0"
                    :src="sliderSlides[0].image"
                    :alt="sliderSlides[0].title"
                    class="w-full h-full object-cover"
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
                    v-if="localSliderConfig.show_title || localSliderConfig.show_description"
                    class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-2"
                  >
                    <h4
                      v-if="localSliderConfig.show_title && sliderSlides.length > 0 && sliderSlides[0].title"
                      class="text-white text-sm font-bold mb-1"
                    >
                      {{ sliderSlides[0].title }}
                    </h4>
                    <p
                      v-if="localSliderConfig.show_description && sliderSlides.length > 0 && sliderSlides[0].description"
                      class="text-white/90 text-xs"
                    >
                      {{ sliderSlides[0].description }}
                    </p>
                  </div>
                </div>

                <!-- بنر دوم -->
                <div
                  class="relative overflow-hidden rounded-lg w-full"
                  :style="{
                    height: `${localSliderConfig.mobile_height || 150}px`,
                    backgroundColor: localSliderConfig.bg_enabled ? localSliderConfig.bg_color : 'transparent'
                  }"
                >
                  <img
                    v-if="sliderSlides.length > 1"
                    :src="sliderSlides[1].image"
                    :alt="sliderSlides[1].title"
                    class="w-full h-full object-cover"
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
                    v-if="localSliderConfig.show_title || localSliderConfig.show_description"
                    class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-2"
                  >
                    <h4
                      v-if="localSliderConfig.show_title && sliderSlides.length > 1 && sliderSlides[1].title"
                      class="text-white text-sm font-bold mb-1"
                    >
                      {{ sliderSlides[1].title }}
                    </h4>
                    <p
                      v-if="localSliderConfig.show_description && sliderSlides.length > 1 && sliderSlides[1].description"
                      class="text-white/90 text-xs"
                    >
                      {{ sliderSlides[1].description }}
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
                    height: `${localSliderConfig.mobile_height || 150}px`,
                    backgroundColor: localSliderConfig.bg_enabled ? localSliderConfig.bg_color : 'transparent'
                  }"
                >
                  <img
                    v-if="sliderSlides.length > 0"
                    :src="sliderSlides[0].image"
                    :alt="sliderSlides[0].title"
                    class="w-full h-full object-cover"
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
                    height: `${localSliderConfig.mobile_height || 150}px`,
                    backgroundColor: localSliderConfig.bg_enabled ? localSliderConfig.bg_color : 'transparent'
                  }"
                >
                  <img
                    v-if="sliderSlides.length > 1"
                    :src="sliderSlides[1].image"
                    :alt="sliderSlides[1].title"
                    class="w-full h-full object-cover"
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
            <p>ارتفاع: {{ localSliderConfig.mobile_height || 150 }}px |
              نمایش: {{ localSliderConfig.mobile_vertical_display ? 'عمودی' : 'افقی' }} |
              پس‌زمینه: {{ localSliderConfig.bg_enabled ? 'فعال' : 'غیرفعال' }}</p>
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

          <div v-if="sliderSlides.length === 0" class="text-gray-400 text-center py-8">
            چیزی برای نمایش وجود ندارد!
          </div>

          <div v-else class="flex flex-col gap-6">
            <div
              v-for="(slide, idx) in sliderSlides"
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
                  class="text-blue-500 hover:text-blue-700 p-1"
                  @click="props.editSlide(idx)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536M9 11l6 6M3 17v4h4l10.293-10.293a1 1 0 00-1.414-1.414L3 17z"></path>
                  </svg>
                </button>
                <button
                  class="text-red-500 hover:text-red-700 p-1"
                  @click="props.removeSlide(idx)"
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
import { computed, ref } from 'vue'
import type { BannerConfig, SliderConfig } from '~/types/widget'

// Props
interface Props {
  sliderConfig: BannerConfig | SliderConfig
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

const emit = defineEmits<{
  'update:sliderConfig': [value: BannerConfig | SliderConfig]
}>()

// Helper function to check if config is SliderConfig
function isSliderConfig(config: BannerConfig | SliderConfig): config is SliderConfig {
  return 'slides' in config
}

// Proxy for sliderConfig to avoid mutating props
const localSliderConfig = computed({
  get: () => new Proxy(props.sliderConfig, {
    set(obj, prop, value) {
      emit('update:sliderConfig', { ...obj, [prop]: value } as BannerConfig | SliderConfig)
      return true
    }
  }) as BannerConfig | SliderConfig,
  set: (val) => emit('update:sliderConfig', val as BannerConfig | SliderConfig)
})

// Computed property for slides with type guard
const sliderSlides = computed(() => {
  return isSliderConfig(localSliderConfig.value) ? localSliderConfig.value.slides : []
})

// Tab state
const activeTab = ref<'desktop' | 'mobile'>('desktop')

// Expose activeTab for parent component
defineExpose({
  activeTab
})

// Computed values with default fallbacks
const mobileSliderHeight = computed({
  get: () => localSliderConfig.value.mobile_height,
  set: val => {
    if (isSliderConfig(localSliderConfig.value)) {
      emit('update:sliderConfig', { ...localSliderConfig.value, mobile_height: val } as SliderConfig)
    } else {
      emit('update:sliderConfig', { ...localSliderConfig.value, mobile_height: val } as BannerConfig)
    }
  }
})
</script>
