<template>
  <div class="seo-image-optimization" dir="rtl">
    <div class="bg-white rounded-lg shadow-md p-6">
      <div class="flex items-center justify-between mb-6">
        <h1 class="text-2xl font-bold text-gray-800">بهینه‌سازی تصاویر</h1>
        <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors">
          اسکن تصاویر جدید
        </button>
      </div>

      <!-- آمار کلی -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
        <div class="bg-gradient-to-r from-blue-500 to-blue-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۵۰۰</div>
          <div class="text-sm">تصاویر کل</div>
        </div>
        <div class="bg-gradient-to-r from-green-500 to-green-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۳۵۰</div>
          <div class="text-sm">بهینه شده</div>
        </div>
        <div class="bg-gradient-to-r from-yellow-500 to-yellow-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۱۰۰</div>
          <div class="text-sm">نیازمند Alt</div>
        </div>
        <div class="bg-gradient-to-r from-red-500 to-red-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۵۰</div>
          <div class="text-sm">حجم بالا</div>
        </div>
      </div>

      <!-- جدول تصاویر -->
      <div class="bg-white border border-gray-200 rounded-lg overflow-hidden">
        <div class="px-6 py-4 border-b border-gray-200">
          <h3 class="text-lg font-semibold text-gray-800">تصاویر سایت</h3>
        </div>
        <table class="w-full">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تصویر</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">حجم</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">Alt Text</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="image in images" :key="image.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <img :src="image.url" :alt="image.alt" class="w-12 h-12 object-cover rounded">
                  <div class="mr-3">
                    <div class="text-sm font-medium text-gray-900">{{ image.name }}</div>
                    <div class="text-sm text-gray-500">{{ image.page }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ image.size }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ image.alt || 'بدون Alt' }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getStatusClass(image.status)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ image.status }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">ویرایش</button>
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

// داده‌های نمونه برای تصاویر
const images = ref([
  {
    id: 1,
    name: 'laptop-gaming.jpg',
    url: '/images/laptop-gaming.jpg',
    page: '/products/laptop-gaming',
    size: '۲۵۰ KB',
    alt: 'لپ تاپ گیمینگ',
    status: 'بهینه'
  },
  {
    id: 2,
    name: 'smartphone.jpg',
    url: '/images/smartphone.jpg',
    page: '/products/smartphone',
    size: '۵۰۰ KB',
    alt: '',
    status: 'نیازمند Alt'
  },
  {
    id: 3,
    name: 'headphone.jpg',
    url: '/images/headphone.jpg',
    page: '/products/headphone',
    size: '۱.۲ MB',
    alt: 'هدفون بی‌سیم',
    status: 'حجم بالا'
  }
])

// تابع تعیین کلاس وضعیت
const getStatusClass = (status: string) => {
  switch (status) {
    case 'بهینه':
      return 'bg-green-100 text-green-800'
    case 'نیازمند Alt':
      return 'bg-yellow-100 text-yellow-800'
    case 'حجم بالا':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}
</script> 
