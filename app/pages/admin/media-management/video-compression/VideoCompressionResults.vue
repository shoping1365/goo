<template>
  <div class="order-5 bg-white rounded-2xl shadow-lg border border-gray-200 mb-6 mx-4 w-full max-w-full min-w-0 overflow-x-hidden">
    <div class="flex items-center justify-between cursor-pointer px-6 py-4 border-b border-orange-300" @click="toggleSection('compressionResults')">
      <div class="flex items-center gap-6">
        <span class="text-base font-bold bg-orange-100 text-orange-700 px-4 py-1 rounded-lg">نتایج فشرده‌سازی</span>
      </div>
      <!-- Filters and Toggle Icon - در سمت راست هدر کانتینر -->
      <div class="flex items-center gap-2">
        <input
          v-model="searchQuery"
          type="text"
          placeholder="جستجو در نتایج..."
          class="px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500"
          @click.stop
        >
        <select
          v-model="filterStatus"
          class="px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500"
          @click.stop
        >
          <option value="">همه وضعیت‌ها</option>
          <option value="completed">تکمیل شده</option>
          <option value="error">خطا</option>
          <option value="pending">در انتظار</option>
          <option value="processing">در حال پردازش</option>
        </select>
        <span class="text-gray-500 text-2xl">{{ sections.compressionResults ? '−' : '+' }}</span>
      </div>
    </div>
         <div v-show="sections.compressionResults" class="p-6 w-full min-w-0">
       <!-- Results Summary -->
       <div class="mb-6">
        <div class="grid grid-cols-2 md:grid-cols-4 gap-6">
          <div class="bg-blue-50 rounded-lg p-6 text-center">
            <div class="text-2xl font-bold text-blue-600">{{ compressionJobs.length }}</div>
            <div class="text-sm text-blue-700">کل نتایج</div>
          </div>
          <div class="bg-green-50 rounded-lg p-6 text-center">
            <div class="text-2xl font-bold text-green-600">{{ formatFileSize(totalOriginalSize) }}</div>
            <div class="text-sm text-green-700">حجم اصلی</div>
          </div>
          <div class="bg-purple-50 rounded-lg p-6 text-center">
            <div class="text-2xl font-bold text-purple-600">{{ formatFileSize(totalCompressedSize) }}</div>
            <div class="text-sm text-purple-700">حجم فشرده</div>
          </div>
          <div class="bg-orange-50 rounded-lg p-6 text-center">
            <div class="text-2xl font-bold text-orange-600">{{ formatFileSize(totalSavedSize) }}</div>
            <div class="text-sm text-orange-700">صرفه‌جویی</div>
          </div>
        </div>
      </div>

             <!-- Results Controls -->
       <div class="flex flex-wrap gap-3 mb-6">
         <button
           v-if="selectedResults.length > 0"
           class="bg-red-500 hover:bg-red-600 text-white px-4 py-2 rounded-lg text-sm font-medium"
           @click="deleteSelectedResults"
         >
           حذف انتخاب شده‌ها
         </button>
         <button
           v-if="selectedResults.length > 0"
           class="bg-green-500 hover:bg-green-600 text-white px-4 py-2 rounded-lg text-sm font-medium"
           @click="batchDownload"
         >
           دانلود انتخاب شده‌ها
         </button>
       </div>

      <!-- Results Table -->
      <div v-if="filteredResults.length === 0" class="text-center py-8">
        <svg class="w-16 h-16 text-gray-300 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
        </svg>
        <h3 class="text-lg font-medium text-gray-900 mb-2">نتیجه‌ای یافت نشد</h3>
        <p class="text-gray-500">ویدیوهایی را فشرده کنید تا نتایج نمایش داده شود</p>
      </div>
      <div v-else class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                <input
                  v-model="selectAllResults"
                  type="checkbox"
                  class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                  @change="toggleSelectAllResults"
                >
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">ویدیو</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">حجم اصلی</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">حجم فشرده</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کاهش</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="job in filteredResults" :key="job.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                                  <input
                    v-model="selectedResults"
                    :value="job.id"
                    type="checkbox"
                    class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                  >
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="w-10 h-8 bg-gray-100 rounded flex items-center justify-center mr-3">
                    <svg class="w-5 h-5 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z"></path>
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
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ formatFileSize(job.compressed_size || 0) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span
v-if="job.original_size && job.compressed_size" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full" :class="{
                  'bg-green-100 text-green-800': calculateReduction(job.original_size, job.compressed_size) > 50,
                  'bg-yellow-100 text-yellow-800': calculateReduction(job.original_size, job.compressed_size) > 20 && calculateReduction(job.original_size, job.compressed_size) <= 50,
                  'bg-red-100 text-red-800': calculateReduction(job.original_size, job.compressed_size) <= 20
                }">
                  {{ calculateReduction(job.original_size, job.compressed_size) }}%
                </span>
                <span v-else class="text-gray-400">-</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getStatusBadgeClass(job.status)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                  {{ getStatusText(job.status) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <div class="flex space-x-2 space-x-reverse">
                  <button
                    v-if="job.status === 'completed' && job.media?.file_path"
                    class="text-blue-600 hover:text-blue-900"
                    title="دانلود"
                    @click="downloadCompressed(job)"
                  >
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
                    </svg>
                  </button>
                  <button
                    v-if="job.status === 'completed'"
                    class="text-green-600 hover:text-green-900"
                    title="مقایسه"
                    @click="compareVideos(job)"
                  >
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
                    </svg>
                  </button>
                  <button
                    class="text-red-600 hover:text-red-900"
                    title="حذف"
                    @click="removeResult(job)"
                  >
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
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
const selectedResults = ref<number[]>([])
const selectAllResults = ref(false)

// Collapsible sections state
const defaultSections = {
  compressionResults: true
} as const

const sections = reactive<Record<keyof typeof defaultSections, boolean>>({ ...defaultSections })

function toggleSection(section: keyof typeof sections) {
  sections[section] = !sections[section]
}

// Computed properties
const filteredResults = computed(() => {
  let results = compressionJobs.value
  if (searchQuery.value) {
    results = results.filter(job => 
      job.media?.filename?.toLowerCase().includes(searchQuery.value.toLowerCase())
    )
  }
  if (filterStatus.value) {
    results = results.filter(job => job.status === filterStatus.value)
  }
  return results
})

const totalOriginalSize = computed(() => 
  compressionJobs.value.reduce((sum, job) => sum + (job.original_size || 0), 0)
)

const totalCompressedSize = computed(() => 
  compressionJobs.value.reduce((sum, job) => sum + (job.compressed_size || 0), 0)
)

const totalSavedSize = computed(() => totalOriginalSize.value - totalCompressedSize.value)

// Methods
const toggleSelectAllResults = () => {
  if (selectAllResults.value) {
    selectedResults.value = filteredResults.value.map(job => job.id)
  } else {
    selectedResults.value = []
  }
}

const deleteSelectedResults = () => {
  if (selectedResults.value.length === 0) {
    alert('هیچ نتیجه‌ای انتخاب نشده است')
    return
  }
  compressionJobs.value = compressionJobs.value.filter(job => !selectedResults.value.includes(job.id))
  selectedResults.value = []
  selectAllResults.value = false
  alert('نتایج انتخاب‌شده حذف شدند')
}

const batchDownload = () => {
  if (selectedResults.value.length === 0) {
    alert('هیچ نتیجه‌ای انتخاب نشده است')
    return
  }
  alert('دانلود همه نتایج انتخاب‌شده (شبیه‌سازی)')
}

const downloadCompressed = (job: CompressionJob) => {
  if (job.media?.file_path) {
    const link = document.createElement('a')
    link.href = job.media.file_path
    link.download = `compressed-${job.media.filename}`
    link.click()
  }
}

const compareVideos = (job: CompressionJob) => {
  alert(`مقایسه ویدیو ${job.media?.filename || 'نامشخص'} (شبیه‌سازی)`)
}

const removeResult = (job: CompressionJob) => {
  if (confirm(`آیا از حذف نتیجه "${job.media?.filename || 'نامشخص'}" اطمینان دارید؟`)) {
    const index = compressionJobs.value.findIndex(j => j.id === job.id)
    if (index > -1) {
      compressionJobs.value.splice(index, 1)
    }
  }
}

// تعریف interface ها
interface ApiResponse<T> {
  data?: T
  success?: boolean
  message?: string
}

interface CompressionJobDisplay {
  id: string
  filename: string
  originalSize: number
  compressedSize: number
  status: string
  createdAt: string
  completedAt?: string
  compressionRatio?: number
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

const calculateReduction = (original: number, compressed: number) => {
  if (!original || !compressed) return 0
  return Math.round(((original - compressed) / original) * 100)
}

const fetchCompressionResults = async () => {
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
    console.error('خطا در دریافت نتایج فشرده‌سازی:', error)
  }
}

let refreshInterval: NodeJS.Timeout | null = null

onMounted(() => {
  fetchCompressionResults()
  // بروزرسانی هر 10 ثانیه
  refreshInterval = setInterval(fetchCompressionResults, 0)
})

onUnmounted(() => {
  if (refreshInterval) {
    clearInterval(refreshInterval)
  }
})

// Expose methods to parent
defineExpose({
  compressionJobs,
  fetchCompressionResults
})
</script>

<!--
  کامپوننت نتایج فشرده‌سازی ویدیو
  - نمایش خلاصه آماری نتایج از API
  - جدول نتایج با قابلیت جستجو و فیلتر
  - عملیات دانلود، مقایسه و حذف
  - انتخاب چندتایی و عملیات دسته‌ای
  - بروزرسانی خودکار هر 10 ثانیه
  - توضیحات کامل به فارسی در کد
--> 
