<template>
  <div class="seo-monitoring" dir="rtl">
    <div class="bg-white rounded-lg shadow-md p-6">
      <div class="flex items-center justify-between mb-6">
        <h1 class="text-2xl font-bold text-gray-800">نظارت و گزارش‌گیری</h1>
        <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors">
          تولید گزارش جدید
        </button>
      </div>

      <!-- آمار کلی -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
        <div class="bg-gradient-to-r from-blue-500 to-blue-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۱۵,۲۵۰</div>
          <div class="text-sm">بازدید ارگانیک</div>
        </div>
        <div class="bg-gradient-to-r from-green-500 to-green-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۸۵</div>
          <div class="text-sm">امتیاز سئو</div>
        </div>
        <div class="bg-gradient-to-r from-yellow-500 to-yellow-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۲۵۰</div>
          <div class="text-sm">کلمات کلیدی</div>
        </div>
        <div class="bg-gradient-to-r from-red-500 to-red-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۱۲</div>
          <div class="text-sm">هشدارها</div>
        </div>
      </div>

      <!-- نمودار روند -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-6">
        <!-- روند ترافیک -->
        <div class="bg-white border border-gray-200 rounded-lg p-6">
          <h3 class="text-lg font-semibold text-gray-800 mb-4">روند ترافیک ارگانیک</h3>
          <div class="h-48 bg-gray-50 rounded-lg flex items-center justify-center">
            <div class="text-center">
              <div class="text-2xl font-bold text-gray-600">نمودار</div>
              <div class="text-sm text-gray-500">روند ۳۰ روز گذشته</div>
            </div>
          </div>
        </div>

        <!-- رتبه‌بندی کلمات کلیدی -->
        <div class="bg-white border border-gray-200 rounded-lg p-6">
          <h3 class="text-lg font-semibold text-gray-800 mb-4">رتبه‌بندی کلمات کلیدی</h3>
          <div class="h-48 bg-gray-50 rounded-lg flex items-center justify-center">
            <div class="text-center">
              <div class="text-2xl font-bold text-gray-600">نمودار</div>
              <div class="text-sm text-gray-500">تغییرات رتبه</div>
            </div>
          </div>
        </div>
      </div>

      <!-- هشدارها -->
      <div class="bg-white border border-gray-200 rounded-lg p-6 mb-6">
        <h3 class="text-lg font-semibold text-gray-800 mb-4">هشدارهای سئو</h3>
        <div class="space-y-3">
          <div class="flex items-center justify-between p-3 bg-red-50 rounded-lg">
            <div class="flex items-center">
              <span class="text-red-600 text-lg ml-2">⚠️</span>
              <span class="text-sm text-gray-700">صفحه محصولات کند بارگذاری می‌شود</span>
            </div>
            <span class="text-xs text-gray-500">۲ ساعت پیش</span>
          </div>
          <div class="flex items-center justify-between p-3 bg-yellow-50 rounded-lg">
            <div class="flex items-center">
              <span class="text-yellow-600 text-lg ml-2">⚠️</span>
              <span class="text-sm text-gray-700">۵ تصویر بدون Alt Text</span>
            </div>
            <span class="text-xs text-gray-500">۱ روز پیش</span>
          </div>
          <div class="flex items-center justify-between p-3 bg-blue-50 rounded-lg">
            <div class="flex items-center">
              <span class="text-blue-600 text-lg ml-2">ℹ️</span>
              <span class="text-sm text-gray-700">کلمه کلیدی "لپ تاپ" بهبود یافت</span>
            </div>
            <span class="text-xs text-gray-500">۳ روز پیش</span>
          </div>
        </div>
      </div>

      <!-- جدول گزارش‌ها -->
      <div class="bg-white border border-gray-200 rounded-lg overflow-hidden">
        <div class="px-6 py-4 border-b border-gray-200">
          <h3 class="text-lg font-semibold text-gray-800">گزارش‌های اخیر</h3>
        </div>
        <table class="w-full">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع گزارش</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">امتیاز</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="report in reports" :key="report.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm font-medium text-gray-900">{{ report.type }}</div>
                <div class="text-sm text-gray-500">{{ report.description }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ report.date }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getStatusClass(report.status)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ report.status }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="w-16 bg-gray-200 rounded-full h-2 mr-2">
                    <div :class="getScoreClass(report.score)" :style="{ width: report.score + '%' }" class="h-2 rounded-full"></div>
                  </div>
                  <span class="text-sm text-gray-900">{{ report.score }}%</span>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">مشاهده</button>
                <button class="text-green-600 hover:text-green-900">دانلود</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
declare const useAuth: () => { user: any; hasPermission: (perm: string) => boolean }
</script>

<script setup lang="ts">
import { ref } from 'vue';

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user, hasPermission } = useAuth()

// داده‌های نمونه برای گزارش‌ها
const reports = ref([
  {
    id: 1,
    type: 'گزارش هفتگی سئو',
    description: 'تحلیل کامل وضعیت سئو سایت',
    date: '۱۴۰۲/۱۰/۱۵',
    status: 'تکمیل شده',
    score: 85
  },
  {
    id: 2,
    type: 'گزارش کلمات کلیدی',
    description: 'تحلیل رتبه‌بندی کلمات کلیدی',
    date: '۱۴۰۲/۱۰/۱۰',
    status: 'در حال پردازش',
    score: 72
  },
  {
    id: 3,
    type: 'گزارش عملکرد',
    description: 'تحلیل سرعت و عملکرد سایت',
    date: '۱۴۰۲/۱۰/۰۵',
    status: 'تکمیل شده',
    score: 90
  }
])

// تابع تعیین کلاس امتیاز
const getScoreClass = (score: number) => {
  if (score >= 80) return 'bg-green-500'
  if (score >= 60) return 'bg-yellow-500'
  return 'bg-red-500'
}

// تابع تعیین کلاس وضعیت
const getStatusClass = (status: string) => {
  switch (status) {
    case 'تکمیل شده':
      return 'bg-green-100 text-green-800'
    case 'در حال پردازش':
      return 'bg-yellow-100 text-yellow-800'
    case 'خطا':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}
</script> 

