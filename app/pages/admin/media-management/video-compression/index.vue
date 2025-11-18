<template>
  <div class="min-h-screen w-full max-w-full min-w-0 overflow-x-hidden">
    <!-- Header -->
    <div class="bg-white shadow-sm border-b border-gray-200 sticky top-0 z-10 w-full max-w-full min-w-0 overflow-x-hidden">
      <div class="px-2 md:px-3 py-2 w-full max-w-full min-w-0">
        <div class="flex items-center justify-between w-full max-w-full min-w-0">
          <div>
            <h1 class="text-lg font-bold text-gray-900">تنظیمات و فشرده‌سازی ویدیو</h1>
            <p class="text-xs text-gray-500 mt-0.5">بهینه‌سازی و کاهش حجم ویدیوها</p>
          </div>
          <div class="flex items-center gap-3">
            <NuxtLink 
              to="/admin/media"
              class="bg-gradient-to-l from-pink-200 via-purple-200 to-blue-200 text-gray-800 px-4 py-2 rounded-lg shadow-md hover:from-pink-300 hover:to-blue-300 hover:shadow-lg transition-all flex items-center gap-2 font-bold border border-blue-100"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path>
              </svg>
              بازگشت به کتابخانه
            </NuxtLink>
          </div>
        </div>
      </div>
    </div>

    <!-- Main Content -->
    <div class="px-0 max-w-full w-full flex flex-col min-w-0 overflow-x-hidden">
      <!-- کارت تنظیمات فشرده‌سازی ویدیو -->
      <VideoCompressionSettings @update:enabled="handleSettingsChange" />
      
      <!-- کارت صف فشرده‌سازی -->
      <VideoCompressionQueue />

      <!-- کارت پیغام های خطای صف فشرده سازی -->
      <div class="order-5 bg-white rounded-2xl shadow-lg border border-red-200 mb-6 mx-4 w-full max-w-full min-w-0 overflow-x-hidden">
        <div class="flex items-center justify-between px-6 py-4 border-b border-red-300">
          <span class="text-base font-bold bg-red-100 text-red-700 px-4 py-1 rounded-lg">پیغام های خطای صف فشرده سازی</span>
          <button 
            class="bg-red-500 hover:bg-red-600 text-white px-4 py-2 rounded-lg text-sm font-medium transition-colors duration-200 flex items-center gap-2"
            @click="clearErrorMessages"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
            </svg>
            پاکسازی پیام‌های خطا
          </button>
        </div>
        <div class="p-6 text-sm text-red-700">
          <!-- اینجا پیغام‌های خطا نمایش داده می‌شود -->
          <div v-if="errorMessages.length === 0" class="text-gray-500">
            هیچ خطایی وجود ندارد.
          </div>
          <div v-else class="space-y-3">
            <div v-for="error in errorMessages" :key="error.id" class="p-3 bg-red-50 border border-red-200 rounded-lg">
              <div class="flex justify-between items-start">
                <div class="flex-1">
                  <div class="font-medium text-red-800">{{ error.error_message }}</div>
                  <div class="text-xs text-red-600 mt-1">
                    Media ID: {{ error.media_id }} | Created: {{ formatDate(error.created_at) }}
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      
      <!-- کارت نتایج فشرده‌سازی -->
      <VideoCompressionResults />
      
      <!-- کارت پیش‌نمایش ویدیو -->
      <VideoPreview />
      
      <!-- کارت مانیتورینگ سیستم -->
      <VideoMonitoring :compression-enabled="compressionEnabled" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { defineAsyncComponent } from 'vue'

// Import کامپوننت‌ها به صورت async
const VideoMonitoring = defineAsyncComponent(() => import('./VideoMonitoring.vue'))
const VideoCompressionSettings = defineAsyncComponent(() => import('./VideoCompressionSettings.vue'))
const VideoPreview = defineAsyncComponent(() => import('./VideoPreview.vue'))
const VideoCompressionQueue = defineAsyncComponent(() => import('./VideoCompressionQueue.vue'))
const VideoCompressionResults = defineAsyncComponent(() => import('./VideoCompressionResults.vue'))

// وضعیت فعال بودن فشرده‌سازی
const compressionEnabled = ref(true)

// پیام‌های خطا
const errorMessages = ref([])

// تابعی برای بروزرسانی وضعیت از تنظیمات
function handleSettingsChange(enabled: boolean) {
  compressionEnabled.value = enabled
}

// تابع بارگذاری پیام‌های خطا
async function loadErrorMessages() {
  try {
    interface ApiResponse {
      success?: boolean
      data?: Array<{ error_message?: string }>
    }
    const response = await $fetch<ApiResponse>('/api/media/compression-jobs?job_type=video_compression&status=error', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    })
    
    if (response.success && Array.isArray(response.data)) {
      errorMessages.value = response.data.filter(job => job.error_message && job.error_message.trim() !== '')
    }
  } catch (error) {
    console.error('خطا در بارگذاری پیام‌های خطا:', error)
  }
}

// تابع پاکسازی پیام‌های خطا
async function clearErrorMessages() {
  try {
    await $fetch('/api/media/compression-jobs/clear-errors', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      }
    })
    
    // بارگذاری مجدد پیام‌های خطا
    await loadErrorMessages()
  } catch (error) {
    console.error('خطا در پاکسازی پیام‌های خطا:', error)
  }
}

// تابع فرمت تاریخ
function formatDate(dateString: string) {
  if (!dateString) return ''
  const date = new Date(dateString)
  return date.toLocaleString('fa-IR')
}

// بارگذاری پیام‌های خطا در mount
onMounted(() => {
  loadErrorMessages()
})

declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
declare const useHead: (head: { title?: string }) => void

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

useHead({
  title: 'فشرده‌سازی ویدیو - پنل مدیریت'
})
</script>

<!--
  صفحه اصلی فشرده‌سازی ویدیو
  - شامل کامپوننت‌های تنظیمات، پیش‌نمایش، صف و نتایج
  - مدیریت layout و meta tags
  - توضیحات کامل به فارسی در کد
--> 
