<template>
  <div class="bg-white rounded-lg shadow-md p-6">
    <h3 class="text-lg font-semibold text-gray-800 mb-4">مدیریت کاربران سرویس Emails</h3>
    <!-- آمار کاربران -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
      <div class="bg-blue-50 p-6 rounded-lg">
        <div class="text-sm text-blue-600">کل کاربران</div>
        <div class="text-2xl font-bold text-blue-800">۸,۴۵۰</div>
      </div>
      <div class="bg-green-50 p-6 rounded-lg">
        <div class="text-sm text-green-600">کاربران فعال</div>
        <div class="text-2xl font-bold text-green-800">۶,۲۳۰</div>
      </div>
      <div class="bg-purple-50 p-6 rounded-lg">
        <div class="text-sm text-purple-600">کاربران جدید</div>
        <div class="text-2xl font-bold text-purple-800">۸۵</div>
      </div>
      <div class="bg-orange-50 p-6 rounded-lg">
        <div class="text-sm text-orange-600">کاربران ویژه</div>
        <div class="text-2xl font-bold text-orange-800">۱,۰۴۰</div>
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
            <option>ویژه</option>
            <option>ادمین</option>
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
                    <div class="text-sm text-gray-500">ali.ahmadi@example.com</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">ali.ahmadi@example.com</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-orange-100 text-orange-800">ویژه</span>
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
                    <div class="text-sm text-gray-500">sara.mohammadi@example.com</div>
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
                    <div class="text-sm text-gray-500">mohammad.rezaei@example.com</div>
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
                    <div class="text-sm text-gray-500">fateme.karimi@example.com</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">fateme.karimi@example.com</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-purple-100 text-purple-800">ویژه</span>
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
// کامپوننت مدیریت کاربران سرویس Emails
// این کامپوننت مدیریت و نظارت بر کاربران را فراهم می‌کند
</script> 
