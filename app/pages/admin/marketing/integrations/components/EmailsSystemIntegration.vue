<template>
  <div class="bg-white rounded-lg shadow-md p-6">
    <h3 class="text-lg font-semibold text-gray-800 mb-4">یکپارچه‌سازی سیستم سرویس Emails</h3>
    <!-- آمار یکپارچه‌سازی -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
      <div class="bg-blue-50 p-6 rounded-lg">
        <div class="text-sm text-blue-600">سیستم‌های متصل</div>
        <div class="text-2xl font-bold text-blue-800">۳</div>
      </div>
      <div class="bg-green-50 p-6 rounded-lg">
        <div class="text-sm text-green-600">درخواست‌های موفق</div>
        <div class="text-2xl font-bold text-green-800">۱,۲۴۰</div>
      </div>
      <div class="bg-purple-50 p-6 rounded-lg">
        <div class="text-sm text-purple-600">درخواست‌های ناموفق</div>
        <div class="text-2xl font-bold text-purple-800">۲۵</div>
      </div>
      <div class="bg-orange-50 p-6 rounded-lg">
        <div class="text-sm text-orange-600">آخرین اتصال</div>
        <div class="text-2xl font-bold text-orange-800">دیروز</div>
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
// کامپوننت یکپارچه‌سازی سیستم سرویس Emails
// این کامپوننت مدیریت و نظارت بر یکپارچه‌سازی سیستم ایمیل را فراهم می‌کند
</script> 
