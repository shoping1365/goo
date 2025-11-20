<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between">
        <div>
          <h2 class="text-xl font-bold text-gray-900">مدیریت خطاهای حسابداری</h2>
          <p class="text-gray-600 mt-1">نظارت، تشخیص و رفع خطاهای حسابداری کیف پول</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'errors'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
            @click="activeTab = 'errors'"
          >
            خطاها
          </button>
          <button
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'reconciliation'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
            @click="activeTab = 'reconciliation'"
          >
            تطبیق حساب
          </button>
          <button
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'validation'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
            @click="activeTab = 'validation'"
          >
            اعتبارسنجی
          </button>
          <button
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'reports'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
            @click="activeTab = 'reports'"
          >
            گزارش‌ها
          </button>
        </div>
      </div>
    </div>

    <!-- محتوای تب‌ها -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200">
      <!-- خطاها -->
      <div v-if="activeTab === 'errors'" class="p-6">
        <div class="space-y-6">
          <!-- آمار خطاها -->
          <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
            <div class="bg-white border border-gray-200 rounded-lg p-6">
              <div class="flex items-center">
                <div class="w-8 h-8 bg-red-100 rounded-full flex items-center justify-center">
                  <svg class="w-4 h-4 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z" />
                  </svg>
                </div>
                <div class="mr-3">
                  <p class="text-sm font-medium text-gray-900">خطاهای بحرانی</p>
                  <p class="text-lg font-bold text-red-600">{{ errorStats.critical }}</p>
                </div>
              </div>
            </div>
            <div class="bg-white border border-gray-200 rounded-lg p-6">
              <div class="flex items-center">
                <div class="w-8 h-8 bg-yellow-100 rounded-full flex items-center justify-center">
                  <svg class="w-4 h-4 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                </div>
                <div class="mr-3">
                  <p class="text-sm font-medium text-gray-900">خطاهای هشدار</p>
                  <p class="text-lg font-bold text-yellow-600">{{ errorStats.warning }}</p>
                </div>
              </div>
            </div>
            <div class="bg-white border border-gray-200 rounded-lg p-6">
              <div class="flex items-center">
                <div class="w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center">
                  <svg class="w-4 h-4 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                </div>
                <div class="mr-3">
                  <p class="text-sm font-medium text-gray-900">رفع شده</p>
                  <p class="text-lg font-bold text-blue-600">{{ errorStats.resolved }}</p>
                </div>
              </div>
            </div>
            <div class="bg-white border border-gray-200 rounded-lg p-6">
              <div class="flex items-center">
                <div class="w-8 h-8 bg-green-100 rounded-full flex items-center justify-center">
                  <svg class="w-4 h-4 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
                  </svg>
                </div>
                <div class="mr-3">
                  <p class="text-sm font-medium text-gray-900">نرخ رفع</p>
                  <p class="text-lg font-bold text-green-600">{{ errorStats.resolutionRate }}%</p>
                </div>
              </div>
            </div>
          </div>

          <!-- لیست خطاها -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">خطاهای اخیر</h3>
            <div class="space-y-3">
              <div v-for="error in recentErrors" :key="error.id" class="border border-gray-200 rounded-lg p-6">
                <div class="flex items-center justify-between mb-3">
                  <div class="flex items-center space-x-3 space-x-reverse">
                    <div :class="`w-4 h-4 rounded-full ${getErrorTypeClass(error.type)}`"></div>
                    <div>
                      <h4 class="font-medium text-gray-900">{{ error.title }}</h4>
                      <p class="text-sm text-gray-600">{{ error.description }}</p>
                    </div>
                  </div>
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <span :class="`px-2 py-1 text-xs rounded-full ${getErrorStatusClass(error.status)}`">
                      {{ getErrorStatusText(error.status) }}
                    </span>
                    <span class="text-xs text-gray-500">{{ error.time }}</span>
                  </div>
                </div>
                <div class="grid grid-cols-2 md:grid-cols-4 gap-6 text-sm">
                  <div>
                    <span class="text-gray-600">شناسه تراکنش:</span>
                    <span class="text-gray-900">{{ error.transactionId }}</span>
                  </div>
                  <div>
                    <span class="text-gray-600">مبلغ:</span>
                    <span class="text-gray-900">{{ formatCurrency(error.amount) }}</span>
                  </div>
                  <div>
                    <span class="text-gray-600">کاربر:</span>
                    <span class="text-gray-900">{{ error.user }}</span>
                  </div>
                  <div>
                    <span class="text-gray-600">اولویت:</span>
                    <span :class="`${getPriorityClass(error.priority)}`">{{ getPriorityText(error.priority) }}</span>
                  </div>
                </div>
                <div class="mt-3 flex space-x-2 space-x-reverse">
                  <button class="bg-blue-600 text-white px-3 py-2 rounded text-sm hover:bg-blue-700 transition-colors">
                    مشاهده جزئیات
                  </button>
                  <button class="bg-green-600 text-white px-3 py-2 rounded text-sm hover:bg-green-700 transition-colors">
                    رفع خطا
                  </button>
                  <button class="bg-gray-600 text-white px-3 py-2 rounded text-sm hover:bg-gray-700 transition-colors">
                    نادیده گرفتن
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- تطبیق حساب -->
      <div v-if="activeTab === 'reconciliation'" class="p-6">
        <div class="space-y-6">
          <!-- وضعیت تطبیق -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">وضعیت تطبیق حساب</h3>
            <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
              <div class="border border-green-200 rounded-lg p-6 bg-green-50">
                <div class="flex items-center mb-2">
                  <div class="w-4 h-4 bg-green-500 rounded-full ml-2"></div>
                  <h4 class="font-medium text-green-800">تطبیق شده</h4>
                </div>
                <p class="text-2xl font-bold text-green-600">{{ reconciliationStats.matched }}</p>
                <p class="text-sm text-green-700">تراکنش تطبیق شده</p>
              </div>
              <div class="border border-yellow-200 rounded-lg p-6 bg-yellow-50">
                <div class="flex items-center mb-2">
                  <div class="w-4 h-4 bg-yellow-500 rounded-full ml-2"></div>
                  <h4 class="font-medium text-yellow-800">نیاز به بررسی</h4>
                </div>
                <p class="text-2xl font-bold text-yellow-600">{{ reconciliationStats.pending }}</p>
                <p class="text-sm text-yellow-700">تراکنش نیازمند بررسی</p>
              </div>
              <div class="border border-red-200 rounded-lg p-6 bg-red-50">
                <div class="flex items-center mb-2">
                  <div class="w-4 h-4 bg-red-500 rounded-full ml-2"></div>
                  <h4 class="font-medium text-red-800">عدم تطبیق</h4>
                </div>
                <p class="text-2xl font-bold text-red-600">{{ reconciliationStats.unmatched }}</p>
                <p class="text-sm text-red-700">تراکنش عدم تطبیق</p>
              </div>
            </div>
          </div>

          <!-- تراکنش‌های عدم تطبیق -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">تراکنش‌های عدم تطبیق</h3>
            <div class="overflow-x-auto">
              <table class="min-w-full divide-y divide-gray-200">
                <thead class="bg-gray-50">
                  <tr>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">شناسه</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مبلغ سیستم</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مبلغ بانک</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تفاوت</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
                  </tr>
                </thead>
                <tbody class="bg-white divide-y divide-gray-200">
                  <tr v-for="transaction in unmatchedTransactions" :key="transaction.id">
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ transaction.id }}</td>
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ transaction.date }}</td>
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ formatCurrency(transaction.systemAmount) }}</td>
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ formatCurrency(transaction.bankAmount) }}</td>
                    <td class="px-6 py-4 whitespace-nowrap">
                      <span :class="`text-sm ${transaction.difference >= 0 ? 'text-red-600' : 'text-green-600'}`">
                        {{ formatCurrency(transaction.difference) }}
                      </span>
                    </td>
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                      <button class="text-blue-600 hover:text-blue-800 ml-2">بررسی</button>
                      <button class="text-green-600 hover:text-green-800">تطبیق</button>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>

      <!-- اعتبارسنجی -->
      <div v-if="activeTab === 'validation'" class="p-6">
        <div class="space-y-6">
          <!-- قوانین اعتبارسنجی -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">قوانین اعتبارسنجی</h3>
            <div class="space-y-4">
              <div v-for="rule in validationRules" :key="rule.id" class="flex items-center justify-between p-6 border border-gray-200 rounded-lg">
                <div class="flex items-center space-x-3 space-x-reverse">
                  <div :class="`w-8 h-8 rounded-full flex items-center justify-center ${rule.active ? 'bg-green-100' : 'bg-gray-100'}`">
                    <svg class="w-4 h-4" :class="rule.active ? 'text-green-600' : 'text-gray-400'" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                  </div>
                  <div>
                    <h4 class="font-medium text-gray-900">{{ rule.name }}</h4>
                    <p class="text-sm text-gray-600">{{ rule.description }}</p>
                  </div>
                </div>
                <div class="flex items-center space-x-2 space-x-reverse">
                  <label class="relative inline-flex items-center cursor-pointer">
                    <input v-model="rule.active" type="checkbox" class="sr-only peer">
                    <div class="w-9 h-5 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-4 after:w-4 after:transition-all peer-checked:bg-blue-600"></div>
                  </label>
                  <button class="text-blue-600 hover:text-blue-800 text-sm">ویرایش</button>
                </div>
              </div>
            </div>
          </div>

          <!-- نتایج اعتبارسنجی -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">نتایج اعتبارسنجی اخیر</h3>
            <div class="space-y-3">
              <div v-for="result in validationResults" :key="result.id" class="flex items-center justify-between p-6 bg-gray-50 rounded-lg">
                <div class="flex items-center space-x-3 space-x-reverse">
                  <div :class="`w-8 h-8 rounded-full flex items-center justify-center ${result.status === 'passed' ? 'bg-green-100' : 'bg-red-100'}`">
                    <svg class="w-4 h-4" :class="result.status === 'passed' ? 'text-green-600' : 'text-red-600'" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                  </div>
                  <div>
                    <h4 class="font-medium text-gray-900">{{ result.ruleName }}</h4>
                    <p class="text-sm text-gray-600">{{ result.description }}</p>
                  </div>
                </div>
                <div class="text-sm text-gray-500">{{ result.time }}</div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- گزارش‌ها -->
      <div v-if="activeTab === 'reports'" class="p-6">
        <div class="space-y-6">
          <!-- گزارش‌های آماده -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">گزارش‌های آماده</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div v-for="report in availableReports" :key="report.id" class="border border-gray-200 rounded-lg p-6">
                <h4 class="font-medium text-gray-900 mb-2">{{ report.name }}</h4>
                <p class="text-sm text-gray-600 mb-3">{{ report.description }}</p>
                <div class="flex items-center justify-between">
                  <span class="text-xs text-gray-500">{{ report.lastGenerated }}</span>
                  <button class="bg-blue-600 text-white px-3 py-1 rounded text-sm hover:bg-blue-700 transition-colors">
                    ایجاد گزارش
                  </button>
                </div>
              </div>
            </div>
          </div>

          <!-- گزارش سفارشی -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">گزارش سفارشی</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <h4 class="font-medium text-gray-900 mb-3">پارامترهای گزارش</h4>
                <div class="space-y-3">
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">نوع خطا</label>
                    <select v-model="customReport.errorType" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                      <option value="">همه</option>
                      <option value="critical">بحرانی</option>
                      <option value="warning">هشدار</option>
                      <option value="info">اطلاعات</option>
                    </select>
                  </div>
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">بازه زمانی</label>
                    <select v-model="customReport.timeRange" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                      <option value="today">امروز</option>
                      <option value="week">هفته گذشته</option>
                      <option value="month">ماه گذشته</option>
                      <option value="custom">سفارشی</option>
                    </select>
                  </div>
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">فرمت گزارش</label>
                    <select v-model="customReport.format" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                      <option value="pdf">PDF</option>
                      <option value="excel">Excel</option>
                      <option value="csv">CSV</option>
                    </select>
                  </div>
                </div>
              </div>
              <div>
                <h4 class="font-medium text-gray-900 mb-3">تنظیمات اضافی</h4>
                <div class="space-y-3">
                  <div class="flex items-center">
                    <input v-model="customReport.includeDetails" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                    <label class="mr-2 text-sm text-gray-700">شامل جزئیات کامل</label>
                  </div>
                  <div class="flex items-center">
                    <input v-model="customReport.includeCharts" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                    <label class="mr-2 text-sm text-gray-700">شامل نمودارها</label>
                  </div>
                  <div class="flex items-center">
                    <input v-model="customReport.sendEmail" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                    <label class="mr-2 text-sm text-gray-700">ارسال به ایمیل</label>
                  </div>
                  <div v-if="customReport.sendEmail">
                    <label class="block text-sm font-medium text-gray-700 mb-1">آدرس ایمیل</label>
                    <input v-model="customReport.email" type="email" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="example@domain.com">
                  </div>
                </div>
              </div>
            </div>
            <div class="mt-6">
              <button class="bg-green-600 text-white px-6 py-2 rounded-lg hover:bg-green-700 transition-colors">
                ایجاد گزارش سفارشی
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// تعریف reactive data
const activeTab = ref('errors')

// آمار خطاها
const errorStats = ref({
  critical: 5,
  warning: 23,
  resolved: 145,
  resolutionRate: 96.5
})

// خطاهای اخیر
const recentErrors = ref([
  {
    id: 1,
    title: 'عدم تطبیق مبلغ تراکنش',
    description: 'مبلغ تراکنش با موجودی کیف پول مطابقت ندارد',
    type: 'critical',
    status: 'open',
    time: '2 ساعت پیش',
    transactionId: 'TXN-001',
    amount: 5000000,
    user: 'user123',
    priority: 'high'
  },
  {
    id: 2,
    title: 'خطای اتصال به بانک',
    description: 'عدم امکان اتصال به درگاه بانکی',
    type: 'warning',
    status: 'in_progress',
    time: '4 ساعت پیش',
    transactionId: 'TXN-002',
    amount: 2000000,
    user: 'user456',
    priority: 'medium'
  },
  {
    id: 3,
    title: 'تراکنش تکراری',
    description: 'تراکنش مشابه قبلاً ثبت شده است',
    type: 'warning',
    status: 'resolved',
    time: '6 ساعت پیش',
    transactionId: 'TXN-003',
    amount: 1000000,
    user: 'user789',
    priority: 'low'
  }
])

// آمار تطبیق حساب
const reconciliationStats = ref({
  matched: 15420,
  pending: 45,
  unmatched: 12
})

// تراکنش‌های عدم تطبیق
const unmatchedTransactions = ref([
  {
    id: 'TXN-001',
    date: '2024-01-31',
    systemAmount: 5000000,
    bankAmount: 5000000,
    difference: 0
  },
  {
    id: 'TXN-002',
    date: '2024-01-31',
    systemAmount: 2000000,
    bankAmount: 2000000,
    difference: 0
  },
  {
    id: 'TXN-003',
    date: '2024-01-31',
    systemAmount: 1000000,
    bankAmount: 1000000,
    difference: 0
  }
])

// قوانین اعتبارسنجی
const validationRules = ref([
  {
    id: 1,
    name: 'تطبیق مبلغ تراکنش',
    description: 'مبلغ تراکنش باید با موجودی کیف پول مطابقت داشته باشد',
    active: true
  },
  {
    id: 2,
    name: 'بررسی تراکنش تکراری',
    description: 'جلوگیری از ثبت تراکنش‌های تکراری',
    active: true
  },
  {
    id: 3,
    name: 'اعتبارسنجی کاربر',
    description: 'بررسی اعتبار و وضعیت کاربر',
    active: true
  },
  {
    id: 4,
    name: 'بررسی محدودیت‌های زمانی',
    description: 'اعمال محدودیت‌های زمانی برای تراکنش‌ها',
    active: false
  }
])

// نتایج اعتبارسنجی
const validationResults = ref([
  {
    id: 1,
    ruleName: 'تطبیق مبلغ تراکنش',
    description: 'تمام تراکنش‌ها با موفقیت اعتبارسنجی شدند',
    status: 'passed',
    time: '2024-01-31 14:30'
  },
  {
    id: 2,
    ruleName: 'بررسی تراکنش تکراری',
    description: '2 تراکنش تکراری شناسایی و مسدود شد',
    status: 'failed',
    time: '2024-01-31 14:25'
  }
])

// گزارش‌های آماده
const availableReports = ref([
  {
    id: 1,
    name: 'گزارش خطاهای روزانه',
    description: 'گزارش کامل خطاهای رخ داده در روز جاری',
    lastGenerated: '2024-01-31 14:00'
  },
  {
    id: 2,
    name: 'گزارش تطبیق حساب',
    description: 'گزارش تطبیق حساب‌ها با سیستم‌های بانکی',
    lastGenerated: '2024-01-31 13:30'
  },
  {
    id: 3,
    name: 'گزارش اعتبارسنجی',
    description: 'نتایج اعتبارسنجی تراکنش‌ها',
    lastGenerated: '2024-01-31 12:00'
  },
  {
    id: 4,
    name: 'گزارش عملکرد سیستم',
    description: 'آمار عملکرد و کارایی سیستم',
    lastGenerated: '2024-01-31 11:00'
  }
])

// گزارش سفارشی
const customReport = ref({
  errorType: '',
  timeRange: 'week',
  format: 'pdf',
  includeDetails: true,
  includeCharts: false,
  sendEmail: false,
  email: ''
})

// توابع کمکی
const getErrorTypeClass = (type: string) => {
  switch (type) {
    case 'critical':
      return 'bg-red-500'
    case 'warning':
      return 'bg-yellow-500'
    case 'info':
      return 'bg-blue-500'
    default:
      return 'bg-gray-500'
  }
}

const getErrorStatusClass = (status: string) => {
  switch (status) {
    case 'open':
      return 'bg-red-100 text-red-800'
    case 'in_progress':
      return 'bg-yellow-100 text-yellow-800'
    case 'resolved':
      return 'bg-green-100 text-green-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getErrorStatusText = (status: string) => {
  switch (status) {
    case 'open':
      return 'باز'
    case 'in_progress':
      return 'در حال بررسی'
    case 'resolved':
      return 'رفع شده'
    default:
      return 'نامشخص'
  }
}

const getPriorityClass = (priority: string) => {
  switch (priority) {
    case 'high':
      return 'text-red-600'
    case 'medium':
      return 'text-yellow-600'
    case 'low':
      return 'text-green-600'
    default:
      return 'text-gray-600'
  }
}

const getPriorityText = (priority: string) => {
  switch (priority) {
    case 'high':
      return 'زیاد'
    case 'medium':
      return 'متوسط'
    case 'low':
      return 'کم'
    default:
      return 'نامشخص'
  }
}

// تابع فرمت‌بندی ارز
const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR', {
    style: 'currency',
    currency: 'IRR',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(amount)
}
</script> 
