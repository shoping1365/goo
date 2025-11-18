<template>
  <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
    <!-- هدر اصلی -->
    <div class="flex flex-col lg:flex-row lg:items-center lg:justify-between gap-6 mb-6">
      <div>
        <h1 class="text-2xl font-bold text-gray-900">اتصال حسابداری</h1>
        <p class="text-gray-600 mt-2">مدیریت اتصال و همگام‌سازی با نرم‌افزارهای حسابداری و ERP</p>
      </div>
      
      <!-- دکمه‌های عملیاتی -->
      <div class="flex flex-wrap gap-3">
        <button 
          @click="testAllConnections"
          class="inline-flex items-center gap-2 px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg transition-colors duration-200"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
          </svg>
          تست تمام اتصالات
        </button>
        
        <button 
          @click="syncAllData"
          class="inline-flex items-center gap-2 px-4 py-2 bg-green-600 hover:bg-green-700 text-white rounded-lg transition-colors duration-200"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
          </svg>
          همگام‌سازی کامل
        </button>
        
        <button 
          @click="showSettings"
          class="inline-flex items-center gap-2 px-4 py-2 bg-gray-600 hover:bg-gray-700 text-white rounded-lg transition-colors duration-200"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"></path>
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
          </svg>
          تنظیمات
        </button>
      </div>
    </div>

    <!-- آمار کلی -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <!-- اتصالات فعال -->
      <div class="bg-blue-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-blue-600">{{ connectionStats.active }}</div>
            <div class="text-sm text-blue-700">اتصالات فعال</div>
          </div>
          <div class="w-10 h-10 bg-blue-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
            </svg>
          </div>
        </div>
        <div class="mt-2 text-xs text-blue-600">
          {{ connectionStats.total }} کل اتصال
        </div>
      </div>

      <!-- تراکنش‌های همگام‌سازی شده -->
      <div class="bg-green-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-green-600">{{ formatNumber(connectionStats.syncedTransactions) }}</div>
            <div class="text-sm text-green-700">تراکنش همگام شده</div>
          </div>
          <div class="w-10 h-10 bg-green-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
        <div class="mt-2 text-xs text-green-600">
          امروز: {{ formatNumber(connectionStats.todayTransactions) }}
        </div>
      </div>

      <!-- نرخ موفقیت -->
      <div class="bg-purple-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-purple-600">{{ connectionStats.successRate }}%</div>
            <div class="text-sm text-purple-700">نرخ موفقیت</div>
          </div>
          <div class="w-10 h-10 bg-purple-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
          </div>
        </div>
        <div class="mt-2 text-xs text-purple-600">
          {{ connectionStats.totalAttempts }} تلاش کل
        </div>
      </div>

      <!-- خطاهای اخیر -->
      <div class="bg-red-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-red-600">{{ connectionStats.recentErrors }}</div>
            <div class="text-sm text-red-700">خطاهای اخیر</div>
          </div>
          <div class="w-10 h-10 bg-red-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
        <div class="mt-2 text-xs text-red-600">
          آخرین: {{ connectionStats.lastErrorTime }}
        </div>
      </div>
    </div>

    <!-- وضعیت اتصالات -->
    <div class="mt-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">وضعیت اتصالات</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div 
          v-for="connection in connectionStatus" 
          :key="connection.id"
          class="bg-white border rounded-lg p-6"
          :class="connection.status === 'active' ? 'border-green-200' : connection.status === 'error' ? 'border-red-200' : 'border-yellow-200'"
        >
          <div class="flex items-center justify-between mb-2">
            <div class="flex items-center gap-2">
              <div 
                class="w-3 h-3 rounded-full"
                :class="connection.status === 'active' ? 'bg-green-500' : connection.status === 'error' ? 'bg-red-500' : 'bg-yellow-500'"
              ></div>
              <span class="font-medium text-gray-900">{{ connection.name }}</span>
            </div>
            <span 
              class="text-xs px-2 py-1 rounded-full"
              :class="connection.status === 'active' ? 'bg-green-100 text-green-700' : connection.status === 'error' ? 'bg-red-100 text-red-700' : 'bg-yellow-100 text-yellow-700'"
            >
              {{ getStatusLabel(connection.status) }}
            </span>
          </div>
          <div class="text-sm text-gray-600">{{ connection.description }}</div>
          <div class="mt-2 text-xs text-gray-500">
            آخرین همگام‌سازی: {{ connection.lastSync }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'

// آمار اتصالات
const connectionStats = ref({
  active: 3,
  total: 5,
  syncedTransactions: 15420,
  todayTransactions: 1250,
  successRate: 98.5,
  totalAttempts: 15680,
  recentErrors: 2,
  lastErrorTime: '2 ساعت پیش'
})

// وضعیت اتصالات
const connectionStatus = ref([
  {
    id: 1,
    name: 'هلو',
    status: 'active',
    description: 'نرم‌افزار حسابداری هلو',
    lastSync: '5 دقیقه پیش'
  },
  {
    id: 2,
    name: 'پارسیان',
    status: 'active',
    description: 'نرم‌افزار حسابداری پارسیان',
    lastSync: '10 دقیقه پیش'
  },
  {
    id: 3,
    name: 'سپیدار',
    status: 'warning',
    description: 'نرم‌افزار حسابداری سپیدار',
    lastSync: '1 ساعت پیش'
  },
  {
    id: 4,
    name: 'قیاس',
    status: 'error',
    description: 'نرم‌افزار حسابداری قیاس',
    lastSync: '2 ساعت پیش'
  },
  {
    id: 5,
    name: 'دیگر نرم‌افزارها',
    status: 'inactive',
    description: 'سایر نرم‌افزارهای حسابداری',
    lastSync: 'هرگز'
  }
])

// فرمت اعداد
const formatNumber = (num: number): string => {
  return new Intl.NumberFormat('fa-IR').format(num)
}

// برچسب وضعیت
const getStatusLabel = (status: string): string => {
  const labels = {
    active: 'فعال',
    warning: 'هشدار',
    error: 'خطا',
    inactive: 'غیرفعال'
  }
  return labels[status] || status
}

// تست تمام اتصالات
const testAllConnections = async () => {
  try {
    // TODO: تست تمام اتصالات
    console.log('تست تمام اتصالات شروع شد')
  } catch (error) {
    console.error('خطا در تست اتصالات:', error)
  }
}

// همگام‌سازی کامل
const syncAllData = async () => {
  try {
    // TODO: همگام‌سازی کامل داده‌ها
    console.log('همگام‌سازی کامل شروع شد')
  } catch (error) {
    console.error('خطا در همگام‌سازی:', error)
  }
}

// نمایش تنظیمات
const showSettings = () => {
  // TODO: نمایش تنظیمات
  console.log('نمایش تنظیمات')
}

// بارگذاری اولیه
onMounted(() => {
  // TODO: بارگذاری آمار اتصالات از API
})
</script>

<!--
  کامپوننت هدر صفحه اتصال حسابداری
  شامل:
  1. عنوان و توضیحات صفحه
  2. دکمه‌های عملیاتی اصلی
  3. آمار کلی اتصالات
  4. وضعیت اتصالات مختلف
  5. طراحی مدرن و کاملاً ریسپانسیو
--> 
