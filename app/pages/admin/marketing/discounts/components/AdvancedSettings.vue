<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200">
    <!-- هدر بخش -->
    <div class="p-6 border-b border-gray-200">
      <div class="flex items-center justify-between">
        <div>
          <h2 class="text-lg font-semibold text-gray-900">تنظیمات پیشرفته و API</h2>
          <p class="text-gray-600 mt-1">تنظیمات پیشرفته، API، وب‌هوک‌ها و یکپارچه‌سازی‌های تخصصی</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button @click="saveAllSettings" class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors">
            <span class="i-heroicons-check ml-2"></span>
            ذخیره همه تنظیمات
          </button>
        </div>
      </div>
    </div>

    <!-- تب‌های تنظیمات -->
    <div class="border-b border-gray-200">
      <div class="flex border-b border-gray-200 overflow-x-auto">
        <button v-for="tab in tabs" :key="tab.value" @click="activeTab = tab.value"
          :class="['px-6 py-3 -mb-px font-medium text-sm focus:outline-none whitespace-nowrap', activeTab === tab.value ? 'border-b-2 border-blue-600 text-blue-700' : 'text-gray-500 hover:text-blue-600']">
          {{ tab.label }}
        </button>
      </div>
    </div>

    <!-- محتوای تب‌ها -->
    <div class="p-6">
      <!-- تنظیمات API -->
      <div v-if="activeTab === 'api'" class="space-y-6">
        <div class="bg-blue-50 p-6 rounded-lg">
          <h4 class="text-md font-medium text-blue-900 mb-2">تنظیمات API</h4>
          <p class="text-sm text-blue-700">مدیریت کلیدهای API، محدودیت‌های نرخ و دسترسی‌ها</p>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <!-- کلیدهای API -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h5 class="font-medium text-gray-900 mb-4">کلیدهای API</h5>
            <div class="space-y-4">
              <div v-for="key in apiKeys" :key="key.id" class="border border-gray-200 rounded-lg p-6">
                <div class="flex items-center justify-between mb-3">
                  <div>
                    <h6 class="font-medium text-gray-900">{{ key.name }}</h6>
                    <p class="text-sm text-gray-500">{{ key.description }}</p>
                  </div>
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <button @click="regenerateKey(key)" class="text-blue-600 hover:text-blue-900 text-sm">
                      <span class="i-heroicons-arrow-path ml-1"></span>
                      تجدید
                    </button>
                    <button @click="deleteKey(key)" class="text-red-600 hover:text-red-900 text-sm">
                      <span class="i-heroicons-trash"></span>
                    </button>
                  </div>
                </div>
                <div class="bg-gray-50 p-3 rounded">
                  <div class="flex items-center justify-between">
                    <code class="text-sm text-gray-700 font-mono">{{ key.key }}</code>
                    <button @click="copyToClipboard(key.key)" class="text-blue-600 hover:text-blue-900">
                      <span class="i-heroicons-clipboard"></span>
                    </button>
                  </div>
                </div>
                <div class="mt-3 flex justify-between text-sm">
                  <span class="text-gray-500">آخرین استفاده:</span>
                  <span class="font-medium">{{ formatDate(key.lastUsed) }}</span>
                </div>
              </div>
              
              <button @click="createNewApiKey" class="w-full p-6 border-2 border-dashed border-gray-300 rounded-lg text-gray-500 hover:border-blue-500 hover:text-blue-500 transition-colors">
                <span class="i-heroicons-plus ml-2"></span>
                افزودن کلید API جدید
              </button>
            </div>
          </div>

          <!-- محدودیت‌های نرخ -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h5 class="font-medium text-gray-900 mb-4">محدودیت‌های نرخ</h5>
            <div class="space-y-4">
              <div v-for="limit in rateLimits" :key="limit.id" class="flex items-center justify-between p-3 bg-gray-50 rounded">
                <div>
                  <span class="text-sm font-medium text-gray-700">{{ limit.endpoint }}</span>
                  <p class="text-xs text-gray-500">{{ limit.description }}</p>
                </div>
                <div class="flex items-center space-x-2 space-x-reverse">
                  <input v-model="limit.limit" type="number" class="w-20 px-2 py-1 border border-gray-300 rounded text-sm">
                  <span class="text-sm text-gray-500">درخواست/دقیقه</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- وب‌هوک‌ها -->
      <div v-if="activeTab === 'webhooks'" class="space-y-6">
        <div class="bg-green-50 p-6 rounded-lg">
          <h4 class="text-md font-medium text-green-900 mb-2">مدیریت وب‌هوک‌ها</h4>
          <p class="text-sm text-green-700">تنظیم وب‌هوک‌ها برای اطلاع‌رسانی رویدادها به سیستم‌های خارجی</p>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <!-- لیست وب‌هوک‌ها -->
          <div class="border border-gray-200 rounded-lg p-6">
            <div class="flex items-center justify-between mb-4">
              <h5 class="font-medium text-gray-900">وب‌هوک‌های فعال</h5>
              <button @click="showWebhookForm = true" class="px-3 py-1 bg-blue-600 text-white rounded text-sm hover:bg-blue-700">
                <span class="i-heroicons-plus ml-1"></span>
                افزودن
              </button>
            </div>
            
            <div class="space-y-4">
              <div v-for="webhook in webhooks" :key="webhook.id" class="border border-gray-200 rounded-lg p-6">
                <div class="flex items-center justify-between mb-3">
                  <div>
                    <h6 class="font-medium text-gray-900">{{ webhook.name }}</h6>
                    <p class="text-sm text-gray-500">{{ webhook.url }}</p>
                  </div>
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <span :class="['px-2 py-1 rounded-full text-xs', webhook.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800']">
                      {{ webhook.status === 'active' ? 'فعال' : 'غیرفعال' }}
                    </span>
                    <button @click="editWebhook(webhook)" class="text-blue-600 hover:text-blue-900">
                      <span class="i-heroicons-pencil-square"></span>
                    </button>
                    <button @click="deleteWebhook(webhook)" class="text-red-600 hover:text-red-900">
                      <span class="i-heroicons-trash"></span>
                    </button>
                  </div>
                </div>
                
                <div class="space-y-2">
                  <div class="flex flex-wrap gap-2">
                    <span v-for="event in webhook.events" :key="event" class="px-2 py-1 bg-blue-100 text-blue-800 rounded text-xs">
                      {{ event }}
                    </span>
                  </div>
                  <div class="flex justify-between text-sm">
                    <span class="text-gray-500">آخرین ارسال:</span>
                    <span class="font-medium">{{ formatDate(webhook.lastSent) }}</span>
                  </div>
                  <div class="flex justify-between text-sm">
                    <span class="text-gray-500">نرخ موفقیت:</span>
                    <span class="font-medium">{{ webhook.successRate }}%</span>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- رویدادهای وب‌هوک -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h5 class="font-medium text-gray-900 mb-4">رویدادهای قابل تنظیم</h5>
            <div class="space-y-3">
              <div v-for="event in webhookEvents" :key="event.name" class="flex items-center justify-between p-3 bg-gray-50 rounded">
                <div>
                  <span class="text-sm font-medium text-gray-700">{{ event.name }}</span>
                  <p class="text-xs text-gray-500">{{ event.description }}</p>
                </div>
                <div class="flex items-center space-x-2 space-x-reverse">
                  <input v-model="event.enabled" type="checkbox" class="rounded border-gray-300">
                  <span class="text-xs text-gray-500">{{ event.frequency }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- تنظیمات پیشرفته -->
      <div v-if="activeTab === 'advanced'" class="space-y-6">
        <div class="bg-purple-50 p-6 rounded-lg">
          <h4 class="text-md font-medium text-purple-900 mb-2">تنظیمات پیشرفته سیستم</h4>
          <p class="text-sm text-purple-700">تنظیمات تخصصی برای عملکرد بهینه سیستم</p>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <!-- تنظیمات کش -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h5 class="font-medium text-gray-900 mb-4">تنظیمات کش</h5>
            <div class="space-y-4">
              <div class="flex items-center justify-between">
                <div>
                  <span class="text-sm font-medium text-gray-700">فعال‌سازی کش</span>
                  <p class="text-xs text-gray-500">بهبود عملکرد با ذخیره‌سازی موقت</p>
                </div>
                <input v-model="advancedSettings.cache.enabled" type="checkbox" class="rounded border-gray-300">
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">زمان انقضای کش (دقیقه)</label>
                <input v-model="advancedSettings.cache.ttl" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر اندازه کش (MB)</label>
                <input v-model="advancedSettings.cache.maxSize" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
              </div>
            </div>
          </div>

          <!-- تنظیمات پردازش -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h5 class="font-medium text-gray-900 mb-4">تنظیمات پردازش</h5>
            <div class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">تعداد کارهای همزمان</label>
                <input v-model="advancedSettings.processing.concurrentJobs" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">تایم‌اوت کار (ثانیه)</label>
                <input v-model="advancedSettings.processing.jobTimeout" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
              </div>
              
              <div class="flex items-center justify-between">
                <div>
                  <span class="text-sm font-medium text-gray-700">پردازش غیرهمزمان</span>
                  <p class="text-xs text-gray-500">اجرای کارها در پس‌زمینه</p>
                </div>
                <input v-model="advancedSettings.processing.asyncEnabled" type="checkbox" class="rounded border-gray-300">
              </div>
            </div>
          </div>
        </div>

        <!-- تنظیمات امنیتی پیشرفته -->
        <div class="border border-gray-200 rounded-lg p-6">
          <h5 class="font-medium text-gray-900 mb-4">تنظیمات امنیتی پیشرفته</h5>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div class="space-y-4">
              <div class="flex items-center justify-between">
                <div>
                  <span class="text-sm font-medium text-gray-700">رمزنگاری داده‌ها</span>
                  <p class="text-xs text-gray-500">رمزنگاری کوپن‌ها و اطلاعات حساس</p>
                </div>
                <input v-model="advancedSettings.security.encryption" type="checkbox" class="rounded border-gray-300">
              </div>
              
              <div class="flex items-center justify-between">
                <div>
                  <span class="text-sm font-medium text-gray-700">احراز هویت دو مرحله‌ای</span>
                  <p class="text-xs text-gray-500">برای دسترسی به تنظیمات حساس</p>
                </div>
                <input v-model="advancedSettings.security.twoFactor" type="checkbox" class="rounded border-gray-300">
              </div>
              
              <div class="flex items-center justify-between">
                <div>
                  <span class="text-sm font-medium text-gray-700">لاگ‌های امنیتی</span>
                  <p class="text-xs text-gray-500">ثبت تمام فعالیت‌های حساس</p>
                </div>
                <input v-model="advancedSettings.security.securityLogs" type="checkbox" class="rounded border-gray-300">
              </div>
            </div>
            
            <div class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر تلاش‌های ناموفق</label>
                <input v-model="advancedSettings.security.maxFailedAttempts" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">زمان قفل حساب (دقیقه)</label>
                <input v-model="advancedSettings.security.lockoutDuration" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">طول حداقل رمز عبور</label>
                <input v-model="advancedSettings.security.minPasswordLength" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- لاگ‌ها و مانیتورینگ -->
      <div v-if="activeTab === 'logs'" class="space-y-6">
        <div class="bg-orange-50 p-6 rounded-lg">
          <h4 class="text-md font-medium text-orange-900 mb-2">لاگ‌ها و مانیتورینگ</h4>
          <p class="text-sm text-orange-700">مشاهده لاگ‌های سیستم و مانیتورینگ عملکرد</p>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <!-- فیلترهای لاگ -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h5 class="font-medium text-gray-900 mb-4">فیلترهای لاگ</h5>
            <div class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">سطح لاگ</label>
                <select v-model="logFilters.level" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
                  <option value="">همه سطوح</option>
                  <option value="error">خطا</option>
                  <option value="warning">هشدار</option>
                  <option value="info">اطلاعات</option>
                  <option value="debug">دیباگ</option>
                </select>
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">نوع رویداد</label>
                <select v-model="logFilters.type" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
                  <option value="">همه انواع</option>
                  <option value="api">API</option>
                  <option value="webhook">وب‌هوک</option>
                  <option value="security">امنیت</option>
                  <option value="system">سیستم</option>
                </select>
              </div>
              
              <div class="grid grid-cols-2 gap-6">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">از تاریخ</label>
                  <input v-model="logFilters.startDate" type="date" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">تا تاریخ</label>
                  <input v-model="logFilters.endDate" type="date" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
                </div>
              </div>
              
              <button @click="loadLogs" class="w-full px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors">
                بارگذاری لاگ‌ها
              </button>
            </div>
          </div>

          <!-- نمایش لاگ‌ها -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h5 class="font-medium text-gray-900 mb-4">لاگ‌های اخیر</h5>
            <div class="space-y-3 max-h-96 overflow-y-auto">
              <div v-for="log in logs" :key="log.id" class="p-3 bg-gray-50 rounded text-sm">
                <div class="flex items-center justify-between mb-2">
                  <span :class="['px-2 py-1 rounded text-xs', getLogLevelClass(log.level)]">
                    {{ log.level }}
                  </span>
                  <span class="text-gray-500">{{ formatDate(log.timestamp) }}</span>
                </div>
                <div class="font-medium text-gray-900">{{ log.message }}</div>
                <div class="text-gray-500 text-xs mt-1">{{ log.type }} • {{ log.user }}</div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- مودال فرم وب‌هوک -->
    <div v-if="showWebhookForm" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg shadow-xl max-w-2xl w-full mx-4 max-h-[90vh] overflow-y-auto">
        <div class="p-6 border-b border-gray-200">
          <div class="flex items-center justify-between">
            <h3 class="text-lg font-semibold text-gray-900">
              {{ editingWebhook ? 'ویرایش وب‌هوک' : 'افزودن وب‌هوک جدید' }}
            </h3>
            <button @click="closeWebhookForm" class="text-gray-400 hover:text-gray-600">
              <span class="i-heroicons-x-mark text-xl"></span>
            </button>
          </div>
        </div>
        
        <form @submit.prevent="handleWebhookSubmit" class="p-6 space-y-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نام وب‌هوک *</label>
            <input v-model="webhookForm.name" type="text" required class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="نام وب‌هوک">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">URL *</label>
            <input v-model="webhookForm.url" type="url" required class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="https://example.com/webhook">
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">رویدادها *</label>
            <div class="space-y-2">
              <label v-for="event in availableEvents" :key="event" class="flex items-center">
                <input v-model="webhookForm.events" :value="event" type="checkbox" class="rounded border-gray-300 ml-2">
                <span class="text-sm text-gray-700">{{ event }}</span>
              </label>
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
            <select v-model="webhookForm.status" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
              <option value="active">فعال</option>
              <option value="inactive">غیرفعال</option>
            </select>
          </div>
        </form>
        
        <div class="p-6 border-t border-gray-200 flex justify-end space-x-3 space-x-reverse">
          <button @click="closeWebhookForm" class="px-6 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200 transition-colors">
            انصراف
          </button>
          <button @click="handleWebhookSubmit" class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors">
            ذخیره
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'

const activeTab = ref('api')
const showWebhookForm = ref(false)
const editingWebhook = ref<any>(null)

const tabs = [
  { value: 'api', label: 'تنظیمات API' },
  { value: 'webhooks', label: 'وب‌هوک‌ها' },
  { value: 'advanced', label: 'تنظیمات پیشرفته' },
  { value: 'logs', label: 'لاگ‌ها و مانیتورینگ' }
]

const apiKeys = ref([
  {
    id: 1,
    name: 'کلید اصلی',
    description: 'دسترسی کامل به API',
    key: 'sk_live_1234567890abcdef',
    lastUsed: '2024-01-15T10:30:00'
  },
  {
    id: 2,
    name: 'کلید تست',
    description: 'فقط برای تست و توسعه',
    key: 'sk_test_abcdef1234567890',
    lastUsed: '2024-01-14T15:45:00'
  }
])

const rateLimits = ref([
  { id: 1, endpoint: '/api/coupons', description: 'ایجاد و ویرایش کوپن‌ها', limit: 100 },
  { id: 2, endpoint: '/api/campaigns', description: 'مدیریت کمپین‌ها', limit: 50 },
  { id: 3, endpoint: '/api/analytics', description: 'دریافت آمار و گزارش‌ها', limit: 200 }
])

const webhooks = ref([
  {
    id: 1,
    name: 'CRM Integration',
    url: 'https://crm.example.com/webhook',
    status: 'active',
    events: ['coupon_created', 'coupon_used'],
    lastSent: '2024-01-15T10:30:00',
    successRate: 98.5
  },
  {
    id: 2,
    name: 'Analytics Service',
    url: 'https://analytics.example.com/webhook',
    status: 'active',
    events: ['campaign_started', 'campaign_ended'],
    lastSent: '2024-01-14T09:15:00',
    successRate: 95.2
  }
])

const webhookEvents = ref([
  { name: 'coupon_created', description: 'ایجاد کوپن جدید', enabled: true, frequency: 'فوری' },
  { name: 'coupon_used', description: 'استفاده از کوپن', enabled: true, frequency: 'فوری' },
  { name: 'coupon_expired', description: 'انقضای کوپن', enabled: false, frequency: 'روزانه' },
  { name: 'campaign_started', description: 'شروع کمپین', enabled: true, frequency: 'فوری' },
  { name: 'campaign_ended', description: 'پایان کمپین', enabled: true, frequency: 'فوری' }
])

const advancedSettings = reactive({
  cache: {
    enabled: true,
    ttl: 30,
    maxSize: 100
  },
  processing: {
    concurrentJobs: 5,
    jobTimeout: 300,
    asyncEnabled: true
  },
  security: {
    encryption: true,
    twoFactor: false,
    securityLogs: true,
    maxFailedAttempts: 5,
    lockoutDuration: 30,
    minPasswordLength: 8
  }
})

const logFilters = reactive({
  level: '',
  type: '',
  startDate: '',
  endDate: ''
})

const logs = ref([
  {
    id: 1,
    level: 'info',
    message: 'کوپن جدید ایجاد شد: WELCOME2024',
    type: 'api',
    user: 'admin@example.com',
    timestamp: '2024-01-15T10:30:00'
  },
  {
    id: 2,
    level: 'warning',
    message: 'محدودیت نرخ API نزدیک به سقف',
    type: 'system',
    user: 'system',
    timestamp: '2024-01-15T10:25:00'
  },
  {
    id: 3,
    level: 'error',
    message: 'خطا در ارسال وب‌هوک به CRM',
    type: 'webhook',
    user: 'system',
    timestamp: '2024-01-15T10:20:00'
  }
])

const availableEvents = [
  'coupon_created',
  'coupon_used',
  'coupon_expired',
  'campaign_started',
  'campaign_ended',
  'user_registered',
  'order_completed'
]

const webhookForm = reactive({
  name: '',
  url: '',
  events: [] as string[],
  status: 'active'
})

const formatDate = (date: string): string => {
  return new Intl.DateTimeFormat('fa-IR').format(new Date(date))
}

const getLogLevelClass = (level: string): string => {
  const classes = {
    error: 'bg-red-100 text-red-800',
    warning: 'bg-yellow-100 text-yellow-800',
    info: 'bg-blue-100 text-blue-800',
    debug: 'bg-gray-100 text-gray-800'
  }
  return classes[level as keyof typeof classes] || 'bg-gray-100 text-gray-800'
}

const regenerateKey = async (key: any) => {
  if (confirm(`آیا از تجدید کلید API "${key.name}" اطمینان دارید؟`)) {
    try {
      key.key = 'sk_' + Math.random().toString(36).substr(2, 9)
      key.lastUsed = new Date().toISOString()
    } catch (error) {
      console.error('خطا در تجدید کلید:', error)
    }
  }
}

const deleteKey = async (key: any) => {
  if (confirm(`آیا از حذف کلید API "${key.name}" اطمینان دارید؟`)) {
    try {
      const index = apiKeys.value.findIndex(k => k.id === key.id)
      if (index !== -1) {
        apiKeys.value.splice(index, 1)
      }
    } catch (error) {
      console.error('خطا در حذف کلید:', error)
    }
  }
}

const copyToClipboard = async (text: string) => {
  try {
    await navigator.clipboard.writeText(text)
    alert('کلید در کلیپ‌بورد کپی شد')
  } catch (error) {
    console.error('خطا در کپی کردن:', error)
  }
}

const createNewApiKey = () => {
  const newKey = {
    id: Date.now(),
    name: 'کلید جدید',
    description: 'توضیحات کلید',
    key: 'sk_' + Math.random().toString(36).substr(2, 9),
    lastUsed: null
  }
  apiKeys.value.unshift(newKey)
}

const editWebhook = (webhook: any) => {
  editingWebhook.value = webhook
  Object.assign(webhookForm, webhook)
  showWebhookForm.value = true
}

const deleteWebhook = async (webhook: any) => {
  if (confirm(`آیا از حذف وب‌هوک "${webhook.name}" اطمینان دارید؟`)) {
    try {
      const index = webhooks.value.findIndex(w => w.id === webhook.id)
      if (index !== -1) {
        webhooks.value.splice(index, 1)
      }
    } catch (error) {
      console.error('خطا در حذف وب‌هوک:', error)
    }
  }
}

const handleWebhookSubmit = async () => {
  if (!webhookForm.name || !webhookForm.url || webhookForm.events.length === 0) {
    alert('لطفاً فیلدهای اجباری را پر کنید')
    return
  }
  
  if (editingWebhook.value) {
    Object.assign(editingWebhook.value, webhookForm)
  } else {
    webhooks.value.unshift({
      id: Date.now(),
      ...webhookForm,
      lastSent: null,
      successRate: 0
    })
  }
  closeWebhookForm()
}

const closeWebhookForm = () => {
  showWebhookForm.value = false
  editingWebhook.value = null
  Object.assign(webhookForm, { name: '', url: '', events: [], status: 'active' })
}

const loadLogs = () => {
  // پیاده‌سازی بارگذاری لاگ‌ها بر اساس فیلترها
  console.log('بارگذاری لاگ‌ها با فیلترها:', logFilters)
}

const saveAllSettings = async () => {
  try {
    // پیاده‌سازی ذخیره همه تنظیمات
    console.log('ذخیره تنظیمات پیشرفته:', advancedSettings)
    alert('تنظیمات با موفقیت ذخیره شد')
  } catch (error) {
    console.error('خطا در ذخیره تنظیمات:', error)
    alert('خطا در ذخیره تنظیمات')
  }
}
</script>

<!--
  کامپوننت تنظیمات پیشرفته و API
  شامل:
  1. مدیریت کلیدهای API و محدودیت‌های نرخ
  2. مدیریت وب‌هوک‌ها و رویدادها
  3. تنظیمات پیشرفته سیستم (کش، پردازش، امنیت)
  4. لاگ‌ها و مانیتورینگ
  توضیحات کامل به فارسی در کد
--> 
