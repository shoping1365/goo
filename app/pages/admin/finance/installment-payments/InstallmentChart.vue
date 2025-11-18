<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
    <div class="flex items-center justify-between mb-6">
      <h3 class="text-lg font-semibold text-gray-900">نمودارها و گزارش‌ها</h3>
      <div class="flex items-center space-x-2 space-x-reverse">
        <select v-model="chartType" class="text-sm border border-gray-300 rounded-md px-3 py-1">
          <option value="line">خطی</option>
          <option value="bar">ستونی</option>
          <option value="pie">دایره‌ای</option>
          <option value="area">ناحیه‌ای</option>
        </select>
        <select v-model="timeRange" class="text-sm border border-gray-300 rounded-md px-3 py-1">
          <option value="7">7 روز گذشته</option>
          <option value="30">30 روز گذشته</option>
          <option value="90">90 روز گذشته</option>
          <option value="365">یک سال گذشته</option>
        </select>
        <button @click="exportChart" class="text-blue-600 hover:text-blue-800">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
          </svg>
        </button>
      </div>
    </div>

    <!-- Chart Container -->
    <div class="mb-6">
      <div class="bg-gray-100 rounded-lg p-6 text-center h-96">
        <div class="text-gray-500 mb-4">
          <svg class="w-16 h-16 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
          </svg>
        </div>
        <p class="text-gray-600">نمودار {{ getChartTypeName() }}</p>
        <p class="text-sm text-gray-500 mt-2">داده‌های {{ timeRange }} روز گذشته</p>
      </div>
    </div>

    <!-- Chart Legend -->
    <div class="mb-6">
      <div class="flex items-center justify-center space-x-6 space-x-reverse">
        <div class="flex items-center">
          <div class="w-4 h-4 bg-blue-500 rounded mr-2"></div>
          <span class="text-sm text-gray-600">درخواست‌ها</span>
        </div>
        <div class="flex items-center">
          <div class="w-4 h-4 bg-green-500 rounded mr-2"></div>
          <span class="text-sm text-gray-600">تایید شده</span>
        </div>
        <div class="flex items-center">
          <div class="w-4 h-4 bg-red-500 rounded mr-2"></div>
          <span class="text-sm text-gray-600">رد شده</span>
        </div>
        <div class="flex items-center">
          <div class="w-4 h-4 bg-yellow-500 rounded mr-2"></div>
          <span class="text-sm text-gray-600">در انتظار</span>
        </div>
      </div>
    </div>

    <!-- Data Summary -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
      <div class="bg-blue-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-blue-600 mb-1">{{ chartData.totalRequests }}</div>
        <div class="text-sm text-gray-600">کل درخواست‌ها</div>
      </div>
      <div class="bg-green-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-green-600 mb-1">{{ chartData.approvedRequests }}</div>
        <div class="text-sm text-gray-600">تایید شده</div>
      </div>
      <div class="bg-red-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-red-600 mb-1">{{ chartData.rejectedRequests }}</div>
        <div class="text-sm text-gray-600">رد شده</div>
      </div>
      <div class="bg-yellow-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-yellow-600 mb-1">{{ chartData.pendingRequests }}</div>
        <div class="text-sm text-gray-600">در انتظار</div>
      </div>
    </div>

    <!-- Multiple Charts -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- Amount Chart -->
      <div class="bg-gray-50 rounded-lg p-6">
        <h4 class="text-md font-semibold text-gray-900 mb-4">نمودار مبالغ</h4>
        <div class="bg-white rounded p-6 h-64 flex items-center justify-center">
          <div class="text-center text-gray-500">
            <svg class="w-12 h-12 mx-auto mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1" />
            </svg>
            <p class="text-sm">نمودار مبالغ اقساط</p>
          </div>
        </div>
      </div>

      <!-- Credit Score Distribution -->
      <div class="bg-gray-50 rounded-lg p-6">
        <h4 class="text-md font-semibold text-gray-900 mb-4">توزیع امتیاز اعتباری</h4>
        <div class="bg-white rounded p-6 h-64 flex items-center justify-center">
          <div class="text-center text-gray-500">
            <svg class="w-12 h-12 mx-auto mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z" />
            </svg>
            <p class="text-sm">توزیع امتیازات</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Chart Controls -->
    <div class="mt-6 flex items-center justify-between">
      <div class="flex items-center space-x-4 space-x-reverse">
        <button @click="previousPeriod" class="text-gray-600 hover:text-gray-800">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
          </svg>
        </button>
        <span class="text-sm text-gray-600">{{ currentPeriodLabel }}</span>
        <button @click="nextPeriod" class="text-gray-600 hover:text-gray-800">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
          </svg>
        </button>
      </div>
      
      <div class="flex items-center space-x-2 space-x-reverse">
        <button @click="refreshChart" class="text-blue-600 hover:text-blue-800 text-sm font-medium">
          به‌روزرسانی
        </button>
        <button @click="fullScreen" class="text-green-600 hover:text-green-800 text-sm font-medium">
          تمام‌صفحه
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue'

const chartType = ref('line')
const timeRange = ref('30')

const chartData = ref({
  totalRequests: 1250,
  approvedRequests: 912,
  rejectedRequests: 203,
  pendingRequests: 135
})

const currentPeriodLabel = computed(() => {
  return `${timeRange.value} روز گذشته`
})

const getChartTypeName = (): string => {
  switch (chartType.value) {
    case 'line': return 'خطی'
    case 'bar': return 'ستونی'
    case 'pie': return 'دایره‌ای'
    case 'area': return 'ناحیه‌ای'
    default: return 'خطی'
  }
}

const exportChart = () => {
  // Export chart as image or PDF
  alert('صادر کردن نمودار')
}

const previousPeriod = () => {
  // Navigate to previous period
}

const nextPeriod = () => {
  // Navigate to next period
}

const refreshChart = () => {
  // Refresh chart data
}

const fullScreen = () => {
  // Open chart in full screen
}

onMounted(() => {
  // Load chart data
})

// Watch for changes
watch([chartType, timeRange], () => {
  // Refresh chart when filters change
})
</script>
