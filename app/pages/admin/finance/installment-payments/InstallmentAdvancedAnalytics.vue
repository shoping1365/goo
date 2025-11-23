<template>
  <div v-if="hasAccess" class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
    <div class="flex items-center justify-between mb-6">
      <h3 class="text-lg font-semibold text-gray-900">تحلیل‌های پیشرفته</h3>
      <div class="flex items-center space-x-2 space-x-reverse">
        <select v-model="selectedPeriod" class="text-sm border border-gray-300 rounded-md px-3 py-1">
          <option value="7">7 روز گذشته</option>
          <option value="30">30 روز گذشته</option>
          <option value="90">90 روز گذشته</option>
          <option value="365">یک سال گذشته</option>
        </select>
        <button class="text-blue-600 hover:text-blue-800" @click="refreshAnalytics">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
        </button>
      </div>
    </div>

    <!-- Key Metrics Grid -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-6">
      <div class="bg-gradient-to-r from-blue-50 to-blue-100 p-6 rounded-lg">
        <div class="flex items-center">
          <div class="p-2 bg-blue-500 rounded-lg">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1" />
            </svg>
          </div>
          <div class="mr-3">
            <p class="text-sm text-gray-600">میانگین مبلغ اقساط</p>
            <p class="text-xl font-bold text-gray-900">{{ formatCurrency(analytics.avgAmount) }}</p>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-green-50 to-green-100 p-6 rounded-lg">
        <div class="flex items-center">
          <div class="p-2 bg-green-500 rounded-lg">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <div class="mr-3">
            <p class="text-sm text-gray-600">نرخ تایید</p>
            <p class="text-xl font-bold text-gray-900">{{ analytics.approvalRate }}%</p>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-yellow-50 to-yellow-100 p-6 rounded-lg">
        <div class="flex items-center">
          <div class="p-2 bg-yellow-500 rounded-lg">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" />
            </svg>
          </div>
          <div class="mr-3">
            <p class="text-sm text-gray-600">میانگین امتیاز اعتباری</p>
            <p class="text-xl font-bold text-gray-900">{{ analytics.avgCreditScore }}/100</p>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-purple-50 to-purple-100 p-6 rounded-lg">
        <div class="flex items-center">
          <div class="p-2 bg-purple-500 rounded-lg">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
            </svg>
          </div>
          <div class="mr-3">
            <p class="text-sm text-gray-600">مشتریان جدید</p>
            <p class="text-xl font-bold text-gray-900">{{ analytics.newCustomers }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Risk Analysis -->
    <div class="mb-6">
      <h4 class="text-md font-semibold text-gray-900 mb-4">تحلیل ریسک</h4>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div class="bg-red-50 border border-red-200 rounded-lg p-6">
          <div class="flex items-center justify-between mb-2">
            <span class="text-sm font-medium text-red-800">ریسک بالا</span>
            <span class="text-lg font-bold text-red-600">{{ analytics.highRiskCount }}</span>
          </div>
          <div class="w-full bg-red-200 rounded-full h-2">
            <div class="bg-red-500 h-2 rounded-full" :style="{ width: analytics.highRiskPercentage + '%' }"></div>
          </div>
          <p class="text-xs text-red-600 mt-1">{{ analytics.highRiskPercentage }}% از کل درخواست‌ها</p>
        </div>

        <div class="bg-yellow-50 border border-yellow-200 rounded-lg p-6">
          <div class="flex items-center justify-between mb-2">
            <span class="text-sm font-medium text-yellow-800">ریسک متوسط</span>
            <span class="text-lg font-bold text-yellow-600">{{ analytics.mediumRiskCount }}</span>
          </div>
          <div class="w-full bg-yellow-200 rounded-full h-2">
            <div class="bg-yellow-500 h-2 rounded-full" :style="{ width: analytics.mediumRiskPercentage + '%' }"></div>
          </div>
          <p class="text-xs text-yellow-600 mt-1">{{ analytics.mediumRiskPercentage }}% از کل درخواست‌ها</p>
        </div>

        <div class="bg-green-50 border border-green-200 rounded-lg p-6">
          <div class="flex items-center justify-between mb-2">
            <span class="text-sm font-medium text-green-800">ریسک پایین</span>
            <span class="text-lg font-bold text-green-600">{{ analytics.lowRiskCount }}</span>
          </div>
          <div class="w-full bg-green-200 rounded-full h-2">
            <div class="bg-green-500 h-2 rounded-full" :style="{ width: analytics.lowRiskPercentage + '%' }"></div>
          </div>
          <p class="text-xs text-green-600 mt-1">{{ analytics.lowRiskPercentage }}% از کل درخواست‌ها</p>
        </div>
      </div>
    </div>

    <!-- Performance Trends -->
    <div class="mb-6">
      <h4 class="text-md font-semibold text-gray-900 mb-4">روند عملکرد</h4>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="bg-gray-50 rounded-lg p-6">
          <div class="flex items-center justify-between mb-2">
            <span class="text-sm text-gray-600">درخواست‌های روزانه</span>
            <span class="text-sm font-medium" :class="analytics.dailyRequestsTrend > 0 ? 'text-green-600' : 'text-red-600'">
              {{ analytics.dailyRequestsTrend > 0 ? '+' : '' }}{{ analytics.dailyRequestsTrend }}%
            </span>
          </div>
          <div class="flex items-center space-x-2 space-x-reverse">
            <div class="flex-1 bg-gray-200 rounded-full h-2">
              <div class="bg-blue-500 h-2 rounded-full" :style="{ width: analytics.dailyRequestsProgress + '%' }"></div>
            </div>
            <span class="text-sm text-gray-500">{{ analytics.dailyRequests }}/{{ analytics.dailyRequestsTarget }}</span>
          </div>
        </div>

        <div class="bg-gray-50 rounded-lg p-6">
          <div class="flex items-center justify-between mb-2">
            <span class="text-sm text-gray-600">مبلغ تایید شده</span>
            <span class="text-sm font-medium" :class="analytics.approvedAmountTrend > 0 ? 'text-green-600' : 'text-red-600'">
              {{ analytics.approvedAmountTrend > 0 ? '+' : '' }}{{ analytics.approvedAmountTrend }}%
            </span>
          </div>
          <div class="flex items-center space-x-2 space-x-reverse">
            <div class="flex-1 bg-gray-200 rounded-full h-2">
              <div class="bg-green-500 h-2 rounded-full" :style="{ width: analytics.approvedAmountProgress + '%' }"></div>
            </div>
            <span class="text-sm text-gray-500">{{ formatCurrency(analytics.approvedAmount) }}</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Top Performing Categories -->
    <div class="mb-6">
      <h4 class="text-md font-semibold text-gray-900 mb-4">دسته‌بندی‌های برتر</h4>
      <div class="space-y-3">
        <div v-for="category in analytics.topCategories" :key="category.name" class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
          <div class="flex items-center">
            <div class="w-3 h-3 rounded-full mr-3" :style="{ backgroundColor: category.color }"></div>
            <span class="text-sm font-medium text-gray-900">{{ category.name }}</span>
          </div>
          <div class="flex items-center space-x-4 space-x-reverse">
            <span class="text-sm text-gray-600">{{ category.count }} درخواست</span>
            <span class="text-sm font-medium text-gray-900">{{ formatCurrency(category.amount) }}</span>
            <span class="text-sm" :class="category.growth > 0 ? 'text-green-600' : 'text-red-600'">
              {{ category.growth > 0 ? '+' : '' }}{{ category.growth }}%
            </span>
          </div>
        </div>
      </div>
    </div>

    <!-- Geographic Distribution -->
    <div class="mb-6">
      <h4 class="text-md font-semibold text-gray-900 mb-4">توزیع جغرافیایی</h4>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div v-for="region in analytics.geographicDistribution" :key="region.name" class="bg-gray-50 rounded-lg p-6">
          <div class="flex items-center justify-between mb-2">
            <span class="text-sm font-medium text-gray-900">{{ region.name }}</span>
            <span class="text-sm text-gray-600">{{ region.percentage }}%</span>
          </div>
          <div class="w-full bg-gray-200 rounded-full h-2">
            <div class="bg-blue-500 h-2 rounded-full" :style="{ width: region.percentage + '%' }"></div>
          </div>
          <div class="flex justify-between text-xs text-gray-500 mt-1">
            <span>{{ region.count }} درخواست</span>
            <span>{{ formatCurrency(region.amount) }}</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Customer Insights -->
    <div>
      <h4 class="text-md font-semibold text-gray-900 mb-4">بینش مشتریان</h4>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div class="text-center p-6 bg-blue-50 rounded-lg">
          <div class="text-2xl font-bold text-blue-600 mb-1">{{ analytics.customerInsights.avgAge }}</div>
          <div class="text-sm text-gray-600">میانگین سن</div>
        </div>
        <div class="text-center p-6 bg-green-50 rounded-lg">
          <div class="text-2xl font-bold text-green-600 mb-1">{{ analytics.customerInsights.repeatCustomers }}%</div>
          <div class="text-sm text-gray-600">مشتریان تکراری</div>
        </div>
        <div class="text-center p-6 bg-purple-50 rounded-lg">
          <div class="text-2xl font-bold text-purple-600 mb-1">{{ analytics.customerInsights.avgInstallments }}</div>
          <div class="text-sm text-gray-600">میانگین تعداد اقساط</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>
</script>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue'
import { useAuth } from '~/composables/useAuth'

// احراز هویت
const { user, isAuthenticated } = useAuth()

// بررسی دسترسی admin
const hasAccess = computed(() => {
  if (!isAuthenticated.value) {
    return false
  }

  const userRole = user.value?.role?.toLowerCase() || ''
  const adminRoles = ['admin', 'developer']
  return adminRoles.includes(userRole)
})

// بررسی احراز هویت و دسترسی admin - نمایش 404 در صورت عدم دسترسی
const checkAuth = async () => {
  if (!hasAccess.value) {
    await navigateTo('/404', { external: false })
  }
}

// بررسی احراز هویت هنگام تغییر وضعیت احراز هویت
watch([isAuthenticated, hasAccess], async () => {
  if (!hasAccess.value) {
    await checkAuth()
  }
})

interface Category {
  name: string
  count: number
  amount: number
  growth: number
  color: string
}

interface Region {
  name: string
  count: number
  amount: number
  percentage: number
}

interface CustomerInsights {
  avgAge: number
  repeatCustomers: number
  avgInstallments: number
}

interface Analytics {
  avgAmount: number
  approvalRate: number
  avgCreditScore: number
  newCustomers: number
  highRiskCount: number
  highRiskPercentage: number
  mediumRiskCount: number
  mediumRiskPercentage: number
  lowRiskCount: number
  lowRiskPercentage: number
  dailyRequests: number
  dailyRequestsTarget: number
  dailyRequestsProgress: number
  dailyRequestsTrend: number
  approvedAmount: number
  approvedAmountProgress: number
  approvedAmountTrend: number
  topCategories: Category[]
  geographicDistribution: Region[]
  customerInsights: CustomerInsights
}

const selectedPeriod = ref('30')
const analytics = ref<Analytics>({
  avgAmount: 0,
  approvalRate: 0,
  avgCreditScore: 0,
  newCustomers: 0,
  highRiskCount: 0,
  highRiskPercentage: 0,
  mediumRiskCount: 0,
  mediumRiskPercentage: 0,
  lowRiskCount: 0,
  lowRiskPercentage: 0,
  dailyRequests: 0,
  dailyRequestsTarget: 0,
  dailyRequestsProgress: 0,
  dailyRequestsTrend: 0,
  approvedAmount: 0,
  approvedAmountProgress: 0,
  approvedAmountTrend: 0,
  topCategories: [],
  geographicDistribution: [],
  customerInsights: {
    avgAge: 0,
    repeatCustomers: 0,
    avgInstallments: 0
  }
})

const formatCurrency = (amount: number): string => {
  return new Intl.NumberFormat('fa-IR', {
    style: 'currency',
    currency: 'IRR',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(amount)
}

const fetchAnalytics = async () => {
  try {
    const response = await $fetch<{ data?: Analytics }>('/api/admin/installment-payments/analytics', {
      query: { period: selectedPeriod.value }
    })
    const data = response.data
    
    if (data) {
      analytics.value = data
    }
  } catch (error) {
    console.error('خطا در دریافت تحلیل‌ها:', error)
    // در حالت واقعی، اینجا از createError استفاده می‌کنیم
  }
}

const refreshAnalytics = () => {
  fetchAnalytics()
}

onMounted(async () => {
  await checkAuth()
  if (!hasAccess.value) {
    return
  }
  fetchAnalytics()
})

// Watch for period changes
watch(selectedPeriod, () => {
  fetchAnalytics()
})
</script> 
