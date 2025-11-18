<template>
  <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
    <!-- هدر -->
    <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-6 mb-6">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">مقایسه دوره‌ای مالیات</h3>
        <p class="text-sm text-gray-600">مقایسه مالیات با دوره‌های قبل</p>
      </div>
      
      <!-- انتخاب دوره مقایسه -->
      <div class="flex gap-2">
        <select 
          v-model="comparisonPeriod"
          class="px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
        >
          <option value="monthly">مقایسه ماهانه</option>
          <option value="quarterly">مقایسه فصلی</option>
          <option value="yearly">مقایسه سالانه</option>
        </select>
      </div>
    </div>

    <!-- کارت‌های مقایسه -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <!-- مقایسه با دوره قبل -->
      <div class="bg-gradient-to-br from-blue-50 to-blue-100 rounded-xl p-6">
        <div class="flex items-center justify-between mb-4">
          <div class="w-10 h-10 bg-blue-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
          </div>
          <div class="text-right">
            <div class="text-sm text-blue-600">دوره قبل</div>
            <div class="text-lg font-bold text-blue-900">{{ formatCurrency(previousPeriodAmount) }}</div>
          </div>
        </div>
        <div class="flex items-center gap-2">
          <div class="flex items-center gap-1" :class="previousPeriodGrowth >= 0 ? 'text-green-600' : 'text-red-600'">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="previousPeriodGrowth >= 0 ? 'M13 7h8m0 0v8m0-8l-8 8-4-4-6 6' : 'M13 17h8m0 0V9m0 8l-8-8-4 4-6-6'"></path>
            </svg>
            <span class="text-sm font-medium">{{ previousPeriodGrowth >= 0 ? '+' : '' }}{{ previousPeriodGrowth }}%</span>
          </div>
          <span class="text-sm text-blue-600">نسبت به دوره قبل</span>
        </div>
      </div>

      <!-- مقایسه با سال قبل -->
      <div class="bg-gradient-to-br from-green-50 to-green-100 rounded-xl p-6">
        <div class="flex items-center justify-between mb-4">
          <div class="w-10 h-10 bg-green-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
            </svg>
          </div>
          <div class="text-right">
            <div class="text-sm text-green-600">سال قبل</div>
            <div class="text-lg font-bold text-green-900">{{ formatCurrency(lastYearAmount) }}</div>
          </div>
        </div>
        <div class="flex items-center gap-2">
          <div class="flex items-center gap-1" :class="lastYearGrowth >= 0 ? 'text-green-600' : 'text-red-600'">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="lastYearGrowth >= 0 ? 'M13 7h8m0 0v8m0-8l-8 8-4-4-6 6' : 'M13 17h8m0 0V9m0 8l-8-8-4 4-6-6'"></path>
            </svg>
            <span class="text-sm font-medium">{{ lastYearGrowth >= 0 ? '+' : '' }}{{ lastYearGrowth }}%</span>
          </div>
          <span class="text-sm text-green-600">نسبت به سال قبل</span>
        </div>
      </div>

      <!-- مقایسه با میانگین -->
      <div class="bg-gradient-to-br from-purple-50 to-purple-100 rounded-xl p-6">
        <div class="flex items-center justify-between mb-4">
          <div class="w-10 h-10 bg-purple-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
          </div>
          <div class="text-right">
            <div class="text-sm text-purple-600">میانگین</div>
            <div class="text-lg font-bold text-purple-900">{{ formatCurrency(averageAmount) }}</div>
          </div>
        </div>
        <div class="flex items-center gap-2">
          <div class="flex items-center gap-1" :class="averageGrowth >= 0 ? 'text-green-600' : 'text-red-600'">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="averageGrowth >= 0 ? 'M13 7h8m0 0v8m0-8l-8 8-4-4-6 6' : 'M13 17h8m0 0V9m0 8l-8-8-4 4-6-6'"></path>
            </svg>
            <span class="text-sm font-medium">{{ averageGrowth >= 0 ? '+' : '' }}{{ averageGrowth }}%</span>
          </div>
          <span class="text-sm text-purple-600">نسبت به میانگین</span>
        </div>
      </div>
    </div>

    <!-- جدول مقایسه تفصیلی -->
    <div class="mt-8">
      <h4 class="text-md font-semibold text-gray-900 mb-4">مقایسه تفصیلی</h4>
      <div class="overflow-x-auto">
        <table class="w-full text-sm">
          <thead>
            <tr class="border-b border-gray-200">
              <th class="text-right py-3 px-4 font-medium text-gray-600">دوره</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">مبلغ مالیات</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">تغییر</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">درصد تغییر</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">وضعیت</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="item in comparisonTable" :key="item.period" class="border-b border-gray-100">
              <td class="py-3 px-4 text-gray-900">{{ item.period }}</td>
              <td class="py-3 px-4 text-gray-900 font-medium">{{ formatCurrency(item.amount) }}</td>
              <td class="py-3 px-4">
                <span :class="item.change >= 0 ? 'text-green-600' : 'text-red-600'">
                  {{ item.change >= 0 ? '+' : '' }}{{ formatCurrency(item.change) }}
                </span>
              </td>
              <td class="py-3 px-4">
                <span :class="item.percentage >= 0 ? 'text-green-600' : 'text-red-600'">
                  {{ item.percentage >= 0 ? '+' : '' }}{{ item.percentage }}%
                </span>
              </td>
              <td class="py-3 px-4">
                <span 
                  :class="[
                    'px-2 py-1 rounded-full text-xs font-medium',
                    item.status === 'up' ? 'bg-green-100 text-green-700' : 
                    item.status === 'down' ? 'bg-red-100 text-red-700' : 
                    'bg-gray-100 text-gray-700'
                  ]"
                >
                  {{ item.statusText }}
                </span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- خلاصه تحلیل -->
    <div class="mt-6 p-6 bg-gray-50 rounded-lg">
      <div class="flex items-start gap-3">
        <div class="w-8 h-8 bg-gray-200 rounded-lg flex items-center justify-center flex-shrink-0">
          <svg class="w-4 h-4 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
          </svg>
        </div>
        <div>
          <h4 class="font-medium text-gray-900">خلاصه تحلیل</h4>
          <p class="text-sm text-gray-600 mt-1">
            {{ analysisSummary }}
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch, onMounted } from 'vue'

// دوره مقایسه
const comparisonPeriod = ref('monthly')

// داده‌های مقایسه
const previousPeriodAmount = ref(0)
const lastYearAmount = ref(0)
const averageAmount = ref(0)
const previousPeriodGrowth = ref(0)
const lastYearGrowth = ref(0)
const averageGrowth = ref(0)

// جدول مقایسه
const comparisonTable = ref([
  {
    period: 'ماه جاری',
    amount: 51000000,
    change: 2000000,
    percentage: 4.1,
    status: 'up',
    statusText: 'افزایش'
  },
  {
    period: 'ماه قبل',
    amount: 49000000,
    change: 7000000,
    percentage: 16.7,
    status: 'up',
    statusText: 'افزایش'
  },
  {
    period: '۲ ماه قبل',
    amount: 42000000,
    change: -3000000,
    percentage: -6.7,
    status: 'down',
    statusText: 'کاهش'
  }
])

// خلاصه تحلیل
const analysisSummary = computed(() => {
  const positiveCount = comparisonTable.value.filter(item => item.status === 'up').length
  const negativeCount = comparisonTable.value.filter(item => item.status === 'down').length
  
  if (positiveCount > negativeCount) {
    return 'روند کلی مالیات در این دوره صعودی بوده و رشد مثبتی را نشان می‌دهد.'
  } else if (negativeCount > positiveCount) {
    return 'روند کلی مالیات در این دوره نزولی بوده و نیاز به بررسی بیشتر دارد.'
  } else {
    return 'روند مالیات در این دوره متغیر بوده و ثبات نسبی دارد.'
  }
})

// فرمت کردن مبلغ به تومان
const formatCurrency = (amount: number): string => {
  return new Intl.NumberFormat('fa-IR', {
    style: 'currency',
    currency: 'IRR',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(amount)
}

// بارگذاری داده‌های مقایسه
const loadComparisonData = async () => {
  try {
    // TODO: دریافت داده‌ها از API بر اساس دوره انتخاب شده
    // const response = await $fetch(`/api/admin/tax/comparison?period=${comparisonPeriod.value}`)
    
    // داده‌های نمونه
    previousPeriodAmount.value = 49000000
    lastYearAmount.value = 42000000
    averageAmount.value = 45500000
    previousPeriodGrowth.value = 4.1
    lastYearGrowth.value = 21.4
    averageGrowth.value = 12.1
  } catch (error) {
    console.error('خطا در بارگذاری داده‌های مقایسه:', error)
  }
}

// نظارت بر تغییر دوره
watch(comparisonPeriod, () => {
  loadComparisonData()
})

// بارگذاری اولیه
onMounted(() => {
  loadComparisonData()
})
</script>

<!--
  کامپوننت مقایسه دوره‌ای مالیات
  شامل:
  1. کارت‌های مقایسه با دوره‌های مختلف
  2. جدول مقایسه تفصیلی
  3. تحلیل و خلاصه روند
  4. فیلترهای دوره مقایسه
  5. طراحی مدرن و کاملاً ریسپانسیو
--> 
