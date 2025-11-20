<template>
  <div class="seo-user-experience" dir="rtl">
    <div class="bg-white rounded-lg shadow-md p-6">
      <div class="flex items-center justify-between mb-6">
        <h1 class="text-2xl font-bold text-gray-800">تجربه کاربری و ریسپانسیو</h1>
        <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors">
          تست تجربه کاربری
        </button>
      </div>

      <!-- آمار کلی -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
        <div class="bg-gradient-to-r from-blue-500 to-blue-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۹۲</div>
          <div class="text-sm">امتیاز UX</div>
        </div>
        <div class="bg-gradient-to-r from-green-500 to-green-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۱۰۰</div>
          <div class="text-sm">ریسپانسیو</div>
        </div>
        <div class="bg-gradient-to-r from-yellow-500 to-yellow-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۸۵</div>
          <div class="text-sm">دسترسی‌پذیری</div>
        </div>
        <div class="bg-gradient-to-r from-red-500 to-red-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۵</div>
          <div class="text-sm">مشکلات UX</div>
        </div>
      </div>

      <!-- تست دستگاه‌ها -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-6">
        <!-- تست ریسپانسیو -->
        <div class="bg-white border border-gray-200 rounded-lg p-6">
          <h3 class="text-lg font-semibold text-gray-800 mb-4">تست ریسپانسیو</h3>
          <div class="space-y-3">
            <div class="flex items-center justify-between p-3 bg-green-50 rounded-lg">
              <span class="text-sm text-gray-700">دسکتاپ (1920px)</span>
              <span class="text-green-600 text-sm">✅ مناسب</span>
            </div>
            <div class="flex items-center justify-between p-3 bg-green-50 rounded-lg">
              <span class="text-sm text-gray-700">تبلت (768px)</span>
              <span class="text-green-600 text-sm">✅ مناسب</span>
            </div>
            <div class="flex items-center justify-between p-3 bg-yellow-50 rounded-lg">
              <span class="text-sm text-gray-700">موبایل (375px)</span>
              <span class="text-yellow-600 text-sm">⚠️ نیاز به بهبود</span>
            </div>
          </div>
        </div>

        <!-- دسترسی‌پذیری -->
        <div class="bg-white border border-gray-200 rounded-lg p-6">
          <h3 class="text-lg font-semibold text-gray-800 mb-4">دسترسی‌پذیری</h3>
          <div class="space-y-3">
            <div class="flex items-center justify-between p-3 bg-green-50 rounded-lg">
              <span class="text-sm text-gray-700">کنتراست رنگ</span>
              <span class="text-green-600 text-sm">✅ مناسب</span>
            </div>
            <div class="flex items-center justify-between p-3 bg-green-50 rounded-lg">
              <span class="text-sm text-gray-700">Alt Text تصاویر</span>
              <span class="text-green-600 text-sm">✅ مناسب</span>
            </div>
            <div class="flex items-center justify-between p-3 bg-yellow-50 rounded-lg">
              <span class="text-sm text-gray-700">ناوبری کیبورد</span>
              <span class="text-yellow-600 text-sm">⚠️ نیاز به بهبود</span>
            </div>
          </div>
        </div>
      </div>

      <!-- جدول صفحات -->
      <div class="bg-white border border-gray-200 rounded-lg overflow-hidden">
        <div class="px-6 py-4 border-b border-gray-200">
          <h3 class="text-lg font-semibold text-gray-800">وضعیت UX صفحات</h3>
        </div>
        <table class="w-full">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">صفحه</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">امتیاز UX</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">ریسپانسیو</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">دسترسی‌پذیری</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="page in pages" :key="page.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm font-medium text-gray-900">{{ page.title }}</div>
                <div class="text-sm text-gray-500">{{ page.url }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="w-16 bg-gray-200 rounded-full h-2 mr-2">
                    <div :class="getScoreClass(page.uxScore)" :style="{ width: page.uxScore + '%' }" class="h-2 rounded-full"></div>
                  </div>
                  <span class="text-sm text-gray-900">{{ page.uxScore }}%</span>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getStatusClass(page.responsive)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ page.responsive }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getStatusClass(page.accessibility)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ page.accessibility }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">مشاهده</button>
                <button class="text-green-600 hover:text-green-900">بهینه‌سازی</button>
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
// declare const useAuth: () => { user: unknown; hasPermission: (perm: string) => boolean }
</script>

<script setup lang="ts">
import { ref } from 'vue';

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
// const { user, hasPermission } = useAuth()

// داده‌های نمونه برای صفحات
const pages = ref([
  {
    id: 1,
    title: 'صفحه اصلی',
    url: '/',
    uxScore: 92,
    responsive: 'بهینه',
    accessibility: 'بهینه'
  },
  {
    id: 2,
    title: 'محصولات',
    url: '/products',
    uxScore: 85,
    responsive: 'بهینه',
    accessibility: 'نیازمند بهبود'
  },
  {
    id: 3,
    title: 'لپ تاپ گیمینگ',
    url: '/products/laptop-gaming',
    uxScore: 78,
    responsive: 'نیازمند بهبود',
    accessibility: 'بهینه'
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
    case 'بهینه':
      return 'bg-green-100 text-green-800'
    case 'نیازمند بهبود':
      return 'bg-yellow-100 text-yellow-800'
    case 'ضعیف':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}
</script> 

