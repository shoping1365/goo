<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">تنظیمات امنیتی</h4>
        <p class="text-sm text-gray-600 mt-1">مدیریت رمزهای عبور و احراز هویت</p>
      </div>
    </div>

    <!-- آمار امنیت -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">وضعیت امنیت</h5>
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">امنیت کلی</h6>
            <div class="w-3 h-3 bg-green-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-green-600">{{ securityStats.overall }}%</div>
          <div class="text-xs text-gray-500 mt-1">عالی</div>
        </div>

        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">احراز هویت دو مرحله‌ای</h6>
            <div class="w-3 h-3 bg-green-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-green-600">{{ securityStats.twoFactor }}</div>
          <div class="text-xs text-gray-500 mt-1">فعال</div>
        </div>

        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">آخرین ورود</h6>
            <div class="w-3 h-3 bg-blue-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-blue-600">{{ securityStats.lastLogin }}</div>
          <div class="text-xs text-gray-500 mt-1">ساعت پیش</div>
        </div>

        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">تلاش‌های ناموفق</h6>
            <div class="w-3 h-3 bg-yellow-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-yellow-600">{{ securityStats.failedAttempts }}</div>
          <div class="text-xs text-gray-500 mt-1">امروز</div>
        </div>
      </div>
    </div>

    <!-- تنظیمات رمز عبور -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">تنظیمات رمز عبور</h5>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حداقل طول رمز عبور</label>
          <input
            v-model="passwordSettings.minLength"
            type="number"
            min="8"
            max="32"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر تلاش ورود</label>
          <input
            v-model="passwordSettings.maxAttempts"
            type="number"
            min="3"
            max="10"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">مدت قفل حساب</label>
          <select
            v-model="passwordSettings.lockoutDuration"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="15">۱۵ دقیقه</option>
            <option value="30">۳۰ دقیقه</option>
            <option value="60">۱ ساعت</option>
            <option value="1440">۲۴ ساعت</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">انقضای رمز عبور</label>
          <select
            v-model="passwordSettings.expiryDays"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="30">۳۰ روز</option>
            <option value="60">۶۰ روز</option>
            <option value="90">۹۰ روز</option>
            <option value="180">۱۸۰ روز</option>
          </select>
        </div>
      </div>

      <div class="mt-4 space-y-3">
        <label class="flex items-center">
          <input
            v-model="passwordSettings.requireUppercase"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">نیاز به حروف بزرگ</span>
        </label>

        <label class="flex items-center">
          <input
            v-model="passwordSettings.requireLowercase"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">نیاز به حروف کوچک</span>
        </label>

        <label class="flex items-center">
          <input
            v-model="passwordSettings.requireNumbers"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">نیاز به اعداد</span>
        </label>

        <label class="flex items-center">
          <input
            v-model="passwordSettings.requireSymbols"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">نیاز به نمادهای خاص</span>
        </label>
      </div>
    </div>

    <!-- احراز هویت دو مرحله‌ای -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">احراز هویت دو مرحله‌ای</h5>
      <div class="space-y-4">
        <div class="flex items-center justify-between p-6 border border-gray-200 rounded-lg">
          <div>
            <h6 class="text-sm font-medium text-gray-900">احراز هویت دو مرحله‌ای</h6>
            <p class="text-xs text-gray-500">استفاده از کد تأیید برای ورود</p>
          </div>
          <div class="flex items-center space-x-3 space-x-reverse">
            <span class="text-sm text-green-600 font-medium">فعال</span>
            <button
              class="relative inline-flex h-6 w-11 flex-shrink-0 cursor-pointer rounded-full border-2 border-transparent transition-colors duration-200 ease-in-out focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 bg-blue-600"
              @click="toggleTwoFactor"
            >
              <span class="translate-x-5 pointer-events-none relative inline-block h-5 w-5 transform rounded-full bg-white shadow ring-0 transition duration-200 ease-in-out"></span>
            </button>
          </div>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">روش احراز هویت</label>
            <select
              v-model="twoFactorSettings.method"
              class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="app">اپلیکیشن احراز هویت</option>
              <option value="sms">پیامک</option>
              <option value="email">ایمیل</option>
            </select>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">مدت اعتبار کد</label>
            <select
              v-model="twoFactorSettings.expiryTime"
              class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="30">۳۰ ثانیه</option>
              <option value="60">۱ دقیقه</option>
              <option value="300">۵ دقیقه</option>
            </select>
          </div>
        </div>
      </div>
    </div>

    <!-- تاریخچه ورود -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">تاریخچه ورود</h5>
      </div>
      <div class="p-6">
        <div class="space-y-4">
          <div
            v-for="login in loginHistory"
            :key="login.id"
            class="flex items-start space-x-4 space-x-reverse p-6 border border-gray-200 rounded-lg"
          >
            <div
              class="w-2 h-2 rounded-full mt-2"
              :class="getLoginStatusColor(login.status)"
            ></div>
            <div class="flex-1">
              <div class="flex items-center justify-between">
                <h6 class="text-sm font-medium text-gray-900">{{ login.user }}</h6>
                <span class="text-xs text-gray-500">{{ login.timestamp }}</span>
              </div>
              <p class="text-sm text-gray-600 mt-1">{{ login.ip }} - {{ login.device }}</p>
              <div class="flex items-center space-x-4 space-x-reverse mt-2 text-xs text-gray-500">
                <span>وضعیت: {{ getLoginStatusText(login.status) }}</span>
                <span>روش: {{ login.method }}</span>
                <span>موقعیت: {{ login.location }}</span>
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

// آمار امنیت
const securityStats = ref({
  overall: 95,
  twoFactor: 'فعال',
  lastLogin: 2,
  failedAttempts: 3
})

// تنظیمات رمز عبور
const passwordSettings = ref({
  minLength: 12,
  maxAttempts: 5,
  lockoutDuration: 30,
  expiryDays: 90,
  requireUppercase: true,
  requireLowercase: true,
  requireNumbers: true,
  requireSymbols: true
})

// تنظیمات احراز هویت دو مرحله‌ای
const twoFactorSettings = ref({
  method: 'app',
  expiryTime: 30
})

// تاریخچه ورود
const loginHistory = ref([
  {
    id: 1,
    user: 'مدیر سیستم',
    ip: '192.168.1.100',
    device: 'Chrome - Windows',
    status: 'success',
    method: 'رمز عبور + 2FA',
    location: 'تهران، ایران',
    timestamp: '۲ ساعت پیش'
  },
  {
    id: 2,
    user: 'کاربر حسابداری',
    ip: '192.168.1.101',
    device: 'Firefox - macOS',
    status: 'success',
    method: 'رمز عبور',
    location: 'اصفهان، ایران',
    timestamp: '۴ ساعت پیش'
  }
])

// رنگ‌های وضعیت
const getLoginStatusColor = (status: string) => {
  const colors = {
    success: 'bg-green-500',
    failed: 'bg-red-500',
    blocked: 'bg-yellow-500'
  }
  return colors[status] || 'bg-gray-500'
}

const getLoginStatusText = (status: string) => {
  const texts = {
    success: 'موفق',
    failed: 'ناموفق',
    blocked: 'مسدود'
  }
  return texts[status] || status
}

// عملیات
const toggleTwoFactor = () => {
  console.log('تغییر وضعیت احراز هویت دو مرحله‌ای')
}
</script>

<!--
  کامپوننت تنظیمات امنیتی
  شامل:
  1. آمار امنیت
  2. تنظیمات رمز عبور
  3. احراز هویت دو مرحله‌ای
  4. تاریخچه ورود
  5. طراحی مدرن و کاملاً ریسپانسیو
--> 
