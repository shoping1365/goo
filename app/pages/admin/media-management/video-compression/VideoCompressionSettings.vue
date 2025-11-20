<template>
  <div class="order-1 bg-white rounded-2xl shadow-lg border border-gray-100 mb-6 mt-6 mx-1 md:mx-4 w-full max-w-full overflow-x-hidden min-w-0">
    <div class="flex flex-col md:flex-row md:items-center md:justify-between px-2 md:px-6 py-4 border-b border-blue-300 relative gap-3 md:gap-0 w-full min-w-0">
      <!-- عنوان تاشو -->
      <div class="flex items-center gap-2 cursor-pointer pr-0 w-full min-w-0" @click="toggleSection('compressionSettings')">
        <span class="text-base font-bold bg-pink-100 text-rose-700 px-2 md:px-4 py-1 rounded-lg ml-auto">تنظیمات فشرده‌سازی ویدیو</span>
      </div>
      <!-- toggle icon moved to far left -->
      <span class="absolute left-2 md:left-4 top-1/2 -translate-y-1/2 text-gray-500 text-2xl select-none cursor-pointer" @click="toggleSection('compressionSettings')">{{ sections.compressionSettings ? '−' : '+' }}</span>

      <!-- سوییچ فعال‌سازی سراسری + دکمه ذخیره -->
      <div class="flex flex-col md:flex-row items-start md:items-center gap-2 md:gap-6 mt-2 md:mt-0 w-full md:w-auto min-w-0">
        <!-- Toggle: Make label and switch always in one line, aligned left -->
        <div class="flex items-center gap-2 whitespace-nowrap justify-start w-full ml-12">
          <span class="text-sm text-gray-700">فشرده‌سازی خودکار</span>
          <label class="inline-flex items-center cursor-pointer select-none">
            <input v-model="compressionSettings.enabled" type="checkbox" class="sr-only peer">
            <div class="w-11 h-6 bg-gray-200 rounded-full relative transition peer-checked:bg-green-500">
              <span class="absolute left-1 top-1 w-4 h-4 bg-white rounded-full transition-transform peer-checked:translate-x-5"></span>
            </div>
          </label>
        </div>



        <!-- Save button (shown only on unsaved changes) -->
        <button
          v-if="unsavedChanges"
          class="bg-green-500 hover:bg-green-600 text-white text-sm px-4 py-1.5 rounded-lg shadow mt-2 md:mt-0 ml-32 md:ml-40 w-full md:w-auto block md:inline-block whitespace-nowrap"
          @click="saveCompressionSettings"
        >
          ذخیره تغییرات
        </button>
      </div>
    </div>
    <div v-show="sections.compressionSettings" class="p-2 md:p-6 w-full min-w-0">
      <div :class="[ 'grid grid-cols-1 gap-6 md:grid-cols-3 md:gap-6 w-full min-w-0', !compressionSettings.enabled && 'opacity-40 pointer-events-none']">
        <!-- Quality Settings -->
        <div class="space-y-4 w-full min-w-0">
          <h3 class="font-medium text-gray-900">کیفیت خروجی</h3>
          <div class="space-y-3 w-full min-w-0">
            <label class="flex items-center w-full min-w-0">
              <input 
                v-model="compressionSettings.quality" 
                type="radio" 
                value="smart" 
                class="w-4 h-4 text-blue-600 border-gray-300 focus:ring-blue-500"
              >
              <span class="mr-2 text-sm text-blue-700 font-semibold cursor-pointer">فشرده‌سازی هوشمند (پیشنهادی)</span>
            </label>
            <label class="flex items-center w-full min-w-0">
              <input 
                v-model="compressionSettings.quality" 
                type="radio" 
                value="1080p" 
                :disabled="compressionSettings.quality === 'smart'"
                class="w-4 h-4 text-blue-600 border-gray-300 focus:ring-blue-500"
              >
              <span class="mr-2 text-sm">1080p (Full HD) - کیفیت عالی</span>
            </label>
            <label class="flex items-center w-full min-w-0">
              <input 
                v-model="compressionSettings.quality" 
                type="radio" 
                value="720p" 
                :disabled="compressionSettings.quality === 'smart'"
                class="w-4 h-4 text-blue-600 border-gray-300 focus:ring-blue-500"
              >
              <span class="mr-2 text-sm">720p (HD) - کیفیت خوب، حجم متوسط</span>
            </label>
            <label class="flex items-center w-full min-w-0">
              <input 
                v-model="compressionSettings.quality" 
                type="radio" 
                value="custom" 
                :disabled="compressionSettings.quality === 'smart'"
                class="w-4 h-4 text-blue-600 border-gray-300 focus:ring-blue-500"
              >
              <span class="mr-2 text-sm">سفارشی</span>
            </label>
            <div v-if="compressionSettings.quality === 'custom'" class="mt-3 space-y-3 w-full min-w-0">
              <div class="grid grid-cols-1 sm:grid-cols-2 gap-3 w-full min-w-0">
                <div class="w-full min-w-0">
                  <label class="block text-sm font-medium text-gray-700 mb-1">عرض</label>
                  <input
                    v-model="compressionSettings.customWidth"
                    type="number"
                    placeholder="1920"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 text-sm"
                  >
                </div>
                <div class="w-full min-w-0">
                  <label class="block text-sm font-medium text-gray-700 mb-1">ارتفاع</label>
                  <input
                    v-model="compressionSettings.customHeight"
                    type="number"
                    placeholder="1080"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 text-sm"
                  >
                </div>
              </div>
              <div class="w-full min-w-0">
                <label class="block text-sm font-medium text-gray-700 mb-1">
                  Bitrate (kbps): {{ compressionSettings.customBitrate }}
                </label>
                <input
                  v-model="compressionSettings.customBitrate"
                  type="range"
                  min="500"
                  max="10000"
                  class="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer"
                >
              </div>
            </div>
          </div>
        </div>
        <!-- Format -->
        <div class="space-y-4 w-full min-w-0">
          <h3 class="font-medium text-gray-900">فرمت خروجی</h3>
          <div class="space-y-3 w-full min-w-0">
            <label class="flex items-center w-full min-w-0">
              <input 
                v-model="compressionSettings.format" 
                type="radio" 
                value="mp4" 
                :disabled="compressionSettings.quality === 'smart'"
                class="w-4 h-4 text-blue-600 border-gray-300 focus:ring-blue-500"
              >
              <span class="mr-2 text-sm">MP4 (H.264) - سازگاری بالا</span>
            </label>
            <label class="flex items-center w-full min-w-0">
              <input 
                v-model="compressionSettings.format" 
                type="radio" 
                value="webm" 
                :disabled="compressionSettings.quality === 'smart'"
                class="w-4 h-4 text-blue-600 border-gray-300 focus:ring-blue-500"
              >
              <span class="mr-2 text-sm">WebM (VP9) - فشرده‌سازی بهتر</span>
            </label>
            <label class="flex items-center w-full min-w-0">
              <input 
                v-model="compressionSettings.format" 
                type="radio" 
                value="mov" 
                :disabled="compressionSettings.quality === 'smart'"
                class="w-4 h-4 text-blue-600 border-gray-300 focus:ring-blue-500"
              >
              <span class="mr-2 text-sm">MOV - کیفیت بالا</span>
            </label>
          </div>
        </div>
        <!-- Worker Count -->
        <div class="mt-4 w-full min-w-0">
          <label class="block text-sm font-medium text-gray-700 mb-2">تعداد ورکر</label>
          <select 
            v-model="compressionSettings.workerCount"
            :disabled="compressionSettings.quality === 'smart'"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 text-sm"
          >
            <option value="1">1 (تک هسته‌ای)</option>
            <option value="2">2</option>
            <option value="4">4</option>
            <option value="8">8</option>
            <option value="16">16</option>
            <option value="auto">خودکار (توصیه شده)</option>
          </select>
        </div>
      </div>
      <!-- Frame Rate, Codec, Frame Analysis (full width, outside grid) -->
      <div class="mt-6 w-full min-w-0">
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6 w-full min-w-0">
          <div class="w-full min-w-0">
            <label class="block text-sm font-medium text-gray-700 mb-2">نرخ فریم</label>
            <select 
              v-model="compressionSettings.frameRate"
              :disabled="compressionSettings.quality === 'smart'"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 text-sm"
            >
              <option value="original">حفظ اصلی</option>
              <option value="60">60 FPS</option>
              <option value="30">30 FPS</option>
              <option value="24">24 FPS</option>
              <option value="15">15 FPS</option>
            </select>
          </div>
          <div class="w-full min-w-0">
            <label class="block text-sm font-medium text-gray-700 mb-2">کدک ویدیو</label>
            <select 
              v-model="compressionSettings.codec"
              :disabled="compressionSettings.quality === 'smart'"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 text-sm"
            >
              <option value="h264">H.264 (AVC)</option>
              <option value="h265">H.265 (HEVC)</option>
              <option value="vp9">VP9</option>
              <option value="av1">AV1</option>
            </select>
          </div>
          <div class="w-full min-w-0">
            <label class="block text-sm font-medium text-gray-700 mb-2">تحلیل فریم</label>
            <select 
              v-model="compressionSettings.frameAnalysisMode"
              :disabled="compressionSettings.quality === 'smart'"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 text-sm"
            >
              <option value="5frames">5 فریم (سریع‌تر)</option>
              <option value="full">فول فریم (کامل)</option>
            </select>
            <p class="text-xs text-gray-600 mt-1">
              تحلیل 5 فریم سریع‌تر است، فول فریم کیفیت بهتری دارد
            </p>
          </div>

        </div>
      </div>
      <!-- Advanced Options -->
      <div class="mt-6 pt-6 border-t border-gray-200 w-full min-w-0">
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-2 sm:gap-6 mt-4 mb-2 px-1 sm:px-4 lg:px-8 w-full min-w-0">
          <label class="flex items-center w-full min-w-0">
            <input 
              v-model="compressionSettings.twoPass" 
              type="checkbox" 
              :disabled="compressionSettings.quality === 'smart'"
              class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
            >
            <span class="mr-2 text-sm">انکود دو مرحله‌ای</span>
          </label>
          <label class="flex items-center w-full min-w-0">
            <input 
              v-model="compressionSettings.removeMetadata" 
              type="checkbox" 
              :disabled="compressionSettings.quality === 'smart'"
              class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
            >
            <span class="mr-2 text-sm">حذف متادیتا</span>
          </label>

          <label class="flex items-center w-full min-w-0">
            <input 
              v-model="compressionSettings.createThumbnail" 
              type="checkbox" 
              :disabled="compressionSettings.quality === 'smart'"
              class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
            >
            <span class="mr-2 text-sm">ایجاد تصویر بندانگشتی</span>
          </label>
          <button 
            class="bg-blue-500 text-white text-xs rounded hover:bg-blue-600 transition-colors"
            style="padding: 0;"
            @click="testScheduler"
          >
            تست Scheduler
          </button>
        </div>
      </div>

      <!-- زمان‌بندی -->
      <div class="mt-6 pt-6 border-t border-gray-200 w-full min-w-0">
        <h3 class="font-medium text-gray-900 mb-4"></h3>
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-6 w-full min-w-0">
          <div class="w-full min-w-0">
            <label class="block text-sm font-medium text-gray-700 mb-2">ساعت شروع فشرده‌سازی</label>
            <select v-model="compressionSettings.startHour" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 text-sm">
              <option v-for="h in 24" :key="h-1" :value="h-1">{{ (h-1).toString().padStart(2, '0') }}:00</option>
            </select>
          </div>
          <div class="w-full min-w-0">
            <label class="block text-sm font-medium text-gray-700 mb-2">ساعت پایان فشرده‌سازی</label>
            <select v-model="compressionSettings.endHour" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 text-sm">
              <option v-for="h in 24" :key="h-1" :value="h-1">{{ (h-1).toString().padStart(2, '0') }}:00</option>
            </select>
          </div>
        </div>

      </div>

    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted, watch } from 'vue'

const emit = defineEmits(['update:enabled'])
// Reactive data
const compressionSettings = reactive({
  quality: '1080p',
  customWidth: 1920,
  customHeight: 1080,
  customBitrate: 2000,
  format: 'mp4',
  codec: 'h264',
  frameRate: 'original',
  twoPass: false,
  removeMetadata: true,
  createThumbnail: true,
  enabled: true,
  workerCount: 'auto',
  frameAnalysisMode: '5frames',
  startHour: 1,
  endHour: 13,
})

// Collapsible sections state
const defaultSections = {
  compressionSettings: true
} as const

const sections = reactive<Record<keyof typeof defaultSections, boolean>>({ ...defaultSections })

function toggleSection(section: keyof typeof sections) {
  sections[section] = !sections[section]
}

// Unsaved changes tracking
const unsavedChanges = ref(false)
const lastLoadedSettings = ref({ ...compressionSettings })

// Watch for changes in compressionSettings
watch(compressionSettings, () => {
  unsavedChanges.value = true
}, { deep: true })

// Watch for changes in compressionSettings.enabled and emit the event
watch(() => compressionSettings.enabled, (val) => {
  emit('update:enabled', val)
})

// Function to save compression settings
async function saveCompressionSettings() {
  const payload = [
    { key: 'video_compression.quality', value: compressionSettings.quality, category: 'video_compression', type: 'string' },
    { key: 'video_compression.format', value: compressionSettings.format, category: 'video_compression', type: 'string' },
    { key: 'video_compression.codec', value: compressionSettings.codec, category: 'video_compression', type: 'string' },
    { key: 'video_compression.two_pass', value: String(compressionSettings.twoPass), category: 'video_compression', type: 'boolean' },
    { key: 'video_compression.remove_metadata', value: String(compressionSettings.removeMetadata), category: 'video_compression', type: 'boolean' },

    { key: 'video_compression.create_thumbnail', value: String(compressionSettings.createThumbnail), category: 'video_compression', type: 'boolean' },
    { key: 'video_compression.frame_rate', value: compressionSettings.frameRate, category: 'video_compression', type: 'string' },
    { key: 'video_compression.custom_width', value: compressionSettings.customWidth.toString(), category: 'video_compression', type: 'string' },
    { key: 'video_compression.custom_height', value: compressionSettings.customHeight.toString(), category: 'video_compression', type: 'string' },
    { key: 'video_compression.custom_bitrate', value: compressionSettings.customBitrate.toString(), category: 'video_compression', type: 'string' },
    { key: 'video_compression.worker_count', value: compressionSettings.workerCount, category: 'video_compression', type: 'string' },
    { key: 'video_compression.frame_analysis_mode', value: compressionSettings.frameAnalysisMode, category: 'video_compression', type: 'string' },
    { key: 'video_compression.enabled', value: String(compressionSettings.enabled), category: 'video_compression', type: 'boolean' },
    { key: 'video_compression.start_hour', value: compressionSettings.startHour.toString(), category: 'video_compression', type: 'string' },
    { key: 'video_compression.end_hour', value: compressionSettings.endHour.toString(), category: 'video_compression', type: 'string' }
  ]
  
  try {
    const headers: Record<string, string> = { 'Content-Type': 'application/json' }
    const res = await fetch('/api/admin/settings', {
      method: 'PUT',
      credentials: 'include',
      headers,
      body: JSON.stringify(payload)
    })
    if (res.ok) {
      unsavedChanges.value = false
      lastLoadedSettings.value = { ...compressionSettings }
      
      // بارگذاری تنظیمات زمان‌بندی پس از ذخیره
      await loadCompressionSchedule()
      
      // showToast('تنظیمات با موفقیت ذخیره شد')
    } else {
      const data = await res.json()
      // showToast(data?.error || 'خطا در ذخیره تنظیمات', 'error')
    }
  } catch (err) {
    // showToast('ارتباط با سرور برقرار نشد', 'error')
  }
}

// تابع بارگذاری تنظیمات زمان‌بندی
const loadCompressionSchedule = async () => {
  try {
    const response = await $fetch('/api/settings/video-compression', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    })
    
    if (Array.isArray(response)) {
      const settings = response.reduce((acc: any, setting: any) => {
        acc[setting.key || setting.Key] = setting.value ?? setting.Value
        return acc
      }, {})
      
      // بروزرسانی تنظیمات زمان‌بندی در کامپوننت‌های دیگر
      // این تابع برای ادغام با saveCompressionSettings اضافه شده است
    }
  } catch (error) {
    console.error('خطا در بارگذاری تنظیمات زمان‌بندی:', error)
  }
}

// تابع تست Scheduler
const testScheduler = async () => {
  try {
    interface ApiResponse {
      success?: boolean
      data?: unknown
    }
    const response = await $fetch<ApiResponse>('/api/media/test-scheduler', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      }
    })
    
    if (response.success) {
      console.log('نتایج تست Scheduler:', response.data)
      // نمایش نتایج در console یا toast
    }
  } catch (error) {
    console.error('خطا در تست Scheduler:', error)
  }
}



// Load settings on mount
onMounted(async () => {
  try {
    const headers: Record<string, string> = {}
    const res = await fetch('/api/settings/video-compression', {
      headers,
      credentials: 'include'
    })
    if (!res.ok) return
    const settings = await res.json()
    const map: Record<string, string> = {}
    
    if (Array.isArray(settings)) {
      for (const s of settings) {
        map[s.key || s.Key] = s.value ?? s.Value
      }
    }

    const toBool = (v: any) => {
      if (typeof v === 'boolean') return v
      return String(v).toLowerCase() === 'true' || v === '1'
    }

    // بارگذاری تنظیمات زمان‌بندی
    if (map['video_compression.start_hour']) compressionSettings.startHour = parseInt(map['video_compression.start_hour'])
    if (map['video_compression.end_hour']) compressionSettings.endHour = parseInt(map['video_compression.end_hour'])

    if (map['video_compression.quality']) compressionSettings.quality = map['video_compression.quality']
    if (map['video_compression.format']) compressionSettings.format = map['video_compression.format']
    if (map['video_compression.codec']) compressionSettings.codec = map['video_compression.codec']
    if (map['video_compression.two_pass']) compressionSettings.twoPass = toBool(map['video_compression.two_pass'])
    if (map['video_compression.remove_metadata']) compressionSettings.removeMetadata = toBool(map['video_compression.remove_metadata'])

    if (map['video_compression.create_thumbnail']) compressionSettings.createThumbnail = toBool(map['video_compression.create_thumbnail'])
    if (map['video_compression.frame_rate']) compressionSettings.frameRate = map['video_compression.frame_rate']
    if (map['video_compression.custom_width']) compressionSettings.customWidth = parseInt(map['video_compression.custom_width'])
    if (map['video_compression.custom_height']) compressionSettings.customHeight = parseInt(map['video_compression.custom_height'])
    if (map['video_compression.custom_bitrate']) compressionSettings.customBitrate = parseInt(map['video_compression.custom_bitrate'])
    if (map['video_compression.worker_count']) compressionSettings.workerCount = map['video_compression.worker_count']
    if (map['video_compression.frame_analysis_mode']) compressionSettings.frameAnalysisMode = map['video_compression.frame_analysis_mode']
    if (map['video_compression.enabled']) compressionSettings.enabled = toBool(map['video_compression.enabled'])

    lastLoadedSettings.value = { ...compressionSettings }
    unsavedChanges.value = false
  } catch (e) {
    // silently fail
  }
})

// Expose settings to parent components
defineExpose({
  compressionSettings
})
</script>

<!--
  کامپوننت تنظیمات فشرده‌سازی ویدیو
  - مدیریت تمام تنظیمات مربوط به فشرده‌سازی
  - ذخیره و بارگذاری تنظیمات از سرور
  - ردیابی تغییرات ذخیره نشده
  - توضیحات کامل به فارسی در کد
--> 
