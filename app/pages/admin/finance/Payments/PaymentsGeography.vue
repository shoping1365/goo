<template>
  <div v-if="hasAccess" class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
    <div class="flex items-center justify-between mb-6">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">آمار جغرافیایی پرداخت‌ها</h3>
        <p class="text-sm text-gray-500 mt-1">Geographic Payment Statistics</p>
      </div>
      <div class="flex items-center space-x-2 space-x-reverse">
        <select class="text-sm border border-gray-300 rounded px-2 py-1">
          <option>آخرین ۳۰ روز</option>
          <option>آخرین ۷ روز</option>
          <option>آخرین ۹۰ روز</option>
        </select>
      </div>
    </div>

    <!-- نقشه ایران (نمایشی) -->
    <div class="mb-6">
      <div class="bg-gray-50 rounded-lg p-6 h-48 flex items-center justify-center">
        <div class="text-center">
          <svg class="w-16 h-16 text-gray-300 mx-auto mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 20l-5.447-2.724A1 1 0 013 16.382V5.618a1 1 0 011.447-.894L9 7m0 13l6-3m-6 3V7m6 10l4.553 2.276A1 1 0 0021 18.382V7.618a1 1 0 00-1.447-.894L15 4m0 13V4m-6 3l6-3"></path>
          </svg>
          <p class="text-sm text-gray-500">نقشه تعاملی ایران</p>
        </div>
      </div>
    </div>

    <!-- آمار استان‌ها -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
      <div>
        <h4 class="font-medium text-gray-900 mb-3">برترین استان‌ها</h4>
        <div class="space-y-3">
          <div 
            v-for="(province, index) in topProvinces" 
            :key="province.id"
            class="flex items-center justify-between p-3 bg-gray-50 rounded-lg"
          >
            <div class="flex items-center">
              <div class="w-6 h-6 rounded-full bg-blue-100 flex items-center justify-center ml-3 text-xs font-medium text-blue-600">
                {{ index + 1 }}
              </div>
              <div>
                <div class="font-medium text-gray-900">{{ province.name }}</div>
                <div class="text-sm text-gray-500">{{ province.count }} تراکنش</div>
              </div>
            </div>
            <div class="text-right">
              <div class="font-semibold text-gray-900">{{ formatPrice(province.amount) }}</div>
              <div class="text-xs text-gray-500">{{ province.percentage }}%</div>
            </div>
          </div>
        </div>
      </div>

      <div>
        <h4 class="font-medium text-gray-900 mb-3">شهرهای برتر</h4>
        <div class="space-y-3">
          <div 
            v-for="(city, index) in topCities" 
            :key="city.id"
            class="flex items-center justify-between p-3 bg-gray-50 rounded-lg"
          >
            <div class="flex items-center">
              <div class="w-6 h-6 rounded-full bg-green-100 flex items-center justify-center ml-3 text-xs font-medium text-green-600">
                {{ index + 1 }}
              </div>
              <div>
                <div class="font-medium text-gray-900">{{ city.name }}</div>
                <div class="text-sm text-gray-500">{{ city.province }}</div>
              </div>
            </div>
            <div class="text-right">
              <div class="font-semibold text-gray-900">{{ formatPrice(city.amount) }}</div>
              <div class="text-xs text-gray-500">{{ city.count }} تراکنش</div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- آمار کلی جغرافیایی -->
    <div class="mt-6 pt-6 border-t border-gray-200">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div class="text-center">
          <div class="text-2xl font-bold text-blue-600">{{ totalProvinces }}</div>
          <div class="text-sm text-gray-500">استان فعال</div>
        </div>
        <div class="text-center">
          <div class="text-2xl font-bold text-green-600">{{ totalCities }}</div>
          <div class="text-sm text-gray-500">شهر فعال</div>
        </div>
        <div class="text-center">
          <div class="text-2xl font-bold text-purple-600">{{ formatPrice(totalGeographicAmount) }}</div>
          <div class="text-sm text-gray-500">مجموع مبالغ</div>
        </div>
        <div class="text-center">
          <div class="text-2xl font-bold text-orange-600">{{ averagePerProvince }}</div>
          <div class="text-sm text-gray-500">میانگین هر استان</div>
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

// داده‌های استان‌های برتر
const topProvinces = ref([
  {
    id: 1,
    name: 'تهران',
    amount: 25000000,
    count: 89,
    percentage: 35
  },
  {
    id: 2,
    name: 'اصفهان',
    amount: 18000000,
    count: 67,
    percentage: 25
  },
  {
    id: 3,
    name: 'خراسان رضوی',
    amount: 12000000,
    count: 45,
    percentage: 17
  },
  {
    id: 4,
    name: 'فارس',
    amount: 9000000,
    count: 34,
    percentage: 13
  },
  {
    id: 5,
    name: 'آذربایجان شرقی',
    amount: 6000000,
    count: 23,
    percentage: 10
  }
])

// داده‌های شهرهای برتر
const topCities = ref([
  {
    id: 1,
    name: 'تهران',
    province: 'تهران',
    amount: 20000000,
    count: 78
  },
  {
    id: 2,
    name: 'اصفهان',
    province: 'اصفهان',
    amount: 15000000,
    count: 56
  },
  {
    id: 3,
    name: 'مشهد',
    province: 'خراسان رضوی',
    amount: 10000000,
    count: 38
  },
  {
    id: 4,
    name: 'شیراز',
    province: 'فارس',
    amount: 7000000,
    count: 28
  },
  {
    id: 5,
    name: 'تبریز',
    province: 'آذربایجان شرقی',
    amount: 5000000,
    count: 19
  }
])

// محاسبه آمار کلی
const totalProvinces = computed(() => topProvinces.value.length)

const totalCities = computed(() => topCities.value.length)

const totalGeographicAmount = computed(() => {
  return topProvinces.value.reduce((sum, province) => sum + province.amount, 0)
})

const averagePerProvince = computed(() => {
  return formatPrice(totalGeographicAmount.value / totalProvinces.value)
})

// فرمت کردن قیمت
const formatPrice = (price: number) => {
  if (price >= 1000000) {
    return (price / 1000000).toFixed(1) + 'M'
  } else if (price >= 1000) {
    return (price / 1000).toFixed(0) + 'K'
  }
  return price.toString()
}
</script> 
