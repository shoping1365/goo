<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
    <div class="flex items-center justify-between mb-6">
      <h3 class="text-lg font-semibold text-gray-900">صادرات داده‌ها</h3>
      <button @click="exportAll" class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700">
        صادرات همه
      </button>
    </div>

    <!-- Export Options -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-6">
      <div class="bg-gray-50 rounded-lg p-6">
        <div class="flex items-center mb-3">
          <svg class="w-8 h-8 text-green-600 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 17v-2m3 2v-4m3 4v-6m2 10H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
          </svg>
          <h4 class="text-lg font-semibold text-gray-900">Excel</h4>
        </div>
        <p class="text-sm text-gray-600 mb-4">صادرات داده‌ها در فرمت Excel برای تحلیل پیشرفته</p>
        <div class="space-y-2">
          <button @click="exportExcel('all')" class="w-full text-left px-3 py-2 text-sm bg-white border border-gray-300 rounded-md hover:bg-gray-50">
            همه درخواست‌ها
          </button>
          <button @click="exportExcel('approved')" class="w-full text-left px-3 py-2 text-sm bg-white border border-gray-300 rounded-md hover:bg-gray-50">
            فقط تایید شده‌ها
          </button>
          <button @click="exportExcel('pending')" class="w-full text-left px-3 py-2 text-sm bg-white border border-gray-300 rounded-md hover:bg-gray-50">
            در انتظار بررسی
          </button>
        </div>
      </div>

      <div class="bg-gray-50 rounded-lg p-6">
        <div class="flex items-center mb-3">
          <svg class="w-8 h-8 text-red-600 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
          </svg>
          <h4 class="text-lg font-semibold text-gray-900">PDF</h4>
        </div>
        <p class="text-sm text-gray-600 mb-4">گزارش‌های قابل چاپ و ارسال</p>
        <div class="space-y-2">
          <button @click="exportPDF('summary')" class="w-full text-left px-3 py-2 text-sm bg-white border border-gray-300 rounded-md hover:bg-gray-50">
            گزارش خلاصه
          </button>
          <button @click="exportPDF('detailed')" class="w-full text-left px-3 py-2 text-sm bg-white border border-gray-300 rounded-md hover:bg-gray-50">
            گزارش تفصیلی
          </button>
          <button @click="exportPDF('analytics')" class="w-full text-left px-3 py-2 text-sm bg-white border border-gray-300 rounded-md hover:bg-gray-50">
            گزارش تحلیلی
          </button>
        </div>
      </div>

      <div class="bg-gray-50 rounded-lg p-6">
        <div class="flex items-center mb-3">
          <svg class="w-8 h-8 text-blue-600 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21h10a2 2 0 002-2V9.414a1 1 0 00-.293-.707l-5.414-5.414A1 1 0 0012.586 3H7a2 2 0 00-2 2v14a2 2 0 002 2z" />
          </svg>
          <h4 class="text-lg font-semibold text-gray-900">CSV</h4>
        </div>
        <p class="text-sm text-gray-600 mb-4">فرمت ساده برای پردازش داده‌ها</p>
        <div class="space-y-2">
          <button @click="exportCSV('customers')" class="w-full text-left px-3 py-2 text-sm bg-white border border-gray-300 rounded-md hover:bg-gray-50">
            اطلاعات مشتریان
          </button>
          <button @click="exportCSV('transactions')" class="w-full text-left px-3 py-2 text-sm bg-white border border-gray-300 rounded-md hover:bg-gray-50">
            تراکنش‌ها
          </button>
          <button @click="exportCSV('products')" class="w-full text-left px-3 py-2 text-sm bg-white border border-gray-300 rounded-md hover:bg-gray-50">
            محصولات
          </button>
        </div>
      </div>
    </div>

    <!-- Export Filters -->
    <div class="mb-6">
      <h4 class="text-md font-semibold text-gray-900 mb-4">فیلترهای صادرات</h4>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">بازه زمانی</label>
          <select v-model="exportFilters.dateRange" class="w-full border border-gray-300 rounded-md px-3 py-2">
            <option value="last_7_days">7 روز گذشته</option>
            <option value="last_30_days">30 روز گذشته</option>
            <option value="last_90_days">90 روز گذشته</option>
            <option value="custom">سفارشی</option>
          </select>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
          <select v-model="exportFilters.status" class="w-full border border-gray-300 rounded-md px-3 py-2">
            <option value="all">همه</option>
            <option value="pending">در انتظار</option>
            <option value="approved">تایید شده</option>
            <option value="rejected">رد شده</option>
          </select>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">بازه مبلغ</label>
          <select v-model="exportFilters.amountRange" class="w-full border border-gray-300 rounded-md px-3 py-2">
            <option value="all">همه مبالغ</option>
            <option value="low">کمتر از 10 میلیون</option>
            <option value="medium">10 تا 50 میلیون</option>
            <option value="high">بیش از 50 میلیون</option>
          </select>
        </div>
      </div>
    </div>

    <!-- Custom Date Range -->
    <div v-if="exportFilters.dateRange === 'custom'" class="mb-6">
      <h4 class="text-md font-semibold text-gray-900 mb-4">بازه زمانی سفارشی</h4>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">از تاریخ</label>
          <input v-model="exportFilters.startDate" type="date" class="w-full border border-gray-300 rounded-md px-3 py-2">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تا تاریخ</label>
          <input v-model="exportFilters.endDate" type="date" class="w-full border border-gray-300 rounded-md px-3 py-2">
        </div>
      </div>
    </div>

    <!-- Export History -->
    <div class="mb-6">
      <h4 class="text-md font-semibold text-gray-900 mb-4">تاریخچه صادرات</h4>
      <div class="space-y-3">
        <div v-for="export_item in exportHistory" :key="export_item.id" class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
          <div class="flex items-center">
            <div class="w-10 h-10 bg-blue-100 rounded-full flex items-center justify-center mr-3">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
              </svg>
            </div>
            <div>
              <h5 class="text-md font-semibold text-gray-900">{{ export_item.name }}</h5>
              <p class="text-sm text-gray-600">{{ export_item.date }} | {{ export_item.format }} | {{ export_item.size }}</p>
            </div>
          </div>
          <div class="flex items-center space-x-2 space-x-reverse">
            <span class="px-2 py-1 text-xs rounded-full" :class="getStatusClass(export_item.status)">
              {{ getStatusLabel(export_item.status) }}
            </span>
            <button v-if="export_item.status === 'completed'" @click="downloadExport(export_item)" class="text-blue-600 hover:text-blue-800">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
              </svg>
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Scheduled Exports -->
    <div>
      <h4 class="text-md font-semibold text-gray-900 mb-4">صادرات زمان‌بندی شده</h4>
      <div class="bg-blue-50 rounded-lg p-6">
        <div class="flex items-center justify-between mb-4">
          <span class="text-md font-medium text-gray-900">گزارش هفتگی</span>
          <label class="flex items-center">
            <input v-model="scheduledExports.weekly" type="checkbox" class="rounded border-gray-300">
            <span class="mr-2 text-sm">فعال</span>
          </label>
        </div>
        
        <div class="flex items-center justify-between mb-4">
          <span class="text-md font-medium text-gray-900">گزارش ماهانه</span>
          <label class="flex items-center">
            <input v-model="scheduledExports.monthly" type="checkbox" class="rounded border-gray-300">
            <span class="mr-2 text-sm">فعال</span>
          </label>
        </div>
        
        <div class="flex items-center justify-between">
          <span class="text-md font-medium text-gray-900">گزارش ساله</span>
          <label class="flex items-center">
            <input v-model="scheduledExports.yearly" type="checkbox" class="rounded border-gray-300">
            <span class="mr-2 text-sm">فعال</span>
          </label>
        </div>
        
        <div class="mt-4">
          <button @click="saveScheduleSettings" class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700">
            ذخیره تنظیمات
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'

interface ExportItem {
  id: number
  name: string
  date: string
  format: string
  size: string
  status: 'pending' | 'processing' | 'completed' | 'failed'
  downloadUrl?: string
}

const exportFilters = ref({
  dateRange: 'last_30_days',
  status: 'all',
  amountRange: 'all',
  startDate: '',
  endDate: ''
})

const scheduledExports = ref({
  weekly: false,
  monthly: true,
  yearly: false
})

const exportHistory = ref<ExportItem[]>([
  {
    id: 1,
    name: 'گزارش اقساط ماهانه',
    date: '1403/08/15',
    format: 'Excel',
    size: '2.5 MB',
    status: 'completed',
    downloadUrl: '/exports/monthly-report.xlsx'
  },
  {
    id: 2,
    name: 'لیست مشتریان',
    date: '1403/08/10',
    format: 'CSV',
    size: '1.2 MB',
    status: 'completed',
    downloadUrl: '/exports/customers.csv'
  }
])

const exportExcel = async (type: string) => {
  try {
    interface ApiResponse {
      success?: boolean
      message?: string
    }
    const response = await $fetch<ApiResponse>('/api/admin/installment-payments/export/excel', {
      method: 'POST',
      body: {
        type,
        filters: exportFilters.value
      }
    })
    
    if (response.success) {
      alert('فایل Excel در حال آماده‌سازی است')
      // Add to export history
    }
  } catch (error) {
    console.error('خطا در صادرات Excel:', error)
    alert('خطا در صادرات فایل')
  }
}

const exportPDF = async (type: string) => {
  try {
    interface ApiResponse {
      success?: boolean
      message?: string
    }
    const response = await $fetch<ApiResponse>('/api/admin/installment-payments/export/pdf', {
      method: 'POST',
      body: {
        type,
        filters: exportFilters.value
      }
    })
    
    if (response.success) {
      alert('فایل PDF در حال آماده‌سازی است')
    }
  } catch (error) {
    console.error('خطا در صادرات PDF:', error)
    alert('خطا در صادرات فایل')
  }
}

const exportCSV = async (type: string) => {
  try {
    interface ApiResponse {
      success?: boolean
      message?: string
    }
    const response = await $fetch<ApiResponse>('/api/admin/installment-payments/export/csv', {
      method: 'POST',
      body: {
        type,
        filters: exportFilters.value
      }
    })
    
    if (response.success) {
      alert('فایل CSV در حال آماده‌سازی است')
    }
  } catch (error) {
    console.error('خطا در صادرات CSV:', error)
    alert('خطا در صادرات فایل')
  }
}

const exportAll = () => {
  // Export all data in multiple formats
  alert('شروع صادرات همه داده‌ها')
}

const downloadExport = (exportItem: ExportItem) => {
  if (exportItem.downloadUrl) {
    window.open(exportItem.downloadUrl, '_blank')
  }
}

const getStatusClass = (status: string): string => {
  switch (status) {
    case 'pending': return 'bg-yellow-100 text-yellow-800'
    case 'processing': return 'bg-blue-100 text-blue-800'
    case 'completed': return 'bg-green-100 text-green-800'
    case 'failed': return 'bg-red-100 text-red-800'
    default: return 'bg-gray-100 text-gray-800'
  }
}

const getStatusLabel = (status: string): string => {
  switch (status) {
    case 'pending': return 'در انتظار'
    case 'processing': return 'در حال پردازش'
    case 'completed': return 'تکمیل شده'
    case 'failed': return 'ناموفق'
    default: return 'نامشخص'
  }
}

const saveScheduleSettings = async () => {
  try {
    await $fetch('/api/admin/installment-payments/export/schedule', {
      method: 'PUT',
      body: scheduledExports.value
    })
    alert('تنظیمات زمان‌بندی ذخیره شد')
  } catch (error) {
    console.error('خطا در ذخیره تنظیمات:', error)
    alert('خطا در ذخیره تنظیمات')
  }
}

onMounted(() => {
  // Load export history and settings
})
</script>
