<template>
  <div v-if="hasAccess" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
    <!-- آمار کلی -->
    <TemplateCard
      title="کل درگاه‌ها"
      :value="stats.totalGateways"
      variant="blue"
    >
      <template #icon>
        <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z" />
        </svg>
      </template>
    </TemplateCard>

    <!-- درگاه‌های فعال -->
    <TemplateCard
      title="درگاه‌های فعال"
      :value="stats.activeGateways"
      variant="green"
    >
      <template #icon>
        <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>
      </template>
    </TemplateCard>

    <!-- تراکنش‌های امروز -->
    <TemplateCard
      title="تراکنش‌های امروز"
      :value="formatNumber(stats.todayTransactions)"
      variant="purple"
    >
      <template #icon>
        <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" />
        </svg>
      </template>
    </TemplateCard>

    <!-- درآمد امروز -->
    <TemplateCard
      title="درآمد امروز"
      :value="formatPrice(stats.todayRevenue)"
      variant="yellow"
    >
      <template #icon>
        <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1" />
        </svg>
      </template>
    </TemplateCard>
  </div>

  <!-- نمودار آمار و وضعیت درگاه‌ها -->
  <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-8">
    <!-- نمودار آمار -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
    <div class="flex items-center justify-between mb-6">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">آمار هفتگی تراکنش‌ها</h3>
        <p class="text-sm text-gray-600 mt-1">نمایش تراکنش‌های روزانه در بازه زمانی انتخاب شده</p>
      </div>
      <div class="flex space-x-2 space-x-reverse">
        <TemplateButton 
          :bg-gradient="chartPeriod === 'week' ? 'bg-gradient-to-r from-blue-500 to-blue-600' : 'bg-gradient-to-r from-gray-100 to-gray-200'" 
          :text-color="chartPeriod === 'week' ? 'text-white' : 'text-gray-600'"
          :border-color="chartPeriod === 'week' ? 'border border-blue-500' : 'border border-gray-300'"
          :hover-class="chartPeriod === 'week' ? 'hover:from-blue-600 hover:to-blue-700' : 'hover:from-gray-200 hover:to-gray-300'"
          :focus-class="chartPeriod === 'week' ? 'focus:ring-2 focus:ring-blue-500 focus:ring-offset-2' : 'focus:ring-2 focus:ring-gray-500 focus:ring-offset-2'"
          size="medium"
          :custom-class="chartPeriod === 'week' ? 'shadow-lg shadow-blue-500/25' : ''"
          @click="chartPeriod = 'week'"
        >
          هفته
        </TemplateButton>
        <TemplateButton 
          :bg-gradient="chartPeriod === 'month' ? 'bg-gradient-to-r from-blue-500 to-blue-600' : 'bg-gradient-to-r from-gray-100 to-gray-200'" 
          :text-color="chartPeriod === 'month' ? 'text-white' : 'text-gray-600'"
          :border-color="chartPeriod === 'month' ? 'border border-blue-500' : 'border border-gray-300'"
          :hover-class="chartPeriod === 'month' ? 'hover:from-blue-600 hover:to-blue-700' : 'hover:from-gray-200 hover:to-gray-300'"
          :focus-class="chartPeriod === 'month' ? 'focus:ring-2 focus:ring-blue-500 focus:ring-offset-2' : 'focus:ring-2 focus:ring-gray-500 focus:ring-offset-2'"
          size="medium"
          :custom-class="chartPeriod === 'month' ? 'shadow-lg shadow-blue-500/25' : ''"
          @click="chartPeriod = 'month'"
        >
          ماه
        </TemplateButton>
      </div>
    </div>
    
    <!-- نمودار پیشرفته -->
    <div class="relative bg-gradient-to-br from-gray-50 to-white rounded-xl p-6 border border-gray-100">
      <!-- خطوط راهنما -->
      <div class="absolute inset-6 flex flex-col justify-between pointer-events-none">
        <div class="border-b border-gray-200/50"></div>
        <div class="border-b border-gray-200/50"></div>
        <div class="border-b border-gray-200/50"></div>
        <div class="border-b border-gray-200/50"></div>
        <div class="border-b border-gray-200/50"></div>
      </div>
      
      <!-- نمودار ستونی -->
      <div class="h-80 flex items-end justify-center relative z-10">
        <div v-for="(item, index) in chartData" :key="index" class="flex flex-col items-center group mx-4">
          <!-- ستون اصلی -->
          <div class="w-full relative">
            <div 
              class="w-full rounded-t-xl transition-all duration-700 ease-out group-hover:shadow-2xl group-hover:scale-110 cursor-pointer relative overflow-hidden"
              :style="{ height: (item.value / maxChartValue) * 280 + 'px' }"
            >
              <!-- گرادیان پس‌زمینه - رنگ‌های پاستیلی -->
              <div class="absolute inset-0 bg-gradient-to-t from-indigo-400/80 via-purple-300/70 to-pink-300/60"></div>
              
              <!-- افکت درخشش -->
              <div class="absolute inset-0 bg-gradient-to-r from-transparent via-white/30 to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-500"></div>
              
              <!-- خط بالای ستون - رنگ‌های پاستیلی -->
              <div class="absolute top-0 left-0 right-0 h-2 bg-gradient-to-r from-pink-200/80 via-purple-200/70 to-indigo-200/60 rounded-t-xl"></div>
              
              <!-- سایه داخلی -->
              <div class="absolute inset-0 bg-gradient-to-b from-transparent via-black/10 to-transparent"></div>
            </div>
            
            <!-- نشانگر مقدار -->
            <div class="absolute -top-10 left-1/2 transform -translate-x-1/2 bg-gradient-to-r from-gray-900 to-gray-800 text-white text-xs px-3 py-2 rounded-lg opacity-0 group-hover:opacity-100 transition-all duration-300 pointer-events-none whitespace-nowrap shadow-lg">
              <div class="font-bold">{{ formatNumber(item.value) }}</div>
              <div class="text-gray-300 text-xs">تراکنش</div>
              <!-- فلش -->
              <div class="absolute top-full left-1/2 transform -translate-x-1/2 w-0 h-0 border-l-4 border-r-4 border-t-4 border-transparent border-t-gray-800"></div>
            </div>
          </div>
          
          <!-- برچسب روز -->
          <div class="mt-6 text-center">
            <div class="text-sm font-semibold text-gray-800 group-hover:text-indigo-500 transition-colors duration-300">
              {{ item.label }}
            </div>
            <div class="text-xs text-gray-500 mt-1 font-medium">
              {{ formatNumber(item.value) }}
            </div>
          </div>
        </div>
      </div>
      
      <!-- محور Y -->
      <div class="absolute left-6 top-6 bottom-6 flex flex-col justify-between text-xs text-gray-600 pointer-events-none font-medium">
        <span class="bg-white/80 px-2 py-1 rounded">{{ formatNumber(maxChartValue) }}</span>
        <span class="bg-white/80 px-2 py-1 rounded">{{ formatNumber(maxChartValue * 0.75) }}</span>
        <span class="bg-white/80 px-2 py-1 rounded">{{ formatNumber(maxChartValue * 0.5) }}</span>
        <span class="bg-white/80 px-2 py-1 rounded">{{ formatNumber(maxChartValue * 0.25) }}</span>
        <span class="bg-white/80 px-2 py-1 rounded">0</span>
      </div>
    </div>
      </div>

    <!-- وضعیت درگاه‌ها -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">وضعیت درگاه‌ها</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="space-y-4">
          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-3 h-3 bg-green-400 rounded-full ml-2"></div>
              <span class="text-gray-700">فعال</span>
            </div>
            <span class="font-medium text-gray-900">{{ stats.activeGateways }} درگاه</span>
          </div>
          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-3 h-3 bg-red-400 rounded-full ml-2"></div>
              <span class="text-gray-700">غیرفعال</span>
            </div>
            <span class="font-medium text-gray-900">{{ stats.inactiveGateways }} درگاه</span>
          </div>
          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-3 h-3 bg-yellow-400 rounded-full ml-2"></div>
              <span class="text-gray-700">نگهداری</span>
            </div>
            <span class="font-medium text-gray-900">{{ stats.maintenanceGateways }} درگاه</span>
          </div>
        </div>
        
        <!-- نمودار دایره‌ای -->
        <div class="flex justify-center">
          <div class="relative w-32 h-32">
            <svg class="w-32 h-32 transform -rotate-90">
              <circle cx="64" cy="64" r="56" stroke="currentColor" stroke-width="8" fill="transparent" class="text-gray-200"></circle>
              <circle
cx="64" cy="64" r="56" stroke="currentColor" stroke-width="8" fill="transparent" 
                :stroke-dasharray="`${(stats.activeGateways / stats.totalGateways) * 352} 352`" 
                class="text-green-400"></circle>
            </svg>
            <div class="absolute inset-0 flex items-center justify-center">
              <span class="text-lg font-bold text-gray-900">{{ stats.activePercentage }}%</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- درگاه‌های برتر امروز -->
  <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-8">
    <h3 class="text-lg font-semibold text-gray-900 mb-4">درگاه‌های برتر امروز</h3>
    <div class="space-y-4">
      <div v-for="(gateway, index) in topGateways" :key="gateway.id" class="flex items-center justify-between">
        <div class="flex items-center">
          <div class="w-8 h-8 rounded-full flex items-center justify-center text-white text-sm font-bold ml-3" :class="gateway.color">
            {{ index + 1 }}
          </div>
          <div>
            <p class="font-medium text-gray-900">{{ gateway.name }}</p>
            <p class="text-sm text-gray-600">{{ gateway.englishName }}</p>
          </div>
        </div>
        <div class="text-right">
          <p class="font-medium text-gray-900">{{ formatNumber(gateway.transactions) }}</p>
          <p class="text-sm text-gray-600">تراکنش</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>
</script>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue';
import { useAuth } from '~/composables/useAuth';
import TemplateCard from '~/components/common/TemplateCard.vue'
import TemplateButton from '~/components/common/TemplateButton.vue'

// احراز هویت
const { user, isAuthenticated } = useAuth();

// بررسی دسترسی admin
const hasAccess = computed(() => {
  if (!isAuthenticated.value) {
    return false;
  }

  const userRole = user.value?.role?.toLowerCase() || '';
  const adminRoles = ['admin', 'developer'];
  return adminRoles.includes(userRole);
});

// بررسی احراز هویت و دسترسی admin - نمایش 404 در صورت عدم دسترسی
const checkAuth = async (): Promise<void> => {
  if (!hasAccess.value) {
    await navigateTo('/404', { external: false });
  }
};

// بررسی احراز هویت در هنگام mount
onMounted(async () => {
  await checkAuth();
});

// بررسی احراز هویت هنگام تغییر وضعیت احراز هویت
watch([isAuthenticated, hasAccess], async () => {
  if (!hasAccess.value) {
    await checkAuth();
  }
});

// تعریف interface برای آمار
interface GatewayStats {
  totalGateways: number
  activeGateways: number
  inactiveGateways: number
  maintenanceGateways: number
  todayTransactions: number
  todayRevenue: number
  gatewaysGrowth: number
  transactionsGrowth: number
  revenueGrowth: number
  activePercentage: number
}

interface TopGateway {
  id: number
  name: string
  englishName: string
  transactions: number
  color: string
}

interface ChartData {
  label: string
  value: number
}

// متغیرهای reactive
const chartPeriod = ref<'week' | 'month'>('week')

// داده‌های آمار
const stats = ref<GatewayStats>({
  totalGateways: 15,
  activeGateways: 12,
  inactiveGateways: 2,
  maintenanceGateways: 1,
  todayTransactions: 4567,
  todayRevenue: 1250000000,
  gatewaysGrowth: 5.2,
  transactionsGrowth: 12.8,
  revenueGrowth: 8.5,
  activePercentage: 80
})

// درگاه‌های برتر
const topGateways = ref<TopGateway[]>([])

// دریافت درگاه‌های برتر از API
const fetchTopGateways = async () => {
  try {
    // دریافت درگاه‌های فعال
    const response = await $fetch('/api/payment-gateways?status=active') as { success?: boolean; data?: unknown[] }
    
    interface Gateway {
      id?: number | string
      name?: string
      english_name?: string
      status?: string
      today_transactions?: number
      [key: string]: unknown
    }
    
    if (response.success && Array.isArray(response.data)) {
      // فیلتر کردن فقط درگاه‌های فعال و مرتب‌سازی بر اساس تراکنش‌های امروز
      const activeGateways = (response.data as Gateway[])
        .filter((gateway: Gateway) => gateway.status === 'active')
        .sort((a: Gateway, b: Gateway) => (b.today_transactions || 0) - (a.today_transactions || 0))
        .slice(0, 5)
        .map((gateway: Gateway, index: number) => ({
          id: Number(gateway.id || 0),
          name: String(gateway.name || ''),
          englishName: String(gateway.english_name || gateway.name || ''),
          transactions: Number(gateway.today_transactions || 0),
          color: ['bg-green-500', 'bg-blue-500', 'bg-orange-500', 'bg-teal-500', 'bg-indigo-500'][index] || 'bg-gray-500'
        }))
      
      topGateways.value = activeGateways as typeof topGateways.value
    }
  } catch (error) {
    console.error('خطا در دریافت درگاه‌های برتر:', error)
    // در صورت خطا، از داده‌های نمونه استفاده کن
    topGateways.value = [
      { id: 1, name: 'زرین‌پال', englishName: 'ZarinPal', transactions: 1250, color: 'bg-green-500' },
      { id: 2, name: 'نکست‌پی', englishName: 'NextPay', transactions: 890, color: 'bg-blue-500' },
      { id: 3, name: 'پی‌ایر', englishName: 'PayIR', transactions: 720, color: 'bg-orange-500' },
      { id: 4, name: 'پی‌پینگ', englishName: 'PayPing', transactions: 650, color: 'bg-teal-500' },
      { id: 5, name: 'ملی‌پی', englishName: 'MelliPay', transactions: 580, color: 'bg-indigo-500' }
    ]
  }
}

// داده‌های نمودار هفتگی
const weeklyChartData = ref<ChartData[]>([
  { label: 'شنبه', value: 1200 },
  { label: 'یکشنبه', value: 1350 },
  { label: 'دوشنبه', value: 1100 },
  { label: 'سه‌شنبه', value: 1450 },
  { label: 'چهارشنبه', value: 1300 },
  { label: 'پنج‌شنبه', value: 1600 },
  { label: 'جمعه', value: 1400 }
])

// داده‌های نمودار ماهانه
const monthlyChartData = ref<ChartData[]>([
  { label: 'هفته 1', value: 8500 },
  { label: 'هفته 2', value: 9200 },
  { label: 'هفته 3', value: 8800 },
  { label: 'هفته 4', value: 9500 }
])

// محاسبه داده‌های نمودار
const chartData = computed(() => {
  return chartPeriod.value === 'week' ? weeklyChartData.value : monthlyChartData.value
})

const maxChartValue = computed(() => {
  return Math.max(...chartData.value.map(item => item.value))
})



// توابع کمکی
const formatNumber = (num: number): string => {
  if (num >= 1000000) {
    return (num / 1000000).toFixed(1) + 'M'
  } else if (num >= 1000) {
    return (num / 1000).toFixed(0) + 'K'
  }
  return new Intl.NumberFormat('fa-IR').format(num)
}

const formatPrice = (price: number): string => {
  if (price >= 1000000000) {
    return (price / 1000000000).toFixed(1) + ' میلیارد'
  } else if (price >= 1000000) {
    return (price / 1000000).toFixed(1) + ' میلیون'
  } else if (price >= 1000) {
    return (price / 1000).toFixed(0) + ' هزار'
  }
  return new Intl.NumberFormat('fa-IR').format(price) + ' تومان'
}

// بارگذاری داده‌ها
onMounted(() => {
  fetchTopGateways()
})
</script>

<!--
  کامپوننت آمار کلی درگاه‌های پرداخت
  شامل:
  1. کارت‌های آمار کلی (کل درگاه‌ها، فعال، تراکنش‌ها، درآمد)
  2. نمودار تراکنش‌های هفتگی/ماهانه
  3. لیست درگاه‌های برتر
  4. نمودار وضعیت درگاه‌ها
  5. نمایش رشد و تغییرات
  6. طراحی ریسپانسیو
  توضیحات کامل به فارسی در کد
--> 
