<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">نمودارهای تعاملی</h4>
        <p class="text-sm text-gray-600 mt-1">مدیریت و تنظیم نمودارهای تعاملی</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          @click="exportCharts"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
          </svg>
          خروجی نمودارها
        </button>
      </div>
    </div>

    <!-- انتخاب نوع نمودار -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">انتخاب نوع نمودار</h5>
      <div class="grid grid-cols-2 md:grid-cols-4 gap-6">
        <button
          v-for="chartType in chartTypes"
          :key="chartType.id"
          class="p-6 border-2 rounded-lg transition-all text-center"
          :class="selectedChartType === chartType.id ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-gray-300'"
          @click="selectedChartType = chartType.id"
        >
          <div class="w-12 h-12 mx-auto mb-2 flex items-center justify-center">
            <svg class="w-8 h-8 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="chartType.icon" />
            </svg>
          </div>
          <div class="text-sm font-medium text-gray-900">{{ chartType.name }}</div>
          <div class="text-xs text-gray-500 mt-1">{{ chartType.description }}</div>
        </button>
      </div>
    </div>

    <!-- تنظیمات نمودار -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">تنظیمات نمودار</h5>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">عنوان نمودار</label>
          <input
            v-model="chartSettings.title"
            type="text"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            placeholder="عنوان نمودار را وارد کنید"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نوع نمودار</label>
          <select
            v-model="chartSettings.type"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="line">خطی</option>
            <option value="bar">ستونی</option>
            <option value="pie">دایره‌ای</option>
            <option value="area">ناحیه‌ای</option>
            <option value="scatter">پراکندگی</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">رنگ‌بندی</label>
          <select
            v-model="chartSettings.colorScheme"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="blue">آبی</option>
            <option value="green">سبز</option>
            <option value="purple">بنفش</option>
            <option value="orange">نارنجی</option>
            <option value="rainbow">رنگین‌کمان</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">اندازه نمودار</label>
          <select
            v-model="chartSettings.size"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="small">کوچک</option>
            <option value="medium">متوسط</option>
            <option value="large">بزرگ</option>
            <option value="full">تمام عرض</option>
          </select>
        </div>
      </div>

      <div class="mt-4 space-y-3">
        <label class="flex items-center">
          <input
            v-model="chartSettings.showLegend"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">نمایش راهنما</span>
        </label>

        <label class="flex items-center">
          <input
            v-model="chartSettings.showGrid"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">نمایش شبکه</span>
        </label>

        <label class="flex items-center">
          <input
            v-model="chartSettings.animate"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">انیمیشن</span>
        </label>

        <label class="flex items-center">
          <input
            v-model="chartSettings.interactive"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">تعاملی</span>
        </label>
      </div>
    </div>

    <!-- پیش‌نمایش نمودار -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">پیش‌نمایش نمودار</h5>
      <div class="bg-gray-50 rounded-lg p-6">
        <div class="h-64 flex items-center justify-center">
          <div class="text-center">
            <svg class="w-16 h-16 mx-auto mb-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
            </svg>
            <p class="text-gray-500">پیش‌نمایش نمودار {{ chartSettings.title || 'بدون عنوان' }}</p>
            <p class="text-sm text-gray-400 mt-1">نوع: {{ getChartTypeName(chartSettings.type) }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- نمودارهای موجود -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">نمودارهای موجود</h5>
      </div>
      <div class="p-6">
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <div
            v-for="chart in existingCharts"
            :key="chart.id"
            class="border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow"
          >
            <div class="flex items-center justify-between mb-3">
              <h6 class="text-sm font-medium text-gray-900">{{ chart.title }}</h6>
              <div class="flex items-center space-x-2 space-x-reverse">
                <button
                  class="text-blue-600 hover:text-blue-800"
                  @click="editChart(chart)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                  </svg>
                </button>
                <button
                  class="text-red-600 hover:text-red-800"
                  @click="deleteChart(chart)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                  </svg>
                </button>
              </div>
            </div>
            
            <div class="h-32 bg-gray-50 rounded flex items-center justify-center mb-3">
              <svg class="w-8 h-8 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="getChartIcon(chart.type)" />
              </svg>
            </div>
            
            <div class="flex items-center justify-between text-xs text-gray-500">
              <span>{{ getChartTypeName(chart.type) }}</span>
              <span>{{ chart.lastUpdated }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات داده -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">تنظیمات داده</h5>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">منبع داده</label>
          <select
            v-model="chartSettings.dataSource"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="transactions">تراکنش‌ها</option>
            <option value="connections">اتصالات</option>
            <option value="errors">خطاها</option>
            <option value="performance">عملکرد</option>
            <option value="custom">داده سفارشی</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">بازه زمانی</label>
          <select
            v-model="chartSettings.timeRange"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="1day">۱ روز</option>
            <option value="7days">۷ روز</option>
            <option value="30days">۳۰ روز</option>
            <option value="3months">۳ ماه</option>
            <option value="1year">۱ سال</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تعداد نقاط داده</label>
          <input
            v-model="chartSettings.dataPoints"
            type="number"
            min="5"
            max="100"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">به‌روزرسانی خودکار</label>
          <select
            v-model="chartSettings.autoRefresh"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="disabled">غیرفعال</option>
            <option value="30s">۳۰ ثانیه</option>
            <option value="1m">۱ دقیقه</option>
            <option value="5m">۵ دقیقه</option>
            <option value="15m">۱۵ دقیقه</option>
          </select>
        </div>
      </div>
    </div>

    <!-- دکمه‌های عملیات -->
    <div class="flex items-center justify-end space-x-3 space-x-reverse">
      <button
        class="px-4 py-2 border border-gray-300 rounded-md text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
        @click="resetSettings"
      >
        بازنشانی
      </button>
      <button
        class="px-4 py-2 border border-transparent rounded-md text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
        @click="previewChart"
      >
        پیش‌نمایش
      </button>
      <button
        class="px-4 py-2 border border-transparent rounded-md text-sm font-medium text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
        @click="saveChart"
      >
        ذخیره نمودار
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// انواع نمودار
const chartTypes = ref([
  {
    id: 'line',
    name: 'نمودار خطی',
    description: 'برای نمایش روند',
    icon: 'M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z'
  },
  {
    id: 'bar',
    name: 'نمودار ستونی',
    description: 'برای مقایسه مقادیر',
    icon: 'M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z'
  },
  {
    id: 'pie',
    name: 'نمودار دایره‌ای',
    description: 'برای نمایش نسبت‌ها',
    icon: 'M11 3.055A9.001 9.001 0 1020.945 13H11V3.055z'
  },
  {
    id: 'area',
    name: 'نمودار ناحیه‌ای',
    description: 'برای نمایش حجم',
    icon: 'M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z'
  }
])

// تنظیمات نمودار
const chartSettings = ref({
  title: '',
  type: 'line',
  colorScheme: 'blue',
  size: 'medium',
  showLegend: true,
  showGrid: true,
  animate: true,
  interactive: true,
  dataSource: 'transactions',
  timeRange: '30days',
  dataPoints: 20,
  autoRefresh: 'disabled'
})

// انتخاب نوع نمودار
const selectedChartType = ref('line')

// نمودارهای موجود
const existingCharts = ref([
  {
    id: 1,
    title: 'روند تراکنش‌ها',
    type: 'line',
    lastUpdated: '۲ ساعت پیش'
  },
  {
    id: 2,
    title: 'توزیع خطاها',
    type: 'pie',
    lastUpdated: '۱ روز پیش'
  },
  {
    id: 3,
    title: 'عملکرد اتصالات',
    type: 'bar',
    lastUpdated: '۳ روز پیش'
  }
])

// نام نوع نمودار
const getChartTypeName = (type: string) => {
  const names = {
    line: 'خطی',
    bar: 'ستونی',
    pie: 'دایره‌ای',
    area: 'ناحیه‌ای',
    scatter: 'پراکندگی'
  }
  return names[type] || type
}

// آیکون نمودار
const getChartIcon = (type: string) => {
  const icons = {
    line: 'M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z',
    bar: 'M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z',
    pie: 'M11 3.055A9.001 9.001 0 1020.945 13H11V3.055z',
    area: 'M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z'
  }
  return icons[type] || icons.line
}

// عملیات
const exportCharts = () => {

}

interface Chart {
  id: number | string
  [key: string]: unknown
}

const editChart = (_chart: Chart) => {

}

const deleteChart = (_chart: Chart) => {

}

const resetSettings = () => {
  chartSettings.value = {
    title: '',
    type: 'line',
    colorScheme: 'blue',
    size: 'medium',
    showLegend: true,
    showGrid: true,
    animate: true,
    interactive: true,
    dataSource: 'transactions',
    timeRange: '30days',
    dataPoints: 20,
    autoRefresh: 'disabled'
  }
}

const previewChart = () => {

}

const saveChart = () => {

}
</script>

<!--
  کامپوننت نمودارهای تعاملی
  شامل:
  1. انتخاب نوع نمودار
  2. تنظیمات نمودار
  3. پیش‌نمایش
  4. نمودارهای موجود
  5. تنظیمات داده
  6. طراحی مدرن و کاملاً ریسپانسیو
--> 
