<template>
  <div class="p-6">
    <div class="max-w-7xl mx-auto">
      <!-- Header -->
      <div class="mb-8">
        <div class="flex items-center justify-between">
          <div>
            <h1 class="text-3xl font-bold text-gray-900 mb-2">انتقال محصولات دیجی‌کالا</h1>
            <p class="text-gray-600">وارد کردن محصولات از لینک دسته‌بندی‌های دیجی‌کالا</p>
          </div>
            <div class="flex items-center gap-4">
              <div class="text-sm text-gray-500">
                آخرین بروزرسانی: {{ lastUpdate }}
              </div>
              <button
                @click="refreshStats"
                :disabled="isRefreshing"
                class="px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 disabled:opacity-50 flex items-center gap-2"
              >
                <span v-if="isRefreshing">⟳</span>
                <span v-else>↻</span>
                بروزرسانی
              </button>
            </div>
          </div>
        </div>

        <!-- آمار کلی -->
        <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
          <div class="bg-white p-6 rounded-lg shadow border">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm font-medium text-gray-600">کل import ها</p>
                <p class="text-2xl font-bold text-green-600">{{ formatNumber(stats?.total_imports || 0) }}</p>
              </div>
              <div class="p-3 bg-green-100 rounded-full">
                <span class="text-green-600 text-xl">✓</span>
              </div>
            </div>
          </div>

          <div class="bg-white p-6 rounded-lg shadow border">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm font-medium text-gray-600">در حال پردازش</p>
                <p class="text-2xl font-bold text-blue-600">{{ formatNumber(stats?.in_progress_imports || 0) }}</p>
              </div>
              <div class="p-3 bg-blue-100 rounded-full">
                <span class="text-blue-600 text-xl">⏳</span>
              </div>
            </div>
          </div>

          <div class="bg-white p-6 rounded-lg shadow border">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm font-medium text-gray-600">ناموفق</p>
                <p class="text-2xl font-bold text-red-600">{{ formatNumber(stats?.failed_imports || 0) }}</p>
              </div>
              <div class="p-3 bg-red-100 rounded-full">
                <span class="text-red-600 text-xl">⚠</span>
              </div>
            </div>
          </div>

          <div class="bg-white p-6 rounded-lg shadow border">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm font-medium text-gray-600">میانگین سرعت</p>
                <p class="text-2xl font-bold text-purple-600">{{ formatNumber(stats?.average_speed || 0) }}</p>
              </div>
              <div class="p-3 bg-purple-100 rounded-full">
                <span class="text-purple-600 text-xl">⚡</span>
              </div>
            </div>
          </div>
        </div>

        <!-- فرم Import جدید -->
        <div class="bg-white rounded-lg shadow border mb-8">
          <div class="p-6 border-b">
            <h2 class="text-lg font-semibold text-gray-900">شروع Import جدید</h2>
            <p class="text-sm text-gray-600 mt-1">لینک دسته‌بندی دیجی‌کالا را وارد کنید</p>
          </div>
          <div class="p-6">
            <form @submit.prevent="startImport" class="space-y-6">
              <!-- URL Input -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">
                  لینک دسته‌بندی دیجی‌کالا
                </label>
                <div class="relative">
                  <input
                    v-model="importForm.categoryUrl"
                    type="url"
                    placeholder="https://www.digikala.com/search/category-mobile-phone/"
                    class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                    required
                  />
                  <button
                    type="button"
                    @click="validateUrl"
                    :disabled="!importForm.categoryUrl || isValidating"
                    class="absolute left-2 top-2 px-3 py-1 text-sm bg-gray-100 text-gray-600 rounded hover:bg-gray-200 disabled:opacity-50"
                  >
                    <span v-if="isValidating">⟳</span>
                    <span v-else>بررسی</span>
                  </button>
                </div>
                <p class="text-sm text-gray-500 mt-1">
                  مثال: https://www.digikala.com/search/category-mobile-phone/
                </p>
              </div>

              <!-- تنظیمات -->
              <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">
                    تعداد محصول در دقیقه
                  </label>
                  <select
                    v-model.number="importForm.itemsPerMinute"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                  >
                    <option :value="10">10 محصول در دقیقه</option>
                    <option :value="20">20 محصول در دقیقه</option>
                    <option :value="30">30 محصول در دقیقه</option>
                    <option :value="50">50 محصول در دقیقه</option>
                    <option :value="100">100 محصول در دقیقه</option>
                  </select>
                </div>

                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">
                    حداکثر تعداد محصول
                  </label>
                  <input
                    v-model.number="importForm.maxItems"
                    type="number"
                    min="1"
                    max="10000"
                    placeholder="1000"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                  />
                </div>

                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">
                    دسته‌بندی مقصد
                  </label>
                  <select
                    v-model="importForm.targetCategoryId"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                  >
                    <option value="">انتخاب دسته‌بندی...</option>
                    <option v-for="cat in flattenedCategories" :key="cat.id" :value="cat.id">
                      {{ cat.indentedName }}
                    </option>
                  </select>
                </div>
              </div>

              <!-- گزینه‌های اضافی -->
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div class="space-y-3">
                  <label class="flex items-center">
                    <input
                      v-model="importForm.skipExisting"
                      type="checkbox"
                      class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                    />
                    <span class="ml-2 text-sm text-gray-700">رد کردن محصولات موجود</span>
                  </label>
                  <label class="flex items-center">
                    <input
                      v-model="importForm.importImages"
                      type="checkbox"
                      class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                    />
                    <span class="ml-2 text-sm text-gray-700">دانلود تصاویر محصولات</span>
                  </label>
                </div>
                <div class="space-y-3">
                  <label class="flex items-center">
                    <input
                      v-model="importForm.updatePrices"
                      type="checkbox"
                      class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                    />
                    <span class="ml-2 text-sm text-gray-700">به‌روزرسانی قیمت‌ها</span>
                  </label>
                  <label class="flex items-center">
                    <input
                      v-model="importForm.importReviews"
                      type="checkbox"
                      class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                    />
                    <span class="ml-2 text-sm text-gray-700">Import نظرات کاربران</span>
                  </label>
                </div>
              </div>

              <!-- دکمه شروع -->
              <div class="flex justify-end">
                <button
                  type="submit"
                  :disabled="isImporting || !importForm.categoryUrl"
                  class="px-6 py-3 bg-green-600 text-white rounded-lg hover:bg-green-700 disabled:opacity-50 disabled:cursor-not-allowed flex items-center gap-2"
                >
                  <span v-if="isImporting">⟳</span>
                  <span v-else>⬇</span>
                  {{ isImporting ? 'در حال Import...' : 'شروع Import' }}
                </button>
              </div>
            </form>
          </div>
        </div>

        <!-- نمایش پیشرفت فعلی -->
        <div v-if="currentImport" class="bg-white rounded-lg shadow border mb-8">
          <div class="p-6 border-b">
            <div class="flex items-center justify-between">
              <h2 class="text-lg font-semibold text-gray-900">پیشرفت Import جاری</h2>
              <button
                @click="cancelImport"
                class="px-4 py-2 bg-red-500 text-white rounded-lg hover:bg-red-600 flex items-center gap-2"
              >
                <span>⏹</span>
                توقف
              </button>
            </div>
          </div>
          <div class="p-6">
            <!-- نوار پیشرفت -->
            <div class="mb-6">
              <div class="flex justify-between text-sm text-gray-600 mb-2">
                <span>{{ currentImport.processed }} از {{ currentImport.total }} محصول</span>
                <span>{{ Math.round(currentImport.progress) }}%</span>
              </div>
              <div class="w-full bg-gray-200 rounded-full h-3">
                <div
                  class="bg-blue-600 h-3 rounded-full transition-all duration-300"
                  :style="{ width: currentImport.progress + '%' }"
                ></div>
              </div>
            </div>

            <!-- آمار لحظه‌ای -->
            <div class="grid grid-cols-1 md:grid-cols-4 gap-4 mb-6">
              <div class="text-center p-4 bg-green-50 rounded-lg">
                <div class="text-2xl font-bold text-green-600">{{ currentImport.successful }}</div>
                <div class="text-sm text-gray-600">موفق</div>
              </div>
              <div class="text-center p-4 bg-red-50 rounded-lg">
                <div class="text-2xl font-bold text-red-600">{{ currentImport.failed }}</div>
                <div class="text-sm text-gray-600">ناموفق</div>
              </div>
              <div class="text-center p-4 bg-yellow-50 rounded-lg">
                <div class="text-2xl font-bold text-yellow-600">{{ currentImport.skipped }}</div>
                <div class="text-sm text-gray-600">رد شده</div>
              </div>
              <div class="text-center p-4 bg-blue-50 rounded-lg">
                <div class="text-2xl font-bold text-blue-600">{{ currentImport.speed }}</div>
                <div class="text-sm text-gray-600">در دقیقه</div>
              </div>
            </div>

            <!-- زمان باقی‌مانده -->
            <div v-if="currentImport.eta" class="text-center text-gray-600">
              زمان تخمینی باقی‌مانده: {{ formatDuration(currentImport.eta) }}
            </div>
          </div>
        </div>

        <!-- لاگ‌های زنده -->
        <div class="bg-white rounded-lg shadow border mb-8">
          <div class="p-6 border-b">
            <div class="flex items-center justify-between">
              <h2 class="text-lg font-semibold text-gray-900">لاگ‌های Import</h2>
              <div class="flex items-center gap-2">
                <label class="flex items-center text-sm">
                  <input
                    v-model="autoScroll"
                    type="checkbox"
                    class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                  />
                  <span class="ml-1">اسکرول خودکار</span>
                </label>
                <button
                  @click="clearLogs"
                  class="px-3 py-1 bg-gray-500 text-white rounded hover:bg-gray-600 text-sm"
                >
                  پاک کردن
                </button>
              </div>
            </div>
          </div>
          <div
            ref="logsContainer"
            class="p-4 bg-gray-900 text-gray-100 font-mono text-sm max-h-96 overflow-y-auto"
          >
            <div
              v-for="(log, index) in logs"
              :key="index"
              :class="getLogClass(log.type)"
              class="mb-1 flex items-start gap-2"
            >
              <span class="text-gray-500 text-xs whitespace-nowrap">{{ formatTime(log.timestamp) }}</span>
              <span :class="getLogIconClass(log.type)">{{ getLogIcon(log.type) }}</span>
              <span class="flex-1">{{ log.message }}</span>
            </div>
            <div v-if="logs.length === 0" class="text-gray-500 text-center py-8">
              هیچ لاگی موجود نیست
            </div>
          </div>
        </div>

        <!-- تاریخچه Import ها -->
        <div class="bg-white rounded-lg shadow border">
          <div class="p-6 border-b">
            <h2 class="text-lg font-semibold text-gray-900">تاریخچه Import ها</h2>
          </div>
          <div class="overflow-x-auto">
            <table class="w-full">
              <thead class="bg-gray-50">
                <tr>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase">شناسه</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase">URL</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase">وضعیت</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase">محصولات</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase">تاریخ</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase">مدت زمان</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase">عملیات</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-gray-200">
                <tr v-for="import_ in importHistory" :key="import_.id">
                  <td class="px-6 py-4 text-sm text-gray-900">#{{ import_.id }}</td>
                  <td class="px-6 py-4 text-sm text-blue-600 max-w-xs truncate">
                    <a :href="import_.url" target="_blank" class="hover:underline">
                      {{ import_.url }}
                    </a>
                  </td>
                  <td class="px-6 py-4">
                    <span :class="getStatusClass(import_.status)" class="px-2 py-1 text-xs rounded-full">
                      {{ getStatusText(import_.status) }}
                    </span>
                  </td>
                  <td class="px-6 py-4 text-sm text-gray-900">
                    {{ import_.successful }}/{{ import_.total }}
                  </td>
                  <td class="px-6 py-4 text-sm text-gray-600">
                    {{ formatDate(import_.created_at) }}
                  </td>
                  <td class="px-6 py-4 text-sm text-gray-600">
                    {{ import_.duration ? formatDuration(import_.duration) : '-' }}
                  </td>
                  <td class="px-6 py-4 text-sm">
                    <div class="flex items-center gap-2">
                      <button
                        @click="viewImportDetails(import_.id)"
                        class="text-blue-600 hover:text-blue-800"
                      >
                        جزئیات
                      </button>
                      <button
                        v-if="import_.status === 'failed'"
                        @click="retryImport(import_.id)"
                        class="text-green-600 hover:text-green-800"
                      >
                        تلاش مجدد
                      </button>
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
</template>

<script setup lang="ts">
// @ts-nocheck
import { ref, nextTick, onMounted } from 'vue'

interface ImportStats {
  total_imports: number
  in_progress_imports: number
  failed_imports: number
  total_products?: number
  imported_products?: number
  failed_products?: number
  success_rate?: number
  average_speed?: number
}

interface ImportSettings {
  itemsPerMinute: number
}

interface ImportForm {
  categoryUrl: string
  itemsPerMinute: number
  maxItems: number
  targetCategoryId: string
  skipExisting: boolean
  importImages: boolean
  updatePrices: boolean
  importReviews: boolean
}

interface CurrentImport {
  id: string
  processed: number
  total: number
  progress: number
  successful: number
  failed: number
  skipped: number
  speed: number
  eta: number | null
}

interface LogEntry {
  type: 'info' | 'success' | 'warning' | 'error'
  message: string
  timestamp: Date
}

interface ImportHistory {
  id: string
  url: string
  status: 'running' | 'completed' | 'failed' | 'cancelled'
  total: number
  successful: number
  created_at: string
  duration: number | null
}

interface LocalCategory {
  id: number
  name: string
  parent_id?: number | null
}

interface FlattenedCategory {
  id: number
  name: string
  indentedName: string
  parent_id?: number | null
}

// متا و لایوت
definePageMeta({
  layout: 'admin-main',
  middleware: 'developer-only'
})

useHead({
  title: 'انتقال محصولات دیجی‌کالا - پنل توسعه‌دهنده'
})

// State
const stats = ref<ImportStats>({
  total_imports: 0,
  in_progress_imports: 0,
  failed_imports: 0
})

const importSettings = ref<ImportSettings>({
  itemsPerMinute: 30
})

const importForm = ref<ImportForm>({
  categoryUrl: '',
  itemsPerMinute: 30,
  maxItems: 1000,
  targetCategoryId: '',
  skipExisting: true,
  importImages: true,
  updatePrices: false,
  importReviews: false
})

const currentImport = ref<CurrentImport | null>(null)
const logs = ref<LogEntry[]>([])
const importHistory = ref<ImportHistory[]>([])
const localCategories = ref<LocalCategory[]>([])
const flattenedCategories = ref<FlattenedCategory[]>([])

// UI State
const isImporting = ref(false)
const isValidating = ref(false)
const isRefreshing = ref(false)
const autoScroll = ref(true)
const lastUpdate = ref('')

// Refs
const logsContainer = ref<HTMLElement | null>(null)

// Methods
const formatTime = (date: Date): string => {
  return date.toLocaleTimeString('fa-IR')
}

const formatDate = (dateString: string): string => {
  return new Date(dateString).toLocaleDateString('fa-IR')
}

const formatDuration = (seconds: number): string => {
  const hours = Math.floor(seconds / 3600)
  const minutes = Math.floor((seconds % 3600) / 60)
  const secs = seconds % 60

  if (hours > 0) {
    return `${hours}:${minutes.toString().padStart(2, '0')}:${secs.toString().padStart(2, '0')}`
  }
  return `${minutes}:${secs.toString().padStart(2, '0')}`
}

const addLog = (type: LogEntry['type'], message: string): void => {
  logs.value.push({
    type,
    message,
    timestamp: new Date()
  })

  if (logs.value.length > 1000) {
    logs.value = logs.value.slice(-500)
  }

  if (autoScroll.value) {
    nextTick(() => {
      if (logsContainer.value) {
        logsContainer.value.scrollTop = logsContainer.value.scrollHeight
      }
    })
  }
}

const getLogClass = (type: LogEntry['type']): string => {
  const classes: Record<LogEntry['type'], string> = {
    info: 'text-blue-400',
    success: 'text-green-400',
    warning: 'text-yellow-400',
    error: 'text-red-400'
  }
  return classes[type] || 'text-gray-400'
}

const getLogIconClass = (type: LogEntry['type']): string => {
  return getLogClass(type)
}

const getLogIcon = (type: LogEntry['type']): string => {
  const icons: Record<LogEntry['type'], string> = {
    info: 'ℹ',
    success: '✓',
    warning: '⚠',
    error: '✗'
  }
  return icons[type] || '•'
}

const getStatusClass = (status: string): string => {
  const classes: Record<string, string> = {
    running: 'bg-blue-100 text-blue-800',
    completed: 'bg-green-100 text-green-800',
    failed: 'bg-red-100 text-red-800',
    cancelled: 'bg-gray-100 text-gray-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status: string): string => {
  const texts: Record<string, string> = {
    running: 'در حال اجرا',
    completed: 'تکمیل شده',
    failed: 'ناموفق',
    cancelled: 'لغو شده'
  }
  return texts[status] || status
}

// فرمت کردن اعداد با کاما
const formatNumber = (num: number): string => {
  try {
    return num?.toLocaleString() || '0'
  } catch {
    return '0'
  }
}

const refreshStats = async (): Promise<void> => {
  isRefreshing.value = true
  try {
    console.log('🔄 Refreshing stats...')
    const response = await $fetch<ImportStats>('/api/admin/digikala/stats')
    console.log('📊 Stats response:', response)
    stats.value = response
    lastUpdate.value = new Date().toLocaleTimeString('fa-IR')
  } catch (error) {
    console.error('❌ Error refreshing stats:', error)
    const err = error as Error
    addLog('error', `خطا در دریافت آمار: ${err.message}`)
    // اگر خطا رخ داد، از مقدار پیش‌فرض استفاده کن
    stats.value = {
      total_imports: 0,
      in_progress_imports: 0,
      failed_imports: 0
    }
  } finally {
    isRefreshing.value = false
  }
}

const flattenCategories = (categories: LocalCategory[]): FlattenedCategory[] => {
  const result: FlattenedCategory[] = []
  
  // Create a map for quick lookup
  const categoryMap = new Map<number, LocalCategory>()
  categories.forEach(cat => categoryMap.set(cat.id, cat))
  
  // Recursive function to build flattened list with indentation
  const addCategory = (cat: LocalCategory, level: number = 0) => {
    const indent = '—'.repeat(level) + (level > 0 ? ' ' : '')
    result.push({
      id: cat.id,
      name: cat.name,
      indentedName: indent + cat.name,
      parent_id: cat.parent_id
    })
    
    // Find and add children
    const children = categories.filter(c => c.parent_id === cat.id)
    children.forEach(child => addCategory(child, level + 1))
  }
  
  // Add root categories (those without parent or with null parent)
  const rootCategories = categories.filter(cat => !cat.parent_id)
  rootCategories.forEach(cat => addCategory(cat, 0))
  
  return result
}

const loadLocalCategories = async (): Promise<void> => {
  try {
    // استفاده از همان endpoint که در صفحه product-categories کار می‌کند
    const response = await $fetch('/api/product-categories?all=1')
    
    console.log('📦 دسته‌بندی‌های دریافتی:', response)
    
    // پاسخ ممکنه {data: []} یا مستقیماً [] باشه
    const categories = Array.isArray(response) ? response : (response?.data || [])
    localCategories.value = Array.isArray(categories) ? categories : []
    console.log('📦 تعداد دسته‌بندی‌ها:', localCategories.value.length)
    
    // Flatten categories for display in select with proper indentation
    flattenedCategories.value = flattenCategories(localCategories.value)
    
    if (localCategories.value.length === 0) {
      addLog('warning', 'هیچ دسته‌بندی یافت نشد! لطفاً ابتدا دسته‌بندی‌ها را ایجاد کنید.')
    } else {
      addLog('info', `${localCategories.value.length} دسته‌بندی بارگذاری شد.`)
    }
  } catch (error) {
    const err = error as Error
    addLog('error', `خطا در دریافت دسته‌بندی‌ها: ${err.message}`)
    console.error('❌ خطا در دریافت دسته‌بندی‌ها:', err)
  }
}

const loadImportHistory = async (): Promise<void> => {
  try {
    // اصلاح مسیر API از go-api به api
    const response = await $fetch<{ data: ImportHistory[] }>('/api/admin/digikala/import-history')
    importHistory.value = response?.data || []
    
    if (importHistory.value.length > 0) {
      addLog('info', `${importHistory.value.length} مورد تاریخچه Import بارگذاری شد.`)
    } else {
      addLog('info', 'هیچ تاریخچه‌ای یافت نشد.')
    }
  } catch (error) {
    const err = error as Error
    addLog('warning', `خطا در بارگذاری تاریخچه: ${err.message}`)
    console.warn('⚠️ خطا در بارگذاری تاریخچه:', err)
  }
}

const validateUrl = async (): Promise<void> => {
  if (!importForm.value.categoryUrl) return

  isValidating.value = true
  addLog('info', `بررسی URL: ${importForm.value.categoryUrl}`)

  try {
    const response = await $fetch<{
      valid: boolean
      totalProducts?: number
      message?: string
    }>('/api/admin/digikala/validate-url', { // مسیر اصلاح شده
      method: 'POST',
      body: { url: importForm.value.categoryUrl }
    })

    if (response.valid) {
      addLog('success', `URL معتبر است. تعداد محصولات یافت شده: ${response.totalProducts || 0}`)
    } else {
      addLog('error', `URL نامعتبر: ${response.message || 'خطای نامشخص'}`)
    }
  } catch (error) {
    const err = error as Error
    addLog('error', `خطا در بررسی URL: ${err.message}`)
  } finally {
    isValidating.value = false
  }
}

const startImport = async (): Promise<void> => {
  if (!importForm.value.categoryUrl) return

  isImporting.value = true
  addLog('info', 'شروع فرآیند Import...')

  try {
    const response = await $fetch<{
      import_id: number
      category_title: string
      success: boolean
    }>('/api/admin/digikala/start-import', {
      method: 'POST',
      body: {
        category_url: importForm.value.categoryUrl,
        items_per_minute: importForm.value.itemsPerMinute,
        max_items: importForm.value.maxItems,
        skip_existing: importForm.value.skipExisting,
        import_images: importForm.value.importImages,
        target_category_id: importForm.value.targetCategoryId
      }
    })

    currentImport.value = {
      id: String(response.import_id),
      processed: 0,
      total: importForm.value.maxItems || 0,
      progress: 0,
      successful: 0,
      failed: 0,
      skipped: 0,
      speed: 0,
      eta: null
    }

    addLog('success', `Import شروع شد. شناسه: ${response.import_id}`)
    
    startProgressMonitoring(String(response.import_id))

  } catch (error) {
    const err = error as Error
    addLog('error', `خطا در شروع Import: ${err.message}`)
    isImporting.value = false
  }
}

const startProgressMonitoring = (importId: string): void => {
  const interval = setInterval(async () => {
    try {
      const progress = await $fetch<{
        processed: number
        total: number
        progress: number
        successful: number
        failed: number
        skipped: number
        speed: number
        eta: number | null
        status: string
        logs?: Array<{ type: LogEntry['type'], message: string }>
      }>(`/api/admin/digikala/import-progress/${importId}`) // مسیر اصلاح شده
      
      if (currentImport.value) {
        Object.assign(currentImport.value, progress)
      }

      if (progress.logs) {
        progress.logs.forEach((log) => {
          addLog(log.type, log.message)
        })
      }

      if (progress.status === 'completed' || progress.status === 'failed' || progress.status === 'cancelled') {
        clearInterval(interval)
        isImporting.value = false
        currentImport.value = null
        
        addLog(
          progress.status === 'completed' ? 'success' : 'error',
          `Import ${getStatusText(progress.status)}`
        )
        
        await refreshStats()
        await loadImportHistory()
      }

    } catch (error) {
      const err = error as Error
      addLog('error', `خطا در دریافت وضعیت Import: ${err.message}`)
      clearInterval(interval)
      isImporting.value = false
    }
  }, 2000)
}

const cancelImport = async (): Promise<void> => {
  if (!currentImport.value) return

  try {
    await $fetch(`/api/admin/digikala/cancel-import/${currentImport.value.id}`, { // مسیر اصلاح شده
      method: 'POST'
    })
    
    addLog('warning', 'درخواست لغو Import ارسال شد')
  } catch (error) {
    const err = error as Error
    addLog('error', `خطا در لغو Import: ${err.message}`)
  }
}

const retryImport = async (importId: string): Promise<void> => {
  try {
    await $fetch(`/api/admin/digikala/retry-import/${importId}`, { // مسیر اصلاح شده
      method: 'POST'
    })
    addLog('info', `Import مجدد شروع شد: ${importId}`)
    await loadImportHistory()
  } catch (error) {
    const err = error as Error
    addLog('error', `خطا در Import مجدد: ${err.message}`)
  }
}

// Lifecycle
onMounted(async () => {
  await refreshStats()
  await loadLocalCategories()
  await loadImportHistory()
})
</script>

<style scoped>
/* استایل‌های اضافی در صورت نیاز */
</style>
