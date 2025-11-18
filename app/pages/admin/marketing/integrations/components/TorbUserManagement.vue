<template>
  <div class="bg-white rounded-lg shadow-md p-6">
    <h3 class="text-lg font-semibold text-gray-800 mb-4">مدیریت کاربران سرویس Torb</h3>
    
    <!-- آمار کاربران -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
      <div class="bg-blue-50 p-6 rounded-lg">
        <div class="text-sm text-blue-600">کل کاربران</div>
        <div class="text-2xl font-bold text-blue-800">۱۲,۴۵۰</div>
      </div>
      <div class="bg-green-50 p-6 rounded-lg">
        <div class="text-sm text-green-600">کاربران فعال</div>
        <div class="text-2xl font-bold text-green-800">۸,۲۳۰</div>
      </div>
      <div class="bg-purple-50 p-6 rounded-lg">
        <div class="text-sm text-purple-600">کاربران جدید</div>
        <div class="text-2xl font-bold text-purple-800">۱۸۵</div>
      </div>
      <div class="bg-orange-50 p-6 rounded-lg">
        <div class="text-sm text-orange-600">کاربران پریمیوم</div>
        <div class="text-2xl font-bold text-orange-800">۱,۲۴۰</div>
      </div>
    </div>

    <!-- فیلترهای کاربران -->
    <div class="mb-6">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نوع کاربر</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option>همه</option>
            <option>عادی</option>
            <option>پریمیوم</option>
            <option>ادمین</option>
            <option>مشتری ویژه</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option>همه</option>
            <option>فعال</option>
            <option>غیرفعال</option>
            <option>در انتظار</option>
            <option>مسدود</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ عضویت</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option>همه</option>
            <option>امروز</option>
            <option>هفته گذشته</option>
            <option>ماه گذشته</option>
            <option>سال گذشته</option>
          </select>
        </div>
        <div class="flex items-end">
          <button class="w-full px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors">
            اعمال فیلتر
          </button>
        </div>
      </div>
    </div>

    <!-- جدول کاربران -->
    <div class="mb-6">
      <h4 class="text-md font-medium text-gray-700 mb-3">لیست کاربران</h4>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کاربر</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">ایمیل</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ عضویت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">آخرین ورود</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="w-10 h-10 bg-blue-500 rounded-full flex items-center justify-center text-white text-sm font-medium">ع</div>
                  <div class="mr-3">
                    <div class="text-sm font-medium text-gray-900">علی احمدی</div>
                    <div class="text-sm text-gray-500">+۹۸ ۹۱۲ ۱۲۳ ۴۵۶۷</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">ali.ahmadi@example.com</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-orange-100 text-orange-800">پریمیوم</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-green-100 text-green-800">فعال</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۱۴۰۳/۰۱/۱۵</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">امروز ۱۰:۳۰</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">مشاهده</button>
                <button class="text-green-600 hover:text-green-900">ویرایش</button>
                <button class="text-red-600 hover:text-red-900">مسدود</button>
              </td>
            </tr>
            <tr>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="w-10 h-10 bg-green-500 rounded-full flex items-center justify-center text-white text-sm font-medium">س</div>
                  <div class="mr-3">
                    <div class="text-sm font-medium text-gray-900">سارا محمدی</div>
                    <div class="text-sm text-gray-500">+۹۸ ۹۱۳ ۲۳۴ ۵۶۷۸</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">sara.mohammadi@example.com</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-blue-100 text-blue-800">عادی</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-green-100 text-green-800">فعال</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۱۴۰۳/۰۲/۰۸</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">دیروز ۱۶:۴۵</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">مشاهده</button>
                <button class="text-green-600 hover:text-green-900">ویرایش</button>
                <button class="text-red-600 hover:text-red-900">مسدود</button>
              </td>
            </tr>
            <tr>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="w-10 h-10 bg-purple-500 rounded-full flex items-center justify-center text-white text-sm font-medium">م</div>
                  <div class="mr-3">
                    <div class="text-sm font-medium text-gray-900">محمد رضایی</div>
                    <div class="text-sm text-gray-500">+۹۸ ۹۱۴ ۳۴۵ ۶۷۸۹</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">mohammad.rezaei@example.com</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-red-100 text-red-800">ادمین</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-green-100 text-green-800">فعال</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۱۴۰۲/۱۱/۲۰</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">امروز ۰۸:۱۵</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">مشاهده</button>
                <button class="text-green-600 hover:text-green-900">ویرایش</button>
                <button class="text-red-600 hover:text-red-900">مسدود</button>
              </td>
            </tr>
            <tr>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="w-10 h-10 bg-yellow-500 rounded-full flex items-center justify-center text-white text-sm font-medium">ف</div>
                  <div class="mr-3">
                    <div class="text-sm font-medium text-gray-900">فاطمه کریمی</div>
                    <div class="text-sm text-gray-500">+۹۸ ۹۱۵ ۴۵۶ ۷۸۹۰</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">fateme.karimi@example.com</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-purple-100 text-purple-800">مشتری ویژه</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-yellow-100 text-yellow-800">در انتظار</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۱۴۰۳/۰۳/۱۲</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">هیچ‌وقت</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">مشاهده</button>
                <button class="text-green-600 hover:text-green-900">ویرایش</button>
                <button class="text-red-600 hover:text-red-900">مسدود</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- نمودار کاربران -->
    <div class="mb-6">
      <h4 class="text-md font-medium text-gray-700 mb-3">نمودار توزیع کاربران</h4>
      <div class="bg-gray-50 p-6 rounded-lg h-64 flex items-center justify-center">
        <div class="text-center">
          <div class="text-gray-500 mb-2">نمودار توزیع</div>
          <div class="text-sm text-gray-400">نمایش توزیع کاربران بر اساس نوع و وضعیت</div>
        </div>
      </div>
    </div>

    <!-- آمار تفصیلی -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-6">
      <div class="border border-gray-200 rounded-lg p-6">
        <h5 class="font-medium text-gray-900 mb-3">کاربران فعال</h5>
        <div class="space-y-3">
          <div class="flex items-center justify-between">
            <div>
              <h6 class="font-medium text-gray-900">کاربران عادی</h6>
              <p class="text-sm text-gray-600">۶,۹۹۰ کاربر</p>
            </div>
            <span class="text-blue-600 font-medium">۵۶.۲٪</span>
          </div>
          <div class="flex items-center justify-between">
            <div>
              <h6 class="font-medium text-gray-900">کاربران پریمیوم</h6>
              <p class="text-sm text-gray-600">۱,۲۴۰ کاربر</p>
            </div>
            <span class="text-orange-600 font-medium">۱۰.۰٪</span>
          </div>
          <div class="flex items-center justify-between">
            <div>
              <h6 class="font-medium text-gray-900">مشتریان ویژه</h6>
              <p class="text-sm text-gray-600">۸۵۰ کاربر</p>
            </div>
            <span class="text-purple-600 font-medium">۶.۸٪</span>
          </div>
        </div>
      </div>

      <div class="border border-gray-200 rounded-lg p-6">
        <h5 class="font-medium text-gray-900 mb-3">آمار کلی</h5>
        <div class="space-y-3">
          <div class="flex items-center justify-between">
            <span class="text-gray-600">نرخ فعال‌سازی</span>
            <span class="font-medium">۶۶.۱٪</span>
          </div>
          <div class="flex items-center justify-between">
            <span class="text-gray-600">میانگین سن</span>
            <span class="font-medium">۲۸.۵ سال</span>
          </div>
          <div class="flex items-center justify-between">
            <span class="text-gray-600">کاربران مسدود</span>
            <span class="font-medium">۱۲۵</span>
          </div>
          <div class="flex items-center justify-between">
            <span class="text-gray-600">نرخ نگهداری</span>
            <span class="font-medium">۸۷.۳٪</span>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات مدیریت کاربران -->
    <div class="mb-6">
      <h4 class="text-md font-medium text-gray-700 mb-3">تنظیمات مدیریت کاربران</h4>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="space-y-4">
          <div class="flex items-center justify-between">
            <div>
              <h5 class="font-medium text-gray-900">تایید خودکار ایمیل</h5>
              <p class="text-sm text-gray-600">تایید خودکار ایمیل کاربران جدید</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input type="checkbox" class="sr-only peer">
              <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
            </label>
          </div>
          
          <div class="flex items-center justify-between">
            <div>
              <h5 class="font-medium text-gray-900">اعلان کاربران جدید</h5>
              <p class="text-sm text-gray-600">اعلان ثبت‌نام کاربران جدید</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input type="checkbox" checked class="sr-only peer">
              <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
            </label>
          </div>

          <div class="flex items-center justify-between">
            <div>
              <h5 class="font-medium text-gray-900">مسدودسازی خودکار</h5>
              <p class="text-sm text-gray-600">مسدودسازی خودکار کاربران مشکوک</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input type="checkbox" checked class="sr-only peer">
              <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
            </label>
          </div>
        </div>

        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر تلاش ورود</label>
            <input type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500" value="5">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">مدت زمان مسدودسازی (دقیقه)</label>
            <input type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500" value="30">
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداقل طول رمز عبور</label>
            <input type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500" value="8">
          </div>
        </div>
      </div>
    </div>

    <!-- دکمه‌های عملیات -->
    <div class="flex justify-end space-x-3 space-x-reverse">
      <button class="px-4 py-2 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400 transition-colors">
        افزودن کاربر جدید
      </button>
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
// کامپوننت مدیریت کاربران سرویس Torb
// این کامپوننت مدیریت و نظارت بر کاربران را فراهم می‌کند
</script> 
