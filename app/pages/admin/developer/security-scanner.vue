<template>
  <div class="min-h-screen bg-gray-50 p-6">
    <div class="max-w-7xl mx-auto">
      <!-- Header -->
      <div class="mb-8">
        <h1 class="text-3xl font-bold text-gray-900 mb-2">اسکنر امنیتی</h1>
        <p class="text-gray-600">بررسی امنیت سیستم، کد و آسیب‌پذیری‌ها</p>
      </div>

      <!-- Security Overview -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-red-100 rounded-lg">
              <svg class="w-6 h-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
              </svg>
            </div>
            <div class="mr-4">
              <p class="text-sm font-medium text-gray-600">آسیب‌پذیری‌های بحرانی</p>
              <p class="text-2xl font-bold text-red-600">{{ securityStats.critical }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-yellow-100 rounded-lg">
              <svg class="w-6 h-6 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
              </svg>
            </div>
            <div class="mr-4">
              <p class="text-sm font-medium text-gray-600">آسیب‌پذیری‌های مهم</p>
              <p class="text-2xl font-bold text-yellow-600">{{ securityStats.high }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-blue-100 rounded-lg">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
            <div class="mr-4">
              <p class="text-sm font-medium text-gray-600">آسیب‌پذیری‌های متوسط</p>
              <p class="text-2xl font-bold text-blue-600">{{ securityStats.medium }}</p>
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
              <p class="text-sm font-medium text-gray-600">امنیت کلی</p>
              <p class="text-2xl font-bold text-green-600">{{ securityStats.score }}%</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Quick Scan Actions -->
      <div class="bg-white rounded-lg shadow mb-8">
        <div class="px-6 py-4 border-b border-gray-200">
          <h2 class="text-lg font-semibold text-gray-900">عملیات اسکن سریع</h2>
        </div>
        <div class="p-6">
          <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
            <button @click="startScan('code')" :disabled="isScanning" 
                    class="flex items-center justify-center px-4 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed transition-colors">
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4"></path>
              </svg>
              اسکن کد
            </button>

            <button @click="startScan('dependencies')" :disabled="isScanning"
                    class="flex items-center justify-center px-4 py-3 bg-green-600 text-white rounded-lg hover:bg-green-700 disabled:opacity-50 disabled:cursor-not-allowed transition-colors">
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                
              </svg>
              اسکن وابستگی‌ها
            </button>

            <button @click="startScan('full')" :disabled="isScanning"
                    class="flex items-center justify-center px-4 py-3 bg-purple-600 text-white rounded-lg hover:bg-purple-700 disabled:opacity-50 disabled:cursor-not-allowed transition-colors">
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"></path>
              </svg>
              اسکن کامل
            </button>
          </div>
        </div>
      </div>

      <!-- Scan Configuration -->
      <div class="bg-white rounded-lg shadow mb-8">
        <div class="px-6 py-4 border-b border-gray-200">
          <h2 class="text-lg font-semibold text-gray-900">تنظیمات اسکن</h2>
        </div>
        <div class="p-6">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <!-- Scan Types -->
            <div>
              <h3 class="text-md font-medium text-gray-900 mb-4">نوع اسکن</h3>
              <div class="space-y-3">
                <div class="flex items-center">
                  <input v-model="scanConfig.types.code" type="checkbox" id="scan-code" 
                         class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded">
                  <label for="scan-code" class="mr-2 text-sm text-gray-700">اسکن کد (SQL Injection, XSS)</label>
                </div>
                <div class="flex items-center">
                  <input v-model="scanConfig.types.dependencies" type="checkbox" id="scan-deps" 
                         class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded">
                  <label for="scan-deps" class="mr-2 text-sm text-gray-700">اسکن وابستگی‌ها (CVE)</label>
                </div>
                <div class="flex items-center">
                  <input v-model="scanConfig.types.config" type="checkbox" id="scan-config" 
                         class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded">
                  <label for="scan-config" class="mr-2 text-sm text-gray-700">اسکن تنظیمات</label>
                </div>
                <div class="flex items-center">
                  <input v-model="scanConfig.types.network" type="checkbox" id="scan-network" 
                         class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded">
                  <label for="scan-network" class="mr-2 text-sm text-gray-700">اسکن شبکه</label>
                </div>
              </div>
            </div>

            <!-- Scan Options -->
            <div>
              <h3 class="text-md font-medium text-gray-900 mb-4">گزینه‌های اسکن</h3>
              <div class="space-y-4">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">سطح اسکن</label>
                  <select v-model="scanConfig.level" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
                    <option value="low">کم</option>
                    <option value="medium">متوسط</option>
                    <option value="high">زیاد</option>
                    <option value="extreme">بسیار زیاد</option>
                  </select>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">مسیرهای شامل</label>
                  <textarea v-model="scanConfig.includePaths" rows="3" 
                            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                            placeholder="/app&#10;/src&#10;/config"></textarea>
                </div>
                <div class="flex items-center">
                  <input v-model="scanConfig.autoFix" type="checkbox" id="auto-fix" 
                         class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded">
                  <label for="auto-fix" class="mr-2 text-sm text-gray-700">تصحیح خودکار</label>
                </div>
              </div>
            </div>
          </div>

          <div class="mt-6 flex justify-end">
            <button @click="saveScanConfiguration" 
                    class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors">
              ذخیره تنظیمات
            </button>
          </div>
        </div>
      </div>

      <!-- Vulnerability Results -->
      <div class="bg-white rounded-lg shadow mb-8">
        <div class="px-6 py-4 border-b border-gray-200 flex justify-between items-center">
          <h2 class="text-lg font-semibold text-gray-900">نتایج آسیب‌پذیری‌ها</h2>
          <div class="flex space-x-2 space-x-reverse">
            <select v-model="filter.severity" class="px-3 py-1 border border-gray-300 rounded-md text-sm">
              <option value="">همه سطوح</option>
              <option value="critical">بحرانی</option>
              <option value="high">مهم</option>
              <option value="medium">متوسط</option>
              <option value="low">کم</option>
            </select>
            <select v-model="filter.type" class="px-3 py-1 border border-gray-300 rounded-md text-sm">
              <option value="">همه انواع</option>
              <option value="code">کد</option>
              <option value="dependency">وابستگی</option>
              <option value="config">تنظیمات</option>
              <option value="network">شبکه</option>
            </select>
          </div>
        </div>
        <div class="p-6">
          <div class="space-y-4">
            <div v-for="vulnerability in filteredVulnerabilities" :key="vulnerability.id" 
                 class="border border-gray-200 rounded-lg p-6"
                 :class="getVulnerabilityClass(vulnerability.severity)">
              <div class="flex items-start justify-between">
                <div class="flex-1">
                  <div class="flex items-center mb-2">
                    <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium mr-3"
                          :class="getSeverityBadgeClass(vulnerability.severity)">
                      {{ getSeverityLabel(vulnerability.severity) }}
                    </span>
                    <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-gray-100 text-gray-800">
                      {{ getTypeLabel(vulnerability.type) }}
                    </span>
                  </div>
                  <h3 class="text-lg font-medium text-gray-900 mb-2">{{ vulnerability.title }}</h3>
                  <p class="text-sm text-gray-600 mb-3">{{ vulnerability.description }}</p>
                  
                  <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-3">
                    <div>
                      <p class="text-xs font-medium text-gray-500 mb-1">مسیر فایل</p>
                      <p class="text-sm text-gray-900 font-mono">{{ vulnerability.filePath }}</p>
                    </div>
                    <div>
                      <p class="text-xs font-medium text-gray-500 mb-1">خط کد</p>
                      <p class="text-sm text-gray-900">{{ vulnerability.lineNumber }}</p>
                    </div>
                  </div>

                  <div v-if="vulnerability.recommendation" class="mb-3">
                    <p class="text-xs font-medium text-gray-500 mb-1">توصیه</p>
                    <p class="text-sm text-gray-900">{{ vulnerability.recommendation }}</p>
                  </div>

                  <div v-if="vulnerability.cve" class="mb-3">
                    <p class="text-xs font-medium text-gray-500 mb-1">CVE</p>
                    <a :href="`https://cve.mitre.org/cgi-bin/cvename.cgi?name=${vulnerability.cve}`" 
                       target="_blank" class="text-sm text-blue-600 hover:text-blue-800">
                      {{ vulnerability.cve }}
                    </a>
                  </div>
                </div>
                
                <div class="flex flex-col space-y-2 mr-4">
                  <button @click="fixVulnerability(vulnerability.id)" 
                          class="px-3 py-1 bg-green-600 text-white text-xs rounded hover:bg-green-700 transition-colors">
                    تصحیح
                  </button>
                  <button @click="ignoreVulnerability(vulnerability.id)" 
                          class="px-3 py-1 bg-gray-600 text-white text-xs rounded hover:bg-gray-700 transition-colors">
                    نادیده گرفتن
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Security Report -->
      <div class="bg-white rounded-lg shadow">
        <div class="px-6 py-4 border-b border-gray-200">
          <h2 class="text-lg font-semibold text-gray-900">گزارش امنیتی</h2>
        </div>
        <div class="p-6">
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
            <!-- Security Score -->
            <div>
              <h3 class="text-md font-medium text-gray-900 mb-4">امتیاز امنیتی</h3>
              <div class="relative">
                <svg class="w-32 h-32 transform -rotate-90" viewBox="0 0 36 36">
                  <path class="text-gray-200" stroke="currentColor" stroke-width="2" fill="none" d="M18 2.0845 a 15.9155 15.9155 0 0 1 0 31.831 a 15.9155 15.9155 0 0 1 0 -31.831"/>
                  <path class="text-green-600" stroke="currentColor" stroke-width="2" stroke-dasharray="100, 100" stroke-dashoffset="25" fill="none" d="M18 2.0845 a 15.9155 15.9155 0 0 1 0 31.831 a 15.9155 15.9155 0 0 1 0 -31.831"/>
                </svg>
                <span class="absolute inset-0 flex items-center justify-center text-2xl font-bold text-gray-900">
                  {{ securityStats.score }}%
                </span>
              </div>
            </div>

            <!-- Security Metrics -->
            <div>
              <h3 class="text-md font-medium text-gray-900 mb-4">معیارهای امنیتی</h3>
              <div class="space-y-4">
                <div v-for="metric in securityMetrics" :key="metric.name" class="flex justify-between items-center">
                  <span class="text-sm text-gray-600">{{ metric.name }}</span>
                  <div class="flex items-center">
                    <div class="w-32 bg-gray-200 rounded-full h-2 mr-3">
                      <div class="h-2 rounded-full" 
                           :class="getMetricColorClass(metric.value)"
                           :style="{ width: `${metric.value}%` }"></div>
                    </div>
                    <span class="text-sm font-medium text-gray-900">{{ metric.value }}%</span>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <div class="mt-6 flex justify-end space-x-2 space-x-reverse">
            <button @click="exportReport" 
                    class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors">
              صادر کردن گزارش
            </button>
            <button @click="scheduleScan" 
                    class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors">
              زمان‌بندی اسکن
            </button>
          </div>
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
const { user, hasPermission } = useAuth()

// Reactive data
const securityStats = ref({
  critical: 2,
  high: 5,
  medium: 8,
  score: 75
})

const isScanning = ref(false)

const scanConfig = ref({
  types: {
    code: true,
    dependencies: true,
    config: false,
    network: false
  },
  level: 'medium',
  includePaths: '/app\n/src\n/config',
  autoFix: false
})

const filter = ref({
  severity: '',
  type: ''
})

const vulnerabilities = ref([
  {
    id: 1,
    title: 'SQL Injection Vulnerability',
    description: 'Potential SQL injection found in user input handling',
    severity: 'critical',
    type: 'code',
    filePath: '/app/controllers/user.js',
    lineNumber: 45,
    recommendation: 'Use parameterized queries or ORM to prevent SQL injection',
    cve: 'CVE-2024-1234'
  },
  {
    id: 2,
    title: 'XSS Vulnerability',
    description: 'Cross-site scripting vulnerability in comment system',
    severity: 'high',
    type: 'code',
    filePath: '/app/views/comments.ejs',
    lineNumber: 23,
    recommendation: 'Sanitize user input and use CSP headers',
    cve: null
  },
  {
    id: 3,
    title: 'Outdated Dependency',
    description: 'lodash package version 4.17.15 has known vulnerabilities',
    severity: 'medium',
    type: 'dependency',
    filePath: '/package.json',
    lineNumber: 12,
    recommendation: 'Update to lodash version 4.17.21 or later',
    cve: 'CVE-2023-4567'
  },
  {
    id: 4,
    title: 'Weak Password Policy',
    description: 'Password policy does not enforce strong passwords',
    severity: 'medium',
    type: 'config',
    filePath: '/config/auth.js',
    lineNumber: 8,
    recommendation: 'Implement minimum password requirements',
    cve: null
  }
])

const securityMetrics = ref([
  { name: 'امنیت کد', value: 85 },
  { name: 'امنیت وابستگی‌ها', value: 70 },
  { name: 'امنیت تنظیمات', value: 90 },
  { name: 'امنیت شبکه', value: 95 }
])

// Computed
const filteredVulnerabilities = computed(() => {
  return vulnerabilities.value.filter(vuln => {
    if (filter.value.severity && vuln.severity !== filter.value.severity) return false
    if (filter.value.type && vuln.type !== filter.value.type) return false
    return true
  })
})

// Methods
const startScan = async (type) => {
  isScanning.value = true
  try {
    // Simulate scan
    await new Promise(resolve => setTimeout(resolve, 3000))
    
    // Update stats based on scan results
    if (type === 'full') {
      securityStats.value.critical = Math.floor(Math.random() * 5)
      securityStats.value.high = Math.floor(Math.random() * 10) + 3
      securityStats.value.medium = Math.floor(Math.random() * 15) + 5
      securityStats.value.score = Math.floor(Math.random() * 20) + 70
    }
    
  } catch (error) {
    console.error('Scan error:', error)
  } finally {
    isScanning.value = false
  }
}

const saveScanConfiguration = async () => {
  try {
    await new Promise(resolve => setTimeout(resolve, 1000))
    console.log('Configuration saved:', scanConfig.value)
  } catch (error) {
    console.error('Error saving configuration:', error)
  }
}

const fixVulnerability = (id) => {
  console.log('Fixing vulnerability:', id)
  // Remove from list
  const index = vulnerabilities.value.findIndex(v => v.id === id)
  if (index > -1) {
    vulnerabilities.value.splice(index, 1)
  }
}

const ignoreVulnerability = (id) => {
  console.log('Ignoring vulnerability:', id)
  // Remove from list
  const index = vulnerabilities.value.findIndex(v => v.id === id)
  if (index > -1) {
    vulnerabilities.value.splice(index, 1)
  }
}

const exportReport = () => {
  console.log('Exporting security report...')
}

const scheduleScan = () => {
  console.log('Scheduling scan...')
}

const getVulnerabilityClass = (severity) => {
  const classes = {
    critical: 'border-red-200 bg-red-50',
    high: 'border-yellow-200 bg-yellow-50',
    medium: 'border-blue-200 bg-blue-50',
    low: 'border-gray-200 bg-gray-50'
  }
  return classes[severity] || 'border-gray-200 bg-gray-50'
}

const getSeverityBadgeClass = (severity) => {
  const classes = {
    critical: 'bg-red-100 text-red-800',
    high: 'bg-yellow-100 text-yellow-800',
    medium: 'bg-blue-100 text-blue-800',
    low: 'bg-gray-100 text-gray-800'
  }
  return classes[severity] || 'bg-gray-100 text-gray-800'
}

const getSeverityLabel = (severity) => {
  const labels = {
    critical: 'بحرانی',
    high: 'مهم',
    medium: 'متوسط',
    low: 'کم'
  }
  return labels[severity] || severity
}

const getTypeLabel = (type) => {
  const labels = {
    code: 'کد',
    dependency: 'وابستگی',
    config: 'تنظیمات',
    network: 'شبکه'
  }
  return labels[type] || type
}

const getMetricColorClass = (value) => {
  if (value >= 90) return 'bg-green-600'
  if (value >= 70) return 'bg-yellow-600'
  return 'bg-red-600'
}
</script> 
