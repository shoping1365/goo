<template>
  <div class="bg-white rounded-lg shadow-md p-6">
    <h3 class="text-lg font-semibold text-gray-800 mb-4">مدیریت محصولات سرویس Torb</h3>
    
    <!-- آمار محصولات -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
      <div class="bg-blue-50 p-6 rounded-lg">
        <div class="text-sm text-blue-600">کل محصولات</div>
        <div class="text-2xl font-bold text-blue-800">۲,۴۵۰</div>
      </div>
      <div class="bg-green-50 p-6 rounded-lg">
        <div class="text-sm text-green-600">محصولات فعال</div>
        <div class="text-2xl font-bold text-green-800">۱,۸۲۰</div>
      </div>
      <div class="bg-purple-50 p-6 rounded-lg">
        <div class="text-sm text-purple-600">محصولات جدید</div>
        <div class="text-2xl font-bold text-purple-800">۱۲۵</div>
      </div>
      <div class="bg-orange-50 p-6 rounded-lg">
        <div class="text-sm text-orange-600">محصولات تخفیف‌دار</div>
        <div class="text-2xl font-bold text-orange-800">۳۴۰</div>
      </div>
    </div>

    <!-- فیلترهای محصولات -->
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
          <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option>همه</option>
            <option>فعال</option>
            <option>غیرفعال</option>
            <option>در انتظار</option>
            <option>ناموجود</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">برند</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option>همه</option>
            <option>سامسونگ</option>
            <option>اپل</option>
            <option>شیائومی</option>
            <option>هوآوی</option>
          </select>
        </div>
        <div class="flex items-end">
          <button class="w-full px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors">
            اعمال فیلتر
          </button>
        </div>
      </div>
    </div>

    <!-- جدول محصولات -->
    <div class="mb-6">
      <h4 class="text-md font-medium text-gray-700 mb-3">لیست محصولات</h4>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">محصول</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">دسته‌بندی</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">برند</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">قیمت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">موجودی</th>
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
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">الکترونیک</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">سامسونگ</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۴۵,۰۰۰,۰۰۰ تومان</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-green-100 text-green-800">فعال</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۱۵</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">مشاهده</button>
                <button class="text-green-600 hover:text-green-900">ویرایش</button>
                <button class="text-red-600 hover:text-red-900">حذف</button>
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
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">الکترونیک</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">اپل</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۸۵,۰۰۰,۰۰۰ تومان</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-green-100 text-green-800">فعال</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۸</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">مشاهده</button>
                <button class="text-green-600 hover:text-green-900">ویرایش</button>
                <button class="text-red-600 hover:text-red-900">حذف</button>
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
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">کتاب</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">نشر فنی</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۲۵۰,۰۰۰ تومان</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-yellow-100 text-yellow-800">در انتظار</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۰</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">مشاهده</button>
                <button class="text-green-600 hover:text-green-900">ویرایش</button>
                <button class="text-red-600 hover:text-red-900">حذف</button>
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
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">پوشاک</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">نایک</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۱,۲۰۰,۰۰۰ تومان</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-green-100 text-green-800">فعال</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۴۵</td>
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

    <!-- نمودار محصولات -->
    <div class="mb-6">
      <h4 class="text-md font-medium text-gray-700 mb-3">نمودار توزیع محصولات</h4>
      <div class="bg-gray-50 p-6 rounded-lg h-64 flex items-center justify-center">
        <div class="text-center">
          <div class="text-gray-500 mb-2">نمودار توزیع</div>
          <div class="text-sm text-gray-400">نمایش توزیع محصولات بر اساس دسته‌بندی</div>
        </div>
      </div>
    </div>

    <!-- آمار تفصیلی -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-6">
      <div class="border border-gray-200 rounded-lg p-6">
        <h5 class="font-medium text-gray-900 mb-3">محصولات پرفروش</h5>
        <div class="space-y-3">
          <div class="flex items-center justify-between">
            <div>
              <h6 class="font-medium text-gray-900">گوشی سامسونگ Galaxy S24</h6>
              <p class="text-sm text-gray-600">فروش: ۲۵۰ عدد</p>
            </div>
            <span class="text-green-600 font-medium">۲۵۰</span>
          </div>
          <div class="flex items-center justify-between">
            <div>
              <h6 class="font-medium text-gray-900">لپ‌تاپ اپل MacBook Pro</h6>
              <p class="text-sm text-gray-600">فروش: ۱۸۰ عدد</p>
            </div>
            <span class="text-green-600 font-medium">۱۸۰</span>
          </div>
          <div class="flex items-center justify-between">
            <div>
              <h6 class="font-medium text-gray-900">کفش ورزشی نایک</h6>
              <p class="text-sm text-gray-600">فروش: ۳۲۰ عدد</p>
            </div>
            <span class="text-green-600 font-medium">۳۲۰</span>
          </div>
        </div>
      </div>

      <div class="border border-gray-200 rounded-lg p-6">
        <h5 class="font-medium text-gray-900 mb-3">آمار کلی</h5>
        <div class="space-y-3">
          <div class="flex items-center justify-between">
            <span class="text-gray-600">میانگین قیمت</span>
            <span class="font-medium">۱۲,۵۰۰,۰۰۰ تومان</span>
          </div>
          <div class="flex items-center justify-between">
            <span class="text-gray-600">تعداد برندها</span>
            <span class="font-medium">۴۵</span>
          </div>
          <div class="flex items-center justify-between">
            <span class="text-gray-600">محصولات تخفیف‌دار</span>
            <span class="font-medium">۱۳.۸٪</span>
          </div>
          <div class="flex items-center justify-between">
            <span class="text-gray-600">نرخ موجودی</span>
            <span class="font-medium">۷۴.۳٪</span>
          </div>
        </div>
      </div>
    </div>

    <!-- دکمه‌های عملیات -->
    <div class="flex justify-end space-x-3 space-x-reverse">
      <button class="px-4 py-2 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400 transition-colors">
        افزودن محصول جدید
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
// کامپوننت مدیریت محصولات سرویس Torb
// این کامپوننت مدیریت و نظارت بر محصولات را فراهم می‌کند
</script> 
