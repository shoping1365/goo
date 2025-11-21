<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">مستندات API</h4>
        <p class="text-sm text-gray-600 mt-1">راهنمای کامل API های اتصال حسابداری</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          @click="downloadDocs"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
          </svg>
          دانلود مستندات
        </button>
      </div>
    </div>

    <!-- جستجو -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <div class="relative">
        <input
          v-model="searchQuery"
          type="text"
          placeholder="جستجو در مستندات API..."
          class="w-full pl-10 pr-4 py-3 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
        />
        <div class="absolute inset-y-0 right-0 pr-3 flex items-center pointer-events-none">
          <svg class="h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
        </div>
      </div>
    </div>

    <!-- دسته‌بندی‌ها -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">دسته‌بندی‌ها</h5>
      <div class="flex flex-wrap gap-2">
        <button
          v-for="category in categories"
          :key="category.id"
          class="px-4 py-2 text-sm font-medium rounded-full transition-colors"
          :class="selectedCategory === category.id
            ? 'bg-blue-100 text-blue-700 border border-blue-200'
            : 'bg-gray-100 text-gray-700 hover:bg-gray-200 border border-gray-200'"
          @click="selectCategory(category.id)"
        >
          {{ category.name }}
          <span class="text-xs text-gray-500 mr-1">({{ category.count }})</span>
        </button>
      </div>
    </div>

    <!-- مستندات API -->
    <div class="space-y-6">
      <div
        v-for="api in filteredAPIs"
        :key="api.id"
        class="bg-white rounded-lg border border-gray-200 overflow-hidden"
      >
        <div class="px-6 py-4 border-b border-gray-200">
          <div class="flex items-center justify-between">
            <div class="flex items-center space-x-3 space-x-reverse">
              <span
                class="px-2 py-1 text-xs font-medium rounded-full"
                :class="getMethodClass(api.method)"
              >
                {{ api.method }}
              </span>
              <h6 class="text-sm font-medium text-gray-900">{{ api.title }}</h6>
              <span
                class="px-2 py-1 text-xs font-medium rounded-full"
                :class="getCategoryClass(api.category)"
              >
                {{ getCategoryLabel(api.category) }}
              </span>
            </div>
            <button
              class="text-gray-400 hover:text-gray-600"
              @click="toggleAPI(api.id)"
            >
              <svg
                class="w-5 h-5 transition-transform"
                :class="{ 'rotate-180': openAPIs.includes(api.id) }"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
              </svg>
            </button>
          </div>
          <p class="text-sm text-gray-600 mt-2">{{ api.description }}</p>
        </div>

        <div
          v-show="openAPIs.includes(api.id)"
          class="px-6 py-4"
        >
          <!-- URL -->
          <div class="mb-4">
            <h6 class="text-sm font-medium text-gray-900 mb-2">URL</h6>
            <div class="bg-gray-50 rounded-lg p-3">
              <code class="text-sm text-gray-800">{{ api.url }}</code>
            </div>
          </div>

          <!-- پارامترها -->
          <div v-if="api.parameters && api.parameters.length > 0" class="mb-4">
            <h6 class="text-sm font-medium text-gray-900 mb-2">پارامترها</h6>
            <div class="overflow-x-auto">
              <table class="min-w-full divide-y divide-gray-200">
                <thead class="bg-gray-50">
                  <tr>
                    <th class="px-3 py-2 text-right text-xs font-medium text-gray-500">نام</th>
                    <th class="px-3 py-2 text-right text-xs font-medium text-gray-500">نوع</th>
                    <th class="px-3 py-2 text-right text-xs font-medium text-gray-500">اجباری</th>
                    <th class="px-3 py-2 text-right text-xs font-medium text-gray-500">توضیحات</th>
                  </tr>
                </thead>
                <tbody class="bg-white divide-y divide-gray-200">
                  <tr v-for="param in api.parameters" :key="param.name">
                    <td class="px-3 py-2 text-sm text-gray-900">{{ param.name }}</td>
                    <td class="px-3 py-2 text-sm text-gray-600">{{ param.type }}</td>
                    <td class="px-3 py-2 text-sm text-gray-600">
                      <span :class="param.required ? 'text-red-600' : 'text-green-600'">
                        {{ param.required ? 'بله' : 'خیر' }}
                      </span>
                    </td>
                    <td class="px-3 py-2 text-sm text-gray-600">{{ param.description }}</td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>

          <!-- نمونه درخواست -->
          <div class="mb-4">
            <h6 class="text-sm font-medium text-gray-900 mb-2">نمونه درخواست</h6>
            <div class="bg-gray-900 rounded-lg p-6">
              <pre class="text-sm text-gray-300 overflow-x-auto"><code>{{ api.requestExample }}</code></pre>
            </div>
          </div>

          <!-- نمونه پاسخ -->
          <div class="mb-4">
            <h6 class="text-sm font-medium text-gray-900 mb-2">نمونه پاسخ</h6>
            <div class="bg-gray-900 rounded-lg p-6">
              <pre class="text-sm text-gray-300 overflow-x-auto"><code>{{ api.responseExample }}</code></pre>
            </div>
          </div>

          <!-- کدهای خطا -->
          <div v-if="api.errorCodes && api.errorCodes.length > 0" class="mb-4">
            <h6 class="text-sm font-medium text-gray-900 mb-2">کدهای خطا</h6>
            <div class="overflow-x-auto">
              <table class="min-w-full divide-y divide-gray-200">
                <thead class="bg-gray-50">
                  <tr>
                    <th class="px-3 py-2 text-right text-xs font-medium text-gray-500">کد</th>
                    <th class="px-3 py-2 text-right text-xs font-medium text-gray-500">پیام</th>
                    <th class="px-3 py-2 text-right text-xs font-medium text-gray-500">توضیحات</th>
                  </tr>
                </thead>
                <tbody class="bg-white divide-y divide-gray-200">
                  <tr v-for="error in api.errorCodes" :key="error.code">
                    <td class="px-3 py-2 text-sm text-gray-900">{{ error.code }}</td>
                    <td class="px-3 py-2 text-sm text-gray-600">{{ error.message }}</td>
                    <td class="px-3 py-2 text-sm text-gray-600">{{ error.description }}</td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>

          <!-- تست API -->
          <div class="border-t border-gray-200 pt-4">
            <h6 class="text-sm font-medium text-gray-900 mb-3">تست API</h6>
            <div class="space-y-3">
              <div class="grid grid-cols-1 md:grid-cols-2 gap-3">
                <input
                  v-model="testParams[api.id]"
                  type="text"
                  :placeholder="`پارامترهای تست برای ${api.title}`"
                  class="px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                />
                <button
                  class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
                  @click="testAPI(api)"
                >
                  <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.828 14.828a4 4 0 01-5.656 0M9 10h1m4 0h1m-6 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                  تست
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- نمونه کدها -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">نمونه کدها</h5>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <h6 class="text-sm font-medium text-gray-900 mb-2">JavaScript</h6>
          <div class="bg-gray-900 rounded-lg p-6">
            <pre class="text-sm text-gray-300 overflow-x-auto"><code>{{ javascriptExample }}</code></pre>
          </div>
        </div>
        <div>
          <h6 class="text-sm font-medium text-gray-900 mb-2">PHP</h6>
          <div class="bg-gray-900 rounded-lg p-6">
            <pre class="text-sm text-gray-300 overflow-x-auto"><code>{{ phpExample }}</code></pre>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

// جستجو و فیلتر
const searchQuery = ref('')
const selectedCategory = ref('')
const openAPIs = ref<number[]>([])
const testParams = ref<Record<number, string>>({})

// دسته‌بندی‌ها
const categories = ref([
  { id: 'authentication', name: 'احراز هویت', count: 5 },
  { id: 'connections', name: 'اتصالات', count: 8 },
  { id: 'sync', name: 'همگام‌سازی', count: 12 },
  { id: 'transactions', name: 'تراکنش‌ها', count: 10 },
  { id: 'reports', name: 'گزارش‌ها', count: 6 }
])

// API ها
const apis = ref([
  {
    id: 1,
    title: 'ایجاد اتصال جدید',
    description: 'ایجاد اتصال جدید به نرم‌افزار حسابداری',
    method: 'POST',
    category: 'connections',
    url: '/api/v1/connections',
    parameters: [
      { name: 'name', type: 'string', required: true, description: 'نام اتصال' },
      { name: 'type', type: 'string', required: true, description: 'نوع نرم‌افزار' },
      { name: 'config', type: 'object', required: true, description: 'تنظیمات اتصال' }
    ],
    requestExample: `{
  "name": "اتصال هلو",
  "type": "helo",
  "config": {
    "api_key": "your_api_key",
    "server_url": "https://server.com"
  }
}`,
    responseExample: `{
  "success": true,
  "data": {
    "id": 123,
    "name": "اتصال هلو",
    "status": "active"
  }
}`,
    errorCodes: [
      { code: 400, message: 'Bad Request', description: 'پارامترهای نامعتبر' },
      { code: 401, message: 'Unauthorized', description: 'کلید API نامعتبر' }
    ]
  },
  {
    id: 2,
    title: 'همگام‌سازی تراکنش‌ها',
    description: 'همگام‌سازی تراکنش‌ها از نرم‌افزار حسابداری',
    method: 'GET',
    category: 'sync',
    url: '/api/v1/sync/transactions',
    parameters: [
      { name: 'connection_id', type: 'integer', required: true, description: 'شناسه اتصال' },
      { name: 'from_date', type: 'string', required: false, description: 'تاریخ شروع' },
      { name: 'to_date', type: 'string', required: false, description: 'تاریخ پایان' }
    ],
    requestExample: `GET /api/v1/sync/transactions?connection_id=123&from_date=2024-01-01`,
    responseExample: `{
  "success": true,
  "data": {
    "transactions": [...],
    "total_count": 150
  }
}`,
    errorCodes: [
      { code: 404, message: 'Not Found', description: 'اتصال یافت نشد' },
      { code: 500, message: 'Internal Server Error', description: 'خطای سرور' }
    ]
  }
])

// نمونه کدها
const javascriptExample = `// نمونه کد JavaScript
const response = await fetch('/api/v1/connections', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
    'Authorization': 'Bearer your_token'
  },
  body: JSON.stringify({
    name: 'اتصال هلو',
    type: 'helo',
    config: {
      api_key: 'your_api_key',
      server_url: 'https://server.com'
    }
  })
});
const data = await response.json();`

const phpExample = `<?php
// نمونه کد PHP
$url = 'https://api.example.com/v1/connections';
$data = [
    'name' => 'اتصال هلو',
    'type' => 'helo',
    'config' => [
        'api_key' => 'your_api_key',
        'server_url' => 'https://server.com'
    ]
];

$ch = curl_init();
curl_setopt($ch, CURLOPT_URL, $url);
curl_setopt($ch, CURLOPT_POST, true);
curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($data));
curl_setopt($ch, CURLOPT_HTTPHEADER, [
    'Content-Type: application/json',
    'Authorization: Bearer your_token'
]);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
$response = curl_exec($ch);
curl_close($ch);
$result = json_decode($response, true);
echo json_encode($result, JSON_PRETTY_PRINT);
?>`

// فیلتر کردن API ها
const filteredAPIs = computed(() => {
  return apis.value.filter(api => {
    if (searchQuery.value && !api.title.includes(searchQuery.value) && !api.description.includes(searchQuery.value)) return false
    if (selectedCategory.value && api.category !== selectedCategory.value) return false
    return true
  })
})

// کلاس متد
const getMethodClass = (method: string) => {
  const classes: Record<string, string> = {
    GET: 'bg-green-100 text-green-700',
    POST: 'bg-blue-100 text-blue-700',
    PUT: 'bg-yellow-100 text-yellow-700',
    DELETE: 'bg-red-100 text-red-700'
  }
  return classes[method] || 'bg-gray-100 text-gray-700'
}

// کلاس دسته‌بندی
const getCategoryClass = (category: string) => {
  const classes: Record<string, string> = {
    authentication: 'bg-purple-100 text-purple-700',
    connections: 'bg-blue-100 text-blue-700',
    sync: 'bg-green-100 text-green-700',
    transactions: 'bg-yellow-100 text-yellow-700',
    reports: 'bg-orange-100 text-orange-700'
  }
  return classes[category] || 'bg-gray-100 text-gray-700'
}

// برچسب دسته‌بندی
const getCategoryLabel = (category: string) => {
  const labels: Record<string, string> = {
    authentication: 'احراز هویت',
    connections: 'اتصالات',
    sync: 'همگام‌سازی',
    transactions: 'تراکنش‌ها',
    reports: 'گزارش‌ها'
  }
  return labels[category] || category
}

// انتخاب دسته‌بندی
const selectCategory = (categoryId: string) => {
  selectedCategory.value = selectedCategory.value === categoryId ? '' : categoryId
}

// باز/بسته کردن API
const toggleAPI = (apiId: number) => {
  const index = openAPIs.value.indexOf(apiId)
  if (index > -1) {
    openAPIs.value.splice(index, 1)
  } else {
    openAPIs.value.push(apiId)
  }
}

// تست API
const testAPI = (_api: unknown) => {
  // TODO: Implement API testing
}

// دانلود مستندات
const downloadDocs = () => {
  // TODO: Implement documentation download
}
</script>

<!--
  کامپوننت مستندات API
  شامل:
  1. جستجو و فیلتر
  2. دسته‌بندی‌ها
  3. مستندات کامل API ها
  4. نمونه کدها
  5. تست API
  6. طراحی مدرن و کاملاً ریسپانسیو
-->
