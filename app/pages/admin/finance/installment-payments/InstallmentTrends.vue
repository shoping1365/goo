<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
    <div class="flex items-center justify-between mb-6">
      <h3 class="text-lg font-semibold text-gray-900">روند و تحلیل‌ها</h3>
      <div class="flex items-center space-x-2 space-x-reverse">
        <select v-model="selectedPeriod" class="text-sm border border-gray-300 rounded-md px-3 py-1">
          <option value="7">7 روز گذشته</option>
          <option value="30">30 روز گذشته</option>
          <option value="90">90 روز گذشته</option>
          <option value="365">یک سال گذشته</option>
        </select>
        <button @click="refreshTrends" class="text-blue-600 hover:text-blue-800">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
        </button>
      </div>
    </div>

    <!-- Trend Cards -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-6">
      <div class="bg-gradient-to-r from-blue-50 to-blue-100 p-6 rounded-lg">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm text-gray-600">درخواست‌های جدید</p>
            <p class="text-2xl font-bold text-gray-900">{{ trends.newRequests }}</p>
          </div>
          <div class="text-right">
            <div class="flex items-center" :class="trends.newRequestsChange >= 0 ? 'text-green-600' : 'text-red-600'">
              <svg v-if="trends.newRequestsChange >= 0" class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" />
              </svg>
              <svg v-else class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 17h8m0 0V9m0 8l-8-8-4 4-6-6" />
              </svg>
              <span class="text-sm font-medium">{{ Math.abs(trends.newRequestsChange) }}%</span>
            </div>
            <p class="text-xs text-gray-500">نسبت به دوره قبل</p>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-green-50 to-green-100 p-6 rounded-lg">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm text-gray-600">تایید شده</p>
            <p class="text-2xl font-bold text-gray-900">{{ trends.approvedRequests }}</p>
          </div>
          <div class="text-right">
            <div class="flex items-center" :class="trends.approvedRequestsChange >= 0 ? 'text-green-600' : 'text-red-600'">
              <svg v-if="trends.approvedRequestsChange >= 0" class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" />
              </svg>
              <svg v-else class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 17h8m0 0V9m0 8l-8-8-4 4-6-6" />
              </svg>
              <span class="text-sm font-medium">{{ Math.abs(trends.approvedRequestsChange) }}%</span>
            </div>
            <p class="text-xs text-gray-500">نسبت به دوره قبل</p>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-yellow-50 to-yellow-100 p-6 rounded-lg">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm text-gray-600">مبلغ کل</p>
            <p class="text-2xl font-bold text-gray-900">{{ formatCurrency(trends.totalAmount) }}</p>
          </div>
          <div class="text-right">
            <div class="flex items-center" :class="trends.totalAmountChange >= 0 ? 'text-green-600' : 'text-red-600'">
              <svg v-if="trends.totalAmountChange >= 0" class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" />
              </svg>
              <svg v-else class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 17h8m0 0V9m0 8l-8-8-4 4-6-6" />
              </svg>
              <span class="text-sm font-medium">{{ Math.abs(trends.totalAmountChange) }}%</span>
            </div>
            <p class="text-xs text-gray-500">نسبت به دوره قبل</p>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-purple-50 to-purple-100 p-6 rounded-lg">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm text-gray-600">نرخ تایید</p>
            <p class="text-2xl font-bold text-gray-900">{{ trends.approvalRate }}%</p>
          </div>
          <div class="text-right">
            <div class="flex items-center" :class="trends.approvalRateChange >= 0 ? 'text-green-600' : 'text-red-600'">
              <svg v-if="trends.approvalRateChange >= 0" class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" />
              </svg>
              <svg v-else class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 17h8m0 0V9m0 8l-8-8-4 4-6-6" />
              </svg>
              <span class="text-sm font-medium">{{ Math.abs(trends.approvalRateChange) }}%</span>
            </div>
            <p class="text-xs text-gray-500">نسبت به دوره قبل</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Daily Trends Chart -->
    <div class="mb-6">
      <h4 class="text-md font-semibold text-gray-900 mb-4">روند روزانه</h4>
      <div class="bg-gray-50 rounded-lg p-6">
        <div class="flex items-center justify-between mb-4">
          <div class="flex items-center space-x-4 space-x-reverse">
            <div class="flex items-center">
              <div class="w-3 h-3 bg-blue-500 rounded-full mr-2"></div>
              <span class="text-sm text-gray-600">درخواست‌ها</span>
            </div>
            <div class="flex items-center">
              <div class="w-3 h-3 bg-green-500 rounded-full mr-2"></div>
              <span class="text-sm text-gray-600">تایید شده</span>
            </div>
            <div class="flex items-center">
              <div class="w-3 h-3 bg-red-500 rounded-full mr-2"></div>
              <span class="text-sm text-gray-600">رد شده</span>
            </div>
          </div>
        </div>
        <div class="h-64 flex items-end justify-between space-x-2 space-x-reverse">
          <div v-for="(day, index) in trends.dailyData" :key="index" class="flex-1 flex flex-col items-center">
            <div class="w-full flex items-end justify-center space-x-1 space-x-reverse mb-2">
              <div class="bg-blue-500 rounded-t" :style="{ height: (day.requests / trends.maxRequests * 100) + '%', width: '4px' }"></div>
              <div class="bg-green-500 rounded-t" :style="{ height: (day.approved / trends.maxApproved * 100) + '%', width: '4px' }"></div>
              <div class="bg-red-500 rounded-t" :style="{ height: (day.rejected / trends.maxRejected * 100) + '%', width: '4px' }"></div>
            </div>
            <span class="text-xs text-gray-500">{{ day.date }}</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Amount Distribution -->
    <div class="mb-6">
      <h4 class="text-md font-semibold text-gray-900 mb-4">توزیع مبالغ</h4>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="bg-gray-50 rounded-lg p-6">
          <h5 class="text-sm font-medium text-gray-900 mb-3">بر اساس بازه مبلغ</h5>
          <div class="space-y-3">
            <div v-for="range in trends.amountRanges" :key="range.label" class="flex items-center justify-between">
              <span class="text-sm text-gray-600">{{ range.label }}</span>
              <div class="flex items-center space-x-2 space-x-reverse">
                <div class="w-20 bg-gray-200 rounded-full h-2">
                  <div class="bg-blue-500 h-2 rounded-full" :style="{ width: range.percentage + '%' }"></div>
                </div>
                <span class="text-sm font-medium text-gray-900">{{ range.percentage }}%</span>
              </div>
            </div>
          </div>
        </div>

        <div class="bg-gray-50 rounded-lg p-6">
          <h5 class="text-sm font-medium text-gray-900 mb-3">بر اساس تعداد اقساط</h5>
          <div class="space-y-3">
            <div v-for="installment in trends.installmentRanges" :key="installment.label" class="flex items-center justify-between">
              <span class="text-sm text-gray-600">{{ installment.label }}</span>
              <div class="flex items-center space-x-2 space-x-reverse">
                <div class="w-20 bg-gray-200 rounded-full h-2">
                  <div class="bg-green-500 h-2 rounded-full" :style="{ width: installment.percentage + '%' }"></div>
                </div>
                <span class="text-sm font-medium text-gray-900">{{ installment.percentage }}%</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Performance Metrics -->
    <div class="mb-6">
      <h4 class="text-md font-semibold text-gray-900 mb-4">معیارهای عملکرد</h4>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div class="bg-blue-50 rounded-lg p-6 text-center">
          <div class="text-2xl font-bold text-blue-600 mb-1">{{ trends.avgProcessingTime }}</div>
          <div class="text-sm text-gray-600">میانگین زمان پردازش (ساعت)</div>
        </div>
        <div class="bg-green-50 rounded-lg p-6 text-center">
          <div class="text-2xl font-bold text-green-600 mb-1">{{ trends.avgCreditScore }}</div>
          <div class="text-sm text-gray-600">میانگین امتیاز اعتباری</div>
        </div>
        <div class="bg-purple-50 rounded-lg p-6 text-center">
          <div class="text-2xl font-bold text-purple-600 mb-1">{{ trends.avgInstallmentAmount }}</div>
          <div class="text-sm text-gray-600">میانگین مبلغ هر قسط</div>
        </div>
      </div>
    </div>

    <!-- Top Performing Days -->
    <div>
      <h4 class="text-md font-semibold text-gray-900 mb-4">بهترین روزهای عملکرد</h4>
      <div class="bg-gray-50 rounded-lg p-6">
        <div class="space-y-3">
          <div v-for="(day, index) in trends.topDays" :key="index" class="flex items-center justify-between p-3 bg-white rounded-lg">
            <div class="flex items-center">
              <div class="w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center mr-3">
                <span class="text-sm font-medium text-blue-600">{{ index + 1 }}</span>
              </div>
              <div>
                <div class="text-sm font-medium text-gray-900">{{ day.date }}</div>
                <div class="text-xs text-gray-500">{{ day.dayName }}</div>
              </div>
            </div>
            <div class="text-right">
              <div class="text-sm font-medium text-gray-900">{{ day.requests }} درخواست</div>
              <div class="text-xs text-green-600">{{ day.approvalRate }}% تایید</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref, watch } from 'vue'

interface DailyData {
  date: string
  requests: number
  approved: number
  rejected: number
}

interface AmountRange {
  label: string
  percentage: number
}

interface InstallmentRange {
  label: string
  percentage: number
}

interface TopDay {
  date: string
  dayName: string
  requests: number
  approvalRate: number
}

interface Trends {
  newRequests: number
  newRequestsChange: number
  approvedRequests: number
  approvedRequestsChange: number
  totalAmount: number
  totalAmountChange: number
  approvalRate: number
  approvalRateChange: number
  dailyData: DailyData[]
  maxRequests: number
  maxApproved: number
  maxRejected: number
  amountRanges: AmountRange[]
  installmentRanges: InstallmentRange[]
  avgProcessingTime: number
  avgCreditScore: number
  avgInstallmentAmount: number
  topDays: TopDay[]
}

const selectedPeriod = ref('30')
const trends = ref<Trends>({
  newRequests: 0,
  newRequestsChange: 0,
  approvedRequests: 0,
  approvedRequestsChange: 0,
  totalAmount: 0,
  totalAmountChange: 0,
  approvalRate: 0,
  approvalRateChange: 0,
  dailyData: [],
  maxRequests: 0,
  maxApproved: 0,
  maxRejected: 0,
  amountRanges: [],
  installmentRanges: [],
  avgProcessingTime: 0,
  avgCreditScore: 0,
  avgInstallmentAmount: 0,
  topDays: []
})

const formatCurrency = (amount: number): string => {
  return new Intl.NumberFormat('fa-IR', {
    style: 'currency',
    currency: 'IRR',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(amount)
}

const fetchTrends = async () => {
  try {
    interface TrendsData {
      [key: string]: unknown
    }
    interface ApiResponse {
      data?: TrendsData
    }
    const response = await $fetch<ApiResponse>('/api/admin/installment-payments/trends', {
      query: { period: selectedPeriod.value }
    })
    
    if (response.data) {
      trends.value = response.data as unknown as Trends
    }
  } catch (error) {
    console.error('خطا در دریافت روندها:', error)
    // در حالت واقعی، اینجا از createError استفاده می‌کنیم
  }
}

const refreshTrends = () => {
  fetchTrends()
}

onMounted(() => {
  fetchTrends()
})

// Watch for period changes
watch(selectedPeriod, () => {
  fetchTrends()
})
</script> 
