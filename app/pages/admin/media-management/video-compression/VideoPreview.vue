<template>
  <div class="order-6 bg-white rounded-2xl shadow-lg border border-gray-200 mb-6 mx-4 w-full max-w-full min-w-0 overflow-x-hidden">
    <div class="flex items-center justify-between cursor-pointer px-6 py-4 border-b border-blue-300" @click="toggleSection && toggleSection('preview')">
      <span class="text-base font-bold bg-purple-100 text-purple-800 px-4 py-1 rounded-lg ml-auto">پیش‌نمایش</span>
      <span class="text-gray-500 text-2xl">{{ sections?.preview ? '−' : '+' }}</span>
    </div>
    <div v-show="!sections || sections.preview" class="p-10 w-full max-w-full min-w-0">
      <div class="flex flex-col md:flex-row gap-10 items-center justify-center w-full max-w-full min-w-0">
        <!-- ویدیو اصلی (قبل) -->
        <div class="flex-1 flex flex-col items-center min-h-[240px] md:min-h-[320px]">
          <div class="relative group w-full h-72 flex items-center justify-center">
            <video v-if="previewVideoUrl && typeof previewVideoUrl === 'string' && previewVideoUrl.length > 0" :src="previewVideoUrl" controls class="w-full h-full object-contain rounded-lg shadow border border-gray-200 bg-gray-50 cursor-pointer" @click="openPreviewUrl && openPreviewUrl(previewVideoUrl, 'ویدیو قبل')" />
          </div>
          <span class="text-xs text-gray-500 mb-2">ویدیو قبل</span>
          <div v-if="previewUrl" class="mt-2 bg-blue-50 rounded-lg px-3 py-2 text-xs text-gray-700 space-y-1 w-full max-w-xs">
            <div>فرمت ورودی: <span class="font-bold">{{ previewInfo?.originalFormat }}</span></div>
            <div>رزولوشن: <span class="font-bold">{{ previewInfo?.originalResolution }}</span></div>
            <div>حجم ویدیو: <span class="font-bold">{{ formatFileSize && formatFileSize(previewInfo?.originalSize) }}</span></div>
            <div>مدت: <span class="font-bold">{{ previewInfo?.originalDuration }}</span></div>
          </div>
        </div>
        <!-- ویدیو فشرده‌شده (بعد) -->
        <div class="flex-1 flex flex-col items-center min-h-[240px] md:min-h-[320px]">
          <div class="relative group w-full h-72 flex items-center justify-center">
            <video v-if="previewUrl && typeof previewUrl === 'string' && previewUrl.length > 0" :src="previewUrl" controls class="w-full h-full object-contain rounded-lg shadow border border-green-200 bg-green-50 cursor-pointer" @click="openPreviewUrl && openPreviewUrl(previewUrl, 'ویدیو بعد')" />
          </div>
          <span class="text-xs text-gray-500 mb-2">ویدیو بعد</span>
          <div v-if="previewUrl" class="mt-2 bg-purple-50 rounded-lg px-3 py-2 text-xs text-gray-700 space-y-1 w-full max-w-xs">
            <div>فرمت خروجی: <span class="font-bold">{{ previewInfo?.outputFormat }}</span></div>
            <div>رزولوشن: <span class="font-bold">{{ previewInfo?.compressedResolution }}</span></div>
            <div>حجم ویدیو: <span class="font-bold">{{ formatFileSize && formatFileSize(previewInfo?.compressedSize) }}</span></div>
            <div>مدت: <span class="font-bold">{{ previewInfo?.compressedDuration }}</span></div>
          </div>
        </div>
      </div>
      <!-- کنترل‌های پایین: همه در یک ردیف و وسط چین -->
      <div class="flex flex-wrap flex-row gap-6 items-center justify-center my-8 bg-gradient-to-l from-purple-50 via-blue-50 to-pink-50 rounded-xl p-6 w-full max-w-full min-w-0">
        <!-- دکمه انتخاب ویدیو -->
        <input ref="fileInput" type="file" accept="video/*" class="hidden" @change="onPreviewFileChange && onPreviewFileChange($event)" />
        <button class="flex items-center gap-2 font-bold px-4 py-2 rounded-lg shadow transition-colors bg-gradient-to-l from-pink-100 via-purple-100 to-blue-100 text-blue-800 hover:from-pink-200 hover:to-blue-200" @click="fileInput && fileInput.click && fileInput.click()">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path></svg>
          انتخاب ویدیو
        </button>
        <!-- انتخاب فرمت خروجی -->
        <div class="flex items-center gap-2 bg-blue-50 rounded-lg px-3 py-1">
          <label class="text-sm font-medium text-gray-700">فرمت خروجی</label>
          <select v-model="previewFormat" class="px-2 py-1 rounded border border-gray-300 focus:ring-2 focus:ring-blue-400 text-xs w-auto bg-white">
            <option value="mp4">MP4</option>
            <option value="webm">WebM</option>
            <option value="mov">MOV</option>
          </select>
        </div>
        <!-- کیفیت خروجی عددی -->
        <div v-if="!smartCompression" class="flex items-center gap-2 bg-gradient-to-l from-pink-100 via-blue-50 to-purple-50 rounded-lg px-3 py-1">
          <label class="text-sm font-medium text-gray-700">بیت‌ریت خروجی</label>
          <input v-model="previewBitrate" type="number" min="500" max="10000" class="w-20 px-2 py-1 rounded border border-gray-300 text-center text-xs focus:ring-2 focus:ring-blue-400 bg-white" />
          <span class="text-xs text-gray-500">kbps</span>
        </div>
        <!-- سوییچ هوشمند -->
        <div class="flex items-center gap-2 bg-purple-50 rounded-lg px-3 py-1">
          <label class="flex items-center">
            <input
              v-model="smartCompression"
              type="checkbox"
              class="w-4 h-4 text-blue-600 bg-white border-gray-300 rounded-full focus:ring-blue-500"
            >
            <span class="mr-1 text-sm text-blue-700 cursor-pointer">فشرده‌سازی هوشمند</span>
          </label>
        </div>
        <!-- دکمه پیش‌نمایش -->
        <button :disabled="isPreviewLoading" class="flex items-center gap-2 font-bold px-6 py-2 rounded-lg shadow transition-colors bg-gradient-to-l from-blue-400 to-purple-400 text-white hover:from-blue-500 hover:to-purple-500 disabled:bg-gray-400 disabled:cursor-not-allowed" @click="onPreviewButtonClick && onPreviewButtonClick()">
          <span v-if="isPreviewLoading">در حال تولید پیش‌نمایش...</span>
          <span v-else>پیش‌نمایش</span>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'

// Preview card placeholders
const previewVideoUrl = ref<string>('')
const previewUrl = ref<string>('')
const previewInfo = reactive({
  originalFormat: '',
  originalResolution: '',
  originalSize: 0,
  originalDuration: '',
  outputFormat: '',
  compressedResolution: '',
  compressedSize: 0,
  compressedDuration: ''
})
const previewFormat = ref('mp4')
const previewBitrate = ref(2000)
const smartCompression = ref(false)
const isPreviewLoading = ref(false)
const fileInput = ref<HTMLInputElement | null>(null)

// Collapsible sections state
const defaultSections = {
  preview: true
} as const

const sections = reactive<Record<keyof typeof defaultSections, boolean>>({ ...defaultSections })

function toggleSection(section: keyof typeof sections) {
  sections[section] = !sections[section]
}

// تابع انتخاب فایل برای پیش‌نمایش
function onPreviewFileChange(e: Event) {
  const target = e.target as HTMLInputElement
  if (!target.files || target.files.length === 0) return
  
  const file = target.files[0]
  if (!file.type.startsWith('video/')) {
    alert('لطفاً یک فایل ویدیو انتخاب کنید')
    return
  }
  
  // ایجاد URL برای پیش‌نمایش
  const url = URL.createObjectURL(file)
  previewVideoUrl.value = url
  
  // استخراج اطلاعات فایل
  previewInfo.originalFormat = file.name.split('.').pop()?.toUpperCase() || ''
  previewInfo.originalSize = file.size
  
  // پاک کردن پیش‌نمایش قبلی
  if (previewUrl.value) {
    URL.revokeObjectURL(previewUrl.value)
    previewUrl.value = ''
  }
  
  // ریست کردن اطلاعات خروجی
  previewInfo.outputFormat = ''
  previewInfo.compressedResolution = ''
  previewInfo.compressedSize = 0
  previewInfo.compressedDuration = ''
  
  // استخراج اطلاعات ویدیو با استفاده از video element
  // Validate URL before using it
  let safeUrl = url
  try {
    // If it's a relative URL, it's safe
    if (url.startsWith('http://') || url.startsWith('https://')) {
      const urlObj = new URL(url)
      // Only allow http and https protocols
      if (!['http:', 'https:'].includes(urlObj.protocol)) {
        console.error('Invalid URL protocol:', urlObj.protocol)
        return
      }
      safeUrl = urlObj.toString()
    }
  } catch (e) {
    console.error('Invalid URL:', url, e)
    return
  }
  
  const video = document.createElement('video')
  video.onloadedmetadata = () => {
    previewInfo.originalResolution = `${video.videoWidth}x${video.videoHeight}`
    previewInfo.originalDuration = formatTime(video.duration * 1000)
  }
  video.src = safeUrl
}

// تابع تولید پیش‌نمایش
async function onPreviewButtonClick() {
  if (!previewVideoUrl.value) {
    alert('لطفاً ابتدا یک ویدیو انتخاب کنید')
    return
  }
  
  isPreviewLoading.value = true
  
  try {
    // اینجا باید API call برای تولید پیش‌نمایش انجام شود
    // فعلاً فقط یک نمونه نمایش داده می‌شود
    await new Promise(resolve => setTimeout(resolve, 2000))
    
    // شبیه‌سازی نتیجه فشرده‌سازی
    previewInfo.outputFormat = previewFormat.value.toUpperCase()
    previewInfo.compressedResolution = '1920x1080'
    previewInfo.compressedSize = Math.floor(previewInfo.originalSize * 0.6) // 40% کاهش
    previewInfo.compressedDuration = previewInfo.originalDuration
    
    // ایجاد URL نمونه برای ویدیو فشرده‌شده
    previewUrl.value = previewVideoUrl.value // در عمل باید URL واقعی باشد
    
  } catch (error) {
    console.error('خطا در تولید پیش‌نمایش:', error)
    alert('خطا در تولید پیش‌نمایش')
  } finally {
    isPreviewLoading.value = false
  }
}

// تابع باز کردن URL در تب جدید
function openPreviewUrl(url: string, _title: string) {
  window.open(url, '_blank', 'noopener,noreferrer')
}

// تابع فرمت کردن اندازه فایل
const formatFileSize = (bytes: number): string => {
  if (bytes === 0) return '0 B'
  const k = 1024
  const sizes = ['B', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

// تابع فرمت کردن زمان
const formatTime = (milliseconds: number): string => {
  const seconds = Math.floor(milliseconds / 1000)
  const minutes = Math.floor(seconds / 60)
  const hours = Math.floor(minutes / 60)
  
  if (hours > 0) {
    return `${hours}:${(minutes % 60).toString().padStart(2, '0')}:${(seconds % 60).toString().padStart(2, '0')}`
  } else {
    return `${minutes}:${(seconds % 60).toString().padStart(2, '0')}`
  }
}
</script>

<!--
  کامپوننت پیش‌نمایش ویدیو
  - نمایش ویدیو قبل و بعد از فشرده‌سازی
  - انتخاب فایل ویدیو برای پیش‌نمایش
  - تنظیمات پیش‌نمایش (فرمت، بیت‌ریت، هوشمند)
  - تولید پیش‌نمایش فشرده‌سازی
  - توضیحات کامل به فارسی در کد
--> 
