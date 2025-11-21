<template>
  <div class="min-h-screen bg-gray-50 p-6">
    <div class="max-w-7xl mx-auto">
      <!-- Header -->
      <div class="mb-8">
        <h1 class="text-3xl font-bold text-gray-900 mb-2">تست پیشرفته API</h1>
        <p class="text-gray-600">ابزارهای تست و دیباگ API endpoints</p>
      </div>

      <!-- Stats Cards -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-blue-100 rounded-lg">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">تست‌های موفق</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.success }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-red-100 rounded-lg">
              <svg class="w-6 h-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">تست‌های ناموفق</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.failed }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-green-100 rounded-lg">
              <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">میانگین زمان پاسخ</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.avgResponseTime }}ms</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-purple-100 rounded-lg">
              <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">کل تست‌ها</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.total }}</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Main Content Grid -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
        <!-- API Tester -->
        <div class="bg-white rounded-lg shadow">
          <div class="p-6 border-b border-gray-200">
            <h2 class="text-xl font-semibold text-gray-900">تست API</h2>
          </div>
          <div class="p-6">
            <div class="space-y-4">
              <!-- Method Selection -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">روش درخواست</label>
                <select v-model="requestConfig.method" class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500">
                  <option value="GET">GET</option>
                  <option value="POST">POST</option>
                  <option value="PUT">PUT</option>
                  <option value="DELETE">DELETE</option>
                  <option value="PATCH">PATCH</option>
                </select>
              </div>

              <!-- URL Input -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">آدرس API</label>
                <input 
                  v-model="requestConfig.url" 
                  type="text" 
                  placeholder="https://api.example.com/endpoint"
                  class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
              </div>

              <!-- Headers -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">Headers</label>
                <textarea 
                  v-model="requestConfig.headers" 
                  placeholder='{"Content-Type": "application/json", "Authorization": "Bearer token"}'
                  rows="3"
                  class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
                ></textarea>
              </div>

              <!-- Body -->
              <div v-if="['POST', 'PUT', 'PATCH'].includes(requestConfig.method)">
                <label class="block text-sm font-medium text-gray-700 mb-2">Body</label>
                <textarea 
                  v-model="requestConfig.body" 
                  placeholder='{"key": "value"}'
                  rows="5"
                  class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
                ></textarea>
              </div>

              <!-- Send Button -->
              <button 
                :disabled="sending" 
                class="w-full bg-blue-600 hover:bg-blue-700 disabled:bg-gray-400 text-white font-medium py-3 px-4 rounded-lg transition-colors"
                @click="sendRequest"
              >
                <span v-if="sending">در حال ارسال...</span>
                <span v-else>ارسال درخواست</span>
              </button>
            </div>
          </div>
        </div>

        <!-- Response Viewer -->
        <div class="bg-white rounded-lg shadow">
          <div class="p-6 border-b border-gray-200">
            <h2 class="text-xl font-semibold text-gray-900">پاسخ</h2>
          </div>
          <div class="p-6">
            <div v-if="response" class="space-y-4">
              <!-- Status -->
              <div class="flex items-center justify-between">
                <span class="text-sm font-medium text-gray-700">وضعیت:</span>
                <span
:class="[
                  'px-2 py-1 rounded text-sm font-medium',
                  response.status >= 200 && response.status < 300 ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'
                ]">
                  {{ response.status }}
                </span>
              </div>

              <!-- Response Time -->
              <div class="flex items-center justify-between">
                <span class="text-sm font-medium text-gray-700">زمان پاسخ:</span>
                <span class="text-sm text-gray-600">{{ response.responseTime }}ms</span>
              </div>

              <!-- Response Headers -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">Headers</label>
                <pre class="bg-gray-100 p-3 rounded text-sm overflow-x-auto">{{ JSON.stringify(response.headers, null, 2) }}</pre>
              </div>

              <!-- Response Body -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">Body</label>
                <pre class="bg-gray-100 p-3 rounded text-sm overflow-x-auto max-h-64 overflow-y-auto">{{ JSON.stringify(response.data, null, 2) }}</pre>
              </div>
            </div>
            <div v-else class="text-center text-gray-500 py-8">
              <svg class="w-12 h-12 mx-auto mb-4 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
              </svg>
              <p>پاسخی برای نمایش وجود ندارد</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Quick Tests -->
      <div class="mt-8 bg-white rounded-lg shadow">
        <div class="p-6 border-b border-gray-200">
          <h2 class="text-xl font-semibold text-gray-900">تست‌های سریع</h2>
        </div>
        <div class="p-6">
          <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
            <button 
              class="bg-green-500 hover:bg-green-600 text-white font-medium py-3 px-4 rounded-lg transition-colors"
              @click="quickTest('health')"
            >
              تست سلامت سیستم
            </button>
            <button 
              class="bg-blue-500 hover:bg-blue-600 text-white font-medium py-3 px-4 rounded-lg transition-colors"
              @click="quickTest('database')"
            >
              تست دیتابیس
            </button>
            <button 
              class="bg-purple-500 hover:bg-purple-600 text-white font-medium py-3 px-4 rounded-lg transition-colors"
              @click="quickTest('auth')"
            >
              تست احراز هویت
            </button>
          </div>
        </div>
      </div>

      <!-- Test History -->
      <div class="mt-8 bg-white rounded-lg shadow">
        <div class="p-6 border-b border-gray-200">
          <h2 class="text-xl font-semibold text-gray-900">تاریخچه تست‌ها</h2>
        </div>
        <div class="p-6">
          <div class="space-y-3">
            <div v-for="test in testHistory" :key="test.id" class="flex items-center justify-between p-6 border rounded-lg">
              <div class="flex items-center space-x-4 space-x-reverse">
                <div
:class="[
                  'w-3 h-3 rounded-full',
                  test.success ? 'bg-green-500' : 'bg-red-500'
                ]"></div>
                <div>
                  <p class="font-medium">{{ test.method }} {{ test.url }}</p>
                  <p class="text-sm text-gray-600">{{ test.responseTime }}ms - {{ test.timestamp }}</p>
                </div>
              </div>
              <button 
                class="text-blue-600 hover:text-blue-800 text-sm font-medium"
                @click="loadTest(test)"
              >
                بارگذاری مجدد
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { onMounted, reactive, ref } from 'vue';

definePageMeta({
  layout: 'admin-main',
  middleware: ['developer-only']
});

// استفاده از useAuth برای چک کردن پرمیژن‌ها
// const { user, hasPermission } = useAuth()

// Reactive data
const sending = ref(false)
const response = ref(null)
const testHistory = ref([])

const stats = reactive({
  success: 0,
  failed: 0,
  total: 0,
  avgResponseTime: 0
})

const requestConfig = reactive({
  method: 'GET',
  url: '',
  headers: '{\n  "Content-Type": "application/json"\n}',
  body: ''
})

// Methods
async function sendRequest() {
  if (!requestConfig.url) {
    alert('لطفاً آدرس API را وارد کنید')
    return
  }

  sending.value = true
  const startTime = Date.now()

  try {
    const headers = JSON.parse(requestConfig.headers || '{}')
    const body = requestConfig.body ? JSON.parse(requestConfig.body) : undefined

    const responseData = await $fetch(requestConfig.url, {
      method: requestConfig.method,
      headers,
      body
    })

    const responseTime = Date.now() - startTime

    response.value = {
      status: 200,
      responseTime,
      headers: {},
      data: responseData
    }

    // Update stats
    stats.success++
    stats.total++
    stats.avgResponseTime = Math.round((stats.avgResponseTime * (stats.total - 1) + responseTime) / stats.total)

    // Add to history
    addToHistory({
      method: requestConfig.method,
      url: requestConfig.url,
      responseTime,
      success: true
    })

  } catch (error) {
    const responseTime = Date.now() - startTime
    
    response.value = {
      status: error.status || 500,
      responseTime,
      headers: {},
      data: { error: error.message }
    }

    // Update stats
    stats.failed++
    stats.total++

    // Add to history
    addToHistory({
      method: requestConfig.method,
      url: requestConfig.url,
      responseTime,
      success: false
    })
  } finally {
    sending.value = false
  }
}

function addToHistory(test) {
  test.id = Date.now()
  test.timestamp = new Date().toLocaleString('fa-IR')
  testHistory.value.unshift(test)
  
  // Keep only last 10 tests
  if (testHistory.value.length > 10) {
    testHistory.value = testHistory.value.slice(0, 10)
  }
}

function loadTest(test) {
  requestConfig.method = test.method
  requestConfig.url = test.url
}

async function quickTest(type) {
  const tests = {
    health: {
      method: 'GET',
      url: '/api/health',
      headers: '{\n  "Content-Type": "application/json"\n}',
      body: ''
    },
    database: {
      method: 'GET',
      url: '/api/db-health',
      headers: '{\n  "Content-Type": "application/json"\n}',
      body: ''
    },
    auth: {
      method: 'POST',
      url: '/api/auth/login',
      headers: '{\n  "Content-Type": "application/json"\n}',
      body: '{\n  "email": "test@example.com",\n  "password": "test123"\n}'
    }
  }

  const test = tests[type]
  if (test) {
    Object.assign(requestConfig, test)
    await sendRequest()
  }
}

// Load initial stats
onMounted(() => {
  // Load stats from localStorage
  const savedStats = localStorage.getItem('api-test-stats')
  if (savedStats) {
    Object.assign(stats, JSON.parse(savedStats))
  }

  // Load test history
  const savedHistory = localStorage.getItem('api-test-history')
  if (savedHistory) {
    testHistory.value = JSON.parse(savedHistory)
  }
})

// Save stats and history to localStorage
watch([stats, testHistory], () => {
  localStorage.setItem('api-test-stats', JSON.stringify(stats))
  localStorage.setItem('api-test-history', JSON.stringify(testHistory.value))
}, { deep: true })
</script> 
