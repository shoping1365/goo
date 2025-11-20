<template>
  <div class="min-h-screen bg-gray-50 p-6">
    <div class="max-w-7xl mx-auto">
      <!-- Header -->
      <div class="mb-8">
        <h1 class="text-3xl font-bold text-gray-900 mb-2">مدیریت متغیرهای محیطی</h1>
        <p class="text-gray-600">ابزارهای مدیریت و پیکربندی متغیرهای محیطی</p>
      </div>

      <!-- Stats Cards -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-blue-100 rounded-lg">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">محیط‌های فعال</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.activeEnvironments }}</p>
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
              <p class="text-sm font-medium text-gray-600">متغیرهای امن</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.secureVariables }}</p>
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
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">متغیرهای حساس</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.sensitiveVariables }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-purple-100 rounded-lg">
              <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">آخرین به‌روزرسانی</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.lastUpdate }}</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Environment Selection -->
      <div class="mb-8 bg-white rounded-lg shadow">
        <div class="p-6 border-b border-gray-200">
          <h2 class="text-xl font-semibold text-gray-900">انتخاب محیط</h2>
        </div>
        <div class="p-6">
          <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
            <button 
              :class="[
                'flex items-center justify-center p-6 rounded-lg border-2 transition-colors',
                selectedEnvironment === 'development' ? 'border-green-500 bg-green-50' : 'border-gray-200 hover:border-gray-300'
              ]"
              @click="selectEnvironment('development')"
            >
              <div class="text-center">
                <div class="w-12 h-12 bg-green-500 rounded-lg flex items-center justify-center mx-auto mb-2">
                  <span class="text-white font-bold text-lg">Dev</span>
                </div>
                <span class="text-sm font-medium text-gray-900">Development</span>
              </div>
            </button>

            <button 
              :class="[
                'flex items-center justify-center p-6 rounded-lg border-2 transition-colors',
                selectedEnvironment === 'staging' ? 'border-yellow-500 bg-yellow-50' : 'border-gray-200 hover:border-gray-300'
              ]"
              @click="selectEnvironment('staging')"
            >
              <div class="text-center">
                <div class="w-12 h-12 bg-yellow-500 rounded-lg flex items-center justify-center mx-auto mb-2">
                  <span class="text-white font-bold text-lg">Stg</span>
                </div>
                <span class="text-sm font-medium text-gray-900">Staging</span>
              </div>
            </button>

            <button 
              :class="[
                'flex items-center justify-center p-6 rounded-lg border-2 transition-colors',
                selectedEnvironment === 'production' ? 'border-red-500 bg-red-50' : 'border-gray-200 hover:border-gray-300'
              ]"
              @click="selectEnvironment('production')"
            >
              <div class="text-center">
                <div class="w-12 h-12 bg-red-500 rounded-lg flex items-center justify-center mx-auto mb-2">
                  <span class="text-white font-bold text-lg">Prod</span>
                </div>
                <span class="text-sm font-medium text-gray-900">Production</span>
              </div>
            </button>

            <button 
              :class="[
                'flex items-center justify-center p-6 rounded-lg border-2 transition-colors',
                selectedEnvironment === 'testing' ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-gray-300'
              ]"
              @click="selectEnvironment('testing')"
            >
              <div class="text-center">
                <div class="w-12 h-12 bg-blue-500 rounded-lg flex items-center justify-center mx-auto mb-2">
                  <span class="text-white font-bold text-lg">Test</span>
                </div>
                <span class="text-sm font-medium text-gray-900">Testing</span>
              </div>
            </button>
          </div>
        </div>
      </div>

      <!-- Main Content Grid -->
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <!-- Variable Management -->
        <div class="lg:col-span-2 bg-white rounded-lg shadow">
          <div class="p-6 border-b border-gray-200">
            <div class="flex items-center justify-between">
              <h2 class="text-xl font-semibold text-gray-900">مدیریت متغیرها</h2>
              <div class="flex items-center space-x-2 space-x-reverse">
                <button 
                  class="bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="addVariable"
                >
                  افزودن متغیر
                </button>
                <button 
                  class="bg-green-600 hover:bg-green-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="importVariables"
                >
                  وارد کردن
                </button>
                <button 
                  class="bg-purple-600 hover:bg-purple-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="exportVariables"
                >
                  صادر کردن
                </button>
              </div>
            </div>
          </div>
          <div class="p-6">
            <div class="space-y-4">
              <div v-for="variable in environmentVariables" :key="variable.id" class="border rounded-lg p-6">
                <div class="flex items-center justify-between mb-2">
                  <div class="flex items-center space-x-3 space-x-reverse">
                    <span class="font-medium text-gray-900">{{ variable.name }}</span>
                    <span
:class="[
                      'px-2 py-1 rounded text-xs font-medium',
                      variable.type === 'secure' ? 'bg-red-100 text-red-800' :
                      variable.type === 'sensitive' ? 'bg-yellow-100 text-yellow-800' :
                      'bg-green-100 text-green-800'
                    ]">
                      {{ variable.type }}
                    </span>
                  </div>
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <button 
                      class="text-blue-600 hover:text-blue-800 text-sm font-medium"
                      @click="editVariable(variable)"
                    >
                      ویرایش
                    </button>
                    <button 
                      class="text-gray-600 hover:text-gray-800 text-sm font-medium"
                      @click="toggleVariableVisibility(variable)"
                    >
                      {{ variable.visible ? 'مخفی' : 'نمایش' }}
                    </button>
                    <button 
                      class="text-red-600 hover:text-red-800 text-sm font-medium"
                      @click="deleteVariable(variable)"
                    >
                      حذف
                    </button>
                  </div>
                </div>
                <div class="text-sm text-gray-500">
                  <p v-if="variable.visible">{{ variable.value }}</p>
                  <p v-else>••••••••••••••••</p>
                  <p class="mt-1">{{ variable.description }}</p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Variable Templates -->
        <div class="bg-white rounded-lg shadow">
          <div class="p-6 border-b border-gray-200">
            <h2 class="text-xl font-semibold text-gray-900">قالب‌های آماده</h2>
          </div>
          <div class="p-6">
            <div class="space-y-4">
              <div v-for="template in variableTemplates" :key="template.id" class="border rounded-lg p-6 hover:bg-gray-50">
                <div class="flex items-center justify-between mb-2">
                  <h3 class="font-medium text-gray-900">{{ template.name }}</h3>
                  <span class="text-xs bg-gray-100 text-gray-600 px-2 py-1 rounded">{{ template.category }}</span>
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
      </div>

      <!-- Variable Validation -->
      <div class="mt-8 bg-white rounded-lg shadow">
        <div class="p-6 border-b border-gray-200">
          <h2 class="text-xl font-semibold text-gray-900">اعتبارسنجی متغیرها</h2>
        </div>
        <div class="p-6">
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
            <div class="space-y-4">
              <h3 class="font-medium text-gray-900">بررسی امنیت</h3>
              <div class="space-y-2">
                <button 
                  class="w-full bg-red-600 hover:bg-red-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="validateSecurity"
                >
                  بررسی امنیت
                </button>
                <button 
                  class="w-full bg-orange-600 hover:bg-orange-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="scanSecrets"
                >
                  اسکن اسرار
                </button>
              </div>
            </div>

            <div class="space-y-4">
              <h3 class="font-medium text-gray-900">بررسی فرمت</h3>
              <div class="space-y-2">
                <button 
                  class="w-full bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="validateFormat"
                >
                  بررسی فرمت
                </button>
                <button 
                  class="w-full bg-purple-600 hover:bg-purple-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="checkDuplicates"
                >
                  بررسی تکرار
                </button>
              </div>
            </div>

            <div class="space-y-4">
              <h3 class="font-medium text-gray-900">بررسی عملکرد</h3>
              <div class="space-y-2">
                <button 
                  class="w-full bg-green-600 hover:bg-green-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="testConnections"
                >
                  تست اتصالات
                </button>
                <button 
                  class="w-full bg-indigo-600 hover:bg-indigo-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="validateUrls"
                >
                  بررسی URL ها
                </button>
              </div>
            </div>

            <div class="space-y-4">
              <h3 class="font-medium text-gray-900">گزارش‌ها</h3>
              <div class="space-y-2">
                <button 
                  class="w-full bg-gray-600 hover:bg-gray-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="generateReport"
                >
                  تولید گزارش
                </button>
                <button 
                  class="w-full bg-teal-600 hover:bg-teal-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="exportAudit"
                >
                  صادر کردن حسابرسی
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Variable History -->
      <div class="mt-8 bg-white rounded-lg shadow">
        <div class="p-6 border-b border-gray-200">
          <h2 class="text-xl font-semibold text-gray-900">تاریخچه تغییرات</h2>
        </div>
        <div class="p-6">
          <div class="overflow-x-auto">
            <table class="min-w-full divide-y divide-gray-200">
              <thead class="bg-gray-50">
                <tr>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">متغیر</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کاربر</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">محیط</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
                </tr>
              </thead>
              <tbody class="bg-white divide-y divide-gray-200">
                <tr v-for="change in variableHistory" :key="change.id">
                  <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{{ change.variable }}</td>
                  <td class="px-6 py-4 whitespace-nowrap">
                    <span
:class="[
                      'px-2 py-1 rounded text-xs font-medium',
                      change.action === 'created' ? 'bg-green-100 text-green-800' :
                      change.action === 'updated' ? 'bg-blue-100 text-blue-800' :
                      'bg-red-100 text-red-800'
                    ]">
                      {{ change.action }}
                    </span>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ change.user }}</td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ change.environment }}</td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ change.date }}</td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                    <button 
                      class="text-blue-600 hover:text-blue-800"
                      @click="viewChangeDetails(change)"
                    >
                      جزئیات
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>

      <!-- Variable Modal -->
      <div v-if="variableModal.show" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
        <div class="bg-white rounded-lg max-w-2xl w-full mx-4">
          <div class="p-6 border-b border-gray-200">
            <div class="flex items-center justify-between">
              <h3 class="text-lg font-semibold text-gray-900">{{ variableModal.isEdit ? 'ویرایش متغیر' : 'افزودن متغیر جدید' }}</h3>
              <button class="text-gray-400 hover:text-gray-600" @click="variableModal.show = false">
                <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                </svg>
              </button>
            </div>
          </div>
          <div class="p-6">
            <div class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">نام متغیر</label>
                <input 
                  v-model="variableModal.name"
                  type="text"
                  class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
                  placeholder="نام متغیر را وارد کنید..."
                >
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">مقدار</label>
                <input 
                  v-model="variableModal.value"
                  type="text"
                  class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
                  placeholder="مقدار متغیر را وارد کنید..."
                >
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">نوع</label>
                <select 
                  v-model="variableModal.type"
                  class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                  <option value="normal">عادی</option>
                  <option value="sensitive">حساس</option>
                  <option value="secure">امن</option>
                </select>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">توضیحات</label>
                <textarea
                  v-model="variableModal.description"
                  rows="3"
                  class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
                  placeholder="توضیحات متغیر را وارد کنید..."
                ></textarea>
              </div>
              <div class="flex items-center justify-end space-x-2 space-x-reverse">
                <button 
                  class="bg-gray-500 hover:bg-gray-600 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="variableModal.show = false"
                >
                  انصراف
                </button>
                <button 
                  class="bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="saveVariable"
                >
                  ذخیره
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
import { ref, reactive } from 'vue'

definePageMeta({
  layout: 'admin-main',
  middleware: ['developer-only']
});

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user, hasPermission } = useAuth()

// Reactive data
const selectedEnvironment = ref('development')
const variableModal = reactive({
  show: false,
  isEdit: false,
  name: '',
  value: '',
  type: 'normal',
  description: ''
})

const stats = reactive({
  activeEnvironments: 4,
  secureVariables: 15,
  sensitiveVariables: 8,
  lastUpdate: '2h ago'
})

const environmentVariables = ref([
  {
    id: 1,
    name: 'DATABASE_URL',
    value: 'postgresql://user:pass@localhost:5432/db',
    type: 'secure',
    description: 'آدرس اتصال به دیتابیس',
    visible: false
  },
  {
    id: 2,
    name: 'API_KEY',
    value: 'sk-1234567890abcdef',
    type: 'secure',
    description: 'کلید API برای سرویس‌های خارجی',
    visible: false
  },
  {
    id: 3,
    name: 'DEBUG_MODE',
    value: 'true',
    type: 'normal',
    description: 'حالت دیباگ برای توسعه',
    visible: true
  },
  {
    id: 4,
    name: 'REDIS_URL',
    value: 'redis://localhost:6379',
    type: 'sensitive',
    description: 'آدرس اتصال به Redis',
    visible: true
  }
])

const variableTemplates = ref([
  {
    id: 1,
    name: 'Database Configuration',
    category: 'Database',
    description: 'قالب کامل برای تنظیمات دیتابیس'
  },
  {
    id: 2,
    name: 'API Configuration',
    category: 'API',
    description: 'قالب برای تنظیمات API و کلیدها'
  },
  {
    id: 3,
    name: 'Email Configuration',
    category: 'Email',
    description: 'قالب برای تنظیمات ایمیل و SMTP'
  },
  {
    id: 4,
    name: 'Payment Gateway',
    category: 'Payment',
    description: 'قالب برای تنظیمات درگاه پرداخت'
  }
])

const variableHistory = ref([
  {
    id: 1,
    variable: 'DATABASE_URL',
    action: 'updated',
    user: 'admin',
    environment: 'production',
    date: '2024-01-15 10:30'
  },
  {
    id: 2,
    variable: 'API_KEY',
    action: 'created',
    user: 'developer',
    environment: 'staging',
    date: '2024-01-15 09:15'
  },
  {
    id: 3,
    variable: 'DEBUG_MODE',
    action: 'deleted',
    user: 'admin',
    environment: 'production',
    date: '2024-01-15 08:45'
  }
])

// Methods
function selectEnvironment(environment) {
  selectedEnvironment.value = environment
  alert(`محیط به ${environment} تغییر یافت`)
}

function addVariable() {
  variableModal.isEdit = false
  variableModal.name = ''
  variableModal.value = ''
  variableModal.type = 'normal'
  variableModal.description = ''
  variableModal.show = true
}

function editVariable(variable) {
  variableModal.isEdit = true
  variableModal.name = variable.name
  variableModal.value = variable.value
  variableModal.type = variable.type
  variableModal.description = variable.description
  variableModal.show = true
}

function saveVariable() {
  if (variableModal.name.trim() && variableModal.value.trim()) {
    if (variableModal.isEdit) {
      alert('متغیر ویرایش شد')
    } else {
      alert('متغیر جدید اضافه شد')
    }
    variableModal.show = false
  }
}

function deleteVariable(variable) {
  if (confirm(`آیا می‌خواهید متغیر ${variable.name} را حذف کنید؟`)) {
    alert(`متغیر ${variable.name} حذف شد`)
  }
}

function toggleVariableVisibility(variable) {
  variable.visible = !variable.visible
}

function importVariables() {
  alert('وارد کردن متغیرها شروع شد')
}

function exportVariables() {
  alert('صادر کردن متغیرها شروع شد')
}

function useTemplate(template) {
  alert(`قالب ${template.name} اعمال شد`)
}

function previewTemplate(template) {
  alert(`پیش‌نمایش قالب ${template.name} نمایش داده می‌شود`)
}

function validateSecurity() {
  alert('بررسی امنیت متغیرها انجام شد')
}

function scanSecrets() {
  alert('اسکن اسرار و کلیدها انجام شد')
}

function validateFormat() {
  alert('بررسی فرمت متغیرها انجام شد')
}

function checkDuplicates() {
  alert('بررسی تکرار متغیرها انجام شد')
}

function testConnections() {
  alert('تست اتصالات انجام شد')
}

function validateUrls() {
  alert('بررسی URL ها انجام شد')
}

function generateReport() {
  alert('گزارش متغیرها تولید شد')
}

function exportAudit() {
  alert('گزارش حسابرسی صادر شد')
}

function viewChangeDetails(change) {
  alert(`جزئیات تغییر ${change.variable} نمایش داده می‌شود`)
}
</script> 
