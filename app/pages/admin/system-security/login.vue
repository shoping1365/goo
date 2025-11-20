<template>
  <div class="p-6" dir="rtl">
    <div class="mb-6 bg-white p-6 rounded-lg shadow-md border border-gray-200">
      <h1 class="text-2xl font-bold text-gray-900 mb-2">ورود</h1>
      <p class="text-gray-600">نظارت بر ورودهای سیستم و تلاش‌های غیرمجاز</p>
    </div>

    <!-- آمار کلی -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
      <TemplateCard
        title="ورودهای موفق"
        :value="successfulLogins"
        variant="green"
      />
      <TemplateCard
        title="ورودهای ناموفق"
        :value="failedLogins"
        variant="red"
      />
      <TemplateCard
        title="تلاش‌های مشکوک"
        :value="suspiciousAttempts"
        variant="yellow"
      />
      <TemplateCard
        title="کاربران آنلاین"
        :value="onlineUsers"
        variant="cyan"
      />
    </div>

    <!-- فیلترها و جستجوی پیشرفته -->
    <div class="bg-gradient-to-br from-blue-50 to-cyan-50 shadow-lg rounded-2xl mb-8 border border-blue-100">
      <div class="px-6 py-4 border-b border-blue-100 flex items-center gap-2">
        <div class="w-8 h-8 bg-gradient-to-br from-purple-400 to-purple-600 rounded-lg flex items-center justify-center shadow-md">
          <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.707A1 1 0 013 7V4z"></path>
          </svg>
        </div>
        <h3 class="text-lg font-bold text-blue-900">فیلترها و جستجوی پیشرفته</h3>
      </div>
      <div class="p-6">
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-6 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نام کاربر</label>
            <input v-model="filters.user" type="text" placeholder="جستجو..." class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-white focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">شناسه کاربر</label>
            <input v-model="filters.userId" type="text" placeholder="شناسه..." class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-white focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع ورود</label>
            <select v-model="filters.type" class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-white focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm">
              <option value="">انتخاب نوع...</option>
              <option value="successful">موفق</option>
              <option value="failed">ناموفق</option>
              <option value="suspicious">مشکوک</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">IP آدرس</label>
            <input v-model="filters.ip" type="text" placeholder="IP آدرس..." class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-white focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
            <select v-model="filters.status" class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-white focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm">
              <option value="">همه</option>
              <option value="active">فعال</option>
              <option value="blocked">مسدود شده</option>
              <option value="suspicious">مشکوک</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ</label>
            <input v-model="filters.date" type="date" class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-white focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm">
          </div>
        </div>
      </div>
    </div>



    <!-- جدول ورودها -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-semibold text-gray-900">تاریخچه ورودها</h3>
      </div>
      <div class="overflow-x-auto overflow-y-auto custom-scrollbar" :style="{ maxHeight: `${itemsPerPage * 60 + 100}px` }">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کاربر</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">IP آدرس</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مرورگر</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع دستگاه</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">موقعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">زمان</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="login in paginatedLoginHistory" :key="login.id" :class="getLoginRowClass(login)">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="flex-shrink-0 h-8 w-8">
                    <div class="h-8 w-8 rounded-full bg-gray-300 flex items-center justify-center">
                      <span class="text-sm font-medium text-gray-700">{{ login.user.charAt(0) }}</span>
                    </div>
                  </div>
                  <div class="mr-3">
                    <div class="text-sm font-medium text-gray-900">{{ login.user }}</div>
                    <div class="text-sm text-gray-500">{{ login.email }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ login.ip }}</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getStatusClass(login.status)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ getStatusText(login.status) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ login.browser }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ login.device }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ login.location }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ login.time }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <div class="flex flex-col items-center space-y-1">
                  <button 
                    class="text-blue-600 hover:text-blue-900 transition-colors" 
                    title="مشاهده جزئیات"
                    @click="viewDetails(login.id)"
                  >
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                    </svg>
                  </button>
                  <button 
                    v-if="(login.status === 'failed' || login.status === 'suspicious') && !login.isBlocked" 
                    class="text-red-600 hover:text-red-900 transition-colors" 
                    title="مسدود کردن IP"
                    @click="blockIP(login)"
                  >
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"></path>
                    </svg>
                  </button>
                  <button 
                    v-if="login.isBlocked" 
                    class="text-green-600 hover:text-green-900 transition-colors" 
                    title="آزاد کردن IP"
                    @click="unblockIP(login)"
                  >
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.707A1 1 0 013 7V4z"></path>
                    </svg>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      
      <!-- کامپوننت صفحه‌بندی -->
      <AdminPagination 
        :current-page="currentPage"
        :total-pages="totalPages"
        :total="loginHistory.length"
        :items-per-page="itemsPerPage"
        @page-changed="handlePageChange"
        @items-per-page-changed="handleItemsPerPageChange"
      />
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
declare const useFetch: <T = unknown>(url: string, options?: { key?: string; server?: boolean; default?: () => T; swr?: boolean; query?: Record<string, string | number>; getCachedData?: (key: string) => unknown; immediate?: boolean }) => { data: { value: T }; error: { value: Error | null }; refresh: () => Promise<void>; execute: () => Promise<void> }
declare const useNuxtData: <T = unknown>(key: string) => { data: { value: T } }
declare const $fetch: <T = unknown>(url: string, options?: { method?: string; body?: unknown }) => Promise<T>
</script>

<script setup lang="ts">
import AdminPagination from '@/components/admin/common/Pagination.vue';
import TemplateCard from '@/components/common/TemplateCard.vue';
import { computed, onMounted, ref, watchEffect } from 'vue';

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// کارت‌های آمار
const successfulLogins = ref(0)
const failedLogins = ref(0)
const suspiciousAttempts = ref(0)
const onlineUsers = ref(0)

// صفحه‌بندی و فیلترها
const currentPage = ref(1)
const itemsPerPage = ref(10)
const totalCount = ref(0)

const filters = ref({
  type: '',
  user: '',
  userId: '',
  ip: '',
  status: '',
  date: ''
})

// داده‌های جدول
type LoginRow = {
  id: number
  user: string
  email: string
  ip: string
  status: 'successful' | 'failed' | 'suspicious'
  browser: string
  device: string
  location: string
  time: string
  isBlocked: boolean
}
const loginHistory = ref<LoginRow[]>([])

// fetch آمار
type LoginStats = {
  successfulLogins: number
  failedLogins: number
  suspiciousActivity: number
  onlineUsers: number
}

const { data: statsData, refresh: _refreshStats } = useFetch<LoginStats>('/api/admin/login-stats', {
  key: 'admin-login-stats',
  server: true,
  default: () => ({ successfulLogins: 0, failedLogins: 0, suspiciousActivity: 0, onlineUsers: 0 }),
  swr: true,
  // SWR
  getCachedData(key) {
    const cached = useNuxtData<LoginStats>(key).data.value
    return cached
  }
})

watchEffect(() => {
  if (statsData?.value) {
    successfulLogins.value = statsData.value.successfulLogins || 0
    failedLogins.value = statsData.value.failedLogins || 0
    suspiciousAttempts.value = statsData.value.suspiciousActivity || 0
    onlineUsers.value = statsData.value.onlineUsers || 0
  }
})

// fetch لیست تاریخچه
type LoginAttempt = {
  id?: number
  mobile?: string
  attempt_ip?: string
  is_successful?: boolean
  user_agent?: string
  created_at?: string
}

type LoginHistoryResponse = {
  attempts: LoginAttempt[]
  total: number
}

const fetchHistory = async () => {
  const query: Record<string, string | number> = {
    page: currentPage.value,
    limit: itemsPerPage.value,
  }
  if (filters.value.user) query.search = filters.value.user
  if (filters.value.type) query.method = filters.value.type
  if (filters.value.ip) query.ip = filters.value.ip
  if (filters.value.date) query.dateFrom = filters.value.date
  if (filters.value.status === 'successful') query.success = 'true'
  if (filters.value.status === 'failed') query.success = 'false'

  const { data, error, execute } = useFetch<LoginHistoryResponse>('/api/admin/login-history', {
    key: `admin-login-history-${currentPage.value}-${itemsPerPage.value}-${JSON.stringify(query)}`,
    server: true,
    query,
    default: () => ({ attempts: [], total: 0 }),
    swr: true,
    immediate: false,
    getCachedData(key) {
      const cached = useNuxtData<LoginHistoryResponse>(key).data.value
      return cached
    }
  })

  await execute()

  if (!error.value && data.value) {
    totalCount.value = data.value.total || 0
    // نگاشت به مدل UI
    loginHistory.value = (data.value.attempts || []).map((a: LoginAttempt, idx: number) => ({
      id: a.id || idx,
      user: a.mobile || '-',
      email: '',
      ip: a.attempt_ip || '-',
      status: a.is_successful ? 'successful' : 'failed',
      browser: a.user_agent ? String(a.user_agent).split(' ')[0] : '-',
      device: '-',
      location: '-',
      time: a.created_at || '-',
      isBlocked: false
    }))
  }
}

onMounted(async () => {
  await fetchHistory()
})

// صفحه‌بندی سمت سرور
const paginatedLoginHistory = computed(() => loginHistory.value)
const totalPages = computed(() => Math.max(1, Math.ceil(totalCount.value / itemsPerPage.value)))

function getStatusClass(status: 'successful' | 'failed' | 'suspicious' | string) {
  switch (status) {
    case 'successful':
      return 'bg-green-100 text-green-800'
    case 'failed':
      return 'bg-red-100 text-red-800'
    case 'suspicious':
      return 'bg-yellow-100 text-yellow-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

function getStatusText(status: 'successful' | 'failed' | 'suspicious' | string) {
  switch (status) {
    case 'successful':
      return 'موفق'
    case 'failed':
      return 'ناموفق'
    case 'suspicious':
      return 'مشکوک'
    default:
      return 'نامشخص'
  }
}

function getLoginRowClass(login: LoginRow) {
  if (login.isBlocked) {
    return 'bg-red-50 border-r-4 border-red-500'
  }
  if (login.status === 'suspicious') {
    return 'bg-yellow-50'
  }
  return ''
}

async function _applyFilters() {
  currentPage.value = 1
  await fetchHistory()
}

async function blockIP(login: LoginRow) {
  if (!login?.ip) return
  try {
    await $fetch<{ success?: boolean }>('/api/admin/traffic/block-ip', { method: 'POST', body: { ip: login.ip } })
    login.isBlocked = true
  } catch (e) {
    // خطا را ساکت مدیریت می‌کنیم
    console.error('Error blocking IP:', e)
  }
}

async function unblockIP(login: LoginRow) {
  if (!login?.ip) return
  try {
    await $fetch<{ success?: boolean }>(`/api/admin/traffic/unblock-ip/${encodeURIComponent(login.ip)}`, { method: 'DELETE' })
    login.isBlocked = false
  } catch (e) {
    // خطا را ساکت مدیریت می‌کنیم
    console.error('Error unblocking IP:', e)
  }
}

function viewDetails(id: number) {
  // منطق نمایش جزئیات
  alert(`جزئیات ورود ${id} نمایش داده می‌شود`)
}

// توابع صفحه‌بندی
async function handlePageChange(page: number) {
  currentPage.value = page
  await fetchHistory()
}

async function handleItemsPerPageChange(newItemsPerPage: number) {
  itemsPerPage.value = newItemsPerPage
  currentPage.value = 1 // بازگشت به صفحه اول
  await fetchHistory()
}
</script>

<style scoped>
.custom-scrollbar {
  scrollbar-width: thin;
  scrollbar-color: #cbd5e0 #f7fafc;
  direction: rtl;
}

.custom-scrollbar::-webkit-scrollbar {
  width: 8px;
  height: 8px;
}

.custom-scrollbar::-webkit-scrollbar-track {
  background: #f7fafc;
  border-radius: 4px;
}

.custom-scrollbar::-webkit-scrollbar-thumb {
  background: #cbd5e0;
  border-radius: 4px;
}

.custom-scrollbar::-webkit-scrollbar-thumb:hover {
  background: #a0aec0;
}

.custom-scrollbar::-webkit-scrollbar-corner {
  background: #f7fafc;
}

/* قرار دادن اسکرول در سمت چپ */
.custom-scrollbar {
  direction: ltr;
}

.custom-scrollbar > * {
  direction: rtl;
}
</style> 
