<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">مقایسه نسخه‌ها</h4>
        <p class="text-sm text-gray-600 mt-1">مقایسه تفاوت‌های بین نسخه‌های مختلف اتصال</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          @click="exportComparison"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
          </svg>
          خروجی مقایسه
        </button>
      </div>
    </div>

    <!-- انتخاب نسخه‌ها -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">انتخاب نسخه‌ها برای مقایسه</h5>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نسخه اول</label>
          <select
            v-model="selectedVersion1"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">انتخاب نسخه...</option>
            <option
              v-for="version in availableVersions"
              :key="version.id"
              :value="version.id"
            >
              {{ version.version }} - {{ version.createdAt }}
            </option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نسخه دوم</label>
          <select
            v-model="selectedVersion2"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">انتخاب نسخه...</option>
            <option
              v-for="version in availableVersions"
              :key="version.id"
              :value="version.id"
            >
              {{ version.version }} - {{ version.createdAt }}
            </option>
          </select>
        </div>
      </div>

      <div class="mt-4 flex items-center justify-center">
        <button
          class="inline-flex items-center px-3 py-2 border border-gray-300 rounded-md text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-blue-500"
          @click="swapVersions"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16V4m0 0L3 8m4-4l4 4m6 0v12m0 0l4-4m-4 4l-4-4" />
          </svg>
          جابجایی نسخه‌ها
        </button>
      </div>
    </div>

    <!-- خلاصه مقایسه -->
    <div v-if="selectedVersion1 && selectedVersion2" class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">خلاصه مقایسه</h5>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div class="bg-blue-50 rounded-lg p-6 text-center">
          <div class="text-2xl font-bold text-blue-600">{{ comparisonStats.added }}</div>
          <div class="text-sm text-blue-700">موارد اضافه شده</div>
        </div>
        <div class="bg-red-50 rounded-lg p-6 text-center">
          <div class="text-2xl font-bold text-red-600">{{ comparisonStats.removed }}</div>
          <div class="text-sm text-red-700">موارد حذف شده</div>
        </div>
        <div class="bg-yellow-50 rounded-lg p-6 text-center">
          <div class="text-2xl font-bold text-yellow-600">{{ comparisonStats.modified }}</div>
          <div class="text-sm text-yellow-700">موارد تغییر یافته</div>
        </div>
      </div>
    </div>

    <!-- نمودار مقایسه -->
    <div v-if="selectedVersion1 && selectedVersion2" class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">نمودار مقایسه</h5>
      <div class="h-64 bg-gray-50 rounded-lg flex items-center justify-center">
        <div class="text-center">
          <svg class="w-16 h-16 mx-auto mb-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
          </svg>
          <p class="text-gray-500">نمودار مقایسه نسخه‌ها</p>
        </div>
      </div>
    </div>

    <!-- تفاوت‌های جزئی -->
    <div v-if="selectedVersion1 && selectedVersion2" class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">تفاوت‌های جزئی</h5>
      </div>
      <div class="divide-y divide-gray-200">
        <div
          v-for="(diff, index) in detailedDifferences"
          :key="index"
          class="p-6 hover:bg-gray-50"
        >
          <div class="flex items-start space-x-4 space-x-reverse">
            <div class="flex-shrink-0">
              <div
                class="w-3 h-3 rounded-full"
                :class="getDiffTypeClass(diff.type)"
              ></div>
            </div>
            <div class="flex-1">
              <div class="flex items-center justify-between">
                <h6 class="text-sm font-medium text-gray-900">{{ diff.title }}</h6>
                <span
                  class="px-2 py-1 text-xs font-medium rounded-full"
                  :class="getDiffTypeClass(diff.type)"
                >
                  {{ getDiffTypeLabel(diff.type) }}
                </span>
              </div>
              <p class="text-sm text-gray-600 mt-1">{{ diff.description }}</p>
              
              <!-- نمایش تفاوت‌ها -->
              <div v-if="diff.details" class="mt-3 bg-gray-50 rounded-lg p-3">
                <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                  <div>
                    <div class="text-xs font-medium text-gray-500 mb-1">نسخه اول</div>
                    <div class="text-sm text-gray-900 bg-white p-2 rounded border">{{ diff.details.version1 }}</div>
                  </div>
                  <div>
                    <div class="text-xs font-medium text-gray-500 mb-1">نسخه دوم</div>
                    <div class="text-sm text-gray-900 bg-white p-2 rounded border">{{ diff.details.version2 }}</div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- فیلتر تفاوت‌ها -->
    <div v-if="selectedVersion1 && selectedVersion2" class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">فیلتر تفاوت‌ها</h5>
      <div class="flex flex-wrap gap-2">
        <button
          v-for="filter in diffFilters"
          :key="filter.type"
          class="px-3 py-1 text-sm font-medium rounded-full transition-colors"
          :class="activeDiffFilters.includes(filter.type)
            ? 'bg-blue-100 text-blue-700 border border-blue-200'
            : 'bg-gray-100 text-gray-700 hover:bg-gray-200 border border-gray-200'"
          @click="toggleDiffFilter(filter.type)"
        >
          {{ filter.label }}
          <span class="text-xs text-gray-500 mr-1">({{ filter.count }})</span>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

// انتخاب نسخه‌ها
const selectedVersion1 = ref('')
const selectedVersion2 = ref('')

// فیلترهای فعال
const activeDiffFilters = ref(['added', 'removed', 'modified'])

// نسخه‌های موجود
const availableVersions = ref([
  {
    id: 1,
    version: 'v2.1.0',
    createdAt: '۱۴۰۲/۱۰/۱۵ - ۱۴:۳۰',
    description: 'نسخه پایدار با بهبودهای عملکرد'
  },
  {
    id: 2,
    version: 'v2.0.5',
    createdAt: '۱۴۰۲/۱۰/۱۰ - ۰۹:۱۵',
    description: 'نسخه خودکار قبل از تغییرات'
  },
  {
    id: 3,
    version: 'v2.0.0',
    createdAt: '۱۴۰۲/۰۹/۲۸ - ۱۶:۴۵',
    description: 'نسخه اصلی با قابلیت‌های جدید'
  },
  {
    id: 4,
    version: 'v1.9.8',
    createdAt: '۱۴۰۲/۰۹/۱۵ - ۱۱:۲۰',
    description: 'آخرین نسخه از سری ۱.x'
  }
])

// آمار مقایسه
const comparisonStats = ref({
  added: 15,
  removed: 8,
  modified: 12
})

// فیلترهای تفاوت
const diffFilters = ref([
  { type: 'added', label: 'اضافه شده', count: 15 },
  { type: 'removed', label: 'حذف شده', count: 8 },
  { type: 'modified', label: 'تغییر یافته', count: 12 }
])

// تفاوت‌های جزئی
const detailedDifferences = ref([
  {
    type: 'added',
    title: 'تنظیمات امنیتی جدید',
    description: 'تنظیمات امنیتی پیشرفته اضافه شده است',
    details: {
      version1: 'تنظیمات امنیتی پایه',
      version2: 'تنظیمات امنیتی پیشرفته + رمزگذاری دوگانه'
    }
  },
  {
    type: 'removed',
    title: 'قابلیت قدیمی',
    description: 'قابلیت قدیمی که دیگر استفاده نمی‌شود حذف شده است',
    details: {
      version1: 'قابلیت قدیمی فعال',
      version2: 'حذف شده'
    }
  },
  {
    type: 'modified',
    title: 'پیکربندی اتصال',
    description: 'پیکربندی اتصال بهبود یافته است',
    details: {
      version1: 'پیکربندی ساده',
      version2: 'پیکربندی پیشرفته با تنظیمات اضافی'
    }
  }
])

// فیلتر کردن تفاوت‌ها
const _filteredDifferences = computed(() => {
  return detailedDifferences.value.filter(diff => 
    activeDiffFilters.value.includes(diff.type)
  )
})

// کلاس نوع تفاوت
const getDiffTypeClass = (type: string) => {
  const classes = {
    added: 'bg-green-500',
    removed: 'bg-red-500',
    modified: 'bg-yellow-500'
  }
  return classes[type] || 'bg-gray-500'
}

// برچسب نوع تفاوت
const getDiffTypeLabel = (type: string) => {
  const labels = {
    added: 'اضافه شده',
    removed: 'حذف شده',
    modified: 'تغییر یافته'
  }
  return labels[type] || type
}

// عملیات
const swapVersions = () => {
  const temp = selectedVersion1.value
  selectedVersion1.value = selectedVersion2.value
  selectedVersion2.value = temp
}

const toggleDiffFilter = (filterType: string) => {
  const index = activeDiffFilters.value.indexOf(filterType)
  if (index > -1) {
    activeDiffFilters.value.splice(index, 1)
  } else {
    activeDiffFilters.value.push(filterType)
  }
}

const exportComparison = () => {

}
</script>

<!--
  کامپوننت مقایسه نسخه‌ها
  شامل:
  1. انتخاب دو نسخه
  2. خلاصه مقایسه
  3. نمودار مقایسه
  4. تفاوت‌های جزئی
  5. فیلتر تفاوت‌ها
  6. طراحی مدرن و کاملاً ریسپانسیو
--> 
