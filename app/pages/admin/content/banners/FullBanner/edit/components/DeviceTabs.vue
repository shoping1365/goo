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
      <!-- تنظیمات بنر -->
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
          <!-- پس‌زمینه فعال -->
          <div>
            <label class="block mb-2 text-sm font-medium text-gray-700">پس‌زمینه فعال</label>
            <select
              v-model="bgEnabled"
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
              v-model="wideBg"
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
              v-model="bgColor"
              type="color"
              class="w-full h-10 border border-gray-300 rounded-md"
            />
          </div>

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



          <!-- عرض بنر -->
          <div>
            <label class="block mb-2 text-sm font-medium text-gray-700">عرض بنر (پیکسل)</label>
            <select
              v-model.number="mobileBannerWidth"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
            >
              <option :value="400">400px</option>
              <option :value="500">500px</option>
              <option :value="600">600px</option>
              <option :value="700">700px</option>
            </select>
          </div>

          <!-- ارتفاع بنر -->
          <div>
            <label class="block mb-2 text-sm font-medium text-gray-700">ارتفاع بنر (پیکسل)</label>
            <select
              v-model.number="mobileBannerHeight"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
            >
              <option :value="100">100px</option>
              <option :value="150" selected>150px</option>
              <option :value="200">200px</option>
              <option :value="250">250px</option>
              <option :value="300">300px</option>
              <option :value="400">400px</option>
              <option :value="500">500px</option>
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
                  height: `${props.bannerConfig.mobile_height || 150}px`,
                  position: 'relative',
                }"
              >
                <!-- نمایش اسلایدها -->
                <div v-if="mobileBanners.length > 0" class="relative h-full">
                  <div 
                    v-for="(banner, index) in mobileBanners" 
                    :key="index"
                    class="absolute inset-0 transition-opacity duration-500"
                    :class="{ 'opacity-100': index === props.currentPreviewBanner, 'opacity-0': index !== props.currentPreviewBanner }"
                  >
                    <img 
                      :src="getMobileImageUrl(banner)" 
                      :alt="(banner as { title?: string }).title || ''"
                      class="w-full h-full"
                      :class="getMobileImageClass()"
                    />
                    <div
                      v-if="props.bannerConfig.show_title || props.bannerConfig.show_description"
                      class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-3"
                    >
                      <h4
                        v-if="props.bannerConfig.show_title && (banner as { title?: string; showTitle?: boolean }).title && ((banner as { showTitle?: boolean }).showTitle !== false)"
                        class="text-white text-base font-bold mb-1"
                      >
                        {{ (banner as { title?: string }).title }}
                      </h4>
                      <p
                        v-if="props.bannerConfig.show_description && (banner as { description?: string }).description"
                        class="text-white/90 text-xs"
                      >
                        {{ (banner as { description?: string }).description }}
                      </p>
                    </div>
                  </div>
                  
                  <!-- دکمه‌های ناوبری - کوچکتر برای موبایل -->
                  <div v-if="props.bannerConfig.show_navigation" class="absolute inset-y-0 left-0 right-0 flex items-center justify-between px-2">
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
                  <div v-if="props.bannerConfig.show_pagination" class="absolute bottom-2 left-1/2 transform -translate-x-1/2 flex gap-1.5">
                    <div 
                      v-for="(banner, index) in props.bannerConfig.banners" 
                      :key="index"
                      class="w-2.5 h-2.5 rounded-full transition-colors"
                      :class="index === props.currentPreviewBanner ? 'bg-white' : 'bg-white/50'"
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
            <p>ارتفاع بنر: {{ props.bannerConfig.mobile_height || 150 }}px |
              پس‌زمینه: {{ props.bannerConfig.bg_enabled ? 'فعال' : 'غیرفعال' }}</p>
          </div>
        </div>

        <!-- نمایش و مدیریت بنرها -->
        <div class="bg-white rounded-xl shadow p-6 flex-1 min-h-[400px]">
          <div class="flex justify-between items-center mb-6">
            <h3 class="text-lg font-bold text-gray-700">بنر اصلی</h3>
            <button
              class="px-4 py-2 rounded-md font-bold transition-colors"
              :class="mobileBanners.length >= 1 
                ? 'bg-gray-400 text-gray-600 cursor-not-allowed' 
                : 'bg-cyan-500 text-white hover:bg-cyan-600'"
              :disabled="mobileBanners.length >= 1"
              @click="props.openAddBannerModal"
            >
              {{ mobileBanners.length >= 1 ? 'حداکثر یک بنر مجاز است' : 'افزودن بنر جدید' }}
            </button>
          </div>

          <div v-if="mobileBanners.length === 0" class="text-gray-400 text-center py-8">
            هیچ بنری برای موبایل ثبت نشده است!
          </div>

          <div v-else class="flex flex-col gap-6">
            <div
              v-for="(banner, idx) in mobileBanners"
              :key="idx"
              class="flex items-center gap-3 p-2 rounded-lg bg-gray-50 w-full"
            >
              <div class="w-28 h-20 flex items-center justify-center rounded border-2 border-purple-200 bg-white overflow-hidden relative">
                <img
                  v-if="(banner as { mobile_image?: string }).mobile_image"
                  :src="(banner as { mobile_image?: string }).mobile_image"
                  alt="بنر موبایل"
                  class="w-full h-full object-cover"
                />
                <span v-else class="text-xs text-gray-500 text-center px-2">بدون تصویر موبایل</span>
              </div>
              <div class="flex flex-col flex-1">
                <div class="font-bold text-sm text-gray-700 mb-1">{{ (banner as { title?: string }).title }}</div>
                <div v-if="(banner as { description?: string }).description" class="text-xs text-gray-600 mb-1">{{ (banner as { description?: string }).description }}</div>
                <div v-if="(banner as { link?: string }).link" class="text-xs text-blue-600 break-all">{{ (banner as { link?: string }).link }}</div>
              </div>
              <div class="flex gap-2">
                <button
                  class="text-blue-500 hover:text-blue-700 p-1"
                  @click="props.editBanner(idx)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536M9 11l6 6M3 17v4h4l10.293-10.293a1 1 0 00-1.414-1.414L3 17z"></path>
                  </svg>
                </button>
                <button
                  class="text-red-500 hover:text-red-700 p-1"
                  @click="props.removeBanner(idx)"
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
import { ref, computed, watch } from 'vue'

// Banner Config interface
interface BannerConfig {
  easy_load_enabled?: boolean
  bg_enabled?: boolean
  wide_bg?: boolean
  bg_color?: string
  padding_top?: number
  padding_bottom?: number
  margin_right?: number
  margin_left?: number
  mobile_banner_width?: number
  mobile_height?: number
  show_title?: boolean
  show_description?: boolean
  show_navigation?: boolean
  show_pagination?: boolean
  banners?: unknown[]
  mobile_banners?: unknown[]
  [key: string]: unknown
}

// Props
interface Props {
  bannerConfig: BannerConfig
  currentPreviewBanner: number
  openAddBannerModal: () => void
  editBanner: (index: number) => void
  removeBanner: (index: number) => void
  getBannerClasses: () => string
  getBannerWidth: () => string
  activeTab?: 'desktop' | 'mobile'
}

const props = withDefaults(defineProps<Props>(), {
  activeTab: 'desktop'
})

// Emits
const emit = defineEmits<{
  'update:activeTab': [value: 'desktop' | 'mobile']
  'update:bannerConfig': [config: BannerConfig]
}>()

// Tab state - synced with parent
const activeTab = ref<'desktop' | 'mobile'>(props.activeTab)

// Watch for prop changes
watch(() => props.activeTab, (newVal) => {
  activeTab.value = newVal
})

// Watch for local changes and emit
watch(activeTab, (newVal) => {
  emit('update:activeTab', newVal)
})

// Expose activeTab for parent component
defineExpose({
  activeTab
})

// Collections split per device
interface MobileBanner {
  title?: string
  description?: string
  mobile_image?: string
  image?: string
  link?: string
  showTitle?: boolean
  [key: string]: unknown
}
const mobileBanners = computed((): MobileBanner[] => {
  const banners = props.bannerConfig?.mobile_banners ?? []
  return Array.isArray(banners) ? banners as MobileBanner[] : []
})

// Helper functions for image handling
const getMobileImageUrl = (banner: MobileBanner): string => {
  // این تابع فقط برای پیش‌نمایش موبایل استفاده میشه
  // اگر عکس موبایل وجود دارد، استفاده کن
  if (banner.mobile_image) {
    return banner.mobile_image
  }
  
  // در غیر این صورت از عکس دسکتاپ استفاده کن
  return banner.image || ''
}

const getMobileImageClass = () => {
  return 'object-cover'
}

// Helper function to emit config updates
const updateConfig = (field: string, value: unknown) => {
  emit('update:bannerConfig', {
    ...props.bannerConfig,
    [field]: value
  })
}

// Computed values with default fallbacks - using emit instead of direct mutation
const easyLoadEnabled = computed({
  get: () => props.bannerConfig?.easy_load_enabled ?? false,
  set: (val: boolean) => updateConfig('easy_load_enabled', val)
})

const bgEnabled = computed({
  get: () => props.bannerConfig?.bg_enabled ?? false,
  set: (val: boolean) => updateConfig('bg_enabled', val)
})

const wideBg = computed({
  get: () => props.bannerConfig?.wide_bg ?? false,
  set: (val: boolean) => updateConfig('wide_bg', val)
})

const bgColor = computed({
  get: () => props.bannerConfig?.bg_color ?? '#ffffff',
  set: (val: string) => updateConfig('bg_color', val)
})

const paddingTop = computed({
  get: () => props.bannerConfig?.padding_top ?? 0,
  set: (val: number) => updateConfig('padding_top', val)
})

const paddingBottom = computed({
  get: () => props.bannerConfig?.padding_bottom ?? 0,
  set: (val: number) => updateConfig('padding_bottom', val)
})

const marginRight = computed({
  get: () => props.bannerConfig?.margin_right ?? 0,
  set: (val: number) => updateConfig('margin_right', val)
})

const marginLeft = computed({
  get: () => props.bannerConfig?.margin_left ?? 0,
  set: (val: number) => updateConfig('margin_left', val)
})

const mobileBannerWidth = computed({
  get: () => props.bannerConfig?.mobile_banner_width ?? 400,
  set: (val: number) => updateConfig('mobile_banner_width', val)
})

const mobileBannerHeight = computed({
  get: () => props.bannerConfig?.mobile_height ?? 150,
  set: (val: number) => updateConfig('mobile_height', val)
})
</script>
