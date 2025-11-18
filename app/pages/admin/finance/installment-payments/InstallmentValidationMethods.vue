<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
    <div class="flex items-center justify-between mb-6">
      <h3 class="text-lg font-semibold text-gray-900">روش‌های اعتبارسنجی</h3>
      <div class="flex items-center space-x-2 space-x-reverse">
        <select v-model="selectedPeriod" class="text-sm border border-gray-300 rounded-md px-3 py-1">
          <option value="7">7 روز گذشته</option>
          <option value="30">30 روز گذشته</option>
          <option value="90">90 روز گذشته</option>
          <option value="365">یک سال گذشته</option>
        </select>
        <button @click="refreshData" class="text-blue-600 hover:text-blue-800">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
        </button>
      </div>
    </div>

    <!-- Validation Methods Overview -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-6">
      <div class="bg-blue-50 rounded-lg p-6">
        <div class="flex items-center">
          <div class="p-2 bg-blue-500 rounded-lg">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <div class="mr-3">
            <p class="text-sm text-gray-600">کل بررسی‌ها</p>
            <p class="text-xl font-bold text-gray-900">{{ overview.totalValidations }}</p>
          </div>
        </div>
      </div>

      <div class="bg-green-50 rounded-lg p-6">
        <div class="flex items-center">
          <div class="p-2 bg-green-500 rounded-lg">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
            </svg>
          </div>
          <div class="mr-3">
            <p class="text-sm text-gray-600">نرخ موفقیت</p>
            <p class="text-xl font-bold text-gray-900">{{ overview.successRate }}%</p>
          </div>
        </div>
      </div>

      <div class="bg-yellow-50 rounded-lg p-6">
        <div class="flex items-center">
          <div class="p-2 bg-yellow-500 rounded-lg">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <div class="mr-3">
            <p class="text-sm text-gray-600">میانگین زمان</p>
            <p class="text-xl font-bold text-gray-900">{{ overview.avgTime }}s</p>
          </div>
        </div>
      </div>

      <div class="bg-purple-50 rounded-lg p-6">
        <div class="flex items-center">
          <div class="p-2 bg-purple-500 rounded-lg">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
            </svg>
          </div>
          <div class="mr-3">
            <p class="text-sm text-gray-600">روش‌های فعال</p>
            <p class="text-xl font-bold text-gray-900">{{ overview.activeMethods }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Validation Methods List -->
    <div class="space-y-4">
      <!-- National ID Validation -->
      <div class="bg-gray-50 rounded-lg p-6">
        <div class="flex items-center justify-between mb-4">
          <div class="flex items-center">
            <div class="w-10 h-10 bg-blue-100 rounded-full flex items-center justify-center mr-3">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H5a2 2 0 00-2 2v9a2 2 0 002 2h14a2 2 0 002-2V8a2 2 0 00-2-2h-5m-4 0V4a2 2 0 114 0v2m-4 0a2 2 0 104 0m-5 8a2 2 0 100-4 2 2 0 000 4zm0 0c1.306 0 2.417.835 2.83 2M9 14a3.001 3.001 0 00-2.83 2M15 11h3m-3 4h2" />
              </svg>
            </div>
            <div>
              <h4 class="text-lg font-semibold text-gray-900">اعتبارسنجی کد ملی</h4>
              <p class="text-sm text-gray-600">بررسی صحت و اعتبار کد ملی از طریق سرویس‌های دولتی</p>
            </div>
          </div>
          <div class="flex items-center">
            <div class="w-3 h-3 rounded-full mr-2" :class="methods.nationalId.status === 'active' ? 'bg-green-500' : 'bg-red-500'"></div>
            <span class="text-sm" :class="methods.nationalId.status === 'active' ? 'text-green-600' : 'text-red-600'">
              {{ methods.nationalId.status === 'active' ? 'فعال' : 'غیرفعال' }}
            </span>
          </div>
        </div>
        
        <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-4">
          <div class="text-center">
            <div class="text-2xl font-bold text-blue-600">{{ methods.nationalId.totalChecks }}</div>
            <div class="text-sm text-gray-600">کل بررسی‌ها</div>
          </div>
          <div class="text-center">
            <div class="text-2xl font-bold text-green-600">{{ methods.nationalId.successRate }}%</div>
            <div class="text-sm text-gray-600">نرخ موفقیت</div>
          </div>
          <div class="text-center">
            <div class="text-2xl font-bold text-yellow-600">{{ methods.nationalId.avgResponseTime }}ms</div>
            <div class="text-sm text-gray-600">زمان پاسخ</div>
          </div>
          <div class="text-center">
            <div class="text-2xl font-bold text-purple-600">{{ methods.nationalId.errorRate }}%</div>
            <div class="text-sm text-gray-600">نرخ خطا</div>
          </div>
        </div>

        <div class="flex items-center space-x-4 space-x-reverse">
          <button @click="testMethod('nationalId')" class="text-blue-600 hover:text-blue-800 text-sm font-medium">
            تست سرویس
          </button>
          <button @click="viewLogs('nationalId')" class="text-green-600 hover:text-green-800 text-sm font-medium">
            مشاهده لاگ‌ها
          </button>
          <button @click="configure('nationalId')" class="text-purple-600 hover:text-purple-800 text-sm font-medium">
            پیکربندی
          </button>
        </div>
      </div>

      <!-- Mobile Phone Validation -->
      <div class="bg-gray-50 rounded-lg p-6">
        <div class="flex items-center justify-between mb-4">
          <div class="flex items-center">
            <div class="w-10 h-10 bg-green-100 rounded-full flex items-center justify-center mr-3">
              <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 18h.01M8 21h8a2 2 0 002-2V5a2 2 0 00-2-2H8a2 2 0 00-2 2v14a2 2 0 002 2z" />
              </svg>
            </div>
            <div>
              <h4 class="text-lg font-semibold text-gray-900">تایید شماره موبایل</h4>
              <p class="text-sm text-gray-600">ارسال کد تایید از طریق پیامک و بررسی صحت شماره</p>
            </div>
          </div>
          <div class="flex items-center">
            <div class="w-3 h-3 rounded-full mr-2" :class="methods.mobile.status === 'active' ? 'bg-green-500' : 'bg-red-500'"></div>
            <span class="text-sm" :class="methods.mobile.status === 'active' ? 'text-green-600' : 'text-red-600'">
              {{ methods.mobile.status === 'active' ? 'فعال' : 'غیرفعال' }}
            </span>
          </div>
        </div>
        
        <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-4">
          <div class="text-center">
            <div class="text-2xl font-bold text-blue-600">{{ methods.mobile.totalChecks }}</div>
            <div class="text-sm text-gray-600">کل بررسی‌ها</div>
          </div>
          <div class="text-center">
            <div class="text-2xl font-bold text-green-600">{{ methods.mobile.successRate }}%</div>
            <div class="text-sm text-gray-600">نرخ موفقیت</div>
          </div>
          <div class="text-center">
            <div class="text-2xl font-bold text-yellow-600">{{ methods.mobile.avgResponseTime }}ms</div>
            <div class="text-sm text-gray-600">زمان پاسخ</div>
          </div>
          <div class="text-center">
            <div class="text-2xl font-bold text-purple-600">{{ methods.mobile.errorRate }}%</div>
            <div class="text-sm text-gray-600">نرخ خطا</div>
          </div>
        </div>

        <div class="flex items-center space-x-4 space-x-reverse">
          <button @click="testMethod('mobile')" class="text-blue-600 hover:text-blue-800 text-sm font-medium">
            تست سرویس
          </button>
          <button @click="viewLogs('mobile')" class="text-green-600 hover:text-green-800 text-sm font-medium">
            مشاهده لاگ‌ها
          </button>
          <button @click="configure('mobile')" class="text-purple-600 hover:text-purple-800 text-sm font-medium">
            پیکربندی
          </button>
        </div>
      </div>

      <!-- Bank Account Validation -->
      <div class="bg-gray-50 rounded-lg p-6">
        <div class="flex items-center justify-between mb-4">
          <div class="flex items-center">
            <div class="w-10 h-10 bg-yellow-100 rounded-full flex items-center justify-center mr-3">
              <svg class="w-6 h-6 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z" />
              </svg>
            </div>
            <div>
              <h4 class="text-lg font-semibold text-gray-900">اعتبارسنجی حساب بانکی</h4>
              <p class="text-sm text-gray-600">بررسی صحت شماره حساب و کارت بانکی از طریق بانک‌ها</p>
            </div>
          </div>
          <div class="flex items-center">
            <div class="w-3 h-3 rounded-full mr-2" :class="methods.bankAccount.status === 'active' ? 'bg-green-500' : 'bg-red-500'"></div>
            <span class="text-sm" :class="methods.bankAccount.status === 'active' ? 'text-green-600' : 'text-red-600'">
              {{ methods.bankAccount.status === 'active' ? 'فعال' : 'غیرفعال' }}
            </span>
          </div>
        </div>
        
        <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-4">
          <div class="text-center">
            <div class="text-2xl font-bold text-blue-600">{{ methods.bankAccount.totalChecks }}</div>
            <div class="text-sm text-gray-600">کل بررسی‌ها</div>
          </div>
          <div class="text-center">
            <div class="text-2xl font-bold text-green-600">{{ methods.bankAccount.successRate }}%</div>
            <div class="text-sm text-gray-600">نرخ موفقیت</div>
          </div>
          <div class="text-center">
            <div class="text-2xl font-bold text-yellow-600">{{ methods.bankAccount.avgResponseTime }}ms</div>
            <div class="text-sm text-gray-600">زمان پاسخ</div>
          </div>
          <div class="text-center">
            <div class="text-2xl font-bold text-purple-600">{{ methods.bankAccount.errorRate }}%</div>
            <div class="text-sm text-gray-600">نرخ خطا</div>
          </div>
        </div>

        <div class="flex items-center space-x-4 space-x-reverse">
          <button @click="testMethod('bankAccount')" class="text-blue-600 hover:text-blue-800 text-sm font-medium">
            تست سرویس
          </button>
          <button @click="viewLogs('bankAccount')" class="text-green-600 hover:text-green-800 text-sm font-medium">
            مشاهده لاگ‌ها
          </button>
          <button @click="configure('bankAccount')" class="text-purple-600 hover:text-purple-800 text-sm font-medium">
            پیکربندی
          </button>
        </div>
      </div>

      <!-- Credit Bureau Check -->
      <div class="bg-gray-50 rounded-lg p-6">
        <div class="flex items-center justify-between mb-4">
          <div class="flex items-center">
            <div class="w-10 h-10 bg-red-100 rounded-full flex items-center justify-center mr-3">
              <svg class="w-6 h-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
              </svg>
            </div>
            <div>
              <h4 class="text-lg font-semibold text-gray-900">استعلام مرکز اطلاعات اعتباری</h4>
              <p class="text-sm text-gray-600">بررسی سابقه اعتباری از طریق مرکز اطلاعات اعتباری بانک مرکزی</p>
            </div>
          </div>
          <div class="flex items-center">
            <div class="w-3 h-3 rounded-full mr-2" :class="methods.creditBureau.status === 'active' ? 'bg-green-500' : 'bg-red-500'"></div>
            <span class="text-sm" :class="methods.creditBureau.status === 'active' ? 'text-green-600' : 'text-red-600'">
              {{ methods.creditBureau.status === 'active' ? 'فعال' : 'غیرفعال' }}
            </span>
          </div>
        </div>
        
        <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-4">
          <div class="text-center">
            <div class="text-2xl font-bold text-blue-600">{{ methods.creditBureau.totalChecks }}</div>
            <div class="text-sm text-gray-600">کل بررسی‌ها</div>
          </div>
          <div class="text-center">
            <div class="text-2xl font-bold text-green-600">{{ methods.creditBureau.successRate }}%</div>
            <div class="text-sm text-gray-600">نرخ موفقیت</div>
          </div>
          <div class="text-center">
            <div class="text-2xl font-bold text-yellow-600">{{ methods.creditBureau.avgResponseTime }}ms</div>
            <div class="text-sm text-gray-600">زمان پاسخ</div>
          </div>
          <div class="text-center">
            <div class="text-2xl font-bold text-purple-600">{{ methods.creditBureau.errorRate }}%</div>
            <div class="text-sm text-gray-600">نرخ خطا</div>
          </div>
        </div>

        <div class="flex items-center space-x-4 space-x-reverse">
          <button @click="testMethod('creditBureau')" class="text-blue-600 hover:text-blue-800 text-sm font-medium">
            تست سرویس
          </button>
          <button @click="viewLogs('creditBureau')" class="text-green-600 hover:text-green-800 text-sm font-medium">
            مشاهده لاگ‌ها
          </button>
          <button @click="configure('creditBureau')" class="text-purple-600 hover:text-purple-800 text-sm font-medium">
            پیکربندی
          </button>
        </div>
      </div>

      <!-- Employment Verification -->
      <div class="bg-gray-50 rounded-lg p-6">
        <div class="flex items-center justify-between mb-4">
          <div class="flex items-center">
            <div class="w-10 h-10 bg-indigo-100 rounded-full flex items-center justify-center mr-3">
              <svg class="w-6 h-6 text-indigo-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 13.255A23.931 23.931 0 0112 15c-3.183 0-6.22-.62-9-1.745M16 6V4a2 2 0 00-2-2h-4a2 2 0 00-2-2v2m8 0v2m0 0v2m0-2h2m-2 0h-2m-4 4h.01M19 20a2 2 0 01-2 2H7a2 2 0 01-2-2V10a2 2 0 012-2h5m5 0v2a2 2 0 002 2h2a2 2 0 002-2v-2a2 2 0 00-2-2h-2z" />
              </svg>
            </div>
            <div>
              <h4 class="text-lg font-semibold text-gray-900">تایید شغل و درآمد</h4>
              <p class="text-sm text-gray-600">بررسی وضعیت شغلی و تأیید درآمد از طریق سازمان‌های مربوطه</p>
            </div>
          </div>
          <div class="flex items-center">
            <div class="w-3 h-3 rounded-full mr-2" :class="methods.employment.status === 'active' ? 'bg-green-500' : 'bg-red-500'"></div>
            <span class="text-sm" :class="methods.employment.status === 'active' ? 'text-green-600' : 'text-red-600'">
              {{ methods.employment.status === 'active' ? 'فعال' : 'غیرفعال' }}
            </span>
          </div>
        </div>
        
        <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-4">
          <div class="text-center">
            <div class="text-2xl font-bold text-blue-600">{{ methods.employment.totalChecks }}</div>
            <div class="text-sm text-gray-600">کل بررسی‌ها</div>
          </div>
          <div class="text-center">
            <div class="text-2xl font-bold text-green-600">{{ methods.employment.successRate }}%</div>
            <div class="text-sm text-gray-600">نرخ موفقیت</div>
          </div>
          <div class="text-center">
            <div class="text-2xl font-bold text-yellow-600">{{ methods.employment.avgResponseTime }}ms</div>
            <div class="text-sm text-gray-600">زمان پاسخ</div>
          </div>
          <div class="text-center">
            <div class="text-2xl font-bold text-purple-600">{{ methods.employment.errorRate }}%</div>
            <div class="text-sm text-gray-600">نرخ خطا</div>
          </div>
        </div>

        <div class="flex items-center space-x-4 space-x-reverse">
          <button @click="testMethod('employment')" class="text-blue-600 hover:text-blue-800 text-sm font-medium">
            تست سرویس
          </button>
          <button @click="viewLogs('employment')" class="text-green-600 hover:text-green-800 text-sm font-medium">
            مشاهده لاگ‌ها
          </button>
          <button @click="configure('employment')" class="text-purple-600 hover:text-purple-800 text-sm font-medium">
            پیکربندی
          </button>
        </div>
      </div>
    </div>

    <!-- Global Settings -->
    <div class="mt-6 bg-blue-50 rounded-lg p-6">
      <h4 class="text-md font-semibold text-gray-900 mb-4">تنظیمات کلی</h4>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر زمان انتظار (ثانیه)</label>
          <input v-model="globalSettings.timeout" type="number" class="w-full border border-gray-300 rounded-md px-3 py-2">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تعداد تلاش مجدد</label>
          <input v-model="globalSettings.retryCount" type="number" class="w-full border border-gray-300 rounded-md px-3 py-2">
        </div>
      </div>
      <div class="mt-4 flex items-center space-x-2 space-x-reverse">
        <button @click="saveGlobalSettings" class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700">
          ذخیره تنظیمات
        </button>
        <button @click="testAllMethods" class="bg-green-600 text-white px-4 py-2 rounded-md hover:bg-green-700">
          تست همه سرویس‌ها
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref, watch } from 'vue'
import { useRouter } from 'vue-router'

interface ValidationMethod {
  status: 'active' | 'inactive'
  totalChecks: number
  successRate: number
  avgResponseTime: number
  errorRate: number
}

interface Overview {
  totalValidations: number
  successRate: number
  avgTime: number
  activeMethods: number
}

interface GlobalSettings {
  timeout: number
  retryCount: number
}

const selectedPeriod = ref('30')

const overview = ref<Overview>({
  totalValidations: 0,
  successRate: 0,
  avgTime: 0,
  activeMethods: 0
})

const methods = ref({
  nationalId: {
    status: 'active' as 'active' | 'inactive',
    totalChecks: 0,
    successRate: 0,
    avgResponseTime: 0,
    errorRate: 0
  },
  mobile: {
    status: 'active' as 'active' | 'inactive',
    totalChecks: 0,
    successRate: 0,
    avgResponseTime: 0,
    errorRate: 0
  },
  bankAccount: {
    status: 'active' as 'active' | 'inactive',
    totalChecks: 0,
    successRate: 0,
    avgResponseTime: 0,
    errorRate: 0
  },
  creditBureau: {
    status: 'active' as 'active' | 'inactive',
    totalChecks: 0,
    successRate: 0,
    avgResponseTime: 0,
    errorRate: 0
  },
  employment: {
    status: 'active' as 'active' | 'inactive',
    totalChecks: 0,
    successRate: 0,
    avgResponseTime: 0,
    errorRate: 0
  }
})

const globalSettings = ref<GlobalSettings>({
  timeout: 30,
  retryCount: 3
})

const fetchValidationData = async () => {
  try {
    interface ValidationMethodsData {
      nationalId: ValidationMethod
      mobile: ValidationMethod
      bankAccount: ValidationMethod
      creditBureau: ValidationMethod
      employment: ValidationMethod
    }
    interface ValidationData {
      overview: Overview
      methods: ValidationMethodsData
      globalSettings: GlobalSettings
    }
    interface ApiResponse {
      data?: ValidationData
    }
    const response = await $fetch<ApiResponse>('/api/admin/installment-payments/validation-methods', {
      query: { period: selectedPeriod.value }
    })
    
    if (response.data) {
      overview.value = response.data.overview
      methods.value = response.data.methods
      globalSettings.value = response.data.globalSettings
    }
  } catch (error) {
    console.error('خطا در دریافت داده‌های اعتبارسنجی:', error)
    // در حالت واقعی، اینجا از createError استفاده می‌کنیم
  }
}

const refreshData = () => {
  fetchValidationData()
}

const testMethod = async (methodName: string) => {
  try {
    await $fetch(`/api/admin/installment-payments/validation-methods/${methodName}/test`)
    // نمایش پیام موفقیت
    alert(`تست ${methodName} با موفقیت انجام شد`)
  } catch (error) {
    console.error(`خطا در تست ${methodName}:`, error)
    alert(`خطا در تست ${methodName}`)
  }
}

const router = useRouter()

const viewLogs = (methodName: string) => {
  router.push(`/admin/finance/installment-payments/validation-methods/${methodName}/logs`)
}

const configure = (methodName: string) => {
  router.push(`/admin/finance/installment-payments/validation-methods/${methodName}/config`)
}

const testAllMethods = async () => {
  try {
    await $fetch('/api/admin/installment-payments/validation-methods/test-all')
    alert('تست همه سرویس‌ها با موفقیت انجام شد')
    fetchValidationData()
  } catch (error) {
    console.error('خطا در تست همه سرویس‌ها:', error)
    alert('خطا در تست سرویس‌ها')
  }
}

const saveGlobalSettings = async () => {
  try {
    await $fetch('/api/admin/installment-payments/validation-methods/settings', {
      method: 'PUT',
      body: globalSettings.value
    })
    alert('تنظیمات با موفقیت ذخیره شد')
  } catch (error) {
    console.error('خطا در ذخیره تنظیمات:', error)
    alert('خطا در ذخیره تنظیمات')
  }
}

onMounted(() => {
  fetchValidationData()
})

// Watch for period changes
watch(selectedPeriod, () => {
  fetchValidationData()
})
</script>
