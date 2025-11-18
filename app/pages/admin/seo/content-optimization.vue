<template>
  <div class="seo-content-optimization" dir="rtl">
    <div class="bg-white rounded-lg shadow-md p-6">
      <div class="flex items-center justify-between mb-6">
        <h1 class="text-2xl font-bold text-gray-800">بهینه‌سازی محتوا و متا</h1>
        <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors">
          تحلیل محتوای جدید
        </button>
      </div>

      <!-- آمار کلی -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
        <div class="bg-gradient-to-r from-blue-500 to-blue-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۱۵۰</div>
          <div class="text-sm">صفحات تحلیل شده</div>
        </div>
        <div class="bg-gradient-to-r from-green-500 to-green-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۸۵</div>
          <div class="text-sm">صفحات بهینه</div>
        </div>
        <div class="bg-gradient-to-r from-yellow-500 to-yellow-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۴۵</div>
          <div class="text-sm">نیازمند بهبود</div>
        </div>
        <div class="bg-gradient-to-r from-red-500 to-red-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۲۰</div>
          <div class="text-sm">مشکل جدی</div>
        </div>
      </div>

      <!-- ابزار تحلیل محتوا -->
      <div class="bg-gray-50 p-6 rounded-lg mb-6">
        <h2 class="text-lg font-semibold text-gray-800 mb-4">ابزار تحلیل محتوا</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">آدرس صفحه</label>
            <input type="url" placeholder="https://example.com/page" class="w-full border border-gray-300 rounded-lg px-3 py-2">
          </div>
          <div class="flex items-end">
            <button class="bg-green-600 text-white px-6 py-2 rounded-lg hover:bg-green-700 transition-colors">
              تحلیل محتوا
            </button>
          </div>
        </div>
      </div>

      <!-- نتایج تحلیل -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-6">
        <!-- تحلیل متا تگ‌ها -->
        <div class="bg-white border border-gray-200 rounded-lg p-6">
          <h3 class="text-lg font-semibold text-gray-800 mb-4">تحلیل متا تگ‌ها</h3>
          <div class="space-y-3">
            <div class="flex items-center justify-between p-3 bg-green-50 rounded-lg">
              <span class="text-sm text-gray-700">عنوان صفحه</span>
              <span class="text-green-600 text-sm">✅ بهینه</span>
            </div>
            <div class="flex items-center justify-between p-3 bg-yellow-50 rounded-lg">
              <span class="text-sm text-gray-700">توضیحات متا</span>
              <span class="text-yellow-600 text-sm">⚠️ نیاز به بهبود</span>
            </div>
            <div class="flex items-center justify-between p-3 bg-green-50 rounded-lg">
              <span class="text-sm text-gray-700">کلمات کلیدی</span>
              <span class="text-green-600 text-sm">✅ بهینه</span>
            </div>
            <div class="flex items-center justify-between p-3 bg-red-50 rounded-lg">
              <span class="text-sm text-gray-700">Open Graph</span>
              <span class="text-red-600 text-sm">❌ مشکل</span>
            </div>
          </div>
        </div>

        <!-- تحلیل محتوا -->
        <div class="bg-white border border-gray-200 rounded-lg p-6">
          <h3 class="text-lg font-semibold text-gray-800 mb-4">تحلیل محتوا</h3>
          <div class="space-y-3">
            <div class="flex items-center justify-between p-3 bg-green-50 rounded-lg">
              <span class="text-sm text-gray-700">طول محتوا</span>
              <span class="text-green-600 text-sm">✅ مناسب</span>
            </div>
            <div class="flex items-center justify-between p-3 bg-yellow-50 rounded-lg">
              <span class="text-sm text-gray-700">تراکم کلمات کلیدی</span>
              <span class="text-yellow-600 text-sm">⚠️ بالا</span>
            </div>
            <div class="flex items-center justify-between p-3 bg-green-50 rounded-lg">
              <span class="text-sm text-gray-700">ساختار هدینگ</span>
              <span class="text-green-600 text-sm">✅ مناسب</span>
            </div>
            <div class="flex items-center justify-between p-3 bg-red-50 rounded-lg">
              <span class="text-sm text-gray-700">تصاویر Alt</span>
              <span class="text-red-600 text-sm">❌ نامناسب</span>
            </div>
          </div>
        </div>
      </div>

      <!-- جدول صفحات -->
      <div class="bg-white border border-gray-200 rounded-lg overflow-hidden">
        <div class="px-6 py-4 border-b border-gray-200">
          <h3 class="text-lg font-semibold text-gray-800">صفحات اخیر</h3>
        </div>
        <table class="w-full">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">صفحه</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">امتیاز سئو</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">آخرین بررسی</th>
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
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ page.lastChecked }}
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
declare const definePageMeta: (meta: { layout?: string }) => void
</script>

<script setup lang="ts">
import { ref } from 'vue';

definePageMeta({
  layout: 'admin-main'
})

// داده‌های نمونه برای صفحات
const pages = ref([
  {
    id: 1,
    title: 'لپ تاپ گیمینگ',
    url: '/products/laptop-gaming',
    score: 85,
    status: 'بهینه',
    lastChecked: '۲ ساعت پیش'
  },
  {
    id: 2,
    title: 'گوشی هوشمند',
    url: '/products/smartphone',
    score: 72,
    status: 'نیازمند بهبود',
    lastChecked: '۱ روز پیش'
  },
  {
    id: 3,
    title: 'هدفون بی‌سیم',
    url: '/products/wireless-headphone',
    score: 45,
    status: 'مشکل جدی',
    lastChecked: '۳ روز پیش'
  },
  {
    id: 4,
    title: 'تبلت اندروید',
    url: '/products/android-tablet',
    score: 90,
    status: 'بهینه',
    lastChecked: '۱ هفته پیش'
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
    case 'مشکل جدی':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}
</script> 
