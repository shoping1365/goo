<template>
  <div class="seo-page-analysis" dir="rtl">
    <div class="bg-white rounded-lg shadow-md p-6">
      <div class="flex items-center justify-between mb-6">
        <h1 class="text-2xl font-bold text-gray-800">آنالیز صفحات</h1>
        <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors">
          آنالیز صفحه جدید
        </button>
      </div>

      <!-- آمار کلی -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
        <div class="bg-gradient-to-r from-blue-500 to-blue-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۲۵۰</div>
          <div class="text-sm">صفحات آنالیز شده</div>
        </div>
        <div class="bg-gradient-to-r from-green-500 to-green-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۱۸۵</div>
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

      <!-- ابزار آنالیز -->
      <div class="bg-gray-50 p-6 rounded-lg mb-6">
        <h2 class="text-lg font-semibold text-gray-800 mb-4">ابزار آنالیز صفحه</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">آدرس صفحه</label>
            <input type="url" placeholder="https://example.com/page" class="w-full border border-gray-300 rounded-lg px-3 py-2">
          </div>
          <div class="flex items-end">
            <button class="bg-green-600 text-white px-6 py-2 rounded-lg hover:bg-green-700 transition-colors">
              آنالیز صفحه
            </button>
          </div>
        </div>
      </div>

      <!-- نتایج آنالیز -->
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-6 mb-6">
        <!-- امتیاز کلی -->
        <div class="bg-white border border-gray-200 rounded-lg p-6">
          <h3 class="text-lg font-semibold text-gray-800 mb-4">امتیاز کلی</h3>
          <div class="text-center">
            <div class="text-4xl font-bold text-green-600 mb-2">۸۵</div>
            <div class="text-sm text-gray-600">از ۱۰۰</div>
            <div class="mt-4">
              <div class="w-full bg-gray-200 rounded-full h-3">
                <div class="bg-green-500 h-3 rounded-full" style="width: 85%"></div>
              </div>
            </div>
          </div>
        </div>

        <!-- مشکلات شناسایی شده -->
        <div class="bg-white border border-gray-200 rounded-lg p-6">
          <h3 class="text-lg font-semibold text-gray-800 mb-4">مشکلات شناسایی شده</h3>
          <div class="space-y-2">
            <div class="flex items-center justify-between p-2 bg-red-50 rounded">
              <span class="text-sm text-gray-700">عنوان تکراری</span>
              <span class="text-red-600 text-sm">❌</span>
            </div>
            <div class="flex items-center justify-between p-2 bg-yellow-50 rounded">
              <span class="text-sm text-gray-700">Meta Description کوتاه</span>
              <span class="text-yellow-600 text-sm">⚠️</span>
            </div>
            <div class="flex items-center justify-between p-2 bg-green-50 rounded">
              <span class="text-sm text-gray-700">ساختار هدینگ</span>
              <span class="text-green-600 text-sm">✅</span>
            </div>
          </div>
        </div>

        <!-- پیشنهادات بهبود -->
        <div class="bg-white border border-gray-200 rounded-lg p-6">
          <h3 class="text-lg font-semibold text-gray-800 mb-4">پیشنهادات بهبود</h3>
          <div class="space-y-2">
            <div class="flex items-center justify-between p-2 bg-blue-50 rounded">
              <span class="text-sm text-gray-700">افزودن Schema Markup</span>
              <span class="text-blue-600 text-sm">💡</span>
            </div>
            <div class="flex items-center justify-between p-2 bg-blue-50 rounded">
              <span class="text-sm text-gray-700">بهینه‌سازی تصاویر</span>
              <span class="text-blue-600 text-sm">💡</span>
            </div>
            <div class="flex items-center justify-between p-2 bg-blue-50 rounded">
              <span class="text-sm text-gray-700">افزودن لینک‌های داخلی</span>
              <span class="text-blue-600 text-sm">💡</span>
            </div>
          </div>
        </div>
      </div>

      <!-- تحلیل جزئیات -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-6">
        <!-- تحلیل تکنیکی -->
        <div class="bg-white border border-gray-200 rounded-lg p-6">
          <h3 class="text-lg font-semibold text-gray-800 mb-4">تحلیل تکنیکی</h3>
          <div class="space-y-3">
            <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
              <span class="text-sm text-gray-700">عنوان صفحه</span>
              <span class="text-green-600 text-sm">✅ ۶۰ کاراکتر</span>
            </div>
            <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
              <span class="text-sm text-gray-700">Meta Description</span>
              <span class="text-yellow-600 text-sm">⚠️ ۱۲۰ کاراکتر</span>
            </div>
            <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
              <span class="text-sm text-gray-700">کلمات کلیدی</span>
              <span class="text-green-600 text-sm">✅ ۲.۵% تراکم</span>
            </div>
            <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
              <span class="text-sm text-gray-700">تصاویر Alt</span>
              <span class="text-red-600 text-sm">❌ ۳ تصویر بدون Alt</span>
            </div>
            <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
              <span class="text-sm text-gray-700">لینک‌های داخلی</span>
              <span class="text-green-600 text-sm">✅ ۸ لینک</span>
            </div>
          </div>
        </div>

        <!-- تحلیل محتوا -->
        <div class="bg-white border border-gray-200 rounded-lg p-6">
          <h3 class="text-lg font-semibold text-gray-800 mb-4">تحلیل محتوا</h3>
          <div class="space-y-3">
            <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
              <span class="text-sm text-gray-700">طول محتوا</span>
              <span class="text-green-600 text-sm">✅ ۱۵۰۰ کلمه</span>
            </div>
            <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
              <span class="text-sm text-gray-700">ساختار هدینگ</span>
              <span class="text-green-600 text-sm">✅ H1, H2, H3</span>
            </div>
            <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
              <span class="text-sm text-gray-700">خوانایی</span>
              <span class="text-green-600 text-sm">✅ ۷۵%</span>
            </div>
            <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
              <span class="text-sm text-gray-700">لیست‌ها</span>
              <span class="text-green-600 text-sm">✅ ۳ لیست</span>
            </div>
            <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
              <span class="text-sm text-gray-700">لینک‌های خارجی</span>
              <span class="text-yellow-600 text-sm">⚠️ ۲ لینک</span>
            </div>
          </div>
        </div>
      </div>

      <!-- جدول صفحات آنالیز شده -->
      <div class="bg-white border border-gray-200 rounded-lg overflow-hidden">
        <div class="px-6 py-4 border-b border-gray-200">
          <h3 class="text-lg font-semibold text-gray-800">صفحات آنالیز شده</h3>
        </div>
        <table class="w-full">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">صفحه</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">امتیاز سئو</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مشکلات</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">آخرین آنالیز</th>
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
                <span :class="getIssuesClass(page.issues)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ page.issues }} مشکل
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ page.lastAnalyzed }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">مشاهده</button>
                <button class="text-green-600 hover:text-green-900">آنالیز مجدد</button>
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

// استفاده از useAuth برای چک کردن پرمیژن‌ها
// const { user, hasPermission } = useAuth()

// داده‌های نمونه برای صفحات
const pages = ref([
  {
    id: 1,
    title: 'صفحه اصلی',
    url: '/',
    score: 85,
    issues: 2,
    lastAnalyzed: '۲ ساعت پیش'
  },
  {
    id: 2,
    title: 'محصولات',
    url: '/products',
    score: 72,
    issues: 5,
    lastAnalyzed: '۱ روز پیش'
  },
  {
    id: 3,
    title: 'لپ تاپ گیمینگ',
    url: '/products/laptop-gaming',
    score: 95,
    issues: 0,
    lastAnalyzed: '۳ روز پیش'
  },
  {
    id: 4,
    title: 'درباره ما',
    url: '/about',
    score: 68,
    issues: 8,
    lastAnalyzed: '۱ هفته پیش'
  }
])

// تابع تعیین کلاس امتیاز
const getScoreClass = (score: number) => {
  if (score >= 80) return 'bg-green-500'
  if (score >= 60) return 'bg-yellow-500'
  return 'bg-red-500'
}

// تابع تعیین کلاس مشکلات
const getIssuesClass = (issues: number) => {
  if (issues === 0) return 'bg-green-100 text-green-800'
  if (issues <= 3) return 'bg-yellow-100 text-yellow-800'
  return 'bg-red-100 text-red-800'
}
</script> 

