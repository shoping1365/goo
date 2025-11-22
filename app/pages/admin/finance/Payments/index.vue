<template>
  <div class="min-h-screen p-6">
    <div class="w-full space-y-6">
      <!-- هدر صفحه -->
      <PaymentsHeader />

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
          <PaymentsSummaryStats />
          
          <!-- نمودار روند پرداخت‌ها -->
          <PaymentsTrends />
          
          <!-- وضعیت پرداخت‌ها -->
          <PaymentsStatus :status="paymentStatus" />
          
          <!-- محصولات برتر از نظر پرداخت -->
          <PaymentsTopProducts />
          
          <!-- روش‌های پرداخت -->
          <PaymentsMethods />
          
          <!-- آمار جغرافیایی -->
          <PaymentsGeography />
          
          <!-- تحلیل زمانی -->
          <PaymentsTimeAnalysis />
          
          <!-- آمار مشتریان -->
          <PaymentsCustomers />
          
          <!-- تنظیمات درگاه‌های پرداخت -->
          <PaymentsGatewaySettings />
        </div>

        <!-- تب پرداخت‌ها -->
        <div v-if="activeTab === 'payments'" class="space-y-6">
          <PaymentsFilters @filter-change="handleFilterChange" />
          <PaymentsTable :payments="filteredPayments" />
        </div>

        <!-- تب گزارشات -->
        <div v-if="activeTab === 'reports'" class="space-y-6">
          <!-- تحلیل‌ها و نمودارها -->
          <div class="grid grid-cols-1 xl:grid-cols-2 gap-6">
            <PaymentsAdvancedAnalytics :analytics="advancedAnalytics" />
            <PaymentsChart />
          </div>

          <!-- توزیع مبالغ پرداخت -->
          <PaymentsAmounts :amounts="paymentAmounts" />

          <!-- گزارش‌گیری و خروجی -->
          <PaymentsExport :payments="payments" />
        </div>

        <!-- تب لاگ‌ها -->
        <div v-if="activeTab === 'logs'" class="space-y-6">
          <PaymentsLogs />
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
</script>

<script setup lang="ts">
import { computed, ref } from 'vue'

// Import کامپوننت‌های payments
// @ts-ignore
import PaymentsHeader from './PaymentsHeader.vue'
// @ts-ignore
import PaymentsSummaryStats from './PaymentsSummaryStats.vue'
// @ts-ignore
import PaymentsAdvancedAnalytics from './PaymentsAdvancedAnalytics.vue'
// @ts-ignore
import PaymentsFilters from './PaymentsFilters.vue'
// @ts-ignore
import PaymentsTable from './PaymentsTable.vue'
// @ts-ignore
import PaymentsChart from './PaymentsChart.vue'
// @ts-ignore
import PaymentsAmounts from './PaymentsAmounts.vue'
// @ts-ignore
import PaymentsExport from './PaymentsExport.vue'
// @ts-ignore
import PaymentsLogs from './PaymentsLogs.vue'
// @ts-ignore
import PaymentsTrends from './PaymentsTrends.vue'
// @ts-ignore
import PaymentsStatus from './PaymentsStatus.vue'
// @ts-ignore
import PaymentsTopProducts from './PaymentsTopProducts.vue'
// @ts-ignore
import PaymentsGatewaySettings from './PaymentsGatewaySettings.vue'
// @ts-ignore
import PaymentsMethods from './PaymentsMethods.vue'
// @ts-ignore
import PaymentsGeography from './PaymentsGeography.vue'
// @ts-ignore
import PaymentsTimeAnalysis from './PaymentsTimeAnalysis.vue'
// @ts-ignore
import PaymentsCustomers from './PaymentsCustomers.vue'

// تعریف متا برای صفحه
definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// تب‌های صفحه
const tabs = ref([
  { id: 'dashboard', name: 'داشبورد' },
  { id: 'payments', name: 'پرداخت‌ها' },
  { id: 'reports', name: 'گزارشات' },
  { id: 'logs', name: 'لاگ‌ها' }
])

const activeTab = ref('dashboard')

// داده‌های نمونه پرداخت‌ها
const payments = ref([
  {
    id: 1,
    transactionId: 'TXN123456',
    amount: 1250000,
    status: 'success',
    paymentMethod: 'zarinpal',
    customerName: 'علی احمدی',
    orderNumber: 'ORD-001',
    date: '2024-01-15T10:30:00',
    gateway: 'زرین‌پال'
  },
  {
    id: 2,
    transactionId: 'TXN123457',
    amount: 850000,
    status: 'failed',
    paymentMethod: 'nextpay',
    customerName: 'فاطمه محمدی',
    orderNumber: 'ORD-002',
    date: '2024-01-15T09:15:00',
    gateway: 'نکست‌پی'
  },
  {
    id: 3,
    transactionId: 'TXN123458',
    amount: 2100000,
    status: 'pending',
    paymentMethod: 'card_to_card',
    customerName: 'محمد رضایی',
    orderNumber: 'ORD-003',
    date: '2024-01-15T08:45:00',
    gateway: 'کارت به کارت'
  },
  {
    id: 4,
    transactionId: 'TXN123459',
    amount: 950000,
    status: 'success',
    paymentMethod: 'zarinpal',
    customerName: 'زهرا کریمی',
    orderNumber: 'ORD-004',
    date: '2024-01-15T07:20:00',
    gateway: 'زرین‌پال'
  },
  {
    id: 5,
    transactionId: 'TXN123460',
    amount: 1750000,
    status: 'refunded',
    paymentMethod: 'nextpay',
    customerName: 'حسن نوری',
    orderNumber: 'ORD-005',
    date: '2024-01-15T06:10:00',
    gateway: 'نکست‌پی'
  }
])

// داده‌های تحلیل‌های پیشرفته
const advancedAnalytics = ref({
  conversionRate: 87.5,
  avgPaymentAmount: 1360000,
  refundRate: 3.2,
  chargebackRate: 0.8
})

// داده‌های وضعیت پرداخت‌ها
const paymentStatus = ref([
  { status: 'success', count: 156, percentage: 78 },
  { status: 'failed', count: 23, percentage: 12 },
  { status: 'pending', count: 15, percentage: 8 },
  { status: 'cancelled', count: 6, percentage: 2 }
])

// داده‌های توزیع مبالغ
const paymentAmounts = ref([
  {
    range: 'کمتر از ۵۰۰ هزار تومان',
    count: 45,
    percentage: 25
  },
  {
    range: '۵۰۰ هزار تا ۱ میلیون تومان',
    count: 78,
    percentage: 35
  },
  {
    range: '۱ تا ۲ میلیون تومان',
    count: 56,
    percentage: 20
  },
  {
    range: 'بیش از ۲ میلیون تومان',
    count: 23,
    percentage: 20
  }
])

// فیلترهای فعلی
const currentFilters = ref({
  status: 'all',
  paymentMethod: 'all',
  dateFrom: '',
  dateTo: '',
  amountMin: '',
  amountMax: '',
  search: ''
})

// فیلتر کردن پرداخت‌ها
const filteredPayments = computed(() => {
  return payments.value.filter(payment => {
    if (currentFilters.value.status !== 'all' && payment.status !== currentFilters.value.status) return false
    if (currentFilters.value.paymentMethod !== 'all' && payment.paymentMethod !== currentFilters.value.paymentMethod) return false
    if (currentFilters.value.search && !payment.transactionId.includes(currentFilters.value.search) && !payment.orderNumber.includes(currentFilters.value.search)) return false
    // TODO: فیلتر تاریخ و مبلغ
    return true
  })
})

interface Filters {
  [key: string]: unknown
}

// مدیریت تغییر فیلترها
const handleFilterChange = (filters: Filters) => {
  currentFilters.value = {
    status: String(filters.status || 'all'),
    paymentMethod: String(filters.paymentMethod || 'all'),
    dateFrom: String(filters.dateFrom || ''),
    dateTo: String(filters.dateTo || ''),
    amountMin: String(filters.amountMin || ''),
    amountMax: String(filters.amountMax || ''),
    search: String(filters.search || '')
  }
}
</script>

<!--
  صفحه اصلی مدیریت پرداخت‌ها
  تقسیم شده به 5 بخش اصلی:
  
  بخش 1: هدر و آمار کلی
  - هدر صفحه با عنوان و دکمه‌های عملیات
  - آمار کلی پرداخت‌ها (تعداد، مبلغ کل، نرخ موفقیت)
  
  بخش 2: تحلیل‌ها و نمودارها
  - تحلیل‌های پیشرفته (نرخ تبدیل، میانگین مبلغ)
  - نمودار روند پرداخت‌ها
  
  بخش 3: فیلترها و جدول اصلی
  - فیلترهای پیشرفته برای جستجو و فیلتر کردن
  - جدول اصلی پرداخت‌ها با قابلیت‌های عملیاتی
  
  بخش 4: آمار تفصیلی
  - روش‌های پرداخت و آمار هر روش
  - وضعیت پرداخت‌ها (موفق، ناموفق، در انتظار)
  - توزیع مبالغ پرداخت‌ها
  
  بخش 5: تنظیمات و گزارش‌گیری
  - تنظیمات درگاه‌های پرداخت
  - گزارش‌گیری و خروجی
  - لاگ و رویدادهای سیستم
  
  طراحی مدرن و کاملاً ریسپانسیو
  توضیحات کامل به فارسی در کد
-->
<style scoped>
.payments-dashboard {
  background: linear-gradient(135deg, #f8fafc 0%, #e2e8f0 100%);
}
</style> 
