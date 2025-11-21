<template>
  <div class="min-h-screen bg-gray-50 p-6">
    <div class="max-w-7xl mx-auto">
      <!-- Header -->
      <div class="mb-8">
        <h1 class="text-3xl font-bold text-gray-900 mb-2">تحلیل‌گر عملکرد</h1>
        <p class="text-gray-600">تحلیل و بهینه‌سازی عملکرد سیستم و کد</p>
      </div>

      <!-- Performance Overview -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-blue-100 rounded-lg">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
              </svg>
            </div>
            <div class="mr-4">
              <p class="text-sm font-medium text-gray-600">زمان پاسخ‌دهی</p>
              <p class="text-2xl font-bold text-gray-900">{{ performanceStats.responseTime }}ms</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-green-100 rounded-lg">
              <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
              </svg>
            </div>
            <div class="mr-4">
              <p class="text-sm font-medium text-gray-600">درخواست در ثانیه</p>
              <p class="text-2xl font-bold text-gray-900">{{ performanceStats.requestsPerSecond }}</p>
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
              <p class="text-sm font-medium text-gray-600">زمان بارگذاری</p>
              <p class="text-2xl font-bold text-gray-900">{{ performanceStats.loadTime }}s</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-purple-100 rounded-lg">
              <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
            <div class="mr-4">
              <p class="text-sm font-medium text-gray-600">امتیاز عملکرد</p>
              <p class="text-2xl font-bold text-gray-900">{{ performanceStats.score }}/100</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Performance Charts -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-8 mb-8">
        <!-- Response Time Chart -->
        <div class="bg-white rounded-lg shadow">
          <div class="px-6 py-4 border-b border-gray-200">
            <h2 class="text-lg font-semibold text-gray-900">زمان پاسخ‌دهی</h2>
          </div>
          <div class="p-6">
            <div class="h-64 flex items-end justify-between space-x-1">
              <div
v-for="(value, index) in responseTimeHistory" :key="index" 
                   class="flex-1 bg-blue-500 rounded-t" 
                   :style="{ height: `${(value / 1000) * 100}%` }">
              </div>
            </div>
            <div class="mt-4 flex justify-between text-sm text-gray-600">
              <span>0ms</span>
              <span>500ms</span>
              <span>1000ms</span>
            </div>
          </div>
        </div>

        <!-- CPU Usage Chart -->
        <div class="bg-white rounded-lg shadow">
          <div class="px-6 py-4 border-b border-gray-200">
            <h2 class="text-lg font-semibold text-gray-900">استفاده از CPU</h2>
          </div>
          <div class="p-6">
            <div class="h-64 flex items-end justify-between space-x-1">
              <div
v-for="(value, index) in cpuUsageHistory" :key="index" 
                   class="flex-1 bg-green-500 rounded-t" 
                   :style="{ height: `${value}%` }">
              </div>
            </div>
            <div class="mt-4 flex justify-between text-sm text-gray-600">
              <span>0%</span>
              <span>50%</span>
              <span>100%</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Performance Analysis -->
      <div class="bg-white rounded-lg shadow mb-8">
        <div class="px-6 py-4 border-b border-gray-200">
          <h2 class="text-lg font-semibold text-gray-900">تحلیل عملکرد</h2>
        </div>
        <div class="p-6">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <!-- Bottlenecks -->
            <div>
              <h3 class="text-md font-medium text-gray-900 mb-4">گلوگاه‌های عملکرد</h3>
              <div class="space-y-3">
                <div
v-for="bottleneck in bottlenecks" :key="bottleneck.id" 
                     class="flex items-center justify-between p-3 border border-gray-200 rounded-lg">
                  <div class="flex items-center">
                    <div
class="w-3 h-3 rounded-full mr-3" 
                         :class="getBottleneckColor(bottleneck.severity)"></div>
                    <div>
                      <p class="font-medium text-gray-900">{{ bottleneck.name }}</p>
                      <p class="text-sm text-gray-600">{{ bottleneck.description }}</p>
                    </div>
                  </div>
                  <span
class="text-sm font-medium" 
                        :class="getBottleneckTextColor(bottleneck.severity)">
                    {{ bottleneck.impact }}
                  </span>
                </div>
              </div>
            </div>

            <!-- Optimization Suggestions -->
            <div>
              <h3 class="text-md font-medium text-gray-900 mb-4">پیشنهادات بهینه‌سازی</h3>
              <div class="space-y-3">
                <div
v-for="suggestion in optimizationSuggestions" :key="suggestion.id" 
                     class="p-3 border border-gray-200 rounded-lg">
                  <div class="flex items-start justify-between">
                    <div class="flex-1">
                      <h4 class="font-medium text-gray-900 mb-1">{{ suggestion.title }}</h4>
                      <p class="text-sm text-gray-600 mb-2">{{ suggestion.description }}</p>
                      <div class="flex items-center text-xs text-gray-500">
                        <span>تأثیر: {{ suggestion.impact }}</span>
                        <span class="mx-2">•</span>
                        <span>پیچیدگی: {{ suggestion.complexity }}</span>
                      </div>
                    </div>
                    <button
class="px-3 py-1 bg-blue-600 text-white text-xs rounded hover:bg-blue-700 transition-colors" 
                            @click="applyOptimization(suggestion.id)">
                      اعمال
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Code Performance -->
      <div class="bg-white rounded-lg shadow mb-8">
        <div class="px-6 py-4 border-b border-gray-200">
          <h2 class="text-lg font-semibold text-gray-900">عملکرد کد</h2>
        </div>
        <div class="p-6">
          <div class="overflow-x-auto">
            <table class="min-w-full divide-y divide-gray-200">
              <thead class="bg-gray-50">
                <tr>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">فایل</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملکرد</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">حجم</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">زمان اجرا</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
                </tr>
              </thead>
              <tbody class="bg-white divide-y divide-gray-200">
                <tr v-for="file in codePerformance" :key="file.name" class="hover:bg-gray-50">
                  <td class="px-6 py-4 whitespace-nowrap">
                    <div class="flex items-center">
                      <span class="text-sm font-medium text-gray-900">{{ file.name }}</span>
                    </div>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap">
                    <div class="flex items-center">
                      <div class="w-16 bg-gray-200 rounded-full h-2 mr-2">
                        <div
class="h-2 rounded-full" 
                             :class="getPerformanceColor(file.performance)"
                             :style="{ width: `${file.performance}%` }"></div>
                      </div>
                      <span class="text-sm text-gray-900">{{ file.performance }}%</span>
                    </div>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                    {{ formatSize(file.size) }}
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                    {{ file.executionTime }}ms
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                    <button
class="text-blue-600 hover:text-blue-900" 
                            @click="optimizeFile(file.name)">بهینه‌سازی</button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>

      <!-- Database Performance -->
      <div class="bg-white rounded-lg shadow mb-8">
        <div class="px-6 py-4 border-b border-gray-200">
          <h2 class="text-lg font-semibold text-gray-900">عملکرد دیتابیس</h2>
        </div>
        <div class="p-6">
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
            <!-- Slow Queries -->
            <div>
              <h3 class="text-md font-medium text-gray-900 mb-4">کوئری‌های کند</h3>
              <div class="space-y-3">
                <div
v-for="query in slowQueries" :key="query.id" 
                     class="p-3 border border-gray-200 rounded-lg">
                  <div class="flex items-start justify-between mb-2">
                    <span class="text-sm font-medium text-gray-900">{{ query.table }}</span>
                    <span class="text-sm text-red-600">{{ query.executionTime }}ms</span>
                  </div>
                  <p class="text-xs text-gray-600 font-mono mb-2">{{ query.sql }}</p>
                  <div class="flex items-center justify-between">
                    <span class="text-xs text-gray-500">تعداد اجرا: {{ query.executionCount }}</span>
                    <button
class="text-xs text-blue-600 hover:text-blue-800" 
                            @click="optimizeQuery(query.id)">بهینه‌سازی</button>
                  </div>
                </div>
              </div>
            </div>

            <!-- Index Analysis -->
            <div>
              <h3 class="text-md font-medium text-gray-900 mb-4">تحلیل ایندکس‌ها</h3>
              <div class="space-y-3">
                <div
v-for="index in indexAnalysis" :key="index.id" 
                     class="p-3 border border-gray-200 rounded-lg">
                  <div class="flex items-start justify-between mb-2">
                    <div>
                      <span class="text-sm font-medium text-gray-900">{{ index.table }}</span>
                      <span class="text-xs text-gray-500 mr-2">.</span>
                      <span class="text-sm text-gray-700">{{ index.column }}</span>
                    </div>
                    <span class="text-sm" :class="index.status === 'missing' ? 'text-red-600' : 'text-green-600'">
                      {{ index.status === 'missing' ? 'مفقود' : 'موجود' }}
                    </span>
                  </div>
                  <p class="text-xs text-gray-600 mb-2">{{ index.description }}</p>
                  <div class="flex items-center justify-between">
                    <span class="text-xs text-gray-500">تأثیر: {{ index.impact }}</span>
                    <button
v-if="index.status === 'missing'" class="text-xs text-green-600 hover:text-green-800" 
                            @click="createIndex(index.id)">ایجاد</button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Performance Report -->
      <div class="bg-white rounded-lg shadow">
        <div class="px-6 py-4 border-b border-gray-200 flex justify-between items-center">
          <h2 class="text-lg font-semibold text-gray-900">گزارش عملکرد</h2>
          <div class="flex space-x-2 space-x-reverse">
            <button
class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors" 
                    @click="generateReport">
              تولید گزارش
            </button>
            <button
class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors" 
                    @click="exportReport">
              صادر کردن
            </button>
          </div>
        </div>
        <div class="p-6">
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
            <!-- Performance Metrics -->
            <div>
              <h3 class="text-md font-medium text-gray-900 mb-4">معیارهای عملکرد</h3>
              <div class="space-y-4">
                <div v-for="metric in performanceMetrics" :key="metric.name" class="flex justify-between items-center">
                  <span class="text-sm text-gray-600">{{ metric.name }}</span>
                  <div class="flex items-center">
                    <div class="w-32 bg-gray-200 rounded-full h-2 mr-3">
                      <div
class="h-2 rounded-full" 
                           :class="getMetricColorClass(metric.value)"
                           :style="{ width: `${metric.value}%` }"></div>
                    </div>
                    <span class="text-sm font-medium text-gray-900">{{ metric.value }}%</span>
                  </div>
                </div>
              </div>
            </div>

            <!-- Recommendations -->
            <div>
              <h3 class="text-md font-medium text-gray-900 mb-4">توصیه‌های کلی</h3>
              <div class="space-y-3">
                <div
v-for="recommendation in recommendations" :key="recommendation.id" 
                     class="flex items-start">
                  <div class="flex-shrink-0 w-2 h-2 bg-blue-500 rounded-full mt-2 mr-3"></div>
                  <div>
                    <p class="text-sm text-gray-900">{{ recommendation.title }}</p>
                    <p class="text-xs text-gray-600">{{ recommendation.description }}</p>
                  </div>
                </div>
              </div>
            </div>
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
// const { user, hasPermission } = useAuth()

// Reactive data
const performanceStats = ref({
  responseTime: 245,
  requestsPerSecond: 1250,
  loadTime: 1.8,
  score: 87
})

const responseTimeHistory = ref([120, 180, 145, 220, 190, 160, 200, 175, 210, 165, 185, 195])
const cpuUsageHistory = ref([45, 52, 38, 67, 43, 58, 49, 61, 44, 53, 47, 55])

const bottlenecks = ref([
  {
    id: 1,
    name: 'کوئری‌های کند دیتابیس',
    description: 'کوئری‌های بدون ایندکس مناسب',
    severity: 'high',
    impact: 'تأثیر بالا'
  },
  {
    id: 2,
    name: 'فایل‌های بزرگ',
    description: 'فایل‌های CSS و JS فشرده نشده',
    severity: 'medium',
    impact: 'تأثیر متوسط'
  },
  {
    id: 3,
    name: 'کش ناکافی',
    description: 'کش Redis پر شده',
    severity: 'low',
    impact: 'تأثیر کم'
  }
])

const optimizationSuggestions = ref([
  {
    id: 1,
    title: 'فشرده‌سازی فایل‌ها',
    description: 'فشرده‌سازی CSS و JS برای کاهش حجم',
    impact: '20% بهبود',
    complexity: 'کم'
  },
  {
    id: 2,
    title: 'بهینه‌سازی کوئری‌ها',
    description: 'اضافه کردن ایندکس‌های مناسب',
    impact: '35% بهبود',
    complexity: 'متوسط'
  },
  {
    id: 3,
    title: 'افزایش کش',
    description: 'افزایش ظرفیت کش Redis',
    impact: '15% بهبود',
    complexity: 'کم'
  }
])

const codePerformance = ref([
  {
    name: 'main.js',
    performance: 85,
    size: 1024 * 1024 * 2.5, // 2.5MB
    executionTime: 45
  },
  {
    name: 'app.css',
    performance: 92,
    size: 1024 * 512, // 512KB
    executionTime: 12
  },
  {
    name: 'api.js',
    performance: 78,
    size: 1024 * 1024 * 1.8, // 1.8MB
    executionTime: 67
  },
  {
    name: 'utils.js',
    performance: 95,
    size: 1024 * 256, // 256KB
    executionTime: 8
  }
])

const slowQueries = ref([
  {
    id: 1,
    table: 'users',
    sql: 'SELECT * FROM users WHERE email = ? AND status = ?',
    executionTime: 450,
    executionCount: 1250
  },
  {
    id: 2,
    table: 'orders',
    sql: 'SELECT o.*, u.name FROM orders o JOIN users u ON o.user_id = u.id WHERE o.created_at > ?',
    executionTime: 320,
    executionCount: 890
  }
])

const indexAnalysis = ref([
  {
    id: 1,
    table: 'users',
    column: 'email',
    status: 'missing',
    description: 'ایندکس برای جستجوی ایمیل',
    impact: 'بالا'
  },
  {
    id: 2,
    table: 'orders',
    column: 'user_id',
    status: 'exists',
    description: 'ایندکس برای ارتباط با کاربران',
    impact: 'متوسط'
  },
  {
    id: 3,
    table: 'products',
    column: 'category_id',
    status: 'missing',
    description: 'ایندکس برای فیلتر محصولات',
    impact: 'بالا'
  }
])

const performanceMetrics = ref([
  { name: 'سرعت بارگذاری', value: 85 },
  { name: 'عملکرد سرور', value: 92 },
  { name: 'عملکرد دیتابیس', value: 78 },
  { name: 'بهینه‌سازی کد', value: 88 }
])

const recommendations = ref([
  {
    id: 1,
    title: 'استفاده از CDN',
    description: 'استفاده از CDN برای فایل‌های استاتیک'
  },
  {
    id: 2,
    title: 'بهینه‌سازی تصاویر',
    description: 'فشرده‌سازی و استفاده از فرمت‌های مدرن'
  },
  {
    id: 3,
    title: 'کش‌گذاری هوشمند',
    description: 'پیاده‌سازی استراتژی‌های کش مناسب'
  }
])

// Methods
const getBottleneckColor = (severity) => {
  const colors = {
    high: 'bg-red-500',
    medium: 'bg-yellow-500',
    low: 'bg-blue-500'
  }
  return colors[severity] || 'bg-gray-500'
}

const getBottleneckTextColor = (severity) => {
  const colors = {
    high: 'text-red-600',
    medium: 'text-yellow-600',
    low: 'text-blue-600'
  }
  return colors[severity] || 'text-gray-600'
}

const getPerformanceColor = (performance) => {
  if (performance >= 90) return 'bg-green-600'
  if (performance >= 70) return 'bg-yellow-600'
  return 'bg-red-600'
}

const getMetricColorClass = (value) => {
  if (value >= 90) return 'bg-green-600'
  if (value >= 70) return 'bg-yellow-600'
  return 'bg-red-600'
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

const applyOptimization = (_id) => {

}

const optimizeFile = (_fileName) => {

}

const optimizeQuery = (_id) => {

}

const createIndex = (_id) => {

}

const generateReport = () => {

}

const exportReport = () => {

}

// Auto-refresh performance stats
onMounted(() => {
  const interval = setInterval(() => {
    performanceStats.value.responseTime = Math.floor(Math.random() * 100) + 200
    performanceStats.value.requestsPerSecond = Math.floor(Math.random() * 500) + 1000
    performanceStats.value.loadTime = (Math.random() * 1 + 1).toFixed(1)
    performanceStats.value.score = Math.floor(Math.random() * 20) + 80
  }, 10000)

  onUnmounted(() => {
    clearInterval(interval)
  })
})
</script> 
