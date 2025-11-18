<template>
  <div class="min-h-screen p-6">
    <div class="max-w-7xl mx-auto space-y-6">
      <!-- هدر صفحه مالیات -->
      <TaxHeader />

      <!-- آمار کلی و داشبورد -->
      <TaxDashboard />

      <!-- نمودار روند مالیات -->
      <TaxTrendChart />

      <!-- مقایسه دوره‌ای -->
      <TaxComparison />

      <!-- مدیریت نرخ‌های مالیاتی -->
      <TaxRatesManagement />

      <!-- تنظیمات معافیت‌ها -->
      <TaxExemptions />

      <!-- نرخ‌های منطقه‌ای -->
      <TaxRegionalRates />

      <!-- محاسبه خودکار مالیات -->
      <TaxCalculator />

      <!-- تنظیمات محاسباتی -->
      <TaxCalculationSettings />

      <!-- تست محاسبات -->
      <TaxCalculationTest />

      <!-- گزارش‌های مالیاتی -->
      <TaxReports />

      <!-- مدیریت فاکتورها -->
      <TaxInvoices />

      <!-- گزارش‌گیری پیشرفته -->
      <TaxAdvancedReports />
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
</script>

<script setup lang="ts">
import { onMounted, ref } from 'vue'

// Import کامپوننت‌های مالیات
// @ts-ignore
import TaxHeader from './TaxHeader.vue'
// @ts-ignore
import TaxDashboard from './TaxDashboard.vue'
// @ts-ignore
import TaxTrendChart from './TaxTrendChart.vue'
// @ts-ignore
import TaxComparison from './TaxComparison.vue'
// @ts-ignore
import TaxRatesManagement from './TaxRatesManagement.vue'
// @ts-ignore
import TaxExemptions from './TaxExemptions.vue'
// @ts-ignore
import TaxRegionalRates from './TaxRegionalRates.vue'
// @ts-ignore
import TaxCalculator from './TaxCalculator.vue'
// @ts-ignore
import TaxCalculationSettings from './TaxCalculationSettings.vue'
// @ts-ignore
import TaxCalculationTest from './TaxCalculationTest.vue'
// @ts-ignore
import TaxReports from './TaxReports.vue'
// @ts-ignore
import TaxInvoices from './TaxInvoices.vue'
// @ts-ignore
import TaxAdvancedReports from './TaxAdvancedReports.vue'

// تعریف متا برای صفحه
definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// داده‌های نمونه مالیات
const taxData = ref({
  totalTaxCollected: 0,
  monthlyTax: 0,
  quarterlyTax: 0,
  yearlyTax: 0,
  growthRate: 0,
  taxTrend: []
})

// بارگذاری داده‌های مالیات
const loadTaxData = async () => {
  try {
    // TODO: دریافت داده‌ها از API
    // const response = await $fetch('/api/admin/tax/dashboard')
    // taxData.value = response.data
    
    // داده‌های نمونه
    taxData.value = {
      totalTaxCollected: 1250000000, // 1.25 میلیارد تومان
      monthlyTax: 45000000, // 45 میلیون تومان
      quarterlyTax: 135000000, // 135 میلیون تومان
      yearlyTax: 540000000, // 540 میلیون تومان
      growthRate: 12.5, // 12.5 درصد رشد
      taxTrend: [
        { month: 'فروردین', amount: 42000000 },
        { month: 'اردیبهشت', amount: 45000000 },
        { month: 'خرداد', amount: 48000000 },
        { month: 'تیر', amount: 52000000 },
        { month: 'مرداد', amount: 49000000 },
        { month: 'شهریور', amount: 51000000 }
      ]
    }
  } catch (error) {
    console.error('خطا در بارگذاری داده‌های مالیات:', error)
  }
}

// بارگذاری اولیه
onMounted(() => {
  loadTaxData()
})
</script>

<!--
  صفحه اصلی مدیریت مالیات
  بخش اول: داشبورد و آمار کلی
  شامل:
  1. هدر صفحه مالیات
  2. آمار کلی و داشبورد
  3. نمودار روند مالیات
  4. مقایسه دوره‌ای
  طراحی مدرن و کاملاً ریسپانسیو
  توضیحات کامل به فارسی در کد
-->
<style scoped>
.tax-dashboard {
  background: linear-gradient(135deg, #f8fafc 0%, #e2e8f0 100%);
}
</style>
