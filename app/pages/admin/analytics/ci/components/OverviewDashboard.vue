<template>
  <div class="mb-8">
    <!-- Section Header -->
    <div class="flex items-center justify-between mb-6">
      <h2 class="text-xl font-semibold text-gray-900">نمای کلی بازار و رقبا</h2>
      <div class="flex items-center space-x-2 space-x-reverse">
        <span class="text-sm text-gray-500">آخرین به‌روزرسانی:</span>
        <span class="text-sm font-medium text-gray-900">{{ lastUpdateTime }}</span>
      </div>
    </div>

    <!-- KPIs Grid -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
      <!-- Market Share KPI -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-gray-600 mb-1">سهم بازار ما</p>
            <p class="text-2xl font-bold text-gray-900">{{ marketShare }}%</p>
            <div class="flex items-center mt-2">
              <span :class="marketShareChange >= 0 ? 'text-green-600' : 'text-red-600'" class="text-sm font-medium">
                {{ marketShareChange >= 0 ? '+' : '' }}{{ marketShareChange }}%
              </span>
              <span class="text-sm text-gray-500 mr-2">نسبت به ماه قبل</span>
            </div>
          </div>
          <div class="w-12 h-12 bg-gradient-to-br from-blue-500 to-blue-600 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
          </div>
        </div>
      </div>

      <!-- Competitor Count KPI -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-gray-600 mb-1">رقبای فعال</p>
            <p class="text-2xl font-bold text-gray-900">{{ activeCompetitors }}</p>
            <div class="flex items-center mt-2">
              <span class="text-orange-600 text-sm font-medium">+{{ newCompetitors }}</span>
              <span class="text-sm text-gray-500 mr-2">رقبای جدید</span>
            </div>
          </div>
          <div class="w-12 h-12 bg-gradient-to-br from-orange-500 to-orange-600 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"></path>
            </svg>
          </div>
        </div>
      </div>

      <!-- Market Growth KPI -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-gray-600 mb-1">رشد بازار</p>
            <p class="text-2xl font-bold text-gray-900">{{ marketGrowth }}%</p>
            <div class="flex items-center mt-2">
              <span :class="marketGrowthChange >= 0 ? 'text-green-600' : 'text-red-600'" class="text-sm font-medium">
                {{ marketGrowthChange >= 0 ? '+' : '' }}{{ marketGrowthChange }}%
              </span>
              <span class="text-sm text-gray-500 mr-2">نسبت به ماه قبل</span>
            </div>
          </div>
          <div class="w-12 h-12 bg-gradient-to-br from-green-500 to-green-600 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
            </svg>
          </div>
        </div>
      </div>

      <!-- Threat Level KPI -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-gray-600 mb-1">سطح تهدید</p>
            <p class="text-2xl font-bold" :class="getThreatLevelClass()">{{ threatLevel }}</p>
            <div class="flex items-center mt-2">
              <span class="text-sm text-gray-500">{{ threatDescription }}</span>
            </div>
          </div>
          <div class="w-12 h-12 rounded-lg flex items-center justify-center" :class="getThreatLevelBgClass()">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- Market Status & Trends -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-8">
      <!-- Market Status -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">وضعیت بازار</h3>
        <div class="space-y-4">
          <div v-for="status in marketStatuses" :key="status.id" class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <div class="flex items-center">
              <span class="w-3 h-3 rounded-full mr-3" :class="getStatusColor(status.level)"></span>
              <span class="text-sm font-medium text-gray-900">{{ status.title }}</span>
            </div>
            <span class="text-sm text-gray-600">{{ status.value }}</span>
          </div>
        </div>
      </div>

      <!-- Key Trends -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">روندهای کلیدی</h3>
        <div class="space-y-4">
          <div v-for="trend in keyTrends" :key="trend.id" class="flex items-start">
            <div class="w-2 h-2 rounded-full mt-2 mr-3" :class="trend.direction === 'up' ? 'bg-green-500' : 'bg-red-500'"></div>
            <div class="flex-1">
              <p class="text-sm font-medium text-gray-900">{{ trend.title }}</p>
              <p class="text-xs text-gray-600 mt-1">{{ trend.description }}</p>
            </div>
            <span class="text-xs text-gray-500">{{ trend.time }}</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Alerts & Important Changes -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-4">
        <h3 class="text-lg font-semibold text-gray-900">هشدارها و تغییرات مهم</h3>
        <button class="text-sm text-blue-600 hover:text-blue-700 font-medium" @click="showAllAlerts = true">
          مشاهده همه
        </button>
      </div>
      
      <div class="space-y-4">
        <div v-for="alert in recentAlerts" :key="alert.id" class="flex items-start p-6 border border-gray-200 rounded-lg" :class="getAlertBorderClass(alert.priority)">
          <div class="w-2 h-2 rounded-full mt-2 mr-3" :class="getAlertColor(alert.priority)"></div>
          <div class="flex-1">
            <div class="flex items-center justify-between">
              <h4 class="text-sm font-medium text-gray-900">{{ alert.title }}</h4>
              <span class="text-xs text-gray-500">{{ formatTime(alert.timestamp) }}</span>
            </div>
            <p class="text-sm text-gray-600 mt-1">{{ alert.description }}</p>
            <div class="flex items-center mt-2 space-x-2 space-x-reverse">
              <span class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium" :class="getPriorityBadgeClass(alert.priority)">
                {{ getPriorityText(alert.priority) }}
              </span>
              <span class="text-xs text-gray-500">{{ alert.competitor }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- All Alerts Modal -->
    <div v-if="showAllAlerts" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg p-6 w-full max-w-4xl mx-4 max-h-[80vh] overflow-y-auto">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">همه هشدارها و تغییرات</h3>
          <button class="text-gray-400 hover:text-gray-600" @click="showAllAlerts = false">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>
        
        <div class="space-y-4">
          <div v-for="alert in allAlerts" :key="alert.id" class="flex items-start p-6 border border-gray-200 rounded-lg" :class="getAlertBorderClass(alert.priority)">
            <div class="w-2 h-2 rounded-full mt-2 mr-3" :class="getAlertColor(alert.priority)"></div>
            <div class="flex-1">
              <div class="flex items-center justify-between">
                <h4 class="text-sm font-medium text-gray-900">{{ alert.title }}</h4>
                <span class="text-xs text-gray-500">{{ formatTime(alert.timestamp) }}</span>
              </div>
              <p class="text-sm text-gray-600 mt-1">{{ alert.description }}</p>
              <div class="flex items-center mt-2 space-x-2 space-x-reverse">
                <span class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium" :class="getPriorityBadgeClass(alert.priority)">
                  {{ getPriorityText(alert.priority) }}
                </span>
                <span class="text-xs text-gray-500">{{ alert.competitor }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'

// Props
interface Props {
  refreshTrigger?: number
}

const props = withDefaults(defineProps<Props>(), {
  refreshTrigger: 0
})

// Reactive data
const showAllAlerts = ref(false)
const lastUpdateTime = ref('')

// Sample data - در پروژه واقعی از API دریافت می‌شود
const marketShare = ref(23.5)
const marketShareChange = ref(2.1)
const activeCompetitors = ref(8)
const newCompetitors = ref(1)
const marketGrowth = ref(12.8)
const marketGrowthChange = ref(-0.5)
const threatLevel = ref('متوسط')
const threatDescription = ref('نیاز به توجه بیشتر')

const marketStatuses = ref([
  { id: 1, title: 'وضعیت کلی بازار', level: 'stable', value: 'پایدار' },
  { id: 2, title: 'رقابت', level: 'high', value: 'شدید' },
  { id: 3, title: 'نوآوری', level: 'medium', value: 'متوسط' },
  { id: 4, title: 'ثبات قیمت', level: 'stable', value: 'پایدار' }
])

const keyTrends = ref([
  { id: 1, title: 'افزایش استفاده از هوش مصنوعی', description: 'رقبا در حال سرمایه‌گذاری روی AI هستند', direction: 'up', time: '2 ساعت پیش' },
  { id: 2, title: 'کاهش قیمت محصولات', description: 'رقبای اصلی قیمت‌ها را کاهش داده‌اند', direction: 'down', time: '1 روز پیش' },
  { id: 3, title: 'ورود به بازارهای جدید', description: 'رقبا در حال گسترش جغرافیایی هستند', direction: 'up', time: '3 روز پیش' }
])

const recentAlerts = ref([
  { id: 1, title: 'محصول جدید رقیب', description: 'شرکت ABC محصول جدیدی در حوزه IoT معرفی کرد', priority: 'high', competitor: 'شرکت ABC', timestamp: new Date() },
  { id: 2, title: 'کمپین تبلیغاتی', description: 'رقبای اصلی کمپین تبلیغاتی جدیدی راه‌اندازی کرده‌اند', priority: 'medium', competitor: 'رقبای اصلی', timestamp: new Date(Date.now() - 3600000) },
  { id: 3, title: 'تغییر قیمت', description: 'شرکت XYZ قیمت محصولات خود را 15% کاهش داد', priority: 'high', competitor: 'شرکت XYZ', timestamp: new Date(Date.now() - 7200000) }
])

const allAlerts = ref([
  ...recentAlerts.value,
  { id: 4, title: 'شراکت جدید', description: 'رقبای اصلی با شرکت‌های فناوری شراکت کرده‌اند', priority: 'low', competitor: 'رقبای اصلی', timestamp: new Date(Date.now() - 86400000) },
  { id: 5, title: 'تغییر استراتژی', description: 'شرکت DEF استراتژی بازاریابی خود را تغییر داده است', priority: 'medium', competitor: 'شرکت DEF', timestamp: new Date(Date.now() - 172800000) }
])

// Computed
const getThreatLevelClass = () => {
  switch (threatLevel.value) {
    case 'کم': return 'text-green-600'
    case 'متوسط': return 'text-yellow-600'
    case 'زیاد': return 'text-red-600'
    default: return 'text-gray-600'
  }
}

const getThreatLevelBgClass = () => {
  switch (threatLevel.value) {
    case 'کم': return 'bg-gradient-to-br from-green-500 to-green-600'
    case 'متوسط': return 'bg-gradient-to-br from-yellow-500 to-yellow-600'
    case 'زیاد': return 'bg-gradient-to-br from-red-500 to-red-600'
    default: return 'bg-gradient-to-br from-gray-500 to-gray-600'
  }
}

// Methods
const getStatusColor = (level: string) => {
  switch (level) {
    case 'stable': return 'bg-green-500'
    case 'high': return 'bg-red-500'
    case 'medium': return 'bg-yellow-500'
    default: return 'bg-gray-500'
  }
}

const getAlertColor = (priority: string) => {
  switch (priority) {
    case 'high': return 'bg-red-500'
    case 'medium': return 'bg-yellow-500'
    case 'low': return 'bg-blue-500'
    default: return 'bg-gray-500'
  }
}

const getAlertBorderClass = (priority: string) => {
  switch (priority) {
    case 'high': return 'border-red-200 bg-red-50'
    case 'medium': return 'border-yellow-200 bg-yellow-50'
    case 'low': return 'border-blue-200 bg-blue-50'
    default: return 'border-gray-200 bg-gray-50'
  }
}

const getPriorityBadgeClass = (priority: string) => {
  switch (priority) {
    case 'high': return 'bg-red-100 text-red-800'
    case 'medium': return 'bg-yellow-100 text-yellow-800'
    case 'low': return 'bg-blue-100 text-blue-800'
    default: return 'bg-gray-100 text-gray-800'
  }
}

const getPriorityText = (priority: string) => {
  switch (priority) {
    case 'high': return 'زیاد'
    case 'medium': return 'متوسط'
    case 'low': return 'کم'
    default: return 'نامشخص'
  }
}

const formatTime = (date: Date) => {
  const now = new Date()
  const diff = now.getTime() - date.getTime()
  const minutes = Math.floor(diff / 60000)
  
  if (minutes < 1) return 'همین الان'
  if (minutes < 60) return `${minutes} دقیقه پیش`
  if (minutes < 1440) return `${Math.floor(minutes / 60)} ساعت پیش`
  return `${Math.floor(minutes / 1440)} روز پیش`
}

const updateLastUpdateTime = () => {
  lastUpdateTime.value = new Date().toLocaleTimeString('fa-IR', { 
    hour: '2-digit', 
    minute: '2-digit',
    hour12: false
  })
}

// Lifecycle
onMounted(() => {
  updateLastUpdateTime()
})

// Watch for refresh trigger
watch(() => props.refreshTrigger, () => {
  updateLastUpdateTime()
})
</script>

<style scoped>
/* Custom styles */
</style>
