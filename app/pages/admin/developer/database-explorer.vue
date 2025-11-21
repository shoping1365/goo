<template>
  <div class="min-h-screen bg-gray-50 p-6">
    <div class="max-w-7xl mx-auto">
      <!-- Header -->
      <div class="mb-8">
        <h1 class="text-3xl font-bold text-gray-900 mb-2">کاوشگر دیتابیس</h1>
        <p class="text-gray-600">ابزارهای مدیریت و کاوش دیتابیس</p>
      </div>

      <!-- Stats Cards -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-blue-100 rounded-lg">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 7v10c0 2.21 3.582 4 8 4s8-1.79 8-4V7M4 7c0 2.21 3.582 4 8 4s8-1.79 8-4M4 7c0-2.21 3.582-4 8-4s8 1.79 8 4"></path>
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">تعداد جداول</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.tables }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-green-100 rounded-lg">
              <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">کل رکوردها</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.totalRecords }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-purple-100 rounded-lg">
              <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">حجم دیتابیس</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.databaseSize }}</p>
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
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">زمان پاسخ</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.responseTime }}ms</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Main Content Grid -->
      <div class="grid grid-cols-1 lg:grid-cols-4 gap-8">
        <!-- Database Schema -->
        <div class="bg-white rounded-lg shadow">
          <div class="p-6 border-b border-gray-200">
            <h2 class="text-xl font-semibold text-gray-900">ساختار دیتابیس</h2>
          </div>
          <div class="p-6">
            <div class="space-y-4">
              <div v-for="table in databaseSchema" :key="table.name" class="space-y-2">
                <div 
                  :class="[
                    'flex items-center justify-between p-3 rounded-lg cursor-pointer hover:bg-gray-50',
                    selectedTable?.name === table.name ? 'bg-blue-50 border border-blue-200' : 'border border-gray-200'
                  ]"
                  @click="selectTable(table)"
                >
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <svg class="w-5 h-5 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"></path>
                    </svg>
                    <span class="font-medium text-gray-900">{{ table.name }}</span>
                  </div>
                  <span class="text-sm text-gray-500">{{ table.records }} رکورد</span>
                </div>
                
                <!-- Table Columns -->
                <div v-if="selectedTable?.name === table.name" class="ml-6 space-y-1">
                  <div v-for="column in table.columns" :key="column.name" class="flex items-center justify-between text-sm">
                    <span class="text-gray-700">{{ column.name }}</span>
                    <span class="text-gray-500">{{ column.type }}</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Query Editor -->
        <div class="lg:col-span-3 bg-white rounded-lg shadow">
          <div class="p-6 border-b border-gray-200">
            <div class="flex items-center justify-between">
              <h2 class="text-xl font-semibold text-gray-900">ویرایشگر کوئری</h2>
              <div class="flex items-center space-x-2 space-x-reverse">
                <button 
                  :disabled="!queryText"
                  class="bg-blue-600 hover:bg-blue-700 disabled:bg-gray-400 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="executeQuery"
                >
                  اجرا
                </button>
                <button 
                  class="bg-gray-500 hover:bg-gray-600 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="clearQuery"
                >
                  پاک کردن
                </button>
              </div>
            </div>
          </div>
          <div class="p-6">
            <div class="space-y-4">
              <!-- Query Input -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">کوئری SQL</label>
                <textarea
                  v-model="queryText"
                  rows="6"
                  class="w-full border border-gray-300 rounded-lg px-3 py-2 font-mono text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
                  placeholder="SELECT * FROM users LIMIT 10;"
                ></textarea>
              </div>

              <!-- Quick Queries -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">کوئری‌های سریع</label>
                <div class="grid grid-cols-2 md:grid-cols-4 gap-2">
                  <button 
                    class="bg-gray-100 hover:bg-gray-200 text-gray-700 font-medium py-2 px-3 rounded text-sm transition-colors"
                    @click="loadQuickQuery('select')"
                  >
                    SELECT
                  </button>
                  <button 
                    class="bg-gray-100 hover:bg-gray-200 text-gray-700 font-medium py-2 px-3 rounded text-sm transition-colors"
                    @click="loadQuickQuery('insert')"
                  >
                    INSERT
                  </button>
                  <button 
                    class="bg-gray-100 hover:bg-gray-200 text-gray-700 font-medium py-2 px-3 rounded text-sm transition-colors"
                    @click="loadQuickQuery('update')"
                  >
                    UPDATE
                  </button>
                  <button 
                    class="bg-gray-100 hover:bg-gray-200 text-gray-700 font-medium py-2 px-3 rounded text-sm transition-colors"
                    @click="loadQuickQuery('delete')"
                  >
                    DELETE
                  </button>
                </div>
              </div>

              <!-- Query Results -->
              <div v-if="queryResults.length > 0">
                <label class="block text-sm font-medium text-gray-700 mb-2">نتایج</label>
                <div class="border rounded-lg overflow-hidden">
                  <div class="overflow-x-auto">
                    <table class="min-w-full divide-y divide-gray-200">
                      <thead class="bg-gray-50">
                        <tr>
                          <th v-for="column in queryColumns" :key="column" class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                            {{ column }}
                          </th>
                        </tr>
                      </thead>
                      <tbody class="bg-white divide-y divide-gray-200">
                        <tr v-for="(row, index) in queryResults" :key="index">
                          <td v-for="column in queryColumns" :key="column" class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                            {{ row[column] }}
                          </td>
                        </tr>
                      </tbody>
                    </table>
                  </div>
                </div>
              </div>

              <!-- Query Error -->
              <div v-if="queryError" class="bg-red-50 border border-red-200 rounded-lg p-6">
                <div class="flex items-center">
                  <svg class="w-5 h-5 text-red-500 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
                  </svg>
                  <span class="text-red-800">{{ queryError }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Database Operations -->
      <div class="mt-8 bg-white rounded-lg shadow">
        <div class="p-6 border-b border-gray-200">
          <h2 class="text-xl font-semibold text-gray-900">عملیات دیتابیس</h2>
        </div>
        <div class="p-6">
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
            <!-- Backup -->
            <div class="space-y-4">
              <h3 class="font-medium text-gray-900">پشتیبان‌گیری</h3>
              <div class="space-y-2">
                <button 
                  class="w-full bg-green-600 hover:bg-green-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="createBackup"
                >
                  ایجاد پشتیبان
                </button>
                <button 
                  class="w-full bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="restoreBackup"
                >
                  بازیابی
                </button>
              </div>
            </div>

            <!-- Optimization -->
            <div class="space-y-4">
              <h3 class="font-medium text-gray-900">بهینه‌سازی</h3>
              <div class="space-y-2">
                <button 
                  class="w-full bg-purple-600 hover:bg-purple-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="optimizeTables"
                >
                  بهینه‌سازی جداول
                </button>
                <button 
                  class="w-full bg-yellow-600 hover:bg-yellow-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="analyzeTables"
                >
                  تحلیل جداول
                </button>
              </div>
            </div>

            <!-- Maintenance -->
            <div class="space-y-4">
              <h3 class="font-medium text-gray-900">نگهداری</h3>
              <div class="space-y-2">
                <button 
                  class="w-full bg-indigo-600 hover:bg-indigo-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="checkIntegrity"
                >
                  بررسی یکپارچگی
                </button>
                <button 
                  class="w-full bg-pink-600 hover:bg-pink-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="vacuumDatabase"
                >
                  پاکسازی
                </button>
              </div>
            </div>

            <!-- Monitoring -->
            <div class="space-y-4">
              <h3 class="font-medium text-gray-900">نظارت</h3>
              <div class="space-y-2">
                <button 
                  class="w-full bg-red-600 hover:bg-red-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="showSlowQueries"
                >
                  کوئری‌های کند
                </button>
                <button 
                  class="w-full bg-orange-600 hover:bg-orange-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="showLocks"
                >
                  قفل‌ها
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Query History -->
      <div class="mt-8 bg-white rounded-lg shadow">
        <div class="p-6 border-b border-gray-200">
          <h2 class="text-xl font-semibold text-gray-900">تاریخچه کوئری‌ها</h2>
        </div>
        <div class="p-6">
          <div class="space-y-3">
            <div v-for="query in queryHistory" :key="query.id" class="flex items-center justify-between p-6 border rounded-lg">
              <div class="flex-1">
                <p class="font-mono text-sm text-gray-900 truncate">{{ query.sql }}</p>
                <p class="text-xs text-gray-500 mt-1">{{ query.timestamp }} - {{ query.executionTime }}ms</p>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <span
:class="[
                  'px-2 py-1 rounded text-xs font-medium',
                  query.success ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'
                ]">
                  {{ query.success ? 'موفق' : 'ناموفق' }}
                </span>
                <button 
                  class="text-blue-600 hover:text-blue-800 text-sm font-medium"
                  @click="loadQuery(query)"
                >
                  بارگذاری
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, reactive, ref } from 'vue';

definePageMeta({
  layout: 'admin-main',
  middleware: ['developer-only']
});

// استفاده از useAuth برای چک کردن پرمیژن‌ها
// const { user, hasPermission } = useAuth()

// Reactive data
const selectedTable = ref(null)
const queryText = ref('')
const queryResults = ref([])
const queryError = ref('')
const queryHistory = ref([])

const stats = reactive({
  tables: 12,
  totalRecords: 15420,
  databaseSize: '2.4 GB',
  responseTime: 45
})

const databaseSchema = ref([
  {
    name: 'users',
    records: 1250,
    columns: [
      { name: 'id', type: 'INTEGER PRIMARY KEY' },
      { name: 'username', type: 'VARCHAR(50)' },
      { name: 'email', type: 'VARCHAR(100)' },
      { name: 'created_at', type: 'TIMESTAMP' }
    ]
  },
  {
    name: 'products',
    records: 3420,
    columns: [
      { name: 'id', type: 'INTEGER PRIMARY KEY' },
      { name: 'name', type: 'VARCHAR(200)' },
      { name: 'price', type: 'DECIMAL(10,2)' },
      { name: 'category_id', type: 'INTEGER' }
    ]
  },
  {
    name: 'orders',
    records: 8900,
    columns: [
      { name: 'id', type: 'INTEGER PRIMARY KEY' },
      { name: 'user_id', type: 'INTEGER' },
      { name: 'total_amount', type: 'DECIMAL(10,2)' },
      { name: 'status', type: 'VARCHAR(20)' }
    ]
  },
  {
    name: 'categories',
    records: 150,
    columns: [
      { name: 'id', type: 'INTEGER PRIMARY KEY' },
      { name: 'name', type: 'VARCHAR(100)' },
      { name: 'description', type: 'TEXT' }
    ]
  }
])

// Computed
const queryColumns = computed(() => {
  if (queryResults.value.length > 0) {
    return Object.keys(queryResults.value[0])
  }
  return []
})

// Methods
function selectTable(table) {
  selectedTable.value = table
  queryText.value = `SELECT * FROM ${table.name} LIMIT 10;`
}

function executeQuery() {
  if (!queryText.value.trim()) return

  const startTime = Date.now()
  
  try {
    // Simulate query execution
    const mockResults = generateMockResults(queryText.value)
    queryResults.value = mockResults
    queryError.value = ''
    
    const executionTime = Date.now() - startTime
    
    // Add to history
    queryHistory.value.unshift({
      id: Date.now(),
      sql: queryText.value,
      success: true,
      executionTime,
      timestamp: new Date().toLocaleString('fa-IR')
    })
    
  } catch (error) {
    queryError.value = error.message
    queryResults.value = []
    
    queryHistory.value.unshift({
      id: Date.now(),
      sql: queryText.value,
      success: false,
      executionTime: Date.now() - startTime,
      timestamp: new Date().toLocaleString('fa-IR')
    })
  }
}

function generateMockResults(query) {
  const lowerQuery = query.toLowerCase()
  
  if (lowerQuery.includes('users')) {
    return [
      { id: 1, username: 'admin', email: 'admin@example.com', created_at: '2024-01-01' },
      { id: 2, username: 'user1', email: 'user1@example.com', created_at: '2024-01-02' },
      { id: 3, username: 'user2', email: 'user2@example.com', created_at: '2024-01-03' }
    ]
  } else if (lowerQuery.includes('products')) {
    return [
      { id: 1, name: 'لپ‌تاپ', price: 25000000, category_id: 1 },
      { id: 2, name: 'موبایل', price: 15000000, category_id: 1 },
      { id: 3, name: 'کتاب', price: 500000, category_id: 2 }
    ]
  } else if (lowerQuery.includes('orders')) {
    return [
      { id: 1, user_id: 1, total_amount: 25000000, status: 'completed' },
      { id: 2, user_id: 2, total_amount: 15000000, status: 'pending' },
      { id: 3, user_id: 3, total_amount: 500000, status: 'cancelled' }
    ]
  }
  
  return []
}

function clearQuery() {
  queryText.value = ''
  queryResults.value = []
  queryError.value = ''
}

function loadQuickQuery(type) {
  const queries = {
    select: 'SELECT * FROM users LIMIT 10;',
    insert: 'INSERT INTO users (username, email) VALUES (\'newuser\', \'newuser@example.com\');',
    update: 'UPDATE users SET email = \'updated@example.com\' WHERE id = 1;',
    delete: 'DELETE FROM users WHERE id = 1;'
  }
  queryText.value = queries[type] || ''
}

function loadQuery(query) {
  queryText.value = query.sql
}

function createBackup() {
  alert('پشتیبان دیتابیس ایجاد شد')
}

function restoreBackup() {
  alert('بازیابی دیتابیس شروع شد')
}

function optimizeTables() {
  alert('بهینه‌سازی جداول انجام شد')
}

function analyzeTables() {
  alert('تحلیل جداول انجام شد')
}

function checkIntegrity() {
  alert('بررسی یکپارچگی دیتابیس انجام شد')
}

function vacuumDatabase() {
  alert('پاکسازی دیتابیس انجام شد')
}

function showSlowQueries() {
  alert('کوئری‌های کند نمایش داده می‌شوند')
}

function showLocks() {
  alert('قفل‌های دیتابیس نمایش داده می‌شوند')
}
</script> 
