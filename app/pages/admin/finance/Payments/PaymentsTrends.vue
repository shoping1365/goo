<template>
  <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
    <div class="flex items-center justify-between mb-6">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">روند ساعتی پرداخت‌ها</h3>
        <p class="text-sm text-gray-500 mt-1">Hourly Payment Trends</p>
      </div>
    </div>

    <!-- نمودار روند ساعتی -->
    <div class="h-48 bg-gray-50 rounded-lg flex items-center justify-center mb-4">
      <div class="text-center">
        <svg class="w-12 h-12 text-gray-300 mx-auto mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
        </svg>
        <p class="text-sm text-gray-500">نمودار روند ساعتی</p>
      </div>
    </div>

    <!-- آمار ساعتی -->
    <div class="grid grid-cols-6 gap-2">
      <div 
        v-for="hour in hourlyData" 
        :key="hour.hour"
        class="text-center p-2 bg-gray-50 rounded-lg"
      >
        <div class="text-xs text-gray-500 mb-1">{{ hour.hour }}:۰۰</div>
        <div class="text-sm font-medium text-gray-700">{{ formatPrice(hour.amount) }}</div>
        <div class="text-xs text-gray-500">{{ hour.count }} تراکنش</div>
      </div>
    </div>

    <!-- آمار خلاصه -->
    <div class="mt-4 pt-4 border-t border-gray-200">
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div class="text-center">
          <div class="text-lg font-bold text-blue-600">{{ formatPrice(peakHourAmount) }}</div>
          <div class="text-sm text-blue-500">بیشترین مبلغ</div>
        </div>
        <div class="text-center">
          <div class="text-lg font-bold text-green-600">{{ peakHourTime }}</div>
          <div class="text-sm text-green-500">ساعت اوج</div>
        </div>
        <div class="text-center">
          <div class="text-lg font-bold text-purple-600">{{ formatPrice(averageHourlyAmount) }}</div>
          <div class="text-sm text-purple-500">میانگین ساعتی</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

// داده‌های ساعتی
const hourlyData = ref([
  { hour: '00', amount: 450000, count: 2 },
  { hour: '01', amount: 320000, count: 1 },
  { hour: '02', amount: 280000, count: 1 },
  { hour: '03', amount: 250000, count: 1 },
  { hour: '04', amount: 310000, count: 1 },
  { hour: '05', amount: 420000, count: 2 },
  { hour: '06', amount: 680000, count: 3 },
  { hour: '07', amount: 890000, count: 4 },
  { hour: '08', amount: 1240000, count: 6 },
  { hour: '09', amount: 1560000, count: 8 },
  { hour: '10', amount: 1890000, count: 10 },
  { hour: '11', amount: 1670000, count: 9 },
  { hour: '12', amount: 1450000, count: 7 },
  { hour: '13', amount: 1340000, count: 6 },
  { hour: '14', amount: 1780000, count: 9 },
  { hour: '15', amount: 1980000, count: 11 },
  { hour: '16', amount: 2340000, count: 13 },
  { hour: '17', amount: 2670000, count: 15 },
  { hour: '18', amount: 2890000, count: 16 },
  { hour: '19', amount: 2980000, count: 17 },
  { hour: '20', amount: 2760000, count: 15 },
  { hour: '21', amount: 2340000, count: 13 },
  { hour: '22', amount: 1670000, count: 9 },
  { hour: '23', amount: 980000, count: 5 }
])

// محاسبه آمار
const peakHourAmount = computed(() => {
  return Math.max(...hourlyData.value.map(hour => hour.amount))
})

const peakHourTime = computed(() => {
  const peakHour = hourlyData.value.find(hour => hour.amount === peakHourAmount.value)
  return peakHour ? `${peakHour.hour}:۰۰` : '-'
})

const averageHourlyAmount = computed(() => {
  const total = hourlyData.value.reduce((sum, hour) => sum + hour.amount, 0)
  return total / hourlyData.value.length
})

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
  کامپوننت روند ساعتی پرداخت‌ها
  - نمایش روند پرداخت‌ها در ساعات مختلف
  - آمار خلاصه
  - کاملاً ریسپانسیو
  توضیحات کامل به فارسی در کد
--> 
