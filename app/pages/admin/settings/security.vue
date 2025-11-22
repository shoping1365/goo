<template>
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8" dir="rtl">
    <div class="space-y-6">
      <!-- reCAPTCHA Settings -->
      <div class="bg-white shadow rounded-lg p-6">
        <h3 class="text-lg font-medium text-gray-900 mb-4">تنظیمات reCAPTCHA</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">نسخه reCAPTCHA</label>
            <select
              v-model="recaptchaSettings.version"
              class="mt-1 block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm rounded-md"
            >
              <option value="v2">reCAPTCHA v2</option>
              <option value="v3">reCAPTCHA v3</option>
            </select>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">کلید سایت</label>
            <input
              v-model="recaptchaSettings.siteKey"
              type="text"
              class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
              placeholder="کلید سایت reCAPTCHA را وارد کنید"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">کلید مخفی</label>
            <input
              v-model="recaptchaSettings.secretKey"
              type="password"
              class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
              placeholder="کلید مخفی reCAPTCHA را وارد کنید"
            />
          </div>

          <div v-if="recaptchaSettings.version === 'v2'">
            <label class="block text-sm font-medium text-gray-700 mb-1">تم</label>
            <select
              v-model="recaptchaSettings.theme"
              class="mt-1 block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm rounded-md"
            >
              <option value="light">روشن</option>
              <option value="dark">تیره</option>
            </select>
          </div>

          <div v-if="recaptchaSettings.version === 'v3'">
            <label class="block text-sm font-medium text-gray-700 mb-1">آستانه امتیاز</label>
            <input
              v-model="recaptchaSettings.scoreThreshold"
              type="number"
              min="0"
              max="1"
              step="0.1"
              class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            />
            <p class="mt-1 text-sm text-gray-500">امتیاز بین 0 تا 1 (پیش‌فرض: 0.5)</p>
          </div>

          <div class="flex justify-end">
            <button
              :disabled="loading"
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 disabled:opacity-50"
              @click="handleSaveRecaptchaSettings"
            >
              <span v-if="loading" class="ml-2">
                <svg class="animate-spin h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                  <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                </svg>
              </span>
              ذخیره تنظیمات
            </button>
          </div>
        </div>
      </div>

      <!-- Password Change -->
      <div class="bg-white shadow rounded-lg p-6">
        <h3 class="text-lg font-medium text-gray-900 mb-4">تغییر رمز عبور</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">رمز عبور فعلی</label>
            <input
              v-model="currentPassword"
              type="password"
              class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">رمز عبور جدید</label>
            <input
              v-model="newPassword"
              type="password"
              class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">تکرار رمز عبور جدید</label>
            <input
              v-model="confirmPassword"
              type="password"
              class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            />
          </div>
          <div class="flex justify-end">
            <button
              :disabled="loading"
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 disabled:opacity-50"
              @click="handleChangePassword"
            >
              <span v-if="loading" class="ml-2">
                <svg class="animate-spin h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                  <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                </svg>
              </span>
              تغییر رمز عبور
            </button>
          </div>
        </div>
      </div>

      <!-- Rate Limiting Settings -->
      <div class="bg-white shadow rounded-lg p-6">
        <h3 class="text-lg font-medium text-gray-900 mb-4">تنظیمات محدودیت نرخ درخواست</h3>
        <div class="space-y-6">
          <!-- پنل ادمین -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h4 class="text-md font-medium text-gray-800 mb-3 flex items-center">
              <svg class="w-5 h-5 text-blue-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"></path>
              </svg>
              پنل ادمین
            </h4>
            <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">درخواست در دقیقه</label>
                <input
                  v-model="rateLimitSettings.admin.requestsPerMinute"
                  type="number"
                  min="10"
                  max="1000"
                  class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                />
                <p class="mt-1 text-xs text-gray-500">پیش‌فرض: 100</p>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">درخواست در ساعت</label>
                <input
                  v-model="rateLimitSettings.admin.requestsPerHour"
                  type="number"
                  min="100"
                  max="10000"
                  class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                />
                <p class="mt-1 text-xs text-gray-500">پیش‌فرض: 1000</p>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">درخواست در روز</label>
                <input
                  v-model="rateLimitSettings.admin.requestsPerDay"
                  type="number"
                  min="1000"
                  max="100000"
                  class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                />
                <p class="mt-1 text-xs text-gray-500">پیش‌فرض: 10000</p>
              </div>
            </div>
          </div>

          <!-- کاربران عمومی و مشتریان -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h4 class="text-md font-medium text-gray-800 mb-3 flex items-center">
              <svg class="w-5 h-5 text-green-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"></path>
              </svg>
              کاربران عمومی و مشتریان
            </h4>
            <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">درخواست در دقیقه</label>
                <input
                  v-model="rateLimitSettings.public.requestsPerMinute"
                  type="number"
                  min="5"
                  max="200"
                  class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                />
                <p class="mt-1 text-xs text-gray-500">پیش‌فرض: 30</p>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">درخواست در ساعت</label>
                <input
                  v-model="rateLimitSettings.public.requestsPerHour"
                  type="number"
                  min="50"
                  max="2000"
                  class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                />
                <p class="mt-1 text-xs text-gray-500">پیش‌فرض: 300</p>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">درخواست در روز</label>
                <input
                  v-model="rateLimitSettings.public.requestsPerDay"
                  type="number"
                  min="500"
                  max="20000"
                  class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                />
                <p class="mt-1 text-xs text-gray-500">پیش‌فرض: 5000</p>
              </div>
            </div>
          </div>

          <div class="flex justify-end">
            <button
              :disabled="loading"
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 disabled:opacity-50"
              @click="handleSaveRateLimitSettings"
            >
              <span v-if="loading" class="ml-2">
                <svg class="animate-spin h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                  <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                </svg>
              </span>
              ذخیره تنظیمات محدودیت
            </button>
          </div>
        </div>
      </div>

      <!-- Two-Factor Authentication -->
      <div class="bg-white shadow rounded-lg p-6">
        <h3 class="text-lg font-medium text-gray-900 mb-4">احراز هویت دو مرحله‌ای</h3>
        <div class="space-y-4">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm font-medium text-gray-900">فعال‌سازی احراز هویت دو مرحله‌ای</p>
              <p class="text-sm text-gray-500">با فعال‌سازی این گزینه، علاوه بر رمز عبور، کد تایید ارسال شده به تلفن همراه نیز نیاز خواهد بود.</p>
            </div>
            <button
              :disabled="loading"
              class="relative inline-flex flex-shrink-0 h-6 w-11 border-2 border-transparent rounded-full cursor-pointer transition-colors ease-in-out duration-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 disabled:opacity-50"
              :class="[twoFactorEnabled ? 'bg-indigo-600' : 'bg-gray-200']"
              @click="handleToggleTwoFactor"
            >
              <span
                class="pointer-events-none inline-block h-5 w-5 rounded-full bg-white shadow transform ring-0 transition ease-in-out duration-200"
                :class="[twoFactorEnabled ? 'translate-x-5' : 'translate-x-0']"
              />
            </button>
          </div>
        </div>
      </div>

      <!-- Active Sessions -->
      <div class="bg-white shadow rounded-lg p-6">
        <h3 class="text-lg font-medium text-gray-900 mb-4">نشست‌های فعال</h3>
        <div class="space-y-4">
          <div v-for="session in activeSessions" :key="session.id" class="flex items-center justify-between py-3 border-b border-gray-200">
            <div>
              <p class="text-sm font-medium text-gray-900">{{ session.device }}</p>
              <p class="text-sm text-gray-500">{{ session.location }} - {{ session.lastActive }}</p>
            </div>
            <button
              :disabled="loading"
              class="text-sm text-red-600 hover:text-red-900 disabled:opacity-50"
              @click="handleTerminateSession(session.id)"
            >
              خاتمه نشست
            </button>
          </div>
        </div>
      </div>

      <!-- Login History -->
      <div class="bg-white shadow rounded-lg p-6">
        <h3 class="text-lg font-medium text-gray-900 mb-4">تاریخچه ورود</h3>
        <div class="space-y-4">
          <div v-for="login in loginHistory" :key="String((login as { id?: string | number }).id ?? Math.random())" class="py-3 border-b border-gray-200">
            <p class="text-sm font-medium text-gray-900">{{ login.device }}</p>
            <p class="text-sm text-gray-500">{{ login.location }} - {{ login.timestamp }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
</script>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { useSecuritySettings } from '~/composables/useSecuritySettings';

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
// const { user, hasPermission } = useAuth()

const {
  loading,
  // error,
  recaptchaSettings,
  saveRecaptchaSettings,
  changePassword,
  twoFactorEnabled,
  toggleTwoFactor,
  activeSessions,
  terminateSession,
  loginHistory,
  loadActiveSessions,
  loadLoginHistory,
  fetchSettings,
  saveRateLimitSettings
} = useSecuritySettings()

// Rate Limiting Settings
const rateLimitSettings = ref({
  admin: {
    requestsPerMinute: 100,
    requestsPerHour: 1000,
    requestsPerDay: 10000
  },
  public: {
    requestsPerMinute: 30,
    requestsPerHour: 300,
    requestsPerDay: 5000
  }
})

// Password Change
const currentPassword = ref('')
const newPassword = ref('')
const confirmPassword = ref('')

const handleSaveRecaptchaSettings = async () => {
  try {
    await saveRecaptchaSettings(recaptchaSettings.value)
  } catch (err) {
    console.error('Error saving reCAPTCHA settings:', err)
  }
}

const handleSaveRateLimitSettings = async () => {
  try {
    const result = await saveRateLimitSettings(rateLimitSettings.value)
    
    if (result.success) {
      alert(result.message)
    } else {
      alert(result.message)
    }
  } catch (err) {
    console.error('Error saving rate limit settings:', err)
    alert('خطا در ذخیره تنظیمات محدودیت نرخ درخواست')
  }
}

const handleChangePassword = async () => {
  if (newPassword.value !== confirmPassword.value) {
    alert('رمز عبور جدید و تکرار آن مطابقت ندارند')
    return
  }

  try {
    await changePassword(currentPassword.value, newPassword.value)
    currentPassword.value = ''
    newPassword.value = ''
    confirmPassword.value = ''
  } catch (err) {
    console.error('Error changing password:', err)
  }
}

const handleToggleTwoFactor = async () => {
  try {
    await toggleTwoFactor()
  } catch (err) {
    console.error('Error toggling two-factor authentication:', err)
  }
}

const handleTerminateSession = async (sessionId: number) => {
  try {
    await terminateSession(sessionId)
  } catch (err) {
    console.error('Error terminating session:', err)
  }
}

onMounted(async () => {
  try {
    // بارگذاری تنظیمات امنیت از API
    await fetchSettings()
    
    await Promise.all([
      loadActiveSessions(),
      loadLoginHistory()
    ])
    
  } catch (err) {
    console.error('Error loading data:', err)
  }
})
</script> 

