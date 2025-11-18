<template>
  <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
    <div class="flex items-center justify-between mb-6">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">تحلیل‌های پیشرفته</h3>
        <p class="text-sm text-gray-500 mt-1">Advanced Analytics & Insights</p>
      </div>
    </div>

    <div class="space-y-6">
      <!-- تحلیل‌های اصلی -->
      <div class="space-y-4">
        <!-- نرخ تبدیل -->
        <div class="flex items-center justify-between p-6 bg-blue-50 rounded-lg border border-blue-200">
          <div class="flex items-center">
            <div class="w-10 h-10 bg-blue-500 rounded-lg flex items-center justify-center ml-3">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
              </svg>
            </div>
            <div>
              <h4 class="font-medium text-blue-800">نرخ تبدیل</h4>
              <p class="text-sm text-blue-600">Conversion Rate</p>
            </div>
          </div>
          <div class="text-left">
            <div class="text-2xl font-bold text-blue-700">{{ analytics.conversionRate }}%</div>
            <div class="text-sm text-blue-600">+2.3% از ماه گذشته</div>
          </div>
        </div>

        <!-- میانگین ارزش تراکنش -->
        <div class="flex items-center justify-between p-6 bg-green-50 rounded-lg border border-green-200">
          <div class="flex items-center">
            <div class="w-10 h-10 bg-green-500 rounded-lg flex items-center justify-center ml-3">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
              </svg>
            </div>
            <div>
              <h4 class="font-medium text-green-800">میانگین ارزش تراکنش</h4>
              <p class="text-sm text-green-600">Average Transaction Value</p>
            </div>
          </div>
          <div class="text-left">
            <div class="text-2xl font-bold text-green-700">{{ formatPrice(analytics.avgPaymentAmount) }}</div>
            <div class="text-sm text-green-600">+5.7% از ماه گذشته</div>
          </div>
        </div>

        <!-- نرخ مرجوعی -->
        <div class="flex items-center justify-between p-6 bg-red-50 rounded-lg border border-red-200">
          <div class="flex items-center">
            <div class="w-10 h-10 bg-red-500 rounded-lg flex items-center justify-center ml-3">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"></path>
              </svg>
            </div>
            <div>
              <h4 class="font-medium text-red-800">نرخ مرجوعی</h4>
              <p class="text-sm text-red-600">Refund Rate</p>
            </div>
          </div>
          <div class="text-left">
            <div class="text-2xl font-bold text-red-700">{{ analytics.refundRate }}%</div>
            <div class="text-sm text-red-600">-0.5% از ماه گذشته</div>
          </div>
        </div>

        <!-- نرخ Chargeback -->
        <div class="flex items-center justify-between p-6 bg-orange-50 rounded-lg border border-orange-200">
          <div class="flex items-center">
            <div class="w-10 h-10 bg-orange-500 rounded-lg flex items-center justify-center ml-3">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
              </svg>
            </div>
            <div>
              <h4 class="font-medium text-orange-800">نرخ Chargeback</h4>
              <p class="text-sm text-orange-600">Chargeback Rate</p>
            </div>
          </div>
          <div class="text-left">
            <div class="text-2xl font-bold text-orange-700">{{ analytics.chargebackRate }}%</div>
            <div class="text-sm text-orange-600">-0.2% از ماه گذشته</div>
          </div>
        </div>
      </div>

      <!-- گزارش تفکیکی بر اساس درگاه -->
      <div class="pt-4 border-t border-gray-200">
        <h4 class="font-medium text-gray-900 mb-3">گزارش تفکیکی بر اساس درگاه</h4>
        <div class="space-y-3">
          <div v-for="gateway in gatewayStats" :key="gateway.name" class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <div class="flex items-center">
              <div class="w-8 h-8 rounded-lg flex items-center justify-center ml-3" :class="gateway.color">
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"></path>
                </svg>
              </div>
              <div>
                <div class="font-medium text-gray-900">{{ gateway.name }}</div>
                <div class="text-sm text-gray-500">{{ gateway.successRate }}% موفقیت</div>
              </div>
            </div>
            <div class="text-left">
              <div class="font-semibold text-gray-900">{{ formatPrice(gateway.totalAmount) }}</div>
              <div class="text-sm text-gray-500">{{ gateway.count }} تراکنش</div>
            </div>
          </div>
        </div>
      </div>

      <!-- گزارش سود و کارمزد -->
      <div class="pt-4 border-t border-gray-200">
        <h4 class="font-medium text-gray-900 mb-3">گزارش سود و کارمزد</h4>
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div class="text-center p-6 bg-emerald-50 rounded-lg border border-emerald-200">
            <div class="text-2xl font-bold text-emerald-700">{{ formatPrice(profitStats.totalRevenue) }}</div>
            <div class="text-sm text-emerald-600">کل درآمد</div>
          </div>
          <div class="text-center p-6 bg-red-50 rounded-lg border border-red-200">
            <div class="text-2xl font-bold text-red-700">{{ formatPrice(profitStats.totalFees) }}</div>
            <div class="text-sm text-red-600">کل کارمزد</div>
          </div>
          <div class="text-center p-6 bg-green-50 rounded-lg border border-green-200">
            <div class="text-2xl font-bold text-green-700">{{ formatPrice(profitStats.netProfit) }}</div>
            <div class="text-sm text-green-600">سود خالص</div>
          </div>
        </div>
        <div class="mt-3 text-center">
          <div class="text-sm text-gray-600">نرخ سود: {{ profitStats.profitMargin }}%</div>
        </div>
      </div>

      <!-- گزارش مقایسه‌ای دوره‌ها -->
      <div class="pt-4 border-t border-gray-200">
        <h4 class="font-medium text-gray-900 mb-3">گزارش مقایسه‌ای دوره‌ها</h4>
        <div class="space-y-3">
          <div v-for="period in periodComparison" :key="period.name" class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <div>
              <div class="font-medium text-gray-900">{{ period.name }}</div>
              <div class="text-sm text-gray-500">{{ period.duration }}</div>
            </div>
            <div class="text-left">
              <div class="font-semibold text-gray-900">{{ formatPrice(period.amount) }}</div>
              <div class="text-sm" :class="period.growth >= 0 ? 'text-green-600' : 'text-red-600'">
                {{ period.growth >= 0 ? '+' : '' }}{{ period.growth }}% رشد
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- نمودار روند نرخ تبدیل -->
      <div class="pt-4 border-t border-gray-200">
        <h4 class="font-medium text-gray-900 mb-3">روند نرخ تبدیل (۳۰ روز گذشته)</h4>
        <div class="h-32 bg-gray-50 rounded-lg flex items-center justify-center">
          <div class="text-center">
            <svg class="w-12 h-12 text-gray-300 mx-auto mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
            <p class="text-sm text-gray-500">نمودار روند نرخ تبدیل</p>
          </div>
        </div>
      </div>

      <!-- نکات کلیدی -->
      <div class="pt-4 border-t border-gray-200">
        <h4 class="font-medium text-gray-900 mb-3">نکات کلیدی</h4>
        <div class="space-y-2">
          <div class="flex items-start">
            <div class="w-2 h-2 bg-green-500 rounded-full mt-2 ml-2"></div>
            <p class="text-sm text-gray-600">نرخ تبدیل در حال بهبود است و ۲.۳٪ افزایش داشته</p>
          </div>
          <div class="flex items-start">
            <div class="w-2 h-2 bg-blue-500 rounded-full mt-2 ml-2"></div>
            <p class="text-sm text-gray-600">میانگین ارزش تراکنش‌ها ۵.۷٪ رشد داشته</p>
          </div>
          <div class="flex items-start">
            <div class="w-2 h-2 bg-orange-500 rounded-full mt-2 ml-2"></div>
            <p class="text-sm text-gray-600">نرخ مرجوعی و Chargeback کاهش یافته</p>
          </div>
          <div class="flex items-start">
            <div class="w-2 h-2 bg-purple-500 rounded-full mt-2 ml-2"></div>
            <p class="text-sm text-gray-600">درگاه زرین‌پال بالاترین نرخ موفقیت را دارد</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface Props {
  analytics: {
    conversionRate: number
    avgPaymentAmount: number
    refundRate: number
    chargebackRate: number
  }
}

const props = defineProps<Props>()

// آمار درگاه‌ها
const gatewayStats = [
  {
    name: 'زرین‌پال',
    totalAmount: 45621000,
    count: 234,
    successRate: 94.2,
    color: 'bg-green-500'
  },
  {
    name: 'نکست‌پی',
    totalAmount: 38934000,
    count: 198,
    successRate: 91.8,
    color: 'bg-blue-500'
  },
  {
    name: 'کارت به کارت',
    totalAmount: 24521000,
    count: 156,
    successRate: 87.5,
    color: 'bg-purple-500'
  },
  {
    name: 'کیف پول',
    totalAmount: 12847000,
    count: 89,
    successRate: 89.3,
    color: 'bg-orange-500'
  }
]

// آمار سود و کارمزد
const profitStats = {
  totalRevenue: 125000000,
  totalFees: 8750000,
  netProfit: 116250000,
  profitMargin: 93.0
}

// مقایسه دوره‌ها
const periodComparison = [
  {
    name: 'این ماه',
    duration: '۱ تا ۳۰ روز گذشته',
    amount: 125000000,
    growth: 15.7
  },
  {
    name: 'ماه گذشته',
    duration: '۱ تا ۳۰ روز قبل',
    amount: 108000000,
    growth: 8.3
  },
  {
    name: 'دو ماه پیش',
    duration: '۳۱ تا ۶۰ روز قبل',
    amount: 99800000,
    growth: -2.1
  }
]

const formatPrice = (price: number) => {
  if (price >= 1000000) {
    return (price / 1000000).toFixed(1) + 'M'
  } else if (price >= 1000) {
    return (price / 1000).toFixed(0) + 'K'
  }
  return new Intl.NumberFormat('fa-IR').format(price)
}
</script>

<!--
  کامپوننت تحلیل‌های پیشرفته
  - نرخ تبدیل، میانگین ارزش، نرخ مرجوعی و Chargeback
  - گزارش تفکیکی بر اساس درگاه
  - گزارش سود و کارمزد
  - گزارش مقایسه‌ای دوره‌ها
  - نمودار روند
  - نکات کلیدی
  - کاملاً ریسپانسیو
  توضیحات کامل به فارسی در کد
--> 
