<template>
  <div class="sales-stats-page p-6">
    <SalesHeader />
    <SalesSummaryCards :stats="stats" />
    <SalesFilters @filter="onFilter" />
    <div class="grid grid-cols-1 xl:grid-cols-3 gap-6 mt-6">
      <SalesChart :data="chartData" class="xl:col-span-2" />
      <SalesTopProducts :products="topProducts" />
    </div>
    <div class="grid grid-cols-1 xl:grid-cols-2 gap-6 mt-6">
      <SalesRevenueTrends :trends="revenueTrends" />
      <SalesCustomerSegments :segments="customerSegments" />
    </div>
    <SalesTable :sales="filteredSales" class="mt-6" />
    <SalesExport :sales="filteredSales" class="mt-6" />
    <SalesAdvancedAnalytics :analytics="advancedAnalytics" class="mt-6" />
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import SalesHeader from './SalesHeader.vue'
import SalesSummaryCards from './SalesSummaryCards.vue'
import SalesFilters from './SalesFilters.vue'
import SalesChart from './SalesChart.vue'
import SalesTopProducts from './SalesTopProducts.vue'
import SalesRevenueTrends from './SalesRevenueTrends.vue'
import SalesCustomerSegments from './SalesCustomerSegments.vue'
import SalesTable from './SalesTable.vue'
import SalesExport from './SalesExport.vue'
import SalesAdvancedAnalytics from './SalesAdvancedAnalytics.vue'

// Sample/fake data for UI demo
const stats = ref({
  totalRevenue: 2450000,
  todayRevenue: 32000,
  weekRevenue: 210000,
  monthRevenue: 890000,
  totalOrders: 1250,
  avgOrderValue: 1960,
  conversionRate: 3.2,
})
const chartData = ref([])
const topProducts = ref([])
const revenueTrends = ref([])
const customerSegments = ref([])
const sales = ref([])
const filteredSales = ref([])
const advancedAnalytics = ref({})

const onFilter = (filters: any) => {
  // TODO: filter sales based on filters
}
</script>

<style scoped>
.sales-stats-page {
  min-height: 100vh;
  background: #f7fafd;
}
</style> 
