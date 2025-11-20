<template>
  <div class="min-h-screen bg-gray-50 p-6">
    <div class="max-w-7xl mx-auto">
      <!-- Header -->
      <div class="mb-8">
        <h1 class="text-3xl font-bold text-gray-900 mb-2">تست پیشرفته اتصال</h1>
        <p class="text-gray-600">ابزارهای تست و دیباگ اتصالات خارجی و سرویس‌ها</p>
      </div>

      <!-- Stats Cards -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-green-100 rounded-lg">
              <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">اتصالات موفق</p>
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
              <p class="text-sm font-medium text-gray-600">اتصالات ناموفق</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.failed }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-blue-100 rounded-lg">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
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
        <!-- Connection Tester -->
        <div class="bg-white rounded-lg shadow">
          <div class="p-6 border-b border-gray-200">
            <h2 class="text-xl font-semibold text-gray-900">تست اتصال</h2>
          </div>
          <div class="p-6">
            <div class="space-y-4">
              <!-- Connection Type -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">نوع اتصال</label>
                <select v-model="connectionConfig.type" class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500">
                  <option value="http">HTTP/HTTPS</option>
                  <option value="database">Database</option>
                  <option value="redis">Redis</option>
                  <option value="smtp">SMTP</option>
                  <option value="ftp">FTP</option>
                  <option value="ssh">SSH</option>
                </select>
              </div>

              <!-- Host Input -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">آدرس سرور</label>
                <input 
                  v-model="connectionConfig.host" 
                  type="text" 
                  placeholder="example.com یا 192.168.1.1"
                  class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
              </div>

              <!-- Port Input -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">پورت</label>
                <input 
                  v-model="connectionConfig.port" 
                  type="number" 
                  placeholder="80, 443, 3306, 6379..."
                  class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
              </div>

              <!-- Credentials -->
              <div v-if="connectionConfig.type !== 'http'">
                <label class="block text-sm font-medium text-gray-700 mb-2">نام کاربری</label>
                <input 
                  v-model="connectionConfig.username" 
                  type="text" 
                  placeholder="نام کاربری"
                  class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
              </div>

              <div v-if="connectionConfig.type !== 'http'">
                <label class="block text-sm font-medium text-gray-700 mb-2">رمز عبور</label>
                <input 
                  v-model="connectionConfig.password" 
                  type="password" 
                  placeholder="رمز عبور"
                  class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
              </div>

              <!-- Database Name (for database connections) -->
              <div v-if="connectionConfig.type === 'database'">
                <label class="block text-sm font-medium text-gray-700 mb-2">نام دیتابیس</label>
                <input 
                  v-model="connectionConfig.database" 
                  type="text" 
                  placeholder="نام دیتابیس"
                  class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
              </div>

              <!-- Test Button -->
              <button 
                :disabled="testing" 
                class="w-full bg-blue-600 hover:bg-blue-700 disabled:bg-gray-400 text-white font-medium py-3 px-4 rounded-lg transition-colors"
                @click="testConnection"
              >
                <span v-if="testing">در حال تست...</span>
                <span v-else>تست اتصال</span>
              </button>
            </div>
          </div>
        </div>

        <!-- Results Viewer -->
        <div class="bg-white rounded-lg shadow">
          <div class="p-6 border-b border-gray-200">
            <h2 class="text-xl font-semibold text-gray-900">نتایج تست</h2>
          </div>
          <div class="p-6">
            <div v-if="testResult" class="space-y-4">
              <!-- Status -->
              <div class="flex items-center justify-between">
                <span class="text-sm font-medium text-gray-700">وضعیت:</span>
                <span
:class="[
                  'px-2 py-1 rounded text-sm font-medium',
                  testResult.success ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'
                ]">
                  {{ testResult.success ? 'موفق' : 'ناموفق' }}
                </span>
              </div>

              <!-- Response Time -->
              <div class="flex items-center justify-between">
                <span class="text-sm font-medium text-gray-700">زمان پاسخ:</span>
                <span class="text-sm text-gray-600">{{ testResult.responseTime }}ms</span>
              </div>

              <!-- Error Message -->
              <div v-if="testResult.error">
                <label class="block text-sm font-medium text-gray-700 mb-2">پیام خطا</label>
                <pre class="bg-red-50 border border-red-200 p-3 rounded text-sm text-red-800 overflow-x-auto">{{ testResult.error }}</pre>
              </div>

              <!-- Connection Details -->
              <div v-if="testResult.details">
                <label class="block text-sm font-medium text-gray-700 mb-2">جزئیات اتصال</label>
                <pre class="bg-gray-100 p-3 rounded text-sm overflow-x-auto">{{ JSON.stringify(testResult.details, null, 2) }}</pre>
              </div>
            </div>
            <div v-else class="text-center text-gray-500 py-8">
              <svg class="w-12 h-12 mx-auto mb-4 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.111 16.404a5.5 5.5 0 017.778 0M12 20h.01m-7.08-7.071c3.904-3.905 10.236-3.905 14.141 0M1.394 9.393c5.857-5.857 15.355-5.857 21.213 0"></path>
              </svg>
              <p>نتیجه‌ای برای نمایش وجود ندارد</p>
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
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
            <button 
              class="bg-blue-500 hover:bg-blue-600 text-white font-medium py-3 px-4 rounded-lg transition-colors"
              @click="quickTest('database')"
            >
              تست دیتابیس
            </button>
            <button 
              class="bg-red-500 hover:bg-red-600 text-white font-medium py-3 px-4 rounded-lg transition-colors"
              @click="quickTest('redis')"
            >
              تست Redis
            </button>
            <button 
              class="bg-green-500 hover:bg-green-600 text-white font-medium py-3 px-4 rounded-lg transition-colors"
              @click="quickTest('smtp')"
            >
              تست SMTP
            </button>
            <button 
              class="bg-purple-500 hover:bg-purple-600 text-white font-medium py-3 px-4 rounded-lg transition-colors"
              @click="quickTest('api')"
            >
              تست API خارجی
            </button>
          </div>
        </div>
      </div>

      <!-- Connection History -->
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
                  <p class="font-medium">{{ test.type }} - {{ test.host }}:{{ test.port }}</p>
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

      <!-- Network Diagnostics -->
      <div class="mt-8 bg-white rounded-lg shadow">
        <div class="p-6 border-b border-gray-200">
          <h2 class="text-xl font-semibold text-gray-900">تشخیص شبکه</h2>
        </div>
        <div class="p-6">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <!-- Ping Test -->
            <div class="space-y-4">
              <h3 class="font-medium text-gray-900">تست Ping</h3>
              <div class="flex space-x-2 space-x-reverse">
                <input 
                  v-model="pingHost"
                  type="text"
                  placeholder="example.com"
                  class="flex-1 border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                                 <button 
                   class="bg-green-500 hover:bg-green-600 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                   @click="performPing"
                 >
                   Ping
                 </button>
              </div>
              <pre v-if="pingResult" class="bg-gray-100 p-3 rounded text-sm overflow-x-auto">{{ pingResult }}</pre>
            </div>

            <!-- DNS Lookup -->
            <div class="space-y-4">
              <h3 class="font-medium text-gray-900">جستجوی DNS</h3>
              <div class="flex space-x-2 space-x-reverse">
                <input 
                  v-model="dnsHost"
                  type="text"
                  placeholder="example.com"
                  class="flex-1 border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                <button 
                  class="bg-blue-500 hover:bg-blue-600 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="lookupDNS"
                >
                  جستجو
                </button>
              </div>
              <pre v-if="dnsResult" class="bg-gray-100 p-3 rounded text-sm overflow-x-auto">{{ dnsResult }}</pre>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, watch } from 'vue'

definePageMeta({
  layout: 'admin-main',
  middleware: ['developer-only']
});

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user, hasPermission } = useAuth()

// Reactive data
const testing = ref(false)
const testResult = ref(null)
const testHistory = ref([])
const pingHost = ref('')
const pingResult = ref('')
const dnsHost = ref('')
const dnsResult = ref('')

const stats = reactive({
  success: 0,
  failed: 0,
  total: 0,
  avgResponseTime: 0
})

const connectionConfig = reactive({
  type: 'http',
  host: '',
  port: '',
  username: '',
  password: '',
  database: ''
})

// Methods
async function testConnection() {
  if (!connectionConfig.host) {
    alert('لطفاً آدرس سرور را وارد کنید')
    return
  }

  testing.value = true
  const startTime = Date.now()

  try {
    // Simulate connection test based on type
    await new Promise(resolve => setTimeout(resolve, 1000 + Math.random() * 2000))

    const responseTime = Date.now() - startTime
    const success = Math.random() > 0.3 // 70% success rate for demo

    testResult.value = {
      success,
      responseTime,
      error: success ? null : 'خطا در اتصال به سرور',
      details: {
        type: connectionConfig.type,
        host: connectionConfig.host,
        port: connectionConfig.port,
        connectionTime: responseTime,
        timestamp: new Date().toISOString()
      }
    }

    // Update stats
    if (success) {
      stats.success++
    } else {
      stats.failed++
    }
    stats.total++
    stats.avgResponseTime = Math.round((stats.avgResponseTime * (stats.total - 1) + responseTime) / stats.total)

    // Add to history
    addToHistory({
      type: connectionConfig.type,
      host: connectionConfig.host,
      port: connectionConfig.port,
      responseTime,
      success
    })

  } catch (error) {
    const responseTime = Date.now() - startTime
    
    testResult.value = {
      success: false,
      responseTime,
      error: error.message,
      details: null
    }

    stats.failed++
    stats.total++
  } finally {
    testing.value = false
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
  connectionConfig.type = test.type
  connectionConfig.host = test.host
  connectionConfig.port = test.port
}

async function quickTest(type) {
  const tests = {
    database: {
      type: 'database',
      host: 'localhost',
      port: '5432',
      username: 'postgres',
      password: 'password',
      database: 'test_db'
    },
    redis: {
      type: 'redis',
      host: 'localhost',
      port: '6379',
      username: '',
      password: '',
      database: ''
    },
    smtp: {
      type: 'smtp',
      host: 'smtp.gmail.com',
      port: '587',
      username: 'test@gmail.com',
      password: 'password',
      database: ''
    },
    api: {
      type: 'http',
      host: 'https://api.example.com',
      port: '443',
      username: '',
      password: '',
      database: ''
    }
  }

  const test = tests[type]
  if (test) {
    Object.assign(connectionConfig, test)
    await testConnection()
  }
}

async function performPing() {
  if (!pingHost.value) {
    alert('لطفاً آدرس را وارد کنید')
    return
  }

  try {
    // Simulate ping
    await new Promise(resolve => setTimeout(resolve, 1000))
    pingResult.value = `PING ${pingHost.value} (192.168.1.1): 56 data bytes\n64 bytes from 192.168.1.1: icmp_seq=0 ttl=64 time=1.234 ms\n64 bytes from 192.168.1.1: icmp_seq=1 ttl=64 time=1.456 ms\n64 bytes from 192.168.1.1: icmp_seq=2 ttl=64 time=1.345 ms\n\n--- ${pingHost.value} ping statistics ---\n3 packets transmitted, 3 packets received, 0.0% packet loss\nround-trip min/avg/max/stddev = 1.234/1.345/1.456/0.111 ms`
  } catch (error) {
    pingResult.value = `Error: ${error.message}`
  }
}

async function lookupDNS() {
  if (!dnsHost.value) {
    alert('لطفاً آدرس را وارد کنید')
    return
  }

  try {
    // Simulate DNS lookup
    await new Promise(resolve => setTimeout(resolve, 500))
    dnsResult.value = `Name: ${dnsHost.value}\nAddress: 192.168.1.1\nAliases: www.${dnsHost.value}\n\nName: ${dnsHost.value}\nAddress: 2001:db8::1`
  } catch (error) {
    dnsResult.value = `Error: ${error.message}`
  }
}

// Load initial stats
onMounted(() => {
  // Load stats from localStorage
  const savedStats = localStorage.getItem('connection-test-stats')
  if (savedStats) {
    Object.assign(stats, JSON.parse(savedStats))
  }

  // Load test history
  const savedHistory = localStorage.getItem('connection-test-history')
  if (savedHistory) {
    testHistory.value = JSON.parse(savedHistory)
  }
})

// Save stats and history to localStorage
watch([stats, testHistory], () => {
  localStorage.setItem('connection-test-stats', JSON.stringify(stats))
  localStorage.setItem('connection-test-history', JSON.stringify(testHistory.value))
}, { deep: true })
</script>

<style scoped>
.group:hover {
  box-shadow: 0 0 0 4px #e0e7ff;
}
</style>
