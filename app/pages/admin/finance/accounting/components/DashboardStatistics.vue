<template>
  <div class="space-y-6">
    <!-- هدر داشبورد -->
    <div class="flex items-center justify-between">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">داشبورد و آمار اتصال حسابداری</h3>
        <p class="text-sm text-gray-600 mt-1">نمای کلی وضعیت و عملکرد اتصال حسابداری</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          :disabled="isRefreshing"
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 disabled:bg-gray-400 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          @click="refreshStats"
        >
          <svg v-if="isRefreshing" class="w-4 h-4 ml-2 animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
          <svg v-else class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          {{ isRefreshing ? 'در حال بروزرسانی...' : 'بروزرسانی' }}
        </button>
      </div>
    </div>

    <!-- آمار کلی -->
    <div class="grid grid-cols-1 md:grid-cols-5 gap-6">
      <div class="bg-blue-50 rounded-lg p-6 flex flex-col items-center">
        <div class="text-2xl font-bold text-blue-600">{{ stats.activeConnections }}</div>
        <div class="text-sm text-blue-700 mt-1">اتصالات فعال</div>
      </div>
      <div class="bg-green-50 rounded-lg p-6 flex flex-col items-center">
        <div class="text-2xl font-bold text-green-600">{{ stats.syncedTransactions }}</div>
        <div class="text-sm text-green-700 mt-1">تراکنش‌های همگام‌سازی شده</div>
      </div>
      <div class="bg-yellow-50 rounded-lg p-6 flex flex-col items-center">
        <div class="text-2xl font-bold text-yellow-600">{{ stats.performanceScore }}%</div>
        <div class="text-sm text-yellow-700 mt-1">عملکرد اتصال</div>
      </div>
      <div class="bg-purple-50 rounded-lg p-6 flex flex-col items-center">
        <div class="text-2xl font-bold text-purple-600">{{ stats.healthStatus }}</div>
        <div class="text-sm text-purple-700 mt-1">وضعیت سلامت</div>
      </div>
      <div class="bg-red-50 rounded-lg p-6 flex flex-col items-center">
        <div class="text-2xl font-bold text-red-600">{{ stats.errorCount }}</div>
        <div class="text-sm text-red-700 mt-1">خطاها و مشکلات</div>
      </div>
    </div>

    <!-- نمودار عملکرد اتصال -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">نمودار عملکرد اتصال در ماه جاری</h5>
      <div class="w-full h-64">
        <canvas ref="performanceChart"></canvas>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
// @ts-ignore
import Chart from 'chart.js/auto'

// وضعیت بروزرسانی
const isRefreshing = ref(false)

// آمار نمونه
const stats = ref({
  activeConnections: 4,
  syncedTransactions: 1280,
  performanceScore: 97,
  healthStatus: 'سالم',
  errorCount: 3
})

// داده‌های نمودار عملکرد
const performanceData = ref({
  labels: [
    '۱', '۲', '۳', '۴', '۵', '۶', '۷', '۸', '۹', '۱۰', '۱۱', '۱۲', '۱۳', '۱۴', '۱۵', '۱۶', '۱۷', '۱۸', '۱۹', '۲۰', '۲۱', '۲۲', '۲۳', '۲۴', '۲۵', '۲۶', '۲۷', '۲۸', '۲۹', '۳۰', '۳۱'
  ],
  values: [97, 98, 96, 95, 97, 98, 99, 97, 96, 98, 97, 97, 98, 99, 97, 96, 97, 98, 99, 97, 98, 97, 97, 98, 99, 97, 96, 97, 98, 99, 97]
})

// رفرنس نمودار
const performanceChart = ref(null)
// eslint-disable-next-line @typescript-eslint/no-explicit-any
let chartInstance: any = null

// رسم نمودار
const renderChart = () => {
  if (chartInstance) {
    chartInstance.destroy()
  }
  // @ts-ignore
  chartInstance = new Chart(performanceChart.value, {
    type: 'line',
    data: {
      labels: performanceData.value.labels,
      datasets: [
        {
          label: 'عملکرد (%)',
          data: performanceData.value.values,
          borderColor: '#3B82F6',
          backgroundColor: 'rgba(59, 130, 246, 0.1)',
          fill: true,
          tension: 0.4
        }
      ]
    },
    options: {
      responsive: true,
      plugins: {
        legend: {
          display: false
        }
      },
      scales: {
        y: {
          min: 90,
          max: 100,
          ticks: {
            stepSize: 2
          }
        }
      }
    }
  })
}

onMounted(() => {
  renderChart()
})

// بروزرسانی آمار
const refreshStats = async () => {
  isRefreshing.value = true
  // TODO: دریافت آمار جدید از سرور
  await new Promise(resolve => setTimeout(resolve, 1500))
  isRefreshing.value = false
  // renderChart() // در صورت تغییر داده‌ها
}
</script>

<!--
  کامپوننت داشبورد و آمار اتصال حسابداری
  شامل:
  1. آمار اتصالات فعال
  2. آمار تراکنش‌های همگام‌سازی شده
  3. نمودار عملکرد اتصال
  4. وضعیت سلامت اتصال
  5. آمار خطاها و مشکلات
  طراحی مدرن و کاملاً ریسپانسیو
  توضیحات کامل به فارسی
--> 
