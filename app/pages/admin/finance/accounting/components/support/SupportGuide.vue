<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">راهنمای اتصال</h4>
        <p class="text-sm text-gray-600 mt-1">مراحل گام‌به‌گام برای اتصال به نرم‌افزارهای حسابداری</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          @click="downloadGuide"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
          </svg>
          دانلود راهنما
        </button>
      </div>
    </div>

    <!-- مراحل اتصال -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">مراحل اتصال</h5>
      <div class="space-y-4">
        <div v-for="(step, index) in connectionSteps" :key="index" class="flex items-start space-x-4 space-x-reverse">
          <div class="flex-shrink-0 w-8 h-8 bg-blue-500 text-white rounded-full flex items-center justify-center text-sm font-medium">
            {{ index + 1 }}
          </div>
          <div class="flex-1">
            <h6 class="font-medium text-gray-900">{{ step.title }}</h6>
            <p class="text-sm text-gray-600 mt-1">{{ step.description }}</p>
            <ul v-if="step.details" class="list-disc pr-6 text-sm text-gray-600 mt-2 space-y-1">
              <li v-for="(detail, idx) in step.details" :key="idx">{{ detail }}</li>
            </ul>
          </div>
        </div>
      </div>
    </div>

    <!-- نکات مهم -->
    <div class="bg-yellow-50 rounded-lg p-6">
      <h5 class="text-md font-medium text-yellow-900 mb-4">نکات مهم</h5>
      <ul class="list-disc pr-6 text-sm text-yellow-800 space-y-2">
        <li>اطمینان از فعال بودن سرویس نرم‌افزار حسابداری</li>
        <li>بررسی تنظیمات فایروال و پورت‌های مورد نیاز</li>
        <li>استفاده از کلیدهای API معتبر و به‌روز</li>
        <li>بررسی سطح دسترسی کاربران در نرم‌افزار حسابداری</li>
        <li>اطمینان از به‌روزرسانی نسخه نرم‌افزارها</li>
      </ul>
    </div>

    <!-- دسترسی سریع -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h5 class="text-md font-medium text-gray-900 mb-4">دسترسی سریع</h5>
        <div class="space-y-3">
          <button
            class="w-full text-right p-3 border border-gray-200 rounded-lg hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-blue-500"
            @click="openVideoGuide"
          >
            <div class="flex items-center justify-between">
              <span class="text-sm font-medium text-gray-900">ویدیوی آموزشی</span>
              <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.828 14.828a4 4 0 01-5.656 0M9 10h1m4 0h1m-6 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </div>
          </button>
          <button
            class="w-full text-right p-3 border border-gray-200 rounded-lg hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-blue-500"
            @click="openFAQ"
          >
            <div class="flex items-center justify-between">
              <span class="text-sm font-medium text-gray-900">سوالات متداول</span>
              <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </div>
          </button>
          <button
            class="w-full text-right p-3 border border-gray-200 rounded-lg hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-blue-500"
            @click="contactSupport"
          >
            <div class="flex items-center justify-between">
              <span class="text-sm font-medium text-gray-900">تماس با پشتیبانی</span>
              <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z" />
              </svg>
            </div>
          </button>
        </div>
      </div>

      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h5 class="text-md font-medium text-gray-900 mb-4">مستندات کامل</h5>
        <div class="space-y-3">
          <a
            href="#"
            class="block p-3 border border-gray-200 rounded-lg hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            <div class="flex items-center justify-between">
              <span class="text-sm font-medium text-gray-900">راهنمای کامل اتصال</span>
              <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
              </svg>
            </div>
          </a>
          <a
            href="#"
            class="block p-3 border border-gray-200 rounded-lg hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            <div class="flex items-center justify-between">
              <span class="text-sm font-medium text-gray-900">مستندات API</span>
              <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
              </svg>
            </div>
          </a>
          <a
            href="#"
            class="block p-3 border border-gray-200 rounded-lg hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            <div class="flex items-center justify-between">
              <span class="text-sm font-medium text-gray-900">نمونه کدها</span>
              <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
              </svg>
            </div>
          </a>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// مراحل اتصال
const connectionSteps = ref([
  {
    title: 'بررسی پیش‌نیازها',
    description: 'اطمینان از نصب و فعال بودن نرم‌افزار حسابداری',
    details: [
      'نصب آخرین نسخه نرم‌افزار حسابداری',
      'فعال کردن سرویس‌های مورد نیاز',
      'بررسی دسترسی‌های شبکه'
    ]
  },
  {
    title: 'دریافت کلید API',
    description: 'دریافت کلید API از پنل مدیریت نرم‌افزار حسابداری',
    details: [
      'ورود به پنل مدیریت',
      'ایجاد کلید API جدید',
      'کپی کردن کلید در جای امن'
    ]
  },
  {
    title: 'تنظیم اتصال',
    description: 'پیکربندی اتصال در سیستم',
    details: [
      'وارد کردن آدرس سرور',
      'تنظیم کلید API',
      'تست اتصال اولیه'
    ]
  },
  {
    title: 'تست و تایید',
    description: 'بررسی عملکرد اتصال و همگام‌سازی',
    details: [
      'ارسال درخواست تست',
      'بررسی پاسخ‌های دریافتی',
      'تایید نهایی اتصال'
    ]
  }
])

// دانلود راهنما
const downloadGuide = () => {
  // TODO: دانلود فایل راهنما

}

// باز کردن ویدیوی آموزشی
const openVideoGuide = () => {
  // TODO: باز کردن ویدیوی آموزشی

}

// باز کردن سوالات متداول
const openFAQ = () => {
  // TODO: باز کردن سوالات متداول

}

// تماس با پشتیبانی
const contactSupport = () => {
  // TODO: تماس با پشتیبانی

}
</script>

<!--
  کامپوننت راهنمای اتصال
  شامل:
  1. مراحل گام‌به‌گام اتصال
  2. نکات مهم
  3. دسترسی سریع
  4. مستندات کامل
  5. طراحی مدرن و کاملاً ریسپانسیو
--> 
