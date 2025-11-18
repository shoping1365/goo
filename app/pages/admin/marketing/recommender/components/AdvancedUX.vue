<template>
  <div class="space-y-8">
    <!-- هدر بخش -->
    <div class="flex justify-between items-center">
      <div>
        <h2 class="text-2xl font-bold text-gray-900">تجربه کاربری پیشرفته</h2>
        <p class="mt-1 text-sm text-gray-600">پیشنهادات صوتی، تصویری، واقعیت افزوده و دستیار شخصی</p>
      </div>
      <button class="bg-indigo-600 text-white px-4 py-2 rounded-lg hover:bg-indigo-700 transition-colors">
        <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
        </svg>
        فعال‌سازی قابلیت‌ها
      </button>
    </div>

    <!-- آمار تجربه کاربری پیشرفته -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-6">
      <div class="bg-gradient-to-r from-indigo-500 to-indigo-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">پیشنهادات صوتی</p>
            <p class="text-2xl font-bold">فعال</p>
            <p class="text-xs opacity-75">در حال اجرا</p>
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
            <p class="text-sm opacity-90">پیشنهادات تصویری</p>
            <p class="text-2xl font-bold">فعال</p>
            <p class="text-xs opacity-75">در حال اجرا</p>
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
            <p class="text-sm opacity-90">پیشنهادات AR/VR</p>
            <p class="text-2xl font-bold">فعال</p>
            <p class="text-xs opacity-75">در حال اجرا</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
      </div>
      <div class="bg-gradient-to-r from-yellow-500 to-yellow-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">چت‌بات هوشمند</p>
            <p class="text-2xl font-bold">فعال</p>
            <p class="text-xs opacity-75">در حال اجرا</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
          </div>
        </div>
      </div>
      <div class="bg-gradient-to-r from-red-500 to-red-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">دستیار شخصی</p>
            <p class="text-2xl font-bold">فعال</p>
            <p class="text-xs opacity-75">در حال اجرا</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- جزئیات تجربه کاربری پیشرفته -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-6">جزئیات تجربه کاربری پیشرفته</h3>
      <ul class="list-disc pr-6 space-y-2 text-gray-700 text-sm">
        <li>پیشنهادات صوتی</li>
        <li>پیشنهادات تصویری</li>
        <li>پیشنهادات AR/VR</li>
        <li>چت‌بات هوشمند</li>
        <li>دستیار شخصی</li>
      </ul>
    </div>
  </div>
</template>

<script setup lang="ts">
// کامپوننت تجربه کاربری پیشرفته
</script> 
