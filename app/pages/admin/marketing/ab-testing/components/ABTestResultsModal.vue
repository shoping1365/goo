<template>
  <div v-if="isOpen" class="fixed inset-0 z-50 overflow-y-auto">
    <div class="flex items-center justify-center min-h-screen px-4 pt-4 pb-20 text-center sm:block sm:p-0">
      <!-- Overlay -->
      <div class="fixed inset-0 transition-opacity bg-gray-500 bg-opacity-75" @click="$emit('close')"></div>

      <!-- Modal -->
      <div class="inline-block w-full max-w-6xl px-6 py-6 my-8 overflow-hidden text-right transition-all transform bg-white rounded-lg shadow-xl align-middle">
        <!-- Header -->
        <div class="flex items-center justify-between mb-6">
          <h3 class="text-lg font-medium text-gray-900">نتایج و تحلیل تست</h3>
          <div class="flex items-center space-x-4 space-x-reverse">
            <button @click="exportResults" class="px-3 py-1 text-sm bg-green-100 text-green-700 rounded-lg hover:bg-green-200">
              صادرات
            </button>
            <button @click="$emit('close')" class="text-gray-400 hover:text-gray-600">
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>
        </div>

        <!-- Content -->
        <div class="space-y-6">
          <!-- آمار تفصیلی -->
          <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
            <div class="bg-blue-50 rounded-lg p-6">
              <div class="text-sm text-blue-600 mb-1">بازدیدکنندگان A</div>
              <div class="text-2xl font-bold text-blue-900">{{ formatNumber(results.visitorsA) }}</div>
              <div class="text-xs text-blue-600 mt-1">{{ results.trafficSplitA }}% ترافیک</div>
            </div>
            <div class="bg-green-50 rounded-lg p-6">
              <div class="text-sm text-green-600 mb-1">بازدیدکنندگان B</div>
              <div class="text-2xl font-bold text-green-900">{{ formatNumber(results.visitorsB) }}</div>
              <div class="text-xs text-green-600 mt-1">{{ results.trafficSplitB }}% ترافیک</div>
            </div>
            <div class="bg-purple-50 rounded-lg p-6">
              <div class="text-sm text-purple-600 mb-1">نرخ تبدیل A</div>
              <div class="text-2xl font-bold text-purple-900">{{ results.conversionRateA }}%</div>
              <div class="text-xs text-purple-600 mt-1">{{ formatNumber(results.conversionsA) }} تبدیل</div>
            </div>
            <div class="bg-orange-50 rounded-lg p-6">
              <div class="text-sm text-orange-600 mb-1">نرخ تبدیل B</div>
              <div class="text-2xl font-bold text-orange-900">{{ results.conversionRateB }}%</div>
              <div class="text-xs text-orange-600 mt-1">{{ formatNumber(results.conversionsB) }} تبدیل</div>
            </div>
          </div>

          <!-- تحلیل آماری -->
          <div class="bg-gray-50 rounded-lg p-6">
            <h4 class="text-lg font-semibold text-gray-900 mb-4">تحلیل آماری</h4>
            
            <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
              <!-- تست معناداری آماری -->
              <div class="bg-white rounded-lg p-6">
                <div class="flex items-center justify-between mb-2">
                  <span class="text-sm font-medium text-gray-700">معناداری آماری</span>
                  <span :class="results.statisticalSignificance ? 'text-green-600' : 'text-red-600'" class="text-sm font-bold">
                    {{ results.statisticalSignificance ? 'معنادار' : 'غیرمعنادار' }}
                  </span>
                </div>
                <div class="text-2xl font-bold text-gray-900">{{ results.pValue }}</div>
                <div class="text-xs text-gray-500 mt-1">P-Value</div>
              </div>

              <!-- فاصله اطمینان -->
              <div class="bg-white rounded-lg p-6">
                <div class="text-sm font-medium text-gray-700 mb-2">فاصله اطمینان</div>
                <div class="text-lg font-bold text-gray-900">{{ results.confidenceInterval.min }}% - {{ results.confidenceInterval.max }}%</div>
                <div class="text-xs text-gray-500 mt-1">95% اطمینان</div>
              </div>

              <!-- قدرت تست -->
              <div class="bg-white rounded-lg p-6">
                <div class="text-sm font-medium text-gray-700 mb-2">قدرت تست</div>
                <div class="text-2xl font-bold text-gray-900">{{ results.testPower }}%</div>
                <div class="text-xs text-gray-500 mt-1">احتمال تشخیص تفاوت</div>
              </div>
            </div>

            <!-- پیشنهاد برنده -->
            <div class="mt-6 p-6 bg-yellow-50 rounded-lg border border-yellow-200">
              <div class="flex items-center">
                <svg class="w-5 h-5 text-yellow-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                <div>
                  <div class="text-sm font-medium text-yellow-800">پیشنهاد برنده</div>
                  <div class="text-sm text-yellow-700">{{ results.winnerSuggestion }}</div>
                </div>
              </div>
            </div>
          </div>

          <!-- نمودارهای تعاملی -->
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <!-- نمودار روند نرخ تبدیل -->
            <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
              <h4 class="text-lg font-semibold text-gray-900 mb-4">روند نرخ تبدیل</h4>
              <div class="h-64 flex items-end justify-between space-x-2 space-x-reverse">
                <div v-for="(day, index) in conversionTrend" :key="index" class="flex-1 flex flex-col items-center">
                  <div class="w-full flex space-x-1 space-x-reverse">
                    <div class="flex-1 bg-blue-500 rounded-t" :style="{ height: `${day.versionA * 2}px` }"></div>
                    <div class="flex-1 bg-green-500 rounded-t" :style="{ height: `${day.versionB * 2}px` }"></div>
                  </div>
                  <span class="text-xs text-gray-500 mt-2">{{ day.date }}</span>
                </div>
              </div>
              <div class="flex justify-center space-x-4 space-x-reverse mt-4">
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

            <!-- نمودار مقایسه نسخه‌ها -->
            <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
              <h4 class="text-lg font-semibold text-gray-900 mb-4">مقایسه نسخه‌ها</h4>
              <div class="space-y-4">
                <div>
                  <div class="flex justify-between text-sm mb-1">
                    <span>نرخ تبدیل</span>
                    <span>{{ results.conversionRateA }}% vs {{ results.conversionRateB }}%</span>
                  </div>
                  <div class="w-full bg-gray-200 rounded-full h-2">
                    <div class="bg-blue-500 h-2 rounded-full" :style="{ width: `${Math.max(results.conversionRateA, results.conversionRateB)}%` }"></div>
                  </div>
                </div>
                <div>
                  <div class="flex justify-between text-sm mb-1">
                    <span>درآمد</span>
                    <span>{{ formatCurrency(results.revenueA) }} vs {{ formatCurrency(results.revenueB) }}</span>
                  </div>
                  <div class="w-full bg-gray-200 rounded-full h-2">
                    <div class="bg-green-500 h-2 rounded-full" :style="{ width: `${Math.max(results.revenueA, results.revenueB) / 1000000}%` }"></div>
                  </div>
                </div>
                <div>
                  <div class="flex justify-between text-sm mb-1">
                    <span>زمان در صفحه</span>
                    <span>{{ results.avgTimeA }}s vs {{ results.avgTimeB }}s</span>
                  </div>
                  <div class="w-full bg-gray-200 rounded-full h-2">
                    <div class="bg-purple-500 h-2 rounded-full" :style="{ width: `${Math.max(results.avgTimeA, results.avgTimeB)}%` }"></div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- نمودار توزیع ترافیک -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
            <h4 class="text-lg font-semibold text-gray-900 mb-4">توزیع ترافیک در طول زمان</h4>
            <div class="h-64 flex items-end justify-between space-x-2 space-x-reverse">
              <div v-for="(hour, index) in trafficDistribution" :key="index" class="flex-1 flex flex-col items-center">
                <div class="w-full flex space-x-1 space-x-reverse">
                  <div class="flex-1 bg-blue-500 rounded-t" :style="{ height: `${hour.versionA}px` }"></div>
                  <div class="flex-1 bg-green-500 rounded-t" :style="{ height: `${hour.versionB}px` }"></div>
                </div>
                <span class="text-xs text-gray-500 mt-2">{{ hour.time }}</span>
              </div>
            </div>
          </div>

          <!-- نمودار درآمد در طول زمان -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
            <h4 class="text-lg font-semibold text-gray-900 mb-4">درآمد در طول زمان</h4>
            <div class="h-64 flex items-end justify-between space-x-2 space-x-reverse">
              <div v-for="(day, index) in revenueTrend" :key="index" class="flex-1 flex flex-col items-center">
                <div class="w-full flex space-x-1 space-x-reverse">
                  <div class="flex-1 bg-blue-500 rounded-t" :style="{ height: `${day.versionA / 10000}px` }"></div>
                  <div class="flex-1 bg-green-500 rounded-t" :style="{ height: `${day.versionB / 10000}px` }"></div>
                </div>
                <span class="text-xs text-gray-500 mt-2">{{ day.date }}</span>
              </div>
            </div>
          </div>

          <!-- آمار اضافی -->
          <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
            <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
              <div class="text-sm text-gray-500 mb-1">زمان متوسط در صفحه</div>
              <div class="text-lg font-bold text-gray-900">
                A: {{ results.avgTimeA }}s / B: {{ results.avgTimeB }}s
              </div>
            </div>
            <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
              <div class="text-sm text-gray-500 mb-1">نرخ پرش</div>
              <div class="text-lg font-bold text-gray-900">
                A: {{ results.bounceRateA }}% / B: {{ results.bounceRateB }}%
              </div>
            </div>
            <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
              <div class="text-sm text-gray-500 mb-1">درآمد کل</div>
              <div class="text-lg font-bold text-gray-900">
                A: {{ formatCurrency(results.revenueA) }} / B: {{ formatCurrency(results.revenueB) }}
              </div>
            </div>
          </div>
        </div>

        <!-- دکمه‌های عملیات -->
        <div class="flex items-center justify-end space-x-4 space-x-reverse pt-6 border-t mt-6">
          <button
            @click="$emit('close')"
            class="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-lg hover:bg-gray-50"
          >
            بستن
          </button>
          <button
            @click="exportResults"
            class="px-4 py-2 text-sm font-medium text-white bg-green-600 border border-transparent rounded-lg hover:bg-green-700"
          >
            صادرات نتایج
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
// Props
const props = defineProps<{
  isOpen: boolean
  testId?: number
}>()

// Events
const emit = defineEmits(['close'])

// داده‌های نتایج (mock data)
const results = ref({
  visitorsA: 1250,
  visitorsB: 1250,
  trafficSplitA: 50,
  trafficSplitB: 50,
  conversionRateA: 12.3,
  conversionRateB: 15.8,
  conversionsA: 154,
  conversionsB: 198,
  revenueA: 15400000,
  revenueB: 19800000,
  avgTimeA: 145,
  avgTimeB: 167,
  bounceRateA: 45.2,
  bounceRateB: 42.1,
  statisticalSignificance: true,
  pValue: 0.023,
  confidenceInterval: {
    min: 2.1,
    max: 4.9
  },
  testPower: 85,
  winnerSuggestion: 'نسخه B با 28.5% بهبود در نرخ تبدیل پیشنهاد می‌شود'
})

// داده‌های نمودار روند نرخ تبدیل
const conversionTrend = ref([
  { date: '1', versionA: 12.1, versionB: 15.2 },
  { date: '2', versionA: 12.5, versionB: 15.6 },
  { date: '3', versionA: 12.8, versionB: 15.9 },
  { date: '4', versionA: 12.2, versionB: 15.7 },
  { date: '5', versionA: 12.6, versionB: 16.1 },
  { date: '6', versionA: 12.4, versionB: 15.8 },
  { date: '7', versionA: 12.3, versionB: 15.8 }
])

// داده‌های نمودار توزیع ترافیک
const trafficDistribution = ref([
  { time: '00', versionA: 45, versionB: 48 },
  { time: '04', versionA: 32, versionB: 35 },
  { time: '08', versionA: 78, versionB: 82 },
  { time: '12', versionA: 95, versionB: 98 },
  { time: '16', versionA: 88, versionB: 92 },
  { time: '20', versionA: 65, versionB: 68 }
])

// داده‌های نمودار درآمد
const revenueTrend = ref([
  { date: '1', versionA: 1200000, versionB: 1500000 },
  { date: '2', versionA: 1250000, versionB: 1580000 },
  { date: '3', versionA: 1280000, versionB: 1590000 },
  { date: '4', versionA: 1220000, versionB: 1570000 },
  { date: '5', versionA: 1260000, versionB: 1610000 },
  { date: '6', versionA: 1240000, versionB: 1580000 },
  { date: '7', versionA: 1230000, versionB: 1580000 }
])

// فرمت کردن اعداد
const formatNumber = (num: number) => {
  return new Intl.NumberFormat('fa-IR').format(num)
}

// فرمت کردن ارز
const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR', {
    style: 'currency',
    currency: 'IRR',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(amount)
}

// صادرات نتایج
const exportResults = () => {
  // منطق صادرات به Excel/PDF/CSV
  console.log('صادرات نتایج تست:', props.testId)
  alert('نتایج با موفقیت صادر شد')
}
</script> 
