<template>
  <div class="integration-manager bg-white rounded-2xl shadow-xl border border-gray-100 p-8">
    <div class="flex items-center mb-6">
      <div class="p-3 bg-gradient-to-r from-violet-400 to-violet-600 rounded-xl shadow-lg">
        <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.367 2.684 3 3 0 00-5.367-2.684z"></path>
        </svg>
      </div>
      <div class="mr-4">
        <h2 class="text-2xl font-bold text-gray-900">مدیریت یکپارچه‌سازی</h2>
        <p class="text-gray-600 mt-1">اتصال به سرویس‌های خارجی و API ها</p>
      </div>
    </div>

    <!-- Integration Status -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
      <div class="bg-gradient-to-r from-violet-50 to-violet-100 rounded-xl p-6 border border-violet-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-violet-600">سرویس‌های فعال</p>
            <p class="text-3xl font-bold text-violet-900">{{ integrationStats.activeServices }}</p>
          </div>
          <div class="p-3 bg-violet-500 rounded-lg">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
            </svg>
          </div>
        </div>
        <div class="mt-4 text-sm text-violet-700">
          از {{ integrationStats.totalServices }} سرویس
        </div>
      </div>

      <div class="bg-gradient-to-r from-green-50 to-green-100 rounded-xl p-6 border border-green-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-green-600">درخواست‌های موفق</p>
            <p class="text-3xl font-bold text-green-900">{{ integrationStats.successfulRequests }}</p>
          </div>
          <div class="p-3 bg-green-500 rounded-lg">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
        <div class="mt-4 text-sm text-green-700">
          {{ integrationStats.successRate }}% نرخ موفقیت
        </div>
      </div>

      <div class="bg-gradient-to-r from-blue-50 to-blue-100 rounded-xl p-6 border border-blue-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-blue-600">داده‌های همگام‌سازی</p>
            <p class="text-3xl font-bold text-blue-900">{{ integrationStats.syncedData }}</p>
          </div>
          <div class="p-3 bg-blue-500 rounded-lg">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
            </svg>
          </div>
        </div>
        <div class="mt-4 text-sm text-blue-700">
          در ۲۴ ساعت گذشته
        </div>
      </div>

      <div class="bg-gradient-to-r from-orange-50 to-orange-100 rounded-xl p-6 border border-orange-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-orange-600">خطاهای اتصال</p>
            <p class="text-3xl font-bold text-orange-900">{{ integrationStats.connectionErrors }}</p>
          </div>
          <div class="p-3 bg-orange-500 rounded-lg">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
        <div class="mt-4 text-sm text-orange-700">
          {{ integrationStats.errorRate }}% نرخ خطا
        </div>
      </div>
    </div>

    <!-- Email Service Integrations -->
    <div class="bg-blue-50 rounded-xl p-6 mb-8">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">سرویس‌های ایمیل</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div v-for="service in emailServices" :key="service.id" class="bg-white rounded-lg p-6 border">
          <div class="flex items-center justify-between mb-3">
            <div class="flex items-center">
              <img :src="service.logo" :alt="service.name" class="w-8 h-8 rounded">
              <h4 class="font-medium text-gray-900 mr-2">{{ service.name }}</h4>
            </div>
            <span :class="getServiceStatusClass(service.status)" class="px-2 py-1 text-xs font-medium rounded-full">
              {{ getServiceStatusText(service.status) }}
            </span>
          </div>
          <p class="text-sm text-gray-600 mb-3">{{ service.description }}</p>
          <div class="space-y-2">
            <div class="flex items-center justify-between text-xs">
              <span class="text-gray-500">API Key:</span>
              <span class="font-mono text-gray-700">{{ service.apiKey ? '***' + service.apiKey.slice(-4) : 'تنظیم نشده' }}</span>
            </div>
            <div class="flex items-center justify-between text-xs">
              <span class="text-gray-500">آخرین ارسال:</span>
              <span class="text-gray-700">{{ service.lastSent || 'هیچ' }}</span>
            </div>
          </div>
          <div class="flex space-x-2 space-x-reverse mt-3">
            <button class="flex-1 px-3 py-1 bg-blue-600 text-white text-sm rounded hover:bg-blue-700" @click="configureService(service)">
              تنظیم
            </button>
            <button class="flex-1 px-3 py-1 bg-gray-600 text-white text-sm rounded hover:bg-gray-700" @click="testService(service)">
              تست
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- SMS Service Integrations -->
    <div class="bg-green-50 rounded-xl p-6 mb-8">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">سرویس‌های پیامک</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div v-for="service in smsServices" :key="service.id" class="bg-white rounded-lg p-6 border">
          <div class="flex items-center justify-between mb-3">
            <div class="flex items-center">
              <img :src="service.logo" :alt="service.name" class="w-8 h-8 rounded">
              <h4 class="font-medium text-gray-900 mr-2">{{ service.name }}</h4>
            </div>
            <span :class="getServiceStatusClass(service.status)" class="px-2 py-1 text-xs font-medium rounded-full">
              {{ getServiceStatusText(service.status) }}
            </span>
          </div>
          <p class="text-sm text-gray-600 mb-3">{{ service.description }}</p>
          <div class="space-y-2">
            <div class="flex items-center justify-between text-xs">
              <span class="text-gray-500">اعتبار باقی‌مانده:</span>
              <span class="font-medium text-gray-700">{{ service.remainingCredit }} تومان</span>
            </div>
            <div class="flex items-center justify-between text-xs">
              <span class="text-gray-500">آخرین ارسال:</span>
              <span class="text-gray-700">{{ service.lastSent || 'هیچ' }}</span>
            </div>
          </div>
          <div class="flex space-x-2 space-x-reverse mt-3">
            <button class="flex-1 px-3 py-1 bg-green-600 text-white text-sm rounded hover:bg-green-700" @click="configureService(service)">
              تنظیم
            </button>
            <button class="flex-1 px-3 py-1 bg-gray-600 text-white text-sm rounded hover:bg-gray-700" @click="testService(service)">
              تست
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- CRM Integrations -->
    <div class="bg-purple-50 rounded-xl p-6 mb-8">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">سیستم‌های CRM</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div v-for="service in crmServices" :key="service.id" class="bg-white rounded-lg p-6 border">
          <div class="flex items-center justify-between mb-3">
            <div class="flex items-center">
              <img :src="service.logo" :alt="service.name" class="w-8 h-8 rounded">
              <h4 class="font-medium text-gray-900 mr-2">{{ service.name }}</h4>
            </div>
            <span :class="getServiceStatusClass(service.status)" class="px-2 py-1 text-xs font-medium rounded-full">
              {{ getServiceStatusText(service.status) }}
            </span>
          </div>
          <p class="text-sm text-gray-600 mb-3">{{ service.description }}</p>
          <div class="space-y-2">
            <div class="flex items-center justify-between text-xs">
              <span class="text-gray-500">مشتریان همگام‌سازی:</span>
              <span class="font-medium text-gray-700">{{ service.syncedCustomers }}</span>
            </div>
            <div class="flex items-center justify-between text-xs">
              <span class="text-gray-500">آخرین همگام‌سازی:</span>
              <span class="text-gray-700">{{ service.lastSync || 'هیچ' }}</span>
            </div>
          </div>
          <div class="flex space-x-2 space-x-reverse mt-3">
            <button class="flex-1 px-3 py-1 bg-purple-600 text-white text-sm rounded hover:bg-purple-700" @click="configureService(service)">
              تنظیم
            </button>
            <button class="flex-1 px-3 py-1 bg-gray-600 text-white text-sm rounded hover:bg-gray-700" @click="syncData(service)">
              همگام‌سازی
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Analytics Integrations -->
    <div class="bg-orange-50 rounded-xl p-6 mb-8">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">سرویس‌های تحلیل</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div v-for="service in analyticsServices" :key="service.id" class="bg-white rounded-lg p-6 border">
          <div class="flex items-center justify-between mb-3">
            <div class="flex items-center">
              <img :src="service.logo" :alt="service.name" class="w-8 h-8 rounded">
              <h4 class="font-medium text-gray-900 mr-2">{{ service.name }}</h4>
            </div>
            <span :class="getServiceStatusClass(service.status)" class="px-2 py-1 text-xs font-medium rounded-full">
              {{ getServiceStatusText(service.status) }}
            </span>
          </div>
          <p class="text-sm text-gray-600 mb-3">{{ service.description }}</p>
          <div class="space-y-2">
            <div class="flex items-center justify-between text-xs">
              <span class="text-gray-500">رویدادهای ارسال:</span>
              <span class="font-medium text-gray-700">{{ service.eventsSent }}</span>
            </div>
            <div class="flex items-center justify-between text-xs">
              <span class="text-gray-500">آخرین ارسال:</span>
              <span class="text-gray-700">{{ service.lastSent || 'هیچ' }}</span>
            </div>
          </div>
          <div class="flex space-x-2 space-x-reverse mt-3">
            <button class="flex-1 px-3 py-1 bg-orange-600 text-white text-sm rounded hover:bg-orange-700" @click="configureService(service)">
              تنظیم
            </button>
            <button class="flex-1 px-3 py-1 bg-gray-600 text-white text-sm rounded hover:bg-gray-700" @click="testService(service)">
              تست
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- API Configuration -->
    <div class="bg-gray-50 rounded-xl p-6 mb-8">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات API</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <h4 class="font-medium text-gray-900 mb-3">کلیدهای API</h4>
          <div class="space-y-3">
            <div v-for="key in apiKeys" :key="key.id" class="flex items-center justify-between p-3 bg-white rounded-lg border">
              <div>
                <span class="text-sm font-medium text-gray-900">{{ key.name }}</span>
                <span class="text-xs text-gray-500 mr-2">{{ key.description }}</span>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <span class="text-xs font-mono text-gray-600">{{ key.key.slice(0, 8) }}...</span>
                <button class="text-blue-600 hover:text-blue-800 text-xs" @click="regenerateKey(key.id)">
                  تجدید
                </button>
              </div>
            </div>
          </div>
        </div>

        <div>
          <h4 class="font-medium text-gray-900 mb-3">تنظیمات Webhook</h4>
          <div class="space-y-3">
            <div v-for="webhook in webhooks" :key="webhook.id" class="p-3 bg-white rounded-lg border">
              <div class="flex items-center justify-between mb-2">
                <span class="text-sm font-medium text-gray-900">{{ webhook.name }}</span>
                <span :class="getWebhookStatusClass(webhook.status)" class="px-2 py-1 text-xs rounded-full">
                  {{ getWebhookStatusText(webhook.status) }}
                </span>
              </div>
              <div class="text-xs text-gray-600 mb-2">{{ webhook.url }}</div>
              <div class="flex space-x-2 space-x-reverse">
                <button class="px-2 py-1 bg-blue-600 text-white text-xs rounded hover:bg-blue-700" @click="testWebhook(webhook.id)">
                  تست
                </button>
                <button class="px-2 py-1 bg-gray-600 text-white text-xs rounded hover:bg-gray-700" @click="editWebhook(webhook.id)">
                  ویرایش
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Integration Logs -->
    <div class="bg-gray-50 rounded-xl p-6 mb-8">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">گزارش‌های یکپارچه‌سازی</h3>
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-gray-100">
            <tr>
              <th class="px-4 py-2 text-right text-xs font-medium text-gray-500">زمان</th>
              <th class="px-4 py-2 text-right text-xs font-medium text-gray-500">سرویس</th>
              <th class="px-4 py-2 text-right text-xs font-medium text-gray-500">عملیات</th>
              <th class="px-4 py-2 text-right text-xs font-medium text-gray-500">وضعیت</th>
              <th class="px-4 py-2 text-right text-xs font-medium text-gray-500">جزئیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="log in integrationLogs" :key="log.id" class="hover:bg-gray-50">
              <td class="px-4 py-2 text-sm text-gray-900">{{ formatTime(log.timestamp) }}</td>
              <td class="px-4 py-2 text-sm text-gray-900">{{ log.service }}</td>
              <td class="px-4 py-2 text-sm text-gray-900">{{ log.action }}</td>
              <td class="px-4 py-2">
                <span :class="getLogStatusClass(log.status)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ getLogStatusText(log.status) }}
                </span>
              </td>
              <td class="px-4 py-2 text-sm text-gray-600">{{ log.details }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Action Buttons -->
    <div class="flex justify-between">
      <button class="px-6 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors" @click="refreshIntegrations">
        بروزرسانی اتصالات
      </button>
      
      <div class="flex space-x-3 space-x-reverse">
        <button class="px-6 py-3 bg-purple-600 text-white rounded-lg hover:bg-purple-700 transition-colors" @click="exportIntegrations">
          خروجی تنظیمات
        </button>
        <button class="px-6 py-3 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors" @click="saveSettings">
          ذخیره تنظیمات
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

interface IntegrationStats {
  activeServices: number
  totalServices: number
  successfulRequests: number
  successRate: number
  syncedData: number
  connectionErrors: number
  errorRate: number
}

interface IntegrationService {
  id: number
  name: string
  description: string
  logo: string
  status: 'active' | 'inactive' | 'error' | 'configuring'
  apiKey?: string
  lastSent?: string
  remainingCredit?: number
  syncedCustomers?: number
  lastSync?: string
  eventsSent?: number
}

interface APIKey {
  id: number
  name: string
  description: string
  key: string
}

interface Webhook {
  id: number
  name: string
  url: string
  status: 'active' | 'inactive' | 'error'
}

interface IntegrationLog {
  id: number
  timestamp: Date
  service: string
  action: string
  status: 'success' | 'error' | 'warning'
  details: string
}

const integrationStats = ref<IntegrationStats>({
  activeServices: 8,
  totalServices: 12,
  successfulRequests: 1247,
  successRate: 98.5,
  syncedData: 2341,
  connectionErrors: 19,
  errorRate: 1.5
})

const emailServices = ref<IntegrationService[]>([
  {
    id: 1,
    name: 'Mailgun',
    description: 'سرویس ارسال ایمیل با قابلیت‌های پیشرفته',
    logo: '/api/mailgun-logo.png',
    status: 'active',
    apiKey: 'key-1234567890abcdef',
    lastSent: '۲ دقیقه پیش'
  },
  {
    id: 2,
    name: 'SendGrid',
    description: 'پلتفرم ارسال ایمیل ابری',
    logo: '/api/sendgrid-logo.png',
    status: 'inactive'
  },
  {
    id: 3,
    name: 'SMTP',
    description: 'ارسال مستقیم از طریق SMTP',
    logo: '/api/smtp-logo.png',
    status: 'configuring'
  }
])

const smsServices = ref<IntegrationService[]>([
  {
    id: 4,
    name: 'کاوه‌نگار',
    description: 'سرویس پیامک ایرانی',
    logo: '/api/kavenegar-logo.png',
    status: 'active',
    remainingCredit: 50000,
    lastSent: '۵ دقیقه پیش'
  },
  {
    id: 5,
    name: 'ملی پیامک',
    description: 'ارائه‌دهنده خدمات پیامک',
    logo: '/api/melipayamak-logo.png',
    status: 'inactive'
  }
])

const crmServices = ref<IntegrationService[]>([
  {
    id: 6,
    name: 'HubSpot',
    description: 'پلتفرم مدیریت ارتباط با مشتری',
    logo: '/api/hubspot-logo.png',
    status: 'active',
    syncedCustomers: 1247,
    lastSync: '۱ ساعت پیش'
  },
  {
    id: 7,
    name: 'Salesforce',
    description: 'سیستم مدیریت فروش',
    logo: '/api/salesforce-logo.png',
    status: 'error'
  }
])

const analyticsServices = ref<IntegrationService[]>([
  {
    id: 8,
    name: 'Google Analytics',
    description: 'تحلیل ترافیک وب‌سایت',
    logo: '/api/ga-logo.png',
    status: 'active',
    eventsSent: 2341,
    lastSent: '۱۰ دقیقه پیش'
  },
  {
    id: 9,
    name: 'Mixpanel',
    description: 'تحلیل رفتار کاربران',
    logo: '/api/mixpanel-logo.png',
    status: 'inactive'
  }
])

const apiKeys = ref<APIKey[]>([
  {
    id: 1,
    name: 'کلید عمومی',
    description: 'برای دسترسی به API عمومی',
    key: 'pk_live_1234567890abcdef'
  },
  {
    id: 2,
    name: 'کلید خصوصی',
    description: 'برای عملیات حساس',
    key: 'sk_live_1234567890abcdef'
  }
])

const webhooks = ref<Webhook[]>([
  {
    id: 1,
    name: 'نظرسنجی جدید',
    url: 'https://webhook.site/123456',
    status: 'active'
  },
  {
    id: 2,
    name: 'پاسخ جدید',
    url: 'https://api.example.com/webhook',
    status: 'error'
  }
])

const integrationLogs = ref<IntegrationLog[]>([
  {
    id: 1,
    timestamp: new Date(Date.now() - 5 * 60 * 1000),
    service: 'Mailgun',
    action: 'ارسال ایمیل',
    status: 'success',
    details: 'ایمیل نظرسنجی ارسال شد'
  },
  {
    id: 2,
    timestamp: new Date(Date.now() - 15 * 60 * 1000),
    service: 'کاوه‌نگار',
    action: 'ارسال پیامک',
    status: 'success',
    details: 'پیامک نظرسنجی ارسال شد'
  },
  {
    id: 3,
    timestamp: new Date(Date.now() - 30 * 60 * 1000),
    service: 'Salesforce',
    action: 'همگام‌سازی',
    status: 'error',
    details: 'خطا در اتصال به API'
  }
])

// Methods
const getServiceStatusClass = (status: string) => {
  switch (status) {
    case 'active':
      return 'bg-green-100 text-green-800'
    case 'inactive':
      return 'bg-gray-100 text-gray-800'
    case 'error':
      return 'bg-red-100 text-red-800'
    case 'configuring':
      return 'bg-yellow-100 text-yellow-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getServiceStatusText = (status: string) => {
  switch (status) {
    case 'active':
      return 'فعال'
    case 'inactive':
      return 'غیرفعال'
    case 'error':
      return 'خطا'
    case 'configuring':
      return 'در حال تنظیم'
    default:
      return 'نامشخص'
  }
}

const getWebhookStatusClass = (status: string) => {
  switch (status) {
    case 'active':
      return 'bg-green-100 text-green-800'
    case 'inactive':
      return 'bg-gray-100 text-gray-800'
    case 'error':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getWebhookStatusText = (status: string) => {
  switch (status) {
    case 'active':
      return 'فعال'
    case 'inactive':
      return 'غیرفعال'
    case 'error':
      return 'خطا'
    default:
      return 'نامشخص'
  }
}

const getLogStatusClass = (status: string) => {
  switch (status) {
    case 'success':
      return 'bg-green-100 text-green-800'
    case 'error':
      return 'bg-red-100 text-red-800'
    case 'warning':
      return 'bg-yellow-100 text-yellow-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getLogStatusText = (status: string) => {
  switch (status) {
    case 'success':
      return 'موفق'
    case 'error':
      return 'خطا'
    case 'warning':
      return 'هشدار'
    default:
      return 'نامشخص'
  }
}

const formatTime = (timestamp: Date) => {
  return new Intl.DateTimeFormat('fa-IR', {
    hour: '2-digit',
    minute: '2-digit'
  }).format(timestamp)
}

const configureService = (_service: IntegrationService) => {
}

const testService = (_service: IntegrationService) => {
}

const syncData = (_service: IntegrationService) => {
}

const regenerateKey = (_keyId: number) => {
  if (confirm('آیا از تجدید کلید API اطمینان دارید؟')) {
  }
}

const testWebhook = (_webhookId: number) => {
}

const editWebhook = (_webhookId: number) => {
}

const refreshIntegrations = () => {
}

const exportIntegrations = () => {
}

const saveSettings = () => {
}

// Expose methods for parent component
defineExpose({
  integrationStats,
  emailServices,
  smsServices,
  crmServices,
  analyticsServices,
  apiKeys,
  webhooks,
  integrationLogs
})
</script>

<style scoped>
.integration-manager {
  transition: all 0.3s ease;
}

/* Service card hover effects */
.bg-white:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
}

/* Status badge animations */
.bg-green-100,
.bg-red-100,
.bg-yellow-100,
.bg-gray-100 {
  transition: all 0.2s ease;
}

/* Button hover effects */
button:hover {
  transform: translateY(-1px);
  transition: transform 0.2s ease;
}

/* Table row hover effects */
.hover\:bg-gray-50:hover {
  background-color: #f9fafb;
}

/* Responsive adjustments */
@media (max-width: 768px) {
  .grid {
    grid-template-columns: 1fr;
  }
  
  .overflow-x-auto {
    overflow-x: auto;
  }
}
</style> 
