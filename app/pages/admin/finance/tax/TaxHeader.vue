<template>
  <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
    <div class="flex flex-col lg:flex-row lg:items-center lg:justify-between gap-6">
      <!-- عنوان و توضیحات -->
      <div class="flex-1">
        <div class="flex items-center gap-3 mb-2">
          <div class="w-10 h-10 bg-gradient-to-br from-green-500 to-emerald-600 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 7h6m0 10v-3m-3 3h.01M9 17h.01M9 14h.01M12 14h.01M15 11h.01M12 11h.01M9 11h.01M7 21h10a2 2 0 002-2V5a2 2 0 00-2-2H7a2 2 0 00-2 2v14a2 2 0 002 2z"></path>
            </svg>
          </div>
          <div>
            <h1 class="text-2xl font-bold text-gray-900">مدیریت مالیات</h1>
            <p class="text-gray-600 text-sm">نظارت و مدیریت کامل سیستم مالیاتی</p>
          </div>
        </div>
      </div>

      <!-- دکمه‌های عملیاتی -->
      <div class="flex flex-wrap gap-3">
        <!-- دکمه تنظیمات -->
        <button class="inline-flex items-center gap-2 px-4 py-2 bg-gray-100 hover:bg-gray-200 text-gray-700 rounded-lg transition-colors duration-200">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"></path>
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
          </svg>
          تنظیمات
        </button>

        <!-- دکمه گزارش‌گیری -->
        <button class="inline-flex items-center gap-2 px-4 py-2 bg-blue-100 hover:bg-blue-200 text-blue-700 rounded-lg transition-colors duration-200">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 17v-2m3 2v-4m3 4v-6m2 10H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
          </svg>
          گزارش‌گیری
        </button>

        <!-- دکمه اظهارنامه -->
        <button class="inline-flex items-center gap-2 px-4 py-2 bg-green-100 hover:bg-green-200 text-green-700 rounded-lg transition-colors duration-200">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
          </svg>
          اظهارنامه
        </button>
      </div>
    </div>

    <!-- نوار وضعیت -->
    <div class="mt-6 pt-6 border-t border-gray-200">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <!-- وضعیت سیستم -->
        <div class="flex items-center gap-3">
          <div class="w-2 h-2 bg-green-500 rounded-full"></div>
          <span class="text-sm text-gray-600">سیستم مالیاتی فعال</span>
        </div>

        <!-- آخرین بروزرسانی -->
        <div class="flex items-center gap-3">
          <svg class="w-4 h-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
          </svg>
          <span class="text-sm text-gray-600">آخرین بروزرسانی: {{ lastUpdate }}</span>
        </div>

        <!-- تعداد فاکتورها -->
        <div class="flex items-center gap-3">
          <svg class="w-4 h-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
          </svg>
          <span class="text-sm text-gray-600">{{ totalInvoices }} فاکتور مالیاتی</span>
        </div>

        <!-- مهلت اظهارنامه -->
        <div class="flex items-center gap-3">
          <svg class="w-4 h-4 text-orange-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
          </svg>
          <span class="text-sm text-orange-600">{{ daysUntilDeclaration }} روز تا اظهارنامه</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'

// داده‌های وضعیت
const lastUpdate = ref('')
const totalInvoices = ref(0)
const daysUntilDeclaration = ref(0)

// بارگذاری داده‌های وضعیت
const loadStatusData = async () => {
  try {
    // TODO: دریافت داده‌ها از API
    // const response = await $fetch('/api/admin/tax/status')
    
    // داده‌های نمونه
    lastUpdate.value = 'امروز ساعت ۱۴:۳۰'
    totalInvoices.value = 1247
    daysUntilDeclaration.value = 15
  } catch (error) {
    console.error('خطا در بارگذاری وضعیت مالیات:', error)
  }
}

// بارگذاری اولیه
onMounted(() => {
  loadStatusData()
})
</script>

<!--
  کامپوننت هدر صفحه مالیات
  شامل:
  1. عنوان و آیکون صفحه
  2. دکمه‌های عملیاتی اصلی
  3. نوار وضعیت سیستم
  4. اطلاعات مهم و هشدارها
  طراحی مدرن و کاملاً ریسپانسیو
--> 
