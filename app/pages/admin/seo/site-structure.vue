<template>
  <div class="seo-site-structure" dir="rtl">
    <div class="bg-white rounded-lg shadow-md p-6">
      <div class="flex items-center justify-between mb-6">
        <h1 class="text-2xl font-bold text-gray-800">ساختار و ناوبری سایت</h1>
        <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors">
          تولید نقشه سایت
        </button>
      </div>

      <!-- آمار کلی -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
        <div class="bg-gradient-to-r from-blue-500 to-blue-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۵۰۰</div>
          <div class="text-sm">صفحات ایندکس شده</div>
        </div>
        <div class="bg-gradient-to-r from-green-500 to-green-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۸</div>
          <div class="text-sm">سطح عمق</div>
        </div>
        <div class="bg-gradient-to-r from-yellow-500 to-yellow-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۲۵</div>
          <div class="text-sm">لینک‌های شکسته</div>
        </div>
        <div class="bg-gradient-to-r from-red-500 to-red-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۱۲</div>
          <div class="text-sm">صفحات یتیم</div>
        </div>
      </div>

      <!-- ساختار سایت -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-6">
        <!-- نقشه سایت -->
        <div class="bg-white border border-gray-200 rounded-lg p-6">
          <h3 class="text-lg font-semibold text-gray-800 mb-4">نقشه سایت</h3>
          <div class="space-y-2">
            <div class="flex items-center justify-between p-2 bg-gray-50 rounded">
              <span class="text-sm text-gray-700">XML Sitemap</span>
              <span class="text-green-600 text-sm">✅ فعال</span>
            </div>
            <div class="flex items-center justify-between p-2 bg-gray-50 rounded">
              <span class="text-sm text-gray-700">HTML Sitemap</span>
              <span class="text-yellow-600 text-sm">⚠️ نیاز به بهبود</span>
            </div>
            <div class="flex items-center justify-between p-2 bg-gray-50 rounded">
              <span class="text-sm text-gray-700">آخرین بروزرسانی</span>
              <span class="text-sm text-gray-600">۲ روز پیش</span>
            </div>
          </div>
        </div>

        <!-- ساختار URL -->
        <div class="bg-white border border-gray-200 rounded-lg p-6">
          <h3 class="text-lg font-semibold text-gray-800 mb-4">ساختار URL</h3>
          <div class="space-y-2">
            <div class="flex items-center justify-between p-2 bg-gray-50 rounded">
              <span class="text-sm text-gray-700">طول URL</span>
              <span class="text-green-600 text-sm">✅ مناسب</span>
            </div>
            <div class="flex items-center justify-between p-2 bg-gray-50 rounded">
              <span class="text-sm text-gray-700">کاراکترهای خاص</span>
              <span class="text-green-600 text-sm">✅ مناسب</span>
            </div>
            <div class="flex items-center justify-between p-2 bg-gray-50 rounded">
              <span class="text-sm text-gray-700">کلمات کلیدی</span>
              <span class="text-yellow-600 text-sm">⚠️ نیاز به بهبود</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Breadcrumb -->
      <div class="bg-white border border-gray-200 rounded-lg p-6 mb-6">
        <h3 class="text-lg font-semibold text-gray-800 mb-4">Breadcrumb Navigation</h3>
        <div class="space-y-3">
          <div class="flex items-center justify-between p-3 bg-green-50 rounded-lg">
            <span class="text-sm text-gray-700">پیاده‌سازی Breadcrumb</span>
            <span class="text-green-600 text-sm">✅ فعال</span>
          </div>
          <div class="flex items-center justify-between p-3 bg-yellow-50 rounded-lg">
            <span class="text-sm text-gray-700">Schema Markup</span>
            <span class="text-yellow-600 text-sm">⚠️ نیاز به بهبود</span>
          </div>
          <div class="flex items-center justify-between p-3 bg-green-50 rounded-lg">
            <span class="text-sm text-gray-700">نمایش در موبایل</span>
            <span class="text-green-600 text-sm">✅ مناسب</span>
          </div>
        </div>
      </div>

      <!-- جدول صفحات -->
      <div class="bg-white border border-gray-200 rounded-lg overflow-hidden">
        <div class="px-6 py-4 border-b border-gray-200">
          <h3 class="text-lg font-semibold text-gray-800">ساختار صفحات</h3>
        </div>
        <table class="w-full">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">صفحه</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">سطح</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">لینک‌های داخلی</th>
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
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="text-sm text-gray-900">{{ page.level }}</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="text-sm text-gray-900">{{ page.internalLinks }}</span>
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

// تعریف interface برای صفحه
interface Page {
  id: number
  title: string
  url: string
  level: number
  internalLinks: number
  status: string
}

// داده‌های نمونه برای صفحات
const pages = ref<Page[]>([
  {
    id: 1,
    title: 'صفحه اصلی',
    url: '/',
    level: 1,
    internalLinks: 25,
    status: 'بهینه'
  },
  {
    id: 2,
    title: 'محصولات',
    url: '/products',
    level: 2,
    internalLinks: 15,
    status: 'بهینه'
  },
  {
    id: 3,
    title: 'لپ تاپ گیمینگ',
    url: '/products/laptop-gaming',
    level: 3,
    internalLinks: 8,
    status: 'نیازمند بهبود'
  },
  {
    id: 4,
    title: 'گوشی هوشمند',
    url: '/products/smartphone',
    level: 3,
    internalLinks: 12,
    status: 'بهینه'
  }
])

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

