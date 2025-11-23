<template>
  <div v-if="hasAccess" class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
    <div class="flex items-center justify-between mb-6">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">تحلیل زمانی پرداخت‌ها</h3>
        <p class="text-sm text-gray-500 mt-1">Time-based Payment Analysis</p>
      </div>
      <div class="flex items-center space-x-2 space-x-reverse">
        <button class="text-sm text-blue-600 hover:text-blue-800">مشاهده جزئیات</button>
      </div>
    </div>

    <!-- نمودار روند روزانه -->
    <div class="mb-6">
      <div class="flex items-center justify-between mb-4">
        <h4 class="font-medium text-gray-900">روند روزانه (آخرین ۷ روز)</h4>
        <div class="flex items-center space-x-4 space-x-reverse">
          <div class="flex items-center">
            <div class="w-3 h-3 bg-blue-500 rounded-full ml-2"></div>
            <span class="text-sm text-gray-600">مبلغ</span>
          </div>
          <div class="flex items-center">
            <div class="w-3 h-3 bg-green-500 rounded-full ml-2"></div>
            <span class="text-sm text-gray-600">تعداد</span>
          </div>
        </div>
      </div>
      
      <div class="h-32 bg-gray-50 rounded-lg p-6">
        <div class="flex items-end justify-between h-full">
          <div 
            v-for="(day, index) in dailyData" 
            :key="index"
            class="flex flex-col items-center"
            style="width: calc(100% / 7 - 8px)"
          >
            <div class="text-xs text-gray-500 mb-1">{{ day.name }}</div>
            <div
class="w-full bg-blue-500 rounded-t mb-1 relative group cursor-pointer"
                 :style="{ height: (day.amount / maxAmount * 100) + '%', minHeight: '10px' }">
              <div class="absolute -top-6 left-1/2 transform -translate-x-1/2 bg-gray-800 text-white text-xs px-2 py-1 rounded opacity-0 group-hover:opacity-100 transition-opacity whitespace-nowrap">
                {{ formatPrice(day.amount) }}
              </div>
            </div>
            <div
class="w-full bg-green-500 rounded-t relative group cursor-pointer"
                 :style="{ height: (day.count / maxCount * 100) + '%', minHeight: '10px' }">
              <div class="absolute -top-6 left-1/2 transform -translate-x-1/2 bg-gray-800 text-white text-xs px-2 py-1 rounded opacity-0 group-hover:opacity-100 transition-opacity whitespace-nowrap">
                {{ day.count }} تراکنش
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- آمار ساعتی -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
      <!-- ساعات اوج -->
      <div>
        <h4 class="font-medium text-gray-900 mb-3">ساعات اوج پرداخت</h4>
        <div class="space-y-2">
          <div 
            v-for="(hour, index) in peakHours" 
            :key="hour.hour"
            class="flex items-center justify-between p-2 bg-gray-50 rounded"
          >
            <div class="flex items-center">
              <div class="w-6 h-6 rounded-full bg-orange-100 flex items-center justify-center ml-2 text-xs font-medium text-orange-600">
                {{ index + 1 }}
              </div>
              <span class="text-sm font-medium">{{ hour.hour }}:۰۰</span>
            </div>
            <div class="text-right">
              <div class="text-sm font-semibold">{{ formatPrice(hour.amount) }}</div>
              <div class="text-xs text-gray-500">{{ hour.count }} تراکنش</div>
            </div>
          </div>
        </div>
      </div>

      <!-- روزهای هفته -->
      <div>
        <h4 class="font-medium text-gray-900 mb-3">عملکرد روزهای هفته</h4>
        <div class="space-y-2">
          <div 
            v-for="day in weekDays" 
            :key="day.name"
            class="flex items-center justify-between p-2 bg-gray-50 rounded"
          >
            <span class="text-sm font-medium">{{ day.name }}</span>
            <div class="text-right">
              <div class="text-sm font-semibold">{{ formatPrice(day.amount) }}</div>
              <div class="text-xs text-gray-500">{{ day.count }} تراکنش</div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- آمار کلی زمانی -->
    <div class="mt-6 pt-6 border-t border-gray-200">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div class="text-center">
          <div class="text-2xl font-bold text-blue-600">{{ formatPrice(averageDailyAmount) }}</div>
          <div class="text-sm text-gray-500">میانگین روزانه</div>
        </div>
        <div class="text-center">
          <div class="text-2xl font-bold text-green-600">{{ averageDailyCount }}</div>
          <div class="text-sm text-gray-500">میانگین تراکنش روزانه</div>
        </div>
        <div class="text-center">
          <div class="text-2xl font-bold text-purple-600">{{ peakHourTime }}</div>
          <div class="text-sm text-gray-500">ساعت اوج</div>
        </div>
        <div class="text-center">
          <div class="text-2xl font-bold text-orange-600">{{ bestDay }}</div>
          <div class="text-sm text-gray-500">بهترین روز</div>
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

// داده‌های روزانه
const dailyData = ref([
  { name: 'شنبه', amount: 8500000, count: 45 },
  { name: 'یکشنبه', amount: 9200000, count: 52 },
  { name: 'دوشنبه', amount: 7800000, count: 41 },
  { name: 'سه‌شنبه', amount: 10500000, count: 58 },
  { name: 'چهارشنبه', amount: 9800000, count: 54 },
  { name: 'پنج‌شنبه', amount: 11200000, count: 62 },
  { name: 'جمعه', amount: 6800000, count: 38 }
])

// ساعات اوج
const peakHours = ref([
  { hour: '14', amount: 4500000, count: 25 },
  { hour: '15', amount: 4200000, count: 23 },
  { hour: '16', amount: 3800000, count: 21 },
  { hour: '17', amount: 3500000, count: 19 },
  { hour: '18', amount: 3200000, count: 18 }
])

// روزهای هفته
const weekDays = ref([
  { name: 'شنبه', amount: 8500000, count: 45 },
  { name: 'یکشنبه', amount: 9200000, count: 52 },
  { name: 'دوشنبه', amount: 7800000, count: 41 },
  { name: 'سه‌شنبه', amount: 10500000, count: 58 },
  { name: 'چهارشنبه', amount: 9800000, count: 54 },
  { name: 'پنج‌شنبه', amount: 11200000, count: 62 },
  { name: 'جمعه', amount: 6800000, count: 38 }
])

// محاسبات
const maxAmount = computed(() => {
  return Math.max(...dailyData.value.map(day => day.amount))
})

const maxCount = computed(() => {
  return Math.max(...dailyData.value.map(day => day.count))
})

const averageDailyAmount = computed(() => {
  const total = dailyData.value.reduce((sum, day) => sum + day.amount, 0)
  return total / dailyData.value.length
})

const averageDailyCount = computed(() => {
  const total = dailyData.value.reduce((sum, day) => sum + day.count, 0)
  return Math.round(total / dailyData.value.length)
})

const peakHourTime = computed(() => {
  const peak = peakHours.value[0]
  return peak ? `${peak.hour}:۰۰` : '-'
})

const bestDay = computed(() => {
  const best = weekDays.value.reduce((max, day) => 
    day.amount > max.amount ? day : max
  )
  return best.name
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
