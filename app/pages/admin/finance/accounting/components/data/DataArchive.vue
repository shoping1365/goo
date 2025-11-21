<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">آرشیو داده‌ها</h4>
        <p class="text-sm text-gray-600 mt-1">مدیریت آرشیو و نگهداری داده‌های قدیمی</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          @click="createArchive"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 8h14M5 8a2 2 0 110-4h14a2 2 0 110 4M5 8v10a2 2 0 002 2h10a2 2 0 002-2V8m-9 4h4" />
          </svg>
          ایجاد آرشیو
        </button>
      </div>
    </div>

    <!-- آمار آرشیو -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">آمار آرشیو</h5>
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">کل آرشیوها</h6>
            <div class="w-3 h-3 bg-blue-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-blue-600">{{ archiveStats.total }}</div>
          <div class="text-xs text-gray-500 mt-1">فایل</div>
        </div>

        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">حجم کل آرشیو</h6>
            <div class="w-3 h-3 bg-green-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-green-600">{{ archiveStats.totalSize }}</div>
          <div class="text-xs text-gray-500 mt-1">گیگابایت</div>
        </div>

        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">آخرین آرشیو</h6>
            <div class="w-3 h-3 bg-yellow-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-yellow-600">{{ archiveStats.lastArchive }}</div>
          <div class="text-xs text-gray-500 mt-1">روز پیش</div>
        </div>

        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">وضعیت</h6>
            <div class="w-3 h-3 bg-green-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-green-600">{{ archiveStats.status }}</div>
          <div class="text-xs text-gray-500 mt-1">سالم</div>
        </div>
      </div>
    </div>

    <!-- لیست آرشیوها -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">آرشیوهای موجود</h5>
      </div>
      <div class="p-6">
        <div class="space-y-4">
          <div
            v-for="archive in archives"
            :key="archive.id"
            class="flex items-center justify-between p-6 border border-gray-200 rounded-lg hover:bg-gray-50"
          >
            <div class="flex items-center">
              <div
                class="w-3 h-3 rounded-full ml-3"
                :class="getArchiveStatusColor(archive.status)"
              ></div>
              <div>
                <h6 class="text-sm font-medium text-gray-900">{{ archive.name }}</h6>
                <p class="text-xs text-gray-500">{{ archive.description }}</p>
              </div>
            </div>
            <div class="flex items-center space-x-4 space-x-reverse">
              <div class="text-right">
                <div class="text-sm font-medium text-gray-900">{{ archive.size }}</div>
                <div class="text-xs text-gray-500">حجم</div>
              </div>
              <div class="text-right">
                <div class="text-sm font-medium text-gray-900">{{ archive.date }}</div>
                <div class="text-xs text-gray-500">تاریخ</div>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <button
                  class="text-blue-600 hover:text-blue-800"
                  @click="restoreArchive(archive)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
                  </svg>
                </button>
                <button
                  class="text-green-600 hover:text-green-800"
                  @click="downloadArchive(archive)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                  </svg>
                </button>
                <button
                  v-if="canDeleteArchive"
                  class="text-red-600 hover:text-red-800"
                  @click="deleteArchive(archive)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تاریخچه آرشیو -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">تاریخچه آرشیو</h5>
      </div>
      <div class="p-6">
        <div class="space-y-4">
          <div
            v-for="history in archiveHistory"
            :key="history.id"
            class="flex items-start space-x-4 space-x-reverse p-6 border border-gray-200 rounded-lg"
          >
            <div
              class="w-2 h-2 rounded-full mt-2"
              :class="getHistoryColor(history.status)"
            ></div>
            <div class="flex-1">
              <div class="flex items-center justify-between">
                <h6 class="text-sm font-medium text-gray-900">{{ history.archiveName }}</h6>
                <span class="text-xs text-gray-500">{{ history.timestamp }}</span>
              </div>
              <p class="text-sm text-gray-600 mt-1">{{ history.description }}</p>
              <div class="flex items-center space-x-4 space-x-reverse mt-2 text-xs text-gray-500">
                <span>نوع: {{ getArchiveTypeText(history.type) }}</span>
                <span>حجم: {{ history.size }}</span>
                <span>کاربر: {{ history.user }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات آرشیو -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">تنظیمات آرشیو</h5>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">سن داده‌ها برای آرشیو</label>
          <select
            v-model="archiveSettings.ageThreshold"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="90">۹۰ روز</option>
            <option value="180">۱۸۰ روز</option>
            <option value="365">۱ سال</option>
            <option value="730">۲ سال</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">فرکانس آرشیو</label>
          <select
            v-model="archiveSettings.frequency"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="monthly">ماهانه</option>
            <option value="quarterly">فصلی</option>
            <option value="yearly">سالانه</option>
            <option value="manual">دستی</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">فشرده‌سازی</label>
          <select
            v-model="archiveSettings.compression"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="none">بدون فشرده‌سازی</option>
            <option value="gzip">GZIP</option>
            <option value="zip">ZIP</option>
            <option value="tar">TAR</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">محل ذخیره</label>
          <select
            v-model="archiveSettings.storageLocation"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="local">محلی</option>
            <option value="cloud">ابر</option>
            <option value="external">خارجی</option>
          </select>
        </div>
      </div>

      <div class="mt-4 space-y-3">
        <label class="flex items-center">
          <input
            v-model="archiveSettings.autoArchive"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">آرشیو خودکار</span>
        </label>

        <label class="flex items-center">
          <input
            v-model="archiveSettings.encryptArchive"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">رمزگذاری آرشیوها</span>
        </label>

        <label class="flex items-center">
          <input
            v-model="archiveSettings.verifyIntegrity"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">بررسی یکپارچگی</span>
        </label>

        <label class="flex items-center">
          <input
            v-model="archiveSettings.notifyOnArchive"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">اعلان پس از آرشیو</span>
        </label>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import { useAuth } from '~/composables/useAuth'

const { user: _user, hasPermission } = useAuth()
const canDeleteArchive = computed(() => hasPermission('data-archive.delete'))

// آمار آرشیو
const archiveStats = ref({
  total: 12,
  totalSize: 3.2,
  lastArchive: 15,
  status: 'سالم'
})

// آرشیوهای موجود
const archives = ref([
  {
    id: 1,
    name: 'آرشیو تراکنش‌های ۱۴۰۱',
    description: 'آرشیو کامل تراکنش‌های سال ۱۴۰۱',
    status: 'success',
    size: '450MB',
    date: '۱۵ دی ۱۴۰۲'
  },
  {
    id: 2,
    name: 'آرشیو لاگ‌های سیستم',
    description: 'آرشیو لاگ‌های سیستم با سن بیش از ۱ سال',
    status: 'success',
    size: '120MB',
    date: '۱۰ دی ۱۴۰۲'
  },
  {
    id: 3,
    name: 'آرشیو گزارش‌های قدیمی',
    description: 'آرشیو گزارش‌های مالی قدیمی',
    status: 'warning',
    size: '85MB',
    date: '۵ دی ۱۴۰۲'
  }
])

// تاریخچه آرشیو
const archiveHistory = ref([
  {
    id: 1,
    archiveName: 'آرشیو تراکنش‌های ۱۴۰۱',
    description: 'آرشیو خودکار تراکنش‌های سال ۱۴۰۱',
    type: 'auto',
    size: '450MB',
    status: 'success',
    user: 'سیستم',
    timestamp: '۱۵ دی ۱۴۰۲'
  },
  {
    id: 2,
    archiveName: 'آرشیو لاگ‌های سیستم',
    description: 'آرشیو دستی لاگ‌های سیستم',
    type: 'manual',
    size: '120MB',
    status: 'success',
    user: 'مدیر سیستم',
    timestamp: '۱۰ دی ۱۴۰۲'
  }
])

// تنظیمات آرشیو
const archiveSettings = ref({
  ageThreshold: 365,
  frequency: 'monthly',
  compression: 'gzip',
  storageLocation: 'local',
  autoArchive: true,
  encryptArchive: true,
  verifyIntegrity: true,
  notifyOnArchive: true
})

// رنگ‌های وضعیت
const getArchiveStatusColor = (status: string) => {
  const colors = {
    success: 'bg-green-500',
    warning: 'bg-yellow-500',
    error: 'bg-red-500'
  }
  return colors[status] || 'bg-gray-500'
}

const getHistoryColor = (status: string) => {
  const colors = {
    success: 'bg-green-500',
    warning: 'bg-yellow-500',
    error: 'bg-red-500'
  }
  return colors[status] || 'bg-gray-500'
}

const getArchiveTypeText = (type: string) => {
  const texts = {
    auto: 'خودکار',
    manual: 'دستی'
  }
  return texts[type] || type
}

interface Archive {
  id: number | string
  name?: string
  [key: string]: unknown
}

// عملیات
const createArchive = () => {

}

const restoreArchive = (_archive: Archive) => {

}

const downloadArchive = (_archive: Archive) => {

}

const deleteArchive = (_archive: Archive) => {

}

const _bulkDelete = () => {
  if (!canDeleteArchive.value) {
    alert('شما مجوز حذف آرشیو را ندارید.');
    return;
  }
  // ... existing code ...
}
</script>

<!--
  کامپوننت آرشیو داده‌ها
  شامل:
  1. آمار آرشیو
  2. لیست آرشیوها
  3. تاریخچه آرشیو
  4. تنظیمات آرشیو
  5. طراحی مدرن و کاملاً ریسپانسیو
--> 

