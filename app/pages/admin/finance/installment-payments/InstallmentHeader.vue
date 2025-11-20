<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
    <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between">
      <div class="flex items-center space-x-4 space-x-reverse">
        <div class="flex-shrink-0">
          <div class="w-12 h-12 bg-gradient-to-r from-blue-500 to-purple-600 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
        <div>
          <h1 class="text-2xl font-bold text-gray-900">خرید اقساطی</h1>
          <p class="text-sm text-gray-500 mt-1">مدیریت خریدهای اقساطی و اعتبارسنجی مشتریان</p>
        </div>
      </div>
      
      <div class="mt-4 sm:mt-0 flex flex-col sm:flex-row space-y-2 sm:space-y-0 sm:space-x-2 sm:space-x-reverse">
        <button
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          @click="openNewInstallmentModal"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
          </svg>
          درخواست جدید
        </button>
        
        <button
          class="inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          @click="openCreditCheckModal"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"></path>
          </svg>
          اعتبارسنجی
        </button>
        
        <button
          class="inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          @click="exportData"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
          </svg>
          خروجی
        </button>
      </div>
    </div>
    
    <!-- آمار سریع -->
    <div class="mt-6 grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
      <div class="bg-blue-50 border border-blue-200 rounded-lg p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <svg class="w-8 h-8 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-blue-600">درخواست‌های تایید شده</p>
            <p class="text-2xl font-bold text-blue-900">{{ stats.approved }}</p>
          </div>
        </div>
      </div>
      
      <div class="bg-yellow-50 border border-yellow-200 rounded-lg p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <svg class="w-8 h-8 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-yellow-600">در انتظار بررسی</p>
            <p class="text-2xl font-bold text-yellow-900">{{ stats.pending }}</p>
          </div>
        </div>
      </div>
      
      <div class="bg-red-50 border border-red-200 rounded-lg p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <svg class="w-8 h-8 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-red-600">معوق</p>
            <p class="text-2xl font-bold text-red-900">{{ stats.overdue }}</p>
          </div>
        </div>
      </div>
      
      <div class="bg-green-50 border border-green-200 rounded-lg p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <svg class="w-8 h-8 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-green-600">تکمیل شده</p>
            <p class="text-2xl font-bold text-green-900">{{ stats.completed }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'

const stats = ref({
  approved: 0,
  pending: 0,
  overdue: 0,
  completed: 0
})

// متدها
const openNewInstallmentModal = () => {
  // باز کردن مودال درخواست جدید
  console.log('باز کردن مودال درخواست جدید')
}

const openCreditCheckModal = () => {
  // باز کردن مودال اعتبارسنجی
  console.log('باز کردن مودال اعتبارسنجی')
}

const exportData = () => {
  // خروجی گرفتن از داده‌ها
  console.log('خروجی گرفتن از داده‌ها')
}

// دریافت آمار
const fetchStats = async () => {
  try {
    interface StatsData {
      approved: number
      pending: number
      overdue: number
      completed: number
    }
    interface ApiResponse {
      data?: StatsData
    }
    const response = await $fetch<ApiResponse>('/api/admin/installment-payments/stats', {
      method: 'GET'
    })
    stats.value = response.data || {
      approved: 120,
      pending: 45,
      overdue: 23,
      completed: 89
    }
  } catch (error) {
    console.error('خطا در دریافت آمار:', error)
  }
}

onMounted(() => {
  fetchStats()
})
</script> 
