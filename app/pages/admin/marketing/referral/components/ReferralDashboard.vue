<template>
  <div v-if="hasAccess" class="bg-white rounded-lg shadow p-6">
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-xl font-semibold text-gray-900">آمار کلی برنامه ارجاع</h2>
      <div class="flex space-x-2 space-x-reverse">
        <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors">
          بروزرسانی
        </button>
      </div>
    </div>

    <!-- کارت‌های آمار -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-6">
      <!-- تعداد کل ارجاعات -->
      <div class="bg-gradient-to-r from-blue-500 to-blue-600 text-white p-6 rounded-lg">
        <div class="flex items-center">
          <div class="p-2 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"></path>
            </svg>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium opacity-90">کل ارجاعات</p>
            <p class="text-2xl font-semibold">{{ referralStats.totalReferrals }}</p>
          </div>
        </div>
      </div>

      <!-- ارجاعات موفق -->
      <div class="bg-gradient-to-r from-green-500 to-green-600 text-white p-6 rounded-lg">
        <div class="flex items-center">
          <div class="p-2 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium opacity-90">ارجاعات موفق</p>
            <p class="text-2xl font-semibold">{{ referralStats.successfulReferrals }}</p>
          </div>
        </div>
      </div>

      <!-- نرخ تبدیل -->
      <div class="bg-gradient-to-r from-yellow-500 to-yellow-600 text-white p-6 rounded-lg">
        <div class="flex items-center">
          <div class="p-2 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
            </svg>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium opacity-90">نرخ تبدیل</p>
            <p class="text-2xl font-semibold">{{ referralStats.conversionRate }}%</p>
          </div>
        </div>
      </div>

      <!-- پاداش‌های پرداخت شده -->
      <div class="bg-gradient-to-r from-purple-500 to-purple-600 text-white p-6 rounded-lg">
        <div class="flex items-center">
          <div class="p-2 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
            </svg>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium opacity-90">پاداش‌های پرداخت شده</p>
            <p class="text-2xl font-semibold">{{ referralStats.totalRewards.toLocaleString() }}</p>
          </div>
        </div>
      </div>

      <!-- میانگین پاداش -->
      <div class="bg-gradient-to-r from-red-500 to-red-600 text-white p-6 rounded-lg">
        <div class="flex items-center">
          <div class="p-2 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z"></path>
            </svg>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium opacity-90">میانگین پاداش</p>
            <p class="text-2xl font-semibold">{{ referralStats.averageReward.toLocaleString() }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- نمودار روند -->
    <div class="mt-8">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">روند رشد برنامه ارجاع</h3>
      <div class="bg-gray-50 p-6 rounded-lg">
        <div class="h-64 flex items-center justify-center text-gray-500">
          نمودار روند رشد (در حال توسعه)
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue';
import { useAuth } from '~/composables/useAuth';

declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>;

// احراز هویت
const { user, isAuthenticated } = useAuth();

// بررسی دسترسی admin
const hasAccess = computed(() => {
  if (!isAuthenticated.value) {
    return false;
  }

  const userRole = user.value?.role?.toLowerCase() || '';
  const adminRoles = ['admin', 'developer', 'super_admin', 'manager', 'operator'];
  return adminRoles.includes(userRole);
});

// بررسی احراز هویت و دسترسی admin - نمایش 404 در صورت عدم دسترسی
const checkAuth = async (): Promise<void> => {
  if (!hasAccess.value) {
    await navigateTo('/404', { external: false });
  }
};


// آمار برنامه ارجاع
const referralStats = ref({
  totalReferrals: 0,
  successfulReferrals: 0,
  conversionRate: 0,
  totalRewards: 0,
  averageReward: 0
})

// بررسی احراز هویت هنگام تغییر وضعیت احراز هویت
watch([isAuthenticated, hasAccess], async () => {
  if (!hasAccess.value) {
    await checkAuth();
  }
});

// بارگذاری آمار
onMounted(async () => {
  await checkAuth();
  // TODO: فراخوانی API برای دریافت آمار
  referralStats.value = {
    totalReferrals: 850,
    successfulReferrals: 425,
    conversionRate: 50,
    totalRewards: 8500000,
    averageReward: 20000
  };
});
</script> 
