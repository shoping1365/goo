<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between">
        <div>
          <h2 class="text-xl font-bold text-gray-900">کنترل‌های داخلی</h2>
          <p class="text-gray-600 mt-1">نظارت، کنترل و مدیریت ریسک‌های کیف پول</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'monitoring'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
            @click="activeTab = 'monitoring'"
          >
            نظارت
          </button>
          <button
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'risk'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
            @click="activeTab = 'risk'"
          >
            مدیریت ریسک
          </button>
          <button
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'compliance'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
            @click="activeTab = 'compliance'"
          >
            انطباق
          </button>
          <button
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'audit'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
            @click="activeTab = 'audit'"
          >
            حسابرسی
          </button>
        </div>
      </div>
    </div>

    <!-- محتوای تب‌ها -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200">
      <!-- نظارت -->
      <div v-if="activeTab === 'monitoring'" class="p-6">
        <div class="space-y-6">
          <!-- داشبورد نظارت -->
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
            <div class="bg-white border border-gray-200 rounded-lg p-6">
              <div class="flex items-center">
                <div class="w-8 h-8 bg-green-100 rounded-full flex items-center justify-center">
                  <svg class="w-4 h-4 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                </div>
                <div class="mr-3">
                  <p class="text-sm font-medium text-gray-900">تراکنش‌های عادی</p>
                  <p class="text-lg font-bold text-green-600">{{ monitoringStats.normalTransactions }}</p>
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
                  <p class="text-sm font-medium text-gray-900">تراکنش‌های مشکوک</p>
                  <p class="text-lg font-bold text-yellow-600">{{ monitoringStats.suspiciousTransactions }}</p>
                </div>
              </div>
            </div>
            <div class="bg-white border border-gray-200 rounded-lg p-6">
              <div class="flex items-center">
                <div class="w-8 h-8 bg-red-100 rounded-full flex items-center justify-center">
                  <svg class="w-4 h-4 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z" />
                  </svg>
                </div>
                <div class="mr-3">
                  <p class="text-sm font-medium text-gray-900">تراکنش‌های مسدود شده</p>
                  <p class="text-lg font-bold text-red-600">{{ monitoringStats.blockedTransactions }}</p>
                </div>
              </div>
            </div>
            <div class="bg-white border border-gray-200 rounded-lg p-6">
              <div class="flex items-center">
                <div class="w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center">
                  <svg class="w-4 h-4 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                  </svg>
                </div>
                <div class="mr-3">
                  <p class="text-sm font-medium text-gray-900">نظارت فعال</p>
                  <p class="text-lg font-bold text-blue-600">{{ monitoringStats.activeMonitoring }}</p>
                </div>
              </div>
            </div>
          </div>

          <!-- هشدارهای فعال -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">هشدارهای فعال</h3>
            <div class="space-y-3">
              <div v-for="alert in activeAlerts" :key="alert.id" class="flex items-center justify-between p-6 bg-red-50 border border-red-200 rounded-lg">
                <div class="flex items-center space-x-3 space-x-reverse">
                  <div class="w-4 h-4 bg-red-500 rounded-full"></div>
                  <div>
                    <p class="font-medium text-red-800">{{ alert.title }}</p>
                    <p class="text-sm text-red-600">{{ alert.description }}</p>
                  </div>
                </div>
                <div class="flex items-center space-x-2 space-x-reverse">
                  <span class="text-xs text-red-600">{{ alert.time }}</span>
                  <button class="text-red-600 hover:text-red-800">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                    </svg>
                  </button>
                </div>
              </div>
            </div>
          </div>

          <!-- تراکنش‌های تحت نظارت -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">تراکنش‌های تحت نظارت</h3>
            <div class="overflow-x-auto">
              <table class="min-w-full divide-y divide-gray-200">
                <thead class="bg-gray-50">
                  <tr>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">شناسه</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مبلغ</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
                  </tr>
                </thead>
                <tbody class="bg-white divide-y divide-gray-200">
                  <tr v-for="transaction in monitoredTransactions" :key="transaction.id">
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ transaction.id }}</td>
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ transaction.type }}</td>
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ transaction.amount.toLocaleString() }} تومان</td>
                    <td class="px-6 py-4 whitespace-nowrap">
                      <span :class="`px-2 py-1 text-xs rounded-full ${getStatusClass(transaction.status)}`">
                        {{ transaction.status }}
                      </span>
                    </td>
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                      <button class="text-blue-600 hover:text-blue-800 ml-2">مشاهده</button>
                      <button class="text-green-600 hover:text-green-800 ml-2">تایید</button>
                      <button class="text-red-600 hover:text-red-800">رد</button>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>

      <!-- مدیریت ریسک -->
      <div v-if="activeTab === 'risk'" class="p-6">
        <div class="space-y-6">
          <!-- سطوح ریسک -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">سطوح ریسک</h3>
            <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
              <div class="border border-green-200 rounded-lg p-6 bg-green-50">
                <div class="flex items-center mb-2">
                  <div class="w-4 h-4 bg-green-500 rounded-full ml-2"></div>
                  <h4 class="font-medium text-green-800">ریسک کم</h4>
                </div>
                <p class="text-sm text-green-700 mb-3">تراکنش‌های عادی و قابل اعتماد</p>
                <div class="text-xs text-green-600">
                  <p>• مبلغ کمتر از 1 میلیون تومان</p>
                  <p>• کاربران تایید شده</p>
                  <p>• تراکنش‌های مکرر</p>
                </div>
              </div>
              <div class="border border-yellow-200 rounded-lg p-6 bg-yellow-50">
                <div class="flex items-center mb-2">
                  <div class="w-4 h-4 bg-yellow-500 rounded-full ml-2"></div>
                  <h4 class="font-medium text-yellow-800">ریسک متوسط</h4>
                </div>
                <p class="text-sm text-yellow-700 mb-3">تراکنش‌های نیازمند بررسی بیشتر</p>
                <div class="text-xs text-yellow-600">
                  <p>• مبلغ 1 تا 10 میلیون تومان</p>
                  <p>• کاربران جدید</p>
                  <p>• تراکنش‌های غیرمعمول</p>
                </div>
              </div>
              <div class="border border-red-200 rounded-lg p-6 bg-red-50">
                <div class="flex items-center mb-2">
                  <div class="w-4 h-4 bg-red-500 rounded-full ml-2"></div>
                  <h4 class="font-medium text-red-800">ریسک بالا</h4>
                </div>
                <p class="text-sm text-red-700 mb-3">تراکنش‌های مشکوک و پرریسک</p>
                <div class="text-xs text-red-600">
                  <p>• مبلغ بیشتر از 10 میلیون تومان</p>
                  <p>• کاربران مشکوک</p>
                  <p>• الگوهای غیرعادی</p>
                </div>
              </div>
            </div>
          </div>

          <!-- قوانین ریسک -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">قوانین مدیریت ریسک</h3>
            <div class="space-y-4">
              <div v-for="rule in riskRules" :key="rule.id" class="flex items-center justify-between p-6 border border-gray-200 rounded-lg">
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
        </div>
      </div>

      <!-- انطباق -->
      <div v-if="activeTab === 'compliance'" class="p-6">
        <div class="space-y-6">
          <!-- استانداردهای انطباق -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">استانداردهای انطباق</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div v-for="standard in complianceStandards" :key="standard.id" class="border border-gray-200 rounded-lg p-6">
                <div class="flex items-center justify-between mb-3">
                  <h4 class="font-medium text-gray-900">{{ standard.name }}</h4>
                  <span :class="`px-2 py-1 text-xs rounded-full ${standard.status === 'compliant' ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'}`">
                    {{ standard.status === 'compliant' ? 'مطابق' : 'نامطابق' }}
                  </span>
                </div>
                <p class="text-sm text-gray-600 mb-3">{{ standard.description }}</p>
                <div class="flex items-center justify-between">
                  <span class="text-xs text-gray-500">آخرین بررسی: {{ standard.lastCheck }}</span>
                  <button class="text-blue-600 hover:text-blue-800 text-sm">جزئیات</button>
                </div>
              </div>
            </div>
          </div>

          <!-- گزارش‌های انطباق -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">گزارش‌های انطباق</h3>
            <div class="space-y-3">
              <div v-for="report in complianceReports" :key="report.id" class="flex items-center justify-between p-6 border border-gray-200 rounded-lg">
                <div>
                  <h4 class="font-medium text-gray-900">{{ report.title }}</h4>
                  <p class="text-sm text-gray-600">{{ report.description }}</p>
                </div>
                <div class="flex items-center space-x-2 space-x-reverse">
                  <span class="text-xs text-gray-500">{{ report.date }}</span>
                  <button class="bg-blue-600 text-white px-3 py-1 rounded text-sm hover:bg-blue-700">
                    دانلود
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- حسابرسی -->
      <div v-if="activeTab === 'audit'" class="p-6">
        <div class="space-y-6">
          <!-- لاگ‌های حسابرسی -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">لاگ‌های حسابرسی</h3>
            <div class="space-y-3">
              <div v-for="log in auditLogs" :key="log.id" class="flex items-center space-x-3 space-x-reverse p-6 bg-gray-50 rounded-lg">
                <div :class="`w-8 h-8 rounded-full flex items-center justify-center ${getLogTypeClass(log.type)}`">
                  <svg class="w-4 h-4" :class="getLogTypeIcon(log.type)" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                </div>
                <div class="flex-1">
                  <p class="text-sm font-medium text-gray-900">{{ log.action }}</p>
                  <p class="text-xs text-gray-500">{{ log.user }} - {{ log.time }}</p>
                </div>
                <div class="text-xs text-gray-500">{{ log.ip }}</div>
              </div>
            </div>
          </div>

          <!-- گزارش‌های حسابرسی -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">گزارش‌های حسابرسی</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div v-for="report in auditReports" :key="report.id" class="border border-gray-200 rounded-lg p-6">
                <h4 class="font-medium text-gray-900 mb-2">{{ report.title }}</h4>
                <p class="text-sm text-gray-600 mb-3">{{ report.description }}</p>
                <div class="flex items-center justify-between">
                  <span class="text-xs text-gray-500">{{ report.period }}</span>
                  <button class="bg-green-600 text-white px-3 py-1 rounded text-sm hover:bg-green-700">
                    ایجاد گزارش
                  </button>
                </div>
              </div>
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
const activeTab = ref('monitoring')

// آمار نظارت
const monitoringStats = ref({
  normalTransactions: 1247,
  suspiciousTransactions: 23,
  blockedTransactions: 5,
  activeMonitoring: 18
})

// هشدارهای فعال
const activeAlerts = ref([
  {
    id: 1,
    title: 'تراکنش مشکوک شناسایی شد',
    description: 'تراکنش با مبلغ 15 میلیون تومان از کاربر جدید',
    time: '5 دقیقه پیش'
  },
  {
    id: 2,
    title: 'تلاش ورود ناموفق',
    description: '5 تلاش ورود ناموفق از IP مشکوک',
    time: '10 دقیقه پیش'
  }
])

// تراکنش‌های تحت نظارت
const monitoredTransactions = ref([
  {
    id: 'TXN-001',
    type: 'شارژ',
    amount: 5000000,
    status: 'مشکوک'
  },
  {
    id: 'TXN-002',
    type: 'برداشت',
    amount: 2000000,
    status: 'در انتظار تایید'
  },
  {
    id: 'TXN-003',
    type: 'انتقال',
    amount: 8000000,
    status: 'مسدود شده'
  }
])

// قوانین ریسک
const riskRules = ref([
  {
    id: 1,
    name: 'محدودیت مبلغ تراکنش',
    description: 'تراکنش‌های بالای 10 میلیون تومان نیاز به تایید دارند',
    active: true
  },
  {
    id: 2,
    name: 'تشخیص تراکنش‌های غیرعادی',
    description: 'تشخیص خودکار تراکنش‌های خارج از الگوی معمول',
    active: true
  },
  {
    id: 3,
    name: 'محدودیت تعداد تراکنش',
    description: 'حداکثر 10 تراکنش در روز برای هر کاربر',
    active: false
  }
])

// استانداردهای انطباق
const complianceStandards = ref([
  {
    id: 1,
    name: 'PCI DSS',
    description: 'استاندارد امنیت داده‌های کارت‌های پرداخت',
    status: 'compliant',
    lastCheck: '2024-01-15'
  },
  {
    id: 2,
    name: 'ISO 27001',
    description: 'مدیریت امنیت اطلاعات',
    status: 'compliant',
    lastCheck: '2024-01-10'
  },
  {
    id: 3,
    name: 'GDPR',
    description: 'مقررات حفاظت از داده‌های عمومی',
    status: 'non-compliant',
    lastCheck: '2024-01-20'
  }
])

// گزارش‌های انطباق
const complianceReports = ref([
  {
    id: 1,
    title: 'گزارش انطباق ماهانه',
    description: 'گزارش کامل انطباق با استانداردهای امنیتی',
    date: '2024-01-31'
  },
  {
    id: 2,
    title: 'گزارش حسابرسی امنیتی',
    description: 'نتایج حسابرسی امنیتی سیستم کیف پول',
    date: '2024-01-25'
  }
])

// لاگ‌های حسابرسی
const auditLogs = ref([
  {
    id: 1,
    type: 'security',
    action: 'تغییر تنظیمات امنیتی',
    user: 'admin',
    time: '2024-01-31 14:30',
    ip: '192.168.1.100'
  },
  {
    id: 2,
    type: 'transaction',
    action: 'تایید تراکنش مشکوک',
    user: 'operator1',
    time: '2024-01-31 14:25',
    ip: '192.168.1.101'
  },
  {
    id: 3,
    type: 'user',
    action: 'افزودن کاربر جدید',
    user: 'admin',
    time: '2024-01-31 14:20',
    ip: '192.168.1.100'
  }
])

// گزارش‌های حسابرسی
const auditReports = ref([
  {
    id: 1,
    title: 'گزارش فعالیت کاربران',
    description: 'گزارش کامل فعالیت‌های کاربران در سیستم',
    period: 'ماه گذشته'
  },
  {
    id: 2,
    title: 'گزارش تراکنش‌های مالی',
    description: 'گزارش تفصیلی تمام تراکنش‌های مالی',
    period: 'هفته گذشته'
  }
])

// توابع کمکی
const getStatusClass = (status: string) => {
  switch (status) {
    case 'مشکوک':
      return 'bg-yellow-100 text-yellow-800'
    case 'در انتظار تایید':
      return 'bg-blue-100 text-blue-800'
    case 'مسدود شده':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getLogTypeClass = (type: string) => {
  switch (type) {
    case 'security':
      return 'bg-red-100'
    case 'transaction':
      return 'bg-blue-100'
    case 'user':
      return 'bg-green-100'
    default:
      return 'bg-gray-100'
  }
}

const getLogTypeIcon = (type: string) => {
  switch (type) {
    case 'security':
      return 'text-red-600'
    case 'transaction':
      return 'text-blue-600'
    case 'user':
      return 'text-green-600'
    default:
      return 'text-gray-600'
  }
}
</script> 
