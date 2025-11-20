<template>
  <div class="p-6" dir="rtl">
    <!-- Header -->
    <div class="mb-6 bg-white p-6 rounded-lg shadow-sm border border-gray-200">
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-2xl font-bold text-gray-900 mb-2">مدیریت API</h1>
          <p class="text-gray-600">مستندات، تست و مدیریت API های سیستم</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button 
            class="px-4 py-2 text-green-600 bg-green-50 border-2 border-green-200 rounded-lg hover:bg-green-100 hover:border-green-300 transition-all duration-200 flex items-center shadow-sm hover:shadow-md" 
            @click="exportApiDocs"
          >
            <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
            </svg>
            خروجی PDF
          </button>
          <button 
            class="px-4 py-2 text-blue-600 bg-blue-50 border-2 border-blue-200 rounded-lg hover:bg-blue-100 hover:border-blue-300 transition-all duration-200 flex items-center shadow-sm hover:shadow-md" 
            @click="refreshApiStatus"
          >
            <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
            </svg>
            به‌روزرسانی
          </button>
        </div>
      </div>
    </div>

    <!-- API Status Cards -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
      <!-- Total Endpoints -->
      <div class="bg-gradient-to-br from-blue-50 to-blue-100 p-6 rounded-xl shadow-md border border-blue-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-blue-600">کل Endpoint ها</p>
            <p class="text-2xl font-bold text-blue-800">{{ apiStats.totalEndpoints }}</p>
          </div>
          <div class="w-12 h-12 bg-blue-500 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 9l3 3-3 3m5 0h3M5 20h14a2 2 0 002-2V6a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
            </svg>
          </div>
        </div>
      </div>

      <!-- Active Endpoints -->
      <div class="bg-gradient-to-br from-green-50 to-green-100 p-6 rounded-xl shadow-md border border-green-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-green-600">Endpoint های فعال</p>
            <p class="text-2xl font-bold text-green-800">{{ apiStats.activeEndpoints }}</p>
          </div>
          <div class="w-12 h-12 bg-green-500 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
      </div>

      <!-- API Calls Today -->
      <div class="bg-gradient-to-br from-purple-50 to-purple-100 p-6 rounded-xl shadow-md border border-purple-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-purple-600">درخواست‌های امروز</p>
            <p class="text-2xl font-bold text-purple-800">{{ apiStats.todayCalls }}</p>
          </div>
          <div class="w-12 h-12 bg-purple-500 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
            </svg>
          </div>
        </div>
      </div>

      <!-- Error Rate -->
      <div class="bg-gradient-to-br from-red-50 to-red-100 p-6 rounded-xl shadow-md border border-red-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-red-600">نرخ خطا</p>
            <p class="text-2xl font-bold text-red-800">{{ apiStats.errorRate }}%</p>
          </div>
          <div class="w-12 h-12 bg-red-500 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- Tabs -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 mb-8">
      <div class="border-b border-gray-200">
        <nav class="flex space-x-8 space-x-reverse px-6">
          <button 
            v-for="tab in tabs" 
            :key="tab.id"
            :class="[
              'py-4 px-1 border-b-2 font-medium text-sm transition-colors duration-200',
              activeTab === tab.id 
                ? 'border-blue-500 text-blue-600' 
                : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
            ]"
            @click="activeTab = tab.id"
          >
            {{ tab.name }}
          </button>
        </nav>
      </div>

      <!-- Tab Content -->
      <div class="p-6">
        <!-- API Documentation Tab -->
        <div v-if="activeTab === 'documentation'" class="space-y-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">مستندات API</h3>
            <div class="flex items-center space-x-3 space-x-reverse">
              <select v-model="selectedVersion" class="border border-gray-300 rounded-md px-3 py-1 text-sm">
                <option value="v1">نسخه 1.0</option>
                <option value="v2">نسخه 2.0</option>
                <option value="beta">نسخه Beta</option>
              </select>
              <button class="text-blue-600 hover:text-blue-800 text-sm" @click="toggleAllEndpoints">
                {{ allExpanded ? 'بستن همه' : 'باز کردن همه' }}
              </button>
            </div>
          </div>

          <!-- API Categories -->
          <div class="space-y-4">
            <div v-for="category in apiCategories" :key="category.name" class="border border-gray-200 rounded-lg">
              <div class="bg-gray-50 px-4 py-3 border-b border-gray-200">
                <h4 class="font-semibold text-gray-900">{{ category.name }}</h4>
                <p class="text-sm text-gray-600">{{ category.description }}</p>
              </div>
              <div class="p-6 space-y-3">
                <div v-for="endpoint in category.endpoints" :key="endpoint.path" class="border border-gray-200 rounded-lg p-6">
                  <div class="flex items-center justify-between mb-3">
                    <div class="flex items-center space-x-3 space-x-reverse">
                      <span 
                        :class="[
                          'px-2 py-1 text-xs font-semibold rounded',
                          getMethodColor(endpoint.method)
                        ]"
                      >
                        {{ endpoint.method }}
                      </span>
                      <code class="bg-gray-100 px-2 py-1 rounded text-sm font-mono">{{ endpoint.path }}</code>
                    </div>
                    <button 
                      class="text-gray-500 hover:text-gray-700"
                      @click="toggleEndpoint(endpoint.path)"
                    >
                      <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
                      </svg>
                    </button>
                  </div>
                  
                  <p class="text-sm text-gray-600 mb-3">{{ endpoint.description }}</p>
                  
                  <div v-if="expandedEndpoints.includes(endpoint.path)" class="space-y-4">
                    <!-- Parameters -->
                    <div v-if="endpoint.parameters && endpoint.parameters.length > 0">
                      <h5 class="font-medium text-gray-900 mb-2">پارامترها:</h5>
                      <div class="bg-gray-50 rounded p-3">
                        <div v-for="param in endpoint.parameters" :key="param.name" class="flex items-center justify-between py-1">
                          <div>
                            <span class="font-mono text-sm">{{ param.name }}</span>
                            <span class="text-xs text-gray-500 mr-2">({{ param.type }})</span>
                          </div>
                          <span class="text-xs" :class="param.required ? 'text-red-600' : 'text-gray-500'">
                            {{ param.required ? 'ضروری' : 'اختیاری' }}
                          </span>
                        </div>
                      </div>
                    </div>

                    <!-- Response Example -->
                    <div>
                      <h5 class="font-medium text-gray-900 mb-2">نمونه پاسخ:</h5>
                      <pre class="bg-gray-900 text-green-400 p-3 rounded text-sm overflow-x-auto">{{ endpoint.responseExample }}</pre>
                    </div>

                    <!-- Test Button -->
                    <button 
                      class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors duration-200"
                      @click="testEndpoint(endpoint)"
                    >
                      تست Endpoint
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- API Testing Tab -->
        <div v-if="activeTab === 'testing'" class="space-y-6">
          <h3 class="text-lg font-semibold text-gray-900">تست API</h3>
          
          <!-- Test Form -->
          <div class="bg-gray-50 rounded-lg p-6">
            <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-4">
              <select v-model="testRequest.method" class="border border-gray-300 rounded-md px-3 py-2">
                <option value="GET">GET</option>
                <option value="POST">POST</option>
                <option value="PUT">PUT</option>
                <option value="DELETE">DELETE</option>
                <option value="PATCH">PATCH</option>
              </select>
              <input 
                v-model="testRequest.url" 
                type="text" 
                placeholder="URL endpoint"
                class="border border-gray-300 rounded-md px-3 py-2 md:col-span-2"
              >
              <button 
                :disabled="isLoading"
                class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 disabled:opacity-50"
                @click="sendTestRequest"
              >
                {{ isLoading ? 'در حال ارسال...' : 'ارسال' }}
              </button>
            </div>

            <!-- Headers -->
            <div class="mb-4">
              <h4 class="font-medium text-gray-900 mb-2">Headers:</h4>
              <div class="space-y-2">
                <div v-for="(header, index) in testRequest.headers" :key="index" class="flex items-center space-x-2 space-x-reverse">
                  <input 
                    v-model="header.key" 
                    type="text" 
                    placeholder="Key"
                    class="border border-gray-300 rounded-md px-3 py-1 flex-1"
                  >
                  <input 
                    v-model="header.value" 
                    type="text" 
                    placeholder="Value"
                    class="border border-gray-300 rounded-md px-3 py-1 flex-1"
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
                  class="text-blue-600 hover:text-blue-800 text-sm"
                  @click="addHeader"
                >
                  + افزودن Header
                </button>
              </div>
            </div>

            <!-- Body -->
            <div v-if="testRequest.method !== 'GET'">
              <h4 class="font-medium text-gray-900 mb-2">Body:</h4>
              <textarea 
                v-model="testRequest.body" 
                rows="6"
                placeholder="JSON body"
                class="w-full border border-gray-300 rounded-md px-3 py-2 font-mono text-sm"
              ></textarea>
            </div>
          </div>

          <!-- Response -->
          <div v-if="testResponse" class="bg-gray-50 rounded-lg p-6">
            <h4 class="font-medium text-gray-900 mb-2">پاسخ:</h4>
            <div class="mb-2">
              <span class="text-sm text-gray-600">Status: </span>
              <span 
                :class="[
                  'px-2 py-1 text-xs font-semibold rounded',
                  testResponse.status >= 200 && testResponse.status < 300 ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'
                ]"
              >
                {{ testResponse.status }}
              </span>
              <span class="text-sm text-gray-600 mr-4">زمان: {{ testResponse.duration }}ms</span>
            </div>
            <pre class="bg-gray-900 text-green-400 p-3 rounded text-sm overflow-x-auto">{{ testResponse.data }}</pre>
          </div>
        </div>

        <!-- API Monitoring Tab -->
        <div v-if="activeTab === 'monitoring'" class="space-y-6">
          <h3 class="text-lg font-semibold text-gray-900">نظارت بر API</h3>
          
          <!-- Real-time Stats -->
          <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
            <div class="bg-blue-50 p-6 rounded-lg border border-blue-200">
              <h4 class="font-medium text-blue-900 mb-2">درخواست‌های فعال</h4>
              <p class="text-2xl font-bold text-blue-800">{{ monitoringStats.activeRequests }}</p>
            </div>
            <div class="bg-green-50 p-6 rounded-lg border border-green-200">
              <h4 class="font-medium text-green-900 mb-2">میانگین زمان پاسخ</h4>
              <p class="text-2xl font-bold text-green-800">{{ monitoringStats.avgResponseTime }}ms</p>
            </div>
            <div class="bg-purple-50 p-6 rounded-lg border border-purple-200">
              <h4 class="font-medium text-purple-900 mb-2">درخواست‌های دقیقه</h4>
              <p class="text-2xl font-bold text-purple-800">{{ monitoringStats.requestsPerMinute }}</p>
            </div>
          </div>

          <!-- Recent API Calls -->
          <div class="bg-gray-50 rounded-lg p-6">
            <h4 class="font-medium text-gray-900 mb-4">آخرین درخواست‌های API</h4>
            <div class="space-y-3">
              <div v-for="call in recentApiCalls" :key="call.id" class="bg-white p-6 rounded-lg border border-gray-200">
                <div class="flex items-center justify-between">
                  <div class="flex items-center space-x-3 space-x-reverse">
                    <span 
                      :class="[
                        'px-2 py-1 text-xs font-semibold rounded',
                        getMethodColor(call.method)
                      ]"
                    >
                      {{ call.method }}
                    </span>
                    <span class="font-mono text-sm">{{ call.path }}</span>
                  </div>
                  <div class="flex items-center space-x-4 space-x-reverse text-sm text-gray-600">
                    <span>{{ call.duration }}ms</span>
                    <span 
                      :class="call.status >= 200 && call.status < 300 ? 'text-green-600' : 'text-red-600'"
                    >
                      {{ call.status }}
                    </span>
                    <span>{{ call.timestamp }}</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- API Keys Tab -->
        <div v-if="activeTab === 'keys'" class="space-y-6">
          <div class="flex items-center justify-between">
            <h3 class="text-lg font-semibold text-gray-900">کلیدهای API</h3>
            <button 
              class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors duration-200"
              @click="generateNewKey"
            >
              تولید کلید جدید
            </button>
          </div>

          <div class="space-y-4">
            <div v-for="key in apiKeys" :key="key.id" class="bg-gray-50 rounded-lg p-6 border border-gray-200">
              <div class="flex items-center justify-between">
                <div>
                  <h4 class="font-medium text-gray-900">{{ key.name }}</h4>
                  <p class="text-sm text-gray-600">{{ key.description }}</p>
                  <p class="text-xs text-gray-500 mt-1">تاریخ ایجاد: {{ key.createdAt }}</p>
                </div>
                <div class="flex items-center space-x-2 space-x-reverse">
                  <span 
                    :class="[
                      'px-2 py-1 text-xs font-semibold rounded',
                      key.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'
                    ]"
                  >
                    {{ key.status === 'active' ? 'فعال' : 'غیرفعال' }}
                  </span>
                  <button 
                    :class="[
                      'px-3 py-1 text-xs rounded transition-colors duration-200',
                      key.status === 'active' 
                        ? 'bg-red-100 text-red-600 hover:bg-red-200' 
                        : 'bg-green-100 text-green-600 hover:bg-green-200'
                    ]"
                    @click="toggleKeyStatus(key.id)"
                  >
                    {{ key.status === 'active' ? 'غیرفعال' : 'فعال' }}
                  </button>
                  <button 
                    class="text-red-600 hover:text-red-800"
                    @click="deleteKey(key.id)"
                  >
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                    </svg>
                  </button>
                </div>
              </div>
              <div class="mt-3">
                <code class="bg-gray-900 text-green-400 px-3 py-2 rounded text-sm font-mono block">{{ key.key }}</code>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
definePageMeta({ layout: 'admin-main' })

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user, hasPermission } = useAuth()

// Reactive data
const activeTab = ref('documentation')
const selectedVersion = ref('v1')
const allExpanded = ref(false)
const expandedEndpoints = ref([])
const isLoading = ref(false)

// API Stats
const apiStats = ref({
  totalEndpoints: 24,
  activeEndpoints: 22,
  todayCalls: 15420,
  errorRate: 2.3
})

// Monitoring Stats
const monitoringStats = ref({
  activeRequests: 8,
  avgResponseTime: 245,
  requestsPerMinute: 156
})

// Test Request
const testRequest = ref({
  method: 'GET',
  url: '',
  headers: [
    { key: 'Content-Type', value: 'application/json' },
    { key: 'Authorization', value: 'Bearer ' }
  ],
  body: ''
})

// Test Response
const testResponse = ref(null)

// API Keys
const apiKeys = ref([
  {
    id: 1,
    name: 'کلید اصلی',
    description: 'کلید API برای دسترسی کامل به سیستم',
    key: 'sk_live_1234567890abcdef',
    status: 'active',
    createdAt: '1402/10/15'
  },
  {
    id: 2,
    name: 'کلید تست',
    description: 'کلید API برای محیط تست',
    key: 'sk_test_abcdef1234567890',
    status: 'active',
    createdAt: '1402/10/10'
  }
])

// Recent API Calls
const recentApiCalls = ref([
  {
    id: 1,
    method: 'GET',
    path: '/api/products',
    status: 200,
    duration: 145,
    timestamp: '14:32:15'
  },
  {
    id: 2,
    method: 'POST',
    path: '/api/orders',
    status: 201,
    duration: 234,
    timestamp: '14:31:42'
  },
  {
    id: 3,
    method: 'GET',
    path: '/api/users/123',
    status: 404,
    duration: 89,
    timestamp: '14:30:18'
  }
])

// Tabs
const tabs = [
  { id: 'documentation', name: 'مستندات' },
  { id: 'testing', name: 'تست API' },
  { id: 'monitoring', name: 'نظارت' },
  { id: 'keys', name: 'کلیدها' }
]

// API Categories
const apiCategories = ref([
  {
    name: 'محصولات',
    description: 'مدیریت محصولات و کاتالوگ',
    endpoints: [
      {
        method: 'GET',
        path: '/api/products',
        description: 'دریافت لیست محصولات',
        parameters: [
          { name: 'page', type: 'integer', required: false },
          { name: 'limit', type: 'integer', required: false },
          { name: 'category', type: 'string', required: false }
        ],
        responseExample: `{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "محصول نمونه",
      "price": 100000,
      "category": "الکترونیک"
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 10
  }
}`
      },
      {
        method: 'POST',
        path: '/api/products',
        description: 'ایجاد محصول جدید',
        parameters: [
          { name: 'name', type: 'string', required: true },
          { name: 'price', type: 'number', required: true },
          { name: 'description', type: 'string', required: false }
        ],
        responseExample: `{
  "success": true,
  "data": {
    "id": 123,
    "name": "محصول جدید",
    "price": 150000,
    "created_at": "2024-01-15T10:30:00Z"
  }
}`
      }
    ]
  },
  {
    name: 'کاربران',
    description: 'مدیریت کاربران و احراز هویت',
    endpoints: [
      {
        method: 'GET',
        path: '/api/users',
        description: 'دریافت لیست کاربران',
        parameters: [
          { name: 'page', type: 'integer', required: false },
          { name: 'role', type: 'string', required: false }
        ],
        responseExample: `{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "احمد محمدی",
      "email": "ahmad@example.com",
      "role": "user"
    }
  ]
}`
      }
    ]
  },
  {
    name: 'سفارشات',
    description: 'مدیریت سفارشات و تراکنش‌ها',
    endpoints: [
      {
        method: 'GET',
        path: '/api/orders',
        description: 'دریافت لیست سفارشات',
        parameters: [
          { name: 'status', type: 'string', required: false },
          { name: 'user_id', type: 'integer', required: false }
        ],
        responseExample: `{
  "success": true,
  "data": [
    {
      "id": 1,
      "user_id": 123,
      "total_amount": 250000,
      "status": "pending"
    }
  ]
}`
      }
    ]
  }
])

// Methods
const getMethodColor = (method) => {
  const colors = {
    GET: 'bg-green-100 text-green-800',
    POST: 'bg-blue-100 text-blue-800',
    PUT: 'bg-yellow-100 text-yellow-800',
    DELETE: 'bg-red-100 text-red-800',
    PATCH: 'bg-purple-100 text-purple-800'
  }
  return colors[method] || 'bg-gray-100 text-gray-800'
}

const toggleEndpoint = (path) => {
  const index = expandedEndpoints.value.indexOf(path)
  if (index > -1) {
    expandedEndpoints.value.splice(index, 1)
  } else {
    expandedEndpoints.value.push(path)
  }
}

const toggleAllEndpoints = () => {
  if (allExpanded.value) {
    expandedEndpoints.value = []
  } else {
    expandedEndpoints.value = apiCategories.value.flatMap(cat => 
      cat.endpoints.map(endpoint => endpoint.path)
    )
  }
  allExpanded.value = !allExpanded.value
}

const testEndpoint = (endpoint) => {
  testRequest.value.method = endpoint.method
  testRequest.value.url = endpoint.path
  activeTab.value = 'testing'
}

const addHeader = () => {
  testRequest.value.headers.push({ key: '', value: '' })
}

const removeHeader = (index) => {
  testRequest.value.headers.splice(index, 1)
}

const sendTestRequest = async () => {
  isLoading.value = true
  try {
    const startTime = Date.now()
    
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 1000))
    
    const duration = Date.now() - startTime
    
    testResponse.value = {
      status: 200,
      duration,
      data: JSON.stringify({
        success: true,
        message: 'درخواست با موفقیت انجام شد',
        data: {
          id: 123,
          name: 'نمونه داده',
          timestamp: new Date().toISOString()
        }
      }, null, 2)
    }
  } catch (error) {
    testResponse.value = {
      status: 500,
      duration: 0,
      data: JSON.stringify({
        success: false,
        error: 'خطا در ارسال درخواست'
      }, null, 2)
    }
  } finally {
    isLoading.value = false
  }
}

const generateNewKey = () => {
  const newKey = {
    id: Date.now(),
    name: `کلید جدید ${apiKeys.value.length + 1}`,
    description: 'کلید API جدید',
    key: `sk_${Math.random().toString(36).substr(2, 9)}_${Date.now()}`,
    status: 'active',
    createdAt: new Date().toLocaleDateString('fa-IR')
  }
  apiKeys.value.unshift(newKey)
}

const toggleKeyStatus = (keyId) => {
  const key = apiKeys.value.find(k => k.id === keyId)
  if (key) {
    key.status = key.status === 'active' ? 'inactive' : 'active'
  }
}

const deleteKey = (keyId) => {
  if (confirm('آیا از حذف این کلید اطمینان دارید؟')) {
    apiKeys.value = apiKeys.value.filter(k => k.id !== keyId)
  }
}

const exportApiDocs = () => {
  alert('خروجی PDF در حال آماده‌سازی است...')
}

const refreshApiStatus = async () => {
  try {
    // Simulate API call to refresh stats
    await new Promise(resolve => setTimeout(resolve, 1000))
    
    // Update stats with new values
    apiStats.value = {
      totalEndpoints: 24 + Math.floor(Math.random() * 5),
      activeEndpoints: 22 + Math.floor(Math.random() * 3),
      todayCalls: 15420 + Math.floor(Math.random() * 1000),
      errorRate: (2.3 + Math.random() * 2).toFixed(1)
    }
    
    alert('آمار API به‌روزرسانی شد')
  } catch (error) {
    alert('خطا در به‌روزرسانی آمار')
  }
}

// Lifecycle
onMounted(() => {
  // Initialize with some endpoints expanded
  expandedEndpoints.value = ['/api/products', '/api/users']
})
</script>

<style scoped>
/* Custom styles for API documentation */
pre {
  direction: ltr;
  text-align: left;
}

code {
  direction: ltr;
  text-align: left;
}
</style> 
