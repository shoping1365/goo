<template>
  <div class="min-h-screen bg-gray-50 p-6">
    <!-- Header -->
    <div class="mb-8">
      <h1 class="text-3xl font-bold text-gray-900">مراقبت سیستم</h1>
      <p class="mt-2 text-gray-600">نظارت بر عملکرد و رویدادهای سیستم</p>
    </div>

    <!-- Health Status -->
    <div class="mb-6 rounded-lg bg-white p-6 shadow">
      <h2 class="text-xl font-bold text-gray-900 mb-4">وضعیت سیستم</h2>
      
      <div v-if="performanceHealth" class="grid grid-cols-1 md:grid-cols-2 gap-4">
        <!-- Status Badge -->
        <div class="flex items-center gap-4">
          <div :class="['w-4 h-4 rounded-full', getHealthStatusColor(performanceHealth.status)]"></div>
          <div>
            <p class="text-sm text-gray-600">وضعیت</p>
            <p class="font-bold">
              {{ performanceHealth.status === 'healthy' ? 'سالم' : performanceHealth.status === 'warning' ? 'اخطار' : 'بحرانی' }}
            </p>
          </div>
        </div>

        <!-- Issues -->
        <div v-if="performanceHealth.issues.length > 0">
          <p class="text-sm text-gray-600 mb-2">مسائل شناسایی شده:</p>
          <div class="space-y-1">
            <p v-for="(issue, idx) in performanceHealth.issues" :key="idx" class="text-sm text-red-600">
              • {{ issue }}
            </p>
          </div>
        </div>
      </div>

      <div v-else class="text-gray-500">بارگیری...</div>
    </div>

    <!-- Performance Metrics -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-6">
      <!-- Response Time -->
      <div class="rounded-lg bg-white p-6 shadow">
        <h3 class="text-lg font-bold text-gray-900 mb-4">زمان پاسخ</h3>
        <div v-if="performance" class="space-y-2">
          <div class="flex justify-between">
            <span class="text-gray-600">میانگین:</span>
            <span class="font-bold">{{ formatDuration(performance.avgDuration) }}</span>
          </div>
        </div>
      </div>

      <!-- Error Rate -->
      <div class="rounded-lg bg-white p-6 shadow">
        <h3 class="text-lg font-bold text-gray-900 mb-4">درصد خطا</h3>
        <div v-if="performance" class="space-y-2">
          <div class="flex justify-between">
            <span class="text-gray-600">درصد:</span>
            <span class="font-bold" :class="performance.errorRate > 5 ? 'text-red-600' : 'text-green-600'">
              {{ performance.errorRate.toFixed(2) }}%
            </span>
          </div>
        </div>
      </div>
    </div>

    <!-- Error Statistics -->
    <div class="rounded-lg bg-white p-6 shadow mb-6">
      <h2 class="text-xl font-bold text-gray-900 mb-4">آمار خطاها</h2>
      
      <div v-if="errorStats" class="grid grid-cols-2 md:grid-cols-4 gap-4">
        <div class="text-center p-4 bg-gray-50 rounded">
          <p class="text-sm text-gray-600">کل</p>
          <p class="text-2xl font-bold">{{ errorStats.total }}</p>
        </div>
        <div class="text-center p-4 bg-red-50 rounded">
          <p class="text-sm text-gray-600">بحرانی</p>
          <p class="text-2xl font-bold text-red-600">{{ errorStats.bySeverity?.critical || 0 }}</p>
        </div>
        <div class="text-center p-4 bg-orange-50 rounded">
          <p class="text-sm text-gray-600">بالا</p>
          <p class="text-2xl font-bold text-orange-600">{{ errorStats.bySeverity?.high || 0 }}</p>
        </div>
        <div class="text-center p-4 bg-green-50 rounded">
          <p class="text-sm text-gray-600">حل شده</p>
          <p class="text-2xl font-bold text-green-600">{{ errorStats.resolved || 0 }}</p>
        </div>
      </div>
    </div>

    <!-- Recent Errors -->
    <div class="rounded-lg bg-white p-6 shadow mb-6">
      <h2 class="text-xl font-bold text-gray-900 mb-4">خطاهای اخیر</h2>
      
      <div v-if="errors.length > 0" class="space-y-2">
        <div v-for="(err, idx) in errors.slice(0, 10)" :key="idx" class="p-4 border border-gray-200 rounded">
          <div class="flex items-start justify-between">
            <div>
              <p class="font-bold">{{ err.error_type }}</p>
              <p class="text-sm text-gray-600">{{ err.message }}</p>
              <p class="text-xs text-gray-500 mt-1">{{ new Date(err.timestamp).toLocaleString('fa-IR') }}</p>
            </div>
            <span :class="['px-3 py-1 rounded text-white text-sm', getErrorSeverityColor(err.severity)]">
              {{ err.severity }}
            </span>
          </div>
        </div>
      </div>
      <div v-else class="text-gray-500">خطایی ثبت نشده است</div>
    </div>

    <!-- Endpoint Performance -->
    <div class="rounded-lg bg-white p-6 shadow">
      <h2 class="text-xl font-bold text-gray-900 mb-4">عملکرد نقاط پایانی</h2>
      
      <div v-if="endpoints.length > 0" class="overflow-x-auto">
        <table class="w-full text-sm">
          <thead class="bg-gray-50 border-b">
            <tr>
              <th class="px-4 py-2 text-right">نقطه پایانی</th>
              <th class="px-4 py-2 text-right">روش</th>
              <th class="px-4 py-2 text-right">متوسط</th>
              <th class="px-4 py-2 text-right">درخواست‌ها</th>
              <th class="px-4 py-2 text-right">خطاها</th>
              <th class="px-4 py-2 text-right">موفقیت</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="ep in endpoints.slice(0, 15)" :key="ep.endpoint" class="border-b hover:bg-gray-50">
              <td class="px-4 py-2">{{ ep.endpoint }}</td>
              <td class="px-4 py-2">{{ ep.method }}</td>
              <td class="px-4 py-2">{{ formatDuration(ep.avgDuration) }}</td>
              <td class="px-4 py-2">{{ ep.totalRequests }}</td>
              <td class="px-4 py-2">{{ ep.errorCount }}</td>
              <td class="px-4 py-2">
                <span :class="ep.successRate > 95 ? 'text-green-600' : 'text-orange-600'">
                  {{ ep.successRate.toFixed(1) }}%
                </span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <div v-else class="text-gray-500">داده‌ای موجود نیست</div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center">
      <div class="bg-white rounded-lg p-6">
        <p class="text-gray-600">بارگیری...</p>
      </div>
    </div>

    <!-- Error Message -->
    <div v-if="error" class="mt-4 p-4 bg-red-50 border border-red-200 rounded text-red-600">
      {{ error }}
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[]; title?: string }) => void
</script>

<script setup lang="ts">
import { onMounted } from 'vue'
import { useMonitoring } from '~/composables/useMonitoring'

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin',
  title: 'مراقبت سیستم',
})

const monitoring = useMonitoring()

const {
  errorStats,
  performance,
  performanceHealth,
  errors,
  endpoints,
  loading,
  error,
  fetchDashboard,
  getErrorSeverityColor,
  getHealthStatusColor,
  formatDuration,
} = monitoring

onMounted(async () => {
  await fetchDashboard()

  // Refresh every 30 seconds
  setInterval(fetchDashboard, 30000)
})
</script>
