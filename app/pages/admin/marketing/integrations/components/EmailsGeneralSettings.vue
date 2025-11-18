<template>
  <div class="bg-white rounded-lg shadow-md p-6">
    <h3 class="text-lg font-semibold text-gray-800 mb-4">تنظیمات عمومی سرویس Emails</h3>
    <!-- تنظیمات اصلی -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-6">
      <div class="space-y-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نام سرویس</label>
          <input type="text" value="Emails Service" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">آدرس سرور</label>
          <input type="text" value="smtp.emails.com" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">پورت</label>
          <input type="number" value="587" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
      </div>
      <div class="space-y-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر ایمیل در روز</label>
          <input type="number" value="10000" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تاخیر بین ارسال (ثانیه)</label>
          <input type="number" value="2" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">زبان پیش‌فرض</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option>فارسی</option>
            <option>انگلیسی</option>
            <option>عربی</option>
          </select>
        </div>
      </div>
    </div>
    <!-- دکمه‌های عملیات -->
    <div class="flex justify-end space-x-3 space-x-reverse">
      <button class="px-4 py-2 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400 transition-colors">
        بازنشانی
      </button>
      <button class="px-4 py-2 bg-green-600 text-white rounded-md hover:bg-green-700 transition-colors">
        ذخیره تنظیمات
      </button>
      <button class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors">
        بروزرسانی
      </button>
    </div>
  </div>
</template>
<script setup>
// کامپوننت تنظیمات عمومی سرویس Emails
</script> 
