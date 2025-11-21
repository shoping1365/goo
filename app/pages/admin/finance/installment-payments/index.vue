<template>
  <div class="min-h-screen p-6">
    <div class="w-full space-y-6">
      <!-- هدر صفحه -->
      <InstallmentHeader />

      <!-- منوی تب‌بندی -->
      <div class="bg-white border-b border-gray-200">
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

      <!-- محتوای تب‌ها -->
      <div class="p-6">
        <!-- تب داشبورد -->
        <div v-if="activeTab === 'dashboard'" class="space-y-6">
          <!-- آمار کلی -->
          <InstallmentSummaryStats />
          
          <!-- نمودار روند خریدهای اقساطی -->
          <InstallmentTrends />
          
          <!-- وضعیت خریدهای اقساطی -->
          <InstallmentStatus :status="installmentStatus" />
          
          <!-- محصولات برتر از نظر خرید اقساطی -->
          <InstallmentTopProducts />
          
          <!-- روش‌های اعتبارسنجی -->
          <InstallmentValidationMethods />
          
          <!-- آمار جغرافیایی -->
          <InstallmentGeography />
          
          <!-- تحلیل زمانی -->
          <InstallmentTimeAnalysis />
          
          <!-- آمار مشتریان -->
          <InstallmentCustomers />
          
          <!-- تنظیمات سیستم اعتبارسنجی -->
          <InstallmentCreditSettings />
        </div>

        <!-- تب خریدهای اقساطی -->
        <div v-if="activeTab === 'installments'" class="space-y-6">
          <InstallmentFilters @filter-change="handleFilterChange" />
          <InstallmentTable :installments="filteredInstallments" />
        </div>

        <!-- تب اعتبارسنجی -->
        <div v-if="activeTab === 'credit-check'" class="space-y-6">
          <InstallmentCreditCheck />
        </div>

        <!-- تب محصولات -->
        <div v-if="activeTab === 'products'" class="space-y-6">
          <InstallmentProducts />
        </div>

        <!-- تب گزارشات -->
        <div v-if="activeTab === 'reports'" class="space-y-6">
          <!-- تحلیل‌ها و نمودارها -->
          <div class="grid grid-cols-1 xl:grid-cols-2 gap-6">
            <InstallmentAdvancedAnalytics :analytics="advancedAnalytics" />
            <InstallmentChart />
          </div>

          <!-- توزیع مبالغ اقساط -->
          <InstallmentAmounts :amounts="installmentAmounts" />

          <!-- گزارش‌گیری و خروجی -->
          <InstallmentExport :installments="installments" />
        </div>

        <!-- تب لاگ‌ها -->
        <div v-if="activeTab === 'logs'" class="space-y-6">
          <InstallmentLogs />
        </div>

        <!-- تب تنظیمات -->
        <div v-if="activeTab === 'settings'" class="space-y-6">
          <InstallmentSettings />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';

// تعریف definePageMeta برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void

// تعریف layout
definePageMeta({
  layout: 'admin-main'
})

// Import کامپوننت‌های installment
// @ts-ignore
import InstallmentHeader from './InstallmentHeader.vue';
// @ts-ignore
import InstallmentSummaryStats from './InstallmentSummaryStats.vue';
// @ts-ignore
import InstallmentAdvancedAnalytics from './InstallmentAdvancedAnalytics.vue';
// @ts-ignore
import InstallmentFilters from './InstallmentFilters.vue';
// @ts-ignore
import InstallmentTable from './InstallmentTable.vue';
// @ts-ignore
import InstallmentTrends from './InstallmentTrends.vue';
// @ts-ignore
import InstallmentStatus from './InstallmentStatus.vue';
// @ts-ignore
import InstallmentTopProducts from './InstallmentTopProducts.vue';
// @ts-ignore
import InstallmentValidationMethods from './InstallmentValidationMethods.vue';
// @ts-ignore
import InstallmentGeography from './InstallmentGeography.vue';
// @ts-ignore
import InstallmentTimeAnalysis from './InstallmentTimeAnalysis.vue';
// @ts-ignore
import InstallmentCustomers from './InstallmentCustomers.vue';
// @ts-ignore
import InstallmentCreditSettings from './InstallmentCreditSettings.vue';
// @ts-ignore
import InstallmentCreditCheck from './InstallmentCreditCheck.vue';
// @ts-ignore
import InstallmentProducts from './InstallmentProducts.vue';
// @ts-ignore
import InstallmentChart from './InstallmentChart.vue';
// @ts-ignore
import InstallmentAmounts from './InstallmentAmounts.vue';
// @ts-ignore
import InstallmentExport from './InstallmentExport.vue';
// @ts-ignore
import InstallmentLogs from './InstallmentLogs.vue';
// @ts-ignore
import InstallmentSettings from './InstallmentSettings.vue';

// تعریف تب‌ها
const tabs = [
  { id: 'dashboard', name: 'داشبورد' },
  { id: 'installments', name: 'خریدهای اقساطی' },
  { id: 'credit-check', name: 'اعتبارسنجی' },
  { id: 'products', name: 'محصولات' },
  { id: 'reports', name: 'گزارشات' },
  { id: 'logs', name: 'لاگ‌ها' },
  { id: 'settings', name: 'تنظیمات' }
]

const activeTab = ref('dashboard')

// داده‌های نمونه
const installments = ref([])
const installmentStatus = ref({
  pending: 45,
  approved: 120,
  rejected: 15,
  completed: 89,
  overdue: 23
})

const advancedAnalytics = ref({
  totalInstallments: 1250,
  totalAmount: 4500000000,
  averageInstallment: 3600000,
  approvalRate: 78.5,
  defaultRate: 3.2
})

const installmentAmounts = ref({
  ranges: [
    { range: '0-1 میلیون', count: 45, percentage: 15 },
    { range: '1-5 میلیون', count: 120, percentage: 40 },
    { range: '5-10 میلیون', count: 89, percentage: 30 },
    { range: '10+ میلیون', count: 45, percentage: 15 }
  ]
})

const filteredInstallments = computed(() => {
  return installments.value
})

interface Filters {
  [key: string]: unknown
}

// متدها
const handleFilterChange = (_filters: Filters) => {

}

// دریافت داده‌ها
const fetchInstallments = async () => {
  try {
    const response = await $fetch<{ data?: unknown[] }>('/api/admin/installment-payments', {
      method: 'GET'
    })
    const data = response.data
    installments.value = data || []
  } catch (error) {
    console.error('خطا در دریافت خریدهای اقساطی:', error)
  }
}

onMounted(() => {
  fetchInstallments()
})
</script> 
