<template>
  <div v-if="hasAccess" class="space-y-6">

    <!-- ردیف دوم: وضعیت پرداخت‌ها -->
    <div class="bg-white rounded-xl shadow-sm border-8 border-white p-6">
      <div class="flex items-center justify-between mb-6">
        <div>
          <h3 class="text-lg font-semibold text-gray-900">وضعیت پرداخت‌ها</h3>
          <p class="text-sm text-gray-500 mt-1">Payment Status Distribution</p>
        </div>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div 
          v-for="item in paymentStatus" 
          :key="item.status"
          class="flex items-center justify-between p-6 rounded-lg"
          :class="getStatusBgClass(item.status)"
        >
          <div class="flex items-center">
            <div 
              class="w-3 h-3 rounded-full ml-3"
              :class="getStatusColorClass(item.status)"
            ></div>
            <div>
              <div class="font-medium" :class="getStatusTextClass(item.status)">
                {{ getStatusName(item.status) }}
              </div>
              <div class="text-sm" :class="getStatusSubTextClass(item.status)">
                {{ item.percentage }}% از کل
              </div>
            </div>
          </div>
          <div class="text-left">
            <div class="font-semibold" :class="getStatusTextClass(item.status)">
              {{ item.count }}
            </div>
            <div class="text-xs" :class="getStatusSubTextClass(item.status)">
              تراکنش
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- ردیف سوم: نمودارها -->
    <div class="grid grid-cols-1 xl:grid-cols-2 gap-6">
      <!-- نمودار روند پرداخت‌ها -->
      <div class="bg-white rounded-xl shadow-sm border-8 border-white p-6">
        <div class="flex items-center justify-between mb-6">
          <div>
            <h3 class="text-xl font-semibold text-gray-900">نمودار روند پرداخت‌های ۷ روز گذشته</h3>
            <p class="text-sm text-gray-500 mt-1">Payment Trends - Last 7 Days</p>
          </div>
          <div class="flex items-center space-x-4 space-x-reverse">
            <div class="flex items-center">
              <div class="w-3 h-3 bg-blue-500 rounded-full ml-2"></div>
              <span class="text-sm text-gray-600">مبلغ (میلیون تومان)</span>
            </div>
            <div class="flex items-center">
              <div class="w-3 h-3 bg-green-500 rounded-full ml-2"></div>
              <span class="text-sm text-gray-600">تعداد تراکنش</span>
            </div>
          </div>
        </div>
        
        <!-- نمودار واقعی -->
        <div class="h-80 bg-gradient-to-br from-gray-50 to-gray-100 rounded-3xl p-6 relative overflow-hidden">
          <!-- محور Y -->
          <div class="absolute left-6 top-6 bottom-6 w-16 flex flex-col justify-between text-xs text-gray-500">
            <div>30M</div>
            <div>25M</div>
            <div>20M</div>
            <div>15M</div>
            <div>10M</div>
            <div>5M</div>
            <div>0M</div>
          </div>
          
          <!-- نمودار ستونی -->
          <div class="absolute left-24 right-6 top-6 bottom-16 flex items-end justify-between">
            <div 
              v-for="(day, index) in chartData" 
              :key="index" 
              class="flex flex-col items-center"
              style="width: calc(100% / 7 - 8px)"
            >
              <!-- ستون مبلغ -->
              <div 
                class="w-full bg-blue-500 rounded-t-lg relative group cursor-pointer transition-all duration-300 hover:bg-blue-600"
                :style="{ height: (day.amount / 30000000 * 100) + '%', minHeight: '20px' }"
              >
                <div class="absolute -top-8 left-1/2 transform -translate-x-1/2 bg-gray-800 text-white text-xs px-2 py-1 rounded opacity-0 group-hover:opacity-100 transition-opacity whitespace-nowrap">
                  {{ formatPrice(day.amount) }}
                </div>
              </div>
              
              <!-- ستون تعداد -->
              <div 
                class="w-full bg-green-500 rounded-t-lg mt-2 relative group cursor-pointer transition-all duration-300 hover:bg-green-600"
                :style="{ height: (day.count / 100 * 100) + '%', minHeight: '10px' }"
              >
                <div class="absolute -top-8 left-1/2 transform -translate-x-1/2 bg-gray-800 text-white text-xs px-2 py-1 rounded opacity-0 group-hover:opacity-100 transition-opacity whitespace-nowrap">
                  {{ day.count }} تراکنش
                </div>
              </div>
            </div>
          </div>
          
          <!-- محور X -->
          <div class="absolute left-24 right-6 bottom-6 flex justify-between text-xs text-gray-500">
            <div v-for="(day, index) in chartData" :key="index" class="text-center" style="width: calc(100% / 7 - 8px)">
              {{ day.date }}
            </div>
          </div>
          
          <!-- خطوط راهنما -->
          <div class="absolute left-24 right-6 top-6 bottom-16 pointer-events-none">
            <div class="h-full flex flex-col justify-between">
              <div v-for="i in 6" :key="i" class="border-t border-gray-200 border-dashed"></div>
            </div>
          </div>
        </div>
      </div>
      
      <!-- توزیع روش‌های پرداخت -->
      <div class="bg-white rounded-xl shadow-sm border-8 border-white p-6">
        <div class="flex items-center justify-between mb-6">
          <div>
            <h3 class="text-lg font-semibold text-gray-900">توزیع روش‌های پرداخت</h3>
            <p class="text-sm text-gray-500 mt-1">Payment Methods Distribution</p>
          </div>
        </div>
        
        <div class="flex items-center justify-center">
          <!-- نمودار دایره‌ای -->
          <div class="relative w-32 h-32">
            <svg class="w-32 h-32 transform -rotate-90" viewBox="0 0 100 100">
              <!-- کارت به کارت -->
              <circle
                cx="50"
                cy="50"
                r="40"
                fill="none"
                stroke="#3B82F6"
                stroke-width="8"
                :stroke-dasharray="`${(paymentMethods.cardToCard.percentage / 100) * 251.2} 251.2`"
                stroke-dashoffset="0"
              />
              <!-- درگاه بانکی -->
              <circle
                cx="50"
                cy="50"
                r="40"
                fill="none"
                stroke="#10B981"
                stroke-width="8"
                :stroke-dasharray="`${(paymentMethods.bankGateway.percentage / 100) * 251.2} 251.2`"
                :stroke-dashoffset="`-${(paymentMethods.cardToCard.percentage / 100) * 251.2}`"
              />
              <!-- کیف پول -->
              <circle
                cx="50"
                cy="50"
                r="40"
                fill="none"
                stroke="#8B5CF6"
                stroke-width="8"
                :stroke-dasharray="`${(paymentMethods.wallet.percentage / 100) * 251.2} 251.2`"
                :stroke-dashoffset="`-${((paymentMethods.cardToCard.percentage + paymentMethods.bankGateway.percentage) / 100) * 251.2}`"
              />
              <!-- پرداخت نقدی -->
              <circle
                cx="50"
                cy="50"
                r="40"
                fill="none"
                stroke="#F59E0B"
                stroke-width="8"
                :stroke-dasharray="`${(paymentMethods.cash.percentage / 100) * 251.2} 251.2`"
                :stroke-dashoffset="`-${((paymentMethods.cardToCard.percentage + paymentMethods.bankGateway.percentage + paymentMethods.wallet.percentage) / 100) * 251.2}`"
              />
            </svg>
            
            <!-- مرکز نمودار -->
            <div class="absolute inset-0 flex items-center justify-center">
              <div class="text-center">
                <div class="text-sm font-bold text-gray-900">{{ totalPaymentMethods }}%</div>
                <div class="text-xs text-gray-500">کل</div>
              </div>
            </div>
          </div>
        </div>
        
        <!-- راهنمای رنگ‌ها -->
        <div class="grid grid-cols-1 gap-2 mt-6">
          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-3 h-3 bg-blue-500 rounded-full ml-2"></div>
              <span class="text-sm text-gray-600">کارت به کارت</span>
            </div>
            <span class="text-sm font-medium text-gray-900">{{ paymentMethods.cardToCard.percentage }}%</span>
          </div>
          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-3 h-3 bg-green-500 rounded-full ml-2"></div>
              <span class="text-sm text-gray-600">درگاه بانکی</span>
            </div>
            <span class="text-sm font-medium text-gray-900">{{ paymentMethods.bankGateway.percentage }}%</span>
          </div>
          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-3 h-3 bg-purple-500 rounded-full ml-2"></div>
              <span class="text-sm text-gray-600">کیف پول</span>
            </div>
            <span class="text-sm font-medium text-gray-900">{{ paymentMethods.wallet.percentage }}%</span>
          </div>
          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-3 h-3 bg-orange-500 rounded-full ml-2"></div>
              <span class="text-sm text-gray-600">پرداخت نقدی</span>
            </div>
            <span class="text-sm font-medium text-gray-900">{{ paymentMethods.cash.percentage }}%</span>
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

// داده‌های وضعیت پرداخت‌ها
const paymentStatus = ref([
  { status: 'success', count: 892, percentage: 89.2 },
  { status: 'failed', count: 67, percentage: 6.7 },
  { status: 'pending', count: 23, percentage: 2.3 }
])

// داده‌های روش‌های پرداخت
const paymentMethods = ref({
  cardToCard: {
    amount: 45621000,
    percentage: 36.7
  },
  bankGateway: {
    amount: 38934000,
    percentage: 31.2
  },
  wallet: {
    amount: 24521000,
    percentage: 19.7
  },
  cash: {
    amount: 12847000,
    percentage: 10.3
  }
})

// محاسبه کل روش‌های پرداخت
const totalPaymentMethods = computed(() => {
  return Object.values(paymentMethods.value).reduce((sum, method) => sum + method.percentage, 0)
})

// محاسبه کل تراکنش‌های پرداخت
const _totalPaymentTransactions = computed(() =>
  paymentStatus.value.reduce((sum, item) => sum + item.count, 0)
)

// توابع کمکی برای وضعیت پرداخت‌ها
const getStatusName = (status: string) => {
  const names = {
    'success': 'موفق',
    'failed': 'ناموفق',
    'pending': 'در انتظار',
    'cancelled': 'لغو شده'
  }
  return names[status as keyof typeof names] || status
}

const getStatusColorClass = (status: string) => {
  const colors = {
    'success': 'bg-green-500',
    'failed': 'bg-red-500',
    'pending': 'bg-yellow-500',
    'cancelled': 'bg-gray-500'
  }
  return colors[status as keyof typeof colors] || 'bg-gray-500'
}

const getStatusBgClass = (status: string) => {
  const colors = {
    'success': 'bg-green-50',
    'failed': 'bg-red-50',
    'pending': 'bg-yellow-50',
    'cancelled': 'bg-gray-50'
  }
  return colors[status as keyof typeof colors] || 'bg-gray-50'
}

const getStatusTextClass = (status: string) => {
  const colors = {
    'success': 'text-green-800',
    'failed': 'text-red-800',
    'pending': 'text-yellow-800',
    'cancelled': 'text-gray-800'
  }
  return colors[status as keyof typeof colors] || 'text-gray-800'
}

const getStatusSubTextClass = (status: string) => {
  const colors = {
    'success': 'text-green-600',
    'failed': 'text-red-600',
    'pending': 'text-yellow-600',
    'cancelled': 'text-gray-600'
  }
  return colors[status as keyof typeof colors] || 'text-gray-600'
}

// داده‌های نمودار
const chartData = ref([
  { date: '01/08', amount: 12500000, count: 45 },
  { date: '02/08', amount: 15800000, count: 52 },
  { date: '03/08', amount: 14200000, count: 48 },
  { date: '04/08', amount: 18900000, count: 63 },
  { date: '05/08', amount: 23400000, count: 78 },
  { date: '06/08', amount: 26700000, count: 89 },
  { date: '07/08', amount: 19800000, count: 66 }
])

// فرمت قیمت به تومان
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
  کامپوننت آمار و خلاصه وضعیت پرداخت‌ها
  - طراحی مدرن با گرادیان‌های رنگی
  - کارت‌های آمار با آیکون‌های مناسب
  - نمایش روند رشد با درصدها
  - نمودار روند پرداخت‌ها (واقعی)
  - نمودار دایره‌ای توزیع روش‌های پرداخت
  - کاملاً ریسپانسیو
  توضیحات کامل به فارسی در کد
--> 
