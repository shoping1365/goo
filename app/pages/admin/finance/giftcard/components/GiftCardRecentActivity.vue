<template>
  <div class="bg-white rounded-2xl border border-gray-200/50 shadow-sm p-6 mb-8">
    <div class="flex items-center justify-between mb-6">
      <h2 class="text-xl font-bold text-gray-900">فعالیت‌های اخیر</h2>
      <button class="text-sm text-blue-600 hover:text-blue-700 font-medium transition-colors">
        مشاهده همه
      </button>
    </div>
    
    <div class="space-y-4">
      <div 
        v-for="activity in activities" 
        :key="activity.id"
        class="flex items-center space-x-4 space-x-reverse p-6 bg-gray-50 rounded-xl hover:bg-gray-100 transition-colors"
      >
        <!-- آیکون فعالیت -->
        <div class="flex-shrink-0">
          <div 
            :class="[
              activity.type === 'created' ? 'bg-green-100 text-green-600' : '',
              activity.type === 'used' ? 'bg-blue-100 text-blue-600' : '',
              activity.type === 'expired' ? 'bg-orange-100 text-orange-600' : '',
              activity.type === 'blocked' ? 'bg-red-100 text-red-600' : '',
              'w-12 h-12 rounded-xl flex items-center justify-center'
            ]"
          >
            <svg v-if="activity.type === 'created'" class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path>
            </svg>
            <svg v-else-if="activity.type === 'used'" class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
            </svg>
            <svg v-else-if="activity.type === 'expired'" class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
            <svg v-else-if="activity.type === 'blocked'" class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728L5.636 5.636m12.728 12.728L18.364 5.636M5.636 18.364l12.728-12.728"></path>
            </svg>
          </div>
        </div>
        
        <!-- جزئیات فعالیت -->
        <div class="flex-1 min-w-0">
          <div class="flex items-center justify-between">
            <h3 class="text-sm font-medium text-gray-900 truncate">{{ activity.description }}</h3>
            <span class="text-sm text-gray-500">{{ formatTime(activity.timestamp) }}</span>
          </div>
          <div class="flex items-center justify-between mt-1">
            <p class="text-sm text-gray-600">{{ activity.user }}</p>
            <span 
              :class="[
                activity.status === 'success' ? 'bg-green-100 text-green-800' : '',
                activity.status === 'warning' ? 'bg-orange-100 text-orange-800' : '',
                activity.status === 'error' ? 'bg-red-100 text-red-800' : '',
                'px-2 py-1 text-xs rounded-full font-medium'
              ]"
            >
              {{ getStatusText(activity.status) }}
            </span>
          </div>
          <div class="flex items-center justify-between mt-2">
            <span class="text-sm font-semibold text-gray-900">{{ formatPrice(activity.amount) }}</span>
            <div class="flex items-center space-x-2 space-x-reverse">
              <button class="text-xs text-blue-600 hover:text-blue-700 transition-colors">
                جزئیات
              </button>
              <button class="text-xs text-gray-500 hover:text-gray-700 transition-colors">
                کپی کد
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
    
    <!-- خلاصه آماری -->
    <div class="mt-6 pt-6 border-t border-gray-200">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div class="text-center">
          <div class="text-2xl font-bold text-green-600">{{ getActivityCount('created') }}</div>
          <div class="text-sm text-gray-600">ایجاد شده</div>
        </div>
        <div class="text-center">
          <div class="text-2xl font-bold text-blue-600">{{ getActivityCount('used') }}</div>
          <div class="text-sm text-gray-600">استفاده شده</div>
        </div>
        <div class="text-center">
          <div class="text-2xl font-bold text-orange-600">{{ getActivityCount('expired') }}</div>
          <div class="text-sm text-gray-600">منقضی شده</div>
        </div>
        <div class="text-center">
          <div class="text-2xl font-bold text-red-600">{{ getActivityCount('blocked') }}</div>
          <div class="text-sm text-gray-600">مسدود شده</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// تعریف interface برای فعالیت‌ها
interface Activity {
  id: number
  type: 'created' | 'used' | 'expired' | 'blocked'
  description: string
  amount: number
  user: string
  timestamp: Date
  status: 'success' | 'warning' | 'error'
}

const props = defineProps<{
  activities: Activity[]
}>()

// تابع فرمت کردن زمان
const formatTime = (date: Date): string => {
  const now = new Date()
  const diff = now.getTime() - date.getTime()
  const minutes = Math.floor(diff / 60000)
  const hours = Math.floor(diff / 3600000)
  const days = Math.floor(diff / 86400000)
  
  if (minutes < 60) {
    return `${minutes} دقیقه پیش`
  } else if (hours < 24) {
    return `${hours} ساعت پیش`
  } else {
    return `${days} روز پیش`
  }
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

// تابع دریافت متن وضعیت
const getStatusText = (status: string): string => {
  const statusMap = {
    success: 'موفق',
    warning: 'هشدار',
    error: 'خطا'
  }
  return statusMap[status] || status
}

// تابع شمارش فعالیت‌ها بر اساس نوع
const getActivityCount = (type: string): number => {
  return props.activities.filter(activity => activity.type === type).length
}
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 
