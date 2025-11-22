<template>
  <div class="bg-white rounded-xl shadow-lg p-6">
    <h3 class="text-lg font-bold text-gray-900 mb-4">خروجی و گزارش‌گیری</h3>
    
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <!-- Excel Export -->
      <button
        class="flex flex-col items-center p-6 border-2 border-dashed border-green-300 rounded-lg hover:border-green-500 hover:bg-green-50 transition-all duration-200 group"
        @click="exportExcel"
      >
        <div class="w-12 h-12 bg-green-100 rounded-lg flex items-center justify-center mb-3 group-hover:bg-green-200 transition-colors">
          <svg class="w-6 h-6 text-green-600" fill="currentColor" viewBox="0 0 20 20">
            <path fill-rule="evenodd" d="M3 17a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm3.293-7.707a1 1 0 011.414 0L9 10.586V3a1 1 0 112 0v7.586l1.293-1.293a1 1 0 111.414 1.414l-3 3a1 1 0 01-1.414 0l-3-3a1 1 0 010-1.414z" clip-rule="evenodd" />
          </svg>
        </div>
        <div class="text-sm font-semibold text-gray-900 mb-1">فایل اکسل</div>
        <div class="text-xs text-gray-500 text-center">خروجی کامل در قالب XLSX</div>
      </button>

      <!-- PDF Export -->
      <button
        class="flex flex-col items-center p-6 border-2 border-dashed border-red-300 rounded-lg hover:border-red-500 hover:bg-red-50 transition-all duration-200 group"
        @click="exportPDF"
      >
        <div class="w-12 h-12 bg-red-100 rounded-lg flex items-center justify-center mb-3 group-hover:bg-red-200 transition-colors">
          <svg class="w-6 h-6 text-red-600" fill="currentColor" viewBox="0 0 20 20">
            <path fill-rule="evenodd" d="M4 4a2 2 0 012-2h4.586A2 2 0 0112 2.586L15.414 6A2 2 0 0116 7.414V16a2 2 0 01-2 2H6a2 2 0 01-2-2V4zm2 6a1 1 0 011-1h6a1 1 0 110 2H7a1 1 0 01-1-1zm1 3a1 1 0 100 2h6a1 1 0 100-2H7z" clip-rule="evenodd" />
          </svg>
        </div>
        <div class="text-sm font-semibold text-gray-900 mb-1">گزارش PDF</div>
        <div class="text-xs text-gray-500 text-center">گزارش تحلیلی کامل</div>
      </button>

      <!-- CSV Export -->
      <button
        class="flex flex-col items-center p-6 border-2 border-dashed border-blue-300 rounded-lg hover:border-blue-500 hover:bg-blue-50 transition-all duration-200 group"
        @click="exportCSV"
      >
        <div class="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center mb-3 group-hover:bg-blue-200 transition-colors">
          <svg class="w-6 h-6 text-blue-600" fill="currentColor" viewBox="0 0 20 20">
            <path fill-rule="evenodd" d="M3 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm0 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm0 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm0 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z" clip-rule="evenodd" />
          </svg>
        </div>
        <div class="text-sm font-semibold text-gray-900 mb-1">فایل CSV</div>
        <div class="text-xs text-gray-500 text-center">داده خام برای تحلیل</div>
      </button>

      <!-- Email Report -->
      <button
        class="flex flex-col items-center p-6 border-2 border-dashed border-purple-300 rounded-lg hover:border-purple-500 hover:bg-purple-50 transition-all duration-200 group"
        @click="emailReport"
      >
        <div class="w-12 h-12 bg-purple-100 rounded-lg flex items-center justify-center mb-3 group-hover:bg-purple-200 transition-colors">
          <svg class="w-6 h-6 text-purple-600" fill="currentColor" viewBox="0 0 20 20">
            <path d="M2.003 5.884L10 9.882l7.997-3.998A2 2 0 0016 4H4a2 2 0 00-1.997 1.884z" />
            <path d="M18 8.118l-8 4-8-4V14a2 2 0 002 2h12a2 2 0 002-2V8.118z" />
          </svg>
        </div>
        <div class="text-sm font-semibold text-gray-900 mb-1">ارسال ایمیل</div>
        <div class="text-xs text-gray-500 text-center">ارسال گزارش به ایمیل</div>
      </button>
    </div>

    <!-- Export Options -->
    <div class="mt-6 p-6 bg-gray-50 rounded-lg">
      <h4 class="text-sm font-semibold text-gray-900 mb-3">تنظیمات خروجی</h4>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div>
          <label class="block text-xs text-gray-600 mb-1">بازه زمانی</label>
          <select class="w-full text-sm border border-gray-300 rounded px-2 py-1">
            <option>همه داده‌ها</option>
            <option>30 روز اخیر</option>
            <option>7 روز اخیر</option>
            <option>امروز</option>
          </select>
        </div>
        <div>
          <label class="block text-xs text-gray-600 mb-1">فرمت تاریخ</label>
          <select class="w-full text-sm border border-gray-300 rounded px-2 py-1">
            <option>شمسی</option>
            <option>میلادی</option>
          </select>
        </div>
        <div>
          <label class="block text-xs text-gray-600 mb-1">شامل نمودارها</label>
          <label class="flex items-center">
            <input type="checkbox" checked class="mr-2">
            <span class="text-sm text-gray-700">بله</span>
          </label>
        </div>
      </div>
    </div>

    <!-- Statistics Summary -->
    <div class="mt-6 grid grid-cols-2 md:grid-cols-4 gap-6">
      <div class="text-center">
        <div class="text-2xl font-bold text-blue-600">{{ visits.length.toLocaleString('fa-IR') }}</div>
        <div class="text-sm text-gray-500">رکورد کل</div>
      </div>
      <div class="text-center">
        <div class="text-2xl font-bold text-green-600">{{ uniqueVisitors.toLocaleString('fa-IR') }}</div>
        <div class="text-sm text-gray-500">بازدیدکننده یکتا</div>
      </div>
      <div class="text-center">
        <div class="text-2xl font-bold text-purple-600">{{ totalPages.toLocaleString('fa-IR') }}</div>
        <div class="text-sm text-gray-500">صفحه مختلف</div>
      </div>
      <div class="text-center">
        <div class="text-2xl font-bold text-orange-600">{{ avgDuration }}</div>
        <div class="text-sm text-gray-500">میانگین مدت</div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
</script>

<script setup lang="ts">
import { computed } from 'vue'

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

interface Visit {
  id: number
  timestamp: string
  ip: string
  city: string
  device: string
  browser: string
  page: string
  referrer: string
  duration: string
}

const props = defineProps<{
  visits: Visit[]
}>()

const uniqueVisitors = computed(() => {
  const uniqueIps = new Set(props.visits.map(v => v.ip))
  return uniqueIps.size
})

const totalPages = computed(() => {
  const uniquePages = new Set(props.visits.map(v => v.page))
  return uniquePages.size
})

const avgDuration = computed(() => {
  if (props.visits.length === 0) return '0:00'
  
  const totalSeconds = props.visits.reduce((sum, visit) => {
    const [minutes, seconds] = visit.duration.split(':').map(Number)
    return sum + (minutes * 60) + seconds
  }, 0)
  
  const avgSeconds = Math.round(totalSeconds / props.visits.length)
  const minutes = Math.floor(avgSeconds / 60)
  const seconds = avgSeconds % 60
  
  return `${minutes}:${seconds.toString().padStart(2, '0')}`
})

const exportExcel = () => {
  // TODO: Implement Excel export
}

const exportPDF = () => {
  // TODO: Implement PDF export
}

const exportCSV = () => {
  // TODO: Implement CSV export
}

const emailReport = () => {
  // TODO: Implement email functionality
}
</script> 
