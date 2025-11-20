<script setup lang="ts">
import { ref, computed, onMounted, reactive } from 'vue'

// تعریف definePageMeta برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

interface SearchLog {
  id: string
  query: string
  timestamp: string
  userId?: string
  resultsCount: number
  clickedResults: number
  searchDuration: number
  filters?: Record<string, any>
  userAgent?: string
  ip?: string
}

interface SearchStats {
  totalSearches: number
  uniqueUsers: number
  avgResultsCount: number
  avgSearchDuration: number
  topQueries: Array<{ query: string; count: number }>
  searchesByDay: Array<{ date: string; count: number }>
  noResultQueries: Array<{ query: string; count: number }>
  popularFilters: Record<string, number>
}

import { useApiClient } from '@/utils/api'

const { api } = useApiClient()
const isLoading = ref(false)
const searchLogs = ref<SearchLog[]>([])
const stats = ref<SearchStats | null>(null)
const dateRange = reactive({
  from: new Date(Date.now() - 7 * 24 * 60 * 60 * 1000).toISOString().split('T')[0],
  to: new Date().toISOString().split('T')[0]
})

const activeTab = ref<'overview' | 'logs' | 'popular' | 'failed'>('overview')

// Pagination
const currentPage = ref(1)
const pageSize = ref(50)
const totalPages = computed(() => Math.ceil(searchLogs.value.length / pageSize.value))

const paginatedLogs = computed(() => {
  const start = (currentPage.value - 1) * pageSize.value
  const end = start + pageSize.value
  return searchLogs.value.slice(start, end)
})

// Fetch data
const fetchSearchAnalytics = async () => {
  isLoading.value = true
  try {
    // Use $fetch directly to Nuxt server API to avoid CORS
    const response = await $fetch<{ logs: SearchLog[]; stats: SearchStats }>(
      '/api/admin/analytics/search',
      {
        params: {
          from: dateRange.from,
          to: dateRange.to
        }
      }
    )
    
    if (response) {
      searchLogs.value = response.logs || []
      stats.value = response.stats || null
    }
  } catch (error) {
    console.error('Error fetching search analytics:', error)
    // Fallback to mock data for development
    generateMockData()
  } finally {
    isLoading.value = false
  }
}

// Mock data for development
const generateMockData = () => {
  const queries = ['موبایل سامسونگ', 'لپ تاپ ایسوس', 'هدفون بی سیم', 'کیبورد مکانیکی', 'مانیتور گیمینگ']
  const mockLogs: SearchLog[] = []
  
  for (let i = 0; i < 100; i++) {
    // استفاده از crypto.getRandomValues برای تولید اعداد تصادفی امن در mock data
    const randomArray = new Uint32Array(8)
    crypto.getRandomValues(randomArray)
    const randomValues = Array.from(randomArray).map(val => val / 4294967296) // Normalize to 0-1
    
    const query = queries[Math.floor(randomValues[0] * queries.length)]
    
    mockLogs.push({
      id: `log_${i}`,
      query: query,
      timestamp: new Date(Date.now() - randomValues[1] * 7 * 24 * 60 * 60 * 1000).toISOString(),
      userId: randomValues[2] > 0.3 ? `user_${Math.floor(randomValues[3] * 50)}` : undefined,
      resultsCount: Math.floor(randomValues[4] * 100),
      clickedResults: Math.floor(randomValues[5] * 5),
      searchDuration: Math.floor(randomValues[6] * 500) + 50,
      filters: randomValues[7] > 0.5 ? { category: 'electronics', priceRange: '1000000-5000000' } : undefined,
      userAgent: 'Mozilla/5.0',
      ip: `192.168.1.${Math.floor(randomValues[0] * 255)}`
    })
  }
  
  searchLogs.value = mockLogs.sort((a, b) => new Date(b.timestamp).getTime() - new Date(a.timestamp).getTime())
  
  const queryCount: Record<string, number> = {}
  mockLogs.forEach(log => {
    queryCount[log.query] = (queryCount[log.query] || 0) + 1
  })
  
  stats.value = {
    totalSearches: mockLogs.length,
    uniqueUsers: new Set(mockLogs.filter(l => l.userId).map(l => l.userId)).size,
    avgResultsCount: mockLogs.reduce((sum, log) => sum + log.resultsCount, 0) / mockLogs.length,
    avgSearchDuration: mockLogs.reduce((sum, log) => sum + log.searchDuration, 0) / mockLogs.length,
    topQueries: Object.entries(queryCount)
      .map(([query, count]) => ({ query, count }))
      .sort((a, b) => b.count - a.count)
      .slice(0, 10),
    searchesByDay: [],
    noResultQueries: mockLogs
      .filter(log => log.resultsCount === 0)
      .reduce((acc, log) => {
        const existing = acc.find(item => item.query === log.query)
        if (existing) {
          existing.count++
        } else {
          acc.push({ query: log.query, count: 1 })
        }
        return acc
      }, [] as Array<{ query: string; count: number }>)
      .sort((a, b) => b.count - a.count)
      .slice(0, 10),
    popularFilters: {
      'category': 45,
      'price': 32,
      'brand': 28,
      'rating': 15
    }
  }
}

// Export to CSV
const exportToCSV = () => {
  const headers = ['تاریخ', 'جستجو', 'تعداد نتایج', 'کلیک‌ها', 'مدت زمان (ms)', 'کاربر']
  const rows = searchLogs.value.map(log => [
    new Date(log.timestamp).toLocaleString('fa-IR'),
    log.query,
    log.resultsCount,
    log.clickedResults,
    log.searchDuration,
    log.userId || 'مهمان'
  ])
  
  const csvContent = [
    headers.join(','),
    ...rows.map(row => row.join(','))
  ].join('\n')
  
  const blob = new Blob(['\ufeff' + csvContent], { type: 'text/csv;charset=utf-8;' })
  const link = document.createElement('a')
  link.href = URL.createObjectURL(blob)
  link.download = `search-analytics-${dateRange.from}-${dateRange.to}.csv`
  link.click()
}

// Format date
const formatDate = (dateStr: string) => {
  return new Date(dateStr).toLocaleString('fa-IR', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}

onMounted(() => {
  fetchSearchAnalytics()
})
</script>

<template>
  <div class="min-h-screen bg-gray-50 p-6">
    <!-- Header -->
    <div class="mb-8">
      <div class="flex justify-between items-center mb-4">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">گزارشات جستجو</h1>
          <p class="text-gray-600 mt-2">تحلیل و بررسی رفتار جستجوی کاربران</p>
        </div>
        <div class="flex gap-3">
          <button
            class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors flex items-center gap-2"
            @click="exportToCSV"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"/>
            </svg>
            خروجی CSV
          </button>
          <NuxtLink
            to="/admin/analytics/Reports"
            class="px-4 py-2 bg-gray-600 text-white rounded-lg hover:bg-gray-700 transition-colors"
          >
            بازگشت
          </NuxtLink>
        </div>
      </div>

      <!-- Date Range Filter -->
      <div class="bg-white rounded-lg shadow p-4 flex items-center gap-4">
        <label class="text-gray-700 font-medium">بازه زمانی:</label>
        <input
          v-model="dateRange.from"
          type="date"
          class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
        />
        <span class="text-gray-500">تا</span>
        <input
          v-model="dateRange.to"
          type="date"
          class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
        />
        <button
          :disabled="isLoading"
          class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors disabled:opacity-50"
          @click="fetchSearchAnalytics"
        >
          {{ isLoading ? 'در حال بارگذاری...' : 'اعمال فیلتر' }}
        </button>
      </div>
    </div>

    <!-- Stats Overview -->
    <div v-if="stats" class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
      <div class="bg-white rounded-lg shadow p-6">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-gray-600 text-sm">کل جستجوها</p>
            <p class="text-3xl font-bold text-gray-900 mt-2">{{ stats.totalSearches.toLocaleString('fa-IR') }}</p>
          </div>
          <div class="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"/>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow p-6">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-gray-600 text-sm">کاربران یکتا</p>
            <p class="text-3xl font-bold text-gray-900 mt-2">{{ stats.uniqueUsers.toLocaleString('fa-IR') }}</p>
          </div>
          <div class="w-12 h-12 bg-green-100 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z"/>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow p-6">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-gray-600 text-sm">میانگین نتایج</p>
            <p class="text-3xl font-bold text-gray-900 mt-2">{{ Math.round(stats.avgResultsCount).toLocaleString('fa-IR') }}</p>
          </div>
          <div class="w-12 h-12 bg-yellow-100 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"/>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow p-6">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-gray-600 text-sm">سرعت میانگین</p>
            <p class="text-3xl font-bold text-gray-900 mt-2">{{ Math.round(stats.avgSearchDuration).toLocaleString('fa-IR') }} ms</p>
          </div>
          <div class="w-12 h-12 bg-purple-100 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"/>
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- Tabs -->
    <div class="bg-white rounded-lg shadow mb-6">
      <div class="border-b border-gray-200">
        <nav class="flex -mb-px">
          <button
            :class="[
              'px-6 py-4 text-sm font-medium border-b-2 transition-colors',
              activeTab === 'overview'
                ? 'border-blue-600 text-blue-600'
                : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
            ]"
            @click="activeTab = 'overview'"
          >
            نمای کلی
          </button>
          <button
            :class="[
              'px-6 py-4 text-sm font-medium border-b-2 transition-colors',
              activeTab === 'logs'
                ? 'border-blue-600 text-blue-600'
                : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
            ]"
            @click="activeTab = 'logs'"
          >
            لاگ‌های جستجو
          </button>
          <button
            :class="[
              'px-6 py-4 text-sm font-medium border-b-2 transition-colors',
              activeTab === 'popular'
                ? 'border-blue-600 text-blue-600'
                : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
            ]"
            @click="activeTab = 'popular'"
          >
            جستجوهای محبوب
          </button>
          <button
            :class="[
              'px-6 py-4 text-sm font-medium border-b-2 transition-colors',
              activeTab === 'failed'
                ? 'border-blue-600 text-blue-600'
                : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
            ]"
            @click="activeTab = 'failed'"
          >
            جستجوهای ناموفق
          </button>
        </nav>
      </div>
    </div>

    <!-- Tab Content -->
    <div v-if="!isLoading">
      <!-- Overview Tab -->
      <div v-if="activeTab === 'overview' && stats" class="space-y-6">
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <!-- Top Queries -->
          <div class="bg-white rounded-lg shadow p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">پرجستجوترین کلمات</h3>
            <div class="space-y-3">
              <div
                v-for="(item, index) in stats.topQueries"
                :key="index"
                class="flex items-center justify-between p-3 bg-gray-50 rounded-lg"
              >
                <span class="text-gray-900 font-medium">{{ item.query }}</span>
                <span class="px-3 py-1 bg-blue-100 text-blue-700 rounded-full text-sm font-medium">
                  {{ item.count.toLocaleString('fa-IR') }}
                </span>
              </div>
            </div>
          </div>

          <!-- Popular Filters -->
          <div class="bg-white rounded-lg shadow p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">فیلترهای پرکاربرد</h3>
            <div class="space-y-3">
              <div
                v-for="(count, filter) in stats.popularFilters"
                :key="filter"
                class="flex items-center justify-between"
              >
                <span class="text-gray-700">{{ filter }}</span>
                <div class="flex items-center gap-2">
                  <div class="w-32 bg-gray-200 rounded-full h-2">
                    <div
                      class="bg-blue-600 h-2 rounded-full"
                      :style="{ width: `${(count / stats.totalSearches) * 100}%` }"
                    ></div>
                  </div>
                  <span class="text-sm text-gray-600 w-12 text-left">{{ count }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Logs Tab -->
      <div v-if="activeTab === 'logs'" class="bg-white rounded-lg shadow overflow-hidden">
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase">زمان</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase">جستجو</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase">نتایج</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase">کلیک‌ها</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase">مدت زمان</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase">کاربر</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="log in paginatedLogs" :key="log.id" class="hover:bg-gray-50">
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  {{ formatDate(log.timestamp) }}
                </td>
                <td class="px-6 py-4 text-sm text-gray-900 font-medium">
                  {{ log.query }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm">
                  <span
                    :class="[
                      'px-2 py-1 rounded-full text-xs font-medium',
                      log.resultsCount > 0 ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'
                    ]"
                  >
                    {{ log.resultsCount.toLocaleString('fa-IR') }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  {{ log.clickedResults.toLocaleString('fa-IR') }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-600">
                  {{ log.searchDuration.toLocaleString('fa-IR') }} ms
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-600">
                  {{ log.userId || 'مهمان' }}
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <!-- Pagination -->
        <div class="bg-gray-50 px-6 py-4 flex items-center justify-between border-t border-gray-200">
          <div class="text-sm text-gray-700">
            نمایش {{ ((currentPage - 1) * pageSize + 1).toLocaleString('fa-IR') }} 
            تا {{ Math.min(currentPage * pageSize, searchLogs.length).toLocaleString('fa-IR') }} 
            از {{ searchLogs.length.toLocaleString('fa-IR') }}
          </div>
          <div class="flex gap-2">
            <button
              :disabled="currentPage === 1"
              class="px-4 py-2 border border-gray-300 rounded-lg text-sm font-medium text-gray-700 hover:bg-gray-50 disabled:opacity-50"
              @click="currentPage--"
            >
              قبلی
            </button>
            <button
              :disabled="currentPage === totalPages"
              class="px-4 py-2 border border-gray-300 rounded-lg text-sm font-medium text-gray-700 hover:bg-gray-50 disabled:opacity-50"
              @click="currentPage++"
            >
              بعدی
            </button>
          </div>
        </div>
      </div>

      <!-- Popular Searches Tab -->
      <div v-if="activeTab === 'popular' && stats" class="bg-white rounded-lg shadow p-6">
        <h3 class="text-xl font-semibold text-gray-900 mb-6">محبوب‌ترین جستجوها</h3>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div
            v-for="(item, index) in stats.topQueries"
            :key="index"
            class="flex items-center justify-between p-4 border border-gray-200 rounded-lg hover:border-blue-300 transition-colors"
          >
            <div class="flex items-center gap-3">
              <span class="text-2xl font-bold text-gray-300">{{ (index + 1).toLocaleString('fa-IR') }}</span>
              <span class="text-gray-900 font-medium">{{ item.query }}</span>
            </div>
            <div class="text-left">
              <div class="text-2xl font-bold text-blue-600">{{ item.count.toLocaleString('fa-IR') }}</div>
              <div class="text-xs text-gray-500">جستجو</div>
            </div>
          </div>
        </div>
      </div>

      <!-- Failed Searches Tab -->
      <div v-if="activeTab === 'failed' && stats" class="bg-white rounded-lg shadow p-6">
        <h3 class="text-xl font-semibold text-gray-900 mb-6">جستجوهای بدون نتیجه</h3>
        <div class="space-y-4">
          <div
            v-for="(item, index) in stats.noResultQueries"
            :key="index"
            class="flex items-center justify-between p-4 bg-red-50 border border-red-200 rounded-lg"
          >
            <div class="flex items-center gap-3">
              <svg class="w-5 h-5 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
              </svg>
              <span class="text-gray-900 font-medium">{{ item.query }}</span>
            </div>
            <span class="px-3 py-1 bg-red-100 text-red-700 rounded-full text-sm font-medium">
              {{ item.count.toLocaleString('fa-IR') }} بار
            </span>
          </div>
        </div>
        <div v-if="stats.noResultQueries.length === 0" class="text-center py-12">
          <svg class="w-16 h-16 text-green-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
          </svg>
          <p class="text-gray-600">همه جستجوها نتیجه داشتند! 🎉</p>
        </div>
      </div>
    </div>

    <!-- Loading State -->
    <div v-else class="bg-white rounded-lg shadow p-12 text-center">
      <div class="inline-block animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
      <p class="mt-4 text-gray-600">در حال بارگذاری داده‌ها...</p>
    </div>
  </div>
</template>

<style scoped>
/* Custom scrollbar for table */
.overflow-x-auto::-webkit-scrollbar {
  height: 8px;
}

.overflow-x-auto::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 10px;
}

.overflow-x-auto::-webkit-scrollbar-thumb {
  background: #888;
  border-radius: 10px;
}

.overflow-x-auto::-webkit-scrollbar-thumb:hover {
  background: #555;
}
</style>
