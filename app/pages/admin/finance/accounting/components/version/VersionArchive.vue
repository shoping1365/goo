<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">آرشیو نسخه‌های قدیمی</h4>
        <p class="text-sm text-gray-600 mt-1">مدیریت و نگهداری نسخه‌های قدیمی اتصال</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          class="inline-flex items-center px-4 py-2 border border-red-300 shadow-sm text-sm leading-4 font-medium rounded-md text-red-700 bg-white hover:bg-red-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500"
          @click="cleanupArchive"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
          </svg>
          پاکسازی آرشیو
        </button>
      </div>
    </div>

    <!-- آمار آرشیو -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div class="bg-blue-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-blue-600">{{ archiveStats.total }}</div>
        <div class="text-sm text-blue-700">کل نسخه‌های آرشیو شده</div>
      </div>
      <div class="bg-green-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-green-600">{{ archiveStats.size }}</div>
        <div class="text-sm text-green-700">حجم کل آرشیو (GB)</div>
      </div>
      <div class="bg-yellow-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-yellow-600">{{ archiveStats.oldest }}</div>
        <div class="text-sm text-yellow-700">قدیمی‌ترین نسخه</div>
      </div>
      <div class="bg-purple-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-purple-600">{{ archiveStats.space }}</div>
        <div class="text-sm text-purple-700">فضای آزاد (GB)</div>
      </div>
    </div>

    <!-- فیلترها -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div>
          <input
            v-model="searchQuery"
            type="text"
            placeholder="جستجو در آرشیو..."
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
        </div>
        <div>
          <select
            v-model="selectedAge"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه سن‌ها</option>
            <option value="1month">کمتر از ۱ ماه</option>
            <option value="3months">کمتر از ۳ ماه</option>
            <option value="6months">کمتر از ۶ ماه</option>
            <option value="1year">کمتر از ۱ سال</option>
            <option value="older">بیش از ۱ سال</option>
          </select>
        </div>
        <div>
          <select
            v-model="selectedSize"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه اندازه‌ها</option>
            <option value="small">کوچک (کمتر از ۱۰۰MB)</option>
            <option value="medium">متوسط (۱۰۰MB تا ۵۰۰MB)</option>
            <option value="large">بزرگ (بیش از ۵۰۰MB)</option>
          </select>
        </div>
        <div>
          <button
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-blue-500"
            @click="clearFilters"
          >
            پاک کردن فیلترها
          </button>
        </div>
      </div>
    </div>

    <!-- جدول آرشیو -->
    <div class="bg-white rounded-lg border border-gray-200 overflow-hidden">
      <div class="px-6 py-4 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <h5 class="text-md font-medium text-gray-900">لیست نسخه‌های آرشیو شده</h5>
          <div class="flex items-center space-x-2 space-x-reverse">
            <button
              class="text-sm text-blue-600 hover:text-blue-800"
              @click="selectAll"
            >
              انتخاب همه
            </button>
            <button
              class="text-sm text-gray-600 hover:text-gray-800"
              @click="deselectAll"
            >
              لغو انتخاب
            </button>
          </div>
        </div>
      </div>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500">
                <input
                  v-model="selectAllItems"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500">نسخه</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500">تاریخ آرشیو</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500">حجم</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500">سن</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="version in filteredArchivedVersions" :key="version.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <input
                  v-model="selectedVersions"
                  :value="version.id"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="flex-shrink-0 w-8 h-8 bg-gray-100 rounded-full flex items-center justify-center">
                    <svg class="w-4 h-4 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 8h14M5 8a2 2 0 110-4h14a2 2 0 110 4M5 8v10a2 2 0 002 2h10a2 2 0 002-2V8m-9 4h4" />
                    </svg>
                  </div>
                  <div class="mr-3">
                    <div class="text-sm font-medium text-gray-900">{{ version.version }}</div>
                    <div class="text-sm text-gray-500">{{ version.connectionName }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ version.archivedAt }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ version.size }} MB
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ version.age }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span
                  class="px-2 py-1 text-xs font-medium rounded-full"
                  :class="getStatusClass(version.status)"
                >
                  {{ getStatusLabel(version.status) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <div class="flex items-center space-x-2 space-x-reverse">
                  <button
                    class="text-blue-600 hover:text-blue-900"
                    @click="restoreVersion(version)"
                  >
                    بازیابی
                  </button>
                  <button
                    class="text-green-600 hover:text-green-900"
                    @click="downloadVersion(version)"
                  >
                    دانلود
                  </button>
                  <button
                    class="text-red-600 hover:text-red-900"
                    @click="deleteVersion(version)"
                  >
                    حذف
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- عملیات دسته‌ای -->
    <div v-if="selectedVersions.length > 0" class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">عملیات دسته‌ای</h5>
      <div class="flex items-center space-x-4 space-x-reverse">
        <span class="text-sm text-gray-600">{{ selectedVersions.length }} نسخه انتخاب شده</span>
        <button
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          @click="restoreSelected"
        >
          بازیابی انتخاب شده‌ها
        </button>
        <button
          class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          @click="downloadSelected"
        >
          دانلود انتخاب شده‌ها
        </button>
        <button
          class="inline-flex items-center px-4 py-2 border border-red-300 shadow-sm text-sm leading-4 font-medium rounded-md text-red-700 bg-white hover:bg-red-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500"
          @click="deleteSelected"
        >
          حذف انتخاب شده‌ها
        </button>
      </div>
    </div>

    <!-- تنظیمات آرشیو -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">تنظیمات آرشیو</h5>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر سن آرشیو</label>
          <select
            v-model="archiveSettings.maxAge"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="1month">۱ ماه</option>
            <option value="3months">۳ ماه</option>
            <option value="6months">۶ ماه</option>
            <option value="1year">۱ سال</option>
            <option value="2years">۲ سال</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر حجم آرشیو</label>
          <select
            v-model="archiveSettings.maxSize"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="1gb">۱ گیگابایت</option>
            <option value="5gb">۵ گیگابایت</option>
            <option value="10gb">۱۰ گیگابایت</option>
            <option value="20gb">۲۰ گیگابایت</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">فاصله آرشیو خودکار</label>
          <select
            v-model="archiveSettings.autoArchiveInterval"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="daily">روزانه</option>
            <option value="weekly">هفتگی</option>
            <option value="monthly">ماهانه</option>
            <option value="disabled">غیرفعال</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">فشرده‌سازی آرشیو</label>
          <div class="flex items-center">
            <input
              v-model="archiveSettings.compression"
              type="checkbox"
              class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
            />
            <label class="mr-3 text-sm text-gray-700">
              فعال کردن فشرده‌سازی برای صرفه‌جویی در فضا
            </label>
          </div>
        </div>
      </div>
      <div class="mt-4 flex items-center justify-end">
        <button
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          @click="saveArchiveSettings"
        >
          ذخیره تنظیمات
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

// فیلترها
const searchQuery = ref('')
const selectedAge = ref('')
const selectedSize = ref('')

// انتخاب نسخه‌ها
const selectedVersions = ref<number[]>([])
const selectAllItems = ref(false)

// تنظیمات آرشیو
const archiveSettings = ref({
  maxAge: '6months',
  maxSize: '10gb',
  autoArchiveInterval: 'weekly',
  compression: true
})

// آمار آرشیو
const archiveStats = ref({
  total: 45,
  size: 2.8,
  oldest: 'v1.5.0',
  space: 15.2
})

// نسخه‌های آرشیو شده
const archivedVersions = ref([
  {
    id: 1,
    version: 'v2.0.5',
    connectionName: 'اتصال هلو',
    archivedAt: '۱۴۰۲/۱۰/۱۰ - ۰۹:۱۵',
    size: 42.8,
    age: '۲ ماه پیش',
    status: 'archived'
  },
  {
    id: 2,
    version: 'v2.0.0',
    connectionName: 'اتصال هلو',
    archivedAt: '۱۴۰۲/۰۹/۲۸ - ۱۶:۴۵',
    size: 38.5,
    age: '۳ ماه پیش',
    status: 'archived'
  },
  {
    id: 3,
    version: 'v1.9.8',
    connectionName: 'اتصال هلو',
    archivedAt: '۱۴۰۲/۰۹/۱۵ - ۱۱:۲۰',
    size: 35.2,
    age: '۴ ماه پیش',
    status: 'archived'
  },
  {
    id: 4,
    version: 'v1.9.5',
    connectionName: 'اتصال هلو',
    archivedAt: '۱۴۰۲/۰۸/۲۰ - ۱۴:۳۰',
    size: 32.1,
    age: '۵ ماه پیش',
    status: 'archived'
  }
])

// فیلتر کردن نسخه‌های آرشیو شده
const filteredArchivedVersions = computed(() => {
  return archivedVersions.value.filter(version => {
    if (searchQuery.value && !version.version.includes(searchQuery.value) && !version.connectionName.includes(searchQuery.value)) return false
    // TODO: فیلتر سن و اندازه
    return true
  })
})

// انتخاب همه
const selectAll = () => {
  selectedVersions.value = filteredArchivedVersions.value.map(v => v.id)
  selectAllItems.value = true
}

// لغو انتخاب همه
const deselectAll = () => {
  selectedVersions.value = []
  selectAllItems.value = false
}

// کلاس وضعیت
const getStatusClass = (status: string) => {
  const classes = {
    archived: 'bg-yellow-100 text-yellow-700',
    compressed: 'bg-blue-100 text-blue-700',
    deleted: 'bg-red-100 text-red-700'
  }
  return classes[status] || 'bg-gray-100 text-gray-700'
}

// برچسب وضعیت
const getStatusLabel = (status: string) => {
  const labels = {
    archived: 'آرشیو شده',
    compressed: 'فشرده شده',
    deleted: 'حذف شده'
  }
  return labels[status] || status
}

// عملیات
const restoreVersion = (version: any) => {
  console.log('بازیابی نسخه:', version)
}

const downloadVersion = (version: any) => {
  console.log('دانلود نسخه:', version)
}

const deleteVersion = (version: any) => {
  console.log('حذف نسخه:', version)
}

const restoreSelected = () => {
  console.log('بازیابی نسخه‌های انتخاب شده:', selectedVersions.value)
}

const downloadSelected = () => {
  console.log('دانلود نسخه‌های انتخاب شده:', selectedVersions.value)
}

const deleteSelected = () => {
  console.log('حذف نسخه‌های انتخاب شده:', selectedVersions.value)
}

const cleanupArchive = () => {
  console.log('پاکسازی آرشیو')
}

const saveArchiveSettings = () => {
  console.log('ذخیره تنظیمات آرشیو:', archiveSettings.value)
}

const clearFilters = () => {
  searchQuery.value = ''
  selectedAge.value = ''
  selectedSize.value = ''
}
</script>

<!--
  کامپوننت آرشیو نسخه‌های قدیمی
  شامل:
  1. آمار آرشیو
  2. فیلتر و جستجو
  3. جدول نسخه‌های آرشیو شده
  4. عملیات دسته‌ای
  5. تنظیمات آرشیو
  6. طراحی مدرن و کاملاً ریسپانسیو
--> 
