<template>
  <div class="bg-white rounded-xl shadow-lg p-6">
    <div class="flex items-center justify-between mb-4">
      <h2 class="text-xl font-bold text-gray-800">📊 بازدیدهای اخیر کاربر (تحلیل بازاریابی)</h2>
      <button 
        @click="cleanupUnknownViews"
        :disabled="cleaning"
        class="px-4 py-2 bg-red-600 hover:bg-red-700 text-white text-sm rounded-lg disabled:bg-gray-400 transition-colors"
      >
        {{ cleaning ? '🔄 در حال پاکسازی...' : '🗑️ پاک کردن Unknown ها' }}
      </button>
    </div>
    
    <div v-if="loading" class="text-center py-8 text-gray-500">در حال بارگذاری...</div>
    
    <div v-else-if="error" class="text-center py-8 text-red-500">{{ error }}</div>
    
    <div v-else-if="!views || views.length === 0" class="text-center py-8 text-gray-500">
      هیچ بازدیدی ثبت نشده است
    </div>
    
    <div v-else>
      <!-- آمار کلی -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-6 p-6 bg-blue-50 rounded-lg">
        <div class="text-center">
          <div class="text-2xl font-bold text-blue-600">{{ analytics.count }}</div>
          <div class="text-sm text-gray-600">تعداد محصولات بازدید شده</div>
        </div>
        <div class="text-center">
          <div class="text-2xl font-bold text-green-600">{{ formatDuration(analytics.total_duration) }}</div>
          <div class="text-sm text-gray-600">مجموع زمان بازدید</div>
        </div>
        <div class="text-center">
          <div class="text-2xl font-bold text-purple-600">{{ formatDuration(analytics.avg_duration) }}</div>
          <div class="text-sm text-gray-600">میانگین زمان بازدید</div>
        </div>
      </div>

      <!-- جدول بازدیدها -->
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr class="text-right">
              <th class="px-4 py-3 text-xs font-medium text-gray-500 uppercase">محصول</th>
              <th class="px-4 py-3 text-xs font-medium text-gray-500 uppercase">دستگاه</th>
              <th class="px-4 py-3 text-xs font-medium text-gray-500 uppercase">مرورگر</th>
              <th class="px-4 py-3 text-xs font-medium text-gray-500 uppercase">اولین بازدید</th>
              <th class="px-4 py-3 text-xs font-medium text-gray-500 uppercase">آخرین بازدید</th>
              <th class="px-4 py-3 text-xs font-medium text-gray-500 uppercase">تعداد بازدید</th>
              <th class="px-4 py-3 text-xs font-medium text-gray-500 uppercase">مدت زمان</th>
              <th class="px-4 py-3 text-xs font-medium text-gray-500 uppercase">علاقه</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="view in views" :key="view.id" class="hover:bg-gray-50">
              <td class="px-4 py-3">
                <div class="flex items-center gap-3">
                  <img 
                    :src="view.product.image_url || view.product.image || '/default-product.svg'" 
                    alt="محصول"
                    class="w-12 h-12 object-cover rounded"
                  />
                  <div class="text-sm">
                    <div class="font-medium text-gray-900">{{ view.product.name }}</div>
                    <div class="text-gray-500">کد: {{ view.product.id }}</div>
                  </div>
                </div>
              </td>
              <td class="px-4 py-3 text-sm">
                <div class="flex items-center gap-2">
                  <span v-if="view.device_type === 'Mobile'">📱</span>
                  <span v-else-if="view.device_type === 'Tablet'">📲</span>
                  <span v-else>💻</span>
                  <div>
                    <div class="font-medium text-gray-900">{{ view.device_model || view.device_type }}</div>
                    <div class="text-xs text-gray-500">{{ view.os }} {{ view.os_version }}</div>
                    <div class="text-xs text-gray-400" dir="ltr">{{ view.ip_address }}</div>
                  </div>
                </div>
              </td>
              <td class="px-4 py-3 text-sm text-gray-700">
                <div>{{ view.browser }}</div>
                <div class="text-xs text-gray-500">{{ view.browser_version }}</div>
              </td>
              <td class="px-4 py-3 text-sm text-gray-700" dir="ltr">
                <div>{{ formatDate(view.viewed_at) }}</div>
                <div class="text-xs text-gray-500">{{ formatTime(view.viewed_at) }}</div>
              </td>
              <td class="px-4 py-3 text-sm text-gray-700" dir="ltr">
                <div>{{ formatDate(view.last_updated_at) }}</div>
                <div class="text-xs text-gray-500">{{ formatTime(view.last_updated_at) }}</div>
              </td>
              <td class="px-4 py-3 text-center">
                <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800">
                  {{ view.view_count }}x
                </span>
              </td>
              <td class="px-4 py-3 text-sm text-gray-700">
                {{ formatDuration(view.duration_seconds) }}
              </td>
              <td class="px-4 py-3 text-center">
                <span 
                  class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                  :class="getInterestLevel(view.duration_seconds, view.view_count).class"
                >
                  {{ getInterestLevel(view.duration_seconds, view.view_count).label }}
                </span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- راهنمای سطح علاقه -->
      <div class="mt-6 p-6 bg-gray-50 rounded-lg">
        <h3 class="text-sm font-semibold mb-2">راهنمای سطح علاقه (برای بازاریابی):</h3>
        <div class="grid grid-cols-1 md:grid-cols-4 gap-2 text-xs">
          <div class="flex items-center gap-2">
            <span class="px-2 py-1 rounded bg-red-100 text-red-800">سرد</span>
            <span class="text-gray-600">کمتر از 10 ثانیه</span>
          </div>
          <div class="flex items-center gap-2">
            <span class="px-2 py-1 rounded bg-yellow-100 text-yellow-800">متوسط</span>
            <span class="text-gray-600">10-30 ثانیه</span>
          </div>
          <div class="flex items-center gap-2">
            <span class="px-2 py-1 rounded bg-green-100 text-green-800">گرم</span>
            <span class="text-gray-600">30-60 ثانیه</span>
          </div>
          <div class="flex items-center gap-2">
            <span class="px-2 py-1 rounded bg-purple-100 text-purple-800">داغ</span>
            <span class="text-gray-600">بیش از 60 ثانیه یا 2+ بازدید</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const $fetch: <T = unknown>(url: string, options?: { credentials?: string; method?: string; body?: unknown }) => Promise<T>
</script>

<script setup lang="ts">
import { ref, onMounted } from 'vue'

const props = defineProps<{
  user: any
}>()

const views = ref<any[]>([])
const analytics = ref<any>({ count: 0, total_duration: 0, avg_duration: 0 })
const loading = ref(true)
const error = ref('')
const cleaning = ref(false)

onMounted(async () => {
  await fetchRecentViews()
})

async function cleanupUnknownViews() {
  if (!confirm('آیا مطمئن هستید که می‌خواهید تمام بازدیدهای Unknown را پاک کنید؟\nاین عملیات برگشت‌پذیر نیست!')) {
    return
  }
  
  cleaning.value = true
  
  try {
    const response: any = await $fetch('/api/admin/recent-views/cleanup-unknown', {
      method: 'DELETE',
      credentials: 'include'
    })
    
    console.log('🔍 Response:', response) // برای debug
    
    const deletedCount = response?.deleted_count || response?.data?.deleted_count || 0
    alert(`✅ ${deletedCount} بازدید قدیمی با موفقیت پاک شد`)
    
    // بروزرسانی لیست
    await fetchRecentViews()
  } catch (err: any) {
    console.error('خطا در پاکسازی:', err)
    alert('❌ خطا در پاکسازی بازدیدهای قدیمی')
  } finally {
    cleaning.value = false
  }
}

async function fetchRecentViews() {
  if (!props.user?.id) return
  
  loading.value = true
  error.value = ''
  
  try {
    const response = await $fetch(`/api/admin/recent-views/user/${props.user.id}?limit=50`, {
      credentials: 'include'
    }) as { data?: unknown[]; count?: number; total_duration?: number; avg_duration?: number }
    
    views.value = (response.data || []) as any[]
    analytics.value = {
      count: response.count || 0,
      total_duration: response.total_duration || 0,
      avg_duration: response.avg_duration || 0
    }
  } catch (err: any) {
    console.error('خطا در دریافت بازدیدها:', err)
    error.value = 'خطا در دریافت بازدیدهای کاربر'
  } finally {
    loading.value = false
  }
}

function formatDate(dateStr: string) {
  if (!dateStr) return '-'
  const d = new Date(dateStr)
  return d.toLocaleDateString('fa-IR')
}

function formatTime(dateStr: string) {
  if (!dateStr) return '-'
  const d = new Date(dateStr)
  return d.toLocaleTimeString('fa-IR', { hour: '2-digit', minute: '2-digit' })
}

function formatDuration(seconds: number) {
  if (!seconds || seconds === 0) return '0 ثانیه'
  
  const hours = Math.floor(seconds / 3600)
  const minutes = Math.floor((seconds % 3600) / 60)
  const secs = seconds % 60
  
  const parts = []
  if (hours > 0) parts.push(`${hours} ساعت`)
  if (minutes > 0) parts.push(`${minutes} دقیقه`)
  if (secs > 0 || parts.length === 0) parts.push(`${secs} ثانیه`)
  
  return parts.join(' و ')
}

function getInterestLevel(duration: number, viewCount: number) {
  // محاسبه Engagement Score
  const score = duration + (viewCount * 30) // هر بازدید اضافی = 30 ثانیه امتیاز
  
  if (score >= 60 || viewCount >= 2) {
    return { label: '🔥 داغ', class: 'bg-purple-100 text-purple-800' }
  } else if (score >= 30) {
    return { label: '🟢 گرم', class: 'bg-green-100 text-green-800' }
  } else if (score >= 10) {
    return { label: '🟡 متوسط', class: 'bg-yellow-100 text-yellow-800' }
  } else {
    return { label: '🔵 سرد', class: 'bg-red-100 text-red-800' }
  }
}
</script>

