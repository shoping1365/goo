<template>
  <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
    <div class="flex items-center justify-between mb-6">
      <h2 class="text-xl font-bold text-gray-900">گزارش‌گیری</h2>
      <div class="flex items-center space-x-4 space-x-reverse">
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 flex items-center" @click="generateReport">
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 17v-2m3 2v-4m3 4v-6m2 10H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
          </svg>
          ایجاد گزارش
        </button>
      </div>
    </div>

    <!-- فیلترهای گزارش -->
    <div class="bg-gray-50 rounded-lg p-6 mb-6">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">نوع گزارش</label>
          <select v-model="reportType" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
            <option value="performance">عملکرد کلی</option>
            <option value="successful">تست‌های موفق</option>
            <option value="unsuccessful">تست‌های ناموفق</option>
            <option value="roi">ROI</option>
            <option value="custom">گزارش سفارشی</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">بازه زمانی</label>
          <select v-model="timeRange" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
            <option value="7">7 روز اخیر</option>
            <option value="30">30 روز اخیر</option>
            <option value="90">90 روز اخیر</option>
            <option value="365">یک سال اخیر</option>
            <option value="custom">سفارشی</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">نوع تست</label>
          <select v-model="testType" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
            <option value="">همه انواع</option>
            <option value="page">صفحه</option>
            <option value="button">دکمه</option>
            <option value="price">قیمت</option>
            <option value="image">تصویر</option>
            <option value="product">محصول</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">فرمت صادرات</label>
          <select v-model="exportFormat" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
            <option value="excel">Excel</option>
            <option value="pdf">PDF</option>
            <option value="csv">CSV</option>
          </select>
        </div>
      </div>
    </div>

    <!-- گزارش‌های آماده -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-6">
      <div class="bg-blue-50 rounded-lg p-6 cursor-pointer hover:bg-blue-100" @click="selectReport('performance')">
        <div class="flex items-center">
          <div class="p-2 bg-blue-100 rounded-lg">
            <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
            </svg>
          </div>
          <div class="mr-3">
            <h3 class="text-sm font-medium text-blue-900">عملکرد کلی</h3>
            <p class="text-xs text-blue-600">آمار کلی تست‌ها</p>
          </div>
        </div>
      </div>

      <div class="bg-green-50 rounded-lg p-6 cursor-pointer hover:bg-green-100" @click="selectReport('successful')">
        <div class="flex items-center">
          <div class="p-2 bg-green-100 rounded-lg">
            <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <div class="mr-3">
            <h3 class="text-sm font-medium text-green-900">تست‌های موفق</h3>
            <p class="text-xs text-green-600">تست‌هایی که بهبود نشان دادند</p>
          </div>
        </div>
      </div>

      <div class="bg-red-50 rounded-lg p-6 cursor-pointer hover:bg-red-100" @click="selectReport('unsuccessful')">
        <div class="flex items-center">
          <div class="p-2 bg-red-100 rounded-lg">
            <svg class="w-6 h-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <div class="mr-3">
            <h3 class="text-sm font-medium text-red-900">تست‌های ناموفق</h3>
            <p class="text-xs text-red-600">تست‌هایی که بهبود نشان ندادند</p>
          </div>
        </div>
      </div>

      <div class="bg-purple-50 rounded-lg p-6 cursor-pointer hover:bg-purple-100" @click="selectReport('roi')">
        <div class="flex items-center">
          <div class="p-2 bg-purple-100 rounded-lg">
            <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1" />
            </svg>
          </div>
          <div class="mr-3">
            <h3 class="text-sm font-medium text-purple-900">ROI</h3>
            <p class="text-xs text-purple-600">بازگشت سرمایه تست‌ها</p>
          </div>
        </div>
      </div>
    </div>

    <!-- پیش‌نمایش گزارش -->
    <div v-if="currentReport" class="bg-gray-50 rounded-lg p-6">
      <div class="flex items-center justify-between mb-4">
        <h3 class="text-lg font-semibold text-gray-900">{{ currentReport.title }}</h3>
        <div class="flex items-center space-x-2 space-x-reverse">
          <button class="px-3 py-1 text-sm bg-green-100 text-green-700 rounded-lg hover:bg-green-200" @click="exportReport">
            صادرات
          </button>
          <button class="px-3 py-1 text-sm bg-blue-100 text-blue-700 rounded-lg hover:bg-blue-200" @click="printReport">
            چاپ
          </button>
        </div>
      </div>

      <!-- محتوای گزارش -->
      <div class="space-y-4">
        <!-- خلاصه اجرایی -->
        <div class="bg-white rounded-lg p-6">
          <h4 class="text-md font-semibold text-gray-900 mb-2">خلاصه اجرایی</h4>
          <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
            <div class="text-center">
              <div class="text-2xl font-bold text-blue-600">{{ currentReport.summary.totalTests }}</div>
              <div class="text-sm text-gray-600">کل تست‌ها</div>
            </div>
            <div class="text-center">
              <div class="text-2xl font-bold text-green-600">{{ currentReport.summary.successfulTests }}</div>
              <div class="text-sm text-gray-600">تست‌های موفق</div>
            </div>
            <div class="text-center">
              <div class="text-2xl font-bold text-purple-600">{{ currentReport.summary.avgImprovement }}%</div>
              <div class="text-sm text-gray-600">متوسط بهبود</div>
            </div>
          </div>
        </div>

        <!-- نمودار عملکرد -->
        <div class="bg-white rounded-lg p-6">
          <h4 class="text-md font-semibold text-gray-900 mb-4">عملکرد در طول زمان</h4>
          <div class="h-64 flex items-end justify-between space-x-2 space-x-reverse">
            <div v-for="(month, index) in currentReport.performance" :key="index" class="flex-1 flex flex-col items-center">
              <div class="w-full bg-blue-500 rounded-t-lg" :style="{ height: `${month.value * 2}px` }"></div>
              <span class="text-xs text-gray-500 mt-2">{{ month.label }}</span>
            </div>
          </div>
        </div>

        <!-- جدول جزئیات -->
        <div class="bg-white rounded-lg p-6">
          <h4 class="text-md font-semibold text-gray-900 mb-4">جزئیات تست‌ها</h4>
          <div class="overflow-x-auto">
            <table class="min-w-full divide-y divide-gray-200">
              <thead class="bg-gray-50">
                <tr>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نام تست</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نتیجه</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">بهبود</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">ROI</th>
                </tr>
              </thead>
              <tbody class="bg-white divide-y divide-gray-200">
                <tr v-for="test in currentReport.details" :key="test.id">
                  <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{{ test.name }}</td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ test.type }}</td>
                  <td class="px-6 py-4 whitespace-nowrap">
                    <span :class="test.result === 'success' ? 'text-green-600' : 'text-red-600'" class="text-sm">
                      {{ test.result === 'success' ? 'موفق' : 'ناموفق' }}
                    </span>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ test.improvement }}%</td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ test.roi }}%</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <!-- توصیه‌ها -->
        <div class="bg-white rounded-lg p-6">
          <h4 class="text-md font-semibold text-gray-900 mb-4">توصیه‌ها</h4>
          <ul class="space-y-2">
            <li v-for="(recommendation, index) in currentReport.recommendations" :key="index" class="flex items-start">
              <svg class="w-4 h-4 text-blue-500 mt-0.5 ml-2 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
              <span class="text-sm text-gray-700">{{ recommendation }}</span>
            </li>
          </ul>
        </div>
      </div>
    </div>

    <!-- تاریخچه گزارش‌ها -->
    <div class="mt-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">تاریخچه گزارش‌ها</h3>
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نام گزارش</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ ایجاد</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="report in reportHistory" :key="report.id">
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{{ report.name }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ report.type }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ formatDate(report.createdAt) }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <div class="flex items-center space-x-3 space-x-reverse">
                  <button class="text-blue-600 hover:text-blue-900" @click="downloadReport(report.id)">دانلود</button>
                  <button class="text-red-600 hover:text-red-900" @click="deleteReport(report.id)">حذف</button>
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
import { onMounted, ref } from 'vue'
// State
const reportType = ref('performance')
const timeRange = ref('30')
const testType = ref('')
const exportFormat = ref('excel')
const currentReport = ref(null)

// تاریخچه گزارش‌ها
const reportHistory = ref([
  {
    id: 1,
    name: 'گزارش عملکرد ماهانه',
    type: 'عملکرد کلی',
    createdAt: '2024-01-15'
  },
  {
    id: 2,
    name: 'گزارش ROI تست‌ها',
    type: 'ROI',
    createdAt: '2024-01-10'
  },
  {
    id: 3,
    name: 'گزارش تست‌های موفق',
    type: 'تست‌های موفق',
    createdAt: '2024-01-05'
  }
])

// انتخاب گزارش
const selectReport = (type: string) => {
  reportType.value = type
  generateReport()
}

// ایجاد گزارش
const generateReport = () => {
  // داده‌های نمونه برای گزارش
  const reports = {
    performance: {
      title: 'گزارش عملکرد کلی',
      summary: {
        totalTests: 24,
        successfulTests: 18,
        avgImprovement: 15.2
      },
      performance: [
        { label: 'فروردین', value: 12 },
        { label: 'اردیبهشت', value: 15 },
        { label: 'خرداد', value: 18 },
        { label: 'تیر', value: 22 },
        { label: 'مرداد', value: 20 },
        { label: 'شهریور', value: 24 }
      ],
      details: [
        { id: 1, name: 'تست دکمه خرید', type: 'دکمه', result: 'success', improvement: 28.5, roi: 340 },
        { id: 2, name: 'تست قیمت محصول', type: 'قیمت', result: 'success', improvement: 12.3, roi: 180 },
        { id: 3, name: 'تست تصویر هدر', type: 'تصویر', result: 'failure', improvement: -5.2, roi: -45 }
      ],
      recommendations: [
        'تست‌های دکمه‌ها بیشترین موفقیت را داشته‌اند',
        'تست‌های قیمت‌گذاری نیاز به بررسی دقیق‌تر دارند',
        'تست‌های تصویری نیاز به استراتژی جدید دارند'
      ]
    },
    successful: {
      title: 'گزارش تست‌های موفق',
      summary: {
        totalTests: 18,
        successfulTests: 18,
        avgImprovement: 22.1
      },
      performance: [
        { label: 'فروردین', value: 8 },
        { label: 'اردیبهشت', value: 10 },
        { label: 'خرداد', value: 12 },
        { label: 'تیر', value: 15 },
        { label: 'مرداد', value: 16 },
        { label: 'شهریور', value: 18 }
      ],
      details: [
        { id: 1, name: 'تست دکمه خرید', type: 'دکمه', result: 'success', improvement: 28.5, roi: 340 },
        { id: 2, name: 'تست قیمت محصول', type: 'قیمت', result: 'success', improvement: 12.3, roi: 180 },
        { id: 3, name: 'تست رنگ هدر', type: 'صفحه', result: 'success', improvement: 8.7, roi: 120 }
      ],
      recommendations: [
        'تست‌های دکمه‌ها بیشترین ROI را داشته‌اند',
        'تست‌های قیمت‌گذاری نیاز به تکرار دارند',
        'تست‌های رنگ‌بندی موثر بوده‌اند'
      ]
    },
    unsuccessful: {
      title: 'گزارش تست‌های ناموفق',
      summary: {
        totalTests: 6,
        successfulTests: 0,
        avgImprovement: -8.3
      },
      performance: [
        { label: 'فروردین', value: 4 },
        { label: 'اردیبهشت', value: 5 },
        { label: 'خرداد', value: 6 },
        { label: 'تیر', value: 7 },
        { label: 'مرداد', value: 4 },
        { label: 'شهریور', value: 6 }
      ],
      details: [
        { id: 1, name: 'تست تصویر هدر', type: 'تصویر', result: 'failure', improvement: -5.2, roi: -45 },
        { id: 2, name: 'تست فونت متن', type: 'صفحه', result: 'failure', improvement: -12.1, roi: -120 },
        { id: 3, name: 'تست چیدمان', type: 'صفحه', result: 'failure', improvement: -7.6, roi: -85 }
      ],
      recommendations: [
        'تست‌های تصویری نیاز به استراتژی جدید دارند',
        'تست‌های فونت باید با دقت بیشتری انجام شوند',
        'تست‌های چیدمان نیاز به تحلیل عمیق‌تر دارند'
      ]
    },
    roi: {
      title: 'گزارش ROI',
      summary: {
        totalTests: 24,
        successfulTests: 18,
        avgImprovement: 15.2
      },
      performance: [
        { label: 'فروردین', value: 120 },
        { label: 'اردیبهشت', value: 180 },
        { label: 'خرداد', value: 220 },
        { label: 'تیر', value: 280 },
        { label: 'مرداد', value: 320 },
        { label: 'شهریور', value: 340 }
      ],
      details: [
        { id: 1, name: 'تست دکمه خرید', type: 'دکمه', result: 'success', improvement: 28.5, roi: 340 },
        { id: 2, name: 'تست قیمت محصول', type: 'قیمت', result: 'success', improvement: 12.3, roi: 180 },
        { id: 3, name: 'تست رنگ هدر', type: 'صفحه', result: 'success', improvement: 8.7, roi: 120 }
      ],
      recommendations: [
        'ROI کلی سیستم مثبت و رو به رشد است',
        'تست‌های دکمه‌ها بیشترین بازگشت سرمایه را داشته‌اند',
        'تست‌های قیمت‌گذاری نیاز به بهینه‌سازی دارند'
      ]
    }
  }

  currentReport.value = reports[reportType.value as keyof typeof reports]
}

// صادرات گزارش
const exportReport = () => {

  alert(`گزارش با موفقیت به فرمت ${exportFormat.value} صادر شد`)
}

// چاپ گزارش
const printReport = () => {
  window.print()
}

// دانلود گزارش
const downloadReport = (_reportId: number) => {

  alert('گزارش با موفقیت دانلود شد')
}

// حذف گزارش
const deleteReport = (reportId: number) => {
  if (confirm('آیا از حذف این گزارش اطمینان دارید؟')) {
    reportHistory.value = reportHistory.value.filter(report => report.id !== reportId)
  }
}

// فرمت کردن تاریخ
const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleDateString('fa-IR')
}

// ایجاد گزارش اولیه
onMounted(() => {
  generateReport()
})
</script> 
