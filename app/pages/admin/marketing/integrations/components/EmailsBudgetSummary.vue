<template>
  <div class="bg-white rounded-lg shadow-md p-6">
    <h3 class="text-lg font-semibold text-gray-800 mb-4">خلاصه بودجه سرویس Emails</h3>
    
    <!-- خلاصه کلی بودجه -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-6">
      <div class="bg-blue-50 p-6 rounded-lg">
        <div class="text-sm text-blue-600">بودجه کل</div>
        <div class="text-2xl font-bold text-blue-800">۸,۰۰۰,۰۰۰ تومان</div>
      </div>
      <div class="bg-green-50 p-6 rounded-lg">
        <div class="text-sm text-green-600">هزینه شده</div>
        <div class="text-2xl font-bold text-green-800">۵,۲۰۰,۰۰۰ تومان</div>
      </div>
      <div class="bg-orange-50 p-6 rounded-lg">
        <div class="text-sm text-orange-600">باقی‌مانده</div>
        <div class="text-2xl font-bold text-orange-800">۲,۸۰۰,۰۰۰ تومان</div>
      </div>
    </div>

    <!-- نمودار توزیع بودجه -->
    <div class="mb-6">
      <h4 class="text-md font-medium text-gray-700 mb-3">توزیع بودجه بر اساس نوع کمپین</h4>
      <div class="space-y-3">
        <div class="flex items-center justify-between">
          <span class="text-sm text-gray-600">خبرنامه‌ها</span>
          <div class="flex items-center space-x-2">
            <div class="w-32 bg-gray-200 rounded-full h-2">
              <div class="bg-blue-500 h-2 rounded-full" style="width: 40%"></div>
            </div>
            <span class="text-sm font-medium">۴۰٪</span>
          </div>
        </div>
        <div class="flex items-center justify-between">
          <span class="text-sm text-gray-600">تبلیغات محصولات</span>
          <div class="flex items-center space-x-2">
            <div class="w-32 bg-gray-200 rounded-full h-2">
              <div class="bg-green-500 h-2 rounded-full" style="width: 35%"></div>
            </div>
            <span class="text-sm font-medium">۳۵٪</span>
          </div>
        </div>
        <div class="flex items-center justify-between">
          <span class="text-sm text-gray-600">اطلاع‌رسانی‌ها</span>
          <div class="flex items-center space-x-2">
            <div class="w-32 bg-gray-200 rounded-full h-2">
              <div class="bg-orange-500 h-2 rounded-full" style="width: 20%"></div>
            </div>
            <span class="text-sm font-medium">۲۰٪</span>
          </div>
        </div>
        <div class="flex items-center justify-between">
          <span class="text-sm text-gray-600">سایر هزینه‌ها</span>
          <div class="flex items-center space-x-2">
            <div class="w-32 bg-gray-200 rounded-full h-2">
              <div class="bg-purple-500 h-2 rounded-full" style="width: 5%"></div>
            </div>
            <span class="text-sm font-medium">۵٪</span>
          </div>
        </div>
      </div>
    </div>

    <!-- جدول جزئیات بودجه -->
    <div>
      <h4 class="text-md font-medium text-gray-700 mb-3">جزئیات هزینه‌ها</h4>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع کمپین</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">بودجه تخصیص یافته</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">هزینه شده</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">باقی‌مانده</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نرخ مصرف</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">خبرنامه‌ها</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">۳,۲۰۰,۰۰۰ تومان</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">۲,۰۸۰,۰۰۰ تومان</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">۱,۱۲۰,۰۰۰ تومان</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-green-100 text-green-800">۶۵٪</span>
              </td>
            </tr>
            <tr>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">تبلیغات محصولات</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">۲,۸۰۰,۰۰۰ تومان</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">۱,۸۲۰,۰۰۰ تومان</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">۹۸۰,۰۰۰ تومان</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-yellow-100 text-yellow-800">۶۵٪</span>
              </td>
            </tr>
            <tr>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">اطلاع‌رسانی‌ها</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">۱,۶۰۰,۰۰۰ تومان</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">۱,۰۴۰,۰۰۰ تومان</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">۵۶۰,۰۰۰ تومان</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-blue-100 text-blue-800">۶۵٪</span>
              </td>
            </tr>
            <tr>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">سایر هزینه‌ها</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">۴۰۰,۰۰۰ تومان</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">۲۶۰,۰۰۰ تومان</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">۱۴۰,۰۰۰ تومان</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-purple-100 text-purple-800">۶۵٪</span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- پیش‌بینی بودجه -->
    <div class="mt-6">
      <h4 class="text-md font-medium text-gray-700 mb-3">پیش‌بینی بودجه</h4>
      <div class="bg-yellow-50 p-6 rounded-lg">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <svg class="h-5 w-5 text-yellow-400" viewBox="0 0 20 20" fill="currentColor">
              <path fill-rule="evenodd" d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z" clip-rule="evenodd" />
            </svg>
          </div>
          <div class="mr-3">
            <p class="text-sm text-yellow-800">
              بر اساس نرخ مصرف فعلی، بودجه شما تا <strong>۲۵ روز دیگر</strong> کافی خواهد بود.
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
// کامپوننت خلاصه بودجه سرویس Emails
// این کامپوننت نمایش خلاصه‌ای از بودجه و هزینه‌های سرویس ایمیل را فراهم می‌کند
</script> 
