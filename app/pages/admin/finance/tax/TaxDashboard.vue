<template>
  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
    <!-- کارت کل مالیات جمع‌آوری شده -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-4">
        <div class="w-12 h-12 bg-gradient-to-br from-blue-500 to-blue-600 rounded-lg flex items-center justify-center">
          <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
          </svg>
        </div>
        <div class="text-right">
          <div class="text-sm text-gray-500">کل مالیات</div>
          <div class="text-2xl font-bold text-gray-900">{{ formatCurrency(totalTaxCollected) }}</div>
        </div>
      </div>
      <div class="flex items-center gap-2">
        <div class="flex items-center gap-1 text-green-600">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
          </svg>
          <span class="text-sm font-medium">+{{ growthRate }}%</span>
        </div>
        <span class="text-sm text-gray-500">نسبت به ماه قبل</span>
      </div>
    </div>

    <!-- کارت مالیات ماهانه -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-4">
        <div class="w-12 h-12 bg-gradient-to-br from-green-500 to-green-600 rounded-lg flex items-center justify-center">
          <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
          </svg>
        </div>
        <div class="text-right">
          <div class="text-sm text-gray-500">مالیات ماهانه</div>
          <div class="text-2xl font-bold text-gray-900">{{ formatCurrency(monthlyTax) }}</div>
        </div>
      </div>
      <div class="flex items-center gap-2">
        <div class="flex items-center gap-1 text-blue-600">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
          </svg>
          <span class="text-sm font-medium">+{{ monthlyGrowth }}%</span>
        </div>
        <span class="text-sm text-gray-500">نسبت به ماه قبل</span>
      </div>
    </div>

    <!-- کارت مالیات فصلی -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-4">
        <div class="w-12 h-12 bg-gradient-to-br from-purple-500 to-purple-600 rounded-lg flex items-center justify-center">
          <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
          </svg>
        </div>
        <div class="text-right">
          <div class="text-sm text-gray-500">مالیات فصلی</div>
          <div class="text-2xl font-bold text-gray-900">{{ formatCurrency(quarterlyTax) }}</div>
        </div>
      </div>
      <div class="flex items-center gap-2">
        <div class="flex items-center gap-1 text-purple-600">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
          </svg>
          <span class="text-sm font-medium">+{{ quarterlyGrowth }}%</span>
        </div>
        <span class="text-sm text-gray-500">نسبت به فصل قبل</span>
      </div>
    </div>

    <!-- کارت مالیات سالانه -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-4">
        <div class="w-12 h-12 bg-gradient-to-br from-orange-500 to-orange-600 rounded-lg flex items-center justify-center">
          <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
          </svg>
        </div>
        <div class="text-right">
          <div class="text-sm text-gray-500">مالیات سالانه</div>
          <div class="text-2xl font-bold text-gray-900">{{ formatCurrency(yearlyTax) }}</div>
        </div>
      </div>
      <div class="flex items-center gap-2">
        <div class="flex items-center gap-1 text-orange-600">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
          </svg>
          <span class="text-sm font-medium">+{{ yearlyGrowth }}%</span>
        </div>
        <span class="text-sm text-gray-500">نسبت به سال قبل</span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'

// داده‌های آمار مالیات
const totalTaxCollected = ref(0)
const monthlyTax = ref(0)
const quarterlyTax = ref(0)
const yearlyTax = ref(0)
const growthRate = ref(0)
const monthlyGrowth = ref(0)
const quarterlyGrowth = ref(0)
const yearlyGrowth = ref(0)

// فرمت کردن مبلغ به تومان
const formatCurrency = (amount: number): string => {
  return new Intl.NumberFormat('fa-IR', {
    style: 'currency',
    currency: 'IRR',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(amount)
}

// بارگذاری داده‌های آمار
const loadDashboardData = async () => {
  try {
    // TODO: دریافت داده‌ها از API
    // const response = await $fetch('/api/admin/tax/dashboard-stats')
    
    // داده‌های نمونه
    totalTaxCollected.value = 1250000000 // 1.25 میلیارد تومان
    monthlyTax.value = 45000000 // 45 میلیون تومان
    quarterlyTax.value = 135000000 // 135 میلیون تومان
    yearlyTax.value = 540000000 // 540 میلیون تومان
    growthRate.value = 12.5
    monthlyGrowth.value = 8.3
    quarterlyGrowth.value = 15.7
    yearlyGrowth.value = 22.1
  } catch (error) {
    console.error('خطا در بارگذاری آمار داشبورد:', error)
  }
}

// بارگذاری اولیه
onMounted(() => {
  loadDashboardData()
})
</script>

<!--
  کامپوننت داشبورد مالیات
  شامل:
  1. کارت کل مالیات جمع‌آوری شده
  2. کارت مالیات ماهانه
  3. کارت مالیات فصلی
  4. کارت مالیات سالانه
  هر کارت شامل:
  - آیکون و عنوان
  - مبلغ اصلی
  - درصد رشد
  - مقایسه با دوره قبل
  طراحی مدرن و کاملاً ریسپانسیو
--> 
