<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">نگهداری نسخه‌ها</h4>
        <p class="text-sm text-gray-600 mt-1">مدیریت نسخه‌های مختلف اتصال حسابداری</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          @click="exportVersions"
          class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
          </svg>
          خروجی
        </button>
      </div>
    </div>

    <!-- فیلترها -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div>
          <input
            v-model="searchQuery"
            type="text"
            placeholder="جستجو در نسخه‌ها..."
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
        </div>
        <div>
          <select
            v-model="selectedType"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه انواع</option>
            <option value="manual">دستی</option>
            <option value="auto">خودکار</option>
            <option value="scheduled">برنامه‌ریزی شده</option>
          </select>
        </div>
        <div>
          <select
            v-model="selectedStatus"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه وضعیت‌ها</option>
            <option value="active">فعال</option>
            <option value="archived">آرشیو شده</option>
            <option value="deleted">حذف شده</option>
          </select>
        </div>
        <div>
          <button
            @click="clearFilters"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            پاک کردن فیلترها
          </button>
        </div>
      </div>
    </div>

    <!-- آمار نسخه‌ها -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div class="bg-blue-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-blue-600">{{ totalVersions }}</div>
        <div class="text-sm text-blue-700">کل نسخه‌ها</div>
      </div>
      <div class="bg-green-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-green-600">{{ activeVersions }}</div>
        <div class="text-sm text-green-700">نسخه‌های فعال</div>
      </div>
      <div class="bg-yellow-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-yellow-600">{{ archivedVersions }}</div>
        <div class="text-sm text-yellow-700">نسخه‌های آرشیو شده</div>
      </div>
      <div class="bg-purple-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-purple-600">{{ totalSize }}</div>
        <div class="text-sm text-purple-700">حجم کل (MB)</div>
      </div>
    </div>

    <!-- جدول نسخه‌ها -->
    <div class="bg-white rounded-lg border border-gray-200 overflow-hidden">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">لیست نسخه‌ها</h5>
      </div>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500">نسخه</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500">نوع</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500">تاریخ ایجاد</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500">حجم</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500">توضیحات</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="version in filteredVersions" :key="version.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="flex-shrink-0 w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center">
                    <svg class="w-4 h-4 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                    </svg>
                  </div>
                  <div class="mr-3">
                    <div class="text-sm font-medium text-gray-900">{{ version.version }}</div>
                    <div class="text-sm text-gray-500">{{ version.connectionName }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span
                  class="px-2 py-1 text-xs font-medium rounded-full"
                  :class="getTypeClass(version.type)"
                >
                  {{ getTypeLabel(version.type) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ version.createdAt }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ version.size }} MB
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span
                  class="px-2 py-1 text-xs font-medium rounded-full"
                  :class="getStatusClass(version.status)"
                >
                  {{ getStatusLabel(version.status) }}
                </span>
              </td>
              <td class="px-6 py-4 text-sm text-gray-900">
                <div class="max-w-xs truncate">{{ version.description }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <div class="flex items-center space-x-2 space-x-reverse">
                  <button
                    @click="viewVersion(version)"
                    class="text-blue-600 hover:text-blue-900"
                  >
                    مشاهده
                  </button>
                  <button
                    @click="downloadVersion(version)"
                    class="text-green-600 hover:text-green-900"
                  >
                    دانلود
                  </button>
                  <button
                    @click="deleteVersion(version)"
                    class="text-red-600 hover:text-red-900"
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

    <!-- صفحه‌بندی -->
    <div class="bg-white rounded-lg border border-gray-200 px-6 py-4">
      <div class="flex items-center justify-between">
        <div class="text-sm text-gray-700">
          نمایش {{ startIndex + 1 }} تا {{ endIndex }} از {{ totalVersions }} نسخه
        </div>
        <div class="flex items-center space-x-2 space-x-reverse">
          <button
            @click="previousPage"
            :disabled="currentPage === 1"
            class="px-3 py-1 border border-gray-300 rounded-md text-sm disabled:opacity-50 disabled:cursor-not-allowed"
          >
            قبلی
          </button>
          <span class="px-3 py-1 text-sm text-gray-700">
            صفحه {{ currentPage }} از {{ totalPages }}
          </span>
          <button
            @click="nextPage"
            :disabled="currentPage === totalPages"
            class="px-3 py-1 border border-gray-300 rounded-md text-sm disabled:opacity-50 disabled:cursor-not-allowed"
          >
            بعدی
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

// فیلترها
const searchQuery = ref('')
const selectedType = ref('')
const selectedStatus = ref('')

// صفحه‌بندی
const currentPage = ref(1)
const itemsPerPage = 10

// نسخه‌ها
const versions = ref([
  {
    id: 1,
    version: 'v2.1.0',
    connectionName: 'اتصال هلو',
    type: 'manual',
    createdAt: '۱۴۰۲/۱۰/۱۵ - ۱۴:۳۰',
    size: 45.2,
    status: 'active',
    description: 'نسخه پایدار با بهبودهای عملکرد'
  },
  {
    id: 2,
    version: 'v2.0.5',
    connectionName: 'اتصال هلو',
    type: 'auto',
    createdAt: '۱۴۰۲/۱۰/۱۰ - ۰۹:۱۵',
    size: 42.8,
    status: 'archived',
    description: 'نسخه خودکار قبل از تغییرات'
  },
  {
    id: 3,
    version: 'v2.0.0',
    connectionName: 'اتصال هلو',
    type: 'scheduled',
    createdAt: '۱۴۰۲/۰۹/۲۸ - ۱۶:۴۵',
    size: 38.5,
    status: 'archived',
    description: 'نسخه اصلی با قابلیت‌های جدید'
  },
  {
    id: 4,
    version: 'v1.9.8',
    connectionName: 'اتصال هلو',
    type: 'manual',
    createdAt: '۱۴۰۲/۰۹/۱۵ - ۱۱:۲۰',
    size: 35.2,
    status: 'deleted',
    description: 'آخرین نسخه از سری ۱.x'
  }
])

// فیلتر کردن نسخه‌ها
const filteredVersions = computed(() => {
  return versions.value.filter(version => {
    if (searchQuery.value && !version.version.includes(searchQuery.value) && !version.connectionName.includes(searchQuery.value)) return false
    if (selectedType.value && version.type !== selectedType.value) return false
    if (selectedStatus.value && version.status !== selectedStatus.value) return false
    return true
  })
})

// آمار
const totalVersions = computed(() => versions.value.length)
const activeVersions = computed(() => versions.value.filter(v => v.status === 'active').length)
const archivedVersions = computed(() => versions.value.filter(v => v.status === 'archived').length)
const totalSize = computed(() => versions.value.reduce((sum, v) => sum + v.size, 0).toFixed(1))

// صفحه‌بندی
const totalPages = computed(() => Math.ceil(filteredVersions.value.length / itemsPerPage))
const startIndex = computed(() => (currentPage.value - 1) * itemsPerPage)
const endIndex = computed(() => Math.min(startIndex.value + itemsPerPage, filteredVersions.value.length))

// کلاس نوع
const getTypeClass = (type: string) => {
  const classes = {
    manual: 'bg-blue-100 text-blue-700',
    auto: 'bg-green-100 text-green-700',
    scheduled: 'bg-yellow-100 text-yellow-700'
  }
  return classes[type] || 'bg-gray-100 text-gray-700'
}

// برچسب نوع
const getTypeLabel = (type: string) => {
  const labels = {
    manual: 'دستی',
    auto: 'خودکار',
    scheduled: 'برنامه‌ریزی شده'
  }
  return labels[type] || type
}

// کلاس وضعیت
const getStatusClass = (status: string) => {
  const classes = {
    active: 'bg-green-100 text-green-700',
    archived: 'bg-yellow-100 text-yellow-700',
    deleted: 'bg-red-100 text-red-700'
  }
  return classes[status] || 'bg-gray-100 text-gray-700'
}

// برچسب وضعیت
const getStatusLabel = (status: string) => {
  const labels = {
    active: 'فعال',
    archived: 'آرشیو شده',
    deleted: 'حذف شده'
  }
  return labels[status] || status
}

// عملیات
const viewVersion = (version: any) => {
  console.log('مشاهده نسخه:', version)
}

const downloadVersion = (version: any) => {
  console.log('دانلود نسخه:', version)
}

const deleteVersion = (version: any) => {
  console.log('حذف نسخه:', version)
}

const exportVersions = () => {
  console.log('خروجی نسخه‌ها')
}

const clearFilters = () => {
  searchQuery.value = ''
  selectedType.value = ''
  selectedStatus.value = ''
}

const previousPage = () => {
  if (currentPage.value > 1) {
    currentPage.value--
  }
}

const nextPage = () => {
  if (currentPage.value < totalPages.value) {
    currentPage.value++
  }
}
</script>

<!--
  کامپوننت نگهداری نسخه‌ها
  شامل:
  1. لیست نسخه‌ها
  2. فیلتر و جستجو
  3. آمار نسخه‌ها
  4. عملیات مدیریتی
  5. صفحه‌بندی
  6. طراحی مدرن و کاملاً ریسپانسیو
--> 
