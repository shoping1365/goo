<template>
  <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
    <!-- هدر -->
    <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-6 mb-6">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">گزارش‌گیری و نظارت</h3>
        <p class="text-sm text-gray-600">گزارش‌های جامع و نظارت بر عملکرد اتصالات حسابداری</p>
      </div>
      
      <!-- دکمه‌های عملیاتی -->
      <div class="flex flex-wrap gap-3">
        <button 
          @click="generateReport"
          class="inline-flex items-center gap-2 px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg transition-colors duration-200"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
          </svg>
          تولید گزارش
        </button>
        
        <button 
          @click="exportReport"
          class="inline-flex items-center gap-2 px-4 py-2 bg-green-600 hover:bg-green-700 text-white rounded-lg transition-colors duration-200"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
          </svg>
          صادر کردن گزارش
        </button>
        
        <button 
          @click="showReportSettings = true"
          class="inline-flex items-center gap-2 px-4 py-2 bg-gray-600 hover:bg-gray-700 text-white rounded-lg transition-colors duration-200"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"></path>
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
          </svg>
          تنظیمات گزارش
        </button>
      </div>
    </div>

    <!-- آمار کلی -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-6">
      <!-- تعداد گزارش‌ها -->
      <div class="bg-blue-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-blue-600">{{ formatNumber(reportStats.totalReports) }}</div>
            <div class="text-sm text-blue-700">کل گزارش‌ها</div>
          </div>
          <div class="w-10 h-10 bg-blue-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
          </div>
        </div>
        <div class="mt-2 text-xs text-blue-600">
          این ماه: {{ formatNumber(reportStats.monthlyReports) }}
        </div>
      </div>

      <!-- گزارش‌های خودکار -->
      <div class="bg-green-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-green-600">{{ formatNumber(reportStats.autoReports) }}</div>
            <div class="text-sm text-green-700">گزارش خودکار</div>
          </div>
          <div class="w-10 h-10 bg-green-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
            </svg>
          </div>
        </div>
        <div class="mt-2 text-xs text-green-600">
          {{ reportStats.autoReportRate }}% از کل
        </div>
      </div>

      <!-- گزارش‌های دستی -->
      <div class="bg-purple-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-purple-600">{{ formatNumber(reportStats.manualReports) }}</div>
            <div class="text-sm text-purple-700">گزارش دستی</div>
          </div>
          <div class="w-10 h-10 bg-purple-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6V4m0 2a2 2 0 100 4m0-4a2 2 0 110 4m-6 8a2 2 0 100-4m0 4a2 2 0 100 4m0-4v2m0-6V4m6 6v10m6-2a2 2 0 100-4m0 4a2 2 0 100 4m0-4v2m0-6V4"></path>
            </svg>
          </div>
        </div>
        <div class="mt-2 text-xs text-purple-600">
          {{ reportStats.manualReportRate }}% از کل
        </div>
      </div>

      <!-- گزارش‌های خطا -->
      <div class="bg-red-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-red-600">{{ formatNumber(reportStats.errorReports) }}</div>
            <div class="text-sm text-red-700">گزارش خطا</div>
          </div>
          <div class="w-10 h-10 bg-red-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
        <div class="mt-2 text-xs text-red-600">
          {{ reportStats.errorReportRate }}% نرخ خطا
        </div>
      </div>
    </div>

    <!-- نمودارها و گزارش‌ها -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-6">
      <!-- نمودار روند گزارش‌ها -->
      <div class="bg-gray-50 rounded-lg p-6">
        <h4 class="font-medium text-gray-900 mb-4">روند گزارش‌ها (7 روز گذشته)</h4>
        <div class="h-64 flex items-end justify-between gap-2">
          <div 
            v-for="(day, index) in reportTrend" 
            :key="index"
            class="flex-1 bg-blue-500 rounded-t"
            :style="{ height: (day.count / maxReportCount) * 100 + '%' }"
          >
            <div class="text-center text-xs text-white mt-1">{{ day.count }}</div>
          </div>
        </div>
        <div class="flex justify-between text-xs text-gray-600 mt-2">
          <span v-for="(day, index) in reportTrend" :key="index">{{ day.date }}</span>
        </div>
      </div>

      <!-- توزیع نوع گزارش‌ها -->
      <div class="bg-gray-50 rounded-lg p-6">
        <h4 class="font-medium text-gray-900 mb-4">توزیع نوع گزارش‌ها</h4>
        <div class="space-y-3">
          <div v-for="type in reportTypes" :key="type.name" class="flex items-center justify-between">
            <div class="flex items-center gap-2">
              <div class="w-3 h-3 rounded-full" :style="{ backgroundColor: type.color }"></div>
              <span class="text-sm text-gray-700">{{ type.name }}</span>
            </div>
            <div class="flex items-center gap-2">
              <div class="w-24 bg-gray-200 rounded-full h-2">
                <div 
                  class="h-2 rounded-full"
                  :style="{ width: type.percentage + '%', backgroundColor: type.color }"
                ></div>
              </div>
              <span class="text-sm text-gray-600">{{ type.percentage }}%</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- جدول گزارش‌های اخیر -->
    <div>
      <h4 class="font-medium text-gray-900 mb-4">گزارش‌های اخیر</h4>
      <div class="overflow-x-auto">
        <table class="w-full text-sm">
          <thead>
            <tr class="border-b border-gray-200 bg-gray-50">
              <th class="text-right py-3 px-4 font-medium text-gray-600">نام گزارش</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">نوع</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">تاریخ تولید</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">حجم</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">وضعیت</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">عملیات</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="report in recentReports" :key="report.id" class="border-b border-gray-100 hover:bg-gray-50">
              <td class="py-3 px-4">
                <div>
                  <div class="font-medium text-gray-900">{{ report.name }}</div>
                  <div class="text-xs text-gray-500">{{ report.description }}</div>
                </div>
              </td>
              <td class="py-3 px-4">
                <span 
                  class="px-2 py-1 rounded-full text-xs font-medium"
                  :class="getReportTypeClass(report.type)"
                >
                  {{ getReportTypeLabel(report.type) }}
                </span>
              </td>
              <td class="py-3 px-4 text-gray-900">{{ report.generatedAt }}</td>
              <td class="py-3 px-4 text-gray-900">{{ report.size }}</td>
              <td class="py-3 px-4">
                <span 
                  class="px-2 py-1 rounded-full text-xs font-medium"
                  :class="getReportStatusClass(report.status)"
                >
                  {{ getReportStatusLabel(report.status) }}
                </span>
              </td>
              <td class="py-3 px-4">
                <div class="flex items-center gap-2">
                  <button 
                    @click="downloadReport(report)"
                    class="p-1 text-blue-600 hover:text-blue-800"
                    title="دانلود"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
                    </svg>
                  </button>
                  <button 
                    @click="viewReport(report)"
                    class="p-1 text-green-600 hover:text-green-800"
                    title="مشاهده"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                    </svg>
                  </button>
                  <button 
                    v-if="canDeleteReport"
                    @click="deleteReport(report)"
                    class="p-1 text-red-600 hover:text-red-800"
                    title="حذف"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
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

    <!-- مودال تنظیمات گزارش -->
    <div v-if="showReportSettings" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-xl p-6 w-full max-w-2xl mx-4 max-h-[90vh] overflow-y-auto">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">تنظیمات گزارش‌گیری</h3>
          <button @click="showReportSettings = false" class="text-gray-400 hover:text-gray-600">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>

        <div class="space-y-6">
          <!-- تنظیمات گزارش خودکار -->
          <div>
            <h4 class="font-medium text-gray-900 mb-3">گزارش‌های خودکار</h4>
            <div class="space-y-3">
              <label class="flex items-center">
                <input 
                  v-model="reportSettings.autoReports"
                  type="checkbox"
                  class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                />
                <span class="mr-2 text-sm text-gray-700">فعال کردن گزارش‌های خودکار</span>
              </label>
              
              <div v-if="reportSettings.autoReports" class="space-y-3">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">فرکانس گزارش‌ها</label>
                  <select 
                    v-model="reportSettings.reportFrequency"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                  >
                    <option value="daily">روزانه</option>
                    <option value="weekly">هفتگی</option>
                    <option value="monthly">ماهانه</option>
                  </select>
                </div>
                
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">فرمت گزارش</label>
                  <select 
                    v-model="reportSettings.reportFormat"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                  >
                    <option value="pdf">PDF</option>
                    <option value="excel">Excel</option>
                    <option value="csv">CSV</option>
                  </select>
                </div>
              </div>
            </div>
          </div>

          <!-- تنظیمات نگهداری -->
          <div>
            <h4 class="font-medium text-gray-900 mb-3">نگهداری گزارش‌ها</h4>
            <div class="space-y-3">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">مدت نگهداری (روز)</label>
                <input 
                  v-model.number="reportSettings.retentionDays"
                  type="number"
                  min="1"
                  max="365"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                />
              </div>
              
              <label class="flex items-center">
                <input 
                  v-model="reportSettings.autoCleanup"
                  type="checkbox"
                  class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                />
                <span class="mr-2 text-sm text-gray-700">پاک‌سازی خودکار گزارش‌های قدیمی</span>
              </label>
            </div>
          </div>

          <!-- دکمه‌ها -->
          <div class="flex gap-3 pt-4">
            <button 
              @click="saveReportSettings"
              class="flex-1 px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg transition-colors duration-200"
            >
              ذخیره تنظیمات
            </button>
            <button 
              @click="showReportSettings = false"
              class="flex-1 px-4 py-2 bg-gray-200 hover:bg-gray-300 text-gray-700 rounded-lg transition-colors duration-200"
            >
              انصراف
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'

const canDeleteReport = ref(false) // This will be set based on user permissions

// وضعیت مودال
const showReportSettings = ref(false)

// آمار گزارش‌ها
const reportStats = ref({
  totalReports: 1250,
  monthlyReports: 45,
  autoReports: 980,
  autoReportRate: 78.4,
  manualReports: 270,
  manualReportRate: 21.6,
  errorReports: 12,
  errorReportRate: 1.0
})

// تنظیمات گزارش
const reportSettings = ref({
  autoReports: true,
  reportFrequency: 'weekly',
  reportFormat: 'pdf',
  retentionDays: 30,
  autoCleanup: true
})

// روند گزارش‌ها
const reportTrend = ref([
  { date: 'شنبه', count: 15 },
  { date: 'یکشنبه', count: 22 },
  { date: 'دوشنبه', count: 18 },
  { date: 'سه‌شنبه', count: 25 },
  { date: 'چهارشنبه', count: 20 },
  { date: 'پنج‌شنبه', count: 16 },
  { date: 'جمعه', count: 12 }
])

// نوع گزارش‌ها
const reportTypes = ref([
  { name: 'گزارش فروش', percentage: 35, color: '#3B82F6' },
  { name: 'گزارش مالی', percentage: 25, color: '#10B981' },
  { name: 'گزارش موجودی', percentage: 20, color: '#F59E0B' },
  { name: 'گزارش مشتریان', percentage: 15, color: '#8B5CF6' },
  { name: 'سایر', percentage: 5, color: '#EF4444' }
])

// گزارش‌های اخیر
const recentReports = ref([
  {
    id: 1,
    name: 'گزارش فروش هفتگی',
    description: 'گزارش فروش هفته گذشته',
    type: 'auto',
    generatedAt: '2024-01-15 14:30',
    size: '2.5 MB',
    status: 'completed'
  },
  {
    id: 2,
    name: 'گزارش مالی ماهانه',
    description: 'گزارش مالی دسامبر 2024',
    type: 'manual',
    generatedAt: '2024-01-14 09:15',
    size: '5.2 MB',
    status: 'completed'
  },
  {
    id: 3,
    name: 'گزارش موجودی',
    description: 'وضعیت موجودی کالاها',
    type: 'auto',
    generatedAt: '2024-01-13 16:45',
    size: '1.8 MB',
    status: 'processing'
  }
])

// حداکثر تعداد گزارش
const maxReportCount = computed(() => {
  return Math.max(...reportTrend.value.map(day => day.count))
})

// فرمت اعداد
const formatNumber = (num: number): string => {
  return new Intl.NumberFormat('fa-IR').format(num)
}

// کلاس نوع گزارش
const getReportTypeClass = (type: string) => {
  const classes = {
    auto: 'bg-green-100 text-green-700',
    manual: 'bg-blue-100 text-blue-700',
    scheduled: 'bg-purple-100 text-purple-700'
  }
  return classes[type] || 'bg-gray-100 text-gray-700'
}

// برچسب نوع گزارش
const getReportTypeLabel = (type: string) => {
  const labels = {
    auto: 'خودکار',
    manual: 'دستی',
    scheduled: 'زمان‌بندی شده'
  }
  return labels[type] || type
}

// کلاس وضعیت گزارش
const getReportStatusClass = (status: string) => {
  const classes = {
    completed: 'bg-green-100 text-green-700',
    processing: 'bg-yellow-100 text-yellow-700',
    failed: 'bg-red-100 text-red-700',
    pending: 'bg-gray-100 text-gray-700'
  }
  return classes[status] || 'bg-gray-100 text-gray-700'
}

// برچسب وضعیت گزارش
const getReportStatusLabel = (status: string) => {
  const labels = {
    completed: 'تکمیل شده',
    processing: 'در حال پردازش',
    failed: 'ناموفق',
    pending: 'در انتظار'
  }
  return labels[status] || status
}

// تولید گزارش
const generateReport = async () => {
  try {
    // TODO: تولید گزارش
    console.log('تولید گزارش شروع شد')
  } catch (error) {
    console.error('خطا در تولید گزارش:', error)
  }
}

// صادر کردن گزارش
const exportReport = () => {
  // TODO: صادر کردن گزارش
  console.log('صادر کردن گزارش')
}

// ذخیره تنظیمات گزارش
const saveReportSettings = async () => {
  try {
    // TODO: ذخیره تنظیمات
    console.log('تنظیمات گزارش ذخیره شد:', reportSettings.value)
    showReportSettings.value = false
  } catch (error) {
    console.error('خطا در ذخیره تنظیمات:', error)
  }
}

// دانلود گزارش
const downloadReport = (report: any) => {
  // TODO: دانلود گزارش
  console.log('دانلود گزارش:', report)
}

// مشاهده گزارش
const viewReport = (report: any) => {
  // TODO: مشاهده گزارش
  console.log('مشاهده گزارش:', report)
}

// حذف گزارش
const deleteReport = async (report: any) => {
  if (confirm('آیا از حذف این گزارش اطمینان دارید؟')) {
    try {
      // TODO: حذف گزارش
      console.log('گزارش حذف شد:', report)
    } catch (error) {
      console.error('خطا در حذف گزارش:', error)
    }
  }
}

// بارگذاری اولیه
onMounted(() => {
  // TODO: بارگذاری آمار و گزارش‌ها از API
})

import { useAuth } from '~/composables/useAuth'

// تعریف definePageMeta برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void

// احراز هویت غیرفعال شده است - محدودیت دسترسی حذف شده
definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user, hasPermission } = useAuth()
</script>

<!--
  کامپوننت گزارش‌گیری و نظارت
  شامل:
  1. آمار کلی گزارش‌ها
  2. نمودار روند گزارش‌ها
  3. توزیع نوع گزارش‌ها
  4. جدول گزارش‌های اخیر
  5. تنظیمات گزارش‌گیری
  6. تولید و صادر کردن گزارش‌ها
  7. طراحی مدرن و کاملاً ریسپانسیو
--> 

