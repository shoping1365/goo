<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="bg-white rounded-lg shadow px-4 py-4">
      <div class="flex items-center justify-between">
        <div>
          <h2 class="text-xl font-bold text-gray-900">سیستم راه‌اندازی و تست</h2>
          <p class="mt-1 text-sm text-gray-500">راه‌اندازی و تست سناریوهای مختلف امتیازدهی</p>
        </div>
        <div class="flex space-x-3 space-x-reverse">
          <button class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition-colors" @click="runFullSetup">
            راه‌اندازی کامل
          </button>
          <button class="bg-green-600 text-white px-4 py-2 rounded-md hover:bg-green-700 transition-colors" @click="runTests">
            اجرای تست‌ها
          </button>
        </div>
      </div>
    </div>

    <!-- Setup Wizard -->
    <div class="bg-white rounded-lg shadow px-4 py-4">
      <h3 class="text-lg font-medium text-gray-900 mb-4">مراحل راه‌اندازی</h3>
      
      <div class="space-y-4">
        <div v-for="(step, index) in setupSteps" :key="step.id" class="flex items-center px-4 py-4 border border-gray-200 rounded-lg" :class="getStepClass(step.status)">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 rounded-full flex items-center justify-center" :class="getStepIconClass(step.status)">
              <svg v-if="step.status === 'completed'" class="w-5 h-5 text-white" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd"></path>
              </svg>
              <svg v-else-if="step.status === 'running'" class="w-5 h-5 text-white animate-spin" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M4 2a1 1 0 011 1v2.101a7.002 7.002 0 0111.601 2.566 1 1 0 11-1.885.666A5.002 5.002 0 005.999 7H9a1 1 0 010 2H4a1 1 0 01-1-1V3a1 1 0 011-1zm.008 9.057a1 1 0 011.276.61A5.002 5.002 0 0014.001 13H11a1 1 0 110-2h5a1 1 0 011 1v5a1 1 0 11-2 0v-2.101a7.002 7.002 0 01-11.601-2.566 1 1 0 01.61-1.276z" clip-rule="evenodd"></path>
              </svg>
              <span v-else class="text-sm font-medium text-gray-500">{{ index + 1 }}</span>
            </div>
          </div>
          <div class="mr-4 flex-1">
            <h4 class="text-sm font-medium text-gray-900">{{ step.name }}</h4>
            <p class="text-sm text-gray-500">{{ step.description }}</p>
            <div v-if="step.status === 'running'" class="mt-2">
              <div class="w-full bg-gray-200 rounded-full h-2">
                <div :style="{ width: step.progress + '%' }" class="bg-blue-600 h-2 rounded-full transition-all duration-300"></div>
              </div>
              <span class="text-xs text-gray-500 mt-1">{{ step.progress }}% تکمیل شده</span>
            </div>
          </div>
          <div class="flex space-x-2 space-x-reverse">
            <button v-if="step.status === 'pending'" class="text-blue-600 hover:text-blue-900 text-sm" @click="runStep(step)">
              اجرا
            </button>
            <button v-if="step.status === 'completed'" class="text-gray-600 hover:text-gray-900 text-sm" @click="resetStep(step)">
              بازنشانی
            </button>
            <button class="text-purple-600 hover:text-purple-900 text-sm" @click="viewStepDetails(step)">
              جزئیات
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Testing Environment -->
    <div class="bg-white rounded-lg shadow px-4 py-4">
      <h3 class="text-lg font-medium text-gray-900 mb-4">محیط تست</h3>
      
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gapx-4 py-4">
        <div v-for="test in testEnvironments" :key="test.id" class="border border-gray-200 rounded-lg px-4 py-4">
          <div class="flex items-center justify-between mb-3">
            <h4 class="text-lg font-medium text-gray-900">{{ test.name }}</h4>
            <span :class="getTestStatusClass(test.status)" class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium">
              {{ getTestStatusText(test.status) }}
            </span>
          </div>
          <p class="text-sm text-gray-600 mb-4">{{ test.description }}</p>
          <div class="space-y-2 mb-4">
            <div class="flex justify-between text-sm">
              <span class="text-gray-500">کاربران تست:</span>
              <span class="text-gray-900">{{ test.testUsers }}</span>
            </div>
            <div class="flex justify-between text-sm">
              <span class="text-gray-500">مدت تست:</span>
              <span class="text-gray-900">{{ test.duration }}</span>
            </div>
            <div class="flex justify-between text-sm">
              <span class="text-gray-500">نرخ موفقیت:</span>
              <span class="text-gray-900">{{ test.successRate }}%</span>
            </div>
          </div>
          <div class="flex space-x-2 space-x-reverse">
            <button :disabled="test.status === 'running'" class="bg-green-600 text-white px-3 py-1 rounded-md text-sm hover:bg-green-700 disabled:opacity-50" @click="startTest(test)">
              شروع تست
            </button>
            <button :disabled="test.status !== 'running'" class="bg-red-600 text-white px-3 py-1 rounded-md text-sm hover:bg-red-700 disabled:opacity-50" @click="stopTest(test)">
              توقف
            </button>
            <button class="bg-blue-600 text-white px-3 py-1 rounded-md text-sm hover:bg-blue-700" @click="viewTestResults(test)">
              نتایج
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Performance Monitoring -->
    <div class="bg-white rounded-lg shadow px-4 py-4">
      <h3 class="text-lg font-medium text-gray-900 mb-4">نظارت بر عملکرد</h3>
      
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gapx-4 py-4 mb-6">
        <div class="bg-gray-50 rounded-lg px-4 py-4">
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <div class="w-8 h-8 bg-green-100 rounded-full flex items-center justify-center">
                <svg class="w-5 h-5 text-green-600" fill="currentColor" viewBox="0 0 20 20">
                  <path d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
              </div>
            </div>
            <div class="mr-4">
              <p class="text-sm font-medium text-gray-500">تست‌های موفق</p>
              <p class="text-2xl font-semibold text-gray-900">{{ monitoring.successfulTests }}</p>
            </div>
          </div>
        </div>

        <div class="bg-gray-50 rounded-lg px-4 py-4">
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <div class="w-8 h-8 bg-red-100 rounded-full flex items-center justify-center">
                <svg class="w-5 h-5 text-red-600" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clip-rule="evenodd"></path>
                </svg>
              </div>
            </div>
            <div class="mr-4">
              <p class="text-sm font-medium text-gray-500">تست‌های ناموفق</p>
              <p class="text-2xl font-semibold text-gray-900">{{ monitoring.failedTests }}</p>
            </div>
          </div>
        </div>

        <div class="bg-gray-50 rounded-lg px-4 py-4">
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <div class="w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center">
                <svg class="w-5 h-5 text-blue-600" fill="currentColor" viewBox="0 0 20 20">
                  <path d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-12a1 1 0 10-2 0v4a1 1 0 00.293.707l2.828 2.829a1 1 0 101.415-1.415L11 9.586V6z"></path>
                </svg>
              </div>
            </div>
            <div class="mr-4">
              <p class="text-sm font-medium text-gray-500">زمان متوسط</p>
              <p class="text-2xl font-semibold text-gray-900">{{ monitoring.avgTime }}s</p>
            </div>
          </div>
        </div>

        <div class="bg-gray-50 rounded-lg px-4 py-4">
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <div class="w-8 h-8 bg-yellow-100 rounded-full flex items-center justify-center">
                <svg class="w-5 h-5 text-yellow-600" fill="currentColor" viewBox="0 0 20 20">
                  <path d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
              </div>
            </div>
            <div class="mr-4">
              <p class="text-sm font-medium text-gray-500">نرخ موفقیت</p>
              <p class="text-2xl font-semibold text-gray-900">{{ monitoring.successRate }}%</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Performance Chart -->
      <div class="h-64 bg-gray-50 rounded-lg flex items-center justify-center">
        <div class="text-center">
          <svg class="w-16 h-16 text-gray-400 mx-auto mb-4" fill="currentColor" viewBox="0 0 20 20">
            <path d="M2 11a1 1 0 011-1h2a1 1 0 011 1v5a1 1 0 01-1 1H3a1 1 0 01-1-1v-5zM8 7a1 1 0 011-1h2a1 1 0 011 1v9a1 1 0 01-1 1H9a1 1 0 01-1-1V7zM14 4a1 1 0 011-1h2a1 1 0 011 1v12a1 1 0 01-1 1h-2a1 1 0 01-1-1V4z"></path>
          </svg>
          <p class="text-gray-500">نمودار عملکرد در اینجا نمایش داده می‌شود</p>
        </div>
      </div>
    </div>

    <!-- Test Results -->
    <div class="bg-white rounded-lg shadow px-4 py-4">
      <h3 class="text-lg font-medium text-gray-900 mb-4">نتایج تست‌ها</h3>
      
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نام تست</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">سناریو</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ اجرا</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مدت زمان</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نتایج</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="result in testResults" :key="result.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm font-medium text-gray-900">{{ result.name }}</div>
                <div class="text-sm text-gray-500">{{ result.description }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ result.scenario }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getResultStatusClass(result.status)" class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium">
                  {{ getResultStatusText(result.status) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ formatDate(result.executedAt) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ result.duration }}s
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm text-gray-900">
                  <div>موفق: {{ result.passed }}</div>
                  <div>ناموفق: {{ result.failed }}</div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3" @click="viewTestDetails(result)">جزئیات</button>
                <button class="text-green-600 hover:text-green-900 ml-3" @click="downloadReport(result)">گزارش</button>
                <button class="text-purple-600 hover:text-purple-900" @click="rerunTest(result)">اجرای مجدد</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

interface Step {
  id: number;
  name: string;
  description: string;
  status: string;
  progress: number;
}

interface TestEnvironment {
  id: number;
  name: string;
  description: string;
  status: string;
  testUsers: number;
  duration: string;
  successRate: number;
}

interface TestResult {
  id: number;
  name: string;
  description: string;
  scenario: string;
  status: string;
  executedAt: string;
  duration: number;
  passed: number;
  failed: number;
}

// Props and Emits
defineProps<{
  scenarios?: Record<string, unknown>[]
}>()

defineEmits<{
  runSetup: [steps: Step[]]
  runTests: [tests: TestEnvironment[]]
  viewResults: [results: TestResult]
}>()

// Reactive data
const setupSteps = ref<Step[]>([
  {
    id: 1,
    name: 'بررسی پیش‌نیازها',
    description: 'بررسی سیستم‌های مورد نیاز و تنظیمات پایه',
    status: 'completed',
    progress: 100
  },
  {
    id: 2,
    name: 'راه‌اندازی پایگاه داده',
    description: 'ایجاد جداول و ساختارهای مورد نیاز',
    status: 'completed',
    progress: 100
  },
  {
    id: 3,
    name: 'تنظیم قوانین امتیازدهی',
    description: 'تعریف قوانین و پارامترهای امتیازدهی',
    status: 'running',
    progress: 65
  },
  {
    id: 4,
    name: 'راه‌اندازی سیستم‌های تحلیل',
    description: 'فعال‌سازی سیستم‌های هوش مصنوعی و تحلیل',
    status: 'pending',
    progress: 0
  },
  {
    id: 5,
    name: 'تست اولیه',
    description: 'اجرای تست‌های اولیه و بررسی عملکرد',
    status: 'pending',
    progress: 0
  },
  {
    id: 6,
    name: 'فعال‌سازی کامل',
    description: 'فعال‌سازی کامل سیستم برای کاربران',
    status: 'pending',
    progress: 0
  }
])

const testEnvironments = ref([
  {
    id: 1,
    name: 'محیط تست پایه',
    description: 'تست قابلیت‌های اصلی سیستم',
    status: 'ready',
    testUsers: 100,
    duration: '2 ساعت',
    successRate: 95
  },
  {
    id: 2,
    name: 'محیط تست پیشرفته',
    description: 'تست سناریوهای پیچیده و عملکرد بالا',
    status: 'running',
    testUsers: 500,
    duration: '6 ساعت',
    successRate: 87
  },
  {
    id: 3,
    name: 'محیط تست استرس',
    description: 'تست عملکرد تحت فشار بالا',
    status: 'ready',
    testUsers: 1000,
    duration: '12 ساعت',
    successRate: 92
  }
])

const testResults = ref([
  {
    id: 1,
    name: 'تست سیستم پایه',
    description: 'تست قابلیت‌های اصلی',
    scenario: 'سناریو پایه فروشگاه',
    status: 'passed',
    executedAt: '2024-01-15T10:30:00Z',
    duration: 45,
    passed: 25,
    failed: 0
  },
  {
    id: 2,
    name: 'تست سیستم پیشرفته',
    description: 'تست سناریوهای پیچیده',
    scenario: 'سناریو پیشرفته وفاداری',
    status: 'running',
    executedAt: '2024-01-15T09:15:00Z',
    duration: 120,
    passed: 18,
    failed: 2
  },
  {
    id: 3,
    name: 'تست عملکرد',
    description: 'تست سرعت و کارایی',
    scenario: 'سناریو ارجاع و مشارکت',
    status: 'failed',
    executedAt: '2024-01-14T16:45:00Z',
    duration: 30,
    passed: 8,
    failed: 12
  }
])

const monitoring = ref({
  successfulTests: 45,
  failedTests: 3,
  avgTime: 2.3,
  successRate: 93.8
})

// Methods
const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleDateString('fa-IR')
}

const getStepClass = (status: string) => {
  switch (status) {
    case 'completed':
      return 'bg-green-50 border-green-200'
    case 'running':
      return 'bg-blue-50 border-blue-200'
    case 'failed':
      return 'bg-red-50 border-red-200'
    default:
      return 'bg-gray-50 border-gray-200'
  }
}

const getStepIconClass = (status: string) => {
  switch (status) {
    case 'completed':
      return 'bg-green-500'
    case 'running':
      return 'bg-blue-500'
    case 'failed':
      return 'bg-red-500'
    default:
      return 'bg-gray-300'
  }
}

const getTestStatusClass = (status: string) => {
  switch (status) {
    case 'ready':
      return 'bg-green-100 text-green-800'
    case 'running':
      return 'bg-blue-100 text-blue-800'
    case 'completed':
      return 'bg-purple-100 text-purple-800'
    case 'failed':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getTestStatusText = (status: string) => {
  switch (status) {
    case 'ready':
      return 'آماده'
    case 'running':
      return 'در حال اجرا'
    case 'completed':
      return 'تکمیل شده'
    case 'failed':
      return 'ناموفق'
    default:
      return 'نامشخص'
  }
}

const getResultStatusClass = (status: string) => {
  switch (status) {
    case 'passed':
      return 'bg-green-100 text-green-800'
    case 'running':
      return 'bg-blue-100 text-blue-800'
    case 'failed':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getResultStatusText = (status: string) => {
  switch (status) {
    case 'passed':
      return 'موفق'
    case 'running':
      return 'در حال اجرا'
    case 'failed':
      return 'ناموفق'
    default:
      return 'نامشخص'
  }
}

const runFullSetup = () => {
  // Run full system setup
}

const runTests = () => {
  // Run all tests
}

const runStep = (step: Step) => {
  step.status = 'running'
  step.progress = 0
  
  // Simulate progress
  const interval = setInterval(() => {
    step.progress += 10
    if (step.progress >= 100) {
      step.status = 'completed'
      clearInterval(interval)
    }
  }, 500)
}

const resetStep = (step: Step) => {
  step.status = 'pending'
  step.progress = 0
}

const viewStepDetails = (_step: Step) => {
  // View step details
}

const startTest = (test: TestEnvironment) => {
  test.status = 'running'
  // Start test
}

const stopTest = (test: TestEnvironment) => {
  test.status = 'completed'
  // Stop test
}

const viewTestResults = (_test: TestEnvironment) => {
  // View test results
}

const viewTestDetails = (_result: TestResult) => {
  // View test details
}

const downloadReport = (_result: TestResult) => {
  // Download test report
}

const rerunTest = (_result: TestResult) => {
  // Rerun test
}

// Lifecycle
// onMounted removed as it only contained console.log
</script> 
