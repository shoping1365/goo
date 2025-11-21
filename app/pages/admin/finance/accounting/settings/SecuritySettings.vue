<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">تنظیمات امنیتی</h4>
        <p class="text-sm text-gray-600 mt-1">تنظیمات امنیتی و محافظتی اتصال حسابداری</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
          @click="testSecuritySettings"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          تست امنیت
        </button>
        <button
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          @click="saveSecuritySettings"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
          </svg>
          ذخیره تنظیمات
        </button>
      </div>
    </div>

    <!-- آمار امنیتی -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div class="bg-green-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-green-600">{{ securityStats.secureConnections }}</div>
            <div class="text-sm text-green-700">اتصالات امن</div>
          </div>
          <div class="w-10 h-10 bg-green-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-blue-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-blue-600">{{ securityStats.encryptedData }}</div>
            <div class="text-sm text-blue-700">داده‌های رمزگذاری شده</div>
          </div>
          <div class="w-10 h-10 bg-blue-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-yellow-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-yellow-600">{{ securityStats.failedAttempts }}</div>
            <div class="text-sm text-yellow-700">تلاش‌های ناموفق</div>
          </div>
          <div class="w-10 h-10 bg-yellow-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-red-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-red-600">{{ securityStats.securityScore }}</div>
            <div class="text-sm text-red-700">امتیاز امنیتی</div>
          </div>
          <div class="w-10 h-10 bg-red-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات امنیتی -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- تنظیمات احراز هویت -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h5 class="text-md font-medium text-gray-900 mb-4">تنظیمات احراز هویت</h5>
        <div class="space-y-4">
          <div>
            <label class="flex items-center">
              <input
                v-model="securitySettings.twoFactorAuth"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">فعال کردن احراز هویت دو مرحله‌ای</span>
            </label>
          </div>

          <div>
            <label class="flex items-center">
              <input
                v-model="securitySettings.sessionTimeout"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">محدودیت زمان نشست</span>
            </label>
          </div>

          <div v-if="securitySettings.sessionTimeout">
            <label class="block text-sm font-medium text-gray-700 mb-2">زمان نشست (دقیقه)</label>
            <input
              v-model.number="securitySettings.sessionTimeoutMinutes"
              type="number"
              min="5"
              max="480"
              class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
          </div>

          <div>
            <label class="flex items-center">
              <input
                v-model="securitySettings.ipWhitelist"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">محدودیت دسترسی بر اساس IP</span>
            </label>
          </div>

          <div v-if="securitySettings.ipWhitelist">
            <label class="block text-sm font-medium text-gray-700 mb-2">آدرس‌های IP مجاز</label>
            <textarea
              v-model="securitySettings.allowedIPs"
              rows="3"
              class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="هر IP در یک خط جداگانه"
            ></textarea>
          </div>
        </div>
      </div>

      <!-- تنظیمات رمزگذاری -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h5 class="text-md font-medium text-gray-900 mb-4">تنظیمات رمزگذاری</h5>
        <div class="space-y-4">
          <div>
            <label class="flex items-center">
              <input
                v-model="securitySettings.encryptSensitiveData"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">رمزگذاری داده‌های حساس</span>
            </label>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">الگوریتم رمزگذاری</label>
            <select
              v-model="securitySettings.encryptionAlgorithm"
              class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="AES-256">AES-256 (پیشنهادی)</option>
              <option value="AES-192">AES-192</option>
              <option value="AES-128">AES-128</option>
              <option value="ChaCha20">ChaCha20</option>
            </select>
          </div>

          <div>
            <label class="flex items-center">
              <input
                v-model="securitySettings.sslVerification"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">تأیید گواهی SSL</span>
            </label>
          </div>

          <div>
            <label class="flex items-center">
              <input
                v-model="securitySettings.secureHeaders"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">هدرهای امنیتی HTTP</span>
            </label>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات لاگ و نظارت -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- تنظیمات لاگ امنیتی -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h5 class="text-md font-medium text-gray-900 mb-4">لاگ امنیتی</h5>
        <div class="space-y-4">
          <div>
            <label class="flex items-center">
              <input
                v-model="securitySettings.securityLogging"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">فعال کردن لاگ امنیتی</span>
            </label>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">سطح لاگ امنیتی</label>
            <select
              v-model="securitySettings.securityLogLevel"
              class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="low">کم (فقط خطاها)</option>
              <option value="medium">متوسط (خطاها و هشدارها)</option>
              <option value="high">زیاد (تمام رویدادها)</option>
            </select>
          </div>

          <div>
            <label class="flex items-center">
              <input
                v-model="securitySettings.auditTrail"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">ردیابی حسابرسی</span>
            </label>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">مدت نگهداری لاگ (روز)</label>
            <input
              v-model.number="securitySettings.logRetentionDays"
              type="number"
              min="1"
              max="365"
              class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
          </div>
        </div>
      </div>

      <!-- تنظیمات هشدار امنیتی -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h5 class="text-md font-medium text-gray-900 mb-4">هشدارهای امنیتی</h5>
        <div class="space-y-4">
          <div>
            <label class="flex items-center">
              <input
                v-model="securitySettings.suspiciousActivityAlert"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">هشدار فعالیت مشکوک</span>
            </label>
          </div>

          <div>
            <label class="flex items-center">
              <input
                v-model="securitySettings.failedLoginAlert"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">هشدار ورود ناموفق</span>
            </label>
          </div>

          <div>
            <label class="flex items-center">
              <input
                v-model="securitySettings.dataBreachAlert"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">هشدار نشت داده</span>
            </label>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">کانال‌های هشدار</label>
            <div class="space-y-2">
              <label class="flex items-center">
                <input
                  v-model="securitySettings.emailAlerts"
                  type="checkbox"
                  class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                />
                <span class="mr-2 text-sm text-gray-700">ایمیل</span>
              </label>
              <label class="flex items-center">
                <input
                  v-model="securitySettings.smsAlerts"
                  type="checkbox"
                  class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                />
                <span class="mr-2 text-sm text-gray-700">پیامک</span>
              </label>
              <label class="flex items-center">
                <input
                  v-model="securitySettings.pushAlerts"
                  type="checkbox"
                  class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                />
                <span class="mr-2 text-sm text-gray-700">اعلان‌های فوری</span>
              </label>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- وضعیت امنیتی -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">وضعیت امنیتی</h5>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div
          v-for="status in securityStatus"
          :key="status.name"
          class="flex items-center justify-between p-6 rounded-lg"
          :class="getStatusClass(status.status)"
        >
          <div class="flex items-center space-x-3 space-x-reverse">
            <div
              class="w-3 h-3 rounded-full"
              :class="getStatusColor(status.status)"
            ></div>
            <span class="text-sm font-medium" :class="getStatusTextClass(status.status)">
              {{ status.name }}
            </span>
          </div>
          <span class="text-sm" :class="getStatusTextClass(status.status)">
            {{ status.value }}
          </span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// آمار امنیتی
const securityStats = ref({
  secureConnections: 8,
  encryptedData: 15420,
  failedAttempts: 12,
  securityScore: 95
})

// تنظیمات امنیتی
const securitySettings = ref({
  twoFactorAuth: true,
  sessionTimeout: true,
  sessionTimeoutMinutes: 30,
  ipWhitelist: false,
  allowedIPs: '',
  encryptSensitiveData: true,
  encryptionAlgorithm: 'AES-256',
  sslVerification: true,
  secureHeaders: true,
  securityLogging: true,
  securityLogLevel: 'medium',
  auditTrail: true,
  logRetentionDays: 90,
  suspiciousActivityAlert: true,
  failedLoginAlert: true,
  dataBreachAlert: true,
  emailAlerts: true,
  smsAlerts: false,
  pushAlerts: true
})

// وضعیت امنیتی
const securityStatus = ref([
  { name: 'احراز هویت دو مرحله‌ای', status: 'active', value: 'فعال' },
  { name: 'رمزگذاری داده‌ها', status: 'active', value: 'فعال' },
  { name: 'تأیید SSL', status: 'active', value: 'فعال' },
  { name: 'لاگ امنیتی', status: 'active', value: 'فعال' },
  { name: 'ردیابی حسابرسی', status: 'active', value: 'فعال' },
  { name: 'هشدارهای امنیتی', status: 'warning', value: 'نیاز به بررسی' }
])

// رنگ وضعیت
const getStatusColor = (status: string) => {
  const colors = {
    active: 'bg-green-500',
    warning: 'bg-yellow-500',
    error: 'bg-red-500'
  }
  return colors[status] || 'bg-gray-500'
}

// کلاس وضعیت
const getStatusClass = (status: string) => {
  const classes = {
    active: 'bg-green-50 border-green-200',
    warning: 'bg-yellow-50 border-yellow-200',
    error: 'bg-red-50 border-red-200'
  }
  return classes[status] || 'bg-gray-50 border-gray-200'
}

// کلاس متن وضعیت
const getStatusTextClass = (status: string) => {
  const classes = {
    active: 'text-green-700',
    warning: 'text-yellow-700',
    error: 'text-red-700'
  }
  return classes[status] || 'text-gray-700'
}

// تست تنظیمات امنیتی
const testSecuritySettings = () => {
  // TODO: تست تنظیمات امنیتی

}

// ذخیره تنظیمات امنیتی
const saveSecuritySettings = () => {
  // TODO: ذخیره تنظیمات امنیتی

}
</script>

<!--
  کامپوننت تنظیمات امنیتی
  شامل:
  1. آمار امنیتی
  2. تنظیمات احراز هویت
  3. تنظیمات رمزگذاری
  4. لاگ امنیتی
  5. هشدارهای امنیتی
  6. وضعیت امنیتی
  7. طراحی مدرن و کاملاً ریسپانسیو
--> 
