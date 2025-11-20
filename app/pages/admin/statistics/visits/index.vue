<template>
  <div class="visits-dashboard min-h-screen" dir="rtl">
    <VisitsHeader />
    
    <!-- Summary Cards -->
    <div class="px-6 mb-6">
      <VisitsSummaryCards :stats="stats" />
    </div>

    <!-- Filters -->
    <div class="px-6 mb-6">
      <VisitsFilters @filter="onFilter" />
    </div>

    <!-- Main Analytics Grid -->
    <div class="px-6 grid grid-cols-1 xl:grid-cols-12 gap-6 mb-6">
      <!-- Chart Section -->
      <div class="xl:col-span-8">
        <VisitsChart :data="chartData" />
      </div>
      <!-- Map Section -->
      <div class="xl:col-span-4">
        <VisitsMap :locations="mapData" />
      </div>
    </div>

    <!-- Secondary Analytics -->
    <div class="px-6 grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-6 mb-6">
      <VisitsTrends :trends="trendsData" />
      <VisitsDevices :devices="devicesData" />
      <VisitsReferrers :referrers="referrersData" />
    </div>

    <!-- Pages and Browser Stats -->
    <div class="px-6 grid grid-cols-1 xl:grid-cols-2 gap-6 mb-6">
      <VisitsTopPages :pages="topPagesData" />
      <VisitsBrowsers :browsers="browsersData" />
    </div>

    <!-- Detailed Table -->
    <div class="px-6 mb-6">
      <VisitsTable :visits="filteredVisits" />
    </div>

    <!-- Export and Advanced Analytics -->
    <div class="px-6 grid grid-cols-1 xl:grid-cols-2 gap-6 mb-6">
      <VisitsExport :visits="filteredVisits" />
      <VisitsAdvancedAnalytics :analytics="advancedAnalytics" />
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
declare const useAuth: () => { user: unknown; hasPermission: (perm: string) => boolean }
</script>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import VisitsAdvancedAnalytics from './VisitsAdvancedAnalytics.vue'
import VisitsBrowsers from './VisitsBrowsers.vue'
import VisitsChart from './VisitsChart.vue'
import VisitsDevices from './VisitsDevices.vue'
import VisitsExport from './VisitsExport.vue'
import VisitsFilters from './VisitsFilters.vue'
import VisitsHeader from './VisitsHeader.vue'
import VisitsMap from './VisitsMap.vue'
import VisitsReferrers from './VisitsReferrers.vue'
import VisitsSummaryCards from './VisitsSummaryCards.vue'
import VisitsTable from './VisitsTable.vue'
import VisitsTopPages from './VisitsTopPages.vue'
import VisitsTrends from './VisitsTrends.vue'

// Set admin layout
definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user: _user, hasPermission: _hasPermission } = useAuth()

// Enhanced data with realistic values
const stats = ref({
  total: 124583,
  today: 1247,
  week: 8932,
  month: 45621,
  unique: 89234,
  bounce: 34.2,
  avgDuration: '4:23',
  growth: 12.5
})

const chartData = ref([
  { date: '1403/08/01', visits: 1200, unique: 850 },
  { date: '1403/08/02', visits: 1350, unique: 920 },
  { date: '1403/08/03', visits: 1180, unique: 800 },
  { date: '1403/08/04', visits: 1420, unique: 980 },
  { date: '1403/08/05', visits: 1680, unique: 1120 },
  { date: '1403/08/06', visits: 1890, unique: 1290 },
  { date: '1403/08/07', visits: 1547, unique: 1050 }
])

const mapData = ref([
  { city: 'تهران', visits: 45621, percentage: 36.7 },
  { city: 'اصفهان', visits: 18934, percentage: 15.2 },
  { city: 'مشهد', visits: 14521, percentage: 11.7 },
  { city: 'شیراز', visits: 12847, percentage: 10.3 },
  { city: 'تبریز', visits: 9856, percentage: 7.9 },
  { city: 'اهواز', visits: 7234, percentage: 5.8 },
  { city: 'کرج', visits: 6892, percentage: 5.5 },
  { city: 'قم', visits: 4567, percentage: 3.7 },
  { city: 'کرمان', visits: 3421, percentage: 2.7 },
  { city: 'ساری', visits: 2690, percentage: 2.2 }
])

const trendsData = ref([
  { hour: '00', visits: 45 },
  { hour: '01', visits: 32 },
  { hour: '02', visits: 28 },
  { hour: '03', visits: 25 },
  { hour: '04', visits: 31 },
  { hour: '05', visits: 42 },
  { hour: '06', visits: 68 },
  { hour: '07', visits: 89 },
  { hour: '08', visits: 124 },
  { hour: '09', visits: 156 },
  { hour: '10', visits: 189 },
  { hour: '11', visits: 167 },
  { hour: '12', visits: 145 },
  { hour: '13', visits: 134 },
  { hour: '14', visits: 178 },
  { hour: '15', visits: 198 },
  { hour: '16', visits: 234 },
  { hour: '17', visits: 267 },
  { hour: '18', visits: 289 },
  { hour: '19', visits: 298 },
  { hour: '20', visits: 276 },
  { hour: '21', visits: 234 },
  { hour: '22', visits: 167 },
  { hour: '23', visits: 98 }
])

const devicesData = ref([
  { device: 'موبایل', visits: 67234, percentage: 54.0 },
  { device: 'دسکتاپ', visits: 43567, percentage: 35.0 },
  { device: 'تبلت', visits: 13782, percentage: 11.0 }
])

const referrersData = ref([
  { source: 'Google', visits: 45621, percentage: 36.6 },
  { source: 'مستقیم', visits: 34892, percentage: 28.0 },
  { source: 'Instagram', visits: 18934, percentage: 15.2 },
  { source: 'Telegram', visits: 12847, percentage: 10.3 },
  { source: 'Facebook', visits: 7856, percentage: 6.3 },
  { source: 'Twitter', visits: 4433, percentage: 3.6 }
])

const topPagesData = ref([
  { page: '/', visits: 23456, percentage: 18.8 },
  { page: '/products', visits: 18934, percentage: 15.2 },
  { page: '/product/123', visits: 12847, percentage: 10.3 },
  { page: '/about', visits: 9856, percentage: 7.9 },
  { page: '/contact', visits: 7234, percentage: 5.8 },
  { page: '/blog', visits: 6892, percentage: 5.5 },
  { page: '/services', visits: 4567, percentage: 3.7 }
])

const browsersData = ref([
  { browser: 'Chrome', visits: 78234, percentage: 62.8 },
  { browser: 'Safari', visits: 23456, percentage: 18.8 },
  { browser: 'Firefox', visits: 12847, percentage: 10.3 },
  { browser: 'Edge', visits: 7856, percentage: 6.3 },
  { browser: 'Opera', visits: 2190, percentage: 1.8 }
])

const _visits = ref([])
const filteredVisits = ref([])
const advancedAnalytics = ref({
  conversionRate: 3.2,
  avgSessionDuration: '4:23',
  pagesPerSession: 2.8,
  newVisitorRate: 67.3
})

const onFilter = (filters: Record<string, unknown>) => {
  // console.log('Filtering with:', filters)
  // TODO: implement actual filtering logic
}

onMounted(() => {
  // Initialize filtered visits with sample data
  filteredVisits.value = Array.from({ length: 20 }, (_, i) => ({
    id: i + 1,
    timestamp: new Date(Date.now() - Math.random() * 86400000).toLocaleString('fa-IR'),
    ip: `192.168.${Math.floor(Math.random() * 255)}.${Math.floor(Math.random() * 255)}`,
    city: mapData.value[Math.floor(Math.random() * mapData.value.length)].city,
    device: devicesData.value[Math.floor(Math.random() * devicesData.value.length)].device,
    browser: browsersData.value[Math.floor(Math.random() * browsersData.value.length)].browser,
    page: topPagesData.value[Math.floor(Math.random() * topPagesData.value.length)].page,
    referrer: referrersData.value[Math.floor(Math.random() * referrersData.value.length)].source,
    duration: `${Math.floor(Math.random() * 10) + 1}:${Math.floor(Math.random() * 60).toString().padStart(2, '0')}`
  }))
})
</script>

<style scoped>
.statistics-visits-page {
  min-height: 100vh;
  background: #f7fafd;
}
</style> 

