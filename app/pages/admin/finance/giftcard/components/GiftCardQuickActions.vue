<template>
  <div class="bg-white rounded-2xl border border-gray-200/50 shadow-sm p-6 mb-8">
    <div class="flex items-center justify-between mb-6">
      <h2 class="text-xl font-bold text-gray-900">عملیات سریع</h2>
      <div class="flex items-center space-x-2 space-x-reverse">
        <span class="text-sm text-gray-500">آخرین بروزرسانی:</span>
        <span class="text-sm font-medium text-gray-900">{{ formatTime(new Date()) }}</span>
      </div>
    </div>
    
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <!-- ایجاد کارت جدید -->
      <button 
        class="group relative bg-gradient-to-r from-purple-500 to-pink-600 rounded-xl p-6 text-white hover:from-purple-600 hover:to-pink-700 transition-all duration-200 shadow-lg hover:shadow-xl transform hover:-translate-y-1"
        @click="$emit('create-card')"
      >
        <div class="flex items-center justify-between">
          <div>
            <h3 class="text-lg font-semibold mb-2">ایجاد کارت جدید</h3>
            <p class="text-purple-100 text-sm">ایجاد گیفت کارت منحصر به فرد</p>
          </div>
          <div class="w-12 h-12 bg-white/20 rounded-lg flex items-center justify-center group-hover:scale-110 transition-transform">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path>
            </svg>
          </div>
        </div>
      </button>

      <!-- ایجاد گروهی -->
      <button 
        class="group relative bg-gradient-to-r from-blue-500 to-cyan-600 rounded-xl p-6 text-white hover:from-blue-600 hover:to-cyan-700 transition-all duration-200 shadow-lg hover:shadow-xl transform hover:-translate-y-1"
        @click="$emit('bulk-create')"
      >
        <div class="flex items-center justify-between">
          <div>
            <h3 class="text-lg font-semibold mb-2">ایجاد گروهی</h3>
            <p class="text-blue-100 text-sm">ایجاد چندین کارت همزمان</p>
          </div>
          <div class="w-12 h-12 bg-white/20 rounded-lg flex items-center justify-center group-hover:scale-110 transition-transform">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              
            </svg>
          </div>
        </div>
      </button>

      <!-- خروجی گزارش -->
      <button 
        class="group relative bg-gradient-to-r from-green-500 to-emerald-600 rounded-xl p-6 text-white hover:from-green-600 hover:to-emerald-700 transition-all duration-200 shadow-lg hover:shadow-xl transform hover:-translate-y-1"
        @click="$emit('export-data')"
      >
        <div class="flex items-center justify-between">
          <div>
            <h3 class="text-lg font-semibold mb-2">خروجی گزارش</h3>
            <p class="text-green-100 text-sm">دانلود گزارش کامل</p>
          </div>
          <div class="w-12 h-12 bg-white/20 rounded-lg flex items-center justify-center group-hover:scale-110 transition-transform">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
            </svg>
          </div>
        </div>
      </button>

      <!-- تنظیمات -->
      <button 
        class="group relative bg-gradient-to-r from-orange-500 to-red-600 rounded-xl p-6 text-white hover:from-orange-600 hover:to-red-700 transition-all duration-200 shadow-lg hover:shadow-xl transform hover:-translate-y-1"
        @click="showSettings = true"
      >
        <div class="flex items-center justify-between">
          <div>
            <h3 class="text-lg font-semibold mb-2">تنظیمات</h3>
            <p class="text-orange-100 text-sm">مدیریت تنظیمات سیستم</p>
          </div>
          <div class="w-12 h-12 bg-white/20 rounded-lg flex items-center justify-center group-hover:scale-110 transition-transform">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"></path>
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
            </svg>
          </div>
        </div>
      </button>
    </div>

    <!-- عملیات اضافی -->
    <div class="mt-6 pt-6 border-t border-gray-200">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">عملیات اضافی</h3>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <!-- بررسی اعتبار -->
        <button class="flex items-center justify-between p-6 bg-gray-50 hover:bg-gray-100 rounded-lg transition-colors">
          <div class="flex items-center space-x-3 space-x-reverse">
            <div class="w-10 h-10 bg-blue-100 rounded-lg flex items-center justify-center">
              <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
            <div>
              <h4 class="font-medium text-gray-900">بررسی اعتبار</h4>
              <p class="text-sm text-gray-600">بررسی وضعیت کارت</p>
            </div>
          </div>
          <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
          </svg>
        </button>

        <!-- مسدود کردن کارت -->
        <button class="flex items-center justify-between p-6 bg-gray-50 hover:bg-gray-100 rounded-lg transition-colors">
          <div class="flex items-center space-x-3 space-x-reverse">
            <div class="w-10 h-10 bg-red-100 rounded-lg flex items-center justify-center">
              <svg class="w-5 h-5 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728L5.636 5.636m12.728 12.728L18.364 5.636M5.636 18.364l12.728-12.728"></path>
              </svg>
            </div>
            <div>
              <h4 class="font-medium text-gray-900">مسدود کردن</h4>
              <p class="text-sm text-gray-600">مسدود کردن کارت</p>
            </div>
          </div>
          <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
          </svg>
        </button>

        <!-- تمدید اعتبار -->
        <button class="flex items-center justify-between p-6 bg-gray-50 hover:bg-gray-100 rounded-lg transition-colors">
          <div class="flex items-center space-x-3 space-x-reverse">
            <div class="w-10 h-10 bg-green-100 rounded-lg flex items-center justify-center">
              <svg class="w-5 h-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
              </svg>
            </div>
            <div>
              <h4 class="font-medium text-gray-900">تمدید اعتبار</h4>
              <p class="text-sm text-gray-600">تمدید تاریخ انقضا</p>
            </div>
          </div>
          <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
          </svg>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// تعریف events
const _emit = defineEmits<{
  'create-card': []
  'bulk-create': []
  'export-data': []
}>()

// Reactive data
const showSettings = ref(false)

// تابع فرمت کردن زمان
const formatTime = (date: Date): string => {
  return date.toLocaleTimeString('fa-IR', { 
    hour: '2-digit', 
    minute: '2-digit',
    second: '2-digit'
  })
}
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 
