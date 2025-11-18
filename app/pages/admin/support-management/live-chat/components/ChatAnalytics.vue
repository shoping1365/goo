<template>
  <div class="bg-white/80 backdrop-blur-md rounded-2xl shadow-lg border border-gray-200/50 p-6">
    <div class="flex items-center justify-between mb-6">
      <h3 class="text-lg font-semibold text-gray-900 flex items-center">
        <svg class="w-5 h-5 ml-2 text-green-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
        </svg>
        آمار و گزارش‌ها
      </h3>
      <div class="flex items-center space-x-2 space-x-reverse">
        <select 
          v-model="selectedPeriod"
          class="text-sm border border-gray-300 rounded-lg px-3 py-1 focus:outline-none focus:ring-2 focus:ring-blue-500"
        >
          <option value="today">امروز</option>
          <option value="week">هفته جاری</option>
          <option value="month">ماه جاری</option>
          <option value="year">سال جاری</option>
        </select>
        <button 
          @click="exportReport"
          class="px-3 py-1 text-sm bg-blue-500 text-white rounded-lg hover:bg-blue-600 transition-colors"
        >
          خروجی
        </button>
      </div>
    </div>

    <!-- Key Metrics -->
    <div class="grid grid-cols-2 lg:grid-cols-4 gap-6 mb-6">
      <div class="bg-gradient-to-r from-blue-50 to-blue-100 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-blue-600">{{ metrics.totalChats }}</div>
            <div class="text-sm text-blue-600">کل چت‌ها</div>
          </div>
          <div class="w-10 h-10 bg-blue-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"></path>
            </svg>
          </div>
        </div>
        <div class="mt-2 text-xs text-blue-600">
          <span :class="metrics.chatGrowth >= 0 ? 'text-green-600' : 'text-red-600'">
            {{ metrics.chatGrowth >= 0 ? '+' : '' }}{{ metrics.chatGrowth }}%
          </span>
          نسبت به دوره قبل
        </div>
      </div>

      <div class="bg-gradient-to-r from-green-50 to-green-100 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-green-600">{{ metrics.avgResponseTime }}</div>
            <div class="text-sm text-green-600">میانگین پاسخ (دقیقه)</div>
          </div>
          <div class="w-10 h-10 bg-green-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
        <div class="mt-2 text-xs text-green-600">
          <span :class="metrics.responseTimeChange >= 0 ? 'text-red-600' : 'text-green-600'">
            {{ metrics.responseTimeChange >= 0 ? '+' : '' }}{{ metrics.responseTimeChange }}%
          </span>
          نسبت به دوره قبل
        </div>
      </div>

      <div class="bg-gradient-to-r from-purple-50 to-purple-100 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-purple-600">{{ metrics.satisfactionRate }}%</div>
            <div class="text-sm text-purple-600">رضایت مشتریان</div>
          </div>
          <div class="w-10 h-10 bg-purple-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"></path>
            </svg>
          </div>
        </div>
        <div class="mt-2 text-xs text-purple-600">
          <span :class="metrics.satisfactionChange >= 0 ? 'text-green-600' : 'text-red-600'">
            {{ metrics.satisfactionChange >= 0 ? '+' : '' }}{{ metrics.satisfactionChange }}%
          </span>
          نسبت به دوره قبل
        </div>
      </div>

      <div class="bg-gradient-to-r from-orange-50 to-orange-100 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-orange-600">{{ metrics.activeUsers }}</div>
            <div class="text-sm text-orange-600">کاربران فعال</div>
          </div>
          <div class="w-10 h-10 bg-orange-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197m13.5-9a2.5 2.5 0 11-5 0 2.5 2.5 0 015 0z"></path>
            </svg>
          </div>
        </div>
        <div class="mt-2 text-xs text-orange-600">
          <span :class="metrics.userGrowth >= 0 ? 'text-green-600' : 'text-red-600'">
            {{ metrics.userGrowth >= 0 ? '+' : '' }}{{ metrics.userGrowth }}%
          </span>
          نسبت به دوره قبل
        </div>
      </div>
    </div>

    <!-- Charts Section -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-6">
      <!-- Chat Volume Chart -->
      <div class="bg-gray-50 rounded-lg p-6">
        <h4 class="text-sm font-medium text-gray-900 mb-3">حجم چت‌ها</h4>
        <div class="h-32 flex items-end space-x-1 space-x-reverse">
          <div 
            v-for="(value, index) in chatVolumeData" 
            :key="index"
            class="flex-1 bg-blue-500 rounded-t"
            :style="{ height: `${(value / Math.max(...chatVolumeData)) * 100}%` }"
          ></div>
        </div>
        <div class="flex justify-between text-xs text-gray-500 mt-2">
          <span>شنبه</span>
          <span>یکشنبه</span>
          <span>دوشنبه</span>
          <span>سه‌شنبه</span>
          <span>چهارشنبه</span>
          <span>پنج‌شنبه</span>
          <span>جمعه</span>
        </div>
      </div>

      <!-- Response Time Chart -->
      <div class="bg-gray-50 rounded-lg p-6">
        <h4 class="text-sm font-medium text-gray-900 mb-3">زمان پاسخ‌دهی</h4>
        <div class="h-32 flex items-end space-x-1 space-x-reverse">
          <div 
            v-for="(value, index) in responseTimeData" 
            :key="index"
            class="flex-1 bg-green-500 rounded-t"
            :style="{ height: `${(value / Math.max(...responseTimeData)) * 100}%` }"
          ></div>
        </div>
        <div class="flex justify-between text-xs text-gray-500 mt-2">
          <span>شنبه</span>
          <span>یکشنبه</span>
          <span>دوشنبه</span>
          <span>سه‌شنبه</span>
          <span>چهارشنبه</span>
          <span>پنج‌شنبه</span>
          <span>جمعه</span>
        </div>
      </div>
    </div>

    <!-- Top Issues -->
    <div class="bg-gray-50 rounded-lg p-6">
      <h4 class="text-sm font-medium text-gray-900 mb-3">مشکلات رایج</h4>
      <div class="space-y-2">
        <div 
          v-for="issue in topIssues" 
          :key="issue.id"
          class="flex items-center justify-between p-2 bg-white rounded-lg"
        >
          <div class="flex items-center space-x-3 space-x-reverse">
            <div class="w-3 h-3 rounded-full" :style="{ backgroundColor: issue.color }"></div>
            <span class="text-sm text-gray-700">{{ issue.name }}</span>
          </div>
          <div class="flex items-center space-x-2 space-x-reverse">
            <span class="text-sm font-medium text-gray-900">{{ issue.count }}</span>
            <span class="text-xs text-gray-500">({{ issue.percentage }}%)</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Performance Summary -->
    <div class="mt-6 grid grid-cols-1 lg:grid-cols-3 gap-6">
      <div class="text-center p-6 bg-blue-50 rounded-lg">
        <div class="text-2xl font-bold text-blue-600">{{ performance.peakHours }}</div>
        <div class="text-sm text-blue-600">ساعات اوج</div>
      </div>
      <div class="text-center p-6 bg-green-50 rounded-lg">
        <div class="text-2xl font-bold text-green-600">{{ performance.avgSessionDuration }}</div>
        <div class="text-sm text-green-600">میانگین مدت جلسه (دقیقه)</div>
      </div>
      <div class="text-center p-6 bg-purple-50 rounded-lg">
        <div class="text-2xl font-bold text-purple-600">{{ performance.resolutionRate }}%</div>
        <div class="text-sm text-purple-600">نرخ حل مشکل</div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'

// Reactive data
const selectedPeriod = ref('week')

// Metrics data
const metrics = ref({
  totalChats: 1247,
  chatGrowth: 12.5,
  avgResponseTime: 2.3,
  responseTimeChange: -8.2,
  satisfactionRate: 94.7,
  satisfactionChange: 2.1,
  activeUsers: 89,
  userGrowth: 15.3
})

// Chart data
const chatVolumeData = ref([45, 52, 38, 67, 73, 58, 49])
const responseTimeData = ref([2.1, 1.8, 2.5, 1.9, 2.2, 1.7, 2.0])

// Top issues
const topIssues = ref([
  { id: 1, name: 'مشکلات فنی', count: 156, percentage: 32, color: '#EF4444' },
  { id: 2, name: 'سوالات محصول', count: 134, percentage: 27, color: '#3B82F6' },
  { id: 3, name: 'مشکلات پرداخت', count: 98, percentage: 20, color: '#10B981' },
  { id: 4, name: 'سوالات ارسال', count: 67, percentage: 14, color: '#F59E0B' },
  { id: 5, name: 'سایر', count: 35, percentage: 7, color: '#8B5CF6' }
])

// Performance data
const performance = ref({
  peakHours: '14:00-16:00',
  avgSessionDuration: 8.5,
  resolutionRate: 87.3
})

// Methods
const exportReport = () => {
  console.log('Exporting analytics report for period:', selectedPeriod.value)
  // Implementation for exporting report
}

// Watch for period changes
watch(selectedPeriod, (newPeriod) => {
  console.log('Period changed to:', newPeriod)
  // Update metrics based on selected period
  updateMetrics(newPeriod)
})

const updateMetrics = (period: string) => {
  // Simulate data update based on period
  const periodMultipliers = {
    today: 0.15,
    week: 1,
    month: 4.5,
    year: 52
  }
  
  const multiplier = periodMultipliers[period as keyof typeof periodMultipliers] || 1
  
  metrics.value = {
    totalChats: Math.round(1247 * multiplier),
    chatGrowth: 12.5,
    avgResponseTime: 2.3,
    responseTimeChange: -8.2,
    satisfactionRate: 94.7,
    satisfactionChange: 2.1,
    activeUsers: Math.round(89 * multiplier),
    userGrowth: 15.3
  }
}
</script>

<style scoped>
/* Custom styles for analytics */
</style> 
