<template>
  <div class="seo-performance" dir="rtl">
    <div class="bg-white rounded-lg shadow-md p-6">
      <div class="flex items-center justify-between mb-6">
        <h1 class="text-2xl font-bold text-gray-800">عملکرد و سرعت</h1>
        <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors">
          تست سرعت جدید
        </button>
      </div>

      <!-- آمار کلی -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
        <div class="bg-gradient-to-r from-blue-500 to-blue-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۲.۳s</div>
          <div class="text-sm">زمان بارگذاری</div>
        </div>
        <div class="bg-gradient-to-r from-green-500 to-green-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۸۵</div>
          <div class="text-sm">امتیاز PageSpeed</div>
        </div>
        <div class="bg-gradient-to-r from-yellow-500 to-yellow-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۲.۱MB</div>
          <div class="text-sm">حجم صفحه</div>
        </div>
        <div class="bg-gradient-to-r from-red-500 to-red-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۱۵</div>
          <div class="text-sm">درخواست‌ها</div>
        </div>
      </div>

      <!-- Core Web Vitals -->
      <div class="bg-white border border-gray-200 rounded-lg p-6 mb-6">
        <h3 class="text-lg font-semibold text-gray-800 mb-4">Core Web Vitals</h3>
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div class="text-center p-6 bg-green-50 rounded-lg">
            <div class="text-2xl font-bold text-green-600">۱.۲s</div>
            <div class="text-sm text-gray-600">LCP</div>
            <div class="text-xs text-green-600">خوب</div>
          </div>
          <div class="text-center p-6 bg-yellow-50 rounded-lg">
            <div class="text-2xl font-bold text-yellow-600">۱۵۰ms</div>
            <div class="text-sm text-gray-600">FID</div>
            <div class="text-xs text-yellow-600">نیاز به بهبود</div>
          </div>
          <div class="text-center p-6 bg-green-50 rounded-lg">
            <div class="text-2xl font-bold text-green-600">۰.۰۵</div>
            <div class="text-sm text-gray-600">CLS</div>
            <div class="text-xs text-green-600">خوب</div>
          </div>
        </div>
      </div>

      <!-- جدول صفحات -->
      <div class="bg-white border border-gray-200 rounded-lg overflow-hidden">
        <div class="px-6 py-4 border-b border-gray-200">
          <h3 class="text-lg font-semibold text-gray-800">عملکرد صفحات</h3>
        </div>
        <table class="w-full">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">صفحه</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">زمان بارگذاری</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">امتیاز</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="page in pages" :key="page.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm font-medium text-gray-900">{{ page.title }}</div>
                <div class="text-sm text-gray-500">{{ page.url }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ page.loadTime }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="w-16 bg-gray-200 rounded-full h-2 mr-2">
                    <div :class="getScoreClass(page.score)" :style="{ width: page.score + '%' }" class="h-2 rounded-full"></div>
                  </div>
                  <span class="text-sm text-gray-900">{{ page.score }}%</span>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getStatusClass(page.status)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ page.status }}
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
    loadTime: '۲.۱s',
    score: 85,
    status: 'خوب'
  },
  {
    id: 2,
    title: 'محصولات',
    url: '/products',
    loadTime: '۳.۲s',
    score: 65,
    status: 'نیازمند بهبود'
  },
  {
    id: 3,
    title: 'لپ تاپ گیمینگ',
    url: '/products/laptop-gaming',
    loadTime: '۴.۵s',
    score: 45,
    status: 'ضعیف'
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
    case 'خوب':
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

