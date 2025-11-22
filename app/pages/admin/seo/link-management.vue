<template>
  <div class="seo-link-management" dir="rtl">
    <div class="bg-white rounded-lg shadow-md p-6">
      <div class="flex items-center justify-between mb-6">
        <h1 class="text-2xl font-bold text-gray-800">مدیریت لینک‌ها</h1>
        <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors">
          اسکن لینک‌های جدید
        </button>
      </div>

      <!-- آمار کلی -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
        <div class="bg-gradient-to-r from-blue-500 to-blue-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۱,۲۵۰</div>
          <div class="text-sm">لینک‌های داخلی</div>
        </div>
        <div class="bg-gradient-to-r from-green-500 to-green-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۸۵</div>
          <div class="text-sm">لینک‌های خارجی</div>
        </div>
        <div class="bg-gradient-to-r from-yellow-500 to-yellow-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۲۵</div>
          <div class="text-sm">لینک‌های شکسته</div>
        </div>
        <div class="bg-gradient-to-r from-red-500 to-red-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۱۲</div>
          <div class="text-sm">لینک‌های مشکوک</div>
        </div>
      </div>

      <!-- جدول لینک‌ها -->
      <div class="bg-white border border-gray-200 rounded-lg overflow-hidden">
        <div class="px-6 py-4 border-b border-gray-200">
          <h3 class="text-lg font-semibold text-gray-800">لینک‌های سایت</h3>
        </div>
        <table class="w-full">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">لینک</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">Anchor Text</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="link in links" :key="link.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm font-medium text-gray-900">{{ link.url }}</div>
                <div class="text-sm text-gray-500">{{ link.source }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getTypeClass(link.type)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ link.type }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getStatusClass(link.status)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ link.status }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ link.anchorText }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">مشاهده</button>
                <button class="text-red-600 hover:text-red-900">حذف</button>
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

// داده‌های نمونه برای لینک‌ها
const links = ref([
  {
    id: 1,
    url: '/products/laptop-gaming',
    source: '/products',
    type: 'داخلی',
    status: 'فعال',
    anchorText: 'لپ تاپ گیمینگ'
  },
  {
    id: 2,
    url: 'https://example.com',
    source: '/about',
    type: 'خارجی',
    status: 'فعال',
    anchorText: 'مثال'
  },
  {
    id: 3,
    url: '/old-page',
    source: '/products',
    type: 'داخلی',
    status: 'شکسته',
    anchorText: 'صفحه قدیمی'
  }
])

// تابع تعیین کلاس نوع
const getTypeClass = (type: string) => {
  switch (type) {
    case 'داخلی':
      return 'bg-blue-100 text-blue-800'
    case 'خارجی':
      return 'bg-green-100 text-green-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

// تابع تعیین کلاس وضعیت
const getStatusClass = (status: string) => {
  switch (status) {
    case 'فعال':
      return 'bg-green-100 text-green-800'
    case 'شکسته':
      return 'bg-red-100 text-red-800'
    case 'مشکوک':
      return 'bg-yellow-100 text-yellow-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}
</script> 
