<template>
  <div class="bg-white rounded-lg shadow-md p-6">
    <h3 class="text-lg font-semibold text-gray-800 mb-4">مدیریت ترافیک سرویس Torb</h3>
    
    <!-- آمار ترافیک -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
      <div class="bg-blue-50 p-6 rounded-lg">
        <div class="text-sm text-blue-600">بازدید امروز</div>
        <div class="text-2xl font-bold text-blue-800">۱۲,۴۵۰</div>
      </div>
      <div class="bg-green-50 p-6 rounded-lg">
        <div class="text-sm text-green-600">کاربران منحصر</div>
        <div class="text-2xl font-bold text-green-800">۸,۲۳۰</div>
      </div>
      <div class="bg-purple-50 p-6 rounded-lg">
        <div class="text-sm text-purple-600">نرخ پرش</div>
        <div class="text-2xl font-bold text-purple-800">۲۳.۵٪</div>
      </div>
      <div class="bg-orange-50 p-6 rounded-lg">
        <div class="text-sm text-orange-600">زمان متوسط</div>
        <div class="text-2xl font-bold text-orange-800">۴:۳۲</div>
      </div>
    </div>

    <!-- نمودار ترافیک -->
    <div class="mb-6">
      <h4 class="text-md font-medium text-gray-700 mb-3">نمودار ترافیک</h4>
      <div class="bg-gray-50 p-6 rounded-lg h-64 flex items-center justify-center">
        <div class="text-center">
          <div class="text-gray-500 mb-2">نمودار ترافیک</div>
          <div class="text-sm text-gray-400">نمایش روند ترافیک در طول زمان</div>
        </div>
      </div>
    </div>

    <!-- آمار تفصیلی -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-6">
      <div class="border border-gray-200 rounded-lg p-6">
        <h5 class="font-medium text-gray-900 mb-3">منابع ترافیک</h5>
        <div class="space-y-3">
          <div class="flex items-center justify-between">
            <div>
              <h6 class="font-medium text-gray-900">جستجوی ارگانیک</h6>
              <p class="text-sm text-gray-600">۴۵.۲٪ از کل ترافیک</p>
            </div>
            <span class="text-green-600 font-medium">۴۵.۲٪</span>
          </div>
          <div class="flex items-center justify-between">
            <div>
              <h6 class="font-medium text-gray-900">شبکه‌های اجتماعی</h6>
              <p class="text-sm text-gray-600">۲۸.۷٪ از کل ترافیک</p>
            </div>
            <span class="text-blue-600 font-medium">۲۸.۷٪</span>
          </div>
          <div class="flex items-center justify-between">
            <div>
              <h6 class="font-medium text-gray-900">تبلیغات</h6>
              <p class="text-sm text-gray-600">۱۸.۳٪ از کل ترافیک</p>
            </div>
            <span class="text-purple-600 font-medium">۱۸.۳٪</span>
          </div>
        </div>
      </div>

      <div class="border border-gray-200 rounded-lg p-6">
        <h5 class="font-medium text-gray-900 mb-3">آمار کلی</h5>
        <div class="space-y-3">
          <div class="flex items-center justify-between">
            <span class="text-gray-600">میانگین بازدید روزانه</span>
            <span class="font-medium">۱۱,۲۴۰</span>
          </div>
          <div class="flex items-center justify-between">
            <span class="text-gray-600">بیشترین بازدید</span>
            <span class="font-medium">۱۸,۵۶۰</span>
          </div>
          <div class="flex items-center justify-between">
            <span class="text-gray-600">کمترین بازدید</span>
            <span class="font-medium">۶,۸۹۰</span>
          </div>
          <div class="flex items-center justify-between">
            <span class="text-gray-600">رشد ماهانه</span>
            <span class="font-medium">+۱۲.۵٪</span>
          </div>
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
// کامپوننت مدیریت ترافیک سرویس Torb
// این کامپوننت مدیریت و نظارت بر ترافیک را فراهم می‌کند
</script> 
