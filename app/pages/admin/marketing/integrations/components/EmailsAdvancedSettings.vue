<template>
  <div class="bg-white rounded-lg shadow-md p-6">
    <h3 class="text-lg font-semibold text-gray-800 mb-4">تنظیمات پیشرفته سرویس Emails</h3>
    
    <!-- تنظیمات امنیتی -->
    <div class="mb-6">
      <h4 class="text-md font-medium text-gray-700 mb-3">تنظیمات امنیتی</h4>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="space-y-4">
          <div class="flex items-center justify-between">
            <div>
              <h5 class="font-medium text-gray-900">احراز هویت دو مرحله‌ای</h5>
              <p class="text-sm text-gray-600">فعال‌سازی احراز هویت اضافی برای امنیت بیشتر</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input type="checkbox" checked class="sr-only peer">
              <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
            </label>
          </div>
          
          <div class="flex items-center justify-between">
            <div>
              <h5 class="font-medium text-gray-900">رمزگذاری ایمیل‌ها</h5>
              <p class="text-sm text-gray-600">رمزگذاری تمام ایمیل‌های ارسالی</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input type="checkbox" checked class="sr-only peer">
              <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
            </label>
          </div>
          
          <div class="flex items-center justify-between">
            <div>
              <h5 class="font-medium text-gray-900">لاگ فعالیت‌ها</h5>
              <p class="text-sm text-gray-600">ثبت تمام فعالیت‌های ارسال ایمیل</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input type="checkbox" checked class="sr-only peer">
              <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
            </label>
          </div>
        </div>

        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">مدت زمان نشست (دقیقه)</label>
            <input type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500" value="90">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر تلاش ورود</label>
            <input type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500" value="3">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">زمان قفل حساب (دقیقه)</label>
            <input type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500" value="15">
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات عملکرد -->
    <div class="mb-6">
      <h4 class="text-md font-medium text-gray-700 mb-3">تنظیمات عملکرد</h4>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر ایمیل در ساعت</label>
            <input type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500" value="1000">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تاخیر بین ارسال (ثانیه)</label>
            <input type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500" value="2">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">اندازه کش (MB)</label>
            <input type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500" value="256">
          </div>
        </div>

        <div class="space-y-4">
          <div class="flex items-center justify-between">
            <div>
              <h5 class="font-medium text-gray-900">فشرده‌سازی ایمیل‌ها</h5>
              <p class="text-sm text-gray-600">فشرده‌سازی ایمیل‌ها برای کاهش حجم</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input type="checkbox" checked class="sr-only peer">
              <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
            </label>
          </div>
          
          <div class="flex items-center justify-between">
            <div>
              <h5 class="font-medium text-gray-900">بهینه‌سازی خودکار</h5>
              <p class="text-sm text-gray-600">بهینه‌سازی خودکار عملکرد ارسال</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input type="checkbox" class="sr-only peer">
              <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
            </label>
          </div>
          
          <div class="flex items-center justify-between">
            <div>
              <h5 class="font-medium text-gray-900">نمایش آمار عملکرد</h5>
              <p class="text-sm text-gray-600">نمایش آمار عملکرد در داشبورد</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input type="checkbox" checked class="sr-only peer">
              <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
            </label>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات اعلان‌ها -->
    <div class="mb-6">
      <h4 class="text-md font-medium text-gray-700 mb-3">تنظیمات اعلان‌ها</h4>
      <div class="space-y-4">
        <div class="flex items-center justify-between">
          <div>
            <h5 class="font-medium text-gray-900">اعلان‌های ارسال موفق</h5>
            <p class="text-sm text-gray-600">اعلان موفقیت‌آمیز بودن ارسال ایمیل</p>
          </div>
          <label class="relative inline-flex items-center cursor-pointer">
            <input type="checkbox" checked class="sr-only peer">
            <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
          </label>
        </div>
        
        <div class="flex items-center justify-between">
          <div>
            <h5 class="font-medium text-gray-900">اعلان‌های خطا</h5>
            <p class="text-sm text-gray-600">اعلان خطاهای ارسال ایمیل</p>
          </div>
          <label class="relative inline-flex items-center cursor-pointer">
            <input type="checkbox" checked class="sr-only peer">
            <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
          </label>
        </div>
        
        <div class="flex items-center justify-between">
          <div>
            <h5 class="font-medium text-gray-900">اعلان‌های محدودیت</h5>
            <p class="text-sm text-gray-600">اعلان رسیدن به محدودیت ارسال</p>
          </div>
          <label class="relative inline-flex items-center cursor-pointer">
            <input type="checkbox" class="sr-only peer">
            <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
          </label>
        </div>
      </div>
    </div>

    <!-- تنظیمات پشتیبان‌گیری -->
    <div class="mb-6">
      <h4 class="text-md font-medium text-gray-700 mb-3">تنظیمات پشتیبان‌گیری</h4>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="space-y-4">
          <div class="flex items-center justify-between">
            <div>
              <h5 class="font-medium text-gray-900">پشتیبان‌گیری خودکار</h5>
              <p class="text-sm text-gray-600">پشتیبان‌گیری خودکار از قالب‌ها و لیست‌ها</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input type="checkbox" checked class="sr-only peer">
              <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
            </label>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">فرکانس پشتیبان‌گیری</label>
            <select class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
              <option>روزانه</option>
              <option>هفتگی</option>
              <option>ماهانه</option>
            </select>
          </div>
        </div>

        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تعداد نسخه‌های نگهداری</label>
            <input type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500" value="5">
          </div>
          
          <div class="flex items-center justify-between">
            <div>
              <h5 class="font-medium text-gray-900">فشرده‌سازی پشتیبان</h5>
              <p class="text-sm text-gray-600">فشرده‌سازی فایل‌های پشتیبان</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input type="checkbox" checked class="sr-only peer">
              <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
            </label>
          </div>
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
        تست تنظیمات
      </button>
    </div>
  </div>
</template>

<script setup>
// کامپوننت تنظیمات پیشرفته سرویس Emails
// این کامپوننت تنظیمات امنیتی، عملکرد و پشتیبان‌گیری سرویس Emails را مدیریت می‌کند
</script> 
