<template>
  <div class="bg-white rounded-lg shadow-md p-6">
    <h3 class="text-lg font-semibold text-gray-800 mb-4">مدیریت امنیت سرویس Torb</h3>
    
    <!-- آمار امنیت -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
      <div class="bg-blue-50 p-6 rounded-lg">
        <div class="text-sm text-blue-600">نمره امنیت</div>
        <div class="text-2xl font-bold text-blue-800">۹۵</div>
      </div>
      <div class="bg-green-50 p-6 rounded-lg">
        <div class="text-sm text-green-600">تهدیدات مسدود شده</div>
        <div class="text-2xl font-bold text-green-800">۲,۴۵۰</div>
      </div>
      <div class="bg-purple-50 p-6 rounded-lg">
        <div class="text-sm text-purple-600">آخرین اسکن</div>
        <div class="text-2xl font-bold text-purple-800">۲ ساعت پیش</div>
      </div>
      <div class="bg-orange-50 p-6 rounded-lg">
        <div class="text-sm text-orange-600">وضعیت سیستم</div>
        <div class="text-2xl font-bold text-orange-800">ایمن</div>
      </div>
    </div>

    <!-- نمودار امنیت -->
    <div class="mb-6">
      <h4 class="text-md font-medium text-gray-700 mb-3">نمودار امنیت</h4>
      <div class="bg-gray-50 p-6 rounded-lg h-64 flex items-center justify-center">
        <div class="text-center">
          <div class="text-gray-500 mb-2">نمودار امنیت</div>
          <div class="text-sm text-gray-400">نمایش وضعیت امنیت سیستم</div>
        </div>
      </div>
    </div>

    <!-- دکمه‌های عملیات -->
    <div class="flex justify-end space-x-3 space-x-reverse">
      <button class="px-4 py-2 bg-green-600 text-white rounded-md hover:bg-green-700 transition-colors">
        صادر کردن گزارش
      </button>
      <button class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors">
        بروزرسانی
      </button>
    </div>
  </div>
</template>

<script setup>
// کامپوننت مدیریت امنیت سرویس Torb
// این کامپوننت مدیریت و نظارت بر امنیت را فراهم می‌کند
</script> 
