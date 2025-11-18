<template>
  <div class="bg-white rounded-lg shadow-md p-6">
    <h3 class="text-lg font-semibold text-gray-800 mb-4">مدیریت رقبا سرویس Torb</h3>
    
    <!-- آمار رقبا -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
      <div class="bg-blue-50 p-6 rounded-lg">
        <div class="text-sm text-blue-600">کل رقبا</div>
        <div class="text-2xl font-bold text-blue-800">۱۲</div>
      </div>
      <div class="bg-green-50 p-6 rounded-lg">
        <div class="text-sm text-green-600">رقبای فعال</div>
        <div class="text-2xl font-bold text-green-800">۸</div>
      </div>
      <div class="bg-orange-50 p-6 rounded-lg">
        <div class="text-sm text-orange-600">تغییرات قیمت</div>
        <div class="text-2xl font-bold text-orange-800">۲۴</div>
      </div>
      <div class="bg-purple-50 p-6 rounded-lg">
        <div class="text-sm text-purple-600">هشدارها</div>
        <div class="text-2xl font-bold text-purple-800">۵</div>
      </div>
    </div>

    <!-- فیلترهای رقبا -->
    <div class="mb-6">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">دسته‌بندی</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option>همه</option>
            <option>الکترونیک</option>
            <option>پوشاک</option>
            <option>کتاب</option>
            <option>خودرو</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option>همه</option>
            <option>فعال</option>
            <option>غیرفعال</option>
            <option>در انتظار</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">اولویت</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option>همه</option>
            <option>بالا</option>
            <option>متوسط</option>
            <option>پایین</option>
          </select>
        </div>
        <div class="flex items-end">
          <button class="w-full px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors">
            اعمال فیلتر
          </button>
        </div>
      </div>
    </div>

    <!-- جدول رقبا -->
    <div class="mb-6">
      <h4 class="text-md font-medium text-gray-700 mb-3">لیست رقبا</h4>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نام رقیب</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">دسته‌بندی</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">اولویت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">آخرین بروزرسانی</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">دیجی‌کالا</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">الکترونیک</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-green-100 text-green-800">فعال</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-red-100 text-red-800">بالا</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۲ ساعت پیش</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">مشاهده</button>
                <button class="text-green-600 hover:text-green-900">ویرایش</button>
                <button class="text-red-600 hover:text-red-900">حذف</button>
              </td>
            </tr>
            <tr>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">بامیلو</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">عمومی</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-green-100 text-green-800">فعال</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-orange-100 text-orange-800">متوسط</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۱ روز پیش</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">مشاهده</button>
                <button class="text-green-600 hover:text-green-900">ویرایش</button>
                <button class="text-red-600 hover:text-red-900">حذف</button>
              </td>
            </tr>
            <tr>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">اسنپ‌مارکت</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">عمومی</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-yellow-100 text-yellow-800">در انتظار</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-blue-100 text-blue-800">پایین</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۳ روز پیش</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">مشاهده</button>
                <button class="text-green-600 hover:text-green-900">ویرایش</button>
                <button class="text-red-600 hover:text-red-900">حذف</button>
              </td>
            </tr>
            <tr>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">کتاب‌راه</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">کتاب</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-green-100 text-green-800">فعال</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-orange-100 text-orange-800">متوسط</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۶ ساعت پیش</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">مشاهده</button>
                <button class="text-green-600 hover:text-green-900">ویرایش</button>
                <button class="text-red-600 hover:text-red-900">حذف</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- نمودار مقایسه -->
    <div class="mb-6">
      <h4 class="text-md font-medium text-gray-700 mb-3">نمودار مقایسه قیمت‌ها</h4>
      <div class="bg-gray-50 p-6 rounded-lg h-64 flex items-center justify-center">
        <div class="text-center">
          <div class="text-gray-500 mb-2">نمودار مقایسه</div>
          <div class="text-sm text-gray-400">نمایش مقایسه قیمت‌ها با رقبا</div>
        </div>
      </div>
    </div>

    <!-- هشدارهای قیمت -->
    <div class="mb-6">
      <h4 class="text-md font-medium text-gray-700 mb-3">هشدارهای قیمت</h4>
      <div class="space-y-3">
        <div class="border border-red-200 bg-red-50 p-6 rounded-lg">
          <div class="flex items-center justify-between">
            <div>
              <h5 class="font-medium text-red-900">کاهش قیمت در دیجی‌کالا</h5>
              <p class="text-sm text-red-700">قیمت گوشی سامسونگ ۲۰٪ کاهش یافته است</p>
            </div>
            <span class="text-xs text-red-600">۲ ساعت پیش</span>
          </div>
        </div>
        <div class="border border-orange-200 bg-orange-50 p-6 rounded-lg">
          <div class="flex items-center justify-between">
            <div>
              <h5 class="font-medium text-orange-900">افزایش قیمت در بامیلو</h5>
              <p class="text-sm text-orange-700">قیمت لپ‌تاپ ۱۵٪ افزایش یافته است</p>
            </div>
            <span class="text-xs text-orange-600">۱ روز پیش</span>
          </div>
        </div>
        <div class="border border-yellow-200 bg-yellow-50 p-6 rounded-lg">
          <div class="flex items-center justify-between">
            <div>
              <h5 class="font-medium text-yellow-900">محصول جدید در اسنپ‌مارکت</h5>
              <p class="text-sm text-yellow-700">محصول جدیدی به رقبای شما اضافه شده است</p>
            </div>
            <span class="text-xs text-yellow-600">۳ روز پیش</span>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات نظارت -->
    <div class="mb-6">
      <h4 class="text-md font-medium text-gray-700 mb-3">تنظیمات نظارت</h4>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="space-y-4">
          <div class="flex items-center justify-between">
            <div>
              <h5 class="font-medium text-gray-900">نظارت خودکار</h5>
              <p class="text-sm text-gray-600">نظارت خودکار بر تغییرات قیمت رقبا</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input type="checkbox" checked class="sr-only peer">
              <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
            </label>
          </div>
          
          <div class="flex items-center justify-between">
            <div>
              <h5 class="font-medium text-gray-900">اعلان‌های ایمیل</h5>
              <p class="text-sm text-gray-600">ارسال اعلان تغییرات قیمت از طریق ایمیل</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input type="checkbox" checked class="sr-only peer">
              <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
            </label>
          </div>
        </div>

        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">فرکانس بررسی (دقیقه)</label>
            <input type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500" value="30">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حد آستانه تغییر قیمت (٪)</label>
            <input type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500" value="10">
          </div>
        </div>
      </div>
    </div>

    <!-- دکمه‌های عملیات -->
    <div class="flex justify-end space-x-3 space-x-reverse">
      <button class="px-4 py-2 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400 transition-colors">
        افزودن رقیب
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
// کامپوننت مدیریت رقبا سرویس Torb
// این کامپوننت نظارت و مدیریت رقبا را فراهم می‌کند
</script> 
