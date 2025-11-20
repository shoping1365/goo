<template>
  <div class="seo-keyword-tracking" dir="rtl">
    <div class="bg-white rounded-lg shadow-md p-6">
      <div class="flex items-center justify-between mb-6">
        <h1 class="text-2xl font-bold text-gray-800">تحلیل و ردیابی کلمات کلیدی</h1>
        <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors">
          افزودن کلمه کلیدی جدید
        </button>
      </div>

      <!-- آمار کلی -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
        <div class="bg-gradient-to-r from-blue-500 to-blue-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۲۵۰</div>
          <div class="text-sm">کلمات کلیدی ردیابی شده</div>
        </div>
        <div class="bg-gradient-to-r from-green-500 to-green-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۴۵</div>
          <div class="text-sm">کلمات کلیدی در صفحه اول</div>
        </div>
        <div class="bg-gradient-to-r from-yellow-500 to-yellow-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۱۲</div>
          <div class="text-sm">کلمات کلیدی در حال بهبود</div>
        </div>
        <div class="bg-gradient-to-r from-red-500 to-red-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۸</div>
          <div class="text-sm">کلمات کلیدی نیازمند توجه</div>
        </div>
      </div>

      <!-- فیلترها -->
      <div class="bg-gray-50 p-6 rounded-lg mb-6">
        <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
            <select class="w-full border border-gray-300 rounded-lg px-3 py-2">
              <option>همه</option>
              <option>صفحه اول</option>
              <option>در حال بهبود</option>
              <option>نیازمند توجه</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">اولویت</label>
            <select class="w-full border border-gray-300 rounded-lg px-3 py-2">
              <option>همه</option>
              <option>بالا</option>
              <option>متوسط</option>
              <option>پایین</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">جستجو</label>
            <input type="text" placeholder="جستجو در کلمات کلیدی..." class="w-full border border-gray-300 rounded-lg px-3 py-2">
          </div>
          <div class="flex items-end">
            <button class="bg-gray-600 text-white px-4 py-2 rounded-lg hover:bg-gray-700 transition-colors">
              اعمال فیلتر
            </button>
          </div>
        </div>
      </div>

      <!-- جدول کلمات کلیدی -->
      <div class="bg-white border border-gray-200 rounded-lg overflow-hidden">
        <table class="w-full">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کلمه کلیدی</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">رتبه فعلی</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تغییر</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">حجم جستجو</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">اولویت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="keyword in keywords" :key="keyword.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm font-medium text-gray-900">{{ keyword.term }}</div>
                <div class="text-sm text-gray-500">{{ keyword.url }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getRankClass(keyword.rank)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ keyword.rank }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <span v-if="keyword.change > 0" class="text-green-600 text-sm">+{{ keyword.change }}</span>
                  <span v-else-if="keyword.change < 0" class="text-red-600 text-sm">{{ keyword.change }}</span>
                  <span v-else class="text-gray-500 text-sm">-</span>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ keyword.volume }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getPriorityClass(keyword.priority)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ keyword.priority }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">ویرایش</button>
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

// داده‌های نمونه برای کلمات کلیدی
const keywords = ref([
  {
    id: 1,
    term: 'لپ تاپ گیمینگ',
    url: '/products/laptop-gaming',
    rank: 3,
    change: 2,
    volume: '۱۲,۵۰۰',
    priority: 'بالا'
  },
  {
    id: 2,
    term: 'گوشی هوشمند',
    url: '/products/smartphone',
    rank: 8,
    change: -1,
    volume: '۸,۲۰۰',
    priority: 'بالا'
  },
  {
    id: 3,
    term: 'هدفون بی‌سیم',
    url: '/products/wireless-headphone',
    rank: 15,
    change: 0,
    volume: '۳,۵۰۰',
    priority: 'متوسط'
  },
  {
    id: 4,
    term: 'تبلت اندروید',
    url: '/products/android-tablet',
    rank: 25,
    change: 5,
    volume: '۲,۱۰۰',
    priority: 'متوسط'
  }
])

// تابع تعیین کلاس رتبه
const getRankClass = (rank: number) => {
  if (rank <= 3) return 'bg-green-100 text-green-800'
  if (rank <= 10) return 'bg-blue-100 text-blue-800'
  if (rank <= 20) return 'bg-yellow-100 text-yellow-800'
  return 'bg-red-100 text-red-800'
}

// تابع تعیین کلاس اولویت
const getPriorityClass = (priority: string) => {
  switch (priority) {
    case 'بالا':
      return 'bg-red-100 text-red-800'
    case 'متوسط':
      return 'bg-yellow-100 text-yellow-800'
    case 'پایین':
      return 'bg-green-100 text-green-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}
</script> 

