<template>
  <div class="space-y-8">
    <!-- بخش مراحل کارت -->
    <section>
      <div class="bg-gradient-to-r from-indigo-50 to-purple-50 rounded-lg p-6 mb-6">
        <h2 class="text-2xl font-bold text-gray-900 mb-2">🔄 مراحل چرخه حیات کارت</h2>
        <p class="text-gray-600">مدیریت کامل مراحل مختلف زندگی گیفت کارت از ایجاد تا انقضا</p>
      </div>
      
      <!-- نمودار چرخه حیات -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">نمودار چرخه حیات</h3>
        <div class="flex flex-wrap justify-center items-center space-x-4 space-x-reverse">
          <div v-for="(stage, index) in lifecycleStages" :key="stage.name" class="flex items-center">
            <div class="text-center">
              <div
class="w-16 h-16 rounded-full flex items-center justify-center text-white text-lg font-bold mb-2"
                   :class="getStageColorClass(stage.status)">
                {{ stage.icon }}
              </div>
              <div class="text-sm font-medium text-gray-900">{{ stage.name }}</div>
              <div class="text-xs text-gray-500">{{ stage.count }} کارت</div>
            </div>
            <div v-if="index < lifecycleStages.length - 1" class="w-8 h-0.5 bg-gray-300 mx-2"></div>
          </div>
        </div>
      </div>
      
      <!-- جزئیات مراحل -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div
v-for="stage in lifecycleStages" :key="stage.name" 
             class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h4 class="text-lg font-semibold text-gray-900">{{ stage.name }}</h4>
            <span
class="text-xs rounded-full px-3 py-1"
                  :class="getStatusBadgeClass(stage.status)">
              {{ stage.status }}
            </span>
          </div>
          
          <div class="space-y-3">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">تعداد کارت‌ها:</span>
              <span class="text-lg font-bold" :class="getStageTextColor(stage.status)">
                {{ stage.count }}
              </span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">درصد کل:</span>
              <span class="text-sm font-medium text-gray-900">{{ stage.percentage }}%</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">میانگین مدت:</span>
              <span class="text-sm font-medium text-gray-900">{{ stage.avgDuration }}</span>
            </div>
          </div>
          
          <!-- نوار پیشرفت -->
          <div class="mt-4">
            <div class="flex justify-between text-sm text-gray-500 mb-1">
              <span>پیشرفت</span>
              <span>{{ stage.progress }}%</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div
class="h-2 rounded-full transition-all duration-300"
                   :class="getStageProgressColor(stage.status)"
                   :style="{ width: stage.progress + '%' }"></div>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-3 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
              مشاهده
            </button>
            <button class="px-3 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              گزارش
            </button>
          </div>
        </div>
      </div>
    </section>

    <!-- بخش عملیات ویژه -->
    <section>
      <div class="bg-gradient-to-r from-red-50 to-pink-50 rounded-lg p-6 mb-6">
        <h2 class="text-2xl font-bold text-gray-900 mb-2">⚡ عملیات ویژه</h2>
        <p class="text-gray-600">عملیات خاص و مدیریت پیشرفته گیفت کارت‌ها</p>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <!-- مسدود کردن کارت -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">مسدود کردن کارت</h3>
            <span class="text-xs bg-red-100 text-red-800 rounded-full px-3 py-1">عملیات حساس</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">کارت‌های مسدود شده:</span>
              <span class="text-lg font-bold text-red-600">{{ specialOpsStats.blockedCards }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">مسدود شده امروز:</span>
              <span class="text-lg font-bold text-orange-600">{{ specialOpsStats.blockedToday }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">در انتظار بررسی:</span>
              <span class="text-lg font-bold text-yellow-600">{{ specialOpsStats.pendingReview }}</span>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-red-600 text-white text-sm rounded-lg hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500">
              مسدود کردن
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              تاریخچه
            </button>
          </div>
        </div>

        <!-- تمدید اعتبار -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">تمدید اعتبار</h3>
            <span class="text-xs bg-green-100 text-green-800 rounded-full px-3 py-1">فعال</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">تمدید شده:</span>
              <span class="text-lg font-bold text-green-600">{{ specialOpsStats.renewedCards }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">در انتظار تمدید:</span>
              <span class="text-lg font-bold text-blue-600">{{ specialOpsStats.pendingRenewal }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">تمدید خودکار:</span>
              <span class="text-lg font-bold text-purple-600">{{ specialOpsStats.autoRenewed }}</span>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-green-600 text-white text-sm rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500">
              تمدید دستی
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              تنظیمات
            </button>
          </div>
        </div>

        <!-- انتقال اعتبار -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">انتقال اعتبار</h3>
            <span class="text-xs bg-blue-100 text-blue-800 rounded-full px-3 py-1">فعال</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">انتقال‌های انجام شده:</span>
              <span class="text-lg font-bold text-blue-600">{{ specialOpsStats.transfersCompleted }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">در انتظار تأیید:</span>
              <span class="text-lg font-bold text-yellow-600">{{ specialOpsStats.pendingTransfers }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">مبلغ کل انتقال‌ها:</span>
              <span class="text-lg font-bold text-green-600">{{ formatCurrency(specialOpsStats.totalTransferAmount) }}</span>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
              انتقال جدید
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              تأیید انتقال‌ها
            </button>
          </div>
        </div>

        <!-- ادغام کارت‌ها -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">ادغام کارت‌ها</h3>
            <span class="text-xs bg-purple-100 text-purple-800 rounded-full px-3 py-1">پیشرفته</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">ادغام‌های انجام شده:</span>
              <span class="text-lg font-bold text-purple-600">{{ specialOpsStats.mergesCompleted }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">در انتظار ادغام:</span>
              <span class="text-lg font-bold text-orange-600">{{ specialOpsStats.pendingMerges }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">کارت‌های ادغام شده:</span>
              <span class="text-lg font-bold text-green-600">{{ specialOpsStats.mergedCards }}</span>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-purple-600 text-white text-sm rounded-lg hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-purple-500">
              ادغام جدید
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              بررسی
            </button>
          </div>
        </div>
      </div>
    </section>

    <!-- داشبورد کلی چرخه حیات -->
    <section>
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-xl font-semibold text-gray-900 mb-4">📊 آمار کلی چرخه حیات</h3>
        
        <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
          <div class="text-center p-6 bg-blue-50 rounded-lg">
            <div class="text-2xl font-bold text-blue-600">{{ overallStats.totalCards }}</div>
            <div class="text-sm text-gray-600">کل کارت‌ها</div>
          </div>
          
          <div class="text-center p-6 bg-green-50 rounded-lg">
            <div class="text-2xl font-bold text-green-600">{{ overallStats.activeCards }}</div>
            <div class="text-sm text-gray-600">کارت‌های فعال</div>
          </div>
          
          <div class="text-center p-6 bg-yellow-50 rounded-lg">
            <div class="text-2xl font-bold text-yellow-600">{{ overallStats.expiredCards }}</div>
            <div class="text-sm text-gray-600">کارت‌های منقضی</div>
          </div>
          
          <div class="text-center p-6 bg-purple-50 rounded-lg">
            <div class="text-2xl font-bold text-purple-600">{{ overallStats.specialOperations }}</div>
            <div class="text-sm text-gray-600">عملیات ویژه</div>
          </div>
        </div>
        
        <!-- نمودار توزیع وضعیت‌ها -->
        <div class="mt-6">
          <h4 class="text-sm font-medium text-gray-900 mb-3">توزیع وضعیت‌ها</h4>
          <div class="flex items-center space-x-4 space-x-reverse">
            <div v-for="status in statusDistribution" :key="status.name" class="flex items-center">
              <div :class="getStatusColorClass(status.name)" class="w-3 h-3 rounded-full mr-2"></div>
              <span class="text-sm text-gray-600">{{ status.label }}: {{ status.count }}</span>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
// مراحل چرخه حیات
const lifecycleStages = [
  {
    name: 'ایجاد',
    icon: '🎫',
    status: 'active',
    count: 15420,
    percentage: 100,
    avgDuration: '1 دقیقه',
    progress: 100
  },
  {
    name: 'فعال‌سازی',
    icon: '✅',
    status: 'active',
    count: 14890,
    percentage: 96.6,
    avgDuration: '5 دقیقه',
    progress: 96
  },
  {
    name: 'استفاده',
    icon: '💳',
    status: 'active',
    count: 12450,
    percentage: 80.8,
    avgDuration: 'متغیر',
    progress: 80
  },
  {
    name: 'انقضا',
    icon: '⏰',
    status: 'warning',
    count: 2340,
    percentage: 15.2,
    avgDuration: '1 سال',
    progress: 15
  },
  {
    name: 'تمدید',
    icon: '🔄',
    status: 'success',
    count: 890,
    percentage: 5.8,
    avgDuration: '30 روز',
    progress: 6
  }
]

// آمار عملیات ویژه
const specialOpsStats = {
  blockedCards: 156,
  blockedToday: 12,
  pendingReview: 23,
  renewedCards: 890,
  pendingRenewal: 234,
  autoRenewed: 567,
  transfersCompleted: 445,
  pendingTransfers: 67,
  totalTransferAmount: 125000000,
  mergesCompleted: 89,
  pendingMerges: 12,
  mergedCards: 178
}

// آمار کلی
const overallStats = {
  totalCards: 15420,
  activeCards: 12450,
  expiredCards: 2340,
  specialOperations: 1567
}

// توزیع وضعیت‌ها
const statusDistribution = [
  { name: 'active', label: 'فعال', count: 12450 },
  { name: 'used', label: 'استفاده شده', count: 8900 },
  { name: 'expired', label: 'منقضی', count: 2340 },
  { name: 'blocked', label: 'مسدود', count: 156 },
  { name: 'pending', label: 'در انتظار', count: 574 }
]

// توابع کمکی
const getStageColorClass = (status: string) => {
  const classes = {
    active: 'bg-green-500',
    warning: 'bg-yellow-500',
    success: 'bg-blue-500',
    danger: 'bg-red-500'
  }
  return classes[status] || 'bg-gray-500'
}

const getStageTextColor = (status: string) => {
  const classes = {
    active: 'text-green-600',
    warning: 'text-yellow-600',
    success: 'text-blue-600',
    danger: 'text-red-600'
  }
  return classes[status] || 'text-gray-600'
}

const getStageProgressColor = (status: string) => {
  const classes = {
    active: 'bg-green-500',
    warning: 'bg-yellow-500',
    success: 'bg-blue-500',
    danger: 'bg-red-500'
  }
  return classes[status] || 'bg-gray-500'
}

const getStatusBadgeClass = (status: string) => {
  const classes = {
    active: 'bg-green-100 text-green-800',
    warning: 'bg-yellow-100 text-yellow-800',
    success: 'bg-blue-100 text-blue-800',
    danger: 'bg-red-100 text-red-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getStatusColorClass = (status: string) => {
  const classes = {
    active: 'bg-green-500',
    used: 'bg-blue-500',
    expired: 'bg-red-500',
    blocked: 'bg-gray-500',
    pending: 'bg-yellow-500'
  }
  return classes[status] || 'bg-gray-500'
}

const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}
</script>

<!--
  مستندسازی:
  این کامپوننت شامل دو بخش اصلی است:
  1. مراحل چرخه حیات: ایجاد، فعال‌سازی، استفاده، انقضا، تمدید
  2. عملیات ویژه: مسدود کردن، تمدید اعتبار، انتقال اعتبار، ادغام کارت‌ها
  
  تمام توضیحات به فارسی و با طراحی ریسپانسیو ارائه شده است.
  نمودار چرخه حیات و آمار دقیق برای نمایش وضعیت‌ها استفاده شده‌اند.
--> 
