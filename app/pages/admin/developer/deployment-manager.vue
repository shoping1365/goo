<template>
  <div class="min-h-screen bg-gray-50 p-6">
    <div class="max-w-7xl mx-auto">
      <!-- Header -->
      <div class="mb-8">
        <h1 class="text-3xl font-bold text-gray-900 mb-2">مدیر استقرار</h1>
        <p class="text-gray-600">مدیریت و نظارت بر فرآیند استقرار نرم‌افزار</p>
      </div>

      <!-- Deployment Overview -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-green-100 rounded-lg">
              <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
            <div class="mr-4">
              <p class="text-sm font-medium text-gray-600">موفق</p>
              <p class="text-2xl font-bold text-gray-900">{{ deploymentStats.successful }}</p>
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
            <div class="mr-4">
              <p class="text-sm font-medium text-gray-600">ناموفق</p>
              <p class="text-2xl font-bold text-gray-900">{{ deploymentStats.failed }}</p>
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
              <p class="text-2xl font-bold text-gray-900">{{ deploymentStats.running }}</p>
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
            <div class="mr-4">
              <p class="text-sm font-medium text-gray-600">متوسط زمان</p>
              <p class="text-2xl font-bold text-gray-900">{{ deploymentStats.avgTime }}min</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Quick Deploy -->
      <div class="bg-white rounded-lg shadow mb-8">
        <div class="px-6 py-4 border-b border-gray-200">
          <h2 class="text-lg font-semibold text-gray-900">استقرار سریع</h2>
        </div>
        <div class="p-6">
          <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
            <button
:disabled="isDeploying" class="flex items-center justify-center px-4 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed transition-colors" 
                    @click="startDeployment('staging')">
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-8l-4-4m0 0L8 8m4-4v12"></path>
              </svg>
              استقرار در Staging
            </button>

            <button
:disabled="isDeploying" class="flex items-center justify-center px-4 py-3 bg-green-600 text-white rounded-lg hover:bg-green-700 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
                    @click="startDeployment('production')">
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
              </svg>
              استقرار در Production
            </button>

            <button
:disabled="isDeploying" class="flex items-center justify-center px-4 py-3 bg-red-600 text-white rounded-lg hover:bg-red-700 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
                    @click="rollbackDeployment()">
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h10a8 8 0 018 8v2M3 10l6 6m-6-6l6-6"></path>
              </svg>
              بازگشت به نسخه قبلی
            </button>
          </div>
        </div>
      </div>

      <!-- Deployment Configuration -->
      <div class="bg-white rounded-lg shadow mb-8">
        <div class="px-6 py-4 border-b border-gray-200">
          <h2 class="text-lg font-semibold text-gray-900">تنظیمات استقرار</h2>
        </div>
        <div class="p-6">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <!-- Environment Settings -->
            <div>
              <h3 class="text-md font-medium text-gray-900 mb-4">تنظیمات محیط</h3>
              <div class="space-y-4">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">محیط هدف</label>
                  <select v-model="deploymentConfig.environment" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
                    <option value="staging">Staging</option>
                    <option value="production">Production</option>
                    <option value="testing">Testing</option>
                  </select>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">استراتژی استقرار</label>
                  <select v-model="deploymentConfig.strategy" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
                    <option value="blue-green">Blue-Green</option>
                    <option value="rolling">Rolling Update</option>
                    <option value="canary">Canary</option>
                    <option value="recreate">Recreate</option>
                  </select>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">تعداد Replica</label>
                  <input
v-model="deploymentConfig.replicas" type="number" min="1" max="10" 
                         class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
                </div>
              </div>
            </div>

            <!-- Build Settings -->
            <div>
              <h3 class="text-md font-medium text-gray-900 mb-4">تنظیمات Build</h3>
              <div class="space-y-4">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">Branch</label>
                  <select v-model="deploymentConfig.branch" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
                    <option value="main">main</option>
                    <option value="develop">develop</option>
                    <option value="feature/new-feature">feature/new-feature</option>
                  </select>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">Commit Hash</label>
                  <input
v-model="deploymentConfig.commitHash" type="text" 
                         class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                         placeholder="a1b2c3d4...">
                </div>
                <div class="flex items-center">
                  <input
id="auto-rollback" v-model="deploymentConfig.autoRollback" type="checkbox" 
                         class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded">
                  <label for="auto-rollback" class="mr-2 text-sm text-gray-700">بازگشت خودکار در صورت خطا</label>
                </div>
              </div>
            </div>
          </div>

          <div class="mt-6 flex justify-end">
            <button
class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors" 
                    @click="saveDeploymentConfiguration">
              ذخیره تنظیمات
            </button>
          </div>
        </div>
      </div>

      <!-- Deployment History -->
      <div class="bg-white rounded-lg shadow mb-8">
        <div class="px-6 py-4 border-b border-gray-200 flex justify-between items-center">
          <h2 class="text-lg font-semibold text-gray-900">تاریخچه استقرار</h2>
          <div class="flex space-x-2 space-x-reverse">
            <select v-model="filter.status" class="px-3 py-1 border border-gray-300 rounded-md text-sm">
              <option value="">همه وضعیت‌ها</option>
              <option value="successful">موفق</option>
              <option value="failed">ناموفق</option>
              <option value="running">در حال اجرا</option>
            </select>
            <select v-model="filter.environment" class="px-3 py-1 border border-gray-300 rounded-md text-sm">
              <option value="">همه محیط‌ها</option>
              <option value="staging">Staging</option>
              <option value="production">Production</option>
              <option value="testing">Testing</option>
            </select>
          </div>
        </div>
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">محیط</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نسخه</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">زمان</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مدت</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="deployment in filteredDeployments" :key="deployment.id" class="hover:bg-gray-50">
                <td class="px-6 py-4 whitespace-nowrap">
                  <span
class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                        :class="getEnvironmentBadgeClass(deployment.environment)">
                    {{ getEnvironmentLabel(deployment.environment) }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <div>
                    <p class="text-sm font-medium text-gray-900">{{ deployment.version }}</p>
                    <p class="text-xs text-gray-500">{{ deployment.commitHash }}</p>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span
class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                        :class="getStatusBadgeClass(deployment.status)">
                    {{ getStatusLabel(deployment.status) }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  {{ formatDate(deployment.timestamp) }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  {{ deployment.duration }}min
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                  <div class="flex space-x-2 space-x-reverse">
                    <button
class="text-blue-600 hover:text-blue-900" 
                            @click="viewDeploymentLogs(deployment.id)">لاگ‌ها</button>
                    <button
v-if="deployment.status === 'successful'" class="text-red-600 hover:text-red-900" 
                            @click="rollbackToVersion(deployment.id)">بازگشت</button>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Environment Status -->
      <div class="bg-white rounded-lg shadow mb-8">
        <div class="px-6 py-4 border-b border-gray-200">
          <h2 class="text-lg font-semibold text-gray-900">وضعیت محیط‌ها</h2>
        </div>
        <div class="p-6">
          <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
            <div
v-for="env in environments" :key="env.name" 
                 class="border border-gray-200 rounded-lg p-6">
              <div class="flex items-center justify-between mb-4">
                <h3 class="text-lg font-medium text-gray-900">{{ env.name }}</h3>
                <span
class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                      :class="getStatusBadgeClass(env.status)">
                  {{ getStatusLabel(env.status) }}
                </span>
              </div>
              
              <div class="space-y-3">
                <div class="flex justify-between text-sm">
                  <span class="text-gray-600">نسخه فعلی:</span>
                  <span class="font-medium">{{ env.currentVersion }}</span>
                </div>
                <div class="flex justify-between text-sm">
                  <span class="text-gray-600">تعداد Pod:</span>
                  <span class="font-medium">{{ env.podCount }}</span>
                </div>
                <div class="flex justify-between text-sm">
                  <span class="text-gray-600">CPU:</span>
                  <span class="font-medium">{{ env.cpuUsage }}%</span>
                </div>
                <div class="flex justify-between text-sm">
                  <span class="text-gray-600">RAM:</span>
                  <span class="font-medium">{{ env.memoryUsage }}%</span>
                </div>
              </div>

              <div class="mt-4 flex space-x-2 space-x-reverse">
                <button
class="flex-1 px-3 py-2 bg-green-600 text-white text-sm rounded hover:bg-green-700 transition-colors" 
                        @click="scaleEnvironment(env.name, 'up')">
                  افزایش
                </button>
                <button
class="flex-1 px-3 py-2 bg-yellow-600 text-white text-sm rounded hover:bg-yellow-700 transition-colors" 
                        @click="scaleEnvironment(env.name, 'down')">
                  کاهش
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Deployment Pipeline -->
      <div class="bg-white rounded-lg shadow">
        <div class="px-6 py-4 border-b border-gray-200">
          <h2 class="text-lg font-semibold text-gray-900">خط لوله استقرار</h2>
        </div>
        <div class="p-6">
          <div class="space-y-6">
            <div
v-for="stage in deploymentPipeline" :key="stage.id" 
                 class="flex items-center">
              <div
class="flex-shrink-0 w-8 h-8 rounded-full flex items-center justify-center mr-4"
                   :class="getStageStatusClass(stage.status)">
                <svg v-if="stage.status === 'completed'" class="w-5 h-5 text-white" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd"></path>
                </svg>
                <svg v-else-if="stage.status === 'running'" class="w-5 h-5 text-white animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
                </svg>
                <span v-else class="text-white text-sm font-medium">{{ stage.order }}</span>
              </div>
              
              <div class="flex-1">
                <h3 class="text-sm font-medium text-gray-900">{{ stage.name }}</h3>
                <p class="text-xs text-gray-600">{{ stage.description }}</p>
              </div>
              
              <div class="text-right">
                <p class="text-sm text-gray-900">{{ stage.duration }}s</p>
                <p class="text-xs text-gray-500">{{ formatDate(stage.timestamp) }}</p>
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
const { user, hasPermission } = useAuth()

// Reactive data
const deploymentStats = ref({
  successful: 24,
  failed: 3,
  running: 1,
  avgTime: 8.5
})

const isDeploying = ref(false)

const deploymentConfig = ref({
  environment: 'staging',
  strategy: 'blue-green',
  replicas: 3,
  branch: 'main',
  commitHash: 'a1b2c3d4e5f6',
  autoRollback: true
})

const filter = ref({
  status: '',
  environment: ''
})

const deployments = ref([
  {
    id: 1,
    environment: 'production',
    version: 'v1.2.3',
    commitHash: 'a1b2c3d4e5f6',
    status: 'successful',
    timestamp: new Date('2024-01-15T10:30:00'),
    duration: 12
  },
  {
    id: 2,
    environment: 'staging',
    version: 'v1.2.4',
    commitHash: 'b2c3d4e5f6g7',
    status: 'running',
    timestamp: new Date('2024-01-15T11:15:00'),
    duration: 8
  },
  {
    id: 3,
    environment: 'production',
    version: 'v1.2.2',
    commitHash: 'c3d4e5f6g7h8',
    status: 'failed',
    timestamp: new Date('2024-01-14T16:45:00'),
    duration: 5
  }
])

const environments = ref([
  {
    name: 'Production',
    status: 'healthy',
    currentVersion: 'v1.2.3',
    podCount: 5,
    cpuUsage: 65,
    memoryUsage: 72
  },
  {
    name: 'Staging',
    status: 'deploying',
    currentVersion: 'v1.2.4',
    podCount: 3,
    cpuUsage: 45,
    memoryUsage: 58
  },
  {
    name: 'Testing',
    status: 'healthy',
    currentVersion: 'v1.2.4',
    podCount: 2,
    cpuUsage: 30,
    memoryUsage: 40
  }
])

const deploymentPipeline = ref([
  {
    id: 1,
    name: 'Build',
    description: 'ساخت کد و ایجاد Docker Image',
    status: 'completed',
    order: 1,
    duration: 120,
    timestamp: new Date('2024-01-15T11:00:00')
  },
  {
    id: 2,
    name: 'Test',
    description: 'اجرای تست‌های خودکار',
    status: 'completed',
    order: 2,
    duration: 45,
    timestamp: new Date('2024-01-15T11:02:00')
  },
  {
    id: 3,
    name: 'Deploy',
    description: 'استقرار در محیط هدف',
    status: 'running',
    order: 3,
    duration: 30,
    timestamp: new Date('2024-01-15T11:03:00')
  },
  {
    id: 4,
    name: 'Health Check',
    description: 'بررسی سلامت سرویس',
    status: 'pending',
    order: 4,
    duration: 15,
    timestamp: null
  }
])

// Computed
const filteredDeployments = computed(() => {
  return deployments.value.filter(deployment => {
    if (filter.value.status && deployment.status !== filter.value.status) return false
    if (filter.value.environment && deployment.environment !== filter.value.environment) return false
    return true
  })
})

// Methods
const startDeployment = async (environment) => {
  isDeploying.value = true
  try {
    // Simulate deployment
    await new Promise(resolve => setTimeout(resolve, 5000))
    
    const newDeployment = {
      id: Date.now(),
      environment,
      version: `v1.2.${Math.floor(Math.random() * 10)}`,
      commitHash: Math.random().toString(36).substring(2, 8),
      status: 'successful',
      timestamp: new Date(),
      duration: Math.floor(Math.random() * 20) + 5
    }
    
    deployments.value.unshift(newDeployment)
    deploymentStats.value.successful++
    
  } catch (error) {
    console.error('Deployment error:', error)
  } finally {
    isDeploying.value = false
  }
}

const rollbackDeployment = async () => {
  isDeploying.value = true
  try {
    await new Promise(resolve => setTimeout(resolve, 3000))
    console.log('Rollback completed')
  } catch (error) {
    console.error('Rollback error:', error)
  } finally {
    isDeploying.value = false
  }
}

const saveDeploymentConfiguration = async () => {
  try {
    await new Promise(resolve => setTimeout(resolve, 1000))
    console.log('Configuration saved:', deploymentConfig.value)
  } catch (error) {
    console.error('Error saving configuration:', error)
  }
}

const viewDeploymentLogs = (id) => {
  console.log('Viewing logs for deployment:', id)
}

const rollbackToVersion = (id) => {
  console.log('Rolling back to version:', id)
}

const scaleEnvironment = (name, direction) => {
  console.log(`Scaling ${name} ${direction}`)
}

const getEnvironmentBadgeClass = (environment) => {
  const classes = {
    production: 'bg-red-100 text-red-800',
    staging: 'bg-yellow-100 text-yellow-800',
    testing: 'bg-blue-100 text-blue-800'
  }
  return classes[environment] || 'bg-gray-100 text-gray-800'
}

const getEnvironmentLabel = (environment) => {
  const labels = {
    production: 'Production',
    staging: 'Staging',
    testing: 'Testing'
  }
  return labels[environment] || environment
}

const getStatusBadgeClass = (status) => {
  const classes = {
    successful: 'bg-green-100 text-green-800',
    failed: 'bg-red-100 text-red-800',
    running: 'bg-yellow-100 text-yellow-800',
    healthy: 'bg-green-100 text-green-800',
    deploying: 'bg-blue-100 text-blue-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getStatusLabel = (status) => {
  const labels = {
    successful: 'موفق',
    failed: 'ناموفق',
    running: 'در حال اجرا',
    healthy: 'سالم',
    deploying: 'در حال استقرار'
  }
  return labels[status] || status
}

const getStageStatusClass = (status) => {
  const classes = {
    completed: 'bg-green-500',
    running: 'bg-blue-500',
    failed: 'bg-red-500',
    pending: 'bg-gray-300'
  }
  return classes[status] || 'bg-gray-300'
}

const formatDate = (date) => {
  if (!date) return '-'
  return new Intl.DateTimeFormat('fa-IR', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  }).format(date)
}
</script> 
