<template>
  <div class="space-y-8">
    <!-- هدر بخش -->
    <div class="flex justify-between items-center">
      <div>
        <h2 class="text-2xl font-bold text-gray-900">چت‌بات و دستیار هوشمند</h2>
        <p class="mt-1 text-sm text-gray-600">پیشنهادات هوشمند، پاسخ‌دهی خودکار و راهنمایی خرید</p>
      </div>
      <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors">
        <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
        </svg>
        شروع چت جدید
      </button>
    </div>

    <!-- وضعیت چت‌بات -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-6">
      <div class="bg-gradient-to-r from-blue-500 to-blue-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">چت‌های فعال</p>
            <p class="text-2xl font-bold">12</p>
            <p class="text-xs opacity-75">در حال گفتگو</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"></path>
            </svg>
          </div>
        </div>
      </div>
      <div class="bg-gradient-to-r from-green-500 to-green-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">پاسخ‌دهی خودکار</p>
            <p class="text-2xl font-bold">98%</p>
            <p class="text-xs opacity-75">دقت پاسخ</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
            </svg>
          </div>
        </div>
      </div>
      <div class="bg-gradient-to-r from-purple-500 to-purple-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">تشخیص قصد</p>
            <p class="text-2xl font-bold">92.5%</p>
            <p class="text-xs opacity-75">دقت</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
            </svg>
          </div>
        </div>
      </div>
      <div class="bg-gradient-to-r from-yellow-500 to-yellow-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">راهنمایی هوشمند</p>
            <p class="text-2xl font-bold">4.8/5</p>
            <p class="text-xs opacity-75">امتیاز کاربران</p>
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
            <p class="text-sm opacity-90">دستیار خرید</p>
            <p class="text-2xl font-bold">1,234</p>
            <p class="text-xs opacity-75">راهنمایی انجام شده</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- چت‌بات -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-6">پنل چت‌بات</h3>
      <div class="h-64 bg-gray-50 rounded-lg flex flex-col justify-end p-6 overflow-y-auto">
        <div class="flex flex-col space-y-2">
          <div class="self-end bg-blue-100 text-blue-900 px-4 py-2 rounded-lg max-w-xs">سلام! چطور می‌تونم کمکتون کنم؟</div>
          <div class="self-start bg-gray-200 text-gray-900 px-4 py-2 rounded-lg max-w-xs">می‌خوام یک گوشی موبایل مناسب بخرم.</div>
          <div class="self-end bg-blue-100 text-blue-900 px-4 py-2 rounded-lg max-w-xs">چه ویژگی‌هایی براتون مهمه؟</div>
        </div>
      </div>
      <div class="flex mt-4">
        <input type="text" class="flex-1 px-4 py-2 border border-gray-300 rounded-l-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="پیام خود را بنویسید...">
        <button class="bg-blue-600 text-white px-6 py-2 rounded-r-lg hover:bg-blue-700 transition-colors">ارسال</button>
      </div>
    </div>

    <!-- قابلیت‌های دستیار -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8 mt-8">
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <h4 class="text-lg font-semibold text-gray-900 mb-4">پاسخ‌دهی خودکار</h4>
        <ul class="list-disc pr-6 space-y-2 text-gray-700 text-sm">
          <li>پاسخ به سوالات متداول کاربران</li>
          <li>پیشنهاد محصولات بر اساس نیاز کاربر</li>
          <li>تشخیص و پاسخ به درخواست‌های خرید</li>
        </ul>
      </div>
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <h4 class="text-lg font-semibold text-gray-900 mb-4">تشخیص قصد و راهنمایی هوشمند</h4>
        <ul class="list-disc pr-6 space-y-2 text-gray-700 text-sm">
          <li>تحلیل پیام کاربر و تشخیص هدف</li>
          <li>ارائه راهنمایی شخصی‌سازی شده</li>
          <li>پیشنهاد بهترین گزینه‌ها بر اساس نیاز</li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// کامپوننت چت‌بات و دستیار هوشمند
import { ref } from 'vue'

const chatMessages = ref([
  { sender: 'bot', text: 'سلام! چطور می‌تونم کمکتون کنم؟' },
  { sender: 'user', text: 'می‌خوام یک گوشی موبایل مناسب بخرم.' },
  { sender: 'bot', text: 'چه ویژگی‌هایی براتون مهمه؟' }
])
</script> 
