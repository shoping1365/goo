<template>
  <div class="min-h-screen bg-gray-50 p-6">
    <div class="max-w-7xl mx-auto">
      <!-- Header -->
      <div class="mb-8">
        <h1 class="text-3xl font-bold text-gray-900 mb-2">ساخت خط لوله CI/CD</h1>
        <p class="text-gray-600">ابزارهای ساخت و مدیریت خط لوله‌های CI/CD</p>
      </div>

      <!-- Stats Cards -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-blue-100 rounded-lg">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">خط لوله‌های فعال</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.activePipelines }}</p>
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
              <p class="text-2xl font-semibold text-gray-900">{{ stats.successfulRuns }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-red-100 rounded-lg">
              <svg class="w-6 h-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">ناموفق</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.failedRuns }}</p>
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
              <p class="text-sm font-medium text-gray-600">زمان متوسط</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.avgDuration }}m</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Main Content Grid -->
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <!-- Pipeline Builder -->
        <div class="bg-white rounded-lg shadow">
          <div class="p-6 border-b border-gray-200">
            <h2 class="text-xl font-semibold text-gray-900">ساخت خط لوله</h2>
          </div>
          <div class="p-6">
            <div class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">نام خط لوله</label>
                <input 
                  v-model="pipelineName"
                  type="text"
                  class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
                  placeholder="نام خط لوله را وارد کنید..."
                >
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">نوع خط لوله</label>
                <select 
                  v-model="pipelineType"
                  class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                  <option value="github">GitHub Actions</option>
                  <option value="gitlab">GitLab CI</option>
                  <option value="jenkins">Jenkins</option>
                  <option value="azure">Azure DevOps</option>
                  <option value="custom">سفارشی</option>
                </select>
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">مراحل خط لوله</label>
                <div class="space-y-2">
                  <div v-for="(stage, index) in pipelineStages" :key="index" class="flex items-center space-x-2 space-x-reverse">
                    <input 
                      v-model="stage.name"
                      type="text"
                      class="flex-1 border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
                      placeholder="نام مرحله..."
                    >
                    <select 
                      v-model="stage.type"
                      class="border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
                    >
                      <option value="build">ساخت</option>
                      <option value="test">تست</option>
                      <option value="deploy">استقرار</option>
                      <option value="custom">سفارشی</option>
                    </select>
                    <button 
                      class="text-red-600 hover:text-red-800"
                      @click="removeStage(index)"
                    >
                      <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                      </svg>
                    </button>
                  </div>
                  <button 
                    class="w-full bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                    @click="addStage"
                  >
                    افزودن مرحله
                  </button>
                </div>
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">تنظیمات پیشرفته</label>
                <div class="space-y-2">
                  <label class="flex items-center">
                    <input 
                      v-model="pipelineConfig.autoTrigger"
                      type="checkbox"
                      class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                    >
                    <span class="mr-2 text-sm text-gray-700">اجرای خودکار</span>
                  </label>
                  <label class="flex items-center">
                    <input 
                      v-model="pipelineConfig.parallelExecution"
                      type="checkbox"
                      class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                    >
                    <span class="mr-2 text-sm text-gray-700">اجرای موازی</span>
                  </label>
                  <label class="flex items-center">
                    <input 
                      v-model="pipelineConfig.notifications"
                      type="checkbox"
                      class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                    >
                    <span class="mr-2 text-sm text-gray-700">اعلان‌ها</span>
                  </label>
                </div>
              </div>

              <button 
                :disabled="!pipelineName"
                class="w-full bg-green-600 hover:bg-green-700 disabled:bg-gray-400 text-white font-medium py-3 px-4 rounded-lg transition-colors"
                @click="createPipeline"
              >
                ایجاد خط لوله
              </button>
            </div>
          </div>
        </div>

        <!-- Pipeline Templates -->
        <div class="bg-white rounded-lg shadow">
          <div class="p-6 border-b border-gray-200">
            <h2 class="text-xl font-semibold text-gray-900">قالب‌های آماده</h2>
          </div>
          <div class="p-6">
            <div class="space-y-4">
              <div v-for="template in pipelineTemplates" :key="template.id" class="border rounded-lg p-6 hover:bg-gray-50">
                <div class="flex items-center justify-between mb-2">
                  <h3 class="font-medium text-gray-900">{{ template.name }}</h3>
                  <span class="text-xs bg-gray-100 text-gray-600 px-2 py-1 rounded">{{ template.type }}</span>
                </div>
                <p class="text-sm text-gray-500 mb-3">{{ template.description }}</p>
                <div class="flex items-center space-x-2 space-x-reverse">
                  <button 
                    class="text-blue-600 hover:text-blue-800 text-sm font-medium"
                    @click="useTemplate(template)"
                  >
                    استفاده
                  </button>
                  <button 
                    class="text-gray-600 hover:text-gray-800 text-sm font-medium"
                    @click="previewTemplate(template)"
                  >
                    پیش‌نمایش
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Pipeline Status -->
        <div class="bg-white rounded-lg shadow">
          <div class="p-6 border-b border-gray-200">
            <h2 class="text-xl font-semibold text-gray-900">وضعیت خط لوله‌ها</h2>
          </div>
          <div class="p-6">
            <div class="space-y-4">
              <div v-for="pipeline in activePipelines" :key="pipeline.id" class="border rounded-lg p-6">
                <div class="flex items-center justify-between mb-2">
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <span class="font-medium text-gray-900">{{ pipeline.name }}</span>
                    <span
:class="[
                      'px-2 py-1 rounded text-xs font-medium',
                      pipeline.status === 'running' ? 'bg-blue-100 text-blue-800' :
                      pipeline.status === 'success' ? 'bg-green-100 text-green-800' :
                      pipeline.status === 'failed' ? 'bg-red-100 text-red-800' :
                      'bg-yellow-100 text-yellow-800'
                    ]">
                      {{ pipeline.status }}
                    </span>
                  </div>
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <button 
                      v-if="pipeline.status !== 'running'"
                      class="text-green-600 hover:text-green-800 text-sm font-medium"
                      @click="runPipeline(pipeline)"
                    >
                      اجرا
                    </button>
                    <button 
                      v-if="pipeline.status === 'running'"
                      class="text-red-600 hover:text-red-800 text-sm font-medium"
                      @click="stopPipeline(pipeline)"
                    >
                      توقف
                    </button>
                    <button 
                      class="text-blue-600 hover:text-blue-800 text-sm font-medium"
                      @click="viewPipelineDetails(pipeline)"
                    >
                      جزئیات
                    </button>
                  </div>
                </div>
                <div class="text-sm text-gray-500 space-y-1">
                  <p>آخرین اجرا: {{ pipeline.lastRun }}</p>
                  <p>مدت زمان: {{ pipeline.duration }}</p>
                  <p>شاخه: {{ pipeline.branch }}</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Pipeline Execution History -->
      <div class="mt-8 bg-white rounded-lg shadow">
        <div class="p-6 border-b border-gray-200">
          <h2 class="text-xl font-semibold text-gray-900">تاریخچه اجراها</h2>
        </div>
        <div class="p-6">
          <div class="overflow-x-auto">
            <table class="min-w-full divide-y divide-gray-200">
              <thead class="bg-gray-50">
                <tr>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">خط لوله</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">شاخه</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مدت زمان</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
                </tr>
              </thead>
              <tbody class="bg-white divide-y divide-gray-200">
                <tr v-for="execution in executionHistory" :key="execution.id">
                  <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{{ execution.pipeline }}</td>
                  <td class="px-6 py-4 whitespace-nowrap">
                    <span
:class="[
                      'px-2 py-1 rounded text-xs font-medium',
                      execution.status === 'success' ? 'bg-green-100 text-green-800' :
                      execution.status === 'failed' ? 'bg-red-100 text-red-800' :
                      'bg-yellow-100 text-yellow-800'
                    ]">
                      {{ execution.status }}
                    </span>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ execution.branch }}</td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ execution.date }}</td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ execution.duration }}</td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                    <button 
                      class="text-blue-600 hover:text-blue-800"
                      @click="viewExecutionLogs(execution)"
                    >
                      لاگ‌ها
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>

      <!-- Pipeline Configuration Modal -->
      <div v-if="configModal.show" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
        <div class="bg-white rounded-lg max-w-4xl max-h-4xl overflow-auto">
          <div class="p-6 border-b border-gray-200">
            <div class="flex items-center justify-between">
              <h3 class="text-lg font-semibold text-gray-900">پیکربندی خط لوله</h3>
              <button class="text-gray-400 hover:text-gray-600" @click="configModal.show = false">
                <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                </svg>
              </button>
            </div>
          </div>
          <div class="p-6">
            <pre class="bg-gray-900 text-gray-100 p-6 rounded-lg overflow-x-auto text-sm">{{ configModal.config }}</pre>
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

// استفاده از useAuth برای چک کردن پرمیژن‌ها
// const { user, hasPermission } = useAuth()

// Reactive data
const pipelineName = ref('')
const pipelineType = ref('github')
const pipelineStages = ref([
  { name: 'Build', type: 'build' },
  { name: 'Test', type: 'test' },
  { name: 'Deploy', type: 'deploy' }
])

const pipelineConfig = reactive({
  autoTrigger: true,
  parallelExecution: false,
  notifications: true
})

const configModal = reactive({
  show: false,
  config: ''
})

const stats = reactive({
  activePipelines: 5,
  successfulRuns: 124,
  failedRuns: 8,
  avgDuration: 12
})

const pipelineTemplates = ref([
  {
    id: 1,
    name: 'Node.js Application',
    type: 'GitHub Actions',
    description: 'قالب کامل برای برنامه‌های Node.js شامل تست و استقرار'
  },
  {
    id: 2,
    name: 'React Frontend',
    type: 'GitLab CI',
    description: 'قالب برای برنامه‌های React با تست و ساخت'
  },
  {
    id: 3,
    name: 'Go Backend',
    type: 'Jenkins',
    description: 'قالب برای سرویس‌های Go با تست و استقرار'
  },
  {
    id: 4,
    name: 'Docker Multi-stage',
    type: 'Azure DevOps',
    description: 'قالب برای ساخت و استقرار کانتینرها'
  }
])

const activePipelines = ref([
  {
    id: 1,
    name: 'Frontend CI/CD',
    status: 'running',
    lastRun: '2 minutes ago',
    duration: '8m 32s',
    branch: 'main'
  },
  {
    id: 2,
    name: 'Backend Pipeline',
    status: 'success',
    lastRun: '1 hour ago',
    duration: '5m 15s',
    branch: 'develop'
  },
  {
    id: 3,
    name: 'Database Migration',
    status: 'failed',
    lastRun: '3 hours ago',
    duration: '2m 45s',
    branch: 'feature/db-update'
  }
])

const executionHistory = ref([
  {
    id: 1,
    pipeline: 'Frontend CI/CD',
    status: 'success',
    branch: 'main',
    date: '2024-01-15 10:30',
    duration: '8m 32s'
  },
  {
    id: 2,
    pipeline: 'Backend Pipeline',
    status: 'success',
    branch: 'develop',
    date: '2024-01-15 09:15',
    duration: '5m 15s'
  },
  {
    id: 3,
    pipeline: 'Database Migration',
    status: 'failed',
    branch: 'feature/db-update',
    date: '2024-01-15 07:45',
    duration: '2m 45s'
  }
])

// Methods
function addStage() {
  pipelineStages.value.push({ name: '', type: 'build' })
}

function removeStage(index) {
  pipelineStages.value.splice(index, 1)
}

function createPipeline() {
  if (pipelineName.value.trim()) {
    alert(`خط لوله ${pipelineName.value} ایجاد شد`)
    pipelineName.value = ''
    pipelineStages.value = [
      { name: 'Build', type: 'build' },
      { name: 'Test', type: 'test' },
      { name: 'Deploy', type: 'deploy' }
    ]
  }
}

function useTemplate(template) {
  alert(`قالب ${template.name} اعمال شد`)
}

function previewTemplate(template) {
  configModal.config = `name: ${template.name}
on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v2
    - name: Use Node.js
      uses: actions/setup-node@v2
      with:
        node-version: '16'
    - run: npm ci
    - run: npm test
    - run: npm run build`
  configModal.show = true
}

function runPipeline(pipeline) {
  alert(`خط لوله ${pipeline.name} اجرا شد`)
  pipeline.status = 'running'
}

function stopPipeline(pipeline) {
  alert(`خط لوله ${pipeline.name} متوقف شد`)
  pipeline.status = 'stopped'
}

function viewPipelineDetails(pipeline) {
  alert(`جزئیات خط لوله ${pipeline.name} نمایش داده می‌شود`)
}

function viewExecutionLogs(execution) {
  alert(`لاگ‌های اجرای ${execution.pipeline} نمایش داده می‌شود`)
}
</script> 
