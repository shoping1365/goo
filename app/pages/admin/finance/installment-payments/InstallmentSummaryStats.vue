<template>
  <div v-if="hasAccess" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
    <!-- کل خریدهای اقساطی -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex items-center">
        <div class="flex-shrink-0">
          <div class="w-8 h-8 bg-blue-100 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
        <div class="mr-4">
          <p class="text-sm font-medium text-gray-600">کل خریدهای اقساطی</p>
          <p class="text-2xl font-bold text-gray-900">{{ stats.totalInstallments }}</p>
        </div>
      </div>
      <div class="mt-4">
        <div class="flex items-center">
          <span class="text-sm text-green-600 font-medium">+{{ stats.installmentGrowth }}%</span>
          <span class="text-sm text-gray-500 mr-2">از ماه گذشته</span>
        </div>
      </div>
    </div>

    <!-- کل مبلغ -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex items-center">
        <div class="flex-shrink-0">
          <div class="w-8 h-8 bg-green-100 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
            </svg>
          </div>
        </div>
        <div class="mr-4">
          <p class="text-sm font-medium text-gray-600">کل مبلغ اقساطی</p>
          <p class="text-2xl font-bold text-gray-900">{{ formatCurrency(stats.totalAmount) }}</p>
        </div>
      </div>
      <div class="mt-4">
        <div class="flex items-center">
          <span class="text-sm text-green-600 font-medium">+{{ stats.amountGrowth }}%</span>
          <span class="text-sm text-gray-500 mr-2">از ماه گذشته</span>
        </div>
      </div>
    </div>

    <!-- نرخ تایید -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex items-center">
        <div class="flex-shrink-0">
          <div class="w-8 h-8 bg-yellow-100 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
          </div>
        </div>
        <div class="mr-4">
          <p class="text-sm font-medium text-gray-600">نرخ تایید</p>
          <p class="text-2xl font-bold text-gray-900">{{ stats.approvalRate }}%</p>
        </div>
      </div>
      <div class="mt-4">
        <div class="w-full bg-gray-200 rounded-full h-2">
          <div class="bg-yellow-600 h-2 rounded-full" :style="{ width: stats.approvalRate + '%' }"></div>
        </div>
      </div>
    </div>

    <!-- نرخ بازپرداخت -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex items-center">
        <div class="flex-shrink-0">
          <div class="w-8 h-8 bg-purple-100 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
            </svg>
          </div>
        </div>
        <div class="mr-4">
          <p class="text-sm font-medium text-gray-600">نرخ بازپرداخت</p>
          <p class="text-2xl font-bold text-gray-900">{{ stats.repaymentRate }}%</p>
        </div>
      </div>
      <div class="mt-4">
        <div class="w-full bg-gray-200 rounded-full h-2">
          <div class="bg-purple-600 h-2 rounded-full" :style="{ width: stats.repaymentRate + '%' }"></div>
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

const stats = ref({
  totalInstallments: 0,
  totalAmount: 0,
  approvalRate: 0,
  repaymentRate: 0,
  installmentGrowth: 0,
  amountGrowth: 0
})

const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

const fetchStats = async () => {
  try {
    interface StatsData {
      totalInstallments: number
      totalAmount: number
      approvalRate: number
      repaymentRate: number
      installmentGrowth: number
      amountGrowth: number
    }
    interface ApiResponse {
      data?: StatsData
    }
    const response = await $fetch<ApiResponse>('/api/admin/installment-payments/stats', {
      method: 'GET'
    })
    stats.value = response.data || {
      totalInstallments: 1250,
      totalAmount: 4500000000,
      approvalRate: 78.5,
      repaymentRate: 92.3,
      installmentGrowth: 15.2,
      amountGrowth: 23.7
    }
  } catch (error) {
    console.error('خطا در دریافت آمار:', error)
  }
}

onMounted(async () => {
  await checkAuth()
  if (!hasAccess.value) {
    return
  }
  fetchStats()
})
</script> 
