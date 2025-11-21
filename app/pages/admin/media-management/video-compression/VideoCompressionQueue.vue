<template>
  <div class="order-4 bg-white rounded-2xl shadow-lg border border-gray-200 mb-6 mx-4 w-full max-w-full min-w-0 overflow-x-hidden">
    <div class="flex items-center justify-between cursor-pointer px-6 py-4 border-b border-green-300" @click="toggleSection('compressionQueue')">
      <div class="flex items-center gap-4">
        <span class="text-base font-bold bg-green-100 text-green-700 px-4 py-1 rounded-lg">صف فشرده‌سازی</span>
      </div>
      <!-- Filters, Refresh Button and Toggle Icon - در سمت راست هدر کانتینر -->
      <div class="flex items-center gap-2">
        <input
          v-model="searchQuery"
          type="text"
          placeholder="جستجو در صف..."
          class="px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500"
          @click.stop
        >
        <select
          v-model="filterStatus"
          class="px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500"
          @click.stop
        >
          <option value="">همه وضعیت‌ها</option>
          <option value="pending">در انتظار</option>
          <option value="processing">در حال پردازش</option>
          <option value="completed">تکمیل شده</option>
          <option value="error">خطا</option>
          <option value="paused">متوقف شده</option>
        </select>
        <button 
          v-if="queueStatus.inProgress === 0 && queueStatus.total > 0" 
          class="bg-green-500 hover:bg-green-600 text-white px-3 py-2 rounded-lg text-sm font-medium"
          @click="processPendingJobs"
          @click.stop
        >
          شروع پردازش
        </button>
        <button 
          class="bg-blue-500 hover:bg-blue-600 text-white px-3 py-2 rounded-lg text-sm font-medium"
          @click="fetchCompressionJobs"
          @click.stop
        >
          بروزرسانی
        </button>
        <span class="text-gray-500 text-2xl">{{ sections.compressionQueue ? '−' : '+' }}</span>
      </div>
    </div>
    <div v-show="sections.compressionQueue" class="p-6 w-full min-w-0">
      <!-- Queue Status -->
      <div class="mb-6">
        <div class="grid grid-cols-2 md:grid-cols-5 gap-4">
          <div class="bg-blue-50 rounded-lg p-4 text-center">
            <div class="text-2xl font-bold text-blue-600">{{ queueStatus.total }}</div>
            <div class="text-sm text-blue-700">کل</div>
          </div>
          <div class="bg-yellow-50 rounded-lg p-4 text-center">
            <div class="text-2xl font-bold text-yellow-600">{{ queueStatus.inProgress }}</div>
            <div class="text-sm text-yellow-700">در حال پردازش</div>
          </div>
          <div class="bg-green-50 rounded-lg p-4 text-center">
            <div class="text-2xl font-bold text-green-600">{{ queueStatus.completed }}</div>
            <div class="text-sm text-green-700">تکمیل شده</div>
          </div>
          <div class="bg-red-50 rounded-lg p-4 text-center">
            <div class="text-2xl font-bold text-red-600">{{ queueStatus.errors }}</div>
            <div class="text-sm text-red-700">خطا</div>
          </div>
          <div class="bg-purple-50 rounded-lg p-4 text-center">
            <div class="text-2xl font-bold text-purple-600">{{ queueStatus.progress }}%</div>
            <div class="text-sm text-purple-700">پیشرفت</div>
          </div>
        </div>
      </div>

      <!-- Queue Controls -->
      <div class="flex flex-wrap gap-3 mb-6">
        <button
          v-if="selectedQueue.length > 0"
          class="bg-red-500 hover:bg-red-600 text-white px-4 py-2 rounded-lg text-sm font-medium"
          @click="cancelSelectedJobs"
        >
          لغو انتخاب شده‌ها
        </button>
        <button
          v-if="selectedQueue.length > 0"
          class="bg-yellow-500 hover:bg-yellow-600 text-white px-4 py-2 rounded-lg text-sm font-medium"
          @click="pauseSelectedJobs"
        >
          توقف انتخاب شده‌ها
        </button>
        <button
          v-if="hasManualCompressibleJobs"
          class="bg-green-500 hover:bg-green-600 text-white px-4 py-2 rounded-lg text-sm font-medium"
          @click="manualCompressSelectedJobs"
        >
          فشرده‌سازی دستی
        </button>
        <button 
          v-if="queueStatus.completed > 0" 
          class="bg-gray-500 hover:bg-gray-600 text-white px-4 py-2 rounded-lg text-sm font-medium"
          @click="clearCompleted"
        >
          پاک کردن تکمیل شده‌ها
        </button>
      </div>

      <!-- Queue Table -->
      <div v-if="filteredQueueItems.length === 0" class="text-center py-8">
        <svg class="w-16 h-16 text-gray-300 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
        </svg>
        <h3 class="text-lg font-medium text-gray-900 mb-2">{{ compressionJobs.length === 0 ? 'صف فشرده‌سازی خالی است' : 'نتیجه‌ای یافت نشد' }}</h3>
        <p class="text-gray-500">{{ compressionJobs.length === 0 ? 'ویدیوهایی را انتخاب کنید تا در صف قرار گیرند' : 'با فیلترهای انتخاب شده نتیجه‌ای یافت نشد' }}</p>
      </div>
      <div v-else class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                <input
                  v-model="selectAllQueue"
                  type="checkbox"
                  class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                  @change="toggleSelectAllQueue"
                >
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">ویدیو</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">حجم اصلی</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">پیشرفت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ ایجاد</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="job in filteredQueueItems" :key="job.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <input
                  v-model="selectedQueue"
                  :value="job.id"
                  type="checkbox"
                  class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                >
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="w-10 h-8 bg-gray-100 rounded flex items-center justify-center mr-3">
                    <svg class="w-5 h-5 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 002 2z"></path>
                    </svg>
                  </div>
                  <div>
                    <div class="text-sm font-medium text-gray-900">{{ job.media?.filename || 'فایل ناشناس' }}</div>
                    <div class="text-sm text-gray-500">Job ID: {{ job.id }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ formatFileSize(job.original_size || 0) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getStatusBadgeClass(job.status)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                  {{ getStatusText(job.status) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div v-if="job.status === 'processing' && job.progress" class="w-32">
                  <div class="w-full bg-gray-200 rounded-full h-2">
                    <div 
                      class="bg-blue-500 h-2 rounded-full transition-all duration-300" 
                      :style="{ width: job.progress + '%' }"
                    ></div>
                  </div>
                  <div class="text-xs text-gray-500 mt-1">{{ job.progress }}%</div>
                </div>
                <span v-else class="text-gray-400">-</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ formatDate(job.created_at) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <div class="flex space-x-2 space-x-reverse">
                  <button 
                    v-if="job.status === 'processing'"
                    class="text-yellow-600 hover:text-yellow-800"
                    title="توقف"
                    @click="pauseJob(job.id)"
                  >
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 9v6m4-6v6m7-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                    </svg>
                  </button>
                  <button 
                    v-if="job.status === 'processing' || job.status === 'pending'"
                    class="text-red-600 hover:text-red-800"
                    title="لغو"
                    @click="cancelJob(job.id)"
                  >
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                    </svg>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted, onUnmounted } from 'vue'

interface ApiResponse<T> {
  data?: T
  success?: boolean
  message?: string
  result?: string
  logs?: string
}

interface CompressionJob {
  id: number
  media_id: number
  job_type: string
  status: string
  progress: number | null
  original_size: number | null
  compressed_size: number | null
  created_at: string
  updated_at: string
  log: string | null
  media?: {
    id: number
    filename: string
    file_path: string
  }
}

// Reactive data
const compressionJobs = ref<CompressionJob[]>([])
const searchQuery = ref('')
const filterStatus = ref('')
const selectedQueue = ref<number[]>([])

// Collapsible sections state
const defaultSections = {
  compressionQueue: true
} as const

const sections = reactive<Record<keyof typeof defaultSections, boolean>>({ ...defaultSections })

function toggleSection(section: keyof typeof sections) {
  sections[section] = !sections[section]
}

// Computed properties for queue status
const queueStatus = computed(() => {
  const total = compressionJobs.value.length
  const inProgress = compressionJobs.value.filter(job => job.status === 'processing').length
  const completed = compressionJobs.value.filter(job => job.status === 'completed').length
  const errors = compressionJobs.value.filter(job => job.status === 'error').length
  const progress = total > 0 ? Math.round((completed / total) * 100) : 0

  return {
    total,
    inProgress,
    completed,
    errors,
    progress
  }
})

// Computed property for filtered queue items
const filteredQueueItems = computed(() => {
  let filtered = compressionJobs.value

  // Filter by status
  if (filterStatus.value) {
    filtered = filtered.filter(job => job.status === filterStatus.value)
  }

  // Filter by search query
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    filtered = filtered.filter(job => 
      job.media?.filename?.toLowerCase().includes(query) ||
      job.id.toString().includes(query)
    )
  }

  return filtered
})

// Computed property for select all checkbox
const selectAllQueue = computed({
  get: () => {
    return filteredQueueItems.value.length > 0 && selectedQueue.value.length === filteredQueueItems.value.length
  },
  set: (value: boolean) => {
    if (value) {
      selectedQueue.value = filteredQueueItems.value.map(job => job.id)
    } else {
      selectedQueue.value = []
    }
  }
})

// Computed property to check if selected jobs are manually compressible
const hasManualCompressibleJobs = computed(() => {
  if (selectedQueue.value.length === 0) return false
  
  const selectedJobs = filteredQueueItems.value.filter(job => selectedQueue.value.includes(job.id))
  return selectedJobs.some(job => job.status === 'error' || job.status === 'pending')
})

// Methods
const pauseJob = async (jobId: number): Promise<void> => {
  try {
    const response = await $fetch(`/api/media/compression-jobs/${jobId}/pause`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      }
    }) as ApiResponse<unknown>
    
    if (response.success) {
      await fetchCompressionJobs()
    }
  } catch (error) {
    console.error('خطا در توقف کار:', error)
  }
}

const cancelJob = async (jobId: number): Promise<void> => {
  try {
    const response = await $fetch(`/api/media/compression-jobs/${jobId}/cancel`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      }
    }) as ApiResponse<unknown>
    
    if (response.success) {
      await fetchCompressionJobs()
    }
  } catch (error) {
    console.error('خطا در لغو کار:', error)
  }
}

const toggleSelectAllQueue = () => {
  if (selectAllQueue.value) {
    selectedQueue.value = []
  } else {
    selectedQueue.value = filteredQueueItems.value.map(job => job.id)
  }
}

const cancelSelectedJobs = async () => {
  if (selectedQueue.value.length === 0) return
  
  if (confirm(`آیا از لغو ${selectedQueue.value.length} کار انتخاب شده اطمینان دارید؟`)) {
    try {
      for (const jobId of selectedQueue.value) {
        await cancelJob(jobId)
      }
      selectedQueue.value = []
    } catch (error) {
      console.error('خطا در لغو کارهای انتخاب شده:', error)
    }
  }
}

const pauseSelectedJobs = async () => {
  if (selectedQueue.value.length === 0) return
  
  if (confirm(`آیا از توقف ${selectedQueue.value.length} کار انتخاب شده اطمینان دارید؟`)) {
    try {
      for (const jobId of selectedQueue.value) {
        await pauseJob(jobId)
      }
      selectedQueue.value = []
    } catch (error) {
      console.error('خطا در توقف کارهای انتخاب شده:', error)
    }
  }
}

const manualCompressSelectedJobs = async () => {
  if (selectedQueue.value.length === 0) return
  
  const selectedJobs = filteredQueueItems.value.filter(job => selectedQueue.value.includes(job.id))
  const compressibleJobs = selectedJobs.filter(job => job.status === 'error' || job.status === 'pending')
  
  if (compressibleJobs.length === 0) {
    alert('هیچ فایل قابل فشرده‌سازی انتخاب نشده است')
    return
  }
  
  if (confirm(`آیا از فشرده‌سازی دستی ${compressibleJobs.length} فایل انتخاب شده اطمینان دارید؟`)) {
    try {
      // تبدیل job های error به pending
      for (const job of compressibleJobs) {
        if (job.status === 'error') {
          await $fetch(`/api/media/compression-jobs/${job.id}/reset`, {
            method: 'POST'
          })
        }
      }
      
      // شروع پردازش
      await processPendingJobs()
      
      selectedQueue.value = []
      alert('فشرده‌سازی دستی شروع شد')
    } catch (error) {
      console.error('خطا در فشرده‌سازی دستی:', error)
      alert('خطا در شروع فشرده‌سازی دستی')
    }
  }
}

const clearCompleted = () => {
  compressionJobs.value = compressionJobs.value.filter(job => job.status !== 'completed')
}

const processPendingJobs = async (): Promise<void> => {
  try {
    const response = await $fetch('/api/media/process-pending-video-compression', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      }
    }) as ApiResponse<unknown>
    
    if (response.success) {
      // بروزرسانی لیست بعد از 2 ثانیه
      setTimeout(() => {
        fetchCompressionJobs()
      }, 2000)
    }
  } catch (error) {
    console.error('خطا در شروع پردازش:', error)
  }
}

const fetchCompressionJobs = async () => {
  try {
    const response: ApiResponse<CompressionJob[]> = await $fetch('/api/media/compression-jobs?job_type=video_compression', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    })
    
    if (response.success) {
      compressionJobs.value = response.data || []
    }
  } catch (error) {
    console.error('خطا در دریافت صف فشرده‌سازی:', error)
  }
}

const formatFileSize = (bytes: number): string => {
  if (bytes === 0) return '0 B'
  const k = 1024
  const sizes = ['B', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

const getStatusText = (status: string) => {
  const statusMap: Record<string, string> = {
    'pending': 'در انتظار',
    'processing': 'در حال پردازش',
    'completed': 'تکمیل شده',
    'error': 'خطا'
  }
  return statusMap[status] || status
}

const getStatusBadgeClass = (status: string) => {
  const classMap: Record<string, string> = {
    'pending': 'bg-yellow-100 text-yellow-800',
    'processing': 'bg-purple-100 text-purple-800',
    'completed': 'bg-green-100 text-green-800',
    'error': 'bg-red-100 text-red-800'
  }
  return classMap[status] || 'bg-gray-100 text-gray-800'
}

const formatDate = (dateString: string) => {
  if (!dateString) return ''
  const date = new Date(dateString)
  return date.toLocaleString('fa-IR')
}

// Auto-refresh functionality
let refreshInterval: NodeJS.Timeout | null = null

onMounted(() => {
  fetchCompressionJobs()
  // بروزرسانی هر 10 ثانیه
  refreshInterval = setInterval(fetchCompressionJobs, 10000)
})

onUnmounted(() => {
  if (refreshInterval) {
    clearInterval(refreshInterval)
  }
})

// Expose methods to parent
defineExpose({
  compressionJobs,
  fetchCompressionJobs
})
</script>

<!--
  کامپوننت صف فشرده‌سازی ویدیو
  - نمایش وضعیت کلی صف فشرده‌سازی
  - کنترل‌های توقف/ادامه/لغو
  - نمایش پیشرفت هر آیتم
  - مدیریت صف فشرده‌سازی
  - توضیحات کامل به فارسی در کد
--> 
