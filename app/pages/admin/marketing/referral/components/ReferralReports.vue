<template>
  <div v-if="hasAccess" class="bg-white rounded-lg shadow p-6">
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-xl font-semibold text-gray-900">گزارش‌ها و تحلیل</h2>
      <div class="flex space-x-2 space-x-reverse">
        <select v-model="reportType" class="border border-gray-300 rounded-md px-3 py-2">
          <option value="daily">روزانه</option>
          <option value="weekly">هفتگی</option>
          <option value="monthly">ماهانه</option>
        </select>
        <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors">
          دانلود گزارش
        </button>
      </div>
    </div>

    <!-- کارت‌های آمار گزارش -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
      <div class="bg-blue-50 p-6 rounded-lg">
        <div class="flex items-center">
          <div class="p-2 bg-blue-100 rounded-lg">
            <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
            </svg>
          </div>
          <div class="mr-3">
            <p class="text-sm text-blue-600">نرخ تبدیل</p>
            <p class="text-xl font-semibold text-blue-900">{{ reportStats.conversionRate }}%</p>
          </div>
        </div>
      </div>

      <div class="bg-green-50 p-6 rounded-lg">
        <div class="flex items-center">
          <div class="p-2 bg-green-100 rounded-lg">
            <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
            </svg>
          </div>
          <div class="mr-3">
            <p class="text-sm text-green-600">میانگین پاداش</p>
            <p class="text-xl font-semibold text-green-900">{{ reportStats.avgReward.toLocaleString() }}</p>
          </div>
        </div>
      </div>

      <div class="bg-purple-50 p-6 rounded-lg">
        <div class="flex items-center">
          <div class="p-2 bg-purple-100 rounded-lg">
            <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"></path>
            </svg>
          </div>
          <div class="mr-3">
            <p class="text-sm text-purple-600">کاربران فعال</p>
            <p class="text-xl font-semibold text-purple-900">{{ reportStats.activeUsers }}</p>
          </div>
        </div>
      </div>

      <div class="bg-orange-50 p-6 rounded-lg">
        <div class="flex items-center">
          <div class="p-2 bg-orange-100 rounded-lg">
            <svg class="w-6 h-6 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
          </div>
          <div class="mr-3">
            <p class="text-sm text-orange-600">درآمد کل</p>
            <p class="text-xl font-semibold text-orange-900">{{ reportStats.totalRevenue.toLocaleString() }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- نمودارها -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <!-- نمودار روند ارجاعات -->
      <div class="bg-gray-50 p-6 rounded-lg">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">روند ارجاعات</h3>
        <div class="h-64 flex items-center justify-center text-gray-500">
          نمودار روند ارجاعات (در حال توسعه)
        </div>
      </div>

      <!-- نمودار توزیع پاداش‌ها -->
      <div class="bg-gray-50 p-6 rounded-lg">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">توزیع پاداش‌ها</h3>
        <div class="h-64 flex items-center justify-center text-gray-500">
          نمودار توزیع پاداش‌ها (در حال توسعه)
        </div>
      </div>
    </div>

    <!-- جدول برترین ارجاع‌دهندگان -->
    <div class="mt-8">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">برترین ارجاع‌دهندگان</h3>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                رتبه
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                کاربر
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                تعداد ارجاعات
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                ارجاعات موفق
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                کل پاداش
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                نرخ موفقیت
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="(referrer, index) in topReferrers" :key="referrer.id">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <span class="text-lg font-semibold text-gray-900">{{ index + 1 }}</span>
                  <div v-if="index < 3" class="mr-2">
                    <svg class="w-5 h-5 text-yellow-500" fill="currentColor" viewBox="0 0 20 20">
                      <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"></path>
                    </svg>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <img class="h-8 w-8 rounded-full ml-2" :src="referrer.avatar" :alt="referrer.name">
                  <div>
                    <div class="text-sm font-medium text-gray-900">{{ referrer.name }}</div>
                    <div class="text-sm text-gray-500">{{ referrer.email }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm text-gray-900">{{ referrer.totalReferrals }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm text-gray-900">{{ referrer.successfulReferrals }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm font-medium text-green-600">{{ referrer.totalReward.toLocaleString() }} تومان</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm text-gray-900">{{ referrer.successRate }}%</div>
              </td>
            </tr>
          </tbody>
        </table>
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

// نوع گزارش
const reportType = ref('monthly')

// آمار گزارش
const reportStats = ref({
  conversionRate: 65,
  avgReward: 25000,
  activeUsers: 1250,
  totalRevenue: 85000000
})

// برترین ارجاع‌دهندگان
const topReferrers = ref([
  {
    id: 1,
    name: 'علی احمدی',
    email: 'ali@example.com',
    avatar: '/default-avatar.png',
    totalReferrals: 45,
    successfulReferrals: 38,
    totalReward: 950000,
    successRate: 84
  },
  {
    id: 2,
    name: 'فاطمه محمدی',
    email: 'fateme@example.com',
    avatar: '/default-avatar.png',
    totalReferrals: 32,
    successfulReferrals: 28,
    totalReward: 700000,
    successRate: 88
  },
  {
    id: 3,
    name: 'محمد رضایی',
    email: 'mohammad@example.com',
    avatar: '/default-avatar.png',
    totalReferrals: 28,
    successfulReferrals: 22,
    totalReward: 550000,
    successRate: 79
  },
  {
    id: 4,
    name: 'زهرا کریمی',
    email: 'zahra@example.com',
    avatar: '/default-avatar.png',
    totalReferrals: 25,
    successfulReferrals: 20,
    totalReward: 500000,
    successRate: 80
  },
  {
    id: 5,
    name: 'احمد نوری',
    email: 'ahmad@example.com',
    avatar: '/default-avatar.png',
    totalReferrals: 22,
    successfulReferrals: 18,
    totalReward: 450000,
    successRate: 82
  }
])
</script> 
