<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200">
    <!-- هدر بخش -->
    <div class="p-6 border-b border-gray-200">
      <div class="flex items-center justify-between">
        <div>
          <h2 class="text-lg font-semibold text-gray-900">تحلیل و بهینه‌سازی کمپین‌ها</h2>
          <p class="text-gray-600 mt-1">بررسی عملکرد، شاخص‌های کلیدی و پیشنهادات بهبود</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <input v-model="filters.dateRange" type="month" class="px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500" />
        </div>
      </div>
    </div>

    <!-- شاخص‌های کلیدی KPI -->
    <div class="p-6 border-b border-gray-200">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div class="bg-blue-50 p-6 rounded-lg">
          <div class="flex items-center">
            <div class="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center ml-3">
              <span class="i-heroicons-ticket text-blue-600 text-xl"></span>
            </div>
            <div>
              <p class="text-sm text-blue-600">تعداد استفاده</p>
              <p class="text-2xl font-bold text-blue-900">{{ kpi.usageCount }}</p>
            </div>
          </div>
        </div>
        <div class="bg-green-50 p-6 rounded-lg">
          <div class="flex items-center">
            <div class="w-12 h-12 bg-green-100 rounded-lg flex items-center justify-center ml-3">
              <span class="i-heroicons-chart-bar text-green-600 text-xl"></span>
            </div>
            <div>
              <p class="text-sm text-green-600">نرخ تبدیل</p>
              <p class="text-2xl font-bold text-green-900">{{ kpi.conversionRate }}%</p>
            </div>
          </div>
        </div>
        <div class="bg-purple-50 p-6 rounded-lg">
          <div class="flex items-center">
            <div class="w-12 h-12 bg-purple-100 rounded-lg flex items-center justify-center ml-3">
              <span class="i-heroicons-currency-dollar text-purple-600 text-xl"></span>
            </div>
            <div>
              <p class="text-sm text-purple-600">درآمد</p>
              <p class="text-2xl font-bold text-purple-900">{{ formatPrice(kpi.revenue) }}</p>
            </div>
          </div>
        </div>
        <div class="bg-orange-50 p-6 rounded-lg">
          <div class="flex items-center">
            <div class="w-12 h-12 bg-orange-100 rounded-lg flex items-center justify-center ml-3">
              <span class="i-heroicons-banknotes text-orange-600 text-xl"></span>
            </div>
            <div>
              <p class="text-sm text-orange-600">مصرف بودجه</p>
              <p class="text-2xl font-bold text-orange-900">{{ formatPrice(kpi.spentBudget) }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- نمودار عملکرد کمپین‌ها -->
    <div class="p-6 border-b border-gray-200">
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <div class="bg-white border border-gray-200 rounded-lg p-6">
          <h4 class="text-md font-medium text-gray-900 mb-4">روند استفاده از کوپن‌ها</h4>
          <!-- نمودار خطی نمونه -->
          <div class="h-48 flex items-center justify-center text-gray-400">[نمودار خطی - استفاده روزانه]</div>
        </div>
        <div class="bg-white border border-gray-200 rounded-lg p-6">
          <h4 class="text-md font-medium text-gray-900 mb-4">درآمد و مصرف بودجه</h4>
          <!-- نمودار میله‌ای نمونه -->
          <div class="h-48 flex items-center justify-center text-gray-400">[نمودار میله‌ای - درآمد/بودجه]</div>
        </div>
      </div>
    </div>

    <!-- جدول مقایسه کمپین‌ها -->
    <div class="p-6 border-b border-gray-200">
      <h4 class="text-md font-medium text-gray-900 mb-4">مقایسه کمپین‌ها</h4>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نام کمپین</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تعداد استفاده</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نرخ تبدیل</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">درآمد</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مصرف بودجه</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">درصد مصرف بودجه</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="row in campaignRows" :key="row.id">
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ row.name }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ row.usageCount }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ row.conversionRate }}%</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ formatPrice(row.revenue) }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ formatPrice(row.spentBudget) }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                <div class="flex items-center">
                  <div class="w-20 bg-gray-200 rounded-full h-2 ml-2">
                    <div class="h-2 rounded-full" :class="row.percentUsed >= 90 ? 'bg-orange-500' : 'bg-blue-600'" :style="{ width: row.percentUsed + '%' }"></div>
                  </div>
                  <span>{{ row.percentUsed }}%</span>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- پیشنهادات بهینه‌سازی -->
    <div class="p-6">
      <h4 class="text-md font-medium text-gray-900 mb-4">پیشنهادات بهبود و بهینه‌سازی</h4>
      <ul class="list-disc pr-6 space-y-2 text-gray-700">
        <li v-for="suggestion in suggestions" :key="suggestion">{{ suggestion }}</li>
      </ul>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const filters = ref({
  dateRange: ''
})

const kpi = ref({
  usageCount: 1247,
  conversionRate: 12.5,
  revenue: 45600000,
  spentBudget: 31200000
})

const campaignRows = ref([
  {
    id: 1,
    name: 'کمپین نوروزی',
    usageCount: 450,
    conversionRate: 10.2,
    revenue: 12000000,
    spentBudget: 9500000,
    percentUsed: 95
  },
  {
    id: 2,
    name: 'کمپین تابستانه',
    usageCount: 320,
    conversionRate: 13.8,
    revenue: 15000000,
    spentBudget: 12000000,
    percentUsed: 60
  },
  {
    id: 3,
    name: 'کوپن ویژه مشتریان وفادار',
    usageCount: 210,
    conversionRate: 15.5,
    revenue: 9000000,
    spentBudget: 4000000,
    percentUsed: 80
  }
])

const suggestions = ref([
  'افزایش بودجه کمپین‌هایی که نزدیک به سقف هستند.',
  'بهبود هدف‌گذاری کاربران برای افزایش نرخ تبدیل.',
  'کاهش سقف روزانه کمپین‌هایی با مصرف بالا و نرخ تبدیل پایین.',
  'استفاده از تست A/B برای مقایسه پیام‌های تخفیف.',
  'بررسی زمان‌بندی ارسال کوپن‌ها برای افزایش اثربخشی.'
])

const formatPrice = (price: number): string => {
  return new Intl.NumberFormat('fa-IR').format(price) + ' تومان'
}
</script>

<!--
  کامپوننت تحلیل و بهینه‌سازی کمپین‌ها
  شامل:
  1. نمایش KPI و نمودارهای عملکرد
  2. جدول مقایسه کمپین‌ها
  3. پیشنهادات هوشمند بهبود
  توضیحات کامل به فارسی در کد
--> 
