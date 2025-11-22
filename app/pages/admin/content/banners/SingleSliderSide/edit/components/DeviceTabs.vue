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
              v-model="easyLoadEnabled"
              type="checkbox"
              class="w-4 h-4 text-blue-600 bg-blue-100 border-blue-300 rounded focus:ring-blue-500 focus:ring-2"
            />
            <label for="easyLoadMobile" class="text-sm font-medium text-blue-700">لیزی لود</label>
          </div>
        </div>
        
        <div class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-6 gap-6">
          <!-- پدینگ بالا -->
          <div>
            <label class="block mb-2 text-sm font-medium text-gray-700">پدینگ بالا (px)</label>
            <input
              v-model.number="paddingTop"
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
              v-model.number="paddingBottom"
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
              v-model.number="marginRight"
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
              v-model.number="marginLeft"
              type="number"
              min="0"
              max="100"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
              placeholder="0"
            />
          </div>

          <!-- پخش خودکار -->
          <div>
            <label class="block mb-2 text-sm font-medium text-gray-700">پخش خودکار</label>
            <select
              v-model="autoplayEnabled"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
            >
              <option :value="true">فعال</option>
              <option :value="false">غیرفعال</option>
            </select>
          </div>

          <!-- تاخیر پخش خودکار -->
          <div>
            <label class="block mb-2 text-sm font-medium text-gray-700">تاخیر پخش خودکار (ثانیه)</label>
            <select
              v-model.number="autoplayDelay"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
            >
              <option :value="3">3 ثانیه</option>
              <option :value="5">5 ثانیه</option>
              <option :value="7">7 ثانیه</option>
              <option :value="10">10 ثانیه</option>
              <option :value="15">15 ثانیه</option>
            </select>
          </div>

          <!-- حلقه اسلایدر -->
          <div>
            <label class="block mb-2 text-sm font-medium text-gray-700">حلقه اسلایدر</label>
            <select
              v-model="loopEnabled"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
            >
              <option :value="true">فعال</option>
              <option :value="false">غیرفعال</option>
            </select>
          </div>

          <!-- نمایش دکمه‌های ناوبری -->
          <div>
            <label class="block mb-2 text-sm font-medium text-gray-700">نمایش دکمه‌های ناوبری</label>
            <select
              v-model="showNavigation"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
            >
              <option :value="true">نمایش</option>
              <option :value="false">مخفی</option>
            </select>
          </div>

          <!-- نمایش نقطه‌های ناوبری -->
          <div>
            <label class="block mb-2 text-sm font-medium text-gray-700">نمایش نقطه‌های ناوبری</label>
            <select
              v-model="showPagination"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
            >
              <option :value="true">نمایش</option>
              <option :value="false">مخفی</option>
            </select>
          </div>

          <!-- عرض اسلایدر در موبایل -->
          <div>
            <label class="block mb-2 text-sm font-medium text-gray-700">عرض اسلایدر در موبایل</label>
            <select
              v-model.number="mobileSliderWidth"
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
              <!-- Slider Section - Full width for mobile -->
              <div
                class="relative overflow-hidden rounded-lg w-full"
                :style="{
                  height: `${props.sliderConfig.mobile_height || 150}px`,
                  position: 'relative',
                }"
              >
                <!-- نمایش اسلایدها -->
                <div v-if="props.sliderConfig.slides.length > 0" class="relative h-full">
                  <div 
                    v-for="(slide, index) in props.sliderConfig.slides" 
                    :key="index"
                    class="absolute inset-0 transition-opacity duration-500"
                    :class="{ 'opacity-100': index === props.currentPreviewSlide, 'opacity-0': index !== props.currentPreviewSlide }"
                  >
                    <img 
                      :src="slide.image" 
                      :alt="slide.title"
                      class="w-full h-full object-cover"
                    />
                    <div
                      v-if="props.sliderConfig.show_title || props.sliderConfig.show_description"
                      class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-3"
                    >
                      <h4
                        v-if="props.sliderConfig.show_title && slide.title && (slide.showTitle !== false)"
                        class="text-white text-base font-bold mb-1"
                      >
                        {{ slide.title }}
                      </h4>
                      <p
                        v-if="props.sliderConfig.show_description && slide.description"
                        class="text-white/90 text-xs"
                      >
                        {{ slide.description }}
                      </p>
                    </div>
                  </div>
                  
                  <!-- دکمه‌های ناوبری - کوچکتر برای موبایل -->
                  <div v-if="props.sliderConfig.show_navigation" class="absolute inset-y-0 left-0 right-0 flex items-center justify-between px-2">
                    <button class="bg-white/20 hover:bg-white/30 text-white rounded-full p-1.5 transition-colors">
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"></path>
                      </svg>
                    </button>
                    <button class="bg-white/20 hover:bg-white/30 text-white rounded-full p-1.5 transition-colors">
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
                      </svg>
                    </button>
                  </div>
                  
                  <!-- نقطه‌های ناوبری - کوچکتر برای موبایل -->
                  <div v-if="(props.sliderConfig as SliderConfig).show_pagination" class="absolute bottom-2 left-1/2 transform -translate-x-1/2 flex gap-1.5">
                    <div 
                      v-for="(slide, index) in (props.sliderConfig as SliderConfig).slides" 
                      :key="index"
                      class="w-2.5 h-2.5 rounded-full transition-colors"
                      :class="index === props.currentPreviewSlide ? 'bg-white' : 'bg-white/50'"
                    ></div>
                  </div>
                </div>
                
                <!-- پیام عدم وجود اسلاید -->
                <div v-else class="flex items-center justify-center h-full text-gray-400">
                  <div class="text-center">
                    <svg class="w-12 h-12 mx-auto mb-3 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                    </svg>
                    <p class="text-sm">هیچ اسلایدی برای نمایش وجود ندارد</p>
                    <p class="text-xs text-gray-500">برای شروع، یک اسلاید جدید اضافه کنید</p>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- اطلاعات تنظیمات موبایل -->
          <div class="mt-4 text-sm text-gray-600 text-center border-t border-gray-200 pt-4">
            <p>ارتفاع اسلایدر: {{ props.sliderConfig.mobile_height || 150 }}px |
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

          <div v-if="!(props.sliderConfig as SliderConfig).slides || (props.sliderConfig as SliderConfig).slides.length === 0" class="text-gray-400 text-center py-8">
            چیزی برای نمایش وجود ندارد!
          </div>

          <div v-else class="flex flex-col gap-6">
            <div
              v-for="(slide, idx) in (props.sliderConfig as SliderConfig).slides"
              :key="idx"
              class="flex items-center gap-3 p-2 rounded-lg bg-gray-50 w-full"
            >
              <img
                :src="(slide as { image?: string }).image || ''"
                alt="اسلایدر"
                class="w-28 h-20 object-cover rounded border-2 border-purple-200"
              />
              <div class="flex flex-col flex-1">
                <div class="font-bold text-sm text-gray-700 mb-1">{{ (slide as { title?: string }).title }}</div>
                <div v-if="(slide as { description?: string }).description" class="text-xs text-gray-600 mb-1">{{ (slide as { description?: string }).description }}</div>
                <div v-if="(slide as { link?: string }).link" class="text-xs text-blue-600 break-all">{{ (slide as { link?: string }).link }}</div>
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
import type { SliderConfig } from '~/types/widget'

// Props
interface Props {
  sliderConfig: SliderConfig
  currentPreviewSlide: number
  openAddSliderModal: () => void
  editSlide: (index: number) => void
  removeSlide: (index: number) => void
  getSliderClasses: () => string
  getSliderWidth: () => string
}

const props = defineProps<Props>()

const emit = defineEmits<{
  'update:sliderConfig': [value: SliderConfig]
}>()

// Helper function to emit config updates
const updateConfig = (field: string, value: unknown) => {
  emit('update:sliderConfig', {
    ...props.sliderConfig,
    [field]: value
  })
}

// Computed properties for config fields using emit instead of direct mutation
const autoplayEnabled = computed({
  get: () => props.sliderConfig?.autoplay_enabled ?? false,
  set: (val: boolean) => updateConfig('autoplay_enabled', val)
})

const autoplayDelay = computed({
  get: () => props.sliderConfig?.autoplay_delay ?? 5,
  set: (val: number) => updateConfig('autoplay_delay', val)
})

const loopEnabled = computed({
  get: () => props.sliderConfig?.loop_enabled ?? false,
  set: (val: boolean) => updateConfig('loop_enabled', val)
})

const showNavigation = computed({
  get: () => props.sliderConfig?.show_navigation ?? false,
  set: (val: boolean) => updateConfig('show_navigation', val)
})

const showPagination = computed({
  get: () => props.sliderConfig?.show_pagination ?? false,
  set: (val: boolean) => updateConfig('show_pagination', val)
})

const mobileSliderWidth = computed({
  get: () => props.sliderConfig?.mobile_slider_width ?? 300,
  set: (val: number) => updateConfig('mobile_slider_width', val)
})

const mobileSliderHeight = computed({
  get: () => props.sliderConfig?.mobile_height ?? 150,
  set: (val: number | string) => updateConfig('mobile_height', val)
})

const easyLoadEnabled = computed({
  get: () => props.sliderConfig?.easy_load_enabled ?? false,
  set: (val: boolean) => updateConfig('easy_load_enabled', val)
})

const paddingTop = computed({
  get: () => props.sliderConfig?.padding_top ?? 0,
  set: (val: number) => updateConfig('padding_top', val)
})

const paddingBottom = computed({
  get: () => props.sliderConfig?.padding_bottom ?? 0,
  set: (val: number) => updateConfig('padding_bottom', val)
})

const marginRight = computed({
  get: () => props.sliderConfig?.margin_right ?? 0,
  set: (val: number) => updateConfig('margin_right', val)
})

const marginLeft = computed({
  get: () => props.sliderConfig?.margin_left ?? 0,
  set: (val: number) => updateConfig('margin_left', val)
})

// Tab state
const activeTab = ref<'desktop' | 'mobile'>('desktop')

// Expose activeTab for parent component
defineExpose({
  activeTab
})
</script>
