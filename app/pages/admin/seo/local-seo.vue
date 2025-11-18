<template>
  <div class="seo-local-seo" dir="rtl">
    <div class="bg-white rounded-lg shadow-md p-6">
      <div class="flex items-center justify-between mb-6">
        <h1 class="text-2xl font-bold text-gray-800">سئو محلی</h1>
        <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors">
          افزودن کسب و کار جدید
        </button>
      </div>

      <!-- آمار کلی -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
        <div class="bg-gradient-to-r from-blue-500 to-blue-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۵</div>
          <div class="text-sm">کسب و کارها</div>
        </div>
        <div class="bg-gradient-to-r from-green-500 to-green-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۸۵</div>
          <div class="text-sm">امتیاز کلی</div>
        </div>
        <div class="bg-gradient-to-r from-yellow-500 to-yellow-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۲۵۰</div>
          <div class="text-sm">نظرات</div>
        </div>
        <div class="bg-gradient-to-r from-red-500 to-red-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۱۲</div>
          <div class="text-sm">نیازمند بررسی</div>
        </div>
      </div>

      <!-- جدول کسب و کارها -->
      <div class="bg-white border border-gray-200 rounded-lg overflow-hidden">
        <div class="px-6 py-4 border-b border-gray-200">
          <h3 class="text-lg font-semibold text-gray-800">کسب و کارهای محلی</h3>
        </div>
        <table class="w-full">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نام کسب و کار</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">آدرس</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">امتیاز</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="business in businesses" :key="business.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm font-medium text-gray-900">{{ business.name }}</div>
                <div class="text-sm text-gray-500">{{ business.category }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ business.address }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="w-16 bg-gray-200 rounded-full h-2 mr-2">
                    <div :class="getScoreClass(business.score)" :style="{ width: business.score + '%' }" class="h-2 rounded-full"></div>
                  </div>
                  <span class="text-sm text-gray-900">{{ business.score }}%</span>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getStatusClass(business.status)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ business.status }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">مشاهده</button>
                <button class="text-green-600 hover:text-green-900">ویرایش</button>
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

// داده‌های نمونه برای کسب و کارها
const businesses = ref([
  {
    id: 1,
    name: 'فروشگاه الکترونیک تهران',
    category: 'فروشگاه الکترونیک',
    address: 'تهران، خیابان ولیعصر',
    score: 85,
    status: 'فعال'
  },
  {
    id: 2,
    name: 'فروشگاه الکترونیک اصفهان',
    category: 'فروشگاه الکترونیک',
    address: 'اصفهان، خیابان چهارباغ',
    score: 72,
    status: 'نیازمند بررسی'
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
    case 'فعال':
      return 'bg-green-100 text-green-800'
    case 'نیازمند بررسی':
      return 'bg-yellow-100 text-yellow-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}
</script> 

