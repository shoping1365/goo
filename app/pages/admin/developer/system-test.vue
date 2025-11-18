<template>
  <div class="space-y-6">
    <div class="bg-white rounded-lg shadow p-6">
      <h1 class="text-2xl font-bold mb-6">تست سیستم</h1>

      <!-- مدیریت سیستم -->
      <div class="mb-8 flex flex-wrap gap-6">
        <button @click="clearCache" class="bg-yellow-500 hover:bg-yellow-600 text-white px-4 py-2 rounded shadow">پاک کردن کامل کش</button>
        <button @click="resetDatabase" class="bg-red-500 hover:bg-red-600 text-white px-4 py-2 rounded shadow">ریست دیتابیس</button>
        <button @click="restartServer" class="bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded shadow">ریست سرور اصلی</button>
        <button @click="clearAllLogs" class="bg-gray-700 hover:bg-gray-900 text-white px-4 py-2 rounded shadow">پاک کردن تمام لاگ‌ها</button>
      </div>

        <!-- API Connections -->
        <div class="mb-8">
          <h2 class="text-xl font-semibold mb-4">وضعیت اتصال‌های API</h2>
          <div v-if="loading.api" class="text-center py-4">
            <span>در حال بارگذاری...</span>
          </div>
          <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            <div v-for="api in apiConnections" :key="api.id"
                 class="border rounded-lg p-6"
                 :class="{
                   'border-green-500 bg-green-50': api.status === 'success',
                   'border-red-500 bg-red-50': api.status === 'error',
                   'border-yellow-500 bg-yellow-50': api.status === 'warning'
                 }">
              <div class="flex items-center justify-between">
                <div>
                  <h3 class="font-medium">{{ api.name }}</h3>
                  <p class="text-sm text-gray-600">{{ api.url }}</p>
                </div>
                <div class="flex items-center">
                  <span class="text-sm"
                        :class="{
                          'text-green-600': api.status === 'success',
                          'text-red-600': api.status === 'error',
                          'text-yellow-600': api.status === 'warning'
                        }">
                    {{ api.status === 'success' ? 'متصل' : api.status === 'error' ? 'خطا' : 'هشدار' }}
                  </span>
                  <button @click="testConnection(api)" class="mr-2 text-blue-600 hover:text-blue-800" title="تست اتصال">
                    🔄
                  </button>
                </div>
              </div>
              <div v-if="api.lastChecked" class="mt-2 text-xs text-gray-500">
                آخرین بررسی: {{ formatDate(api.lastChecked) }}
              </div>
            </div>
          </div>
        </div>

        <!-- Database Status -->
        <div class="mb-8">
          <h2 class="text-xl font-semibold mb-4">وضعیت پایگاه داده</h2>
          <div v-if="loading.database" class="text-center py-4">
            <span>در حال بارگذاری...</span>
          </div>
          <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            <div v-for="db in databaseStatus" :key="db.id"
                 class="border rounded-lg p-6"
                 :class="{
                   'border-green-500 bg-green-50': db.status === 'success',
                   'border-red-500 bg-red-50': db.status === 'error'
                 }">
              <div class="flex items-center justify-between">
                <div>
                  <h3 class="font-medium">{{ db.name }}</h3>
                  <p class="text-sm text-gray-600">{{ db.type }}</p>
                </div>
                <div class="flex items-center">
                  <span class="text-sm"
                        :class="{
                          'text-green-600': db.status === 'success',
                          'text-red-600': db.status === 'error'
                        }">
                    {{ db.status === 'success' ? 'فعال' : 'خطا' }}
                  </span>
                  <button @click="testDatabase(db)" class="mr-2 text-blue-600 hover:text-blue-800" title="تست اتصال">
                    🔄
                  </button>
                </div>
              </div>
              <div class="mt-2">
                <div class="text-sm text-gray-600">حجم: {{ db.size }}</div>
                <div class="text-sm text-gray-600">تعداد جداول: {{ db.tables }}</div>
              </div>
            </div>
          </div>
        </div>

        <!-- Server Status -->
        <div class="mb-8">
          <h2 class="text-xl font-semibold mb-4">وضعیت سرور</h2>
          <div v-if="loading.server" class="text-center py-4">
            <span>در حال بارگذاری...</span>
          </div>
          <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            <div v-for="server in serverStatus" :key="server.id"
                 class="border rounded-lg p-6"
                 :class="{
                   'border-green-500 bg-green-50': server.status === 'success',
                   'border-red-500 bg-red-50': server.status === 'error'
                 }">
              <div class="flex items-center justify-between">
                <div>
                  <h3 class="font-medium">{{ server.name }}</h3>
                  <p class="text-sm text-gray-600">{{ server.ip }}</p>
                </div>
                <div class="flex items-center">
                  <span class="text-sm"
                        :class="{
                          'text-green-600': server.status === 'success',
                          'text-red-600': server.status === 'error'
                        }">
                    {{ server.status === 'success' ? 'فعال' : 'خطا' }}
                  </span>
                  <button @click="testServer(server)" class="mr-2 text-blue-600 hover:text-blue-800" title="تست اتصال">
                    🔄
                  </button>
                </div>
              </div>
              <div class="mt-2">
                <div class="text-sm text-gray-600">پردازنده: {{ server.cpu }}%</div>
                <div class="text-sm text-gray-600">حافظه: {{ server.ram }}%</div>
                <div class="text-sm text-gray-600">فضای ذخیره‌سازی: {{ server.disk }}%</div>
              </div>
            </div>
          </div>
        </div>

        <!-- System Logs -->
        <div>
          <h2 class="text-xl font-semibold mb-4">لاگ‌های سیستم</h2>
          <div class="border rounded-lg overflow-hidden">
            <div class="bg-gray-50 px-4 py-2 border-b flex justify-between items-center">
              <div class="flex space-x-4 space-x-reverse">
                <button @click="refreshLogs" class="text-blue-600 hover:text-blue-800" title="بروزرسانی لاگ‌ها">🔄</button>
                <button @click="clearLogs" class="text-red-600 hover:text-red-800" title="پاک کردن لاگ‌ها">🗑️</button>
              </div>
              <div class="flex items-center space-x-4 space-x-reverse">
                <select v-model="logLevel" class="border rounded px-2 py-1 text-sm">
                  <option value="all">همه</option>
                  <option value="error">خطا</option>
                  <option value="warning">هشدار</option>
                  <option value="info">اطلاعات</option>
                </select>
                <input type="text" v-model="logSearch" placeholder="جستجو در لاگ‌ها..." class="border rounded px-2 py-1 text-sm w-64">
              </div>
            </div>
            <div v-if="loading.logs" class="text-center py-4">
              <span>در حال بارگذاری...</span>
            </div>
            <div v-else class="h-96 overflow-y-auto">
              <table class="min-w-full divide-y divide-gray-200">
                <thead class="bg-gray-50">
                  <tr>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">زمان</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">سطح</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">پیام</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">جزئیات</th>
                  </tr>
                </thead>
                <tbody class="bg-white divide-y divide-gray-200">
                  <tr v-for="log in filteredLogs" :key="log.id"
                      :class="{
                        'bg-red-50': log.level === 'error',
                        'bg-yellow-50': log.level === 'warning',
                        'bg-blue-50': log.level === 'info'
                      }">
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ formatDate(log.timestamp) }}</td>
                    <td class="px-6 py-4 whitespace-nowrap">
                      <span class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full"
                            :class="{
                              'bg-red-100 text-red-800': log.level === 'error',
                              'bg-yellow-100 text-yellow-800': log.level === 'warning',
                              'bg-blue-100 text-blue-800': log.level === 'info'
                            }">
                        {{ log.level === 'error' ? 'خطا' : log.level === 'warning' ? 'هشدار' : 'اطلاعات' }}
                      </span>
                    </td>
                    <td class="px-6 py-4 text-sm text-gray-500">{{ log.message }}</td>
                    <td class="px-6 py-4 text-sm text-gray-500">
                      <button @click="showLogDetails(log)" class="text-blue-600 hover:text-blue-800">
                        مشاهده جزئیات
                      </button>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
    </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'

definePageMeta({
  layout: 'admin-main',
  middleware: ['developer-only']
});

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user, hasPermission } = useAuth()

const apiConnections = ref([])
const databaseStatus = ref([])
const serverStatus = ref([])
const logs = ref([])
const logLevel = ref('all')
const logSearch = ref('')
const loading = ref({
  api: false,
  database: false,
  server: false,
  logs: false
})

const fetchApiStatus = async () => {
  loading.value.api = true
  try {
    const response = await fetch('/api/admin/system/api-status')
    apiConnections.value = await response.json()
  } catch (e) { apiConnections.value = [] }
  loading.value.api = false
}
const fetchDatabaseStatus = async () => {
  loading.value.database = true
  try {
    const response = await fetch('/api/admin/system/database-status')
    databaseStatus.value = await response.json()
  } catch (e) { databaseStatus.value = [] }
  loading.value.database = false
}
const fetchServerStatus = async () => {
  loading.value.server = true
  try {
    const response = await fetch('/api/admin/system/server-status')
    serverStatus.value = await response.json()
  } catch (e) { serverStatus.value = [] }
  loading.value.server = false
}
const fetchLogs = async () => {
  loading.value.logs = true
  try {
    const response = await fetch('/api/admin/system/logs')
    logs.value = await response.json()
  } catch (e) { logs.value = [] }
  loading.value.logs = false
}

const formatDate = (date) => {
  return new Intl.DateTimeFormat('fa-IR', {
    year: 'numeric', month: 'long', day: 'numeric',
    hour: '2-digit', minute: '2-digit', second: '2-digit'
  }).format(new Date(date))
}

const testConnection = async (api) => {
  await fetch(`/api/admin/system/test-api/${api.id}`)
  fetchApiStatus()
}
const testDatabase = async (db) => {
  await fetch(`/api/admin/system/test-database/${db.id}`)
  fetchDatabaseStatus()
}
const testServer = async (server) => {
  await fetch(`/api/admin/system/test-server/${server.id}`)
  fetchServerStatus()
}
const refreshLogs = async () => { await fetchLogs() }
const clearLogs = async () => {
  await fetch('/api/admin/system/clear-logs', { method: 'POST' })
  await fetchLogs()
}
const showLogDetails = (log) => {
  alert(log.details)
}

const clearCache = async () => {
  if (confirm('آیا از پاک کردن کامل کش مطمئن هستید؟')) {
    await fetch('/api/admin/system/clear-cache', { method: 'POST' })
    alert('کش پاک شد!')
  }
}
const resetDatabase = async () => {
  if (confirm('آیا از ریست دیتابیس مطمئن هستید؟')) {
    await fetch('/api/admin/system/reset-database', { method: 'POST' })
    alert('دیتابیس ریست شد!')
  }
}
const restartServer = async () => {
  if (confirm('آیا از ریست سرور مطمئن هستید؟')) {
    await fetch('/api/admin/system/restart-server', { method: 'POST' })
    alert('درخواست ریست سرور ارسال شد!')
  }
}

const clearAllLogs = async () => {
  if (confirm('آیا از پاک کردن تمام لاگ‌ها مطمئن هستید؟')) {
    await fetch('/api/admin/system/clear-logs', { method: 'POST' })
    await fetchLogs()
    alert('تمام لاگ‌ها پاک شد!')
  }
}

const filteredLogs = computed(() => {
  return logs.value.filter(log => {
    const matchesLevel = logLevel.value === 'all' || log.level === logLevel.value
    const matchesSearch = log.message?.toLowerCase().includes(logSearch.value.toLowerCase()) ||
      (log.details && log.details.toLowerCase().includes(logSearch.value.toLowerCase()))
    return matchesLevel && matchesSearch
  })
})

onMounted(async () => {
  await Promise.all([
    fetchApiStatus(),
    fetchDatabaseStatus(),
    fetchServerStatus(),
    fetchLogs()
  ])
})
</script> 
