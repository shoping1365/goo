<template>
  <div class="min-h-screen bg-gray-50 p-6">
    <div class="max-w-7xl mx-auto">
      <!-- Header -->
      <div class="mb-8">
        <h1 class="text-3xl font-bold text-gray-900 mb-2">مدیریت پشتیبان‌گیری</h1>
        <p class="text-gray-600">مدیریت و ایجاد پشتیبان‌گیری از دیتابیس، فایل‌ها و تنظیمات سیستم</p>
      </div>

      <!-- Stats Cards -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-blue-100 rounded-lg">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-8l-4-4m0 0L8 8m4-4v12"></path>
              </svg>
            </div>
            <div class="mr-4">
              <p class="text-sm font-medium text-gray-600">کل پشتیبان‌ها</p>
              <p class="text-2xl font-bold text-gray-900">{{ stats.totalBackups }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-green-100 rounded-lg">
              <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
            <div class="mr-4">
              <p class="text-sm font-medium text-gray-600">موفق</p>
              <p class="text-2xl font-bold text-gray-900">{{ stats.successfulBackups }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-yellow-100 rounded-lg">
              <svg class="w-6 h-6 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
            <div class="mr-4">
              <p class="text-sm font-medium text-gray-600">در حال اجرا</p>
              <p class="text-2xl font-bold text-gray-900">{{ stats.runningBackups }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-red-100 rounded-lg">
              <svg class="w-6 h-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
              </svg>
            </div>
            <div class="mr-4">
              <p class="text-sm font-medium text-gray-600">ناموفق</p>
              <p class="text-2xl font-bold text-gray-900">{{ stats.failedBackups }}</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Quick Actions -->
      <div class="bg-white rounded-lg shadow mb-8">
        <div class="px-6 py-4 border-b border-gray-200">
          <h2 class="text-lg font-semibold text-gray-900">عملیات سریع</h2>
        </div>
        <div class="p-6">
          <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
            <button
:disabled="isCreatingBackup" class="flex items-center justify-center px-4 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed transition-colors" 
                    @click="createBackup('database')">
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-8l-4-4m0 0L8 8m4-4v12"></path>
              </svg>
              پشتیبان‌گیری دیتابیس
            </button>

            <button
:disabled="isCreatingBackup" class="flex items-center justify-center px-4 py-3 bg-green-600 text-white rounded-lg hover:bg-green-700 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
                    @click="createBackup('files')">
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2H5a2 2 0 00-2-2z"></path>
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 5a2 2 0 012-2h4a2 2 0 012 2v2H8V5z"></path>
              </svg>
              پشتیبان‌گیری فایل‌ها
            </button>

            <button
:disabled="isCreatingBackup" class="flex items-center justify-center px-4 py-3 bg-purple-600 text-white rounded-lg hover:bg-purple-700 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
                    @click="createBackup('full')">
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                
              </svg>
              پشتیبان‌گیری کامل
            </button>
          </div>
        </div>
      </div>

      <!-- Backup Configuration -->
      <div class="bg-white rounded-lg shadow mb-8">
        <div class="px-6 py-4 border-b border-gray-200">
          <h2 class="text-lg font-semibold text-gray-900">تنظیمات پشتیبان‌گیری</h2>
        </div>
        <div class="p-6">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <!-- Database Backup Settings -->
            <div>
              <h3 class="text-md font-medium text-gray-900 mb-4">تنظیمات دیتابیس</h3>
              <div class="space-y-4">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">فرکانس پشتیبان‌گیری</label>
                  <select v-model="config.database.frequency" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
                    <option value="daily">روزانه</option>
                    <option value="weekly">هفتگی</option>
                    <option value="monthly">ماهانه</option>
                  </select>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">نگهداری (روز)</label>
                  <input
v-model="config.database.retention" type="number" min="1" max="365" 
                         class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
                </div>
                <div class="flex items-center">
                  <input
id="db-compression" v-model="config.database.compression" type="checkbox" 
                         class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded">
                  <label for="db-compression" class="mr-2 text-sm text-gray-700">فشرده‌سازی</label>
                </div>
              </div>
            </div>

            <!-- File Backup Settings -->
            <div>
              <h3 class="text-md font-medium text-gray-900 mb-4">تنظیمات فایل‌ها</h3>
              <div class="space-y-4">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">مسیرهای شامل</label>
                  <textarea
v-model="config.files.includePaths" rows="3" 
                            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                            placeholder="/uploads&#10;/public&#10;/config"></textarea>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">مسیرهای مستثنی</label>
                  <textarea
v-model="config.files.excludePaths" rows="3" 
                            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                            placeholder="/uploads/temp&#10;/logs&#10;*.tmp"></textarea>
                </div>
                <div class="flex items-center">
                  <input
id="file-incremental" v-model="config.files.incremental" type="checkbox" 
                         class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded">
                  <label for="file-incremental" class="mr-2 text-sm text-gray-700">پشتیبان‌گیری افزایشی</label>
                </div>
              </div>
            </div>
          </div>

          <div class="mt-6 flex justify-end">
            <button
class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors" 
                    @click="saveConfiguration">
              ذخیره تنظیمات
            </button>
          </div>
        </div>
      </div>

      <!-- Backup History -->
      <div class="bg-white rounded-lg shadow">
        <div class="px-6 py-4 border-b border-gray-200 flex justify-between items-center">
          <h2 class="text-lg font-semibold text-gray-900">تاریخچه پشتیبان‌گیری</h2>
          <div class="flex space-x-2 space-x-reverse">
            <select v-model="filter.status" class="px-3 py-1 border border-gray-300 rounded-md text-sm">
              <option value="">همه وضعیت‌ها</option>
              <option value="completed">موفق</option>
              <option value="failed">ناموفق</option>
              <option value="running">در حال اجرا</option>
            </select>
            <select v-model="filter.type" class="px-3 py-1 border border-gray-300 rounded-md text-sm">
              <option value="">همه انواع</option>
              <option value="database">دیتابیس</option>
              <option value="files">فایل‌ها</option>
              <option value="full">کامل</option>
            </select>
          </div>
        </div>
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">حجم</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="backup in filteredBackups" :key="backup.id" class="hover:bg-gray-50">
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="flex items-center">
                    <span
class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                          :class="getTypeBadgeClass(backup.type)">
                      {{ getTypeLabel(backup.type) }}
                    </span>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  {{ formatDate(backup.createdAt) }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  {{ formatSize(backup.size) }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span
class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                        :class="getStatusBadgeClass(backup.status)">
                    {{ getStatusLabel(backup.status) }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                  <div class="flex space-x-2 space-x-reverse">
                    <button
class="text-blue-600 hover:text-blue-900" 
                            @click="downloadBackup(backup.id)">دانلود</button>
                    <button
class="text-green-600 hover:text-green-900" 
                            @click="restoreBackup(backup.id)">بازیابی</button>
                    <button
class="text-red-600 hover:text-red-900" 
                            @click="deleteBackup(backup.id)">حذف</button>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
definePageMeta({
  layout: 'admin-main',
  middleware: ['developer-only']
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
// const { user, hasPermission } = useAuth()

// Reactive data
const stats = ref({
  totalBackups: 24,
  successfulBackups: 20,
  runningBackups: 1,
  failedBackups: 3
})

const isCreatingBackup = ref(false)

const config = ref({
  database: {
    frequency: 'daily',
    retention: 30,
    compression: true
  },
  files: {
    includePaths: '/uploads\n/public\n/config',
    excludePaths: '/uploads/temp\n/logs\n*.tmp',
    incremental: true
  }
})

const filter = ref({
  status: '',
  type: ''
})

const backups = ref([
  {
    id: 1,
    type: 'database',
    createdAt: new Date('2024-01-15T10:30:00'),
    size: 1024 * 1024 * 50, // 50MB
    status: 'completed'
  },
  {
    id: 2,
    type: 'files',
    createdAt: new Date('2024-01-14T15:45:00'),
    size: 1024 * 1024 * 200, // 200MB
    status: 'completed'
  },
  {
    id: 3,
    type: 'full',
    createdAt: new Date('2024-01-13T08:20:00'),
    size: 1024 * 1024 * 500, // 500MB
    status: 'failed'
  },
  {
    id: 4,
    type: 'database',
    createdAt: new Date('2024-01-12T22:10:00'),
    size: 1024 * 1024 * 45, // 45MB
    status: 'running'
  }
])

// Computed
const filteredBackups = computed(() => {
  return backups.value.filter(backup => {
    if (filter.value.status && backup.status !== filter.value.status) return false
    if (filter.value.type && backup.type !== filter.value.type) return false
    return true
  })
})

// Methods
const createBackup = async (type) => {
  isCreatingBackup.value = true
  try {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 2000))
    
    const newBackup = {
      id: Date.now(),
      type,
      createdAt: new Date(),
      size: 0,
      status: 'running'
    }
    
    backups.value.unshift(newBackup)
    stats.value.runningBackups++
    
    // Simulate backup completion
    setTimeout(() => {
      newBackup.status = 'completed'
      newBackup.size = Math.random() * 1024 * 1024 * 100 + 1024 * 1024 * 10
      stats.value.runningBackups--
      stats.value.successfulBackups++
    }, 5000)
    
  } catch (error) {
    console.error('Error creating backup:', error)
  } finally {
    isCreatingBackup.value = false
  }
}

const saveConfiguration = async () => {
  try {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 1000))

  } catch (error) {
    console.error('Error saving configuration:', error)
  }
}

const downloadBackup = (_id) => {

}

const restoreBackup = (_id) => {

}

const deleteBackup = (id) => {
  const index = backups.value.findIndex(b => b.id === id)
  if (index > -1) {
    backups.value.splice(index, 1)
  }
}

const getTypeLabel = (type) => {
  const labels = {
    database: 'دیتابیس',
    files: 'فایل‌ها',
    full: 'کامل'
  }
  return labels[type] || type
}

const getTypeBadgeClass = (type) => {
  const classes = {
    database: 'bg-blue-100 text-blue-800',
    files: 'bg-green-100 text-green-800',
    full: 'bg-purple-100 text-purple-800'
  }
  return classes[type] || 'bg-gray-100 text-gray-800'
}

const getStatusLabel = (status) => {
  const labels = {
    completed: 'موفق',
    failed: 'ناموفق',
    running: 'در حال اجرا'
  }
  return labels[status] || status
}

const getStatusBadgeClass = (status) => {
  const classes = {
    completed: 'bg-green-100 text-green-800',
    failed: 'bg-red-100 text-red-800',
    running: 'bg-yellow-100 text-yellow-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const formatDate = (date) => {
  return new Intl.DateTimeFormat('fa-IR', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  }).format(date)
}

const formatSize = (bytes) => {
  const sizes = ['B', 'KB', 'MB', 'GB']
  let i = 0
  let size = bytes
  while (size >= 1024 && i < sizes.length - 1) {
    size /= 1024
    i++
  }
  return `${size.toFixed(1)} ${sizes[i]}`
}
</script> 
