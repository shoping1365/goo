<template>
  <div v-if="hasAccess" class="mb-8">
    <h2 class="text-xl font-bold text-gray-900 mb-6">آمار کلی گیفت کارت‌ها</h2>
    
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <!-- کل کارت‌ها -->
      <div class="bg-gradient-to-r from-blue-500 to-blue-600 rounded-2xl p-6 text-white shadow-lg">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-blue-100 text-sm font-medium">کل کارت‌ها</p>
            <p class="text-3xl font-bold mt-2">{{ formatNumber(stats.totalCards) }}</p>
            <p class="text-blue-100 text-sm mt-1">
              <span class="text-green-300">+{{ stats.monthlyGrowth }}%</span> رشد ماهانه
            </p>
          </div>
          <div class="w-12 h-12 bg-white/20 rounded-xl flex items-center justify-center">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"></path>
            </svg>
          </div>
        </div>
      </div>

      <!-- کارت‌های فعال -->
      <div class="bg-gradient-to-r from-green-500 to-green-600 rounded-2xl p-6 text-white shadow-lg">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-green-100 text-sm font-medium">کارت‌های فعال</p>
            <p class="text-3xl font-bold mt-2">{{ formatNumber(stats.activeCards) }}</p>
            <p class="text-green-100 text-sm mt-1">
              <span class="text-green-300">{{ Math.round((stats.activeCards / stats.totalCards) * 100) }}%</span> از کل
            </p>
          </div>
          <div class="w-12 h-12 bg-white/20 rounded-xl flex items-center justify-center">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
      </div>

      <!-- کارت‌های استفاده شده -->
      <div class="bg-gradient-to-r from-purple-500 to-purple-600 rounded-2xl p-6 text-white shadow-lg">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-purple-100 text-sm font-medium">کارت‌های استفاده شده</p>
            <p class="text-3xl font-bold mt-2">{{ formatNumber(stats.usedCards) }}</p>
            <p class="text-purple-100 text-sm mt-1">
              <span class="text-purple-300">{{ stats.conversionRate }}%</span> نرخ تبدیل
            </p>
          </div>
          <div class="w-12 h-12 bg-white/20 rounded-xl flex items-center justify-center">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
            </svg>
          </div>
        </div>
      </div>

      <!-- کارت‌های منقضی شده -->
      <div class="bg-gradient-to-r from-orange-500 to-orange-600 rounded-2xl p-6 text-white shadow-lg">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-orange-100 text-sm font-medium">کارت‌های منقضی شده</p>
            <p class="text-3xl font-bold mt-2">{{ formatNumber(stats.expiredCards) }}</p>
            <p class="text-orange-100 text-sm mt-1">
              <span class="text-orange-300">{{ Math.round((stats.expiredCards / stats.totalCards) * 100) }}%</span> از کل
            </p>
          </div>
          <div class="w-12 h-12 bg-white/20 rounded-xl flex items-center justify-center">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- آمار مالی -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mt-6">
      <!-- کل ارزش -->
      <div class="bg-white rounded-2xl p-6 border border-gray-200/50 shadow-sm">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">کل ارزش گیفت کارت‌ها</h3>
          <div class="w-10 h-10 bg-green-100 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
            </svg>
          </div>
        </div>
        <div class="space-y-3">
          <div class="flex justify-between items-center">
            <span class="text-gray-600">کل ارزش:</span>
            <span class="text-2xl font-bold text-green-600">{{ formatPrice(stats.totalValue) }}</span>
          </div>
          <div class="flex justify-between items-center">
            <span class="text-gray-600">استفاده شده:</span>
            <span class="text-lg font-semibold text-gray-900">{{ formatPrice(stats.usedValue) }}</span>
          </div>
          <div class="flex justify-between items-center">
            <span class="text-gray-600">باقی‌مانده:</span>
            <span class="text-lg font-semibold text-blue-600">{{ formatPrice(stats.totalValue - stats.usedValue) }}</span>
          </div>
        </div>
      </div>

      <!-- نمودار وضعیت -->
      <div class="bg-white rounded-2xl p-6 border border-gray-200/50 shadow-sm">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">وضعیت کارت‌ها</h3>
          <div class="w-10 h-10 bg-purple-100 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
          </div>
        </div>
        <div class="space-y-4">
          <!-- نوار پیشرفت فعال -->
          <div>
            <div class="flex justify-between items-center mb-2">
              <span class="text-sm text-gray-600">فعال</span>
              <span class="text-sm font-medium text-gray-900">{{ Math.round((stats.activeCards / stats.totalCards) * 100) }}%</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div class="bg-green-500 h-2 rounded-full" :style="{ width: (stats.activeCards / stats.totalCards) * 100 + '%' }"></div>
            </div>
          </div>
          
          <!-- نوار پیشرفت استفاده شده -->
          <div>
            <div class="flex justify-between items-center mb-2">
              <span class="text-sm text-gray-600">استفاده شده</span>
              <span class="text-sm font-medium text-gray-900">{{ Math.round((stats.usedCards / stats.totalCards) * 100) }}%</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div class="bg-purple-500 h-2 rounded-full" :style="{ width: (stats.usedCards / stats.totalCards) * 100 + '%' }"></div>
            </div>
          </div>
          
          <!-- نوار پیشرفت منقضی شده -->
          <div>
            <div class="flex justify-between items-center mb-2">
              <span class="text-sm text-gray-600">منقضی شده</span>
              <span class="text-sm font-medium text-gray-900">{{ Math.round((stats.expiredCards / stats.totalCards) * 100) }}%</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div class="bg-orange-500 h-2 rounded-full" :style="{ width: (stats.expiredCards / stats.totalCards) * 100 + '%' }"></div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>
</script>

<script setup lang="ts">
import { computed, onMounted, watch } from 'vue';
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
// تعریف props برای آمار گیفت کارت
interface GiftCardStats {
  totalCards: number
  activeCards: number
  usedCards: number
  expiredCards: number
  totalValue: number
  usedValue: number
  monthlyGrowth: number
  conversionRate: number
}

const _props = defineProps<{
  stats: GiftCardStats
}>()

// تابع فرمت کردن اعداد
const formatNumber = (num: number): string => {
  return new Intl.NumberFormat('fa-IR').format(num)
}

// تابع فرمت کردن قیمت
const formatPrice = (price: number): string => {
  return new Intl.NumberFormat('fa-IR', {
    style: 'currency',
    currency: 'IRR',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(price)
}
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 
