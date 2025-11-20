<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200">
    <!-- هدر بخش -->
    <div class="px-6 py-4 border-b border-gray-200">
      <div class="flex items-center justify-between">
        <div>
          <h3 class="text-lg font-semibold text-gray-900">گزارش‌گیری و نظارت</h3>
          <p class="text-sm text-gray-600 mt-1">گزارش‌گیری و نظارت بر عملکرد اتصال حسابداری</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button
            class="inline-flex items-center px-3 py-2 border border-gray-300 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
            @click="refreshData"
          >
            <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
            </svg>
            بروزرسانی
          </button>
          <button
            class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
            @click="exportReport"
          >
            <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
            </svg>
            صادر کردن گزارش
          </button>
        </div>
      </div>
    </div>

    <!-- تب‌های بخش‌های مختلف -->
    <div class="border-b border-gray-200">
      <nav class="-mb-px flex space-x-8 space-x-reverse px-6">
        <button
          v-for="tab in tabs"
          :key="tab.id"
          :class="[
            activeTab === tab.id
              ? 'border-blue-500 text-blue-600'
              : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300',
            'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm'
          ]"
          @click="activeTab = tab.id"
        >
          {{ tab.name }}
        </button>
      </nav>
    </div>

    <!-- محتوای تب‌ها -->
    <div class="p-6">
      <!-- تب گزارش وضعیت اتصال -->
      <div v-if="activeTab === 'connection'" class="space-y-6">
        <ConnectionStatusReport />
      </div>

      <!-- تب گزارش تراکنش‌های همگام‌سازی شده -->
      <div v-if="activeTab === 'transactions'" class="space-y-6">
        <SyncedTransactionsReport />
      </div>

      <!-- تب گزارش خطاهای همگام‌سازی -->
      <div v-if="activeTab === 'errors'" class="space-y-6">
        <SyncErrorsReport />
      </div>

      <!-- تب لاگ فعالیت‌های اتصال -->
      <div v-if="activeTab === 'logs'" class="space-y-6">
        <ConnectionActivityLogs />
      </div>

      <!-- تب نظارت بر عملکرد اتصال -->
      <div v-if="activeTab === 'monitoring'" class="space-y-6">
        <ConnectionPerformanceMonitoring />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// Import کامپوننت‌های زیربخش‌ها
// @ts-ignore
import ConnectionStatusReport from './reporting/ConnectionStatusReport.vue'
// @ts-ignore
import SyncedTransactionsReport from './reporting/SyncedTransactionsReport.vue'
// @ts-ignore
import SyncErrorsReport from './reporting/SyncErrorsReport.vue'
// @ts-ignore
import ConnectionActivityLogs from './reporting/ConnectionActivityLogs.vue'
// @ts-ignore
import ConnectionPerformanceMonitoring from './reporting/ConnectionPerformanceMonitoring.vue'

// تعریف تب‌های مختلف
const tabs = ref([
  { id: 'connection', name: 'وضعیت اتصال' },
  { id: 'transactions', name: 'تراکنش‌های همگام شده' },
  { id: 'errors', name: 'خطاهای همگام‌سازی' },
  { id: 'logs', name: 'لاگ فعالیت‌ها' },
  { id: 'monitoring', name: 'نظارت بر عملکرد' }
])

// تب فعال
const activeTab = ref('connection')

// بروزرسانی داده‌ها
const refreshData = () => {
  // TODO: بروزرسانی داده‌ها از سرور
  console.log('بروزرسانی داده‌های گزارش‌گیری و نظارت')
}

// صادر کردن گزارش
const exportReport = () => {
  // TODO: صادر کردن گزارش
  console.log('صادر کردن گزارش')
}
</script>

<!--
  کامپوننت گزارش‌گیری و نظارت
  شامل ۵ بخش اصلی:
  1. گزارش وضعیت اتصال
  2. گزارش تراکنش‌های همگام‌سازی شده
  3. گزارش خطاهای همگام‌سازی
  4. لاگ فعالیت‌های اتصال
  5. نظارت بر عملکرد اتصال
  طراحی مدرن با تب‌بندی
  کاملاً ریسپانسیو
--> 
