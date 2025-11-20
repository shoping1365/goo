<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between">
        <div>
          <h2 class="text-xl font-bold text-gray-900">گزارش‌های تحلیلی مالی</h2>
          <p class="text-gray-600 mt-1">تحلیل‌های پیشرفته و گزارش‌های مالی کیف پول</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'overview'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
            @click="activeTab = 'overview'"
          >
            نمای کلی
          </button>
          <button
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'trends'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
            @click="activeTab = 'trends'"
          >
            روندها
          </button>
          <button
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'performance'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
            @click="activeTab = 'performance'"
          >
            عملکرد
          </button>
          <button
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'predictions'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
            @click="activeTab = 'predictions'"
          >
            پیش‌بینی‌ها
          </button>
        </div>
      </div>
    </div>

    <!-- محتوای تب‌ها -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200">
      <!-- نمای کلی -->
      <div v-if="activeTab === 'overview'" class="p-6">
        <div class="space-y-6">
          <!-- کارت‌های آماری -->
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
            <div class="bg-white border border-gray-200 rounded-lg p-6">
              <div class="flex items-center">
                <div class="w-8 h-8 bg-green-100 rounded-full flex items-center justify-center">
                  <svg class="w-4 h-4 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1" />
                  </svg>
                </div>
                <div class="mr-3">
                  <p class="text-sm font-medium text-gray-900">کل موجودی</p>
                  <p class="text-lg font-bold text-green-600">{{ formatCurrency(overviewStats.totalBalance) }}</p>
                </div>
              </div>
            </div>
            <div class="bg-white border border-gray-200 rounded-lg p-6">
              <div class="flex items-center">
                <div class="w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center">
                  <svg class="w-4 h-4 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" />
                  </svg>
                </div>
                <div class="mr-3">
                  <p class="text-sm font-medium text-gray-900">تراکنش‌های ماهانه</p>
                  <p class="text-lg font-bold text-blue-600">{{ overviewStats.monthlyTransactions }}</p>
                </div>
              </div>
            </div>
            <div class="bg-white border border-gray-200 rounded-lg p-6">
              <div class="flex items-center">
                <div class="w-8 h-8 bg-purple-100 rounded-full flex items-center justify-center">
                  <svg class="w-4 h-4 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
                  </svg>
                </div>
                <div class="mr-3">
                  <p class="text-sm font-medium text-gray-900">کاربران فعال</p>
                  <p class="text-lg font-bold text-purple-600">{{ overviewStats.activeUsers }}</p>
                </div>
              </div>
            </div>
            <div class="bg-white border border-gray-200 rounded-lg p-6">
              <div class="flex items-center">
                <div class="w-8 h-8 bg-orange-100 rounded-full flex items-center justify-center">
                  <svg class="w-4 h-4 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
                  </svg>
                </div>
                <div class="mr-3">
                  <p class="text-sm font-medium text-gray-900">نرخ رشد</p>
                  <p class="text-lg font-bold text-orange-600">{{ overviewStats.growthRate }}%</p>
                </div>
              </div>
            </div>
          </div>

          <!-- نمودار توزیع تراکنش‌ها -->
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <div class="border border-gray-200 rounded-lg p-6">
              <h3 class="text-lg font-semibold text-gray-900 mb-4">توزیع تراکنش‌ها</h3>
              <div class="space-y-4">
                <div v-for="category in transactionDistribution" :key="category.name" class="flex items-center justify-between">
                  <div class="flex items-center space-x-3 space-x-reverse">
                    <div class="w-4 h-4 rounded-full" :style="{ backgroundColor: category.color }"></div>
                    <span class="text-sm font-medium text-gray-900">{{ category.name }}</span>
                  </div>
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <div class="w-32 bg-gray-200 rounded-full h-2">
                      <div class="bg-blue-600 h-2 rounded-full" :style="{ width: category.percentage + '%' }"></div>
                    </div>
                    <span class="text-sm text-gray-600">{{ category.percentage }}%</span>
                  </div>
                </div>
              </div>
            </div>
            <div class="border border-gray-200 rounded-lg p-6">
              <h3 class="text-lg font-semibold text-gray-900 mb-4">وضعیت کیف پول‌ها</h3>
              <div class="space-y-4">
                <div v-for="status in walletStatus" :key="status.name" class="flex items-center justify-between">
                  <div class="flex items-center space-x-3 space-x-reverse">
                    <div class="w-4 h-4 rounded-full" :style="{ backgroundColor: status.color }"></div>
                    <span class="text-sm font-medium text-gray-900">{{ status.name }}</span>
                  </div>
                  <span class="text-sm text-gray-600">{{ status.count }} کیف پول</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- روندها -->
      <div v-if="activeTab === 'trends'" class="p-6">
        <div class="space-y-6">
          <!-- نمودار روند تراکنش‌ها -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">روند تراکنش‌ها در 30 روز گذشته</h3>
            <div class="h-64 bg-gray-50 rounded-lg flex items-center justify-center">
              <div class="text-center">
                <svg class="w-16 h-16 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
                </svg>
                <p class="text-gray-500">نمودار روند تراکنش‌ها</p>
                <p class="text-sm text-gray-400">نمایش روند روزانه تراکنش‌ها</p>
              </div>
            </div>
          </div>

          <!-- تحلیل روندها -->
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div class="border border-gray-200 rounded-lg p-6">
              <h3 class="text-lg font-semibold text-gray-900 mb-4">روند موجودی</h3>
              <div class="space-y-3">
                <div v-for="trend in balanceTrends" :key="trend.period" class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
                  <span class="text-sm font-medium text-gray-900">{{ trend.period }}</span>
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <span class="text-sm text-gray-600">{{ formatCurrency(trend.amount) }}</span>
                    <span :class="`text-xs px-2 py-1 rounded ${trend.change >= 0 ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'}`">
                      {{ trend.change >= 0 ? '+' : '' }}{{ trend.change }}%
                    </span>
                  </div>
                </div>
              </div>
            </div>
            <div class="border border-gray-200 rounded-lg p-6">
              <h3 class="text-lg font-semibold text-gray-900 mb-4">روند کاربران</h3>
              <div class="space-y-3">
                <div v-for="trend in userTrends" :key="trend.period" class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
                  <span class="text-sm font-medium text-gray-900">{{ trend.period }}</span>
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <span class="text-sm text-gray-600">{{ trend.count }} کاربر</span>
                    <span :class="`text-xs px-2 py-1 rounded ${trend.change >= 0 ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'}`">
                      {{ trend.change >= 0 ? '+' : '' }}{{ trend.change }}%
                    </span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- عملکرد -->
      <div v-if="activeTab === 'performance'" class="p-6">
        <div class="space-y-6">
          <!-- شاخص‌های عملکرد -->
          <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
            <div class="border border-gray-200 rounded-lg p-6">
              <h4 class="font-medium text-gray-900 mb-2">نرخ موفقیت تراکنش‌ها</h4>
              <div class="flex items-center space-x-2 space-x-reverse">
                <div class="text-2xl font-bold text-green-600">{{ performanceMetrics.successRate }}%</div>
                <div class="text-sm text-green-600">↑ +2.5%</div>
              </div>
              <p class="text-xs text-gray-500 mt-1">نسبت به ماه گذشته</p>
            </div>
            <div class="border border-gray-200 rounded-lg p-6">
              <h4 class="font-medium text-gray-900 mb-2">زمان متوسط پردازش</h4>
              <div class="flex items-center space-x-2 space-x-reverse">
                <div class="text-2xl font-bold text-blue-600">{{ performanceMetrics.avgProcessingTime }}s</div>
                <div class="text-sm text-red-600">↓ -0.3s</div>
              </div>
              <p class="text-xs text-gray-500 mt-1">نسبت به ماه گذشته</p>
            </div>
            <div class="border border-gray-200 rounded-lg p-6">
              <h4 class="font-medium text-gray-900 mb-2">رضایت کاربران</h4>
              <div class="flex items-center space-x-2 space-x-reverse">
                <div class="text-2xl font-bold text-purple-600">{{ performanceMetrics.userSatisfaction }}/5</div>
                <div class="text-sm text-green-600">↑ +0.2</div>
              </div>
              <p class="text-xs text-gray-500 mt-1">نسبت به ماه گذشته</p>
            </div>
          </div>

          <!-- مقایسه عملکرد -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">مقایسه عملکرد</h3>
            <div class="overflow-x-auto">
              <table class="min-w-full divide-y divide-gray-200">
                <thead class="bg-gray-50">
                  <tr>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">شاخص</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">ماه جاری</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">ماه گذشته</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تغییر</th>
                  </tr>
                </thead>
                <tbody class="bg-white divide-y divide-gray-200">
                  <tr v-for="metric in performanceComparison" :key="metric.name">
                    <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{{ metric.name }}</td>
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ metric.current }}</td>
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ metric.previous }}</td>
                    <td class="px-6 py-4 whitespace-nowrap">
                      <span :class="`text-sm ${metric.change >= 0 ? 'text-green-600' : 'text-red-600'}`">
                        {{ metric.change >= 0 ? '+' : '' }}{{ metric.change }}%
                      </span>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>

      <!-- پیش‌بینی‌ها -->
      <div v-if="activeTab === 'predictions'" class="p-6">
        <div class="space-y-6">
          <!-- پیش‌بینی تراکنش‌ها -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">پیش‌بینی تراکنش‌ها (3 ماه آینده)</h3>
            <div class="h-64 bg-gray-50 rounded-lg flex items-center justify-center">
              <div class="text-center">
                <svg class="w-16 h-16 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z" />
                </svg>
                <p class="text-gray-500">نمودار پیش‌بینی تراکنش‌ها</p>
                <p class="text-sm text-gray-400">پیش‌بینی بر اساس الگوهای تاریخی</p>
              </div>
            </div>
          </div>

          <!-- پیش‌بینی‌های عددی -->
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div class="border border-gray-200 rounded-lg p-6">
              <h3 class="text-lg font-semibold text-gray-900 mb-4">پیش‌بینی موجودی</h3>
              <div class="space-y-3">
                <div v-for="prediction in balancePredictions" :key="prediction.month" class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
                  <span class="text-sm font-medium text-gray-900">{{ prediction.month }}</span>
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <span class="text-sm text-gray-600">{{ formatCurrency(prediction.amount) }}</span>
                    <span class="text-xs text-blue-600">±{{ prediction.confidence }}%</span>
                  </div>
                </div>
              </div>
            </div>
            <div class="border border-gray-200 rounded-lg p-6">
              <h3 class="text-lg font-semibold text-gray-900 mb-4">پیش‌بینی کاربران</h3>
              <div class="space-y-3">
                <div v-for="prediction in userPredictions" :key="prediction.month" class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
                  <span class="text-sm font-medium text-gray-900">{{ prediction.month }}</span>
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <span class="text-sm text-gray-600">{{ prediction.count }} کاربر</span>
                    <span class="text-xs text-blue-600">±{{ prediction.confidence }}%</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
// تعریف reactive data
const activeTab = ref('overview')

// آمار نمای کلی
const overviewStats = ref({
  totalBalance: 12500000000, // 12.5 میلیارد تومان
  monthlyTransactions: 15420,
  activeUsers: 3420,
  growthRate: 15.8
})

// توزیع تراکنش‌ها
const transactionDistribution = ref([
  { name: 'شارژ', percentage: 45, color: '#10B981' },
  { name: 'برداشت', percentage: 30, color: '#3B82F6' },
  { name: 'انتقال', percentage: 15, color: '#8B5CF6' },
  { name: 'سایر', percentage: 10, color: '#F59E0B' }
])

// وضعیت کیف پول‌ها
const walletStatus = ref([
  { name: 'فعال', count: 2840, color: '#10B981' },
  { name: 'غیرفعال', count: 580, color: '#6B7280' },
  { name: 'معلق', count: 120, color: '#F59E0B' },
  { name: 'مسدود', count: 45, color: '#EF4444' }
])

// روند موجودی
const balanceTrends = ref([
  { period: 'هفته اول', amount: 11500000000, change: 5.2 },
  { period: 'هفته دوم', amount: 11800000000, change: 2.6 },
  { period: 'هفته سوم', amount: 12200000000, change: 3.4 },
  { period: 'هفته چهارم', amount: 12500000000, change: 2.5 }
])

// روند کاربران
const userTrends = ref([
  { period: 'هفته اول', count: 3100, change: 8.5 },
  { period: 'هفته دوم', count: 3200, change: 3.2 },
  { period: 'هفته سوم', count: 3300, change: 3.1 },
  { period: 'هفته چهارم', count: 3420, change: 3.6 }
])

// شاخص‌های عملکرد
const performanceMetrics = ref({
  successRate: 98.5,
  avgProcessingTime: 1.2,
  userSatisfaction: 4.6
})

// مقایسه عملکرد
const performanceComparison = ref([
  { name: 'نرخ موفقیت', current: '98.5%', previous: '96.0%', change: 2.5 },
  { name: 'زمان پردازش', current: '1.2s', previous: '1.5s', change: -20.0 },
  { name: 'رضایت کاربران', current: '4.6/5', previous: '4.4/5', change: 4.5 },
  { name: 'تعداد تراکنش‌ها', current: '15,420', previous: '13,200', change: 16.8 }
])

// پیش‌بینی موجودی
const balancePredictions = ref([
  { month: 'ماه آینده', amount: 13500000000, confidence: 5 },
  { month: '2 ماه آینده', amount: 14500000000, confidence: 8 },
  { month: '3 ماه آینده', amount: 15500000000, confidence: 12 }
])

// پیش‌بینی کاربران
const userPredictions = ref([
  { month: 'ماه آینده', count: 3700, confidence: 3 },
  { month: '2 ماه آینده', count: 4000, confidence: 5 },
  { month: '3 ماه آینده', count: 4300, confidence: 8 }
])

// تابع فرمت‌بندی ارز
const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR', {
    style: 'currency',
    currency: 'IRR',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(amount)
}
</script> 
