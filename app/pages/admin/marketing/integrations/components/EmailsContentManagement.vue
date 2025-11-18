<template>
  <div class="bg-white rounded-lg shadow-md p-6">
    <h3 class="text-lg font-semibold text-gray-800 mb-4">مدیریت محتوا سرویس Emails</h3>
    
    <!-- آمار محتوا -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
      <div class="bg-blue-50 p-6 rounded-lg">
        <div class="text-sm text-blue-600">قالب‌های ایمیل</div>
        <div class="text-2xl font-bold text-blue-800">۲۴</div>
      </div>
      <div class="bg-green-50 p-6 rounded-lg">
        <div class="text-sm text-green-600">ایمیل‌های ارسال شده</div>
        <div class="text-2xl font-bold text-green-800">۱۱,۸۲۰</div>
      </div>
      <div class="bg-purple-50 p-6 rounded-lg">
        <div class="text-sm text-purple-600">نرخ باز شدن</div>
        <div class="text-2xl font-bold text-purple-800">۷۸.۵٪</div>
      </div>
      <div class="bg-orange-50 p-6 rounded-lg">
        <div class="text-sm text-orange-600">نرخ کلیک</div>
        <div class="text-2xl font-bold text-orange-800">۱۲.۳٪</div>
      </div>
    </div>

    <!-- دکمه‌های عملیات -->
    <div class="flex justify-end space-x-3 space-x-reverse">
      <button class="px-4 py-2 bg-green-600 text-white rounded-md hover:bg-green-700 transition-colors">
        ایجاد قالب جدید
      </button>
      <button class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors">
        بروزرسانی
      </button>
    </div>
  </div>
</template>

<script setup>
// کامپوننت مدیریت محتوا سرویس Emails
// این کامپوننت مدیریت محتوای ایمیل را فراهم می‌کند
</script> 
