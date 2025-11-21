<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">اعتبارسنجی داده‌ها</h4>
        <p class="text-sm text-gray-600 mt-1">بررسی و اعتبارسنجی داده‌های حسابداری</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          @click="runValidation"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          اجرای اعتبارسنجی
        </button>
      </div>
    </div>

    <!-- آمار اعتبارسنجی -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">آمار اعتبارسنجی</h5>
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">داده‌های معتبر</h6>
            <div class="w-3 h-3 bg-green-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-green-600">{{ validationStats.valid }}</div>
          <div class="text-xs text-gray-500 mt-1">از {{ validationStats.total }} رکورد</div>
        </div>

        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">خطاهای اعتبارسنجی</h6>
            <div class="w-3 h-3 bg-red-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-red-600">{{ validationStats.errors }}</div>
          <div class="text-xs text-gray-500 mt-1">نیاز به اصلاح</div>
        </div>

        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">هشدارها</h6>
            <div class="w-3 h-3 bg-yellow-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-yellow-600">{{ validationStats.warnings }}</div>
          <div class="text-xs text-gray-500 mt-1">نیاز به بررسی</div>
        </div>

        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">نرخ موفقیت</h6>
            <div class="w-3 h-3 bg-blue-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-blue-600">{{ validationStats.successRate }}%</div>
          <div class="text-xs text-gray-500 mt-1">داده‌های معتبر</div>
        </div>
      </div>
    </div>

    <!-- قوانین اعتبارسنجی -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">قوانین اعتبارسنجی</h5>
      </div>
      <div class="p-6">
        <div class="space-y-4">
          <div
            v-for="rule in validationRules"
            :key="rule.id"
            class="flex items-center justify-between p-6 border border-gray-200 rounded-lg hover:bg-gray-50"
          >
            <div class="flex items-center">
              <div
                class="w-3 h-3 rounded-full ml-3"
                :class="getRuleStatusColor(rule.status)"
              ></div>
              <div>
                <h6 class="text-sm font-medium text-gray-900">{{ rule.name }}</h6>
                <p class="text-xs text-gray-500">{{ rule.description }}</p>
              </div>
            </div>
            <div class="flex items-center space-x-4 space-x-reverse">
              <div class="text-right">
                <div class="text-sm font-medium text-gray-900">{{ rule.priority }}</div>
                <div class="text-xs text-gray-500">اولویت</div>
              </div>
              <div class="text-right">
                <div class="text-sm font-medium text-gray-900">{{ rule.errorCount }}</div>
                <div class="text-xs text-gray-500">خطا</div>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <button
                  class="text-blue-600 hover:text-blue-800"
                  @click="editRule(rule)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                  </svg>
                </button>
                <button
                  class="text-gray-600 hover:text-gray-800"
                  @click="toggleRule(rule)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="rule.enabled ? 'M5 13l4 4L19 7' : 'M18 9v12m0 0l-4-4m4 4l4-4'" />
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- خطاهای اعتبارسنجی -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">خطاهای اعتبارسنجی</h5>
      </div>
      <div class="p-6">
        <div class="space-y-4">
          <div
            v-for="error in validationErrors"
            :key="error.id"
            class="flex items-start space-x-4 space-x-reverse p-6 border border-gray-200 rounded-lg"
          >
            <div
              class="w-2 h-2 rounded-full mt-2"
              :class="getErrorColor(error.severity)"
            ></div>
            <div class="flex-1">
              <div class="flex items-center justify-between">
                <h6 class="text-sm font-medium text-gray-900">{{ error.field }}</h6>
                <span class="text-xs text-gray-500">{{ error.timestamp }}</span>
              </div>
              <p class="text-sm text-gray-600 mt-1">{{ error.message }}</p>
              <div class="flex items-center space-x-4 space-x-reverse mt-2 text-xs text-gray-500">
                <span>مقدار: {{ error.value }}</span>
                <span>قانون: {{ error.rule }}</span>
                <span>شدت: {{ getSeverityText(error.severity) }}</span>
              </div>
              <div class="mt-2 flex items-center space-x-2 space-x-reverse">
                <button
                  class="text-blue-600 hover:text-blue-800 text-xs"
                  @click="fixError(error)"
                >
                  اصلاح خودکار
                </button>
                <button
                  class="text-gray-600 hover:text-gray-800 text-xs"
                  @click="ignoreError(error)"
                >
                  نادیده گرفتن
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات اعتبارسنجی -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">تنظیمات اعتبارسنجی</h5>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">سطح اعتبارسنجی</label>
          <select
            v-model="validationSettings.level"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="strict">سختگیرانه</option>
            <option value="normal">عادی</option>
            <option value="lenient">آسان‌گیر</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر خطاها</label>
          <input
            v-model="validationSettings.maxErrors"
            type="number"
            min="1"
            max="1000"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">بررسی خودکار</label>
          <select
            v-model="validationSettings.autoCheck"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="disabled">غیرفعال</option>
            <option value="onSave">در زمان ذخیره</option>
            <option value="onChange">در زمان تغییر</option>
            <option value="periodic">دوره‌ای</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">فاصله بررسی دوره‌ای</label>
          <select
            v-model="validationSettings.checkInterval"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="5">۵ دقیقه</option>
            <option value="15">۱۵ دقیقه</option>
            <option value="30">۳۰ دقیقه</option>
            <option value="60">۱ ساعت</option>
          </select>
        </div>
      </div>

      <div class="mt-4 space-y-3">
        <label class="flex items-center">
          <input
            v-model="validationSettings.autoFix"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">اصلاح خودکار خطاها</span>
        </label>

        <label class="flex items-center">
          <input
            v-model="validationSettings.notifyOnError"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">اعلان در صورت خطا</span>
        </label>

        <label class="flex items-center">
          <input
            v-model="validationSettings.stopOnError"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">توقف در صورت خطا</span>
        </label>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// آمار اعتبارسنجی
const validationStats = ref({
  valid: 1245,
  errors: 23,
  warnings: 8,
  total: 1276,
  successRate: 97.6
})

// قوانین اعتبارسنجی
const validationRules = ref([
  {
    id: 1,
    name: 'اعتبارسنجی مبلغ',
    description: 'بررسی صحت مبالغ تراکنش‌ها',
    status: 'active',
    priority: 'بالا',
    errorCount: 5,
    enabled: true
  },
  {
    id: 2,
    name: 'اعتبارسنجی تاریخ',
    description: 'بررسی صحت تاریخ‌های تراکنش',
    status: 'active',
    priority: 'متوسط',
    errorCount: 3,
    enabled: true
  },
  {
    id: 3,
    name: 'اعتبارسنجی شناسه',
    description: 'بررسی یکتا بودن شناسه‌ها',
    status: 'warning',
    priority: 'بالا',
    errorCount: 12,
    enabled: true
  },
  {
    id: 4,
    name: 'اعتبارسنجی فرمت',
    description: 'بررسی فرمت صحیح داده‌ها',
    status: 'error',
    priority: 'متوسط',
    errorCount: 8,
    enabled: false
  }
])

// خطاهای اعتبارسنجی
const validationErrors = ref([
  {
    id: 1,
    field: 'مبلغ تراکنش',
    message: 'مبلغ نمی‌تواند منفی باشد',
    value: '-50000',
    rule: 'اعتبارسنجی مبلغ',
    severity: 'error',
    timestamp: '۲ دقیقه پیش'
  },
  {
    id: 2,
    field: 'تاریخ تراکنش',
    message: 'تاریخ در آینده نمی‌تواند باشد',
    value: '2025-01-15',
    rule: 'اعتبارسنجی تاریخ',
    severity: 'warning',
    timestamp: '۵ دقیقه پیش'
  },
  {
    id: 3,
    field: 'شناسه تراکنش',
    message: 'شناسه تکراری است',
    value: 'TXN123456',
    rule: 'اعتبارسنجی شناسه',
    severity: 'error',
    timestamp: '۱۰ دقیقه پیش'
  }
])

// تنظیمات اعتبارسنجی
const validationSettings = ref({
  level: 'normal',
  maxErrors: 100,
  autoCheck: 'onSave',
  checkInterval: 30,
  autoFix: true,
  notifyOnError: true,
  stopOnError: false
})

// رنگ‌های وضعیت
const getRuleStatusColor = (status: string) => {
  const colors = {
    active: 'bg-green-500',
    warning: 'bg-yellow-500',
    error: 'bg-red-500'
  }
  return colors[status] || 'bg-gray-500'
}

const getErrorColor = (severity: string) => {
  const colors = {
    error: 'bg-red-500',
    warning: 'bg-yellow-500',
    info: 'bg-blue-500'
  }
  return colors[severity] || 'bg-gray-500'
}

const getSeverityText = (severity: string) => {
  const texts = {
    error: 'خطا',
    warning: 'هشدار',
    info: 'اطلاعات'
  }
  return texts[severity] || severity
}

// عملیات
interface Rule {
  id: number | string
  enabled?: boolean
  [key: string]: unknown
}

interface ErrorItem {
  id: number | string
  [key: string]: unknown
}

const runValidation = () => {

}

const editRule = (_rule: Rule) => {

}

const toggleRule = (rule: Rule) => {
  rule.enabled = !rule.enabled

}

const fixError = (_error: ErrorItem) => {

}

const ignoreError = (_error: ErrorItem) => {

}
</script>

<!--
  کامپوننت اعتبارسنجی داده‌ها
  شامل:
  1. آمار اعتبارسنجی
  2. قوانین اعتبارسنجی
  3. خطاهای اعتبارسنجی
  4. تنظیمات اعتبارسنجی
  5. طراحی مدرن و کاملاً ریسپانسیو
--> 
