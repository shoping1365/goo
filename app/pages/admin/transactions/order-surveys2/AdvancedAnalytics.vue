<template>
  <div class="space-y-8">
    <!-- Header -->
    <div>
      <h3 class="text-lg font-semibold text-gray-800">تحلیل پیشرفته و بهینه‌سازی</h3>
      <p class="text-gray-600 text-sm">تحلیل رفتار، زمان‌بندی بهینه و گزارشات مقایسه‌ای</p>
    </div>
    <!-- Optimal Timing Analysis -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h4 class="text-md font-semibold text-gray-800 mb-4">تحلیل زمان‌بندی بهینه</h4>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <h5 class="text-sm font-medium text-gray-700 mb-3">بهترین ساعات پاسخ‌دهی</h5>
          <div class="space-y-2">
            <div v-for="hour in bestHours" :key="hour.hour" class="flex items-center justify-between p-2 bg-gray-50 rounded">
              <span class="text-sm text-gray-700">{{ hour.hour }}</span>
              <div class="flex items-center space-x-2 space-x-reverse">
                <div class="w-20 bg-gray-200 rounded-full h-2">
                  <div :style="{width: hour.rate + '%'}" class="bg-green-500 h-2 rounded-full"></div>
                </div>
                <span class="text-xs text-gray-600">{{ hour.rate }}%</span>
              </div>
            </div>
          </div>
        </div>
        <div>
          <h5 class="text-sm font-medium text-gray-700 mb-3">پیشنهادات بهبود</h5>
          <div class="space-y-3">
            <div v-for="suggestion in suggestions" :key="suggestion.id" class="p-3 border border-blue-200 bg-blue-50 rounded-lg">
              <div class="flex items-start">
                <svg class="w-5 h-5 text-blue-600 mt-0.5 mr-2" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z" clip-rule="evenodd"></path>
                </svg>
                <div>
                  <div class="text-sm font-medium text-blue-900">{{ suggestion.title }}</div>
                  <div class="text-xs text-blue-700 mt-1">{{ suggestion.description }}</div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- Period Comparison -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h4 class="text-md font-semibold text-gray-800 mb-4">مقایسه دوره‌های مختلف</h4>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">دوره اول</label>
          <select v-model="period1" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
            <option value="2024-06">خرداد 1403</option>
            <option value="2024-05">اردیبهشت 1403</option>
            <option value="2024-04">فروردین 1403</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">دوره دوم</label>
          <select v-model="period2" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
            <option value="2024-05">اردیبهشت 1403</option>
            <option value="2024-06">خرداد 1403</option>
            <option value="2024-04">فروردین 1403</option>
          </select>
        </div>
        <div class="flex items-end">
          <button class="w-full px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg text-sm" @click="comparePeriods">مقایسه</button>
        </div>
      </div>
      <div v-if="comparisonData" class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div class="text-center p-6 bg-gray-50 rounded-lg">
          <div class="text-2xl font-bold text-gray-900">{{ comparisonData.responseRate1 }}%</div>
          <div class="text-xs text-gray-600">نرخ پاسخ دوره اول</div>
        </div>
        <div class="text-center p-6 bg-gray-50 rounded-lg">
          <div class="text-2xl font-bold text-gray-900">{{ comparisonData.responseRate2 }}%</div>
          <div class="text-xs text-gray-600">نرخ پاسخ دوره دوم</div>
        </div>
        <div class="text-center p-6 bg-gray-50 rounded-lg">
          <div class="text-2xl font-bold text-gray-900">{{ comparisonData.satisfaction1 }}%</div>
          <div class="text-xs text-gray-600">رضایت دوره اول</div>
        </div>
        <div class="text-center p-6 bg-gray-50 rounded-lg">
          <div class="text-2xl font-bold text-gray-900">{{ comparisonData.satisfaction2 }}%</div>
          <div class="text-xs text-gray-600">رضایت دوره دوم</div>
        </div>
      </div>
    </div>
    <!-- Product Comparison -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h4 class="text-md font-semibold text-gray-800 mb-4">مقایسه محصولات</h4>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-4 py-3 text-right text-xs font-medium text-gray-500">محصول</th>
              <th class="px-4 py-3 text-right text-xs font-medium text-gray-500">تعداد نظرسنجی</th>
              <th class="px-4 py-3 text-right text-xs font-medium text-gray-500">میانگین امتیاز</th>
              <th class="px-4 py-3 text-right text-xs font-medium text-gray-500">نرخ پاسخ</th>
              <th class="px-4 py-3 text-right text-xs font-medium text-gray-500">روند</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="product in productComparison" :key="product.id" class="hover:bg-gray-50">
              <td class="px-4 py-3 whitespace-nowrap text-sm text-gray-900">{{ product.name }}</td>
              <td class="px-4 py-3 whitespace-nowrap text-sm text-gray-900">{{ product.surveys }}</td>
              <td class="px-4 py-3 whitespace-nowrap">
                <div class="flex items-center">
                  <span class="text-sm text-gray-900">{{ product.avgScore }}</span>
                  <div class="flex mr-2">
                    <span v-for="i in Math.round(product.avgScore)" :key="i" class="text-yellow-400 text-xs">★</span>
                  </div>
                </div>
              </td>
              <td class="px-4 py-3 whitespace-nowrap text-sm text-gray-900">{{ product.responseRate }}%</td>
              <td class="px-4 py-3 whitespace-nowrap">
                <span :class="product.trend === 'up' ? 'text-green-600' : 'text-red-600'" class="text-sm">
                  {{ product.trend === 'up' ? '↗' : '↘' }} {{ product.trendValue }}%
                </span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
    <!-- Improvement Trends -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h4 class="text-md font-semibold text-gray-800 mb-4">تحلیل روند بهبود</h4>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div class="text-center p-6 bg-gradient-to-r from-green-50 to-blue-50 rounded-lg">
          <div class="text-3xl font-bold text-green-600 mb-2">+15%</div>
          <div class="text-sm text-gray-700">بهبود نرخ پاسخ</div>
          <div class="text-xs text-gray-500 mt-1">در 3 ماه گذشته</div>
        </div>
        <div class="text-center p-6 bg-gradient-to-r from-blue-50 to-purple-50 rounded-lg">
          <div class="text-3xl font-bold text-blue-600 mb-2">+8%</div>
          <div class="text-sm text-gray-700">بهبود رضایت</div>
          <div class="text-xs text-gray-500 mt-1">در 3 ماه گذشته</div>
        </div>
        <div class="text-center p-6 bg-gradient-to-r from-purple-50 to-pink-50 rounded-lg">
          <div class="text-3xl font-bold text-purple-600 mb-2">-25%</div>
          <div class="text-sm text-gray-700">کاهش خطاهای ارسال</div>
          <div class="text-xs text-gray-500 mt-1">در 3 ماه گذشته</div>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import { ref } from 'vue'

interface ComparisonData {
  responseRate1: number
  responseRate2: number
  satisfaction1: number
  satisfaction2: number
}

const period1 = ref('2024-06')
const period2 = ref('2024-05')
const comparisonData = ref<ComparisonData | null>(null)
const bestHours = ref([
  { hour: '09:00-11:00', rate: 85 },
  { hour: '14:00-16:00', rate: 78 },
  { hour: '19:00-21:00', rate: 72 },
  { hour: '11:00-13:00', rate: 65 },
  { hour: '16:00-18:00', rate: 58 }
])
const suggestions = ref([
  {
    id: 1,
    title: 'ارسال در ساعات بهینه',
    description: 'ارسال پیامک‌ها در ساعات 9-11 صبح و 2-4 بعدازظهر'
  },
  {
    id: 2,
    title: 'بهبود قالب پیام',
    description: 'استفاده از قالب‌های شخصی‌سازی شده برای افزایش نرخ پاسخ'
  },
  {
    id: 3,
    title: 'پیگیری خودکار',
    description: 'ارسال پیامک یادآوری برای مشتریانی که پاسخ نداده‌اند'
  }
])
const productComparison = ref([
  { id: 1, name: 'لپ‌تاپ گیمینگ', surveys: 150, avgScore: 4.2, responseRate: 75, trend: 'up', trendValue: 12 },
  { id: 2, name: 'گوشی هوشمند', surveys: 200, avgScore: 4.5, responseRate: 82, trend: 'up', trendValue: 8 },
  { id: 3, name: 'هدفون بی‌سیم', surveys: 80, avgScore: 3.8, responseRate: 65, trend: 'down', trendValue: 5 },
  { id: 4, name: 'تبلت', surveys: 120, avgScore: 4.1, responseRate: 70, trend: 'up', trendValue: 15 }
])
const comparePeriods = () => {
  // شبیه‌سازی مقایسه دوره‌ها
  comparisonData.value = {
    responseRate1: Math.floor(Math.random() * 30) + 60,
    responseRate2: Math.floor(Math.random() * 30) + 60,
    satisfaction1: Math.floor(Math.random() * 20) + 70,
    satisfaction2: Math.floor(Math.random() * 20) + 70
  }
}
</script> 
