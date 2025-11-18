<template>
  <div class="bg-white rounded-lg shadow-md p-6">
    <h3 class="text-lg font-semibold text-gray-800 mb-4">کرالر سرویس Emails</h3>
    
    <!-- آمار کرالر -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
      <div class="bg-blue-50 p-6 rounded-lg">
        <div class="text-sm text-blue-600">ایمیل‌های کرال شده</div>
        <div class="text-2xl font-bold text-blue-800">۲,۴۵۰</div>
      </div>
      <div class="bg-green-50 p-6 rounded-lg">
        <div class="text-sm text-green-600">وضعیت کرالر</div>
        <div class="text-2xl font-bold text-green-800">فعال</div>
      </div>
      <div class="bg-purple-50 p-6 rounded-lg">
        <div class="text-sm text-purple-600">آخرین کرال</div>
        <div class="text-2xl font-bold text-purple-800">۲ ساعت پیش</div>
      </div>
      <div class="bg-orange-50 p-6 rounded-lg">
        <div class="text-sm text-orange-600">سرعت کرال</div>
        <div class="text-2xl font-bold text-orange-800">۱۰۰/دقیقه</div>
      </div>
    </div>

    <!-- دکمه‌های عملیات -->
    <div class="flex justify-end space-x-3 space-x-reverse">
      <button class="px-4 py-2 bg-green-600 text-white rounded-md hover:bg-green-700 transition-colors">
        شروع کرال
      </button>
      <button class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors">
        بروزرسانی
      </button>
    </div>
  </div>
</template>

<script setup>
// کامپوننت کرالر سرویس Emails
// این کامپوننت کرالر ایمیل را فراهم می‌کند
</script> 
