<template>
  <div class="grid grid-cols-2 md:grid-cols-3 xl:grid-cols-7 gap-6">
    <SalesSummaryCard title="کل درآمد" :value="formatCurrency(stats.totalRevenue)" icon="mdi-cash" color="green" />
    <SalesSummaryCard title="درآمد امروز" :value="formatCurrency(stats.todayRevenue)" icon="mdi-calendar-today" color="blue" />
    <SalesSummaryCard title="درآمد هفته" :value="formatCurrency(stats.weekRevenue)" icon="mdi-calendar-week" color="purple" />
    <SalesSummaryCard title="درآمد ماه" :value="formatCurrency(stats.monthRevenue)" icon="mdi-calendar-month" color="orange" />
    <SalesSummaryCard title="تعداد سفارش" :value="stats.totalOrders" icon="mdi-cart" color="teal" />
    <SalesSummaryCard title="میانگین سفارش" :value="formatCurrency(stats.avgOrderValue)" icon="mdi-calculator" color="indigo" />
    <SalesSummaryCard title="نرخ تبدیل" :value="stats.conversionRate + '%'" icon="mdi-chart-line" color="red" />
  </div>
</template>
<script setup lang="ts">
import SalesSummaryCard from './SalesSummaryCard.vue';

interface SalesStats {
  totalRevenue: number;
  todayRevenue: number;
  weekRevenue: number;
  monthRevenue: number;
  totalOrders: number;
  avgOrderValue: number;
  conversionRate: number;
}

defineProps<{ stats: SalesStats }>()

const formatCurrency = (value: number) => {
  return new Intl.NumberFormat('fa-IR').format(value) + ' تومان'
}
</script> 
