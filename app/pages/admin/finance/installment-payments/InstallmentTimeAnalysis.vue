<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
    <div class="flex items-center justify-between mb-6">
      <h3 class="text-lg font-semibold text-gray-900">تحلیل زمانی</h3>
      <div class="flex items-center space-x-2 space-x-reverse">
        <select v-model="timeframe" class="text-sm border border-gray-300 rounded-md px-3 py-1">
          <option value="hourly">ساعتی</option>
          <option value="daily">روزانه</option>
          <option value="weekly">هفتگی</option>
          <option value="monthly">ماهانه</option>
        </select>
        <button class="text-blue-600 hover:text-blue-800" @click="refreshData">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
        </button>
      </div>
    </div>

    <!-- Summary Stats -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
      <div class="bg-blue-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-blue-600 mb-1">{{ analysis.peakHour }}</div>
        <div class="text-sm text-gray-600">ساعت اوج</div>
      </div>
      <div class="bg-green-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-green-600 mb-1">{{ analysis.avgProcessingTime }}h</div>
        <div class="text-sm text-gray-600">میانگین پردازش</div>
      </div>
      <div class="bg-yellow-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-yellow-600 mb-1">{{ analysis.busyDays }}</div>
        <div class="text-sm text-gray-600">روزهای شلوغ</div>
      </div>
      <div class="bg-purple-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-purple-600 mb-1">{{ analysis.seasonalTrend }}%</div>
        <div class="text-sm text-gray-600">روند فصلی</div>
      </div>
    </div>

    <!-- Hourly Pattern -->
    <div class="mb-6">
      <h4 class="text-md font-semibold text-gray-900 mb-4">الگوی ساعتی</h4>
      <div class="bg-gray-50 rounded-lg p-6">
        <div class="grid grid-cols-24 gap-1">
          <div v-for="hour in 24" :key="hour" class="text-center">
            <div 
              class="w-full bg-blue-200 rounded mb-1" 
              :style="{ height: getHourlyHeight(hour) + 'px' }"
            ></div>
            <span class="text-xs text-gray-500">{{ hour }}</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Time Analysis Data -->
    <div class="space-y-4">
      <div v-for="period in timeData" :key="period.label" class="bg-gray-50 rounded-lg p-6">
        <div class="flex items-center justify-between mb-3">
          <h5 class="text-md font-semibold text-gray-900">{{ period.label }}</h5>
          <span class="text-sm font-medium text-gray-600">{{ period.requests }} درخواست</span>
        </div>
        
        <div class="grid grid-cols-1 md:grid-cols-4 gap-6 text-sm">
          <div>
            <span class="text-gray-600 block">مبلغ کل</span>
            <span class="font-medium">{{ formatCurrency(period.totalAmount) }}</span>
          </div>
          <div>
            <span class="text-gray-600 block">نرخ تایید</span>
            <span class="font-medium">{{ period.approvalRate }}%</span>
          </div>
          <div>
            <span class="text-gray-600 block">زمان پردازش</span>
            <span class="font-medium">{{ period.processingTime }}h</span>
          </div>
          <div>
            <span class="text-gray-600 block">تغییرات</span>
            <span class="font-medium" :class="period.change > 0 ? 'text-green-600' : 'text-red-600'">
              {{ period.change > 0 ? '+' : '' }}{{ period.change }}%
            </span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'

const timeframe = ref('daily')
const analysis = ref({
  peakHour: '14:00',
  avgProcessingTime: 2.5,
  busyDays: 'دوشنبه-چهارشنبه',
  seasonalTrend: 15
})

const timeData = ref([
  {
    label: 'شنبه',
    requests: 150,
    totalAmount: 5000000000,
    approvalRate: 75,
    processingTime: 2.3,
    change: 12
  }
])

const formatCurrency = (amount: number): string => {
  return new Intl.NumberFormat('fa-IR', {
    style: 'currency',
    currency: 'IRR',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(amount)
}

const getHourlyHeight = (hour: number): number => {
  return Math.random() * 50 + 10
}

const refreshData = () => {
  // Refresh logic
}

onMounted(() => {
  refreshData()
})
</script>
