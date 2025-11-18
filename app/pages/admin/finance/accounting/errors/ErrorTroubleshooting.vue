<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">راهنمای رفع خطا</h4>
        <p class="text-sm text-gray-600 mt-1">راهنمای گام‌به‌گام برای رفع خطاهای رایج اتصال حسابداری</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          @click="openHelpCenter"
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 10h.01M12 10h.01M16 10h.01M9 16h6m2 4H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
          </svg>
          مرکز راهنما
        </button>
      </div>
    </div>

    <!-- لیست خطاهای رایج و راه‌حل‌ها -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">خطاهای رایج و راه‌حل‌ها</h5>
      <div class="space-y-6">
        <div v-for="item in troubleshootingList" :key="item.id" class="border-b border-gray-100 pb-4">
          <div class="flex items-center space-x-2 space-x-reverse mb-2">
            <svg class="w-5 h-5 text-blue-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 10h.01M12 10h.01M16 10h.01M9 16h6m2 4H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
            </svg>
            <span class="font-semibold text-gray-900">{{ item.title }}</span>
          </div>
          <div class="text-sm text-gray-700 mb-2">{{ item.description }}</div>
          <ul class="list-disc pr-6 text-sm text-gray-600 space-y-1">
            <li v-for="(step, idx) in item.steps" :key="idx">{{ step }}</li>
          </ul>
        </div>
      </div>
    </div>

    <!-- جستجوی راهنما -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">جستجوی راهنما</h5>
      <input
        v-model="searchQuery"
        type="text"
        class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
        placeholder="عبارت یا کد خطا را وارد کنید..."
      />
      <div v-if="filteredList.length > 0" class="mt-4 space-y-4">
        <div v-for="item in filteredList" :key="item.id" class="border-b border-gray-100 pb-2">
          <div class="font-semibold text-blue-700">{{ item.title }}</div>
          <div class="text-xs text-gray-600">{{ item.description }}</div>
        </div>
      </div>
      <div v-else class="mt-4 text-sm text-gray-500">موردی یافت نشد.</div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

// لیست خطاهای رایج و راه‌حل‌ها
const troubleshootingList = ref([
  {
    id: 1,
    title: 'Timeout در اتصال به نرم‌افزار حسابداری',
    description: 'اتصال به نرم‌افزار حسابداری پس از مدت زمان مشخص قطع می‌شود.',
    steps: [
      'بررسی وضعیت شبکه و اینترنت سرور حسابداری',
      'اطمینان از فعال بودن سرویس حسابداری',
      'بررسی تنظیمات فایروال و پورت‌ها',
      'افزایش مقدار timeout در تنظیمات اتصال'
    ]
  },
  {
    id: 2,
    title: 'کلید API نامعتبر یا منقضی شده',
    description: 'در هنگام احراز هویت با خطای کلید API مواجه می‌شوید.',
    steps: [
      'بررسی اعتبار کلید API در پنل حسابداری',
      'در صورت نیاز، ایجاد کلید جدید و جایگزینی آن',
      'بررسی سطح دسترسی کلید API',
      'بررسی تاریخ انقضای کلید'
    ]
  },
  {
    id: 3,
    title: 'خطا در همگام‌سازی داده‌ها',
    description: 'داده‌ها به درستی بین سیستم‌ها همگام نمی‌شوند.',
    steps: [
      'بررسی فرمت و ساختار داده‌های ارسالی',
      'بررسی لاگ خطاهای همگام‌سازی',
      'اطمینان از به‌روزرسانی نسخه نرم‌افزارها',
      'بررسی محدودیت‌های API مقصد'
    ]
  },
  {
    id: 4,
    title: 'کندی یا کاهش عملکرد اتصال',
    description: 'عملکرد اتصال حسابداری کند شده یا با تاخیر همراه است.',
    steps: [
      'بررسی مصرف منابع سرور (CPU/RAM)',
      'بررسی وضعیت شبکه و پهنای باند',
      'بهینه‌سازی کوئری‌ها و درخواست‌ها',
      'استفاده از کش مناسب در صورت امکان'
    ]
  },
  {
    id: 5,
    title: 'خطاهای امنیتی و دسترسی',
    description: 'دسترسی به برخی بخش‌ها یا داده‌ها با خطا مواجه می‌شود.',
    steps: [
      'بررسی سطح دسترسی کاربران و کلیدهای API',
      'بررسی تنظیمات مجوزها در نرم‌افزار حسابداری',
      'اطمینان از فعال بودن رمزگذاری ارتباط',
      'بررسی لاگ‌های امنیتی سیستم'
    ]
  }
])

// جستجو
const searchQuery = ref('')
const filteredList = computed(() => {
  if (!searchQuery.value) return troubleshootingList.value
  return troubleshootingList.value.filter(item =>
    item.title.includes(searchQuery.value) ||
    item.description.includes(searchQuery.value)
  )
})

// مرکز راهنما
const openHelpCenter = () => {
  // TODO: باز کردن مرکز راهنما
  window.open('https://help.example.com', '_blank')
}
</script>

<!--
  کامپوننت راهنمای رفع خطا
  شامل:
  1. لیست خطاهای رایج و راه‌حل‌ها
  2. جستجوی راهنما
  3. دکمه مرکز راهنما
  4. طراحی مدرن و کاملاً ریسپانسیو
--> 
