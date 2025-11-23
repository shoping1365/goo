<template>
  <div v-if="hasAccess" class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
    <div class="flex items-center justify-between mb-6">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">نمودار روند پرداخت‌ها</h3>
        <p class="text-sm text-gray-500 mt-1">Payment Trends Chart</p>
      </div>
      <div class="flex items-center space-x-2 space-x-reverse">
        <button 
          :class="[
            'px-3 py-1 text-sm rounded-lg transition-colors',
            period === 'week' ? 'bg-blue-500 text-white' : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
          ]"
          @click="changePeriod('week')"
        >
          هفته
        </button>
        <button 
          :class="[
            'px-3 py-1 text-sm rounded-lg transition-colors',
            period === 'month' ? 'bg-blue-500 text-white' : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
          ]"
          @click="changePeriod('month')"
        >
          ماه
        </button>
        <button 
          :class="[
            'px-3 py-1 text-sm rounded-lg transition-colors',
            period === 'year' ? 'bg-blue-500 text-white' : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
          ]"
          @click="changePeriod('year')"
        >
          سال
        </button>
      </div>
    </div>

    <!-- نمودار -->
    <div class="h-80 bg-gradient-to-br from-gray-50 to-gray-100 rounded-xl p-6 relative overflow-hidden">
      <div class="absolute inset-0 flex items-center justify-center">
        <div class="text-center">
          <div class="w-24 h-24 bg-gradient-to-br from-blue-100 to-blue-200 rounded-full flex items-center justify-center mb-4">
            <svg class="w-12 h-12 text-blue-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
          </div>
          <h4 class="text-lg font-semibold text-gray-700 mb-2">نمودار روند پرداخت‌ها</h4>
          <p class="text-gray-500 text-sm">Chart.js یا ApexCharts برای نمایش نمودار واقعی</p>
        </div>
      </div>
      
      <!-- نمونه داده‌های نمودار -->
      <div class="absolute bottom-4 left-4 right-4">
        <div class="grid grid-cols-7 gap-2">
          <div v-for="(day, index) in chartData" :key="index" class="text-center">
            <div class="text-xs text-gray-500 mb-1">{{ day.date }}</div>
            <div class="text-sm font-medium text-gray-700">{{ formatPrice(day.amount) }}</div>
            <div class="text-xs text-gray-500">{{ day.count }} تراکنش</div>
          </div>
        </div>
      </div>
    </div>

    <!-- آمار خلاصه -->
    <div class="mt-6 grid grid-cols-1 md:grid-cols-3 gap-6">
      <div class="text-center p-6 bg-blue-50 rounded-lg">
        <div class="text-2xl font-bold text-blue-600">{{ formatPrice(totalAmount) }}</div>
        <div class="text-sm text-blue-500">مجموع مبلغ</div>
      </div>
      <div class="text-center p-6 bg-green-50 rounded-lg">
        <div class="text-2xl font-bold text-green-600">{{ totalCount }}</div>
        <div class="text-sm text-green-500">تعداد تراکنش</div>
      </div>
      <div class="text-center p-6 bg-purple-50 rounded-lg">
        <div class="text-2xl font-bold text-purple-600">{{ formatPrice(averageAmount) }}</div>
        <div class="text-sm text-purple-500">میانگین مبلغ</div>
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

const period = ref('week')

// داده‌های نمودار بر اساس دوره
const chartData = computed(() => {
  if (period.value === 'week') {
    return [
      { date: 'شنبه', amount: 12500000, count: 45 },
      { date: 'یکشنبه', amount: 15800000, count: 52 },
      { date: 'دوشنبه', amount: 14200000, count: 48 },
      { date: 'سه‌شنبه', amount: 18900000, count: 63 },
      { date: 'چهارشنبه', amount: 23400000, count: 78 },
      { date: 'پنج‌شنبه', amount: 26700000, count: 89 },
      { date: 'جمعه', amount: 19800000, count: 66 }
    ]
  } else if (period.value === 'month') {
    return [
      { date: 'هفته ۱', amount: 85000000, count: 312 },
      { date: 'هفته ۲', amount: 92000000, count: 345 },
      { date: 'هفته ۳', amount: 78000000, count: 289 },
      { date: 'هفته ۴', amount: 95000000, count: 356 }
    ]
  } else {
    return [
      { date: 'فروردین', amount: 320000000, count: 1245 },
      { date: 'اردیبهشت', amount: 380000000, count: 1456 },
      { date: 'خرداد', amount: 420000000, count: 1589 },
      { date: 'تیر', amount: 450000000, count: 1678 },
      { date: 'مرداد', amount: 480000000, count: 1789 },
      { date: 'شهریور', amount: 520000000, count: 1923 }
    ]
  }
})

// محاسبه آمار
const totalAmount = computed(() => {
  return chartData.value.reduce((sum, day) => sum + day.amount, 0)
})

const totalCount = computed(() => {
  return chartData.value.reduce((sum, day) => sum + day.count, 0)
})

const averageAmount = computed(() => {
  return totalCount.value > 0 ? totalAmount.value / totalCount.value : 0
})

const changePeriod = (newPeriod: string) => {
  period.value = newPeriod
}

const formatPrice = (price: number) => {
  if (price >= 1000000) {
    return (price / 1000000).toFixed(1) + 'M'
  } else if (price >= 1000) {
    return (price / 1000).toFixed(0) + 'K'
  }
  return new Intl.NumberFormat('fa-IR').format(price)
}
</script>

<!--
  کامپوننت نمودار روند پرداخت‌ها
  - نمایش روند پرداخت‌ها در دوره‌های مختلف
  - آمار خلاصه
  - کاملاً ریسپانسیو
  توضیحات کامل به فارسی در کد
--> 
