<template>
  <div class="bg-white rounded-lg shadow-md p-6">
    <h3 class="text-lg font-semibold text-gray-800 mb-4">مدیریت قیمت‌ها سرویس Emails</h3>
    <!-- آمار قیمت‌ها -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
      <div class="bg-blue-50 p-6 rounded-lg">
        <div class="text-sm text-blue-600">کل محصولات</div>
        <div class="text-2xl font-bold text-blue-800">۱,۸۵۰</div>
      </div>
      <div class="bg-green-50 p-6 rounded-lg">
        <div class="text-sm text-green-600">تغییرات امروز</div>
        <div class="text-2xl font-bold text-green-800">۲۸</div>
      </div>
      <div class="bg-purple-50 p-6 rounded-lg">
        <div class="text-sm text-purple-600">محصولات تخفیف‌دار</div>
        <div class="text-2xl font-bold text-purple-800">۲۴۰</div>
      </div>
      <div class="bg-orange-50 p-6 rounded-lg">
        <div class="text-sm text-orange-600">میانگین تخفیف</div>
        <div class="text-2xl font-bold text-orange-800">۱۵.۵٪</div>
      </div>
    </div>
    <!-- فیلترهای قیمت -->
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
            <option>لوازم خانگی</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نوع تغییر</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option>همه</option>
            <option>افزایش قیمت</option>
            <option>کاهش قیمت</option>
            <option>تخفیف</option>
            <option>بدون تغییر</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">بازه قیمت</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option>همه</option>
            <option>کمتر از ۱ میلیون</option>
            <option>۱ تا ۵ میلیون</option>
            <option>۵ تا ۲۰ میلیون</option>
            <option>بیش از ۲۰ میلیون</option>
          </select>
        </div>
        <div class="flex items-end">
          <button class="w-full px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors">
            اعمال فیلتر
          </button>
        </div>
      </div>
    </div>
    <!-- جدول قیمت‌ها -->
    <div class="mb-6">
      <h4 class="text-md font-medium text-gray-700 mb-3">لیست تغییرات قیمت</h4>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">محصول</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">قیمت قبلی</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">قیمت جدید</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تغییر</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="w-10 h-10 bg-gray-200 rounded-lg flex items-center justify-center">
                    <span class="text-xs text-gray-500">تصویر</span>
                  </div>
                  <div class="mr-3">
                    <div class="text-sm font-medium text-gray-900">گوشی سامسونگ Galaxy S24</div>
                    <div class="text-sm text-gray-500">SKU: SAMS-S24-128</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۴۸,۰۰۰,۰۰۰ تومان</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۴۵,۰۰۰,۰۰۰ تومان</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-green-100 text-green-800">-۶.۲۵٪</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-blue-100 text-blue-800">کاهش قیمت</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">امروز ۱۴:۳۰</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">مشاهده</button>
                <button class="text-green-600 hover:text-green-900">ویرایش</button>
                <button class="text-red-600 hover:text-red-900">بازگشت</button>
              </td>
            </tr>
            <tr>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="w-10 h-10 bg-gray-200 rounded-lg flex items-center justify-center">
                    <span class="text-xs text-gray-500">تصویر</span>
                  </div>
                  <div class="mr-3">
                    <div class="text-sm font-medium text-gray-900">لپ‌تاپ اپل MacBook Pro</div>
                    <div class="text-sm text-gray-500">SKU: APPLE-MBP-256</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۸۲,۰۰۰,۰۰۰ تومان</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۸۵,۰۰۰,۰۰۰ تومان</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-red-100 text-red-800">+۳.۶۶٪</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-red-100 text-red-800">افزایش قیمت</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">دیروز ۱۶:۴۵</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">مشاهده</button>
                <button class="text-green-600 hover:text-green-900">ویرایش</button>
                <button class="text-red-600 hover:text-red-900">بازگشت</button>
              </td>
            </tr>
            <tr>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="w-10 h-10 bg-gray-200 rounded-lg flex items-center justify-center">
                    <span class="text-xs text-gray-500">تصویر</span>
                  </div>
                  <div class="mr-3">
                    <div class="text-sm font-medium text-gray-900">کفش ورزشی نایک</div>
                    <div class="text-sm text-gray-500">SKU: NIKE-SHOE-001</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۱,۵۰۰,۰۰۰ تومان</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۱,۲۰۰,۰۰۰ تومان</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-green-100 text-green-800">-۲۰٪</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-purple-100 text-purple-800">تخفیف</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۲ روز پیش</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">مشاهده</button>
                <button class="text-green-600 hover:text-green-900">ویرایش</button>
                <button class="text-red-600 hover:text-red-900">بازگشت</button>
              </td>
            </tr>
            <tr>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="w-10 h-10 bg-gray-200 rounded-lg flex items-center justify-center">
                    <span class="text-xs text-gray-500">تصویر</span>
                  </div>
                  <div class="mr-3">
                    <div class="text-sm font-medium text-gray-900">کتاب برنامه‌نویسی React</div>
                    <div class="text-sm text-gray-500">SKU: BOOK-REACT-001</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۳۰۰,۰۰۰ تومان</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۲۵۰,۰۰۰ تومان</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-green-100 text-green-800">-۱۶.۶۷٪</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-purple-100 text-purple-800">تخفیف</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۳ روز پیش</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">مشاهده</button>
                <button class="text-green-600 hover:text-green-900">ویرایش</button>
                <button class="text-red-600 hover:text-red-900">بازگشت</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
    <!-- نمودار تغییرات قیمت -->
    <div class="mb-6">
      <h4 class="text-md font-medium text-gray-700 mb-3">نمودار تغییرات قیمت</h4>
      <div class="bg-gray-50 p-6 rounded-lg h-64 flex items-center justify-center">
        <div class="text-center">
          <div class="text-gray-500 mb-2">نمودار تغییرات</div>
          <div class="text-sm text-gray-400">نمایش روند تغییرات قیمت در طول زمان</div>
        </div>
      </div>
    </div>
    <!-- آمار تفصیلی -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-6">
      <div class="border border-gray-200 rounded-lg p-6">
        <h5 class="font-medium text-gray-900 mb-3">تغییرات اخیر</h5>
        <div class="space-y-3">
          <div class="flex items-center justify-between">
            <div>
              <h6 class="font-medium text-gray-900">کاهش قیمت</h6>
              <p class="text-sm text-gray-600">۱۸ محصول</p>
            </div>
            <span class="text-green-600 font-medium">-۹.۵٪</span>
          </div>
          <div class="flex items-center justify-between">
            <div>
              <h6 class="font-medium text-gray-900">افزایش قیمت</h6>
              <p class="text-sm text-gray-600">۹ محصول</p>
            </div>
            <span class="text-red-600 font-medium">+۵.۳٪</span>
          </div>
          <div class="flex items-center justify-between">
            <div>
              <h6 class="font-medium text-gray-900">تخفیف‌ها</h6>
              <p class="text-sm text-gray-600">۳ محصول</p>
            </div>
            <span class="text-purple-600 font-medium">-۲۵٪</span>
          </div>
        </div>
      </div>
      <div class="border border-gray-200 rounded-lg p-6">
        <h5 class="font-medium text-gray-900 mb-3">آمار کلی</h5>
        <div class="space-y-3">
          <div class="flex items-center justify-between">
            <span class="text-gray-600">میانگین تغییر قیمت</span>
            <span class="font-medium">-۲.۱٪</span>
          </div>
          <div class="flex items-center justify-between">
            <span class="text-gray-600">بیشترین کاهش</span>
            <span class="font-medium">-۲۸٪</span>
          </div>
          <div class="flex items-center justify-between">
            <span class="text-gray-600">بیشترین افزایش</span>
            <span class="font-medium">+۱۷٪</span>
          </div>
          <div class="flex items-center justify-between">
            <span class="text-gray-600">محصولات بدون تغییر</span>
            <span class="font-medium">۱,۷۸۷</span>
          </div>
        </div>
      </div>
    </div>
    <!-- تنظیمات قیمت‌گذاری -->
    <div class="mb-6">
      <h4 class="text-md font-medium text-gray-700 mb-3">تنظیمات قیمت‌گذاری</h4>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="space-y-4">
          <div class="flex items-center justify-between">
            <div>
              <h5 class="font-medium text-gray-900">قیمت‌گذاری خودکار</h5>
              <p class="text-sm text-gray-600">تنظیم خودکار قیمت بر اساس رقبا</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input type="checkbox" checked class="sr-only peer">
              <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
            </label>
          </div>
          <div class="flex items-center justify-between">
            <div>
              <h5 class="font-medium text-gray-900">اعلان تغییرات قیمت</h5>
              <p class="text-sm text-gray-600">اعلان تغییرات قیمت رقبا</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input type="checkbox" checked class="sr-only peer">
              <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
            </label>
          </div>
          <div class="flex items-center justify-between">
            <div>
              <h5 class="font-medium text-gray-900">تخفیف خودکار</h5>
              <p class="text-sm text-gray-600">اعمال تخفیف خودکار بر اساس موجودی</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input type="checkbox" class="sr-only peer">
              <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
            </label>
          </div>
        </div>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حد آستانه تغییر قیمت (٪)</label>
            <input type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500" value="5">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر تخفیف (٪)</label>
            <input type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500" value="30">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">فرکانس بررسی قیمت (ساعت)</label>
            <input type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500" value="2">
          </div>
        </div>
      </div>
    </div>
    <!-- دکمه‌های عملیات -->
    <div class="flex justify-end space-x-3 space-x-reverse">
      <button class="px-4 py-2 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400 transition-colors">
        تغییر دسته‌ای قیمت
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
// کامپوننت مدیریت قیمت‌ها سرویس Emails
// این کامپوننت مدیریت و نظارت بر تغییرات قیمت را فراهم می‌کند
</script> 
