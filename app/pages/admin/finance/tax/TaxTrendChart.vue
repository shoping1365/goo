<template>
  <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
    <!-- هدر نمودار -->
    <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-6 mb-6">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">روند مالیات در طول زمان</h3>
        <p class="text-sm text-gray-600">نمودار تغییرات مالیات در ۶ ماه گذشته</p>
      </div>
      
      <!-- فیلترهای دوره -->
      <div class="flex gap-2">
        <button 
          v-for="period in periods" 
          :key="period.value"
          @click="selectedPeriod = period.value"
          :class="[
            'px-3 py-1.5 text-sm rounded-lg transition-colors duration-200',
            selectedPeriod === period.value 
              ? 'bg-blue-100 text-blue-700' 
              : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
          ]"
        >
          {{ period.label }}
        </button>
      </div>
    </div>

    <!-- نمودار -->
    <div class="relative">
      <!-- نمودار خطی ساده -->
      <div class="h-64 flex items-end justify-between gap-2">
        <div 
          v-for="(item, index) in chartData" 
          :key="index"
          class="flex-1 flex flex-col items-center"
        >
          <!-- ستون نمودار -->
          <div class="w-full max-w-16 relative">
            <div 
              class="bg-gradient-to-t from-blue-500 to-blue-400 rounded-t-lg transition-all duration-300 hover:from-blue-600 hover:to-blue-500"
              :style="{ height: `${(item.amount / maxAmount) * 200}px` }"
            ></div>
            
            <!-- مقدار روی ستون -->
            <div class="absolute -top-8 left-1/2 transform -translate-x-1/2 text-xs text-gray-600 font-medium">
              {{ formatCompactCurrency(item.amount) }}
            </div>
          </div>
          
          <!-- برچسب محور X -->
          <div class="mt-2 text-xs text-gray-500 text-center">
            {{ item.month }}
          </div>
        </div>
      </div>

      <!-- خط مرجع -->
      <div class="absolute bottom-16 left-0 right-0 h-px bg-gray-200"></div>
    </div>

    <!-- آمار اضافی -->
    <div class="mt-6 pt-6 border-t border-gray-200">
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <!-- میانگین ماهانه -->
        <div class="text-center">
          <div class="text-2xl font-bold text-gray-900">{{ formatCurrency(averageMonthly) }}</div>
          <div class="text-sm text-gray-500">میانگین ماهانه</div>
        </div>
        
        <!-- بیشترین مقدار -->
        <div class="text-center">
          <div class="text-2xl font-bold text-green-600">{{ formatCurrency(maxAmount) }}</div>
          <div class="text-sm text-gray-500">بیشترین مقدار</div>
        </div>
        
        <!-- کمترین مقدار -->
        <div class="text-center">
          <div class="text-2xl font-bold text-red-600">{{ formatCurrency(minAmount) }}</div>
          <div class="text-sm text-gray-500">کمترین مقدار</div>
        </div>
      </div>
    </div>

    <!-- تحلیل روند -->
    <div class="mt-6 p-6 bg-blue-50 rounded-lg">
      <div class="flex items-start gap-3">
        <div class="w-8 h-8 bg-blue-100 rounded-lg flex items-center justify-center flex-shrink-0">
          <svg class="w-4 h-4 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"></path>
          </svg>
        </div>
        <div>
          <h4 class="font-medium text-blue-900">تحلیل روند</h4>
          <p class="text-sm text-blue-700 mt-1">
            {{ trendAnalysis }}
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue'

// فیلترهای دوره
const periods = [
  { label: '۶ ماه', value: '6months' },
  { label: '۳ ماه', value: '3months' },
  { label: '۱ ماه', value: '1month' }
]

const selectedPeriod = ref('6months')

// داده‌های نمودار
const chartData = ref([
  { month: 'فروردین', amount: 42000000 },
  { month: 'اردیبهشت', amount: 45000000 },
  { month: 'خرداد', amount: 48000000 },
  { month: 'تیر', amount: 52000000 },
  { month: 'مرداد', amount: 49000000 },
  { month: 'شهریور', amount: 51000000 }
])

// محاسبه آمار
const maxAmount = computed(() => Math.max(...chartData.value.map(item => item.amount)))
const minAmount = computed(() => Math.min(...chartData.value.map(item => item.amount)))
const averageMonthly = computed(() => {
  const sum = chartData.value.reduce((acc, item) => acc + item.amount, 0)
  return Math.round(sum / chartData.value.length)
})

// تحلیل روند
const trendAnalysis = computed(() => {
  const firstAmount = chartData.value[0].amount
  const lastAmount = chartData.value[chartData.value.length - 1].amount
  const growth = ((lastAmount - firstAmount) / firstAmount) * 100
  
  if (growth > 0) {
    return `مالیات در این دوره ${growth.toFixed(1)}% رشد داشته است. روند کلی صعودی است.`
  } else if (growth < 0) {
    return `مالیات در این دوره ${Math.abs(growth).toFixed(1)}% کاهش داشته است. روند کلی نزولی است.`
  } else {
    return 'مالیات در این دوره ثابت بوده است. روند کلی پایدار است.'
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

// فرمت کردن مبلغ فشرده
const formatCompactCurrency = (amount: number): string => {
  if (amount >= 1000000000) {
    return `${(amount / 1000000000).toFixed(1)} میلیارد`
  } else if (amount >= 1000000) {
    return `${(amount / 1000000).toFixed(1)} میلیون`
  } else if (amount >= 1000) {
    return `${(amount / 1000).toFixed(1)} هزار`
  }
  return amount.toString()
}

// بارگذاری داده‌های نمودار
const loadChartData = async () => {
  try {
    // TODO: دریافت داده‌ها از API بر اساس دوره انتخاب شده
    // const response = await $fetch(`/api/admin/tax/trend?period=${selectedPeriod.value}`)
    // chartData.value = response.data
  } catch (error) {
    console.error('خطا در بارگذاری داده‌های نمودار:', error)
  }
}

// نظارت بر تغییر دوره
watch(selectedPeriod, () => {
  loadChartData()
})

// بارگذاری اولیه
onMounted(() => {
  loadChartData()
})
</script>

<!--
  کامپوننت نمودار روند مالیات
  شامل:
  1. نمودار خطی/ستونی تعاملی
  2. فیلترهای دوره زمانی
  3. آمار اضافی (میانگین، بیشترین، کمترین)
  4. تحلیل روند خودکار
  5. طراحی مدرن و کاملاً ریسپانسیو
--> 
