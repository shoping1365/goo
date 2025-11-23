<template>
  <div v-if="hasAccess" class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
    <div class="flex items-center justify-between mb-6">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">لاگ و رویدادها</h3>
        <p class="text-sm text-gray-500 mt-1">Payment Logs & Events</p>
      </div>
      <div class="flex items-center space-x-2 space-x-reverse">
        <button 
          class="inline-flex items-center px-3 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200 transition-colors"
          @click="refreshLogs"
        >
          <svg class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
          </svg>
          بروزرسانی
        </button>
        <button 
          class="inline-flex items-center px-3 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 transition-colors"
          @click="exportLogs"
        >
          <svg class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
          </svg>
          خروجی
        </button>
      </div>
    </div>

    <!-- فیلترهای لاگ -->
    <div class="mb-4 p-6 bg-gray-50 rounded-lg">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">سطح لاگ</label>
          <select 
            v-model="logFilters.level"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="all">همه سطوح</option>
            <option value="error">خطا</option>
            <option value="warning">هشدار</option>
            <option value="info">اطلاعات</option>
            <option value="debug">دیباگ</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">نوع رویداد</label>
          <select 
            v-model="logFilters.type"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="all">همه رویدادها</option>
            <option value="payment">پرداخت</option>
            <option value="refund">مرجوع</option>
            <option value="gateway">درگاه</option>
            <option value="system">سیستم</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">از تاریخ</label>
          <input 
            v-model="logFilters.dateFrom" 
            type="date"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">تا تاریخ</label>
          <input 
            v-model="logFilters.dateTo" 
            type="date"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
        </div>
      </div>
    </div>

    <!-- لیست لاگ‌ها -->
    <div class="space-y-3 max-h-96 overflow-y-auto">
      <div 
        v-for="log in filteredLogs" 
        :key="log.id"
        class="p-6 border border-gray-200 rounded-lg hover:bg-gray-50 transition-colors"
      >
        <div class="flex items-start justify-between">
          <div class="flex items-start space-x-3 space-x-reverse">
            <!-- آیکون سطح لاگ -->
            <div 
              :class="{
                'w-8 h-8 rounded-full flex items-center justify-center': true,
                'bg-red-100': log.level === 'error',
                'bg-yellow-100': log.level === 'warning',
                'bg-blue-100': log.level === 'info',
                'bg-gray-100': log.level === 'debug'
              }"
            >
              <svg 
                v-if="log.level === 'error'"
                class="w-4 h-4 text-red-600" 
                fill="none" 
                stroke="currentColor" 
                viewBox="0 0 24 24"
              >
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
              </svg>
              <svg 
                v-else-if="log.level === 'warning'"
                class="w-4 h-4 text-yellow-600" 
                fill="none" 
                stroke="currentColor" 
                viewBox="0 0 24 24"
              >
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
              </svg>
              <svg 
                v-else-if="log.level === 'info'"
                class="w-4 h-4 text-blue-600" 
                fill="none" 
                stroke="currentColor" 
                viewBox="0 0 24 24"
              >
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
              <svg 
                v-else
                class="w-4 h-4 text-gray-600" 
                fill="none" 
                stroke="currentColor" 
                viewBox="0 0 24 24"
              >
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
            
            <div class="flex-1">
              <div class="flex items-center space-x-2 space-x-reverse mb-1">
                <span 
                  :class="{
                    'text-xs font-medium px-2 py-1 rounded-full': true,
                    'bg-red-100 text-red-800': log.level === 'error',
                    'bg-yellow-100 text-yellow-800': log.level === 'warning',
                    'bg-blue-100 text-blue-800': log.level === 'info',
                    'bg-gray-100 text-gray-800': log.level === 'debug'
                  }"
                >
                  {{ log.level.toUpperCase() }}
                </span>
                <span class="text-xs text-gray-500">{{ log.type }}</span>
                <span class="text-xs text-gray-400">{{ formatDate(log.timestamp) }}</span>
              </div>
              <p class="text-sm text-gray-900 mb-1">{{ log.message }}</p>
              <div v-if="log.details" class="text-xs text-gray-600 bg-gray-50 p-2 rounded">
                {{ log.details }}
              </div>
            </div>
          </div>
          
          <button 
            class="text-gray-400 hover:text-gray-600"
            @click="viewLogDetails(log)"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
            </svg>
          </button>
        </div>
      </div>
    </div>

    <!-- آمار لاگ‌ها -->
    <div class="mt-4 pt-4 border-t border-gray-200">
      <div class="grid grid-cols-2 md:grid-cols-4 gap-6">
        <div class="text-center">
          <div class="text-2xl font-bold text-red-600">{{ logStats.errors }}</div>
          <div class="text-sm text-gray-500">خطاها</div>
        </div>
        <div class="text-center">
          <div class="text-2xl font-bold text-yellow-600">{{ logStats.warnings }}</div>
          <div class="text-sm text-gray-500">هشدارها</div>
        </div>
        <div class="text-center">
          <div class="text-2xl font-bold text-blue-600">{{ logStats.info }}</div>
          <div class="text-sm text-gray-500">اطلاعات</div>
        </div>
        <div class="text-center">
          <div class="text-2xl font-bold text-gray-600">{{ logStats.total }}</div>
          <div class="text-sm text-gray-500">کل</div>
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

// داده‌های نمونه لاگ‌ها
const logs = ref([
  {
    id: 1,
    level: 'error',
    type: 'gateway',
    message: 'خطا در اتصال به درگاه زرین‌پال',
    details: 'Timeout error: Connection failed after 30 seconds',
    timestamp: new Date('2024-01-15T10:30:00')
  },
  {
    id: 2,
    level: 'warning',
    type: 'payment',
    message: 'پرداخت با مبلغ غیرمعمول',
    details: 'Amount: 50,000,000 تومان - User: user123',
    timestamp: new Date('2024-01-15T10:25:00')
  },
  {
    id: 3,
    level: 'info',
    type: 'payment',
    message: 'پرداخت موفق انجام شد',
    details: 'Transaction ID: TXN123456 - Amount: 1,250,000 تومان',
    timestamp: new Date('2024-01-15T10:20:00')
  },
  {
    id: 4,
    level: 'info',
    type: 'refund',
    message: 'مرجوع وجه انجام شد',
    details: 'Refund ID: REF789 - Original TXN: TXN123456',
    timestamp: new Date('2024-01-15T10:15:00')
  },
  {
    id: 5,
    level: 'debug',
    type: 'system',
    message: 'بررسی اتصال درگاه‌ها',
    details: 'Checking gateway connections...',
    timestamp: new Date('2024-01-15T10:10:00')
  }
])

const logFilters = ref({
  level: 'all',
  type: 'all',
  dateFrom: '',
  dateTo: ''
})

// فیلتر کردن لاگ‌ها
const filteredLogs = computed(() => {
  return logs.value.filter(log => {
    if (logFilters.value.level !== 'all' && log.level !== logFilters.value.level) return false
    if (logFilters.value.type !== 'all' && log.type !== logFilters.value.type) return false
    // TODO: فیلتر تاریخ
    return true
  })
})

// آمار لاگ‌ها
const logStats = computed(() => {
  const stats = {
    errors: 0,
    warnings: 0,
    info: 0,
    total: logs.value.length
  }
  
  logs.value.forEach(log => {
    switch (log.level) {
      case 'error':
        stats.errors++
        break
      case 'warning':
        stats.warnings++
        break
      case 'info':
        stats.info++
        break
    }
  })
  
  return stats
})

const formatDate = (date: Date) => {
  return new Intl.DateTimeFormat('fa-IR', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit'
  }).format(date)
}

const refreshLogs = () => {
  // TODO: بروزرسانی لاگ‌ها از سرور
}

const exportLogs = () => {
  // TODO: خروجی لاگ‌ها
}

interface Log {
  [key: string]: unknown
}

const viewLogDetails = (_log: Log) => {
  // TODO: نمایش جزئیات لاگ
}
</script>

<!--
  کامپوننت لاگ و رویدادها
  - نمایش لاگ‌های سیستم
  - فیلتر بر اساس سطح و نوع
  - آمار لاگ‌ها
  - کاملاً ریسپانسیو
  توضیحات کامل به فارسی در کد
--> 
