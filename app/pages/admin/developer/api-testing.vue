<template>
  <div class="min-h-screen bg-gray-50 p-6">
    <div class="max-w-7xl mx-auto">
      <!-- Header -->
      <div class="mb-8">
        <h1 class="text-3xl font-bold text-gray-900 mb-2">تست و مستندسازی API</h1>
        <p class="text-gray-600">ابزارهای تست، مستندسازی و مانیتورینگ API</p>
      </div>

      <!-- Stats Cards -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-blue-100 rounded-lg">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">کل API ها</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.totalApis }}</p>
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
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">موفق</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.successRate }}%</p>
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
              <p class="text-2xl font-semibold text-gray-900">{{ stats.avgResponseTime }}ms</p>
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
              <p class="text-sm font-medium text-gray-600">درخواست/دقیقه</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.requestsPerMinute }}</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Main Content Grid -->
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <!-- API Testing -->
        <div class="lg:col-span-2 bg-white rounded-lg shadow">
          <div class="p-6 border-b border-gray-200">
            <h2 class="text-xl font-semibold text-gray-900">تست API</h2>
          </div>
          <div class="p-6">
            <div class="space-y-6">
              <!-- Request Configuration -->
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">متد HTTP</label>
                  <select 
                    v-model="requestConfig.method"
                    class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
                  >
                    <option value="GET">GET</option>
                    <option value="POST">POST</option>
                    <option value="PUT">PUT</option>
                    <option value="DELETE">DELETE</option>
                    <option value="PATCH">PATCH</option>
                  </select>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">URL</label>
                  <input 
                    v-model="requestConfig.url"
                    type="text"
                    class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
                    placeholder="https://api.example.com/endpoint"
                  >
                </div>
              </div>

              <!-- Headers -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">Headers</label>
                <div class="space-y-2">
                  <div v-for="(header, index) in requestConfig.headers" :key="index" class="flex items-center space-x-2 space-x-reverse">
                    <input 
                      v-model="header.key"
                      type="text"
                      class="flex-1 border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
                      placeholder="Key"
                    >
                    <input 
                      v-model="header.value"
                      type="text"
                      class="flex-1 border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
                      placeholder="Value"
                    >
                    <button 
                      class="text-red-600 hover:text-red-800"
                      @click="removeHeader(index)"
                    >
                      <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                      </svg>
                    </button>
                  </div>
                  <button 
                    class="text-blue-600 hover:text-blue-800 text-sm font-medium"
                    @click="addHeader"
                  >
                    + افزودن Header
                  </button>
                </div>
              </div>

              <!-- Request Body -->
              <div v-if="['POST', 'PUT', 'PATCH'].includes(requestConfig.method)">
                <label class="block text-sm font-medium text-gray-700 mb-2">Request Body</label>
                <textarea
                  v-model="requestConfig.body"
                  rows="6"
                  class="w-full border border-gray-300 rounded-lg px-3 py-2 font-mono text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
                  placeholder='{"key": "value"}'
                ></textarea>
              </div>

              <!-- Send Request -->
              <div class="flex items-center space-x-2 space-x-reverse">
                <button 
                  class="bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 px-6 rounded-lg transition-colors"
                  @click="sendRequest"
                >
                  ارسال درخواست
                </button>
                <button 
                  class="bg-green-600 hover:bg-green-700 text-white font-medium py-2 px-6 rounded-lg transition-colors"
                  @click="saveRequest"
                >
                  ذخیره درخواست
                </button>
                <button 
                  class="bg-gray-600 hover:bg-gray-700 text-white font-medium py-2 px-6 rounded-lg transition-colors"
                  @click="clearRequest"
                >
                  پاک کردن
                </button>
              </div>

              <!-- Response -->
              <div v-if="response">
                <div class="flex items-center justify-between mb-2">
                  <h3 class="font-medium text-gray-900">پاسخ</h3>
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <span
:class="[
                      'px-2 py-1 rounded text-xs font-medium',
                      response.status >= 200 && response.status < 300 ? 'bg-green-100 text-green-800' :
                      response.status >= 400 && response.status < 500 ? 'bg-yellow-100 text-yellow-800' :
                      'bg-red-100 text-red-800'
                    ]">
                      {{ response.status }}
                    </span>
                    <span class="text-sm text-gray-500">{{ response.time }}ms</span>
                  </div>
                </div>
                <div class="bg-gray-900 rounded-lg p-6 overflow-x-auto">
                  <pre class="text-gray-100 font-mono text-sm">{{ response.data }}</pre>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- API Collections -->
        <div class="bg-white rounded-lg shadow">
          <div class="p-6 border-b border-gray-200">
            <h2 class="text-xl font-semibold text-gray-900">مجموعه‌های API</h2>
          </div>
          <div class="p-6">
            <div class="space-y-4">
              <div v-for="collection in apiCollections" :key="collection.id" class="border rounded-lg p-6 hover:bg-gray-50">
                <div class="flex items-center justify-between mb-2">
                  <h3 class="font-medium text-gray-900">{{ collection.name }}</h3>
                  <span class="text-xs bg-gray-100 text-gray-600 px-2 py-1 rounded">{{ collection.endpoints }} endpoints</span>
                </div>
                <p class="text-sm text-gray-500 mb-3">{{ collection.description }}</p>
                <div class="flex items-center space-x-2 space-x-reverse">
                  <button 
                    class="text-blue-600 hover:text-blue-800 text-sm font-medium"
                    @click="loadCollection(collection)"
                  >
                    بارگذاری
                  </button>
                  <button 
                    class="text-green-600 hover:text-green-800 text-sm font-medium"
                    @click="exportCollection(collection)"
                  >
                    صادر کردن
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- API Documentation -->
      <div class="mt-8 bg-white rounded-lg shadow">
        <div class="p-6 border-b border-gray-200">
          <div class="flex items-center justify-between">
            <h2 class="text-xl font-semibold text-gray-900">مستندات API</h2>
            <div class="flex items-center space-x-2 space-x-reverse">
              <button 
                class="bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                @click="generateDocs"
              >
                تولید مستندات
              </button>
              <button 
                class="bg-green-600 hover:bg-green-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                @click="exportDocs"
              >
                صادر کردن
              </button>
            </div>
          </div>
        </div>
        <div class="p-6">
          <div class="space-y-6">
            <div v-for="endpoint in apiEndpoints" :key="endpoint.id" class="border rounded-lg p-6">
              <div class="flex items-center justify-between mb-4">
                <div class="flex items-center space-x-3 space-x-reverse">
                  <span
:class="[
                    'px-3 py-1 rounded text-sm font-medium',
                    endpoint.method === 'GET' ? 'bg-green-100 text-green-800' :
                    endpoint.method === 'POST' ? 'bg-blue-100 text-blue-800' :
                    endpoint.method === 'PUT' ? 'bg-yellow-100 text-yellow-800' :
                    'bg-red-100 text-red-800'
                  ]">
                    {{ endpoint.method }}
                  </span>
                  <span class="font-mono text-gray-900">{{ endpoint.path }}</span>
                </div>
                <div class="flex items-center space-x-2 space-x-reverse">
                  <button 
                    class="text-blue-600 hover:text-blue-800 text-sm font-medium"
                    @click="testEndpoint(endpoint)"
                  >
                    تست
                  </button>
                  <button 
                    class="text-gray-600 hover:text-gray-800 text-sm font-medium"
                    @click="editEndpoint(endpoint)"
                  >
                    ویرایش
                  </button>
                </div>
              </div>
              <p class="text-gray-600 mb-4">{{ endpoint.description }}</p>
              
              <!-- Parameters -->
              <div v-if="endpoint.parameters.length > 0" class="mb-4">
                <h4 class="font-medium text-gray-900 mb-2">پارامترها</h4>
                <div class="space-y-2">
                  <div v-for="param in endpoint.parameters" :key="param.name" class="flex items-center space-x-4 space-x-reverse text-sm">
                    <span class="font-mono text-gray-900">{{ param.name }}</span>
                    <span class="text-gray-500">{{ param.type }}</span>
                    <span
:class="[
                      'px-2 py-1 rounded text-xs',
                      param.required ? 'bg-red-100 text-red-800' : 'bg-gray-100 text-gray-600'
                    ]">
                      {{ param.required ? 'ضروری' : 'اختیاری' }}
                    </span>
                    <span class="text-gray-500">{{ param.description }}</span>
                  </div>
                </div>
              </div>

              <!-- Response Examples -->
              <div>
                <h4 class="font-medium text-gray-900 mb-2">نمونه پاسخ</h4>
                <div class="bg-gray-900 rounded-lg p-6 overflow-x-auto">
                  <pre class="text-gray-100 font-mono text-sm">{{ endpoint.responseExample }}</pre>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- API Monitoring -->
      <div class="mt-8 bg-white rounded-lg shadow">
        <div class="p-6 border-b border-gray-200">
          <h2 class="text-xl font-semibold text-gray-900">مانیتورینگ API</h2>
        </div>
        <div class="p-6">
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
            <div class="space-y-4">
              <h3 class="font-medium text-gray-900">آمار درخواست‌ها</h3>
              <div class="space-y-2">
                <div class="flex justify-between items-center">
                  <span class="text-sm text-gray-600">موفق:</span>
                  <span class="text-sm font-medium text-green-600">{{ monitoring.successCount }}</span>
                </div>
                <div class="flex justify-between items-center">
                  <span class="text-sm text-gray-600">خطا:</span>
                  <span class="text-sm font-medium text-red-600">{{ monitoring.errorCount }}</span>
                </div>
                <div class="flex justify-between items-center">
                  <span class="text-sm text-gray-600">کل:</span>
                  <span class="text-sm font-medium">{{ monitoring.totalCount }}</span>
                </div>
              </div>
            </div>

            <div class="space-y-4">
              <h3 class="font-medium text-gray-900">زمان پاسخ</h3>
              <div class="space-y-2">
                <div class="flex justify-between items-center">
                  <span class="text-sm text-gray-600">میانگین:</span>
                  <span class="text-sm font-medium">{{ monitoring.avgResponseTime }}ms</span>
                </div>
                <div class="flex justify-between items-center">
                  <span class="text-sm text-gray-600">حداقل:</span>
                  <span class="text-sm font-medium">{{ monitoring.minResponseTime }}ms</span>
                </div>
                <div class="flex justify-between items-center">
                  <span class="text-sm text-gray-600">حداکثر:</span>
                  <span class="text-sm font-medium">{{ monitoring.maxResponseTime }}ms</span>
                </div>
              </div>
            </div>

            <div class="space-y-4">
              <h3 class="font-medium text-gray-900">خطاها</h3>
              <div class="space-y-2">
                <div v-for="error in monitoring.errors" :key="error.code" class="flex justify-between items-center">
                  <span class="text-sm text-gray-600">{{ error.code }}:</span>
                  <span class="text-sm font-medium text-red-600">{{ error.count }}</span>
                </div>
              </div>
            </div>

            <div class="space-y-4">
              <h3 class="font-medium text-gray-900">عملکرد</h3>
              <div class="space-y-2">
                <button 
                  class="w-full bg-green-600 hover:bg-green-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="startMonitoring"
                >
                  شروع مانیتورینگ
                </button>
                <button 
                  class="w-full bg-red-600 hover:bg-red-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="stopMonitoring"
                >
                  توقف مانیتورینگ
                </button>
                <button 
                  class="w-full bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="exportMetrics"
                >
                  صادر کردن آمار
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
import { reactive, ref } from 'vue';

definePageMeta({
  layout: 'admin-main',
  middleware: ['developer-only']
});

// Authentication setup
const { checkAuth } = useAuth()

// Check authentication on mount
if (import.meta.client) {
  checkAuth()
}

// Reactive data
const requestConfig = reactive({
  method: 'GET',
  url: 'https://api.example.com/users',
  headers: [
    { key: 'Content-Type', value: 'application/json' },
    { key: 'Authorization', value: 'Bearer token' }
  ],
  body: ''
})

const response = ref(null)

const stats = reactive({
  totalApis: 45,
  successRate: 98.5,
  avgResponseTime: 245,
  requestsPerMinute: 1250
})

const apiCollections = ref([
  {
    id: 1,
    name: 'User Management',
    description: 'API های مدیریت کاربران',
    endpoints: 8
  },
  {
    id: 2,
    name: 'Product Catalog',
    description: 'API های کاتالوگ محصولات',
    endpoints: 12
  },
  {
    id: 3,
    name: 'Order Processing',
    description: 'API های پردازش سفارشات',
    endpoints: 6
  },
  {
    id: 4,
    name: 'Authentication',
    description: 'API های احراز هویت',
    endpoints: 4
  }
])

const apiEndpoints = ref([
  {
    id: 1,
    method: 'GET',
    path: '/api/users',
    description: 'دریافت لیست کاربران',
    parameters: [
      { name: 'page', type: 'integer', required: false, description: 'شماره صفحه' },
      { name: 'limit', type: 'integer', required: false, description: 'تعداد در هر صفحه' }
    ],
    responseExample: `{
  "users": [
    {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com"
    }
  ],
  "total": 100,
  "page": 1
}`
  },
  {
    id: 2,
    method: 'POST',
    path: '/api/users',
    description: 'ایجاد کاربر جدید',
    parameters: [
      { name: 'name', type: 'string', required: true, description: 'نام کاربر' },
      { name: 'email', type: 'string', required: true, description: 'ایمیل کاربر' }
    ],
    responseExample: `{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "created_at": "2024-01-15T10:30:00Z"
}`
  },
  {
    id: 3,
    method: 'PUT',
    path: '/api/users/{id}',
    description: 'به‌روزرسانی کاربر',
    parameters: [
      { name: 'id', type: 'integer', required: true, description: 'شناسه کاربر' },
      { name: 'name', type: 'string', required: false, description: 'نام کاربر' },
      { name: 'email', type: 'string', required: false, description: 'ایمیل کاربر' }
    ],
    responseExample: `{
  "id": 1,
  "name": "John Doe Updated",
  "email": "john.updated@example.com",
  "updated_at": "2024-01-15T10:30:00Z"
}`
  }
])

const monitoring = reactive({
  successCount: 12450,
  errorCount: 125,
  totalCount: 12575,
  avgResponseTime: 245,
  minResponseTime: 45,
  maxResponseTime: 1200,
  errors: [
    { code: '400', count: 45 },
    { code: '401', count: 30 },
    { code: '404', count: 25 },
    { code: '500', count: 25 }
  ]
})

// Methods
function addHeader() {
  requestConfig.headers.push({ key: '', value: '' })
}

function removeHeader(index) {
  requestConfig.headers.splice(index, 1)
}

function sendRequest() {
  if (!requestConfig.url.trim()) {
    alert('لطفا URL را وارد کنید')
    return
  }

  // Simulate API request
  const startTime = Date.now()
  
  setTimeout(() => {
    const endTime = Date.now()
    const responseTime = endTime - startTime
    
    response.value = {
      status: 200,
      time: responseTime,
      data: `{
  "success": true,
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com"
  },
  "message": "Request successful"
}`
    }
  }, 500)
}

function saveRequest() {
  alert('درخواست ذخیره شد')
}

function clearRequest() {
  requestConfig.method = 'GET'
  requestConfig.url = ''
  requestConfig.headers = [
    { key: 'Content-Type', value: 'application/json' }
  ]
  requestConfig.body = ''
  response.value = null
}

function loadCollection(collection) {
  alert(`مجموعه ${collection.name} بارگذاری شد`)
}

function exportCollection(collection) {
  alert(`مجموعه ${collection.name} صادر شد`)
}

function generateDocs() {
  alert('مستندات API تولید شد')
}

function exportDocs() {
  alert('مستندات API صادر شد')
}

function testEndpoint(endpoint) {
  requestConfig.method = endpoint.method
  requestConfig.url = `https://api.example.com${endpoint.path}`
  alert(`Endpoint ${endpoint.method} ${endpoint.path} برای تست آماده شد`)
}

function editEndpoint(endpoint) {
  alert(`ویرایش endpoint ${endpoint.method} ${endpoint.path}`)
}

function startMonitoring() {
  alert('مانیتورینگ API شروع شد')
}

function stopMonitoring() {
  alert('مانیتورینگ API متوقف شد')
}

function exportMetrics() {
  alert('آمار API صادر شد')
}
</script> 
