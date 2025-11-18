<template>
  <div class="space-y-8">
    <!-- هدر بخش -->
    <div class="flex justify-between items-center">
      <div>
        <h2 class="text-2xl font-bold text-gray-900">اخلاق و شفافیت</h2>
        <p class="mt-1 text-sm text-gray-600">توضیح‌پذیری، عدالت، حفاظت از حریم خصوصی و ردپای تصمیمات</p>
      </div>
      <button class="bg-yellow-600 text-white px-4 py-2 rounded-lg hover:bg-yellow-700 transition-colors">
        <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
        </svg>
        مشاهده گزارش
      </button>
    </div>

    <!-- آمار اخلاق و شفافیت -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-6">
      <div class="bg-gradient-to-r from-yellow-500 to-yellow-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">توضیح‌پذیری</p>
            <p class="text-2xl font-bold">98%</p>
            <p class="text-xs opacity-75">تصمیمات</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
            </svg>
          </div>
        </div>
      </div>
      <div class="bg-gradient-to-r from-blue-500 to-blue-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">تشخیص سوگیری</p>
            <p class="text-2xl font-bold">1.2%</p>
            <p class="text-xs opacity-75">موارد شناسایی شده</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"></path>
            </svg>
          </div>
        </div>
      </div>
      <div class="bg-gradient-to-r from-green-500 to-green-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">معیارهای عدالت</p>
            <p class="text-2xl font-bold">99%</p>
            <p class="text-xs opacity-75">پوشش داده شده</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
      </div>
      <div class="bg-gradient-to-r from-purple-500 to-purple-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">حفاظت از حریم خصوصی</p>
            <p class="text-2xl font-bold">100%</p>
            <p class="text-xs opacity-75">اطلاعات رمزنگاری شده</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z"></path>
            </svg>
          </div>
        </div>
      </div>
      <div class="bg-gradient-to-r from-red-500 to-red-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">Audit Trail</p>
            <p class="text-2xl font-bold">فعال</p>
            <p class="text-xs opacity-75">ردپای تصمیمات</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- جزئیات اخلاق و شفافیت -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-6">جزئیات اخلاق و شفافیت</h3>
      <ul class="list-disc pr-6 space-y-2 text-gray-700 text-sm">
        <li>توضیح‌پذیری تصمیمات (Explainable AI)</li>
        <li>تشخیص سوگیری (Bias Detection)</li>
        <li>معیارهای عدالت (Fairness Metrics)</li>
        <li>حفاظت از حریم خصوصی (Privacy Protection)</li>
        <li>ردپای تصمیمات (Audit Trail)</li>
      </ul>
    </div>
  </div>
</template>

<script setup lang="ts">
// کامپوننت اخلاق و شفافیت
</script> 
