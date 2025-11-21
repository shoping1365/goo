<template>
  <div class="order-1 bg-white rounded-2xl shadow-lg border border-gray-200 mb-6 mx-4 w-full max-w-full min-w-0 overflow-x-hidden">
    <div class="flex items-center justify-between cursor-pointer px-6 py-4 border-b border-blue-300" @click="toggleSection('monitoring')">
      <div class="flex items-center gap-6">
        <span class="text-base font-bold bg-blue-100 text-blue-700 px-4 py-1 rounded-lg">مانیتورینگ سیستم</span>
      </div>
      <!-- Refresh Button and Toggle Icon - در سمت راست هدر کانتینر -->
      <div class="flex items-center gap-2">
        <button 
          class="bg-blue-500 hover:bg-blue-600 text-white px-3 py-2 rounded-lg text-sm font-medium"
          @click="refreshMonitoring"
          @click.stop
        >
          بروزرسانی
        </button>
        <span class="text-gray-500 text-2xl">{{ sections.monitoring ? '−' : '+' }}</span>
      </div>
    </div>
    <div v-show="sections.monitoring" class="p-6 w-full min-w-0">
      <!-- System Status Overview -->
      <div class="mb-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">وضعیت کلی سیستم</h3>
        <div class="grid grid-cols-2 md:grid-cols-4 gap-6">
          <div class="bg-blue-50 rounded-lg p-6 text-center border border-blue-200">
            <div class="text-2xl font-bold text-blue-600">{{ systemStatus.totalJobs }}</div>
            <div class="text-sm text-blue-700">کل کارها</div>
          </div>
          <div class="bg-yellow-50 rounded-lg p-6 text-center border border-yellow-200">
            <div class="text-2xl font-bold text-yellow-600">{{ systemStatus.activeJobs }}</div>
            <div class="text-sm text-yellow-700">کارهای فعال</div>
          </div>
          <div class="bg-green-50 rounded-lg p-6 text-center border border-green-200">
            <div class="text-2xl font-bold text-green-600">{{ systemStatus.completedJobs }}</div>
            <div class="text-sm text-green-700">تکمیل شده</div>
          </div>
          <div class="bg-red-50 rounded-lg p-6 text-center border border-red-200">
            <div class="text-2xl font-bold text-red-600">{{ systemStatus.errorJobs }}</div>
            <div class="text-sm text-red-700">خطا</div>
          </div>
        </div>
      </div>

      <!-- Performance Metrics -->
      <div class="mb-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">معیارهای عملکرد</h3>
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div class="bg-purple-50 rounded-lg p-6 border border-purple-200">
            <div class="flex items-center justify-between">
              <div>
                <div class="text-sm text-purple-600">میانگین زمان فشرده‌سازی</div>
                <div class="text-xl font-bold text-purple-700">{{ performanceMetrics.avgCompressionTime }}</div>
              </div>
              <svg class="w-8 h-8 text-purple-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
          </div>
          <div class="bg-indigo-50 rounded-lg p-6 border border-indigo-200">
            <div class="flex items-center justify-between">
              <div>
                <div class="text-sm text-indigo-600">نرخ موفقیت</div>
                <div class="text-xl font-bold text-indigo-700">{{ performanceMetrics.successRate }}%</div>
              </div>
              <svg class="w-8 h-8 text-indigo-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
          </div>
          <div class="bg-teal-50 rounded-lg p-6 border border-teal-200">
            <div class="flex items-center justify-between">
              <div>
                <div class="text-sm text-teal-600">فضای ذخیره شده</div>
                <div class="text-xl font-bold text-teal-700">{{ performanceMetrics.savedSpace }}</div>
              </div>
              <svg class="w-8 h-8 text-teal-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 7v10c0 2.21 3.582 4 8 4s8-1.79 8-4V7M4 7c0 2.21 3.582 4 8 4s8-1.79 8-4M4 7c0-2.21 3.582-4 8-4s8 1.79 8 4"></path>
              </svg>
            </div>
          </div>
        </div>
      </div>

      <!-- Real-time Activity -->
      <div class="mb-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">فعالیت لحظه‌ای</h3>
        <div class="bg-gray-50 rounded-lg p-6 border border-gray-200">
          <div class="space-y-3 max-h-64 overflow-y-auto">
            <div v-if="recentActivities.length === 0" class="text-center py-8">
              <svg class="w-12 h-12 text-gray-300 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
              <p class="text-gray-500">هیچ فعالیتی ثبت نشده است</p>
            </div>
            <div 
              v-for="activity in recentActivities" 
              :key="activity.id"
              class="flex items-center justify-between p-3 bg-white rounded-lg border border-gray-100"
            >
              <div class="flex items-center space-x-3 space-x-reverse">
                <div :class="getActivityIconClass(activity.type)" class="w-8 h-8 rounded-full flex items-center justify-center">
                  <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path v-if="activity.type === 'start'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.828 14.828a4 4 0 01-5.656 0M9 10h1m4 0h1m-6 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                    <path v-else-if="activity.type === 'complete'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                    <path v-else-if="activity.type === 'error'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                    <path v-else stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                  </svg>
                </div>
                <div>
                  <div class="text-sm font-medium text-gray-900">{{ activity.message }}</div>
                  <div class="text-xs text-gray-500">{{ formatTime(activity.timestamp) }}</div>
                </div>
              </div>
              <div class="text-xs text-gray-400">{{ activity.duration }}</div>
            </div>
          </div>
        </div>
      </div>

      <!-- System Health -->
      <div class="mb-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">سلامت سیستم</h3>
                 <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
           <div class="bg-white rounded-lg p-6 border border-gray-200">
             <div class="flex items-center justify-between mb-3">
               <span class="text-sm font-medium text-gray-700">CPU Usage</span>
               <span class="text-sm text-gray-500">{{ systemHealth.cpu }}%</span>
             </div>
                   <div class="w-full bg-gray-200 rounded-full h-2 flex flex-row-reverse">
        <div 
          class="bg-blue-500 h-2 rounded-full transition-all duration-300"
          :style="{ width: systemHealth.cpu + '%' }"
        ></div>
      </div>
           </div>
           <div class="bg-white rounded-lg p-6 border border-gray-200">
             <div class="flex items-center justify-between mb-3">
               <span class="text-sm font-medium text-gray-700">Memory Usage</span>
               <span class="text-sm text-gray-500">{{ systemHealth.memory }}%</span>
             </div>
             <div class="w-full bg-gray-200 rounded-full h-2 flex flex-row-reverse">
               <div 
                 class="bg-green-500 h-2 rounded-full transition-all duration-300" 
                 :style="{ width: systemHealth.memory + '%' }"
               ></div>
             </div>
           </div>
           <div class="bg-white rounded-lg p-6 border border-gray-200">
             <div class="flex items-center justify-between mb-3">
               <span class="text-sm font-medium text-gray-700">Disk Usage</span>
               <span class="text-sm text-gray-500">{{ systemHealth.disk }}%</span>
             </div>
             <div class="w-full bg-gray-200 rounded-full h-2 flex flex-row-reverse">
               <div 
                 class="bg-orange-500 h-2 rounded-full transition-all duration-300" 
 
                 :style="{ width: systemHealth.disk + '%' }"
               ></div>
             </div>
           </div>
         </div>
      </div>


    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, onUnmounted } from 'vue'

interface Activity {
  id: number
  type: 'start' | 'complete' | 'error' | 'info'
  message: string
  timestamp: string
  duration: string
}

interface SystemHealth {
  cpu: number
  memory: number
  disk: number
}

interface Job {
  id?: number
  status: string
  created_at?: string
  updated_at?: string
  original_size?: number
  compressed_size?: number
  [key: string]: unknown
}

const _props = defineProps({
  compressionEnabled: {
    type: Boolean,
    required: true
  }
})

// Reactive data
const systemStatus = ref({
  totalJobs: 0,
  activeJobs: 0,
  completedJobs: 0,
  errorJobs: 0
})

const performanceMetrics = ref({
  avgCompressionTime: '0 دقیقه',
  successRate: 0,
  savedSpace: '0 MB'
})

const recentActivities = ref<Activity[]>([])

const systemHealth = ref<SystemHealth>({
  cpu: 0,
  memory: 0,
  disk: 0
})

// Compression schedule data
const compressionSchedule = ref({
  startHour: 1,
  endHour: 13
})

const currentTime = ref('')
const scheduleStatus = ref('inactive')

// Collapsible sections state
const defaultSections = {
  monitoring: true
} as const

const sections = reactive<Record<keyof typeof defaultSections, boolean>>({ ...defaultSections })

function toggleSection(section: keyof typeof defaultSections) {
  sections[section] = !sections[section]
}

// Methods
const refreshMonitoring = async () => {
  await Promise.all([
    fetchSystemStatus(),
    fetchPerformanceMetrics(),
    fetchRecentActivities(),
    fetchSystemHealth()
  ])
}

const fetchSystemStatus = async () => {
  try {
    interface ApiResponse {
      success?: boolean
      data?: Array<{ status?: string }>
    }
    const response = await $fetch<ApiResponse>('/api/media/compression-jobs?job_type=video_compression', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    })
    
    if (response.success) {
      const jobs = (response.data || []) as Job[]
      systemStatus.value = {
        totalJobs: jobs.length,
        activeJobs: jobs.filter((job: Job) => job.status === 'processing').length,
        completedJobs: jobs.filter((job: Job) => job.status === 'completed').length,
        errorJobs: jobs.filter((job: Job) => job.status === 'error').length
      }
    }
  } catch (error) {
    console.error('خطا در دریافت وضعیت سیستم:', error)
  }
}

const fetchPerformanceMetrics = async () => {
  try {
    interface ApiResponse {
      success?: boolean
      data?: Array<{ status?: string; created_at?: string; updated_at?: string; original_size?: number; compressed_size?: number }>
    }
    const response = await $fetch<ApiResponse>('/api/media/compression-jobs?job_type=video_compression', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    })
    
    if (response.success) {
      const jobs = (response.data || []) as Job[]
      const completedJobs = jobs.filter((job: Job) => job.status === 'completed')
      const totalJobs = jobs.length
      
      // محاسبه میانگین زمان فشرده‌سازی
      let totalTime = 0
      let avgTime = 0
      
      if (completedJobs.length > 0) {
        completedJobs.forEach((job: Job) => {
          if (job.created_at && job.updated_at) {
            const startTime = new Date(job.created_at).getTime()
            const endTime = new Date(job.updated_at).getTime()
            const duration = (endTime - startTime) / 1000 / 60 // به دقیقه
            totalTime += duration
          }
        })
        avgTime = Math.round(totalTime / completedJobs.length)
      }
      
      // محاسبه فضای ذخیره شده
      let totalSavedSpace = 0
      completedJobs.forEach((job: Job) => {
        if (job.original_size && job.compressed_size) {
          totalSavedSpace += (job.original_size - job.compressed_size)
        }
      })
      
      performanceMetrics.value = {
        avgCompressionTime: avgTime > 0 ? `${avgTime} دقیقه` : '0 دقیقه',
        successRate: totalJobs > 0 ? Math.round((completedJobs.length / totalJobs) * 100) : 0,
        savedSpace: totalSavedSpace > 0 ? formatFileSize(totalSavedSpace) : '0 B'
      }
    }
  } catch (error) {
    console.error('خطا در دریافت معیارهای عملکرد:', error)
  }
}

const fetchRecentActivities = async () => {
  try {
    interface ApiResponse {
      success?: boolean
      data?: Array<{ id?: string; status?: string; media?: { filename?: string }; updated_at?: string; created_at?: string }>
    }
    const response = await $fetch<ApiResponse>('/api/media/compression-jobs?job_type=video_compression', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    })
    
    if (response.success) {
      const _jobs = response.data || []
      const activities: Activity[] = []
      
      // تبدیل کارها به فعالیت‌ها
      const typedJobs = (response.data || []) as Job[]
      typedJobs.slice(0, 10).forEach((job: Job) => {
        const activity: Activity = {
          id: job.id,
          type: job.status === 'completed' ? 'complete' : 
                job.status === 'error' ? 'error' : 
                job.status === 'processing' ? 'start' : 'info',
          message: job.status === 'completed' ? `فشرده‌سازی ${job.media?.filename || 'فایل'} تکمیل شد` :
                   job.status === 'error' ? `خطا در فشرده‌سازی ${job.media?.filename || 'فایل'}` :
                   job.status === 'processing' ? `شروع فشرده‌سازی ${job.media?.filename || 'فایل'}` :
                   `در انتظار فشرده‌سازی ${job.media?.filename || 'فایل'}`,
          timestamp: job.updated_at || job.created_at,
          duration: getTimeAgo(job.updated_at || job.created_at)
        }
        activities.push(activity)
      })
      
      recentActivities.value = activities
    }
  } catch (error) {
    console.error('خطا در دریافت فعالیت‌های اخیر:', error)
  }
}

const fetchSystemHealth = async () => {
  try {
    interface ApiResponse {
      success?: boolean
      data?: { cpu?: number; memory?: number; disk?: number }
    }
    const response = await $fetch<ApiResponse>('/api/media/system-health', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    })
    
    if (response.success && response.data) {
      systemHealth.value = {
        cpu: Math.round(response.data.cpu || 0),
        memory: Math.round(response.data.memory || 0),
        disk: Math.round(response.data.disk || 0)
      }
    }
  } catch (error) {
    console.error('خطا در دریافت سلامت سیستم:', error)
    // در صورت خطا، مقادیر پیش‌فرض نمایش داده می‌شود
    systemHealth.value = {
      cpu: 0,
      memory: 0,
      disk: 0
    }
  }
}

const getActivityIconClass = (type: string) => {
  const classes = {
    start: 'bg-blue-500',
    complete: 'bg-green-500',
    error: 'bg-red-500',
    info: 'bg-gray-500'
  }
  return classes[type as keyof typeof classes] || 'bg-gray-500'
}

const formatTime = (timestamp: string) => {
  const date = new Date(timestamp)
  return date.toLocaleTimeString('fa-IR', {
    hour: '2-digit',
    minute: '2-digit'
  })
}

const getTimeAgo = (timestamp: string) => {
  const now = new Date().getTime()
  const time = new Date(timestamp).getTime()
  const diff = now - time
  
  const minutes = Math.floor(diff / (1000 * 60))
  const hours = Math.floor(diff / (1000 * 60 * 60))
  const days = Math.floor(diff / (1000 * 60 * 60 * 24))
  
  if (days > 0) return `${days} روز پیش`
  if (hours > 0) return `${hours} ساعت پیش`
  if (minutes > 0) return `${minutes} دقیقه پیش`
  return 'لحظاتی پیش'
}

const formatFileSize = (bytes: number): string => {
  if (bytes === 0) return '0 B'
  const k = 1024
  const sizes = ['B', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

// Compression schedule methods
const _updateCompressionSchedule = async () => {
  try {
    interface ApiResponse {
      success?: boolean
    }
    const payload = [
      { key: 'video_compression.start_hour', value: compressionSchedule.value.startHour.toString(), category: 'video_compression', type: 'string' },
      { key: 'video_compression.end_hour', value: compressionSchedule.value.endHour.toString(), category: 'video_compression', type: 'string' }
    ]
    
    const response = await $fetch<ApiResponse>('/api/admin/settings', {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: payload
    })
    
    if (response.success) {
    }
  } catch (error) {
    console.error('خطا در ذخیره تنظیمات زمان‌بندی:', error)
  }
}



const updateCurrentTime = () => {
  const now = new Date()
  currentTime.value = now.toLocaleTimeString('fa-IR', {
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit'
  })
  
  // بررسی وضعیت زمان‌بندی
  const currentHour = now.getHours()
  const startHour = compressionSchedule.value.startHour
  const endHour = compressionSchedule.value.endHour
  
  let inTimeWindow = false
  if (startHour <= endHour) {
    // بازه عادی (مثل 1 تا 13)
    inTimeWindow = currentHour >= startHour && currentHour <= endHour
  } else {
    // بازه شبانه‌روزی (مثل 22 تا 6)
    inTimeWindow = currentHour >= startHour || currentHour <= endHour
  }
  
  scheduleStatus.value = inTimeWindow ? 'active' : 'inactive'
}

const _getScheduleStatusText = () => {
  return scheduleStatus.value === 'active' 
    ? 'فشرده‌سازی فعال است' 
    : 'فشرده‌سازی غیرفعال است'
}

const _getScheduleStatusClass = () => {
  return scheduleStatus.value === 'active' 
    ? 'bg-green-500' 
    : 'bg-red-500'
}



// Auto-refresh functionality
let refreshInterval: NodeJS.Timeout | null = null

onMounted(() => {
  refreshMonitoring()
  updateCurrentTime()
  
  // بروزرسانی هر 30 ثانیه
  refreshInterval = setInterval(refreshMonitoring, 30000)
  
  // بروزرسانی زمان فعلی هر ثانیه
  const timeInterval = setInterval(updateCurrentTime, 1000)
  
  onUnmounted(() => {
    if (timeInterval) {
      clearInterval(timeInterval)
    }
  })
})

onUnmounted(() => {
  if (refreshInterval) {
    clearInterval(refreshInterval)
  }
})

// Expose methods to parent
defineExpose({
  refreshMonitoring
})
</script>

<!--
  کامپوننت مانیتورینگ سیستم فشرده‌سازی ویدیو
  - نمایش وضعیت کلی سیستم
  - معیارهای عملکرد
  - فعالیت‌های لحظه‌ای
  - سلامت سیستم
  - بروزرسانی خودکار
  - توضیحات کامل به فارسی در کد
--> 
