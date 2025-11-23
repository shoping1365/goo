<template>
  <ClientOnly>
    <!-- Page Header -->
    <div class="bg-white shadow-sm border-b border-gray-200 mb-4">
      <div class="w-full px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center py-3">
          <div>
            <h1 class="text-lg font-bold text-gray-900">سفارشات در حال انجام</h1>
            <p class="text-xs text-gray-600 mt-1">مدیریت و پیگیری سفارشات در حال انجام</p>
          </div>
          <div class="flex space-x-2 space-x-reverse">
            <button
              class="inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-lg text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 shadow-md transition-all duration-200"
              @click="refreshData"
            >
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
              </svg>
              بروزرسانی
            </button>
          </div>
        </div>
      </div>
    </div>

    <div class="w-full px-4 py-4">
      <!-- تب‌های اصلی -->
      <div class="bg-white border-b border-gray-200 mb-6">
        <nav class="-mb-px flex px-6 overflow-x-auto">
          <button
            v-for="tab in tabs"
            :key="tab.id"
            :class="[
              activeTab === tab.id
                ? 'border-blue-500 text-blue-600'
                : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300',
            'flex-1 whitespace-nowrap py-4 px-4 border-b-2 font-medium text-sm text-center'
            ]"
            @click="activeTab = tab.id"
          >
            {{ tab.name }}
          </button>
        </nav>
      </div>

      <div class="px-4 py-4">
        <!-- تب داشبورد -->
        <div v-if="activeTab === 'dashboard'" data-component="dashboard">
          <ProcessingDashboard />
        </div>

        <!-- تب سفارشات -->
        <div v-if="activeTab === 'processing-orders'" data-component="orders">
          <ProcessingOrdersTable />
        </div>

        <!-- تب گزارشات -->
        <div v-if="activeTab === 'reports'">
          <ProcessingReports />
        </div>
      </div>
    </div>
  </ClientOnly>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
declare const useHead: (head: { title?: string; meta?: Array<{ name?: string; content?: string }> }) => void
</script>

<script setup lang="ts">
import { ref } from 'vue'
import ProcessingDashboard from './ProcessingDashboard.vue'
import ProcessingOrdersTable from './ProcessingOrdersTable.vue'
import ProcessingReports from './ProcessingReports.vue'

definePageMeta({ layout: 'admin-main', middleware: 'admin' })

// تب‌های صفحه
const tabs = ref([
  { id: 'dashboard', name: 'داشبورد' },
  { id: 'processing-orders', name: 'سفارشات در حال انجام' },
  { id: 'reports', name: 'گزارشات' }
])

const activeTab = ref('dashboard')

// تابع بروزرسانی داده‌ها
/**
 * این تابع داده‌های صفحه را بروزرسانی می‌کند.
 * در Vue 3، باید از emit یا props برای ارتباط با کامپوننت‌های فرزند استفاده کرد.
 */
const refreshData = () => {
  // TODO: استفاده از emit یا ref برای بروزرسانی کامپوننت‌های فرزند
  // این کد قبلی از __vue__ استفاده می‌کرد که در Vue 3 وجود ندارد
}

// تنظیم عنوان صفحه
useHead({
  title: 'سفارشات در حال انجام - پنل مدیریت',
  meta: [
    { name: 'description', content: 'مدیریت سفارشات در حال انجام و گزارشات مربوطه' }
  ]
})
</script> 
