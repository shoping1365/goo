<template>
  <div class="mb-10">
    <!-- آمار کلی تست‌ها -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-sm text-gray-500 mb-1">کل تست‌ها</div>
            <div class="text-3xl font-bold text-gray-900">{{ stats.totalTests }}</div>
            <div class="text-xs text-green-600 mt-1">+12% از ماه گذشته</div>
          </div>
          <div class="p-3 bg-blue-100 rounded-lg">
            <svg class="w-8 h-8 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-sm text-gray-500 mb-1">تست‌های فعال</div>
            <div class="text-3xl font-bold text-green-600">{{ stats.activeTests }}</div>
            <div class="text-xs text-green-600 mt-1">+5 تست جدید</div>
          </div>
          <div class="p-3 bg-green-100 rounded-lg">
            <svg class="w-8 h-8 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-sm text-gray-500 mb-1">در انتظار</div>
            <div class="text-3xl font-bold text-yellow-600">{{ stats.pendingTests }}</div>
            <div class="text-xs text-yellow-600 mt-1">آماده شروع</div>
          </div>
          <div class="p-3 bg-yellow-100 rounded-lg">
            <svg class="w-8 h-8 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-sm text-gray-500 mb-1">تکمیل شده</div>
            <div class="text-3xl font-bold text-blue-600">{{ stats.completedTests }}</div>
            <div class="text-xs text-blue-600 mt-1">نرخ موفقیت {{ stats.successRate }}%</div>
          </div>
          <div class="p-3 bg-blue-100 rounded-lg">
            <svg class="w-8 h-8 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- نمودارهای آماری -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-8">
      <!-- نمودار تست‌های فعال در طول زمان -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">تست‌های فعال در طول زمان</h3>
          <select class="text-sm border border-gray-300 rounded-lg px-3 py-1">
            <option>7 روز اخیر</option>
            <option>30 روز اخیر</option>
            <option>90 روز اخیر</option>
          </select>
        </div>
        <div class="h-64 flex items-end justify-between space-x-2 space-x-reverse">
          <div v-for="(day, index) in activeTestsChart" :key="index" class="flex-1 flex flex-col items-center">
            <div class="w-full bg-blue-500 rounded-t-lg" :style="{ height: `${day.value * 4}px` }"></div>
            <span class="text-xs text-gray-500 mt-2">{{ day.label }}</span>
          </div>
        </div>
      </div>

      <!-- نمودار نرخ تبدیل -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">نرخ تبدیل (Conversion Rate)</h3>
          <div class="flex space-x-2 space-x-reverse">
            <div class="flex items-center">
              <div class="w-3 h-3 bg-blue-500 rounded-full ml-1"></div>
              <span class="text-xs text-gray-600">نسخه A</span>
            </div>
            <div class="flex items-center">
              <div class="w-3 h-3 bg-green-500 rounded-full ml-1"></div>
              <span class="text-xs text-gray-600">نسخه B</span>
            </div>
          </div>
        </div>
        <div class="h-64 flex items-end justify-center space-x-8 space-x-reverse">
          <div class="flex flex-col items-center">
            <div class="w-16 bg-blue-500 rounded-t-lg h-[120px]">
              <div class="text-white text-xs text-center pt-2 font-bold">12.3%</div>
            </div>
            <span class="text-sm text-gray-600 mt-2">نسخه A</span>
          </div>
          <div class="flex flex-col items-center">
            <div class="w-16 bg-green-500 rounded-t-lg h-[150px]">
              <div class="text-white text-xs text-center pt-2 font-bold">15.8%</div>
            </div>
            <span class="text-sm text-gray-600 mt-2">نسخه B</span>
          </div>
        </div>
      </div>
    </div>

    <!-- نمودارهای ترافیک و درآمد -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- نمودار ترافیک تقسیم شده -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">ترافیک تقسیم شده</h3>
        <div class="flex items-center justify-center">
          <div class="relative w-48 h-48">
            <!-- دایره نمودار -->
            <svg class="w-48 h-48 transform -rotate-90" viewBox="0 0 100 100">
              <circle cx="50" cy="50" r="40" stroke="#e5e7eb" stroke-width="10" fill="none" />
              <circle
cx="50" cy="50" r="40" stroke="#3b82f6" stroke-width="10" fill="none" 
                      stroke-dasharray="125.6" stroke-dashoffset="62.8" />
              <circle
cx="50" cy="50" r="40" stroke="#10b981" stroke-width="10" fill="none" 
                      stroke-dasharray="125.6" stroke-dashoffset="0" />
            </svg>
            <div class="absolute inset-0 flex items-center justify-center">
              <div class="text-center">
                <div class="text-2xl font-bold text-gray-900">{{ stats.totalTraffic }}</div>
                <div class="text-sm text-gray-500">کل ترافیک</div>
              </div>
            </div>
          </div>
        </div>
        <div class="flex justify-center space-x-6 space-x-reverse mt-4">
          <div class="flex items-center">
            <div class="w-3 h-3 bg-blue-500 rounded-full ml-2"></div>
            <span class="text-sm text-gray-600">نسخه A (50%)</span>
          </div>
          <div class="flex items-center">
            <div class="w-3 h-3 bg-green-500 rounded-full ml-2"></div>
            <span class="text-sm text-gray-600">نسخه B (50%)</span>
          </div>
        </div>
      </div>

      <!-- نمودار درآمد تاثیر گرفته -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">درآمد تاثیر گرفته</h3>
        <div class="space-y-4">
          <div class="flex items-center justify-between p-3 bg-green-50 rounded-lg">
            <div>
              <div class="text-sm text-gray-600">افزایش درآمد</div>
              <div class="text-lg font-bold text-green-600">+{{ formatCurrency(stats.revenueIncrease) }}</div>
            </div>
            <div class="text-green-600">
              <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" />
              </svg>
            </div>
          </div>
          <div class="flex items-center justify-between p-3 bg-blue-50 rounded-lg">
            <div>
              <div class="text-sm text-gray-600">ROI تست‌ها</div>
              <div class="text-lg font-bold text-blue-600">{{ stats.roi }}%</div>
            </div>
            <div class="text-blue-600">
              <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1" />
              </svg>
            </div>
          </div>
          <div class="flex items-center justify-between p-3 bg-purple-50 rounded-lg">
            <div>
              <div class="text-sm text-gray-600">متوسط بهبود</div>
              <div class="text-lg font-bold text-purple-600">+{{ stats.avgImprovement }}%</div>
            </div>
            <div class="text-purple-600">
              <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
              </svg>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
// داده‌های آماری (mock data)
const stats = ref({
  totalTests: 24,
  activeTests: 8,
  pendingTests: 3,
  completedTests: 13,
  successRate: 85,
  totalTraffic: 15420,
  revenueIncrease: 2850000,
  roi: 340,
  avgImprovement: 18.5
})

// داده‌های نمودار تست‌های فعال
const activeTestsChart = ref([
  { label: 'ش', value: 12 },
  { label: 'ی', value: 15 },
  { label: 'د', value: 8 },
  { label: 'س', value: 18 },
  { label: 'چ', value: 22 },
  { label: 'پ', value: 16 },
  { label: 'ج', value: 25 }
])

// فرمت کردن ارز
const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR', {
    style: 'currency',
    currency: 'IRR',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(amount)
}
</script> 
