<template>
  <div>
    <div class="p-6">
      <!-- Header -->
      <div class="mb-6">
        <div class="flex items-center gap-4 mb-4">
          <button
            @click="$router.back()"
            class="p-2 hover:bg-gray-100 rounded-lg transition-colors"
          >
            <Icon name="mdi:arrow-right" class="text-xl text-gray-600" />
          </button>
          <div class="flex-1">
            <h1 class="text-2xl font-bold text-gray-900">جزئیات Import #{{ importId }}</h1>
            <p class="text-sm text-gray-600 mt-1">اطلاعات کامل فرآیند Import از دیجی‌کالا</p>
          </div>
          <div v-if="importData" class="flex items-center gap-3">
            <span :class="getStatusClass(importData.status)" class="px-3 py-1 rounded-full text-sm font-medium">
              {{ getStatusText(importData.status) }}
            </span>
          </div>
        </div>
      </div>

      <!-- Loading State -->
      <div v-if="!importData" class="space-y-6">
        <div class="bg-white rounded-lg shadow p-6">
          <div class="animate-pulse space-y-4">
            <div class="h-4 bg-gray-200 rounded w-3/4"></div>
            <div class="h-4 bg-gray-200 rounded w-1/2"></div>
            <div class="h-4 bg-gray-200 rounded w-2/3"></div>
          </div>
        </div>
      </div>

      <!-- Content -->
      <div v-else class="space-y-6">
        <!-- آمار پیشرفت -->
        <div class="bg-white rounded-lg shadow">
          <div class="p-6 border-b">
            <h2 class="text-lg font-semibold text-gray-900">آمار پیشرفت</h2>
          </div>
          <div class="p-6">
            <!-- نوار پیشرفت -->
            <div class="mb-8">
              <div class="flex justify-between text-sm text-gray-600 mb-3">
                <span class="font-medium">پیشرفت کلی</span>
                <span class="font-bold text-blue-600">
                  {{ importData.processed?.toLocaleString() || 0 }} از {{ importData.totalProducts?.toLocaleString() || 0 }}
                </span>
              </div>
              <div class="w-full bg-gray-200 rounded-full h-4 overflow-hidden">
                <div
                  class="bg-gradient-to-r from-blue-500 to-blue-600 h-4 rounded-full transition-all duration-500 flex items-center justify-end px-2"
                  :style="{ width: Math.round(((importData.processed || 0) / (importData.totalProducts || 1)) * 100) + '%' }"
                >
                  <span class="text-xs text-white font-bold">
                    {{ Math.round(((importData.processed || 0) / (importData.totalProducts || 1)) * 100) }}%
                  </span>
                </div>
              </div>
            </div>

            <!-- آمار تفصیلی -->
            <div class="grid grid-cols-1 md:grid-cols-4 gap-4">
              <div class="text-center p-4 bg-green-50 rounded-lg border">
                <div class="text-3xl font-bold text-green-600">{{ (importData.successful || 0).toLocaleString() }}</div>
                <div class="text-sm text-gray-600 mt-1">موفق</div>
                <div class="text-xs text-gray-500">
                  {{ Math.round(((importData.successful || 0) / (importData.processed || 1)) * 100) }}%
                </div>
              </div>

              <div class="text-center p-4 bg-red-50 rounded-lg border">
                <div class="text-3xl font-bold text-red-600">{{ (importData.failed || 0).toLocaleString() }}</div>
                <div class="text-sm text-gray-600 mt-1">ناموفق</div>
                <div class="text-xs text-gray-500">
                  {{ Math.round(((importData.failed || 0) / (importData.processed || 1)) * 100) }}%
                </div>
              </div>

              <div class="text-center p-4 bg-yellow-50 rounded-lg border">
                <div class="text-3xl font-bold text-yellow-600">{{ (importData.skipped || 0).toLocaleString() }}</div>
                <div class="text-sm text-gray-600 mt-1">رد شده</div>
                <div class="text-xs text-gray-500">
                  {{ Math.round(((importData.skipped || 0) / (importData.processed || 1)) * 100) }}%
                </div>
              </div>

              <div class="text-center p-4 bg-purple-50 rounded-lg border">
                <div class="text-3xl font-bold text-purple-600">{{ importData.speed || 0 }}</div>
                <div class="text-sm text-gray-600 mt-1">سرعت</div>
                <div class="text-xs text-gray-500">در دقیقه</div>
              </div>
            </div>
          </div>
        </div>

        <!-- اطلاعات کلی -->
        <div class="bg-white rounded-lg shadow">
          <div class="p-6 border-b">
            <h2 class="text-lg font-semibold text-gray-900">اطلاعات کلی</h2>
          </div>
          <div class="p-6">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <!-- URL -->
              <div>
                <label class="block text-sm font-medium text-gray-600 mb-1">URL منبع</label>
                <a :href="importData.url" target="_blank" class="text-blue-600 hover:underline break-all">
                  {{ importData.url }}
                </a>
              </div>

              <!-- تاریخ شروع -->
              <div>
                <label class="block text-sm font-medium text-gray-600 mb-1">تاریخ شروع</label>
                <p class="text-gray-900">{{ formatDateTime(importData.startTime) }}</p>
              </div>

              <!-- مدت زمان -->
              <div>
                <label class="block text-sm font-medium text-gray-600 mb-1">مدت زمان</label>
                <p class="text-gray-900">{{ getDuration(importData) }}</p>
              </div>

              <!-- دسته‌بندی مقصد -->
              <div>
                <label class="block text-sm font-medium text-gray-600 mb-1">دسته‌بندی مقصد</label>
                <p class="text-gray-900">{{ targetCategoryName || 'تعیین نشده' }}</p>
              </div>
            </div>
          </div>
        </div>

        <!-- تنظیمات Import -->
        <div class="bg-white rounded-lg shadow">
          <div class="p-6 border-b">
            <h2 class="text-lg font-semibold text-gray-900">تنظیمات Import</h2>
          </div>
          <div class="p-6">
            <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
              <div class="text-center p-4 bg-gray-50 rounded-lg">
                <Icon name="mdi:speedometer" class="text-2xl text-blue-600 mx-auto mb-2" />
                <div class="text-2xl font-bold text-gray-900">{{ importData.settings?.itemsPerMinute || 0 }}</div>
                <div class="text-xs text-gray-600 mt-1">در دقیقه</div>
              </div>
              <div class="text-center p-4 bg-gray-50 rounded-lg">
                <Icon name="mdi:package" class="text-2xl text-green-600 mx-auto mb-2" />
                <div class="text-2xl font-bold text-gray-900">{{ importData.settings?.maxItems?.toLocaleString() || 0 }}</div>
                <div class="text-xs text-gray-600 mt-1">حداکثر</div>
              </div>
              <div class="text-center p-4 bg-gray-50 rounded-lg">
                <Icon :name="importData.settings?.skipExisting ? 'mdi:check-circle' : 'mdi:close-circle'" 
                      :class="importData.settings?.skipExisting ? 'text-green-600' : 'text-red-600'" 
                      class="text-2xl mx-auto mb-2" />
                <div class="text-xs text-gray-600">رد موجود</div>
                <div class="text-sm font-medium" :class="importData.settings?.skipExisting ? 'text-green-600' : 'text-red-600'">
                  {{ importData.settings?.skipExisting ? 'بله' : 'خیر' }}
                </div>
              </div>
              <div class="text-center p-4 bg-gray-50 rounded-lg">
                <Icon :name="importData.settings?.importImages ? 'mdi:check-circle' : 'mdi:close-circle'" 
                      :class="importData.settings?.importImages ? 'text-green-600' : 'text-red-600'" 
                      class="text-2xl mx-auto mb-2" />
                <div class="text-xs text-gray-600">تصاویر</div>
                <div class="text-sm font-medium" :class="importData.settings?.importImages ? 'text-green-600' : 'text-red-600'">
                  {{ importData.settings?.importImages ? 'بله' : 'خیر' }}
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- لاگ‌های Import -->
        <div class="bg-white rounded-lg shadow">
          <div class="p-6 border-b">
            <div class="flex items-center justify-between">
              <h2 class="text-lg font-semibold text-gray-900">لاگ‌های Import</h2>
              <div class="flex items-center gap-2">
                <button
                  @click="refreshLogs"
                  :disabled="isRefreshingLogs"
                  class="px-3 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 disabled:opacity-50 text-sm flex items-center gap-2"
                >
                  <Icon :name="isRefreshingLogs ? 'mdi:loading' : 'mdi:refresh'" :class="{ 'animate-spin': isRefreshingLogs }" />
                  بروزرسانی
                </button>
                <button
                  @click="downloadLogs"
                  class="px-3 py-2 bg-gray-600 text-white rounded-lg hover:bg-gray-700 text-sm flex items-center gap-2"
                >
                  <Icon name="mdi:download" />
                  دانلود
                </button>
              </div>
            </div>
          </div>
          <div class="p-4 max-h-96 overflow-y-auto border-t">
            <div v-if="detailedLogs.length === 0" class="text-center py-12 text-gray-500">
              <p>هیچ لاگی موجود نیست</p>
              <p class="text-sm mt-1">لاگ‌ها به صورت لحظه‌ای در اینجا نمایش داده خواهند شد</p>
            </div>
            <div v-else class="space-y-1">
              <div
                v-for="(log, index) in detailedLogs"
                :key="index"
                class="flex items-start gap-2 text-sm p-2"
              >
                <span class="text-gray-500 text-xs whitespace-nowrap">{{ formatTime(log.timestamp) }}</span>
                <span :class="getLogTypeColor(log.type)" class="font-bold">{{ getLogIcon(log.type) }}</span>
                <span class="flex-1 text-gray-700">{{ log.message }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// @ts-nocheck
interface ImportData {
  id: string
  url: string
  status: string
  totalProducts: number
  processed: number
  successful: number
  failed: number
  skipped: number
  speed: number
  startTime: string
  endTime?: string
  settings: {
    itemsPerMinute: number
    maxItems: number
    targetCategoryId: string
    skipExisting: boolean
    importImages: boolean
    updatePrices: boolean
    importReviews: boolean
  }
}

interface LogEntry {
  type: 'info' | 'success' | 'warning' | 'error'
  message: string
  timestamp: string
}

// متا و لایوت
definePageMeta({
  layout: 'admin-main',
  middleware: 'developer-only'
})

// Router params
const route = useRoute()
const importId = route.params.importId as string

useHead({
  title: `جزئیات Import #${importId} - پنل توسعه‌دهنده`
})

// State
const importData = ref<ImportData | null>(null)
const detailedLogs = ref<LogEntry[]>([])
const targetCategoryName = ref<string>('')
const isRefreshingLogs = ref(false)

// Methods
const formatDateTime = (dateString: string) => {
  if (!dateString) return 'نامشخص'
  const date = new Date(dateString)
  return date.toLocaleDateString('fa-IR') + ' ' + date.toLocaleTimeString('fa-IR')
}

const formatTime = (dateString: string) => {
  if (!dateString) return ''
  return new Date(dateString).toLocaleTimeString('fa-IR')
}

const getStatusClass = (status: string) => {
  const classes = {
    running: 'bg-blue-100 text-blue-800 border border-blue-200',
    completed: 'bg-green-100 text-green-800 border border-green-200',
    failed: 'bg-red-100 text-red-800 border border-red-200',
    cancelled: 'bg-gray-100 text-gray-800 border border-gray-200'
  }
  return classes[status as keyof typeof classes] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status: string) => {
  const texts = {
    running: 'در حال اجرا',
    completed: 'تکمیل شده',
    failed: 'ناموفق',
    cancelled: 'لغو شده'
  }
  return texts[status as keyof typeof texts] || status
}

const getDuration = (data: ImportData) => {
  if (!data.startTime) return 'نامشخص'
  
  const start = new Date(data.startTime)
  const end = data.endTime ? new Date(data.endTime) : new Date()
  const diff = Math.floor((end.getTime() - start.getTime()) / 1000)
  
  const hours = Math.floor(diff / 3600)
  const minutes = Math.floor((diff % 3600) / 60)
  const seconds = diff % 60
  
  if (hours > 0) {
    return `${hours} ساعت و ${minutes} دقیقه`
  }
  if (minutes > 0) {
    return `${minutes} دقیقه و ${seconds} ثانیه`
  }
  return `${seconds} ثانیه`
}

const getLogTypeColor = (type: string) => {
  const classes = {
    info: 'text-blue-600',
    success: 'text-green-600',
    warning: 'text-yellow-600',
    error: 'text-red-600'
  }
  return classes[type as keyof typeof classes] || 'text-gray-600'
}

const getLogIcon = (type: string) => {
  const icons = {
    info: 'ℹ',
    success: '✓',
    warning: '⚠',
    error: '✗'
  }
  return icons[type as keyof typeof icons] || '•'
}

const loadImportDetails = async () => {
  try {
    const response = await $fetch<ImportData>(`/go-api/admin/digikala/import-details/${importId}`)
    importData.value = response
    
    // بارگذاری نام دسته‌بندی
    if (response.settings?.targetCategoryId) {
      try {
        const categories = await $fetch<{ data: any[] }>('/api/admin/product-categories?all=1')
        const category = categories.data?.find(c => c.id.toString() === response.settings.targetCategoryId)
        targetCategoryName.value = category?.name || 'نامشخص'
      } catch (error) {
        targetCategoryName.value = 'خطا در بارگذاری'
      }
    }
  } catch (error: any) {
    console.error('Error loading import details:', error)
  }
}

const loadDetailedLogs = async () => {
  try {
    const response = await $fetch<{ logs: LogEntry[] }>(`/go-api/admin/digikala/import-logs/${importId}`)
    detailedLogs.value = response.logs || []
  } catch (error: any) {
    console.error('Error loading logs:', error)
    detailedLogs.value = []
  }
}

const refreshLogs = async () => {
  isRefreshingLogs.value = true
  try {
    await loadDetailedLogs()
  } finally {
    isRefreshingLogs.value = false
  }
}

const downloadLogs = () => {
  const logText = detailedLogs.value
    .map(log => `[${formatTime(log.timestamp)}] ${log.type.toUpperCase()}: ${log.message}`)
    .join('\n')
  
  const blob = new Blob([logText], { type: 'text/plain;charset=utf-8' })
  const url = URL.createObjectURL(blob)
  const link = document.createElement('a')
  link.href = url
  link.download = `import-${importId}-logs.txt`
  document.body.appendChild(link)
  link.click()
  document.body.removeChild(link)
  URL.revokeObjectURL(url)
}

// Lifecycle
onMounted(async () => {
  await loadImportDetails()
  await loadDetailedLogs()
})
</script>

<style scoped>
/* استایل‌های اضافی در صورت نیاز */
</style>
